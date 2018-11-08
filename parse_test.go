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
			`@foo public without sharing class Foo extends Bar implements Baz {}`,
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
						Name: "foo",
						Position: &ast.Position{
							Column: 0,
							Line:   1,
						},
					},
				},
				Name: "Foo",
				Position: &ast.Position{
					Column: 28,
					Line:   1,
				},
				SuperClass: ast.Type{
					Name: []string{
						"Bar",
					},
					Parameters: []ast.Node{},
					Position: &ast.Position{
						FileName: "",
						Column:   46,
						Line:     1,
					},
				},
				ImplementClasses: []ast.Type{
					{
						Name: []string{
							"Baz",
						},
						Parameters: []ast.Node{},
						Position: &ast.Position{
							FileName: "",
							Column:   61,
							Line:     1,
						},
					},
				},
				Declarations: []ast.Node{},
			},
		},
		{
			`class Foo {
public Integer field;
public Double field_with_init = 2;
public static String static_field;
public static Boolean static_field_with_init = 1;
public static String static_method(Boolean p1) { }
public void static_method(){ }
}`,
			&ast.ClassDeclaration{
				Modifiers:   []ast.Modifier{},
				Annotations: []ast.Annotation{},
				Name:        "Foo",
				Position: &ast.Position{
					Column: 0,
					Line:   1,
				},
				Declarations: []ast.Node{
					&ast.FieldDeclaration{
						Type: ast.Type{
							Name: []string{
								"Integer",
							},
							Parameters: []ast.Node{},
							Position: &ast.Position{
								FileName: "",
								Column:   7,
								Line:     2,
							},
						},
						Modifiers: []ast.Modifier{
							ast.Modifier{
								Name: "public",
								Position: &ast.Position{
									FileName: "",
									Column:   0,
									Line:     2,
								},
							},
						},
						Declarators: []ast.VariableDeclarator{
							ast.VariableDeclarator{
								Name: "field",
								Expression: &ast.NullLiteral{
									Position: &ast.Position{
										FileName: "",
										Column:   15,
										Line:     2,
									},
								},
								Position: &ast.Position{
									FileName: "",
									Column:   15,
									Line:     2,
								},
							},
						},
						Position: nil,
					},
					&ast.FieldDeclaration{
						Type: ast.Type{
							Name: []string{
								"Double",
							},
							Parameters: []ast.Node{},
							Position: &ast.Position{
								FileName: "",
								Column:   7,
								Line:     3,
							},
						},
						Modifiers: []ast.Modifier{
							ast.Modifier{
								Name: "public",
								Position: &ast.Position{
									FileName: "",
									Column:   0,
									Line:     3,
								},
							},
						},
						Declarators: []ast.VariableDeclarator{
							ast.VariableDeclarator{
								Name: "field_with_init",
								Expression: &ast.IntegerLiteral{
									Value: 2,
									Position: &ast.Position{
										FileName: "",
										Column:   32,
										Line:     3,
									},
								},
								Position: &ast.Position{
									FileName: "",
									Column:   14,
									Line:     3,
								},
							},
						},
						Position: nil,
					},
					&ast.FieldDeclaration{
						Type: ast.Type{
							Name: []string{
								"String",
							},
							Parameters: []ast.Node{},
							Position: &ast.Position{
								FileName: "",
								Column:   14,
								Line:     4,
							},
						},
						Modifiers: []ast.Modifier{
							ast.Modifier{
								Name: "public",
								Position: &ast.Position{
									FileName: "",
									Column:   0,
									Line:     4,
								},
							},
							ast.Modifier{
								Name: "static",
								Position: &ast.Position{
									FileName: "",
									Column:   7,
									Line:     4,
								},
							},
						},
						Declarators: []ast.VariableDeclarator{
							ast.VariableDeclarator{
								Name: "static_field",
								Expression: &ast.NullLiteral{
									Position: &ast.Position{
										FileName: "",
										Column:   21,
										Line:     4,
									},
								},
								Position: &ast.Position{
									FileName: "",
									Column:   21,
									Line:     4,
								},
							},
						},
						Position: nil,
					},
					&ast.FieldDeclaration{
						Type: ast.Type{
							Name: []string{
								"Boolean",
							},
							Parameters: []ast.Node{},
							Position: &ast.Position{
								FileName: "",
								Column:   14,
								Line:     5,
							},
						},
						Modifiers: []ast.Modifier{
							ast.Modifier{
								Name: "public",
								Position: &ast.Position{
									FileName: "",
									Column:   0,
									Line:     5,
								},
							},
							ast.Modifier{
								Name: "static",
								Position: &ast.Position{
									FileName: "",
									Column:   7,
									Line:     5,
								},
							},
						},
						Declarators: []ast.VariableDeclarator{
							ast.VariableDeclarator{
								Name: "static_field_with_init",
								Expression: &ast.IntegerLiteral{
									Value: 1,
									Position: &ast.Position{
										FileName: "",
										Column:   47,
										Line:     5,
									},
								},
								Position: &ast.Position{
									FileName: "",
									Column:   22,
									Line:     5,
								},
							},
						},
						Position: nil,
					},
					&ast.MethodDeclaration{
						Name: "static_method",
						Modifiers: []ast.Modifier{
							ast.Modifier{
								Name: "public",
								Position: &ast.Position{
									FileName: "",
									Column:   0,
									Line:     6,
								},
							},
							ast.Modifier{
								Name: "static",
								Position: &ast.Position{
									FileName: "",
									Column:   7,
									Line:     6,
								},
							},
						},
						ReturnType: ast.Type{
							Name: []string{
								"String",
							},
							Parameters: []ast.Node{},
							Position: &ast.Position{
								FileName: "",
								Column:   14,
								Line:     6,
							},
						},
						Parameters: []ast.Parameter{
							ast.Parameter{
								Modifiers: []ast.Modifier{},
								Type: ast.Type{
									Name: []string{
										"Boolean",
									},
									Parameters: []ast.Node{},
									Position: &ast.Position{
										FileName: "",
										Column:   35,
										Line:     6,
									},
								},
								Name: "p1",
								Position: &ast.Position{
									FileName: "",
									Column:   35,
									Line:     6,
								},
							},
						},
						Throws: []ast.Node{},
						Statements: ast.Block{
							Statements: []ast.Node{},
							Position: &ast.Position{
								FileName: "",
								Column:   47,
								Line:     6,
							},
						},
						NativeFunction: nil,
						Position: &ast.Position{
							FileName: "",
							Column:   14,
							Line:     6,
						},
					},
					&ast.MethodDeclaration{
						Name: "static_method",
						Modifiers: []ast.Modifier{
							ast.Modifier{
								Name: "public",
								Position: &ast.Position{
									FileName: "",
									Column:   0,
									Line:     7,
								},
							},
						},
						ReturnType: *ast.VoidType,
						Parameters: []ast.Parameter{},
						Throws:     []ast.Node{},
						Statements: ast.Block{
							Statements: []ast.Node{},
							Position: &ast.Position{
								FileName: "",
								Column:   27,
								Line:     7,
							},
						},
						NativeFunction: nil,
						Position: &ast.Position{
							FileName: "",
							Column:   7,
							Line:     7,
						},
					},
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
