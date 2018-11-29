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
	Data map[string]map[string]*builtin.ObjectMap
}

func NewStaticFieldMap() *StaticFieldMap {
	return &StaticFieldMap{
		Data: map[string]map[string]*builtin.ObjectMap{},
	}
}

func (m *StaticFieldMap) Add(ns, k, f string, n *builtin.Object) {
	objMap, _ := m.Get(ns, k)
	objMap.Set(f, n)
}

func (m *StaticFieldMap) Set(ns, k string, n *builtin.ObjectMap) {
	ns = strings.ToLower(ns)
	k = strings.ToLower(k)
	if _, ok := m.Data[ns]; !ok {
		m.Data[ns] = map[string]*builtin.ObjectMap{}
	}
	m.Data[ns][k] = n
}

func (m *StaticFieldMap) Get(ns, k string) (*builtin.ObjectMap, bool) {
	ns = strings.ToLower(ns)
	k = strings.ToLower(k)
	if objMap, ok := m.Data[ns]; ok {
		if n, ok := objMap[k]; ok {
			return n, ok
		}
	}
	return nil, false
}
