package builtin

import (
	"strings"

	"github.com/tzmfreedom/goland/ast"
)

var stringTypeRef = &ast.TypeRef{
	Name:       []string{"String"},
	Parameters: []ast.Node{},
}

var StringType *ast.ClassType

func createStringType() *ast.ClassType {
	instanceMethods := ast.NewMethodMap()
	method := ast.CreateMethod(
		"split",
		CreateListTypeRef(stringTypeRef),
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
	method.ReturnTypeRef.Parameters = []ast.Node{stringTypeRef}

	instanceMethods.Set("split", []*ast.Method{method})
	staticMethods := ast.NewMethodMap()
	staticMethods.Set("valueOf", []*ast.Method{
		ast.CreateMethod(
			"valueOf",
			stringTypeRef,
			[]*ast.Parameter{objectTypeParameter},
			func(this *ast.Object, params []*ast.Object, extra map[string]interface{}) interface{} {
				toConvert := params[0]
				return NewString(String(toConvert))
			},
		),
	})
	classType := ast.CreateClass(
		"String",
		nil,
		instanceMethods,
		staticMethods,
	)
	classType.ToString = func(o *ast.Object) string {
		return o.Value().(string)
	}
	return classType
}

var stringTypeParameter = &ast.Parameter{
	TypeRef: &ast.TypeRef{
		Name:       []string{"Object"},
		Parameters: []ast.Node{},
	},
	Name: "_",
}

func init() {
	StringType = createStringType()
	primitiveClassMap.Set("String", StringType)
}
