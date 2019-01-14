package compiler

import (
	"errors"
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
			if !t.(*builtin.ClassType).Equals(e.(*builtin.ClassType)) {
				v.AddError(fmt.Sprintf("expression <%s> does not match <%s>", e.(*builtin.ClassType).String(), t.(*builtin.ClassType).String()), f.Expression)
			}
		}
	}

	if n.InstanceFields != nil {
		for _, f := range n.InstanceFields.Data {
			t, _ := f.Type.Accept(v)
			e, _ := f.Expression.Accept(v)
			if e == nil {
				continue
			}
			if !t.(*builtin.ClassType).Equals(e.(*builtin.ClassType)) {
				v.AddError(fmt.Sprintf("expression <%s> does not match <%s>", e.(*builtin.ClassType).String(), t.(*builtin.ClassType).String()), f.Expression)
			}
		}
	}

	if n.StaticMethods != nil {
		for _, methods := range n.StaticMethods.All() {
			for _, m := range methods {
				_, err := v.VisitMethod(m)
				if err != nil {
					return nil, err
				}
			}
		}
	}

	if n.InstanceMethods != nil {
		for _, methods := range n.InstanceMethods.All() {
			for _, m := range methods {
				_, err := v.VisitMethod(m)
				if err != nil {
					return nil, err
				}
			}
		}
	}

	if n.Constructors != nil {
		for _, m := range n.Constructors {
			_, err := v.VisitMethod(m)
			if err != nil {
				return nil, err
			}
		}
	}

	if n.SuperClass != nil {
		_, err := n.SuperClass.Accept(v)
		if err != nil {
			return nil, err
		}
	}

	if n.ImplementClasses != nil {
		for _, impl := range n.ImplementClasses {
			_, err := impl.Accept(v)
			if err != nil {
				return nil, err
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
	k, err := n.Receiver.Accept(v)
	if err != nil {
		return nil, err
	}
	klass := k.(*builtin.ClassType)
	t, _ := n.Key.Accept(v)
	if t != builtin.IntegerType && t != builtin.StringType {
		v.AddError(fmt.Sprintf("array key <%v> must be Integer or string", t.(*builtin.ClassType).String()), n.Key)
	}
	generics := klass.Extra["generics"].([]*builtin.ClassType)
	if len(generics) == 0 {
		return nil, v.compileError("generics is not specified", n)
	}
	return generics[0], nil
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
	t, err := n.Type.Accept(v)
	if err != nil {
		return nil, err
	}
	v.Context.Env.Set(n.Identifier, t.(*builtin.ClassType)) // TODO: append scope
	_, err = n.Block.Accept(v)
	if err != nil {
		return nil, err
	}
	return nil, nil
}

func (v *TypeChecker) VisitFinally(n *ast.Finally) (interface{}, error) {
	n.Block.Accept(v)
	return nil, nil
}

func (v *TypeChecker) VisitFor(n *ast.For) (interface{}, error) {
	return v.NewEnv(func() (interface{}, error) {
		_, err := n.Control.Accept(v)
		if err != nil {
			return nil, err
		}
		if n.Statements != nil {
			_, err := n.Statements.Accept(v)
			if err != nil {
				return nil, err
			}
		}
		return nil, nil
	})
}

func (v *TypeChecker) VisitForControl(n *ast.ForControl) (interface{}, error) {
	if n.ForInit != nil {
		for _, forInit := range n.ForInit {
			_, err := forInit.Accept(v)
			if err != nil {
				return nil, err
			}
		}
	}
	if n.ForUpdate != nil {
		for _, u := range n.ForUpdate {
			_, err := u.Accept(v)
			if err != nil {
				return nil, err
			}
		}
	}
	if n.Expression != nil {
		t, err := n.Expression.Accept(v)
		if err != nil {
			return nil, err
		}
		if t != builtin.BooleanType {
			v.AddError(fmt.Sprintf("condition <%s> must be Boolean expression", t.(*builtin.ClassType).String()), n.Expression)
		}
	}
	return nil, nil
}

func (v *TypeChecker) VisitEnhancedForControl(n *ast.EnhancedForControl) (interface{}, error) {
	t, _ := n.Type.Accept(v)
	declClassType := t.(*builtin.ClassType)
	exp, _ := n.Expression.Accept(v)
	expClassType := exp.(*builtin.ClassType)
	v.Context.Env.Set(n.VariableDeclaratorId, declClassType)

	if expClassType.Name != "List" && expClassType.Name != "Set" {
		v.AddError(fmt.Sprintf("expression <%s> must be List or Set expression", expClassType.Name), n)
		return nil, nil
	}

	genericsType := expClassType.Extra["generics"].([]*builtin.ClassType)[0]
	if !declClassType.Equals(genericsType) {
		v.AddError(fmt.Sprintf("expression <%s> must be <%s> expression", declClassType.Name, expClassType.Name), n)
	}
	return nil, nil
}

func (v *TypeChecker) VisitIf(n *ast.If) (interface{}, error) {
	t, _ := n.Condition.Accept(v)
	if t != builtin.BooleanType {
		v.AddError(fmt.Sprintf("condition <%s> must be Boolean expression", t.(*builtin.ClassType).String()), n.Condition)
	}
	if _, err := n.IfStatement.Accept(v); err != nil {
		return nil, err
	}
	if n.ElseStatement != nil {
		if _, err := n.ElseStatement.Accept(v); err != nil {
			return nil, err
		}
	}
	return nil, nil
}

func (v *TypeChecker) VisitMethodDeclaration(n *ast.MethodDeclaration) (interface{}, error) {
	panic(123)
	// v.Context.CurrentMethod = n
	env := newTypeEnv(nil)
	v.Context.Env = env
	classType, ok := v.Context.ClassTypes.Get(n.Parent.(*ast.ClassDeclaration).Name)
	if !ok {
		panic("not found")
	}
	v.Context.Env.Set("this", classType)
	for _, param := range n.Parameters {
		p := param.(*ast.Parameter)
		t, _ := p.Type.Accept(v)
		env.Set(p.Name, t.(*builtin.ClassType))
	}
	r, err := n.Statements.Accept(v)
	if err != nil {
		v.Context.CurrentMethod = nil
		return nil, err
	}
	if n.ReturnType != nil && r == nil {
		retType, _ := n.ReturnType.Accept(v)
		v.AddError(fmt.Sprintf("return type <void> does not match %v", retType.(*builtin.ClassType).String()), n)
	}

	v.Context.CurrentMethod = nil
	return nil, nil
}

func (v *TypeChecker) VisitMethodInvocation(n *ast.MethodInvocation) (interface{}, error) {
	resolver := TypeResolver{Context: v.Context}

	nameOrExp := n.NameOrExpression
	types := make([]*builtin.ClassType, len(n.Parameters))
	for i, p := range n.Parameters {
		t, err := p.Accept(v)
		if err != nil {
			return nil, err
		}
		types[i] = t.(*builtin.ClassType)
	}
	if name, ok := nameOrExp.(*ast.Name); ok {
		// TODO: implement
		if name.Value[0] == "Debugger" {
			return nil, nil
		}
		receiverType, method, err := resolver.ResolveMethod(name.Value, types)
		if err != nil {
			return nil, v.compileError(err.Error(), n)
		}
		if method.ReturnType != nil {
			retType := method.ReturnType.(*ast.TypeRef).Name[0]
			if retType == "T:1" || retType == "T:2" {
				generics := receiverType.Extra["generics"].([]*builtin.ClassType)
				if retType == "T:1" {
					return generics[0], nil
				} else {
					return generics[1], nil
				}
			}
			return method.ReturnType.Accept(v)
		}
	} else if fieldAccess, ok := nameOrExp.(*ast.FieldAccess); ok {
		classType, err := fieldAccess.Expression.Accept(v)
		if err != nil {
			return nil, err
		}
		receiverType, method, err := resolver.FindInstanceMethod(
			classType.(*builtin.ClassType),
			fieldAccess.FieldName,
			types,
			MODIFIER_PUBLIC_ONLY,
		)
		if err != nil {
			return nil, v.compileError(err.Error(), n)
		}
		if method.ReturnType != nil {
			// TODO: duplicate code
			retType := method.ReturnType.(*ast.TypeRef).Name[0]
			if retType == "T:1" || retType == "T:2" {
				generics := receiverType.Extra["generics"].([]*builtin.ClassType)
				if retType == "T:1" {
					return generics[0], nil
				} else {
					return generics[1], nil
				}
			}
			return method.ReturnType.Accept(v)
		}
	}
	return nil, nil
}

func (v *TypeChecker) VisitNew(n *ast.New) (interface{}, error) {
	t, err := n.Type.Accept(v)
	if err != nil {
		return nil, err
	}
	params := make([]*builtin.ClassType, len(n.Parameters))
	for i, p := range n.Parameters {
		param, err := p.Accept(v)
		if err != nil {
			return nil, err
		}
		params[i] = param.(*builtin.ClassType)
	}
	typeResolver := &TypeResolver{Context: v.Context}
	classType := t.(*builtin.ClassType)

	if classType.Constructors == nil {
		v.AddError(fmt.Sprintf("Type cannot be constructed: %s", classType.Name), n)
	} else if len(classType.Constructors) > 0 {
		_, method, err := typeResolver.SearchConstructor(classType, params)
		if err != nil {
			return nil, err
		}
		if method == nil && typeResolver.HasConstructor(classType) {
			v.AddError(fmt.Sprintf("constructor <%s> not found", classType.Name), n)
		} else {
			// TODO: for protected impl
			if method.IsPrivate() && v.Context.CurrentClass != classType {
				v.AddError(fmt.Sprintf("constructor <%s> not found", classType.Name), n)
			}
		}
	}
	return t, nil
}

func (v *TypeChecker) VisitNullLiteral(n *ast.NullLiteral) (interface{}, error) {
	return builtin.NullType, nil
}

func (v *TypeChecker) VisitUnaryOperator(n *ast.UnaryOperator) (interface{}, error) {
	t, _ := n.Expression.Accept(v)
	if t != builtin.IntegerType {
		v.AddError(fmt.Sprintf("expression <%s> must be Integer", t.(*builtin.ClassType).String()), n.Expression)
	}
	return nil, nil
}

func (v *TypeChecker) VisitBinaryOperator(n *ast.BinaryOperator) (interface{}, error) {
	r, _ := n.Right.Accept(v)
	if n.Op == "=" ||
		n.Op == "=+" ||
		n.Op == "=-" ||
		n.Op == "=*" ||
		n.Op == "=/" {

		var l *builtin.ClassType
		var err error
		resolver := &TypeResolver{Context: v.Context}
		switch leftNode := n.Left.(type) {
		case *ast.Name:
			l, err = resolver.ResolveVariable(leftNode.Value, true)
			if err != nil {
				return nil, v.compileError(err.Error(), n)
			}
		case *ast.FieldAccess:
			classType, _ := leftNode.Expression.Accept(v)
			f, _ := resolver.findInstanceField(classType.(*builtin.ClassType), leftNode.FieldName, MODIFIER_PUBLIC_ONLY, true)
			l, err = resolver.ResolveType(f.Type.(*ast.TypeRef).Name)
			if err != nil {
				return nil, v.compileError(err.Error(), n)
			}
		}
		if r != nil && !l.Equals(r.(*builtin.ClassType)) {
			v.AddError(fmt.Sprintf("expression <%s> does not match <%s>", l.String(), r.(*builtin.ClassType).String()), n.Left)
		}
		return l, nil
	} else {
		l, _ := n.Left.Accept(v)
		if n.Op == "+" {
			if l != builtin.IntegerType && l != builtin.StringType && l != builtin.DoubleType {
				v.AddError(fmt.Sprintf("expression <%s> must be Integer, String or Double", l.(*builtin.ClassType).String()), n.Left)
			}
			if (l == builtin.StringType || r == builtin.StringType) && l != r {
				v.AddError(fmt.Sprintf("expression <%s> does not match <%s>", l.(*builtin.ClassType).String(), r.(*builtin.ClassType).String()), n.Left)
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
				v.AddError(fmt.Sprintf("expression <%s> must be Integer or Double", l.(*builtin.ClassType).String()), n.Left)
			} else if r != builtin.IntegerType && r != builtin.DoubleType {
				v.AddError(fmt.Sprintf("expression <%s> must be Integer or Double", r.(*builtin.ClassType).String()), n.Right)
			}
			if l == builtin.DoubleType || r == builtin.DoubleType {
				return builtin.DoubleType, nil
			} else {
				return builtin.IntegerType, nil
			}
		}
		if n.Op == "==" || n.Op == "!=" || n.Op == "<" || n.Op == "<=" || n.Op == ">" || n.Op == ">=" {
			return builtin.BooleanType, nil
		}
	}
	return nil, nil
}

func (v *TypeChecker) VisitReturn(n *ast.Return) (interface{}, error) {
	if v.Context.CurrentMethod.ReturnType == nil {
		if n.Expression != nil {
			exp, _ := n.Expression.Accept(v)
			v.AddError(fmt.Sprintf("return type <%s> does not match void", exp.(*builtin.ClassType).String()), n.Expression)
		}
		return nil, nil
	}

	retType, _ := v.Context.CurrentMethod.ReturnType.Accept(v)
	if n.Expression == nil {
		v.AddError(fmt.Sprintf("return type <void> does not match %v", retType.(*builtin.ClassType).String()), n.Expression)
	}
	exp, _ := n.Expression.Accept(v)
	if !retType.(*builtin.ClassType).Equals(exp.(*builtin.ClassType)) {
		v.AddError(fmt.Sprintf("return type <%s> does not match %v", exp.(*builtin.ClassType).String(), retType.(*builtin.ClassType).String()), n.Expression)
	}
	return exp, nil
}

func (v *TypeChecker) VisitThrow(n *ast.Throw) (interface{}, error) {
	r, err := n.Expression.Accept(v)
	if err != nil {
		return nil, err
	}
	// Check Subclass of Exception
	baseClass := r.(*builtin.ClassType)
	super := baseClass.SuperClass
	if super == nil {
		v.AddError(fmt.Sprintf("Throw expression must be of type exception: %s", baseClass.Name), n)
	} else {
		typeResolver := &TypeResolver{Context: v.Context}
		t, err := typeResolver.ResolveType(super.(*ast.TypeRef).Name)
		if err != nil {
			v.AddError(err.Error(), n)
		} else if t != builtin.ExceptionType {
			v.AddError(fmt.Sprintf("Throw expression must be of type exception: %s", baseClass.Name), n)
		}
	}
	return nil, nil
}

func (v *TypeChecker) VisitSoql(n *ast.Soql) (interface{}, error) {
	resolver := &TypeResolver{Context: v.Context}
	t, err := resolver.ResolveType([]string{n.FromObject})
	if err != nil {
		return nil, v.compileError(err.Error(), n)
	}
	return &builtin.ClassType{
		Name: "List",
		Extra: map[string]interface{}{
			"generics": []*builtin.ClassType{t},
		},
	}, nil
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
	declType, err := n.Type.Accept(v)
	if err != nil {
		return nil, err
	}
	for _, d := range n.Declarators {
		t, err := d.Accept(v)
		if err != nil {
			return nil, err
		}
		decl := d.(*ast.VariableDeclarator)
		if _, ok := v.Context.Env.Get(decl.Name); ok {
			v.AddError(fmt.Sprintf("variable declaration is duplicated <%s>", decl.Name), n.Type)
			continue
		}
		v.Context.Env.Set(decl.Name, declType.(*builtin.ClassType))
		if !declType.(*builtin.ClassType).Equals(t.(*builtin.ClassType)) {
			v.AddError(fmt.Sprintf("expression <%s> does not match <%s>", declType.(*builtin.ClassType).String(), t.(*builtin.ClassType).String()), n.Type)
		}
	}
	return nil, nil
}

func (v *TypeChecker) VisitVariableDeclarator(n *ast.VariableDeclarator) (interface{}, error) {
	if n.Expression != nil {
		return n.Expression.Accept(v)
	}
	return builtin.NullType, nil
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
		v.AddError(fmt.Sprintf("condition <%s> must be Boolean expression", t.(*builtin.ClassType).String()), n.Condition)
	}
	_, err := n.Statements.Accept(v)
	if err != nil {
		return nil, err
	}
	return nil, nil
}

func (v *TypeChecker) VisitNothingStatement(n *ast.NothingStatement) (interface{}, error) {
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
		return nil, v.compileError(fmt.Sprintf("field <%s> does not exist", n.FieldName), n)
	}
	return f.Type.(*ast.TypeRef).Accept(v)
}

func (v *TypeChecker) VisitType(n *ast.TypeRef) (interface{}, error) {
	resolver := &TypeResolver{Context: v.Context}
	t, err := resolver.ResolveType(n.Name)
	if err != nil {
		return nil, v.compileError(err.Error(), n)
	}
	if t.IsGeneric() {
		types := make([]*builtin.ClassType, len(n.Parameters))
		for i, p := range n.Parameters {
			classType, err := p.(*ast.TypeRef).Accept(v)
			if err != nil {
				return nil, v.compileError(err.Error(), n)
			}
			types[i] = classType.(*builtin.ClassType)
		}
		return &builtin.ClassType{
			Name:            t.Name,
			InstanceMethods: t.InstanceMethods,
			StaticMethods:   t.StaticMethods,
			Extra: map[string]interface{}{
				"generics": types,
			},
		}, nil
	}
	return t, nil
}

func (v *TypeChecker) VisitBlock(n *ast.Block) (interface{}, error) {
	var r interface{}
	var err error
	return v.NewEnv(func() (interface{}, error) {
		for _, stmt := range n.Statements {
			r, err = stmt.Accept(v)
			if err != nil {
				return nil, err
			}
		}
		if len(n.Statements) > 0 {
			if _, ok := n.Statements[len(n.Statements)-1].(*ast.Return); ok {
				return r, nil
			}
		}
		return nil, nil
	})
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
		v.AddError(fmt.Sprintf("condition <%s> must be Boolean expression", c.(*builtin.ClassType).String()), n.Condition)
	}
	t, _ := n.TrueExpression.Accept(v)
	f, _ := n.FalseExpression.Accept(v)
	if !t.(*builtin.ClassType).Equals(f.(*builtin.ClassType)) {
		v.AddError(fmt.Sprintf("expression does not match %s != %s", t.(*builtin.ClassType).String(), f.(*builtin.ClassType).String()), n.TrueExpression)
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
	t, err := resolver.ResolveVariable(n.Value, false)
	if err != nil {
		return nil, v.compileError(err.Error(), n)
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

func (v *TypeChecker) VisitMethod(n *builtin.Method) (interface{}, error) {
	v.Context.CurrentMethod = n
	env := newTypeEnv(nil)
	v.Context.Env = env
	classType, ok := v.Context.ClassTypes.Get(n.Parent.(*ast.ClassDeclaration).Name)
	if !ok {
		panic("not found")
	}
	v.Context.Env.Set("this", classType)
	for _, param := range n.Parameters {
		p := param.(*ast.Parameter)
		t, _ := p.Type.Accept(v)
		env.Set(p.Name, t.(*builtin.ClassType))
	}
	r, err := n.Statements.Accept(v)
	if err != nil {
		v.Context.CurrentMethod = nil
		return nil, err
	}
	if n.ReturnType != nil && r == nil {
		retType, _ := n.ReturnType.Accept(v)
		v.AddError(fmt.Sprintf("return type <void> does not match %v", retType.(*builtin.ClassType).String()), n.ReturnType)
	}

	v.Context.CurrentMethod = nil
	return nil, nil
}

func (v *TypeChecker) AddError(msg string, node ast.Node) {
	v.Errors = append(v.Errors, &Error{Message: msg, Node: node})
}

func (v *TypeChecker) compileError(msg string, n ast.Node) error {
	v.AddError(msg, n)
	return errors.New(msg)
}

func (v *TypeChecker) NewEnv(f func() (interface{}, error)) (interface{}, error) {
	prevEnv := v.Context.Env
	v.Context.Env = newTypeEnv(prevEnv)
	r, err := f()
	v.Context.Env = prevEnv
	return r, err
}

type Error struct {
	Message string
	Node    ast.Node
}
