package builtin

import (
	"github.com/tzmfreedom/goland/ast"
)

var MapType = createMapType()

func createMapType() *ClassType {
	instanceMethods := NewMethodMap()
	instanceMethods.Set(
		"get",
		[]ast.Node{
			CreateMethod(
				"get",
				[]string{"T:2"},
				func(this interface{}, params []interface{}, options ...interface{}) interface{} {
					key := params[0].(*Object).StringValue()
					thisObj := this.(*Object)
					values := thisObj.Extra["values"].(map[string]*Object)
					return values[key]
				},
			),
		},
	)
	instanceMethods.Set(
		"put",
		[]ast.Node{
			CreateMethod(
				"put",
				nil,
				func(this interface{}, params []interface{}, options ...interface{}) interface{} {
					key := params[0].(*Object).StringValue()
					thisObj := this.(*Object)
					values := thisObj.Extra["values"].(map[string]*Object)
					values[key] = params[1].(*Object)
					return nil
				},
			),
		},
	)
	instanceMethods.Set(
		"size",
		[]ast.Node{
			CreateMethod(
				"size",
				[]string{"Integer"},
				func(this interface{}, params []interface{}, options ...interface{}) interface{} {
					thisObj := this.(*Object)
					return NewInteger(len(thisObj.Extra["values"].(map[string]*Object)))
				},
			),
		},
	)

	return CreateClass(
		"Map",
		[]*ast.ConstructorDeclaration{},
		instanceMethods,
		nil,
	)
}

func init() {
	primitiveClassMap.Set("Map", MapType)
}
