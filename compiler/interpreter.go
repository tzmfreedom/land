package compiler

import (
	"github.com/k0kubun/pp"
	"github.com/tzmfreedom/goland/ast"
)

type Interpreter struct{}

func (v *Interpreter) VisitClassDeclaration(n *ast.ClassDeclaration) (interface{}, error) {
	return ast.VisitClassDeclaration(v, n)
}

func (v *Interpreter) VisitModifier(n *ast.Modifier) (interface{}, error) {
	return ast.VisitModifier(v, n)
}

func (v *Interpreter) VisitAnnotation(n *ast.Annotation) (interface{}, error) {
	return ast.VisitAnnotation(v, n)
}

func (v *Interpreter) VisitInterfaceDeclaration(n *ast.InterfaceDeclaration) (interface{}, error) {
	return ast.VisitInterfaceDeclaration(v, n)
}

func (v *Interpreter) VisitIntegerLiteral(n *ast.IntegerLiteral) (interface{}, error) {
	return &Integer{Value: n.Value}, nil
}

func (v *Interpreter) VisitParameter(n *ast.Parameter) (interface{}, error) {
	return ast.VisitParameter(v, n)
}

func (v *Interpreter) VisitArrayAccess(n *ast.ArrayAccess) (interface{}, error) {
	return ast.VisitArrayAccess(v, n)
}

func (v *Interpreter) VisitBooleanLiteral(n *ast.BooleanLiteral) (interface{}, error) {
	return &Boolean{n.Value}, nil
}

func (v *Interpreter) VisitBreak(n *ast.Break) (interface{}, error) {
	return &BreakValue{}, nil
}

func (v *Interpreter) VisitContinue(n *ast.Continue) (interface{}, error) {
	return &ContinueValue{}, nil
}

func (v *Interpreter) VisitDml(n *ast.Dml) (interface{}, error) {
	return ast.VisitDml(v, n)
}

func (v *Interpreter) VisitDoubleLiteral(n *ast.DoubleLiteral) (interface{}, error) {
	return &Double{n.Value}, nil
}

func (v *Interpreter) VisitFieldDeclaration(n *ast.FieldDeclaration) (interface{}, error) {
	return ast.VisitFieldDeclaration(v, n)
}

func (v *Interpreter) VisitTry(n *ast.Try) (interface{}, error) {
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

func (v *Interpreter) VisitCatch(n *ast.Catch) (interface{}, error) {
	return ast.VisitCatch(v, n)
}

func (v *Interpreter) VisitFinally(n *ast.Finally) (interface{}, error) {
	return ast.VisitFinally(v, n)
}

func (v *Interpreter) VisitFor(n *ast.For) (interface{}, error) {
	control := n.Control.(*ast.ForControl)
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
	return ast.VisitFor(v, n)
}

func (v *Interpreter) VisitForControl(n *ast.ForControl) (interface{}, error) {
	return ast.VisitForControl(v, n)
}

func (v *Interpreter) VisitEnhancedForControl(n *ast.EnhancedForControl) (interface{}, error) {
	return ast.VisitEnhancedForControl(v, n)
}

func (v *Interpreter) VisitIf(n *ast.If) (interface{}, error) {
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

func (v *Interpreter) VisitMethodDeclaration(n *ast.MethodDeclaration) (interface{}, error) {
	return ast.VisitMethodDeclaration(v, n)
}

func (v *Interpreter) VisitMethodInvocation(n *ast.MethodInvocation) (interface{}, error) {
	switch exp := n.NameOrExpression.(type) {
	case *ast.Name:
		if len(exp.Value) == 2 &&
			exp.Value[0] == "System" &&
			exp.Value[1] == "debug" {
			for _, p := range n.Parameters {
				res, err := p.Accept(v)
				if err != nil {
					return nil, err
				}
				pp.Println(res)
			}
		}
	case *ast.FieldAccess:
	}
	return ast.VisitMethodInvocation(v, n)
}

func (v *Interpreter) VisitNew(n *ast.New) (interface{}, error) {
	return ast.VisitNew(v, n)
}

func (v *Interpreter) VisitNullLiteral(n *ast.NullLiteral) (interface{}, error) {
	return &Null{}, nil
}

func (v *Interpreter) VisitUnaryOperator(n *ast.UnaryOperator) (interface{}, error) {
	return ast.VisitUnaryOperator(v, n)
}

func (v *Interpreter) VisitBinaryOperator(n *ast.BinaryOperator) (interface{}, error) {
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

func (v *Interpreter) VisitReturn(n *ast.Return) (interface{}, error) {
	res, err := n.Expression.Accept(v)
	if err != nil {
		return nil, err
	}
	return &ReturnValue{res}, nil
}

func (v *Interpreter) VisitThrow(n *ast.Throw) (interface{}, error) {
	res, err := n.Expression.Accept(v)
	if err != nil {
		return nil, err
	}
	return &RaiseValue{res}, nil
}

func (v *Interpreter) VisitSoql(n *ast.Soql) (interface{}, error) {
	return ast.VisitSoql(v, n)
}

func (v *Interpreter) VisitSosl(n *ast.Sosl) (interface{}, error) {
	return ast.VisitSosl(v, n)
}

func (v *Interpreter) VisitStringLiteral(n *ast.StringLiteral) (interface{}, error) {
	return &String{n.Value}, nil
}

func (v *Interpreter) VisitSwitch(n *ast.Switch) (interface{}, error) {
	return ast.VisitSwitch(v, n)
}

func (v *Interpreter) VisitTrigger(n *ast.Trigger) (interface{}, error) {
	return ast.VisitTrigger(v, n)
}

func (v *Interpreter) VisitTriggerTiming(n *ast.TriggerTiming) (interface{}, error) {
	return ast.VisitTriggerTiming(v, n)
}

func (v *Interpreter) VisitVariableDeclaration(n *ast.VariableDeclaration) (interface{}, error) {
	return ast.VisitVariableDeclaration(v, n)
}

func (v *Interpreter) VisitVariableDeclarator(n *ast.VariableDeclarator) (interface{}, error) {
	return ast.VisitVariableDeclarator(v, n)
}

func (v *Interpreter) VisitWhen(n *ast.When) (interface{}, error) {
	return ast.VisitWhen(v, n)
}

func (v *Interpreter) VisitWhenType(n *ast.WhenType) (interface{}, error) {
	return ast.VisitWhenType(v, n)
}

func (v *Interpreter) VisitWhile(n *ast.While) (interface{}, error) {
	return ast.VisitWhile(v, n)
}

func (v *Interpreter) VisitNothingStatement(n *ast.NothingStatement) (interface{}, error) {
	return ast.VisitNothingStatement(v, n)
}

func (v *Interpreter) VisitCastExpression(n *ast.CastExpression) (interface{}, error) {
	return ast.VisitCastExpression(v, n)
}

func (v *Interpreter) VisitFieldAccess(n *ast.FieldAccess) (interface{}, error) {
	return ast.VisitFieldAccess(v, n)
}

func (v *Interpreter) VisitType(n *ast.TypeRef) (interface{}, error) {
	return ast.VisitType(v, n)
}

func (v *Interpreter) VisitBlock(n *ast.Block) (interface{}, error) {
	for _, stmt := range n.Statements {
		_, err := stmt.Accept(v)
		if err != nil {
			return nil, err
		}
	}
	return nil, nil
}

func (v *Interpreter) VisitGetterSetter(n *ast.GetterSetter) (interface{}, error) {
	return ast.VisitGetterSetter(v, n)
}

func (v *Interpreter) VisitPropertyDeclaration(n *ast.PropertyDeclaration) (interface{}, error) {
	return ast.VisitPropertyDeclaration(v, n)
}

func (v *Interpreter) VisitArrayInitializer(n *ast.ArrayInitializer) (interface{}, error) {
	return ast.VisitArrayInitializer(v, n)
}

func (v *Interpreter) VisitArrayCreator(n *ast.ArrayCreator) (interface{}, error) {
	return ast.VisitArrayCreator(v, n)
}

func (v *Interpreter) VisitSoqlBindVariable(n *ast.SoqlBindVariable) (interface{}, error) {
	return ast.VisitSoqlBindVariable(v, n)
}

func (v *Interpreter) VisitTernalyExpression(n *ast.TernalyExpression) (interface{}, error) {
	res, err := n.Condition.Accept(v)
	if err != nil {
		return nil, err
	}
	if res.(*Boolean).Value {
		return n.TrueExpression.Accept(v)
	}
	return n.FalseExpression.Accept(v)
}

func (v *Interpreter) VisitMapCreator(n *ast.MapCreator) (interface{}, error) {
	return ast.VisitMapCreator(v, n)
}

func (v *Interpreter) VisitSetCreator(n *ast.SetCreator) (interface{}, error) {
	return ast.VisitSetCreator(v, n)
}

func (v *Interpreter) VisitName(n *ast.Name) (interface{}, error) {
	return ast.VisitName(v, n)
}

func (v *Interpreter) VisitConstructorDeclaration(n *ast.ConstructorDeclaration) (interface{}, error) {
	return ast.VisitConstructorDeclaration(v, n)
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

type Null struct{}

type Object struct {
	ClassType      ast.Node
	InstanceFields []ast.Node
	GenericType    string
	Parent         ast.Node
}

type ReturnValue struct {
	Value interface{}
}

type BreakValue struct{}

type ContinueValue struct{}

type RaiseValue struct {
	Value interface{}
}
