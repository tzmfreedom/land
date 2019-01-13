package builtin

import "fmt"

var DoubleType = &ClassType{
	Name: "Double",
	ToString: func(o *Object) string {
		return fmt.Sprintf("%f", o.Value().(float64))
	},
}

func init() {
	primitiveClassMap.Set("Double", DoubleType)
}
