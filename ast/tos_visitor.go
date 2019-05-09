package ast

import (
	"fmt"
	"strings"
)

type TosVisitor struct {
	Indent int
}

func (v *TosVisitor) AddIndent(f func()) {
	v.Indent += 4
	f()
	v.Indent -= 4
}

func (v *TosVisitor) withIndent(src string) string {
	return strings.Repeat(" ", v.Indent) + src
}

func (v *TosVisitor) VisitClassDeclaration(n *ClassDeclaration) (interface{}, error) {
	annotations := make([]string, len(n.Annotations))
	for i, a := range n.Annotations {
		r, err := a.Accept(v)
		if err != nil {
			return nil, err
		}
		annotations[i] = r.(string)
	}
	annotationStr := ""
	if len(annotations) != 0 {
		annotationStr = fmt.Sprintf("%s\n", strings.Join(annotations, "\n"))
	}
	modifiers := make([]string, len(n.Modifiers))
	for i, m := range n.Modifiers {
		r, err := m.Accept(v)
		if err != nil {
			return nil, err
		}
		modifiers[i] = r.(string)
	}
	declarations := make([]string, len(n.Declarations))
	v.AddIndent(func() {
		for i, d := range n.Declarations {
			r, err := d.Accept(v)
			if err != nil {
				panic(err)
			}
			declarations[i] = r.(string)
		}
	})
	super := ""
	if n.SuperClassRef != nil {
		r, err := n.SuperClassRef.Accept(v)
		if err != nil {
			return nil, err
		}
		super = "extends " + r.(string)
	}
	implements := make([]string, len(n.ImplementClassRefs))
	for i, impl := range n.ImplementClassRefs {
		r, err := impl.Accept(v)
		if err != nil {
			return nil, err
		}
		implements[i] = r.(string)
	}
	implString := ""
	if len(implements) != 0 {
		implString = "implements " + strings.Join(implements, ", ")
	}
	body := ""
	if len(declarations) != 0 {
		body = fmt.Sprintf("%s\n", strings.Join(declarations, "\n"))
	}
	return fmt.Sprintf(
		`%s%s class %s %s %s {
%s%s`,
		annotationStr,
		strings.Join(modifiers, " "),
		n.Name,
		super,
		implString,
		body,
		v.withIndent("}"),
	), nil
}

func (v *TosVisitor) VisitModifier(n *Modifier) (interface{}, error) {
	return n.Name, nil
}

func (v *TosVisitor) VisitAnnotation(n *Annotation) (interface{}, error) {
	return n.Name, nil
}

func (v *TosVisitor) VisitInterfaceDeclaration(n *InterfaceDeclaration) (interface{}, error) {
	modifiers := make([]string, len(n.Modifiers))
	for i, m := range n.Modifiers {
		r, err := m.Accept(v)
		if err != nil {
			return nil, err
		}
		modifiers[i] = r.(string)
	}
	methods := make([]string, len(n.Methods))
	v.AddIndent(func() {
		for i, m := range n.Methods {
			r, err := m.Accept(v)
			if err != nil {
				panic(err)
			}
			methods[i] = r.(string)
		}
	})
	body := ""
	if len(methods) != 0 {
		body = fmt.Sprintf("%s\n", strings.Join(methods, "\n"))
	}

	return fmt.Sprintf(
		`%s interface %s {
%s%s`,
		strings.Join(modifiers, " "),
		n.Name,
		body,
		v.withIndent("}"),
	), nil
}

func (v *TosVisitor) VisitIntegerLiteral(n *IntegerLiteral) (interface{}, error) {
	return fmt.Sprintf("%d", n.Value), nil
}

func (v *TosVisitor) VisitParameter(n *Parameter) (interface{}, error) {
	r, err := n.TypeRef.Accept(v)
	if err != nil {
		return nil, err
	}
	return fmt.Sprintf(
		"%s %s",
		r.(string),
		n.Name,
	), nil
}

func (v *TosVisitor) VisitArrayAccess(n *ArrayAccess) (interface{}, error) {
	r, err := n.Receiver.Accept(v)
	if err != nil {
		return nil, err
	}
	k, err := n.Key.Accept(v)
	if err != nil {
		return nil, err
	}
	return fmt.Sprintf(
		"%s[%s]",
		r.(string),
		k.(string),
	), nil
}

func (v *TosVisitor) VisitBooleanLiteral(n *BooleanLiteral) (interface{}, error) {
	val := "false"
	if n.Value {
		val = "true"
	}
	return val, nil
}

func (v *TosVisitor) VisitBreak(n *Break) (interface{}, error) {
	return "break", nil
}

func (v *TosVisitor) VisitContinue(n *Continue) (interface{}, error) {
	return "continue", nil
}

func (v *TosVisitor) VisitDml(n *Dml) (interface{}, error) {
	r, err := n.Expression.Accept(v)
	if err != nil {
		return nil, err
	}
	return fmt.Sprintf("%s %s", n.Type, r.(string)), nil
}

func (v *TosVisitor) VisitDoubleLiteral(n *DoubleLiteral) (interface{}, error) {
	return fmt.Sprintf("%f", n.Value), nil
}

func (v *TosVisitor) VisitFieldDeclaration(n *FieldDeclaration) (interface{}, error) {
	modifiers := make([]string, len(n.Modifiers))
	for i, m := range n.Modifiers {
		r, err := m.Accept(v)
		if err != nil {
			return nil, err
		}
		modifiers[i] = r.(string)
	}
	r, err := n.TypeRef.Accept(v)
	if err != nil {
		return nil, err
	}
	declarators := make([]string, len(n.Declarators))
	for i, decl := range n.Declarators {
		r, err := decl.Accept(v)
		if err != nil {
			return nil, err
		}
		declarators[i] = r.(string)
	}

	return fmt.Sprintf(
		`%s %s %s;`,
		v.withIndent(strings.Join(modifiers, " ")),
		r.(string),
		strings.Join(declarators, ", "),
	), nil

}

func (v *TosVisitor) VisitTry(n *Try) (interface{}, error) {
	stmt := ""
	v.AddIndent(func() {
		r, err := n.Block.Accept(v)
		if err != nil {
			panic(err)
		}
		stmt = r.(string)
	})
	catches := make([]string, len(n.CatchClause))
	for i, c := range n.CatchClause {
		r, err := c.Accept(v)
		if err != nil {
			return nil, err
		}
		catches[i] = r.(string)
	}
	f, err := n.FinallyBlock.Accept(v)
	if err != nil {
		return nil, err
	}
	return fmt.Sprintf(
		`try {
%s%s%s
%s`,
		stmt,
		strings.Join(catches, "\n"),
		f.(string),
		v.withIndent("}"),
	), nil
}

func (v *TosVisitor) VisitCatch(n *Catch) (interface{}, error) {
	t, err := n.TypeRef.Accept(v)
	if err != nil {
		return nil, err
	}
	stmt := ""
	v.AddIndent(func() {
		r, err := n.Block.Accept(v)
		if err != nil {
			panic(err)
		}
		stmt = r.(string)
	})
	if stmt != "" {
		stmt = fmt.Sprintf("%s\n", stmt)
	}
	return fmt.Sprintf(
		` catch (%s %s) {
%s%s`,
		t.(string),
		n.Identifier,
		stmt,
		v.withIndent("}"),
	), nil
}

func (v *TosVisitor) VisitFinally(n *Finally) (interface{}, error) {
	stmt := ""
	v.AddIndent(func() {
		r, err := n.Block.Accept(v)
		if err != nil {
			panic(err)
		}
		stmt = r.(string)
	})
	if stmt != "" {
		stmt = fmt.Sprintf("%s\n", stmt)
	}
	return fmt.Sprintf(
		` finally {
%s%s`,
		stmt,
		v.withIndent("}"),
	), nil
}

func (v *TosVisitor) VisitFor(n *For) (interface{}, error) {
	control, err := n.Control.Accept(v)
	if err != nil {
		return nil, err
	}
	stmt := ""
	v.AddIndent(func() {
		r, err := n.Statements.Accept(v)
		if err != nil {
			panic(err)
		}
		stmt = r.(string)
	})
	if stmt != "" {
		stmt = fmt.Sprintf("%s\n", stmt)
	}
	return fmt.Sprintf(
		`for (%s) {
%s%s`,
		control.(string),
		stmt,
		v.withIndent("}"),
	), nil
}

func (v *TosVisitor) VisitForControl(n *ForControl) (interface{}, error) {
	inits := make([]string, len(n.ForInit))
	for i, forInit := range n.ForInit {
		exp, err := forInit.Accept(v)
		if err != nil {
			return nil, err
		}
		inits[i] = exp.(string)
	}
	exp, err := n.Expression.Accept(v)
	if err != nil {
		return nil, err
	}
	updates := make([]string, len(n.ForUpdate))
	for i, u := range n.ForUpdate {
		r, err := u.Accept(v)
		if err != nil {
			return nil, err
		}
		updates[i] = r.(string)
	}
	return fmt.Sprintf(
		`%s %s; %s`,
		strings.Join(inits, ", "),
		exp.(string),
		strings.Join(updates, ","),
	), nil
}

func (v *TosVisitor) VisitEnhancedForControl(n *EnhancedForControl) (interface{}, error) {
	t, err := n.TypeRef.Accept(v)
	if err != nil {
		return nil, err
	}
	exp, err := n.Expression.Accept(v)
	if err != nil {
		return nil, err
	}
	return fmt.Sprintf(
		`%s %s : %s`,
		t.(string),
		n.VariableDeclaratorId,
		exp.(string),
	), nil
}

func (v *TosVisitor) VisitIf(n *If) (interface{}, error) {
	cond, err := n.Condition.Accept(v)
	if err != nil {
		return nil, err
	}
	ifStmt := ""
	v.AddIndent(func() {
		r, err := n.IfStatement.Accept(v)
		if err != nil {
			panic(err)
		}
		ifStmt = r.(string)
	})
	if ifStmt != "" {
		ifStmt = fmt.Sprintf("%s\n", ifStmt)
	}
	elseStmt := ""
	if n.ElseStatement != nil {
		v.AddIndent(func() {
			r, err := n.IfStatement.Accept(v)
			if err != nil {
				panic(err)
			}
			elseStmt = r.(string)
		})
		if elseStmt != "" {
			elseStmt = fmt.Sprintf("%s\n", elseStmt)
		}
		elseStmt = fmt.Sprintf(
			` else {
%s
%s`,
			elseStmt,
			v.withIndent("}"),
		)
	}
	return fmt.Sprintf(
		`if (%s) {
%s%s%s`,
		cond.(string),
		ifStmt,
		v.withIndent("}"),
		elseStmt,
	), nil
}

func (v *TosVisitor) VisitMethodDeclaration(n *MethodDeclaration) (interface{}, error) {
	annotations := make([]string, len(n.Annotations))
	for i, a := range n.Annotations {
		r, err := a.Accept(v)
		if err != nil {
			return nil, err
		}
		annotations[i] = r.(string)
	}
	annotationStr := ""
	if len(annotations) != 0 {
		annotationStr = fmt.Sprintf("%s\n", strings.Join(annotations, "\n"))
	}
	modifiers := make([]string, len(n.Modifiers))
	for i, m := range n.Modifiers {
		r, err := m.Accept(v)
		if err != nil {
			return nil, err
		}
		modifiers[i] = r.(string)
	}
	returnType := "void"
	if n.ReturnType != nil {
		r, err := n.ReturnType.Accept(v)
		if err != nil {
			return nil, err
		}
		returnType = r.(string)
	}
	parameters := make([]string, len(n.Parameters))
	for i, p := range n.Parameters {
		r, err := p.Accept(v)
		if err != nil {
			return nil, err
		}
		parameters[i] = r.(string)
	}
	block := ""
	v.AddIndent(func() {
		r, err := n.Statements.Accept(v)
		if err != nil {
			panic(err)
		}
		block = r.(string)
	})
	if block != "" {
		block = fmt.Sprintf("%s\n", block)
	}
	return fmt.Sprintf(
		`%s%s %s %s (%s) {
%s%s`,
		annotationStr,
		v.withIndent(strings.Join(modifiers, " ")),
		returnType,
		n.Name,
		strings.Join(parameters, ", "),
		block,
		v.withIndent("}"),
	), nil
}

func (v *TosVisitor) VisitMethodInvocation(n *MethodInvocation) (interface{}, error) {
	exp, err := n.NameOrExpression.Accept(v)
	if err != nil {
		return nil, err
	}
	parameters := make([]string, len(n.Parameters))
	for i, p := range n.Parameters {
		r, err := p.Accept(v)
		if err != nil {
			return nil, err
		}
		parameters[i] = r.(string)
	}
	return fmt.Sprintf(
		"%s(%s)",
		exp.(string),
		strings.Join(parameters, ", "),
	), nil
}

func (v *TosVisitor) VisitNew(n *New) (interface{}, error) {
	t, err := n.TypeRef.Accept(v)
	if err != nil {
		return nil, err
	}
	parameters := make([]string, len(n.Parameters))
	for i, p := range n.Parameters {
		r, err := p.Accept(v)
		if err != nil {
			return nil, err
		}
		parameters[i] = r.(string)
	}
	return fmt.Sprintf(
		"new %s(%s)",
		t.(string),
		strings.Join(parameters, ", "),
	), nil
}

func (v *TosVisitor) VisitNullLiteral(n *NullLiteral) (interface{}, error) {
	return "null", nil
}

func (v *TosVisitor) VisitUnaryOperator(n *UnaryOperator) (interface{}, error) {
	val, err := n.Expression.Accept(v)
	if err != nil {
		return nil, err
	}
	if n.IsPrefix {
		return fmt.Sprintf("%s%s", n.Op, val.(string)), nil
	}
	return fmt.Sprintf("%s%s", val.(string), n.Op), nil
}

func (v *TosVisitor) VisitBinaryOperator(n *BinaryOperator) (interface{}, error) {
	l, err := n.Left.Accept(v)
	if err != nil {
		return nil, err
	}
	r, err := n.Right.Accept(v)
	if err != nil {
		return nil, err
	}
	return fmt.Sprintf("%s %s %s", l.(string), n.Op, r.(string)), nil
}

func (v *TosVisitor) VisitInstanceofOperator(n *InstanceofOperator) (interface{}, error) {
	l, err := n.Expression.Accept(v)
	if err != nil {
		return nil, err
	}
	r, err := n.TypeRef.Accept(v)
	if err != nil {
		return nil, err
	}
	return fmt.Sprintf("%s instanceof %s", l.(string), r.(string)), nil
}

func (v *TosVisitor) VisitReturn(n *Return) (interface{}, error) {
	if n.Expression != nil {
		exp, err := n.Expression.Accept(v)
		if err != nil {
			return nil, err
		}
		return fmt.Sprintf("return %s", exp.(string)), nil
	}
	return "return", nil
}

func (v *TosVisitor) VisitThrow(n *Throw) (interface{}, error) {
	if n.Expression != nil {
		exp, err := n.Expression.Accept(v)
		if err != nil {
			return nil, err
		}
		return fmt.Sprintf("throw %s", exp.(string)), nil
	}
	return "throw", nil
}

func (v *TosVisitor) VisitSoql(n *Soql) (interface{}, error) {
	where := ""
	fields := make([]string, len(n.SelectFields))
	from := ""
	v.AddIndent(func() {
		v.AddIndent(func() {
			for i, f := range n.SelectFields {
				switch val := f.(type) {
				case *SelectField:
					fields[i] = v.withIndent(strings.Join(val.Value, "."))
				case *SoqlFunction:
					fields[i] = v.withIndent(val.Name + "()")
				}
			}

			from = v.withIndent(n.FromObject)

			if n.Where != nil {
				where = v.withIndent(v.createWhere(n.Where))
			}
		})
	})

	indent := ""
	v.AddIndent(func() {
		indent = v.withIndent("")
	})
	if where != "" {
		where = "\n" + indent + "WHERE\n" + where
	}
	orderBy := ""
	groupBy := ""
	limit := ""
	if n.Limit != nil {
		i, err := n.Limit.Accept(v)
		if err != nil {
			return nil, err
		}
		v.AddIndent(func() {
			v.AddIndent(func() {
				limit = "\n" + indent + "LIMIT\n" + v.withIndent(i.(string))
			})
		})
	}

	return fmt.Sprintf(`[
%sSELECT
%s
%sFROM
%s%s%s%s%s%s`,
		indent,
		strings.Join(fields, ",\n"),
		indent,
		from,
		where,
		orderBy,
		groupBy,
		limit,
		"\n"+v.withIndent("]"),
	), nil
}

func (v *TosVisitor) createWhere(n Node) string {
	switch val := n.(type) {
	case *WhereCondition:
		var field string
		switch f := val.Field.(type) {
		case *SelectField:
			field = strings.Join(f.Value, ".")
		case *SoqlFunction:
			field = f.Name + "()"
		}
		value, err := val.Expression.Accept(v)
		if err != nil {
			panic(err)
		}
		return fmt.Sprintf("%s %s %s", field, val.Op, value.(string))
	case *WhereBinaryOperator:
		where := ""
		if val.Left != nil {
			where += v.createWhere(val.Left)
		}
		if val.Right != nil {
			where += fmt.Sprintf("\n%s %s", v.withIndent(val.Op), v.createWhere(val.Right))
		}
		return where
	}
	return ""
}

func (v *TosVisitor) VisitSosl(n *Sosl) (interface{}, error) {
	return VisitSosl(v, n)
}

func (v *TosVisitor) VisitStringLiteral(n *StringLiteral) (interface{}, error) {
	return "'" + n.Value + "'", nil
}

func (v *TosVisitor) VisitSwitch(n *Switch) (interface{}, error) {
	exp, err := n.Expression.Accept(v)
	if err != nil {
		return nil, err
	}
	whenStmts := make([]string, len(n.WhenStatements))
	v.AddIndent(func() {
		for i, stmt := range n.WhenStatements {
			r, err := stmt.Accept(v)
			if err != nil {
				panic(err)
			}
			whenStmts[i] = r.(string)
		}
	})
	elseStmt := ""
	v.AddIndent(func() {
		r, err := n.ElseStatement.Accept(v)
		if err != nil {
			panic(err)
		}
		elseStmt = r.(string)
	})
	if elseStmt != "" {
		elseStmt = fmt.Sprintf(
			` else {
%s
%s`,
			elseStmt,
			v.withIndent("}"),
		)
	}
	return fmt.Sprintf(
		`switch on %s {
%s
%s
%s`,
		exp.(string),
		strings.Join(whenStmts, "\n"),
		elseStmt,
		v.withIndent("}"),
	), nil
}

func (v *TosVisitor) VisitTrigger(n *Trigger) (interface{}, error) {
	timings := make([]string, len(n.TriggerTimings))
	for i, t := range n.TriggerTimings {
		r, err := t.Accept(v)
		if err != nil {
			return nil, err
		}
		timings[i] = r.(string)
	}
	stmt, err := n.Statements.Accept(v)
	if err != nil {
		return nil, err
	}
	return fmt.Sprintf(
		`trigger %s on %s (%s) {
%s
%s`,
		n.Name,
		n.Object,
		strings.Join(timings, ", "),
		stmt.(string),
		v.withIndent("}"),
	), nil
}

func (v *TosVisitor) VisitTriggerTiming(n *TriggerTiming) (interface{}, error) {
	return fmt.Sprintf("%s %s", n.Timing, n.Dml), nil
}

func (v *TosVisitor) VisitVariableDeclaration(n *VariableDeclaration) (interface{}, error) {
	t, err := n.TypeRef.Accept(v)
	if err != nil {
		return nil, err
	}
	declarators := make([]string, len(n.Declarators))
	for i, decl := range n.Declarators {
		r, err := decl.Accept(v)
		if err != nil {
			return nil, err
		}
		declarators[i] = r.(string)
	}
	return fmt.Sprintf(
		"%s %s",
		t.(string),
		strings.Join(declarators, ", "),
	), nil
}

func (v *TosVisitor) VisitVariableDeclarator(n *VariableDeclarator) (interface{}, error) {
	if n.Expression == nil {
		return fmt.Sprintf("%s", n.Name), nil
	}
	exp, err := n.Expression.Accept(v)
	if err != nil {
		return nil, err
	}
	return fmt.Sprintf("%s = %s", n.Name, exp.(string)), nil
}

func (v *TosVisitor) VisitWhen(n *When) (interface{}, error) {
	conditions := make([]string, len(n.Condition))
	for i, cond := range n.Condition {
		r, err := cond.Accept(v)
		if err != nil {
			return nil, err
		}
		conditions[i] = r.(string)
	}
	stmt := ""
	v.AddIndent(func() {
		r, err := n.Statements.Accept(v)
		if err != nil {
			panic(err)
		}
		stmt = r.(string)
	})
	return fmt.Sprintf(
		`when %s {
%s
%s`,
		strings.Join(conditions, ", "),
		stmt,
		v.withIndent("}"),
	), nil

	return VisitWhen(v, n)
}

func (v *TosVisitor) VisitWhenType(n *WhenType) (interface{}, error) {
	r, err := n.TypeRef.Accept(v)
	if err != nil {
		return nil, err
	}
	return fmt.Sprintf(
		"%s %s",
		r.(string),
		n.Identifier,
	), nil
}

func (v *TosVisitor) VisitWhile(n *While) (interface{}, error) {
	cond, err := n.Condition.Accept(v)
	if err != nil {
		return nil, err
	}
	statements := ""
	v.AddIndent(func() {
		r, err := n.Statements.Accept(v)
		if err != nil {
			panic(err)
		}
		statements = r.(string)
	})
	return fmt.Sprintf(
		`while (%s) {
%s
%s`,
		cond.(string),
		statements,
		v.withIndent("}"),
	), nil
}

func (v *TosVisitor) VisitNothingStatement(n *NothingStatement) (interface{}, error) {
	return "", nil
}

func (v *TosVisitor) VisitCastExpression(n *CastExpression) (interface{}, error) {
	t, err := n.CastTypeRef.Accept(v)
	if err != nil {
		return nil, err
	}
	exp, err := n.Expression.Accept(v)
	if err != nil {
		return nil, err
	}
	return fmt.Sprintf("(%s)%s", t.(string), exp.(string)), nil
}

func (v *TosVisitor) VisitFieldAccess(n *FieldAccess) (interface{}, error) {
	exp, err := n.Expression.Accept(v)
	if err != nil {
		return nil, err
	}
	return fmt.Sprintf("%s.%s", exp.(string), n.FieldName), nil
}

func (v *TosVisitor) VisitType(n *TypeRef) (interface{}, error) {
	paramString := ""
	params := make([]string, len(n.Parameters))
	for i, param := range n.Parameters {
		r, err := param.Accept(v)
		if err != nil {
			return nil, err
		}
		params[i] = r.(string)
	}
	if len(params) != 0 {
		paramString = fmt.Sprintf("<%s>", strings.Join(params, ", "))
	}
	return fmt.Sprintf(
		"%s%s",
		strings.Join(n.Name, "."),
		paramString,
	), nil
}

func (v *TosVisitor) VisitBlock(n *Block) (interface{}, error) {
	statements := make([]string, len(n.Statements))
	for i, s := range n.Statements {
		r, err := s.Accept(v)
		if err != nil {
			return nil, err
		}
		statements[i] = v.withIndent(r.(string)) + ";"
	}
	return strings.Join(statements, "\n"), nil
}

func (v *TosVisitor) VisitGetterSetter(n *GetterSetter) (interface{}, error) {
	return VisitGetterSetter(v, n)
}

func (v *TosVisitor) VisitPropertyDeclaration(n *PropertyDeclaration) (interface{}, error) {
	return VisitPropertyDeclaration(v, n)
}

func (v *TosVisitor) VisitArrayInitializer(n *ArrayInitializer) (interface{}, error) {
	return VisitArrayInitializer(v, n)
}

func (v *TosVisitor) VisitArrayCreator(n *ArrayCreator) (interface{}, error) {
	return VisitArrayCreator(v, n)
}

func (v *TosVisitor) VisitSoqlBindVariable(n *SoqlBindVariable) (interface{}, error) {
	return VisitSoqlBindVariable(v, n)
}

func (v *TosVisitor) VisitTernalyExpression(n *TernalyExpression) (interface{}, error) {
	return VisitTernalyExpression(v, n)
}

func (v *TosVisitor) VisitMapCreator(n *MapCreator) (interface{}, error) {
	return VisitMapCreator(v, n)
}

func (v *TosVisitor) VisitSetCreator(n *SetCreator) (interface{}, error) {
	return VisitSetCreator(v, n)
}

func (v *TosVisitor) VisitName(n *Name) (interface{}, error) {
	return strings.Join(n.Value, "."), nil
}

func (v *TosVisitor) VisitConstructorDeclaration(n *ConstructorDeclaration) (interface{}, error) {
	return VisitConstructorDeclaration(v, n)
}

func ToString(n Node) string {
	visitor := &TosVisitor{}
	r, err := n.Accept(visitor)
	if err != nil {
		panic(err)
	}
	return r.(string)
}
