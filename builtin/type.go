package builtin

import (
	"fmt"
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
	Extra            map[string]interface{}
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

func (t *ClassType) Equals(other *ClassType) bool {
	if t.IsGeneric() && other.IsGeneric() {
		if t.Name != other.Name {
			return false
		}
		types := t.Extra["generics"].([]*ClassType)
		otherTypes := other.Extra["generics"].([]*ClassType)
		if len(types) != len(otherTypes) {
			return false
		}
		for i, classType := range types {
			if !classType.Equals(otherTypes[i]) {
				return false
			}
		}
		return true
	}
	if !t.IsGeneric() && !other.IsGeneric() {
		return t == other
	}
	return false
}

func (t *ClassType) IsGeneric() bool {
	return t.Name == "List" ||
		t.Name == "Map" ||
		t.Name == "Set"
}

func (t *ClassType) String() string {
	if t.IsGeneric() {
		classTypes := t.Extra["generics"].([]*ClassType)
		generics := make([]string, len(classTypes))
		for i, classType := range classTypes {
			generics[i] = classType.String()
		}
		return fmt.Sprintf("%s<%s>", t.Name, strings.Join(generics, ", "))
	}
	return t.Name
}

type Field struct {
	Type       ast.Node
	Modifiers  []ast.Node
	Name       string
	Expression ast.Node
	Getter     ast.Node
	Setter     ast.Node
	Location   *ast.Location
	Parent     ast.Node
}

func (f *Field) IsAccessor(modifier string, checkSetter bool) bool {
	is := f.Is(modifier)
	if checkSetter {
		if f.Setter != nil && !f.Setter.(*ast.GetterSetter).IsModifierBlank() {
			is = f.Setter.(*ast.GetterSetter).Is(modifier)
		}
	} else {
		if f.Getter != nil && !f.Getter.(*ast.GetterSetter).IsModifierBlank() {
			is = f.Getter.(*ast.GetterSetter).Is(modifier)
		}
	}
	return is
}

func (f *Field) IsPublic(checkSetter bool) bool {
	return f.IsAccessor("public", checkSetter)
}

func (f *Field) IsPrivate(checkSetter bool) bool {
	return f.IsAccessor("private", checkSetter)
}

func (f *Field) IsProtected(checkSetter bool) bool {
	return f.IsAccessor("protected", checkSetter)
}

func (f *Field) AccessModifier(checkSetter bool) string {
	if f.IsPublic(checkSetter) {
		return "public"
	}
	if f.IsPrivate(checkSetter) {
		return "private"
	}
	if f.IsProtected(checkSetter) {
		return "protected"
	}
	return ""
}

func (f *Field) IsOverride() bool {
	return f.Is("override")
}

func (f *Field) IsAbstract() bool {
	return f.Is("abstract")
}

func (f *Field) IsVirtual() bool {
	return f.Is("virtual")
}

func (f *Field) Is(name string) bool {
	name = strings.ToLower(name)
	for _, modifier := range f.Modifiers {
		modifierName := strings.ToLower(modifier.(*ast.Modifier).Name)
		if modifierName == name {
			return true
		}
	}
	return false
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

func (m *ClassMap) Clear() {
	m.Data = map[string]*ClassType{}
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

func TypeName(v interface{}) string {
	return v.(*ClassType).Name
}

var Null = &Object{
	ClassType:      nil,
	InstanceFields: NewObjectMap(),
	GenericType:    []*ClassType{},
	Extra:          map[string]interface{}{},
}

func NewInteger(value int) *Object {
	t := CreateObject(IntegerType)
	t.Extra["value"] = value
	return t
}

func NewDouble(value float64) *Object {
	t := CreateObject(DoubleType)
	t.Extra["value"] = value
	return t
}

func NewString(value string) *Object {
	t := CreateObject(StringType)
	t.Extra["value"] = value
	return t
}

func NewBoolean(value bool) *Object {
	t := CreateObject(BooleanType)
	t.Extra["value"] = value
	return t
}
