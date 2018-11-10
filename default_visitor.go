package main

import (
	"github.com/tzmfreedom/goland/ast"
)

type DefaultVisitor struct{}

func (v *DefaultVisitor) VisitClassDeclaration(n *ast.ClassDeclaration) interface{} {
	return nil
}

func (v *DefaultVisitor) VisitModifier(n *ast.Modifier) interface{} {
	return nil
}

func (v *DefaultVisitor) VisitAnnotation(n *ast.Annotation) interface{} {
	return nil
}

func (v *DefaultVisitor) VisitInterface(n *ast.Interface) interface{} {
	return nil
}

func (v *DefaultVisitor) VisitIntegerLiteral(n *ast.IntegerLiteral) interface{} {
	return nil
}

func (v *DefaultVisitor) VisitParameter(n *ast.Parameter) interface{} {
	return nil
}

func (v *DefaultVisitor) VisitArrayAccess(n *ast.ArrayAccess) interface{} {
	return nil
}

func (v *DefaultVisitor) VisitBooleanLiteral(n *ast.BooleanLiteral) interface{} {
	return nil
}

func (v *DefaultVisitor) VisitBreak(n *ast.Break) interface{} {
	return nil
}

func (v *DefaultVisitor) VisitContinue(n *ast.Continue) interface{} {
	return nil
}

func (v *DefaultVisitor) VisitDml(n *ast.Dml) interface{} {
	return nil
}

func (v *DefaultVisitor) VisitDoubleLiteral(n *ast.DoubleLiteral) interface{} {
	return nil
}

func (v *DefaultVisitor) VisitFieldDeclaration(n *ast.FieldDeclaration) interface{} {
	return nil
}

func (v *DefaultVisitor) VisitFieldVariable(n *ast.FieldVariable) interface{} {
	return nil
}

func (v *DefaultVisitor) VisitTry(n *ast.Try) interface{} {
	return nil
}

func (v *DefaultVisitor) VisitCatch(n *ast.Catch) interface{} {
	return nil
}

func (v *DefaultVisitor) VisitFinally(n *ast.Finally) interface{} {
	return nil
}

func (v *DefaultVisitor) VisitFor(n *ast.For) interface{} {
	return nil
}

func (v *DefaultVisitor) VisitForEnum(n *ast.ForEnum) interface{} {
	return nil
}

func (v *DefaultVisitor) VisitForControl(n *ast.ForControl) interface{} {
	return nil
}

func (v *DefaultVisitor) VisitEnhancedForControl(n *ast.EnhancedForControl) interface{} {
	return nil
}

func (v *DefaultVisitor) VisitIf(n *ast.If) interface{} {
	return nil
}

func (v *DefaultVisitor) VisitMethodDeclaration(n *ast.MethodDeclaration) interface{} {
	return nil
}

func (v *DefaultVisitor) VisitMethodInvocation(n *ast.MethodInvocation) interface{} {
	return nil
}

func (v *DefaultVisitor) VisitNew(n *ast.New) interface{} {
	return nil
}

func (v *DefaultVisitor) VisitNullLiteral(n *ast.NullLiteral) interface{} {
	return nil
}

func (v *DefaultVisitor) VisitObject(n *ast.Object) interface{} {
	return nil
}

func (v *DefaultVisitor) VisitUnaryOperator(n *ast.UnaryOperator) interface{} {
	return nil
}

func (v *DefaultVisitor) VisitBinaryOperator(n *ast.BinaryOperator) interface{} {
	return nil
}

func (v *DefaultVisitor) VisitReturn(n *ast.Return) interface{} {
	return nil
}

func (v *DefaultVisitor) VisitThrow(n *ast.Throw) interface{} {
	return nil
}

func (v *DefaultVisitor) VisitSoql(n *ast.Soql) interface{} {
	return nil
}

func (v *DefaultVisitor) VisitSosl(n *ast.Sosl) interface{} {
	return nil
}

func (v *DefaultVisitor) VisitStringLiteral(n *ast.StringLiteral) interface{} {
	return nil
}

func (v *DefaultVisitor) VisitSwitch(n *ast.Switch) interface{} {
	return nil
}

func (v *DefaultVisitor) VisitTrigger(n *ast.Trigger) interface{} {
	return nil
}

func (v *DefaultVisitor) VisitTriggerTiming(n *ast.TriggerTiming) interface{} {
	return nil
}

func (v *DefaultVisitor) VisitVariableDeclaration(n *ast.VariableDeclaration) interface{} {
	return nil
}

func (v *DefaultVisitor) VisitVariableDeclarator(n *ast.VariableDeclarator) interface{} {
	return nil
}

func (v *DefaultVisitor) VisitWhen(n *ast.When) interface{} {
	return nil
}

func (v *DefaultVisitor) VisitWhenType(n *ast.WhenType) interface{} {
	return nil
}

func (v *DefaultVisitor) VisitWhile(n *ast.While) interface{} {
	return nil
}

func (v *DefaultVisitor) VisitNothingStatement(n *ast.NothingStatement) interface{} {
	return nil
}

func (v *DefaultVisitor) VisitCastExpression(n *ast.CastExpression) interface{} {
	return nil
}

func (v *DefaultVisitor) VisitFieldAccess(n *ast.FieldAccess) interface{} {
	return nil
}

func (v *DefaultVisitor) VisitType(n *ast.Type) interface{} {
	return nil
}

func (v *DefaultVisitor) VisitBlock(n *ast.Block) interface{} {
	return nil
}

func (v *DefaultVisitor) VisitGetterSetter(n *ast.GetterSetter) interface{} {
	return nil
}

func (v *DefaultVisitor) VisitPropertyDeclaration(n *ast.PropertyDeclaration) interface{} {
	return nil
}

func (v *DefaultVisitor) VisitArrayInitializer(n *ast.ArrayInitializer) interface{} {
	return nil
}

func (v *DefaultVisitor) VisitArrayCreator(n *ast.ArrayCreator) interface{} {
	return nil
}

func (v *DefaultVisitor) VisitBlob(n *ast.Blob) interface{} {
	return nil
}

func (v *DefaultVisitor) VisitSoqlBindVariable(n *ast.SoqlBindVariable) interface{} {
	return nil
}

func (v *DefaultVisitor) VisitTernalyExpression(n *ast.TernalyExpression) interface{} {
	return nil
}

func (v *DefaultVisitor) VisitMapCreator(n *ast.MapCreator) interface{} {
	return nil
}

func (v *DefaultVisitor) VisitSetCreator(n *ast.SetCreator) interface{} {
	return nil
}

func (v *DefaultVisitor) VisitName(n *ast.Name) interface{} {
	return nil
}

func (v *DefaultVisitor) VisitConstructorDeclaration(n *ast.ConstructorDeclaration) interface{} {
	return nil
}
