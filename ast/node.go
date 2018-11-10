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
}

type Modifier struct {
	Name     string
	Position *Position
}

type Annotation struct {
	Name       string
	Parameters []*Parameter
	Position   *Position
}

type Interface struct {
	Annotations []Node
	Modifiers   []Node
	Name        Name
	SuperClass  []Node
	Methods     map[string][]MethodDeclaration
	Position    *Position
}

type IntegerLiteral struct {
	Value    int
	Position *Position
}

type Parameter struct {
	Modifiers []Node
	Type      Node
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
	Type       string
	Expression Node
	UpsertKey  string
	Position   *Position
}

type DoubleLiteral struct {
	Value    float64
	Position *Position
}

type FieldDeclaration struct {
	Type        Node
	Modifiers   []Node
	Declarators []Node
	Position    *Position
}

type FieldVariable struct {
	Type       Node
	Modifiers  []Node
	Expression Node
	Position   *Position
	Getter     Node
	Setter     Node
}

type Try struct {
	Block        *Block
	CatchClause  []Node
	FinallyBlock *Block
	Position     *Position
}

type Catch struct {
	Modifiers  []Node
	Type       Node
	Identifier string
	Block      *Block
	Position   *Position
}

type Finally struct {
	Block    Node
	Position *Position
}

type For struct {
	Control    Node
	Statements Node
	Position   *Position
}

type ForEnum struct {
	Type           Node
	Identifier     Node
	ListExpression Node
	Statements     []Node
	Position       *Position
}

type ForControl struct {
	ForInit    Node
	Expression Node
	ForUpdate  []Node
	Position   *Position
}

type EnhancedForControl struct {
	Modifiers            []Node
	Type                 Node
	VariableDeclaratorId string
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
	Modifiers      []Node
	ReturnType     Node
	Parameters     []*Parameter
	Throws         []Node
	Statements     *Block
	NativeFunction Node
	Position       *Position
}

type MethodInvocation struct {
	NameOrExpression Node
	Parameters       []Node
	Position         *Position
}

type New struct {
	Type       Node
	Parameters []Node
	Position   *Position
}

type NullLiteral struct {
	Position *Position
}

type Object struct {
	ClassType      Node
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
	TriggerTimings []*TriggerTiming
	Statements     []Node
	Position       *Position
}

type TriggerTiming struct {
	Timing   string
	Dml      string
	Position *Position
}

type VariableDeclaration struct {
	Modifiers   []Node
	Type        Node
	Declarators []Node
	Position    *Position
}

type VariableDeclarator struct {
	Name       string
	Expression Node
	Position   *Position
}

type When struct {
	Condition  []Node
	Statements *Block
	Position   *Position
}

type WhenType struct {
	Type       Node
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
type NothingStatement struct {
	Position *Position
}

type CastExpression struct {
	CastType   Node
	Expression Node
	Position   *Position
}

type FieldAccess struct {
	Expression Node
	FieldName  string
	Position   *Position
}

type Type struct {
	Name       []string
	Parameters []Node
	Position   *Position
}

type Block struct {
	Statements []Node
	Position   *Position
}

type GetterSetter struct {
	Type       string
	Modifiers  []Node
	MethodBody *Block
	Position   *Position
}

type PropertyDeclaration struct {
	Modifiers     []Node
	Type          Node
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
	Value    string
	Position *Position
}

type ConstructorDeclaration struct {
	Modifiers      []Node
	ReturnType     Node
	Parameters     []*Parameter
	Throws         []Node
	Statements     []Node
	NativeFunction Node
	Position       *Position
}

type InterfaceDeclaration struct {
	Modifiers []Node
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

var VoidType = &Type{}

type Node interface {
	Accept(Visitor) interface{}
	GetChildren() []interface{}
	GetType() string
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
