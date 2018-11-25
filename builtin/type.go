package builtin

import (
	"strings"

	"github.com/tzmfreedom/goland/ast"
)

type ClassType struct {
	Annotations      []ast.Node
	Modifiers        []ast.Node
	Name             string
	SuperClass       ast.Node
	ImplementClasses []ast.Node
	Constructors     []*ast.ConstructorDeclaration
	InstanceFields   *FieldMap
	StaticFields     *FieldMap
	InstanceMethods  *MethodMap
	StaticMethods    *MethodMap
	InnerClasses     *ClassMap
	Location         *ast.Location
	Parent           ast.Node
}

func (t *ClassType) IsPrimitive() bool {
	if t == IntegerType ||
		t == StringType ||
		t == BooleanType ||
		t == DoubleType {
		return true
	}
	return false
}

type Field struct {
	Type       ast.Node
	Modifiers  []ast.Node
	Name       string
	Expression ast.Node
	Location   *ast.Location
	Parent     ast.Node
}

func (f *Field) IsPublic() bool {
	for _, m := range f.Modifiers {
		if m.(*ast.Modifier).Name == "public" {
			return true
		}
	}
	return false
}

func (f *Field) IsPrivate() bool {
	for _, m := range f.Modifiers {
		if m.(*ast.Modifier).Name == "private" {
			return true
		}
	}
	return false
}

func (f *Field) IsProtected() bool {
	for _, m := range f.Modifiers {
		if m.(*ast.Modifier).Name == "protected" {
			return true
		}
	}
	return false
}

func (f *Field) AccessModifier() string {
	if f.IsPublic() {
		return "public"
	}
	if f.IsPrivate() {
		return "private"
	}
	if f.IsProtected() {
		return "protected"
	}
	return ""
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

/**
 * NameSpaces
 */
type NameSpaceStore struct {
	Data map[string]*ClassMap
}

func NewNameSpaceStore() *NameSpaceStore {
	return &NameSpaceStore{
		Data: map[string]*ClassMap{},
	}
}

func (m *NameSpaceStore) Add(k string, n *ClassType) {
	classMap, _ := m.Get(k)
	classMap.Set(k, n)
}

func (m *NameSpaceStore) Set(k string, n *ClassMap) {
	m.Data[strings.ToLower(k)] = n
}

func (m *NameSpaceStore) Get(k string) (*ClassMap, bool) {
	n, ok := m.Data[strings.ToLower(k)]
	return n, ok
}

const (
	VoidType = iota
)

func TypeName(v interface{}) string {
	return v.(*ClassType).Name
}
