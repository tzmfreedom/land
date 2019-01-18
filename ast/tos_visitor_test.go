package ast

import (
	"testing"
)

func TestToString(t *testing.T) {
	testCases := []struct {
		Input    Node
		Expected string
	}{
		{
			&StringLiteral{Value: "foo"},
			"'foo'",
		},
		{
			&IntegerLiteral{Value: 1},
			"1",
		},
		{
			&DoubleLiteral{Value: 1.23},
			"1.230000",
		},
		{
			&BooleanLiteral{Value: true},
			"true",
		},
		{
			&BooleanLiteral{Value: false},
			"false",
		},
		{
			&VariableDeclaration{
				TypeRef: &TypeRef{
					Name: []string{"String"},
				},
				Declarators: []*VariableDeclarator{
					{
						Name:       "s",
						Expression: nil,
					},
				},
			},
			"String s",
		},
		{
			&VariableDeclaration{
				TypeRef: &TypeRef{
					Name: []string{"Integer"},
				},
				Declarators: []*VariableDeclarator{
					{
						Name:       "i",
						Expression: &IntegerLiteral{Value: 1},
					},
				},
			},
			"Integer i = 1",
		},
		{
			&BinaryOperator{
				Left:  &IntegerLiteral{Value: 1},
				Right: &IntegerLiteral{Value: 2},
				Op:    "+",
			},
			"1 + 2",
		},
		{
			&ArrayAccess{
				Receiver: &Name{Value: []string{"foo"}},
				Key:      &StringLiteral{Value: "bar"},
			},
			"foo['bar']",
		},
		{
			&ArrayAccess{
				Receiver: &Name{Value: []string{"foo"}},
				Key:      &IntegerLiteral{Value: 1},
			},
			"foo[1]",
		},
		{
			&Break{},
			"break",
		},
		{
			&Continue{},
			"continue",
		},
		{
			&Return{
				Expression: &StringLiteral{
					Value: "foo",
				},
			},
			"return 'foo'",
		},
		{
			&Dml{
				Type:       "insert",
				Expression: &Name{Value: []string{"foo"}},
			},
			"insert foo",
		},
		{
			&ClassDeclaration{
				Name:          "foo",
				SuperClassRef: &TypeRef{Name: []string{"bar"}},
				ImplementClassRefs: []*TypeRef{
					{Name: []string{"baz"}},
				},
				Modifiers: []*Modifier{
					{Name: "public"},
				},
				Annotations: []*Annotation{
					{Name: "@annotation"},
				},
			},
			`@annotation
public class foo extends bar implements baz {
}`,
		},
		{
			&MethodDeclaration{
				Name:       "foo",
				ReturnType: &TypeRef{Name: []string{"Integer"}},
				Modifiers: []*Modifier{
					{Name: "public"},
				},
				Annotations: []*Annotation{
					{Name: "@annotation"},
				},
				Parameters: []*Parameter{
					{
						TypeRef: &TypeRef{Name: []string{"String"}},
						Name:    "s",
					},
				},
				Statements: &Block{
					Statements: []Node{},
				},
			},
			`@annotation
public Integer foo (String s) {
}`,
		},
	}

	for _, testCase := range testCases {
		actual := ToString(testCase.Input)

		if testCase.Expected != actual {
			t.Errorf("expected %s, actual %s", testCase.Expected, actual)
		}
	}
}
