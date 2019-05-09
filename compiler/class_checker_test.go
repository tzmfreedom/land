package compiler

import (
	"testing"

	"errors"

	"github.com/tzmfreedom/land/ast"
	"github.com/tzmfreedom/land/builtin"
)

func TestClassChecker(t *testing.T) {
	testCases := []struct {
		Input         *ast.ClassType
		ExpectedError error
	}{
		// different parameter type
		{
			&ast.ClassType{
				Modifiers:      []*ast.Modifier{ast.PublicModifier()},
				Annotations:    []*ast.Annotation{},
				Name:           "Foo",
				SuperClassRef:  nil,
				InstanceFields: ast.NewFieldMap(),
				StaticFields:   ast.NewFieldMap(),
				InstanceMethods: &ast.MethodMap{
					Data: map[string][]*ast.Method{
						"bar": {
							&ast.Method{
								Name:       "bar",
								ReturnType: builtin.IntegerType,
								Modifiers: []*ast.Modifier{
									{
										Name: "public",
									},
								},
								Parameters: []*ast.Parameter{
									{
										Type: builtin.IntegerType,
										Name: "a",
									},
								},
							},
							&ast.Method{
								Name:       "bar",
								ReturnType: builtin.IntegerType,
								Modifiers: []*ast.Modifier{
									{
										Name: "public",
									},
								},
								Parameters: []*ast.Parameter{
									{
										Type: builtin.StringType,
										Name: "a",
									},
								},
							},
						},
					},
				},
				StaticMethods: ast.NewMethodMap(),
			},
			nil,
		},
		// same parameter and name signature, difference return type
		{
			&ast.ClassType{
				Modifiers:      []*ast.Modifier{ast.PublicModifier()},
				Annotations:    []*ast.Annotation{},
				Name:           "Foo",
				SuperClassRef:  nil,
				InstanceFields: ast.NewFieldMap(),
				StaticFields:   ast.NewFieldMap(),
				InstanceMethods: &ast.MethodMap{
					Data: map[string][]*ast.Method{
						"bar": {
							&ast.Method{
								Name:       "bar",
								ReturnType: builtin.IntegerType,
								Modifiers: []*ast.Modifier{
									{
										Name: "public",
									},
								},
								Parameters: []*ast.Parameter{
									{
										Type: builtin.IntegerType,
										Name: "a",
									},
								},
							},
							&ast.Method{
								Name:       "bar",
								ReturnType: builtin.StringType,
								Modifiers: []*ast.Modifier{
									{
										Name: "public",
									},
								},
								Parameters: []*ast.Parameter{
									{
										Type: builtin.IntegerType,
										Name: "a",
									},
								},
							},
						},
					},
				},
				StaticMethods: ast.NewMethodMap(),
			},
			errors.New("method bar is duplicated"),
		},
		// different parameter number
		{
			&ast.ClassType{
				Modifiers:      []*ast.Modifier{ast.PublicModifier()},
				Annotations:    []*ast.Annotation{},
				Name:           "Foo",
				SuperClassRef:  nil,
				InstanceFields: ast.NewFieldMap(),
				StaticFields:   ast.NewFieldMap(),
				InstanceMethods: &ast.MethodMap{
					Data: map[string][]*ast.Method{
						"bar": {
							&ast.Method{
								Name:       "bar",
								ReturnType: builtin.IntegerType,
								Modifiers: []*ast.Modifier{
									{
										Name: "public",
									},
								},
								Parameters: []*ast.Parameter{
									{
										Type: builtin.IntegerType,
										Name: "a",
									},
								},
							},
							&ast.Method{
								Name:       "bar",
								ReturnType: builtin.IntegerType,
								Modifiers: []*ast.Modifier{
									{
										Name: "public",
									},
								},
								Parameters: []*ast.Parameter{
									{
										Type: builtin.IntegerType,
										Name: "a",
									},
									{
										Type: builtin.IntegerType,
										Name: "b",
									},
								},
							},
						},
					},
				},
				StaticMethods: ast.NewMethodMap(),
			},
			nil,
		},
		// same parameter name
		{
			&ast.ClassType{
				Modifiers:      []*ast.Modifier{ast.PublicModifier()},
				Annotations:    []*ast.Annotation{},
				Name:           "Foo",
				SuperClass:     nil,
				InstanceFields: ast.NewFieldMap(),
				StaticFields:   ast.NewFieldMap(),
				InstanceMethods: &ast.MethodMap{
					Data: map[string][]*ast.Method{
						"bar": {
							&ast.Method{
								Name:       "bar",
								ReturnType: builtin.IntegerType,
								Modifiers: []*ast.Modifier{
									{
										Name: "public",
									},
								},
								Parameters: []*ast.Parameter{
									{
										Type: builtin.IntegerType,
										Name: "a",
									},
									{
										Type: builtin.IntegerType,
										Name: "a",
									},
								},
							},
						},
					},
				},
				StaticMethods: ast.NewMethodMap(),
			},
			errors.New("parameter name is duplicated: a"),
		},
	}
	for i, testCase := range testCases {
		err := CheckClass(testCase.Input)
		if testCase.ExpectedError == nil {
			if err != nil {
				t.Errorf("%d: expect nil, actual %s", i, err.Error())
			}
			continue
		}
		if err == nil {
			t.Errorf("error is not raised, expected %s", testCase.ExpectedError.Error())
			continue
		}
		if testCase.ExpectedError.Error() != err.Error() {
			t.Errorf("%d: expected %s, actual %s", i, testCase.ExpectedError.Error(), err.Error())
		}
	}
}
