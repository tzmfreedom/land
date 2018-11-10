package ast

import (
	"fmt"
	"reflect"
	"strings"
)

type Position struct {
	FileName string
	Column   int
	Line     int
}

type ClassDeclaration struct {
	Annotations      []Node
	Modifiers        []Node
	Name             string
	SuperClass       Node
	ImplementClasses []Node
	Declarations     []Node
	InnerClasses     []Node
	Position         *Position
	Parent           Node
}

type Modifier struct {
	Name     string
	Position *Position
	Parent   Node
}

type Annotation struct {
	Name       string
	Parameters []*Parameter
	Position   *Position
	Parent     Node
}

type Interface struct {
	Annotations []Node
	Modifiers   []Node
	Name        Name
	SuperClass  []Node
	Methods     map[string][]MethodDeclaration
	Position    *Position
	Parent      Node
}

type IntegerLiteral struct {
	Value    int
	Position *Position
	Parent   Node
}

type Parameter struct {
	Modifiers []Node
	Type      Node
	Name      string
	Position  *Position
	Parent    Node
}

type ArrayAccess struct {
	Receiver Node
	Key      Node
	Position *Position
	Parent   Node
}

type BooleanLiteral struct {
	Value    bool
	Position *Position
	Parent   Node
}

type Break struct {
	Position *Position
	Parent   Node
}

type Continue struct {
	Position *Position
	Parent   Node
}

type Dml struct {
	Type       string
	Expression Node
	UpsertKey  string
	Position   *Position
	Parent     Node
}

type DoubleLiteral struct {
	Value    float64
	Position *Position
	Parent   Node
}

type FieldDeclaration struct {
	Type        Node
	Modifiers   []Node
	Declarators []Node
	Position    *Position
	Parent      Node
}

type FieldVariable struct {
	Type       Node
	Modifiers  []Node
	Expression Node
	Position   *Position
	Parent     Node
	Getter     Node
	Setter     Node
}

type Try struct {
	Block        *Block
	CatchClause  []Node
	FinallyBlock *Block
	Position     *Position
	Parent       Node
}

type Catch struct {
	Modifiers  []Node
	Type       Node
	Identifier string
	Block      *Block
	Position   *Position
	Parent     Node
}

type Finally struct {
	Block    Node
	Position *Position
	Parent   Node
}

type For struct {
	Control    Node
	Statements Node
	Position   *Position
	Parent     Node
}

type ForEnum struct {
	Type           Node
	Identifier     Node
	ListExpression Node
	Statements     []Node
	Position       *Position
	Parent         Node
}

type ForControl struct {
	ForInit    Node
	Expression Node
	ForUpdate  []Node
	Position   *Position
	Parent     Node
}

type EnhancedForControl struct {
	Modifiers            []Node
	Type                 Node
	VariableDeclaratorId string
	Expression           Node
	Position             *Position
	Parent               Node
}

type If struct {
	Condition     Node
	IfStatement   Node
	ElseStatement Node
	Position      *Position
	Parent        Node
}

type MethodDeclaration struct {
	Name           string
	Modifiers      []Node
	ReturnType     Node
	Parameters     []*Parameter
	Throws         []Node
	Statements     *Block
	NativeFunction Node
	Position       *Position
	Parent         Node
}

type MethodInvocation struct {
	NameOrExpression Node
	Parameters       []Node
	Position         *Position
	Parent           Node
}

type New struct {
	Type       Node
	Parameters []Node
	Position   *Position
	Parent     Node
}

type NullLiteral struct {
	Position *Position
	Parent   Node
}

type Object struct {
	ClassType      Node
	InstanceFields []Node
	GenericType    string
	Position       *Position
	Parent         Node
}

type UnaryOperator struct {
	Op         string
	Expression Node
	IsPrefix   bool
	Position   *Position
	Parent     Node
}

type BinaryOperator struct {
	Op       string
	Left     Node
	Right    Node
	Position *Position
	Parent   Node
}

type Return struct {
	Expression Node
	Position   *Position
	Parent     Node
}

type Throw struct {
	Expression Node
	Position   *Position
	Parent     Node
}

type Soql struct {
	SelectFields []string
	FromObject   string
	Where        []string
	Order        string
	Limit        int
	Position     *Position
	Parent       Node
}

type Sosl struct {
	Position *Position
	Parent   Node
}

type StringLiteral struct {
	Value    string
	Position *Position
	Parent   Node
}

type Switch struct {
	Expression     Node
	WhenStatements []Node
	ElseStatement  Node
	Position       *Position
	Parent         Node
}

type Trigger struct {
	Name           string
	Object         string
	TriggerTimings []*TriggerTiming
	Statements     []Node
	Position       *Position
	Parent         Node
}

type TriggerTiming struct {
	Timing   string
	Dml      string
	Position *Position
	Parent   Node
}

type VariableDeclaration struct {
	Modifiers   []Node
	Type        Node
	Declarators []Node
	Position    *Position
	Parent      Node
}

type VariableDeclarator struct {
	Name       string
	Expression Node
	Position   *Position
	Parent     Node
}

type When struct {
	Condition  []Node
	Statements *Block
	Position   *Position
	Parent     Node
}

type WhenType struct {
	Type       Node
	Identifier string
	Position   *Position
	Parent     Node
}

type While struct {
	Condition  Node
	Statements []Node
	IsDo       bool
	Position   *Position
	Parent     Node
}

// TOTO: when to use?
type NothingStatement struct {
	Position *Position
	Parent   Node
}

type CastExpression struct {
	CastType   Node
	Expression Node
	Position   *Position
	Parent     Node
}

type FieldAccess struct {
	Expression Node
	FieldName  string
	Position   *Position
	Parent     Node
}

type Type struct {
	Name       []string
	Parameters []Node
	Position   *Position
	Parent     Node
}

type Block struct {
	Statements []Node
	Position   *Position
	Parent     Node
}

type GetterSetter struct {
	Type       string
	Modifiers  []Node
	MethodBody *Block
	Position   *Position
	Parent     Node
}

type PropertyDeclaration struct {
	Modifiers     []Node
	Type          Node
	Identifier    string
	GetterSetters Node
	Position      *Position
	Parent        Node
}

type ArrayInitializer struct {
	Initializers []Node
	Position     *Position
	Parent       Node
}

type ArrayCreator struct {
	Dim              int
	Expressions      []Node
	ArrayInitializer Node
	Position         *Position
	Parent           Node
}

type Blob struct {
	Value    []byte
	Position *Position
	Parent   Node
}

type SoqlBindVariable struct {
	Expression Node
	Position   *Position
	Parent     Node
}

type TernalyExpression struct {
	Condition       Node
	TrueExpression  Node
	FalseExpression Node
	Position        *Position
	Parent          Node
}

type MapCreator struct {
	Position *Position
	Parent   Node
}

type SetCreator struct {
	Position *Position
	Parent   Node
}

type Name struct {
	Value    string
	Position *Position
	Parent   Node
}

type ConstructorDeclaration struct {
	Modifiers      []Node
	ReturnType     Node
	Parameters     []*Parameter
	Throws         []Node
	Statements     []Node
	NativeFunction Node
	Position       *Position
	Parent         Node
}

type InterfaceDeclaration struct {
	Modifiers []Node
}

type Visitor interface {
	VisitClassDeclaration(*ClassDeclaration) interface{}
	VisitModifier(*Modifier) interface{}
	VisitAnnotation(*Annotation) interface{}
	VisitInterface(*Interface) interface{}
	VisitIntegerLiteral(*IntegerLiteral) interface{}
	VisitParameter(*Parameter) interface{}
	VisitArrayAccess(*ArrayAccess) interface{}
	VisitBooleanLiteral(*BooleanLiteral) interface{}
	VisitBreak(*Break) interface{}
	VisitContinue(*Continue) interface{}
	VisitDml(*Dml) interface{}
	VisitDoubleLiteral(*DoubleLiteral) interface{}
	VisitFieldDeclaration(*FieldDeclaration) interface{}
	VisitFieldVariable(*FieldVariable) interface{}
	VisitTry(*Try) interface{}
	VisitCatch(*Catch) interface{}
	VisitFinally(*Finally) interface{}
	VisitFor(*For) interface{}
	VisitForEnum(*ForEnum) interface{}
	VisitForControl(*ForControl) interface{}
	VisitEnhancedForControl(*EnhancedForControl) interface{}
	VisitIf(*If) interface{}
	VisitMethodDeclaration(*MethodDeclaration) interface{}
	VisitMethodInvocation(*MethodInvocation) interface{}
	VisitNew(*New) interface{}
	VisitNullLiteral(*NullLiteral) interface{}
	VisitObject(*Object) interface{}
	VisitUnaryOperator(*UnaryOperator) interface{}
	VisitBinaryOperator(*BinaryOperator) interface{}
	VisitReturn(*Return) interface{}
	VisitThrow(*Throw) interface{}
	VisitSoql(*Soql) interface{}
	VisitSosl(*Sosl) interface{}
	VisitStringLiteral(*StringLiteral) interface{}
	VisitSwitch(*Switch) interface{}
	VisitTrigger(*Trigger) interface{}
	VisitTriggerTiming(*TriggerTiming) interface{}
	VisitVariableDeclaration(*VariableDeclaration) interface{}
	VisitVariableDeclarator(*VariableDeclarator) interface{}
	VisitWhen(*When) interface{}
	VisitWhenType(*WhenType) interface{}
	VisitWhile(*While) interface{}
	VisitNothingStatement(*NothingStatement) interface{}
	VisitCastExpression(*CastExpression) interface{}
	VisitFieldAccess(*FieldAccess) interface{}
	VisitType(*Type) interface{}
	VisitBlock(*Block) interface{}
	VisitGetterSetter(*GetterSetter) interface{}
	VisitPropertyDeclaration(*PropertyDeclaration) interface{}
	VisitArrayInitializer(*ArrayInitializer) interface{}
	VisitArrayCreator(*ArrayCreator) interface{}
	VisitBlob(*Blob) interface{}
	VisitSoqlBindVariable(*SoqlBindVariable) interface{}
	VisitTernalyExpression(*TernalyExpression) interface{}
	VisitMapCreator(*MapCreator) interface{}
	VisitSetCreator(*SetCreator) interface{}
	VisitName(*Name) interface{}
	VisitConstructorDeclaration(*ConstructorDeclaration) interface{}
}

var VoidType = &Type{}

type Node interface {
	Accept(Visitor) interface{}
	GetChildren() []interface{}
	GetType() string
	GetParent() Node
	SetParent(Node)
}

func (n *ClassDeclaration) Accept(v Visitor) interface{} {
	return v.VisitClassDeclaration(n)
}

func (n *ClassDeclaration) GetChildren() []interface{} {
	return []interface{}{
		n.Name,
		n.ImplementClasses,
		n.SuperClass,
		n.Annotations,
		n.Declarations,
		n.InnerClasses,
	}
}

func (n *Modifier) Accept(v Visitor) interface{} {
	return v.VisitModifier(n)
}

func (n *Modifier) GetChildren() []interface{} {
	return []interface{}{
		n.Name,
	}
}

func (n *Annotation) Accept(v Visitor) interface{} {
	return v.VisitAnnotation(n)
}

func (n *Annotation) GetChildren() []interface{} {
	return []interface{}{
		n.Name,
	}
}

func (n *Interface) Accept(v Visitor) interface{} {
	return v.VisitInterface(n)
}

func (n *Interface) GetChildren() []interface{} {
	return []interface{}{
		n.Name,
		n.Annotations,
		n.Methods,
		n.Modifiers,
		n.SuperClass,
	}
}

func (n *IntegerLiteral) Accept(v Visitor) interface{} {
	return v.VisitIntegerLiteral(n)
}

func (n *IntegerLiteral) GetChildren() []interface{} {
	return []interface{}{
		n.Value,
	}
}

func (n *Parameter) Accept(v Visitor) interface{} {
	return v.VisitParameter(n)
}

func (n *Parameter) GetChildren() []interface{} {
	return []interface{}{
		n.Type,
		n.Name,
		n.Modifiers,
	}
}

func (n *ArrayAccess) Accept(v Visitor) interface{} {
	return v.VisitArrayAccess(n)
}

func (n *ArrayAccess) GetChildren() []interface{} {
	return []interface{}{
		n.Receiver,
		n.Key,
	}
}

func (n *BooleanLiteral) Accept(v Visitor) interface{} {
	return v.VisitBooleanLiteral(n)
}

func (n *BooleanLiteral) GetChildren() []interface{} {
	return nil
}

func (n *Break) Accept(v Visitor) interface{} {
	return v.VisitBreak(n)
}

func (n *Break) GetChildren() []interface{} {
	return nil
}

func (n *Continue) Accept(v Visitor) interface{} {
	return v.VisitContinue(n)
}

func (n *Continue) GetChildren() []interface{} {
	return nil
}

func (n *Dml) Accept(v Visitor) interface{} {
	return v.VisitDml(n)
}

func (n *Dml) GetChildren() []interface{} {
	return []interface{}{
		n.Type,
		n.Expression,
	}
}

func (n *DoubleLiteral) Accept(v Visitor) interface{} {
	return v.VisitDoubleLiteral(n)
}

func (n *DoubleLiteral) GetChildren() []interface{} {
	return []interface{}{
		n.Value,
	}
}

func (n *FieldDeclaration) Accept(v Visitor) interface{} {
	return v.VisitFieldDeclaration(n)
}

func (n *FieldDeclaration) GetChildren() []interface{} {
	return []interface{}{
		n.Type,
		n.Declarators,
	}
}

func (n *FieldVariable) Accept(v Visitor) interface{} {
	return v.VisitFieldVariable(n)
}

func (n *FieldVariable) GetChildren() []interface{} {
	return []interface{}{
		n.Modifiers,
		n.Type,
		n.Expression,
		n.Getter,
		n.Setter,
	}
}

func (n *Try) Accept(v Visitor) interface{} {
	return v.VisitTry(n)
}

func (n *Try) GetChildren() []interface{} {
	return []interface{}{
		n.Block,
		n.CatchClause,
		n.FinallyBlock,
	}
}

func (n *Catch) Accept(v Visitor) interface{} {
	return v.VisitCatch(n)
}

func (n *Catch) GetChildren() []interface{} {
	return []interface{}{
		n.Type,
		n.Identifier,
		n.Modifiers,
		n.Block,
	}
}

func (n *Finally) Accept(v Visitor) interface{} {
	return v.VisitFinally(n)
}

func (n *Finally) GetChildren() []interface{} {
	return []interface{}{
		n.Block,
	}
}

func (n *For) Accept(v Visitor) interface{} {
	return v.VisitFor(n)
}

func (n *For) GetChildren() []interface{} {
	return []interface{}{
		n.Statements,
		n.Control,
	}
}

func (n *ForEnum) Accept(v Visitor) interface{} {
	return v.VisitForEnum(n)
}

func (n *ForEnum) GetChildren() []interface{} {
	return []interface{}{
		n.Type,
		n.Identifier,
		n.ListExpression,
		n.Statements,
	}
}

func (n *ForControl) Accept(v Visitor) interface{} {
	return v.VisitForControl(n)
}

func (n *ForControl) GetChildren() []interface{} {
	return []interface{}{
		n.ForInit,
		n.Expression,
		n.ForUpdate,
	}
}

func (n *EnhancedForControl) Accept(v Visitor) interface{} {
	return v.VisitEnhancedForControl(n)
}

func (n *EnhancedForControl) GetChildren() []interface{} {
	return []interface{}{
		n.Type,
		n.VariableDeclaratorId,
		n.Expression,
	}
}

func (n *If) Accept(v Visitor) interface{} {
	return v.VisitIf(n)
}

func (n *If) GetChildren() []interface{} {
	return []interface{}{
		n.IfStatement,
		n.Condition,
	}
}

func (n *MethodDeclaration) Accept(v Visitor) interface{} {
	return v.VisitMethodDeclaration(n)
}

func (n *MethodDeclaration) GetChildren() []interface{} {
	return []interface{}{
		n.Name,
		n.Modifiers,
		n.ReturnType,
		n.Throws,
		n.Parameters,
		n.NativeFunction,
		n.Statements,
	}
}

func (n *MethodInvocation) Accept(v Visitor) interface{} {
	return v.VisitMethodInvocation(n)
}

func (n *MethodInvocation) GetChildren() []interface{} {
	return []interface{}{
		n.NameOrExpression,
		n.Parameters,
	}
}

func (n *New) Accept(v Visitor) interface{} {
	return v.VisitNew(n)
}

func (n *New) GetChildren() []interface{} {
	return []interface{}{
		n.Type,
		n.Parameters,
	}
}

func (n *NullLiteral) Accept(v Visitor) interface{} {
	return v.VisitNullLiteral(n)
}

func (n *NullLiteral) GetChildren() []interface{} {
	return nil
}

func (n *Object) Accept(v Visitor) interface{} {
	return v.VisitObject(n)
}

func (n *Object) GetChildren() []interface{} {
	return []interface{}{
		n.ClassType,
		n.GenericType,
		n.InstanceFields,
	}
}

func (n *UnaryOperator) Accept(v Visitor) interface{} {
	return v.VisitUnaryOperator(n)
}

func (n *UnaryOperator) GetChildren() []interface{} {
	return []interface{}{
		n.Op,
		n.Expression,
		n.IsPrefix,
	}
}

func (n *BinaryOperator) Accept(v Visitor) interface{} {
	return v.VisitBinaryOperator(n)
}

func (n *BinaryOperator) GetChildren() []interface{} {
	return []interface{}{
		n.Op,
		n.Left,
		n.Right,
	}
}

func (n *Return) Accept(v Visitor) interface{} {
	return v.VisitReturn(n)
}

func (n *Return) GetChildren() []interface{} {
	return []interface{}{
		n.Expression,
	}
}

func (n *Throw) Accept(v Visitor) interface{} {
	return v.VisitThrow(n)
}

func (n *Throw) GetChildren() []interface{} {
	return []interface{}{
		n.Expression,
	}
}

func (n *Soql) Accept(v Visitor) interface{} {
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

func (n *Sosl) Accept(v Visitor) interface{} {
	return v.VisitSosl(n)
}

func (n *Sosl) GetChildren() []interface{} {
	return []interface{}{}
}

func (n *StringLiteral) Accept(v Visitor) interface{} {
	return v.VisitStringLiteral(n)
}

func (n *StringLiteral) GetChildren() []interface{} {
	return []interface{}{
		n.Value,
	}
}

func (n *Switch) Accept(v Visitor) interface{} {
	return v.VisitSwitch(n)
}

func (n *Switch) GetChildren() []interface{} {
	return []interface{}{
		n.Expression,
		n.WhenStatements,
		n.ElseStatement,
	}
}

func (n *Trigger) Accept(v Visitor) interface{} {
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

func (n *TriggerTiming) Accept(v Visitor) interface{} {
	return v.VisitTriggerTiming(n)
}

func (n *TriggerTiming) GetChildren() []interface{} {
	return []interface{}{
		n.Dml,
		n.Timing,
	}
}

func (n *VariableDeclaration) Accept(v Visitor) interface{} {
	return v.VisitVariableDeclaration(n)
}

func (n *VariableDeclaration) GetChildren() []interface{} {
	return []interface{}{
		n.Type,
		n.Modifiers,
		n.Declarators,
	}
}

func (n *VariableDeclarator) Accept(v Visitor) interface{} {
	return v.VisitVariableDeclarator(n)
}

func (n *VariableDeclarator) GetChildren() []interface{} {
	return []interface{}{
		n.Name,
		n.Expression,
	}
}

func (n *When) Accept(v Visitor) interface{} {
	return v.VisitWhen(n)
}

func (n *When) GetChildren() []interface{} {
	return []interface{}{
		n.Condition,
		n.Statements,
	}
}

func (n *WhenType) Accept(v Visitor) interface{} {
	return v.VisitWhenType(n)
}

func (n *WhenType) GetChildren() []interface{} {
	return []interface{}{
		n.Type,
		n.Identifier,
	}
}

func (n *While) Accept(v Visitor) interface{} {
	return v.VisitWhile(n)
}

func (n *While) GetChildren() []interface{} {
	return []interface{}{
		n.Condition,
		n.Statements,
		n.IsDo,
	}
}

func (n *NothingStatement) Accept(v Visitor) interface{} {
	return v.VisitNothingStatement(n)
}

func (n *NothingStatement) GetChildren() []interface{} {
	return nil
}

func (n *CastExpression) Accept(v Visitor) interface{} {
	return v.VisitCastExpression(n)
}

func (n *CastExpression) GetChildren() []interface{} {
	return []interface{}{
		n.CastType,
		n.Expression,
	}
}

func (n *FieldAccess) Accept(v Visitor) interface{} {
	return v.VisitFieldAccess(n)
}

func (n *FieldAccess) GetChildren() []interface{} {
	return []interface{}{
		n.Expression,
		n.FieldName,
	}
}

func (n *Type) Accept(v Visitor) interface{} {
	return v.VisitType(n)
}

func (n *Type) GetChildren() []interface{} {
	return []interface{}{
		n.Name,
		n.Parameters,
	}
}

func (n *Block) Accept(v Visitor) interface{} {
	return v.VisitBlock(n)
}

func (n *Block) GetChildren() []interface{} {
	return []interface{}{
		n.Statements,
	}
}

func (n *GetterSetter) Accept(v Visitor) interface{} {
	return v.VisitGetterSetter(n)
}

func (n *GetterSetter) GetChildren() []interface{} {
	return []interface{}{
		n.Type,
		n.Modifiers,
		n.MethodBody,
	}
}

func (n *PropertyDeclaration) Accept(v Visitor) interface{} {
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

func (n *ArrayInitializer) Accept(v Visitor) interface{} {
	return v.VisitArrayInitializer(n)
}

func (n *ArrayInitializer) GetChildren() []interface{} {
	return []interface{}{
		n.Initializers,
	}
}

func (n *ArrayCreator) Accept(v Visitor) interface{} {
	return v.VisitArrayCreator(n)
}

func (n *ArrayCreator) GetChildren() []interface{} {
	return []interface{}{
		n.Dim,
		n.ArrayInitializer,
		n.Expressions,
	}
}

func (n *Blob) Accept(v Visitor) interface{} {
	return v.VisitBlob(n)
}

func (n *Blob) GetChildren() []interface{} {
	return []interface{}{
		n.Value,
	}
}

func (n *SoqlBindVariable) Accept(v Visitor) interface{} {
	return v.VisitSoqlBindVariable(n)
}

func (n *SoqlBindVariable) GetChildren() []interface{} {
	return []interface{}{
		n.Expression,
	}
}

func (n *TernalyExpression) Accept(v Visitor) interface{} {
	return v.VisitTernalyExpression(n)
}

func (n *TernalyExpression) GetChildren() []interface{} {
	return []interface{}{
		n.Condition,
		n.TrueExpression,
		n.FalseExpression,
	}
}

func (n *MapCreator) Accept(v Visitor) interface{} {
	return v.VisitMapCreator(n)
}

func (n *MapCreator) GetChildren() []interface{} {
	return []interface{}{}
}

func (n *SetCreator) Accept(v Visitor) interface{} {
	return v.VisitSetCreator(n)
}

func (n *SetCreator) GetChildren() []interface{} {
	return []interface{}{}
}

func (n *Name) Accept(v Visitor) interface{} {
	return v.VisitName(n)
}

func (n *Name) GetChildren() []interface{} {
	return []interface{}{
		n.Value,
	}
}

func (n *ConstructorDeclaration) Accept(v Visitor) interface{} {
	return v.VisitConstructorDeclaration(n)
}

func (n *ConstructorDeclaration) GetChildren() []interface{} {
	return []interface{}{
		n.Parameters,
		n.ReturnType,
		n.Modifiers,
		n.NativeFunction,
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
func (n *Interface) GetType() string {
	return "Interface"
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
func (n *FieldVariable) GetType() string {
	return "FieldVariable"
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
func (n *ForEnum) GetType() string {
	return "ForEnum"
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
func (n *Object) GetType() string {
	return "Object"
}
func (n *UnaryOperator) GetType() string {
	return "UnaryOperator"
}
func (n *BinaryOperator) GetType() string {
	return "BinaryOperator"
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
func (n *Type) GetType() string {
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
func (n *Blob) GetType() string {
	return "Blob"
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

func (n *ClassDeclaration) GetParent() Node {
	return n.Parent
}
func (n *Modifier) GetParent() Node {
	return n.Parent
}
func (n *Annotation) GetParent() Node {
	return n.Parent
}
func (n *Interface) GetParent() Node {
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
func (n *FieldVariable) GetParent() Node {
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
func (n *ForEnum) GetParent() Node {
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
func (n *Object) GetParent() Node {
	return n.Parent
}
func (n *UnaryOperator) GetParent() Node {
	return n.Parent
}
func (n *BinaryOperator) GetParent() Node {
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
func (n *Type) GetParent() Node {
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
func (n *Blob) GetParent() Node {
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

func (n *ClassDeclaration) SetParent(parent Node) {
	n.Parent = parent
}

func (n *Modifier) SetParent(parent Node) {
	n.Parent = parent
}

func (n *Annotation) SetParent(parent Node) {
	n.Parent = parent
}

func (n *Interface) SetParent(parent Node) {
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

func (n *FieldVariable) SetParent(parent Node) {
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

func (n *ForEnum) SetParent(parent Node) {
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

func (n *Object) SetParent(parent Node) {
	n.Parent = parent
}

func (n *UnaryOperator) SetParent(parent Node) {
	n.Parent = parent
}

func (n *BinaryOperator) SetParent(parent Node) {
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

func (n *Type) SetParent(parent Node) {
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

func (n *Blob) SetParent(parent Node) {
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

func Dump(n Node, ident int) string {
	if n == nil || reflect.ValueOf(n).IsNil() {
		return "nil"
	}
	children := n.GetChildren()
	if len(children) != 0 {
		properties := make([]string, len(children))
		for i, child := range children {
			if nodes, ok := child.([]Node); ok {
				properties[i] = DumpArray(nodes, ident+2)
			} else if node, ok := child.(Node); ok {
				properties[i] = Dump(node, ident+2)
			} else {
				properties[i] = strings.Repeat(" ", ident) +
					fmt.Sprintf("%v", child)
			}
		}
		return strings.Repeat(" ", ident) +
			"(" +
			n.GetType() + "\n" +
			strings.Repeat(" ", ident+2) +
			strings.Join(properties, "\n") +
			")"
	}
	return strings.Repeat(" ", ident) + "(" + n.GetType() + ")"
}

func DumpArray(nodes []Node, ident int) string {
	properties := make([]string, len(nodes))
	for i, n := range nodes {
		properties[i] = Dump(n, 0)
	}
	return strings.Repeat(" ", ident) +
		"[" + strings.Join(properties, ",") + "]"
}
