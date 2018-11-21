package compiler

import (
	"errors"
	"testing"

	"github.com/google/go-cmp/cmp"
	"github.com/k0kubun/pp"
	"github.com/tzmfreedom/goland/ast"
)

func TestResolveVariable(t *testing.T) {
	testCases := []struct {
		Input    []string
		Context  *Context
		Expected Type
		Error    error
	}{
		{
			[]string{"i"},
			&Context{
				ClassTypes: &ClassMap{
					Data: map[string]*ClassType{
						"integer": IntegerType,
					},
				},
				Env: &TypeEnv{
					Data: &TypeMap{
						Data: map[string]Type{
							"i": IntegerType,
						},
					},
				},
			},
			IntegerType,
			nil,
		},
		{
			[]string{"i"},
			&Context{
				ClassTypes: &ClassMap{
					Data: map[string]*ClassType{
						"integer": IntegerType,
					},
				},
				Env: &TypeEnv{
					Data: &TypeMap{
						Data: map[string]Type{},
					},
				},
			},
			nil,
			errors.New("i is not found in this scope"),
		},
		{
			[]string{"i", "j"},
			&Context{
				ClassTypes: &ClassMap{
					Data: map[string]*ClassType{
						"integer": IntegerType,
					},
				},
				Env: &TypeEnv{
					Data: &TypeMap{
						Data: map[string]Type{
							"i": &ClassType{
								InstanceFields: &FieldMap{
									Data: map[string]*Field{
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
			IntegerType,
			nil,
		},
		{
			[]string{"i", "j"},
			&Context{
				ClassTypes: &ClassMap{
					Data: map[string]*ClassType{
						"integer": IntegerType,
					},
				},
				Env: &TypeEnv{
					Data: &TypeMap{
						Data: map[string]Type{
							"i": &ClassType{
								InstanceFields: &FieldMap{
									Data: map[string]*Field{},
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
					Data: map[string]*ClassType{
						"integer": IntegerType,
						"i": {
							StaticFields: &FieldMap{
								Data: map[string]*Field{
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
			IntegerType,
			nil,
		},
		{
			[]string{"i", "j", "k"},
			&Context{
				Env: &TypeEnv{
					Data: &TypeMap{},
				},
				ClassTypes: &ClassMap{
					Data: map[string]*ClassType{
						"integer": IntegerType,
						"foo": {
							InstanceFields: &FieldMap{
								Data: map[string]*Field{
									"k": {
										Type: &ast.TypeRef{
											Name: []string{"Integer"},
										},
									},
								},
							},
						},
						"i": {
							StaticFields: &FieldMap{
								Data: map[string]*Field{
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
			IntegerType,
			nil,
		},
		{
			[]string{"i", "j", "k"},
			&Context{
				Env: &TypeEnv{
					Data: &TypeMap{},
				},
				ClassTypes: &ClassMap{
					Data: map[string]*ClassType{
						"integer": IntegerType,
					},
				},
				NameSpaces: &NameSpaceStore{
					Data: map[string]*ClassMap{
						"i": {
							Data: map[string]*ClassType{
								"j": {
									StaticFields: &FieldMap{
										Data: map[string]*Field{
											"k": {
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
			},
			IntegerType,
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

func TestResolveType(t *testing.T) {
	testCases := []struct {
		Input    []string
		Context  *Context
		Expected Type
		Error    error
	}{
		{
			[]string{"i"},
			&Context{
				ClassTypes: &ClassMap{
					Data: map[string]*ClassType{
						"i": {
							Name: "i",
						},
					},
				},
			},
			&ClassType{Name: "i"},
			nil,
		},
		{
			[]string{"i"},
			&Context{
				ClassTypes: NewClassMap(),
				NameSpaces: &NameSpaceStore{
					Data: map[string]*ClassMap{
						"system": {
							Data: map[string]*ClassType{
								"i": {
									Name: "i",
								},
							},
						},
					},
				},
			},
			&ClassType{Name: "i"},
			nil,
		},
		{
			[]string{"i", "j"},
			&Context{
				ClassTypes: &ClassMap{
					Data: map[string]*ClassType{
						"i": {
							Name: "i",
							InnerClasses: &ClassMap{
								Data: map[string]*ClassType{
									"j": {
										Name: "j",
									},
								},
							},
						},
					},
				},
			},
			&ClassType{Name: "j"},
			nil,
		},
		{
			[]string{"i", "j"},
			&Context{
				ClassTypes: NewClassMap(),
				NameSpaces: &NameSpaceStore{
					Data: map[string]*ClassMap{
						"i": {
							Data: map[string]*ClassType{
								"j": {
									Name: "j",
								},
							},
						},
					},
				},
			},
			&ClassType{Name: "j"},
			nil,
		},
		{
			[]string{"i", "j", "k"},
			&Context{
				ClassTypes: NewClassMap(),
				NameSpaces: &NameSpaceStore{
					Data: map[string]*ClassMap{
						"i": {
							Data: map[string]*ClassType{
								"j": {
									Name: "j",
									InnerClasses: &ClassMap{
										Data: map[string]*ClassType{
											"k": {
												Name: "k",
											},
										},
									},
								},
							},
						},
					},
				},
			},
			&ClassType{Name: "k"},
			nil,
		},
	}
	typeResolver := &TypeResolver{}
	for _, testCase := range testCases {
		actual, err := typeResolver.ResolveType(testCase.Input, testCase.Context)
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
