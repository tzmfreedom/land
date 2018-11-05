package main

import (
	"github.com/google/go-cmp/cmp"
	"github.com/k0kubun/pp"
	"github.com/tzmfreedom/goland/ast"
	"testing"
)

func TestParse(t *testing.T) {
	testCases := []struct {
		Code     string
		Expected *ast.ClassDeclaration
	}{
		{
			`@foo public without sharing class Hoge {}`,
			&ast.ClassDeclaration{
				Modifiers: []ast.Modifier{
					{
						Name: "public",
						Position: &ast.Position{
							Column: 5,
							Line:   1,
						},
					},
					{
						Name: "without sharing",
						Position: &ast.Position{
							Column: 12,
							Line:   1,
						},
					},
				},
				Annotations: []ast.Annotation{
					{
						Name: ast.Name{
							Value: []string{"foo"},
							Position: &ast.Position{
								Column: 1,
								Line:   1,
							},
						},
						Position: &ast.Position{
							Column: 0,
							Line:   1,
						},
					},
				},
				Name: "Hoge",
				Position: &ast.Position{
					Column: 28,
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
			pp.Print(out)
			t.Errorf(diff)
		}
	}
}
