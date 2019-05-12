package interpreter

import (
	"strings"

	"github.com/pkg/errors"
	"github.com/tzmfreedom/land/ast"
	"github.com/tzmfreedom/land/builtin"
	"github.com/tzmfreedom/land/compiler"
)

type TypeResolver struct {
	Context  *Context
	resolver *compiler.TypeResolver
}

func NewTypeResolver(ctx *Context) *TypeResolver {
	compilerContext := compiler.NewContext()
	compilerContext.ClassTypes = ctx.ClassTypes
	compilerContext.NameSpaces = ctx.NameSpaces
	compilerContext.CurrentClass = ctx.CurrentClass
	return &TypeResolver{
		Context:  ctx,
		resolver: compiler.NewTypeResolver(compilerContext),
	}
}

func (r *TypeResolver) SetVariable(names []string, setValue *ast.Object) error {
	if len(names) == 1 {
		if _, ok := r.Context.Env.Get(names[0]); ok {
			r.Context.Env.Update(names[0], setValue)
			return nil
		}
		// this
		if val, ok := r.Context.Env.Get("this"); ok {
			return setVariable(val, names[0], setValue)
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
			return setVariable(val, last, setValue)
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
			return setVariable(val, last, setValue)
		}
		if v, ok := r.Context.StaticField.Get("_", name); ok {
			if val, ok := v.Get(names[1]); ok {
				for _, f := range names[2 : len(names)-1] {
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
				return setVariable(val, last, setValue)
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

func setVariable(receiver *ast.Object, name string, value *ast.Object) error {
	v, ok := receiver.InstanceFields.Get(name)
	if !ok {
		panic("InstanceFields#Get failed")
	}
	if v.Final {
		return errors.New("Final variable has already been initialized")
	}
	receiver.InstanceFields.Set(name, value)
	return nil
}

func (r *TypeResolver) ResolveVariable(names []string) (*ast.Object, error) {
	if len(names) == 1 {
		if v, ok := r.Context.Env.Get(names[0]); ok {
			return v, nil
		}
		// this
		if val, ok := r.Context.Env.Get("this"); ok {
			if val, ok = val.InstanceFields.Get(names[0]); ok {
				if val == nil {
					return nil, builtin.NewNullPointerException(names[0])
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
					return nil, builtin.NewNullPointerException(f)
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
					return nil, builtin.NewNullPointerException(f)
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
						return nil, builtin.NewNullPointerException(f)
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
						return nil, builtin.NewNullPointerException(f)
					}
				}
				return val, nil
			}
		}
	}
	return nil, nil
}

func (r *TypeResolver) ResolveMethod(names []string, parameters []*ast.Object) (interface{}, *ast.Method, error) {
	if len(names) == 1 {
		methodName := names[0]
		if v, ok := r.Context.Env.Get("this"); ok {
			_, method, err := FindInstanceMethod(v, methodName, parameters, compiler.MODIFIER_ALL_OK)
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
			return FindInstanceMethod(val, methodName, parameters, compiler.MODIFIER_ALL_OK)
		}
		if len(names) == 2 {
			if v, ok := r.Context.ClassTypes.Get(first); ok {
				return FindStaticMethod(v, methodName, parameters, compiler.MODIFIER_ALL_OK)
			}
		}
		if len(names) >= 3 {
			// namespace.class.static_method()
			if len(names) == 3 {
				if v, ok := r.Context.NameSpaces.Get(first); ok {
					if classType, ok := v.Get(names[1]); ok {
						return FindStaticMethod(classType, methodName, parameters, compiler.MODIFIER_ALL_OK)
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
					return FindInstanceMethod(val, methodName, parameters, compiler.MODIFIER_ALL_OK)
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
						return FindInstanceMethod(val, methodName, parameters, compiler.MODIFIER_ALL_OK)
					}
				}
			}
		}
	}
	return nil, nil, errors.Errorf("%s is not found in this scope", strings.Join(names, "."))
}

func (r *TypeResolver) ResolveType(names []string) (*ast.ClassType, error) {
	return r.resolver.ResolveType(names)
}

func FindInstanceMethod(object *ast.Object, methodName string, parameters []*ast.Object, allowedModifier int) (*ast.Object, *ast.Method, error) {
	_, method, err := compiler.FindInstanceMethod(object.ClassType, methodName, convertClassTypes(parameters), allowedModifier)
	return object, method, err
}

func FindStaticMethod(classType *ast.ClassType, methodName string, parameters []*ast.Object, allowedModifier int) (*ast.ClassType, *ast.Method, error) {
	return compiler.FindStaticMethod(classType, methodName, convertClassTypes(parameters), allowedModifier)
}

func (r *TypeResolver) SearchMethod(receiverClass *ast.ClassType, methods []*ast.Method, parameters []*ast.Object) *ast.Method {
	return builtin.SearchMethod(receiverClass, methods, convertClassTypes(parameters))
}

func (r *TypeResolver) SearchConstructor(receiverClass *ast.ClassType, parameters []*ast.Object) (*ast.ClassType, *ast.Method, error) {
	return r.resolver.SearchConstructor(receiverClass, convertClassTypes(parameters))
}

func (r *TypeResolver) ConvertType(n *ast.TypeRef) (*ast.ClassType, error) {
	return r.resolver.ConvertType(n)
}

func convertClassTypes(parameters []*ast.Object) []*ast.ClassType {
	inputParameters := make([]*ast.ClassType, len(parameters))
	for i, parameter := range parameters {
		inputParameters[i] = parameter.ClassType
	}
	return inputParameters
}
