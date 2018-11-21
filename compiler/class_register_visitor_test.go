package compiler

import (
	"testing"

	"github.com/google/go-cmp/cmp"
	"github.com/k0kubun/pp"
	"github.com/tzmfreedom/goland/ast"
)

func TestClassRegister(t *testing.T) {
	testCases := []struct {
		Input    ast.Node
		Expected *ClassType
	}{
		{
			&ast.ClassDeclaration{
				Modifiers:   []ast.Node{},
				Annotations: []ast.Node{},
				Name:        "Foo",
				Declarations: []ast.Node{
					&ast.FieldDeclaration{
						Type: &ast.TypeRef{
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
						Type: &ast.TypeRef{
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
						Type: &ast.TypeRef{
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
						Type: &ast.TypeRef{
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
						ReturnType: &ast.TypeRef{
							Name: []string{
								"String",
							},
							Parameters: []ast.Node{},
						},
						Parameters: []ast.Node{
							&ast.Parameter{
								Modifiers: []ast.Node{},
								Type: &ast.TypeRef{
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
						ReturnType: nil,
						Parameters: []ast.Node{},
						Throws:     []ast.Node{},
						Statements: &ast.Block{
							Statements: []ast.Node{},
						},
						NativeFunction: nil,
					},
				},
			},
			&ClassType{
				Modifiers:    []ast.Node{},
				Annotations:  []ast.Node{},
				Name:         "Foo",
				InnerClasses: NewClassMap(),
				InstanceMethods: &MethodMap{
					Data: map[string][]ast.Node{
						"method": {
							&ast.MethodDeclaration{
								Name: "method",
								Modifiers: []ast.Node{
									&ast.Modifier{
										Name: "public",
									},
								},
								ReturnType: nil,
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
				StaticMethods: &MethodMap{
					Data: map[string][]ast.Node{
						"static_method": {
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
								ReturnType: &ast.TypeRef{
									Name: []string{
										"String",
									},
									Parameters: []ast.Node{},
								},
								Parameters: []ast.Node{
									&ast.Parameter{
										Modifiers: []ast.Node{},
										Type: &ast.TypeRef{
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
						},
					},
				},
				InstanceFields: &FieldMap{
					Data: map[string]*Field{
						"field": {
							Type: &ast.TypeRef{
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
							Name:       "field",
							Expression: nil,
						},
						"field_with_init": {
							Type: &ast.TypeRef{
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
							Name: "field_with_init",
							Expression: &ast.IntegerLiteral{
								Value: 2,
							},
						},
					},
				},
				StaticFields: &FieldMap{
					Data: map[string]*Field{
						"static_field": {
							Type: &ast.TypeRef{
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
							Name:       "static_field",
							Expression: nil,
						},
						"static_field_with_init": {
							Type: &ast.TypeRef{
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
							Name: "static_field_with_init",
							Expression: &ast.IntegerLiteral{
								Value: 1,
							},
						},
					},
				},
			},
		},
	}
	for _, testCase := range testCases {
		classRegister := &ClassRegisterVisitor{}
		actual, err := testCase.Input.Accept(classRegister)
		if err != nil {
			panic(err)
		}

		equalNode(t, testCase.Expected, actual.(*ClassType))
	}
}

func equalNode(t *testing.T, expected *ClassType, actual *ClassType) {
	if ok := cmp.Equal(expected, actual); !ok {
		diff := cmp.Diff(expected, actual)
		pp.Print(actual)
		t.Errorf(diff)
	}
}
