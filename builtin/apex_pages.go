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
				pageReferenceTypeRef,
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

	instanceMethods = ast.NewMethodMap()
	instanceMethods.Set(
		"getRecord",
		[]*ast.Method{
			ast.CreateMethod(
				"getRecord",
				&ast.TypeRef{Name: []string{"Account"}}, // TODO: SObject
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
				stringTypeRef, // TODO: SObject
				[]*ast.Parameter{},
				func(this *ast.Object, params []*ast.Object, extra map[string]interface{}) interface{} {
					record := this.Extra["record"].(*ast.Object)
					value, _ := record.InstanceFields.Get("Id")
					return value
				},
			),
		},
	)

	staticMethods = ast.NewMethodMap()
	classType = ast.CreateClass(
		"StandardController",
		[]*ast.Method{
			{
				Modifiers: []*ast.Modifier{ast.PublicModifier()},
				Parameters: []*ast.Parameter{
					{
						TypeRef: &ast.TypeRef{
							Name:       []string{"Account"}, // TODO: Sobject
							Parameters: []ast.Node{},
						},
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
