package ast

func VisitClassDeclaration(v Visitor, n *ClassDeclaration) (interface{}, error) {
	for _, d := range n.Declarations {
		d.Accept(v)
	}
	return nil, nil
}

func VisitModifier(v Visitor, n *Modifier) (interface{}, error) {
	return nil, nil
}

func VisitAnnotation(v Visitor, n *Annotation) (interface{}, error) {
	return nil, nil
}

func VisitInterface(v Visitor, n *Interface) (interface{}, error) {
	return nil, nil
}

func VisitIntegerLiteral(v Visitor, n *IntegerLiteral) (interface{}, error) {
	return nil, nil
}

func VisitParameter(v Visitor, n *Parameter) (interface{}, error) {
	return nil, nil
}

func VisitArrayAccess(v Visitor, n *ArrayAccess) (interface{}, error) {
	return nil, nil
}

func VisitBooleanLiteral(v Visitor, n *BooleanLiteral) (interface{}, error) {
	return nil, nil
}

func VisitBreak(v Visitor, n *Break) (interface{}, error) {
	return nil, nil
}

func VisitContinue(v Visitor, n *Continue) (interface{}, error) {
	return nil, nil
}

func VisitDml(v Visitor, n *Dml) (interface{}, error) {
	return nil, nil
}

func VisitDoubleLiteral(v Visitor, n *DoubleLiteral) (interface{}, error) {
	return nil, nil
}

func VisitFieldDeclaration(v Visitor, n *FieldDeclaration) (interface{}, error) {
	return nil, nil
}

func VisitFieldVariable(v Visitor, n *FieldVariable) (interface{}, error) {
	return nil, nil
}

func VisitTry(v Visitor, n *Try) (interface{}, error) {
	return nil, nil
}

func VisitCatch(v Visitor, n *Catch) (interface{}, error) {
	return nil, nil
}

func VisitFinally(v Visitor, n *Finally) (interface{}, error) {
	return nil, nil
}

func VisitFor(v Visitor, n *For) (interface{}, error) {
	n.Statements.Accept(v)
	return nil, nil
}

func VisitForEnum(v Visitor, n *ForEnum) (interface{}, error) {
	n.Statements.Accept(v)
	n.ListExpression.Accept(v)
	return nil, nil
}

func VisitForControl(v Visitor, n *ForControl) (interface{}, error) {
	return nil, nil
}

func VisitEnhancedForControl(v Visitor, n *EnhancedForControl) (interface{}, error) {
	return nil, nil
}

func VisitIf(v Visitor, n *If) (interface{}, error) {
	return nil, nil
}

func VisitMethodDeclaration(v Visitor, n *MethodDeclaration) (interface{}, error) {
	n.Statements.Accept(v)
	return nil, nil
}

func VisitMethodInvocation(v Visitor, n *MethodInvocation) (interface{}, error) {
	return nil, nil
}

func VisitNew(v Visitor, n *New) (interface{}, error) {
	return nil, nil
}

func VisitNullLiteral(v Visitor, n *NullLiteral) (interface{}, error) {
	return nil, nil
}

func VisitObject(v Visitor, n *Object) (interface{}, error) {
	return nil, nil
}

func VisitUnaryOperator(v Visitor, n *UnaryOperator) (interface{}, error) {
	return nil, nil
}

func VisitBinaryOperator(v Visitor, n *BinaryOperator) (interface{}, error) {
	return nil, nil
}

func VisitReturn(v Visitor, n *Return) (interface{}, error) {
	return nil, nil
}

func VisitThrow(v Visitor, n *Throw) (interface{}, error) {
	return nil, nil
}

func VisitSoql(v Visitor, n *Soql) (interface{}, error) {
	return nil, nil
}

func VisitSosl(v Visitor, n *Sosl) (interface{}, error) {
	return nil, nil
}

func VisitStringLiteral(v Visitor, n *StringLiteral) (interface{}, error) {
	return nil, nil
}

func VisitSwitch(v Visitor, n *Switch) (interface{}, error) {
	return nil, nil
}

func VisitTrigger(v Visitor, n *Trigger) (interface{}, error) {
	return nil, nil
}

func VisitTriggerTiming(v Visitor, n *TriggerTiming) (interface{}, error) {
	return nil, nil
}

func VisitVariableDeclaration(v Visitor, n *VariableDeclaration) (interface{}, error) {
	return nil, nil
}

func VisitVariableDeclarator(v Visitor, n *VariableDeclarator) (interface{}, error) {
	return nil, nil
}

func VisitWhen(v Visitor, n *When) (interface{}, error) {
	return nil, nil
}

func VisitWhenType(v Visitor, n *WhenType) (interface{}, error) {
	return nil, nil
}

func VisitWhile(v Visitor, n *While) (interface{}, error) {
	for _, s := range n.Statements {
		s.Accept(v)
	}
	return nil, nil
}

func VisitNothingStatement(v Visitor, n *NothingStatement) (interface{}, error) {
	return nil, nil
}

func VisitCastExpression(v Visitor, n *CastExpression) (interface{}, error) {
	return nil, nil
}

func VisitFieldAccess(v Visitor, n *FieldAccess) (interface{}, error) {
	return nil, nil
}

func VisitType(v Visitor, n *Type) (interface{}, error) {
	return nil, nil
}

func VisitBlock(v Visitor, n *Block) (interface{}, error) {
	return nil, nil
}

func VisitGetterSetter(v Visitor, n *GetterSetter) (interface{}, error) {
	return nil, nil
}

func VisitPropertyDeclaration(v Visitor, n *PropertyDeclaration) (interface{}, error) {
	return nil, nil
}

func VisitArrayInitializer(v Visitor, n *ArrayInitializer) (interface{}, error) {
	return nil, nil
}

func VisitArrayCreator(v Visitor, n *ArrayCreator) (interface{}, error) {
	return nil, nil
}

func VisitBlob(v Visitor, n *Blob) (interface{}, error) {
	return nil, nil
}

func VisitSoqlBindVariable(v Visitor, n *SoqlBindVariable) (interface{}, error) {
	return nil, nil
}

func VisitTernalyExpression(v Visitor, n *TernalyExpression) (interface{}, error) {
	return nil, nil
}

func VisitMapCreator(v Visitor, n *MapCreator) (interface{}, error) {
	return nil, nil
}

func VisitSetCreator(v Visitor, n *SetCreator) (interface{}, error) {
	return nil, nil
}

func VisitName(v Visitor, n *Name) (interface{}, error) {
	return nil, nil
}

func VisitConstructorDeclaration(v Visitor, n *ConstructorDeclaration) (interface{}, error) {
	return nil, nil
}
