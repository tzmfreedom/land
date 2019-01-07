package builtin

var DoubleType = &ClassType{
	Name: "Double",
}

func init() {
	primitiveClassMap.Set("Double", DoubleType)
}
