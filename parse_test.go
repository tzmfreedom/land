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
					},
					&ast.Modifier{
						Name: "without sharing",
					},
				},
				Annotations: []ast.Node{
					&ast.Annotation{
						Name: "foo",
					},
				},
				Name: "Foo",
				SuperClass: &ast.Type{
					Name: []string{
						"Bar",
					},
					Parameters: []ast.Node{},
				},
				ImplementClasses: []ast.Node{
					&ast.Type{
						Name: []string{
							"Baz",
						},
						Parameters: []ast.Node{},
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
				Declarations: []ast.Node{
					&ast.FieldDeclaration{
						Type: &ast.Type{
							Name: []string{
								"Integer",
							},
							Parameters: []ast.Node{},
						},
						Modifiers: []ast.Node{
							&ast.Modifier{
								Name: "public",
							},
						},
						Declarators: []ast.Node{
							&ast.VariableDeclarator{
								Name:       "field",
								Expression: nil,
							},
						},
					},
					&ast.FieldDeclaration{
						Type: &ast.Type{
							Name: []string{
								"Double",
							},
							Parameters: []ast.Node{},
						},
						Modifiers: []ast.Node{
							&ast.Modifier{
								Name: "public",
							},
						},
						Declarators: []ast.Node{
							&ast.VariableDeclarator{
								Name: "field_with_init",
								Expression: &ast.IntegerLiteral{
									Value: 2,
								},
							},
						},
					},
					&ast.FieldDeclaration{
						Type: &ast.Type{
							Name: []string{
								"String",
							},
							Parameters: []ast.Node{},
						},
						Modifiers: []ast.Node{
							&ast.Modifier{
								Name: "public",
							},
							&ast.Modifier{
								Name: "static",
							},
						},
						Declarators: []ast.Node{
							&ast.VariableDeclarator{
								Name:       "static_field",
								Expression: nil,
							},
						},
					},
					&ast.FieldDeclaration{
						Type: &ast.Type{
							Name: []string{
								"Boolean",
							},
							Parameters: []ast.Node{},
						},
						Modifiers: []ast.Node{
							&ast.Modifier{
								Name: "public",
							},
							&ast.Modifier{
								Name: "static",
							},
						},
						Declarators: []ast.Node{
							&ast.VariableDeclarator{
								Name: "static_field_with_init",
								Expression: &ast.IntegerLiteral{
									Value: 1,
								},
							},
						},
					},
					&ast.MethodDeclaration{
						Name: "static_method",
						Modifiers: []ast.Node{
							&ast.Modifier{
								Name: "public",
							},
							&ast.Modifier{
								Name: "static",
							},
						},
						ReturnType: &ast.Type{
							Name: []string{
								"String",
							},
							Parameters: []ast.Node{},
						},
						Parameters: []ast.Node{
							&ast.Parameter{
								Modifiers: []ast.Node{},
								Type: &ast.Type{
									Name: []string{
										"Boolean",
									},
									Parameters: []ast.Node{},
								},
								Name: "p1",
							},
						},
						Throws: []ast.Node{},
						Statements: &ast.Block{
							Statements: []ast.Node{},
						},
						NativeFunction: nil,
					},
					&ast.MethodDeclaration{
						Name: "method",
						Modifiers: []ast.Node{
							&ast.Modifier{
								Name: "public",
							},
						},
						ReturnType: ast.VoidType,
						Parameters: []ast.Node{},
						Throws:     []ast.Node{},
						Statements: &ast.Block{
							Statements: []ast.Node{},
						},
						NativeFunction: nil,
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
					},
					Declarators: []ast.Node{
						&ast.VariableDeclarator{
							Name: "i",
							Expression: &ast.IntegerLiteral{
								Value: 0,
							},
						},
					},
				},
				&ast.VariableDeclaration{
					Modifiers: []ast.Node{},
					Type: &ast.Type{
						Name: []string{
							"String",
						},
						Parameters: []ast.Node{},
					},
					Declarators: []ast.Node{
						&ast.VariableDeclarator{
							Name: "s",
							Expression: &ast.StringLiteral{
								Value: "abc",
							},
						},
					},
				},
				&ast.VariableDeclaration{
					Modifiers: []ast.Node{},
					Type: &ast.Type{
						Name: []string{
							"Double",
						},
						Parameters: []ast.Node{},
					},
					Declarators: []ast.Node{
						&ast.VariableDeclarator{
							Name: "d",
							Expression: &ast.DoubleLiteral{
								Value: 1.230000,
							},
						},
					},
				},
				&ast.VariableDeclaration{
					Modifiers: []ast.Node{},
					Type: &ast.Type{
						Name: []string{
							"Boolean",
						},
						Parameters: []ast.Node{},
					},
					Declarators: []ast.Node{
						&ast.VariableDeclarator{
							Name: "b",
							Expression: &ast.BooleanLiteral{
								Value: true,
							},
						},
					},
				},
				&ast.VariableDeclaration{
					Modifiers: []ast.Node{},
					Type: &ast.Type{
						Name: []string{
							"Integer",
						},
						Parameters: []ast.Node{},
					},
					Declarators: []ast.Node{
						&ast.VariableDeclarator{
							Name:       "i",
							Expression: nil,
						},
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
					},
					Right: &ast.BinaryOperator{
						Op: "+",
						Left: &ast.BinaryOperator{
							Op: "+",
							Left: &ast.IntegerLiteral{
								Value: 1,
							},
							Right: &ast.IntegerLiteral{
								Value: 2,
							},
						},
						Right: &ast.IntegerLiteral{
							Value: 3,
						},
					},
				},
				&ast.BinaryOperator{
					Op: "=",
					Left: &ast.Name{
						Value: "i",
					},
					Right: &ast.BinaryOperator{
						Op: "+",
						Left: &ast.IntegerLiteral{
							Value: 1,
						},
						Right: &ast.BinaryOperator{
							Op: "*",
							Left: &ast.IntegerLiteral{
								Value: 2,
							},
							Right: &ast.IntegerLiteral{
								Value: 3,
							},
						},
					},
				},
				&ast.BinaryOperator{
					Op: "=",
					Left: &ast.Name{
						Value: "i",
					},
					Right: &ast.BinaryOperator{
						Op: "*",
						Left: &ast.BinaryOperator{
							Op: "+",
							Left: &ast.IntegerLiteral{
								Value: 1,
							},
							Right: &ast.IntegerLiteral{
								Value: 2,
							},
						},
						Right: &ast.IntegerLiteral{
							Value: 3,
						},
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
						},
						FieldName: "bar",
					},
					FieldName: "baz",
				},
				&ast.MethodInvocation{
					NameOrExpression: &ast.FieldAccess{
						Expression: &ast.FieldAccess{
							Expression: &ast.Name{
								Value: "foo",
							},
							FieldName: "bar",
						},
						FieldName: "baz",
					},
					Parameters: []ast.Node{},
				},
				&ast.MethodInvocation{
					NameOrExpression: &ast.Name{
						Value: "foo",
					},
					Parameters: []ast.Node{
						&ast.IntegerLiteral{
							Value: 1,
						},
						&ast.Name{
							Value: "s",
						},
						&ast.BooleanLiteral{
							Value: false,
						},
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
						},
						Right: &ast.IntegerLiteral{
							Value: 1,
						},
					},
					IfStatement: &ast.Block{
						Statements: []ast.Node{
							&ast.BooleanLiteral{
								Value: true,
							},
						},
					},
					ElseStatement: &ast.If{
						Condition: &ast.BinaryOperator{
							Op: "==",
							Left: &ast.Name{
								Value: "i",
							},
							Right: &ast.IntegerLiteral{
								Value: 2,
							},
						},
						IfStatement: &ast.Block{
							Statements: []ast.Node{
								&ast.BooleanLiteral{
									Value: true,
								},
							},
						},
						ElseStatement: &ast.Block{
							Statements: []ast.Node{
								&ast.BooleanLiteral{
									Value: false,
								},
							},
						},
					},
				},
				&ast.Switch{
					Expression: &ast.Name{
						Value: "i",
					},
					WhenStatements: []ast.Node{
						&ast.When{
							Condition: []ast.Node{
								&ast.IntegerLiteral{
									Value: 1,
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
						&ast.When{
							Condition: []ast.Node{
								&ast.IntegerLiteral{
									Value: 2,
								},
								&ast.IntegerLiteral{
									Value: 3,
								},
							},
							Statements: &ast.Block{
								Statements: []ast.Node{
									&ast.BooleanLiteral{
										Value: false,
									},
								},
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
									},
									Identifier: "a",
								},
							},
							Statements: &ast.Block{
								Statements: []ast.Node{
									&ast.BooleanLiteral{
										Value: false,
									},
								},
							},
						},
					},
					ElseStatement: &ast.Block{
						Statements: []ast.Node{
							&ast.IntegerLiteral{
								Value: 1,
							},
						},
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
							},
							Declarators: []ast.Node{
								&ast.VariableDeclarator{
									Name: "i",
									Expression: &ast.IntegerLiteral{
										Value: 0,
									},
								},
							},
						},
						Expression: &ast.BinaryOperator{
							Op: "<",
							Left: &ast.Name{
								Value: "i",
							},
							Right: &ast.Name{
								Value: "imax",
							},
						},
						ForUpdate: []ast.Node{
							&ast.UnaryOperator{
								Op: "++",
								Expression: &ast.Name{
									Value: "i",
								},
								IsPrefix: false,
							},
						},
					},
					Statements: &ast.Block{
						Statements: []ast.Node{
							&ast.Continue{},
						},
					},
				},
				&ast.While{
					Condition: &ast.BooleanLiteral{
						Value: true,
					},
					Statements: []ast.Node{
						&ast.Block{
							Statements: []ast.Node{
								&ast.Break{},
							},
						},
					},
					IsDo: false,
				},
				&ast.While{
					Condition: &ast.BooleanLiteral{
						Value: false,
					},
					Statements: []ast.Node{
						&ast.Block{
							Statements: []ast.Node{
								&ast.Return{
									Expression: nil,
								},
							},
						},
					},
					IsDo: true,
				},
				&ast.For{
					Control: &ast.EnhancedForControl{
						Modifiers: []ast.Node{},
						Type: &ast.Type{
							Name: []string{
								"Account",
							},
							Parameters: []ast.Node{},
						},
						VariableDeclaratorId: "acc",
						Expression: &ast.Name{
							Value: "accounts",
						},
					},
					Statements: &ast.Block{
						Statements: []ast.Node{
							&ast.Return{
								Expression: &ast.IntegerLiteral{
									Value: 1,
								},
							},
						},
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
								},
							},
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
							},
							Identifier: "e",
							Block: &ast.Block{
								Statements: []ast.Node{
									&ast.Return{
										Expression: nil,
									},
								},
							},
						},
					},
					FinallyBlock: &ast.Block{
						Statements: []ast.Node{
							&ast.Return{
								Expression: nil,
							},
						},
					},
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
		actual := parseString(testCase.Code)

		equalNode(t, testCase.Expected, actual)
	}
}

func equalNode(t *testing.T, expected ast.Node, actual ast.Node) {
	e := ast.ToString(expected)
	a := ast.ToString(actual)
	if e != a {
		pp.Print(actual)
		diff := cmp.Diff(a, e)
		t.Errorf(diff)
	}
}

func createExpectedClass(statements []ast.Node) *ast.ClassDeclaration {
	return &ast.ClassDeclaration{
		Modifiers:   []ast.Node{},
		Annotations: []ast.Node{},
		Name:        "Foo",
		Declarations: []ast.Node{
			&ast.MethodDeclaration{
				Name: "action",
				Modifiers: []ast.Node{
					&ast.Modifier{
						Name: "public",
					},
				},
				ReturnType: ast.VoidType,
				Parameters: []ast.Node{},
				Throws:     []ast.Node{},
				Statements: &ast.Block{
					Statements: statements,
				},
				NativeFunction: nil,
			},
		},
	}
}
