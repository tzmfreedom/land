package compiler

import "github.com/tzmfreedom/goland/ast"

type TypeResolver struct{}

func (r *TypeResolver) ResolveVariable(names []string, ctx *Context) (interface{}, error) {
	if len(names) == 1 {
		if v, ok := ctx.Env.Get(names[0]); ok {
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

	}
	return nil, nil
}
