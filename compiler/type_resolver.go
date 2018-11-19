package compiler

type TypeResolver struct{}

func (r *TypeResolver) ResolveVariable(names []string, ctx *Context) (interface{}, error) {
	if len(names) == 1 {
		if v, ok := ctx.Env.Get(names[0]); ok {
			return v, nil
		}
	} else {
		name := names[0]
		if v, ok := ctx.Env.Get(name); ok {
			// lookup
			return v, nil
		} else if v, ok := ctx.ClassTypes.Get(name); ok {
			// lookup class.static
			return v, nil
		} else if v, ok := ctx.NameSpaces.Get(name); ok {
			// lookup namespace.class
			return v, nil
		}
	}
	return nil, nil
}

func (r *TypeResolver) ResolveMethod(names []string, ctx *Context) (interface{}, error) {
	// name node or field name node
	if len(names) == 1 {
		if v, ok := ctx.Env.Get(names[0]); ok {
			return v, nil
		}
	} else {
		name := names[0]
		if v, ok := ctx.Env.Get(name); ok {
			// lookup
			return v, nil
		} else if v, ok := ctx.ClassTypes.Get(name); ok {
			// lookup class.static
			return v, nil
		} else if v, ok := ctx.NameSpaces.Get(name); ok {
			// lookup namespace.class
			return v, nil
		}
	}
	return nil, nil
}
