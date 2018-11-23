package builtin

import (
	"fmt"
	"strings"

	"github.com/tzmfreedom/goland/ast"
)

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

var ListType = &ClassType{
	Name: "List",
}

var MapType = &ClassType{
	Name: "Map",
}

var SetType = &ClassType{
	Name: "Set",
}

var System = &ClassType{
	Name: "System",
	StaticMethods: &MethodMap{
		Data: map[string][]ast.Node{
			"debug": {
				&ast.MethodDeclaration{
					Name: "debug",
					NativeFunction: func(parameter []interface{}) ast.Node {
						o := parameter[0].(*Object)
						fmt.Println(String(o))
						return nil
					},
				},
			},
		},
	},
}

func PrimitiveClassMap() *ClassMap {
	return &ClassMap{
		Data: map[string]*ClassType{
			"integer": IntegerType,
			"string":  StringType,
			"double":  DoubleType,
			"boolean": BooleanType,
			"list":    ListType,
			"set":     SetType,
			"map":     MapType,
			// temporary
			"system": System,
		},
	}
}

type Object struct {
	ClassType      *ClassType
	InstanceFields *ObjectMap
	GenericType    []*ClassType
	Extra          map[string]interface{}
	ToString       func(*Object) string
}

func CreateObject(t *ClassType) *Object {
	return &Object{
		ClassType:      t,
		InstanceFields: NewObjectMap(),
		GenericType:    []*ClassType{},
		Extra:          map[string]interface{}{},
	}
}

func (o *Object) Value() interface{} {
	return o.Extra["value"]
}

func (o *Object) IntegerValue() int {
	return o.Value().(int)
}

func (o *Object) DoubleValue() float64 {
	return o.Value().(float64)
}

func (o *Object) BoolValue() bool {
	return o.Value().(bool)
}

func (o *Object) StringValue() string {
	return o.Value().(string)
}

/**
 * ObjectMap
 */
type ObjectMap struct {
	Data map[string]*Object
}

func NewObjectMap() *ObjectMap {
	return &ObjectMap{
		Data: map[string]*Object{},
	}
}

func (m *ObjectMap) Set(k string, n *Object) {
	m.Data[strings.ToLower(k)] = n
}

func (m *ObjectMap) Get(k string) (*Object, bool) {
	n, ok := m.Data[strings.ToLower(k)]
	return n, ok
}

func (m *ObjectMap) All() map[string]*Object {
	return m.Data
}
