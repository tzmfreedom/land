package compiler

import (
	"strings"

	"github.com/tzmfreedom/goland/ast"
)

type Env struct {
	Data   []NodeMap
	Parent *Env
}

type NodeMap struct {
	Data map[string]ast.Node
}

func (m *NodeMap) Set(k string, n ast.Node) {
	m.Data[strings.ToLower(k)] = n
}

func (m *NodeMap) Get(k string) ast.Node {
	return m.Data[strings.ToLower(k)]
}

func (m *NodeMap) Contains(k string) bool {
	_, ok := m.Data[strings.ToLower(k)]
	return ok
}
