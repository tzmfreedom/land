package compiler

import (
	"strings"

	"github.com/tzmfreedom/goland/ast"
)

type Env interface{}

type Context struct {
	Env         *TypeEnv
	StaticField *TypeMap
	ClassTypes  *ClassMap       // loaded User Class
	NameSpaces  *NameSpaceStore // NameSpaces and Related Classes

	CurrentMethod *ast.MethodDeclaration
	CurrentClass  *ClassType
}

/**
 * VarEnv
 */
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

/**
 * NodeMap
 */
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

/**
 * TypeEnv Map
 */
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

func (e *TypeEnv) Get(k string) (Type, bool) {
	n, ok := e.Data.Get(k)
	if ok {
		return n, true
	}
	if e.Parent != nil {
		return e.Parent.Get(k)
	}
	return nil, false
}

func (e *TypeEnv) Set(k string, n Type) {
	e.Data.Set(k, n)
}

/**
 * TypeMap
 */
type TypeMap struct {
	Data map[string]Type
}

func newTypeMap() *TypeMap {
	return &TypeMap{
		Data: map[string]Type{},
	}
}

func (m *TypeMap) Set(k string, n Type) {
	m.Data[strings.ToLower(k)] = n
}

func (m *TypeMap) Get(k string) (Type, bool) {
	n, ok := m.Data[strings.ToLower(k)]
	return n, ok
}

/**
 * NameSpaces
 */
type NameSpaceStore struct {
	Data map[string]*ClassMap
}

func NewNameSpaceStore() *NameSpaceStore {
	return &NameSpaceStore{
		Data: map[string]*ClassMap{},
	}
}

func (m *NameSpaceStore) Add(k string, n *ClassType) {
	classMap, _ := m.Get(k)
	classMap.Set(k, n)
}

func (m *NameSpaceStore) Set(k string, n *ClassMap) {
	m.Data[strings.ToLower(k)] = n
}

func (m *NameSpaceStore) Get(k string) (*ClassMap, bool) {
	n, ok := m.Data[strings.ToLower(k)]
	return n, ok
}

/**
 * StaticFieldMap
 */
type StaticFieldMap struct {
	Data map[string]*TypeMap
}

func NewStaticFieldMap() *StaticFieldMap {
	return &StaticFieldMap{
		Data: map[string]*TypeMap{},
	}
}

func (m *StaticFieldMap) Add(k string, n *Type) {
	typeMap, _ := m.Get(k)
	typeMap.Set(k, n)
}

func (m *StaticFieldMap) Set(k string, n *TypeMap) {
	m.Data[strings.ToLower(k)] = n
}

func (m *StaticFieldMap) Get(k string) (*TypeMap, bool) {
	n, ok := m.Data[strings.ToLower(k)]
	return n, ok
}
