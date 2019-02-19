package builtin

import (
	"fmt"

	"github.com/k0kubun/pp"
	"github.com/tzmfreedom/land/ast"
)

var PrimitiveClasses []*ast.ClassType

func NewClassMapWithPrimivie(classTypes []*ast.ClassType) *ast.ClassMap {
	classMap := PrimitiveClassMap()
	for _, classType := range classTypes {
		classMap.Set(classType.Name, classType)
	}
	return classMap
}

var primitiveClassMap = &ast.ClassMap{
	Data: map[string]*ast.ClassType{
		"integer": IntegerType,
		"string":  StringType,
		"double":  DoubleType,
		"boolean": BooleanType,
	},
}

func PrimitiveClassMap() *ast.ClassMap {
	return primitiveClassMap
}

var nameSpaceStore = NewNameSpaceStore()

func GetNameSpaceStore() *NameSpaceStore {
	return nameSpaceStore
}

var ReturnType = &ast.ClassType{Name: "Return"}
var RaiseType = &ast.ClassType{Name: "Raise"}
var BreakType = &ast.ClassType{Name: "Break"}
var Break = &ast.Object{ClassType: BreakType}
var ContinueType = &ast.ClassType{Name: "Continue"}
var Continue = &ast.Object{ClassType: ContinueType}

func CreateReturn(value *ast.Object) *ast.Object {
	return &ast.Object{
		ClassType: ReturnType,
		Extra: map[string]interface{}{
			"value": value,
		},
	}
}

func CreateRaise(value *ast.Object) *ast.Object {
	return &ast.Object{
		ClassType: RaiseType,
		Extra: map[string]interface{}{
			"value": value,
		},
	}
}

func Debug(obj interface{}) {
	switch o := obj.(type) {
	case *ast.Object:
		fmt.Println(o.ClassType)
		fmt.Println(o.Extra)
	case ast.Node:
		fmt.Println(o.GetLocation())
		fmt.Println(o.GetType())
	default:
		pp.Println(o)
	}
}
