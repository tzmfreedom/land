package interpreter

import (
	"strings"

	"github.com/pkg/errors"
	"github.com/tzmfreedom/goland/ast"
)

type TypeResolver struct{}

func (r *TypeResolver) ResolveVariable(names []string, ctx *Context) (*Object, error) {
	if len(names) == 1 {
		if v, ok := ctx.Env.Get(names[0]); ok {
			return v, nil
		}
		return nil, errors.Errorf("%s is not found in this scope", names[0])
	} else {
		name := names[0]
		if val, ok := ctx.Env.Get(name); ok {
			for _, f := range names[1:] {
				val, ok = val.InstanceFields.Get(f)
				if !ok {
					return nil, errors.Errorf("%s is not found in this scope", f)
				}
			}
			return val, nil
		}
		if v, ok := ctx.StaticField.Get(name); ok {
			if val, ok := v.Get(names[1]); ok {
				for _, f := range names[2:] {
					val, ok = val.InstanceFields.Get(f)
					if !ok {
						return nil, errors.Errorf("%s is not found in this scope", f)
					}
				}
				return val, nil
			}
		}
		//if v, ok := ctx.NameSpaces.Get(name); ok {
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
	return nil, nil
}

func (r *TypeResolver) ResolveMethod(names []string, ctx *Context) (ast.Node, error) {
	if len(names) == 1 {
		methodName := names[0]
		if v, ok := ctx.Env.Get("this"); ok {
			if methods, ok := v.ClassType.InstanceMethods.Get(methodName); ok {
				return methods[0], nil
			}
			return nil, errors.Errorf("%s is not found in this scope", methodName)
		}
	} else {
		first := names[0]
		methodName := names[len(names)-1]
		fields := names[1 : len(names)-1]
		if val, ok := ctx.Env.Get(first); ok {
			for _, f := range fields {
				val, ok = val.InstanceFields.Get(f)
				if !ok {
					return nil, errors.Errorf("%s is not found in this scope", f)
				}
			}
			methods, ok := val.ClassType.InstanceMethods.Get(methodName)
			if ok {
				return methods[0], nil
			}
			return nil, errors.Errorf("%s is not found in this scope", methodName)
		}
		if len(names) == 2 {
			if v, ok := ctx.ClassTypes.Get(first); ok {
				if methods, ok := v.StaticMethods.Get(methodName); ok {
					return methods[0], nil
				}
				return nil, errors.Errorf("%s is not found in this scope", methodName)
			}
		}
		if v, ok := ctx.StaticField.Get(first); ok {
			if val, ok := v.Get(names[1]); ok {
				for _, f := range names[2 : len(names)-1] {
					val, ok = val.InstanceFields.Get(f)
					if !ok {
						return nil, errors.Errorf("%s is not found in this scope", f)
					}
				}
				methods, ok := val.ClassType.InstanceMethods.Get(methodName)
				if ok {
					return methods[0], nil
				}
				return nil, errors.Errorf("%s is not found in this scope", methodName)
			}
		}
		//if v, ok := ctx.NameSpaces.Get(first); ok {
		//	if classType, ok := v.Get(names[1]); ok {
		//		if len(names) > 3 {
		//			if field, ok := classType.StaticFields.Get(names[2]); ok {
		//				t := field.Type.(*ast.TypeRef)
		//				fieldType, _ := r.ResolveType(t.Name, ctx)
		//				for _, f := range names[3 : len(names)-1] {
		//					instanceField, ok := val.InstanceFields.Get(f)
		//					if !ok {
		//						return nil, errors.Errorf("%s is not found in this scope", f)
		//					}
		//					fieldType, _ = r.ResolveType(instanceField.Type.(*ast.TypeRef).Name, ctx)
		//				}
		//				methods, ok := fieldType.(*ClassType).InstanceMethods.Get(methodName)
		//				if ok {
		//					return methods[0], nil
		//				}
		//				return nil, errors.Errorf("%s is not found in this scope", methodName)
		//			}
		//		} else {
		//			methods, ok := classType.StaticMethods.Get(methodName)
		//			if ok {
		//				return methods[0], nil
		//			}
		//			return nil, errors.Errorf("%s is not found in this scope", methodName)
		//		}
		//	}
		//}
	}
	return nil, errors.Errorf("%s is not found in this scope", strings.Join(names, "."))
}

func (r *TypeResolver) ResolveType(names []string, ctx *Context) (interface{}, error) {
	if len(names) == 1 {
		className := names[0]
		if class, ok := ctx.ClassTypes.Get(className); ok {
			return class, nil
		}
		if classTypes, ok := ctx.NameSpaces.Get("System"); ok {
			if class, ok := classTypes.Get(className); ok {
				return class, nil
			}
		}
	} else if len(names) == 2 {
		// search for UserClass.InnerClass
		if class, ok := ctx.ClassTypes.Get(names[0]); ok {
			if inner, ok := class.InnerClasses.Get(names[1]); ok {
				return inner, nil
			}
		}
		// search for NameSpace.UserClass
		if classTypes, ok := ctx.NameSpaces.Get(names[0]); ok {
			if class, ok := classTypes.Get(names[1]); ok {
				return class, nil
			}
		}
	} else if len(names) == 3 {
		// search for NameSpace.UserClass.InnerClass
		if classTypes, ok := ctx.NameSpaces.Get(names[0]); ok {
			if class, ok := classTypes.Get(names[1]); ok {
				if inner, ok := class.InnerClasses.Get(names[2]); ok {
					return inner, nil
				}
			}
		}
	}
	return nil, nil
}
