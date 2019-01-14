package builtin

import (
	"fmt"
	"strings"

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
				[]ast.Node{t1Parameter},
				func(this interface{}, params []interface{}, extra map[string]interface{}) interface{} {
					key := params[0].(*Object).StringValue()
					thisObj := this.(*Object)
					values := thisObj.Extra["values"].(map[string]*Object)
					if v := values[key]; v != nil {
						return v
					}
					return Null
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
				[]ast.Node{t1Parameter, t2Parameter},
				func(this interface{}, params []interface{}, extra map[string]interface{}) interface{} {
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
				[]ast.Node{},
				func(this interface{}, params []interface{}, extra map[string]interface{}) interface{} {
					thisObj := this.(*Object)
					return NewInteger(len(thisObj.Extra["values"].(map[string]*Object)))
				},
			),
		},
	)
	instanceMethods.Set(
		"keySet",
		[]ast.Node{
			CreateMethod(
				"keySet",
				[]string{"T:1"},
				[]ast.Node{},
				func(this interface{}, params []interface{}, extra map[string]interface{}) interface{} {
					thisObj := this.(*Object)
					keySets := map[string]struct{}{}
					for key, _ := range thisObj.Extra["values"].(map[string]*Object) {
						keySets[key] = struct{}{}
					}
					setClass, _ := PrimitiveClassMap().Get("Set")
					object := CreateObject(setClass)
					object.Extra["values"] = keySets
					object.Extra["generices"] = nil // TODO: implement
					return object
				},
			),
		},
	)

	classType := CreateClass(
		"Map",
		[]*ast.ConstructorDeclaration{},
		instanceMethods,
		nil,
	)
	classType.ToString = func(o *Object) string {
		values := o.Extra["values"].(map[string]*Object)
		parameters := make([]string, len(values))
		i := 0
		for k, v := range values {
			parameters[i] = fmt.Sprintf("%s => %s", k, String(v))
			i++
		}
		return strings.Join(parameters, ", ")
	}
	return classType
}

func init() {
	primitiveClassMap.Set("Map", MapType)
}
