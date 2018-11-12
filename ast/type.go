package ast

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
	Position         *Position
	Parent           Node
}

var (
	BooleanType = &Type{Name: []string{"Boolean"}}
	IntegerType = &Type{Name: []string{"Integer"}}
	StringType  = &Type{Name: []string{"String"}}
	DoubleType  = &Type{Name: []string{"Double"}}
)
