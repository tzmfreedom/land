package ast

type Type interface{}

type ClassType struct {
	Annotations      []Node
	Modifiers        []Node
	Name             string
	SuperClass       Node
	ImplementClasses []Node
	InstanceFields   []Node
	StaticFields     []Node
	InstanceMethods  []Node
	StaticMethods    []Node
	InnerClasses     []Node
	Location         *Location
	Parent           Node
}

const (
	VoidType = iota
	BooleanType
	IntegerType
	StringType
	DoubleType
)

func TypeName(v interface{}) string {
	switch v {
	case VoidType:
		return "void"
	case BooleanType:
		return "boolean"
	case IntegerType:
		return "integer"
	case StringType:
		return "string"
	case DoubleType:
		return "double"
	default:
		return ""
	}
}
