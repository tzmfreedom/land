package builtin

import "github.com/tzmfreedom/goland/ast"

func init() {
	instanceMethods := NewMethodMap()
	staticMethods := NewMethodMap()
	staticMethods.Set(
		"currentPage",
		[]*Method{
			CreateMethod(
				"currentPage",
				[]string{"PageReference"},
				[]ast.Node{},
				func(this *Object, params []*Object, extra map[string]interface{}) interface{} {
					return extra["current_page"]
				},
			),
		},
	)

	classType := CreateClass(
		"ApexPages",
		[]*Method{},
		instanceMethods,
		staticMethods,
	)

	primitiveClassMap.Set("ApexPages", classType)

	instanceMethods = NewMethodMap()
	instanceMethods.Set(
		"getRecord",
		[]*Method{
			CreateMethod(
				"getRecord",
				[]string{"Account"}, // TODO: SObject
				[]ast.Node{},
				func(this *Object, params []*Object, extra map[string]interface{}) interface{} {
					return this.Extra["record"]
				},
			),
		},
	)
	instanceMethods.Set(
		"getId",
		[]*Method{
			CreateMethod(
				"getId",
				[]string{"String"}, // TODO: SObject
				[]ast.Node{},
				func(this *Object, params []*Object, extra map[string]interface{}) interface{} {
					record := this.Extra["record"].(*Object)
					value, _ := record.InstanceFields.Get("Id")
					return value
				},
			),
		},
	)

	staticMethods = NewMethodMap()
	classType = CreateClass(
		"StandardController",
		[]*Method{
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
				NativeFunction: func(this *Object, params []*Object, extra map[string]interface{}) interface{} {
					this.Extra["record"] = params[0]
					return nil
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
