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

type EvalRequest struct {
	String string
	Method string
}

type EvalResult struct {
	String string
}

func (s *Server) Run() {
	classMap := builtin.NewClassMapWithPrimivie(s.ClassTypes)
	interpreter := interpreter.NewInterpreter(classMap)

	http.HandleFunc("/eval", func(w http.ResponseWriter, r *http.Request) {
		req := &EvalRequest{}
		buf := new(bytes.Buffer)
		buf.ReadFrom(r.Body)
		json.Unmarshal(buf.Bytes(), &req)
		b, err := base64.StdEncoding.DecodeString(req.String)
		if err != nil {
			fmt.Fprintf(os.Stderr, "Error")
		}
		root, err := ast.ParseString(string(b))
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
			fmt.Fprintf(os.Stderr, "Error")
		}
		if len(typeChecker.Errors) != 0 {
			for _, e := range typeChecker.Errors {
				fmt.Fprintf(os.Stderr, "%s\n", e.Message)
			}
		}
		invoke := &ast.MethodInvocation{
			NameOrExpression: &ast.Name{
				Value: []string{classType.Name, req.Method},
			},
		}
		interpreter.LoadStaticField()
		stdout := new(bytes.Buffer)
		interpreter.Stdout = stdout
		interpreter.Stderr = new(bytes.Buffer)
		_, err = invoke.Accept(interpreter)
		if err != nil {
			fmt.Fprintf(os.Stderr, "Error")
		}

		b64body := base64.StdEncoding.EncodeToString(stdout.Bytes())
		body, err := json.Marshal(&EvalResult{b64body})
		if err != nil {
			fmt.Fprintf(os.Stderr, "Error")
		}
		fmt.Fprint(w, string(body))
	})
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
	err := http.ListenAndServe(":8080", nil)
	if err != nil {
		panic(err)
	}
}

func Run(classTypes []*builtin.ClassType) {
	server := &Server{ClassTypes: classTypes}
	server.Run()
}
