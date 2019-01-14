package builtin

import (
	"github.com/tzmfreedom/goland/ast"
)

var ListType = createListType()

var t1Parameter = &ast.Parameter{
	Type: &ast.TypeRef{
		Name:       []string{"T:1"},
		Parameters: []ast.Node{},
	},
	Name: "_",
}

var t2Parameter = &ast.Parameter{
	Type: &ast.TypeRef{
		Name:       []string{"T:2"},
		Parameters: []ast.Node{},
	},
	Name: "_",
}

func createListType() *ClassType {
	instanceMethods := NewMethodMap()
	instanceMethods.Set(
		"add",
		[]ast.Node{
			CreateMethod(
				"add",
				nil,
				[]ast.Node{t1Parameter},
				func(this interface{}, params []interface{}, extra map[string]interface{}) interface{} {
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
				[]ast.Node{},
				func(this interface{}, params []interface{}, extra map[string]interface{}) interface{} {
					thisObj := this.(*Object)
					return NewInteger(len(thisObj.Extra["records"].([]*Object)))
				},
			),
		},
	)

	classType := CreateClass(
		"List",
		[]*ast.ConstructorDeclaration{
			{
				Modifiers: []ast.Node{
					&ast.Modifier{
						Name: "public",
					},
				},
				Parameters:     []ast.Node{},
				NativeFunction: func(this interface{}, params []interface{}) {},
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
				NativeFunction: func(this interface{}, params []interface{}) {
					listObj := params[0].(*Object)
					listParams := listObj.Extra["records"].([]*Object)
					newListParams := make([]*Object, len(listParams))
					for i, listParam := range listParams {
						newListParams[i] = &Object{
							ClassType: listParam.ClassType,
						}
					}
					thisObj := this.(*Object)
					thisObj.Extra = map[string]interface{}{
						"records": newListParams,
					}
				},
			},
		},
		instanceMethods,
		nil,
	)
	return classType
}

func init() {
	primitiveClassMap.Set("list", ListType)
}
