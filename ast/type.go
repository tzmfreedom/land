package ast

import "strings"

type Type interface{}

type ClassType struct {
	Annotations      []Node
	Modifiers        []Node
	Name             string
	SuperClass       Node
	ImplementClasses []Node
	InstanceFields   *FieldMap
	StaticFields     *FieldMap
	InstanceMethods  *MethodMap
	StaticMethods    *MethodMap
	InnerClasses     map[string]*ClassType
	Location         *Location
	Parent           Node
}

type Field struct {
	Type       Node
	Modifiers  []Node
	Name       string
	Expression Node
	Location   *Location
	Parent     Node
}

type FieldMap struct {
	Data map[string]*Field
}

func NewFieldMap() *FieldMap {
	return &FieldMap{
		Data: map[string]*Field{},
	}
}

func (m *FieldMap) Set(k string, n *Field) {
	m.Data[strings.ToLower(k)] = n
}

func (m *FieldMap) Get(k string) (*Field, bool) {
	n, ok := m.Data[strings.ToLower(k)]
	return n, ok
}

func (m *FieldMap) All() []*Field {
	fields := make([]*Field, len(m.Data))
	for _, v := range m.Data {
		fields = append(fields, v)
	}
	return fields
}

type MethodMap struct {
	Data map[string][]Node
}

func NewMethodMap() *MethodMap {
	return &MethodMap{
		Data: map[string][]Node{},
	}
}

func (m *MethodMap) Add(k string, n Node) {
	if data, ok := m.Get(k); ok {
		data = append(data, n)
		m.Set(k, data)
	} else {
		m.Set(k, []Node{n})
	}
}

func (m *MethodMap) Set(k string, n []Node) {
	m.Data[strings.ToLower(k)] = n
}

func (m *MethodMap) Get(k string) ([]Node, bool) {
	n, ok := m.Data[strings.ToLower(k)]
	return n, ok
}

func (m *MethodMap) All() [][]Node {
	fields := make([][]Node, len(m.Data))
	for _, v := range m.Data {
		fields = append(fields, v)
	}
	return fields
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
