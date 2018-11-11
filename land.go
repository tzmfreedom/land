package main

import (
	"os"

	"github.com/antlr/antlr4/runtime/Go/antlr"
	"github.com/k0kubun/pp"
	"github.com/tzmfreedom/goland/ast"
	"github.com/tzmfreedom/goland/parser"
	"github.com/tzmfreedom/goland/visitor"
)

func main() {
	f := os.Args[1]
	t := parseFile(f)
	pp.Print(t)
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
	n := t.(ast.Node)
	checker := &visitor.SoqlChecker{}
	n.Accept(checker)
	return n
}
