package ast

func visitChildren(v Visitor, n Node) (interface{}, error) {
	children := n.GetChildren()
	for _, child := range children {
		if child == nil {
			continue
		}
		if nodes, ok := child.([]Node); ok {
			for _, node := range nodes {
				_, err := node.Accept(v)
				if err != nil {
					return nil, err
				}
			}
		} else if node, ok := child.(Node); ok {
			_, err := node.Accept(v)
			if err != nil {
				return nil, err
			}
		}
	}
	return nil, nil
}

func VisitClassDeclaration(v Visitor, n *ClassDeclaration) (interface{}, error) {
	return visitChildren(v, n)
}

func VisitModifier(v Visitor, n *Modifier) (interface{}, error) {
	return visitChildren(v, n)
}

func VisitAnnotation(v Visitor, n *Annotation) (interface{}, error) {
	return visitChildren(v, n)
}

func VisitInterfaceDeclaration(v Visitor, n *InterfaceDeclaration) (interface{}, error) {
	return visitChildren(v, n)
}

func VisitIntegerLiteral(v Visitor, n *IntegerLiteral) (interface{}, error) {
	return visitChildren(v, n)
}

func VisitParameter(v Visitor, n *Parameter) (interface{}, error) {
	return visitChildren(v, n)
}

func VisitArrayAccess(v Visitor, n *ArrayAccess) (interface{}, error) {
	return visitChildren(v, n)
}

func VisitBooleanLiteral(v Visitor, n *BooleanLiteral) (interface{}, error) {
	return visitChildren(v, n)
}

func VisitBreak(v Visitor, n *Break) (interface{}, error) {
	return visitChildren(v, n)
}

func VisitContinue(v Visitor, n *Continue) (interface{}, error) {
	return visitChildren(v, n)
}

func VisitDml(v Visitor, n *Dml) (interface{}, error) {
	return visitChildren(v, n)
}

func VisitDoubleLiteral(v Visitor, n *DoubleLiteral) (interface{}, error) {
	return visitChildren(v, n)
}

func VisitFieldDeclaration(v Visitor, n *FieldDeclaration) (interface{}, error) {
	return visitChildren(v, n)
}

func VisitTry(v Visitor, n *Try) (interface{}, error) {
	return visitChildren(v, n)
}

func VisitCatch(v Visitor, n *Catch) (interface{}, error) {
	return visitChildren(v, n)
}

func VisitFinally(v Visitor, n *Finally) (interface{}, error) {
	return visitChildren(v, n)
}

func VisitFor(v Visitor, n *For) (interface{}, error) {
	return visitChildren(v, n)
}

func VisitForControl(v Visitor, n *ForControl) (interface{}, error) {
	return visitChildren(v, n)
}

func VisitEnhancedForControl(v Visitor, n *EnhancedForControl) (interface{}, error) {
	return visitChildren(v, n)
}

func VisitIf(v Visitor, n *If) (interface{}, error) {
	return visitChildren(v, n)
}

func VisitMethodDeclaration(v Visitor, n *MethodDeclaration) (interface{}, error) {
	return visitChildren(v, n)
}

func VisitMethodInvocation(v Visitor, n *MethodInvocation) (interface{}, error) {
	return visitChildren(v, n)
}

func VisitNew(v Visitor, n *New) (interface{}, error) {
	return visitChildren(v, n)
}

func VisitNullLiteral(v Visitor, n *NullLiteral) (interface{}, error) {
	return visitChildren(v, n)
}

func VisitUnaryOperator(v Visitor, n *UnaryOperator) (interface{}, error) {
	return visitChildren(v, n)
}

func VisitBinaryOperator(v Visitor, n *BinaryOperator) (interface{}, error) {
	return visitChildren(v, n)
}

func VisitInstanceofOperator(v Visitor, n *InstanceofOperator) (interface{}, error) {
	return VisitInstanceofOperator(v, n)
}

func VisitReturn(v Visitor, n *Return) (interface{}, error) {
	return visitChildren(v, n)
}

func VisitThrow(v Visitor, n *Throw) (interface{}, error) {
	return visitChildren(v, n)
}

func VisitSoql(v Visitor, n *Soql) (interface{}, error) {
	return visitChildren(v, n)
}

func VisitSosl(v Visitor, n *Sosl) (interface{}, error) {
	return visitChildren(v, n)
}

func VisitStringLiteral(v Visitor, n *StringLiteral) (interface{}, error) {
	return visitChildren(v, n)
}

func VisitSwitch(v Visitor, n *Switch) (interface{}, error) {
	return visitChildren(v, n)
}

func VisitTrigger(v Visitor, n *Trigger) (interface{}, error) {
	return visitChildren(v, n)
}

func VisitTriggerTiming(v Visitor, n *TriggerTiming) (interface{}, error) {
	return visitChildren(v, n)
}

func VisitVariableDeclaration(v Visitor, n *VariableDeclaration) (interface{}, error) {
	return visitChildren(v, n)
}

func VisitVariableDeclarator(v Visitor, n *VariableDeclarator) (interface{}, error) {
	return visitChildren(v, n)
}

func VisitWhen(v Visitor, n *When) (interface{}, error) {
	return visitChildren(v, n)
}

func VisitWhenType(v Visitor, n *WhenType) (interface{}, error) {
	return visitChildren(v, n)
}

func VisitWhile(v Visitor, n *While) (interface{}, error) {
	return visitChildren(v, n)
}

func VisitNothingStatement(v Visitor, n *NothingStatement) (interface{}, error) {
	return visitChildren(v, n)
}

func VisitCastExpression(v Visitor, n *CastExpression) (interface{}, error) {
	return visitChildren(v, n)
}

func VisitFieldAccess(v Visitor, n *FieldAccess) (interface{}, error) {
	return visitChildren(v, n)
}

func VisitType(v Visitor, n *TypeRef) (interface{}, error) {
	return visitChildren(v, n)
}

func VisitBlock(v Visitor, n *Block) (interface{}, error) {
	return visitChildren(v, n)
}

func VisitGetterSetter(v Visitor, n *GetterSetter) (interface{}, error) {
	return visitChildren(v, n)
}

func VisitPropertyDeclaration(v Visitor, n *PropertyDeclaration) (interface{}, error) {
	return visitChildren(v, n)
}

func VisitArrayInitializer(v Visitor, n *ArrayInitializer) (interface{}, error) {
	return visitChildren(v, n)
}

func VisitArrayCreator(v Visitor, n *ArrayCreator) (interface{}, error) {
	return visitChildren(v, n)
}

func VisitSoqlBindVariable(v Visitor, n *SoqlBindVariable) (interface{}, error) {
	return visitChildren(v, n)
}

func VisitTernalyExpression(v Visitor, n *TernalyExpression) (interface{}, error) {
	return visitChildren(v, n)
}

func VisitMapCreator(v Visitor, n *MapCreator) (interface{}, error) {
	return visitChildren(v, n)
}

func VisitSetCreator(v Visitor, n *SetCreator) (interface{}, error) {
	return visitChildren(v, n)
}

func VisitName(v Visitor, n *Name) (interface{}, error) {
	return visitChildren(v, n)
}

func VisitConstructorDeclaration(v Visitor, n *ConstructorDeclaration) (interface{}, error) {
	return visitChildren(v, n)
}
