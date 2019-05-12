package builtin

import "github.com/tzmfreedom/land/ast"

type NullPointerException struct {
	name string
	location *ast.Location
}

func NewNullPointerException(name string) *NullPointerException {
	return &NullPointerException{name: name}
}

func (e *NullPointerException) Error() string {
	return "null pointer exception"
}

func (e *NullPointerException) GetName() string {
	return e.name
}
