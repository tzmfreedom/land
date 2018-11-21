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
		// Array Access
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
		// If, Else
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
		// While
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
		// Ternaly
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
		// Return Type
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
		// For
		{
			&ClassType{
				InstanceMethods: &MethodMap{
					Data: map[string][]ast.Node{
						"foo": {
							&ast.MethodDeclaration{
								Statements: &ast.Block{
									Statements: []ast.Node{
										func() *ast.For {
											t := &ast.For{
												Control: &ast.ForControl{
													Expression: &ast.IntegerLiteral{},
												},
											}
											b := &ast.Block{
												Statements: []ast.Node{},
												Parent:     t,
											}
											b.Statements = append(b.Statements, &ast.Break{Parent: b})
											b.Statements = append(b.Statements, &ast.Continue{Parent: b})
											t.Statements = b
											return t
										}(),
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
										&ast.Break{},
										&ast.Continue{},
									},
								},
							},
						},
					},
				},
			},
			[]*Error{
				{
					Message: "break must be in for/while loop",
				},
				{
					Message: "condition <Double> must be Boolean expression",
				},
				{
					Message: "condition <Integer> must be Boolean expression",
				},
				{
					Message: "condition <String> must be Boolean expression",
				},
				{
					Message: "continue must be in for/while loop",
				},
			},
		},
		// Unary Operator
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
		// Variable Declaration, Variable Assignment
		{
			&ClassType{
				InstanceMethods: &MethodMap{
					Data: map[string][]ast.Node{
						"foo": {
							&ast.MethodDeclaration{
								Statements: &ast.Block{
									Statements: []ast.Node{
										&ast.VariableDeclaration{
											Type: &ast.TypeRef{Name: []string{"Integer"}},
											Declarators: []ast.Node{
												&ast.VariableDeclarator{
													Name:       "i",
													Expression: &ast.IntegerLiteral{},
												},
											},
										},
										&ast.VariableDeclaration{
											Type: &ast.TypeRef{Name: []string{"String"}},
											Declarators: []ast.Node{
												&ast.VariableDeclarator{
													Name:       "j",
													Expression: &ast.IntegerLiteral{},
												},
											},
										},
										&ast.BinaryOperator{
											Op:    "=",
											Left:  &ast.Name{Value: []string{"i"}},
											Right: &ast.IntegerLiteral{},
										},
										&ast.BinaryOperator{
											Op:    "=",
											Left:  &ast.Name{Value: []string{"i"}},
											Right: &ast.StringLiteral{},
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
					Message: "expression <Integer> does not match <String>",
				},
				{
					Message: "expression <String> does not match <Integer>",
				},
			},
		},
		//
		{
			&ClassType{
				InstanceMethods: &MethodMap{
					Data: map[string][]ast.Node{
						"foo": {
							&ast.MethodDeclaration{
								Statements: &ast.Block{
									Statements: []ast.Node{
										&ast.BinaryOperator{
											Op:    "+",
											Left:  &ast.IntegerLiteral{},
											Right: &ast.IntegerLiteral{},
										},
										&ast.BinaryOperator{
											Op:    "-",
											Left:  &ast.IntegerLiteral{},
											Right: &ast.IntegerLiteral{},
										},
										&ast.BinaryOperator{
											Op:    "*",
											Left:  &ast.IntegerLiteral{},
											Right: &ast.IntegerLiteral{},
										},
										&ast.BinaryOperator{
											Op:    "/",
											Left:  &ast.IntegerLiteral{},
											Right: &ast.IntegerLiteral{},
										},
										&ast.BinaryOperator{
											Op:    "%",
											Left:  &ast.IntegerLiteral{},
											Right: &ast.IntegerLiteral{},
										},
										// failure
										&ast.BinaryOperator{
											Op:    "+",
											Left:  &ast.StringLiteral{},
											Right: &ast.IntegerLiteral{},
										},
										// failure
										&ast.BinaryOperator{
											Op:    "-",
											Left:  &ast.StringLiteral{},
											Right: &ast.IntegerLiteral{},
										},
										// failure
										&ast.BinaryOperator{
											Op:    "*",
											Left:  &ast.StringLiteral{},
											Right: &ast.IntegerLiteral{},
										},
										// failure
										&ast.BinaryOperator{
											Op:    "/",
											Left:  &ast.StringLiteral{},
											Right: &ast.IntegerLiteral{},
										},
										// failure
										&ast.BinaryOperator{
											Op:    "%",
											Left:  &ast.StringLiteral{},
											Right: &ast.IntegerLiteral{},
										},
										&ast.BinaryOperator{
											Op:    "+",
											Left:  &ast.StringLiteral{},
											Right: &ast.StringLiteral{},
										},
										// failure
										&ast.BinaryOperator{
											Op:    "+",
											Left:  &ast.BooleanLiteral{},
											Right: &ast.BooleanLiteral{},
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
					Message: "expression <Boolean> must be Integer, String or Double",
				},
				{
					Message: "expression <String> does not match <Integer>",
				},
				{
					Message: "expression <String> must be Integer or Double",
				},
				{
					Message: "expression <String> must be Integer or Double",
				},
				{
					Message: "expression <String> must be Integer or Double",
				},
				{
					Message: "expression <String> must be Integer or Double",
				},
			},
		},
	}
	for _, testCase := range testCases {
		checker := NewTypeChecker()
		checker.Context = &Context{}
		checker.Context.ClassTypes = &ClassMap{
			Data: map[string]*ClassType{
				"integer": IntegerType,
				"string":  StringType,
			},
		}
		checker.VisitClassType(testCase.Input)

		messages := make([]string, len(checker.Errors))
		for i, err := range checker.Errors {
			messages[i] = err.Message
		}
		sort.Slice(messages, func(i, j int) bool {
			return messages[i] < messages[j]
		})

		if len(messages) != len(testCase.ExpectErrors) {
			t.Errorf("error size is not match: %d != %d", len(messages), len(testCase.ExpectErrors))
		} else {
			for i, message := range messages {
				expected := testCase.ExpectErrors[i]
				if expected.Message != message {
					t.Errorf("expected: %s, actual: %s", expected.Message, message)
				}
			}
		}
	}
}
