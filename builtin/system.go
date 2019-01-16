package builtin

import (
	"fmt"
	"io"

	"github.com/tzmfreedom/goland/ast"
)

const (
	InfoColor    = "\033[1;34m%s\033[0m"
	NoticeColor  = "\033[1;36m%s\033[0m"
	WarningColor = "\033[1;33m%s\033[0m"
	ErrorColor   = "\033[1;31m%s\033[0m"
	DebugColor   = "\033[0;36m%s\033[0m"
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
								node := extra["node"].(ast.Node)
								errors := extra["errors"].([]*TestError)
								message := fmt.Sprintf("      expected: %s\n      actual:   %s", expected.Value(), actual.Value())
								extra["errors"] = append(errors, &TestError{
									Node:    node,
									Message: message,
								})
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

type TestError struct {
	Node    ast.Node
	Message string
}
