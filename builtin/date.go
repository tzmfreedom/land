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

func init() {
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
