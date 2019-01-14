package interpreter

import (
	"strings"

	"github.com/pkg/errors"
	"github.com/tzmfreedom/goland/builtin"
	"github.com/tzmfreedom/goland/compiler"
)

type TypeResolver struct {
	Context  *Context
	resolver *compiler.TypeResolver
}

func NewTypeResolver(ctx *Context) *TypeResolver {
	compilerContext := compiler.NewContext()
	compilerContext.ClassTypes = ctx.ClassTypes
	compilerContext.NameSpaces = ctx.NameSpaces
	return &TypeResolver{
		Context:  ctx,
		resolver: &compiler.TypeResolver{Context: compilerContext},
	}
}

func (r *TypeResolver) SetVariable(names []string, setValue *builtin.Object) error {
	if len(names) == 1 {
		if _, ok := r.Context.Env.Get(names[0]); ok {
			r.Context.Env.Set(names[0], setValue)
			return nil
		}
		// this
		if val, ok := r.Context.Env.Get("this"); ok {
			val.InstanceFields.Set(names[0], setValue)
			return nil
		}
		return errors.Errorf("%s is not found in this scope", names[0])
	} else {
		name := names[0]
		if val, ok := r.Context.Env.Get(name); ok {
			for _, f := range names[1 : len(names)-1] {
				val, ok = val.InstanceFields.Get(f)
				if !ok {
					return errors.Errorf("%s is not found in this scope", f)
				}
			}
			last := names[len(names)-1]
			_, ok = val.InstanceFields.Get(last)
			if !ok {
				return errors.Errorf("%s is not found in this scope", last)
			}
			val.InstanceFields.Set(last, setValue)
			return nil
		}
		// this
		if val, ok := r.Context.Env.Get("this"); ok {
			for _, f := range names[0 : len(names)-1] {
				val, ok = val.InstanceFields.Get(f)
				if !ok {
					return errors.Errorf("%s is not found in this scope", f)
				}
			}
			last := names[len(names)-1]
			_, ok = val.InstanceFields.Get(last)
			if !ok {
				return errors.Errorf("%s is not found in this scope", last)
			}
			val.InstanceFields.Set(last, setValue)
			return nil
		}
		if v, ok := r.Context.StaticField.Get("_", name); ok {
			if val, ok := v.Get(names[1]); ok {
				for _, f := range names[2:] {
					val, ok = val.InstanceFields.Get(f)
					if !ok {
						return errors.Errorf("%s is not found in this scope", f)
					}
				}
				return nil
			}
		}
		//if v, ok := r.Context.NameSpaces.Get(name); ok {
		//	if classType, ok := v.Get(names[1]); ok {
		//		if field, ok := classType.StaticFields.Get(names[2]); ok {
		//			t := field.Type.(*ast.TypeRef)
		//			fieldType, _ := r.ResolveType(t.Name, ctx)
		//			for _, f := range names[3:] {
		//				instanceField, ok := fieldType.InstanceFields.Get(f)
		//				if !ok {
		//					return nil, errors.Errorf("%s is not found in this scope", f)
		//				}
		//				fieldType, _ = r.ResolveType(instanceField.Type.(*ast.TypeRef).Name, ctx)
		//			}
		//			return fieldType, nil
		//		}
		//	}
		//}
	}
	return nil
}

func (r *TypeResolver) ResolveVariable(names []string) (*builtin.Object, error) {
	if len(names) == 1 {
		if v, ok := r.Context.Env.Get(names[0]); ok {
			return v, nil
		}
		// this
		if val, ok := r.Context.Env.Get("this"); ok {
			if val, ok = val.InstanceFields.Get(names[0]); ok {
				if val == nil {
					return nil, errors.Errorf("null pointer exception: %s", names[0])
				}
				return val, nil
			}
		}
		return nil, errors.Errorf("%s is not found in this scope", names[0])
	} else {
		name := names[0]
		if val, ok := r.Context.Env.Get(name); ok {
			for _, f := range names[1:] {
				if val == builtin.Null {
					return nil, errors.Errorf("null pointer exception: %s", f)
				}
				val, ok = val.InstanceFields.Get(f)
				if !ok {
					return nil, errors.Errorf("%s is not found in this scope", f)
				}
			}
			return val, nil
		}
		// this
		if val, ok := r.Context.Env.Get("this"); ok {
			for _, f := range names[0:] {
				val, ok = val.InstanceFields.Get(f)
				if !ok {
					return nil, errors.Errorf("%s is not found in this scope", f)
				}
				if val == nil {
					return nil, errors.Errorf("null pointer exception: %s", f)
				}
			}
			return val, nil
		}
		if v, ok := r.Context.StaticField.Get("_", name); ok {
			if val, ok := v.Get(names[1]); ok {
				for _, f := range names[2:] {
					val, ok = val.InstanceFields.Get(f)
					if !ok {
						return nil, errors.Errorf("%s is not found in this scope", f)
					}
					if val == nil {
						return nil, errors.Errorf("null pointer exception: %s", f)
					}
				}
				return val, nil
			}
		}
		if objMap, ok := r.Context.StaticField.Get(name, names[1]); ok {
			if val, ok := objMap.Get(names[2]); ok {
				for _, f := range names[3:] {
					val, ok = val.InstanceFields.Get(f)
					if !ok {
						return nil, errors.Errorf("%s is not found in this scope", f)
					}
					if val == nil {
						return nil, errors.Errorf("null pointer exception: %s", f)
					}
				}
				return val, nil
			}
		}
	}
	return nil, nil
}

func (r *TypeResolver) ResolveMethod(names []string, parameters []*builtin.Object) (interface{}, *builtin.Method, error) {
	if len(names) == 1 {
		methodName := names[0]
		if v, ok := r.Context.Env.Get("this"); ok {
			_, method, err := r.FindInstanceMethod(v, methodName, parameters, compiler.MODIFIER_ALL_OK)
			if err != nil {
				return nil, nil, err
			}
			if method == nil {
				return nil, nil, errors.Errorf("%s is not found in this scope", methodName)
			}
			return v, method, nil
		}
	} else {
		first := names[0]
		methodName := names[len(names)-1]
		fields := names[1 : len(names)-1]
		if val, ok := r.Context.Env.Get(first); ok {
			for _, f := range fields {
				val, ok = val.InstanceFields.Get(f)
				if !ok {
					return nil, nil, errors.Errorf("%s is not found in this scope", f)
				}
			}
			return r.FindInstanceMethod(val, methodName, parameters, compiler.MODIFIER_PUBLIC_ONLY)
		}
		if len(names) == 2 {
			if v, ok := r.Context.ClassTypes.Get(first); ok {
				return r.FindStaticMethod(v, methodName, parameters, compiler.MODIFIER_PUBLIC_ONLY)
			}
		}
		if len(names) >= 3 {
			// namespace.class.static_method()
			if len(names) == 3 {
				if v, ok := r.Context.NameSpaces.Get(first); ok {
					if classType, ok := v.Get(names[1]); ok {
						return r.FindStaticMethod(classType, methodName, parameters, compiler.MODIFIER_PUBLIC_ONLY)
					}
				}
			}

			if v, ok := r.Context.StaticField.Get("_", first); ok {
				if val, ok := v.Get(names[1]); ok {
					for _, f := range names[2 : len(names)-1] {
						val, ok = val.InstanceFields.Get(f)
						if !ok {
							return nil, nil, errors.Errorf("%s is not found in this scope", f)
						}
					}
					return r.FindInstanceMethod(val, methodName, parameters, compiler.MODIFIER_PUBLIC_ONLY)
				}
			}

			// namespace.class.static_field...instance_method()
			if len(names) > 3 {
				if objMap, ok := r.Context.StaticField.Get(first, names[1]); ok {
					if val, ok := objMap.Get(names[2]); ok {
						for _, f := range names[3 : len(names)-1] {
							val, ok = val.InstanceFields.Get(f)
							if !ok {
								return nil, nil, errors.Errorf("%s is not found in this scope", f)
							}
						}
						return r.FindInstanceMethod(val, methodName, parameters, compiler.MODIFIER_PUBLIC_ONLY)
					}
				}
			}
		}
	}
	return nil, nil, errors.Errorf("%s is not found in this scope", strings.Join(names, "."))
}

func (r *TypeResolver) ResolveType(names []string) (*builtin.ClassType, error) {
	if len(names) == 1 {
		className := names[0]
		if class, ok := r.Context.ClassTypes.Get(className); ok {
			return class, nil
		}
		if classTypes, ok := r.Context.NameSpaces.Get("System"); ok {
			if class, ok := classTypes.Get(className); ok {
				return class, nil
			}
		}
	} else if len(names) == 2 {
		// search for UserClass.InnerClass
		if class, ok := r.Context.ClassTypes.Get(names[0]); ok {
			if inner, ok := class.InnerClasses.Get(names[1]); ok {
				return inner, nil
			}
		}
		// search for NameSpace.UserClass
		if classTypes, ok := r.Context.NameSpaces.Get(names[0]); ok {
			if class, ok := classTypes.Get(names[1]); ok {
				return class, nil
			}
		}
	} else if len(names) == 3 {
		// search for NameSpace.UserClass.InnerClass
		if classTypes, ok := r.Context.NameSpaces.Get(names[0]); ok {
			if class, ok := classTypes.Get(names[1]); ok {
				if inner, ok := class.InnerClasses.Get(names[2]); ok {
					return inner, nil
				}
			}
		}
	}
	return nil, nil
}

func (r *TypeResolver) FindInstanceMethod(object *builtin.Object, methodName string, parameters []*builtin.Object, allowedModifier int) (*builtin.Object, *builtin.Method, error) {
	inputParameters := make([]*builtin.ClassType, len(parameters))
	for i, parameter := range parameters {
		inputParameters[i] = parameter.ClassType
	}
	_, method, err := r.resolver.FindInstanceMethod(object.ClassType, methodName, inputParameters, allowedModifier)
	return object, method, err
}

func (r *TypeResolver) FindStaticMethod(classType *builtin.ClassType, methodName string, parameters []*builtin.Object, allowedModifier int) (*builtin.ClassType, *builtin.Method, error) {
	inputParameters := make([]*builtin.ClassType, len(parameters))
	for i, parameter := range parameters {
		inputParameters[i] = parameter.ClassType
	}
	return r.resolver.FindStaticMethod(classType, methodName, inputParameters, allowedModifier)
}
