package builtin

import "github.com/tzmfreedom/goland/ast"

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

	primitiveClassMap.Set("HttpResponse", httpResponseType)
}
