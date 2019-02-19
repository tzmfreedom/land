package builtin

import (
	"regexp"

	"github.com/tzmfreedom/land/ast"
)

func init() {
	instanceMethods := ast.NewMethodMap()
	staticMethods := ast.NewMethodMap()
	patternType := ast.CreateClass(
		"Pattern",
		[]*ast.Method{
			ast.CreateMethod(
				"Pattern",
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

	staticMethods.Set(
		"matches",
		[]*ast.Method{
			ast.CreateMethod(
				"matches",
				BooleanType,
				[]*ast.Parameter{
					stringTypeParameter,
					stringTypeParameter,
				},
				func(this *ast.Object, params []*ast.Object, extra map[string]interface{}) interface{} {
					regExp := params[0].StringValue()
					stringtoMatch := params[1].StringValue()
					m, err := regexp.MatchString(regExp, stringtoMatch)
					if err != nil {
						panic(err) // TODO: impl
					}
					return NewBoolean(m)
				},
			),
		},
	)

	primitiveClassMap.Set("Pattern", patternType)
}
