package interpreter

import (
	"errors"
	"testing"

	"github.com/google/go-cmp/cmp"
	"github.com/google/go-cmp/cmp/cmpopts"
	"github.com/k0kubun/pp"
	"github.com/tzmfreedom/land/ast"
	"github.com/tzmfreedom/land/builtin"
)

func TestResolveVariable(t *testing.T) {
	testCases := []struct {
		Input    []string
		Context  *Context
		Expected *ast.Object
		Error    error
	}{
		{
			[]string{"i"},
			&Context{
				ClassTypes: &ast.ClassMap{
					Data: map[string]*ast.ClassType{
						"integer": builtin.IntegerType,
					},
				},
				Env: &Env{
					Data: &ast.ObjectMap{
						Data: map[string]*ast.Object{
							"i": builtin.NewInteger(1),
						},
					},
				},
			},
			builtin.NewInteger(1),
			nil,
		},
		{
			[]string{"i"},
			&Context{
				ClassTypes: &ast.ClassMap{
					Data: map[string]*ast.ClassType{
						"integer": builtin.IntegerType,
					},
				},
				Env: &Env{
					Data: &ast.ObjectMap{
						Data: map[string]*ast.Object{},
					},
				},
			},
			nil,
			errors.New("i is not found in this scope"),
		},
		{
			[]string{"i", "j"},
			&Context{
				Env: &Env{
					Data: &ast.ObjectMap{
						Data: map[string]*ast.Object{
							"i": {
								InstanceFields: &ast.ObjectMap{
									Data: map[string]*ast.Object{
										"j": builtin.NewInteger(2),
									},
								},
							},
						},
					},
				},
			},
			builtin.NewInteger(2),
			nil,
		},
		{
			[]string{"i", "j"},
			&Context{
				Env: &Env{
					Data: &ast.ObjectMap{
						Data: map[string]*ast.Object{
							"i": {
								InstanceFields: &ast.ObjectMap{
									Data: map[string]*ast.Object{},
								},
							},
						},
					},
				},
			},
			nil,
			errors.New("j is not found in this scope"),
		},
		//{
		//	[]string{"i", "j"},
		//	&Context{
		//		Env: &Env{
		//			Data: &ast.ObjectMap{},
		//		},
		//		ClassTypes: &ast.ClassMap{
		//			Data: map[string]*ast.ClassType{
		//				"integer": builtin.IntegerType,
		//				"i": {
		//					StaticFields: &builtin.FieldMap{
		//						Data: map[string]*builtin.Field{
		//							"j": {
		//								Type: &ast.TypeRef{
		//									Name: []string{"Integer"},
		//								},
		//							},
		//						},
		//					},
		//				},
		//			},
		//		},
		//	},
		//	builtin.NewInteger(2),
		//	nil,
		//},
		//{
		//	[]string{"i", "j", "k"},
		//	&Context{
		//		Env: &Env{
		//			Data: &ast.ObjectMap{},
		//		},
		//		ClassTypes: &ast.ClassMap{
		//			Data: map[string]*ast.ClassType{
		//				"integer": builtin.IntegerType,
		//				"foo": {
		//					InstanceFields: &builtin.FieldMap{
		//						Data: map[string]*builtin.Field{
		//							"k": {
		//								Type: &ast.TypeRef{
		//									Name: []string{"Integer"},
		//								},
		//							},
		//						},
		//					},
		//				},
		//				"i": {
		//					StaticFields: &builtin.FieldMap{
		//						Data: map[string]*builtin.Field{
		//							"j": {
		//								Type: &ast.TypeRef{
		//									Name: []string{"foo"},
		//								},
		//							},
		//						},
		//					},
		//				},
		//			},
		//		},
		//	},
		//	builtin.IntegerType,
		//	nil,
		//},
		//{
		//	[]string{"i", "j", "k"},
		//	&Context{
		//		Env: &Env{
		//			Data: &ast.ObjectMap{},
		//		},
		//		ClassTypes: &ast.ClassMap{
		//			Data: map[string]*ast.ClassType{
		//				"integer": builtin.IntegerType,
		//			},
		//		},
		//		NameSpaces: &builtin.NameSpaceStore{
		//			Data: map[string]*ast.ClassMap{
		//				"i": {
		//					Data: map[string]*ast.ClassType{
		//						"j": {
		//							StaticFields: &builtin.FieldMap{
		//								Data: map[string]*builtin.Field{
		//									"k": {
		//										Type: &ast.TypeRef{
		//											Name: []string{"Integer"},
		//										},
		//									},
		//								},
		//							},
		//						},
		//					},
		//				},
		//			},
		//		},
		//	},
		//	builtin.IntegerType,
		//	nil,
		//},
	}

	ignore := cmpopts.IgnoreTypes(func(*ast.Object) string { return "" })

	for _, testCase := range testCases {
		typeResolver := NewTypeResolver(testCase.Context)
		actual, err := typeResolver.ResolveVariable(testCase.Input)
		if testCase.Error != nil && testCase.Error.Error() != err.Error() {
			diff := cmp.Diff(testCase.Error.Error(), err.Error())
			t.Errorf(diff)
		}

		if ok := cmp.Equal(testCase.Expected, actual, ignore); !ok {
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
		Expected *ast.Method
		Error    error
	}{
		{
			[]string{"foo"},
			&Context{
				Env: &Env{
					Data: &ast.ObjectMap{
						Data: map[string]*ast.Object{
							"this": {
								ClassType: &ast.ClassType{
									InstanceMethods: &ast.MethodMap{
										Data: map[string][]*ast.Method{
											"foo": {
												{
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
			&ast.Method{Name: "foo"},
			nil,
		},
		{
			[]string{"foo", "bar"},
			&Context{
				Env: &Env{
					Data: &ast.ObjectMap{
						Data: map[string]*ast.Object{
							"foo": {
								ClassType: &ast.ClassType{
									InstanceMethods: &ast.MethodMap{
										Data: map[string][]*ast.Method{
											"bar": {
												{
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
			},
			&ast.Method{Name: "bar"},
			nil,
		},
		{
			[]string{"klass", "foo"},
			&Context{
				ClassTypes: &ast.ClassMap{
					Data: map[string]*ast.ClassType{
						"klass": {
							StaticMethods: &ast.MethodMap{
								Data: map[string][]*ast.Method{
									"foo": {
										{
											Name: "foo",
										},
									},
								},
							},
						},
					},
				},
				Env: NewEnv(nil),
			},
			&ast.Method{Name: "foo"},
			nil,
		},
		//{
		//	[]string{"klass", "foo", "bar"},
		//	&Context{
		//		ClassTypes: &ast.ClassMap{
		//			Data: map[string]*ast.ClassType{
		//				"klass": {
		//					StaticFields: &builtin.FieldMap{
		//						Data: map[string]*builtin.Field{
		//							"foo": {
		//								Name: "foo",
		//								Type: &ast.TypeRef{
		//									Name: []string{"klass2"},
		//								},
		//							},
		//						},
		//					},
		//				},
		//				"klass2": {
		//					InstanceMethods: &ast.MethodMap{
		//						Data: map[string][]*ast.Method{
		//							"bar": {
		//								&ast.Method{
		//									Name: "bar",
		//								},
		//							},
		//						},
		//					},
		//				},
		//			},
		//		},
		//		Env: newTypeEnv(nil),
		//	},
		//	&ast.Method{Name: "bar"},
		//	nil,
		//},
		//{
		//	[]string{"namespace", "klass", "foo"},
		//	&Context{
		//		Env:        newTypeEnv(nil),
		//		ClassTypes: ast.NewClassMap(),
		//		NameSpaces: &builtin.NameSpaceStore{
		//			Data: map[string]*ast.ClassMap{
		//				"namespace": {
		//					Data: map[string]*ast.ClassType{
		//						"klass": {
		//							StaticMethods: &ast.MethodMap{
		//								Data: map[string][]*ast.Method{
		//									"foo": {
		//										&ast.Method{
		//											Name: "foo",
		//										},
		//									},
		//								},
		//							},
		//						},
		//					},
		//				},
		//			},
		//		},
		//	},
		//	&ast.Method{Name: "foo"},
		//	nil,
		//},
		//{
		//	[]string{"namespace", "klass", "foo", "bar"},
		//	&Context{
		//		Env: newTypeEnv(nil),
		//		ClassTypes: &ast.ClassMap{
		//			Data: map[string]*ast.ClassType{
		//				"klass2": {
		//					InstanceMethods: &ast.MethodMap{
		//						Data: map[string][]*ast.Method{
		//							"bar": {
		//								&ast.Method{
		//									Name: "bar",
		//								},
		//							},
		//						},
		//					},
		//				},
		//			},
		//		},
		//		NameSpaces: &builtin.NameSpaceStore{
		//			Data: map[string]*ast.ClassMap{
		//				"namespace": {
		//					Data: map[string]*ast.ClassType{
		//						"klass": {
		//							StaticFields: &builtin.FieldMap{
		//								Data: map[string]*builtin.Field{
		//									"foo": {
		//										Type: &ast.TypeRef{
		//											Name: []string{"klass2"},
		//										},
		//										Name: "foo",
		//									},
		//								},
		//							},
		//						},
		//					},
		//				},
		//			},
		//		},
		//	},
		//	&ast.Method{Name: "bar"},
		//	nil,
		//},
	}
	for _, testCase := range testCases {
		typeResolver := NewTypeResolver(testCase.Context)
		// TODO: test receiver
		_, actual, err := typeResolver.ResolveMethod(testCase.Input, []*ast.Object{})
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
