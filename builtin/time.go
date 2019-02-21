package builtin

import (
	"time"

	"github.com/tzmfreedom/land/ast"
)

var timeTypeRef = &ast.TypeRef{
	Name:       []string{"Time"},
	Parameters: []*ast.TypeRef{},
}

var timeType = ast.CreateClass(
	"Time",
	[]*ast.Method{},
	ast.NewMethodMap(),
	ast.NewMethodMap(),
)

var timeTypeParameter = &ast.Parameter{
	Type: timeType,
	Name: "_",
}

func init() {
	timeType.InstanceMethods.Set(
		"addHours",
		[]*ast.Method{
			ast.CreateMethod(
				"addHours",
				timeType,
				[]*ast.Parameter{IntegerTypeParameter},
				func(this *ast.Object, params []*ast.Object, extra map[string]interface{}) interface{} {
					additionalHours := params[0].IntegerValue()
					tm := this.Extra["value"].(time.Time)
					obj := ast.CreateObject(timeType)
					obj.Extra["value"] = tm.Add(time.Hour * time.Duration(additionalHours))
					return obj
				},
			),
		},
	)

	timeType.InstanceMethods.Set(
		"addMilliseconds",
		[]*ast.Method{
			ast.CreateMethod(
				"addMilliseconds",
				timeType,
				[]*ast.Parameter{IntegerTypeParameter},
				func(this *ast.Object, params []*ast.Object, extra map[string]interface{}) interface{} {
					additionalMilliseconds := params[0].IntegerValue()
					tm := this.Extra["value"].(time.Time)
					obj := ast.CreateObject(timeType)
					obj.Extra["value"] = tm.Add(time.Millisecond * time.Duration(additionalMilliseconds))
					return obj
				},
			),
		},
	)

	timeType.InstanceMethods.Set(
		"addMinutes",
		[]*ast.Method{
			ast.CreateMethod(
				"addMinutes",
				timeType,
				[]*ast.Parameter{IntegerTypeParameter},
				func(this *ast.Object, params []*ast.Object, extra map[string]interface{}) interface{} {
					additionalMinutes := params[0].IntegerValue()
					tm := this.Extra["value"].(time.Time)
					obj := ast.CreateObject(timeType)
					obj.Extra["value"] = tm.Add(time.Minute * time.Duration(additionalMinutes))
					return obj
				},
			),
		},
	)
	timeType.InstanceMethods.Set(
		"addSeconds",
		[]*ast.Method{
			ast.CreateMethod(
				"addSeconds",
				timeType,
				[]*ast.Parameter{IntegerTypeParameter},
				func(this *ast.Object, params []*ast.Object, extra map[string]interface{}) interface{} {
					additionalSeconds := params[0].IntegerValue()
					tm := this.Extra["value"].(time.Time)
					obj := ast.CreateObject(timeType)
					obj.Extra["value"] = tm.Add(time.Second * time.Duration(additionalSeconds))
					return obj
				},
			),
		},
	)
	timeType.InstanceMethods.Set(
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
	timeType.InstanceMethods.Set(
		"millisecond",
		[]*ast.Method{
			ast.CreateMethod(
				"millisecond",
				IntegerType,
				[]*ast.Parameter{},
				func(this *ast.Object, params []*ast.Object, extra map[string]interface{}) interface{} {
					tm := this.Extra["value"].(time.Time)
					return NewInteger(tm.Nanosecond()/1000)
				},
			),
		},
	)
	timeType.InstanceMethods.Set(
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
	timeType.InstanceMethods.Set(
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
	timeType.StaticMethods.Set(
		"newInstance",
		[]*ast.Method{
			ast.CreateMethod(
				"newInstance",
				timeType,
				[]*ast.Parameter{
					IntegerTypeParameter,
					IntegerTypeParameter,
					IntegerTypeParameter,
					IntegerTypeParameter,
				},
				func(this *ast.Object, params []*ast.Object, extra map[string]interface{}) interface{} {
					hour := params[0].IntegerValue()
					minutes := params[1].IntegerValue()
					seconds := params[2].IntegerValue()
					milliseconds := params[3].IntegerValue()

					obj := ast.CreateObject(timeType)
					obj.Extra["value"] = time.Date(2020, time.Month(1), 1, hour, minutes, seconds, milliseconds * 1000, time.UTC)
					return obj
				},
			),
		},
	)
	primitiveClassMap.Set("Time", timeType)
}
