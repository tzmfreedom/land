package ast

type ClassRegisterVisitor struct {
	ClassEntries  []Node
	StaticEntries []Node
}

func (v *ClassRegisterVisitor) VisitClassDeclaration(n *ClassDeclaration) (interface{}, error) {
	return VisitClassDeclaration(v, n)
}

func (v *ClassRegisterVisitor) VisitModifier(n *Modifier) (interface{}, error) {
	return VisitModifier(v, n)
}

func (v *ClassRegisterVisitor) VisitAnnotation(n *Annotation) (interface{}, error) {
	return VisitAnnotation(v, n)
}

func (v *ClassRegisterVisitor) VisitInterface(n *Interface) (interface{}, error) {
	return VisitInterface(v, n)
}

func (v *ClassRegisterVisitor) VisitIntegerLiteral(n *IntegerLiteral) (interface{}, error) {
	return VisitIntegerLiteral(v, n)
}

func (v *ClassRegisterVisitor) VisitParameter(n *Parameter) (interface{}, error) {
	return VisitParameter(v, n)
}

func (v *ClassRegisterVisitor) VisitArrayAccess(n *ArrayAccess) (interface{}, error) {
	return VisitArrayAccess(v, n)
}

func (v *ClassRegisterVisitor) VisitBooleanLiteral(n *BooleanLiteral) (interface{}, error) {
	return VisitBooleanLiteral(v, n)
}

func (v *ClassRegisterVisitor) VisitBreak(n *Break) (interface{}, error) {
	return VisitBreak(v, n)
}

func (v *ClassRegisterVisitor) VisitContinue(n *Continue) (interface{}, error) {
	return VisitContinue(v, n)
}

func (v *ClassRegisterVisitor) VisitDml(n *Dml) (interface{}, error) {
	return VisitDml(v, n)
}

func (v *ClassRegisterVisitor) VisitDoubleLiteral(n *DoubleLiteral) (interface{}, error) {
	return VisitDoubleLiteral(v, n)
}

func (v *ClassRegisterVisitor) VisitFieldDeclaration(n *FieldDeclaration) (interface{}, error) {
	return VisitFieldDeclaration(v, n)
}

func (v *ClassRegisterVisitor) VisitFieldVariable(n *FieldVariable) (interface{}, error) {
	return VisitFieldVariable(v, n)
}

func (v *ClassRegisterVisitor) VisitTry(n *Try) (interface{}, error) {
	return VisitTry(v, n)
}

func (v *ClassRegisterVisitor) VisitCatch(n *Catch) (interface{}, error) {
	return VisitCatch(v, n)
}

func (v *ClassRegisterVisitor) VisitFinally(n *Finally) (interface{}, error) {
	return VisitFinally(v, n)
}

func (v *ClassRegisterVisitor) VisitFor(n *For) (interface{}, error) {
	return VisitFor(v, n)
}

func (v *ClassRegisterVisitor) VisitForEnum(n *ForEnum) (interface{}, error) {
	return VisitForEnum(v, n)
}

func (v *ClassRegisterVisitor) VisitForControl(n *ForControl) (interface{}, error) {
	return VisitForControl(v, n)
}

func (v *ClassRegisterVisitor) VisitEnhancedForControl(n *EnhancedForControl) (interface{}, error) {
	return VisitEnhancedForControl(v, n)
}

func (v *ClassRegisterVisitor) VisitIf(n *If) (interface{}, error) {
	return VisitIf(v, n)
}

func (v *ClassRegisterVisitor) VisitMethodDeclaration(n *MethodDeclaration) (interface{}, error) {
	return VisitMethodDeclaration(v, n)
}

func (v *ClassRegisterVisitor) VisitMethodInvocation(n *MethodInvocation) (interface{}, error) {
	return VisitMethodInvocation(v, n)
}

func (v *ClassRegisterVisitor) VisitNew(n *New) (interface{}, error) {
	return VisitNew(v, n)
}

func (v *ClassRegisterVisitor) VisitNullLiteral(n *NullLiteral) (interface{}, error) {
	return VisitNullLiteral(v, n)
}

func (v *ClassRegisterVisitor) VisitObject(n *Object) (interface{}, error) {
	return VisitObject(v, n)
}

func (v *ClassRegisterVisitor) VisitUnaryOperator(n *UnaryOperator) (interface{}, error) {
	return VisitUnaryOperator(v, n)
}

func (v *ClassRegisterVisitor) VisitBinaryOperator(n *BinaryOperator) (interface{}, error) {
	return VisitBinaryOperator(v, n)
}

func (v *ClassRegisterVisitor) VisitReturn(n *Return) (interface{}, error) {
	return VisitReturn(v, n)
}

func (v *ClassRegisterVisitor) VisitThrow(n *Throw) (interface{}, error) {
	return VisitThrow(v, n)
}

func (v *ClassRegisterVisitor) VisitSoql(n *Soql) (interface{}, error) {
	return VisitSoql(v, n)
}

func (v *ClassRegisterVisitor) VisitSosl(n *Sosl) (interface{}, error) {
	return VisitSosl(v, n)
}

func (v *ClassRegisterVisitor) VisitStringLiteral(n *StringLiteral) (interface{}, error) {
	return VisitStringLiteral(v, n)
}

func (v *ClassRegisterVisitor) VisitSwitch(n *Switch) (interface{}, error) {
	return VisitSwitch(v, n)
}

func (v *ClassRegisterVisitor) VisitTrigger(n *Trigger) (interface{}, error) {
	return VisitTrigger(v, n)
}

func (v *ClassRegisterVisitor) VisitTriggerTiming(n *TriggerTiming) (interface{}, error) {
	return VisitTriggerTiming(v, n)
}

func (v *ClassRegisterVisitor) VisitVariableDeclaration(n *VariableDeclaration) (interface{}, error) {
	return VisitVariableDeclaration(v, n)
}

func (v *ClassRegisterVisitor) VisitVariableDeclarator(n *VariableDeclarator) (interface{}, error) {
	return VisitVariableDeclarator(v, n)
}

func (v *ClassRegisterVisitor) VisitWhen(n *When) (interface{}, error) {
	return VisitWhen(v, n)
}

func (v *ClassRegisterVisitor) VisitWhenType(n *WhenType) (interface{}, error) {
	return VisitWhenType(v, n)
}

func (v *ClassRegisterVisitor) VisitWhile(n *While) (interface{}, error) {
	return VisitWhile(v, n)
}

func (v *ClassRegisterVisitor) VisitNothingStatement(n *NothingStatement) (interface{}, error) {
	return VisitNothingStatement(v, n)
}

func (v *ClassRegisterVisitor) VisitCastExpression(n *CastExpression) (interface{}, error) {
	return VisitCastExpression(v, n)
}

func (v *ClassRegisterVisitor) VisitFieldAccess(n *FieldAccess) (interface{}, error) {
	return VisitFieldAccess(v, n)
}

func (v *ClassRegisterVisitor) VisitType(n *Type) (interface{}, error) {
	return VisitType(v, n)
}

func (v *ClassRegisterVisitor) VisitBlock(n *Block) (interface{}, error) {
	return VisitBlock(v, n)
}

func (v *ClassRegisterVisitor) VisitGetterSetter(n *GetterSetter) (interface{}, error) {
	return VisitGetterSetter(v, n)
}

func (v *ClassRegisterVisitor) VisitPropertyDeclaration(n *PropertyDeclaration) (interface{}, error) {
	return VisitPropertyDeclaration(v, n)
}

func (v *ClassRegisterVisitor) VisitArrayInitializer(n *ArrayInitializer) (interface{}, error) {
	return VisitArrayInitializer(v, n)
}

func (v *ClassRegisterVisitor) VisitArrayCreator(n *ArrayCreator) (interface{}, error) {
	return VisitArrayCreator(v, n)
}

func (v *ClassRegisterVisitor) VisitBlob(n *Blob) (interface{}, error) {
	return VisitBlob(v, n)
}

func (v *ClassRegisterVisitor) VisitSoqlBindVariable(n *SoqlBindVariable) (interface{}, error) {
	return VisitSoqlBindVariable(v, n)
}

func (v *ClassRegisterVisitor) VisitTernalyExpression(n *TernalyExpression) (interface{}, error) {
	return VisitTernalyExpression(v, n)
}

func (v *ClassRegisterVisitor) VisitMapCreator(n *MapCreator) (interface{}, error) {
	return VisitMapCreator(v, n)
}

func (v *ClassRegisterVisitor) VisitSetCreator(n *SetCreator) (interface{}, error) {
	return VisitSetCreator(v, n)
}

func (v *ClassRegisterVisitor) VisitName(n *Name) (interface{}, error) {
	return VisitName(v, n)
}

func (v *ClassRegisterVisitor) VisitConstructorDeclaration(n *ConstructorDeclaration) (interface{}, error) {
	return VisitConstructorDeclaration(v, n)
}
