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
				Modifiers: []*Modifier{
					{
						Name: "public",
					},
					{
						Name: "without sharing",
					},
				},
				Annotations: []*Annotation{
					{
						Name: "foo",
					},
				},
				Name: "Foo",
				SuperClassRef: &TypeRef{
					Name: []string{
						"Bar",
					},
					Parameters: []*TypeRef{},
				},
				ImplementClassRefs: []*TypeRef{
					{
						Name: []string{
							"Baz",
						},
						Parameters: []*TypeRef{},
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
				Modifiers:   []*Modifier{},
				Annotations: []*Annotation{},
				Name:        "Foo",
				Declarations: []Node{
					&FieldDeclaration{
						TypeRef: &TypeRef{
							Name: []string{
								"Integer",
							},
							Parameters: []*TypeRef{},
						},
						Modifiers: []*Modifier{
							{
								Name: "public",
							},
						},
						Declarators: []*VariableDeclarator{
							{
								Name:       "field",
								Expression: nil,
							},
						},
					},
					&FieldDeclaration{
						TypeRef: &TypeRef{
							Name: []string{
								"Double",
							},
							Parameters: []*TypeRef{},
						},
						Modifiers: []*Modifier{
							{
								Name: "public",
							},
						},
						Declarators: []*VariableDeclarator{
							{
								Name: "field_with_init",
								Expression: &IntegerLiteral{
									Value: 2,
								},
							},
						},
					},
					&FieldDeclaration{
						TypeRef: &TypeRef{
							Name: []string{
								"String",
							},
							Parameters: []*TypeRef{},
						},
						Modifiers: []*Modifier{
							{
								Name: "public",
							},
							{
								Name: "static",
							},
						},
						Declarators: []*VariableDeclarator{
							{
								Name:       "static_field",
								Expression: nil,
							},
						},
					},
					&FieldDeclaration{
						TypeRef: &TypeRef{
							Name: []string{
								"Boolean",
							},
							Parameters: []*TypeRef{},
						},
						Modifiers: []*Modifier{
							{
								Name: "public",
							},
							{
								Name: "static",
							},
						},
						Declarators: []*VariableDeclarator{
							{
								Name: "static_field_with_init",
								Expression: &IntegerLiteral{
									Value: 1,
								},
							},
						},
					},
					&MethodDeclaration{
						Name: "static_method",
						Modifiers: []*Modifier{
							{
								Name: "public",
							},
							{
								Name: "static",
							},
						},
						ReturnType: &TypeRef{
							Name: []string{
								"String",
							},
							Parameters: []*TypeRef{},
						},
						Parameters: []*Parameter{
							{
								Modifiers: []*Modifier{},
								TypeRef: &TypeRef{
									Name: []string{
										"Boolean",
									},
									Parameters: []*TypeRef{},
								},
								Name: "p1",
							},
						},
						Throws: []Node{},
						Statements: &Block{
							Statements: []Node{},
						},
					},
					&MethodDeclaration{
						Name: "method",
						Modifiers: []*Modifier{
							&Modifier{
								Name: "public",
							},
						},
						ReturnType: nil,
						Parameters: []*Parameter{},
						Throws:     []Node{},
						Statements: &Block{
							Statements: []Node{},
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
			createExpectedClass([]Node{
				&VariableDeclaration{
					Modifiers: []*Modifier{},
					TypeRef: &TypeRef{
						Name: []string{
							"Integer",
						},
						Parameters: []*TypeRef{},
					},
					Declarators: []*VariableDeclarator{
						{
							Name: "i",
							Expression: &IntegerLiteral{
								Value: 0,
							},
						},
					},
				},
				&VariableDeclaration{
					Modifiers: []*Modifier{},
					TypeRef: &TypeRef{
						Name: []string{
							"String",
						},
						Parameters: []*TypeRef{},
					},
					Declarators: []*VariableDeclarator{
						{
							Name: "s",
							Expression: &StringLiteral{
								Value: "abc",
							},
						},
					},
				},
				&VariableDeclaration{
					Modifiers: []*Modifier{},
					TypeRef: &TypeRef{
						Name: []string{
							"Double",
						},
						Parameters: []*TypeRef{},
					},
					Declarators: []*VariableDeclarator{
						{
							Name: "d",
							Expression: &DoubleLiteral{
								Value: 1.230000,
							},
						},
					},
				},
				&VariableDeclaration{
					Modifiers: []*Modifier{},
					TypeRef: &TypeRef{
						Name: []string{
							"Boolean",
						},
						Parameters: []*TypeRef{},
					},
					Declarators: []*VariableDeclarator{
						{
							Name: "b",
							Expression: &BooleanLiteral{
								Value: true,
							},
						},
					},
				},
				&VariableDeclaration{
					Modifiers: []*Modifier{},
					TypeRef: &TypeRef{
						Name: []string{
							"Integer",
						},
						Parameters: []*TypeRef{},
					},
					Declarators: []*VariableDeclarator{
						{
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
    when else {
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
					WhenStatements: []*When{
						{
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
						{
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
						{
							Condition: []Node{
								&WhenType{
									TypeRef: &TypeRef{
										Name: []string{
											"Account",
										},
										Parameters: []*TypeRef{},
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
						ForInit: []Node{
							&VariableDeclaration{
								Modifiers: []*Modifier{},
								TypeRef: &TypeRef{
									Name: []string{
										"Integer",
									},
									Parameters: []*TypeRef{},
								},
								Declarators: []*VariableDeclarator{
									{
										Name: "i",
										Expression: &IntegerLiteral{
											Value: 0,
										},
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
						Modifiers: []*Modifier{},
						TypeRef: &TypeRef{
							Name: []string{
								"Account",
							},
							Parameters: []*TypeRef{},
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
					CatchClause: []*Catch{
						{
							Modifiers: []*Modifier{},
							TypeRef: &TypeRef{
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
			`class Foo {
public void action(){
[
  SELECT
    id,
    Name,
    count()
  FROM
    Account
  WHERE
    Name = ''
    AND A__c = 1
    AND B__c = true
    OR  C__c = :foo
    OR  D__c = :bar()
  GROUP BY
    E__c,
    F__c
  ORDER BY
    id,
    Name ASC
  OFFSET
    1000
];
}
}`,
			createExpectedClass([]Node{
				&Soql{
					SelectFields: []Node{
						&SelectField{
							Value: []string{"id"},
						},
						&SelectField{
							Value: []string{"Name"},
						},
						&SoqlFunction{
							Name: "count",
						},
					},
					FromObject: "Account",
					Where: &WhereBinaryOperator{
						Left: &WhereBinaryOperator{
							Left: &WhereBinaryOperator{
								Left: &WhereBinaryOperator{
									Left: &WhereCondition{
										Field: &SelectField{
											Value: []string{"Name"},
										},
										Op: "=",
										Expression: &StringLiteral{
											Value: "",
										},
										Not: false,
									},
									Right: &WhereCondition{
										Field: &SelectField{
											Value: []string{"A__c"},
										},
										Op: "=",
										Expression: &IntegerLiteral{
											Value: 1,
										},
										Not: false,
									},
									Op: "AND",
								},
								Right: &WhereCondition{
									Field: &SelectField{
										Value: []string{"B__c"},
									},
									Op: "=",
									Expression: &BooleanLiteral{
										Value: true,
									},
									Not: false,
								},
								Op: "AND",
							},
							Right: &WhereCondition{
								Field: &SelectField{
									Value: []string{"C__c"},
								},
								Op: "=",
								Expression: &Name{
									Value: []string{"foo"},
								},
								Not: false,
							},
							Op: "OR",
						},
						Right: &WhereCondition{
							Field: &SelectField{
								Value: []string{"D__c"},
							},
							Op: "=",
							Expression: &MethodInvocation{
								NameOrExpression: &Name{
									Value: []string{"bar"},
								},
							},
							Not: false,
						},
						Op: "OR",
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
		Modifiers:   []*Modifier{},
		Annotations: []*Annotation{},
		Name:        "Foo",
		Declarations: []Node{
			&MethodDeclaration{
				Name: "action",
				Modifiers: []*Modifier{
					{
						Name: "public",
					},
				},
				ReturnType: nil,
				Parameters: []*Parameter{},
				Throws:     []Node{},
				Statements: &Block{
					Statements: statements,
				},
			},
		},
	}
}
