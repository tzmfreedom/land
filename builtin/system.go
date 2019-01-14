package builtin

import (
	"fmt"
	"io"

	"github.com/tzmfreedom/goland/ast"
)

func init() {
	system := CreateClass(
		"System",
		nil,
		nil,
		&MethodMap{
			Data: map[string][]*Method{
				"debug": {
					&Method{
						Name: "debug",
						Modifiers: []ast.Node{
							&ast.Modifier{Name: "public"},
						},
						Parameters: []ast.Node{objectTypeParameter},
						NativeFunction: func(this *Object, parameter []*Object, extra map[string]interface{}) interface{} {
							o := parameter[0]
							stdout := extra["stdout"].(io.Writer)
							fmt.Fprintln(stdout, String(o))
							return nil
						},
					},
				},
				"assertequals": {
					&Method{
						Name: "assertequals",
						Modifiers: []ast.Node{
							&ast.Modifier{Name: "public"},
						},
						Parameters: []ast.Node{
							objectTypeParameter,
							objectTypeParameter,
						},
						NativeFunction: func(this *Object, parameter []*Object, extra map[string]interface{}) interface{} {
							expected := parameter[0]
							actual := parameter[1]
							if expected.Value() != actual.Value() {
								fmt.Printf("expected: %s, actual: %s\n", String(expected), String(actual))
							}
							return nil
						},
					},
				},
			},
		},
	)

	primitiveClassMap.Set("system", system)
}
