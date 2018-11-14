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
