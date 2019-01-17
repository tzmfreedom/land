package compiler

import (
	"fmt"

	"github.com/tzmfreedom/goland/ast"
)

type ClassChecker struct {
	Context *Context
}

func (c *ClassChecker) Check(t *ast.ClassType) error {
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

func (c *ClassChecker) checkSameParameterName(m *ast.MethodMap) error {
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

func (c *ClassChecker) checkDuplicated(m *ast.MethodMap) error {
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

// if override modifier is specified, check modifier virtual/abstract on SuperClassRef
// if SuperClassRef is abstract/virtual, check override method implementation
func (c *ClassChecker) checkOverrideMethod(t *ast.ClassType) error {
	resolver := NewTypeResolver(c.Context, false)
	for _, methods := range t.InstanceMethods.Data {
		for _, m := range methods {
			if m.IsOverride() {
				if t.SuperClass == nil {
					return fmt.Errorf("@Override specified for non-overriding method: %s", m.Name)
				}

				types := make([]*ast.ClassType, len(m.Parameters))
				for i, p := range m.Parameters {
					types[i] = p.Type
				}
				_, _, err := resolver.FindInstanceMethod(t.SuperClass, m.Name, types, MODIFIER_NO_CHECK)
				if err != nil {
					return fmt.Errorf("method %s missing on super class", m.Name)
				}
			} else {
				types := make([]*ast.ClassType, len(m.Parameters))
				for i, p := range m.Parameters {
					types[i] = p.Type
				}
				if t.SuperClassRef == nil {
					continue
				}
				_, method, _ := resolver.FindInstanceMethod(t.SuperClass, m.Name, types, MODIFIER_NO_CHECK)
				if method != nil {
				}
			}
		}
	}

	for _, methods := range t.StaticMethods.Data {
		for _, m := range methods {
			if m.IsOverride() {
				if t.SuperClassRef == nil {
					return fmt.Errorf("override %s is required super class", m.Name)
				}

				types := make([]*ast.ClassType, len(m.Parameters))
				for i, p := range m.Parameters {
					types[i] = p.Type
				}
				_, _, err := resolver.FindStaticMethod(t.SuperClass, m.Name, types, MODIFIER_NO_CHECK)
				if err != nil {
					return fmt.Errorf("method %s missing on super class", m.Name)
				}
			}
		}
	}
	return nil
}

func (c *ClassChecker) checkOverrideField(t *ast.ClassType) error {
	for _, field := range t.InstanceFields.Data {
		if field.IsOverride() {
			if t.SuperClassRef == nil {
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
	if t.SuperClass != nil {
		// TODO: implement
	}
	for _, field := range t.StaticFields.Data {
		if field.IsOverride() {
			if t.SuperClassRef == nil {
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
