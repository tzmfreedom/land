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
		[]*Method{
			CreateMethod(
				"get",
				[]string{"T:2"},
				[]ast.Node{t1Parameter},
				func(this *Object, params []*Object, extra map[string]interface{}) interface{} {
					key := params[0].StringValue()
					values := this.Extra["values"].(map[string]*Object)
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
		[]*Method{
			CreateMethod(
				"put",
				nil,
				[]ast.Node{t1Parameter, t2Parameter},
				func(this *Object, params []*Object, extra map[string]interface{}) interface{} {
					key := params[0].StringValue()
					values := this.Extra["values"].(map[string]*Object)
					values[key] = params[1]
					return nil
				},
			),
		},
	)
	instanceMethods.Set(
		"size",
		[]*Method{
			CreateMethod(
				"size",
				[]string{"Integer"},
				[]ast.Node{},
				func(this *Object, params []*Object, extra map[string]interface{}) interface{} {
					return NewInteger(len(this.Extra["values"].(map[string]*Object)))
				},
			),
		},
	)
	instanceMethods.Set(
		"keySet",
		[]*Method{
			CreateMethod(
				"keySet",
				[]string{"T:1"},
				[]ast.Node{},
				func(this *Object, params []*Object, extra map[string]interface{}) interface{} {
					keySets := map[string]struct{}{}
					for key, _ := range this.Extra["values"].(map[string]*Object) {
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
		[]*Method{},
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
