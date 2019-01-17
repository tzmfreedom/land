package main

import (
	"errors"
	"fmt"
	"os"
	"time"

	"flag"
	"strings"

	"io/ioutil"

	"regexp"

	"bytes"

	"github.com/chzyer/readline"
	"github.com/fsnotify/fsnotify"
	"github.com/joho/godotenv"
	"github.com/mattn/go-colorable"
	"github.com/tzmfreedom/goland/ast"
	"github.com/tzmfreedom/goland/builtin"
	"github.com/tzmfreedom/goland/compiler"
	"github.com/tzmfreedom/goland/interpreter"
	"github.com/tzmfreedom/goland/server"
	"github.com/tzmfreedom/goland/visitor"
)

var classMap = ast.NewClassMap()
var preprocessors = []ast.PreProcessor{
	func(src string) string {
		src = strings.Replace(src, "// #debugger", "_Debugger.run();", -1)
		r := regexp.MustCompile(`// #debug\((.+)\)`)
		src = r.ReplaceAllString(src, "_Debugger.debug($1);")
		return src
	},
}

type option struct {
	SubCommand   string
	Action       string
	Files        []string
	Interactive  bool
	LoaderSource string
}

func main() {
	err := godotenv.Load()
	option, err := parseOption(os.Args)
	if err != nil {
		handleError(err)
		return
	}

	// TODO: namespace(not builtin?)
	if option.SubCommand != "setup" {
		builtin.Load(option.LoaderSource)
	}

	switch option.SubCommand {
	case "db:setup":
		builtin.Setup()
		return
	case "db:seed":
		builtin.Seed()
		return
	case "setup":
		err := builtin.CreateMetadataFile(option.LoaderSource)
		if err != nil {
			handleError(err)
			return
		}
	case "test":
		trees, err := parseFiles(option.Files)
		if err != nil {
			handleError(err)
		}
		classTypes, err := buildAllFile(trees)
		if err != nil {
			handleError(err)
		}
		var i = 1
		for _, classType := range classTypes {
			for _, methods := range classType.StaticMethods.All() {
				for _, m := range methods {
					if m.IsTestMethod() {
						action := fmt.Sprintf("%s#%s", classType.Name, m.Name)
						fmt.Printf("(%d) %s: ", i, action)
						var ret *interpreter.Interpreter
						err = run(action, classTypes, func(i *interpreter.Interpreter) {
							ret = i
							i.Extra["stdout"] = new(bytes.Buffer)
						})
						if err != nil {
							handleError(err)
						}
						stdout := colorable.NewColorableStdout()
						errors := ret.Extra["errors"].([]*builtin.TestError)
						if len(errors) > 0 {
							fmt.Println("")
							for _, error := range errors {
								loc := error.Node.GetLocation()
								str := fmt.Sprintf("  %s at %d:%d\n", loc.FileName, loc.Line, loc.Column)
								fmt.Fprintf(stdout, builtin.NoticeColor, str)
								str = fmt.Sprintf(`    Failure/Error: %s

%s
`, ast.ToString(error.Node), error.Message)
								fmt.Fprintf(stdout, builtin.ErrorColor, str)
							}
						} else {
							fmt.Fprintf(stdout, builtin.InfoColor, "pass\n")
						}
						fmt.Println("")
						i++
					}
				}
			}
		}
	case "watch":
		trees, err := parseFiles(option.Files)
		if err != nil {
			handleError(err)
		}
		classTypes, err := buildAllFile(trees)
		if err != nil {
			handleError(err)
		}
		watchAndRunTest(classTypes, option)
	case "server":
		trees, err := parseFiles(option.Files)
		if err != nil {
			handleError(err)
		}
		classTypes, err := buildAllFile(trees)
		if err != nil {
			handleError(err)
		}
		server.Run(classTypes)
	case "eval-server":
		s := &server.EvalServer{}
		s.Run()
	case "format":
		trees, err := parseFiles(option.Files)
		if err != nil {
			handleError(err)
		}
		for _, t := range trees {
			tos(t)
		}
	case "run":
		trees, err := parseFiles(option.Files)
		if err != nil {
			handleError(err)
		}
		classTypes, err := buildAllFile(trees)
		if err != nil {
			handleError(err)
		}
		if option.Interactive {
			err = interactiveRun(classTypes, option)
			if err != nil {
				handleError(err)
			}
		} else {
			err = run(option.Action, classTypes)
			if err != nil {
				handleError(err)
			}
		}
	case "check":
		trees, err := parseFiles(option.Files)
		if err != nil {
			handleError(err)
		}
		classTypes, err := buildAllFile(trees)
		if err != nil {
			handleError(err)
		}
		for _, classType := range classTypes {
			check(classType)
		}
	}
}

func parseFiles(files []string) ([]ast.Node, error) {
	trees := make([]ast.Node, len(files))
	var err error
	for i, file := range files {
		trees[i], err = ast.ParseFile(file, preprocessors...)
		if err != nil {
			return nil, err
		}
	}
	return trees, nil
}

func parseOption(args []string) (*option, error) {
	flg := flag.NewFlagSet(args[0], flag.ExitOnError)
	fileName := flg.String("f", "", "file")
	directory := flg.String("d", "", "directory")
	action := flg.String("a", "", "action")
	interactive := flg.Bool("i", false, "interactive")
	inputLoaderSource := flg.String("l", "", "loader")

	loaderSource := *inputLoaderSource
	if loaderSource == "" {
		if env := os.Getenv("SALESFORCE_METADATA"); env != "" {
			loaderSource = env
		} else {
			loaderSource = builtin.DefaultMetafileName
		}
	}

	err := flg.Parse(args[2:])
	if err != nil {
		return nil, err
	}

	if fileName == nil && directory == nil {
		return nil, errors.New("-f FILE or -d DIRECTORY is required")
	}

	cmd := os.Args[1]
	if cmd != "format" &&
		cmd != "eval-server" &&
		cmd != "server" &&
		cmd != "setup" &&
		cmd != "db:setup" &&
		cmd != "db:seed" &&
		cmd != "test" &&
		*action == "" &&
		*interactive == false {
		return nil, errors.New("-a CLASS#METHOD is required")
	}
	var files []string
	if *fileName != "" {
		files = []string{*fileName}
	} else if *directory != "" {
		filesInDirectory, err := ioutil.ReadDir(*directory)
		if err != nil {
			handleError(err)
		}
		files = []string{}
		for _, f := range filesInDirectory {
			if f.IsDir() {
				continue
			}
			files = append(files, fmt.Sprintf("%s/%s", *directory, f.Name()))
		}
	}
	return &option{
		SubCommand:   cmd,
		Action:       *action,
		Files:        files,
		Interactive:  *interactive,
		LoaderSource: loaderSource,
	}, nil
}

func convert(n *ast.ClassType, classMap *ast.ClassMap) (*ast.ClassType, error) {
	resolver := compiler.NewTypeRefResolver(classMap, builtin.GetNameSpaceStore())
	return resolver.Resolve(n)
}

func check(n *ast.ClassType) error {
	checker := &visitor.SoqlChecker{}
	_, err := checker.VisitClassType(n)
	return err
}

func register(n ast.Node) (*ast.ClassType, error) {
	register := &compiler.ClassRegisterVisitor{}
	t, err := n.Accept(register)
	if err != nil {
		return nil, err
	}
	classType := t.(*ast.ClassType)
	if _, ok := classMap.Get(classType.Name); ok {
		return nil, fmt.Errorf("Class %s is already defined", classType.Name)
	}
	classMap.Set(classType.Name, classType)
	return classType, nil
}

func semanticAnalysis(t *ast.ClassType) error {
	typeChecker := compiler.NewTypeChecker()
	typeChecker.Context.ClassTypes = builtin.PrimitiveClassMap()
	for _, class := range classMap.Data {
		typeChecker.Context.ClassTypes.Set(class.Name, class)
	}
	typeChecker.Context.NameSpaces = builtin.GetNameSpaceStore()
	_, err := typeChecker.VisitClassType(t)
	if len(typeChecker.Errors) != 0 {
		for _, e := range typeChecker.Errors {
			//pp.Println(e)
			loc := e.Node.GetLocation()
			fmt.Fprintf(os.Stderr, "%s at %d:%d in %s\n", e.Message, loc.Line, loc.Column, loc.FileName)
		}
		return errors.New("compile error")
	}
	return err
}

func run(action string, classTypes []*ast.ClassType, options ...func(*interpreter.Interpreter)) error {
	method := "action"
	args := strings.Split(action, "#")
	if len(args) > 1 {
		method = args[1]
	}
	interpreter := interpreter.NewInterpreterWithBuiltin(classTypes)
	invoke := &ast.MethodInvocation{
		NameOrExpression: &ast.Name{
			Value: []string{args[0], method},
		},
	}
	for _, option := range options {
		option(interpreter)
	}
	interpreter.LoadStaticField()
	_, err := invoke.Accept(interpreter)
	return err
}

func interactiveRun(classTypes []*ast.ClassType, option *option) error {
	lastReloadedAt := time.Now()
	landInterpreter := interpreter.NewInterpreterWithBuiltin(classTypes)
	l, _ := readline.NewEx(&readline.Config{
		Prompt:          "\033[31m>>\033[0m ",
		HistoryFile:     "/tmp/land.tmp",
		InterruptPrompt: "^C",
		EOFPrompt:       "exit",
	})
	watcher, err := fsnotify.NewWatcher()
	if err != nil {
		return err
	}
	defer watcher.Close()

	ch := make(chan bool, 1)
	go func() {
		for {
			select {
			case event, ok := <-watcher.Events:
				if !ok {
					return
				}
				ch <- true
				if event.Op&fsnotify.Write == fsnotify.Write {
					buildFile(landInterpreter, event.Name)
				} else if event.Op&fsnotify.Create == fsnotify.Create {
					buildFile(landInterpreter, event.Name)
				}
				<-ch
			case err, ok := <-watcher.Errors:
				if !ok {
					return
				}
				fmt.Println("error:", err)
			}
		}
	}()
	err = watcher.Add("./fixtures")
	if err != nil {
		return err
	}

	// TODO: implement
	env := interpreter.NewEnv(nil)
	interpreter.Subscribe("method_end", func(ctx *interpreter.Context, n ast.Node) {
		env = ctx.Env
	})
	for {
		line, err := l.Readline()
		if err != nil {
			panic(err)
		}
		inputs := strings.Split(line, " ")
		cmd := inputs[0]
		args := inputs[1:]
		switch cmd {
		case "execute":
			if len(args) == 0 {
				fmt.Println("Error: execute command required argument")
				continue
			}
			ch <- true
			run(args[0], classTypes)
			<-ch
		case "reload":
			ch <- true
			if len(args) == 0 {
				reloadAll(landInterpreter, option.Files)
			} else {
				_, err := buildFile(landInterpreter, args[0])
				if err != nil {
					fmt.Println(err.Error())
				}
			}
			<-ch
		case "run":
			if len(args) == 0 {
				fmt.Println("Error: run command required argument")
				continue
			}
			isReload := false
			for _, f := range option.Files {
				info, err := os.Stat(f)
				if err != nil {
					return err
				}
				if info.ModTime().After(lastReloadedAt) {
					isReload = true
					break
				}
			}
			ch <- true
			if isReload {
				lastReloadedAt = time.Now()
				reloadAll(landInterpreter, option.Files)
			}
			run(args[0], classTypes)
			<-ch
		case "exit":
			return nil
		default:
			code := createTempClass(line)
			execFile(code, env)
		}
	}
	return nil
}

func createTempClass(statement string) string {
	return fmt.Sprintf(`public class Temporary {
public static void action() { %s; }
}`, statement)
}

func watchAndRunTest(classTypes []*ast.ClassType, option *option) error {
	interpreter := interpreter.NewInterpreterWithBuiltin(classTypes)

	watcher, err := fsnotify.NewWatcher()
	if err != nil {
		return err
	}
	defer watcher.Close()

	err = watcher.Add("./classes")
	for {
		select {
		case event, ok := <-watcher.Events:
			if !ok {
				return err
			}
			if event.Op&fsnotify.Write == fsnotify.Write ||
				event.Op&fsnotify.Create == fsnotify.Create {
				classType, err := buildFile(interpreter, event.Name)
				if err != nil {
					fmt.Printf("Error: %s\n", err.Error())
				} else {
					for _, methods := range classType.StaticMethods.Data {
						for _, m := range methods {
							fmt.Println(m.IsTestMethod())
							if m.IsTestMethod() {
								runAction(interpreter, []string{classType.Name, m.Name})
							}
						}
					}
				}
			}
		case err, ok := <-watcher.Errors:
			if !ok {
				return err
			}
			fmt.Println("error:", err)
		}
	}

	return err
}

func runAction(interpreter *interpreter.Interpreter, expression []string) error {
	invoke := &ast.MethodInvocation{
		NameOrExpression: &ast.Name{
			Value: expression,
		},
	}
	interpreter.LoadStaticField()
	_, err := invoke.Accept(interpreter)
	return err
}

func buildFile(interpreter *interpreter.Interpreter, file string) (*ast.ClassType, error) {
	t, err := ast.ParseFile(file, preprocessors...)
	classType, err := register(t)
	if err != nil {
		return nil, fmt.Errorf("Build Error: %s\n", err.Error())
	}
	tmpClassMap := builtin.PrimitiveClassMap()
	for _, classType := range classMap.Data {
		tmpClassMap.Set(classType.Name, classType)
	}

	classType, err = convert(classType, tmpClassMap)
	if err != nil {
		return nil, fmt.Errorf("Build Error: %s\n", err.Error())
	}
	if err = semanticAnalysis(classType); err != nil {
		return nil, fmt.Errorf("Build Error: %s\n", err.Error())
	}
	interpreter.Context.ClassTypes.Set(classType.Name, classType)
	return classType, nil
}

func buildAllFile(trees []ast.Node) ([]*ast.ClassType, error) {
	classTypes := make([]*ast.ClassType, len(trees))
	var err error
	for i, t := range trees {
		classTypes[i], err = register(t)
		if err != nil {
			return nil, err
		}
	}
	tmpClassMap := builtin.PrimitiveClassMap()
	for _, classType := range classMap.Data {
		tmpClassMap.Set(classType.Name, classType)
	}

	for i, classType := range classTypes {
		classTypes[i], err = convert(classType, tmpClassMap)
		if err != nil {
			return nil, err
		}
	}
	for _, t := range classTypes {
		if err := semanticAnalysis(t); err != nil {
			return nil, err
		}
	}
	return classTypes, nil
}

func execFile(code string, env *interpreter.Env) *interpreter.Env {
	t, err := ast.ParseString(code, preprocessors...)
	classType, err := register(t)
	if err != nil {
		panic(err)
	}
	classTypes := builtin.PrimitiveClassMap()
	for _, classType := range classMap.Data {
		classTypes.Set(classType.Name, classType)
	}

	classType, err = convert(classType, classTypes)
	if err != nil {
		panic(err)
	}
	if err = semanticAnalysis(classType); err != nil {
		panic(err)
	}
	interpreter := interpreter.NewInterpreterWithBuiltin([]*ast.ClassType{classType})
	interpreter.Context.Env = env
	invoke := &ast.MethodInvocation{
		NameOrExpression: &ast.Name{
			Value: []string{"Temporary", "action"},
		},
	}
	interpreter.LoadStaticField()
	_, err = invoke.Accept(interpreter)
	if err != nil {
		panic(err)
	}
	return env
}

func reloadAll(interpreter *interpreter.Interpreter, files []string) {
	var err error
	trees := make([]ast.Node, len(files))
	for i, file := range files {
		trees[i], err = ast.ParseFile(file, preprocessors...)
		if err != nil {
			handleError(err)
		}
	}
	classTypes := make([]*ast.ClassType, len(trees))
	for i, t := range trees {
		classTypes[i], err = register(t)
		if err != nil {
			handleError(err)
		}
	}

	classMap := builtin.PrimitiveClassMap()
	for _, classType := range classTypes {
		classMap.Set(classType.Name, classType)
	}

	for i, classType := range classTypes {
		classTypes[i], err = convert(classType, classMap)
		if err != nil {
			handleError(err)
		}
	}

	for _, t := range classTypes {
		if err = semanticAnalysis(t); err != nil {
			handleError(err)
		}
	}
	interpreter.Context.ClassTypes.Clear()
	for _, classType := range classTypes {
		interpreter.Context.ClassTypes.Set(classType.Name, classType)
	}
}

func tos(n ast.Node) {
	visitor := &ast.TosVisitor{}
	r, _ := n.Accept(visitor)
	fmt.Println(r)
}

func handleError(err error) {
	fmt.Fprintln(os.Stderr, err.Error())
	os.Exit(1)
}
