package builtin

import (
	"strings"

	"github.com/tzmfreedom/goland/ast"
)

var stringTypeRef = &ast.TypeRef{
	Name:       []string{"String"},
	Parameters: []*ast.TypeRef{},
}

var StringType = &ast.ClassType{Name: "String"}

func createStringType(c *ast.ClassType) *ast.ClassType {
	instanceMethods := ast.NewMethodMap()
	method := ast.CreateMethod(
		"split",
		CreateListType(StringType),
		[]*ast.Parameter{stringTypeParameter},
		func(this *ast.Object, params []*ast.Object, extra map[string]interface{}) interface{} {
			split := params[0].StringValue()
			src := this.StringValue()
			parts := strings.Split(src, split)
			records := make([]*ast.Object, len(parts))
			for i, part := range parts {
				records[i] = NewString(part)
			}
			listType := ast.CreateObject(ListType)
			listType.Extra["records"] = records
			return listType
		},
	)

	instanceMethods.Set("split", []*ast.Method{method})
	staticMethods := ast.NewMethodMap()
	staticMethods.Set("valueOf", []*ast.Method{
		ast.CreateMethod(
			"valueOf",
			StringType,
			[]*ast.Parameter{objectTypeParameter},
			func(this *ast.Object, params []*ast.Object, extra map[string]interface{}) interface{} {
				toConvert := params[0]
				return NewString(String(toConvert))
			},
		),
	})
	c.InstanceMethods = instanceMethods
	c.StaticMethods = staticMethods
	c.ToString = func(o *ast.Object) string {
		return o.Value().(string)
	}
	return c
}

var stringTypeParameter = &ast.Parameter{
	Type: StringType,
	Name: "_",
}

func init() {
	createStringType(StringType)
	primitiveClassMap.Set("String", StringType)
}
