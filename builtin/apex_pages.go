package builtin

import "github.com/tzmfreedom/goland/ast"

func init() {
	instanceMethods := NewMethodMap()
	staticMethods := NewMethodMap()
	staticMethods.Set(
		"currentPage",
		[]ast.Node{
			CreateMethod(
				"currentPage",
				[]string{"PageReference"},
				[]ast.Node{},
				func(this interface{}, params []interface{}, extra map[string]interface{}) interface{} {
					return extra["current_page"]
				},
			),
		},
	)

	classType := CreateClass(
		"ApexPages",
		[]*ast.ConstructorDeclaration{},
		instanceMethods,
		staticMethods,
	)

	primitiveClassMap.Set("ApexPages", classType)

	instanceMethods = NewMethodMap()
	instanceMethods.Set(
		"getRecord",
		[]ast.Node{
			CreateMethod(
				"getRecord",
				[]string{"Account"}, // TODO: SObject
				[]ast.Node{},
				func(this interface{}, params []interface{}, extra map[string]interface{}) interface{} {
					thisObj := this.(*Object)
					return thisObj.Extra["record"]
				},
			),
		},
	)
	instanceMethods.Set(
		"getId",
		[]ast.Node{
			CreateMethod(
				"getId",
				[]string{"String"}, // TODO: SObject
				[]ast.Node{},
				func(this interface{}, params []interface{}, extra map[string]interface{}) interface{} {
					thisObj := this.(*Object)
					record := thisObj.Extra["record"].(*Object)
					value, _ := record.InstanceFields.Get("Id")
					return value
				},
			),
		},
	)

	staticMethods = NewMethodMap()
	classType = CreateClass(
		"StandardController",
		[]*ast.ConstructorDeclaration{
			{
				Modifiers: []ast.Node{
					&ast.Modifier{
						Name: "public",
					},
				},
				Parameters: []ast.Node{
					&ast.Parameter{
						Type: &ast.TypeRef{
							Name:       []string{"Account"}, // TODO: Sobject
							Parameters: []ast.Node{},
						},
						Name: "_",
					},
				},
				NativeFunction: func(this interface{}, params []interface{}) {
					thisObj := this.(*Object)
					thisObj.Extra["record"] = params[0]
				},
			},
		},
		instanceMethods,
		staticMethods,
	)

	classMap := NewClassMap()
	classMap.Set("StandardController", classType)

	nameSpaceStore.Set("ApexPages", classMap)
}
