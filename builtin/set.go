package builtin

import "github.com/tzmfreedom/goland/ast"

var setType = createSetType()

func createSetType() *ClassType {
	instanceMethods := NewMethodMap()
	instanceMethods.Set(
		"size",
		[]ast.Node{
			CreateMethod(
				"size",
				[]string{"Integer"},
				func(this interface{}, params []interface{}, options ...interface{}) interface{} {
					thisObj := this.(*Object)
					return NewInteger(len(thisObj.Extra["values"].(map[string]struct{})))
				},
			),
		},
	)
	instanceMethods.Set(
		"add",
		[]ast.Node{
			CreateMethod(
				"add",
				[]string{"T:1"},
				func(this interface{}, params []interface{}, options ...interface{}) interface{} {
					key := params[0].(*Object).StringValue()
					thisObj := this.(*Object)
					values := thisObj.Extra["values"].(map[string]struct{})
					values[key] = struct{}{}
					return nil
				},
			),
		},
	)
	instanceMethods.Set(
		"clear",
		[]ast.Node{
			CreateMethod(
				"clear",
				nil,
				func(this interface{}, params []interface{}, options ...interface{}) interface{} {
					thisObj := this.(*Object)
					thisObj.Extra["values"] = map[string]struct{}{}
					return nil
				},
			),
		},
	)

	return &ClassType{
		Name: "Set",
		Constructors: []*ast.ConstructorDeclaration{
			{
				Modifiers: []ast.Node{
					&ast.Modifier{
						Name: "public",
					},
				},
				Parameters:     []ast.Node{},
				NativeFunction: func(params []interface{}) {},
			},
		},
		InstanceMethods: instanceMethods,
	}
}

func init() {
	primitiveClassMap.Set("set", setType)
}
