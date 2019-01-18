package builtin

import "github.com/tzmfreedom/goland/ast"

func init() {
	instanceMethods := ast.NewMethodMap()
	staticMethods := ast.NewMethodMap()
	staticMethods.Set(
		"currentPage",
		[]*ast.Method{
			ast.CreateMethod(
				"currentPage",
				pageReferenceType,
				[]*ast.Parameter{},
				func(this *ast.Object, params []*ast.Object, extra map[string]interface{}) interface{} {
					return extra["current_page"]
				},
			),
		},
	)

	classType := ast.CreateClass(
		"ApexPages",
		[]*ast.Method{},
		instanceMethods,
		staticMethods,
	)

	primitiveClassMap.Set("ApexPages", classType)

	account, _ := PrimitiveClassMap().Get("Account")
	instanceMethods = ast.NewMethodMap()
	instanceMethods.Set(
		"getRecord",
		[]*ast.Method{
			ast.CreateMethod(
				"getRecord",
				account, // TODO: SObject
				[]*ast.Parameter{},
				func(this *ast.Object, params []*ast.Object, extra map[string]interface{}) interface{} {
					return this.Extra["record"]
				},
			),
		},
	)
	instanceMethods.Set(
		"getId",
		[]*ast.Method{
			ast.CreateMethod(
				"getId",
				StringType, // TODO: SObject
				[]*ast.Parameter{},
				func(this *ast.Object, params []*ast.Object, extra map[string]interface{}) interface{} {
					record := this.Extra["record"].(*ast.Object)
					value, _ := record.InstanceFields.Get("Id")
					return value
				},
			),
		},
	)
	accountType, _ := PrimitiveClassMap().Get("Account")

	staticMethods = ast.NewMethodMap()
	classType = ast.CreateClass(
		"StandardController",
		[]*ast.Method{
			{
				Modifiers: []*ast.Modifier{ast.PublicModifier()},
				Parameters: []*ast.Parameter{
					{
						Type: accountType, // TODO: sobject
						Name: "_",
					},
				},
				NativeFunction: func(this *ast.Object, params []*ast.Object, extra map[string]interface{}) interface{} {
					this.Extra["record"] = params[0]
					return nil
				},
			},
		},
		instanceMethods,
		staticMethods,
	)

	classMap := ast.NewClassMap()
	classMap.Set("StandardController", classType)

	nameSpaceStore.Set("ApexPages", classMap)
}
