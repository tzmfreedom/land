package compiler

import (
	"fmt"
	"strings"

	"github.com/pkg/errors"
	"github.com/tzmfreedom/goland/ast"
	"github.com/tzmfreedom/goland/builtin"
)

type TypeResolver struct {
	Context        *Context
	resolver       *builtin.TypeRefResolver
	IgnoreGenerics bool
}

func NewTypeResolver(ctx *Context, ignoreGenerics bool) *TypeResolver {
	return &TypeResolver{
		Context: ctx,
		resolver: &builtin.TypeRefResolver{
			ClassTypes:   ctx.ClassTypes,
			NameSpaces:   ctx.NameSpaces,
			CurrentClass: ctx.CurrentClass,
		},
		IgnoreGenerics: ignoreGenerics,
	}
}

const (
	MODIFIER_PUBLIC_ONLY = iota
	MODIFIER_ALLOW_PROTECTED
	MODIFIER_ALL_OK
	MODIFIER_NO_CHECK
)

// TODO: resolve on static context

func (r *TypeResolver) ResolveVariable(names []string, checkSetter bool) (*builtin.ClassType, error) {
	// TODO: return *ast.TypeRef smart solution
	if len(names) == 1 {
		if v, ok := r.Context.Env.Get(names[0]); ok {
			return v, nil
		}
		return nil, errors.Errorf("%s is not found in this scope", names[0])
	} else {
		name := names[0]
		if fieldType, ok := r.Context.Env.Get(name); ok {
			var instanceField *builtin.Field
			var err error
			for i, f := range names[1:] {
				var allowedModifier int
				if i == 0 && f == "this" {
					allowedModifier = MODIFIER_ALL_OK
				} else {
					allowedModifier = MODIFIER_PUBLIC_ONLY
				}
				check := false
				if len(names)-2 == i {
					check = checkSetter
				}
				instanceField, err = r.findInstanceField(fieldType, f, allowedModifier, check)
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
		if v, ok := r.Context.ClassTypes.Get(name); ok {
			check := false
			if len(names) == 2 {
				check = checkSetter
			}
			n, err := r.findStaticField(v, names[1], MODIFIER_PUBLIC_ONLY, check)
			if err != nil {
				return nil, err
			}
			if n != nil {
				var instanceField *builtin.Field
				var err error
				fieldType := n.Type
				for i, f := range names[2:] {
					check := false
					if len(names)-3 == i {
						check = checkSetter
					}
					instanceField, err = r.findInstanceField(fieldType, f, MODIFIER_PUBLIC_ONLY, check)
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
				field, err := r.findStaticField(classType, names[2], MODIFIER_PUBLIC_ONLY, check)
				if err != nil {
					return nil, err
				}
				if field != nil {
					var instanceField *builtin.Field
					var err error

					fieldType := field.Type
					for i, f := range names[3:] {
						check := false
						if len(names)-4 == i {
							check = checkSetter
						}
						instanceField, err = r.findInstanceField(fieldType, f, MODIFIER_PUBLIC_ONLY, check)
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

func (r *TypeResolver) ResolveMethod(names []string, parameters []*builtin.ClassType) (*builtin.ClassType, *builtin.Method, error) {
	if len(names) == 1 {
		methodName := names[0]
		if v, ok := r.Context.Env.Get("this"); ok {
			classType, method, err := r.FindInstanceMethod(v, methodName, parameters, MODIFIER_ALL_OK)
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
				instanceField, err := r.findInstanceField(fieldType, f, allowedModifier, false)
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
			return r.FindInstanceMethod(fieldType, methodName, parameters, allowedModifier)
		}
		if len(names) == 2 {
			if v, ok := r.Context.ClassTypes.Get(first); ok {
				return r.FindStaticMethod(v, methodName, parameters, MODIFIER_PUBLIC_ONLY)
			}
		}
		if v, ok := r.Context.ClassTypes.Get(first); ok {
			n, err := r.findStaticField(v, names[1], MODIFIER_PUBLIC_ONLY, false)
			if err == nil {
				fieldType := n.Type
				for _, f := range names[2 : len(names)-1] {
					instanceField, err := r.findInstanceField(fieldType, f, MODIFIER_PUBLIC_ONLY, false)
					if err != nil {
						return nil, nil, err
					}
					fieldType = instanceField.Type
				}
				return r.FindInstanceMethod(fieldType, methodName, parameters, MODIFIER_PUBLIC_ONLY)
			}
		}

		if v, ok := r.Context.NameSpaces.Get(first); ok {
			if classType, ok := v.Get(names[1]); ok {
				// namespace.class.static_field.instance_field...instance_method()
				if len(names) > 3 {
					if field, err := r.findStaticField(classType, names[2], MODIFIER_PUBLIC_ONLY, false); err == nil {
						fieldType := field.Type
						for _, f := range names[3 : len(names)-1] {
							instanceField, err := r.findInstanceField(fieldType, f, MODIFIER_PUBLIC_ONLY, false)
							if err != nil {
								return nil, nil, err
							}
							fieldType = instanceField.Type
						}
						return r.FindInstanceMethod(fieldType, methodName, parameters, MODIFIER_PUBLIC_ONLY)
					}
					// namespace.class.static_method()
				} else {
					return r.FindStaticMethod(classType, methodName, parameters, MODIFIER_PUBLIC_ONLY)
				}
			}
		}
	}
	return nil, nil, errors.Errorf("%s is not found in this scope", strings.Join(names, "."))
}

func (r *TypeResolver) ResolveType(names []string) (*builtin.ClassType, error) {
	return r.resolver.ResolveType(names)
}

func (r *TypeResolver) FindInstanceMethod(classType *builtin.ClassType, methodName string, parameters []*builtin.ClassType, allowedModifier int) (*builtin.ClassType, *builtin.Method, error) {
	methods, ok := classType.InstanceMethods.Get(methodName)
	if ok {
		method := r.SearchMethod(classType, methods, parameters)
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
		return r.FindInstanceMethod(classType.SuperClass, methodName, parameters, allowedModifier)
	}
	return nil, nil, methodNotFoundError(classType, methodName, parameters)
}

func (r *TypeResolver) FindStaticMethod(classType *builtin.ClassType, methodName string, parameters []*builtin.ClassType, allowedModifier int) (*builtin.ClassType, *builtin.Method, error) {
	methods, ok := classType.StaticMethods.Get(methodName)
	if ok {
		method := r.SearchMethod(classType, methods, parameters)
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
		return r.FindStaticMethod(classType.SuperClass, methodName, parameters, allowedModifier)
	}
	return nil, nil, methodNotFoundError(classType, methodName, parameters)
}

func (r *TypeResolver) findInstanceField(classType *builtin.ClassType, fieldName string, allowedModifier int, checkSetter bool) (*builtin.Field, error) {
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
		return r.findInstanceField(classType.SuperClass, fieldName, allowedModifier, checkSetter)
	}
	return nil, nil
}

func (r *TypeResolver) findStaticField(classType *builtin.ClassType, fieldName string, allowedModifier int, checkSetter bool) (*builtin.Field, error) {
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
		return r.findStaticField(classType.SuperClass, fieldName, allowedModifier, checkSetter)
	}
	return nil, nil
}

func (r *TypeResolver) SearchMethod(receiverClass *builtin.ClassType, methods []*builtin.Method, parameters []*builtin.ClassType) *builtin.Method {
	l := len(parameters)
	for _, m := range methods {
		if len(m.Parameters) != l {
			continue
		}
		match := true

		for i, p := range m.Parameters {
			inputParam := parameters[i]
			typeRef := p.(*ast.Parameter).Type.(*ast.TypeRef)

			var methodParam *builtin.ClassType
			if typeRef.IsGenerics() {
				// TODO: implement better solution
				if r.IgnoreGenerics {
					continue
				}
				generics := receiverClass.Extra["generics"].([]*builtin.ClassType)
				if typeRef.IsGenericsNumber(1) {
					methodParam = generics[0]
				} else {
					methodParam = generics[1]
				}
			} else {
				methodParam, _ = r.ConvertType(typeRef)
			}
			// TODO: implement
			// extend, implements, Object
			if methodParam == builtin.ObjectType {
				continue
			}
			if !inputParam.Equals(methodParam) {
				match = false
				break
			}
		}
		if match {
			return m
		}
	}
	return nil
}

func (r *TypeResolver) SearchConstructor(classType *builtin.ClassType, parameters []*builtin.ClassType) (*builtin.ClassType, *builtin.Method, error) {
	method := r.SearchMethod(classType, classType.Constructors, parameters)
	if method != nil {
		return classType, method, nil
	}
	if classType.SuperClass != nil {
		return r.SearchConstructor(classType.SuperClass, parameters)
	}
	return nil, nil, nil
}

func (r *TypeResolver) ConvertType(n *ast.TypeRef) (*builtin.ClassType, error) {
	return r.resolver.ConvertType(n)
}

func methodNotFoundError(classType *builtin.ClassType, methodName string, parameters []*builtin.ClassType) error {
	parameterStrings := make([]string, len(parameters))
	for i, parameter := range parameters {
		parameterStrings[i] = parameter.String()
	}
	return fmt.Errorf("Method not found: %s.%s(%s)", classType.Name, methodName, strings.Join(parameterStrings, ", "))
}
