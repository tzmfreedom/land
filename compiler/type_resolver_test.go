package compiler

import (
	"errors"
	"testing"

	"github.com/google/go-cmp/cmp"
	"github.com/google/go-cmp/cmp/cmpopts"
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
			[]string{"local"},
			&Context{
				ClassTypes: &builtin.ClassMap{
					Data: map[string]*builtin.ClassType{
						"integer": builtin.IntegerType,
					},
				},
				Env: &TypeEnv{
					Data: &TypeMap{
						Data: map[string]*builtin.ClassType{
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
			errors.New("local is not found in this scope"),
		},
		{
			[]string{"local", "instance_field"},
			&Context{
				ClassTypes: &builtin.ClassMap{
					Data: map[string]*builtin.ClassType{
						"integer": builtin.IntegerType,
					},
				},
				Env: &TypeEnv{
					Data: &TypeMap{
						Data: map[string]*builtin.ClassType{
							"local": {
								InstanceFields: &builtin.FieldMap{
									Data: map[string]*builtin.Field{
										"instance_field": {
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
			[]string{"local", "instance_field"},
			&Context{
				ClassTypes: &builtin.ClassMap{
					Data: map[string]*builtin.ClassType{
						"integer": builtin.IntegerType,
					},
				},
				Env: &TypeEnv{
					Data: &TypeMap{
						Data: map[string]*builtin.ClassType{
							"local": {
								InstanceFields: &builtin.FieldMap{
									Data: map[string]*builtin.Field{},
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
				ClassTypes: &builtin.ClassMap{
					Data: map[string]*builtin.ClassType{
						"integer": builtin.IntegerType,
						"class": {
							StaticFields: &builtin.FieldMap{
								Data: map[string]*builtin.Field{
									"static_field": {
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
			[]string{"class", "static_field", "instance_field"},
			&Context{
				Env: &TypeEnv{
					Data: &TypeMap{},
				},
				ClassTypes: &builtin.ClassMap{
					Data: map[string]*builtin.ClassType{
						"integer": builtin.IntegerType,
						"class2": {
							InstanceFields: &builtin.FieldMap{
								Data: map[string]*builtin.Field{
									"instance_field": {
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
						"class": {
							StaticFields: &builtin.FieldMap{
								Data: map[string]*builtin.Field{
									"static_field": {
										Type: &ast.TypeRef{
											Name: []string{"class2"},
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
			[]string{"namespace", "class", "static_field"},
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
						"namespace": {
							Data: map[string]*builtin.ClassType{
								"class": {
									StaticFields: &builtin.FieldMap{
										Data: map[string]*builtin.Field{
											"static_field": {
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
		// protected/private
		{
			[]string{"local", "instance_field_protected"},
			&Context{
				ClassTypes: &builtin.ClassMap{
					Data: map[string]*builtin.ClassType{
						"integer": builtin.IntegerType,
					},
				},
				Env: &TypeEnv{
					Data: &TypeMap{
						Data: map[string]*builtin.ClassType{
							"local": {
								InstanceFields: &builtin.FieldMap{
									Data: map[string]*builtin.Field{
										"instance_field_protected": {
											Type: &ast.TypeRef{
												Name: []string{"Integer"},
											},
											Modifiers: []ast.Node{
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
				ClassTypes: &builtin.ClassMap{
					Data: map[string]*builtin.ClassType{
						"integer": builtin.IntegerType,
					},
				},
				Env: &TypeEnv{
					Data: &TypeMap{
						Data: map[string]*builtin.ClassType{
							"local": {
								InstanceFields: &builtin.FieldMap{
									Data: map[string]*builtin.Field{
										"instance_field_private": {
											Type: &ast.TypeRef{
												Name: []string{"Integer"},
											},
											Modifiers: []ast.Node{
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
				ClassTypes: &builtin.ClassMap{
					Data: map[string]*builtin.ClassType{
						"integer": builtin.IntegerType,
						"class": {
							StaticFields: &builtin.FieldMap{
								Data: map[string]*builtin.Field{
									"static_field_protected": {
										Type: &ast.TypeRef{
											Name: []string{"Integer"},
										},
										Modifiers: []ast.Node{
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
				ClassTypes: &builtin.ClassMap{
					Data: map[string]*builtin.ClassType{
						"integer": builtin.IntegerType,
						"class": {
							StaticFields: &builtin.FieldMap{
								Data: map[string]*builtin.Field{
									"static_field_private": {
										Type: &ast.TypeRef{
											Name: []string{"Integer"},
										},
										Modifiers: []ast.Node{
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
		//		ClassTypes: &builtin.ClassMap{
		//			Data: map[string]*builtin.ClassType{
		//				"integer": builtin.IntegerType,
		//				"class2": {
		//					InstanceFields: &builtin.FieldMap{
		//						Data: map[string]*builtin.Field{
		//							"instance_field": {
		//								Type: &ast.TypeRef{
		//									Name: []string{"Integer"},
		//								},
		//								Modifiers: []ast.Node{
		//									&ast.Modifier{
		//										Name: "public",
		//									},
		//								},
		//							},
		//						},
		//					},
		//				},
		//				"class": {
		//					StaticFields: &builtin.FieldMap{
		//						Data: map[string]*builtin.Field{
		//							"static_field": {
		//								Type: &ast.TypeRef{
		//									Name: []string{"class2"},
		//								},
		//								Modifiers: []ast.Node{
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
		//		ClassTypes: &builtin.ClassMap{
		//			Data: map[string]*builtin.ClassType{
		//				"integer": builtin.IntegerType,
		//				"class2": {
		//					InstanceFields: &builtin.FieldMap{
		//						Data: map[string]*builtin.Field{
		//							"instance_field": {
		//								Type: &ast.TypeRef{
		//									Name: []string{"Integer"},
		//								},
		//								Modifiers: []ast.Node{
		//									&ast.Modifier{
		//										Name: "public",
		//									},
		//								},
		//							},
		//						},
		//					},
		//				},
		//				"class": {
		//					StaticFields: &builtin.FieldMap{
		//						Data: map[string]*builtin.Field{
		//							"static_field": {
		//								Type: &ast.TypeRef{
		//									Name: []string{"class2"},
		//								},
		//								Modifiers: []ast.Node{
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
				ClassTypes: &builtin.ClassMap{
					Data: map[string]*builtin.ClassType{
						"integer": builtin.IntegerType,
						"class2": {
							InstanceFields: &builtin.FieldMap{
								Data: map[string]*builtin.Field{
									"instance_field_protected": {
										Type: &ast.TypeRef{
											Name: []string{"Integer"},
										},
										Modifiers: []ast.Node{
											&ast.Modifier{
												Name: "protected",
											},
										},
									},
								},
							},
						},
						"class": {
							StaticFields: &builtin.FieldMap{
								Data: map[string]*builtin.Field{
									"static_field": {
										Type: &ast.TypeRef{
											Name: []string{"class2"},
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
			nil,
			errors.New("Field access modifier must be public but protected"),
		},
		{
			[]string{"class", "static_field", "instance_field_private"},
			&Context{
				Env: &TypeEnv{
					Data: &TypeMap{},
				},
				ClassTypes: &builtin.ClassMap{
					Data: map[string]*builtin.ClassType{
						"integer": builtin.IntegerType,
						"class2": {
							InstanceFields: &builtin.FieldMap{
								Data: map[string]*builtin.Field{
									"instance_field_private": {
										Type: &ast.TypeRef{
											Name: []string{"Integer"},
										},
										Modifiers: []ast.Node{
											&ast.Modifier{
												Name: "private",
											},
										},
									},
								},
							},
						},
						"class": {
							StaticFields: &builtin.FieldMap{
								Data: map[string]*builtin.Field{
									"static_field": {
										Type: &ast.TypeRef{
											Name: []string{"class2"},
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
			nil,
			errors.New("Field access modifier must be public but private"),
		},
		{
			[]string{"namespace", "class", "static_field_protected"},
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
						"namespace": {
							Data: map[string]*builtin.ClassType{
								"class": {
									StaticFields: &builtin.FieldMap{
										Data: map[string]*builtin.Field{
											"static_field_protected": {
												Type: &ast.TypeRef{
													Name: []string{"Integer"},
												},
												Modifiers: []ast.Node{
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
				ClassTypes: &builtin.ClassMap{
					Data: map[string]*builtin.ClassType{
						"integer": builtin.IntegerType,
					},
				},
				NameSpaces: &builtin.NameSpaceStore{
					Data: map[string]*builtin.ClassMap{
						"namespace": {
							Data: map[string]*builtin.ClassType{
								"class": {
									StaticFields: &builtin.FieldMap{
										Data: map[string]*builtin.Field{
											"static_field_private": {
												Type: &ast.TypeRef{
													Name: []string{"Integer"},
												},
												Modifiers: []ast.Node{
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
			},
			nil,
			errors.New("Field access modifier must be public but private"),
		},
	}
	ignore := cmpopts.IgnoreTypes(func(*builtin.Object) string { return "" })

	for _, testCase := range testCases {
		typeResolver := &TypeResolver{Context: testCase.Context}
		actual, _, err := typeResolver.ResolveVariable(testCase.Input, false)
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
		Expected *builtin.Method
		Error    error
	}{
		{
			[]string{"instance"},
			&Context{
				Env: &TypeEnv{
					Data: &TypeMap{
						Data: map[string]*builtin.ClassType{
							"this": {
								Name: "class",
								InstanceMethods: &builtin.MethodMap{
									Data: map[string][]*builtin.Method{
										"instance": {
											{
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
			&builtin.Method{
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
								Name: "class",
								InstanceMethods: &builtin.MethodMap{
									Data: map[string][]*builtin.Method{
										"instance": {
											{
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
			&builtin.Method{
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
			[]string{"class", "static"},
			&Context{
				ClassTypes: &builtin.ClassMap{
					Data: map[string]*builtin.ClassType{
						"class": {
							StaticMethods: &builtin.MethodMap{
								Data: map[string][]*builtin.Method{
									"static": {
										{
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
			&builtin.Method{
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
			[]string{"class", "static", "instance"},
			&Context{
				ClassTypes: &builtin.ClassMap{
					Data: map[string]*builtin.ClassType{
						"class": {
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
											Name: []string{"class2"},
										},
									},
								},
							},
						},
						"class2": {
							InstanceMethods: &builtin.MethodMap{
								Data: map[string][]*builtin.Method{
									"instance": {
										{
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
			&builtin.Method{
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
			[]string{"namespace", "class", "static"},
			&Context{
				Env:        newTypeEnv(nil),
				ClassTypes: builtin.NewClassMap(),
				NameSpaces: &builtin.NameSpaceStore{
					Data: map[string]*builtin.ClassMap{
						"namespace": {
							Data: map[string]*builtin.ClassType{
								"class": {
									StaticMethods: &builtin.MethodMap{
										Data: map[string][]*builtin.Method{
											"static": {
												{
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
			&builtin.Method{
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
			[]string{"namespace", "class", "static", "instance"},
			&Context{
				Env: newTypeEnv(nil),
				ClassTypes: &builtin.ClassMap{
					Data: map[string]*builtin.ClassType{
						"class2": {
							InstanceMethods: &builtin.MethodMap{
								Data: map[string][]*builtin.Method{
									"instance": {
										{
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
								"class": {
									StaticFields: &builtin.FieldMap{
										Data: map[string]*builtin.Field{
											"static": {
												Type: &ast.TypeRef{
													Name: []string{"class2"},
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
			&builtin.Method{
				Name: "instance",
				Modifiers: []ast.Node{
					&ast.Modifier{
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
						Data: map[string]*builtin.ClassType{
							"local": {
								Name: "class",
								InstanceMethods: &builtin.MethodMap{
									Data: map[string][]*builtin.Method{
										"instance": {
											{
												Name: "instance",
												Modifiers: []ast.Node{
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
			},
			nil,
			errors.New("Method access modifier must be public but protected"),
		},
		{
			[]string{"local", "instance"},
			&Context{
				Env: &TypeEnv{
					Data: &TypeMap{
						Data: map[string]*builtin.ClassType{
							"local": {
								Name: "class",
								InstanceMethods: &builtin.MethodMap{
									Data: map[string][]*builtin.Method{
										"instance": {
											{
												Name: "instance",
												Modifiers: []ast.Node{
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
			},
			nil,
			errors.New("Method access modifier must be public but private"),
		},
		{
			[]string{"class", "static"},
			&Context{
				ClassTypes: &builtin.ClassMap{
					Data: map[string]*builtin.ClassType{
						"class": {
							StaticMethods: &builtin.MethodMap{
								Data: map[string][]*builtin.Method{
									"static": {
										{
											Name: "static",
											Modifiers: []ast.Node{
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
				Env: newTypeEnv(nil),
			},
			nil,
			errors.New("Method access modifier must be public but protected"),
		},
		{
			[]string{"class", "static"},
			&Context{
				ClassTypes: &builtin.ClassMap{
					Data: map[string]*builtin.ClassType{
						"class": {
							StaticMethods: &builtin.MethodMap{
								Data: map[string][]*builtin.Method{
									"static": {
										{
											Name: "static",
											Modifiers: []ast.Node{
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
				Env: newTypeEnv(nil),
			},
			nil,
			errors.New("Method access modifier must be public but private"),
		},
		//{
		//	[]string{"class", "static", "instance"},
		//	&Context{
		//		ClassTypes: &builtin.ClassMap{
		//			Data: map[string]*builtin.ClassType{
		//				"class": {
		//					StaticFields: &builtin.FieldMap{
		//						Data: map[string]*builtin.Field{
		//							"static": {
		//								Name: "static",
		//								Modifiers: []ast.Node{
		//									&ast.Modifier{
		//										Name: "public",
		//									},
		//								},
		//								Type: &ast.TypeRef{
		//									Name: []string{"class2"},
		//								},
		//							},
		//						},
		//					},
		//				},
		//				"class2": {
		//					InstanceMethods: &builtin.MethodMap{
		//						Data: map[string][]*builtin.Method{
		//							"instance": {
		//								{
		//									Name: "instance",
		//									Modifiers: []ast.Node{
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
		//		Modifiers: []ast.Node{
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
		//		ClassTypes: &builtin.ClassMap{
		//			Data: map[string]*builtin.ClassType{
		//				"class": {
		//					StaticFields: &builtin.FieldMap{
		//						Data: map[string]*builtin.Field{
		//							"static": {
		//								Name: "static",
		//								Modifiers: []ast.Node{
		//									&ast.Modifier{
		//										Name: "public",
		//									},
		//								},
		//								Type: &ast.TypeRef{
		//									Name: []string{"class2"},
		//								},
		//							},
		//						},
		//					},
		//				},
		//				"class2": {
		//					InstanceMethods: &builtin.MethodMap{
		//						Data: map[string][]*builtin.Method{
		//							"instance": {
		//								{
		//									Name: "instance",
		//									Modifiers: []ast.Node{
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
		//		Modifiers: []ast.Node{
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
				ClassTypes: builtin.NewClassMap(),
				NameSpaces: &builtin.NameSpaceStore{
					Data: map[string]*builtin.ClassMap{
						"namespace": {
							Data: map[string]*builtin.ClassType{
								"class": {
									StaticMethods: &builtin.MethodMap{
										Data: map[string][]*builtin.Method{
											"static": {
												{
													Name: "static",
													Modifiers: []ast.Node{
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
				},
			},
			nil,
			errors.New("Method access modifier must be public but protected"),
		},
		{
			[]string{"namespace", "class", "static"},
			&Context{
				Env:        newTypeEnv(nil),
				ClassTypes: builtin.NewClassMap(),
				NameSpaces: &builtin.NameSpaceStore{
					Data: map[string]*builtin.ClassMap{
						"namespace": {
							Data: map[string]*builtin.ClassType{
								"class": {
									StaticMethods: &builtin.MethodMap{
										Data: map[string][]*builtin.Method{
											"static": {
												{
													Name: "static",
													Modifiers: []ast.Node{
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
				},
			},
			nil,
			errors.New("Method access modifier must be public but private"),
		},
		{
			[]string{"namespace", "class", "static", "instance"},
			&Context{
				Env: newTypeEnv(nil),
				ClassTypes: &builtin.ClassMap{
					Data: map[string]*builtin.ClassType{
						"class2": {
							InstanceMethods: &builtin.MethodMap{
								Data: map[string][]*builtin.Method{
									"instance": {
										{
											Name: "instance",
											Modifiers: []ast.Node{
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
				NameSpaces: &builtin.NameSpaceStore{
					Data: map[string]*builtin.ClassMap{
						"namespace": {
							Data: map[string]*builtin.ClassType{
								"class": {
									StaticFields: &builtin.FieldMap{
										Data: map[string]*builtin.Field{
											"static": {
												Type: &ast.TypeRef{
													Name: []string{"class2"},
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
			nil,
			errors.New("Method access modifier must be public but protected"),
		},
		{
			[]string{"namespace", "class", "static", "instance"},
			&Context{
				Env: newTypeEnv(nil),
				ClassTypes: &builtin.ClassMap{
					Data: map[string]*builtin.ClassType{
						"class2": {
							InstanceMethods: &builtin.MethodMap{
								Data: map[string][]*builtin.Method{
									"instance": {
										{
											Name: "instance",
											Modifiers: []ast.Node{
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
				NameSpaces: &builtin.NameSpaceStore{
					Data: map[string]*builtin.ClassMap{
						"namespace": {
							Data: map[string]*builtin.ClassType{
								"class": {
									StaticFields: &builtin.FieldMap{
										Data: map[string]*builtin.Field{
											"static": {
												Type: &ast.TypeRef{
													Name: []string{"class2"},
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
			nil,
			errors.New("Method access modifier must be public but private"),
		},
	}
	for _, testCase := range testCases {
		typeResolver := &TypeResolver{Context: testCase.Context}
		_, actual, err := typeResolver.ResolveMethod(testCase.Input, []*builtin.ClassType{})
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
