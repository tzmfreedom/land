package main

import (
	"github.com/google/go-cmp/cmp"
	"github.com/tzmfreedom/goland/ast"
	"testing"
)

func TestParse(t *testing.T) {
	testCases := []struct {
		Code     string
		Expected *ast.ClassDeclaration
	}{
		{
			`public without sharing class Hoge {}`,
			&ast.ClassDeclaration{
				Modifiers: []ast.Modifier{
					{
						Name: "public",
						Position: &ast.Position{
							Column: 0,
							Line:   1,
						},
					},
					{
						Name: "without sharing",
						Position: &ast.Position{
							Column: 7,
							Line:   1,
						},
					},
				},
				Name: "Hoge",
				Position: &ast.Position{
					Column: 23,
					Line:   1,
				},
			},
		},
	}
	for _, testCase := range testCases {
		out := parseString(testCase.Code)

		ok := cmp.Equal(testCase.Expected, out)
		if !ok {
			diff := cmp.Diff(testCase.Expected, out)
			t.Errorf(diff)
		}
	}
}
