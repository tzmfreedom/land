package builtin

import (
	"fmt"

	"github.com/tzmfreedom/goland/ast"
)

var pageReferenceType = createPageReferenceType()
var pageReferenceParameter = &ast.Parameter{
	TypeRef: &ast.TypeRef{
		Name:       []string{"PageReference"},
		Parameters: []ast.Node{},
	},
	Name: "_",
}

var pageReferenceTypeRef = &ast.TypeRef{
	Name:       []string{"PageReference"},
	Parameters: []ast.Node{},
}

func createPageReferenceType() *ast.ClassType {
	instanceMethods := ast.NewMethodMap()
	instanceMethods.Set(
		"getUrl",
		[]*ast.Method{
			ast.CreateMethod(
				"getUrl",
				stringTypeRef,
				[]*ast.Parameter{},
				func(this *ast.Object, params []*ast.Object, extra map[string]interface{}) interface{} {
					return this.Extra["url"]
				},
			),
		},
	)
	method := ast.CreateMethod(
		"getParameters",
		nil,
		[]*ast.Parameter{},
		func(this *ast.Object, params []*ast.Object, extra map[string]interface{}) interface{} {
			return this.Extra["parameters"]
		},
	)
	method.ReturnTypeRef = &ast.TypeRef{
		Name: []string{"Map"},
		Parameters: []ast.Node{
			stringTypeRef,
			stringTypeRef,
		},
	}

	instanceMethods.Set(
		"getParameters",
		[]*ast.Method{method},
	)

	classType := ast.CreateClass(
		"PageReference",
		[]*ast.Method{
			{
				Modifiers:  []*ast.Modifier{ast.PublicModifier()},
				Parameters: []*ast.Parameter{stringTypeParameter},
				NativeFunction: func(this *ast.Object, params []*ast.Object, extra map[string]interface{}) interface{} {
					parameters := ast.CreateObject(MapType)
					parameters.Extra["values"] = map[string]*ast.Object{}
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
	classType.ToString = func(o *ast.Object) string {
		url := o.Extra["url"].(*ast.Object)
		parameters := o.Extra["parameters"].(*ast.Object)
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
