package builtin

import (
	"fmt"

	"github.com/tzmfreedom/goland/ast"
)

var BooleanType = &ast.ClassType{
	Name:            "Boolean",
	InstanceMethods: ast.NewMethodMap(),
	StaticMethods:   ast.NewMethodMap(),
	ToString: func(o *ast.Object) string {
		return fmt.Sprintf("%t", o.Value().(bool))
	},
}

var booleanTypeParameter = &ast.Parameter{
	Type: BooleanType,
	Name: "_",
}

func init() {
	primitiveClassMap.Set("Boolean", BooleanType)
}
