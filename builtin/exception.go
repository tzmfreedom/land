package builtin

import (
	"fmt"

	"github.com/tzmfreedom/goland/ast"
)

var ExceptionType = createExceptionType()

var exceptionTypeParameter = &ast.Parameter{
	Type: &ast.TypeRef{
		Name:       []string{"Exception"},
		Parameters: []ast.Node{},
	},
	Name: "_",
}

func createExceptionType() *ClassType {
	instanceMethods := NewMethodMap()
	instanceMethods.Set(
		"getMessage",
		[]*Method{
			CreateMethod(
				"getMessage",
				stringTypeRef,
				[]ast.Node{},
				func(this *Object, params []*Object, extra map[string]interface{}) interface{} {
					return this.Extra["message"]
				},
			),
		},
	)

	classType := CreateClass(
		"Exception",
		[]*Method{
			{
				Modifiers:  []ast.Node{publicModifier},
				Parameters: []ast.Node{},
				NativeFunction: func(this *Object, params []*Object, extra map[string]interface{}) interface{} {
					this.Extra["message"] = Null
					this.Extra["exception"] = Null
					return nil
				},
			},
			{
				Modifiers:  []ast.Node{publicModifier},
				Parameters: []ast.Node{stringTypeParameter},
				NativeFunction: func(this *Object, params []*Object, extra map[string]interface{}) interface{} {
					this.Extra["message"] = params[0]
					this.Extra["exception"] = Null
					return nil
				},
			},
			{
				Modifiers:  []ast.Node{publicModifier},
				Parameters: []ast.Node{exceptionTypeParameter},
				NativeFunction: func(this *Object, params []*Object, extra map[string]interface{}) interface{} {
					this.Extra["message"] = Null
					this.Extra["exception"] = params[0]
					return nil
				},
			},
		},
		instanceMethods,
		nil,
	)
	classType.ToString = func(o *Object) string {
		return fmt.Sprintf("<%s> { message => %s } ", classType.Name, String(o.Extra["message"].(*Object)))
	}
	return classType
}

func init() {
	primitiveClassMap.Set("Exception", ExceptionType)
}
