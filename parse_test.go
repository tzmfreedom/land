package main

import (
	"testing"

	"github.com/google/go-cmp/cmp"
	"github.com/k0kubun/pp"
	"github.com/tzmfreedom/goland/ast"
)

func TestParse(t *testing.T) {
	testCases := []struct {
		Code     string
		Expected ast.Node
	}{
		{
			`@foo public without sharing class Foo extends Bar implements Baz {}`,
			&ast.ClassDeclaration{
				Modifiers: []ast.Node{
					&ast.Modifier{
						Name: "public",
						Position: &ast.Position{
							Column: 5,
							Line:   1,
						},
					},
					&ast.Modifier{
						Name: "without sharing",
						Position: &ast.Position{
							Column: 12,
							Line:   1,
						},
					},
				},
				Annotations: []ast.Node{
					&ast.Annotation{
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
				ImplementClasses: []ast.Node{
					&ast.Type{
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
				Modifiers:   []ast.Node{},
				Annotations: []ast.Node{},
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
						Modifiers: []ast.Node{
							&ast.Modifier{
								Name: "public",
								Position: &ast.Position{
									FileName: "",
									Column:   0,
									Line:     2,
								},
							},
						},
						Declarators: []ast.Node{
							&ast.VariableDeclarator{
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
						Modifiers: []ast.Node{
							&ast.Modifier{
								Name: "public",
								Position: &ast.Position{
									FileName: "",
									Column:   0,
									Line:     3,
								},
							},
						},
						Declarators: []ast.Node{
							&ast.VariableDeclarator{
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
						Modifiers: []ast.Node{
							&ast.Modifier{
								Name: "public",
								Position: &ast.Position{
									FileName: "",
									Column:   0,
									Line:     4,
								},
							},
							&ast.Modifier{
								Name: "static",
								Position: &ast.Position{
									FileName: "",
									Column:   7,
									Line:     4,
								},
							},
						},
						Declarators: []ast.Node{
							&ast.VariableDeclarator{
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
						Modifiers: []ast.Node{
							&ast.Modifier{
								Name: "public",
								Position: &ast.Position{
									FileName: "",
									Column:   0,
									Line:     5,
								},
							},
							&ast.Modifier{
								Name: "static",
								Position: &ast.Position{
									FileName: "",
									Column:   7,
									Line:     5,
								},
							},
						},
						Declarators: []ast.Node{
							&ast.VariableDeclarator{
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
						Modifiers: []ast.Node{
							&ast.Modifier{
								Name: "public",
								Position: &ast.Position{
									FileName: "",
									Column:   0,
									Line:     6,
								},
							},
							&ast.Modifier{
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
						Parameters: []ast.Node{
							&ast.Parameter{
								Modifiers: []ast.Node{},
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
						Modifiers: []ast.Node{
							&ast.Modifier{
								Name: "public",
								Position: &ast.Position{
									FileName: "",
									Column:   0,
									Line:     7,
								},
							},
						},
						ReturnType: ast.VoidType,
						Parameters: []ast.Node{},
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
					Modifiers: []ast.Node{},
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
					Declarators: []ast.Node{
						&ast.VariableDeclarator{
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
					Modifiers: []ast.Node{},
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
					Declarators: []ast.Node{
						&ast.VariableDeclarator{
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
					Modifiers: []ast.Node{},
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
					Declarators: []ast.Node{
						&ast.VariableDeclarator{
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
					Modifiers: []ast.Node{},
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
					Declarators: []ast.Node{
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
					Modifiers: []ast.Node{},
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
					Declarators: []ast.Node{
						&ast.VariableDeclarator{
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
		{
			`class Foo {
public void action(){
  if (i == 1) {
    true;
  } else if (i == 2) {
    true;
  } else {
    false;
  }
  switch ON i {
    when 1 {
      true;
    }
    when 2, 3 {
      false;
    }
    when Account a {
      false;
    }
    else {
      1;
    }
  }
}
}`,
			createExpectedClass([]ast.Node{
				&ast.If{
					Condition: &ast.BinaryOperator{
						Op: "==",
						Left: &ast.Name{
							Value: "i",
							Position: &ast.Position{
								FileName: "",
								Column:   6,
								Line:     3,
							},
						},
						Right: &ast.IntegerLiteral{
							Value: 1,
							Position: &ast.Position{
								FileName: "",
								Column:   11,
								Line:     3,
							},
						},
						Position: &ast.Position{
							FileName: "",
							Column:   6,
							Line:     3,
						},
					},
					IfStatement: &ast.Block{
						Statements: []ast.Node{
							&ast.BooleanLiteral{
								Value: true,
								Position: &ast.Position{
									FileName: "",
									Column:   4,
									Line:     4,
								},
							},
						},
						Position: &ast.Position{
							FileName: "",
							Column:   14,
							Line:     3,
						},
					},
					ElseStatement: &ast.If{
						Condition: &ast.BinaryOperator{
							Op: "==",
							Left: &ast.Name{
								Value: "i",
								Position: &ast.Position{
									FileName: "",
									Column:   13,
									Line:     5,
								},
							},
							Right: &ast.IntegerLiteral{
								Value: 2,
								Position: &ast.Position{
									FileName: "",
									Column:   18,
									Line:     5,
								},
							},
							Position: &ast.Position{
								FileName: "",
								Column:   13,
								Line:     5,
							},
						},
						IfStatement: &ast.Block{
							Statements: []ast.Node{
								&ast.BooleanLiteral{
									Value: true,
									Position: &ast.Position{
										FileName: "",
										Column:   4,
										Line:     6,
									},
								},
							},
							Position: &ast.Position{
								FileName: "",
								Column:   21,
								Line:     5,
							},
						},
						ElseStatement: &ast.Block{
							Statements: []ast.Node{
								&ast.BooleanLiteral{
									Value: false,
									Position: &ast.Position{
										FileName: "",
										Column:   4,
										Line:     8,
									},
								},
							},
							Position: &ast.Position{
								FileName: "",
								Column:   9,
								Line:     7,
							},
						},
						Position: &ast.Position{
							FileName: "",
							Column:   9,
							Line:     5,
						},
					},
					Position: &ast.Position{
						FileName: "",
						Column:   2,
						Line:     3,
					},
				},
				&ast.Switch{
					Expression: &ast.Name{
						Value: "i",
						Position: &ast.Position{
							FileName: "",
							Column:   12,
							Line:     10,
						},
					},
					WhenStatements: []ast.Node{
						&ast.When{
							Condition: []ast.Node{
								&ast.IntegerLiteral{
									Value: 1,
									Position: &ast.Position{
										FileName: "",
										Column:   9,
										Line:     11,
									},
								},
							},
							Statements: &ast.Block{
								Statements: []ast.Node{
									&ast.BooleanLiteral{
										Value: true,
										Position: &ast.Position{
											FileName: "",
											Column:   6,
											Line:     12,
										},
									},
								},
								Position: &ast.Position{
									FileName: "",
									Column:   11,
									Line:     11,
								},
							},
							Position: &ast.Position{
								FileName: "",
								Column:   4,
								Line:     11,
							},
						},
						&ast.When{
							Condition: []ast.Node{
								&ast.IntegerLiteral{
									Value: 2,
									Position: &ast.Position{
										FileName: "",
										Column:   9,
										Line:     14,
									},
								},
								&ast.IntegerLiteral{
									Value: 3,
									Position: &ast.Position{
										FileName: "",
										Column:   12,
										Line:     14,
									},
								},
							},
							Statements: &ast.Block{
								Statements: []ast.Node{
									&ast.BooleanLiteral{
										Value: false,
										Position: &ast.Position{
											FileName: "",
											Column:   6,
											Line:     15,
										},
									},
								},
								Position: &ast.Position{
									FileName: "",
									Column:   14,
									Line:     14,
								},
							},
							Position: &ast.Position{
								FileName: "",
								Column:   4,
								Line:     14,
							},
						},
						&ast.When{
							Condition: []ast.Node{
								&ast.WhenType{
									Type: &ast.Type{
										Name: []string{
											"Account",
										},
										Parameters: []ast.Node{},
										Position: &ast.Position{
											FileName: "",
											Column:   9,
											Line:     17,
										},
									},
									Identifier: "a",
									Position: &ast.Position{
										FileName: "",
										Column:   9,
										Line:     17,
									},
								},
							},
							Statements: &ast.Block{
								Statements: []ast.Node{
									&ast.BooleanLiteral{
										Value: false,
										Position: &ast.Position{
											FileName: "",
											Column:   6,
											Line:     18,
										},
									},
								},
								Position: &ast.Position{
									FileName: "",
									Column:   19,
									Line:     17,
								},
							},
							Position: &ast.Position{
								FileName: "",
								Column:   4,
								Line:     17,
							},
						},
					},
					ElseStatement: &ast.Block{
						Statements: []ast.Node{
							&ast.IntegerLiteral{
								Value: 1,
								Position: &ast.Position{
									FileName: "",
									Column:   6,
									Line:     21,
								},
							},
						},
						Position: &ast.Position{
							FileName: "",
							Column:   9,
							Line:     20,
						},
					},
					Position: &ast.Position{
						FileName: "",
						Column:   2,
						Line:     10,
					},
				},
			}),
		},
		{
			`class Foo {
public void action(){
  for (Integer i = 0; i < imax; i++) {
    continue;
  }
  while (true) {
    break;
  }
  do {
    return;
  } while (false)
  for (Account acc : accounts) {
    return 1;
  }
}
}`,
			createExpectedClass([]ast.Node{
				&ast.For{
					Control: &ast.ForControl{
						ForInit: &ast.VariableDeclaration{
							Modifiers: []ast.Node{},
							Type: &ast.Type{
								Name: []string{
									"Integer",
								},
								Parameters: []ast.Node{},
								Position: &ast.Position{
									FileName: "",
									Column:   7,
									Line:     3,
								},
							},
							Declarators: []ast.Node{
								&ast.VariableDeclarator{
									Name: "i",
									Expression: &ast.IntegerLiteral{
										Value: 0,
										Position: &ast.Position{
											FileName: "",
											Column:   19,
											Line:     3,
										},
									},
									Position: &ast.Position{
										FileName: "",
										Column:   15,
										Line:     3,
									},
								},
							},
							Position: &ast.Position{
								FileName: "",
								Column:   7,
								Line:     3,
							},
						},
						Expression: &ast.BinaryOperator{
							Op: "<",
							Left: &ast.Name{
								Value: "i",
								Position: &ast.Position{
									FileName: "",
									Column:   22,
									Line:     3,
								},
							},
							Right: &ast.Name{
								Value: "imax",
								Position: &ast.Position{
									FileName: "",
									Column:   26,
									Line:     3,
								},
							},
							Position: &ast.Position{
								FileName: "",
								Column:   22,
								Line:     3,
							},
						},
						ForUpdate: []ast.Node{
							&ast.UnaryOperator{
								Op: "++",
								Expression: &ast.Name{
									Value: "i",
									Position: &ast.Position{
										FileName: "",
										Column:   32,
										Line:     3,
									},
								},
								IsPrefix: false,
								Position: &ast.Position{
									FileName: "",
									Column:   32,
									Line:     3,
								},
							},
						},
						Position: &ast.Position{
							FileName: "",
							Column:   7,
							Line:     3,
						},
					},
					Statements: &ast.Block{
						Statements: []ast.Node{
							&ast.Continue{
								Position: &ast.Position{
									FileName: "",
									Column:   4,
									Line:     4,
								},
							},
						},
						Position: &ast.Position{
							FileName: "",
							Column:   37,
							Line:     3,
						},
					},
					Position: &ast.Position{
						FileName: "",
						Column:   2,
						Line:     3,
					},
				},
				&ast.While{
					Condition: &ast.BooleanLiteral{
						Value: true,
						Position: &ast.Position{
							FileName: "",
							Column:   9,
							Line:     6,
						},
					},
					Statements: []ast.Node{
						&ast.Block{
							Statements: []ast.Node{
								&ast.Break{
									Position: &ast.Position{
										FileName: "",
										Column:   4,
										Line:     7,
									},
								},
							},
							Position: &ast.Position{
								FileName: "",
								Column:   15,
								Line:     6,
							},
						},
					},
					IsDo: false,
					Position: &ast.Position{
						FileName: "",
						Column:   2,
						Line:     6,
					},
				},
				&ast.While{
					Condition: &ast.BooleanLiteral{
						Value: false,
						Position: &ast.Position{
							FileName: "",
							Column:   11,
							Line:     11,
						},
					},
					Statements: []ast.Node{
						&ast.Block{
							Statements: []ast.Node{
								&ast.Return{
									Expression: nil,
									Position: &ast.Position{
										FileName: "",
										Column:   4,
										Line:     10,
									},
								},
							},
							Position: &ast.Position{
								FileName: "",
								Column:   5,
								Line:     9,
							},
						},
					},
					IsDo: true,
					Position: &ast.Position{
						FileName: "",
						Column:   2,
						Line:     9,
					},
				},
				&ast.For{
					Control: &ast.EnhancedForControl{
						Modifiers: []ast.Node{},
						Type: &ast.Type{
							Name: []string{
								"Account",
							},
							Parameters: []ast.Node{},
							Position: &ast.Position{
								FileName: "",
								Column:   7,
								Line:     12,
							},
						},
						VariableDeclaratorId: "acc",
						Expression: &ast.Name{
							Value: "accounts",
							Position: &ast.Position{
								FileName: "",
								Column:   21,
								Line:     12,
							},
						},
						Position: &ast.Position{
							FileName: "",
							Column:   7,
							Line:     12,
						},
					},
					Statements: &ast.Block{
						Statements: []ast.Node{
							&ast.Return{
								Expression: &ast.IntegerLiteral{
									Value: 1,
									Position: &ast.Position{
										FileName: "",
										Column:   11,
										Line:     13,
									},
								},
								Position: &ast.Position{
									FileName: "",
									Column:   4,
									Line:     13,
								},
							},
						},
						Position: &ast.Position{
							FileName: "",
							Column:   31,
							Line:     12,
						},
					},
					Position: &ast.Position{
						FileName: "",
						Column:   2,
						Line:     12,
					},
				},
			}),
		},
		{
			`class Foo {
public void action(){
try {
  throw a;
} catch (Exception e) {
  return;
} finally {
  return;
}
}
}`,
			createExpectedClass([]ast.Node{
				&ast.Try{
					Block: &ast.Block{
						Statements: []ast.Node{
							&ast.Throw{
								Expression: &ast.Name{
									Value: "a",
									Position: &ast.Position{
										FileName: "",
										Column:   8,
										Line:     4,
									},
								},
								Position: &ast.Position{
									FileName: "",
									Column:   2,
									Line:     4,
								},
							},
						},
						Position: &ast.Position{
							FileName: "",
							Column:   4,
							Line:     3,
						},
					},
					CatchClause: []ast.Node{
						&ast.Catch{
							Modifiers: []ast.Node{},
							Type: &ast.Type{
								Name: []string{
									"Exception",
								},
								Parameters: nil,
								Position: &ast.Position{
									FileName: "",
									Column:   9,
									Line:     5,
								},
							},
							Identifier: "e",
							Block: &ast.Block{
								Statements: []ast.Node{
									&ast.Return{
										Expression: nil,
										Position: &ast.Position{
											FileName: "",
											Column:   2,
											Line:     6,
										},
									},
								},
								Position: &ast.Position{
									FileName: "",
									Column:   22,
									Line:     5,
								},
							},
							Position: &ast.Position{
								FileName: "",
								Column:   2,
								Line:     5,
							},
						},
					},
					FinallyBlock: &ast.Block{
						Statements: []ast.Node{
							&ast.Return{
								Expression: nil,
								Position: &ast.Position{
									FileName: "",
									Column:   2,
									Line:     8,
								},
							},
						},
						Position: &ast.Position{
							FileName: "",
							Column:   10,
							Line:     7,
						},
					},
					Position: (*ast.Position)(nil),
				},
			}),
		},
		{
			`trigger Foo on Account(before insert, after update) {
true;
}`,
			&ast.Trigger{
				Name:   "Foo",
				Object: "Account",
				TriggerTimings: []ast.Node{
					&ast.TriggerTiming{
						Dml:    "insert",
						Timing: "before",
					},
					&ast.TriggerTiming{
						Dml:    "update",
						Timing: "after",
					},
				},
				Statements: &ast.Block{
					Statements: []ast.Node{
						&ast.BooleanLiteral{
							Value: true,
						},
					},
				},
			},
		},
	}
	for _, testCase := range testCases {
		out := parseString(testCase.Code)

		expected := ast.Dump(testCase.Expected, 0)
		actual := ast.Dump(out, 0)
		ok := cmp.Equal(expected, actual)
		if !ok {
			diff := cmp.Diff(testCase.Expected, out)
			pp.Print(out)
			t.Errorf(diff)
		}
	}
}

func createExpectedClass(statements []ast.Node) *ast.ClassDeclaration {
	return &ast.ClassDeclaration{
		Modifiers:   []ast.Node{},
		Annotations: []ast.Node{},
		Name:        "Foo",
		Position: &ast.Position{
			Column: 0,
			Line:   1,
		},
		Declarations: []ast.Node{
			&ast.MethodDeclaration{
				Name: "action",
				Modifiers: []ast.Node{
					&ast.Modifier{
						Name: "public",
						Position: &ast.Position{
							FileName: "",
							Column:   0,
							Line:     2,
						},
					},
				},
				ReturnType: ast.VoidType,
				Parameters: []ast.Node{},
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
