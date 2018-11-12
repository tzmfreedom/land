package ast

import (
	"strings"
)

type Context struct {
	Env           *Env
	CurrentMethod MethodDeclaration
	CurrentClass  ClassType
}

type Env struct {
	Data   []NodeMap
	Parent *Env
}

type NodeMap struct {
	Data map[string]Node
}

func (m *NodeMap) Set(k string, n Node) {
	m.Data[strings.ToLower(k)] = n
}

func (m *NodeMap) Get(k string) Node {
	return m.Data[strings.ToLower(k)]
}

func (m *NodeMap) Contains(k string) bool {
	_, ok := m.Data[strings.ToLower(k)]
	return ok
}
