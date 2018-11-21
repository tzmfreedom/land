package compiler

import (
	"strings"

	"github.com/tzmfreedom/goland/ast"
)

type Type interface{}

type ClassType struct {
	Annotations      []ast.Node
	Modifiers        []ast.Node
	Name             string
	SuperClass       ast.Node
	ImplementClasses []ast.Node
	InstanceFields   *FieldMap
	StaticFields     *FieldMap
	InstanceMethods  *MethodMap
	StaticMethods    *MethodMap
	InnerClasses     *ClassMap
	Location         *ast.Location
	Parent           ast.Node
}

type Field struct {
	Type       ast.Node
	Modifiers  []ast.Node
	Name       string
	Expression ast.Node
	Location   *ast.Location
	Parent     ast.Node
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
	Data map[string][]ast.Node
}

func NewMethodMap() *MethodMap {
	return &MethodMap{
		Data: map[string][]ast.Node{},
	}
}

func (m *MethodMap) Add(k string, n ast.Node) {
	if data, ok := m.Get(k); ok {
		data = append(data, n)
		m.Set(k, data)
	} else {
		m.Set(k, []ast.Node{n})
	}
}

func (m *MethodMap) Set(k string, n []ast.Node) {
	m.Data[strings.ToLower(k)] = n
}

func (m *MethodMap) Get(k string) ([]ast.Node, bool) {
	n, ok := m.Data[strings.ToLower(k)]
	return n, ok
}

func (m *MethodMap) All() [][]ast.Node {
	fields := make([][]ast.Node, len(m.Data))
	for _, v := range m.Data {
		fields = append(fields, v)
	}
	return fields
}

/**
 * ClassMap
 */
type ClassMap struct {
	Data map[string]*ClassType
}

func NewClassMap() *ClassMap {
	return &ClassMap{
		Data: map[string]*ClassType{},
	}
}

func (m *ClassMap) Set(k string, n *ClassType) {
	m.Data[strings.ToLower(k)] = n
}

func (m *ClassMap) Get(k string) (*ClassType, bool) {
	n, ok := m.Data[strings.ToLower(k)]
	return n, ok
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
