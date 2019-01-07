package builtin

var BooleanType = &ClassType{
	Name: "Boolean",
}

func init() {
	primitiveClassMap.Set("Boolean", BooleanType)
}
