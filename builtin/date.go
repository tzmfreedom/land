package builtin

import (
	"time"

	"github.com/tzmfreedom/land/ast"
)

var dateTypeRef = &ast.TypeRef{
	Name:       []string{"Date"},
	Parameters: []*ast.TypeRef{},
}

var DateType = ast.CreateClass(
	"Date",
	[]*ast.Method{},
	ast.NewMethodMap(),
	ast.NewMethodMap(),
)

var dateTypeParameter = &ast.Parameter{
	Type: DateType,
	Name: "_",
}

func init() {
	DateType.InstanceMethods.Set(
		"addDays",
		[]*ast.Method{
			ast.CreateMethod(
				"addDays",
				DateType,
				[]*ast.Parameter{IntegerTypeParameter},
				func(this *ast.Object, params []*ast.Object, extra map[string]interface{}) interface{} {
					tm := this.Extra["value"].(time.Time)
					days := params[0].IntegerValue()

					obj := ast.CreateObject(DateType)
					obj.Extra["value"] = tm.AddDate(0, 0, days)
					return obj
				},
			),
		},
	)
	DateType.InstanceMethods.Set(
		"addMonths",
		[]*ast.Method{
			ast.CreateMethod(
				"addMonths",
				DateType,
				[]*ast.Parameter{IntegerTypeParameter},
				func(this *ast.Object, params []*ast.Object, extra map[string]interface{}) interface{} {
					tm := this.Extra["value"].(time.Time)
					months := params[0].IntegerValue()

					obj := ast.CreateObject(DateType)
					obj.Extra["value"] = tm.AddDate(0, months, 0)
					return obj
				},
			),
		},
	)
	DateType.InstanceMethods.Set(
		"addYears",
		[]*ast.Method{
			ast.CreateMethod(
				"addYears",
				DateType,
				[]*ast.Parameter{IntegerTypeParameter},
				func(this *ast.Object, params []*ast.Object, extra map[string]interface{}) interface{} {
					tm := this.Extra["value"].(time.Time)
					years := params[0].IntegerValue()

					obj := ast.CreateObject(DateType)
					obj.Extra["value"] = tm.AddDate(years, 0, 0)
					return obj
				},
			),
		},
	)
	DateType.InstanceMethods.Set(
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
	DateType.InstanceMethods.Set(
		"dayOfYear",
		[]*ast.Method{
			ast.CreateMethod(
				"dayOfYear",
				IntegerType,
				[]*ast.Parameter{},
				func(this *ast.Object, params []*ast.Object, extra map[string]interface{}) interface{} {
					tm := this.Extra["value"].(time.Time)
					return NewInteger(tm.YearDay())
				},
			),
		},
	)
	DateType.InstanceMethods.Set(
		"daysBetween",
		[]*ast.Method{
			ast.CreateMethod(
				"daysBetween",
				IntegerType,
				[]*ast.Parameter{dateTypeParameter},
				func(this *ast.Object, params []*ast.Object, extra map[string]interface{}) interface{} {
					secondDate := params[0].Extra["value"].(time.Time)
					tm := this.Extra["value"].(time.Time)
					return NewInteger(secondDate.Day() - tm.Day())
				},
			),
		},
	)
	DateType.InstanceMethods.Set(
		"daysInMonth",
		[]*ast.Method{
			ast.CreateMethod(
				"daysInMonth",
				IntegerType,
				[]*ast.Parameter{
					IntegerTypeParameter,
					IntegerTypeParameter,
				},
				func(this *ast.Object, params []*ast.Object, extra map[string]interface{}) interface{} {
					year := params[0].IntegerValue()
					month := params[1].IntegerValue()
					tm := time.Date(year, time.Month(month+1), 0, 0, 0, 0, 0, time.UTC)
					return NewInteger(tm.Day())
				},
			),
		},
	)
	DateType.InstanceMethods.Set(
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
	DateType.InstanceMethods.Set(
		"isLeapYear",
		[]*ast.Method{
			ast.CreateMethod(
				"isLeapYear",
				BooleanType,
				[]*ast.Parameter{IntegerTypeParameter},
				func(this *ast.Object, params []*ast.Object, extra map[string]interface{}) interface{} {
					year := params[0].IntegerValue()
					var isLeapYear bool
					if year % 400 == 0 {
						isLeapYear = true
					} else if year % 100 == 0 {
						isLeapYear = false
					} else if year % 4 == 0 {
						isLeapYear = true
					}
					return NewBoolean(isLeapYear)
				},
			),
		},
	)
	DateType.InstanceMethods.Set(
		"isSameDay",
		[]*ast.Method{
			ast.CreateMethod(
				"isSameDay",
				BooleanType,
				[]*ast.Parameter{dateTypeParameter},
				func(this *ast.Object, params []*ast.Object, extra map[string]interface{}) interface{} {
					dateToCompare := params[0].Extra["value"].(time.Time)
					tm := this.Extra["value"].(time.Time)
					return NewBoolean(dateToCompare.Equal(tm))
				},
			),
		},
	)
	DateType.InstanceMethods.Set(
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
	DateType.InstanceMethods.Set(
		"monthsBetween",
		[]*ast.Method{
			ast.CreateMethod(
				"monthsBetween",
				StringType,
				[]*ast.Parameter{},
				func(this *ast.Object, params []*ast.Object, extra map[string]interface{}) interface{} {
					secondDate := params[0].Extra["value"].(time.Time)
					thisDate := this.Extra["value"].(time.Time)
					sy := secondDate.Year()
					sm := int(secondDate.Month())
					ty := thisDate.Year()
					tm := int(thisDate.Month())
					return NewInteger((sy - ty) * 12 + sm - tm)
				},
			),
		},
	)
	DateType.StaticMethods.Set(
		"newInstance",
		[]*ast.Method{
			ast.CreateMethod(
				"newInstance",
				StringType,
				[]*ast.Parameter{},
				func(this *ast.Object, params []*ast.Object, extra map[string]interface{}) interface{} {
					tm := this.Extra["value"].(time.Time)
					return NewString(tm.Format("2006/01/02"))
				},
			),
		},
	)
	DateType.StaticMethods.Set(
		"parse",
		[]*ast.Method{
			ast.CreateMethod(
				"parse",
				DateType,
				[]*ast.Parameter{stringTypeParameter},
				func(this *ast.Object, params []*ast.Object, extra map[string]interface{}) interface{} {
					tm, err := time.Parse("2006/01/02", params[0].StringValue())
					if err != nil {
						panic(err)
					}
					obj := ast.CreateObject(DateType)
					obj.Extra["value"] = tm
					return obj
				},
			),
		},
	)
	DateType.InstanceMethods.Set(
		"toStartOfWeek",
		[]*ast.Method{
			ast.CreateMethod(
				"toStartOfWeek",
				DateType,
				[]*ast.Parameter{},
				func(this *ast.Object, params []*ast.Object, extra map[string]interface{}) interface{} {
					tm := this.Extra["value"].(time.Time)
					obj := ast.CreateObject(DateType)
					wk := int(tm.Weekday())
					obj.Extra["value"] = tm.AddDate(0, 0, -wk)
					return obj
				},
			),
		},
	)
	DateType.StaticMethods.Set(
		"valueOf",
		[]*ast.Method{
			ast.CreateMethod(
				"valueOf",
				DateType,
				[]*ast.Parameter{stringTypeParameter},
				func(this *ast.Object, params []*ast.Object, extra map[string]interface{}) interface{} {
					stringDate := params[0].StringValue()
					obj := ast.CreateObject(DateType)
					value, err := time.Parse("2016-01-02 03:04:05", stringDate)
					if err != nil {
						panic(err)
					}
					obj.Extra["value"] = value
					return obj
				},
			),
			ast.CreateMethod(
				"valueOf",
				DateType,
				[]*ast.Parameter{objectTypeParameter},
				func(this *ast.Object, params []*ast.Object, extra map[string]interface{}) interface{} {
					if this.ClassType != DateType {
						panic("no expected type")
					}
					return this
				},
			),
		},
	)
	DateType.InstanceMethods.Set(
		"year",
		[]*ast.Method{
			ast.CreateMethod(
				"year",
				StringType,
				[]*ast.Parameter{},
				func(this *ast.Object, params []*ast.Object, extra map[string]interface{}) interface{} {
					tm := this.Extra["value"].(time.Time)
					return NewInteger(tm.Year())
				},
			),
		},
	)

	DateType.StaticMethods.Set(
		"today",
		[]*ast.Method{
			ast.CreateMethod(
				"today",
				DateType,
				[]*ast.Parameter{},
				func(this *ast.Object, params []*ast.Object, extra map[string]interface{}) interface{} {
					obj := ast.CreateObject(DateType)
					obj.Extra["value"] = time.Now()
					return obj
				},
			),
		},
	)

	primitiveClassMap.Set("Date", DateType)
}
