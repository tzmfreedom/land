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

func (c *ClassChecker) checkDuplicated(methods *builtin.MethodMap) error {
	resolver := &TypeResolver{Context: c.Context}
	for _, methods := range methods.Data {
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
					otherParam := otherDeclaration.Parameters[i]
					otherParamType, _ := resolver.ResolveType(otherParam.(*ast.TypeRef).Name)
					methodParamType, _ := resolver.ResolveType(p.(*ast.TypeRef).Name)
					if methodParamType != otherParamType {
						match = false
						break
					}
				}
				if match {
					return fmt.Errorf("method %s is duplicated", m.Name)
				}
			}
			return nil
		}
	}
	return nil
}

// if override modifier is specified, check modifier virtual/abstract on superclass
// if superclass is abstract/virtual, check override method implementation
func (c *ClassChecker) checkOverrideMethod(t *builtin.ClassType) error {
	return nil
}

func (c *ClassChecker) checkOverrideField(t *builtin.ClassType) error {
	return nil
}
