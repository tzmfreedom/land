package compiler

import (
	"testing"

	"errors"

	"github.com/tzmfreedom/goland/ast"
)

func TestClassChecker(t *testing.T) {
	testCases := []struct {
		Input         *ast.ClassType
		ExpectedError error
	}{
		// difference parameter type
		{
			&ast.ClassType{
				Modifiers:      []ast.Node{},
				Annotations:    []ast.Node{},
				Name:           "Foo",
				SuperClassRef:  nil,
				InstanceFields: ast.NewFieldMap(),
				StaticFields:   ast.NewFieldMap(),
				InstanceMethods: &ast.MethodMap{
					Data: map[string][]*ast.Method{
						"bar": {
							&ast.Method{
								Name: "bar",
								ReturnType: &ast.TypeRef{
									Name: []string{
										"Integer",
									},
									Parameters: []*ast.TypeRef{},
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
							&ast.Method{
								Name: "bar",
								ReturnType: &ast.TypeRef{
									Name: []string{
										"Integer",
									},
									Parameters: []*ast.TypeRef{},
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
				StaticMethods: ast.NewMethodMap(),
			},
			nil,
		},
		// same parameter and name signature, difference return type
		{
			&ast.ClassType{
				Modifiers:      []ast.Node{},
				Annotations:    []ast.Node{},
				Name:           "Foo",
				SuperClassRef:  nil,
				InstanceFields: ast.NewFieldMap(),
				StaticFields:   ast.NewFieldMap(),
				InstanceMethods: &ast.MethodMap{
					Data: map[string][]*ast.Method{
						"bar": {
							&ast.Method{
								Name: "bar",
								ReturnType: &ast.TypeRef{
									Name: []string{
										"Integer",
									},
									Parameters: []*ast.TypeRef{},
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
							&ast.Method{
								Name: "bar",
								ReturnType: &ast.TypeRef{
									Name: []string{
										"String",
									},
									Parameters: []*ast.TypeRef{},
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
				StaticMethods: ast.NewMethodMap(),
			},
			errors.New("method bar is duplicated"),
		},
		// different parameter number
		{
			&ast.ClassType{
				Modifiers:      []ast.Node{},
				Annotations:    []ast.Node{},
				Name:           "Foo",
				SuperClassRef:  nil,
				InstanceFields: ast.NewFieldMap(),
				StaticFields:   ast.NewFieldMap(),
				InstanceMethods: &ast.MethodMap{
					Data: map[string][]*ast.Method{
						"bar": {
							&ast.Method{
								Name: "bar",
								ReturnType: &ast.TypeRef{
									Name: []string{
										"Integer",
									},
									Parameters: []*ast.TypeRef{},
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
							&ast.Method{
								Name: "bar",
								ReturnType: &ast.TypeRef{
									Name: []string{
										"Integer",
									},
									Parameters: []*ast.TypeRef{},
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
				StaticMethods: ast.NewMethodMap(),
			},
			nil,
		},
		// same parameter name
		{
			&ast.ClassType{
				Modifiers:      []ast.Node{},
				Annotations:    []ast.Node{},
				Name:           "Foo",
				SuperClassRef:  nil,
				InstanceFields: ast.NewFieldMap(),
				StaticFields:   ast.NewFieldMap(),
				InstanceMethods: &ast.MethodMap{
					Data: map[string][]*ast.Method{
						"bar": {
							&ast.Method{
								Name: "bar",
								ReturnType: &ast.TypeRef{
									Name: []string{
										"Integer",
									},
									Parameters: []*ast.TypeRef{},
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
				StaticMethods: ast.NewMethodMap(),
			},
			errors.New("parameter name is duplicated: a"),
		},
	}
	for _, testCase := range testCases {
		checker := &ClassChecker{}
		checker.Context = &Context{}
		checker.Context.ClassTypes = ast.NewClassMapWithPrimivie(nil)
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
