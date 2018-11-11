package ast

import "github.com/k0kubun/pp"

type Interpreter struct {
	Env []interface{}
}

func (v *Interpreter) VisitClassDeclaration(n *ClassDeclaration) (interface{}, error) {
	return VisitClassDeclaration(v, n)
}

func (v *Interpreter) VisitModifier(n *Modifier) (interface{}, error) {
	return VisitModifier(v, n)
}

func (v *Interpreter) VisitAnnotation(n *Annotation) (interface{}, error) {
	return VisitAnnotation(v, n)
}

func (v *Interpreter) VisitInterfaceDeclaration(n *InterfaceDeclaration) (interface{}, error) {
	return VisitInterfaceDeclaration(v, n)
}

func (v *Interpreter) VisitIntegerLiteral(n *IntegerLiteral) (interface{}, error) {
	return &Integer{Value: n.Value}, nil
}

func (v *Interpreter) VisitParameter(n *Parameter) (interface{}, error) {
	return VisitParameter(v, n)
}

func (v *Interpreter) VisitArrayAccess(n *ArrayAccess) (interface{}, error) {
	return VisitArrayAccess(v, n)
}

func (v *Interpreter) VisitBooleanLiteral(n *BooleanLiteral) (interface{}, error) {
	return &Boolean{n.Value}, nil
}

func (v *Interpreter) VisitBreak(n *Break) (interface{}, error) {
	return &BreakValue{}, nil
}

func (v *Interpreter) VisitContinue(n *Continue) (interface{}, error) {
	return &ContinueValue{}, nil
}

func (v *Interpreter) VisitDml(n *Dml) (interface{}, error) {
	return VisitDml(v, n)
}

func (v *Interpreter) VisitDoubleLiteral(n *DoubleLiteral) (interface{}, error) {
	return &Double{n.Value}, nil
}

func (v *Interpreter) VisitFieldDeclaration(n *FieldDeclaration) (interface{}, error) {
	return VisitFieldDeclaration(v, n)
}

func (v *Interpreter) VisitTry(n *Try) (interface{}, error) {
	res, err := n.Block.Accept(v)
	if err != nil {
		return nil, err
	}
	switch res.(type) {
	case *RaiseValue:
		_ = res.(*RaiseValue)
		// TODO: implement
	default:
		res, err = n.FinallyBlock.Accept(v)
		return res, nil
	}
	if _, err = n.FinallyBlock.Accept(v); err != nil {
		return nil, err
	}
	return nil, nil
}

func (v *Interpreter) VisitCatch(n *Catch) (interface{}, error) {
	return VisitCatch(v, n)
}

func (v *Interpreter) VisitFinally(n *Finally) (interface{}, error) {
	return VisitFinally(v, n)
}

func (v *Interpreter) VisitFor(n *For) (interface{}, error) {
	control := n.Control.(*ForControl)
	_, err := control.ForInit.Accept(v)
	if err != nil {
		return nil, err
	}
	for {
		res, err := control.Expression.Accept(v)
		if err != nil {
			return nil, err
		}
		if res.(*Boolean).Value {
			res, err = n.Statements.Accept(v)
			for _, stmt := range control.ForUpdate {
				stmt.Accept(v)
			}
		} else {
			break
		}
	}
	return VisitFor(v, n)
}

func (v *Interpreter) VisitForControl(n *ForControl) (interface{}, error) {
	return VisitForControl(v, n)
}

func (v *Interpreter) VisitEnhancedForControl(n *EnhancedForControl) (interface{}, error) {
	return VisitEnhancedForControl(v, n)
}

func (v *Interpreter) VisitIf(n *If) (interface{}, error) {
	res, err := n.Condition.Accept(v)
	if err != nil {
		return nil, err
	}
	if res.(*Boolean).Value {
		n.IfStatement.Accept(v)
	} else {
		n.ElseStatement.Accept(v)
	}
	return nil, nil
}

func (v *Interpreter) VisitMethodDeclaration(n *MethodDeclaration) (interface{}, error) {
	return VisitMethodDeclaration(v, n)
}

func (v *Interpreter) VisitMethodInvocation(n *MethodInvocation) (interface{}, error) {
	switch exp := n.NameOrExpression.(type) {
	case *FieldAccess:
		if exp.Expression.(*Name).Value == "System" &&
			exp.FieldName == "debug" {
			for _, p := range n.Parameters {
				res, err := p.Accept(v)
				if err != nil {
					return nil, err
				}
				pp.Println(res)
			}
		}
	}
	return VisitMethodInvocation(v, n)
}

func (v *Interpreter) VisitNew(n *New) (interface{}, error) {
	return VisitNew(v, n)
}

func (v *Interpreter) VisitNullLiteral(n *NullLiteral) (interface{}, error) {
	return &Null{}, nil
}

func (v *Interpreter) VisitUnaryOperator(n *UnaryOperator) (interface{}, error) {
	return VisitUnaryOperator(v, n)
}

func (v *Interpreter) VisitBinaryOperator(n *BinaryOperator) (interface{}, error) {
	left, err := n.Left.Accept(v)
	if err != nil {
		return nil, err
	}
	right, err := n.Right.Accept(v)
	if err != nil {
		return nil, err
	}
	switch left.(type) {
	case *Integer:
		switch n.Op {
		case "+":
			l := left.(*Integer).Value
			r := right.(*Integer).Value
			return &Integer{Value: l + r}, nil
		case "-":
			l := left.(*Integer).Value
			r := right.(*Integer).Value
			return &Integer{Value: l - r}, nil
		case "*":
			l := left.(*Integer).Value
			r := right.(*Integer).Value
			return &Integer{Value: l * r}, nil
		case "/":
			l := left.(*Integer).Value
			r := right.(*Integer).Value
			return &Integer{Value: l / r}, nil
		case "%":
			l := left.(*Integer).Value
			r := right.(*Integer).Value
			return &Integer{Value: l % r}, nil
		case "<":
			l := left.(*Integer).Value
			r := right.(*Integer).Value
			return &Boolean{Value: l < r}, nil
		case ">":
			l := left.(*Integer).Value
			r := right.(*Integer).Value
			return &Boolean{Value: l > r}, nil
		case "<=":
			l := left.(*Integer).Value
			r := right.(*Integer).Value
			return &Boolean{Value: l <= r}, nil
		case ">=":
			l := left.(*Integer).Value
			r := right.(*Integer).Value
			return &Boolean{Value: l >= r}, nil
		case "==":
			l := left.(*Integer).Value
			r := right.(*Integer).Value
			return &Boolean{Value: l == r}, nil
		case "!=":
			l := left.(*Integer).Value
			r := right.(*Integer).Value
			return &Boolean{Value: l != r}, nil
		}
	case *Double:
		switch n.Op {
		case "+":
			l := left.(*Double).Value
			r := right.(*Double).Value
			return &Double{Value: l + r}, nil
		case "-":
			l := left.(*Double).Value
			r := right.(*Double).Value
			return &Double{Value: l - r}, nil
		case "*":
			l := left.(*Double).Value
			r := right.(*Double).Value
			return &Double{Value: l * r}, nil
		case "/":
			l := left.(*Double).Value
			r := right.(*Double).Value
			return &Double{Value: l / r}, nil
		}
	case *String:
		switch n.Op {
		case "+":
			l := left.(*String).Value
			r := right.(*String).Value
			return &String{Value: l + r}, nil
		}
	}
	return nil, nil
}

func (v *Interpreter) VisitReturn(n *Return) (interface{}, error) {
	res, err := n.Expression.Accept(v)
	if err != nil {
		return nil, err
	}
	return &ReturnValue{res}, nil
}

func (v *Interpreter) VisitThrow(n *Throw) (interface{}, error) {
	res, err := n.Expression.Accept(v)
	if err != nil {
		return nil, err
	}
	return &RaiseValue{res}, nil
}

func (v *Interpreter) VisitSoql(n *Soql) (interface{}, error) {
	return VisitSoql(v, n)
}

func (v *Interpreter) VisitSosl(n *Sosl) (interface{}, error) {
	return VisitSosl(v, n)
}

func (v *Interpreter) VisitStringLiteral(n *StringLiteral) (interface{}, error) {
	return &String{n.Value}, nil
}

func (v *Interpreter) VisitSwitch(n *Switch) (interface{}, error) {
	return VisitSwitch(v, n)
}

func (v *Interpreter) VisitTrigger(n *Trigger) (interface{}, error) {
	return VisitTrigger(v, n)
}

func (v *Interpreter) VisitTriggerTiming(n *TriggerTiming) (interface{}, error) {
	return VisitTriggerTiming(v, n)
}

func (v *Interpreter) VisitVariableDeclaration(n *VariableDeclaration) (interface{}, error) {
	return VisitVariableDeclaration(v, n)
}

func (v *Interpreter) VisitVariableDeclarator(n *VariableDeclarator) (interface{}, error) {
	return VisitVariableDeclarator(v, n)
}

func (v *Interpreter) VisitWhen(n *When) (interface{}, error) {
	return VisitWhen(v, n)
}

func (v *Interpreter) VisitWhenType(n *WhenType) (interface{}, error) {
	return VisitWhenType(v, n)
}

func (v *Interpreter) VisitWhile(n *While) (interface{}, error) {
	return VisitWhile(v, n)
}

func (v *Interpreter) VisitNothingStatement(n *NothingStatement) (interface{}, error) {
	return VisitNothingStatement(v, n)
}

func (v *Interpreter) VisitCastExpression(n *CastExpression) (interface{}, error) {
	return VisitCastExpression(v, n)
}

func (v *Interpreter) VisitFieldAccess(n *FieldAccess) (interface{}, error) {
	return VisitFieldAccess(v, n)
}

func (v *Interpreter) VisitType(n *Type) (interface{}, error) {
	return VisitType(v, n)
}

func (v *Interpreter) VisitBlock(n *Block) (interface{}, error) {
	for _, stmt := range n.Statements {
		_, err := stmt.Accept(v)
		if err != nil {
			return nil, err
		}
	}
	return nil, nil
}

func (v *Interpreter) VisitGetterSetter(n *GetterSetter) (interface{}, error) {
	return VisitGetterSetter(v, n)
}

func (v *Interpreter) VisitPropertyDeclaration(n *PropertyDeclaration) (interface{}, error) {
	return VisitPropertyDeclaration(v, n)
}

func (v *Interpreter) VisitArrayInitializer(n *ArrayInitializer) (interface{}, error) {
	return VisitArrayInitializer(v, n)
}

func (v *Interpreter) VisitArrayCreator(n *ArrayCreator) (interface{}, error) {
	return VisitArrayCreator(v, n)
}

func (v *Interpreter) VisitSoqlBindVariable(n *SoqlBindVariable) (interface{}, error) {
	return VisitSoqlBindVariable(v, n)
}

func (v *Interpreter) VisitTernalyExpression(n *TernalyExpression) (interface{}, error) {
	res, err := n.Condition.Accept(v)
	if err != nil {
		return nil, err
	}
	if res.(*Boolean).Value {
		return n.TrueExpression.Accept(v)
	}
	return n.FalseExpression.Accept(v)
}

func (v *Interpreter) VisitMapCreator(n *MapCreator) (interface{}, error) {
	return VisitMapCreator(v, n)
}

func (v *Interpreter) VisitSetCreator(n *SetCreator) (interface{}, error) {
	return VisitSetCreator(v, n)
}

func (v *Interpreter) VisitName(n *Name) (interface{}, error) {
	return VisitName(v, n)
}

func (v *Interpreter) VisitConstructorDeclaration(n *ConstructorDeclaration) (interface{}, error) {
	return VisitConstructorDeclaration(v, n)
}

type Integer struct {
	Value int
}

type String struct {
	Value string
}

type Double struct {
	Value float64
}

type Boolean struct {
	Value bool
}

type ReturnValue struct {
	Value interface{}
}

type BreakValue struct{}

type ContinueValue struct{}

type RaiseValue struct {
	Value interface{}
}

type Null struct{}

type Object struct {
	ClassType      Node
	InstanceFields []Node
	GenericType    string
	Parent         Node
}

type Value interface {
	Inspect()
}
