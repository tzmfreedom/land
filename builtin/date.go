package builtin

import (
	"time"

	"github.com/tzmfreedom/goland/ast"
)

var dateTypeRef = &ast.TypeRef{
	Name:       []string{"Date"},
	Parameters: []ast.Node{},
}

func init() {
	instanceMethods := ast.NewMethodMap()
	staticMethods := ast.NewMethodMap()
	dateType := ast.CreateClass(
		"Date",
		[]*ast.Method{},
		instanceMethods,
		staticMethods,
	)

	instanceMethods.Set(
		"format",
		[]*ast.Method{
			ast.CreateMethod(
				"format",
				stringTypeRef,
				[]*ast.Parameter{},
				func(this *ast.Object, params []*ast.Object, extra map[string]interface{}) interface{} {
					tm := this.Extra["value"].(time.Time)
					return NewString(tm.Format("2006/01/02"))
				},
			),
		},
	)

	staticMethods.Set(
		"today",
		[]*ast.Method{
			ast.CreateMethod(
				"today",
				dateTypeRef,
				[]*ast.Parameter{},
				func(this *ast.Object, params []*ast.Object, extra map[string]interface{}) interface{} {
					obj := ast.CreateObject(dateType)
					obj.Extra["value"] = time.Now()
					return obj
				},
			),
		},
	)

	primitiveClassMap.Set("Date", dateType)
}
