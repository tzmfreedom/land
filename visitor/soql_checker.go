package visitor

import (
	"github.com/tzmfreedom/goland/ast"
)

type SoqlChecker struct{}

func (v *SoqlChecker) VisitClassDeclaration(n *ast.ClassDeclaration) (interface{}, error) {
	return ast.VisitClassDeclaration(v, n)
}

func (v *SoqlChecker) VisitModifier(n *ast.Modifier) (interface{}, error) {
	return ast.VisitModifier(v, n)
}

func (v *SoqlChecker) VisitAnnotation(n *ast.Annotation) (interface{}, error) {
	return ast.VisitAnnotation(v, n)
}

func (v *SoqlChecker) VisitInterface(n *ast.Interface) (interface{}, error) {
	return ast.VisitInterface(v, n)
}

func (v *SoqlChecker) VisitIntegerLiteral(n *ast.IntegerLiteral) (interface{}, error) {
	return ast.VisitIntegerLiteral(v, n)
}

func (v *SoqlChecker) VisitParameter(n *ast.Parameter) (interface{}, error) {
	return ast.VisitParameter(v, n)
}

func (v *SoqlChecker) VisitArrayAccess(n *ast.ArrayAccess) (interface{}, error) {
	return ast.VisitArrayAccess(v, n)
}

func (v *SoqlChecker) VisitBooleanLiteral(n *ast.BooleanLiteral) (interface{}, error) {
	return ast.VisitBooleanLiteral(v, n)
}

func (v *SoqlChecker) VisitBreak(n *ast.Break) (interface{}, error) {
	return ast.VisitBreak(v, n)
}

func (v *SoqlChecker) VisitContinue(n *ast.Continue) (interface{}, error) {
	return ast.VisitContinue(v, n)
}

func (v *SoqlChecker) VisitDml(n *ast.Dml) (interface{}, error) {
	return ast.VisitDml(v, n)
}

func (v *SoqlChecker) VisitDoubleLiteral(n *ast.DoubleLiteral) (interface{}, error) {
	return ast.VisitDoubleLiteral(v, n)
}

func (v *SoqlChecker) VisitFieldDeclaration(n *ast.FieldDeclaration) (interface{}, error) {
	return ast.VisitFieldDeclaration(v, n)
}

func (v *SoqlChecker) VisitFieldVariable(n *ast.FieldVariable) (interface{}, error) {
	return ast.VisitFieldVariable(v, n)
}

func (v *SoqlChecker) VisitTry(n *ast.Try) (interface{}, error) {
	return ast.VisitTry(v, n)
}

func (v *SoqlChecker) VisitCatch(n *ast.Catch) (interface{}, error) {
	return ast.VisitCatch(v, n)
}

func (v *SoqlChecker) VisitFinally(n *ast.Finally) (interface{}, error) {
	return ast.VisitFinally(v, n)
}

func (v *SoqlChecker) VisitFor(n *ast.For) (interface{}, error) {
	return ast.VisitFor(v, n)
}

func (v *SoqlChecker) VisitForEnum(n *ast.ForEnum) (interface{}, error) {
	return ast.VisitForEnum(v, n)
}

func (v *SoqlChecker) VisitForControl(n *ast.ForControl) (interface{}, error) {
	return ast.VisitForControl(v, n)
}

func (v *SoqlChecker) VisitEnhancedForControl(n *ast.EnhancedForControl) (interface{}, error) {
	return ast.VisitEnhancedForControl(v, n)
}

func (v *SoqlChecker) VisitIf(n *ast.If) (interface{}, error) {
	return ast.VisitIf(v, n)
}

func (v *SoqlChecker) VisitMethodDeclaration(n *ast.MethodDeclaration) (interface{}, error) {
	return ast.VisitMethodDeclaration(v, n)
}

func (v *SoqlChecker) VisitMethodInvocation(n *ast.MethodInvocation) (interface{}, error) {
	return ast.VisitMethodInvocation(v, n)
}

func (v *SoqlChecker) VisitNew(n *ast.New) (interface{}, error) {
	return ast.VisitNew(v, n)
}

func (v *SoqlChecker) VisitNullLiteral(n *ast.NullLiteral) (interface{}, error) {
	return ast.VisitNullLiteral(v, n)
}

func (v *SoqlChecker) VisitUnaryOperator(n *ast.UnaryOperator) (interface{}, error) {
	return ast.VisitUnaryOperator(v, n)
}

func (v *SoqlChecker) VisitBinaryOperator(n *ast.BinaryOperator) (interface{}, error) {
	return ast.VisitBinaryOperator(v, n)
}

func (v *SoqlChecker) VisitReturn(n *ast.Return) (interface{}, error) {
	return ast.VisitReturn(v, n)
}

func (v *SoqlChecker) VisitThrow(n *ast.Throw) (interface{}, error) {
	return ast.VisitThrow(v, n)
}

func (v *SoqlChecker) VisitSoql(n *ast.Soql) (interface{}, error) {
	if isDecendants(n, "For") &&
		!isParent(n, "Return") &&
		!isParent(n, "For") {
		panic("SOQL IN FOR LOOP")
	}

	if isDecendants(n, "While") {
		panic("SOQL IN WHILE LOOP")
	}
	return nil, nil
}

func (v *SoqlChecker) VisitSosl(n *ast.Sosl) (interface{}, error) {
	return ast.VisitSosl(v, n)
}

func (v *SoqlChecker) VisitStringLiteral(n *ast.StringLiteral) (interface{}, error) {
	return ast.VisitStringLiteral(v, n)
}

func (v *SoqlChecker) VisitSwitch(n *ast.Switch) (interface{}, error) {
	return ast.VisitSwitch(v, n)
}

func (v *SoqlChecker) VisitTrigger(n *ast.Trigger) (interface{}, error) {
	return ast.VisitTrigger(v, n)
}

func (v *SoqlChecker) VisitTriggerTiming(n *ast.TriggerTiming) (interface{}, error) {
	return ast.VisitTriggerTiming(v, n)
}

func (v *SoqlChecker) VisitVariableDeclaration(n *ast.VariableDeclaration) (interface{}, error) {
	return ast.VisitVariableDeclaration(v, n)
}

func (v *SoqlChecker) VisitVariableDeclarator(n *ast.VariableDeclarator) (interface{}, error) {
	return ast.VisitVariableDeclarator(v, n)
}

func (v *SoqlChecker) VisitWhen(n *ast.When) (interface{}, error) {
	return ast.VisitWhen(v, n)
}

func (v *SoqlChecker) VisitWhenType(n *ast.WhenType) (interface{}, error) {
	return ast.VisitWhenType(v, n)
}

func (v *SoqlChecker) VisitWhile(n *ast.While) (interface{}, error) {
	return ast.VisitWhile(v, n)
}

func (v *SoqlChecker) VisitNothingStatement(n *ast.NothingStatement) (interface{}, error) {
	return ast.VisitNothingStatement(v, n)
}

func (v *SoqlChecker) VisitCastExpression(n *ast.CastExpression) (interface{}, error) {
	return ast.VisitCastExpression(v, n)
}

func (v *SoqlChecker) VisitFieldAccess(n *ast.FieldAccess) (interface{}, error) {
	return ast.VisitFieldAccess(v, n)
}

func (v *SoqlChecker) VisitType(n *ast.Type) (interface{}, error) {
	return ast.VisitType(v, n)
}

func (v *SoqlChecker) VisitBlock(n *ast.Block) (interface{}, error) {
	return ast.VisitBlock(v, n)
}

func (v *SoqlChecker) VisitGetterSetter(n *ast.GetterSetter) (interface{}, error) {
	return ast.VisitGetterSetter(v, n)
}

func (v *SoqlChecker) VisitPropertyDeclaration(n *ast.PropertyDeclaration) (interface{}, error) {
	return ast.VisitPropertyDeclaration(v, n)
}

func (v *SoqlChecker) VisitArrayInitializer(n *ast.ArrayInitializer) (interface{}, error) {
	return ast.VisitArrayInitializer(v, n)
}

func (v *SoqlChecker) VisitArrayCreator(n *ast.ArrayCreator) (interface{}, error) {
	return ast.VisitArrayCreator(v, n)
}

func (v *SoqlChecker) VisitSoqlBindVariable(n *ast.SoqlBindVariable) (interface{}, error) {
	return ast.VisitSoqlBindVariable(v, n)
}

func (v *SoqlChecker) VisitTernalyExpression(n *ast.TernalyExpression) (interface{}, error) {
	return ast.VisitTernalyExpression(v, n)
}

func (v *SoqlChecker) VisitMapCreator(n *ast.MapCreator) (interface{}, error) {
	return ast.VisitMapCreator(v, n)
}

func (v *SoqlChecker) VisitSetCreator(n *ast.SetCreator) (interface{}, error) {
	return ast.VisitSetCreator(v, n)
}

func (v *SoqlChecker) VisitName(n *ast.Name) (interface{}, error) {
	return ast.VisitName(v, n)
}

func (v *SoqlChecker) VisitConstructorDeclaration(n *ast.ConstructorDeclaration) (interface{}, error) {
	return ast.VisitConstructorDeclaration(v, n)
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
