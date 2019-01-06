package builtin

import (
	"strings"

	"github.com/tzmfreedom/goland/ast"
)

var BooleanType = &ClassType{
	Name: "Boolean",
}

func init() {
	primitiveClassMap.Set("Boolean", BooleanType)
}
