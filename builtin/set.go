package builtin

import "github.com/tzmfreedom/land/ast"

var setType = createSetType()

func createSetType() *ast.ClassType {
	instanceMethods := ast.NewMethodMap()
	instanceMethods.Set(
		"size",
		[]*ast.Method{
			ast.CreateMethod(
				"size",
				IntegerType,
				[]*ast.Parameter{},
				func(this *ast.Object, params []*ast.Object, extra map[string]interface{}) interface{} {
					return NewInteger(len(this.Extra["values"].(map[string]struct{})))
				},
			),
		},
	)
	instanceMethods.Set(
		"add",
		[]*ast.Method{
			ast.CreateMethod(
				"add",
				IntegerType,
				[]*ast.Parameter{t1Parameter},
				func(this *ast.Object, params []*ast.Object, extra map[string]interface{}) interface{} {
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
		[]*ast.Method{
			ast.CreateMethod(
				"clear",
				nil,
				[]*ast.Parameter{},
				func(this *ast.Object, params []*ast.Object, extra map[string]interface{}) interface{} {
					this.Extra["values"] = map[string]struct{}{}
					return nil
				},
			),
		},
	)

	return &ast.ClassType{
		Name: "Set",
		Constructors: []*ast.Method{
			{
				Modifiers:  []*ast.Modifier{ast.PublicModifier()},
				Parameters: []*ast.Parameter{},
				NativeFunction: func(this *ast.Object, params []*ast.Object, extra map[string]interface{}) interface{} {
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
