package compiler

import (
	"strings"

	"github.com/tzmfreedom/goland/ast"
)

type Env interface{}

type Context struct {
	Env           *TypeEnv
	ClassTypes    *ClassMap
	NameSpaces    *TypeEnv
	CurrentMethod *ast.MethodDeclaration
	CurrentClass  *ast.ClassType
}

type VarEnv struct {
	Data   *NodeMap
	Parent *VarEnv
}

func newVarEnv(p *VarEnv) *VarEnv {
	return &VarEnv{
		Data:   &NodeMap{},
		Parent: p,
	}
}

func (e *VarEnv) Get(k string) (ast.Node, bool) {
	n, ok := e.Data.Get(k)
	if ok {
		return n, true
	}
	if e.Parent != nil {
		return e.Parent.Get(k)
	}
	return nil, false
}

func (e *VarEnv) Set(k string, n ast.Node) {
	e.Data.Set(k, n)
}

type NodeMap struct {
	Data map[string]ast.Node
}

func (m *NodeMap) Set(k string, n ast.Node) {
	m.Data[strings.ToLower(k)] = n
}

func (m *NodeMap) Get(k string) (ast.Node, bool) {
	n, ok := m.Data[strings.ToLower(k)]
	return n, ok
}

type TypeEnv struct {
	Data   *TypeMap
	Parent *TypeEnv
}

func newTypeEnv(p *TypeEnv) *TypeEnv {
	return &TypeEnv{
		Data:   newTypeMap(),
		Parent: p,
	}
}

func (e *TypeEnv) Get(k string) (ast.Type, bool) {
	n, ok := e.Data.Get(k)
	if ok {
		return n, true
	}
	if e.Parent != nil {
		return e.Parent.Get(k)
	}
	return nil, false
}

func (e *TypeEnv) Set(k string, n ast.Type) {
	e.Data.Set(k, n)
}

type TypeMap struct {
	Data map[string]ast.Type
}

func newTypeMap() *TypeMap {
	return &TypeMap{
		Data: map[string]ast.Type{},
	}
}

func (m *TypeMap) Set(k string, n ast.Type) {
	m.Data[strings.ToLower(k)] = n
}

func (m *TypeMap) Get(k string) (ast.Type, bool) {
	n, ok := m.Data[strings.ToLower(k)]
	return n, ok
}

type ClassMap struct {
	Data map[string]*ast.ClassType
}

func NewClassMap() *ClassMap {
	return &ClassMap{
		Data: map[string]*ast.ClassType{},
	}
}

func (m *ClassMap) Set(k string, n *ast.ClassType) {
	m.Data[strings.ToLower(k)] = n
}

func (m *ClassMap) Get(k string) (*ast.ClassType, bool) {
	n, ok := m.Data[strings.ToLower(k)]
	return n, ok
}
