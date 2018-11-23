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
	Data   *ObjectMap
	Parent *Env
}

func NewEnv(p *Env) *Env {
	return &Env{
		Data: &ObjectMap{
			Data: map[string]*Object{},
		},
		Parent: p,
	}
}

func (e *Env) Get(k string) (*Object, bool) {
	n, ok := e.Data.Get(k)
	if ok {
		return n, true
	}
	if e.Parent != nil {
		return e.Parent.Get(k)
	}
	return nil, false
}

func (e *Env) Set(k string, n *Object) {
	e.Data.Set(k, n)
}

/**
 * ObjectMap
 */
type ObjectMap struct {
	Data map[string]*Object
}

func NewObjectMap() *ObjectMap {
	return &ObjectMap{
		Data: map[string]*Object{},
	}
}

func (m *ObjectMap) Set(k string, n *Object) {
	m.Data[strings.ToLower(k)] = n
}

func (m *ObjectMap) Get(k string) (*Object, bool) {
	n, ok := m.Data[strings.ToLower(k)]
	return n, ok
}

/**
 * StaticFieldMap
 */
type StaticFieldMap struct {
	Data map[string]*ObjectMap
}

func NewStaticFieldMap() *StaticFieldMap {
	return &StaticFieldMap{
		Data: map[string]*ObjectMap{},
	}
}

func (m *StaticFieldMap) Add(k string, n *Object) {
	objMap, _ := m.Get(k)
	objMap.Set(k, n)
}

func (m *StaticFieldMap) Set(k string, n *ObjectMap) {
	m.Data[strings.ToLower(k)] = n
}

func (m *StaticFieldMap) Get(k string) (*ObjectMap, bool) {
	n, ok := m.Data[strings.ToLower(k)]
	return n, ok
}
