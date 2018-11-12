package ast

import (
	"github.com/antlr/antlr4/runtime/Go/antlr"
	"github.com/tzmfreedom/goland/parser"
)

func ParseFile(f string) (Node, error) {
	input, err := antlr.NewFileStream(f)
	if err != nil {
		return nil, err
	}
	return parse(input, f), nil
}

func ParseString(c string) (Node, error) {
	input := antlr.NewInputStream(c)
	return parse(input, "<string>"), nil
}

func parse(input antlr.CharStream, src string) Node {
	lexer := parser.NewapexLexer(input)
	stream := antlr.NewCommonTokenStream(lexer, 0)
	p := parser.NewapexParser(stream)
	p.AddErrorListener(antlr.NewDiagnosticErrorListener(true))
	p.BuildParseTrees = true
	tree := p.CompilationUnit()
	t := tree.Accept(&Builder{
		Source: src,
	})
	return t.(Node)
}
