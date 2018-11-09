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
				Modifiers: []*ast.Modifier{
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
				Annotations: []*ast.Annotation{
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
				SuperClass: &ast.Type{
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
				ImplementClasses: []*ast.Type{
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
public void method(){ }
}`,
			&ast.ClassDeclaration{
				Modifiers:   []*ast.Modifier{},
				Annotations: []*ast.Annotation{},
				Name:        "Foo",
				Position: &ast.Position{
					Column: 0,
					Line:   1,
				},
				Declarations: []ast.Node{
					&ast.FieldDeclaration{
						Type: &ast.Type{
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
						Modifiers: []*ast.Modifier{
							{
								Name: "public",
								Position: &ast.Position{
									FileName: "",
									Column:   0,
									Line:     2,
								},
							},
						},
						Declarators: []*ast.VariableDeclarator{
							{
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
						Type: &ast.Type{
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
						Modifiers: []*ast.Modifier{
							{
								Name: "public",
								Position: &ast.Position{
									FileName: "",
									Column:   0,
									Line:     3,
								},
							},
						},
						Declarators: []*ast.VariableDeclarator{
							{
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
						Type: &ast.Type{
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
						Modifiers: []*ast.Modifier{
							{
								Name: "public",
								Position: &ast.Position{
									FileName: "",
									Column:   0,
									Line:     4,
								},
							},
							{
								Name: "static",
								Position: &ast.Position{
									FileName: "",
									Column:   7,
									Line:     4,
								},
							},
						},
						Declarators: []*ast.VariableDeclarator{
							{
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
						Type: &ast.Type{
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
						Modifiers: []*ast.Modifier{
							{
								Name: "public",
								Position: &ast.Position{
									FileName: "",
									Column:   0,
									Line:     5,
								},
							},
							{
								Name: "static",
								Position: &ast.Position{
									FileName: "",
									Column:   7,
									Line:     5,
								},
							},
						},
						Declarators: []*ast.VariableDeclarator{
							{
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
						Modifiers: []*ast.Modifier{
							{
								Name: "public",
								Position: &ast.Position{
									FileName: "",
									Column:   0,
									Line:     6,
								},
							},
							{
								Name: "static",
								Position: &ast.Position{
									FileName: "",
									Column:   7,
									Line:     6,
								},
							},
						},
						ReturnType: &ast.Type{
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
						Parameters: []*ast.Parameter{
							{
								Modifiers: []*ast.Modifier{},
								Type: &ast.Type{
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
						Statements: &ast.Block{
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
						Name: "method",
						Modifiers: []*ast.Modifier{
							{
								Name: "public",
								Position: &ast.Position{
									FileName: "",
									Column:   0,
									Line:     7,
								},
							},
						},
						ReturnType: ast.VoidType,
						Parameters: []*ast.Parameter{},
						Throws:     []ast.Node{},
						Statements: &ast.Block{
							Statements: []ast.Node{},
							Position: &ast.Position{
								FileName: "",
								Column:   20,
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
		{
			`class Foo {
public void action(){
Integer i = 0;
String s = 'abc';
Double d = 1.23;
Boolean b = true;
Integer i;
}
}`,
			createExpectedClass([]ast.Node{
				&ast.VariableDeclaration{
					Modifiers: []*ast.Modifier{},
					Type: &ast.Type{
						Name: []string{
							"Integer",
						},
						Parameters: []ast.Node{},
						Position: &ast.Position{
							FileName: "",
							Column:   0,
							Line:     3,
						},
					},
					Declarators: []*ast.VariableDeclarator{
						{
							Name: "i",
							Expression: &ast.IntegerLiteral{
								Value: 0,
								Position: &ast.Position{
									FileName: "",
									Column:   12,
									Line:     3,
								},
							},
							Position: &ast.Position{
								FileName: "",
								Column:   8,
								Line:     3,
							},
						},
					},
					Position: &ast.Position{
						FileName: "",
						Column:   0,
						Line:     3,
					},
				},
				&ast.VariableDeclaration{
					Modifiers: []*ast.Modifier{},
					Type: &ast.Type{
						Name: []string{
							"String",
						},
						Parameters: []ast.Node{},
						Position: &ast.Position{
							FileName: "",
							Column:   0,
							Line:     4,
						},
					},
					Declarators: []*ast.VariableDeclarator{
						{
							Name: "s",
							Expression: &ast.StringLiteral{
								Value: "'ab",
								Position: &ast.Position{
									FileName: "",
									Column:   11,
									Line:     4,
								},
							},
							Position: &ast.Position{
								FileName: "",
								Column:   7,
								Line:     4,
							},
						},
					},
					Position: &ast.Position{
						FileName: "",
						Column:   0,
						Line:     4,
					},
				},
				&ast.VariableDeclaration{
					Modifiers: []*ast.Modifier{},
					Type: &ast.Type{
						Name: []string{
							"Double",
						},
						Parameters: []ast.Node{},
						Position: &ast.Position{
							FileName: "",
							Column:   0,
							Line:     5,
						},
					},
					Declarators: []*ast.VariableDeclarator{
						{
							Name: "d",
							Expression: &ast.DoubleLiteral{
								Value: 1.230000,
								Position: &ast.Position{
									FileName: "",
									Column:   11,
									Line:     5,
								},
							},
							Position: &ast.Position{
								FileName: "",
								Column:   7,
								Line:     5,
							},
						},
					},
					Position: &ast.Position{
						FileName: "",
						Column:   0,
						Line:     5,
					},
				},
				&ast.VariableDeclaration{
					Modifiers: []*ast.Modifier{},
					Type: &ast.Type{
						Name: []string{
							"Boolean",
						},
						Parameters: []ast.Node{},
						Position: &ast.Position{
							FileName: "",
							Column:   0,
							Line:     6,
						},
					},
					Declarators: []*ast.VariableDeclarator{
						&ast.VariableDeclarator{
							Name: "b",
							Expression: &ast.BooleanLiteral{
								Value: true,
								Position: &ast.Position{
									FileName: "",
									Column:   12,
									Line:     6,
								},
							},
							Position: &ast.Position{
								FileName: "",
								Column:   8,
								Line:     6,
							},
						},
					},
					Position: &ast.Position{
						FileName: "",
						Column:   0,
						Line:     6,
					},
				},
				&ast.VariableDeclaration{
					Modifiers: []*ast.Modifier{},
					Type: &ast.Type{
						Name: []string{
							"Integer",
						},
						Parameters: []ast.Node{},
						Position: &ast.Position{
							FileName: "",
							Column:   0,
							Line:     7,
						},
					},
					Declarators: []*ast.VariableDeclarator{
						{
							Name: "i",
							Expression: &ast.NullLiteral{
								Position: &ast.Position{
									FileName: "",
									Column:   8,
									Line:     7,
								},
							},
							Position: &ast.Position{
								FileName: "",
								Column:   8,
								Line:     7,
							},
						},
					},
					Position: &ast.Position{
						FileName: "",
						Column:   0,
						Line:     7,
					},
				},
			}),
		},
		{
			`class Foo {
public void action(){
i = 1 + 2 + 3;
i = 1 + 2 * 3;
i = (1 + 2) * 3;
}
}`,
			createExpectedClass([]ast.Node{
				&ast.BinaryOperator{
					Op: "=",
					Left: &ast.Name{
						Value: "i",
						Position: &ast.Position{
							FileName: "",
							Column:   0,
							Line:     3,
						},
					},
					Right: &ast.BinaryOperator{
						Op: "+",
						Left: &ast.BinaryOperator{
							Op: "+",
							Left: &ast.IntegerLiteral{
								Value: 1,
								Position: &ast.Position{
									FileName: "",
									Column:   4,
									Line:     3,
								},
							},
							Right: &ast.IntegerLiteral{
								Value: 2,
								Position: &ast.Position{
									FileName: "",
									Column:   8,
									Line:     3,
								},
							},
							Position: &ast.Position{
								FileName: "",
								Column:   4,
								Line:     3,
							},
						},
						Right: &ast.IntegerLiteral{
							Value: 3,
							Position: &ast.Position{
								FileName: "",
								Column:   12,
								Line:     3,
							},
						},
						Position: &ast.Position{
							FileName: "",
							Column:   4,
							Line:     3,
						},
					},
					Position: &ast.Position{
						FileName: "",
						Column:   0,
						Line:     3,
					},
				},
				&ast.BinaryOperator{
					Op: "=",
					Left: &ast.Name{
						Value: "i",
						Position: &ast.Position{
							FileName: "",
							Column:   0,
							Line:     4,
						},
					},
					Right: &ast.BinaryOperator{
						Op: "+",
						Left: &ast.IntegerLiteral{
							Value: 1,
							Position: &ast.Position{
								FileName: "",
								Column:   4,
								Line:     4,
							},
						},
						Right: &ast.BinaryOperator{
							Op: "*",
							Left: &ast.IntegerLiteral{
								Value: 2,
								Position: &ast.Position{
									FileName: "",
									Column:   8,
									Line:     4,
								},
							},
							Right: &ast.IntegerLiteral{
								Value: 3,
								Position: &ast.Position{
									FileName: "",
									Column:   12,
									Line:     4,
								},
							},
							Position: &ast.Position{
								FileName: "",
								Column:   8,
								Line:     4,
							},
						},
						Position: &ast.Position{
							FileName: "",
							Column:   4,
							Line:     4,
						},
					},
					Position: &ast.Position{
						FileName: "",
						Column:   0,
						Line:     4,
					},
				},
				&ast.BinaryOperator{
					Op: "=",
					Left: &ast.Name{
						Value: "i",
						Position: &ast.Position{
							FileName: "",
							Column:   0,
							Line:     5,
						},
					},
					Right: &ast.BinaryOperator{
						Op: "*",
						Left: &ast.BinaryOperator{
							Op: "+",
							Left: &ast.IntegerLiteral{
								Value: 1,
								Position: &ast.Position{
									FileName: "",
									Column:   5,
									Line:     5,
								},
							},
							Right: &ast.IntegerLiteral{
								Value: 2,
								Position: &ast.Position{
									FileName: "",
									Column:   9,
									Line:     5,
								},
							},
							Position: &ast.Position{
								FileName: "",
								Column:   5,
								Line:     5,
							},
						},
						Right: &ast.IntegerLiteral{
							Value: 3,
							Position: &ast.Position{
								FileName: "",
								Column:   14,
								Line:     5,
							},
						},
						Position: &ast.Position{
							FileName: "",
							Column:   4,
							Line:     5,
						},
					},
					Position: &ast.Position{
						FileName: "",
						Column:   0,
						Line:     5,
					},
				},
			}),
		},
		{
			`class Foo {
public void action(){
foo.bar.baz;
foo.bar.baz();
foo(1, s, false);
}
}`,
			createExpectedClass([]ast.Node{
				&ast.FieldAccess{
					Expression: &ast.FieldAccess{
						Expression: &ast.Name{
							Value: "foo",
							Position: &ast.Position{
								FileName: "",
								Column:   0,
								Line:     3,
							},
						},
						FieldName: "bar",
						Position: &ast.Position{
							FileName: "",
							Column:   0,
							Line:     3,
						},
					},
					FieldName: "baz",
					Position: &ast.Position{
						FileName: "",
						Column:   0,
						Line:     3,
					},
				},
				&ast.MethodInvocation{
					NameOrExpression: &ast.FieldAccess{
						Expression: &ast.FieldAccess{
							Expression: &ast.Name{
								Value: "foo",
								Position: &ast.Position{
									FileName: "",
									Column:   0,
									Line:     4,
								},
							},
							FieldName: "bar",
							Position: &ast.Position{
								FileName: "",
								Column:   0,
								Line:     4,
							},
						},
						FieldName: "baz",
						Position: &ast.Position{
							FileName: "",
							Column:   0,
							Line:     4,
						},
					},
					Parameters: []ast.Node{},
					Position: &ast.Position{
						FileName: "",
						Column:   0,
						Line:     4,
					},
				},
				&ast.MethodInvocation{
					NameOrExpression: &ast.Name{
						Value: "foo",
						Position: &ast.Position{
							FileName: "",
							Column:   0,
							Line:     5,
						},
					},
					Parameters: []ast.Node{
						&ast.IntegerLiteral{
							Value: 1,
							Position: &ast.Position{
								FileName: "",
								Column:   4,
								Line:     5,
							},
						},
						&ast.Name{
							Value: "s",
							Position: &ast.Position{
								FileName: "",
								Column:   7,
								Line:     5,
							},
						},
						&ast.BooleanLiteral{
							Value: false,
							Position: &ast.Position{
								FileName: "",
								Column:   10,
								Line:     5,
							},
						},
					},
					Position: &ast.Position{
						FileName: "",
						Column:   0,
						Line:     5,
					},
				},
			}),
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

func createExpectedClass(statements []ast.Node) *ast.ClassDeclaration {
	return &ast.ClassDeclaration{
		Modifiers:   []*ast.Modifier{},
		Annotations: []*ast.Annotation{},
		Name:        "Foo",
		Position: &ast.Position{
			Column: 0,
			Line:   1,
		},
		Declarations: []ast.Node{
			&ast.MethodDeclaration{
				Name: "action",
				Modifiers: []*ast.Modifier{
					{
						Name: "public",
						Position: &ast.Position{
							FileName: "",
							Column:   0,
							Line:     2,
						},
					},
				},
				ReturnType: ast.VoidType,
				Parameters: []*ast.Parameter{},
				Throws:     []ast.Node{},
				Statements: &ast.Block{
					Statements: statements,
					Position: &ast.Position{
						FileName: "",
						Column:   20,
						Line:     2,
					},
				},
				NativeFunction: nil,
				Position: &ast.Position{
					FileName: "",
					Column:   7,
					Line:     2,
				},
			},
		},
	}
}
