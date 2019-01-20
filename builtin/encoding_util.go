package builtin

import (
	"encoding/base64"

	"github.com/tzmfreedom/goland/ast"
)

func init() {
	instanceMethods := ast.NewMethodMap()
	staticMethods := ast.NewMethodMap()
	dateType := ast.CreateClass(
		"EncodingUtil",
		[]*ast.Method{},
		instanceMethods,
		staticMethods,
	)

	staticMethods.Set(
		"base64Decode",
		[]*ast.Method{
			ast.CreateMethod(
				"base64Decode",
				StringType,
				[]*ast.Parameter{stringTypeParameter},
				func(this *ast.Object, params []*ast.Object, extra map[string]interface{}) interface{} {
					msg := params[0].StringValue()
					encoded, err := base64.StdEncoding.DecodeString(msg)
					if err != nil {
						panic(err) // TODO: impl
					}
					return NewBlob(encoded)
				},
			),
		},
	)

	staticMethods.Set(
		"base64Encode",
		[]*ast.Method{
			ast.CreateMethod(
				"base64Encode",
				dateType,
				[]*ast.Parameter{BlobTypeParameter},
				func(this *ast.Object, params []*ast.Object, extra map[string]interface{}) interface{} {
					msg := params[0].Value().([]byte)
					encoded := base64.StdEncoding.EncodeToString(msg)
					return NewString(encoded)
				},
			),
		},
	)

	primitiveClassMap.Set("Date", dateType)
}
