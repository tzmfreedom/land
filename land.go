package main

import (
	"fmt"
	"os"

	"flag"

	"github.com/tzmfreedom/goland/ast"
	"github.com/tzmfreedom/goland/compiler"
	"github.com/tzmfreedom/goland/visitor"
)

func main() {
	f := flag.String("f", "", "file")
	_ = flag.String("d", "", "directory")
	cmd := os.Args[1]
	os.Args = os.Args[1:]

	flag.Parse()

	t, err := ast.ParseFile(*f)
	if err != nil {
		handleError(err)
	}
	switch cmd {
	case "format":
		tos(t)
	case "run":
		root, err := convert(t)
		if err != nil {
			handleError(err)
		}
		err = run(root)
		if err != nil {
			handleError(err)
		}
	case "check":
		root, err := convert(t)
		if err != nil {
			handleError(err)
		}
		err = check(root)
		if err != nil {
			handleError(err)
		}
	}
}

func convert(n ast.Node) (ast.Node, error) {
	return n, nil
}

func check(n ast.Node) error {
	checker := &visitor.SoqlChecker{}
	_, err := n.Accept(checker)
	return err
}

func semantic_analysis(n ast.Node) error {
	register := &compiler.ClassRegisterVisitor{}
	t, err := n.Accept(register)
	if err != nil {
		return err
	}
	classTypes := make([]compiler.ClassType, 1)
	if tp, ok := t.(compiler.ClassType); ok {
		classTypes[1] = tp
	}
	typeChecker := &compiler.TypeChecker{
		ClassTypes: classTypes,
	}
	_, err = n.Accept(typeChecker)
	if err != nil {
		return err
	}
	return nil
}

func run(n ast.Node) error {
	interpreter := &compiler.Interpreter{}
	_, err := n.Accept(interpreter)
	return err
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
