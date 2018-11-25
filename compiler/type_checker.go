package compiler

import (
	"fmt"

	"github.com/tzmfreedom/goland/ast"
	"github.com/tzmfreedom/goland/builtin"
)

type TypeChecker struct {
	Context *Context
	Errors  []*Error
}

func NewTypeChecker() *TypeChecker {
	checker := &TypeChecker{
		Context: NewContext(),
		Errors:  []*Error{},
	}
	return checker
}

func (v *TypeChecker) VisitClassType(n *builtin.ClassType) (interface{}, error) {
	v.Context.CurrentClass = n
	if n.StaticFields != nil {
		for _, f := range n.StaticFields.Data {
			t, _ := f.Type.Accept(v)
			e, _ := f.Expression.Accept(v)
			if t != e {
				v.AddError(fmt.Sprintf("expression <%s> does not match <%s>", builtin.TypeName(e), builtin.TypeName(t)), f.Expression)
			}
		}
	}

	if n.InstanceFields != nil {
		for _, f := range n.InstanceFields.Data {
			t, _ := f.Type.Accept(v)
			e, _ := f.Expression.Accept(v)
			if t != e {
				v.AddError(fmt.Sprintf("expression <%s> does not match <%s>", builtin.TypeName(e), builtin.TypeName(t)), f.Expression)
			}
		}
	}

	if n.StaticMethods != nil {
		for _, methods := range n.StaticMethods.All() {
			for _, m := range methods {
				m.Accept(v)
			}
		}
	}

	if n.InstanceMethods != nil {
		for _, methods := range n.InstanceMethods.All() {
			for _, m := range methods {
				m.Accept(v)
			}
		}
	}
	v.Context.CurrentClass = nil
	return nil, nil
}

func (v *TypeChecker) VisitClassDeclaration(n *ast.ClassDeclaration) (interface{}, error) {
	panic("Not pass")
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
	panic("Not pass")
	return ast.VisitInterfaceDeclaration(v, n)
}

func (v *TypeChecker) VisitIntegerLiteral(n *ast.IntegerLiteral) (interface{}, error) {
	return builtin.IntegerType, nil
}

func (v *TypeChecker) VisitParameter(n *ast.Parameter) (interface{}, error) {
	n.Type.Accept(v)
	return nil, nil
}

func (v *TypeChecker) VisitArrayAccess(n *ast.ArrayAccess) (interface{}, error) {
	n.Receiver.Accept(v)
	t, _ := n.Key.Accept(v)
	if t != builtin.IntegerType && t != builtin.StringType {
		v.AddError(fmt.Sprintf("array key <%v> must be Integer or string", builtin.TypeName(t)), n.Key)
	}
	return nil, nil
}

func (v *TypeChecker) VisitBooleanLiteral(n *ast.BooleanLiteral) (interface{}, error) {
	return builtin.BooleanType, nil
}

func (v *TypeChecker) VisitBreak(n *ast.Break) (interface{}, error) {
	if !ast.IsDecendants(n, "For") && !ast.IsDecendants(n, "While") {
		v.AddError("break must be in for/while loop", n)
	}
	return nil, nil
}

func (v *TypeChecker) VisitContinue(n *ast.Continue) (interface{}, error) {
	if !ast.IsDecendants(n, "For") && !ast.IsDecendants(n, "While") {
		v.AddError("continue must be in for/while loop", n)
	}
	return nil, nil
}

func (v *TypeChecker) VisitDml(n *ast.Dml) (interface{}, error) {
	t, _ := n.Expression.Accept(v)
	if t != builtin.ListType {

	}
	return nil, nil
}

func (v *TypeChecker) VisitDoubleLiteral(n *ast.DoubleLiteral) (interface{}, error) {
	return builtin.DoubleType, nil
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
	v.Context.Env.Set(n.Identifier, t.(*builtin.ClassType)) // TODO: append scope
	n.Block.Accept(v)
	return nil, nil
}

func (v *TypeChecker) VisitFinally(n *ast.Finally) (interface{}, error) {
	n.Block.Accept(v)
	return nil, nil
}

func (v *TypeChecker) VisitFor(n *ast.For) (interface{}, error) {
	n.Control.Accept(v)
	if n.Statements != nil {
		n.Statements.Accept(v)
	}
	return nil, nil
}

func (v *TypeChecker) VisitForControl(n *ast.ForControl) (interface{}, error) {
	if n.ForInit != nil {
		n.ForInit.Accept(v)
	}
	if n.ForUpdate != nil {
		for _, u := range n.ForUpdate {
			u.Accept(v)
		}
	}
	if n.Expression != nil {
		t, _ := n.Expression.Accept(v)
		if t != builtin.BooleanType {
			v.AddError(fmt.Sprintf("condition <%s> must be Boolean expression", builtin.TypeName(t)), n.Expression)
		}
	}
	return nil, nil
}

func (v *TypeChecker) VisitEnhancedForControl(n *ast.EnhancedForControl) (interface{}, error) {
	return ast.VisitEnhancedForControl(v, n)
}

func (v *TypeChecker) VisitIf(n *ast.If) (interface{}, error) {
	t, _ := n.Condition.Accept(v)
	if t != builtin.BooleanType {
		v.AddError(fmt.Sprintf("condition <%s> must be Boolean expression", builtin.TypeName(t)), n.Condition)
	}
	n.IfStatement.Accept(v)
	if n.ElseStatement != nil {
		n.ElseStatement.Accept(v)
	}
	return nil, nil
}

func (v *TypeChecker) VisitMethodDeclaration(n *ast.MethodDeclaration) (interface{}, error) {
	v.Context.CurrentMethod = n
	env := newTypeEnv(nil)
	v.Context.Env = env
	for _, param := range n.Parameters {
		p := param.(*ast.Parameter)
		t, _ := p.Type.Accept(v)
		env.Set(p.Name, t.(*builtin.ClassType))
	}
	n.Statements.Accept(v)
	v.Context.CurrentMethod = nil
	return nil, nil
}

func (v *TypeChecker) VisitMethodInvocation(n *ast.MethodInvocation) (interface{}, error) {
	resolver := TypeResolver{Context: v.Context}

	nameOrExp := n.NameOrExpression
	types := make([]*builtin.ClassType, len(n.Parameters))
	for i, p := range n.Parameters {
		t, _ := p.Accept(v)
		types[i] = t.(*builtin.ClassType)
	}
	if name, ok := nameOrExp.(*ast.Name); ok {
		// TODO: implement
		if name.Value[0] == "Debugger" {
			return nil, nil
		}
		method, err := resolver.ResolveMethod(name.Value, types)
		if err != nil {
			// TODO: implement
			return nil, nil
		}
		if method.ReturnType != nil {
			return method.ReturnType.Accept(v)
		}
		return nil, nil
	} else if fieldAccess, ok := nameOrExp.(*ast.FieldAccess); ok {
		classType, _ := fieldAccess.Expression.Accept(v)
		method, err := resolver.FindInstanceMethod(
			classType.(*builtin.ClassType),
			fieldAccess.FieldName,
			types,
			MODIFIER_PUBLIC_ONLY,
		)
		if err != nil {
			// TODO: implmenet
			return nil, nil
		}
		if method.ReturnType != nil {
			return method.ReturnType.Accept(v)
		}
		return nil, nil
	}
	return nil, nil
}

func (v *TypeChecker) VisitNew(n *ast.New) (interface{}, error) {
	t, _ := n.Type.Accept(v)
	for _, p := range n.Parameters {
		p.Accept(v)
	}
	return t, nil
}

func (v *TypeChecker) VisitNullLiteral(n *ast.NullLiteral) (interface{}, error) {
	return nil, nil
}

func (v *TypeChecker) VisitUnaryOperator(n *ast.UnaryOperator) (interface{}, error) {
	t, _ := n.Expression.Accept(v)
	if t != builtin.IntegerType {
		v.AddError(fmt.Sprintf("expression <%s> must be Integer", builtin.TypeName(t)), n.Expression)
	}
	return nil, nil
}

func (v *TypeChecker) VisitBinaryOperator(n *ast.BinaryOperator) (interface{}, error) {
	l, _ := n.Left.Accept(v)
	r, _ := n.Right.Accept(v)
	if n.Op == "+" {
		if l != builtin.IntegerType && l != builtin.StringType && l != builtin.DoubleType {
			v.AddError(fmt.Sprintf("expression <%s> must be Integer, String or Double", builtin.TypeName(l)), n.Left)
		}
		if (l == builtin.StringType || r == builtin.StringType) && l != r {
			v.AddError(fmt.Sprintf("expression <%s> does not match <%s>", builtin.TypeName(l), builtin.TypeName(r)), n.Left)
		}
		if l == builtin.DoubleType || r == builtin.DoubleType {
			return builtin.DoubleType, nil
		} else if l == builtin.StringType {
			return builtin.StringType, nil
		} else {
			return builtin.IntegerType, nil
		}
	}
	if n.Op == "-" || n.Op == "*" || n.Op == "/" || n.Op == "%" {
		if l != builtin.IntegerType && l != builtin.DoubleType {
			v.AddError(fmt.Sprintf("expression <%s> must be Integer or Double", builtin.TypeName(l)), n.Left)
		} else if r != builtin.IntegerType && r != builtin.DoubleType {
			v.AddError(fmt.Sprintf("expression <%s> must be Integer or Double", builtin.TypeName(r)), n.Right)
		}
		if l == builtin.DoubleType || r == builtin.DoubleType {
			return builtin.DoubleType, nil
		} else {
			return builtin.IntegerType, nil
		}
	}
	if n.Op == "=" || n.Op == "+=" || n.Op == "-=" || n.Op == "*=" || n.Op == "/=" || n.Op == "%=" {
		if l != r {
			v.AddError(fmt.Sprintf("expression <%s> does not match <%s>", builtin.TypeName(l), builtin.TypeName(r)), n.Left)
		}
		return l, nil
	}
	if n.Op == "==" || n.Op == "!=" || n.Op == "<" || n.Op == "<=" || n.Op == ">" || n.Op == ">=" {
		return builtin.BooleanType, nil
	}
	return nil, nil
}

func (v *TypeChecker) VisitReturn(n *ast.Return) (interface{}, error) {
	if v.Context.CurrentMethod.ReturnType == nil {
		if n.Expression != nil {
			exp, _ := n.Expression.Accept(v)
			v.AddError(fmt.Sprintf("return type <%s> does not match void", builtin.TypeName(exp)), n.Expression)
		}
	} else {
		retType, _ := v.Context.CurrentMethod.ReturnType.Accept(v)
		if n.Expression == nil {
			v.AddError(fmt.Sprintf("return type <void> does not match %v", builtin.TypeName(retType)), n.Expression)
			return nil, nil
		}
		exp, _ := n.Expression.Accept(v)
		if retType != exp {
			v.AddError(fmt.Sprintf("return type <%s> does not match %v", builtin.TypeName(exp), builtin.TypeName(retType)), n.Expression)
		}
	}
	return nil, nil
}

func (v *TypeChecker) VisitThrow(n *ast.Throw) (interface{}, error) {
	_, _ = n.Expression.Accept(v)
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
	return builtin.StringType, nil
}

func (v *TypeChecker) VisitSwitch(n *ast.Switch) (interface{}, error) {
	exp, _ := n.Expression.Accept(v)
	for _, w := range n.WhenStatements {
		t, _ := w.Accept(v)
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
		t, _ := d.Accept(v)
		decl := d.(*ast.VariableDeclarator)
		v.Context.Env.Set(decl.Name, declType.(*builtin.ClassType))
		if declType != t {
			v.AddError(fmt.Sprintf("expression <%s> does not match <%s>", builtin.TypeName(declType), builtin.TypeName(t)), n.Type)
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
	t, _ := n.Condition.Accept(v)
	if t != builtin.BooleanType {
		v.AddError(fmt.Sprintf("condition <%s> must be Boolean expression", builtin.TypeName(t)), n.Condition)
	}
	n.Statements.Accept(v)
	return nil, nil
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
	classType, _ := n.Expression.Accept(v)
	f, ok := classType.(*builtin.ClassType).InstanceFields.Get(n.FieldName)
	if !ok {
		v.AddError(fmt.Sprintf("field <%s> does not exist", n.FieldName), n)
	}
	resolver := &TypeResolver{Context: v.Context}
	t, err := resolver.ResolveType(f.Type.(*ast.TypeRef).Name)
	if err != nil {
		// TODO: ErrorHandling
	}
	return t, nil
}

func (v *TypeChecker) VisitType(n *ast.TypeRef) (interface{}, error) {
	resolver := &TypeResolver{Context: v.Context}
	t, err := resolver.ResolveType(n.Name)
	if err != nil {
		v.AddError(err.Error(), n)
	}
	return t, nil
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
	c, _ := n.Condition.Accept(v)
	if c != builtin.BooleanType {
		v.AddError(fmt.Sprintf("condition <%s> must be Boolean expression", builtin.TypeName(c)), n.Condition)
	}
	t, _ := n.TrueExpression.Accept(v)
	f, _ := n.FalseExpression.Accept(v)
	if t != f {
		v.AddError(fmt.Sprintf("condition does not match %s != %s", builtin.TypeName(t), builtin.TypeName(f)), n.TrueExpression)
	}
	return t, nil
}

func (v *TypeChecker) VisitMapCreator(n *ast.MapCreator) (interface{}, error) {
	return ast.VisitMapCreator(v, n)
}

func (v *TypeChecker) VisitSetCreator(n *ast.SetCreator) (interface{}, error) {
	return ast.VisitSetCreator(v, n)
}

func (v *TypeChecker) VisitName(n *ast.Name) (interface{}, error) {
	resolver := TypeResolver{Context: v.Context}
	t, err := resolver.ResolveVariable(n.Value)
	if err != nil {
		v.AddError(err.Error(), n)
	}
	return t, nil
}

func (v *TypeChecker) VisitConstructorDeclaration(n *ast.ConstructorDeclaration) (interface{}, error) {
	env := newTypeEnv(nil)
	v.Context.Env = env
	for _, param := range n.Parameters {
		p := param.(*ast.Parameter)
		t, _ := p.Type.Accept(v)
		env.Set(p.Name, t.(*builtin.ClassType))
	}
	n.Statements.Accept(v)
	return nil, nil
}

func (v *TypeChecker) AddError(msg string, node ast.Node) {
	v.Errors = append(v.Errors, &Error{Message: msg, Node: node})
}

type Error struct {
	Message string
	Node    ast.Node
}
