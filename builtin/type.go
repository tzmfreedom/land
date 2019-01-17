package builtin

import (
	"strings"

	"github.com/tzmfreedom/goland/ast"
)

var NullType = &ast.ClassType{
	Name:     "null",
	ToString: func(o *ast.Object) string { return "null" },
}

var Null = &ast.Object{
	ClassType:      NullType,
	InstanceFields: ast.NewObjectMap(),
	GenericType:    []*ast.ClassType{},
	Extra:          map[string]interface{}{},
}

func NewInteger(value int) *ast.Object {
	t := ast.CreateObject(IntegerType)
	t.Extra["value"] = value
	return t
}

func NewDouble(value float64) *ast.Object {
	t := ast.CreateObject(DoubleType)
	t.Extra["value"] = value
	return t
}

func NewString(value string) *ast.Object {
	t := ast.CreateObject(StringType)
	t.Extra["value"] = value
	return t
}

func NewBoolean(value bool) *ast.Object {
	t := ast.CreateObject(BooleanType)
	t.Extra["value"] = value
	return t
}

/**
 * NameSpaces
 */
type NameSpaceStore struct {
	Data map[string]*ast.ClassMap
}

func NewNameSpaceStore() *NameSpaceStore {
	return &NameSpaceStore{
		Data: map[string]*ast.ClassMap{},
	}
}

func (m *NameSpaceStore) Add(k string, n *ast.ClassType) {
	classMap, _ := m.Get(k)
	classMap.Set(k, n)
}

func (m *NameSpaceStore) Set(k string, n *ast.ClassMap) {
	m.Data[strings.ToLower(k)] = n
}

func (m *NameSpaceStore) Get(k string) (*ast.ClassMap, bool) {
	n, ok := m.Data[strings.ToLower(k)]
	return n, ok
}

func TypeName(v interface{}) string {
	return v.(*ast.ClassType).Name
}

func Equals(t, other *ast.ClassType) bool {
	if other == NullType {
		return true
	}
	if t.IsGeneric() && other.IsGeneric() {
		if t.Name != other.Name {
			return false
		}
		types := t.Extra["generics"].([]*ast.ClassType)
		otherTypes := other.Extra["generics"].([]*ast.ClassType)
		if len(types) != len(otherTypes) {
			return false
		}
		for i, classType := range types {
			if !Equals(classType, otherTypes[i]) {
				return false
			}
		}
		return true
	}
	if !t.IsGeneric() && !other.IsGeneric() {
		if t == other {
			return true
		}
		if other.SuperClass != nil {
			if Equals(t, other.SuperClass) {
				return true
			}
		}
		if other.ImplementClasses != nil {
			for _, impl := range other.ImplementClasses {
				if t == impl {
					return true
				}
			}
		}
	}
	return false
}
