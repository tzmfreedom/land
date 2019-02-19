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
	Extra:          map[string]interface{}{},
}

func NewInteger(value int) *ast.Object {
	t := ast.CreateObject(IntegerType)
	t.Extra["value"] = value
	return t
}

func NewLong(value int) *ast.Object {
	t := ast.CreateObject(LongType)
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
	classMap, ok := m.Get(k)
	if !ok {
		panic("NameSpaceStore#Add failed")
	}
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
	if t == ObjectType {
		return true
	}
	if t.IsGenerics() && other.IsGenerics() {
		if t.Name != other.Name {
			return false
		}
		types := t.Generics
		otherTypes := other.Generics
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
	if !t.IsGenerics() && !other.IsGenerics() {
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

func SearchMethod(receiverClass *ast.ClassType, methods []*ast.Method, parameters []*ast.ClassType) *ast.Method {
	l := len(parameters)
	for _, m := range methods {
		if len(m.Parameters) != l {
			continue
		}
		match := true

		for i, p := range m.Parameters {
			inputParam := convertGenericsType(receiverClass, parameters[i])
			methodParam := convertGenericsType(receiverClass, p.Type)

			// TODO: implement
			// extend, implements, Object
			if methodParam == ObjectType {
				continue
			}
			if !Equals(inputParam, methodParam) {
				match = false
				break
			}
		}
		if match {
			return m
		}
	}
	return nil
}

func convertGenericsType(receiverClass *ast.ClassType, classType *ast.ClassType) *ast.ClassType {
	generics := receiverClass.Generics
	if classType == T1type {
		return generics[0]
	}
	if classType == T2type {
		return generics[1]
	}
	if classType.IsGenerics() {
		newClassType := ast.CreateClass(classType.Name, classType.Constructors, classType.InstanceMethods, classType.StaticMethods)
		newClassType.Generics = make([]*ast.ClassType, len(classType.Generics))
		for i, g := range classType.Generics {
			newClassType.Generics[i] = convertGenericsType(receiverClass, g)
		}
		return newClassType
	}
	return classType
}
