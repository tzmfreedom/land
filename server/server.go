package server

import (
	"bytes"
	"fmt"
	"net/http"
	"os"

	"encoding/json"

	"regexp"

	"encoding/base64"

	"github.com/tzmfreedom/goland/ast"
	"github.com/tzmfreedom/goland/builtin"
	"github.com/tzmfreedom/goland/compiler"
	"github.com/tzmfreedom/goland/interpreter"
)

type Server struct {
	ClassTypes []*builtin.ClassType
}

func (s *Server) Run() {
	classMap := builtin.NewClassMapWithPrimivie(s.ClassTypes)
	interpreter := interpreter.NewInterpreter(classMap)

	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		reg := regexp.MustCompile("^/([^/]+)/([^/]+)")
		match := reg.FindStringSubmatch(r.RequestURI)
		klass := match[1]
		method := match[2]
		params := []interface{}{}
		buf := new(bytes.Buffer)
		buf.ReadFrom(r.Body)
		json.Unmarshal(buf.Bytes(), &params)

		parameters := make([]ast.Node, len(params))
		for i, param := range params {
			var p ast.Node
			switch n := param.(type) {
			case string:
				p = &ast.StringLiteral{Value: n}
			case float64:
				p = &ast.DoubleLiteral{Value: n}
			case int:
				p = &ast.IntegerLiteral{Value: n}
			case bool:
				p = &ast.BooleanLiteral{Value: n}
			}
			parameters[i] = p
		}
		name := []string{klass, method}
		invoke := &ast.MethodInvocation{
			NameOrExpression: &ast.Name{
				Value: name,
			},
			Parameters: parameters,
		}
		//resolver := &compiler.TypeResolver{
		//	Context: &compiler.Context{ClassTypes: classMap},
		//}
		//if _, err := resolver.ResolveMethod(name, parameters); err != nil {
		//	panic(err)
		//}
		res, err := invoke.Accept(interpreter)
		if err != nil {
			panic(err)
		}
		fmt.Fprintf(w, builtin.String(res.(*builtin.Object)))
		fmt.Fprintln(w)
	})
	port := getServerPort()
	fmt.Println("listening to 0.0.0.0:" + port)
	err := http.ListenAndServe(":"+port, nil)
	if err != nil {
		panic(err)
	}
}

type EvalServer struct{}

type EvalRequest struct {
	String    string
	Method    string
	WithClass bool
}

type EvalResult struct {
	String string
	Result bool
	Error  string
}

type FormatRequest struct {
	String string
}

type FormatResponse struct {
	String string
}

func (s *EvalServer) Run() {
	http.HandleFunc("/eval", func(w http.ResponseWriter, r *http.Request) {
		eval(w, r)
	})
	http.HandleFunc("/format", func(w http.ResponseWriter, r *http.Request) {
		// parse request
		req := &FormatRequest{}
		buf := new(bytes.Buffer)
		buf.ReadFrom(r.Body)
		json.Unmarshal(buf.Bytes(), &req)
		b, err := base64.StdEncoding.DecodeString(req.String)
		if err != nil {
			fmt.Fprintf(os.Stderr, "Error")
		}

		// parse source
		root, err := ast.ParseString(string(b))
		if err != nil {
			fmt.Fprintf(os.Stderr, "Error")
		}
		// format
		visitor := &ast.TosVisitor{}
		str, _ := root.Accept(visitor)
		if err != nil {
			fmt.Fprintf(os.Stderr, "Error")
		}

		// response
		b64body := base64.StdEncoding.EncodeToString([]byte(str.(string)))
		body, err := json.Marshal(&FormatResponse{
			String: b64body,
		})
		if err != nil {
			fmt.Fprintf(os.Stderr, "Error")
		}
		fmt.Fprintf(w, string(body))
	})
	http.Handle("/", http.FileServer(http.Dir("eval-server")))

	port := getServerPort()
	fmt.Println("listening to 0.0.0.0:" + port)
	err := http.ListenAndServe(":"+port, nil)
	if err != nil {
		panic(err)
	}
}

func eval(w http.ResponseWriter, r *http.Request) {
	classMap := builtin.NewClassMapWithPrimivie([]*builtin.ClassType{})

	req := &EvalRequest{}
	buf := new(bytes.Buffer)
	buf.ReadFrom(r.Body)
	json.Unmarshal(buf.Bytes(), &req)
	b, err := base64.StdEncoding.DecodeString(req.String)
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error")
	}

	classBody := string(b)
	if !req.WithClass {
		classBody = fmt.Sprintf(`public class Land { public static void action() { %s; } }`, classBody)
	}
	// compile
	root, err := ast.ParseString(classBody)
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error")
	}
	register := &compiler.ClassRegisterVisitor{}
	t, err := root.Accept(register)
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error")
	}
	classType := t.(*builtin.ClassType)
	classMap.Set(classType.Name, classType)
	typeChecker := compiler.NewTypeChecker()
	typeChecker.Context.ClassTypes = classMap
	_, err = typeChecker.VisitClassType(classType)
	if err != nil {
		fmt.Fprintf(os.Stderr, err.Error())
		body, err := json.Marshal(&EvalResult{
			Error:  err.Error(),
			Result: false,
		})
		if err != nil {
			fmt.Fprintf(os.Stderr, "Error")
		}
		w.WriteHeader(400)
		fmt.Fprint(w, string(body))
		return
	}
	if len(typeChecker.Errors) != 0 {
		for _, e := range typeChecker.Errors {
			fmt.Fprintf(os.Stderr, "%s\n", e.Message)
		}
	}
	// interpreter
	invoke := &ast.MethodInvocation{
		NameOrExpression: &ast.Name{
			Value: []string{classType.Name, req.Method},
		},
	}

	interpreter := interpreter.NewInterpreter(classMap)
	interpreter.LoadStaticField()
	stdout := new(bytes.Buffer)
	interpreter.Stdout = stdout
	interpreter.Stderr = new(bytes.Buffer)
	_, err = invoke.Accept(interpreter)
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error")
	}

	b64body := base64.StdEncoding.EncodeToString(stdout.Bytes())
	body, err := json.Marshal(&EvalResult{
		String: b64body,
		Result: true,
	})
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error")
	}
	fmt.Fprint(w, string(body))
}

func getServerPort() string {
	port := "8080"
	if os.Getenv("PORT") != "" {
		port = os.Getenv("PORT")
	}
	return port
}

func Run(classTypes []*builtin.ClassType) {
	server := &Server{ClassTypes: classTypes}
	server.Run()
}
