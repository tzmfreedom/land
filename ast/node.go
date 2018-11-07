package ast

type Position struct {
	FileName string
	Column   int
	Line     int
}

type ClassDeclaration struct {
	Annotations      []Annotation
	Modifiers        []Modifier
	Name             string
	SuperClass       []Node
	ImplementClasses []string
	InstanceFields   []map[string]string
	InstanceMethods  []map[string]string
	StaticFields     []map[string]string
	StaticMethods    []map[string]string
	InnerClasses     []string
	Position         *Position
}

type Modifier struct {
	Name     string
	Position *Position
}

type Annotation struct {
	Name       Name
	Parameters []Parameter
	Position   *Position
}

type Interface struct {
	Annotations     []Annotation
	Modifiers       []Modifier
	Name            Name
	SuperClass      []Node
	InstanceMethods []map[string]string
	StaticMethods   []map[string]string
	Position        *Position
}

type IntegerLiteral struct {
	Value    int
	Position *Position
}

type Parameter struct {
	Modifiers []Modifier
	Type      Type
	Name      string
	Position  *Position
}

type ArrayAccess struct {
	Receiver Node
	Key      Node
	Position *Position
}

type BooleanLiteral struct {
	Value    bool
	Position *Position
}

type Break struct {
	Position *Position
}

type Continue struct {
	Position *Position
}

type Dml struct {
	Type       Type
	Expression Node
	UpsertKey  string
	Position   *Position
}

type DoubleLiteral struct {
	Value    float32
	Position *Position
}

type FieldDeclaration struct {
	Type        Type
	Modifiers   []Modifier
	Declarators []Node
	Position    *Position
}

type FieldVariable struct {
	Type       Type
	Modifiers  []Modifier
	Expression Node
	Position   *Position
	Getter     Node
	Setter     Node
}

type Try struct {
	Block        Block
	CatchClause  []Node
	FinallyBlock Block
	Position     *Position
}

type Catch struct {
	Modifiers  []Modifier
	Type       Type
	Identifier string
	Block      Block
	Position   *Position
}

type Finally struct {
	Black    []Node
	Position *Position
}

type For struct {
	Control    ForControl
	Statements []Node
	Position   *Position
}

type ForEnum struct {
	Type           Type
	Identifier     Node
	ListExpression Node
	Statements     []Node
	Position       *Position
}

type ForControl struct {
	ForInit    Node
	Expression Node
	ForUpdate  Node
	Position   *Position
}

type EnhancedForControl struct {
	Modifiers            []Modifier
	Type                 Type
	VariableDeclaratorId Node
	Expression           Node
	Position             *Position
}

type If struct {
	Condition     Node
	IfStatement   Node
	ElseStatement Node
	Position      *Position
}

type MethodDeclaration struct {
	Name           string
	Modifiers      []Modifier
	ReturnType     Type
	Parameters     []Parameter
	Throws         []Node
	Statements     []Node
	NativeFunction Node
	Position       *Position
}

type MethodInvocation struct {
	NameOrExpression Node
	Parameters       []Parameter
	Position         *Position
}

type New struct {
	Type       Type
	Parameters []Parameter
	Position   *Position
}

type NullLiteral struct {
	Position *Position
}

type Object struct {
	ClassType      Type
	InstanceFields []Node
	GenericType    string
	Position       *Position
}

type UnaryOperator struct {
	Op         string
	Expression Node
	IsPrefix   bool
	Position   *Position
}

type BinaryOperator struct {
	Op       string
	Left     Node
	Right    Node
	Position *Position
}

type Return struct {
	Expression Node
	Position   *Position
}

type Throw struct {
	Expression Node
	Position   *Position
}

type Soql struct {
	SelectFields []string
	FromObject   string
	Where        []string
	Order        string
	Limit        int
	Position     *Position
}

type Sosl struct{}

type StringLiteral struct {
	Value    string
	Position *Position
}

type Switch struct {
	Expression     Node
	WhenStatements []Node
	ElseStatement  Node
	Position       *Position
}

type Trigger struct {
	Name           string
	Object         string
	TriggerTimings []TriggerTiming
	Statements     []Node
	Position       *Position
}

type TriggerTiming struct {
	Timing   string
	Dml      string
	Position *Position
}

type VariableDeclaration struct {
	Modifiers   []Modifier
	Type        Type
	Declarators []VariableDeclarator
	Position    *Position
}

type VariableDeclarator struct {
	Name       string
	Expression []Node
	Position   *Position
}

type When struct {
	Condition  Node
	Statements []Node
	Position   *Position
}

type WhenType struct {
	Type       Type
	Identifier string
	Position   *Position
}

type While struct {
	Condition  Node
	Statements []Node
	IsDo       bool
	Position   *Position
}

// TOTO: when to use?
type NothingStatement struct{}

type CastExpression struct {
	CastType   Type
	Expression Node
	Position   *Position
}

type FieldAccess struct {
	Expression Node
	FieldName  string
	Position   *Position
}

type Type struct {
	Name       string
	Parameters []Node
	Position   *Position
}

type Block struct {
	Statements []Node
	Position   *Position
}

type GetterSetter struct {
	Type       Type
	Modifiers  []Modifier
	MethodBody Block
	Position   *Position
}

type PropertyDeclaration struct {
	Modifiers     []Modifier
	Type          Type
	Identifier    string
	GetterSetters Node
	Position      *Position
}

type ArrayInitializer struct {
	Initializers []Node
	Position     *Position
}

type ArrayCreator struct {
	Dim              int
	Expressions      []Node
	ArrayInitializer Node
	Position         *Position
}

type Blob struct {
	Value    []byte
	Position *Position
}

type SoqlBindVariable struct {
	Expression Node
	Position   *Position
}

type TernalyExpression struct {
	Condition       Node
	TrueExpression  Node
	FalseExpression Node
	Position        *Position
}

type MapCreator struct {
	Position *Position
}

type SetCreator struct {
	Position *Position
}

type Name struct {
	Value    []string
	Position *Position
}

type ConstructorDeclaration struct {
	Modifiers      []Modifier
	ReturnType     Type
	Parameters     []Parameter
	Throws         []Node
	Statements     []Node
	NativeFunction Node
	Position       *Position
}

type InterfaceDeclaration struct {
	Modifiers []Modifier
}

type Visitor interface {
	VisitClassDeclaration(Node) interface{}
	VisitModifier(Node) interface{}
	VisitAnnotation(Node) interface{}
	VisitInterface(Node) interface{}
	VisitIntegerLiteral(Node) interface{}
	VisitParameter(Node) interface{}
	VisitArrayAccess(Node) interface{}
	VisitBooleanLiteral(Node) interface{}
	VisitBreak(Node) interface{}
	VisitContinue(Node) interface{}
	VisitDml(Node) interface{}
	VisitDoubleLiteral(Node) interface{}
	VisitFieldDeclaration(Node) interface{}
	VisitFieldVariable(Node) interface{}
	VisitTry(Node) interface{}
	VisitCatch(Node) interface{}
	VisitFinally(Node) interface{}
	VisitFor(Node) interface{}
	VisitForEnum(Node) interface{}
	VisitForControl(Node) interface{}
	VisitEnhancedForControl(Node) interface{}
	VisitIf(Node) interface{}
	VisitMethodDeclaration(Node) interface{}
	VisitMethodInvocation(Node) interface{}
	VisitNew(Node) interface{}
	VisitNullLiteral(Node) interface{}
	VisitObject(Node) interface{}
	VisitUnaryOperator(Node) interface{}
	VisitBinaryOperator(Node) interface{}
	VisitReturn(Node) interface{}
	VisitThrow(Node) interface{}
	VisitSoql(Node) interface{}
	VisitSosl(Node) interface{}
	VisitStringLiteral(Node) interface{}
	VisitSwitch(Node) interface{}
	VisitTrigger(Node) interface{}
	VisitTriggerTiming(Node) interface{}
	VisitVariableDeclaration(Node) interface{}
	VisitVariableDeclarator(Node) interface{}
	VisitWhen(Node) interface{}
	VisitWhenType(Node) interface{}
	VisitWhile(Node) interface{}
	VisitNothingStatement(Node) interface{}
	VisitCastExpression(Node) interface{}
	VisitFieldAccess(Node) interface{}
	VisitType(Node) interface{}
	VisitBlock(Node) interface{}
	VisitGetterSetter(Node) interface{}
	VisitPropertyDeclaration(Node) interface{}
	VisitArrayInitializer(Node) interface{}
	VisitArrayCreator(Node) interface{}
	VisitBlob(Node) interface{}
	VisitSoqlBindVariable(Node) interface{}
	VisitTernalyExpression(Node) interface{}
	VisitMapCreator(Node) interface{}
	VisitSetCreator(Node) interface{}
	VisitName(Node) interface{}
	VisitConstructorDeclaration(Node) interface{}
}

var VoidType = &Type{Name: "void"}

type Node interface {
	Accept(Visitor) interface{}
}

func (n *ClassDeclaration) Accept(v Visitor) interface{} {
	return v.VisitClassDeclaration(n)
}

func (n *Modifier) Accept(v Visitor) interface{} {
	return v.VisitModifier(n)
}

func (n *Annotation) Accept(v Visitor) interface{} {
	return v.VisitAnnotation(n)
}

func (n *Interface) Accept(v Visitor) interface{} {
	return v.VisitInterface(n)
}

func (n *IntegerLiteral) Accept(v Visitor) interface{} {
	return v.VisitIntegerLiteral(n)
}

func (n *Parameter) Accept(v Visitor) interface{} {
	return v.VisitParameter(n)
}

func (n *ArrayAccess) Accept(v Visitor) interface{} {
	return v.VisitArrayAccess(n)
}

func (n *BooleanLiteral) Accept(v Visitor) interface{} {
	return v.VisitBooleanLiteral(n)
}

func (n *Break) Accept(v Visitor) interface{} {
	return v.VisitBreak(n)
}

func (n *Continue) Accept(v Visitor) interface{} {
	return v.VisitContinue(n)
}

func (n *Dml) Accept(v Visitor) interface{} {
	return v.VisitDml(n)
}

func (n *DoubleLiteral) Accept(v Visitor) interface{} {
	return v.VisitDoubleLiteral(n)
}

func (n *FieldDeclaration) Accept(v Visitor) interface{} {
	return v.VisitFieldDeclaration(n)
}

func (n *FieldVariable) Accept(v Visitor) interface{} {
	return v.VisitFieldVariable(n)
}

func (n *Try) Accept(v Visitor) interface{} {
	return v.VisitTry(n)
}

func (n *Catch) Accept(v Visitor) interface{} {
	return v.VisitCatch(n)
}

func (n *Finally) Accept(v Visitor) interface{} {
	return v.VisitFinally(n)
}

func (n *For) Accept(v Visitor) interface{} {
	return v.VisitFor(n)
}

func (n *ForEnum) Accept(v Visitor) interface{} {
	return v.VisitForEnum(n)
}

func (n *ForControl) Accept(v Visitor) interface{} {
	return v.VisitForControl(n)
}

func (n *EnhancedForControl) Accept(v Visitor) interface{} {
	return v.VisitEnhancedForControl(n)
}

func (n *If) Accept(v Visitor) interface{} {
	return v.VisitIf(n)
}

func (n *MethodDeclaration) Accept(v Visitor) interface{} {
	return v.VisitMethodDeclaration(n)
}

func (n *MethodInvocation) Accept(v Visitor) interface{} {
	return v.VisitMethodInvocation(n)
}

func (n *New) Accept(v Visitor) interface{} {
	return v.VisitNew(n)
}

func (n *NullLiteral) Accept(v Visitor) interface{} {
	return v.VisitNullLiteral(n)
}

func (n *Object) Accept(v Visitor) interface{} {
	return v.VisitObject(n)
}

func (n *UnaryOperator) Accept(v Visitor) interface{} {
	return v.VisitUnaryOperator(n)
}

func (n *BinaryOperator) Accept(v Visitor) interface{} {
	return v.VisitBinaryOperator(n)
}

func (n *Return) Accept(v Visitor) interface{} {
	return v.VisitReturn(n)
}

func (n *Throw) Accept(v Visitor) interface{} {
	return v.VisitThrow(n)
}

func (n *Soql) Accept(v Visitor) interface{} {
	return v.VisitSoql(n)
}

func (n *Sosl) Accept(v Visitor) interface{} {
	return v.VisitSosl(n)
}

func (n *StringLiteral) Accept(v Visitor) interface{} {
	return v.VisitStringLiteral(n)
}

func (n *Switch) Accept(v Visitor) interface{} {
	return v.VisitSwitch(n)
}

func (n *Trigger) Accept(v Visitor) interface{} {
	return v.VisitTrigger(n)
}

func (n *TriggerTiming) Accept(v Visitor) interface{} {
	return v.VisitTriggerTiming(n)
}

func (n *VariableDeclaration) Accept(v Visitor) interface{} {
	return v.VisitVariableDeclaration(n)
}

func (n *VariableDeclarator) Accept(v Visitor) interface{} {
	return v.VisitVariableDeclarator(n)
}

func (n *When) Accept(v Visitor) interface{} {
	return v.VisitWhen(n)
}

func (n *WhenType) Accept(v Visitor) interface{} {
	return v.VisitWhenType(n)
}

func (n *While) Accept(v Visitor) interface{} {
	return v.VisitWhile(n)
}

func (n *NothingStatement) Accept(v Visitor) interface{} {
	return v.VisitNothingStatement(n)
}

func (n *CastExpression) Accept(v Visitor) interface{} {
	return v.VisitCastExpression(n)
}

func (n *FieldAccess) Accept(v Visitor) interface{} {
	return v.VisitFieldAccess(n)
}

func (n *Type) Accept(v Visitor) interface{} {
	return v.VisitType(n)
}

func (n *Block) Accept(v Visitor) interface{} {
	return v.VisitBlock(n)
}

func (n *GetterSetter) Accept(v Visitor) interface{} {
	return v.VisitGetterSetter(n)
}

func (n *PropertyDeclaration) Accept(v Visitor) interface{} {
	return v.VisitPropertyDeclaration(n)
}

func (n *ArrayInitializer) Accept(v Visitor) interface{} {
	return v.VisitArrayInitializer(n)
}

func (n *ArrayCreator) Accept(v Visitor) interface{} {
	return v.VisitArrayCreator(n)
}

func (n *Blob) Accept(v Visitor) interface{} {
	return v.VisitBlob(n)
}

func (n *SoqlBindVariable) Accept(v Visitor) interface{} {
	return v.VisitSoqlBindVariable(n)
}

func (n *TernalyExpression) Accept(v Visitor) interface{} {
	return v.VisitTernalyExpression(n)
}

func (n *MapCreator) Accept(v Visitor) interface{} {
	return v.VisitMapCreator(n)
}

func (n *SetCreator) Accept(v Visitor) interface{} {
	return v.VisitSetCreator(n)
}

func (n *Name) Accept(v Visitor) interface{} {
	return v.VisitName(n)
}

func (n *ConstructorDeclaration) Accept(v Visitor) interface{} {
	return v.VisitConstructorDeclaration(n)
}
