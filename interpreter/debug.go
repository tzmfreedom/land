package interpreter

import (
	"github.com/k0kubun/pp"
)

func debug(args ...interface{}) {
	pp.Println(args...)
}

