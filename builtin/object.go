package builtin

import (
	"github.com/tzmfreedom/goland/ast"
)

var ObjectType = CreateClass(
	"Object",
	[]*ast.ConstructorDeclaration{},
	nil,
	nil,
)

var objectTypeParameter = &ast.Parameter{
	Type: &ast.TypeRef{
		Name:       []string{"Object"},
		Parameters: []ast.Node{},
	},
	Name: "_",
}

func init() {
	primitiveClassMap.Set("Object", ObjectType)
}
