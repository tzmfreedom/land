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

var pageReferenceTypeRef = &ast.TypeRef{
	Name:       []string{"PageReference"},
	Parameters: []ast.Node{},
}

func createPageReferenceType() *ClassType {
	instanceMethods := NewMethodMap()
	instanceMethods.Set(
		"getUrl",
		[]*Method{
			CreateMethod(
				"getUrl",
				stringTypeRef,
				[]ast.Node{},
				func(this *Object, params []*Object, extra map[string]interface{}) interface{} {
					return this.Extra["url"]
				},
			),
		},
	)
	method := CreateMethod(
		"getParameters",
		nil,
		[]ast.Node{},
		func(this *Object, params []*Object, extra map[string]interface{}) interface{} {
			return this.Extra["parameters"]
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
		[]*Method{method},
	)

	classType := CreateClass(
		"PageReference",
		[]*Method{
			{
				Modifiers: []ast.Node{
					&ast.Modifier{
						Name: "public",
					},
				},
				Parameters: []ast.Node{stringTypeParameter},
				NativeFunction: func(this *Object, params []*Object, extra map[string]interface{}) interface{} {
					parameters := CreateObject(MapType)
					parameters.Extra["values"] = map[string]*Object{}
					this.Extra = map[string]interface{}{
						"url":        params[0],
						"parameters": parameters,
					}
					return nil
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
