package interpreter

import (
	"errors"
	"testing"

	"github.com/google/go-cmp/cmp"
	"github.com/google/go-cmp/cmp/cmpopts"
	"github.com/k0kubun/pp"
	"github.com/tzmfreedom/goland/builtin"
)

func TestResolveVariable(t *testing.T) {
	testCases := []struct {
		Input    []string
		Context  *Context
		Expected *builtin.Object
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
				Env: &Env{
					Data: &builtin.ObjectMap{
						Data: map[string]*builtin.Object{
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
				ClassTypes: &builtin.ClassMap{
					Data: map[string]*builtin.ClassType{
						"integer": builtin.IntegerType,
					},
				},
				Env: &Env{
					Data: &builtin.ObjectMap{
						Data: map[string]*builtin.Object{},
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
					Data: &builtin.ObjectMap{
						Data: map[string]*builtin.Object{
							"i": {
								InstanceFields: &builtin.ObjectMap{
									Data: map[string]*builtin.Object{
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
					Data: &builtin.ObjectMap{
						Data: map[string]*builtin.Object{
							"i": {
								InstanceFields: &builtin.ObjectMap{
									Data: map[string]*builtin.Object{},
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
		//			Data: &builtin.ObjectMap{},
		//		},
		//		ClassTypes: &builtin.ClassMap{
		//			Data: map[string]*builtin.ClassType{
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
		//			Data: &builtin.ObjectMap{},
		//		},
		//		ClassTypes: &builtin.ClassMap{
		//			Data: map[string]*builtin.ClassType{
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
		//			Data: &builtin.ObjectMap{},
		//		},
		//		ClassTypes: &builtin.ClassMap{
		//			Data: map[string]*builtin.ClassType{
		//				"integer": builtin.IntegerType,
		//			},
		//		},
		//		NameSpaces: &builtin.NameSpaceStore{
		//			Data: map[string]*builtin.ClassMap{
		//				"i": {
		//					Data: map[string]*builtin.ClassType{
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

	ignore := cmpopts.IgnoreTypes(func(*builtin.Object) string { return "" })

	for _, testCase := range testCases {
		typeResolver := &TypeResolver{Context: testCase.Context}
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
		Expected *builtin.Method
		Error    error
	}{
		{
			[]string{"foo"},
			&Context{
				Env: &Env{
					Data: &builtin.ObjectMap{
						Data: map[string]*builtin.Object{
							"this": {
								ClassType: &builtin.ClassType{
									InstanceMethods: &builtin.MethodMap{
										Data: map[string][]*builtin.Method{
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
			&builtin.Method{Name: "foo"},
			nil,
		},
		{
			[]string{"foo", "bar"},
			&Context{
				Env: &Env{
					Data: &builtin.ObjectMap{
						Data: map[string]*builtin.Object{
							"foo": {
								ClassType: &builtin.ClassType{
									InstanceMethods: &builtin.MethodMap{
										Data: map[string][]*builtin.Method{
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
			&builtin.Method{Name: "bar"},
			nil,
		},
		{
			[]string{"klass", "foo"},
			&Context{
				ClassTypes: &builtin.ClassMap{
					Data: map[string]*builtin.ClassType{
						"klass": {
							StaticMethods: &builtin.MethodMap{
								Data: map[string][]*builtin.Method{
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
			&builtin.Method{Name: "foo"},
			nil,
		},
		//{
		//	[]string{"klass", "foo", "bar"},
		//	&Context{
		//		ClassTypes: &builtin.ClassMap{
		//			Data: map[string]*builtin.ClassType{
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
		//					InstanceMethods: &builtin.MethodMap{
		//						Data: map[string][]*builtin.Method{
		//							"bar": {
		//								&builtin.Method{
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
		//	&builtin.Method{Name: "bar"},
		//	nil,
		//},
		//{
		//	[]string{"namespace", "klass", "foo"},
		//	&Context{
		//		Env:        newTypeEnv(nil),
		//		ClassTypes: builtin.NewClassMap(),
		//		NameSpaces: &builtin.NameSpaceStore{
		//			Data: map[string]*builtin.ClassMap{
		//				"namespace": {
		//					Data: map[string]*builtin.ClassType{
		//						"klass": {
		//							StaticMethods: &builtin.MethodMap{
		//								Data: map[string][]*builtin.Method{
		//									"foo": {
		//										&builtin.Method{
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
		//	&builtin.Method{Name: "foo"},
		//	nil,
		//},
		//{
		//	[]string{"namespace", "klass", "foo", "bar"},
		//	&Context{
		//		Env: newTypeEnv(nil),
		//		ClassTypes: &builtin.ClassMap{
		//			Data: map[string]*builtin.ClassType{
		//				"klass2": {
		//					InstanceMethods: &builtin.MethodMap{
		//						Data: map[string][]*builtin.Method{
		//							"bar": {
		//								&builtin.Method{
		//									Name: "bar",
		//								},
		//							},
		//						},
		//					},
		//				},
		//			},
		//		},
		//		NameSpaces: &builtin.NameSpaceStore{
		//			Data: map[string]*builtin.ClassMap{
		//				"namespace": {
		//					Data: map[string]*builtin.ClassType{
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
		//	&builtin.Method{Name: "bar"},
		//	nil,
		//},
	}
	for _, testCase := range testCases {
		typeResolver := &TypeResolver{Context: testCase.Context}
		// TODO: test receiver
		_, actual, err := typeResolver.ResolveMethod(testCase.Input, []*builtin.Object{})
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
