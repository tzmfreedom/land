package builtin

import (
	"strings"

	"github.com/tzmfreedom/goland/ast"
)

var stringTypeRef = &ast.TypeRef{
	Name:       []string{"String"},
	Parameters: []ast.Node{},
}

var StringType *ClassType

func createStringType() *ClassType {
	instanceMethods := NewMethodMap()
	method := CreateMethod(
		"split",
		[]string{"List"},
		[]ast.Node{stringTypeParameter},
		func(this interface{}, params []interface{}, extra map[string]interface{}) interface{} {
			split := params[0].(*Object).StringValue()
			src := this.(*Object).StringValue()
			parts := strings.Split(src, split)
			records := make([]*Object, len(parts))
			for i, part := range parts {
				records[i] = NewString(part)
			}
			listType := CreateObject(ListType)
			listType.Extra["records"] = records
			return listType
		},
	)
	method.ReturnType.(*ast.TypeRef).Parameters = []ast.Node{
		&ast.TypeRef{Name: []string{"String"}},
	}

	instanceMethods.Set("split", []ast.Node{method})
	staticMethods := NewMethodMap()
	staticMethods.Set("valueOf", []ast.Node{
		CreateMethod(
			"valueOf",
			[]string{"String"},
			[]ast.Node{objectTypeParameter},
			func(this interface{}, params []interface{}, extra map[string]interface{}) interface{} {
				toConvert := params[0].(*Object)
				return NewString(String(toConvert))
			},
		),
	})
	return CreateClass(
		"String",
		nil,
		instanceMethods,
		staticMethods,
	)
}

var stringTypeParameter = &ast.Parameter{
	Type: &ast.TypeRef{
		Name:       []string{"Object"},
		Parameters: []ast.Node{},
	},
	Name: "_",
}

func init() {
	StringType = createStringType()
	primitiveClassMap.Set("String", StringType)
}
