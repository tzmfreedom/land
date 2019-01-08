package compiler

import (
	"fmt"

	"github.com/tzmfreedom/goland/ast"
	"github.com/tzmfreedom/goland/builtin"
)

type ClassChecker struct {
	Context *Context
}

func (c *ClassChecker) Check(t *builtin.ClassType) error {
	if err := c.checkSameParameterName(t.InstanceMethods); err != nil {
		return err
	}
	if err := c.checkSameParameterName(t.StaticMethods); err != nil {
		return err
	}
	if err := c.checkDuplicated(t.InstanceMethods); err != nil {
		return err
	}
	if err := c.checkDuplicated(t.StaticMethods); err != nil {
		return err
	}
	if err := c.checkOverrideMethod(t); err != nil {
		return err
	}
	if err := c.checkOverrideField(t); err != nil {
		return err
	}
	return nil
}

func (c *ClassChecker) checkSameParameterName(m *builtin.MethodMap) error {
	for _, methods := range m.Data {
		for _, method := range methods {
			m := method.(*ast.MethodDeclaration)
			parameterNames := map[string]struct{}{}
			for _, p := range m.Parameters {
				name := p.(*ast.Parameter).Name
				if _, ok := parameterNames[name]; ok {
					return fmt.Errorf("parameter name is duplicated: %s", name)
				}
				parameterNames[name] = struct{}{}
			}
		}
	}
	return nil
}

func (c *ClassChecker) checkDuplicated(m *builtin.MethodMap) error {
	resolver := &TypeResolver{Context: c.Context}
	for _, methods := range m.Data {
		if len(methods) == 1 {
			continue
		}
		for i, method := range methods {
			m := method.(*ast.MethodDeclaration)
			l := len(m.Parameters)
			for _, other := range methods[i+1:] {
				otherDeclaration := other.(*ast.MethodDeclaration)
				if len(otherDeclaration.Parameters) != l {
					continue
				}
				match := true
				for i, p := range m.Parameters {
					otherParam := otherDeclaration.Parameters[i].(*ast.Parameter)
					otherParamType, _ := resolver.ResolveType(otherParam.Type.(*ast.TypeRef).Name)
					methodParam := p.(*ast.Parameter)
					methodParamType, _ := resolver.ResolveType(methodParam.Type.(*ast.TypeRef).Name)
					if methodParamType != otherParamType {
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

// if override modifier is specified, check modifier virtual/abstract on superclass
// if superclass is abstract/virtual, check override method implementation
func (c *ClassChecker) checkOverrideMethod(t *builtin.ClassType) error {
	resolver := &TypeResolver{Context: c.Context}
	var super *builtin.ClassType
	if t.SuperClass != nil {
		super, _ = resolver.ResolveType(t.SuperClass.(*ast.TypeRef).Name)
	}
	for _, methods := range t.InstanceMethods.Data {
		for _, method := range methods {
			m := method.(*ast.MethodDeclaration)
			if m.IsOverride() {
				if t.SuperClass == nil {
					return fmt.Errorf("@Override specified for non-overriding method: %s", m.Name)
				}

				types := make([]*builtin.ClassType, len(m.Parameters))
				for i, p := range m.Parameters {
					types[i], _ = resolver.ResolveType(p.(*ast.Parameter).Type.(*ast.TypeRef).Name)
				}
				_, err := resolver.FindInstanceMethod(super, m.Name, types, MODIFIER_NO_CHECK)
				if err != nil {
					return fmt.Errorf("method %s missing on super class", m.Name)
				}
			} else {
				types := make([]*builtin.ClassType, len(m.Parameters))
				for i, p := range m.Parameters {
					types[i], _ = resolver.ResolveType(p.(*ast.Parameter).Type.(*ast.TypeRef).Name)
				}
				method, _ := resolver.FindInstanceMethod(super, m.Name, types, MODIFIER_NO_CHECK)
				if method != nil {
				}
			}
		}
	}

	for _, methods := range t.StaticMethods.Data {
		for _, method := range methods {
			m := method.(*ast.MethodDeclaration)
			if m.IsOverride() {
				if t.SuperClass == nil {
					return fmt.Errorf("override %s is required super class", m.Name)
				}

				types := make([]*builtin.ClassType, len(m.Parameters))
				for i, p := range m.Parameters {
					types[i], _ = resolver.ResolveType(p.(*ast.Parameter).Type.(*ast.TypeRef).Name)
				}
				_, err := resolver.FindStaticMethod(super, m.Name, types, MODIFIER_NO_CHECK)
				if err != nil {
					return fmt.Errorf("method %s missing on super class", m.Name)
				}
			}
		}
	}
	return nil
}

func (c *ClassChecker) checkOverrideField(t *builtin.ClassType) error {
	resolver := &TypeResolver{Context: c.Context}
	var super *builtin.ClassType
	if t.SuperClass != nil {
		super, _ = resolver.ResolveType(t.SuperClass.(*ast.TypeRef).Name)
	}

	for _, field := range t.InstanceFields.Data {
		if field.IsOverride() {
			if t.SuperClass == nil {
				return fmt.Errorf("override %s is required super class", field.Name)
			}

			f, ok := super.InstanceFields.Get(field.Name)
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

			_, ok := super.InstanceFields.Get(field.Name)
			if ok {
				return fmt.Errorf("field %s is not defined in %s", field.Name, t.Name)
			}
		}
	}
	if t.SuperClass != nil {

	}
	for _, field := range t.StaticFields.Data {
		if field.IsOverride() {
			if t.SuperClass == nil {
				return fmt.Errorf("override %s is required super class", field.Name)
			}

			f, ok := super.StaticFields.Get(field.Name)
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
			_, ok := super.StaticFields.Get(field.Name)
			if ok {
				return fmt.Errorf("field %s is not defined in %s", field.Name, t.Name)
			}
		}
	}
	return nil
}
