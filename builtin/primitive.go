package builtin

import (
	"github.com/k0kubun/pp"
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

var SystemType = &ClassType{
	Name: "System",
	StaticMethods: &MethodMap{
		Data: map[string][]ast.Node{
			"debug": {
				&ast.MethodDeclaration{
					Name: "debug",
					NativeFunction: func(parameter []ast.Node) ast.Node {
						pp.Println(parameter[0])
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
		},
	}
}
