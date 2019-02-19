package compiler

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
		Expected *ast.ClassType
		Error    error
	}{
		{
			[]string{"local"},
			&Context{
				ClassTypes: &ast.ClassMap{
					Data: map[string]*ast.ClassType{
						"integer": builtin.IntegerType,
					},
				},
				Env: &TypeEnv{
					Data: &TypeMap{
						Data: map[string]*ast.ClassType{
							"local": builtin.IntegerType,
						},
					},
				},
			},
			builtin.IntegerType,
			nil,
		},
		{
			[]string{"local"},
			&Context{
				ClassTypes: &ast.ClassMap{
					Data: map[string]*ast.ClassType{
						"integer": builtin.IntegerType,
					},
				},
				Env: &TypeEnv{
					Data: &TypeMap{
						Data: map[string]*ast.ClassType{},
					},
				},
			},
			nil,
			errors.New("local is not found in this scope"),
		},
		{
			[]string{"local", "instance_field"},
			&Context{
				ClassTypes: &ast.ClassMap{
					Data: map[string]*ast.ClassType{
						"integer": builtin.IntegerType,
					},
				},
				Env: &TypeEnv{
					Data: &TypeMap{
						Data: map[string]*ast.ClassType{
							"local": {
								InstanceFields: &ast.FieldMap{
									Data: map[string]*ast.Field{
										"instance_field": {
											Type: builtin.IntegerType,
											Modifiers: []*ast.Modifier{
												{
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
			[]string{"local", "instance_field"},
			&Context{
				ClassTypes: &ast.ClassMap{
					Data: map[string]*ast.ClassType{
						"integer": builtin.IntegerType,
					},
				},
				Env: &TypeEnv{
					Data: &TypeMap{
						Data: map[string]*ast.ClassType{
							"local": {
								InstanceFields: &ast.FieldMap{
									Data: map[string]*ast.Field{},
								},
							},
						},
					},
				},
			},
			nil,
			errors.New("Field instance_field is not found"),
		},
		{
			[]string{"class", "static_field"},
			&Context{
				Env: &TypeEnv{
					Data: &TypeMap{},
				},
				ClassTypes: &ast.ClassMap{
					Data: map[string]*ast.ClassType{
						"integer": builtin.IntegerType,
						"class": {
							StaticFields: &ast.FieldMap{
								Data: map[string]*ast.Field{
									"static_field": {
										Type: builtin.IntegerType,
										Modifiers: []*ast.Modifier{
											{
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
			[]string{"class", "static_field", "instance_field"},
			&Context{
				Env: &TypeEnv{
					Data: &TypeMap{},
				},
				ClassTypes: &ast.ClassMap{
					Data: func() map[string]*ast.ClassType {
						classTypeMap := map[string]*ast.ClassType{
							"integer": builtin.IntegerType,
							"class2": {
								InstanceFields: &ast.FieldMap{
									Data: map[string]*ast.Field{
										"instance_field": {
											Type: builtin.IntegerType,
											Modifiers: []*ast.Modifier{
												{
													Name: "public",
												},
											},
										},
									},
								},
							},
						}
						classTypeMap["class"] = &ast.ClassType{
							StaticFields: &ast.FieldMap{
								Data: map[string]*ast.Field{
									"static_field": {
										Type: classTypeMap["class2"],
										Modifiers: []*ast.Modifier{
											{
												Name: "public",
											},
										},
									},
								},
							},
						}
						return classTypeMap
					}(),
				},
			},
			builtin.IntegerType,
			nil,
		},
		{
			[]string{"namespace", "class", "static_field"},
			&Context{
				Env: &TypeEnv{
					Data: &TypeMap{},
				},
				ClassTypes: &ast.ClassMap{
					Data: map[string]*ast.ClassType{
						"integer": builtin.IntegerType,
					},
				},
				NameSpaces: &builtin.NameSpaceStore{
					Data: map[string]*ast.ClassMap{
						"namespace": {
							Data: map[string]*ast.ClassType{
								"class": {
									StaticFields: &ast.FieldMap{
										Data: map[string]*ast.Field{
											"static_field": {
												Type: builtin.IntegerType,
												Modifiers: []*ast.Modifier{
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
		// protected/private
		{
			[]string{"local", "instance_field_protected"},
			&Context{
				ClassTypes: &ast.ClassMap{
					Data: map[string]*ast.ClassType{
						"integer": builtin.IntegerType,
					},
				},
				Env: &TypeEnv{
					Data: &TypeMap{
						Data: map[string]*ast.ClassType{
							"local": {
								InstanceFields: &ast.FieldMap{
									Data: map[string]*ast.Field{
										"instance_field_protected": {
											Type: builtin.IntegerType,
											Modifiers: []*ast.Modifier{
												&ast.Modifier{
													Name: "protected",
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
			nil,
			errors.New("Field access modifier must be public but protected"),
		},
		{
			[]string{"local", "instance_field_private"},
			&Context{
				ClassTypes: &ast.ClassMap{
					Data: map[string]*ast.ClassType{
						"integer": builtin.IntegerType,
					},
				},
				Env: &TypeEnv{
					Data: &TypeMap{
						Data: map[string]*ast.ClassType{
							"local": {
								InstanceFields: &ast.FieldMap{
									Data: map[string]*ast.Field{
										"instance_field_private": {
											Type: builtin.IntegerType,
											Modifiers: []*ast.Modifier{
												&ast.Modifier{
													Name: "private",
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
			nil,
			errors.New("Field access modifier must be public but private"),
		},
		{
			[]string{"class", "static_field_protected"},
			&Context{
				Env: &TypeEnv{
					Data: &TypeMap{},
				},
				ClassTypes: &ast.ClassMap{
					Data: map[string]*ast.ClassType{
						"integer": builtin.IntegerType,
						"class": {
							StaticFields: &ast.FieldMap{
								Data: map[string]*ast.Field{
									"static_field_protected": {
										Type: builtin.IntegerType,
										Modifiers: []*ast.Modifier{
											&ast.Modifier{
												Name: "protected",
											},
										},
									},
								},
							},
						},
					},
				},
			},
			nil,
			errors.New("Field access modifier must be public but protected"),
		},
		{
			[]string{"class", "static_field_private"},
			&Context{
				Env: &TypeEnv{
					Data: &TypeMap{},
				},
				ClassTypes: &ast.ClassMap{
					Data: map[string]*ast.ClassType{
						"integer": builtin.IntegerType,
						"class": {
							StaticFields: &ast.FieldMap{
								Data: map[string]*ast.Field{
									"static_field_private": {
										Type: builtin.IntegerType,
										Modifiers: []*ast.Modifier{
											&ast.Modifier{
												Name: "private",
											},
										},
									},
								},
							},
						},
					},
				},
			},
			nil,
			errors.New("Field access modifier must be public but private"),
		},
		//{
		//	[]string{"class", "static_field_protected", "instance_field"},
		//	&Context{
		//		Env: &TypeEnv{
		//			Data: &TypeMap{},
		//		},
		//		ClassTypes: &ast.ClassMap{
		//			Data: map[string]*ast.ClassType{
		//				"integer": builtin.IntegerType,
		//				"class2": {
		//					InstanceFields: &ast.FieldMap{
		//						Data: map[string]*ast.Field{
		//							"instance_field": {
		//								TypeRef: &ast.TypeRef{
		//									Name: []string{"Integer"},
		//								},
		//								Modifiers: []*ast.Modifier{
		//									&ast.Modifier{
		//										Name: "public",
		//									},
		//								},
		//							},
		//						},
		//					},
		//				},
		//				"class": {
		//					StaticFields: &ast.FieldMap{
		//						Data: map[string]*ast.Field{
		//							"static_field": {
		//								TypeRef: &ast.TypeRef{
		//									Name: []string{"class2"},
		//								},
		//								Modifiers: []*ast.Modifier{
		//									&ast.Modifier{
		//										Name: "protected",
		//									},
		//								},
		//							},
		//						},
		//					},
		//				},
		//			},
		//		},
		//	},
		//	nil,
		//	nil,
		//},
		//{
		//	[]string{"class", "static_field_private", "instance_field"},
		//	&Context{
		//		Env: &TypeEnv{
		//			Data: &TypeMap{},
		//		},
		//		ClassTypes: &ast.ClassMap{
		//			Data: map[string]*ast.ClassType{
		//				"integer": builtin.IntegerType,
		//				"class2": {
		//					InstanceFields: &ast.FieldMap{
		//						Data: map[string]*ast.Field{
		//							"instance_field": {
		//								TypeRef: &ast.TypeRef{
		//									Name: []string{"Integer"},
		//								},
		//								Modifiers: []*ast.Modifier{
		//									&ast.Modifier{
		//										Name: "public",
		//									},
		//								},
		//							},
		//						},
		//					},
		//				},
		//				"class": {
		//					StaticFields: &ast.FieldMap{
		//						Data: map[string]*ast.Field{
		//							"static_field": {
		//								TypeRef: &ast.TypeRef{
		//									Name: []string{"class2"},
		//								},
		//								Modifiers: []*ast.Modifier{
		//									&ast.Modifier{
		//										Name: "private",
		//									},
		//								},
		//							},
		//						},
		//					},
		//				},
		//			},
		//		},
		//	},
		//	nil,
		//	nil,
		//},
		{
			[]string{"class", "static_field", "instance_field_protected"},
			&Context{
				Env: &TypeEnv{
					Data: &TypeMap{},
				},
				ClassTypes: &ast.ClassMap{
					Data: func() map[string]*ast.ClassType {
						classTypeMap := map[string]*ast.ClassType{
							"integer": builtin.IntegerType,
							"class2": {
								InstanceFields: &ast.FieldMap{
									Data: map[string]*ast.Field{
										"instance_field_protected": {
											Type: builtin.IntegerType,
											Modifiers: []*ast.Modifier{
												{
													Name: "protected",
												},
											},
										},
									},
								},
							},
						}
						classTypeMap["class"] = &ast.ClassType{
							StaticFields: &ast.FieldMap{
								Data: map[string]*ast.Field{
									"static_field": {
										Type: classTypeMap["class2"],
										Modifiers: []*ast.Modifier{
											{
												Name: "public",
											},
										},
									},
								},
							},
						}
						return classTypeMap
					}(),
				},
			},
			nil,
			errors.New("Field access modifier must be public but protected"),
		},
		{
			[]string{"class", "static_field", "instance_field_private"},
			&Context{
				Env: &TypeEnv{
					Data: &TypeMap{},
				},
				ClassTypes: &ast.ClassMap{
					Data: func() map[string]*ast.ClassType {
						classTypeMap := map[string]*ast.ClassType{
							"integer": builtin.IntegerType,
							"class2": {
								InstanceFields: &ast.FieldMap{
									Data: map[string]*ast.Field{
										"instance_field_private": {
											Type: builtin.IntegerType,
											Modifiers: []*ast.Modifier{
												{
													Name: "private",
												},
											},
										},
									},
								},
							},
						}
						classTypeMap["class"] = &ast.ClassType{
							StaticFields: &ast.FieldMap{
								Data: map[string]*ast.Field{
									"static_field": {
										Type: classTypeMap["class2"],
										Modifiers: []*ast.Modifier{
											{
												Name: "public",
											},
										},
									},
								},
							},
						}
						return classTypeMap
					}(),
				},
			},
			nil,
			errors.New("Field access modifier must be public but private"),
		},
		{
			[]string{"namespace", "class", "static_field_protected"},
			&Context{
				Env: &TypeEnv{
					Data: &TypeMap{},
				},
				ClassTypes: &ast.ClassMap{
					Data: map[string]*ast.ClassType{
						"integer": builtin.IntegerType,
					},
				},
				NameSpaces: &builtin.NameSpaceStore{
					Data: map[string]*ast.ClassMap{
						"namespace": {
							Data: map[string]*ast.ClassType{
								"class": {
									StaticFields: &ast.FieldMap{
										Data: map[string]*ast.Field{
											"static_field_protected": {
												Type: builtin.IntegerType,
												Modifiers: []*ast.Modifier{
													{
														Name: "protected",
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
			nil,
			errors.New("Field access modifier must be public but protected"),
		},
		{
			[]string{"namespace", "class", "static_field_private"},
			&Context{
				Env: &TypeEnv{
					Data: &TypeMap{},
				},
				ClassTypes: &ast.ClassMap{
					Data: map[string]*ast.ClassType{
						"integer": builtin.IntegerType,
					},
				},
				NameSpaces: &builtin.NameSpaceStore{
					Data: map[string]*ast.ClassMap{
						"namespace": {
							Data: map[string]*ast.ClassType{
								"class": {
									StaticFields: &ast.FieldMap{
										Data: map[string]*ast.Field{
											"static_field_private": {
												Type: builtin.IntegerType,
												Modifiers: []*ast.Modifier{
													{
														Name: "private",
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
			nil,
			errors.New("Field access modifier must be public but private"),
		},
	}
	ignore := cmpopts.IgnoreTypes(func(*ast.Object) string { return "" })

	for _, testCase := range testCases {
		typeResolver := NewTypeResolver(testCase.Context)
		actual, err := typeResolver.ResolveVariable(testCase.Input, false)
		if testCase.Error == nil && err != nil {
			diff := cmp.Diff(testCase.Error, err.Error())
			t.Errorf(diff)
		} else if testCase.Error != nil && err == nil {
			diff := cmp.Diff(testCase.Error.Error(), err)
			t.Errorf(diff)
		} else if testCase.Error != nil && err != nil &&
			testCase.Error.Error() != err.Error() {
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

func TestResolveType(t *testing.T) {
	testCases := []struct {
		Input    []string
		Context  *Context
		Expected *ast.ClassType
		Error    error
	}{
		{
			[]string{"i"},
			&Context{
				ClassTypes: &ast.ClassMap{
					Data: map[string]*ast.ClassType{
						"i": {
							Name: "i",
						},
					},
				},
			},
			&ast.ClassType{Name: "i"},
			nil,
		},
		{
			[]string{"i"},
			&Context{
				ClassTypes: ast.NewClassMap(),
				NameSpaces: &builtin.NameSpaceStore{
					Data: map[string]*ast.ClassMap{
						"system": {
							Data: map[string]*ast.ClassType{
								"i": {
									Name: "i",
								},
							},
						},
					},
				},
			},
			&ast.ClassType{Name: "i"},
			nil,
		},
		{
			[]string{"i", "j"},
			&Context{
				ClassTypes: &ast.ClassMap{
					Data: map[string]*ast.ClassType{
						"i": {
							Name: "i",
							InnerClasses: &ast.ClassMap{
								Data: map[string]*ast.ClassType{
									"j": {
										Name: "j",
									},
								},
							},
						},
					},
				},
			},
			&ast.ClassType{Name: "j"},
			nil,
		},
		{
			[]string{"i", "j"},
			&Context{
				ClassTypes: ast.NewClassMap(),
				NameSpaces: &builtin.NameSpaceStore{
					Data: map[string]*ast.ClassMap{
						"i": {
							Data: map[string]*ast.ClassType{
								"j": {
									Name: "j",
								},
							},
						},
					},
				},
			},
			&ast.ClassType{Name: "j"},
			nil,
		},
		{
			[]string{"i", "j", "k"},
			&Context{
				ClassTypes: ast.NewClassMap(),
				NameSpaces: &builtin.NameSpaceStore{
					Data: map[string]*ast.ClassMap{
						"i": {
							Data: map[string]*ast.ClassType{
								"j": {
									Name: "j",
									InnerClasses: &ast.ClassMap{
										Data: map[string]*ast.ClassType{
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
			&ast.ClassType{Name: "k"},
			nil,
		},
	}
	for _, testCase := range testCases {
		typeResolver := NewTypeResolver(testCase.Context)
		actual, err := typeResolver.ResolveType(testCase.Input)
		if testCase.Error == nil && err != nil {
			diff := cmp.Diff(testCase.Error, err.Error())
			t.Errorf(diff)
		} else if testCase.Error != nil && err == nil {
			diff := cmp.Diff(testCase.Error.Error(), err)
			t.Errorf(diff)
		} else if testCase.Error != nil && err != nil &&
			testCase.Error.Error() != err.Error() {
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
		Expected *ast.Method
		Error    error
	}{
		{
			[]string{"instance"},
			&Context{
				Env: &TypeEnv{
					Data: &TypeMap{
						Data: map[string]*ast.ClassType{
							"this": {
								Name: "class",
								InstanceMethods: &ast.MethodMap{
									Data: map[string][]*ast.Method{
										"instance": {
											{
												Name: "instance",
												Modifiers: []*ast.Modifier{
													{
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
			&ast.Method{
				Name: "instance",
				Modifiers: []*ast.Modifier{
					{
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
						Data: map[string]*ast.ClassType{
							"local": {
								Name: "class",
								InstanceMethods: &ast.MethodMap{
									Data: map[string][]*ast.Method{
										"instance": {
											{
												Name: "instance",
												Modifiers: []*ast.Modifier{
													{
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
			&ast.Method{
				Name: "instance",
				Modifiers: []*ast.Modifier{
					{
						Name: "public",
					},
				},
			},
			nil,
		},
		{
			[]string{"class", "static"},
			&Context{
				ClassTypes: &ast.ClassMap{
					Data: map[string]*ast.ClassType{
						"class": {
							StaticMethods: &ast.MethodMap{
								Data: map[string][]*ast.Method{
									"static": {
										{
											Name: "static",
											Modifiers: []*ast.Modifier{
												{
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
			&ast.Method{
				Name: "static",
				Modifiers: []*ast.Modifier{
					{
						Name: "public",
					},
				},
			},
			nil,
		},
		{
			[]string{"class", "static", "instance"},
			&Context{
				ClassTypes: &ast.ClassMap{
					Data: func() map[string]*ast.ClassType {
						classTypeMap := map[string]*ast.ClassType{
							"class2": {
								InstanceMethods: &ast.MethodMap{
									Data: map[string][]*ast.Method{
										"instance": {
											{
												Name: "instance",
												Modifiers: []*ast.Modifier{
													{
														Name: "public",
													},
												},
											},
										},
									},
								},
							},
						}
						classTypeMap["class"] = &ast.ClassType{
							StaticFields: &ast.FieldMap{
								Data: map[string]*ast.Field{
									"static": {
										Name: "static",
										Modifiers: []*ast.Modifier{
											{
												Name: "public",
											},
										},
										Type: classTypeMap["class2"],
									},
								},
							},
						}
						return classTypeMap
					}(),
				},
				Env: newTypeEnv(nil),
			},
			&ast.Method{
				Name: "instance",
				Modifiers: []*ast.Modifier{
					{
						Name: "public",
					},
				},
			},
			nil,
		},
		{
			[]string{"namespace", "class", "static"},
			&Context{
				Env:        newTypeEnv(nil),
				ClassTypes: ast.NewClassMap(),
				NameSpaces: &builtin.NameSpaceStore{
					Data: map[string]*ast.ClassMap{
						"namespace": {
							Data: map[string]*ast.ClassType{
								"class": {
									StaticMethods: &ast.MethodMap{
										Data: map[string][]*ast.Method{
											"static": {
												{
													Name: "static",
													Modifiers: []*ast.Modifier{
														{
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
			&ast.Method{
				Name: "static",
				Modifiers: []*ast.Modifier{
					{
						Name: "public",
					},
				},
			},
			nil,
		},
		{
			[]string{"namespace", "class", "static", "instance"},
			func() *Context {
				context := &Context{
					Env: newTypeEnv(nil),
					ClassTypes: &ast.ClassMap{
						Data: map[string]*ast.ClassType{
							"class2": {
								InstanceMethods: &ast.MethodMap{
									Data: map[string][]*ast.Method{
										"instance": {
											{
												Name: "instance",
												Modifiers: []*ast.Modifier{
													{
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
				}
				context.NameSpaces = &builtin.NameSpaceStore{
					Data: map[string]*ast.ClassMap{
						"namespace": {
							Data: map[string]*ast.ClassType{
								"class": {
									StaticFields: &ast.FieldMap{
										Data: map[string]*ast.Field{
											"static": {
												Type: context.ClassTypes.Data["class2"],
												Modifiers: []*ast.Modifier{
													{
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
				}
				return context
			}(),
			&ast.Method{
				Name: "instance",
				Modifiers: []*ast.Modifier{
					{
						Name: "public",
					},
				},
			},
			nil,
		},
		// private/protected
		{
			[]string{"local", "instance"},
			&Context{
				Env: &TypeEnv{
					Data: &TypeMap{
						Data: map[string]*ast.ClassType{
							"local": {
								Name: "class",
								InstanceMethods: &ast.MethodMap{
									Data: map[string][]*ast.Method{
										"instance": {
											{
												Name: "instance",
												Modifiers: []*ast.Modifier{
													{
														Name: "protected",
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
			nil,
			errors.New("Method access modifier must be public but protected"),
		},
		{
			[]string{"local", "instance"},
			&Context{
				Env: &TypeEnv{
					Data: &TypeMap{
						Data: map[string]*ast.ClassType{
							"local": {
								Name: "class",
								InstanceMethods: &ast.MethodMap{
									Data: map[string][]*ast.Method{
										"instance": {
											{
												Name: "instance",
												Modifiers: []*ast.Modifier{
													{
														Name: "private",
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
			nil,
			errors.New("Method access modifier must be public but private"),
		},
		{
			[]string{"class", "static"},
			&Context{
				ClassTypes: &ast.ClassMap{
					Data: map[string]*ast.ClassType{
						"class": {
							StaticMethods: &ast.MethodMap{
								Data: map[string][]*ast.Method{
									"static": {
										{
											Name: "static",
											Modifiers: []*ast.Modifier{
												{
													Name: "protected",
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
			nil,
			errors.New("Method access modifier must be public but protected"),
		},
		{
			[]string{"class", "static"},
			&Context{
				ClassTypes: &ast.ClassMap{
					Data: map[string]*ast.ClassType{
						"class": {
							StaticMethods: &ast.MethodMap{
								Data: map[string][]*ast.Method{
									"static": {
										{
											Name: "static",
											Modifiers: []*ast.Modifier{
												{
													Name: "private",
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
			nil,
			errors.New("Method access modifier must be public but private"),
		},
		//{
		//	[]string{"class", "static", "instance"},
		//	&Context{
		//		ClassTypes: &ast.ClassMap{
		//			Data: map[string]*ast.ClassType{
		//				"class": {
		//					StaticFields: &ast.FieldMap{
		//						Data: map[string]*ast.Field{
		//							"static": {
		//								Name: "static",
		//								Modifiers: []*ast.Modifier{
		//									&ast.Modifier{
		//										Name: "public",
		//									},
		//								},
		//								TypeRef: &ast.TypeRef{
		//									Name: []string{"class2"},
		//								},
		//							},
		//						},
		//					},
		//				},
		//				"class2": {
		//					InstanceMethods: &ast.MethodMap{
		//						Data: map[string][]*ast.Method{
		//							"instance": {
		//								{
		//									Name: "instance",
		//									Modifiers: []*ast.Modifier{
		//										&ast.Modifier{
		//											Name: "public",
		//										},
		//									},
		//								},
		//							},
		//						},
		//					},
		//				},
		//			},
		//		},
		//		Env: newTypeEnv(nil),
		//	},
		//	{
		//		Name: "instance",
		//		Modifiers: []*ast.Modifier{
		//			&ast.Modifier{
		//				Name: "public",
		//			},
		//		},
		//	},
		//	nil,
		//},
		//{
		//	[]string{"class", "static", "instance"},
		//	&Context{
		//		ClassTypes: &ast.ClassMap{
		//			Data: map[string]*ast.ClassType{
		//				"class": {
		//					StaticFields: &ast.FieldMap{
		//						Data: map[string]*ast.Field{
		//							"static": {
		//								Name: "static",
		//								Modifiers: []*ast.Modifier{
		//									&ast.Modifier{
		//										Name: "public",
		//									},
		//								},
		//								TypeRef: &ast.TypeRef{
		//									Name: []string{"class2"},
		//								},
		//							},
		//						},
		//					},
		//				},
		//				"class2": {
		//					InstanceMethods: &ast.MethodMap{
		//						Data: map[string][]*ast.Method{
		//							"instance": {
		//								{
		//									Name: "instance",
		//									Modifiers: []*ast.Modifier{
		//										&ast.Modifier{
		//											Name: "public",
		//										},
		//									},
		//								},
		//							},
		//						},
		//					},
		//				},
		//			},
		//		},
		//		Env: newTypeEnv(nil),
		//	},
		//	{
		//		Name: "instance",
		//		Modifiers: []*ast.Modifier{
		//			&ast.Modifier{
		//				Name: "public",
		//			},
		//		},
		//	},
		//	nil,
		//},
		{
			[]string{"namespace", "class", "static"},
			&Context{
				Env:        newTypeEnv(nil),
				ClassTypes: ast.NewClassMap(),
				NameSpaces: &builtin.NameSpaceStore{
					Data: map[string]*ast.ClassMap{
						"namespace": {
							Data: map[string]*ast.ClassType{
								"class": {
									StaticMethods: &ast.MethodMap{
										Data: map[string][]*ast.Method{
											"static": {
												{
													Name: "static",
													Modifiers: []*ast.Modifier{
														{
															Name: "protected",
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
			nil,
			errors.New("Method access modifier must be public but protected"),
		},
		{
			[]string{"namespace", "class", "static"},
			&Context{
				Env:        newTypeEnv(nil),
				ClassTypes: ast.NewClassMap(),
				NameSpaces: &builtin.NameSpaceStore{
					Data: map[string]*ast.ClassMap{
						"namespace": {
							Data: map[string]*ast.ClassType{
								"class": {
									StaticMethods: &ast.MethodMap{
										Data: map[string][]*ast.Method{
											"static": {
												{
													Name: "static",
													Modifiers: []*ast.Modifier{
														{
															Name: "private",
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
			nil,
			errors.New("Method access modifier must be public but private"),
		},
		{
			[]string{"namespace", "class", "static", "instance"},
			func() *Context {
				context := &Context{
					Env: newTypeEnv(nil),
					ClassTypes: &ast.ClassMap{
						Data: map[string]*ast.ClassType{
							"class2": {
								InstanceMethods: &ast.MethodMap{
									Data: map[string][]*ast.Method{
										"instance": {
											{
												Name: "instance",
												Modifiers: []*ast.Modifier{
													{
														Name: "protected",
													},
												},
											},
										},
									},
								},
							},
						},
					},
				}
				context.NameSpaces = &builtin.NameSpaceStore{
					Data: map[string]*ast.ClassMap{
						"namespace": {
							Data: map[string]*ast.ClassType{
								"class": {
									StaticFields: &ast.FieldMap{
										Data: map[string]*ast.Field{
											"static": {
												Type: context.ClassTypes.Data["class2"],
												Modifiers: []*ast.Modifier{
													{
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
				}
				return context
			}(),
			nil,
			errors.New("Method access modifier must be public but protected"),
		},
		{
			[]string{"namespace", "class", "static", "instance"},
			func() *Context {
				context := &Context{
					Env: newTypeEnv(nil),
					ClassTypes: &ast.ClassMap{
						Data: map[string]*ast.ClassType{
							"class2": {
								InstanceMethods: &ast.MethodMap{
									Data: map[string][]*ast.Method{
										"instance": {
											{
												Name: "instance",
												Modifiers: []*ast.Modifier{
													{
														Name: "private",
													},
												},
											},
										},
									},
								},
							},
						},
					},
				}
				context.NameSpaces = &builtin.NameSpaceStore{
					Data: map[string]*ast.ClassMap{
						"namespace": {
							Data: map[string]*ast.ClassType{
								"class": {
									StaticFields: &ast.FieldMap{
										Data: map[string]*ast.Field{
											"static": {
												Type: context.ClassTypes.Data["class2"],
												Modifiers: []*ast.Modifier{
													{
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
				}
				return context
			}(),
			nil,
			errors.New("Method access modifier must be public but private"),
		},
	}
	for _, testCase := range testCases {
		typeResolver := NewTypeResolver(testCase.Context)
		_, actual, err := typeResolver.ResolveMethod(testCase.Input, []*ast.ClassType{})
		if testCase.Error == nil && err != nil {
			diff := cmp.Diff(testCase.Error, err.Error())
			t.Errorf(diff)
		} else if testCase.Error != nil && err == nil {
			diff := cmp.Diff(testCase.Error.Error(), err)
			t.Errorf(diff)
		} else if testCase.Error != nil && err != nil &&
			testCase.Error.Error() != err.Error() {
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
