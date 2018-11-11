package ast

type TypeChecker struct {
	Env []interface{}
}

func (v *TypeChecker) VisitClassDeclaration(n *ClassDeclaration) (interface{}, error) {
	return VisitClassDeclaration(v, n)
}

func (v *TypeChecker) VisitModifier(n *Modifier) (interface{}, error) {
	return VisitModifier(v, n)
}

func (v *TypeChecker) VisitAnnotation(n *Annotation) (interface{}, error) {
	return VisitAnnotation(v, n)
}

func (v *TypeChecker) VisitInterface(n *Interface) (interface{}, error) {
	return VisitInterface(v, n)
}

func (v *TypeChecker) VisitIntegerLiteral(n *IntegerLiteral) (interface{}, error) {
	return VisitIntegerLiteral(v, n)
}

func (v *TypeChecker) VisitParameter(n *Parameter) (interface{}, error) {
	return VisitParameter(v, n)
}

func (v *TypeChecker) VisitArrayAccess(n *ArrayAccess) (interface{}, error) {
	return VisitArrayAccess(v, n)
}

func (v *TypeChecker) VisitBooleanLiteral(n *BooleanLiteral) (interface{}, error) {
	return VisitBooleanLiteral(v, n)
}

func (v *TypeChecker) VisitBreak(n *Break) (interface{}, error) {
	return VisitBreak(v, n)
}

func (v *TypeChecker) VisitContinue(n *Continue) (interface{}, error) {
	return VisitContinue(v, n)
}

func (v *TypeChecker) VisitDml(n *Dml) (interface{}, error) {
	return VisitDml(v, n)
}

func (v *TypeChecker) VisitDoubleLiteral(n *DoubleLiteral) (interface{}, error) {
	return VisitDoubleLiteral(v, n)
}

func (v *TypeChecker) VisitFieldDeclaration(n *FieldDeclaration) (interface{}, error) {
	return VisitFieldDeclaration(v, n)
}

func (v *TypeChecker) VisitFieldVariable(n *FieldVariable) (interface{}, error) {
	return VisitFieldVariable(v, n)
}

func (v *TypeChecker) VisitTry(n *Try) (interface{}, error) {
	return VisitTry(v, n)
}

func (v *TypeChecker) VisitCatch(n *Catch) (interface{}, error) {
	return VisitCatch(v, n)
}

func (v *TypeChecker) VisitFinally(n *Finally) (interface{}, error) {
	return VisitFinally(v, n)
}

func (v *TypeChecker) VisitFor(n *For) (interface{}, error) {
	return VisitFor(v, n)
}

func (v *TypeChecker) VisitForEnum(n *ForEnum) (interface{}, error) {
	return VisitForEnum(v, n)
}

func (v *TypeChecker) VisitForControl(n *ForControl) (interface{}, error) {
	return VisitForControl(v, n)
}

func (v *TypeChecker) VisitEnhancedForControl(n *EnhancedForControl) (interface{}, error) {
	return VisitEnhancedForControl(v, n)
}

func (v *TypeChecker) VisitIf(n *If) (interface{}, error) {
	return VisitIf(v, n)
}

func (v *TypeChecker) VisitMethodDeclaration(n *MethodDeclaration) (interface{}, error) {
	return VisitMethodDeclaration(v, n)
}

func (v *TypeChecker) VisitMethodInvocation(n *MethodInvocation) (interface{}, error) {
	return VisitMethodInvocation(v, n)
}

func (v *TypeChecker) VisitNew(n *New) (interface{}, error) {
	return VisitNew(v, n)
}

func (v *TypeChecker) VisitNullLiteral(n *NullLiteral) (interface{}, error) {
	return VisitNullLiteral(v, n)
}

func (v *TypeChecker) VisitObject(n *Object) (interface{}, error) {
	return VisitObject(v, n)
}

func (v *TypeChecker) VisitUnaryOperator(n *UnaryOperator) (interface{}, error) {
	return VisitUnaryOperator(v, n)
}

func (v *TypeChecker) VisitBinaryOperator(n *BinaryOperator) (interface{}, error) {
	return VisitBinaryOperator(v, n)
}

func (v *TypeChecker) VisitReturn(n *Return) (interface{}, error) {
	return VisitReturn(v, n)
}

func (v *TypeChecker) VisitThrow(n *Throw) (interface{}, error) {
	return VisitThrow(v, n)
}

func (v *TypeChecker) VisitSoql(n *Soql) (interface{}, error) {
	return VisitSoql(v, n)
}

func (v *TypeChecker) VisitSosl(n *Sosl) (interface{}, error) {
	return VisitSosl(v, n)
}

func (v *TypeChecker) VisitStringLiteral(n *StringLiteral) (interface{}, error) {
	return VisitStringLiteral(v, n)
}

func (v *TypeChecker) VisitSwitch(n *Switch) (interface{}, error) {
	return VisitSwitch(v, n)
}

func (v *TypeChecker) VisitTrigger(n *Trigger) (interface{}, error) {
	return VisitTrigger(v, n)
}

func (v *TypeChecker) VisitTriggerTiming(n *TriggerTiming) (interface{}, error) {
	return VisitTriggerTiming(v, n)
}

func (v *TypeChecker) VisitVariableDeclaration(n *VariableDeclaration) (interface{}, error) {
	return VisitVariableDeclaration(v, n)
}

func (v *TypeChecker) VisitVariableDeclarator(n *VariableDeclarator) (interface{}, error) {
	return VisitVariableDeclarator(v, n)
}

func (v *TypeChecker) VisitWhen(n *When) (interface{}, error) {
	return VisitWhen(v, n)
}

func (v *TypeChecker) VisitWhenType(n *WhenType) (interface{}, error) {
	return VisitWhenType(v, n)
}

func (v *TypeChecker) VisitWhile(n *While) (interface{}, error) {
	return VisitWhile(v, n)
}

func (v *TypeChecker) VisitNothingStatement(n *NothingStatement) (interface{}, error) {
	return VisitNothingStatement(v, n)
}

func (v *TypeChecker) VisitCastExpression(n *CastExpression) (interface{}, error) {
	return VisitCastExpression(v, n)
}

func (v *TypeChecker) VisitFieldAccess(n *FieldAccess) (interface{}, error) {
	return VisitFieldAccess(v, n)
}

func (v *TypeChecker) VisitType(n *Type) (interface{}, error) {
	return VisitType(v, n)
}

func (v *TypeChecker) VisitBlock(n *Block) (interface{}, error) {
	return VisitBlock(v, n)
}

func (v *TypeChecker) VisitGetterSetter(n *GetterSetter) (interface{}, error) {
	return VisitGetterSetter(v, n)
}

func (v *TypeChecker) VisitPropertyDeclaration(n *PropertyDeclaration) (interface{}, error) {
	return VisitPropertyDeclaration(v, n)
}

func (v *TypeChecker) VisitArrayInitializer(n *ArrayInitializer) (interface{}, error) {
	return VisitArrayInitializer(v, n)
}

func (v *TypeChecker) VisitArrayCreator(n *ArrayCreator) (interface{}, error) {
	return VisitArrayCreator(v, n)
}

func (v *TypeChecker) VisitBlob(n *Blob) (interface{}, error) {
	return VisitBlob(v, n)
}

func (v *TypeChecker) VisitSoqlBindVariable(n *SoqlBindVariable) (interface{}, error) {
	return VisitSoqlBindVariable(v, n)
}

func (v *TypeChecker) VisitTernalyExpression(n *TernalyExpression) (interface{}, error) {
	return VisitTernalyExpression(v, n)
}

func (v *TypeChecker) VisitMapCreator(n *MapCreator) (interface{}, error) {
	return VisitMapCreator(v, n)
}

func (v *TypeChecker) VisitSetCreator(n *SetCreator) (interface{}, error) {
	return VisitSetCreator(v, n)
}

func (v *TypeChecker) VisitName(n *Name) (interface{}, error) {
	return VisitName(v, n)
}

func (v *TypeChecker) VisitConstructorDeclaration(n *ConstructorDeclaration) (interface{}, error) {
	return VisitConstructorDeclaration(v, n)
}
