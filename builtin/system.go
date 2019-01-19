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

type EqualChecker interface {
	Equals(*ast.Object, *ast.Object) bool
}

func init() {
	system := ast.CreateClass(
		"System",
		nil,
		nil,
		&ast.MethodMap{
			Data: map[string][]*ast.Method{
				"debug": {
					&ast.Method{
						Name:       "debug",
						Modifiers:  []*ast.Modifier{ast.PublicModifier()},
						Parameters: []*ast.Parameter{objectTypeParameter},
						NativeFunction: func(this *ast.Object, parameter []*ast.Object, extra map[string]interface{}) interface{} {
							o := parameter[0]
							stdout := extra["stdout"].(io.Writer)
							fmt.Fprintln(stdout, String(o))
							return nil
						},
					},
				},
				"assertequals": {
					&ast.Method{
						Name:      "assertequals",
						Modifiers: []*ast.Modifier{ast.PublicModifier()},
						Parameters: []*ast.Parameter{
							objectTypeParameter,
							objectTypeParameter,
						},
						NativeFunction: func(this *ast.Object, parameter []*ast.Object, extra map[string]interface{}) interface{} {
							expected := parameter[0]
							actual := parameter[1]
							checker := extra["interpreter"].(EqualChecker)
							if !checker.Equals(expected, actual) {
								node := extra["node"].(ast.Node)
								errors := extra["errors"].([]*TestError)
								message := fmt.Sprintf("      expected: %s\n      actual:   %s", String(expected), String(actual))
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
