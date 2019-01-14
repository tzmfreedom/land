package builtin

import "github.com/tzmfreedom/goland/ast"

var testType = createTestType()

func createTestType() *ClassType {
	instanceMethods := NewMethodMap()
	staticMethods := NewMethodMap()
	staticMethods.Set(
		"setCurrentPage",
		[]ast.Node{
			CreateMethod(
				"setCurrentPage",
				[]string{"PageReference"},
				[]ast.Node{pageReferenceParameter},
				func(this interface{}, params []interface{}, extra map[string]interface{}) interface{} {
					extra["current_page"] = params[0]
					return nil
				},
			),
		},
	)

	classType := CreateClass(
		"Test",
		[]*ast.ConstructorDeclaration{},
		instanceMethods,
		staticMethods,
	)
	return classType
}

func init() {
	primitiveClassMap.Set("Test", testType)
}
