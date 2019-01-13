package builtin

import "fmt"

var IntegerType = &ClassType{
	Name: "Integer",
	ToString: func(o *Object) string {
		return fmt.Sprintf("%d", o.Value().(int))
	},
}

func init() {
	primitiveClassMap.Set("Integer", IntegerType)
}
