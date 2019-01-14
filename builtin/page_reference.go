package builtin

import (
	"fmt"

	"github.com/tzmfreedom/goland/ast"
)

var pageReferenceType = createPageReferenceType()
var pageReferenceParameter = &ast.Parameter{
	Type: &ast.TypeRef{
		Name:       []string{"PageReference"},
		Parameters: []ast.Node{},
	},
	Name: "_",
}

func createPageReferenceType() *ClassType {
	instanceMethods := NewMethodMap()
	instanceMethods.Set(
		"getUrl",
		[]ast.Node{
			CreateMethod(
				"getUrl",
				[]string{"String"},
				[]ast.Node{},
				func(this interface{}, params []interface{}, extra map[string]interface{}) interface{} {
					thisObj := this.(*Object)
					return thisObj.Extra["url"]
				},
			),
		},
	)
	method := CreateMethod(
		"getParameters",
		nil,
		[]ast.Node{},
		func(this interface{}, params []interface{}, extra map[string]interface{}) interface{} {
			thisObj := this.(*Object)
			return thisObj.Extra["parameters"]
		},
	)
	method.ReturnType = &ast.TypeRef{
		Name: []string{"Map"},
		Parameters: []ast.Node{
			stringTypeRef,
			stringTypeRef,
		},
	}

	instanceMethods.Set(
		"getParameters",
		[]ast.Node{method},
	)

	classType := CreateClass(
		"PageReference",
		[]*ast.ConstructorDeclaration{
			{
				Modifiers: []ast.Node{
					&ast.Modifier{
						Name: "public",
					},
				},
				Parameters: []ast.Node{stringTypeParameter},
				NativeFunction: func(this interface{}, params []interface{}) {
					thisObj := this.(*Object)
					parameters := CreateObject(MapType)
					parameters.Extra["values"] = map[string]*Object{}
					thisObj.Extra = map[string]interface{}{
						"url":        params[0].(*Object),
						"parameters": parameters,
					}
				},
			},
		},
		instanceMethods,
		nil,
	)
	classType.ToString = func(o *Object) string {
		url := o.Extra["url"].(*Object)
		parameters := o.Extra["parameters"].(*Object)
		return fmt.Sprintf(
			"<%s> { url => %s, parameters => (%s) }",
			o.ClassType.Name,
			url.StringValue(),
			String(parameters),
		)
	}
	return classType
}

func init() {
	primitiveClassMap.Set("PageReference", pageReferenceType)
}
