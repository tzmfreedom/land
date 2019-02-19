package builtin

import (
	"math"
	"math/rand"
	"time"

	"github.com/tzmfreedom/land/ast"
)

func init() {
	instanceMethods := ast.NewMethodMap()
	staticMethods := ast.NewMethodMap()
	mathType := ast.CreateClass(
		"Math",
		[]*ast.Method{
			ast.CreateMethod(
				"Math",
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
		"floor",
		[]*ast.Method{
			ast.CreateMethod(
				"floor",
				DoubleType,
				[]*ast.Parameter{doubleTypeParameter},
				func(this *ast.Object, params []*ast.Object, extra map[string]interface{}) interface{} {
					return NewDouble(math.Floor(params[0].DoubleValue()))
				},
			),
		},
	)
	staticMethods.Set(
		"random",
		[]*ast.Method{
			ast.CreateMethod(
				"random",
				DoubleType,
				[]*ast.Parameter{},
				func(this *ast.Object, params []*ast.Object, extra map[string]interface{}) interface{} {
					rand.Seed(time.Now().UnixNano())
					return NewDouble(rand.Float64())
				},
			),
		},
	)

	primitiveClassMap.Set("Math", mathType)
}
