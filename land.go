package main

import (
	"os"

	"github.com/antlr/antlr4/runtime/Go/antlr"
	"github.com/tzmfreedom/goland/ast"
	"github.com/tzmfreedom/goland/parser"
	"github.com/tzmfreedom/goland/visitor"
)

func main() {
	f := os.Args[1]
	t := parseFile(f)
	root, err := convert(t)
	if err != nil {
		panic(err)
	}
	err = check(root)
	if err != nil {
		panic(err)
	}
	err = run(t)
	if err != nil {
		panic(err)
	}
}

func parseFile(f string) ast.Node {
	input, err := antlr.NewFileStream(f)
	if err != nil {
		panic(err)
	}
	return parse(input, f)
}

func parseString(c string) ast.Node {
	input := antlr.NewInputStream(c)
	return parse(input, "")
}

func parse(input antlr.CharStream, f string) ast.Node {
	lexer := parser.NewapexLexer(input)
	stream := antlr.NewCommonTokenStream(lexer, 0)
	p := parser.NewapexParser(stream)
	p.AddErrorListener(antlr.NewDiagnosticErrorListener(true))
	p.BuildParseTrees = true
	tree := p.CompilationUnit()
	t := tree.Accept(&AstBuilder{
		CurrentFile: f,
	})
	return t.(ast.Node)
}

func convert(n ast.Node) (ast.Node, error) {
	return n, nil
}

func check(n ast.Node) error {
	checker := &visitor.SoqlChecker{}
	n.Accept(checker)
	return nil
}

func run(n ast.Node) error {
	interpreter := &ast.Interpreter{}
	_, err := n.Accept(interpreter)
	return err
}
