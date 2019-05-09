package ast

import (
	"fmt"
)

type Location struct {
	FileName string
	Column   int
	Line     int
}

type ClassDeclaration struct {
	Annotations        []*Annotation
	Modifiers          []*Modifier
	Name               string
	SuperClassRef      *TypeRef
	ImplementClassRefs []*TypeRef
	Declarations       []Node
	InnerClasses       []*ClassDeclaration
	Location           *Location
	Parent             Node
}

type Modifier struct {
	Name     string
	Location *Location
	Parent   Node
}

type Annotation struct {
	Name       string
	Parameters []Node
	Location   *Location
	Parent     Node
}

type InterfaceDeclaration struct {
	Annotations []*Annotation
	Modifiers   []*Modifier
	Name        string
	Methods     []*MethodDeclaration
	Location    *Location
	Parent      Node
}

type IntegerLiteral struct {
	Value    int
	Location *Location
	Parent   Node
}

type Parameter struct {
	Modifiers []*Modifier
	TypeRef   *TypeRef
	Type      *ClassType
	Name      string
	Location  *Location
	Parent    Node
}

type ArrayAccess struct {
	Receiver Node
	Key      Node
	Location *Location
	Parent   Node
}

type BooleanLiteral struct {
	Value    bool
	Location *Location
	Parent   Node
}

type Break struct {
	Location *Location
	Parent   Node
}

type Continue struct {
	Location *Location
	Parent   Node
}

type Dml struct {
	Type       string
	Expression Node
	UpsertKey  string
	Location   *Location
	Parent     Node
}

type DoubleLiteral struct {
	Value    float64
	Location *Location
	Parent   Node
}

type FieldDeclaration struct {
	TypeRef     *TypeRef
	Type        *ClassType
	Modifiers   []*Modifier
	Annotations []*Annotation
	Declarators []*VariableDeclarator
	Location    *Location
	Parent      Node
}

type Try struct {
	Block        *Block
	CatchClause  []*Catch
	FinallyBlock *Block
	Location     *Location
	Parent       Node
}

type Catch struct {
	Modifiers  []*Modifier
	TypeRef    *TypeRef
	Type       *ClassType
	Identifier string
	Block      *Block
	Location   *Location
	Parent     Node
}

type Finally struct {
	Block    *Block
	Location *Location
	Parent   Node
}

type For struct {
	Control    Node
	Statements *Block
	Location   *Location
	Parent     Node
}

type ForControl struct {
	ForInit    []Node
	Expression Node
	ForUpdate  []Node
	Location   *Location
	Parent     Node
}

type EnhancedForControl struct {
	Modifiers            []*Modifier
	TypeRef              *TypeRef
	Type                 *ClassType
	VariableDeclaratorId string
	Expression           Node
	Location             *Location
	Parent               Node
}

type If struct {
	Condition     Node
	IfStatement   Node
	ElseStatement Node
	Location      *Location
	Parent        Node
}

type MethodDeclaration struct {
	Name        string
	Annotations []*Annotation
	Modifiers   []*Modifier
	ReturnType  *TypeRef
	Parameters  []*Parameter
	Throws      []Node
	Statements  *Block
	Location    *Location
	Parent      Node
}

type MethodInvocation struct {
	NameOrExpression Node
	Parameters       []Node
	Location         *Location
	Parent           Node
}

type New struct {
	TypeRef    *TypeRef
	Type       *ClassType
	Parameters []Node
	Init       *Init
	Location   *Location
	Parent     Node
}

type Init struct {
	Records []Node
	Values  map[Node]Node
	Sizes   []Node
}

type NullLiteral struct {
	Location *Location
	Parent   Node
}

type UnaryOperator struct {
	Op         string
	Expression Node
	IsPrefix   bool
	Location   *Location
	Parent     Node
}

type BinaryOperator struct {
	Op       string
	Left     Node
	Right    Node
	Location *Location
	Parent   Node
}

type InstanceofOperator struct {
	Op         string
	Expression Node
	TypeRef    *TypeRef
	Type       *ClassType
	Location   *Location
	Parent     Node
}

type Return struct {
	Expression Node
	Location   *Location
	Parent     Node
}

type Throw struct {
	Expression Node
	Location   *Location
	Parent     Node
}

type NoopAccepter struct{}

func (n *NoopAccepter) Accept(v Visitor) (interface{}, error) {
	return nil, nil
}

func (n *NoopAccepter) GetChildren() []interface{} {
	return []interface{}{}
}

type Soql struct {
	SelectFields []Node
	FromObject   string
	Where        Node
	Group        *Group
	Order        Node
	Limit        Node
	Offset       Node
	ExactlyOne   bool
	Location     *Location
	Parent       Node
}

type SelectField struct {
	Value    []string
	Location *Location
	Parent   Node
	*NoopAccepter
}

type SoqlFunction struct {
	Name     string
	Location *Location
	Parent   Node
	*NoopAccepter
}

type WhereBinaryOperator struct {
	Left     Node
	Right    Node
	Op       string
	Location *Location
	Parent   Node
	*NoopAccepter
}

type WhereCondition struct {
	Field      Node
	Op         string
	Expression Node
	Not        bool
	Location   *Location
	Parent     Node
	*NoopAccepter
}

type Order struct {
	Field    []Node
	Asc      bool
	Nulls    string
	Location *Location
	Parent   Node
	*NoopAccepter
}

type Group struct {
	Fields []Node
	Having Node
}

type Sosl struct {
	Location *Location
	Parent   Node
}

type StringLiteral struct {
	Value    string
	Location *Location
	Parent   Node
}

type Switch struct {
	Expression     Node
	WhenStatements []*When
	ElseStatement  *Block
	Location       *Location
	Parent         Node
}

type Trigger struct {
	Name           string
	Object         string
	TriggerTimings []Node
	Statements     *Block
	Location       *Location
	Parent         Node
}

type TriggerTiming struct {
	Timing   string
	Dml      string
	Location *Location
	Parent   Node
}

type VariableDeclaration struct {
	Modifiers   []*Modifier
	TypeRef     *TypeRef
	Type        *ClassType
	Declarators []*VariableDeclarator
	Location    *Location
	Parent      Node
}

type VariableDeclarator struct {
	Name       string
	Expression Node
	Location   *Location
	Parent     Node
}

type When struct {
	Condition  []Node
	Statements *Block
	Location   *Location
	Parent     Node
}

type WhenType struct {
	TypeRef    *TypeRef
	Type       *ClassType
	Identifier string
	Location   *Location
	Parent     Node
}

type While struct {
	Condition  Node
	Statements *Block
	IsDo       bool
	Location   *Location
	Parent     Node
}

// TOTO: when to use?
type NothingStatement struct {
	Location *Location
	Parent   Node
}

type CastExpression struct {
	CastTypeRef *TypeRef
	CastType    *ClassType
	Expression  Node
	Location    *Location
	Parent      Node
}

type FieldAccess struct {
	Expression Node
	FieldName  string
	Location   *Location
	Parent     Node
}

type TypeRef struct {
	Name       []string
	Parameters []*TypeRef
	Dimmension int
	Location   *Location
	Parent     Node
}

type Block struct {
	Statements []Node
	Location   *Location
	Parent     Node
}

type GetterSetter struct {
	Type       string
	Modifiers  []*Modifier
	MethodBody *Block
	Location   *Location
	Parent     Node
}

type PropertyDeclaration struct {
	Modifiers     []*Modifier
	Annotations   []*Annotation
	TypeRef       *TypeRef
	Type          *ClassType
	Identifier    string
	GetterSetters []*GetterSetter
	Location      *Location
	Parent        Node
}

type ArrayInitializer struct {
	Initializers []Node
	Location     *Location
	Parent       Node
}

type ArrayCreator struct {
	Dim              int
	Expressions      []Node
	ArrayInitializer Node
	Location         *Location
	Parent           Node
}

type Blob struct {
	Value    []byte
	Location *Location
	Parent   Node
}

type SoqlBindVariable struct {
	Expression Node
	Location   *Location
	Parent     Node
}

type TernalyExpression struct {
	Condition       Node
	TrueExpression  Node
	FalseExpression Node
	Location        *Location
	Parent          Node
}

type MapCreator struct {
	Location *Location
	Parent   Node
}

type SetCreator struct {
	Location *Location
	Parent   Node
}

type Name struct {
	Value    []string
	Location *Location
	Parent   Node
}

type ConstructorDeclaration struct {
	Modifiers   []*Modifier
	Annotations []*Annotation
	Name        string
	ReturnType  Node
	Parameters  []*Parameter
	Throws      []Node
	Statements  *Block
	Location    *Location
	Parent      Node
}

type Visitor interface {
	VisitClassDeclaration(*ClassDeclaration) (interface{}, error)
	VisitModifier(*Modifier) (interface{}, error)
	VisitAnnotation(*Annotation) (interface{}, error)
	VisitInterfaceDeclaration(*InterfaceDeclaration) (interface{}, error)
	VisitIntegerLiteral(*IntegerLiteral) (interface{}, error)
	VisitParameter(*Parameter) (interface{}, error)
	VisitArrayAccess(*ArrayAccess) (interface{}, error)
	VisitBooleanLiteral(*BooleanLiteral) (interface{}, error)
	VisitBreak(*Break) (interface{}, error)
	VisitContinue(*Continue) (interface{}, error)
	VisitDml(*Dml) (interface{}, error)
	VisitDoubleLiteral(*DoubleLiteral) (interface{}, error)
	VisitFieldDeclaration(*FieldDeclaration) (interface{}, error)
	VisitTry(*Try) (interface{}, error)
	VisitCatch(*Catch) (interface{}, error)
	VisitFinally(*Finally) (interface{}, error)
	VisitFor(*For) (interface{}, error)
	VisitForControl(*ForControl) (interface{}, error)
	VisitEnhancedForControl(*EnhancedForControl) (interface{}, error)
	VisitIf(*If) (interface{}, error)
	VisitMethodDeclaration(*MethodDeclaration) (interface{}, error)
	VisitMethodInvocation(*MethodInvocation) (interface{}, error)
	VisitNew(*New) (interface{}, error)
	VisitNullLiteral(*NullLiteral) (interface{}, error)
	VisitUnaryOperator(*UnaryOperator) (interface{}, error)
	VisitBinaryOperator(*BinaryOperator) (interface{}, error)
	VisitInstanceofOperator(*InstanceofOperator) (interface{}, error)
	VisitReturn(*Return) (interface{}, error)
	VisitThrow(*Throw) (interface{}, error)
	VisitSoql(*Soql) (interface{}, error)
	VisitSosl(*Sosl) (interface{}, error)
	VisitStringLiteral(*StringLiteral) (interface{}, error)
	VisitSwitch(*Switch) (interface{}, error)
	VisitTrigger(*Trigger) (interface{}, error)
	VisitTriggerTiming(*TriggerTiming) (interface{}, error)
	VisitVariableDeclaration(*VariableDeclaration) (interface{}, error)
	VisitVariableDeclarator(*VariableDeclarator) (interface{}, error)
	VisitWhen(*When) (interface{}, error)
	VisitWhenType(*WhenType) (interface{}, error)
	VisitWhile(*While) (interface{}, error)
	VisitNothingStatement(*NothingStatement) (interface{}, error)
	VisitCastExpression(*CastExpression) (interface{}, error)
	VisitFieldAccess(*FieldAccess) (interface{}, error)
	VisitType(*TypeRef) (interface{}, error)
	VisitBlock(*Block) (interface{}, error)
	VisitGetterSetter(*GetterSetter) (interface{}, error)
	VisitPropertyDeclaration(*PropertyDeclaration) (interface{}, error)
	VisitArrayInitializer(*ArrayInitializer) (interface{}, error)
	VisitArrayCreator(*ArrayCreator) (interface{}, error)
	VisitSoqlBindVariable(*SoqlBindVariable) (interface{}, error)
	VisitTernalyExpression(*TernalyExpression) (interface{}, error)
	VisitMapCreator(*MapCreator) (interface{}, error)
	VisitSetCreator(*SetCreator) (interface{}, error)
	VisitName(*Name) (interface{}, error)
	VisitConstructorDeclaration(*ConstructorDeclaration) (interface{}, error)
}

type Node interface {
	Accept(Visitor) (interface{}, error)
	GetChildren() []interface{}
	GetType() string
	GetParent() Node
	SetParent(Node)
	GetLocation() *Location
}

func (n *ClassDeclaration) Accept(v Visitor) (interface{}, error) {
	return v.VisitClassDeclaration(n)
}

func (n *ClassDeclaration) GetChildren() []interface{} {
	return []interface{}{
		n.Name,
		n.ImplementClassRefs,
		n.SuperClassRef,
		n.Annotations,
		n.Declarations,
		n.InnerClasses,
	}
}

func (n *Modifier) Accept(v Visitor) (interface{}, error) {
	return v.VisitModifier(n)
}

func (n *Modifier) GetChildren() []interface{} {
	return []interface{}{
		n.Name,
	}
}

func (n *Annotation) Accept(v Visitor) (interface{}, error) {
	return v.VisitAnnotation(n)
}

func (n *Annotation) GetChildren() []interface{} {
	return []interface{}{
		n.Name,
	}
}

func (n *InterfaceDeclaration) Accept(v Visitor) (interface{}, error) {
	return v.VisitInterfaceDeclaration(n)
}

func (n *InterfaceDeclaration) GetChildren() []interface{} {
	return []interface{}{
		n.Name,
		n.Annotations,
		n.Methods,
		n.Modifiers,
	}
}

func (n *IntegerLiteral) Accept(v Visitor) (interface{}, error) {
	return v.VisitIntegerLiteral(n)
}

func (n *IntegerLiteral) GetChildren() []interface{} {
	return []interface{}{
		n.Value,
	}
}

func (n *Parameter) Accept(v Visitor) (interface{}, error) {
	return v.VisitParameter(n)
}

func (n *Parameter) GetChildren() []interface{} {
	return []interface{}{
		n.Type,
		n.Name,
		n.Modifiers,
	}
}

func (n *ArrayAccess) Accept(v Visitor) (interface{}, error) {
	return v.VisitArrayAccess(n)
}

func (n *ArrayAccess) GetChildren() []interface{} {
	return []interface{}{
		n.Receiver,
		n.Key,
	}
}

func (n *BooleanLiteral) Accept(v Visitor) (interface{}, error) {
	return v.VisitBooleanLiteral(n)
}

func (n *BooleanLiteral) GetChildren() []interface{} {
	return nil
}

func (n *Break) Accept(v Visitor) (interface{}, error) {
	return v.VisitBreak(n)
}

func (n *Break) GetChildren() []interface{} {
	return nil
}

func (n *Continue) Accept(v Visitor) (interface{}, error) {
	return v.VisitContinue(n)
}

func (n *Continue) GetChildren() []interface{} {
	return nil
}

func (n *Dml) Accept(v Visitor) (interface{}, error) {
	return v.VisitDml(n)
}

func (n *Dml) GetChildren() []interface{} {
	return []interface{}{
		n.Type,
		n.Expression,
	}
}

func (n *DoubleLiteral) Accept(v Visitor) (interface{}, error) {
	return v.VisitDoubleLiteral(n)
}

func (n *DoubleLiteral) GetChildren() []interface{} {
	return []interface{}{
		n.Value,
	}
}

func (n *FieldDeclaration) Accept(v Visitor) (interface{}, error) {
	return v.VisitFieldDeclaration(n)
}

func (n *FieldDeclaration) GetChildren() []interface{} {
	return []interface{}{
		n.TypeRef,
		n.Declarators,
	}
}

func (n *FieldDeclaration) IsStatic() bool {
	for _, m := range n.Modifiers {
		if m.Name == "static" {
			return true
		}
	}
	return false
}

func (n *Try) Accept(v Visitor) (interface{}, error) {
	return v.VisitTry(n)
}

func (n *Try) GetChildren() []interface{} {
	return []interface{}{
		n.Block,
		n.CatchClause,
		n.FinallyBlock,
	}
}

func (n *Catch) Accept(v Visitor) (interface{}, error) {
	return v.VisitCatch(n)
}

func (n *Catch) GetChildren() []interface{} {
	return []interface{}{
		n.TypeRef,
		n.Identifier,
		n.Modifiers,
		n.Block,
	}
}

func (n *Finally) Accept(v Visitor) (interface{}, error) {
	return v.VisitFinally(n)
}

func (n *Finally) GetChildren() []interface{} {
	return []interface{}{
		n.Block,
	}
}

func (n *For) Accept(v Visitor) (interface{}, error) {
	return v.VisitFor(n)
}

func (n *For) GetChildren() []interface{} {
	return []interface{}{
		n.Statements,
		n.Control,
	}
}

func (n *ForControl) Accept(v Visitor) (interface{}, error) {
	return v.VisitForControl(n)
}

func (n *ForControl) GetChildren() []interface{} {
	return []interface{}{
		n.ForInit,
		n.Expression,
		n.ForUpdate,
	}
}

func (n *EnhancedForControl) Accept(v Visitor) (interface{}, error) {
	return v.VisitEnhancedForControl(n)
}

func (n *EnhancedForControl) GetChildren() []interface{} {
	return []interface{}{
		n.Type,
		n.VariableDeclaratorId,
		n.Expression,
	}
}

func (n *If) Accept(v Visitor) (interface{}, error) {
	return v.VisitIf(n)
}

func (n *If) GetChildren() []interface{} {
	return []interface{}{
		n.IfStatement,
		n.Condition,
	}
}

func (n *MethodDeclaration) Accept(v Visitor) (interface{}, error) {
	return v.VisitMethodDeclaration(n)
}

func (n *MethodDeclaration) GetChildren() []interface{} {
	return []interface{}{
		n.Name,
		n.Modifiers,
		n.ReturnType,
		n.Throws,
		n.Parameters,
		n.Statements,
	}
}

func (n *MethodDeclaration) IsStatic() bool {
	for _, m := range n.Modifiers {
		if m.Name == "static" {
			return true
		}
	}
	return false
}

func (n *MethodInvocation) Accept(v Visitor) (interface{}, error) {
	return v.VisitMethodInvocation(n)
}

func (n *MethodInvocation) GetChildren() []interface{} {
	return []interface{}{
		n.NameOrExpression,
		n.Parameters,
	}
}

func (n *New) Accept(v Visitor) (interface{}, error) {
	return v.VisitNew(n)
}

func (n *New) GetChildren() []interface{} {
	return []interface{}{
		n.Type,
		n.Parameters,
	}
}

func (n *NullLiteral) Accept(v Visitor) (interface{}, error) {
	return v.VisitNullLiteral(n)
}

func (n *NullLiteral) GetChildren() []interface{} {
	return nil
}

func (n *UnaryOperator) Accept(v Visitor) (interface{}, error) {
	return v.VisitUnaryOperator(n)
}

func (n *UnaryOperator) GetChildren() []interface{} {
	return []interface{}{
		n.Op,
		n.Expression,
		n.IsPrefix,
	}
}

func (n *BinaryOperator) Accept(v Visitor) (interface{}, error) {
	return v.VisitBinaryOperator(n)
}

func (n *BinaryOperator) GetChildren() []interface{} {
	return []interface{}{
		n.Op,
		n.Left,
		n.Right,
	}
}

func (n *InstanceofOperator) Accept(v Visitor) (interface{}, error) {
	return v.VisitInstanceofOperator(n)
}

func (n *InstanceofOperator) GetChildren() []interface{} {
	return []interface{}{
		n.Op,
		n.Expression,
		n.TypeRef,
	}
}

func (n *Return) Accept(v Visitor) (interface{}, error) {
	return v.VisitReturn(n)
}

func (n *Return) GetChildren() []interface{} {
	return []interface{}{
		n.Expression,
	}
}

func (n *Throw) Accept(v Visitor) (interface{}, error) {
	return v.VisitThrow(n)
}

func (n *Throw) GetChildren() []interface{} {
	return []interface{}{
		n.Expression,
	}
}

func (n *Soql) Accept(v Visitor) (interface{}, error) {
	return v.VisitSoql(n)
}

func (n *Soql) GetChildren() []interface{} {
	return []interface{}{
		n.SelectFields,
		n.FromObject,
		n.Where,
		n.Order,
		n.Limit,
	}
}

func (n *Sosl) Accept(v Visitor) (interface{}, error) {
	return v.VisitSosl(n)
}

func (n *Sosl) GetChildren() []interface{} {
	return []interface{}{}
}

func (n *StringLiteral) Accept(v Visitor) (interface{}, error) {
	return v.VisitStringLiteral(n)
}

func (n *StringLiteral) GetChildren() []interface{} {
	return []interface{}{
		n.Value,
	}
}

func (n *Switch) Accept(v Visitor) (interface{}, error) {
	return v.VisitSwitch(n)
}

func (n *Switch) GetChildren() []interface{} {
	return []interface{}{
		n.Expression,
		n.WhenStatements,
		n.ElseStatement,
	}
}

func (n *Trigger) Accept(v Visitor) (interface{}, error) {
	return v.VisitTrigger(n)
}

func (n *Trigger) GetChildren() []interface{} {
	return []interface{}{
		n.Name,
		n.Object,
		n.TriggerTimings,
		n.Statements,
	}
}

func (n *TriggerTiming) Accept(v Visitor) (interface{}, error) {
	return v.VisitTriggerTiming(n)
}

func (n *TriggerTiming) GetChildren() []interface{} {
	return []interface{}{
		n.Dml,
		n.Timing,
	}
}

func (n *VariableDeclaration) Accept(v Visitor) (interface{}, error) {
	return v.VisitVariableDeclaration(n)
}

func (n *VariableDeclaration) GetChildren() []interface{} {
	return []interface{}{
		n.Type,
		n.Modifiers,
		n.Declarators,
	}
}

func (n *VariableDeclarator) Accept(v Visitor) (interface{}, error) {
	return v.VisitVariableDeclarator(n)
}

func (n *VariableDeclarator) GetChildren() []interface{} {
	return []interface{}{
		n.Name,
		n.Expression,
	}
}

func (n *When) Accept(v Visitor) (interface{}, error) {
	return v.VisitWhen(n)
}

func (n *When) GetChildren() []interface{} {
	return []interface{}{
		n.Condition,
		n.Statements,
	}
}

func (n *WhenType) Accept(v Visitor) (interface{}, error) {
	return v.VisitWhenType(n)
}

func (n *WhenType) GetChildren() []interface{} {
	return []interface{}{
		n.Type,
		n.Identifier,
	}
}

func (n *While) Accept(v Visitor) (interface{}, error) {
	return v.VisitWhile(n)
}

func (n *While) GetChildren() []interface{} {
	return []interface{}{
		n.Condition,
		n.Statements,
		n.IsDo,
	}
}

func (n *NothingStatement) Accept(v Visitor) (interface{}, error) {
	return v.VisitNothingStatement(n)
}

func (n *NothingStatement) GetChildren() []interface{} {
	return nil
}

func (n *CastExpression) Accept(v Visitor) (interface{}, error) {
	return v.VisitCastExpression(n)
}

func (n *CastExpression) GetChildren() []interface{} {
	return []interface{}{
		n.CastType,
		n.Expression,
	}
}

func (n *FieldAccess) Accept(v Visitor) (interface{}, error) {
	return v.VisitFieldAccess(n)
}

func (n *FieldAccess) GetChildren() []interface{} {
	return []interface{}{
		n.Expression,
		n.FieldName,
	}
}

func (n *TypeRef) Accept(v Visitor) (interface{}, error) {
	return v.VisitType(n)
}

func (n *TypeRef) GetChildren() []interface{} {
	return []interface{}{
		n.Name,
		n.Parameters,
	}
}

func (n *TypeRef) IsGenerics() bool {
	name := n.Name[0]
	return name == "T:1" || name == "T:2"
}

func (n *TypeRef) IsGenericsNumber(number int) bool {
	name := n.Name[0]
	typeref := fmt.Sprintf("T:%d", number)
	return name == typeref
}

func (n *Block) Accept(v Visitor) (interface{}, error) {
	return v.VisitBlock(n)
}

func (n *Block) GetChildren() []interface{} {
	return []interface{}{
		n.Statements,
	}
}

func (n *GetterSetter) Accept(v Visitor) (interface{}, error) {
	return v.VisitGetterSetter(n)
}

func (n *GetterSetter) GetChildren() []interface{} {
	return []interface{}{
		n.Type,
		n.Modifiers,
		n.MethodBody,
	}
}

func (n *GetterSetter) IsGet() bool {
	return n.Type == "get"
}

func (n *GetterSetter) IsSet() bool {
	return n.Type == "set"
}

func (n *GetterSetter) Is(modifier string) bool {
	for _, m := range n.Modifiers {
		if m.Name == modifier {
			return true
		}
	}
	return false
}

func (n *GetterSetter) IsModifierBlank() bool {
	return n.Modifiers == nil || len(n.Modifiers) == 0
}

func (n *PropertyDeclaration) Accept(v Visitor) (interface{}, error) {
	return v.VisitPropertyDeclaration(n)
}

func (n *PropertyDeclaration) GetChildren() []interface{} {
	return []interface{}{
		n.Type,
		n.Identifier,
		n.Modifiers,
		n.GetterSetters,
	}
}

func (n *PropertyDeclaration) IsStatic() bool {
	for _, m := range n.Modifiers {
		if m.Name == "static" {
			return true
		}
	}
	return false
}

func (n *ArrayInitializer) Accept(v Visitor) (interface{}, error) {
	return v.VisitArrayInitializer(n)
}

func (n *ArrayInitializer) GetChildren() []interface{} {
	return []interface{}{
		n.Initializers,
	}
}

func (n *ArrayCreator) Accept(v Visitor) (interface{}, error) {
	return v.VisitArrayCreator(n)
}

func (n *ArrayCreator) GetChildren() []interface{} {
	return []interface{}{
		n.Dim,
		n.ArrayInitializer,
		n.Expressions,
	}
}

func (n *SoqlBindVariable) Accept(v Visitor) (interface{}, error) {
	return v.VisitSoqlBindVariable(n)
}

func (n *SoqlBindVariable) GetChildren() []interface{} {
	return []interface{}{
		n.Expression,
	}
}

func (n *TernalyExpression) Accept(v Visitor) (interface{}, error) {
	return v.VisitTernalyExpression(n)
}

func (n *TernalyExpression) GetChildren() []interface{} {
	return []interface{}{
		n.Condition,
		n.TrueExpression,
		n.FalseExpression,
	}
}

func (n *MapCreator) Accept(v Visitor) (interface{}, error) {
	return v.VisitMapCreator(n)
}

func (n *MapCreator) GetChildren() []interface{} {
	return []interface{}{}
}

func (n *SetCreator) Accept(v Visitor) (interface{}, error) {
	return v.VisitSetCreator(n)
}

func (n *SetCreator) GetChildren() []interface{} {
	return []interface{}{}
}

func (n *Name) Accept(v Visitor) (interface{}, error) {
	return v.VisitName(n)
}

func (n *Name) GetChildren() []interface{} {
	return []interface{}{
		n.Value,
	}
}

func (n *ConstructorDeclaration) Accept(v Visitor) (interface{}, error) {
	return v.VisitConstructorDeclaration(n)
}

func (n *ConstructorDeclaration) GetChildren() []interface{} {
	return []interface{}{
		n.Parameters,
		n.ReturnType,
		n.Modifiers,
		n.Throws,
		n.Statements,
	}
}

func (n *ClassDeclaration) GetType() string {
	return "ClassDeclaration"
}
func (n *Modifier) GetType() string {
	return "Modifier"
}
func (n *Annotation) GetType() string {
	return "Annotation"
}
func (n *InterfaceDeclaration) GetType() string {
	return "InterfaceDeclaration"
}
func (n *IntegerLiteral) GetType() string {
	return "Integer"
}
func (n *Parameter) GetType() string {
	return "Parameter"
}
func (n *ArrayAccess) GetType() string {
	return "ArrayAccess"
}
func (n *BooleanLiteral) GetType() string {
	return "Boolean"
}
func (n *Break) GetType() string {
	return "Break"
}
func (n *Continue) GetType() string {
	return "Continue"
}
func (n *Dml) GetType() string {
	return "Dml"
}
func (n *DoubleLiteral) GetType() string {
	return "Double"
}
func (n *FieldDeclaration) GetType() string {
	return "FieldDeclaration"
}
func (n *Try) GetType() string {
	return "Try"
}
func (n *Catch) GetType() string {
	return "Catch"
}
func (n *Finally) GetType() string {
	return "Finally"
}
func (n *For) GetType() string {
	return "For"
}
func (n *ForControl) GetType() string {
	return "ForControl"
}
func (n *EnhancedForControl) GetType() string {
	return "EnhancedForControl"
}
func (n *If) GetType() string {
	return "If"
}
func (n *MethodDeclaration) GetType() string {
	return "MethodDeclaration"
}
func (n *MethodInvocation) GetType() string {
	return "MethodInvocation"
}
func (n *New) GetType() string {
	return "New"
}
func (n *NullLiteral) GetType() string {
	return "Null"
}
func (n *UnaryOperator) GetType() string {
	return "UnaryOperator"
}
func (n *BinaryOperator) GetType() string {
	return "BinaryOperator"
}
func (n *InstanceofOperator) GetType() string {
	return "InstanceofOperator"
}
func (n *Return) GetType() string {
	return "Return"
}
func (n *Throw) GetType() string {
	return "Throw"
}
func (n *Soql) GetType() string {
	return "Soql"
}
func (n *Sosl) GetType() string {
	return "Sosl"
}
func (n *StringLiteral) GetType() string {
	return "StringLiteral"
}
func (n *Switch) GetType() string {
	return "Switch"
}
func (n *Trigger) GetType() string {
	return "Trigger"
}
func (n *TriggerTiming) GetType() string {
	return "TriggerTiming"
}
func (n *VariableDeclaration) GetType() string {
	return "VariableDeclaration"
}
func (n *VariableDeclarator) GetType() string {
	return "VariableDeclarator"
}
func (n *When) GetType() string {
	return "When"
}
func (n *WhenType) GetType() string {
	return "WhenType"
}
func (n *While) GetType() string {
	return "While"
}
func (n *NothingStatement) GetType() string {
	return "NothingStatement"
}
func (n *CastExpression) GetType() string {
	return "CastExpression"
}
func (n *FieldAccess) GetType() string {
	return "FieldAccess"
}
func (n *TypeRef) GetType() string {
	return "Type"
}
func (n *Block) GetType() string {
	return "Block"
}
func (n *GetterSetter) GetType() string {
	return "GetterSetter"
}
func (n *PropertyDeclaration) GetType() string {
	return "PropertyDeclaration"
}
func (n *ArrayInitializer) GetType() string {
	return "ArrayInitializer"
}
func (n *ArrayCreator) GetType() string {
	return "ArrayCreator"
}
func (n *SoqlBindVariable) GetType() string {
	return "SoqlBindVariable"
}
func (n *TernalyExpression) GetType() string {
	return "TernalyExpression"
}
func (n *MapCreator) GetType() string {
	return "MapCreator"
}
func (n *SetCreator) GetType() string {
	return "SetCreator"
}
func (n *Name) GetType() string {
	return "Name"
}
func (n *ConstructorDeclaration) GetType() string {
	return "ConstructorDeclaration"
}

func (n *WhereBinaryOperator) GetType() string {
	return "WhereBinaryOperator"
}

func (n *WhereCondition) GetType() string {
	return "WhereCondition"
}

func (n *SelectField) GetType() string {
	return "SelectField"
}

func (n *SoqlFunction) GetType() string {
	return "SoqlFunction"
}

func (n *Order) GetType() string {
	return "Order"
}

func (n *ClassDeclaration) GetParent() Node {
	return n.Parent
}
func (n *Modifier) GetParent() Node {
	return n.Parent
}
func (n *Annotation) GetParent() Node {
	return n.Parent
}
func (n *InterfaceDeclaration) GetParent() Node {
	return n.Parent
}
func (n *IntegerLiteral) GetParent() Node {
	return n.Parent
}
func (n *Parameter) GetParent() Node {
	return n.Parent
}
func (n *ArrayAccess) GetParent() Node {
	return n.Parent
}
func (n *BooleanLiteral) GetParent() Node {
	return n.Parent
}
func (n *Break) GetParent() Node {
	return n.Parent
}
func (n *Continue) GetParent() Node {
	return n.Parent
}
func (n *Dml) GetParent() Node {
	return n.Parent
}
func (n *DoubleLiteral) GetParent() Node {
	return n.Parent
}
func (n *FieldDeclaration) GetParent() Node {
	return n.Parent
}
func (n *Try) GetParent() Node {
	return n.Parent
}
func (n *Catch) GetParent() Node {
	return n.Parent
}
func (n *Finally) GetParent() Node {
	return n.Parent
}
func (n *For) GetParent() Node {
	return n.Parent
}
func (n *ForControl) GetParent() Node {
	return n.Parent
}
func (n *EnhancedForControl) GetParent() Node {
	return n.Parent
}
func (n *If) GetParent() Node {
	return n.Parent
}
func (n *MethodDeclaration) GetParent() Node {
	return n.Parent
}
func (n *MethodInvocation) GetParent() Node {
	return n.Parent
}
func (n *New) GetParent() Node {
	return n.Parent
}
func (n *NullLiteral) GetParent() Node {
	return n.Parent
}
func (n *UnaryOperator) GetParent() Node {
	return n.Parent
}
func (n *BinaryOperator) GetParent() Node {
	return n.Parent
}
func (n *InstanceofOperator) GetParent() Node {
	return n.Parent
}
func (n *Return) GetParent() Node {
	return n.Parent
}
func (n *Throw) GetParent() Node {
	return n.Parent
}
func (n *Soql) GetParent() Node {
	return n.Parent
}
func (n *Sosl) GetParent() Node {
	return n.Parent
}
func (n *StringLiteral) GetParent() Node {
	return n.Parent
}
func (n *Switch) GetParent() Node {
	return n.Parent
}
func (n *Trigger) GetParent() Node {
	return n.Parent
}
func (n *TriggerTiming) GetParent() Node {
	return n.Parent
}
func (n *VariableDeclaration) GetParent() Node {
	return n.Parent
}
func (n *VariableDeclarator) GetParent() Node {
	return n.Parent
}
func (n *When) GetParent() Node {
	return n.Parent
}
func (n *WhenType) GetParent() Node {
	return n.Parent
}
func (n *While) GetParent() Node {
	return n.Parent
}
func (n *NothingStatement) GetParent() Node {
	return n.Parent
}
func (n *CastExpression) GetParent() Node {
	return n.Parent
}
func (n *FieldAccess) GetParent() Node {
	return n.Parent
}
func (n *TypeRef) GetParent() Node {
	return n.Parent
}
func (n *Block) GetParent() Node {
	return n.Parent
}
func (n *GetterSetter) GetParent() Node {
	return n.Parent
}
func (n *PropertyDeclaration) GetParent() Node {
	return n.Parent
}
func (n *ArrayInitializer) GetParent() Node {
	return n.Parent
}
func (n *ArrayCreator) GetParent() Node {
	return n.Parent
}
func (n *SoqlBindVariable) GetParent() Node {
	return n.Parent
}
func (n *TernalyExpression) GetParent() Node {
	return n.Parent
}
func (n *MapCreator) GetParent() Node {
	return n.Parent
}
func (n *SetCreator) GetParent() Node {
	return n.Parent
}
func (n *Name) GetParent() Node {
	return n.Parent
}
func (n *ConstructorDeclaration) GetParent() Node {
	return n.Parent
}

func (n *WhereBinaryOperator) GetParent() Node {
	return n.Parent
}

func (n *WhereCondition) GetParent() Node {
	return n.Parent
}

func (n *SelectField) GetParent() Node {
	return n.Parent
}

func (n *SoqlFunction) GetParent() Node {
	return n.Parent
}

func (n *Order) GetParent() Node {
	return n.Parent
}

func (n *ClassDeclaration) SetParent(parent Node) {
	n.Parent = parent
}

func (n *Modifier) SetParent(parent Node) {
	n.Parent = parent
}

func (n *Annotation) SetParent(parent Node) {
	n.Parent = parent
}

func (n *InterfaceDeclaration) SetParent(parent Node) {
	n.Parent = parent
}

func (n *IntegerLiteral) SetParent(parent Node) {
	n.Parent = parent
}

func (n *Parameter) SetParent(parent Node) {
	n.Parent = parent
}

func (n *ArrayAccess) SetParent(parent Node) {
	n.Parent = parent
}

func (n *BooleanLiteral) SetParent(parent Node) {
	n.Parent = parent
}

func (n *Break) SetParent(parent Node) {
	n.Parent = parent
}

func (n *Continue) SetParent(parent Node) {
	n.Parent = parent
}

func (n *Dml) SetParent(parent Node) {
	n.Parent = parent
}

func (n *DoubleLiteral) SetParent(parent Node) {
	n.Parent = parent
}

func (n *FieldDeclaration) SetParent(parent Node) {
	n.Parent = parent
}

func (n *Try) SetParent(parent Node) {
	n.Parent = parent
}

func (n *Catch) SetParent(parent Node) {
	n.Parent = parent
}

func (n *Finally) SetParent(parent Node) {
	n.Parent = parent
}

func (n *For) SetParent(parent Node) {
	n.Parent = parent
}

func (n *ForControl) SetParent(parent Node) {
	n.Parent = parent
}

func (n *EnhancedForControl) SetParent(parent Node) {
	n.Parent = parent
}

func (n *If) SetParent(parent Node) {
	n.Parent = parent
}

func (n *MethodDeclaration) SetParent(parent Node) {
	n.Parent = parent
}

func (n *MethodInvocation) SetParent(parent Node) {
	n.Parent = parent
}

func (n *New) SetParent(parent Node) {
	n.Parent = parent
}

func (n *NullLiteral) SetParent(parent Node) {
	n.Parent = parent
}

func (n *UnaryOperator) SetParent(parent Node) {
	n.Parent = parent
}

func (n *BinaryOperator) SetParent(parent Node) {
	n.Parent = parent
}

func (n *InstanceofOperator) SetParent(parent Node) {
	n.Parent = parent
}

func (n *Return) SetParent(parent Node) {
	n.Parent = parent
}

func (n *Throw) SetParent(parent Node) {
	n.Parent = parent
}

func (n *Soql) SetParent(parent Node) {
	n.Parent = parent
}

func (n *Sosl) SetParent(parent Node) {
	n.Parent = parent
}

func (n *StringLiteral) SetParent(parent Node) {
	n.Parent = parent
}

func (n *Switch) SetParent(parent Node) {
	n.Parent = parent
}

func (n *Trigger) SetParent(parent Node) {
	n.Parent = parent
}

func (n *TriggerTiming) SetParent(parent Node) {
	n.Parent = parent
}

func (n *VariableDeclaration) SetParent(parent Node) {
	n.Parent = parent
}

func (n *VariableDeclarator) SetParent(parent Node) {
	n.Parent = parent
}

func (n *When) SetParent(parent Node) {
	n.Parent = parent
}

func (n *WhenType) SetParent(parent Node) {
	n.Parent = parent
}

func (n *While) SetParent(parent Node) {
	n.Parent = parent
}

func (n *NothingStatement) SetParent(parent Node) {
	n.Parent = parent
}

func (n *CastExpression) SetParent(parent Node) {
	n.Parent = parent
}

func (n *FieldAccess) SetParent(parent Node) {
	n.Parent = parent
}

func (n *TypeRef) SetParent(parent Node) {
	n.Parent = parent
}

func (n *Block) SetParent(parent Node) {
	n.Parent = parent
}

func (n *GetterSetter) SetParent(parent Node) {
	n.Parent = parent
}

func (n *PropertyDeclaration) SetParent(parent Node) {
	n.Parent = parent
}

func (n *ArrayInitializer) SetParent(parent Node) {
	n.Parent = parent
}

func (n *ArrayCreator) SetParent(parent Node) {
	n.Parent = parent
}

func (n *SoqlBindVariable) SetParent(parent Node) {
	n.Parent = parent
}

func (n *TernalyExpression) SetParent(parent Node) {
	n.Parent = parent
}

func (n *MapCreator) SetParent(parent Node) {
	n.Parent = parent
}

func (n *SetCreator) SetParent(parent Node) {
	n.Parent = parent
}

func (n *Name) SetParent(parent Node) {
	n.Parent = parent
}

func (n *ConstructorDeclaration) SetParent(parent Node) {
	n.Parent = parent
}

func (n *WhereBinaryOperator) SetParent(parent Node) {
	n.Parent = parent
}

func (n *WhereCondition) SetParent(parent Node) {
	n.Parent = parent
}

func (n *SelectField) SetParent(parent Node) {
	n.Parent = parent
}

func (n *SoqlFunction) SetParent(parent Node) {
	n.Parent = parent
}

func (n *Order) SetParent(parent Node) {
	n.Parent = parent
}

func (n *ClassDeclaration) GetLocation() *Location {
	return n.Location
}

func (n *Modifier) GetLocation() *Location {
	return n.Location
}

func (n *Annotation) GetLocation() *Location {
	return n.Location
}

func (n *InterfaceDeclaration) GetLocation() *Location {
	return n.Location
}

func (n *IntegerLiteral) GetLocation() *Location {
	return n.Location
}

func (n *Parameter) GetLocation() *Location {
	return n.Location
}

func (n *ArrayAccess) GetLocation() *Location {
	return n.Location
}

func (n *BooleanLiteral) GetLocation() *Location {
	return n.Location
}

func (n *Break) GetLocation() *Location {
	return n.Location
}

func (n *Continue) GetLocation() *Location {
	return n.Location
}

func (n *Dml) GetLocation() *Location {
	return n.Location
}

func (n *DoubleLiteral) GetLocation() *Location {
	return n.Location
}

func (n *FieldDeclaration) GetLocation() *Location {
	return n.Location
}

func (n *Try) GetLocation() *Location {
	return n.Location
}

func (n *Catch) GetLocation() *Location {
	return n.Location
}

func (n *Finally) GetLocation() *Location {
	return n.Location
}

func (n *For) GetLocation() *Location {
	return n.Location
}

func (n *ForControl) GetLocation() *Location {
	return n.Location
}

func (n *EnhancedForControl) GetLocation() *Location {
	return n.Location
}

func (n *If) GetLocation() *Location {
	return n.Location
}

func (n *MethodDeclaration) GetLocation() *Location {
	return n.Location
}

func (n *MethodInvocation) GetLocation() *Location {
	return n.Location
}

func (n *New) GetLocation() *Location {
	return n.Location
}

func (n *NullLiteral) GetLocation() *Location {
	return n.Location
}

func (n *UnaryOperator) GetLocation() *Location {
	return n.Location
}

func (n *BinaryOperator) GetLocation() *Location {
	return n.Location
}

func (n *InstanceofOperator) GetLocation() *Location {
	return n.Location
}

func (n *Return) GetLocation() *Location {
	return n.Location
}

func (n *Throw) GetLocation() *Location {
	return n.Location
}

func (n *Soql) GetLocation() *Location {
	return n.Location
}

func (n *Sosl) GetLocation() *Location {
	return n.Location
}

func (n *StringLiteral) GetLocation() *Location {
	return n.Location
}

func (n *Switch) GetLocation() *Location {
	return n.Location
}

func (n *Trigger) GetLocation() *Location {
	return n.Location
}

func (n *TriggerTiming) GetLocation() *Location {
	return n.Location
}

func (n *VariableDeclaration) GetLocation() *Location {
	return n.Location
}

func (n *VariableDeclarator) GetLocation() *Location {
	return n.Location
}

func (n *When) GetLocation() *Location {
	return n.Location
}

func (n *WhenType) GetLocation() *Location {
	return n.Location
}

func (n *While) GetLocation() *Location {
	return n.Location
}

func (n *NothingStatement) GetLocation() *Location {
	return n.Location
}

func (n *CastExpression) GetLocation() *Location {
	return n.Location
}

func (n *FieldAccess) GetLocation() *Location {
	return n.Location
}

func (n *TypeRef) GetLocation() *Location {
	return n.Location
}

func (n *Block) GetLocation() *Location {
	return n.Location
}

func (n *GetterSetter) GetLocation() *Location {
	return n.Location
}

func (n *PropertyDeclaration) GetLocation() *Location {
	return n.Location
}

func (n *ArrayInitializer) GetLocation() *Location {
	return n.Location
}

func (n *ArrayCreator) GetLocation() *Location {
	return n.Location
}

func (n *Blob) GetLocation() *Location {
	return n.Location
}

func (n *SoqlBindVariable) GetLocation() *Location {
	return n.Location
}

func (n *TernalyExpression) GetLocation() *Location {
	return n.Location
}

func (n *MapCreator) GetLocation() *Location {
	return n.Location
}

func (n *SetCreator) GetLocation() *Location {
	return n.Location
}

func (n *Name) GetLocation() *Location {
	return n.Location
}

func (n *ConstructorDeclaration) GetLocation() *Location {
	return n.Location
}

func (n *WhereBinaryOperator) GetLocation() *Location {
	return n.Location
}

func (n *WhereCondition) GetLocation() *Location {
	return n.Location
}

func (n *SelectField) GetLocation() *Location {
	return n.Location
}

func (n *SoqlFunction) GetLocation() *Location {
	return n.Location
}

func (n *Order) GetLocation() *Location {
	return n.Location
}
