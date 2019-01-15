package builtin

import (
	"fmt"

	"github.com/tzmfreedom/goland/ast"
)

var BooleanType = &ClassType{
	Name: "Boolean",
	ToString: func(o *Object) string {
		return fmt.Sprintf("%t", o.Value().(bool))
	},
}

var booleanTypeRef = &ast.TypeRef{
	Name: []string{"Boolean"},
}

func init() {
	primitiveClassMap.Set("Boolean", BooleanType)
}
