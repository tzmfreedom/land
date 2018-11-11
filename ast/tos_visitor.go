package ast

import (
	"fmt"
	"strings"
)

type TosVisitor struct {
	Content []string
}

func (v *TosVisitor) Dump(n Node, ident int) string {
	children := n.GetChildren()
	if len(children) != 0 {
		properties := make([]string, len(children))
		for i, child := range children {
			if nodes, ok := child.([]Node); ok {
				properties[i] = v.DumpArray(nodes, ident+2)
			} else if node, ok := child.(Node); ok {
				properties[i] = v.Dump(node, ident+2)
			} else {
				properties[i] = strings.Repeat(" ", ident) +
					fmt.Sprintf("%v", child)
			}
		}
		return strings.Repeat(" ", ident) +
			"(" +
			n.GetType() + "\n" +
			strings.Repeat(" ", ident+2) +
			strings.Join(properties, "\n") +
			")"
	}
	return strings.Repeat(" ", ident) + "(" + n.GetType() + ")"
}

func (v *TosVisitor) DumpArray(nodes []Node, ident int) string {
	properties := make([]string, len(nodes))
	for i, n := range nodes {
		properties[i] = v.Dump(n, 0)
	}
	return strings.Repeat(" ", ident) +
		"[" + "\n" +
		strings.Repeat(" ", ident+2) +
		strings.Join(properties, "\n") + "]"
}

func (v *TosVisitor) VisitClassDeclaration(n *ClassDeclaration) (interface{}, error) {
	return VisitClassDeclaration(v, n)
}

func (v *TosVisitor) VisitModifier(n *Modifier) (interface{}, error) {
	return VisitModifier(v, n)
}

func (v *TosVisitor) VisitAnnotation(n *Annotation) (interface{}, error) {
	return VisitAnnotation(v, n)
}

func (v *TosVisitor) VisitInterface(n *Interface) (interface{}, error) {
	return VisitInterface(v, n)
}

func (v *TosVisitor) VisitIntegerLiteral(n *IntegerLiteral) (interface{}, error) {
	return VisitIntegerLiteral(v, n)
}

func (v *TosVisitor) VisitParameter(n *Parameter) (interface{}, error) {
	return VisitParameter(v, n)
}

func (v *TosVisitor) VisitArrayAccess(n *ArrayAccess) (interface{}, error) {
	return VisitArrayAccess(v, n)
}

func (v *TosVisitor) VisitBooleanLiteral(n *BooleanLiteral) (interface{}, error) {
	return VisitBooleanLiteral(v, n)
}

func (v *TosVisitor) VisitBreak(n *Break) (interface{}, error) {
	return VisitBreak(v, n)
}

func (v *TosVisitor) VisitContinue(n *Continue) (interface{}, error) {
	return VisitContinue(v, n)
}

func (v *TosVisitor) VisitDml(n *Dml) (interface{}, error) {
	return VisitDml(v, n)
}

func (v *TosVisitor) VisitDoubleLiteral(n *DoubleLiteral) (interface{}, error) {
	return VisitDoubleLiteral(v, n)
}

func (v *TosVisitor) VisitFieldDeclaration(n *FieldDeclaration) (interface{}, error) {
	return VisitFieldDeclaration(v, n)
}

func (v *TosVisitor) VisitFieldVariable(n *FieldVariable) (interface{}, error) {
	return VisitFieldVariable(v, n)
}

func (v *TosVisitor) VisitTry(n *Try) (interface{}, error) {
	return VisitTry(v, n)
}

func (v *TosVisitor) VisitCatch(n *Catch) (interface{}, error) {
	return VisitCatch(v, n)
}

func (v *TosVisitor) VisitFinally(n *Finally) (interface{}, error) {
	return VisitFinally(v, n)
}

func (v *TosVisitor) VisitFor(n *For) (interface{}, error) {
	return VisitFor(v, n)
}

func (v *TosVisitor) VisitForEnum(n *ForEnum) (interface{}, error) {
	return VisitForEnum(v, n)
}

func (v *TosVisitor) VisitForControl(n *ForControl) (interface{}, error) {
	return VisitForControl(v, n)
}

func (v *TosVisitor) VisitEnhancedForControl(n *EnhancedForControl) (interface{}, error) {
	return VisitEnhancedForControl(v, n)
}

func (v *TosVisitor) VisitIf(n *If) (interface{}, error) {
	return VisitIf(v, n)
}

func (v *TosVisitor) VisitMethodDeclaration(n *MethodDeclaration) (interface{}, error) {
	return VisitMethodDeclaration(v, n)
}

func (v *TosVisitor) VisitMethodInvocation(n *MethodInvocation) (interface{}, error) {
	return VisitMethodInvocation(v, n)
}

func (v *TosVisitor) VisitNew(n *New) (interface{}, error) {
	return VisitNew(v, n)
}

func (v *TosVisitor) VisitNullLiteral(n *NullLiteral) (interface{}, error) {
	return VisitNullLiteral(v, n)
}

func (v *TosVisitor) VisitObject(n *Object) (interface{}, error) {
	return VisitObject(v, n)
}

func (v *TosVisitor) VisitUnaryOperator(n *UnaryOperator) (interface{}, error) {
	return VisitUnaryOperator(v, n)
}

func (v *TosVisitor) VisitBinaryOperator(n *BinaryOperator) (interface{}, error) {
	return VisitBinaryOperator(v, n)
}

func (v *TosVisitor) VisitReturn(n *Return) (interface{}, error) {
	return VisitReturn(v, n)
}

func (v *TosVisitor) VisitThrow(n *Throw) (interface{}, error) {
	return VisitThrow(v, n)
}

func (v *TosVisitor) VisitSoql(n *Soql) (interface{}, error) {
	return VisitSoql(v, n)
}

func (v *TosVisitor) VisitSosl(n *Sosl) (interface{}, error) {
	return VisitSosl(v, n)
}

func (v *TosVisitor) VisitStringLiteral(n *StringLiteral) (interface{}, error) {
	return VisitStringLiteral(v, n)
}

func (v *TosVisitor) VisitSwitch(n *Switch) (interface{}, error) {
	return VisitSwitch(v, n)
}

func (v *TosVisitor) VisitTrigger(n *Trigger) (interface{}, error) {
	return VisitTrigger(v, n)
}

func (v *TosVisitor) VisitTriggerTiming(n *TriggerTiming) (interface{}, error) {
	return VisitTriggerTiming(v, n)
}

func (v *TosVisitor) VisitVariableDeclaration(n *VariableDeclaration) (interface{}, error) {
	return VisitVariableDeclaration(v, n)
}

func (v *TosVisitor) VisitVariableDeclarator(n *VariableDeclarator) (interface{}, error) {
	return VisitVariableDeclarator(v, n)
}

func (v *TosVisitor) VisitWhen(n *When) (interface{}, error) {
	return VisitWhen(v, n)
}

func (v *TosVisitor) VisitWhenType(n *WhenType) (interface{}, error) {
	return VisitWhenType(v, n)
}

func (v *TosVisitor) VisitWhile(n *While) (interface{}, error) {
	return VisitWhile(v, n)
}

func (v *TosVisitor) VisitNothingStatement(n *NothingStatement) (interface{}, error) {
	return VisitNothingStatement(v, n)
}

func (v *TosVisitor) VisitCastExpression(n *CastExpression) (interface{}, error) {
	return VisitCastExpression(v, n)
}

func (v *TosVisitor) VisitFieldAccess(n *FieldAccess) (interface{}, error) {
	return VisitFieldAccess(v, n)
}

func (v *TosVisitor) VisitType(n *Type) (interface{}, error) {
	return VisitType(v, n)
}

func (v *TosVisitor) VisitBlock(n *Block) (interface{}, error) {
	return VisitBlock(v, n)
}

func (v *TosVisitor) VisitGetterSetter(n *GetterSetter) (interface{}, error) {
	return VisitGetterSetter(v, n)
}

func (v *TosVisitor) VisitPropertyDeclaration(n *PropertyDeclaration) (interface{}, error) {
	return VisitPropertyDeclaration(v, n)
}

func (v *TosVisitor) VisitArrayInitializer(n *ArrayInitializer) (interface{}, error) {
	return VisitArrayInitializer(v, n)
}

func (v *TosVisitor) VisitArrayCreator(n *ArrayCreator) (interface{}, error) {
	return VisitArrayCreator(v, n)
}

func (v *TosVisitor) VisitBlob(n *Blob) (interface{}, error) {
	return VisitBlob(v, n)
}

func (v *TosVisitor) VisitSoqlBindVariable(n *SoqlBindVariable) (interface{}, error) {
	return VisitSoqlBindVariable(v, n)
}

func (v *TosVisitor) VisitTernalyExpression(n *TernalyExpression) (interface{}, error) {
	return VisitTernalyExpression(v, n)
}

func (v *TosVisitor) VisitMapCreator(n *MapCreator) (interface{}, error) {
	return VisitMapCreator(v, n)
}

func (v *TosVisitor) VisitSetCreator(n *SetCreator) (interface{}, error) {
	return VisitSetCreator(v, n)
}

func (v *TosVisitor) VisitName(n *Name) (interface{}, error) {
	return VisitName(v, n)
}

func (v *TosVisitor) VisitConstructorDeclaration(n *ConstructorDeclaration) (interface{}, error) {
	return VisitConstructorDeclaration(v, n)
}
