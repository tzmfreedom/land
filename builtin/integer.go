package builtin

import (
	"fmt"

	"github.com/tzmfreedom/goland/ast"
)

var IntegerType = &ast.ClassType{
	Name:            "Integer",
	InstanceMethods: ast.NewMethodMap(),
	StaticMethods:   ast.NewMethodMap(),
	ToString: func(o *ast.Object) string {
		return fmt.Sprintf("%d", o.Value().(int))
	},
}

var IntegerTypeParameter = &ast.Parameter{
	Type: IntegerType,
	Name: "_",
}

var integerTypeRef = &ast.TypeRef{
	Name:       []string{"Integer"},
	Parameters: []*ast.TypeRef{},
}

func init() {
	primitiveClassMap.Set("Integer", IntegerType)
}
