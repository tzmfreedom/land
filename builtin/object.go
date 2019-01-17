package builtin

import (
	"github.com/tzmfreedom/goland/ast"
)

var ObjectType = ast.CreateClass(
	"Object",
	[]*ast.Method{},
	nil,
	nil,
)

var objectTypeParameter = &ast.Parameter{
	TypeRef: &ast.TypeRef{
		Name:       []string{"Object"},
		Parameters: []ast.Node{},
	},
	Name: "_",
}

func init() {
	primitiveClassMap.Set("Object", ObjectType)
}
