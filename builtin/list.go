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
		[]*Method{
			CreateMethod(
				"add",
				nil,
				[]ast.Node{t1Parameter},
				func(this *Object, params []*Object, extra map[string]interface{}) interface{} {
					records := this.Extra["records"].([]*Object)
					listElement := params[0]
					records = append(records, listElement)
					return nil
				},
			),
		},
	)
	instanceMethods.Set(
		"size",
		[]*Method{
			CreateMethod(
				"size",
				[]string{"Integer"},
				[]ast.Node{},
				func(this *Object, params []*Object, extra map[string]interface{}) interface{} {
					return NewInteger(len(this.Extra["records"].([]*Object)))
				},
			),
		},
	)

	classType := CreateClass(
		"List",
		[]*Method{
			{
				Modifiers: []ast.Node{
					&ast.Modifier{
						Name: "public",
					},
				},
				Parameters: []ast.Node{},
				NativeFunction: func(this *Object, params []*Object, extra map[string]interface{}) interface{} {
					return nil
				},
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
				NativeFunction: func(this *Object, params []*Object, extra map[string]interface{}) interface{} {
					listObj := params[0]
					listParams := listObj.Extra["records"].([]*Object)
					newListParams := make([]*Object, len(listParams))
					for i, listParam := range listParams {
						newListParams[i] = &Object{
							ClassType: listParam.ClassType,
						}
					}
					this.Extra = map[string]interface{}{
						"records": newListParams,
					}
					return nil
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
