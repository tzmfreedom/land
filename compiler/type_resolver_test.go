package compiler

import (
	"errors"
	"testing"

	"github.com/google/go-cmp/cmp"
	"github.com/k0kubun/pp"
	"github.com/tzmfreedom/goland/ast"
	"github.com/tzmfreedom/goland/builtin"
)

func TestResolveVariable(t *testing.T) {
	testCases := []struct {
		Input    []string
		Context  *Context
		Expected *builtin.ClassType
		Error    error
	}{
		{
			[]string{"i"},
			&Context{
				ClassTypes: &builtin.ClassMap{
					Data: map[string]*builtin.ClassType{
						"integer": builtin.IntegerType,
					},
				},
				Env: &TypeEnv{
					Data: &TypeMap{
						Data: map[string]*builtin.ClassType{
							"i": builtin.IntegerType,
						},
					},
				},
			},
			builtin.IntegerType,
			nil,
		},
		{
			[]string{"i"},
			&Context{
				ClassTypes: &builtin.ClassMap{
					Data: map[string]*builtin.ClassType{
						"integer": builtin.IntegerType,
					},
				},
				Env: &TypeEnv{
					Data: &TypeMap{
						Data: map[string]*builtin.ClassType{},
					},
				},
			},
			nil,
			errors.New("i is not found in this scope"),
		},
		{
			[]string{"i", "j"},
			&Context{
				ClassTypes: &builtin.ClassMap{
					Data: map[string]*builtin.ClassType{
						"integer": builtin.IntegerType,
					},
				},
				Env: &TypeEnv{
					Data: &TypeMap{
						Data: map[string]*builtin.ClassType{
							"i": {
								InstanceFields: &builtin.FieldMap{
									Data: map[string]*builtin.Field{
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
			builtin.IntegerType,
			nil,
		},
		{
			[]string{"i", "j"},
			&Context{
				ClassTypes: &builtin.ClassMap{
					Data: map[string]*builtin.ClassType{
						"integer": builtin.IntegerType,
					},
				},
				Env: &TypeEnv{
					Data: &TypeMap{
						Data: map[string]*builtin.ClassType{
							"i": {
								InstanceFields: &builtin.FieldMap{
									Data: map[string]*builtin.Field{},
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
				ClassTypes: &builtin.ClassMap{
					Data: map[string]*builtin.ClassType{
						"integer": builtin.IntegerType,
						"i": {
							StaticFields: &builtin.FieldMap{
								Data: map[string]*builtin.Field{
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
			builtin.IntegerType,
			nil,
		},
		{
			[]string{"i", "j", "k"},
			&Context{
				Env: &TypeEnv{
					Data: &TypeMap{},
				},
				ClassTypes: &builtin.ClassMap{
					Data: map[string]*builtin.ClassType{
						"integer": builtin.IntegerType,
						"foo": {
							InstanceFields: &builtin.FieldMap{
								Data: map[string]*builtin.Field{
									"k": {
										Type: &ast.TypeRef{
											Name: []string{"Integer"},
										},
									},
								},
							},
						},
						"i": {
							StaticFields: &builtin.FieldMap{
								Data: map[string]*builtin.Field{
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
			builtin.IntegerType,
			nil,
		},
		{
			[]string{"i", "j", "k"},
			&Context{
				Env: &TypeEnv{
					Data: &TypeMap{},
				},
				ClassTypes: &builtin.ClassMap{
					Data: map[string]*builtin.ClassType{
						"integer": builtin.IntegerType,
					},
				},
				NameSpaces: &builtin.NameSpaceStore{
					Data: map[string]*builtin.ClassMap{
						"i": {
							Data: map[string]*builtin.ClassType{
								"j": {
									StaticFields: &builtin.FieldMap{
										Data: map[string]*builtin.Field{
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
			builtin.IntegerType,
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
		Expected *builtin.ClassType
		Error    error
	}{
		{
			[]string{"i"},
			&Context{
				ClassTypes: &builtin.ClassMap{
					Data: map[string]*builtin.ClassType{
						"i": {
							Name: "i",
						},
					},
				},
			},
			&builtin.ClassType{Name: "i"},
			nil,
		},
		{
			[]string{"i"},
			&Context{
				ClassTypes: builtin.NewClassMap(),
				NameSpaces: &builtin.NameSpaceStore{
					Data: map[string]*builtin.ClassMap{
						"system": {
							Data: map[string]*builtin.ClassType{
								"i": {
									Name: "i",
								},
							},
						},
					},
				},
			},
			&builtin.ClassType{Name: "i"},
			nil,
		},
		{
			[]string{"i", "j"},
			&Context{
				ClassTypes: &builtin.ClassMap{
					Data: map[string]*builtin.ClassType{
						"i": {
							Name: "i",
							InnerClasses: &builtin.ClassMap{
								Data: map[string]*builtin.ClassType{
									"j": {
										Name: "j",
									},
								},
							},
						},
					},
				},
			},
			&builtin.ClassType{Name: "j"},
			nil,
		},
		{
			[]string{"i", "j"},
			&Context{
				ClassTypes: builtin.NewClassMap(),
				NameSpaces: &builtin.NameSpaceStore{
					Data: map[string]*builtin.ClassMap{
						"i": {
							Data: map[string]*builtin.ClassType{
								"j": {
									Name: "j",
								},
							},
						},
					},
				},
			},
			&builtin.ClassType{Name: "j"},
			nil,
		},
		{
			[]string{"i", "j", "k"},
			&Context{
				ClassTypes: builtin.NewClassMap(),
				NameSpaces: &builtin.NameSpaceStore{
					Data: map[string]*builtin.ClassMap{
						"i": {
							Data: map[string]*builtin.ClassType{
								"j": {
									Name: "j",
									InnerClasses: &builtin.ClassMap{
										Data: map[string]*builtin.ClassType{
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
			&builtin.ClassType{Name: "k"},
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

func TestResolveMethod(t *testing.T) {
	testCases := []struct {
		Input    []string
		Context  *Context
		Expected *ast.MethodDeclaration
		Error    error
	}{
		{
			[]string{"foo"},
			&Context{
				Env: &TypeEnv{
					Data: &TypeMap{
						Data: map[string]*builtin.ClassType{
							"this": {
								Name: "klass",
								InstanceMethods: &builtin.MethodMap{
									Data: map[string][]ast.Node{
										"foo": {
											&ast.MethodDeclaration{
												Name: "foo",
											},
										},
									},
								},
							},
						},
					},
				},
			},
			&ast.MethodDeclaration{Name: "foo"},
			nil,
		},
		{
			[]string{"foo", "bar"},
			&Context{
				Env: &TypeEnv{
					Data: &TypeMap{
						Data: map[string]*builtin.ClassType{
							"foo": {
								Name: "klass",
								InstanceMethods: &builtin.MethodMap{
									Data: map[string][]ast.Node{
										"bar": {
											&ast.MethodDeclaration{
												Name: "bar",
											},
										},
									},
								},
							},
						},
					},
				},
			},
			&ast.MethodDeclaration{Name: "bar"},
			nil,
		},
		{
			[]string{"klass", "foo"},
			&Context{
				ClassTypes: &builtin.ClassMap{
					Data: map[string]*builtin.ClassType{
						"klass": {
							StaticMethods: &builtin.MethodMap{
								Data: map[string][]ast.Node{
									"foo": {
										&ast.MethodDeclaration{
											Name: "foo",
										},
									},
								},
							},
						},
					},
				},
				Env: newTypeEnv(nil),
			},
			&ast.MethodDeclaration{Name: "foo"},
			nil,
		},
		{
			[]string{"klass", "foo", "bar"},
			&Context{
				ClassTypes: &builtin.ClassMap{
					Data: map[string]*builtin.ClassType{
						"klass": {
							StaticFields: &builtin.FieldMap{
								Data: map[string]*builtin.Field{
									"foo": {
										Name: "foo",
										Type: &ast.TypeRef{
											Name: []string{"klass2"},
										},
									},
								},
							},
						},
						"klass2": {
							InstanceMethods: &builtin.MethodMap{
								Data: map[string][]ast.Node{
									"bar": {
										&ast.MethodDeclaration{
											Name: "bar",
										},
									},
								},
							},
						},
					},
				},
				Env: newTypeEnv(nil),
			},
			&ast.MethodDeclaration{Name: "bar"},
			nil,
		},
		{
			[]string{"namespace", "klass", "foo"},
			&Context{
				Env:        newTypeEnv(nil),
				ClassTypes: builtin.NewClassMap(),
				NameSpaces: &builtin.NameSpaceStore{
					Data: map[string]*builtin.ClassMap{
						"namespace": {
							Data: map[string]*builtin.ClassType{
								"klass": {
									StaticMethods: &builtin.MethodMap{
										Data: map[string][]ast.Node{
											"foo": {
												&ast.MethodDeclaration{
													Name: "foo",
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
			&ast.MethodDeclaration{Name: "foo"},
			nil,
		},
		{
			[]string{"namespace", "klass", "foo", "bar"},
			&Context{
				Env: newTypeEnv(nil),
				ClassTypes: &builtin.ClassMap{
					Data: map[string]*builtin.ClassType{
						"klass2": {
							InstanceMethods: &builtin.MethodMap{
								Data: map[string][]ast.Node{
									"bar": {
										&ast.MethodDeclaration{
											Name: "bar",
										},
									},
								},
							},
						},
					},
				},
				NameSpaces: &builtin.NameSpaceStore{
					Data: map[string]*builtin.ClassMap{
						"namespace": {
							Data: map[string]*builtin.ClassType{
								"klass": {
									StaticFields: &builtin.FieldMap{
										Data: map[string]*builtin.Field{
											"foo": {
												Type: &ast.TypeRef{
													Name: []string{"klass2"},
												},
												Name: "foo",
											},
										},
									},
								},
							},
						},
					},
				},
			},
			&ast.MethodDeclaration{Name: "bar"},
			nil,
		},
	}
	typeResolver := &TypeResolver{}
	for _, testCase := range testCases {
		actual, err := typeResolver.ResolveMethod(testCase.Input, testCase.Context)
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
