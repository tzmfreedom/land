package ast

import (
	"testing"

	"github.com/google/go-cmp/cmp"
	"github.com/k0kubun/pp"
)

func TestParse(t *testing.T) {
	testCases := []struct {
		Code     string
		Expected Node
	}{
		{
			`@foo public without sharing class Foo extends Bar implements Baz {}`,
			&ClassDeclaration{
				Modifiers: []Node{
					&Modifier{
						Name: "public",
					},
					&Modifier{
						Name: "without sharing",
					},
				},
				Annotations: []Node{
					&Annotation{
						Name: "foo",
					},
				},
				Name: "Foo",
				SuperClass: &TypeRef{
					Name: []string{
						"Bar",
					},
					Parameters: []Node{},
				},
				ImplementClasses: []Node{
					&TypeRef{
						Name: []string{
							"Baz",
						},
						Parameters: []Node{},
					},
				},
				Declarations: []Node{},
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
			&ClassDeclaration{
				Modifiers:   []Node{},
				Annotations: []Node{},
				Name:        "Foo",
				Declarations: []Node{
					&FieldDeclaration{
						Type: &TypeRef{
							Name: []string{
								"Integer",
							},
							Parameters: []Node{},
						},
						Modifiers: []Node{
							&Modifier{
								Name: "public",
							},
						},
						Declarators: []Node{
							&VariableDeclarator{
								Name:       "field",
								Expression: nil,
							},
						},
					},
					&FieldDeclaration{
						Type: &TypeRef{
							Name: []string{
								"Double",
							},
							Parameters: []Node{},
						},
						Modifiers: []Node{
							&Modifier{
								Name: "public",
							},
						},
						Declarators: []Node{
							&VariableDeclarator{
								Name: "field_with_init",
								Expression: &IntegerLiteral{
									Value: 2,
								},
							},
						},
					},
					&FieldDeclaration{
						Type: &TypeRef{
							Name: []string{
								"String",
							},
							Parameters: []Node{},
						},
						Modifiers: []Node{
							&Modifier{
								Name: "public",
							},
							&Modifier{
								Name: "static",
							},
						},
						Declarators: []Node{
							&VariableDeclarator{
								Name:       "static_field",
								Expression: nil,
							},
						},
					},
					&FieldDeclaration{
						Type: &TypeRef{
							Name: []string{
								"Boolean",
							},
							Parameters: []Node{},
						},
						Modifiers: []Node{
							&Modifier{
								Name: "public",
							},
							&Modifier{
								Name: "static",
							},
						},
						Declarators: []Node{
							&VariableDeclarator{
								Name: "static_field_with_init",
								Expression: &IntegerLiteral{
									Value: 1,
								},
							},
						},
					},
					&MethodDeclaration{
						Name: "static_method",
						Modifiers: []Node{
							&Modifier{
								Name: "public",
							},
							&Modifier{
								Name: "static",
							},
						},
						ReturnType: &TypeRef{
							Name: []string{
								"String",
							},
							Parameters: []Node{},
						},
						Parameters: []Node{
							&Parameter{
								Modifiers: []Node{},
								Type: &TypeRef{
									Name: []string{
										"Boolean",
									},
									Parameters: []Node{},
								},
								Name: "p1",
							},
						},
						Throws: []Node{},
						Statements: &Block{
							Statements: []Node{},
						},
						NativeFunction: nil,
					},
					&MethodDeclaration{
						Name: "method",
						Modifiers: []Node{
							&Modifier{
								Name: "public",
							},
						},
						ReturnType: nil,
						Parameters: []Node{},
						Throws:     []Node{},
						Statements: &Block{
							Statements: []Node{},
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
			createExpectedClass([]Node{
				&VariableDeclaration{
					Modifiers: []Node{},
					Type: &TypeRef{
						Name: []string{
							"Integer",
						},
						Parameters: []Node{},
					},
					Declarators: []Node{
						&VariableDeclarator{
							Name: "i",
							Expression: &IntegerLiteral{
								Value: 0,
							},
						},
					},
				},
				&VariableDeclaration{
					Modifiers: []Node{},
					Type: &TypeRef{
						Name: []string{
							"String",
						},
						Parameters: []Node{},
					},
					Declarators: []Node{
						&VariableDeclarator{
							Name: "s",
							Expression: &StringLiteral{
								Value: "abc",
							},
						},
					},
				},
				&VariableDeclaration{
					Modifiers: []Node{},
					Type: &TypeRef{
						Name: []string{
							"Double",
						},
						Parameters: []Node{},
					},
					Declarators: []Node{
						&VariableDeclarator{
							Name: "d",
							Expression: &DoubleLiteral{
								Value: 1.230000,
							},
						},
					},
				},
				&VariableDeclaration{
					Modifiers: []Node{},
					Type: &TypeRef{
						Name: []string{
							"Boolean",
						},
						Parameters: []Node{},
					},
					Declarators: []Node{
						&VariableDeclarator{
							Name: "b",
							Expression: &BooleanLiteral{
								Value: true,
							},
						},
					},
				},
				&VariableDeclaration{
					Modifiers: []Node{},
					Type: &TypeRef{
						Name: []string{
							"Integer",
						},
						Parameters: []Node{},
					},
					Declarators: []Node{
						&VariableDeclarator{
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
			createExpectedClass([]Node{
				&BinaryOperator{
					Op: "=",
					Left: &Name{
						Value: []string{"i"},
					},
					Right: &BinaryOperator{
						Op: "+",
						Left: &BinaryOperator{
							Op: "+",
							Left: &IntegerLiteral{
								Value: 1,
							},
							Right: &IntegerLiteral{
								Value: 2,
							},
						},
						Right: &IntegerLiteral{
							Value: 3,
						},
					},
				},
				&BinaryOperator{
					Op: "=",
					Left: &Name{
						Value: []string{"i"},
					},
					Right: &BinaryOperator{
						Op: "+",
						Left: &IntegerLiteral{
							Value: 1,
						},
						Right: &BinaryOperator{
							Op: "*",
							Left: &IntegerLiteral{
								Value: 2,
							},
							Right: &IntegerLiteral{
								Value: 3,
							},
						},
					},
				},
				&BinaryOperator{
					Op: "=",
					Left: &Name{
						Value: []string{"i"},
					},
					Right: &BinaryOperator{
						Op: "*",
						Left: &BinaryOperator{
							Op: "+",
							Left: &IntegerLiteral{
								Value: 1,
							},
							Right: &IntegerLiteral{
								Value: 2,
							},
						},
						Right: &IntegerLiteral{
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
			createExpectedClass([]Node{
				&FieldAccess{
					Expression: &FieldAccess{
						Expression: &Name{
							Value: []string{"foo"},
						},
						FieldName: "bar",
					},
					FieldName: "baz",
				},
				&MethodInvocation{
					NameOrExpression: &FieldAccess{
						Expression: &FieldAccess{
							Expression: &Name{
								Value: []string{"foo"},
							},
							FieldName: "bar",
						},
						FieldName: "baz",
					},
					Parameters: []Node{},
				},
				&MethodInvocation{
					NameOrExpression: &Name{
						Value: []string{"foo"},
					},
					Parameters: []Node{
						&IntegerLiteral{
							Value: 1,
						},
						&Name{
							Value: []string{"s"},
						},
						&BooleanLiteral{
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
			createExpectedClass([]Node{
				&If{
					Condition: &BinaryOperator{
						Op: "==",
						Left: &Name{
							Value: []string{"i"},
						},
						Right: &IntegerLiteral{
							Value: 1,
						},
					},
					IfStatement: &Block{
						Statements: []Node{
							&BooleanLiteral{
								Value: true,
							},
						},
					},
					ElseStatement: &If{
						Condition: &BinaryOperator{
							Op: "==",
							Left: &Name{
								Value: []string{"i"},
							},
							Right: &IntegerLiteral{
								Value: 2,
							},
						},
						IfStatement: &Block{
							Statements: []Node{
								&BooleanLiteral{
									Value: true,
								},
							},
						},
						ElseStatement: &Block{
							Statements: []Node{
								&BooleanLiteral{
									Value: false,
								},
							},
						},
					},
				},
				&Switch{
					Expression: &Name{
						Value: []string{"i"},
					},
					WhenStatements: []Node{
						&When{
							Condition: []Node{
								&IntegerLiteral{
									Value: 1,
								},
							},
							Statements: &Block{
								Statements: []Node{
									&BooleanLiteral{
										Value: true,
									},
								},
							},
						},
						&When{
							Condition: []Node{
								&IntegerLiteral{
									Value: 2,
								},
								&IntegerLiteral{
									Value: 3,
								},
							},
							Statements: &Block{
								Statements: []Node{
									&BooleanLiteral{
										Value: false,
									},
								},
							},
						},
						&When{
							Condition: []Node{
								&WhenType{
									Type: &TypeRef{
										Name: []string{
											"Account",
										},
										Parameters: []Node{},
									},
									Identifier: "a",
								},
							},
							Statements: &Block{
								Statements: []Node{
									&BooleanLiteral{
										Value: false,
									},
								},
							},
						},
					},
					ElseStatement: &Block{
						Statements: []Node{
							&IntegerLiteral{
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
			createExpectedClass([]Node{
				&For{
					Control: &ForControl{
						ForInit: &VariableDeclaration{
							Modifiers: []Node{},
							Type: &TypeRef{
								Name: []string{
									"Integer",
								},
								Parameters: []Node{},
							},
							Declarators: []Node{
								&VariableDeclarator{
									Name: "i",
									Expression: &IntegerLiteral{
										Value: 0,
									},
								},
							},
						},
						Expression: &BinaryOperator{
							Op: "<",
							Left: &Name{
								Value: []string{"i"},
							},
							Right: &Name{
								Value: []string{"imax"},
							},
						},
						ForUpdate: []Node{
							&UnaryOperator{
								Op: "++",
								Expression: &Name{
									Value: []string{"i"},
								},
								IsPrefix: false,
							},
						},
					},
					Statements: &Block{
						Statements: []Node{
							&Continue{},
						},
					},
				},
				&While{
					Condition: &BooleanLiteral{
						Value: true,
					},
					Statements: &Block{
						Statements: []Node{
							&Break{},
						},
					},
					IsDo: false,
				},
				&While{
					Condition: &BooleanLiteral{
						Value: false,
					},
					Statements: &Block{
						Statements: []Node{
							&Return{
								Expression: nil,
							},
						},
					},
					IsDo: true,
				},
				&For{
					Control: &EnhancedForControl{
						Modifiers: []Node{},
						Type: &TypeRef{
							Name: []string{
								"Account",
							},
							Parameters: []Node{},
						},
						VariableDeclaratorId: "acc",
						Expression: &Name{
							Value: []string{"accounts"},
						},
					},
					Statements: &Block{
						Statements: []Node{
							&Return{
								Expression: &IntegerLiteral{
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
			createExpectedClass([]Node{
				&Try{
					Block: &Block{
						Statements: []Node{
							&Throw{
								Expression: &Name{
									Value: []string{"a"},
								},
							},
						},
					},
					CatchClause: []Node{
						&Catch{
							Modifiers: []Node{},
							Type: &TypeRef{
								Name: []string{
									"Exception",
								},
								Parameters: nil,
							},
							Identifier: "e",
							Block: &Block{
								Statements: []Node{
									&Return{
										Expression: nil,
									},
								},
							},
						},
					},
					FinallyBlock: &Block{
						Statements: []Node{
							&Return{
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
			&Trigger{
				Name:   "Foo",
				Object: "Account",
				TriggerTimings: []Node{
					&TriggerTiming{
						Dml:    "insert",
						Timing: "before",
					},
					&TriggerTiming{
						Dml:    "update",
						Timing: "after",
					},
				},
				Statements: &Block{
					Statements: []Node{
						&BooleanLiteral{
							Value: true,
						},
					},
				},
			},
		},
	}
	for _, testCase := range testCases {
		actual, err := ParseString(testCase.Code)
		if err != nil {
			panic(err)
		}

		equalNode(t, testCase.Expected, actual)
	}
}

func equalNode(t *testing.T, expected Node, actual Node) {
	e := ToString(expected)
	a := ToString(actual)
	if e != a {
		pp.Print(actual)
		diff := cmp.Diff(a, e)
		t.Errorf(diff)
	}
}

func createExpectedClass(statements []Node) *ClassDeclaration {
	return &ClassDeclaration{
		Modifiers:   []Node{},
		Annotations: []Node{},
		Name:        "Foo",
		Declarations: []Node{
			&MethodDeclaration{
				Name: "action",
				Modifiers: []Node{
					&Modifier{
						Name: "public",
					},
				},
				ReturnType: nil,
				Parameters: []Node{},
				Throws:     []Node{},
				Statements: &Block{
					Statements: statements,
				},
				NativeFunction: nil,
			},
		},
	}
}
