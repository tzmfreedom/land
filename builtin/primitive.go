package builtin

import (
	"fmt"
	"io"
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
	Modifiers: []ast.Node{
		&ast.Modifier{
			Name: "public",
		},
	},
	Constructors: []*ast.ConstructorDeclaration{
		{
			Modifiers: []ast.Node{
				&ast.Modifier{
					Name: "public",
				},
			},
			Parameters:     []ast.Node{},
			NativeFunction: func(params []interface{}) {},
		},
		{
			Modifiers: []ast.Node{
				&ast.Modifier{
					Name: "public",
				},
			},
			Parameters: []ast.Node{
				&ast.Parameter{
					Type: &ast.TypeRef{
						Name: []string{"List"},
						Parameters: []ast.Node{
							&ast.TypeRef{
								Name: []string{"T:1"},
							},
						},
					},
					Name: "list",
				},
			},
			NativeFunction: func(params []interface{}) {
				listObj := params[1].(*Object)
				listParams := listObj.Extra["records"].([]*Object)
				newListParams := make([]*Object, len(listParams))
				for i, listParam := range listParams {
					newListParams[i] = &Object{
						ClassType: listParam.ClassType,
					}
				}
				this := params[0].(*Object)
				this.Extra = map[string]interface{}{
					"records": newListParams,
				}
			},
		},
	},
	InstanceFields: NewFieldMap(),
	InstanceMethods: &MethodMap{
		Data: map[string][]ast.Node{
			"size": {
				&ast.MethodDeclaration{
					Name: "size",
					Modifiers: []ast.Node{
						&ast.Modifier{Name: "public"},
					},
					NativeFunction: func(this interface{}, params []interface{}, options ...interface{}) interface{} {
						thisObj := this.(*Object)
						return len(thisObj.Extra["records"].([]*Object))
					},
				},
			},
			"isNext": {
				&ast.MethodDeclaration{
					Name: "next",
					NativeFunction: func(this interface{}, params []interface{}, options ...interface{}) interface{} {
						thisObj := this.(*Object)
						counter := thisObj.Extra["counter"].(int)
						return thisObj.Extra["records"].([]*Object)[counter]
					},
				},
			},
			"next": {
				&ast.MethodDeclaration{
					Name: "next",
					NativeFunction: func(this interface{}, params []interface{}, options ...interface{}) interface{} {
						thisObj := this.(*Object)
						counter := thisObj.Extra["counter"].(int)
						return thisObj.Extra["records"].([]*Object)[counter]
					},
				},
			},
		},
	},
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
					Modifiers: []ast.Node{
						&ast.Modifier{Name: "public"},
					},
					NativeFunction: func(this interface{}, parameter []interface{}, options ...interface{}) interface{} {
						o := parameter[0].(*Object)
						stdout := options[0].(io.Writer)
						fmt.Fprintln(stdout, String(o))
						return nil
					},
				},
			},
			"assertequals": {
				&ast.MethodDeclaration{
					Name: "assertequals",
					Modifiers: []ast.Node{
						&ast.Modifier{Name: "public"},
					},
					NativeFunction: func(this interface{}, parameter []interface{}, options ...interface{}) interface{} {
						expected := parameter[0].(*Object)
						actual := parameter[1].(*Object)
						if expected.Value() != actual.Value() {
							fmt.Printf("expected: %s, actual: %s\n", String(expected), String(actual))
						}
						return nil
					},
				},
			},
		},
	},
}

var AccountType = &ClassType{
	Name: "Account",
}

func NewClassMapWithPrimivie(classTypes []*ClassType) *ClassMap {
	classMap := PrimitiveClassMap()
	for _, classType := range classTypes {
		classMap.Set(classType.Name, classType)
	}
	return classMap
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
			"system":  System,
			"account": AccountType,
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
