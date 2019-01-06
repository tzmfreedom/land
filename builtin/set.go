package builtin

var SetType = &ClassType{
	Name: "Set",
}

func init() {
	primitiveClassMap.Set("set", SetType)
}
