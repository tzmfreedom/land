package builtin

import (
	"strings"

	"github.com/tzmfreedom/goland/ast"
)

var IntegerType = &ClassType{
	Name: "Integer",
}

func init() {
	primitiveClassMap.Set("Integer", IntegerType)
}
