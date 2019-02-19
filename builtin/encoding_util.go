package builtin

import (
	"encoding/base64"

	"net/url"

	"github.com/tzmfreedom/land/ast"
)

func init() {
	instanceMethods := ast.NewMethodMap()
	staticMethods := ast.NewMethodMap()
	encodingUtilType := ast.CreateClass(
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
				BlobType,
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
				StringType,
				[]*ast.Parameter{BlobTypeParameter},
				func(this *ast.Object, params []*ast.Object, extra map[string]interface{}) interface{} {
					msg := params[0].Value().([]byte)
					encoded := base64.StdEncoding.EncodeToString(msg)
					return NewString(encoded)
				},
			),
		},
	)

	staticMethods.Set(
		"urlEncode",
		[]*ast.Method{
			ast.CreateMethod(
				"urlEncode",
				StringType,
				[]*ast.Parameter{
					stringTypeParameter,
					stringTypeParameter,
				},
				func(this *ast.Object, params []*ast.Object, extra map[string]interface{}) interface{} {
					inputString := params[0].StringValue()
					// encodingScheme := params[1].StringValue()
					return NewString(url.QueryEscape(inputString))
				},
			),
		},
	)
	primitiveClassMap.Set("EncodingUtil", encodingUtilType)
}
