package main

import (
	"errors"
	"fmt"
	"os"
	"time"

	"flag"
	"strings"

	"io/ioutil"

	"github.com/chzyer/readline"
	"github.com/fsnotify/fsnotify"
	"github.com/tzmfreedom/goland/ast"
	"github.com/tzmfreedom/goland/builtin"
	"github.com/tzmfreedom/goland/compiler"
	"github.com/tzmfreedom/goland/interpreter"
	"github.com/tzmfreedom/goland/server"
	"github.com/tzmfreedom/goland/visitor"
)

var classMap = builtin.NewClassMap()
var preprocessors = []ast.PreProcessor{
	func(src string) string {
		return strings.Replace(src, "// #debugger", "Debugger.debug();", -1)
	},
}

type option struct {
	SubCommand  string
	Action      string
	Files       []string
	Interactive bool
}

func main() {
	option, err := parseOption(os.Args)
	if err != nil {
		handleError(err)
		return
	}

	trees := make([]ast.Node, len(option.Files))
	for i, file := range option.Files {
		trees[i], err = ast.ParseFile(file, preprocessors...)
		if err != nil {
			handleError(err)
		}
	}
	switch option.SubCommand {
	case "watch":
		classTypes, err := buildAllFile(trees)
		if err != nil {
			handleError(err)
		}
		watchAndRunTest(classTypes, option)
	case "server":
		classTypes, err := buildAllFile(trees)
		if err != nil {
			handleError(err)
		}
		server.Run(classTypes)
	case "format":
		for _, t := range trees {
			tos(t)
		}
	case "run":
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
		newTrees := make([]*builtin.ClassType, len(trees))
		for i, t := range trees {
			root, err := convert(t)
			if err != nil {
				handleError(err)
			}
			newTrees[i], err = register(root)
		}
		for _, t := range newTrees {
			err = semanticAnalysis(t)
			if err != nil {
				handleError(err)
			}
		}
	}
}

func parseOption(args []string) (*option, error) {
	flg := flag.NewFlagSet(args[0], flag.ExitOnError)
	fileName := flg.String("f", "", "file")
	directory := flg.String("d", "", "directory")
	action := flg.String("a", "", "action")
	interactive := flg.Bool("i", false, "interactive")

	err := flg.Parse(args[2:])
	if err != nil {
		return nil, err
	}

	if fileName == nil && directory == nil {
		return nil, errors.New("-f FILE or -d DIRECTORY is required")
	}

	cmd := os.Args[1]
	if cmd != "format" && *action == "" {
		return nil, errors.New("-a CLASS#METHOD is required")
	}
	var files []string
	if *fileName != "" {
		files = []string{*fileName}
	} else {
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
		SubCommand:  cmd,
		Action:      *action,
		Files:       files,
		Interactive: *interactive,
	}, nil
}

func convert(n ast.Node) (ast.Node, error) {
	return n, nil
}

func check(n ast.Node) error {
	checker := &visitor.SoqlChecker{}
	_, err := n.Accept(checker)
	return err
}

func register(n ast.Node) (*builtin.ClassType, error) {
	register := &compiler.ClassRegisterVisitor{}
	t, err := n.Accept(register)
	if err != nil {
		return nil, err
	}
	classType := t.(*builtin.ClassType)
	classMap.Set(classType.Name, classType)
	return classType, nil
}

func semanticAnalysis(t *builtin.ClassType) error {
	typeChecker := compiler.NewTypeChecker()
	typeChecker.Context.ClassTypes = builtin.PrimitiveClassMap()
	for _, class := range classMap.Data {
		typeChecker.Context.ClassTypes.Set(class.Name, class)
	}
	_, err := typeChecker.VisitClassType(t)
	if len(typeChecker.Errors) != 0 {
		for _, e := range typeChecker.Errors {
			fmt.Fprintf(os.Stderr, "%s\n", e.Message)
		}
	}
	return err
}

func run(action string, classTypes []*builtin.ClassType) error {
	method := "action"
	args := strings.Split(action, "#")
	if len(args) > 1 {
		method = args[1]
	}
	interpreter := interpreter.NewInterpreter(builtin.PrimitiveClassMap())
	for _, classType := range classTypes {
		interpreter.Context.ClassTypes.Set(classType.Name, classType)
	}
	invoke := &ast.MethodInvocation{
		NameOrExpression: &ast.Name{
			Value: []string{args[0], method},
		},
	}
	interpreter.LoadStaticField()
	_, err := invoke.Accept(interpreter)
	return err
}

func interactiveRun(classTypes []*builtin.ClassType, option *option) error {
	lastReloadedAt := time.Now()
	interpreter := interpreter.NewInterpreter(builtin.PrimitiveClassMap())
	for _, classType := range classTypes {
		interpreter.Context.ClassTypes.Set(classType.Name, classType)
	}

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
					buildFile(interpreter, event.Name)
				} else if event.Op&fsnotify.Create == fsnotify.Create {
					buildFile(interpreter, event.Name)
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
				reloadAll(interpreter, option.Files)
			} else {
				_, err := buildFile(interpreter, args[0])
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
				reloadAll(interpreter, option.Files)
			}
			run(args[0], classTypes)
			<-ch
		case "exit":
			return nil
		}
	}
	return nil
}

func watchAndRunTest(classTypes []*builtin.ClassType, option *option) error {
	interpreter := interpreter.NewInterpreter(builtin.PrimitiveClassMap())
	for _, classType := range classTypes {
		interpreter.Context.ClassTypes.Set(classType.Name, classType)
	}

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
							decl := m.(*ast.MethodDeclaration)
							fmt.Println(decl.IsTestMethod())
							if decl.IsTestMethod() {
								runAction(interpreter, []string{classType.Name, decl.Name})
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

func buildFile(interpreter *interpreter.Interpreter, file string) (*builtin.ClassType, error) {
	t, err := ast.ParseFile(file, preprocessors...)
	root, err := convert(t)
	if err != nil {
		return nil, fmt.Errorf("Build Error: %s\n", err.Error())
	}
	classType, err := register(root)
	if err = semanticAnalysis(classType); err != nil {
		return nil, fmt.Errorf("Build Error: %s\n", err.Error())
	}
	interpreter.Context.ClassTypes.Set(classType.Name, classType)
	return classType, nil
}

func buildAllFile(trees []ast.Node) ([]*builtin.ClassType, error) {
	classTypes := make([]*builtin.ClassType, len(trees))
	for i, t := range trees {
		root, err := convert(t)
		if err != nil {
			return nil, err
		}
		classTypes[i], err = register(root)
	}
	for _, t := range classTypes {
		if err := semanticAnalysis(t); err != nil {
			return nil, err
		}
	}
	return classTypes, nil
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
	classTypes := make([]*builtin.ClassType, len(trees))
	for i, t := range trees {
		root, err := convert(t)
		if err != nil {
			handleError(err)
		}
		classTypes[i], err = register(root)
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
	fmt.Fprintf(os.Stderr, err.Error())
	os.Exit(1)
}

func validate() {
	return
}
