package ast

import (
	"io/ioutil"

	"github.com/antlr/antlr4/runtime/Go/antlr"
	"github.com/tzmfreedom/land/parser"
)

type PreProcessor func(string) string

func ParseFile(f string, processors ...PreProcessor) (Node, error) {
	bytes, err := ioutil.ReadFile(f)
	if err != nil {
		return nil, err
	}
	src := string(bytes)
	for _, processor := range processors {
		src = processor(src)
	}
	input := antlr.NewInputStream(src)
	return parse(input, f), nil
}

func ParseString(src string, processors ...PreProcessor) (Node, error) {
	for _, processor := range processors {
		src = processor(src)
	}
	input := antlr.NewInputStream(src)
	return parse(input, "<string>"), nil
}

func parse(input antlr.CharStream, src string) Node {
	lexer := parser.NewapexLexer(input)
	stream := antlr.NewCommonTokenStream(lexer, 0)
	p := parser.NewapexParser(stream)
	// p.AddErrorListener(antlr.NewDiagnosticErrorListener(true))
	p.BuildParseTrees = true
	tree := p.CompilationUnit()
	t := tree.Accept(&Builder{
		Source: src,
	})
	return t.(Node)
}
