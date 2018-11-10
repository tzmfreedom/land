package main

import (
	"github.com/tzmfreedom/goland/ast"
)

type SoqlChecker struct {
	*DefaultVisitor
}

func (v *SoqlChecker) VisitClassDeclaration(n *ast.ClassDeclaration) interface{} {
	for _, d := range n.Declarations {
		d.Accept(v)
	}
	return nil
}

func (v *SoqlChecker) VisitMethodDeclaration(n *ast.MethodDeclaration) interface{} {
	n.Statements.Accept(v)
	return nil
}

func (v *SoqlChecker) VisitWhile(n *ast.While) interface{} {
	for _, s := range n.Statements {
		s.Accept(v)
	}
	return nil
}

func (v *SoqlChecker) VisitFor(n *ast.For) interface{} {
	n.Statements.Accept(v)
	return nil
}

func (v *SoqlChecker) VisitSoql(n *ast.Soql) interface{} {
	if isDecendants(n, "For") &&
		!isParent(n, "Return") &&
		!isParent(n, "For") {
		panic("SOQL IN FOR LOOP")
	}

	if isDecendants(n, "While") {
		panic("SOQL IN WHILE LOOP")
	}
	return nil
}

func isDecendants(n ast.Node, typeName string) bool {
	parent := n.GetParent()
	if parent == nil {
		return false
	}
	if parent.GetType() == typeName {
		return true
	}
	return isDecendants(parent, typeName)
}

func isParent(n ast.Node, typeName string) bool {
	parent := n.GetParent()
	if parent == nil {
		return false
	}
	if parent.GetType() == typeName {
		return true
	}
	return false
}

func (v *SoqlChecker) VisitBlock(n *ast.Block) interface{} {
	for _, stmt := range n.Statements {
		stmt.Accept(v)
	}
	return nil
}
