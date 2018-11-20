package compiler

import (
	"strings"

	"github.com/pkg/errors"
	"github.com/tzmfreedom/goland/ast"
)

type TypeResolver struct{}

func (r *TypeResolver) ResolveVariable(names []string, ctx *Context) (interface{}, error) {
	if len(names) == 1 {
		if v, ok := ctx.Env.Get(names[0]); ok {
			return v, nil
		}
		return nil, errors.Errorf("%s is not found in this scope", names[0])
	} else {
		name := names[0]
		if fieldType, ok := ctx.Env.Get(name); ok {
			for _, f := range names[1:] {
				instanceField, ok := fieldType.(*ast.ClassType).InstanceFields.Get(f)
				if !ok {
					return nil, errors.Errorf("%s is not found in this scope", f)
				}
				fieldType, _ = r.ResolveType(instanceField.Type.(*ast.TypeRef).Name, ctx)
			}
			return fieldType, nil
		}
		if v, ok := ctx.ClassTypes.Get(name); ok {
			if n, ok := v.StaticFields.Get(names[1]); ok {
				t := n.Type.(*ast.TypeRef)
				fieldType, _ := r.ResolveType(t.Name, ctx)
				for _, f := range names[2:] {
					instanceField, ok := fieldType.(*ast.ClassType).InstanceFields.Get(f)
					if !ok {
						return nil, errors.Errorf("%s is not found in this scope", f)
					}
					fieldType, _ = r.ResolveType(instanceField.Type.(*ast.TypeRef).Name, ctx)
				}
				return fieldType, nil
			}
		}
		if v, ok := ctx.NameSpaces.Get(name); ok {
			if classType, ok := v.Get(names[1]); ok {
				if field, ok := classType.StaticFields.Get(names[2]); ok {
					t := field.Type.(*ast.TypeRef)
					fieldType, _ := r.ResolveType(t.Name, ctx)
					for _, f := range names[3:] {
						instanceField, ok := fieldType.(*ast.ClassType).InstanceFields.Get(f)
						if !ok {
							return nil, errors.Errorf("%s is not found in this scope", f)
						}
						fieldType, _ = r.ResolveType(instanceField.Type.(*ast.TypeRef).Name, ctx)
					}
					return fieldType, nil
				}
			}
		}
	}
	return nil, nil
}

func (r *TypeResolver) ResolveMethod(names []string, ctx *Context) (interface{}, error) {
	if len(names) == 1 {
		// this.method()
		if v, ok := ctx.Env.Get("this"); ok {
			return v, nil
		}
	} else {
		name := names[0]
		if fieldType, ok := ctx.Env.Get(name); ok {
			for _, f := range names[1:] {
				instanceField, _ := fieldType.(*ast.ClassType).InstanceFields.Get(f)
				fieldType, _ = r.ResolveType(instanceField.Type.(*ast.TypeRef).Name, ctx)
			}
			return fieldType, nil
		}
		if v, ok := ctx.ClassTypes.Get(name); ok {
			n, _ := v.StaticFields.Get(names[1])
			t := n.Type.(*ast.TypeRef)
			fieldType, _ := r.ResolveType(t.Name, ctx)
			for _, f := range names[2:] {
				instanceField, _ := fieldType.(*ast.ClassType).InstanceFields.Get(f)
				fieldType, _ = r.ResolveType(instanceField.Type.(*ast.TypeRef).Name, ctx)
			}
			return fieldType, nil
		}
		if v, ok := ctx.NameSpaces.Get(name); ok {
			classType, _ := v.Get(names[1])
			field, _ := classType.StaticFields.Get(names[2])
			t := field.Type.(*ast.TypeRef)
			fieldType, _ := r.ResolveType(t.Name, ctx)
			for _, f := range names[3:] {
				instanceField, _ := fieldType.(*ast.ClassType).InstanceFields.Get(f)
				fieldType, _ = r.ResolveType(instanceField.Type.(*ast.TypeRef).Name, ctx)
			}
			return fieldType, nil
		}
	}
	return nil, nil
}

func (r *TypeResolver) ResolveType(names []string, ctx *Context) (interface{}, error) {
	if len(names) == 1 {
		if strings.ToLower(names[0]) == "integer" {
			return ast.IntegerType, nil
		}
		class, _ := ctx.ClassTypes.Get(names[0])
		return class, nil
		// search for Primitive or System or UserClass
	} else if len(names) == 2 {
		// search for UserClass.InnerClass
		// search for NameSpace.UserClass
	} else if len(names) == 3 {
		// search for NameSpace.UserClass.InnerClass
	}
	return nil, nil
}
