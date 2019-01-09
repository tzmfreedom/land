package builtin

import (
	"github.com/tzmfreedom/goland/ast"
)

var ListType = createListType()

func createListType() *ClassType {
	instanceMethods := NewMethodMap()
	instanceMethods.Set(
		"add",
		[]ast.Node{
			CreateMethod(
				"add",
				nil,
				func(this interface{}, params []interface{}, options ...interface{}) interface{} {
					thisObj := this.(*Object)
					records := thisObj.Extra["records"].([]*Object)
					listElement := params[0].(*Object)
					records = append(records, listElement)
					return nil
				},
			),
		},
	)
	instanceMethods.Set(
		"size",
		[]ast.Node{
			CreateMethod(
				"size",
				[]string{"Integer"},
				func(this interface{}, params []interface{}, options ...interface{}) interface{} {
					thisObj := this.(*Object)
					return NewInteger(len(thisObj.Extra["records"].([]*Object)))
				},
			),
		},
	)
	instanceMethods.Set(
		"isNext",
		[]ast.Node{
			CreateMethod(
				"next",
				[]string{"Boolean"},
				func(this interface{}, params []interface{}, options ...interface{}) interface{} {
					thisObj := this.(*Object)
					counter := thisObj.Extra["counter"].(int)
					return thisObj.Extra["records"].([]*Object)[counter]
				},
			),
		},
	)
	instanceMethods.Set(
		"next",
		[]ast.Node{
			CreateMethod(
				"next",
				[]string{"T:1"},
				func(this interface{}, params []interface{}, options ...interface{}) interface{} {
					thisObj := this.(*Object)
					counter := thisObj.Extra["counter"].(int)
					return thisObj.Extra["records"].([]*Object)[counter]
				},
			),
		},
	)

	return CreateClass(
		"List",
		[]*ast.ConstructorDeclaration{
			{
				Modifiers: []ast.Node{
					&ast.Modifier{
						Name: "public",
					},
				},
				Parameters:     []ast.Node{},
				NativeFunction: func(params []interface{}) {},
			},
			{
				Modifiers: []ast.Node{
					&ast.Modifier{
						Name: "public",
					},
				},
				Parameters: []ast.Node{
					&ast.Parameter{
						Type: &ast.TypeRef{
							Name: []string{"List"},
							Parameters: []ast.Node{
								&ast.TypeRef{
									Name: []string{"T:1"},
								},
							},
						},
						Name: "list",
					},
				},
				NativeFunction: func(params []interface{}) {
					listObj := params[1].(*Object)
					listParams := listObj.Extra["records"].([]*Object)
					newListParams := make([]*Object, len(listParams))
					for i, listParam := range listParams {
						newListParams[i] = &Object{
							ClassType: listParam.ClassType,
						}
					}
					this := params[0].(*Object)
					this.Extra = map[string]interface{}{
						"records": newListParams,
					}
				},
			},
		},
		instanceMethods,
		nil,
	)
}

func init() {
	primitiveClassMap.Set("list", ListType)
}
