package builtin

import (
	"fmt"

	"github.com/tzmfreedom/goland/ast"
)

var DoubleType = &ast.ClassType{
	Name: "Double",
	ToString: func(o *ast.Object) string {
		return fmt.Sprintf("%f", o.Value().(float64))
	},
}

func init() {
	primitiveClassMap.Set("Double", DoubleType)
}
