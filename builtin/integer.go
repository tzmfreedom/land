package builtin

import (
	"fmt"
	"github.com/dustin/go-humanize"
	"github.com/tzmfreedom/goland/ast"
	"strconv"
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
	IntegerType.InstanceMethods.Set(
		"format",
		[]*ast.Method{
			ast.CreateMethod(
				"format",
				StringType,
				[]*ast.Parameter{},
				func(this *ast.Object, params []*ast.Object, extra map[string]interface{}) interface{} {
					return NewString(humanize.Comma(int64(this.IntegerValue())))
				},
			),
		},
	)
	IntegerType.StaticMethods.Set(
		"valueOf",
		[]*ast.Method{
			ast.CreateMethod(
				"valueOf",
				IntegerType,
				[]*ast.Parameter{stringTypeParameter},
				func(this *ast.Object, params []*ast.Object, extra map[string]interface{}) interface{} {
					value, err := strconv.Atoi(params[0].StringValue())
					if err != nil {
						panic(err)
					}
					return NewInteger(value)
				},
			),
			ast.CreateMethod(
				"valueOf",
				IntegerType,
				[]*ast.Parameter{objectTypeParameter},
				func(this *ast.Object, params []*ast.Object, extra map[string]interface{}) interface{} {
					switch this.ClassType {
					case IntegerType:
						return NewInteger(this.IntegerValue())
					}
					panic("not expected type")
				},
			),
		},
	)

	primitiveClassMap.Set("Integer", IntegerType)
}
