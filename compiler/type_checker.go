package compiler

import (
	"github.com/tzmfreedom/goland/ast"
)

type TypeChecker struct {
	Context *Context
	Errors  []error
}

func (v *TypeChecker) VisitClassType(n *ast.ClassType) (interface{}, error) {
	v.Context.CurrentClass = n
	for _, f := range n.StaticFields {
		_, err := f.Accept(v)
		v.Errors = append(v.Errors, err)
	}

	for _, f := range n.InstanceFields {
		_, err := f.Accept(v)
		v.Errors = append(v.Errors, err)
	}

	for _, m := range n.StaticMethods {
		_, err := m.Accept(v)
		v.Errors = append(v.Errors, err)
	}

	for _, m := range n.InstanceMethods {
		_, err := m.Accept(v)
		v.Errors = append(v.Errors, err)
	}
	v.Context.CurrentClass = nil
	return nil, nil
}

func (v *TypeChecker) VisitClassDeclaration(n *ast.ClassDeclaration) (interface{}, error) {
	return ast.VisitClassDeclaration(v, n)
}

func (v *TypeChecker) VisitModifier(n *ast.Modifier) (interface{}, error) {
	panic("Not pass")
	return ast.VisitModifier(v, n)
}

func (v *TypeChecker) VisitAnnotation(n *ast.Annotation) (interface{}, error) {
	panic("Not pass")
	return ast.VisitAnnotation(v, n)
}

func (v *TypeChecker) VisitInterfaceDeclaration(n *ast.InterfaceDeclaration) (interface{}, error) {
	return ast.VisitInterfaceDeclaration(v, n)
}

func (v *TypeChecker) VisitIntegerLiteral(n *ast.IntegerLiteral) (interface{}, error) {
	return ast.IntegerType, nil
}

func (v *TypeChecker) VisitParameter(n *ast.Parameter) (interface{}, error) {
	return ast.VisitParameter(v, n)
}

func (v *TypeChecker) VisitArrayAccess(n *ast.ArrayAccess) (interface{}, error) {
	return ast.VisitArrayAccess(v, n)
}

func (v *TypeChecker) VisitBooleanLiteral(n *ast.BooleanLiteral) (interface{}, error) {
	return ast.BooleanType, nil
}

func (v *TypeChecker) VisitBreak(n *ast.Break) (interface{}, error) {
	// Check For/While Loop
	return nil, nil
}

func (v *TypeChecker) VisitContinue(n *ast.Continue) (interface{}, error) {
	// Check For/While Loop
	return nil, nil
}

func (v *TypeChecker) VisitDml(n *ast.Dml) (interface{}, error) {
	_, err := n.Expression.Accept(v)
	if err != nil {
		// v.Errors
	}
	return nil, nil
}

func (v *TypeChecker) VisitDoubleLiteral(n *ast.DoubleLiteral) (interface{}, error) {
	return ast.DoubleType, nil
}

func (v *TypeChecker) VisitFieldDeclaration(n *ast.FieldDeclaration) (interface{}, error) {
	return ast.VisitFieldDeclaration(v, n)
}

func (v *TypeChecker) VisitTry(n *ast.Try) (interface{}, error) {
	n.Block.Accept(v)
	for _, c := range n.CatchClause {
		c.Accept(v)
	}
	n.FinallyBlock.Accept(v)
	return nil, nil
}

func (v *TypeChecker) VisitCatch(n *ast.Catch) (interface{}, error) {
	t, _ := n.Type.Accept(v)
	v.Context.Env.Set(n.Identifier, t) // TODO: append scope
	n.Block.Accept(v)
	return nil, nil
}

func (v *TypeChecker) VisitFinally(n *ast.Finally) (interface{}, error) {
	n.Block.Accept(v)
	return nil, nil
}

func (v *TypeChecker) VisitFor(n *ast.For) (interface{}, error) {
	n.Control.Accept(v)
	n.Statements.Accept(v)
	return nil, nil
}

func (v *TypeChecker) VisitForControl(n *ast.ForControl) (interface{}, error) {
	n.ForInit.Accept(v)
	for _, u := range n.ForUpdate {
		u.Accept(v)
	}
	n.Expression.Accept(v)
	return nil, nil
}

func (v *TypeChecker) VisitEnhancedForControl(n *ast.EnhancedForControl) (interface{}, error) {
	return ast.VisitEnhancedForControl(v, n)
}

func (v *TypeChecker) VisitIf(n *ast.If) (interface{}, error) {
	t, err := n.Condition.Accept(v)
	if err != nil {
		// v.Errors
	}
	if t != ast.BooleanType {
		// v.Errors
	}
	n.IfStatement.Accept(v)
	n.ElseStatement.Accept(v)
	return nil, nil
}

func (v *TypeChecker) VisitMethodDeclaration(n *ast.MethodDeclaration) (interface{}, error) {
	v.Context.CurrentMethod = n
	env := newTypeEnv(nil)
	v.Context.Env = env
	for _, param := range n.Parameters {
		p := param.(*ast.Parameter)
		t, _ := p.Type.Accept(v)
		env.Set(p.Name, t)
	}
	n.Statements.Accept(v)
	v.Context.CurrentMethod = nil
	return ast.VisitMethodDeclaration(v, n)
}

func (v *TypeChecker) VisitMethodInvocation(n *ast.MethodInvocation) (interface{}, error) {
	// Use Resolver
	return ast.VisitMethodInvocation(v, n)
}

func (v *TypeChecker) VisitNew(n *ast.New) (interface{}, error) {
	return ast.VisitNew(v, n)
}

func (v *TypeChecker) VisitNullLiteral(n *ast.NullLiteral) (interface{}, error) {
	return nil, nil
}

func (v *TypeChecker) VisitUnaryOperator(n *ast.UnaryOperator) (interface{}, error) {
	return ast.VisitUnaryOperator(v, n)
}

func (v *TypeChecker) VisitBinaryOperator(n *ast.BinaryOperator) (interface{}, error) {
	l, err := n.Left.Accept(v)
	if err != nil {
		return nil, err
	}
	r, err := n.Right.Accept(v)
	if err != nil {
		return nil, err
	}
	if n.Op == "+" {
		if l != ast.IntegerType && l != ast.StringType && l != ast.DoubleType {
			// v.Errors
		}
	}
	if r != l {
		// v.Errors
	}
	return nil, nil
}

func (v *TypeChecker) VisitReturn(n *ast.Return) (interface{}, error) {
	exp, err := n.Expression.Accept(v)
	if err != nil {
		v.Errors = append(v.Errors, err)
	}
	if v.Context.CurrentMethod.ReturnType != exp {
		// v.Errors
	}
	return nil, nil
}

func (v *TypeChecker) VisitThrow(n *ast.Throw) (interface{}, error) {
	_, err := n.Expression.Accept(v)
	if err != nil {
		// v.Errors
	}
	// Check Subclass of Exception
	return nil, nil
}

func (v *TypeChecker) VisitSoql(n *ast.Soql) (interface{}, error) {
	return ast.VisitSoql(v, n)
}

func (v *TypeChecker) VisitSosl(n *ast.Sosl) (interface{}, error) {
	return ast.VisitSosl(v, n)
}

func (v *TypeChecker) VisitStringLiteral(n *ast.StringLiteral) (interface{}, error) {
	return ast.StringType, nil
}

func (v *TypeChecker) VisitSwitch(n *ast.Switch) (interface{}, error) {
	exp, err := n.Expression.Accept(v)
	if err != nil {
		// v.Errors
	}
	for _, w := range n.WhenStatements {
		t, err := w.Accept(v)
		if err != nil {
			// v.Errors
		}
		if t != exp {
			// v.Errors
		}
	}
	return nil, nil
}

func (v *TypeChecker) VisitTrigger(n *ast.Trigger) (interface{}, error) {
	panic("Not pass")
	return nil, nil
}

func (v *TypeChecker) VisitTriggerTiming(n *ast.TriggerTiming) (interface{}, error) {
	panic("Not pass")
	return nil, nil
}

func (v *TypeChecker) VisitVariableDeclaration(n *ast.VariableDeclaration) (interface{}, error) {
	declType, _ := n.Type.Accept(v)
	for _, d := range n.Declarators {
		t, err := d.Accept(v)
		if err != nil {
			// v.Errors
		}
		decl := d.(*ast.VariableDeclarator)
		v.Context.Env.Set(decl.Name, declType)
		if declType != t {
			// v.Errors
		}
	}
	return nil, nil
}

func (v *TypeChecker) VisitVariableDeclarator(n *ast.VariableDeclarator) (interface{}, error) {
	return n.Expression.Accept(v)
}

func (v *TypeChecker) VisitWhen(n *ast.When) (interface{}, error) {
	return ast.VisitWhen(v, n)
}

func (v *TypeChecker) VisitWhenType(n *ast.WhenType) (interface{}, error) {
	return ast.VisitWhenType(v, n)
}

func (v *TypeChecker) VisitWhile(n *ast.While) (interface{}, error) {
	t, err := n.Condition.Accept(v)
	if err != nil {
		// v.Errors
	}
	if t != ast.BooleanType {
		// v.Errors
	}
	for _, stmt := range n.Statements {
		stmt.Accept(v)
	}
	return ast.VisitWhile(v, n)
}

func (v *TypeChecker) VisitNothingStatement(n *ast.NothingStatement) (interface{}, error) {
	panic("Not pass")
	return nil, nil
}

func (v *TypeChecker) VisitCastExpression(n *ast.CastExpression) (interface{}, error) {
	t, _ := n.CastType.Accept(v)
	exp, _ := n.Expression.Accept(v)
	if t != exp {
		// v.Errors
	}
	return t, nil
}

func (v *TypeChecker) VisitFieldAccess(n *ast.FieldAccess) (interface{}, error) {
	return v.ResolveVariable(n), nil
}

func (v *TypeChecker) VisitType(n *ast.TypeRef) (interface{}, error) {
	return v.ResolveType(n), nil
}

func (v *TypeChecker) VisitBlock(n *ast.Block) (interface{}, error) {
	for _, stmt := range n.Statements {
		stmt.Accept(v)
	}
	return nil, nil
}

func (v *TypeChecker) VisitGetterSetter(n *ast.GetterSetter) (interface{}, error) {
	return ast.VisitGetterSetter(v, n)
}

func (v *TypeChecker) VisitPropertyDeclaration(n *ast.PropertyDeclaration) (interface{}, error) {
	return ast.VisitPropertyDeclaration(v, n)
}

func (v *TypeChecker) VisitArrayInitializer(n *ast.ArrayInitializer) (interface{}, error) {
	return ast.VisitArrayInitializer(v, n)
}

func (v *TypeChecker) VisitArrayCreator(n *ast.ArrayCreator) (interface{}, error) {
	return ast.VisitArrayCreator(v, n)
}

func (v *TypeChecker) VisitSoqlBindVariable(n *ast.SoqlBindVariable) (interface{}, error) {
	return ast.VisitSoqlBindVariable(v, n)
}

func (v *TypeChecker) VisitTernalyExpression(n *ast.TernalyExpression) (interface{}, error) {
	return ast.VisitTernalyExpression(v, n)
}

func (v *TypeChecker) VisitMapCreator(n *ast.MapCreator) (interface{}, error) {
	return ast.VisitMapCreator(v, n)
}

func (v *TypeChecker) VisitSetCreator(n *ast.SetCreator) (interface{}, error) {
	return ast.VisitSetCreator(v, n)
}

func (v *TypeChecker) VisitName(n *ast.Name) (interface{}, error) {
	// Use Resolver
	return ast.VisitName(v, n)
}

func (v *TypeChecker) VisitConstructorDeclaration(n *ast.ConstructorDeclaration) (interface{}, error) {
	return ast.VisitConstructorDeclaration(v, n)
}

func (v *TypeChecker) ResolveType(n ast.Node) ast.Type {
	t := n.(*ast.TypeRef)
	if len(t.Name) == 1 {
		val, ok := PrimitiveMap[t.Name[0]]
		if ok {
			return val
		}
		classType, ok := v.Context.ClassTypes.Get(t.Name[0])
		if ok {
			return classType
		}
	}
	return nil
}

func (v *TypeChecker) ResolveVariable(n ast.Node) error {
	// VariableResolver.resolve(n, v.Context)
	return nil
}

var PrimitiveMap = map[string]ast.Type{
	"Integer": ast.IntegerType,
	"String":  ast.StringType,
	"Double":  ast.DoubleType,
	"Boolean": ast.BooleanType,
}
