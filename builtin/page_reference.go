package builtin

import (
	"fmt"

	"github.com/tzmfreedom/goland/ast"
)

var pageReferenceType = createPageReferenceType()

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
					thisObj.Extra = map[string]interface{}{
						"url": params[0].(*Object),
					}
				},
			},
		},
		instanceMethods,
		nil,
	)
	classType.ToString = func(o *Object) string {
		url := o.Extra["url"].(*Object)
		return fmt.Sprintf("<%s> { url => %s }", o.ClassType.Name, url.StringValue())
	}
	return classType
}

func init() {
	primitiveClassMap.Set("PageReference", pageReferenceType)
}
