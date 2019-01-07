package builtin

var IntegerType = &ClassType{
	Name: "Integer",
}

func init() {
	primitiveClassMap.Set("Integer", IntegerType)
}
