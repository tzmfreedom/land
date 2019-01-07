package builtin

import "github.com/tzmfreedom/goland/ast"

var StringType = createStringType()

func createStringType() *ClassType {
	instanceMethods := NewMethodMap()
	staticMethods := NewMethodMap()
	instanceMethods.Set("valueOf", []ast.Node{
		&ast.MethodDeclaration{
			Name: "size",
		},
	})
	return CreateClass(
		"String",
		nil,
		instanceMethods,
		staticMethods,
	)
}

func init() {
	primitiveClassMap.Set("String", StringType)
}
