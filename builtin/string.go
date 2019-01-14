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
		func(this *Object, params []*Object, extra map[string]interface{}) interface{} {
			split := params[0].StringValue()
			src := this.StringValue()
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

	instanceMethods.Set("split", []*Method{method})
	staticMethods := NewMethodMap()
	staticMethods.Set("valueOf", []*Method{
		CreateMethod(
			"valueOf",
			[]string{"String"},
			[]ast.Node{objectTypeParameter},
			func(this *Object, params []*Object, extra map[string]interface{}) interface{} {
				toConvert := params[0]
				return NewString(String(toConvert))
			},
		),
	})
	classType := CreateClass(
		"String",
		nil,
		instanceMethods,
		staticMethods,
	)
	classType.ToString = func(o *Object) string {
		return o.Value().(string)
	}
	return classType
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
