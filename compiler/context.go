package compiler

import (
	"strings"

	"github.com/tzmfreedom/land/ast"
	"github.com/tzmfreedom/land/builtin"
)

type Env interface{}

type Context struct {
	Env         *TypeEnv
	StaticField *TypeMap
	ClassTypes  *ast.ClassMap           // loaded User Class
	NameSpaces  *builtin.NameSpaceStore // NameSpaces and Related Classes

	CurrentMethod *ast.Method
	CurrentClass  *ast.ClassType
}

func NewContext() *Context {
	ctx := &Context{}
	ctx.StaticField = newTypeMap()
	ctx.ClassTypes = ast.NewClassMap()
	ctx.NameSpaces = builtin.NewNameSpaceStore()
	ctx.Env = newTypeEnv(nil)
	return ctx
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

func (e *TypeEnv) Get(k string) (*ast.ClassType, bool) {
	n, ok := e.Data.Get(k)
	if ok {
		return n, true
	}
	if e.Parent != nil {
		return e.Parent.Get(k)
	}
	return nil, false
}

func (e *TypeEnv) Set(k string, n *ast.ClassType) {
	e.Data.Set(k, n)
}

/**
 * TypeMap
 */
type TypeMap struct {
	Data map[string]*ast.ClassType
}

func newTypeMap() *TypeMap {
	return &TypeMap{
		Data: map[string]*ast.ClassType{},
	}
}

func (m *TypeMap) Set(k string, n *ast.ClassType) {
	m.Data[strings.ToLower(k)] = n
}

func (m *TypeMap) Get(k string) (*ast.ClassType, bool) {
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

func (m *StaticFieldMap) Add(k string, n *ast.ClassType) {
	typeMap, ok := m.Get(k)
	if !ok {
		panic("StaticFieldMap#Add failed")
	}
	typeMap.Set(k, n)
}

func (m *StaticFieldMap) Set(k string, n *TypeMap) {
	m.Data[strings.ToLower(k)] = n
}

func (m *StaticFieldMap) Get(k string) (*TypeMap, bool) {
	n, ok := m.Data[strings.ToLower(k)]
	return n, ok
}
