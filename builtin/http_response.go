package builtin

import "github.com/tzmfreedom/land/ast"

var httpResponseType = &ast.ClassType{Name: "HttpResponse"}
var httpResponseTypeParameter = &ast.Parameter{
	Type: httpRequestType,
	Name: "_",
}

func init() {
	instanceMethods := ast.NewMethodMap()
	staticMethods := ast.NewMethodMap()
	httpResponseType.Constructors = []*ast.Method{
		ast.CreateMethod(
			"HttpResponse",
			nil,
			[]*ast.Parameter{},
			func(this *ast.Object, params []*ast.Object, extra map[string]interface{}) interface{} {
				return nil
			},
		),
	}
	httpResponseType.InstanceMethods = instanceMethods
	httpResponseType.StaticMethods = staticMethods

	instanceMethods.Set(
		"getBody",
		[]*ast.Method{
			ast.CreateMethod(
				"getBody",
				StringType,
				[]*ast.Parameter{},
				func(this *ast.Object, params []*ast.Object, extra map[string]interface{}) interface{} {
					return NewString(this.Extra["body"].(string))
				},
			),
		},
	)

	primitiveClassMap.Set("HttpResponse", httpResponseType)
}
