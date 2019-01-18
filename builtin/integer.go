package builtin

import (
	"fmt"

	"github.com/tzmfreedom/goland/ast"
)

var IntegerType = &ast.ClassType{
	Name: "Integer",
	ToString: func(o *ast.Object) string {
		return fmt.Sprintf("%d", o.Value().(int))
	},
}

var integerTypeRef = &ast.TypeRef{
	Name:       []string{"Integer"},
	Parameters: []*ast.TypeRef{},
}

func init() {
	primitiveClassMap.Set("Integer", IntegerType)
}
