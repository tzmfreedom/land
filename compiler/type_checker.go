package compiler

import (
	"errors"
	"fmt"

	"strings"

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

func (v *TypeChecker) VisitClassType(n *ast.ClassType) (interface{}, error) {
	v.Context.CurrentClass = n
	if n.StaticFields != nil {
		for _, f := range n.StaticFields.Data {
			e, err := f.Expression.Accept(v)
			if err != nil {
				return nil, err
			}
			if !builtin.Equals(f.Type, e.(*ast.ClassType)) {
				v.AddError(fmt.Sprintf("expression <%s> does not match <%s>", e.(*ast.ClassType).String(), f.Type.String()), f.Expression)
			}
		}
	}

	if n.InstanceFields != nil {
		for _, f := range n.InstanceFields.Data {
			e, err := f.Expression.Accept(v)
			if err != nil {
				return nil, err
			}
			if e == nil {
				continue
			}
			if !builtin.Equals(f.Type, e.(*ast.ClassType)) {
				v.AddError(fmt.Sprintf("expression <%s> does not match <%s>", e.(*ast.ClassType).String(), f.Type.String()), f.Expression)
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

	if n.ImplementClassRefs != nil {
		for _, impl := range n.ImplementClassRefs {
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
	panic("not pass")
	return nil, nil
}

func (v *TypeChecker) VisitArrayAccess(n *ast.ArrayAccess) (interface{}, error) {
	k, err := n.Receiver.Accept(v)
	if err != nil {
		return nil, err
	}
	klass := k.(*ast.ClassType)
	if !klass.IsGenerics() {
		v.AddError("receiver should be list or map", n.Key)
		return nil, nil
	}
	t, err := n.Key.Accept(v)
	if err != nil {
		return nil, err
	}
	generics := klass.Generics
	if len(generics) == 0 {
		return nil, v.compileError("generics is not specified", n)
	}
	if klass.Name == "Map" {
		if t != builtin.StringType {
			v.AddError(fmt.Sprintf("map key <%v> must be String", t.(*ast.ClassType).String()), n.Key)
		}
		return generics[1], nil
	}
	if t != builtin.IntegerType {
		v.AddError(fmt.Sprintf("list key <%v> must be Integer", t.(*ast.ClassType).String()), n.Key)
	}
	// TODO: implement set
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
	t, err := n.Expression.Accept(v)
	if err != nil {
		return nil, err
	}
	if t != builtin.ListType {
// TODO: impl
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
	v.Context.Env.Set(n.Identifier, n.Type) // TODO: append scope
	_, err := n.Block.Accept(v)
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
			v.AddError(fmt.Sprintf("condition <%s> must be Boolean expression", t.(*ast.ClassType).String()), n.Expression)
		}
	}
	return nil, nil
}

func (v *TypeChecker) VisitEnhancedForControl(n *ast.EnhancedForControl) (interface{}, error) {
	declClassType := n.Type
	exp, err := n.Expression.Accept(v)
	if err != nil {
		return nil, err
	}
	expClassType := exp.(*ast.ClassType)
	v.Context.Env.Set(n.VariableDeclaratorId, declClassType)

	if expClassType.Name != "List" && expClassType.Name != "Set" {
		v.AddError(fmt.Sprintf("expression <%s> must be List or Set expression", expClassType.Name), n)
		return nil, nil
	}

	genericsType := expClassType.Generics[0]
	if !builtin.Equals(declClassType, genericsType) {
		v.AddError(fmt.Sprintf("expression <%s> must be <%s> expression", declClassType.Name, expClassType.Name), n)
	}
	return nil, nil
}

func (v *TypeChecker) VisitIf(n *ast.If) (interface{}, error) {
	t, err := n.Condition.Accept(v)
	if err != nil {
		return nil, err
	}
	if t != builtin.BooleanType {
		v.AddError(fmt.Sprintf("condition <%s> must be Boolean expression", t.(*ast.ClassType).String()), n.Condition)
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
		env.Set(param.Name, param.Type)
	}
	r, err := n.Statements.Accept(v)
	if err != nil {
		v.Context.CurrentMethod = nil
		return nil, err
	}
	if n.ReturnType != nil && r == nil {
		retType, err := n.ReturnType.Accept(v)
		if err != nil {
			return nil, err
		}
		v.AddError(fmt.Sprintf("return type <void> does not match %v", retType.(*ast.ClassType).String()), n)
	}

	v.Context.CurrentMethod = nil
	return nil, nil
}

func (v *TypeChecker) VisitMethodInvocation(n *ast.MethodInvocation) (interface{}, error) {
	resolver := NewTypeResolver(v.Context)

	nameOrExp := n.NameOrExpression
	types := make([]*ast.ClassType, len(n.Parameters))
	for i, p := range n.Parameters {
		t, err := p.Accept(v)
		if err != nil {
			return nil, err
		}
		types[i] = t.(*ast.ClassType)
	}
	if name, ok := nameOrExp.(*ast.Name); ok {
		// TODO: implement
		if name.Value[0] == "_Debugger" {
			return nil, nil
		}
		receiverType, method, err := resolver.ResolveMethod(name.Value, types)
		if err != nil {
			return nil, v.compileError(err.Error(), n)
		}
		if err := v.checkAddError(n, receiverType, method); err != nil {
			return nil, err
		}
		if method.ReturnType != nil {
			retType := method.ReturnType
			if retType == builtin.T1type {
				return receiverType.Generics[0], nil
			}
			if retType == builtin.T2type {
				return receiverType.Generics[1], nil
			}
			return method.ReturnType, nil
		}
	} else if fieldAccess, ok := nameOrExp.(*ast.FieldAccess); ok {
		classType, err := fieldAccess.Expression.Accept(v)
		if err != nil {
			return nil, err
		}
		receiverType, method, err := FindInstanceMethod(
			classType.(*ast.ClassType),
			fieldAccess.FieldName,
			types,
			MODIFIER_PUBLIC_ONLY,
		)
		if err != nil {
			return nil, v.compileError(err.Error(), n)
		}
		if err := v.checkAddError(n, receiverType, method); err != nil {
			return nil, err
		}
		if method.ReturnType != nil {
			// TODO: duplicate code
			retType := method.ReturnType
			if retType == builtin.T1type {
				return receiverType.Generics[0], nil
			}
			if retType == builtin.T2type {
				return receiverType.Generics[1], nil
			}
			return method.ReturnType, nil
		}
	}
	return nil, nil
}

func (v *TypeChecker) VisitNew(n *ast.New) (interface{}, error) {
	params := make([]*ast.ClassType, len(n.Parameters))
	for i, p := range n.Parameters {
		param, err := p.Accept(v)
		if err != nil {
			return nil, err
		}
		params[i] = param.(*ast.ClassType)
	}
	typeResolver := NewTypeResolver(v.Context)
	classType := n.Type

	if classType.Name == "List" {
		elemClass := classType.Generics[0]
		for _, record := range n.Init.Records {
			r, err := record.Accept(v)
			paramElemClass := r.(*ast.ClassType)
			if err != nil {
				return nil, err
			}
			if builtin.Equals(paramElemClass, elemClass) {
				v.AddError(fmt.Sprintf("initialization is not match type %s != %s", elemClass.Name, paramElemClass.Name), n)
			}
		}
	}

	if classType.Name == "Map" {
		keyClass := classType.Generics[0]
		valueClass := classType.Generics[1]
		for key, value := range n.Init.Values {
			r, err := key.Accept(v)
			paramKeyClass := r.(*ast.ClassType)
			if err != nil {
				return nil, err
			}
			if builtin.Equals(paramKeyClass, keyClass) {
				v.AddError(fmt.Sprintf("initialization is not match type %s != %s", keyClass.Name, paramKeyClass.Name), n)
			}
			r, err = value.Accept(v)
			paramValueClass := r.(*ast.ClassType)
			if err != nil {
				return nil, err
			}
			if builtin.Equals(paramValueClass, valueClass) {
				v.AddError(fmt.Sprintf("initialization is not match type %s != %s", valueClass.Name, paramValueClass.Name), n)
			}
		}
	}

	if classType.IsAbstract() || classType.Constructors == nil {
		v.AddError(fmt.Sprintf("Type cannot be constructed: %s", classType.Name), n)
		return n.Type, nil
	}
	if !classType.HasConstructor() {
		return n.Type, nil
	}
	_, method, err := typeResolver.SearchConstructor(classType, params)
	if err != nil {
		return nil, err
	}
	if method == nil {
		v.AddError(fmt.Sprintf("constructor <%s> not found", classType.Name), n)
		return n.Type, nil
	}
	// TODO: for protected impl
	if method.IsPrivate() && v.Context.CurrentClass != classType {
		v.AddError(fmt.Sprintf("constructor <%s> not found", classType.Name), n)
	}
	return n.Type, nil
}

func (v *TypeChecker) VisitNullLiteral(n *ast.NullLiteral) (interface{}, error) {
	return builtin.NullType, nil
}

func (v *TypeChecker) VisitUnaryOperator(n *ast.UnaryOperator) (interface{}, error) {
	t, err := n.Expression.Accept(v)
	if err != nil {
		return nil, err
	}
	if t != builtin.IntegerType {
		v.AddError(fmt.Sprintf("expression <%s> must be Integer", t.(*ast.ClassType).String()), n.Expression)
	}
	return nil, nil
}

func (v *TypeChecker) VisitBinaryOperator(n *ast.BinaryOperator) (interface{}, error) {
	r, err := n.Right.Accept(v)
	if err != nil {
		return nil, err
	}
	if n.Op == "=" ||
		n.Op == "=+" ||
		n.Op == "=-" ||
		n.Op == "=*" ||
		n.Op == "=/" {

		var l *ast.ClassType
		resolver := NewTypeResolver(v.Context)
		switch leftNode := n.Left.(type) {
		case *ast.Name:
			left, err := resolver.ResolveVariable(leftNode.Value, true)
			if err != nil {
				return nil, v.compileError(err.Error(), n)
			}
			l = left
		case *ast.FieldAccess:
			classType, err := leftNode.Expression.Accept(v)
			if err != nil {
				return nil, err
			}
			f, err := FindInstanceField(classType.(*ast.ClassType), leftNode.FieldName, MODIFIER_PUBLIC_ONLY, true)
			if err != nil {
				return nil, err
			}
			l = f.Type
		case *ast.ArrayAccess:
			left, err := leftNode.Accept(v)
			if err != nil {
				return nil, err
			}
			l = left.(*ast.ClassType)
		}
		if r != nil && !builtin.Equals(l, r.(*ast.ClassType)) {
			v.AddError(fmt.Sprintf("expression <%s> does not match <%s>", l.String(), r.(*ast.ClassType).String()), n.Left)
		}
		return l, nil
	} else {
		l, err := n.Left.Accept(v)
		if err != nil {
			return nil, err
		}
		if n.Op == "+" {
			if l != builtin.IntegerType && l != builtin.StringType && l != builtin.DoubleType {
				v.AddError(fmt.Sprintf("expression <%s> must be Integer, String or Double", l.(*ast.ClassType).String()), n.Left)
			}
			if (l == builtin.StringType || r == builtin.StringType) && l != r {
				v.AddError(fmt.Sprintf("expression <%s> does not match <%s>", l.(*ast.ClassType).String(), r.(*ast.ClassType).String()), n.Left)
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
				v.AddError(fmt.Sprintf("expression <%s> must be Integer or Double", l.(*ast.ClassType).String()), n.Left)
			} else if r != builtin.IntegerType && r != builtin.DoubleType {
				v.AddError(fmt.Sprintf("expression <%s> must be Integer or Double", r.(*ast.ClassType).String()), n.Right)
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
			exp, err := n.Expression.Accept(v)
			if err != nil {
				return nil, err
			}
			v.AddError(fmt.Sprintf("return type <%s> does not match void", exp.(*ast.ClassType).String()), n.Expression)
		}
		return nil, nil
	}

	retType := v.Context.CurrentMethod.ReturnType
	if n.Expression == nil {
		v.AddError(fmt.Sprintf("return type <void> does not match %v", retType.String()), n.Expression)
		return nil, nil
	}
	exp, err := n.Expression.Accept(v)
	if err != nil {
		return nil, err
	}
	if !builtin.Equals(retType, exp.(*ast.ClassType)) {
		v.AddError(fmt.Sprintf("return type <%s> does not match %v", exp.(*ast.ClassType).String(), retType.String()), n.Expression)
	}
	return exp, nil
}

func (v *TypeChecker) VisitThrow(n *ast.Throw) (interface{}, error) {
	r, err := n.Expression.Accept(v)
	if err != nil {
		return nil, err
	}
	// Check Subclass of Exception
	baseClass := r.(*ast.ClassType)
	super := baseClass.SuperClass
	if super == nil {
		v.AddError(fmt.Sprintf("Throw expression must be of type exception: %s", baseClass.Name), n)
	} else {
		if err != nil {
			v.AddError(err.Error(), n)
		} else if super != builtin.ExceptionType {
			v.AddError(fmt.Sprintf("Throw expression must be of type exception: %s", baseClass.Name), n)
		}
	}
	return nil, nil
}

func (v *TypeChecker) VisitSoql(n *ast.Soql) (interface{}, error) {
	resolver := NewTypeResolver(v.Context)
	t, err := resolver.ResolveType([]string{n.FromObject})
	if err != nil {
		return nil, v.compileError(err.Error(), n)
	}
	return &ast.ClassType{
		Name:     "List",
		Generics: []*ast.ClassType{t},
	}, nil
}

func (v *TypeChecker) VisitSosl(n *ast.Sosl) (interface{}, error) {
	return ast.VisitSosl(v, n)
}

func (v *TypeChecker) VisitStringLiteral(n *ast.StringLiteral) (interface{}, error) {
	return builtin.StringType, nil
}

func (v *TypeChecker) VisitSwitch(n *ast.Switch) (interface{}, error) {
	exp, err := n.Expression.Accept(v)
	if err != nil {
		return nil, err
	}
	for _, w := range n.WhenStatements {
		t, err := w.Accept(v)
		if err != nil {
			return nil, err
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
	for _, d := range n.Declarators {
		t, err := d.Accept(v)
		if err != nil {
			return nil, err
		}
		if _, ok := v.Context.Env.Get(d.Name); ok {
			v.AddError(fmt.Sprintf("variable declaration is duplicated <%s>", d.Name), n)
			continue
		}
		v.Context.Env.Set(d.Name, n.Type)
		if !builtin.Equals(n.Type, t.(*ast.ClassType)) {
			v.AddError(fmt.Sprintf("expression <%s> does not match <%s>", n.Type.String(), t.(*ast.ClassType).String()), n)
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
	t, err := n.Condition.Accept(v)
	if err != nil {
		return nil, err
	}
	if t != builtin.BooleanType {
		v.AddError(fmt.Sprintf("condition <%s> must be Boolean expression", t.(*ast.ClassType).String()), n.Condition)
	}
	return n.Statements.Accept(v)
}

func (v *TypeChecker) VisitNothingStatement(n *ast.NothingStatement) (interface{}, error) {
	return nil, nil
}

func (v *TypeChecker) VisitCastExpression(n *ast.CastExpression) (interface{}, error) {
	exp, err := n.Expression.Accept(v)
	if err != nil {
		return nil, err
	}
	expClassType := exp.(*ast.ClassType)
	if !builtin.Equals(expClassType, n.CastType) {
		v.AddError(fmt.Sprintf("invalid cast (%s)%s", n.CastType.Name, expClassType.Name), n)
	}
	return n.CastType, nil
}

func (v *TypeChecker) VisitFieldAccess(n *ast.FieldAccess) (interface{}, error) {
	classType, err := n.Expression.Accept(v)
	if err != nil {
		return nil, err
	}
	f, ok := classType.(*ast.ClassType).InstanceFields.Get(n.FieldName)
	if !ok {
		return nil, v.compileError(fmt.Sprintf("field <%s> does not exist", n.FieldName), n)
	}
	return f.Type, nil
}

func (v *TypeChecker) VisitType(n *ast.TypeRef) (interface{}, error) {
	resolver := NewTypeResolver(v.Context)
	return resolver.ConvertType(n)
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
	c, err := n.Condition.Accept(v)
	if err != nil {
		return nil, err
	}
	if c != builtin.BooleanType {
		v.AddError(fmt.Sprintf("condition <%s> must be Boolean expression", c.(*ast.ClassType).String()), n.Condition)
	}
	t, err := n.TrueExpression.Accept(v)
	if err != nil {
		return nil, err
	}
	f, err := n.FalseExpression.Accept(v)
	if err != nil {
		return nil, err
	}
	if !builtin.Equals(t.(*ast.ClassType), f.(*ast.ClassType)) {
		v.AddError(fmt.Sprintf("expression does not match %s != %s", t.(*ast.ClassType).String(), f.(*ast.ClassType).String()), n.TrueExpression)
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
	resolver := NewTypeResolver(v.Context)
	classType, err := resolver.ResolveVariable(n.Value, false)
	if err != nil {
		return nil, v.compileError(err.Error(), n)
	}
	return classType, nil
}

func (v *TypeChecker) VisitConstructorDeclaration(n *ast.ConstructorDeclaration) (interface{}, error) {
	env := newTypeEnv(nil)
	v.Context.Env = env
	for _, param := range n.Parameters {
		env.Set(param.Name, param.Type)
	}
	n.Statements.Accept(v)
	return nil, nil
}

func (v *TypeChecker) VisitMethod(n *ast.Method) (interface{}, error) {
	if n.Parent.IsInterface() {
		return nil, nil
	}
	if n.Parent.IsAbstract() {
		return nil, nil
	}
	v.Context.CurrentMethod = n
	env := newTypeEnv(nil)
	v.Context.Env = env
	classType := n.Parent
	v.Context.Env.Set("this", classType)
	for _, param := range n.Parameters {
		env.Set(param.Name, param.Type)
	}
	r, err := n.Statements.Accept(v)
	if err != nil {
		v.Context.CurrentMethod = nil
		return nil, err
	}
	if n.ReturnType != nil && r == nil {
		v.AddError(fmt.Sprintf("return type <void> does not match %v", n.ReturnType.String()), n.Statements)
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

var notImplementAddErrorError = errors.New("not implement add error")

func (v *TypeChecker) checkAddError(n *ast.MethodInvocation, receiver *ast.ClassType, method *ast.Method) error {
	if strings.ToLower(method.Name) != "adderror" {
		return nil
	}
	if isTypeSObjectField(receiver) {
		resolver := NewTypeResolver(v.Context)
		switch exp := n.NameOrExpression.(type) {
		case *ast.Name:
			if len(exp.Value) < 3 {
				return notImplementAddErrorError
			}
			classType, err := resolver.ResolveVariable(exp.Value[:len(exp.Value)-2], false)
			if err != nil {
				return err
			}
			if classType.SuperClass != nil && classType.SuperClass == builtin.SObjectType {
				return nil
			}
			return notImplementAddErrorError
		case *ast.FieldAccess:
			switch fieldExp := exp.Expression.(type) {
			case *ast.Name:
				classType, err := resolver.ResolveVariable(fieldExp.Value[:len(fieldExp.Value)-1], false)
				if err != nil {
					return err
				}
				if classType.SuperClass != nil && classType.SuperClass == builtin.SObjectType {
					return nil
				}
			case *ast.FieldAccess:
				c, err := fieldExp.Expression.Accept(v)
				if err != nil {
					return err
				}
				classType := c.(*ast.ClassType)
				if classType.SuperClass != nil && classType.SuperClass == builtin.SObjectType {
					return nil
				}
			}
			return notImplementAddErrorError
		}
		panic("not pass")
	}
	return nil
}

func isTypeSObjectField(classType *ast.ClassType) bool {
	return classType == builtin.IntegerType ||
		classType == builtin.StringType ||
		classType == builtin.BooleanType ||
		classType == builtin.DateType ||
		classType == builtin.DoubleType
}
