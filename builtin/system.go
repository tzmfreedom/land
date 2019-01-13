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
			Data: map[string][]ast.Node{
				"debug": {
					&ast.MethodDeclaration{
						Name: "debug",
						Modifiers: []ast.Node{
							&ast.Modifier{Name: "public"},
						},
						Parameters: []ast.Node{objectTypeParameter},
						NativeFunction: func(this interface{}, parameter []interface{}, options ...interface{}) interface{} {
							o := parameter[0].(*Object)
							stdout := options[0].(io.Writer)
							fmt.Fprintln(stdout, String(o))
							return nil
						},
					},
				},
				"assertequals": {
					&ast.MethodDeclaration{
						Name: "assertequals",
						Modifiers: []ast.Node{
							&ast.Modifier{Name: "public"},
						},
						NativeFunction: func(this interface{}, parameter []interface{}, options ...interface{}) interface{} {
							expected := parameter[0].(*Object)
							actual := parameter[1].(*Object)
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
