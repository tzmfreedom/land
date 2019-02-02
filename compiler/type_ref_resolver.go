package compiler

import (
	"github.com/tzmfreedom/goland/ast"
	"github.com/tzmfreedom/goland/builtin"
)

type TypeRefResolver struct {
	resolver *builtin.TypeRefResolver
}

func NewTypeRefResolver(classTypes *ast.ClassMap, nameSpaceStore *builtin.NameSpaceStore) *TypeRefResolver {
	return &TypeRefResolver{
		resolver: &builtin.TypeRefResolver{
			ClassTypes: classTypes,
			NameSpaces: nameSpaceStore,
		},
	}
}

func (v *TypeRefResolver) Resolve(n *ast.ClassType) (*ast.ClassType, error) {
	var err error
	if n.SuperClassRef != nil {
		n.SuperClass, err = v.resolver.ResolveType(n.SuperClassRef.Name)
		if err != nil {
			return nil, err
		}
	}

	if n.ImplementClassRefs != nil {
		n.ImplementClasses = make([]*ast.ClassType, len(n.ImplementClassRefs))
		for i, impl := range n.ImplementClassRefs {
			n.ImplementClasses[i], err = v.resolver.ResolveType(impl.Name)
			if err != nil {
				return nil, err
			}
		}
	}

	for _, c := range n.InnerClasses.Data {
		_, err = v.Resolve(c)
		if err != nil {
			return nil, err
		}
	}

	v.resolver.CurrentClass = n
	for _, f := range n.InstanceFields.Data {
		classType, err := f.TypeRef.Accept(v)
		if err != nil {
			return nil, err
		}
		f.Type = classType.(*ast.ClassType)
	}

	for _, f := range n.StaticFields.Data {
		f.Type, err = v.resolver.ResolveType(f.TypeRef.Name)
		if err != nil {
			return nil, err
		}
	}

	for _, m := range n.Constructors {
		if m.ReturnTypeRef != nil {
			retType, err := m.ReturnTypeRef.Accept(v)
			m.ReturnType = retType.(*ast.ClassType)
			if err != nil {
				return nil, err
			}
		}
		for _, param := range m.Parameters {
			param.Accept(v)
		}
		m.Statements.Accept(v)
	}

	for _, methods := range n.InstanceMethods.All() {
		for _, m := range methods {
			if m.ReturnTypeRef != nil {
				retType, err := m.ReturnTypeRef.Accept(v)
				m.ReturnType = retType.(*ast.ClassType)
				if err != nil {
					return nil, err
				}
			}
			for _, param := range m.Parameters {
				param.Accept(v)
			}
			if !n.Interface && !m.IsAbstract() {
				m.Statements.Accept(v)
			}
		}
	}

	for _, methods := range n.StaticMethods.All() {
		for _, m := range methods {
			if m.ReturnTypeRef != nil {
				retType, err := m.ReturnTypeRef.Accept(v)
				m.ReturnType = retType.(*ast.ClassType)
				if err != nil {
					return nil, err
				}
			}
			for _, param := range m.Parameters {
				param.Accept(v)
			}
			m.Statements.Accept(v)
		}
	}
	return n, nil
}

func (v *TypeRefResolver) VisitClassDeclaration(n *ast.ClassDeclaration) (interface{}, error) {
	return ast.VisitClassDeclaration(v, n)
}

func (v *TypeRefResolver) VisitModifier(n *ast.Modifier) (interface{}, error) {
	return ast.VisitModifier(v, n)
}

func (v *TypeRefResolver) VisitAnnotation(n *ast.Annotation) (interface{}, error) {
	return ast.VisitAnnotation(v, n)
}

func (v *TypeRefResolver) VisitInterfaceDeclaration(n *ast.InterfaceDeclaration) (interface{}, error) {
	return ast.VisitInterfaceDeclaration(v, n)
}

func (v *TypeRefResolver) VisitIntegerLiteral(n *ast.IntegerLiteral) (interface{}, error) {
	return ast.VisitIntegerLiteral(v, n)
}

func (v *TypeRefResolver) VisitParameter(n *ast.Parameter) (interface{}, error) {
	classType, err := n.TypeRef.Accept(v)
	n.Type = classType.(*ast.ClassType)
	if err != nil {
		return nil, err
	}
	return nil, nil
}

func (v *TypeRefResolver) VisitArrayAccess(n *ast.ArrayAccess) (interface{}, error) {
	n.Key.Accept(v)
	n.Receiver.Accept(v)
	return nil, nil
}

func (v *TypeRefResolver) VisitBooleanLiteral(n *ast.BooleanLiteral) (interface{}, error) {
	return ast.VisitBooleanLiteral(v, n)
}

func (v *TypeRefResolver) VisitBreak(n *ast.Break) (interface{}, error) {
	return ast.VisitBreak(v, n)
}

func (v *TypeRefResolver) VisitContinue(n *ast.Continue) (interface{}, error) {
	return ast.VisitContinue(v, n)
}

func (v *TypeRefResolver) VisitDml(n *ast.Dml) (interface{}, error) {
	return ast.VisitDml(v, n)
}

func (v *TypeRefResolver) VisitDoubleLiteral(n *ast.DoubleLiteral) (interface{}, error) {
	return ast.VisitDoubleLiteral(v, n)
}

func (v *TypeRefResolver) VisitFieldDeclaration(n *ast.FieldDeclaration) (interface{}, error) {
	classType, err := n.TypeRef.Accept(v)
	if err != nil {
		return nil, err
	}
	n.Type = classType.(*ast.ClassType)
	for _, d := range n.Declarators {
		d.Accept(v)
	}
	return nil, nil
}

func (v *TypeRefResolver) VisitTry(n *ast.Try) (interface{}, error) {
	n.Block.Accept(v)
	n.FinallyBlock.Accept(v)
	for _, c := range n.CatchClause {
		c.Accept(v)
	}
	return ast.VisitTry(v, n)
}

func (v *TypeRefResolver) VisitCatch(n *ast.Catch) (interface{}, error) {
	classType, err := n.TypeRef.Accept(v)
	n.Type = classType.(*ast.ClassType)
	if err != nil {
		return nil, err
	}
	n.Block.Accept(v)
	return nil, nil
}

func (v *TypeRefResolver) VisitFinally(n *ast.Finally) (interface{}, error) {
	return n.Block.Accept(v)
}

func (v *TypeRefResolver) VisitFor(n *ast.For) (interface{}, error) {
	n.Statements.Accept(v)
	n.Control.Accept(v)
	return nil, nil
}

func (v *TypeRefResolver) VisitForControl(n *ast.ForControl) (interface{}, error) {
	n.Expression.Accept(v)
	for _, init := range n.ForInit {
		init.Accept(v)
	}
	for _, update := range n.ForUpdate {
		update.Accept(v)
	}
	return nil, nil
}

func (v *TypeRefResolver) VisitEnhancedForControl(n *ast.EnhancedForControl) (interface{}, error) {
	classType, err := n.TypeRef.Accept(v)
	n.Type = classType.(*ast.ClassType)
	if err != nil {
		return nil, err
	}
	n.Expression.Accept(v)
	return nil, nil
}

func (v *TypeRefResolver) VisitIf(n *ast.If) (interface{}, error) {
	n.Condition.Accept(v)
	n.IfStatement.Accept(v)
	if n.ElseStatement != nil {
		n.ElseStatement.Accept(v)
	}
	return ast.VisitIf(v, n)
}

func (v *TypeRefResolver) VisitMethodDeclaration(n *ast.MethodDeclaration) (interface{}, error) {
	return ast.VisitMethodDeclaration(v, n)
}

func (v *TypeRefResolver) VisitMethodInvocation(n *ast.MethodInvocation) (interface{}, error) {
	n.NameOrExpression.Accept(v)
	for _, param := range n.Parameters {
		param.Accept(v)
	}
	return nil, nil
}

func (v *TypeRefResolver) VisitNew(n *ast.New) (interface{}, error) {
	classType, err := n.TypeRef.Accept(v)
	n.Type = classType.(*ast.ClassType)
	if err != nil {
		return nil, err
	}
	for _, p := range n.Parameters {
		p.Accept(v)
	}
	return nil, nil
}

func (v *TypeRefResolver) VisitNullLiteral(n *ast.NullLiteral) (interface{}, error) {
	return ast.VisitNullLiteral(v, n)
}

func (v *TypeRefResolver) VisitUnaryOperator(n *ast.UnaryOperator) (interface{}, error) {
	n.Expression.Accept(v)
	return nil, nil
}

func (v *TypeRefResolver) VisitBinaryOperator(n *ast.BinaryOperator) (interface{}, error) {
	n.Left.Accept(v)
	n.Right.Accept(v)
	return nil, nil
}

func (v *TypeRefResolver) VisitReturn(n *ast.Return) (interface{}, error) {
	if n.Expression != nil {
		n.Expression.Accept(v)
	}
	return nil, nil
}

func (v *TypeRefResolver) VisitThrow(n *ast.Throw) (interface{}, error) {
	n.Expression.Accept(v)
	return nil, nil
}

func (v *TypeRefResolver) VisitSoql(n *ast.Soql) (interface{}, error) {
	return ast.VisitSoql(v, n)
}

func (v *TypeRefResolver) VisitSosl(n *ast.Sosl) (interface{}, error) {
	return ast.VisitSosl(v, n)
}

func (v *TypeRefResolver) VisitStringLiteral(n *ast.StringLiteral) (interface{}, error) {
	return ast.VisitStringLiteral(v, n)
}

func (v *TypeRefResolver) VisitSwitch(n *ast.Switch) (interface{}, error) {
	n.Expression.Accept(v)
	for _, w := range n.WhenStatements {
		w.Accept(v)
	}
	if n.ElseStatement != nil {
		n.ElseStatement.Accept(v)
	}
	return nil, nil
}

func (v *TypeRefResolver) VisitTrigger(n *ast.Trigger) (interface{}, error) {
	n.Statements.Accept(v)
	return nil, nil
}

func (v *TypeRefResolver) VisitTriggerTiming(n *ast.TriggerTiming) (interface{}, error) {
	return ast.VisitTriggerTiming(v, n)
}

func (v *TypeRefResolver) VisitVariableDeclaration(n *ast.VariableDeclaration) (interface{}, error) {
	classType, err := n.TypeRef.Accept(v)
	n.Type = classType.(*ast.ClassType)
	if err != nil {
		return nil, err
	}
	for _, d := range n.Declarators {
		d.Accept(v)
	}
	return nil, nil
}

func (v *TypeRefResolver) VisitVariableDeclarator(n *ast.VariableDeclarator) (interface{}, error) {
	n.Expression.Accept(v)
	return nil, nil
}

func (v *TypeRefResolver) VisitWhen(n *ast.When) (interface{}, error) {
	n.Statements.Accept(v)
	for _, c := range n.Condition {
		c.Accept(v)
	}
	return nil, nil
}

func (v *TypeRefResolver) VisitWhenType(n *ast.WhenType) (interface{}, error) {
	classType, err := n.TypeRef.Accept(v)
	if err != nil {
		return nil, err
	}
	n.Type = classType.(*ast.ClassType)
	return nil, nil
}

func (v *TypeRefResolver) VisitWhile(n *ast.While) (interface{}, error) {
	n.Statements.Accept(v)
	n.Condition.Accept(v)
	return nil, nil
}

func (v *TypeRefResolver) VisitNothingStatement(n *ast.NothingStatement) (interface{}, error) {
	return ast.VisitNothingStatement(v, n)
}

func (v *TypeRefResolver) VisitCastExpression(n *ast.CastExpression) (interface{}, error) {
	classType, err := n.CastTypeRef.Accept(v)
	if err != nil {
		return nil, err
	}
	n.CastType = classType.(*ast.ClassType)
	n.Expression.Accept(v)
	return nil, nil
}

func (v *TypeRefResolver) VisitFieldAccess(n *ast.FieldAccess) (interface{}, error) {
	n.Expression.Accept(v)
	return nil, nil
}

func (v *TypeRefResolver) VisitType(n *ast.TypeRef) (interface{}, error) {
	for n.Dimmension > 0 {
		name := n.Name
		params := n.Parameters
		n.Name = []string{"List"}
		n.Parameters = []*ast.TypeRef{
			{
				Name:       name,
				Parameters: params,
			},
		}
		n.Dimmension--
	}

	classType, err := v.resolver.ResolveType(n.Name)
	if err != nil {
		return nil, err
	}
	if classType == builtin.ListType || classType == builtin.MapType {
		paramTypes := make([]*ast.ClassType, len(n.Parameters))
		for i, param := range n.Parameters {
			paramType, err := param.Accept(v)
			if err != nil {
				return nil, err
			}
			paramTypes[i] = paramType.(*ast.ClassType)
		}
		return &ast.ClassType{
			Name:            classType.Name,
			Constructors:    classType.Constructors,
			InstanceFields:  classType.InstanceFields,
			StaticFields:    classType.StaticFields,
			InstanceMethods: classType.InstanceMethods,
			StaticMethods:   classType.StaticMethods,
			Generics:        paramTypes,
			ToString:        classType.ToString,
		}, nil
	}
	return classType, nil
}

func (v *TypeRefResolver) VisitBlock(n *ast.Block) (interface{}, error) {
	for _, stmt := range n.Statements {
		stmt.Accept(v)
	}
	return nil, nil
}

func (v *TypeRefResolver) VisitGetterSetter(n *ast.GetterSetter) (interface{}, error) {
	return ast.VisitGetterSetter(v, n)
}

func (v *TypeRefResolver) VisitPropertyDeclaration(n *ast.PropertyDeclaration) (interface{}, error) {
	classType, err := n.TypeRef.Accept(v)
	n.Type = classType.(*ast.ClassType)
	if err != nil {
		return nil, err
	}
	return nil, nil
}

func (v *TypeRefResolver) VisitArrayInitializer(n *ast.ArrayInitializer) (interface{}, error) {
	for _, n := range n.Initializers {
		n.Accept(v)
	}
	return nil, nil
}

func (v *TypeRefResolver) VisitArrayCreator(n *ast.ArrayCreator) (interface{}, error) {
	for _, exp := range n.Expressions {
		exp.Accept(v)
	}
	n.ArrayInitializer.Accept(v)
	return nil, nil
}

func (v *TypeRefResolver) VisitSoqlBindVariable(n *ast.SoqlBindVariable) (interface{}, error) {
	return ast.VisitSoqlBindVariable(v, n)
}

func (v *TypeRefResolver) VisitTernalyExpression(n *ast.TernalyExpression) (interface{}, error) {
	n.Condition.Accept(v)
	n.FalseExpression.Accept(v)
	n.TrueExpression.Accept(v)
	return nil, nil
}

func (v *TypeRefResolver) VisitMapCreator(n *ast.MapCreator) (interface{}, error) {
	return ast.VisitMapCreator(v, n)
}

func (v *TypeRefResolver) VisitSetCreator(n *ast.SetCreator) (interface{}, error) {
	return ast.VisitSetCreator(v, n)
}

func (v *TypeRefResolver) VisitName(n *ast.Name) (interface{}, error) {
	return ast.VisitName(v, n)
}

func (v *TypeRefResolver) VisitConstructorDeclaration(n *ast.ConstructorDeclaration) (interface{}, error) {
	return ast.VisitConstructorDeclaration(v, n)
}
