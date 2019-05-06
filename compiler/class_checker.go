package compiler

import (
	"fmt"
	"strings"

	"errors"

	"github.com/tzmfreedom/land/ast"
	"github.com/tzmfreedom/land/builtin"
)

func CheckClass(t *ast.ClassType) error {
	if err := checkTopLevelType(t); err != nil {
		return err
	}
	if err := checkConstructorName(t); err != nil {
		return err
	}
	if err := checkImplements(t); err != nil {
		return err
	}
	if err := checkExtends(t); err != nil {
		return err
	}
	if err := checkAbstractMethods(t.InstanceMethods); err != nil {
		return err
	}
	if err := checkAbstractMethods(t.StaticMethods); err != nil {
		return err
	}
	if t.SuperClass != nil {
		if err := checkNotImplementedForAbstractMethod(t, t.InstanceMethods, t.SuperClass.InstanceMethods); err != nil {
			return err
		}
		if err := checkNotImplementedForAbstractMethod(t, t.StaticMethods, t.SuperClass.StaticMethods); err != nil {
			return err
		}
	}
	if err := checkSameParameterName(t.InstanceMethods); err != nil {
		return err
	}
	if err := checkSameParameterName(t.StaticMethods); err != nil {
		return err
	}
	if err := checkMethodSignatureDuplicated(t.InstanceMethods); err != nil {
		return err
	}
	if err := checkMethodSignatureDuplicated(t.StaticMethods); err != nil {
		return err
	}
	if err := checkOverrideMethod(t, t.InstanceMethods, FindInstanceMethod); err != nil {
		return err
	}
	if err := checkOverrideMethod(t, t.StaticMethods, FindStaticMethod); err != nil {
		return err
	}
	if err := checkOverrideField(t); err != nil {
		return err
	}
	return nil
}

func checkTopLevelType(t *ast.ClassType) error {
	if t.Is("public") || t.Is("global") {
		return nil
	}
	return fmt.Errorf("Top-level type must have public or global visibility: %s", t.Name)
}

func checkSameParameterName(m *ast.MethodMap) error {
	for _, methods := range m.Data {
		for _, m := range methods {
			parameterNames := map[string]struct{}{}
			for _, p := range m.Parameters {
				name := p.Name
				if _, ok := parameterNames[name]; ok {
					return fmt.Errorf("parameter name is duplicated: %s", name)
				}
				parameterNames[name] = struct{}{}
			}
		}
	}
	return nil
}

func checkMethodSignatureDuplicated(m *ast.MethodMap) error {
	for _, methods := range m.Data {
		if len(methods) == 1 {
			continue
		}
		for i, m := range methods {
			l := len(m.Parameters)
			for _, other := range methods[i+1:] {
				if len(other.Parameters) != l {
					continue
				}
				match := true
				for i, p := range m.Parameters {
					otherParam := other.Parameters[i]
					if p.Type != otherParam.Type {
						match = false
						break
					}
				}
				if match {
					return fmt.Errorf("method %s is duplicated", m.Name)
				}
			}
		}
	}
	return nil
}

func checkExtends(t *ast.ClassType) error {
	if t.SuperClass == nil {
		return nil
	}
	super := t.SuperClass
	if !super.IsAbstract() && !super.IsVirtual() {
		return fmt.Errorf("Non-virtual and non-abstract type cannot be extended: %s", super.Name)
	}
	return nil
}

func checkImplements(t *ast.ClassType) error {
	if len(t.ImplementClasses) == 0 {
		return nil
	}
	for _, impl := range t.ImplementClasses {
		for _, methods := range impl.InstanceMethods.All() {
			var matchedMethod *ast.Method
			for _, method := range methods {
				instanceMethods, ok := t.InstanceMethods.Get(method.Name)
				if !ok {
					return fmt.Errorf("Class %s must implement the method: %s", t.Name, MethodSignature(impl, method))
				}
				matchedMethod = builtin.SearchMethod(impl, instanceMethods, ParameterClassTypes(method.Parameters))
				if matchedMethod == nil {
					return fmt.Errorf("Class %s must implement the method: %s", t.Name, MethodSignature(impl, method))
				}
			}
		}
	}
	return nil
}

func checkAbstractMethods(m *ast.MethodMap) error {
	for _, methods := range m.All() {
		for _, method := range methods {
			if !method.IsAbstract() {
				continue
			}
			if method.Statements != nil {
				return errors.New("Abstract methods cannot have a body")
			}
		}
	}
	return nil
}

// if override modifier is specified, check modifier virtual/abstract on SuperClass
// if SuperClass is abstract/virtual, check override method implementation
func checkOverrideMethod(
	t *ast.ClassType,
	methodMap *ast.MethodMap,
	findMethod func(*ast.ClassType, string, []*ast.ClassType, int) (*ast.ClassType, *ast.Method, error),
) error {
	for _, methods := range methodMap.All() {
		for _, m := range methods {
			if m.IsOverride() {
				if t.SuperClass == nil {
					return fmt.Errorf("@Override specified for non-overriding method: %s", m.Name)
				}

				_, method, err := findMethod(t.SuperClass, m.Name, ParameterClassTypes(m.Parameters), MODIFIER_NO_CHECK)
				if err != nil {
					return fmt.Errorf("@Override specified for non-overriding method: %s", MethodSignature(t, m))
				}
				if !method.IsAbstract() && !method.IsVirtual() {
					return fmt.Errorf("Non-virtual, non-abstract methods cannot be overridden: %s", MethodSignature(t, m))
				}
			} else {
				if t.SuperClass == nil {
					continue
				}
				types := make([]*ast.ClassType, len(m.Parameters))
				for i, p := range m.Parameters {
					types[i] = p.Type
				}
				_, method, _ := findMethod(t.SuperClass, m.Name, types, MODIFIER_NO_CHECK)
				if method != nil {
					if method.IsAbstract() || method.IsVirtual() {
						return fmt.Errorf("Method must use the override keyword: %s", MethodSignature(t, method))
					}
					return fmt.Errorf("Non-virtual, non-abstract methods cannot be overridden: %s", MethodSignature(t, method))
				}
			}
		}
	}

	return nil
}

func checkNotImplementedForAbstractMethod(t *ast.ClassType, implemented *ast.MethodMap, extended *ast.MethodMap) error {
	if t.SuperClass == nil {
		return nil
	}
	for _, methods := range extended.All() {
		for _, method := range methods {
			if !method.IsAbstract() {
				continue
			}
			implementedMethods, ok := implemented.Get(method.Name)
			if !ok {
				return fmt.Errorf("Class %s must implement the abstract method: %s", t.Name, MethodSignature(t.SuperClass, method))
			}
			matchedMethod := builtin.SearchMethod(nil, implementedMethods, ParameterClassTypes(method.Parameters))
			if matchedMethod == nil {
				return fmt.Errorf("Class %s must implement the abstract method: %s", t.Name, MethodSignature(t.SuperClass, method))
			}
		}
	}
	return nil
}

func checkOverrideField(t *ast.ClassType) error {
	for _, field := range t.InstanceFields.Data {
		if field.IsOverride() {
			if t.SuperClass == nil {
				return fmt.Errorf("override %s is required super class", field.Name)
			}
			f, ok := t.SuperClass.InstanceFields.Get(field.Name)
			if !ok {
				return fmt.Errorf("field %s missing on super class", field.Name)
			}
			if !f.IsAbstract() && !f.IsVirtual() {
				return fmt.Errorf("field %s must be abstract/virtual on super class", field.Name)
			}
		} else {
			if t.SuperClass == nil {
				continue
			}
			_, ok := t.SuperClass.InstanceFields.Get(field.Name)
			if ok {
				return fmt.Errorf("field %s is not defined in %s", field.Name, t.Name)
			}
		}
	}
	for _, field := range t.StaticFields.Data {
		if field.IsOverride() {
			if t.SuperClass == nil {
				return fmt.Errorf("override %s is required super class", field.Name)
			}

			f, ok := t.SuperClass.StaticFields.Get(field.Name)
			if !ok {
				return fmt.Errorf("field %s missing on super class", field.Name)
			}
			if !f.IsAbstract() && !f.IsVirtual() {
				return fmt.Errorf("field %s must be abstract/virtual on super class", field.Name)
			}
		} else {
			if t.SuperClass == nil {
				continue
			}
			_, ok := t.SuperClass.StaticFields.Get(field.Name)
			if ok {
				return fmt.Errorf("field %s is not defined in %s", field.Name, t.Name)
			}
		}
	}
	return nil
}

func checkConstructorName(t *ast.ClassType) error {
	for _, constructor := range t.Constructors {
		if t.Name != constructor.Name {
			return fmt.Errorf("Invalid constructor name: %s", constructor.Name)
		}
	}
	return nil
}

func ParameterClassTypes(parameters []*ast.Parameter) []*ast.ClassType {
	types := make([]*ast.ClassType, len(parameters))
	for i, p := range parameters {
		types[i] = p.Type
	}
	return types
}

func MethodSignature(owner *ast.ClassType, m *ast.Method) string {
	typeStrings := make([]string, len(m.Parameters))
	for i, param := range m.Parameters {
		typeStrings[i] = param.Type.Name
	}
	returnTypeName := "void"
	if m.ReturnType != nil {
		returnTypeName = m.ReturnType.Name
	}
	return fmt.Sprintf("%s %s.%s(%s)", returnTypeName, owner.Name, m.Name, strings.Join(typeStrings, ", "))
}
