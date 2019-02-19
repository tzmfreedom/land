package builtin

import (
	"io/ioutil"
	"net/http"

	"strings"

	"github.com/tzmfreedom/land/ast"
)

func init() {
	instanceMethods := ast.NewMethodMap()
	staticMethods := ast.NewMethodMap()
	httpType := ast.CreateClass(
		"Http",
		[]*ast.Method{
			ast.CreateMethod(
				"Http",
				nil,
				[]*ast.Parameter{},
				func(this *ast.Object, params []*ast.Object, extra map[string]interface{}) interface{} {
					return nil
				},
			),
		},
		instanceMethods,
		staticMethods,
	)

	instanceMethods.Set(
		"send",
		[]*ast.Method{
			ast.CreateMethod(
				"send",
				httpResponseType,
				[]*ast.Parameter{
					httpRequestTypeParameter,
				},
				func(this *ast.Object, params []*ast.Object, extra map[string]interface{}) interface{} {
					request := params[0]
					endpoint := request.Extra["endpoint"].(string)
					method := request.Extra["method"].(string)
					headers := request.Extra["headers"].(map[string]*ast.Object)
					body := request.Extra["body"].(string)
					req, err := http.NewRequest(method, endpoint, strings.NewReader(body))
					if err != nil {
						panic(err) // TODO: impl
					}
					for header, value := range headers {
						req.Header.Add(header, value.StringValue())
					}
					client := &http.Client{}
					res, err := client.Do(req)
					if err != nil {
						panic(err) // TODO: impl
					}
					buf, err := ioutil.ReadAll(res.Body)
					if err != nil {
						panic(err)
					}
					responseObj := ast.CreateObject(httpResponseType)
					responseObj.Extra["body"] = string(buf)
					return responseObj
				},
			),
		},
	)

	primitiveClassMap.Set("Http", httpType)
}
