package visitor

import (
	"errors"

	"github.com/tzmfreedom/land/ast"
)

type SoqlChecker struct{}

func (v *SoqlChecker) VisitClassType(n *ast.ClassType) (interface{}, error) {
	for _, methods := range n.InstanceMethods.All() {
		for _, method := range methods {
			_, err := method.Statements.Accept(v)
			if err != nil {
				return nil, err
			}
		}
	}
	for _, methods := range n.StaticMethods.All() {
		for _, method := range methods {
			_, err := method.Statements.Accept(v)
			if err != nil {
				return nil, err
			}
		}
	}
	return nil, nil
}

func (v *SoqlChecker) VisitClassDeclaration(n *ast.ClassDeclaration) (interface{}, error) {
	return ast.VisitClassDeclaration(v, n)
}

func (v *SoqlChecker) VisitModifier(n *ast.Modifier) (interface{}, error) {
	return ast.VisitModifier(v, n)
}

func (v *SoqlChecker) VisitAnnotation(n *ast.Annotation) (interface{}, error) {
	return ast.VisitAnnotation(v, n)
}

func (v *SoqlChecker) VisitInterfaceDeclaration(n *ast.InterfaceDeclaration) (interface{}, error) {
	return ast.VisitInterfaceDeclaration(v, n)
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
	return n.Statements.Accept(v)
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

func (v *SoqlChecker) VisitInstanceofOperator(n *ast.InstanceofOperator) (interface{}, error) {
	return ast.VisitInstanceofOperator(v, n)
}

func (v *SoqlChecker) VisitReturn(n *ast.Return) (interface{}, error) {
	return ast.VisitReturn(v, n)
}

func (v *SoqlChecker) VisitThrow(n *ast.Throw) (interface{}, error) {
	return ast.VisitThrow(v, n)
}

func (v *SoqlChecker) VisitSoql(n *ast.Soql) (interface{}, error) {
	if ast.IsDecendants(n, "For") &&
		!ast.IsParent(n, "Return") &&
		!ast.IsParent(n, "For") {
		return nil, errors.New("SOQL IN FOR LOOP")
	}

	if ast.IsDecendants(n, "While") {
		return nil, errors.New("SOQL IN WHILE LOOP")
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
	for _, d := range n.Declarators {
		_, err := d.Accept(v)
		if err != nil {
			return nil, err
		}
	}
	return nil, nil
}

func (v *SoqlChecker) VisitVariableDeclarator(n *ast.VariableDeclarator) (interface{}, error) {
	_, err := n.Expression.Accept(v)
	if err != nil {
		return nil, err
	}
	return nil, nil
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

func (v *SoqlChecker) VisitType(n *ast.TypeRef) (interface{}, error) {
	return ast.VisitType(v, n)
}

func (v *SoqlChecker) VisitBlock(n *ast.Block) (interface{}, error) {
	for _, stmt := range n.Statements {
		_, err := stmt.Accept(v)
		if err != nil {
			return nil, err
		}
	}
	return nil, nil
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
