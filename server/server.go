package server

import (
	"bytes"
	"fmt"
	"net/http"

	"encoding/json"

	"regexp"

	"github.com/tzmfreedom/goland/ast"
	"github.com/tzmfreedom/goland/builtin"
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
	err := http.ListenAndServe(":8080", nil)
	if err != nil {
		panic(err)
	}
}

func Run(classTypes []*builtin.ClassType) {
	server := &Server{ClassTypes: classTypes}
	server.Run()
}
