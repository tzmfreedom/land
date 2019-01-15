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
	instanceMethods := NewMethodMap()
	staticMethods := NewMethodMap()
	dateType := CreateClass(
		"Date",
		[]*Method{},
		instanceMethods,
		staticMethods,
	)

	instanceMethods.Set(
		"format",
		[]*Method{
			CreateMethod(
				"format",
				stringTypeRef,
				[]ast.Node{},
				func(this *Object, params []*Object, extra map[string]interface{}) interface{} {
					tm := this.Extra["value"].(time.Time)
					return NewString(tm.Format("2006/01/02"))
				},
			),
		},
	)

	staticMethods.Set(
		"today",
		[]*Method{
			CreateMethod(
				"today",
				dateTypeRef,
				[]ast.Node{},
				func(this *Object, params []*Object, extra map[string]interface{}) interface{} {
					obj := CreateObject(dateType)
					obj.Extra["value"] = time.Now()
					return obj
				},
			),
		},
	)

	primitiveClassMap.Set("Date", dateType)
}
