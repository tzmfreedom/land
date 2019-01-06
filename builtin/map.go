package builtin

var MapType = &ClassType{
	Name: "Map",
}

func init() {
	primitiveClassMap.Set("map", MapType)
}
