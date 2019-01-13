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
					pageRef := CreateObject(pageReferenceType)
					pageRef.Extra["url"] = extra["current_page"].(string)
					pageRef.Extra["parameters"] = extra["parameters"]
					return pageRef
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
}
