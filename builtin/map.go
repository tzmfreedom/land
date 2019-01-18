package builtin

import (
	"fmt"
	"strings"

	"github.com/tzmfreedom/goland/ast"
)

var MapType = createMapType()

func CreateMapType(keyClass, valueClass *ast.ClassType) *ast.ClassType {
	return &ast.ClassType{
		Name:            "Map",
		Modifiers:       MapType.Modifiers,
		Constructors:    MapType.Constructors,
		InstanceFields:  MapType.InstanceFields,
		InstanceMethods: MapType.InstanceMethods,
		StaticFields:    MapType.StaticFields,
		StaticMethods:   MapType.StaticMethods,
		Generics:        []*ast.ClassType{keyClass, valueClass},
	}
}

func createMapType() *ast.ClassType {
	instanceMethods := ast.NewMethodMap()
	instanceMethods.Set(
		"get",
		[]*ast.Method{
			ast.CreateMethod(
				"get",
				T2type,
				[]*ast.Parameter{t1Parameter},
				func(this *ast.Object, params []*ast.Object, extra map[string]interface{}) interface{} {
					key := params[0].StringValue()
					values := this.Extra["values"].(map[string]*ast.Object)
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
		[]*ast.Method{
			ast.CreateMethod(
				"put",
				T2type,
				[]*ast.Parameter{t1Parameter, t2Parameter},
				func(this *ast.Object, params []*ast.Object, extra map[string]interface{}) interface{} {
					key := params[0].StringValue()
					values := this.Extra["values"].(map[string]*ast.Object)
					values[key] = params[1]
					return nil
				},
			),
		},
	)
	instanceMethods.Set(
		"size",
		[]*ast.Method{
			ast.CreateMethod(
				"size",
				nil,
				[]*ast.Parameter{},
				func(this *ast.Object, params []*ast.Object, extra map[string]interface{}) interface{} {
					return NewInteger(len(this.Extra["values"].(map[string]*ast.Object)))
				},
			),
		},
	)
	instanceMethods.Set(
		"keySet",
		[]*ast.Method{
			ast.CreateMethod(
				"keySet",
				nil,
				[]*ast.Parameter{},
				func(this *ast.Object, params []*ast.Object, extra map[string]interface{}) interface{} {
					keySets := map[string]struct{}{}
					for key, _ := range this.Extra["values"].(map[string]*ast.Object) {
						keySets[key] = struct{}{}
					}
					setClass, _ := PrimitiveClassMap().Get("Set")
					object := ast.CreateObject(setClass)
					object.Extra["values"] = keySets
					object.Extra["generices"] = nil // TODO: implement
					return object
				},
			),
		},
	)

	classType := ast.CreateClass(
		"Map",
		[]*ast.Method{},
		instanceMethods,
		nil,
	)
	classType.ToString = func(o *ast.Object) string {
		values := o.Extra["values"].(map[string]*ast.Object)
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
