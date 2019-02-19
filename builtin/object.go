package builtin

import (
	"github.com/tzmfreedom/land/ast"
)

var ObjectType = ast.CreateClass(
	"Object",
	[]*ast.Method{},
	nil,
	nil,
)

var objectTypeParameter = &ast.Parameter{
	Type: ObjectType,
	Name: "_",
}

func init() {
	primitiveClassMap.Set("Object", ObjectType)
}
