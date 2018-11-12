package compiler

import "github.com/tzmfreedom/goland/ast"

type TypeChecker struct {
	ClassTypes    []ast.ClassType
	CurrentMethod string
	CurrentClass  string
	Env           *ast.Env
	Errors        []error
}

func (v *TypeChecker) VisitClassType(n *ast.ClassType) (interface{}, error) {
	for _, f := range n.StaticFields {
		_, err := f.Accept(v)
		v.Errors = append(v.Errors, err)
	}

	for _, f := range n.InstanceFields {
		_, err := f.Accept(v)
		v.Errors = append(v.Errors, err)
	}

	for _, m := range n.StaticMethods {
		_, err := m.Accept(v)
		v.Errors = append(v.Errors, err)
	}

	for _, m := range n.InstanceMethods {
		_, err := m.Accept(v)
		v.Errors = append(v.Errors, err)
	}
	return nil, nil
}

func (v *TypeChecker) VisitClassDeclaration(n *ast.ClassDeclaration) (interface{}, error) {
	return ast.VisitClassDeclaration(v, n)
}

func (v *TypeChecker) VisitModifier(n *ast.Modifier) (interface{}, error) {
	return ast.VisitModifier(v, n)
}

func (v *TypeChecker) VisitAnnotation(n *ast.Annotation) (interface{}, error) {
	return ast.VisitAnnotation(v, n)
}

func (v *TypeChecker) VisitInterfaceDeclaration(n *ast.InterfaceDeclaration) (interface{}, error) {
	return ast.VisitInterfaceDeclaration(v, n)
}

func (v *TypeChecker) VisitIntegerLiteral(n *ast.IntegerLiteral) (interface{}, error) {
	return ast.VisitIntegerLiteral(v, n)
}

func (v *TypeChecker) VisitParameter(n *ast.Parameter) (interface{}, error) {
	return ast.VisitParameter(v, n)
}

func (v *TypeChecker) VisitArrayAccess(n *ast.ArrayAccess) (interface{}, error) {
	return ast.VisitArrayAccess(v, n)
}

func (v *TypeChecker) VisitBooleanLiteral(n *ast.BooleanLiteral) (interface{}, error) {
	return ast.VisitBooleanLiteral(v, n)
}

func (v *TypeChecker) VisitBreak(n *ast.Break) (interface{}, error) {
	return ast.VisitBreak(v, n)
}

func (v *TypeChecker) VisitContinue(n *ast.Continue) (interface{}, error) {
	return ast.VisitContinue(v, n)
}

func (v *TypeChecker) VisitDml(n *ast.Dml) (interface{}, error) {
	return ast.VisitDml(v, n)
}

func (v *TypeChecker) VisitDoubleLiteral(n *ast.DoubleLiteral) (interface{}, error) {
	return ast.VisitDoubleLiteral(v, n)
}

func (v *TypeChecker) VisitFieldDeclaration(n *ast.FieldDeclaration) (interface{}, error) {
	return ast.VisitFieldDeclaration(v, n)
}

func (v *TypeChecker) VisitTry(n *ast.Try) (interface{}, error) {
	return ast.VisitTry(v, n)
}

func (v *TypeChecker) VisitCatch(n *ast.Catch) (interface{}, error) {
	return ast.VisitCatch(v, n)
}

func (v *TypeChecker) VisitFinally(n *ast.Finally) (interface{}, error) {
	return ast.VisitFinally(v, n)
}

func (v *TypeChecker) VisitFor(n *ast.For) (interface{}, error) {
	return ast.VisitFor(v, n)
}

func (v *TypeChecker) VisitForControl(n *ast.ForControl) (interface{}, error) {
	return ast.VisitForControl(v, n)
}

func (v *TypeChecker) VisitEnhancedForControl(n *ast.EnhancedForControl) (interface{}, error) {
	return ast.VisitEnhancedForControl(v, n)
}

func (v *TypeChecker) VisitIf(n *ast.If) (interface{}, error) {
	return ast.VisitIf(v, n)
}

func (v *TypeChecker) VisitMethodDeclaration(n *ast.MethodDeclaration) (interface{}, error) {
	return ast.VisitMethodDeclaration(v, n)
}

func (v *TypeChecker) VisitMethodInvocation(n *ast.MethodInvocation) (interface{}, error) {
	return ast.VisitMethodInvocation(v, n)
}

func (v *TypeChecker) VisitNew(n *ast.New) (interface{}, error) {
	return ast.VisitNew(v, n)
}

func (v *TypeChecker) VisitNullLiteral(n *ast.NullLiteral) (interface{}, error) {
	return ast.VisitNullLiteral(v, n)
}

func (v *TypeChecker) VisitUnaryOperator(n *ast.UnaryOperator) (interface{}, error) {
	return ast.VisitUnaryOperator(v, n)
}

func (v *TypeChecker) VisitBinaryOperator(n *ast.BinaryOperator) (interface{}, error) {
	return ast.VisitBinaryOperator(v, n)
}

func (v *TypeChecker) VisitReturn(n *ast.Return) (interface{}, error) {
	return ast.VisitReturn(v, n)
}

func (v *TypeChecker) VisitThrow(n *ast.Throw) (interface{}, error) {
	return ast.VisitThrow(v, n)
}

func (v *TypeChecker) VisitSoql(n *ast.Soql) (interface{}, error) {
	return ast.VisitSoql(v, n)
}

func (v *TypeChecker) VisitSosl(n *ast.Sosl) (interface{}, error) {
	return ast.VisitSosl(v, n)
}

func (v *TypeChecker) VisitStringLiteral(n *ast.StringLiteral) (interface{}, error) {
	return ast.VisitStringLiteral(v, n)
}

func (v *TypeChecker) VisitSwitch(n *ast.Switch) (interface{}, error) {
	return ast.VisitSwitch(v, n)
}

func (v *TypeChecker) VisitTrigger(n *ast.Trigger) (interface{}, error) {
	return ast.VisitTrigger(v, n)
}

func (v *TypeChecker) VisitTriggerTiming(n *ast.TriggerTiming) (interface{}, error) {
	return ast.VisitTriggerTiming(v, n)
}

func (v *TypeChecker) VisitVariableDeclaration(n *ast.VariableDeclaration) (interface{}, error) {
	return ast.VisitVariableDeclaration(v, n)
}

func (v *TypeChecker) VisitVariableDeclarator(n *ast.VariableDeclarator) (interface{}, error) {
	return ast.VisitVariableDeclarator(v, n)
}

func (v *TypeChecker) VisitWhen(n *ast.When) (interface{}, error) {
	return ast.VisitWhen(v, n)
}

func (v *TypeChecker) VisitWhenType(n *ast.WhenType) (interface{}, error) {
	return ast.VisitWhenType(v, n)
}

func (v *TypeChecker) VisitWhile(n *ast.While) (interface{}, error) {
	return ast.VisitWhile(v, n)
}

func (v *TypeChecker) VisitNothingStatement(n *ast.NothingStatement) (interface{}, error) {
	return ast.VisitNothingStatement(v, n)
}

func (v *TypeChecker) VisitCastExpression(n *ast.CastExpression) (interface{}, error) {
	return ast.VisitCastExpression(v, n)
}

func (v *TypeChecker) VisitFieldAccess(n *ast.FieldAccess) (interface{}, error) {
	return ast.VisitFieldAccess(v, n)
}

func (v *TypeChecker) VisitType(n *ast.Type) (interface{}, error) {
	return ast.VisitType(v, n)
}

func (v *TypeChecker) VisitBlock(n *ast.Block) (interface{}, error) {
	return ast.VisitBlock(v, n)
}

func (v *TypeChecker) VisitGetterSetter(n *ast.GetterSetter) (interface{}, error) {
	return ast.VisitGetterSetter(v, n)
}

func (v *TypeChecker) VisitPropertyDeclaration(n *ast.PropertyDeclaration) (interface{}, error) {
	return ast.VisitPropertyDeclaration(v, n)
}

func (v *TypeChecker) VisitArrayInitializer(n *ast.ArrayInitializer) (interface{}, error) {
	return ast.VisitArrayInitializer(v, n)
}

func (v *TypeChecker) VisitArrayCreator(n *ast.ArrayCreator) (interface{}, error) {
	return ast.VisitArrayCreator(v, n)
}

func (v *TypeChecker) VisitSoqlBindVariable(n *ast.SoqlBindVariable) (interface{}, error) {
	return ast.VisitSoqlBindVariable(v, n)
}

func (v *TypeChecker) VisitTernalyExpression(n *ast.TernalyExpression) (interface{}, error) {
	return ast.VisitTernalyExpression(v, n)
}

func (v *TypeChecker) VisitMapCreator(n *ast.MapCreator) (interface{}, error) {
	return ast.VisitMapCreator(v, n)
}

func (v *TypeChecker) VisitSetCreator(n *ast.SetCreator) (interface{}, error) {
	return ast.VisitSetCreator(v, n)
}

func (v *TypeChecker) VisitName(n *ast.Name) (interface{}, error) {
	return ast.VisitName(v, n)
}

func (v *TypeChecker) VisitConstructorDeclaration(n *ast.ConstructorDeclaration) (interface{}, error) {
	return ast.VisitConstructorDeclaration(v, n)
}
