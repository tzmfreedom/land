package builtin

import (
	"time"

	"github.com/tzmfreedom/land/ast"
)

var datetimeTypeRef = &ast.TypeRef{
	Name:       []string{"Datetime"},
	Parameters: []*ast.TypeRef{},
}

var DatetimeType = ast.CreateClass(
	"Datetime",
	[]*ast.Method{},
	ast.NewMethodMap(),
	ast.NewMethodMap(),
)

func init() {
	DatetimeType.InstanceMethods.Set(
		"year",
		[]*ast.Method{
			ast.CreateMethod(
				"year",
				IntegerType,
				[]*ast.Parameter{},
				func(this *ast.Object, params []*ast.Object, extra map[string]interface{}) interface{} {
					tm := this.Extra["value"].(time.Time)
					return NewInteger(tm.Year())
				},
			),
		},
	)

	DatetimeType.InstanceMethods.Set(
		"month",
		[]*ast.Method{
			ast.CreateMethod(
				"month",
				IntegerType,
				[]*ast.Parameter{},
				func(this *ast.Object, params []*ast.Object, extra map[string]interface{}) interface{} {
					tm := this.Extra["value"].(time.Time)
					return NewInteger(int(tm.Month()))
				},
			),
		},
	)

	DatetimeType.InstanceMethods.Set(
		"day",
		[]*ast.Method{
			ast.CreateMethod(
				"day",
				IntegerType,
				[]*ast.Parameter{},
				func(this *ast.Object, params []*ast.Object, extra map[string]interface{}) interface{} {
					tm := this.Extra["value"].(time.Time)
					return NewInteger(tm.Day())
				},
			),
		},
	)

	DatetimeType.InstanceMethods.Set(
		"hour",
		[]*ast.Method{
			ast.CreateMethod(
				"hour",
				IntegerType,
				[]*ast.Parameter{},
				func(this *ast.Object, params []*ast.Object, extra map[string]interface{}) interface{} {
					tm := this.Extra["value"].(time.Time)
					return NewInteger(tm.Hour())
				},
			),
		},
	)

	DatetimeType.InstanceMethods.Set(
		"minute",
		[]*ast.Method{
			ast.CreateMethod(
				"minute",
				IntegerType,
				[]*ast.Parameter{},
				func(this *ast.Object, params []*ast.Object, extra map[string]interface{}) interface{} {
					tm := this.Extra["value"].(time.Time)
					return NewInteger(tm.Minute())
				},
			),
		},
	)

	DatetimeType.InstanceMethods.Set(
		"second",
		[]*ast.Method{
			ast.CreateMethod(
				"second",
				IntegerType,
				[]*ast.Parameter{},
				func(this *ast.Object, params []*ast.Object, extra map[string]interface{}) interface{} {
					tm := this.Extra["value"].(time.Time)
					return NewInteger(tm.Second())
				},
			),
		},
	)

	DatetimeType.InstanceMethods.Set(
		"millisecond",
		[]*ast.Method{
			ast.CreateMethod(
				"millisecond",
				IntegerType,
				[]*ast.Parameter{},
				func(this *ast.Object, params []*ast.Object, extra map[string]interface{}) interface{} {
					tm := this.Extra["value"].(time.Time)
					return NewInteger(tm.Nanosecond() / 1000)
				},
			),
		},
	)

	DatetimeType.InstanceMethods.Set(
		"format",
		[]*ast.Method{
			ast.CreateMethod(
				"format",
				StringType,
				[]*ast.Parameter{},
				func(this *ast.Object, params []*ast.Object, extra map[string]interface{}) interface{} {
					tm := this.Extra["value"].(time.Time)
					return NewString(tm.Format("2006/01/02"))
				},
			),
		},
	)

	DatetimeType.StaticMethods.Set(
		"now",
		[]*ast.Method{
			ast.CreateMethod(
				"now",
				DateType,
				[]*ast.Parameter{},
				func(this *ast.Object, params []*ast.Object, extra map[string]interface{}) interface{} {
					obj := ast.CreateObject(DatetimeType)
					obj.Extra["value"] = time.Now()
					return obj
				},
			),
		},
	)

	primitiveClassMap.Set("Datetime", DatetimeType)
}
