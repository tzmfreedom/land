package compiler

import (
	"testing"

	"errors"

	"github.com/google/go-cmp/cmp"
	"github.com/google/go-cmp/cmp/cmpopts"
	"github.com/k0kubun/pp"
	"github.com/tzmfreedom/land/ast"
)

func TestClassRegister(t *testing.T) {
	testCases := []struct {
		Input    ast.Node
		Expected *ast.ClassType
	}{
		{
			&ast.ClassDeclaration{
				Modifiers:   []*ast.Modifier{},
				Annotations: []*ast.Annotation{},
				Name:        "Foo",
				Declarations: []ast.Node{
					&ast.FieldDeclaration{
						TypeRef: &ast.TypeRef{
							Name: []string{
								"Integer",
							},
							Parameters: []*ast.TypeRef{},
						},
						Modifiers: []*ast.Modifier{
							{
								Name: "public",
							},
						},
						Declarators: []*ast.VariableDeclarator{
							&ast.VariableDeclarator{
								Name:       "field",
								Expression: nil,
							},
						},
					},
					&ast.FieldDeclaration{
						TypeRef: &ast.TypeRef{
							Name: []string{
								"Double",
							},
							Parameters: []*ast.TypeRef{},
						},
						Modifiers: []*ast.Modifier{
							{
								Name: "public",
							},
						},
						Declarators: []*ast.VariableDeclarator{
							&ast.VariableDeclarator{
								Name: "field_with_init",
								Expression: &ast.IntegerLiteral{
									Value: 2,
								},
							},
						},
					},
					&ast.FieldDeclaration{
						TypeRef: &ast.TypeRef{
							Name: []string{
								"String",
							},
							Parameters: []*ast.TypeRef{},
						},
						Modifiers: []*ast.Modifier{
							{
								Name: "public",
							},
							{
								Name: "static",
							},
						},
						Declarators: []*ast.VariableDeclarator{
							&ast.VariableDeclarator{
								Name:       "static_field",
								Expression: nil,
							},
						},
					},
					&ast.FieldDeclaration{
						TypeRef: &ast.TypeRef{
							Name: []string{
								"Boolean",
							},
							Parameters: []*ast.TypeRef{},
						},
						Modifiers: []*ast.Modifier{
							{
								Name: "public",
							},
							{
								Name: "static",
							},
						},
						Declarators: []*ast.VariableDeclarator{
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
						Modifiers: []*ast.Modifier{
							{
								Name: "public",
							},
							{
								Name: "static",
							},
						},
						ReturnType: &ast.TypeRef{
							Name: []string{
								"String",
							},
							Parameters: []*ast.TypeRef{},
						},
						Parameters: []*ast.Parameter{
							{
								Modifiers: []*ast.Modifier{},
								TypeRef: &ast.TypeRef{
									Name: []string{
										"Boolean",
									},
									Parameters: []*ast.TypeRef{},
								},
								Name: "p1",
							},
						},
						Throws: []ast.Node{},
						Statements: &ast.Block{
							Statements: []ast.Node{},
						},
					},
					&ast.MethodDeclaration{
						Name: "method",
						Modifiers: []*ast.Modifier{
							{
								Name: "public",
							},
						},
						ReturnType: nil,
						Parameters: []*ast.Parameter{},
						Throws:     []ast.Node{},
						Statements: &ast.Block{
							Statements: []ast.Node{},
						},
					},
				},
			},
			&ast.ClassType{
				Modifiers:    []*ast.Modifier{},
				Annotations:  []*ast.Annotation{},
				Name:         "Foo",
				InnerClasses: ast.NewClassMap(),
				Constructors: []*ast.Method{},
				InstanceMethods: &ast.MethodMap{
					Data: map[string][]*ast.Method{
						"method": {
							&ast.Method{
								Name: "method",
								Modifiers: []*ast.Modifier{
									{
										Name: "public",
									},
								},
								ReturnTypeRef: nil,
								Parameters:    []*ast.Parameter{},
								Throws:        []ast.Node{},
								Statements: &ast.Block{
									Statements: []ast.Node{},
								},
							},
						},
					},
				},
				StaticMethods: &ast.MethodMap{
					Data: map[string][]*ast.Method{
						"static_method": {
							&ast.Method{
								Name: "static_method",
								Modifiers: []*ast.Modifier{
									{
										Name: "public",
									},
									{
										Name: "static",
									},
								},
								ReturnTypeRef: &ast.TypeRef{
									Name: []string{
										"String",
									},
									Parameters: []*ast.TypeRef{},
								},
								Parameters: []*ast.Parameter{
									{
										Modifiers: []*ast.Modifier{},
										TypeRef: &ast.TypeRef{
											Name: []string{
												"Boolean",
											},
											Parameters: []*ast.TypeRef{},
										},
										Name: "p1",
									},
								},
								Throws: []ast.Node{},
								Statements: &ast.Block{
									Statements: []ast.Node{},
								},
							},
						},
					},
				},
				InstanceFields: &ast.FieldMap{
					Data: map[string]*ast.Field{
						"field": {
							TypeRef: &ast.TypeRef{
								Name: []string{
									"Integer",
								},
								Parameters: []*ast.TypeRef{},
							},
							Modifiers: []*ast.Modifier{
								{
									Name: "public",
								},
							},
							Name:       "field",
							Expression: nil,
						},
						"field_with_init": {
							TypeRef: &ast.TypeRef{
								Name: []string{
									"Double",
								},
								Parameters: []*ast.TypeRef{},
							},
							Modifiers: []*ast.Modifier{
								{
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
				StaticFields: &ast.FieldMap{
					Data: map[string]*ast.Field{
						"static_field": {
							TypeRef: &ast.TypeRef{
								Name: []string{
									"String",
								},
								Parameters: []*ast.TypeRef{},
							},
							Modifiers: []*ast.Modifier{
								{
									Name: "public",
								},
								{
									Name: "static",
								},
							},
							Name:       "static_field",
							Expression: nil,
						},
						"static_field_with_init": {
							TypeRef: &ast.TypeRef{
								Name: []string{
									"Boolean",
								},
								Parameters: []*ast.TypeRef{},
							},
							Modifiers: []*ast.Modifier{
								{
									Name: "public",
								},
								{
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
	for i, testCase := range testCases {
		classRegister := &ClassRegisterVisitor{}
		actual, err := testCase.Input.Accept(classRegister)
		if err != nil {
			panic(err)
		}

		equalNode(t, i, testCase.Expected, actual.(*ast.ClassType))
	}
}

func equalNode(t *testing.T, i int, expected *ast.ClassType, actual *ast.ClassType) {
	ignore := cmpopts.IgnoreFields(ast.Method{}, "Parent")
	if ok := cmp.Equal(expected, actual, ignore); !ok {
		diff := cmp.Diff(expected, actual, ignore)
		pp.Print(actual)
		t.Errorf(diff)
	}
}

func TestClassRegisterDuplicateClass(t *testing.T) {
	classNode := &ast.ClassDeclaration{
		Modifiers:   []*ast.Modifier{},
		Annotations: []*ast.Annotation{},
		Name:        "Foo",
		Declarations: []ast.Node{
			&ast.ClassDeclaration{
				Modifiers:    []*ast.Modifier{},
				Annotations:  []*ast.Annotation{},
				Name:         "Bar",
				Declarations: []ast.Node{},
			},
			&ast.ClassDeclaration{
				Modifiers:    []*ast.Modifier{},
				Annotations:  []*ast.Annotation{},
				Name:         "Bar",
				Declarations: []ast.Node{},
			},
		},
		InnerClasses: []*ast.ClassDeclaration{},
	}
	classRegister := &ClassRegisterVisitor{}
	classNode.Accept(classRegister)
	_, err := classNode.Accept(classRegister)
	expected := "Class Bar is already defined"
	if err == nil {
		t.Errorf("error is not raised, expected %s", expected)
		return
	}
	if err.Error() != expected {
		t.Errorf("expected %s, actual %s", expected, err.Error())
	}
}

func TestClassRegisterError(t *testing.T) {
	testCases := []struct {
		Input         ast.Node
		ExpectedError error
	}{
		{
			&ast.ClassDeclaration{
				Modifiers:   []*ast.Modifier{},
				Annotations: []*ast.Annotation{},
				Name:        "Foo",
				Declarations: []ast.Node{
					&ast.FieldDeclaration{
						TypeRef: &ast.TypeRef{
							Name: []string{
								"Integer",
							},
							Parameters: []*ast.TypeRef{},
						},
						Declarators: []*ast.VariableDeclarator{
							&ast.VariableDeclarator{
								Name:       "f",
								Expression: nil,
							},
						},
					},
					&ast.FieldDeclaration{
						TypeRef: &ast.TypeRef{
							Name: []string{
								"Double",
							},
							Parameters: []*ast.TypeRef{},
						},
						Declarators: []*ast.VariableDeclarator{
							{
								Name: "f",
								Expression: &ast.IntegerLiteral{
									Value: 2,
								},
							},
						},
					},
				},
			},
			errors.New("Field f is already defined"),
		},
		{
			&ast.ClassDeclaration{
				Modifiers:   []*ast.Modifier{},
				Annotations: []*ast.Annotation{},
				Name:        "Foo",
				Declarations: []ast.Node{
					&ast.FieldDeclaration{
						TypeRef: &ast.TypeRef{
							Name: []string{
								"Integer",
							},
							Parameters: []*ast.TypeRef{},
						},
						Modifiers: []*ast.Modifier{
							{
								Name: "static",
							},
						},
						Declarators: []*ast.VariableDeclarator{
							{
								Name:       "f",
								Expression: nil,
							},
						},
					},
					&ast.FieldDeclaration{
						TypeRef: &ast.TypeRef{
							Name: []string{
								"Double",
							},
							Parameters: []*ast.TypeRef{},
						},
						Modifiers: []*ast.Modifier{
							{
								Name: "static",
							},
						},
						Declarators: []*ast.VariableDeclarator{
							{
								Name: "f",
								Expression: &ast.IntegerLiteral{
									Value: 2,
								},
							},
						},
					},
				},
			},
			errors.New("Field f is already defined"),
		},
		{
			&ast.ClassDeclaration{
				Modifiers:   []*ast.Modifier{},
				Annotations: []*ast.Annotation{},
				Name:        "Foo",
				Declarations: []ast.Node{
					&ast.FieldDeclaration{
						TypeRef: &ast.TypeRef{
							Name: []string{
								"Integer",
							},
							Parameters: []*ast.TypeRef{},
						},
						Modifiers: []*ast.Modifier{
							{
								Name: "static",
							},
						},
						Declarators: []*ast.VariableDeclarator{
							{
								Name:       "f",
								Expression: nil,
							},
						},
					},
					&ast.FieldDeclaration{
						TypeRef: &ast.TypeRef{
							Name: []string{
								"Double",
							},
							Parameters: []*ast.TypeRef{},
						},
						Declarators: []*ast.VariableDeclarator{
							{
								Name: "f",
								Expression: &ast.IntegerLiteral{
									Value: 2,
								},
							},
						},
					},
				},
			},
			nil,
		},
	}

	for _, testCase := range testCases {
		classRegister := &ClassRegisterVisitor{}
		_, err := testCase.Input.Accept(classRegister)
		if testCase.ExpectedError == nil {
			if err != nil {
				t.Errorf("expect nil, actual %s", err.Error())
			}
			continue
		}
		if err == nil {
			t.Errorf("error is not raised, expected %s", testCase.ExpectedError.Error())
			continue
		}
		if testCase.ExpectedError.Error() != err.Error() {
			t.Errorf("expected %s, actual %s", testCase.ExpectedError.Error(), err.Error())
		}
	}
}
