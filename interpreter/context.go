package interpreter

import (
	"strings"

	"github.com/tzmfreedom/goland/ast"
	"github.com/tzmfreedom/goland/builtin"
)

type Context struct {
	Env         *Env
	StaticField *StaticFieldMap
	ClassTypes  *builtin.ClassMap       // loaded User Class
	NameSpaces  *builtin.NameSpaceStore // NameSpaces and Related Classes

	CurrentMethod *ast.MethodDeclaration
	CurrentClass  *builtin.ClassType
}

func NewContext() *Context {
	ctx := &Context{}
	ctx.StaticField = NewStaticFieldMap()
	ctx.ClassTypes = builtin.NewClassMap()
	ctx.NameSpaces = builtin.NewNameSpaceStore()
	ctx.Env = NewEnv(nil)
	return ctx
}

type Env struct {
	Data   *builtin.ObjectMap
	Parent *Env
}

func NewEnv(p *Env) *Env {
	return &Env{
		Data: &builtin.ObjectMap{
			Data: map[string]*builtin.Object{},
		},
		Parent: p,
	}
}

func (e *Env) Get(k string) (*builtin.Object, bool) {
	n, ok := e.Data.Get(k)
	if ok {
		return n, true
	}
	if e.Parent != nil {
		return e.Parent.Get(k)
	}
	return nil, false
}

func (e *Env) Set(k string, n *builtin.Object) {
	e.Data.Set(k, n)
}

/**
 * StaticFieldMap
 */
type StaticFieldMap struct {
	Data map[string]*builtin.ObjectMap
}

func NewStaticFieldMap() *StaticFieldMap {
	return &StaticFieldMap{
		Data: map[string]*builtin.ObjectMap{},
	}
}

func (m *StaticFieldMap) Add(k string, n *builtin.Object) {
	objMap, _ := m.Get(k)
	objMap.Set(k, n)
}

func (m *StaticFieldMap) Set(k string, n *builtin.ObjectMap) {
	m.Data[strings.ToLower(k)] = n
}

func (m *StaticFieldMap) Get(k string) (*builtin.ObjectMap, bool) {
	n, ok := m.Data[strings.ToLower(k)]
	return n, ok
}
