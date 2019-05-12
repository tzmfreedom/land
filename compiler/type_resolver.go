package compiler

import (
	"fmt"
	"strings"

	"github.com/pkg/errors"
	"github.com/tzmfreedom/land/ast"
	"github.com/tzmfreedom/land/builtin"
)

type TypeResolver struct {
	Context  *Context
	resolver *builtin.TypeRefResolver
}

func NewTypeResolver(ctx *Context) *TypeResolver {
	return &TypeResolver{
		Context: ctx,
		resolver: &builtin.TypeRefResolver{
			ClassTypes:   ctx.ClassTypes,
			NameSpaces:   ctx.NameSpaces,
			CurrentClass: ctx.CurrentClass,
		},
	}
}

const (
	MODIFIER_PUBLIC_ONLY = iota
	MODIFIER_ALLOW_PROTECTED
	MODIFIER_ALL_OK
	MODIFIER_NO_CHECK
)

// TODO: resolve on static context

func (r *TypeResolver) ResolveVariable(names []string, checkSetter bool) (*ast.ClassType, error) {
	// TODO: return *ast.TypeRef smart solution
	if len(names) == 1 {
		if v, ok := r.Context.Env.Get(names[0]); ok {
			return v, nil
		}
		if v, ok := r.Context.Env.Get("this"); ok {
			instanceField, err := FindInstanceField(v, names[0], MODIFIER_ALL_OK, checkSetter)
			if err != nil {
				return nil, err
			}
			if instanceField != nil {
				return instanceField.Type, nil
			}
		}
		return nil, errors.Errorf("%s is not found in this scope", names[0])
	} else {
		name := names[0]
		if fieldType, ok := r.Context.Env.Get(name); ok {
			var instanceField *ast.Field
			var err error
			for i, f := range names[1:] {
				var allowedModifier int
				if i == 0 && name == "this" {
					allowedModifier = MODIFIER_ALL_OK
				} else {
					allowedModifier = MODIFIER_PUBLIC_ONLY
				}
				check := false
				if len(names)-2 == i {
					check = checkSetter
				}
				instanceField, err = FindInstanceField(fieldType, f, allowedModifier, check)
				if err != nil {
					return nil, err
				}
				if instanceField == nil {
					return nil, fmt.Errorf("Field %s is not found", f)
				}
				fieldType = instanceField.Type
			}
			if checkSetter && instanceField.IsFinal() {
				if r.Context.CurrentMethod.IsConstructor || r.Context.CurrentMethod == nil {
					if len(names) == 2 && names[0] == "this" {
						// OK
					} else {
						// NG
					}
				}
			}
			return fieldType, nil
		}
		if v, ok := r.Context.ClassTypes.Get(name); ok {
			check := false
			if len(names) == 2 {
				check = checkSetter
			}
			n, err := FindStaticField(v, names[1], MODIFIER_PUBLIC_ONLY, check)
			if err != nil {
				return nil, err
			}
			if n != nil {
				var instanceField *ast.Field
				var err error
				fieldType := n.Type
				for i, f := range names[2:] {
					check := false
					if len(names)-3 == i {
						check = checkSetter
					}
					instanceField, err = FindInstanceField(fieldType, f, MODIFIER_PUBLIC_ONLY, check)
					if err != nil {
						return nil, err
					}
					if instanceField == nil {
						return nil, fmt.Errorf("Field %s is not found", f)
					}
					fieldType = instanceField.Type
				}
				return fieldType, nil
			}
		}
		if v, ok := r.Context.NameSpaces.Get(name); ok {
			if classType, ok := v.Get(names[1]); ok {
				check := false
				if len(names) == 3 {
					check = checkSetter
				}
				field, err := FindStaticField(classType, names[2], MODIFIER_PUBLIC_ONLY, check)
				if err != nil {
					return nil, err
				}
				if field != nil {
					var instanceField *ast.Field
					var err error

					fieldType := field.Type
					for i, f := range names[3:] {
						check := false
						if len(names)-4 == i {
							check = checkSetter
						}
						instanceField, err = FindInstanceField(fieldType, f, MODIFIER_PUBLIC_ONLY, check)
						if err != nil {
							return nil, err
						}
						if instanceField == nil {
							return nil, fmt.Errorf("Field %s is not found", f)
						}
						fieldType = instanceField.Type
					}
					return fieldType, nil
				}
			}
		}
	}
	return nil, fmt.Errorf("local variable %s is not found", names[0])
}

func (r *TypeResolver) ResolveMethod(names []string, parameters []*ast.ClassType) (*ast.ClassType, *ast.Method, error) {
	if len(names) == 1 {
		methodName := names[0]
		if v, ok := r.Context.Env.Get("this"); ok {
			classType, method, err := FindInstanceMethod(v, methodName, parameters, MODIFIER_ALL_OK)
			if err != nil {
				return nil, nil, err
			}
			if method == nil {
				return nil, nil, errors.Errorf("%s is not found in this scope", methodName)
			}
			return classType, method, nil
		}
	} else {
		first := names[0]
		methodName := names[len(names)-1]
		fields := names[1 : len(names)-1]
		if fieldType, ok := r.Context.Env.Get(first); ok {
			for i, f := range fields {
				var allowedModifier int
				if first == "this" && i == 0 {
					allowedModifier = MODIFIER_ALL_OK
				} else {
					allowedModifier = MODIFIER_PUBLIC_ONLY
				}
				instanceField, err := FindInstanceField(fieldType, f, allowedModifier, false)
				if err != nil {
					return nil, nil, err
				}
				fieldType = instanceField.Type
			}
			var allowedModifier int
			if first == "this" && len(fields) == 0 {
				allowedModifier = MODIFIER_ALL_OK
			} else {
				allowedModifier = MODIFIER_PUBLIC_ONLY
			}
			return FindInstanceMethod(fieldType, methodName, parameters, allowedModifier)
		}
		if len(names) == 2 {
			if v, ok := r.Context.ClassTypes.Get(first); ok {
				return FindStaticMethod(v, methodName, parameters, MODIFIER_PUBLIC_ONLY)
			}
		}
		if v, ok := r.Context.ClassTypes.Get(first); ok {
			n, err := FindStaticField(v, names[1], MODIFIER_PUBLIC_ONLY, false)
			if err == nil {
				fieldType := n.Type
				for _, f := range names[2 : len(names)-1] {
					instanceField, err := FindInstanceField(fieldType, f, MODIFIER_PUBLIC_ONLY, false)
					if err != nil {
						return nil, nil, err
					}
					fieldType = instanceField.Type
				}
				return FindInstanceMethod(fieldType, methodName, parameters, MODIFIER_PUBLIC_ONLY)
			}
		}

		if v, ok := r.Context.NameSpaces.Get(first); ok {
			if classType, ok := v.Get(names[1]); ok {
				// namespace.class.static_field.instance_field...instance_method()
				if len(names) > 3 {
					if field, err := FindStaticField(classType, names[2], MODIFIER_PUBLIC_ONLY, false); err == nil {
						fieldType := field.Type
						for _, f := range names[3 : len(names)-1] {
							instanceField, err := FindInstanceField(fieldType, f, MODIFIER_PUBLIC_ONLY, false)
							if err != nil {
								return nil, nil, err
							}
							fieldType = instanceField.Type
						}
						return FindInstanceMethod(fieldType, methodName, parameters, MODIFIER_PUBLIC_ONLY)
					}
					// namespace.class.static_method()
				} else {
					return FindStaticMethod(classType, methodName, parameters, MODIFIER_PUBLIC_ONLY)
				}
			}
		}
	}
	return nil, nil, errors.Errorf("%s is not found in this scope", strings.Join(names, "."))
}

func (r *TypeResolver) ResolveType(names []string) (*ast.ClassType, error) {
	return r.resolver.ResolveType(names)
}

func FindInstanceMethod(classType *ast.ClassType, methodName string, parameters []*ast.ClassType, allowedModifier int) (*ast.ClassType, *ast.Method, error) {
	if classType == builtin.NullType {
		return nil, nil, builtin.NewNullPointerException(methodName)
	}
	methods, ok := classType.InstanceMethods.Get(methodName)
	if ok {
		method := builtin.SearchMethod(classType, methods, parameters)
		if method != nil {
			if allowedModifier == MODIFIER_PUBLIC_ONLY && !method.IsPublic() {
				return nil, nil, fmt.Errorf("Method access modifier must be public but %s", method.AccessModifier())
			}
			if allowedModifier == MODIFIER_ALLOW_PROTECTED && method.IsPrivate() {
				return nil, nil, fmt.Errorf("Method access modifier must be public/protected but private")
			}
			return classType, method, nil
		}
	}
	if classType.SuperClass != nil {
		if allowedModifier == MODIFIER_ALL_OK {
			allowedModifier = MODIFIER_ALLOW_PROTECTED
		}
		return FindInstanceMethod(classType.SuperClass, methodName, parameters, allowedModifier)
	}
	return nil, nil, methodNotFoundError(classType, methodName, parameters)
}

func FindStaticMethod(classType *ast.ClassType, methodName string, parameters []*ast.ClassType, allowedModifier int) (*ast.ClassType, *ast.Method, error) {
	methods, ok := classType.StaticMethods.Get(methodName)
	if ok {
		method := builtin.SearchMethod(classType, methods, parameters)
		if method != nil {
			if allowedModifier == MODIFIER_PUBLIC_ONLY && !method.IsPublic() {
				return nil, nil, fmt.Errorf("Method access modifier must be public but %s", method.AccessModifier())
			}
			if allowedModifier == MODIFIER_ALLOW_PROTECTED && method.IsPrivate() {
				return nil, nil, fmt.Errorf("Method access modifier must be public/protected but private")
			}
			return classType, method, nil
		}
	}
	if classType.SuperClass != nil {
		if allowedModifier == MODIFIER_ALL_OK {
			allowedModifier = MODIFIER_ALLOW_PROTECTED
		}
		return FindStaticMethod(classType.SuperClass, methodName, parameters, allowedModifier)
	}
	return nil, nil, methodNotFoundError(classType, methodName, parameters)
}

func FindInstanceField(classType *ast.ClassType, fieldName string, allowedModifier int, checkSetter bool) (*ast.Field, error) {
	fieldType, ok := classType.InstanceFields.Get(fieldName)
	if ok {
		if allowedModifier == MODIFIER_PUBLIC_ONLY && !fieldType.IsPublic(checkSetter) {
			return nil, fmt.Errorf("Field access modifier must be public but %s", fieldType.AccessModifier(checkSetter))
		}
		if allowedModifier == MODIFIER_ALLOW_PROTECTED && fieldType.IsPrivate(checkSetter) {
			return nil, fmt.Errorf("Field access modifier must be public/protected but private")
		}
		return fieldType, nil
	}
	if classType.SuperClass != nil {
		if allowedModifier == MODIFIER_ALL_OK {
			allowedModifier = MODIFIER_ALLOW_PROTECTED
		}
		f, err := FindInstanceField(classType.SuperClass, fieldName, allowedModifier, checkSetter)
		if err == nil {
			return f, nil
		}
	}
	return nil, fieldNotFoundError(classType, fieldName)
}

func FindStaticField(classType *ast.ClassType, fieldName string, allowedModifier int, checkSetter bool) (*ast.Field, error) {
	fieldType, ok := classType.StaticFields.Get(fieldName)
	if ok {
		if allowedModifier == MODIFIER_PUBLIC_ONLY && !fieldType.IsPublic(checkSetter) {
			return nil, fmt.Errorf("Field access modifier must be public but %s", fieldType.AccessModifier(checkSetter))
		}
		if allowedModifier == MODIFIER_ALLOW_PROTECTED && fieldType.IsPrivate(checkSetter) {
			return nil, fmt.Errorf("Field access modifier must be public/protected but private")
		}
		return fieldType, nil
	}
	if classType.SuperClass != nil {
		if allowedModifier == MODIFIER_ALL_OK {
			allowedModifier = MODIFIER_ALLOW_PROTECTED
		}
		f, err := FindStaticField(classType.SuperClass, fieldName, allowedModifier, checkSetter)
		if err == nil {
			return f, nil
		}
	}
	return nil, fieldNotFoundError(classType, fieldName)
}

func (r *TypeResolver) SearchConstructor(classType *ast.ClassType, parameters []*ast.ClassType) (*ast.ClassType, *ast.Method, error) {
	method := builtin.SearchMethod(classType, classType.Constructors, parameters)
	if method != nil {
		return classType, method, nil
	}
	if classType.SuperClass != nil {
		return r.SearchConstructor(classType.SuperClass, parameters)
	}
	return nil, nil, nil
}

func (r *TypeResolver) ConvertType(n *ast.TypeRef) (*ast.ClassType, error) {
	return r.resolver.ConvertType(n)
}

func methodNotFoundError(classType *ast.ClassType, methodName string, parameters []*ast.ClassType) error {
	parameterStrings := make([]string, len(parameters))
	for i, parameter := range parameters {
		parameterStrings[i] = parameter.String()
	}
	return fmt.Errorf("Method not found: %s.%s(%s)", classType.Name, methodName, strings.Join(parameterStrings, ", "))
}

func fieldNotFoundError(classType *ast.ClassType, fieldName string) error {
	return fmt.Errorf("Field not found: %s.%s", classType.Name, fieldName)
}
