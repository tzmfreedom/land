package compiler

import (
	"testing"

	"errors"

	"github.com/tzmfreedom/goland/ast"
	"github.com/tzmfreedom/goland/builtin"
)

func TestClassChecker(t *testing.T) {
	testCases := []struct {
		Input         *builtin.ClassType
		ExpectedError error
	}{
		// difference parameter type
		{
			&builtin.ClassType{
				Modifiers:      []ast.Node{},
				Annotations:    []ast.Node{},
				Name:           "Foo",
				SuperClass:     nil,
				InstanceFields: builtin.NewFieldMap(),
				StaticFields:   builtin.NewFieldMap(),
				InstanceMethods: &builtin.MethodMap{
					Data: map[string][]*builtin.Method{
						"bar": {
							&builtin.Method{
								Name: "bar",
								ReturnType: &ast.TypeRef{
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
								Parameters: []ast.Node{
									&ast.Parameter{
										Type: &ast.TypeRef{Name: []string{"Integer"}},
										Name: "a",
									},
								},
							},
							&builtin.Method{
								Name: "bar",
								ReturnType: &ast.TypeRef{
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
								Parameters: []ast.Node{
									&ast.Parameter{
										Type: &ast.TypeRef{Name: []string{"String"}},
										Name: "a",
									},
								},
							},
						},
					},
				},
				StaticMethods: builtin.NewMethodMap(),
			},
			nil,
		},
		// same parameter and name signature, difference return type
		{
			&builtin.ClassType{
				Modifiers:      []ast.Node{},
				Annotations:    []ast.Node{},
				Name:           "Foo",
				SuperClass:     nil,
				InstanceFields: builtin.NewFieldMap(),
				StaticFields:   builtin.NewFieldMap(),
				InstanceMethods: &builtin.MethodMap{
					Data: map[string][]*builtin.Method{
						"bar": {
							&builtin.Method{
								Name: "bar",
								ReturnType: &ast.TypeRef{
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
								Parameters: []ast.Node{
									&ast.Parameter{
										Type: &ast.TypeRef{Name: []string{"Integer"}},
										Name: "a",
									},
								},
							},
							&builtin.Method{
								Name: "bar",
								ReturnType: &ast.TypeRef{
									Name: []string{
										"String",
									},
									Parameters: []ast.Node{},
								},
								Modifiers: []ast.Node{
									&ast.Modifier{
										Name: "public",
									},
								},
								Parameters: []ast.Node{
									&ast.Parameter{
										Type: &ast.TypeRef{Name: []string{"Integer"}},
										Name: "a",
									},
								},
							},
						},
					},
				},
				StaticMethods: builtin.NewMethodMap(),
			},
			errors.New("method bar is duplicated"),
		},
		// different parameter number
		{
			&builtin.ClassType{
				Modifiers:      []ast.Node{},
				Annotations:    []ast.Node{},
				Name:           "Foo",
				SuperClass:     nil,
				InstanceFields: builtin.NewFieldMap(),
				StaticFields:   builtin.NewFieldMap(),
				InstanceMethods: &builtin.MethodMap{
					Data: map[string][]*builtin.Method{
						"bar": {
							&builtin.Method{
								Name: "bar",
								ReturnType: &ast.TypeRef{
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
								Parameters: []ast.Node{
									&ast.Parameter{
										Type: &ast.TypeRef{Name: []string{"Integer"}},
										Name: "a",
									},
								},
							},
							&builtin.Method{
								Name: "bar",
								ReturnType: &ast.TypeRef{
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
								Parameters: []ast.Node{
									&ast.Parameter{
										Type: &ast.TypeRef{Name: []string{"Integer"}},
										Name: "a",
									},
									&ast.Parameter{
										Type: &ast.TypeRef{Name: []string{"Integer"}},
										Name: "b",
									},
								},
							},
						},
					},
				},
				StaticMethods: builtin.NewMethodMap(),
			},
			nil,
		},
		// same parameter name
		{
			&builtin.ClassType{
				Modifiers:      []ast.Node{},
				Annotations:    []ast.Node{},
				Name:           "Foo",
				SuperClass:     nil,
				InstanceFields: builtin.NewFieldMap(),
				StaticFields:   builtin.NewFieldMap(),
				InstanceMethods: &builtin.MethodMap{
					Data: map[string][]*builtin.Method{
						"bar": {
							&builtin.Method{
								Name: "bar",
								ReturnType: &ast.TypeRef{
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
								Parameters: []ast.Node{
									&ast.Parameter{
										Type: &ast.TypeRef{Name: []string{"Integer"}},
										Name: "a",
									},
									&ast.Parameter{
										Type: &ast.TypeRef{Name: []string{"Integer"}},
										Name: "a",
									},
								},
							},
						},
					},
				},
				StaticMethods: builtin.NewMethodMap(),
			},
			errors.New("parameter name is duplicated: a"),
		},
	}
	for _, testCase := range testCases {
		checker := &ClassChecker{}
		checker.Context = &Context{}
		checker.Context.ClassTypes = builtin.NewClassMapWithPrimivie(nil)
		err := checker.Check(testCase.Input)
		if testCase.ExpectedError == nil {
			if err != nil {
				t.Fatalf("expect nil, actual %s", err.Error())
			}
			continue
		}
		if err == nil {
			t.Fatalf("error is not raised, expected %s", testCase.ExpectedError.Error())
			continue
		}
		if testCase.ExpectedError.Error() != err.Error() {
			t.Fatalf("expected %s, actual %s", testCase.ExpectedError.Error(), err.Error())
		}
	}
}
