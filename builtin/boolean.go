package builtin

import "fmt"

var BooleanType = &ClassType{
	Name: "Boolean",
	ToString: func(o *Object) string {
		return fmt.Sprintf("%t", o.Value().(bool))
	},
}

func init() {
	primitiveClassMap.Set("Boolean", BooleanType)
}
