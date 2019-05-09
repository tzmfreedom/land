package compiler

import (
	"github.com/tzmfreedom/land/ast"
	"github.com/tzmfreedom/land/builtin"
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
	if n.SuperClassRef != nil {
		superClass, err := n.SuperClassRef.Accept(v)
		if err != nil {
			return nil, err
		}
		n.SuperClass = superClass.(*ast.ClassType)
	}

	if n.ImplementClassRefs != nil {
		n.ImplementClasses = make([]*ast.ClassType, len(n.ImplementClassRefs))
		for i, impl := range n.ImplementClassRefs {
			implementClass, err := impl.Accept(v)
			if err != nil {
				return nil, err
			}
			n.ImplementClasses[i] = implementClass.(*ast.ClassType)
		}
	}

	for _, c := range n.InnerClasses.Data {
		_, err := v.Resolve(c)
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
		if f.Expression != nil {
			f.Expression.Accept(v)
		}
	}

	for _, f := range n.StaticFields.Data {
		classType, err := f.TypeRef.Accept(v)
		if err != nil {
			return nil, err
		}
		f.Type = classType.(*ast.ClassType)
		if f.Expression != nil {
			f.Expression.Accept(v)
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
				_, err := param.Accept(v)
				if err != nil {
					return nil, err
				}
			}
			if !n.Interface && !m.IsAbstract() {
				_, err := m.Statements.Accept(v)
				if err != nil {
					return nil, err
				}
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
				_, err := param.Accept(v)
				if err != nil {
					return nil, err
				}
			}
			_, err := m.Statements.Accept(v)
			if err != nil {
				return nil, err
			}
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
	_, err := n.Key.Accept(v)
	if err != nil {
		return nil, err
	}
	return n.Receiver.Accept(v)
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
		_, err := d.Accept(v)
		if err != nil {
			return nil, err
		}
	}
	return nil, nil
}

func (v *TypeRefResolver) VisitTry(n *ast.Try) (interface{}, error) {
	_, err := n.Block.Accept(v)
	if err != nil {
		return nil, err
	}
	_, err = n.FinallyBlock.Accept(v)
	if err != nil {
		return nil, err
	}
	for _, c := range n.CatchClause {
		_, err := c.Accept(v)
		if err != nil {
			return nil, err
		}
	}
	return ast.VisitTry(v, n)
}

func (v *TypeRefResolver) VisitCatch(n *ast.Catch) (interface{}, error) {
	classType, err := n.TypeRef.Accept(v)
	n.Type = classType.(*ast.ClassType)
	if err != nil {
		return nil, err
	}
	return n.Block.Accept(v)
}

func (v *TypeRefResolver) VisitFinally(n *ast.Finally) (interface{}, error) {
	return n.Block.Accept(v)
}

func (v *TypeRefResolver) VisitFor(n *ast.For) (interface{}, error) {
	_, err := n.Statements.Accept(v)
	if err != nil {
		return nil, err
	}
	return n.Control.Accept(v)
}

func (v *TypeRefResolver) VisitForControl(n *ast.ForControl) (interface{}, error) {
	_, err := n.Expression.Accept(v)
	if err != nil {
		return nil, err
	}
	for _, init := range n.ForInit {
		_, err := init.Accept(v)
		if err != nil {
			return nil, err
		}
	}
	for _, update := range n.ForUpdate {
		_, err := update.Accept(v)
		if err != nil {
			return nil, err
		}
	}
	return nil, nil
}

func (v *TypeRefResolver) VisitEnhancedForControl(n *ast.EnhancedForControl) (interface{}, error) {
	classType, err := n.TypeRef.Accept(v)
	if err != nil {
		return nil, err
	}
	n.Type = classType.(*ast.ClassType)
	return n.Expression.Accept(v)
}

func (v *TypeRefResolver) VisitIf(n *ast.If) (interface{}, error) {
	_, err := n.Condition.Accept(v)
	if err != nil {
		return nil, err
	}
	_, err = n.IfStatement.Accept(v)
	if err != nil {
		return nil, err
	}
	if n.ElseStatement != nil {
		_, err = n.ElseStatement.Accept(v)
		if err != nil {
			return nil, err
		}
	}
	return ast.VisitIf(v, n)
}

func (v *TypeRefResolver) VisitMethodDeclaration(n *ast.MethodDeclaration) (interface{}, error) {
	return ast.VisitMethodDeclaration(v, n)
}

func (v *TypeRefResolver) VisitMethodInvocation(n *ast.MethodInvocation) (interface{}, error) {
	_, err := n.NameOrExpression.Accept(v)
	if err != nil {
		return nil, err
	}
	for _, param := range n.Parameters {
		_, err = param.Accept(v)
		if err != nil {
			return nil, err
		}
	}
	return nil, nil
}

func (v *TypeRefResolver) VisitNew(n *ast.New) (interface{}, error) {
	classType, err := n.TypeRef.Accept(v)
	if err != nil {
		return nil, err
	}
	n.Type = classType.(*ast.ClassType)
	for _, p := range n.Parameters {
		_, err := p.Accept(v)
		if err != nil {
			return nil, err
		}
	}
	return nil, nil
}

func (v *TypeRefResolver) VisitNullLiteral(n *ast.NullLiteral) (interface{}, error) {
	return ast.VisitNullLiteral(v, n)
}

func (v *TypeRefResolver) VisitUnaryOperator(n *ast.UnaryOperator) (interface{}, error) {
	return n.Expression.Accept(v)
}

func (v *TypeRefResolver) VisitBinaryOperator(n *ast.BinaryOperator) (interface{}, error) {
	_, err := n.Left.Accept(v)
	if err != nil {
		return nil, err
	}
	return n.Right.Accept(v)
}

func (v *TypeRefResolver) VisitInstanceofOperator(n *ast.InstanceofOperator) (interface{}, error) {
	_, err := n.Expression.Accept(v)
	if err != nil {
		return nil, err
	}
	classType, err := n.TypeRef.Accept(v)
	if err != nil {
		return nil, err
	}
	n.Type = classType.(*ast.ClassType)
	return nil, nil
}

func (v *TypeRefResolver) VisitReturn(n *ast.Return) (interface{}, error) {
	if n.Expression != nil {
		return n.Expression.Accept(v)
	}
	return nil, nil
}

func (v *TypeRefResolver) VisitThrow(n *ast.Throw) (interface{}, error) {
	return n.Expression.Accept(v)
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
	_, err := n.Expression.Accept(v)
	if err != nil {
		return nil, err
	}
	for _, w := range n.WhenStatements {
		_, err := w.Accept(v)
		if err != nil {
			return nil, err
		}
	}
	if n.ElseStatement != nil {
		return n.ElseStatement.Accept(v)
	}
	return nil, nil
}

func (v *TypeRefResolver) VisitTrigger(n *ast.Trigger) (interface{}, error) {
	return n.Statements.Accept(v)
}

func (v *TypeRefResolver) VisitTriggerTiming(n *ast.TriggerTiming) (interface{}, error) {
	return ast.VisitTriggerTiming(v, n)
}

func (v *TypeRefResolver) VisitVariableDeclaration(n *ast.VariableDeclaration) (interface{}, error) {
	classType, err := n.TypeRef.Accept(v)
	if err != nil {
		return nil, err
	}
	n.Type = classType.(*ast.ClassType)
	for _, d := range n.Declarators {
		_, err := d.Accept(v)
		if err != nil {
			return nil, err
		}
	}
	return nil, nil
}

func (v *TypeRefResolver) VisitVariableDeclarator(n *ast.VariableDeclarator) (interface{}, error) {
	if n.Expression != nil {
		return n.Expression.Accept(v)
	}
	return nil, nil
}

func (v *TypeRefResolver) VisitWhen(n *ast.When) (interface{}, error) {
	_, err := n.Statements.Accept(v)
	if err != nil {
		return nil, err
	}
	for _, c := range n.Condition {
		_, err := c.Accept(v)
		if err != nil {
			return nil, err
		}
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
	_, err := n.Statements.Accept(v)
	if err != nil {
		return nil, err
	}
	return n.Condition.Accept(v)
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
	return n.Expression.Accept(v)
}

func (v *TypeRefResolver) VisitFieldAccess(n *ast.FieldAccess) (interface{}, error) {
	return n.Expression.Accept(v)
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
	if classType.IsGenerics() {
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
		_, err := stmt.Accept(v)
		if err != nil {
			return nil, err
		}
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
		_, err := n.Accept(v)
		if err != nil {
			return nil, err
		}
	}
	return nil, nil
}

func (v *TypeRefResolver) VisitArrayCreator(n *ast.ArrayCreator) (interface{}, error) {
	for _, exp := range n.Expressions {
		_, err := exp.Accept(v)
		if err != nil {
			return nil, err
		}
	}
	return n.ArrayInitializer.Accept(v)
}

func (v *TypeRefResolver) VisitSoqlBindVariable(n *ast.SoqlBindVariable) (interface{}, error) {
	return ast.VisitSoqlBindVariable(v, n)
}

func (v *TypeRefResolver) VisitTernalyExpression(n *ast.TernalyExpression) (interface{}, error) {
	_, err := n.Condition.Accept(v)
	if err != nil {
		return nil, err
	}
	_, err = n.FalseExpression.Accept(v)
	if err != nil {
		return nil, err
	}
	return n.TrueExpression.Accept(v)
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
