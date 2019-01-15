package builtin

import "github.com/tzmfreedom/goland/ast"

var setType = createSetType()

func createSetType() *ClassType {
	instanceMethods := NewMethodMap()
	instanceMethods.Set(
		"size",
		[]*Method{
			CreateMethod(
				"size",
				integerTypeRef,
				[]ast.Node{},
				func(this *Object, params []*Object, extra map[string]interface{}) interface{} {
					return NewInteger(len(this.Extra["values"].(map[string]struct{})))
				},
			),
		},
	)
	instanceMethods.Set(
		"add",
		[]*Method{
			CreateMethod(
				"add",
				integerTypeRef,
				[]ast.Node{t1Parameter},
				func(this *Object, params []*Object, extra map[string]interface{}) interface{} {
					key := params[0].StringValue()
					values := this.Extra["values"].(map[string]struct{})
					values[key] = struct{}{}
					return nil
				},
			),
		},
	)
	instanceMethods.Set(
		"clear",
		[]*Method{
			CreateMethod(
				"clear",
				nil,
				[]ast.Node{},
				func(this *Object, params []*Object, extra map[string]interface{}) interface{} {
					this.Extra["values"] = map[string]struct{}{}
					return nil
				},
			),
		},
	)

	return &ClassType{
		Name: "Set",
		Constructors: []*Method{
			{
				Modifiers: []ast.Node{
					&ast.Modifier{
						Name: "public",
					},
				},
				Parameters: []ast.Node{},
				NativeFunction: func(this *Object, params []*Object, extra map[string]interface{}) interface{} {
					return nil
				},
			},
		},
		InstanceMethods: instanceMethods,
	}
}

func init() {
	primitiveClassMap.Set("set", setType)
}
