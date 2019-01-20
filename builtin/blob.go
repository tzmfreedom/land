package builtin

import (
	"time"

	"github.com/tzmfreedom/goland/ast"
)

var BlobType = &ast.ClassType{Name: "Blob"}

func init() {
	instanceMethods := ast.NewMethodMap()
	staticMethods := ast.NewMethodMap()
	BlobType.Constructors = []*ast.Method{}
	BlobType.InstanceMethods = instanceMethods
	BlobType.StaticMethods = staticMethods

	instanceMethods.Set(
		"toString",
		[]*ast.Method{
			ast.CreateMethod(
				"toString",
				StringType,
				[]*ast.Parameter{},
				func(this *ast.Object, params []*ast.Object, extra map[string]interface{}) interface{} {
					bytes := this.Extra["value"].([]byte)
					return NewString(string(bytes))
				},
			),
		},
	)
	instanceMethods.Set(
		"size",
		[]*ast.Method{
			ast.CreateMethod(
				"size",
				IntegerType,
				[]*ast.Parameter{},
				func(this *ast.Object, params []*ast.Object, extra map[string]interface{}) interface{} {
					bytes := this.Extra["value"].([]byte)
					return NewInteger(len(bytes))
				},
			),
		},
	)
	staticMethods.Set(
		"valueOf",
		[]*ast.Method{
			ast.CreateMethod(
				"valueOf",
				BlobType,
				[]*ast.Parameter{
					stringTypeParameter,
				},
				func(this *ast.Object, params []*ast.Object, extra map[string]interface{}) interface{} {
					obj := ast.CreateObject(BlobType)
					value := params[0].StringValue()
					obj.Extra["value"] = []byte(value)
					return obj
				},
			),
		},
	)

	primitiveClassMap.Set("Blob", BlobType)
}
