package server

import (
	"fmt"
	"net/http"

	"github.com/tzmfreedom/goland/ast"
	"github.com/tzmfreedom/goland/builtin"
	"github.com/tzmfreedom/goland/interpreter"
)

type Server struct {
	ClassTypes []*builtin.ClassType
}

func (s *Server) Run() {
	//method := "action"
	//args := strings.Split(action, "#")
	//if len(args) > 1 {
	//	method = args[1]
	//}
	interpreter := interpreter.NewInterpreter(builtin.PrimitiveClassMap())
	for _, classType := range s.ClassTypes {
		interpreter.Context.ClassTypes.Set(classType.Name, classType)
	}
	invoke := &ast.MethodInvocation{
		NameOrExpression: &ast.Name{
			Value: []string{"Foo", "rest"},
		},
	}

	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
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
