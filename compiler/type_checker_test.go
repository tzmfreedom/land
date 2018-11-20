package compiler

import (
	"testing"

	"github.com/tzmfreedom/goland/ast"
)

func TestTypeChecker(t *testing.T) {
	testCases := []struct {
		Input        *ast.ClassType
		ExpectErrors []*Error
	}{
		{
			&ast.ClassType{
				InstanceMethods: &ast.MethodMap{
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
					Message: "array key <boolean> must be integer or string",
				},
			},
		},
		{
			&ast.ClassType{
				InstanceMethods: &ast.MethodMap{
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
					Message: "condition <string> must be boolean expression",
				},
				{
					Message: "condition <integer> must be boolean expression",
				},
				{
					Message: "condition <double> must be boolean expression",
				},
			},
		},
		{
			&ast.ClassType{
				InstanceMethods: &ast.MethodMap{
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
					Message: "condition <string> must be boolean expression",
				},
				{
					Message: "condition <integer> must be boolean expression",
				},
				{
					Message: "condition <double> must be boolean expression",
				},
			},
		},
		{
			&ast.ClassType{
				InstanceMethods: &ast.MethodMap{
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
					Message: "condition <string> must be boolean expression",
				},
				{
					Message: "condition <integer> must be boolean expression",
				},
				{
					Message: "condition <double> must be boolean expression",
				},
			},
		},
		{
			&ast.ClassType{
				InstanceMethods: &ast.MethodMap{
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
					Message: "return type <string> does not match integer",
				},
				{
					Message: "return type <double> does not match integer",
				},
				{
					Message: "return type <boolean> does not match integer",
				},
			},
		},
		{
			&ast.ClassType{
				InstanceMethods: &ast.MethodMap{
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
					Message: "condition <integer> must be boolean expression",
				},
				{
					Message: "condition <string> must be boolean expression",
				},
				{
					Message: "condition <double> must be boolean expression",
				},
			},
		},
		{
			&ast.ClassType{
				InstanceMethods: &ast.MethodMap{
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
					Message: "expression <string> must be integer",
				},
				{
					Message: "expression <double> must be integer",
				},
				{
					Message: "expression <boolean> must be integer",
				},
			},
		},
	}
	for _, testCase := range testCases {
		checker := NewTypeChecker()
		checker.VisitClassType(testCase.Input)

		for i, actual := range checker.Errors {
			expected := testCase.ExpectErrors[i]
			if expected.Message != actual.Message {
				t.Errorf("expected: %s, actual: %s", expected.Message, actual.Message)
			}
		}
	}
}
