package builtin

import (
	"fmt"
	"strings"

	"github.com/tzmfreedom/land/ast"
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
	BooleanType.StaticMethods.Set(
		"valueOf",
		[]*ast.Method{
			ast.CreateMethod(
				"valueOf",
				BooleanType,
				[]*ast.Parameter{stringTypeParameter},
				func(this *ast.Object, params []*ast.Object, extra map[string]interface{}) interface{} {
					return NewBoolean(strings.ToLower(params[0].StringValue()) == "true")
				},
			),
			ast.CreateMethod(
				"valueOf",
				BooleanType,
				[]*ast.Parameter{objectTypeParameter},
				func(this *ast.Object, params []*ast.Object, extra map[string]interface{}) interface{} {
					switch this.ClassType {
					case StringType:
						return NewBoolean(strings.ToLower(params[0].StringValue()) == "true")
					case BooleanType:
						return NewBoolean(this.BoolValue())
					}
					panic("not expected argument")
				},
			),
		},
	)

	primitiveClassMap.Set("Boolean", BooleanType)
}
