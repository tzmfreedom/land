package builtin

import (
	"fmt"

	"github.com/tzmfreedom/goland/ast"
)

var ExceptionType = createExceptionType()

var exceptionTypeParameter = &ast.Parameter{
	TypeRef: &ast.TypeRef{
		Name:       []string{"Exception"},
		Parameters: []ast.Node{},
	},
	Name: "_",
}

func createExceptionType() *ast.ClassType {
	instanceMethods := ast.NewMethodMap()
	instanceMethods.Set(
		"getMessage",
		[]*ast.Method{
			ast.CreateMethod(
				"getMessage",
				stringTypeRef,
				[]*ast.Parameter{},
				func(this *ast.Object, params []*ast.Object, extra map[string]interface{}) interface{} {
					return this.Extra["message"]
				},
			),
		},
	)

	classType := ast.CreateClass(
		"Exception",
		[]*ast.Method{
			{
				Modifiers:  []*ast.Modifier{ast.PublicModifier()},
				Parameters: []*ast.Parameter{},
				NativeFunction: func(this *ast.Object, params []*ast.Object, extra map[string]interface{}) interface{} {
					this.Extra["message"] = Null
					this.Extra["exception"] = Null
					return nil
				},
			},
			{
				Modifiers:  []*ast.Modifier{ast.PublicModifier()},
				Parameters: []*ast.Parameter{stringTypeParameter},
				NativeFunction: func(this *ast.Object, params []*ast.Object, extra map[string]interface{}) interface{} {
					this.Extra["message"] = params[0]
					this.Extra["exception"] = Null
					return nil
				},
			},
			{
				Modifiers:  []*ast.Modifier{ast.PublicModifier()},
				Parameters: []*ast.Parameter{exceptionTypeParameter},
				NativeFunction: func(this *ast.Object, params []*ast.Object, extra map[string]interface{}) interface{} {
					this.Extra["message"] = Null
					this.Extra["exception"] = params[0]
					return nil
				},
			},
		},
		instanceMethods,
		nil,
	)
	classType.ToString = func(o *ast.Object) string {
		return fmt.Sprintf("<%s> { message => %s } ", classType.Name, String(o.Extra["message"].(*ast.Object)))
	}
	return classType
}

func init() {
	primitiveClassMap.Set("Exception", ExceptionType)
}
