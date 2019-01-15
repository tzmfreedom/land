package compiler

import (
	"errors"
	"sort"
	"testing"

	"strings"

	"github.com/tzmfreedom/goland/ast"
	"github.com/tzmfreedom/goland/builtin"
)

func TestTypeChecker(t *testing.T) {
	testCases := []struct {
		Input        *builtin.ClassType
		ExpectErrors []*Error
	}{
		// Array key must be integer or string
		{
			newTestClassType([]ast.Node{
				&ast.ArrayAccess{
					Receiver: &ast.TypeRef{
						Name: []string{"List"},
						Parameters: []ast.Node{
							&ast.TypeRef{
								Name: []string{"Integer"},
							},
						},
					},
					Key: &ast.IntegerLiteral{},
				},
				&ast.ArrayAccess{
					Receiver: &ast.TypeRef{
						Name: []string{"List"},
						Parameters: []ast.Node{
							&ast.TypeRef{
								Name: []string{"Integer"},
							},
						},
					},
					Key: &ast.StringLiteral{},
				},
				&ast.ArrayAccess{
					Receiver: &ast.TypeRef{
						Name: []string{"List"},
						Parameters: []ast.Node{
							&ast.TypeRef{
								Name: []string{"Integer"},
							},
						},
					},
					Key: &ast.BooleanLiteral{},
				},
			}),
			[]*Error{
				{
					Message: "list key <Boolean> must be Integer",
				},
				{
					Message: "list key <String> must be Integer",
				},
			},
		},
		// `if` condition type must be boolean
		{
			newTestClassType([]ast.Node{
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
			}),
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
		// `while` condition must be boolean
		{
			newTestClassType([]ast.Node{
				&ast.While{
					Condition:  &ast.BooleanLiteral{},
					Statements: &ast.Block{},
				},
				&ast.While{
					Condition:  &ast.StringLiteral{},
					Statements: &ast.Block{},
				},
				&ast.While{
					Condition:  &ast.IntegerLiteral{},
					Statements: &ast.Block{},
				},
				&ast.While{
					Condition:  &ast.DoubleLiteral{},
					Statements: &ast.Block{},
				},
			}),
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
		// ternaly condition must be boolean
		{
			newTestClassType([]ast.Node{
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
			}),
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
		// method return type must be return expression type
		{
			&builtin.ClassType{
				Name: "klass",
				InstanceMethods: &builtin.MethodMap{
					Data: map[string][]*builtin.Method{
						"foo": {
							&builtin.Method{
								ReturnType: &ast.TypeRef{Name: []string{"Integer"}},
								Statements: &ast.Block{
									Statements: []ast.Node{
										&ast.Return{
											Expression: &ast.StringLiteral{},
										},
									},
								},
								Parent: &ast.ClassDeclaration{
									Name: "klass",
								},
							},
						},
						"bar": {
							&builtin.Method{
								ReturnType: &ast.TypeRef{Name: []string{"Integer"}},
								Statements: &ast.Block{
									Statements: []ast.Node{
										&ast.Return{
											Expression: &ast.IntegerLiteral{},
										},
									},
								},
								Parent: &ast.ClassDeclaration{
									Name: "klass",
								},
							},
						},
						"baz": {
							&builtin.Method{
								ReturnType: &ast.TypeRef{Name: []string{"Integer"}},
								Statements: &ast.Block{
									Statements: []ast.Node{
										&ast.Return{
											Expression: &ast.DoubleLiteral{},
										},
									},
								},
								Parent: &ast.ClassDeclaration{
									Name: "klass",
								},
							},
						},
						"qux": {
							&builtin.Method{
								ReturnType: &ast.TypeRef{Name: []string{"Integer"}},
								Statements: &ast.Block{
									Statements: []ast.Node{
										&ast.Return{
											Expression: &ast.BooleanLiteral{},
										},
									},
								},
								Parent: &ast.ClassDeclaration{
									Name: "klass",
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
		// for
		// * break must be in for/while loop
		// * condition must be boolean
		{
			newTestClassType([]ast.Node{
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
			}),
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
			newTestClassType([]ast.Node{
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
			}),
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
		// types must equal on variable declaration, variable assignment
		{
			&builtin.ClassType{
				Name: "klass",
				InstanceMethods: &builtin.MethodMap{
					Data: map[string][]*builtin.Method{
						"foo": {
							&builtin.Method{
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
								Parent: &ast.ClassDeclaration{
									Name: "klass",
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
		// arithmetic expression type
		{
			newTestClassType([]ast.Node{
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
			}),
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
	for i, testCase := range testCases {
		checker := NewTypeChecker()
		checker.Context = &Context{}
		checker.Context.ClassTypes = builtin.NewClassMapWithPrimivie([]*builtin.ClassType{testCase.Input})
		checker.VisitClassType(testCase.Input) // TODO: test error

		messages := make([]string, len(checker.Errors))
		for i, err := range checker.Errors {
			messages[i] = err.Message
		}
		sort.Slice(messages, func(i, j int) bool {
			return messages[i] < messages[j]
		})

		if len(messages) != len(testCase.ExpectErrors) {
			t.Errorf("%d: error size is not match: %d != %d", i, len(messages), len(testCase.ExpectErrors))
			t.Errorf("%s", strings.Join(messages, ", "))
		} else {
			for i, message := range messages {
				expected := testCase.ExpectErrors[i]
				if expected.Message != message {
					t.Errorf("%d: expected: %s, actual: %s", i, expected.Message, message)
				}
			}
		}
	}
}

func TestModifier(t *testing.T) {
	testCases := []struct {
		Input       *builtin.ClassType
		ExpectError error
	}{
		// method call on `this` context
		{
			&builtin.ClassType{
				Name: "klass",
				InstanceMethods: &builtin.MethodMap{
					Data: map[string][]*builtin.Method{
						"private_method": {
							&builtin.Method{
								Modifiers: []ast.Node{
									&ast.Modifier{Name: "private"},
								},
								ReturnType: nil,
								Statements: &ast.Block{},
								Parent: &ast.ClassDeclaration{
									Name: "klass",
								},
							},
						},
						"caller": {
							&builtin.Method{
								ReturnType: nil,
								Statements: &ast.Block{
									Statements: []ast.Node{
										&ast.MethodInvocation{
											NameOrExpression: &ast.Name{Value: []string{"this", "private_method"}},
											Parameters:       []ast.Node{},
										},
									},
								},
								Parent: &ast.ClassDeclaration{
									Name: "klass",
								},
							},
						},
					},
				},
			},
			nil,
		},
		//
		{
			&builtin.ClassType{
				Name: "klass",
				InstanceMethods: &builtin.MethodMap{
					Data: map[string][]*builtin.Method{
						"protected_method": {
							&builtin.Method{
								Modifiers: []ast.Node{
									&ast.Modifier{Name: "protected"},
								},
								ReturnType: nil,
								Statements: &ast.Block{},
								Parent: &ast.ClassDeclaration{
									Name: "klass",
								},
							},
						},
						"caller": {
							&builtin.Method{
								ReturnType: nil,
								Statements: &ast.Block{
									Statements: []ast.Node{
										&ast.MethodInvocation{
											NameOrExpression: &ast.FieldAccess{
												Expression: &ast.New{
													Type: &ast.TypeRef{Name: []string{"klass"}},
												},
												FieldName: "protected_method",
											},
											Parameters: []ast.Node{},
										},
									},
								},
								Parent: &ast.ClassDeclaration{
									Name: "klass",
								},
							},
						},
					},
				},
			},
			errors.New("Method access modifier must be public but protected"),
		},
		{
			&builtin.ClassType{
				Name: "klass",
				InstanceMethods: &builtin.MethodMap{
					Data: map[string][]*builtin.Method{
						"public_method": {
							&builtin.Method{
								Modifiers: []ast.Node{
									&ast.Modifier{Name: "public"},
								},
								ReturnType: nil,
								Statements: &ast.Block{},
								Parent: &ast.ClassDeclaration{
									Name: "klass",
								},
							},
						},
						"caller": {
							&builtin.Method{
								ReturnType: nil,
								Statements: &ast.Block{
									Statements: []ast.Node{
										&ast.MethodInvocation{
											NameOrExpression: &ast.FieldAccess{
												Expression: &ast.New{
													Type: &ast.TypeRef{Name: []string{"klass"}},
												},
												FieldName: "public_method",
											},
											Parameters: []ast.Node{},
										},
									},
								},
								Parent: &ast.ClassDeclaration{
									Name: "klass",
								},
							},
						},
					},
				},
			},
			nil,
		},
	}
	for i, testCase := range testCases {
		checker := NewTypeChecker()
		checker.Context = &Context{}
		checker.Context.ClassTypes = builtin.NewClassMapWithPrimivie([]*builtin.ClassType{testCase.Input})
		_, err := checker.VisitClassType(testCase.Input)

		expected := testCase.ExpectError
		if expected == nil {
			if err != nil {
				t.Errorf("%d: unexpected error raised: %s", i, err.Error())
			}
			continue
		}
		if expected != nil && err == nil {
			t.Errorf("%d: error is expected, but not raised: %s", i, expected.Error())
			continue
		}

		if expected.Error() != err.Error() {
			t.Errorf("%d: expected: %s, actual: %s", i, expected.Error(), err.Error())
		}
	}
}

func newTestClassType(statements []ast.Node) *builtin.ClassType {
	return &builtin.ClassType{
		Name: "klass",
		InstanceMethods: &builtin.MethodMap{
			Data: map[string][]*builtin.Method{
				"foo": {
					&builtin.Method{
						Statements: &ast.Block{
							Statements: statements,
						},
						Parent: &ast.ClassDeclaration{
							Name: "klass",
						},
					},
				},
			},
		},
	}
}
