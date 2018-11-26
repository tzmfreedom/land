package compiler

import (
	"fmt"
	"strings"

	"github.com/pkg/errors"
	"github.com/tzmfreedom/goland/ast"
	"github.com/tzmfreedom/goland/builtin"
)

type TypeResolver struct {
	Context *Context
}

const (
	MODIFIER_PUBLIC_ONLY = iota
	MODIFIER_ALLOW_PROTECTED
	MODIFIER_ALL_OK
	MODIFIER_NO_CHECK
)

func (r *TypeResolver) ResolveVariable(names []string) (*builtin.ClassType, error) {
	if len(names) == 1 {
		if v, ok := r.Context.Env.Get(names[0]); ok {
			return v, nil
		}
		return nil, errors.Errorf("%s is not found in this scope", names[0])
	} else {
		name := names[0]
		if fieldType, ok := r.Context.Env.Get(name); ok {
			for i, f := range names[1:] {
				var allowedModifier int
				if i == 0 && f == "this" {
					allowedModifier = MODIFIER_ALL_OK
				} else {
					allowedModifier = MODIFIER_PUBLIC_ONLY
				}
				instanceField, err := r.findInstanceField(fieldType, f, allowedModifier)
				if err != nil {
					return nil, err
				}
				if instanceField == nil {
					return nil, fmt.Errorf("Field %s is not found", f)
				}
				fieldType, _ = r.ResolveType(instanceField.Type.(*ast.TypeRef).Name)
			}
			return fieldType, nil
		}
		if v, ok := r.Context.ClassTypes.Get(name); ok {
			n, err := r.findStaticField(v, names[1], MODIFIER_PUBLIC_ONLY)
			if err != nil {
				return nil, err
			}
			if n != nil {
				t := n.Type.(*ast.TypeRef)
				fieldType, _ := r.ResolveType(t.Name)
				for _, f := range names[2:] {
					instanceField, err := r.findInstanceField(fieldType, f, MODIFIER_PUBLIC_ONLY)
					if err != nil {
						return nil, err
					}
					if instanceField == nil {
						return nil, fmt.Errorf("Field %s is not found", f)
					}
					fieldType, _ = r.ResolveType(instanceField.Type.(*ast.TypeRef).Name)
				}
				return fieldType, nil
			}
		}
		if v, ok := r.Context.NameSpaces.Get(name); ok {
			if classType, ok := v.Get(names[1]); ok {
				field, err := r.findStaticField(classType, names[2], MODIFIER_PUBLIC_ONLY)
				if err != nil {
					return nil, err
				}
				if field != nil {
					t := field.Type.(*ast.TypeRef)
					fieldType, _ := r.ResolveType(t.Name)
					for _, f := range names[3:] {
						instanceField, err := r.findInstanceField(fieldType, f, MODIFIER_PUBLIC_ONLY)
						if err != nil {
							return nil, err
						}
						if instanceField == nil {
							return nil, fmt.Errorf("Field %s is not found", f)
						}
						fieldType, _ = r.ResolveType(instanceField.Type.(*ast.TypeRef).Name)
					}
					return fieldType, nil
				}
			}
		}
	}
	return nil, fmt.Errorf("local variable %s is not found", names[0])
}

func (r *TypeResolver) ResolveMethod(names []string, parameters []*builtin.ClassType) (*ast.MethodDeclaration, error) {
	if len(names) == 1 {
		methodName := names[0]
		if v, ok := r.Context.Env.Get("this"); ok {
			method, err := r.FindInstanceMethod(v, methodName, parameters, MODIFIER_ALL_OK)
			if err != nil {
				return nil, err
			}
			if method == nil {
				return nil, errors.Errorf("%s is not found in this scope", methodName)
			}
			return method, nil
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
				instanceField, err := r.findInstanceField(fieldType, f, allowedModifier)
				if err != nil {
					return nil, err
				}
				fieldType, _ = r.ResolveType(instanceField.Type.(*ast.TypeRef).Name)
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
			n, err := r.findStaticField(v, names[1], MODIFIER_PUBLIC_ONLY)
			if err == nil {
				t := n.Type.(*ast.TypeRef)
				fieldType, _ := r.ResolveType(t.Name)
				for _, f := range names[2 : len(names)-1] {
					instanceField, err := r.findInstanceField(fieldType, f, MODIFIER_PUBLIC_ONLY)
					if err != nil {
						return nil, err
					}
					fieldType, _ = r.ResolveType(instanceField.Type.(*ast.TypeRef).Name)
				}
				return r.FindInstanceMethod(fieldType, methodName, parameters, MODIFIER_PUBLIC_ONLY)
			}
		}
		if v, ok := r.Context.NameSpaces.Get(first); ok {
			if classType, ok := v.Get(names[1]); ok {
				if len(names) > 3 {
					if field, err := r.findStaticField(classType, names[2], MODIFIER_PUBLIC_ONLY); err == nil {
						t := field.Type.(*ast.TypeRef)
						fieldType, _ := r.ResolveType(t.Name)
						for _, f := range names[3 : len(names)-1] {
							instanceField, err := r.findInstanceField(fieldType, f, MODIFIER_PUBLIC_ONLY)
							if err != nil {
								return nil, err
							}
							fieldType, _ = r.ResolveType(instanceField.Type.(*ast.TypeRef).Name)
						}
						return r.FindInstanceMethod(fieldType, methodName, parameters, MODIFIER_PUBLIC_ONLY)
					}
				} else {
					return r.FindStaticMethod(classType, methodName, parameters, MODIFIER_PUBLIC_ONLY)
				}
			}
		}
	}
	return nil, errors.Errorf("%s is not found in this scope", strings.Join(names, "."))
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

func (r *TypeResolver) FindInstanceMethod(classType *builtin.ClassType, methodName string, parameters []*builtin.ClassType, allowedModifier int) (*ast.MethodDeclaration, error) {
	methods, ok := classType.InstanceMethods.Get(methodName)
	if ok {
		method := r.searchMethod(methods, parameters)
		if method != nil {
			if allowedModifier == MODIFIER_PUBLIC_ONLY && !method.IsPublic() {
				return nil, fmt.Errorf("Method access modifier must be public but %s", method.AccessModifier())
			}
			if allowedModifier == MODIFIER_ALLOW_PROTECTED && method.IsPrivate() {
				return nil, fmt.Errorf("Method access modifier must be public/protected but private")
			}
			return method, nil
		}
	}
	if classType.SuperClass != nil {
		super, err := r.ResolveType(classType.SuperClass.(*ast.TypeRef).Name)
		if err != nil {
			return nil, errors.New("Method not found")
		}
		if allowedModifier == MODIFIER_ALL_OK {
			allowedModifier = MODIFIER_ALLOW_PROTECTED
		}
		return r.FindInstanceMethod(super, methodName, parameters, allowedModifier)
	}
	return nil, errors.New("Method not found")
}

func (r *TypeResolver) FindStaticMethod(classType *builtin.ClassType, methodName string, parameters []*builtin.ClassType, allowedModifier int) (*ast.MethodDeclaration, error) {
	methods, ok := classType.StaticMethods.Get(methodName)
	if ok {
		method := r.searchMethod(methods, parameters)
		if method != nil {
			if allowedModifier == MODIFIER_PUBLIC_ONLY && !method.IsPublic() {
				return nil, fmt.Errorf("Method access modifier must be public but %s", method.AccessModifier())
			}
			if allowedModifier == MODIFIER_ALLOW_PROTECTED && method.IsPrivate() {
				return nil, fmt.Errorf("Method access modifier must be public/protected but private")
			}
			return method, nil
		}
	}
	if classType.SuperClass != nil {
		super, err := r.ResolveType(classType.SuperClass.(*ast.TypeRef).Name)
		if err != nil {
			return nil, errors.New("Method not found")
		}
		if allowedModifier == MODIFIER_ALL_OK {
			allowedModifier = MODIFIER_ALLOW_PROTECTED
		}
		return r.FindStaticMethod(super, methodName, parameters, allowedModifier)
	}
	return nil, errors.New("Method not found")
}

func (r *TypeResolver) findInstanceField(classType *builtin.ClassType, fieldName string, allowedModifier int) (*builtin.Field, error) {
	fieldType, ok := classType.InstanceFields.Get(fieldName)
	if ok {
		if allowedModifier == MODIFIER_PUBLIC_ONLY && !fieldType.IsPublic() {
			return nil, fmt.Errorf("Field access modifier must be public but %s", fieldType.AccessModifier())
		}
		if allowedModifier == MODIFIER_ALLOW_PROTECTED && fieldType.IsPrivate() {
			return nil, fmt.Errorf("Field access modifier must be public/protected but private")
		}
		return fieldType, nil
	}
	if classType.SuperClass != nil {
		super, err := r.ResolveType(classType.SuperClass.(*ast.TypeRef).Name)
		if err != nil {
			return nil, err
		}
		if allowedModifier == MODIFIER_ALL_OK {
			allowedModifier = MODIFIER_ALLOW_PROTECTED
		}
		return r.findInstanceField(super, fieldName, allowedModifier)
	}
	return nil, nil
}

func (r *TypeResolver) findStaticField(classType *builtin.ClassType, fieldName string, allowedModifier int) (*builtin.Field, error) {
	fieldType, ok := classType.StaticFields.Get(fieldName)
	if ok {
		if allowedModifier == MODIFIER_PUBLIC_ONLY && !fieldType.IsPublic() {
			return nil, fmt.Errorf("Field access modifier must be public but %s", fieldType.AccessModifier())
		}
		if allowedModifier == MODIFIER_ALLOW_PROTECTED && fieldType.IsPrivate() {
			return nil, fmt.Errorf("Field access modifier must be public/protected but private")
		}
		return fieldType, nil
	}
	if classType.SuperClass != nil {
		super, err := r.ResolveType(classType.SuperClass.(*ast.TypeRef).Name)
		if err != nil {
			return nil, err
		}
		if allowedModifier == MODIFIER_ALL_OK {
			allowedModifier = MODIFIER_ALLOW_PROTECTED
		}
		return r.findStaticField(super, fieldName, allowedModifier)
	}
	return nil, nil
}

func (r *TypeResolver) searchMethod(methods []ast.Node, parameters []*builtin.ClassType) *ast.MethodDeclaration {
	l := len(parameters)
	for _, method := range methods {
		m := method.(*ast.MethodDeclaration)
		if len(m.Parameters) != l {
			continue
		}
		match := true
		for i, p := range m.Parameters {
			inputParam := parameters[i]
			methodParam, _ := r.ResolveType(p.(*ast.TypeRef).Name)
			if inputParam != methodParam {
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
