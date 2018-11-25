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
											Modifiers: []ast.Node{
												&ast.Modifier{
													Name: "public",
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
			errors.New("Field j is not found"),
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
										Modifiers: []ast.Node{
											&ast.Modifier{
												Name: "public",
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
										Modifiers: []ast.Node{
											&ast.Modifier{
												Name: "public",
											},
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
										Modifiers: []ast.Node{
											&ast.Modifier{
												Name: "public",
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
												Modifiers: []ast.Node{
													&ast.Modifier{
														Name: "public",
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
			},
			builtin.IntegerType,
			nil,
		},
	}
	for _, testCase := range testCases {
		typeResolver := &TypeResolver{Context: testCase.Context}
		actual, err := typeResolver.ResolveVariable(testCase.Input)
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
	for _, testCase := range testCases {
		typeResolver := &TypeResolver{Context: testCase.Context}
		actual, err := typeResolver.ResolveType(testCase.Input)
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
			[]string{"instance"},
			&Context{
				Env: &TypeEnv{
					Data: &TypeMap{
						Data: map[string]*builtin.ClassType{
							"this": {
								Name: "klass",
								InstanceMethods: &builtin.MethodMap{
									Data: map[string][]ast.Node{
										"instance": {
											&ast.MethodDeclaration{
												Name: "instance",
												Modifiers: []ast.Node{
													&ast.Modifier{
														Name: "public",
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
			},
			&ast.MethodDeclaration{
				Name: "instance",
				Modifiers: []ast.Node{
					&ast.Modifier{
						Name: "public",
					},
				},
			},
			nil,
		},
		{
			[]string{"local", "instance"},
			&Context{
				Env: &TypeEnv{
					Data: &TypeMap{
						Data: map[string]*builtin.ClassType{
							"local": {
								Name: "klass",
								InstanceMethods: &builtin.MethodMap{
									Data: map[string][]ast.Node{
										"instance": {
											&ast.MethodDeclaration{
												Name: "instance",
												Modifiers: []ast.Node{
													&ast.Modifier{
														Name: "public",
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
			},
			&ast.MethodDeclaration{
				Name: "instance",
				Modifiers: []ast.Node{
					&ast.Modifier{
						Name: "public",
					},
				},
			},
			nil,
		},
		{
			[]string{"klass", "static"},
			&Context{
				ClassTypes: &builtin.ClassMap{
					Data: map[string]*builtin.ClassType{
						"klass": {
							StaticMethods: &builtin.MethodMap{
								Data: map[string][]ast.Node{
									"static": {
										&ast.MethodDeclaration{
											Name: "static",
											Modifiers: []ast.Node{
												&ast.Modifier{
													Name: "public",
												},
											},
										},
									},
								},
							},
						},
					},
				},
				Env: newTypeEnv(nil),
			},
			&ast.MethodDeclaration{
				Name: "static",
				Modifiers: []ast.Node{
					&ast.Modifier{
						Name: "public",
					},
				},
			},
			nil,
		},
		{
			[]string{"klass", "static", "instance"},
			&Context{
				ClassTypes: &builtin.ClassMap{
					Data: map[string]*builtin.ClassType{
						"klass": {
							StaticFields: &builtin.FieldMap{
								Data: map[string]*builtin.Field{
									"static": {
										Name: "static",
										Modifiers: []ast.Node{
											&ast.Modifier{
												Name: "public",
											},
										},
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
									"instance": {
										&ast.MethodDeclaration{
											Name: "instance",
											Modifiers: []ast.Node{
												&ast.Modifier{
													Name: "public",
												},
											},
										},
									},
								},
							},
						},
					},
				},
				Env: newTypeEnv(nil),
			},
			&ast.MethodDeclaration{
				Name: "instance",
				Modifiers: []ast.Node{
					&ast.Modifier{
						Name: "public",
					},
				},
			},
			nil,
		},
		{
			[]string{"namespace", "klass", "static"},
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
											"static": {
												&ast.MethodDeclaration{
													Name: "static",
													Modifiers: []ast.Node{
														&ast.Modifier{
															Name: "public",
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
				},
			},
			&ast.MethodDeclaration{
				Name: "static",
				Modifiers: []ast.Node{
					&ast.Modifier{
						Name: "public",
					},
				},
			},
			nil,
		},
		{
			[]string{"namespace", "klass", "static", "instance"},
			&Context{
				Env: newTypeEnv(nil),
				ClassTypes: &builtin.ClassMap{
					Data: map[string]*builtin.ClassType{
						"klass2": {
							InstanceMethods: &builtin.MethodMap{
								Data: map[string][]ast.Node{
									"instance": {
										&ast.MethodDeclaration{
											Name: "instance",
											Modifiers: []ast.Node{
												&ast.Modifier{
													Name: "public",
												},
											},
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
											"static": {
												Type: &ast.TypeRef{
													Name: []string{"klass2"},
												},
												Modifiers: []ast.Node{
													&ast.Modifier{
														Name: "public",
													},
												},
												Name: "static",
											},
										},
									},
								},
							},
						},
					},
				},
			},
			&ast.MethodDeclaration{
				Name: "instance",
				Modifiers: []ast.Node{
					&ast.Modifier{
						Name: "public",
					},
				},
			},
			nil,
		},
	}
	for i, testCase := range testCases {
		pp.Println(i)
		typeResolver := &TypeResolver{Context: testCase.Context}
		actual, err := typeResolver.ResolveMethod(testCase.Input, []*builtin.ClassType{})
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
