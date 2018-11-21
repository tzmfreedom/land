package compiler

import (
	"sort"
	"testing"

	"github.com/tzmfreedom/goland/ast"
)

func TestTypeChecker(t *testing.T) {
	testCases := []struct {
		Input        *ClassType
		ExpectErrors []*Error
	}{
		{
			&ClassType{
				InstanceMethods: &MethodMap{
					Data: map[string][]ast.Node{
						"foo": {
							&ast.MethodDeclaration{
								Statements: &ast.Block{
									Statements: []ast.Node{
										&ast.ArrayAccess{
											Receiver: &ast.IntegerLiteral{},
											Key:      &ast.IntegerLiteral{},
										},
										&ast.ArrayAccess{
											Receiver: &ast.IntegerLiteral{},
											Key:      &ast.StringLiteral{},
										},
										&ast.ArrayAccess{
											Receiver: &ast.IntegerLiteral{},
											Key:      &ast.BooleanLiteral{},
										},
									},
								},
							},
						},
					},
				},
			},
			[]*Error{
				{
					Message: "array key <Boolean> must be Integer or string",
				},
			},
		},
		{
			&ClassType{
				InstanceMethods: &MethodMap{
					Data: map[string][]ast.Node{
						"foo": {
							&ast.MethodDeclaration{
								Statements: &ast.Block{
									Statements: []ast.Node{
										&ast.If{
											Condition:     &ast.BooleanLiteral{},
											IfStatement:   &ast.Block{},
											ElseStatement: &ast.Block{},
										},
										&ast.If{
											Condition:     &ast.StringLiteral{},
											IfStatement:   &ast.Block{},
											ElseStatement: &ast.Block{},
										},
										&ast.If{
											Condition:     &ast.IntegerLiteral{},
											IfStatement:   &ast.Block{},
											ElseStatement: &ast.Block{},
										},
										&ast.If{
											Condition:     &ast.DoubleLiteral{},
											IfStatement:   &ast.Block{},
											ElseStatement: &ast.Block{},
										},
									},
								},
							},
						},
					},
				},
			},
			[]*Error{
				{
					Message: "condition <Double> must be Boolean expression",
				},
				{
					Message: "condition <Integer> must be Boolean expression",
				},
				{
					Message: "condition <String> must be Boolean expression",
				},
			},
		},
		{
			&ClassType{
				InstanceMethods: &MethodMap{
					Data: map[string][]ast.Node{
						"foo": {
							&ast.MethodDeclaration{
								Statements: &ast.Block{
									Statements: []ast.Node{
										&ast.While{
											Condition:  &ast.BooleanLiteral{},
											Statements: []ast.Node{},
										},
										&ast.While{
											Condition:  &ast.StringLiteral{},
											Statements: []ast.Node{},
										},
										&ast.While{
											Condition:  &ast.IntegerLiteral{},
											Statements: []ast.Node{},
										},
										&ast.While{
											Condition:  &ast.DoubleLiteral{},
											Statements: []ast.Node{},
										},
									},
								},
							},
						},
					},
				},
			},
			[]*Error{
				{
					Message: "condition <Double> must be Boolean expression",
				},
				{
					Message: "condition <Integer> must be Boolean expression",
				},
				{
					Message: "condition <String> must be Boolean expression",
				},
			},
		},
		{
			&ClassType{
				InstanceMethods: &MethodMap{
					Data: map[string][]ast.Node{
						"foo": {
							&ast.MethodDeclaration{
								Statements: &ast.Block{
									Statements: []ast.Node{
										&ast.TernalyExpression{
											Condition:       &ast.BooleanLiteral{},
											TrueExpression:  &ast.IntegerLiteral{},
											FalseExpression: &ast.IntegerLiteral{},
										},
										&ast.TernalyExpression{
											Condition:       &ast.StringLiteral{},
											TrueExpression:  &ast.IntegerLiteral{},
											FalseExpression: &ast.IntegerLiteral{},
										},
										&ast.TernalyExpression{
											Condition:       &ast.IntegerLiteral{},
											TrueExpression:  &ast.IntegerLiteral{},
											FalseExpression: &ast.IntegerLiteral{},
										},
										&ast.TernalyExpression{
											Condition:       &ast.DoubleLiteral{},
											TrueExpression:  &ast.IntegerLiteral{},
											FalseExpression: &ast.IntegerLiteral{},
										},
									},
								},
							},
						},
					},
				},
			},
			[]*Error{
				{
					Message: "condition <Double> must be Boolean expression",
				},
				{
					Message: "condition <Integer> must be Boolean expression",
				},
				{
					Message: "condition <String> must be Boolean expression",
				},
			},
		},
		{
			&ClassType{
				InstanceMethods: &MethodMap{
					Data: map[string][]ast.Node{
						"foo": {
							&ast.MethodDeclaration{
								ReturnType: &ast.TypeRef{Name: []string{"Integer"}},
								Statements: &ast.Block{
									Statements: []ast.Node{
										&ast.Return{
											Expression: &ast.StringLiteral{},
										},
									},
								},
							},
						},
						"bar": {
							&ast.MethodDeclaration{
								ReturnType: &ast.TypeRef{Name: []string{"Integer"}},
								Statements: &ast.Block{
									Statements: []ast.Node{
										&ast.Return{
											Expression: &ast.IntegerLiteral{},
										},
									},
								},
							},
						},
						"baz": {
							&ast.MethodDeclaration{
								ReturnType: &ast.TypeRef{Name: []string{"Integer"}},
								Statements: &ast.Block{
									Statements: []ast.Node{
										&ast.Return{
											Expression: &ast.DoubleLiteral{},
										},
									},
								},
							},
						},
						"qux": {
							&ast.MethodDeclaration{
								ReturnType: &ast.TypeRef{Name: []string{"Integer"}},
								Statements: &ast.Block{
									Statements: []ast.Node{
										&ast.Return{
											Expression: &ast.BooleanLiteral{},
										},
									},
								},
							},
						},
					},
				},
			},
			[]*Error{
				{
					Message: "return type <Boolean> does not match Integer",
				},
				{
					Message: "return type <Double> does not match Integer",
				},
				{
					Message: "return type <String> does not match Integer",
				},
			},
		},
		{
			&ClassType{
				InstanceMethods: &MethodMap{
					Data: map[string][]ast.Node{
						"foo": {
							&ast.MethodDeclaration{
								Statements: &ast.Block{
									Statements: []ast.Node{
										&ast.For{
											Control: &ast.ForControl{
												Expression: &ast.IntegerLiteral{},
											},
										},
										&ast.For{
											Control: &ast.ForControl{
												Expression: &ast.StringLiteral{},
											},
										},
										&ast.For{
											Control: &ast.ForControl{
												Expression: &ast.DoubleLiteral{},
											},
										},
										&ast.For{
											Control: &ast.ForControl{
												Expression: &ast.BooleanLiteral{},
											},
										},
									},
								},
							},
						},
					},
				},
			},
			[]*Error{
				{
					Message: "condition <Double> must be Boolean expression",
				},
				{
					Message: "condition <Integer> must be Boolean expression",
				},
				{
					Message: "condition <String> must be Boolean expression",
				},
			},
		},
		{
			&ClassType{
				InstanceMethods: &MethodMap{
					Data: map[string][]ast.Node{
						"foo": {
							&ast.MethodDeclaration{
								Statements: &ast.Block{
									Statements: []ast.Node{
										&ast.UnaryOperator{
											Expression: &ast.IntegerLiteral{},
										},
										&ast.UnaryOperator{
											Expression: &ast.StringLiteral{},
										},
										&ast.UnaryOperator{
											Expression: &ast.DoubleLiteral{},
										},
										&ast.UnaryOperator{
											Expression: &ast.BooleanLiteral{},
										},
									},
								},
							},
						},
					},
				},
			},
			[]*Error{
				{
					Message: "expression <Boolean> must be Integer",
				},
				{
					Message: "expression <Double> must be Integer",
				},
				{
					Message: "expression <String> must be Integer",
				},
			},
		},
	}
	for _, testCase := range testCases {
		checker := NewTypeChecker()
		checker.VisitClassType(testCase.Input)

		messages := make([]string, len(checker.Errors))
		for i, err := range checker.Errors {
			messages[i] = err.Message
		}
		sort.Slice(messages, func(i, j int) bool {
			return messages[i] < messages[j]
		})

		for i, message := range messages {
			expected := testCase.ExpectErrors[i]
			if expected.Message != message {
				t.Errorf("expected: %s, actual: %s", expected.Message, message)
			}
		}
	}
}
