package builtin

import (
	"fmt"

	"github.com/tzmfreedom/goland/ast"
)

var BooleanType = &ast.ClassType{
	Name: "Boolean",
	ToString: func(o *ast.Object) string {
		return fmt.Sprintf("%t", o.Value().(bool))
	},
}

var booleanTypeRef = &ast.TypeRef{
	Name: []string{"Boolean"},
}

func init() {
	primitiveClassMap.Set("Boolean", BooleanType)
}
