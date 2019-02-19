package interpreter

import (
	"strings"

	"github.com/tzmfreedom/land/ast"
	"github.com/tzmfreedom/land/builtin"
)

type Context struct {
	Env         *Env
	StaticField *StaticFieldMap
	ClassTypes  *ast.ClassMap           // loaded User Class
	NameSpaces  *builtin.NameSpaceStore // NameSpaces and Related Classes

	CurrentMethod *ast.MethodDeclaration
	CurrentClass  *ast.ClassType
}

func NewContext() *Context {
	ctx := &Context{}
	ctx.StaticField = NewStaticFieldMap()
	ctx.ClassTypes = ast.NewClassMap()
	ctx.NameSpaces = builtin.NewNameSpaceStore()
	ctx.Env = NewEnv(nil)
	return ctx
}

type Env struct {
	Data   *ast.ObjectMap
	Parent *Env
}

func NewEnv(p *Env) *Env {
	return &Env{
		Data:   ast.NewObjectMap(),
		Parent: p,
	}
}

func (e *Env) Get(k string) (*ast.Object, bool) {
	n, ok := e.Data.Get(k)
	if ok {
		return n, true
	}
	if e.Parent != nil {
		return e.Parent.Get(k)
	}
	return nil, false
}

func (e *Env) Update(k string, n *ast.Object) {
	env := e.FindEnv(k)
	if env == nil {
		panic("field not found: " + k)
	}
	env.Data.Set(k, n)
}

func (e *Env) Define(k string, n *ast.Object) {
	e.Data.Set(k, n)
}

func (e *Env) FindEnv(k string) *Env {
	_, ok := e.Data.Get(k)
	if ok {
		return e
	}
	if e.Parent != nil {
		return e.Parent.FindEnv(k)
	}
	return nil
}

/**
 * StaticFieldMap
 */
type StaticFieldMap struct {
	Data map[string]map[string]*ast.ObjectMap
}

func NewStaticFieldMap() *StaticFieldMap {
	return &StaticFieldMap{
		Data: map[string]map[string]*ast.ObjectMap{},
	}
}

func (m *StaticFieldMap) Add(ns, k, f string, n *ast.Object) {
	objMap, ok := m.Get(ns, k)
	if !ok {
		panic("StaticFieldMap#Add failed")
	}
	objMap.Set(f, n)
}

func (m *StaticFieldMap) Set(ns, k string, n *ast.ObjectMap) {
	ns = strings.ToLower(ns)
	k = strings.ToLower(k)
	if _, ok := m.Data[ns]; !ok {
		m.Data[ns] = map[string]*ast.ObjectMap{}
	}
	m.Data[ns][k] = n
}

func (m *StaticFieldMap) Get(ns, k string) (*ast.ObjectMap, bool) {
	ns = strings.ToLower(ns)
	k = strings.ToLower(k)
	if objMap, ok := m.Data[ns]; ok {
		if n, ok := objMap[k]; ok {
			return n, ok
		}
	}
	return nil, false
}
