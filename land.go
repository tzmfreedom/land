package main

import (
	"os"

	"github.com/antlr/antlr4/runtime/Go/antlr"
	"github.com/k0kubun/pp"
	"github.com/tzmfreedom/goland/ast"
	"github.com/tzmfreedom/goland/parser"
)

func main() {
	f := os.Args[1]
	t := parseFile(f)
	pp.Print(t)
}

func parseFile(f string) *ast.ClassDeclaration {
	input, err := antlr.NewFileStream(f)
	if err != nil {
		panic(err)
	}
	return parse(input, f)
}

func parseString(c string) *ast.ClassDeclaration {
	input := antlr.NewInputStream(c)
	return parse(input, "")
}

func parse(input antlr.CharStream, f string) *ast.ClassDeclaration {
	lexer := parser.NewapexLexer(input)
	stream := antlr.NewCommonTokenStream(lexer, 0)
	p := parser.NewapexParser(stream)
	p.AddErrorListener(antlr.NewDiagnosticErrorListener(true))
	p.BuildParseTrees = true
	tree := p.CompilationUnit()
	t := tree.Accept(&AstBuilder{
		CurrentFile: f,
	})
	if cd, ok := t.(ast.ClassDeclaration); ok {
		return &cd
	}
	return nil
}
