package compiler

var PrimitiveClasses []*ClassType

var IntegerType = &ClassType{
	Name: "Integer",
}

var StringType = &ClassType{
	Name: "String",
}

var DoubleType = &ClassType{
	Name: "Double",
}

var BooleanType = &ClassType{
	Name: "Boolean",
}

func init() {
	PrimitiveClasses = []*ClassType{
		IntegerType,
		StringType,
		DoubleType,
		BooleanType,
	}
}
