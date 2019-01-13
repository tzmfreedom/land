package builtin

import (
	"time"

	"github.com/tzmfreedom/goland/ast"
)

func init() {
	instanceMethods := NewMethodMap()
	staticMethods := NewMethodMap()
	dateType := CreateClass(
		"Date",
		[]*ast.ConstructorDeclaration{},
		instanceMethods,
		staticMethods,
	)

	instanceMethods.Set(
		"format",
		[]ast.Node{
			CreateMethod(
				"format",
				[]string{"String"},
				[]ast.Node{},
				func(this interface{}, params []interface{}, extra map[string]interface{}) interface{} {
					thisObj := this.(*Object)
					tm := thisObj.Extra["value"].(time.Time)
					return NewString(tm.Format("2006/01/02"))
				},
			),
		},
	)

	staticMethods.Set(
		"today",
		[]ast.Node{
			CreateMethod(
				"today",
				[]string{"Date"},
				[]ast.Node{},
				func(this interface{}, params []interface{}, extra map[string]interface{}) interface{} {
					obj := CreateObject(dateType)
					obj.Extra["value"] = time.Now()
					return obj
				},
			),
		},
	)

	primitiveClassMap.Set("Date", dateType)
}
