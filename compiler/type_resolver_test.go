package compiler

import (
	"errors"
	"testing"

	"github.com/google/go-cmp/cmp"
	"github.com/k0kubun/pp"
	"github.com/tzmfreedom/goland/ast"
)

func TestResolve(t *testing.T) {
	testCases := []struct {
		Input    []string
		Context  *Context
		Expected ast.Type
		Error    error
	}{
		{
			[]string{"i"},
			&Context{
				Env: &TypeEnv{
					Data: &TypeMap{
						Data: map[string]ast.Type{
							"i": ast.IntegerType,
						},
					},
				},
			},
			ast.IntegerType,
			nil,
		},
		{
			[]string{"i"},
			&Context{
				Env: &TypeEnv{
					Data: &TypeMap{
						Data: map[string]ast.Type{},
					},
				},
			},
			nil,
			errors.New("i is not found in this scope"),
		},
		{
			[]string{"i", "j"},
			&Context{
				Env: &TypeEnv{
					Data: &TypeMap{
						Data: map[string]ast.Type{
							"i": &ast.ClassType{
								InstanceFields: &ast.FieldMap{
									Data: map[string]*ast.Field{
										"j": {
											Type: &ast.TypeRef{
												Name: []string{"Integer"},
											},
										},
									},
								},
							},
						},
					},
				},
			},
			ast.IntegerType,
			nil,
		},
		{
			[]string{"i", "j"},
			&Context{
				Env: &TypeEnv{
					Data: &TypeMap{
						Data: map[string]ast.Type{
							"i": &ast.ClassType{
								InstanceFields: &ast.FieldMap{
									Data: map[string]*ast.Field{},
								},
							},
						},
					},
				},
			},
			nil,
			errors.New("j is not found in this scope"),
		},
		{
			[]string{"i", "j"},
			&Context{
				Env: &TypeEnv{
					Data: &TypeMap{},
				},
				ClassTypes: &ClassMap{
					Data: map[string]*ast.ClassType{
						"i": {
							StaticFields: &ast.FieldMap{
								Data: map[string]*ast.Field{
									"j": {
										Type: &ast.TypeRef{
											Name: []string{"Integer"},
										},
									},
								},
							},
						},
					},
				},
			},
			ast.IntegerType,
			nil,
		},
		{
			[]string{"i", "j", "k"},
			&Context{
				Env: &TypeEnv{
					Data: &TypeMap{},
				},
				ClassTypes: &ClassMap{
					Data: map[string]*ast.ClassType{
						"foo": {
							InstanceFields: &ast.FieldMap{
								Data: map[string]*ast.Field{
									"k": {
										Type: &ast.TypeRef{
											Name: []string{"Integer"},
										},
									},
								},
							},
						},
						"i": {
							StaticFields: &ast.FieldMap{
								Data: map[string]*ast.Field{
									"j": {
										Type: &ast.TypeRef{
											Name: []string{"foo"},
										},
									},
								},
							},
						},
					},
				},
			},
			ast.IntegerType,
			nil,
		},
		{
			[]string{"i", "j", "k"},
			&Context{
				Env: &TypeEnv{
					Data: &TypeMap{},
				},
				ClassTypes: &ClassMap{},
				NameSpaces: &NameSpaceStore{
					Data: map[string]*ClassMap{
						"i": {
							Data: map[string]*ast.ClassType{
								"j": {
									StaticFields: &ast.FieldMap{
										Data: map[string]*ast.Field{
											"k": {
												Type: &ast.TypeRef{
													Name: []string{"integer"},
												},
											},
										},
									},
								},
							},
						},
					},
				},
			},
			ast.IntegerType,
			nil,
		},
	}
	typeResolver := &TypeResolver{}
	for _, testCase := range testCases {
		actual, err := typeResolver.ResolveVariable(testCase.Input, testCase.Context)
		if testCase.Error != nil && testCase.Error.Error() != err.Error() {
			diff := cmp.Diff(testCase.Error.Error(), err.Error())
			t.Errorf(diff)
		}

		if ok := cmp.Equal(testCase.Expected, actual); !ok {
			diff := cmp.Diff(testCase.Expected, actual)
			pp.Print(actual)
			t.Errorf(diff)
		}
	}
}
