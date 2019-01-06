package builtin

import (
	"strings"

	"github.com/tzmfreedom/goland/ast"
)

var DoubleType = &ClassType{
	Name: "Double",
}

func init() {
	primitiveClassMap.Set("Double", DoubleType)
}
