package builtin

import (
	"fmt"

	"github.com/tzmfreedom/land/ast"
)

var pageReferenceType = &ast.ClassType{Name: "PageReference"}
var pageReferenceParameter = &ast.Parameter{
	Type: pageReferenceType,
	Name: "_",
}

var pageReferenceTypeRef = &ast.TypeRef{
	Name:       []string{"PageReference"},
	Parameters: []*ast.TypeRef{},
}

func createPageReferenceType() {
	instanceMethods := ast.NewMethodMap()
	instanceMethods.Set(
		"getUrl",
		[]*ast.Method{
			ast.CreateMethod(
				"getUrl",
				StringType,
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
	method.ReturnType = CreateMapType(StringType, StringType)

	instanceMethods.Set(
		"getParameters",
		[]*ast.Method{method},
	)
	instanceMethods.Set(
		"Equals",
		[]*ast.Method{
			ast.CreateMethod(
				"Equals",
				BooleanType,
				[]*ast.Parameter{pageReferenceParameter},
				func(this *ast.Object, params []*ast.Object, extra map[string]interface{}) interface{} {
					if params[0] == Null {
						return NewBoolean(false)
					}
					self := this.Extra["url"].(*ast.Object).StringValue()
					other := params[0].Extra["url"].(*ast.Object).StringValue()
					return NewBoolean(self == other)
				},
			),
		},
	)
	pageReferenceType.Constructors = []*ast.Method{
		{
			Modifiers:  []*ast.Modifier{ast.PublicModifier()},
			Parameters: []*ast.Parameter{stringTypeParameter},
			NativeFunction: func(this *ast.Object, params []*ast.Object, extra map[string]interface{}) interface{} {
				parameters := ast.CreateObject(CreateMapType(StringType, StringType))
				parameters.Extra["values"] = map[string]*ast.Object{}
				this.Extra = map[string]interface{}{
					"url":        params[0],
					"parameters": parameters,
				}
				return nil
			},
		},
	}
	pageReferenceType.InstanceFields = ast.NewFieldMap()
	pageReferenceType.InstanceMethods = instanceMethods
	pageReferenceType.StaticFields = ast.NewFieldMap()
	pageReferenceType.StaticMethods = ast.NewMethodMap()
	pageReferenceType.ToString = func(o *ast.Object) string {
		url := o.Extra["url"].(*ast.Object)
		parameters := o.Extra["parameters"].(*ast.Object)
		return fmt.Sprintf(
			"<%s> { url => %s, parameters => (%s) }",
			o.ClassType.Name,
			url.StringValue(),
			String(parameters),
		)
	}
}

func init() {
	createPageReferenceType()
	primitiveClassMap.Set("PageReference", pageReferenceType)
}
