package builtin

import "github.com/tzmfreedom/land/ast"

var testType = createTestType()

func createTestType() *ast.ClassType {
	instanceMethods := ast.NewMethodMap()
	staticMethods := ast.NewMethodMap()
	staticMethods.Set(
		"setCurrentPage",
		[]*ast.Method{
			ast.CreateMethod(
				"setCurrentPage",
				pageReferenceType,
				[]*ast.Parameter{pageReferenceParameter},
				func(this *ast.Object, params []*ast.Object, extra map[string]interface{}) interface{} {
					extra["current_page"] = params[0]
					return nil
				},
			),
		},
	)

	classType := ast.CreateClass(
		"Test",
		[]*ast.Method{},
		instanceMethods,
		staticMethods,
	)
	return classType
}

func init() {
	primitiveClassMap.Set("Test", testType)
}
