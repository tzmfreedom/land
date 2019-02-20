package builtin

import (
	"fmt"
	"math"
	"strconv"

	"github.com/dustin/go-humanize"
	"github.com/tzmfreedom/land/ast"
)

var DoubleType = &ast.ClassType{
	Name:            "Double",
	InstanceMethods: ast.NewMethodMap(),
	StaticMethods:   ast.NewMethodMap(),
	ToString: func(o *ast.Object) string {
		return fmt.Sprintf("%f", o.Value().(float64))
	},
}

var doubleTypeParameter = &ast.Parameter{
	Type: DoubleType,
	Name: "_",
}

func init() {
	DoubleType.InstanceMethods.Set(
		"format",
		[]*ast.Method{
			ast.CreateMethod(
				"format",
				StringType,
				[]*ast.Parameter{},
				func(this *ast.Object, params []*ast.Object, extra map[string]interface{}) interface{} {
					value := this.DoubleValue()
					return NewString(humanize.Commaf(value))
				},
			),
		},
	)
	DoubleType.InstanceMethods.Set(
		"intValue",
		[]*ast.Method{
			ast.CreateMethod(
				"intValue",
				IntegerType,
				[]*ast.Parameter{},
				func(this *ast.Object, params []*ast.Object, extra map[string]interface{}) interface{} {
					value := this.DoubleValue()
					return NewInteger(int(value))
				},
			),
		},
	)
	DoubleType.InstanceMethods.Set(
		"longValue",
		[]*ast.Method{
			ast.CreateMethod(
				"longValue",
				LongType,
				[]*ast.Parameter{},
				func(this *ast.Object, params []*ast.Object, extra map[string]interface{}) interface{} {
					value := this.DoubleValue()
					return NewLong(int(value))
				},
			),
		},
	)
	DoubleType.InstanceMethods.Set(
		"round",
		[]*ast.Method{
			ast.CreateMethod(
				"round",
				LongType,
				[]*ast.Parameter{},
				func(this *ast.Object, params []*ast.Object, extra map[string]interface{}) interface{} {
					value := this.DoubleValue()
					return NewLong(int(math.Round(value)))
				},
			),
		},
	)
	DoubleType.StaticMethods.Set(
		"valueOf",
		[]*ast.Method{
			ast.CreateMethod(
				"valueOf",
				StringType,
				[]*ast.Parameter{},
				func(this *ast.Object, params []*ast.Object, extra map[string]interface{}) interface{} {
					value, err := strconv.ParseFloat(this.StringValue(), 64)
					if err != nil {
						panic(err)
					}
					return NewDouble(value)
				},
			),
			ast.CreateMethod(
				"valueOf",
				ObjectType,
				[]*ast.Parameter{},
				func(this *ast.Object, params []*ast.Object, extra map[string]interface{}) interface{} {
					if this.ClassType == DoubleType {
						return this
					}
					panic("no expected type")
				},
			),
		},
	)

	primitiveClassMap.Set("Double", DoubleType)
}
