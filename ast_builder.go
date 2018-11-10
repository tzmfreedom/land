package main

import (
	"strings"

	"strconv"

	"github.com/antlr/antlr4/runtime/Go/antlr"
	"github.com/tzmfreedom/goland/ast"
	"github.com/tzmfreedom/goland/parser"
)

type AstBuilder struct {
	*parser.BaseapexVisitor
	CurrentFile string
}

func (v *AstBuilder) VisitCompilationUnit(ctx *parser.CompilationUnitContext) interface{} {
	return ctx.TypeDeclaration().Accept(v)
}

func (v *AstBuilder) VisitTypeDeclaration(ctx *parser.TypeDeclarationContext) interface{} {
	classOrInterfaceModifiers := ctx.AllClassOrInterfaceModifier()
	modifiers := []ast.Node{}
	annotations := []ast.Node{}
	for _, classOrInterfaceModifier := range classOrInterfaceModifiers {
		r := classOrInterfaceModifier.Accept(v)
		switch n := r.(type) {
		case *ast.Modifier:
			modifiers = append(modifiers, n)
		case *ast.Annotation:
			annotations = append(annotations, n)
		}
	}

	if ctx.ClassDeclaration() != nil {
		cd := ctx.ClassDeclaration().Accept(v)
		classDeclaration, _ := cd.(*ast.ClassDeclaration)
		classDeclaration.Modifiers = modifiers
		classDeclaration.Annotations = annotations
		return classDeclaration
	}
	return nil
}

func (v *AstBuilder) VisitTriggerDeclaration(ctx *parser.TriggerDeclarationContext) interface{} {
	timings := ctx.TriggerTimings().Accept(v)

	name := ctx.ApexIdentifier(0).GetText()
	object := ctx.ApexIdentifier(1).GetText()
	block := ctx.Block().Accept(v)
	return &ast.Trigger{
		Name:           name,
		TriggerTimings: timings.([]*ast.TriggerTiming),
		Object:         object,
		Statements:     block.([]ast.Node),
		Position:       v.newPosition(ctx),
	}
}

func (v *AstBuilder) VisitTriggerTimings(ctx *parser.TriggerTimingsContext) interface{} {
	allTimings := ctx.AllTriggerTiming()
	timings := make([]ast.Node, len(allTimings))
	for i, timing := range allTimings {
		timings[i] = timing.Accept(v).(ast.Node)
	}
	return timings
}

func (v *AstBuilder) VisitTriggerTiming(ctx *parser.TriggerTimingContext) interface{} {
	return &ast.TriggerTiming{
		Timing:   ctx.GetTiming().GetText(),
		Dml:      ctx.GetDml().GetText(),
		Position: v.newPosition(ctx),
	}
}

func (v *AstBuilder) VisitModifier(ctx *parser.ModifierContext) interface{} {
	m := ctx.ClassOrInterfaceModifier()
	if m != nil {
		return m.Accept(v)
	}
	return &ast.Modifier{
		Name:     ctx.GetText(),
		Position: v.newPosition(ctx),
	}
}

func (v *AstBuilder) VisitClassOrInterfaceModifier(ctx *parser.ClassOrInterfaceModifierContext) interface{} {
	annotation := ctx.Annotation()
	if annotation != nil {
		return ctx.Annotation().Accept(v)
	}
	return &ast.Modifier{
		Name:     ctx.GetText(),
		Position: v.newPosition(ctx),
	}
}

func (v *AstBuilder) VisitVariableModifier(ctx *parser.VariableModifierContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *AstBuilder) VisitClassDeclaration(ctx *parser.ClassDeclarationContext) interface{} {
	declarations := ctx.ClassBody().Accept(v).([]ast.Node)

	n := &ast.ClassDeclaration{
		Name:     ctx.ApexIdentifier().GetText(),
		Position: v.newPosition(ctx),
	}
	if t := ctx.ApexType(); t != nil {
		n.SuperClass = t.Accept(v).(ast.Node)
	}
	if tl := ctx.TypeList(); tl != nil {
		n.ImplementClasses = tl.Accept(v).([]ast.Node)
	}
	n.Declarations = make([]ast.Node, len(declarations))
	for i, d := range declarations {
		n.Declarations[i] = d
	}
	return n
}

func (v *AstBuilder) VisitEnumDeclaration(ctx *parser.EnumDeclarationContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *AstBuilder) VisitEnumConstants(ctx *parser.EnumConstantsContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *AstBuilder) VisitEnumConstant(ctx *parser.EnumConstantContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *AstBuilder) VisitEnumBodyDeclarations(ctx *parser.EnumBodyDeclarationsContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *AstBuilder) VisitInterfaceDeclaration(ctx *parser.InterfaceDeclarationContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *AstBuilder) VisitTypeList(ctx *parser.TypeListContext) interface{} {
	apexTypes := ctx.AllApexType()
	types := make([]ast.Node, len(apexTypes))
	for i, t := range apexTypes {
		types[i] = t.Accept(v).(ast.Node)
	}
	return types
}

func (v *AstBuilder) VisitClassBody(ctx *parser.ClassBodyContext) interface{} {
	bodyDeclarations := ctx.AllClassBodyDeclaration()
	declarations := make([]ast.Node, len(bodyDeclarations))
	for i, d := range bodyDeclarations {
		declarations[i] = d.Accept(v).(ast.Node)
	}
	return declarations
}

func (v *AstBuilder) VisitInterfaceBody(ctx *parser.InterfaceBodyContext) interface{} {
	bodyDeclarations := ctx.AllInterfaceBodyDeclaration()
	declarations := make([]ast.Node, len(bodyDeclarations))
	for i, d := range bodyDeclarations {
		declarations[i] = d.Accept(v).(ast.Node)
	}
	return declarations
}

func (v *AstBuilder) VisitClassBodyDeclaration(ctx *parser.ClassBodyDeclarationContext) interface{} {
	memberDeclaration := ctx.MemberDeclaration()
	if memberDeclaration != nil {
		declaration := memberDeclaration.Accept(v)

		modifiers := ctx.AllModifier()
		declarationModifiers := make([]ast.Node, len(modifiers))
		for i, m := range modifiers {
			declarationModifiers[i] = m.Accept(v).(ast.Node)
		}
		switch decl := declaration.(type) {
		case *ast.MethodDeclaration:
			decl.Modifiers = declarationModifiers
			return decl
		case *ast.FieldDeclaration:
			decl.Modifiers = declarationModifiers
			return decl
		case *ast.ConstructorDeclaration:
			decl.Modifiers = declarationModifiers
			return decl
		case *ast.InterfaceDeclaration:
			decl.Modifiers = declarationModifiers
			return decl
		case *ast.ClassDeclaration:
			decl.Modifiers = declarationModifiers
			return decl
		//case *ast.EnumDeclaration:
		//	decl.Modifiers = declarationModifiers
		//	return decl
		case *ast.PropertyDeclaration:
			decl.Modifiers = declarationModifiers
			return decl
		}
	}
	return nil
}

func (v *AstBuilder) VisitMemberDeclaration(ctx *parser.MemberDeclarationContext) interface{} {
	if d := ctx.MethodDeclaration(); d != nil {
		return d.Accept(v)
	} else if d := ctx.FieldDeclaration(); d != nil {
		return d.Accept(v)
	} else if d := ctx.ConstructorDeclaration(); d != nil {
		return d.Accept(v)
	} else if d := ctx.InterfaceDeclaration(); d != nil {
		return d.Accept(v)
	} else if d := ctx.ClassDeclaration(); d != nil {
		return d.Accept(v)
	} else if d := ctx.EnumDeclaration(); d != nil {
		return d.Accept(v)
	} else if d := ctx.PropertyDeclaration(); d != nil {
		return d.Accept(v)
	}
	return nil
}

func (v *AstBuilder) VisitMethodDeclaration(ctx *parser.MethodDeclarationContext) interface{} {
	n := &ast.MethodDeclaration{Position: v.newPosition(ctx)}
	n.Name = ctx.ApexIdentifier().GetText()
	if ctx.ApexType() != nil {
		n.ReturnType = ctx.ApexType().Accept(v).(ast.Node)
	} else {
		n.ReturnType = ast.VoidType
	}
	n.Parameters = ctx.FormalParameters().Accept(v).([]*ast.Parameter)
	if ctx.QualifiedNameList() != nil {
		n.Throws = ctx.QualifiedNameList().Accept(v).([]ast.Node)
	} else {
		n.Throws = []ast.Node{}
	}
	if ctx.MethodBody() != nil {
		n.Statements = ctx.MethodBody().Accept(v).(*ast.Block)
	} else {
		n.Statements = &ast.Block{}
	}
	return n
}

func (v *AstBuilder) VisitConstructorDeclaration(ctx *parser.ConstructorDeclarationContext) interface{} {
	parameters := ctx.FormalParameters().Accept(v).([]*ast.Parameter)
	var throws []ast.Node
	if q := ctx.QualifiedNameList(); q != nil {
		throws = q.Accept(v).([]ast.Node)
	} else {
		throws = []ast.Node{}
	}
	body := ctx.ConstructorBody().Accept(v).([]ast.Node)
	return &ast.ConstructorDeclaration{
		Parameters: parameters,
		Throws:     throws,
		Statements: body,
		Position:   v.newPosition(ctx),
	}
}

func (v *AstBuilder) VisitFieldDeclaration(ctx *parser.FieldDeclarationContext) interface{} {
	t := ctx.ApexType().Accept(v).(ast.Node)
	d := ctx.VariableDeclarators().Accept(v).([]ast.Node)
	return &ast.FieldDeclaration{
		Type:        t,
		Declarators: d,
	}
}

func (v *AstBuilder) VisitPropertyDeclaration(ctx *parser.PropertyDeclarationContext) interface{} {
	t := ctx.ApexType().Accept(v).(ast.Node)
	d := ctx.VariableDeclaratorId().Accept(v).(string)
	b := ctx.PropertyBodyDeclaration().Accept(v).(ast.Node)
	return &ast.PropertyDeclaration{
		Type:          t,
		Identifier:    d,
		GetterSetters: b,
	}
}

func (v *AstBuilder) VisitPropertyBodyDeclaration(ctx *parser.PropertyBodyDeclarationContext) interface{} {
	blocks := ctx.AllPropertyBlock()
	declarations := make([]*ast.Block, len(blocks))
	for i, b := range blocks {
		declarations[i] = b.Accept(v).(*ast.Block)
	}
	return declarations
}

func (v *AstBuilder) VisitInterfaceBodyDeclaration(ctx *parser.InterfaceBodyDeclarationContext) interface{} {
	d := ctx.InterfaceMemberDeclaration().Accept(v).(ast.Interface)
	modifiers := ctx.AllModifier()
	d.Modifiers = make([]ast.Node, len(modifiers)+1)
	for i, m := range modifiers {
		d.Modifiers[i] = m.Accept(v).(ast.Node)
	}
	d.Modifiers[len(modifiers)] = &ast.Modifier{
		Name:     "public",
		Position: v.newPosition(ctx),
	}
	return d
}

func (v *AstBuilder) VisitInterfaceMemberDeclaration(ctx *parser.InterfaceMemberDeclarationContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *AstBuilder) VisitConstDeclaration(ctx *parser.ConstDeclarationContext) interface{} {
	_ = ctx.ApexType().Accept(v)
	_ = ctx.AllConstantDeclarator()

	// TODO: implement
	return nil
}

func (v *AstBuilder) VisitConstantDeclarator(ctx *parser.ConstantDeclaratorContext) interface{} {
	_ = ctx.ApexIdentifier().Accept(v)
	_ = ctx.VariableInitializer().Accept(v)

	// TODO: implement
	return nil
}

func (v *AstBuilder) VisitInterfaceMethodDeclaration(ctx *parser.InterfaceMethodDeclarationContext) interface{} {
	decl := &ast.MethodDeclaration{Position: v.newPosition(ctx)}
	decl.Name = ctx.ApexIdentifier().Accept(v).(string)

	if t := ctx.ApexType(); t != nil {
		decl.ReturnType = t.Accept(v).(ast.Node)
	} else {
		// TODO: implement void
	}
	decl.Parameters = ctx.FormalParameters().Accept(v).([]*ast.Parameter)
	if q := ctx.QualifiedNameList(); q != nil {
		decl.Throws = q.Accept(v).([]ast.Node)
	} else {
		decl.Throws = []ast.Node{}
	}
	return decl
}

func (v *AstBuilder) VisitVariableDeclarators(ctx *parser.VariableDeclaratorsContext) interface{} {
	variableDeclarators := ctx.AllVariableDeclarator()
	declarators := make([]ast.Node, len(variableDeclarators))
	for i, d := range variableDeclarators {
		declarators[i] = d.Accept(v).(ast.Node)
	}
	return declarators
}

func (v *AstBuilder) VisitVariableDeclarator(ctx *parser.VariableDeclaratorContext) interface{} {
	decl := &ast.VariableDeclarator{Position: v.newPosition(ctx)}
	decl.Name = ctx.VariableDeclaratorId().Accept(v).(string)
	if init := ctx.VariableInitializer(); init != nil {
		decl.Expression = init.Accept(v).(ast.Node)
	} else {
		decl.Expression = &ast.NullLiteral{Position: v.newPosition(ctx)}
	}
	return decl
}

func (v *AstBuilder) VisitVariableDeclaratorId(ctx *parser.VariableDeclaratorIdContext) interface{} {
	return ctx.ApexIdentifier().GetText()
}

func (v *AstBuilder) VisitVariableInitializer(ctx *parser.VariableInitializerContext) interface{} {
	if init := ctx.ArrayInitializer(); init != nil {
		return init.Accept(v)
	}
	return ctx.Expression().Accept(v)
}

func (v *AstBuilder) VisitArrayInitializer(ctx *parser.ArrayInitializerContext) interface{} {
	if inits := ctx.AllVariableInitializer(); len(inits) != 0 {
		initializers := make([]ast.Node, len(inits))
		for i, init := range inits {
			initializers[i] = init.Accept(v).(ast.Node)
		}
		return initializers
	}
	return nil
}

func (v *AstBuilder) VisitEnumConstantName(ctx *parser.EnumConstantNameContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *AstBuilder) VisitApexType(ctx *parser.ApexTypeContext) interface{} {
	if interfaceType := ctx.ClassOrInterfaceType(); interfaceType != nil {
		t := interfaceType.Accept(v).(ast.Node)
		// TODO: implement Array
		return t
	} else if primitiveType := ctx.PrimitiveType(); primitiveType != nil {
		t := primitiveType.Accept(v).(ast.Node)
		// TODO: implement Array
		return t
	}
	return nil
}

func (v *AstBuilder) VisitTypedArray(ctx *parser.TypedArrayContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *AstBuilder) VisitClassOrInterfaceType(ctx *parser.ClassOrInterfaceTypeContext) interface{} {
	t := &ast.Type{Position: v.newPosition(ctx)}
	arguments := ctx.AllTypeArguments()
	t.Parameters = make([]ast.Node, len(arguments))
	for i, argument := range arguments {
		t.Parameters[i] = argument.Accept(v).(ast.Node)
	}
	if idents := ctx.AllTypeIdentifier(); len(idents) != 0 {
		t.Name = make([]string, len(idents))
		for i, ident := range idents {
			t.Name[i] = ident.Accept(v).(string)
		}
	} else {
		t.Name = []string{ctx.SET().GetText()}
	}
	return t
}

func (v *AstBuilder) VisitPrimitiveType(ctx *parser.PrimitiveTypeContext) interface{} {
	return &ast.Type{
		Name:       []string{ctx.GetText()},
		Parameters: []ast.Node{},
		Position:   v.newPosition(ctx),
	}
}

func (v *AstBuilder) VisitTypeArguments(ctx *parser.TypeArgumentsContext) interface{} {
	arguments := ctx.AllTypeArgument()
	typeArguments := make([]ast.Node, len(arguments))
	for i, a := range arguments {
		typeArguments[i] = a.Accept(v).(ast.Node)
	}
	return typeArguments
}

func (v *AstBuilder) VisitTypeArgument(ctx *parser.TypeArgumentContext) interface{} {
	return ctx.ApexType().Accept(v)
}

func (v *AstBuilder) VisitQualifiedNameList(ctx *parser.QualifiedNameListContext) interface{} {
	qualifiedNames := ctx.AllQualifiedName()
	names := make([]ast.Node, len(qualifiedNames))
	for i, qn := range qualifiedNames {
		names[i] = qn.Accept(v).(ast.Node)
	}
	return names
}

func (v *AstBuilder) VisitFormalParameters(ctx *parser.FormalParametersContext) interface{} {
	if p := ctx.FormalParameterList(); p != nil {
		return p.Accept(v)
	}
	return []*ast.Parameter{}
}

func (v *AstBuilder) VisitFormalParameterList(ctx *parser.FormalParameterListContext) interface{} {
	formalParameters := ctx.AllFormalParameter()
	parameters := make([]*ast.Parameter, len(formalParameters))
	for i, p := range formalParameters {
		parameters[i] = p.Accept(v).(*ast.Parameter)
	}
	return parameters
}

func (v *AstBuilder) VisitFormalParameter(ctx *parser.FormalParameterContext) interface{} {
	p := &ast.Parameter{Position: v.newPosition(ctx)}
	modifiers := ctx.AllVariableModifier()
	p.Modifiers = make([]ast.Node, len(modifiers))
	for i, m := range modifiers {
		p.Modifiers[i] = m.Accept(v).(ast.Node)
	}
	p.Type = ctx.ApexType().Accept(v).(ast.Node)
	p.Name = ctx.VariableDeclaratorId().Accept(v).(string)
	return p
}

func (v *AstBuilder) VisitLastFormalParameter(ctx *parser.LastFormalParameterContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *AstBuilder) VisitMethodBody(ctx *parser.MethodBodyContext) interface{} {
	return ctx.Block().Accept(v)
}

func (v *AstBuilder) VisitConstructorBody(ctx *parser.ConstructorBodyContext) interface{} {
	return ctx.Block().Accept(v)
}

func (v *AstBuilder) VisitQualifiedName(ctx *parser.QualifiedNameContext) interface{} {
	allIdentifiers := ctx.AllApexIdentifier()
	identifiers := make([]string, len(allIdentifiers))
	for i, identifier := range allIdentifiers {
		ident := identifier.Accept(v)
		identifiers[i], _ = ident.(string)
	}
	n := &ast.Type{Position: v.newPosition(ctx)}
	n.Name = identifiers
	return n
}

func (v *AstBuilder) VisitLiteral(ctx *parser.LiteralContext) interface{} {
	if lit := ctx.IntegerLiteral(); lit != nil {
		val, err := strconv.Atoi(lit.GetText())
		if err != nil {
			panic(err)
		}
		return &ast.IntegerLiteral{Value: val, Position: v.newPosition(ctx)}
	} else if lit := ctx.FloatingPointLiteral(); lit != nil {
		val, err := strconv.ParseFloat(lit.GetText(), 64)
		if err != nil {
			panic(err)
		}
		return &ast.DoubleLiteral{Value: val, Position: v.newPosition(ctx)}
	} else if lit := ctx.StringLiteral(); lit != nil {
		str := lit.GetText()
		return &ast.StringLiteral{Value: str[0 : len(str)-2], Position: v.newPosition(ctx)}
	} else if lit := ctx.BooleanLiteral(); lit != nil {
		// TODO: implement caseinsensitive value
		return &ast.BooleanLiteral{Value: lit.GetText() == "true", Position: v.newPosition(ctx)}
	} else if lit := ctx.NullLiteral(); lit != nil {
		return &ast.NullLiteral{Position: v.newPosition(ctx)}
	}
	return nil
}

func (v *AstBuilder) VisitAnnotation(ctx *parser.AnnotationContext) interface{} {
	name := ctx.AnnotationName().Accept(v).(*ast.Type)
	annotation := &ast.Annotation{}
	// TODO: implement annotationName
	annotation.Name = name.Name[0]
	annotation.Position = v.newPosition(ctx)
	return annotation
}

func (v *AstBuilder) VisitAnnotationName(ctx *parser.AnnotationNameContext) interface{} {
	return ctx.QualifiedName().Accept(v)
}

func (v *AstBuilder) VisitElementValuePairs(ctx *parser.ElementValuePairsContext) interface{} {
	elementValuePairs := ctx.AllElementValuePair()
	pairs := make([]ast.Node, len(elementValuePairs))
	for i, p := range elementValuePairs {
		pairs[i] = p.Accept(v).(ast.Node)
	}
	return pairs
}

func (v *AstBuilder) VisitElementValuePair(ctx *parser.ElementValuePairContext) interface{} {
	ctx.ApexIdentifier().GetText()
	ctx.ElementValue().Accept(v)
	// TODO: implement
	return nil
}

func (v *AstBuilder) VisitElementValue(ctx *parser.ElementValueContext) interface{} {
	if e := ctx.Expression(); e != nil {
		return e.Accept(v)
	} else if a := ctx.Annotation(); a != nil {
		return a.Accept(v)
	} else if init := ctx.ElementValueArrayInitializer(); init != nil {
		return init.Accept(v)
	}
	return nil
}

func (v *AstBuilder) VisitElementValueArrayInitializer(ctx *parser.ElementValueArrayInitializerContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *AstBuilder) VisitBlock(ctx *parser.BlockContext) interface{} {
	blk := &ast.Block{Position: v.newPosition(ctx)}
	statements := ctx.AllBlockStatement()
	blk.Statements = make([]ast.Node, len(statements))
	for i, statement := range statements {
		s := statement.Accept(v)
		blk.Statements[i] = s.(ast.Node)
	}
	return blk
}

func (v *AstBuilder) VisitBlockStatement(ctx *parser.BlockStatementContext) interface{} {
	if s := ctx.Statement(); s != nil {
		return s.Accept(v)
	} else if s := ctx.LocalVariableDeclarationStatement(); s != nil {
		return s.Accept(v)
	} else if t := ctx.TypeDeclaration(); t != nil {
		return t.Accept(v)
	}
	return nil
}

func (v *AstBuilder) VisitLocalVariableDeclarationStatement(ctx *parser.LocalVariableDeclarationStatementContext) interface{} {
	return ctx.LocalVariableDeclaration().Accept(v)
}

func (v *AstBuilder) VisitLocalVariableDeclaration(ctx *parser.LocalVariableDeclarationContext) interface{} {
	decl := &ast.VariableDeclaration{Position: v.newPosition(ctx)}
	modifiers := ctx.AllVariableModifier()
	decl.Modifiers = make([]ast.Node, len(modifiers))
	for i, m := range modifiers {
		decl.Modifiers[i] = m.Accept(v).(ast.Node)
	}
	decl.Type = ctx.ApexType().Accept(v).(ast.Node)
	decl.Declarators = ctx.VariableDeclarators().Accept(v).([]ast.Node)
	return decl
}

func (v *AstBuilder) VisitStatement(ctx *parser.StatementContext) interface{} {
	if t := ctx.TRY(); t != nil {
		try := &ast.Try{}
		try.Block = ctx.Block().Accept(v).(*ast.Block)
		if clauses := ctx.AllCatchClause(); len(clauses) != 0 {
			try.CatchClause = make([]ast.Node, len(clauses))
			for i, c := range clauses {
				try.CatchClause[i] = c.Accept(v).(ast.Node)
			}
		} else {
			try.CatchClause = []ast.Node{}
		}
		if b := ctx.FinallyBlock(); b != nil {
			try.FinallyBlock = b.Accept(v).(*ast.Block)
		}
		return try
	} else if t := ctx.IF(); t != nil {
		n := &ast.If{Position: v.newPosition(ctx)}
		n.Condition = ctx.ParExpression().Accept(v).(ast.Node)
		n.IfStatement = ctx.Statement(0).Accept(v).(ast.Node)
		if ctx.Statement(1) != nil {
			n.ElseStatement = ctx.Statement(1).Accept(v).(ast.Node)
		}
		return n
	} else if s := ctx.SWITCH(); s != nil {
		n := &ast.Switch{Position: v.newPosition(ctx)}
		n.Expression = ctx.Expression().Accept(v).(ast.Node)
		n.WhenStatements = ctx.WhenStatements().Accept(v).([]ast.Node)
		if b := ctx.Block(); b != nil {
			n.ElseStatement = b.Accept(v).(ast.Node)
		}
		return n
	} else if s := ctx.FOR(); s != nil {
		n := &ast.For{Position: v.newPosition(ctx)}
		n.Control = ctx.ForControl().Accept(v).(ast.Node)
		n.Statements = ctx.Statement(0).Accept(v).(ast.Node)
		return n
	} else if s := ctx.WHILE(); s != nil {
		n := &ast.While{Position: v.newPosition(ctx)}
		n.Condition = ctx.ParExpression().Accept(v).(ast.Node)
		statements := ctx.AllStatement()
		n.Statements = make([]ast.Node, len(statements))
		for i, statement := range statements {
			n.Statements[i] = statement.Accept(v).(ast.Node)
		}
		n.IsDo = ctx.DO() != nil
		return n
	} else if s := ctx.RETURN(); s != nil {
		n := &ast.Return{Position: v.newPosition(ctx)}
		if e := ctx.Expression(); e != nil {
			n.Expression = e.Accept(v).(ast.Node)
		}
		return n
	} else if s := ctx.THROW(); s != nil {
		n := &ast.Throw{Position: v.newPosition(ctx)}
		n.Expression = ctx.Expression().Accept(v).(ast.Node)
		return n
	} else if s := ctx.BREAK(); s != nil {
		return &ast.Break{Position: v.newPosition(ctx)}
	} else if s := ctx.CONTINUE(); s != nil {
		return &ast.Continue{Position: v.newPosition(ctx)}
	} else if s := ctx.BREAK(); s != nil {
		return &ast.Break{Position: v.newPosition(ctx)}
	} else if s := ctx.StatementExpression(); s != nil {
		return s.Accept(v)
	} else if s := ctx.ApexDbExpression(); s != nil {
		return s.Accept(v)
	} else if s := ctx.Block(); s != nil {
		return s.Accept(v)
	}
	return &ast.NothingStatement{Position: v.newPosition(ctx)}
}

// goal

func (v *AstBuilder) VisitPropertyBlock(ctx *parser.PropertyBlockContext) interface{} {
	n := &ast.GetterSetter{Position: v.newPosition(ctx)}
	if ctx.Getter() != nil {
		n.Type = ctx.Getter().Accept(v).(string)
	} else {
		n.Type = ctx.Setter().Accept(v).(string)
	}
	modifiers := ctx.AllModifier()
	n.Modifiers = make([]ast.Node, len(modifiers))
	for i, m := range modifiers {
		n.Modifiers[i] = m.Accept(v).(ast.Node)
	}
	return n
}

func (v *AstBuilder) VisitGetter(ctx *parser.GetterContext) interface{} {
	return ctx.GetText()
}

func (v *AstBuilder) VisitSetter(ctx *parser.SetterContext) interface{} {
	return ctx.GetText()
}

func (v *AstBuilder) VisitCatchClause(ctx *parser.CatchClauseContext) interface{} {
	c := &ast.Catch{Position: v.newPosition(ctx)}
	c.Type = ctx.CatchType().Accept(v).(ast.Node)
	c.Identifier = ctx.ApexIdentifier().GetText()
	modifiers := ctx.AllVariableModifier()
	c.Modifiers = make([]ast.Node, len(modifiers))
	for i, m := range modifiers {
		c.Modifiers[i] = m.Accept(v).(ast.Node)
	}
	c.Block = ctx.Block().Accept(v).(*ast.Block)
	return c
}

func (v *AstBuilder) VisitCatchType(ctx *parser.CatchTypeContext) interface{} {
	names := ctx.AllQualifiedName()
	return names[0].Accept(v)
}

func (v *AstBuilder) VisitFinallyBlock(ctx *parser.FinallyBlockContext) interface{} {
	return ctx.Block().Accept(v)
}

func (v *AstBuilder) VisitWhenStatements(ctx *parser.WhenStatementsContext) interface{} {
	whenStatements := ctx.AllWhenStatement()
	statements := make([]ast.Node, len(whenStatements))
	for i, s := range whenStatements {
		statements[i] = s.Accept(v).(ast.Node)
	}
	return statements
}

func (v *AstBuilder) VisitWhenStatement(ctx *parser.WhenStatementContext) interface{} {
	n := &ast.When{Position: v.newPosition(ctx)}
	n.Condition = ctx.WhenExpression().Accept(v).([]ast.Node)
	n.Statements = ctx.Block().Accept(v).(*ast.Block)
	return n
}

func (v *AstBuilder) VisitWhenExpression(ctx *parser.WhenExpressionContext) interface{} {
	if literals := ctx.AllLiteral(); len(literals) != 0 {
		expressions := make([]ast.Node, len(literals))
		for i, l := range literals {
			expressions[i] = l.Accept(v).(ast.Node)
		}
		return expressions
	}
	n := &ast.WhenType{Position: v.newPosition(ctx)}
	n.Type = ctx.ApexType().Accept(v).(ast.Node)
	n.Identifier = ctx.ApexIdentifier().GetText()
	return []ast.Node{n}
}

func (v *AstBuilder) VisitForControl(ctx *parser.ForControlContext) interface{} {
	if c := ctx.EnhancedForControl(); c != nil {
		return c.Accept(v)
	}
	c := &ast.ForControl{Position: v.newPosition(ctx)}
	if f := ctx.ForInit(); f != nil {
		c.ForInit = f.Accept(v).(ast.Node)
	}
	if e := ctx.Expression(); e != nil {
		c.Expression = e.Accept(v).(ast.Node)
	}
	if u := ctx.ForUpdate(); u != nil {
		c.ForUpdate = u.Accept(v).([]ast.Node)
	}
	return c
}

func (v *AstBuilder) VisitForInit(ctx *parser.ForInitContext) interface{} {
	if d := ctx.LocalVariableDeclaration(); d != nil {
		return d.Accept(v)
	}
	return ctx.ExpressionList().Accept(v)
}

func (v *AstBuilder) VisitEnhancedForControl(ctx *parser.EnhancedForControlContext) interface{} {
	n := &ast.EnhancedForControl{Position: v.newPosition(ctx)}
	n.Type = ctx.ApexType().Accept(v).(ast.Node)
	n.VariableDeclaratorId = ctx.VariableDeclaratorId().Accept(v).(string)
	n.Expression = ctx.Expression().Accept(v).(ast.Node)
	modifiers := ctx.AllVariableModifier()
	n.Modifiers = make([]ast.Node, len(modifiers))
	for i, m := range modifiers {
		n.Modifiers[i] = m.Accept(v).(ast.Node)
	}
	return n
}

func (v *AstBuilder) VisitForUpdate(ctx *parser.ForUpdateContext) interface{} {
	return ctx.ExpressionList().Accept(v)
}

func (v *AstBuilder) VisitParExpression(ctx *parser.ParExpressionContext) interface{} {
	return ctx.Expression().Accept(v)
}

func (v *AstBuilder) VisitExpressionList(ctx *parser.ExpressionListContext) interface{} {
	expressions := ctx.AllExpression()
	nodes := make([]ast.Node, len(expressions))
	for i, e := range expressions {
		nodes[i] = e.Accept(v).(ast.Node)
	}
	return nodes
}

func (v *AstBuilder) VisitStatementExpression(ctx *parser.StatementExpressionContext) interface{} {
	return ctx.Expression().Accept(v)
}

func (v *AstBuilder) VisitConstantExpression(ctx *parser.ConstantExpressionContext) interface{} {
	return ctx.Expression().Accept(v)
}

func (v *AstBuilder) VisitApexDbExpressionShort(ctx *parser.ApexDbExpressionShortContext) interface{} {
	n := &ast.Dml{Position: v.newPosition(ctx)}
	n.Type = ctx.GetDml().GetText()
	if ident := ctx.ApexIdentifier(); ident != nil {
		n.UpsertKey = ident.Accept(v).(string)
	}
	n.Expression = ctx.Expression().Accept(v).(ast.Node)
	return n
}

func (v *AstBuilder) VisitApexDbExpression(ctx *parser.ApexDbExpressionContext) interface{} {
	return ctx.ApexDbExpressionShort().Accept(v)
}

func (v *AstBuilder) VisitTernalyExpression(ctx *parser.TernalyExpressionContext) interface{} {
	n := &ast.TernalyExpression{Position: v.newPosition(ctx)}
	n.Condition = ctx.Expression(0).Accept(v).(ast.Node)
	n.TrueExpression = ctx.Expression(1).Accept(v).(ast.Node)
	n.FalseExpression = ctx.Expression(2).Accept(v).(ast.Node)
	return n
}

func (v *AstBuilder) VisitPreUnaryExpression(ctx *parser.PreUnaryExpressionContext) interface{} {
	n := &ast.UnaryOperator{Position: v.newPosition(ctx)}
	n.Op = ctx.GetOp().GetText()
	n.Expression = ctx.Expression().Accept(v).(ast.Node)
	n.IsPrefix = true
	return n
}

func (v *AstBuilder) VisitArrayAccess(ctx *parser.ArrayAccessContext) interface{} {
	n := &ast.ArrayAccess{Position: v.newPosition(ctx)}
	n.Receiver = ctx.Expression(0).Accept(v).(ast.Node)
	n.Key = ctx.Expression(1).Accept(v).(ast.Node)
	return n
}

func (v *AstBuilder) VisitPostUnaryExpression(ctx *parser.PostUnaryExpressionContext) interface{} {
	n := &ast.UnaryOperator{Position: v.newPosition(ctx)}
	n.Op = ctx.GetOp().GetText()
	n.Expression = ctx.Expression().Accept(v).(ast.Node)
	n.IsPrefix = false
	return n
}

func (v *AstBuilder) VisitPrimaryExpression(ctx *parser.PrimaryExpressionContext) interface{} {
	return ctx.Primary().Accept(v)
}

func (v *AstBuilder) VisitOpExpression(ctx *parser.OpExpressionContext) interface{} {
	n := &ast.BinaryOperator{Position: v.newPosition(ctx)}
	n.Op = ctx.GetOp().GetText()
	n.Left = ctx.Expression(0).Accept(v).(ast.Node)
	n.Right = ctx.Expression(1).Accept(v).(ast.Node)
	return n
}

func (v *AstBuilder) VisitNewExpression(ctx *parser.NewObjectExpressionContext) interface{} {
	return ctx.Creator().Accept(v)
}

func (v *AstBuilder) VisitUnaryExpression(ctx *parser.UnaryExpressionContext) interface{} {
	n := &ast.UnaryOperator{Position: v.newPosition(ctx)}
	n.Op = ctx.GetOp().GetText()
	n.Expression = ctx.Expression().Accept(v).(ast.Node)
	return n
}

func (v *AstBuilder) VisitMethodInvocation(ctx *parser.MethodInvocationContext) interface{} {
	n := &ast.MethodInvocation{Position: v.newPosition(ctx)}
	n.NameOrExpression = ctx.Expression().Accept(v).(ast.Node)
	if list := ctx.ExpressionList(); list != nil {
		n.Parameters = list.Accept(v).([]ast.Node)
	} else {
		n.Parameters = []ast.Node{}
	}
	return n
}

func (v *AstBuilder) VisitCastExpression(ctx *parser.CastExpressionContext) interface{} {
	n := &ast.CastExpression{Position: v.newPosition(ctx)}
	n.CastType = ctx.ApexType().Accept(v).(ast.Node)
	n.Expression = ctx.Expression().Accept(v).(ast.Node)
	return n
}

func (v *AstBuilder) VisitShiftExpression(ctx *parser.ShiftExpressionContext) interface{} {
	n := &ast.BinaryOperator{Position: v.newPosition(ctx)}
	ops := []string{}
	for i, o := range ctx.GetOp() {
		ops[i] = o.GetText()
	}
	n.Op = strings.Join(ops, "")
	n.Left = ctx.Expression(0).Accept(v).(ast.Node)
	n.Right = ctx.Expression(1).Accept(v).(ast.Node)
	return n
}

func (v *AstBuilder) VisitFieldAccess(ctx *parser.FieldAccessContext) interface{} {
	expression := ctx.Expression().Accept(v).(ast.Node)
	fieldName := ctx.ApexIdentifier().Accept(v).(string)
	return &ast.FieldAccess{
		Expression: expression,
		FieldName:  fieldName,
		Position:   v.newPosition(ctx),
	}
}

func (v *AstBuilder) VisitPrimary(ctx *parser.PrimaryContext) interface{} {
	if e := ctx.Expression(); e != nil {
		return e.Accept(v)
	} else if t := ctx.THIS(); t != nil {
		return t.Accept(v)
	} else if s := ctx.SUPER(); s != nil {
		return s.Accept(v)
	} else if l := ctx.Literal(); l != nil {
		return l.Accept(v)
	} else if i := ctx.ApexIdentifier(); i != nil {
		n := &ast.Name{Position: v.newPosition(ctx)}
		n.Value = i.Accept(v).(string)
		return n
	} else if l := ctx.SoqlLiteral(); l != nil {
		return l.Accept(v)
	} else if l := ctx.SoslLiteral(); l != nil {
		return l.Accept(v)
	} else if t := ctx.ApexType(); t != nil {
		return t.Accept(v)
	}
	n := &ast.Name{Position: v.newPosition(ctx)}
	n.Value = ctx.PrimitiveType().GetText()
	return n
}

func (v *AstBuilder) VisitCreator(ctx *parser.CreatorContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *AstBuilder) VisitCreatedName(ctx *parser.CreatedNameContext) interface{} {
	if identifiers := ctx.AllApexIdentifier(); len(identifiers) != 0 {
		n := &ast.Type{Position: v.newPosition(ctx)}
		if types := ctx.AllTypeArgumentsOrDiamond(); len(types) != 0 {
			// n.Parameters = ctx.TypeArgumentsOrDiamond(0).Accept(v)
		}
		names := make([]ast.Node, len(identifiers))
		for i, ident := range identifiers {
			names[i] = ident.Accept(v).(ast.Node)
		}
		// TODO: implement
		return n
	}
	return nil
}

func (v *AstBuilder) VisitInnerCreator(ctx *parser.InnerCreatorContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *AstBuilder) VisitArrayCreatorRest(ctx *parser.ArrayCreatorRestContext) interface{} {
	n := &ast.ArrayCreator{Position: v.newPosition(ctx)}
	n.Dim = len(ctx.AllTypedArray())
	if init := ctx.ArrayInitializer(); init != nil {
		n.ArrayInitializer = init.Accept(v).(ast.Node)
	}
	expressions := ctx.AllExpression()
	n.Expressions = make([]ast.Node, len(expressions))
	for i, e := range expressions {
		n.Expressions[i] = e.Accept(v).(ast.Node)
	}
	return n
}

func (v *AstBuilder) VisitMapCreatorRest(ctx *parser.MapCreatorRestContext) interface{} {
	return &ast.MapCreator{Position: v.newPosition(ctx)}
}

func (v *AstBuilder) VisitSetCreatorRest(ctx *parser.SetCreatorRestContext) interface{} {
	return &ast.SetCreator{Position: v.newPosition(ctx)}
}

func (v *AstBuilder) VisitClassCreatorRest(ctx *parser.ClassCreatorRestContext) interface{} {
	return ctx.Arguments().Accept(v)
}

func (v *AstBuilder) VisitExplicitGenericInvocation(ctx *parser.ExplicitGenericInvocationContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *AstBuilder) VisitNonWildcardTypeArguments(ctx *parser.NonWildcardTypeArgumentsContext) interface{} {
	return ctx.TypeList().Accept(v)
}

func (v *AstBuilder) VisitTypeArgumentsOrDiamond(ctx *parser.TypeArgumentsOrDiamondContext) interface{} {
	return ctx.TypeArguments().Accept(v)
}

func (v *AstBuilder) VisitNonWildcardTypeArgumentsOrDiamond(ctx *parser.NonWildcardTypeArgumentsOrDiamondContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *AstBuilder) VisitSuperSuffix(ctx *parser.SuperSuffixContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *AstBuilder) VisitExplicitGenericInvocationSuffix(ctx *parser.ExplicitGenericInvocationSuffixContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *AstBuilder) VisitArguments(ctx *parser.ArgumentsContext) interface{} {
	if list := ctx.ExpressionList(); list != nil {
		return list.Accept(v)
	}
	return []ast.Node{}
}

func (v *AstBuilder) VisitApexIdentifier(ctx *parser.ApexIdentifierContext) interface{} {
	if t := ctx.Identifier(); t != nil {
		return t.GetText()
	} else if t := ctx.GET(); t != nil {
		return t.GetText()
	} else if t := ctx.SET(); t != nil {
		return t.GetText()
	} else if t := ctx.DATA(); t != nil {
		return t.GetText()
	} else if t := ctx.GROUP(); t != nil {
		return t.GetText()
	} else if t := ctx.DELETE(); t != nil {
		return t.GetText()
	} else if t := ctx.INSERT(); t != nil {
		return t.GetText()
	} else if t := ctx.UPDATE(); t != nil {
		return t.GetText()
	} else if t := ctx.UPSERT(); t != nil {
		return t.GetText()
	} else if t := ctx.SCOPE(); t != nil {
		return t.GetText()
	} else if t := ctx.CATEGORY(); t != nil {
		return t.GetText()
	} else if t := ctx.REFERENCE(); t != nil {
		return t.GetText()
	} else if t := ctx.OFFSET(); t != nil {
		return t.GetText()
	} else if t := ctx.THEN(); t != nil {
		return t.GetText()
	} else if t := ctx.FIND(); t != nil {
		return t.GetText()
	} else if t := ctx.RETURNING(); t != nil {
		return t.GetText()
	} else if t := ctx.ALL(); t != nil {
		return t.GetText()
	} else if t := ctx.FIELDS(); t != nil {
		return t.GetText()
	} else if t := ctx.RUNAS(); t != nil {
		return t.GetText()
	} else if t := ctx.SYSTEM(); t != nil {
		return t.GetText()
	} else if t := ctx.PrimitiveType(); t != nil {
		return t.Accept(v)
	}
	return nil
}

func (v *AstBuilder) VisitTypeIdentifier(ctx *parser.TypeIdentifierContext) interface{} {
	if i := ctx.Identifier(); i != nil {
		return i.GetText()
	} else if t := ctx.GET(); t != nil {
		return t.GetText()
	} else if t := ctx.SET(); t != nil {
		return t.GetText()
	} else if t := ctx.DATA(); t != nil {
		return t.GetText()
	} else if t := ctx.GROUP(); t != nil {
		return t.GetText()
	} else if t := ctx.SCOPE(); t != nil {
		return t.GetText()
	}
	return nil
}

/**
 * SOQL
 */
func (v *AstBuilder) VisitSoqlLiteral(ctx *parser.SoqlLiteralContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *AstBuilder) VisitQuery(ctx *parser.QueryContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *AstBuilder) VisitSelectClause(ctx *parser.SelectClauseContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *AstBuilder) VisitFieldList(ctx *parser.FieldListContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *AstBuilder) VisitSelectField(ctx *parser.SelectFieldContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *AstBuilder) VisitFromClause(ctx *parser.FromClauseContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *AstBuilder) VisitFilterScope(ctx *parser.FilterScopeContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *AstBuilder) VisitSoqlFieldReference(ctx *parser.SoqlFieldReferenceContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *AstBuilder) VisitSoqlFunctionCall(ctx *parser.SoqlFunctionCallContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *AstBuilder) VisitSubquery(ctx *parser.SubqueryContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *AstBuilder) VisitWhereClause(ctx *parser.WhereClauseContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *AstBuilder) VisitWhereFields(ctx *parser.WhereFieldsContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *AstBuilder) VisitWhereField(ctx *parser.WhereFieldContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *AstBuilder) VisitLimitClause(ctx *parser.LimitClauseContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *AstBuilder) VisitOrderClause(ctx *parser.OrderClauseContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *AstBuilder) VisitBindVariable(ctx *parser.BindVariableContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *AstBuilder) VisitSoqlValue(ctx *parser.SoqlValueContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *AstBuilder) VisitWithClause(ctx *parser.WithClauseContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *AstBuilder) VisitSoqlFilteringExpression(ctx *parser.SoqlFilteringExpressionContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *AstBuilder) VisitGroupClause(ctx *parser.GroupClauseContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *AstBuilder) VisitFieldGroupList(ctx *parser.FieldGroupListContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *AstBuilder) VisitHavingConditionExpression(ctx *parser.HavingConditionExpressionContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *AstBuilder) VisitOffsetClause(ctx *parser.OffsetClauseContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *AstBuilder) VisitViewClause(ctx *parser.ViewClauseContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *AstBuilder) VisitSoslLiteral(ctx *parser.SoslLiteralContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *AstBuilder) VisitSoslQuery(ctx *parser.SoslQueryContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *AstBuilder) VisitSoslReturningObject(ctx *parser.SoslReturningObjectContext) interface{} {
	return v.VisitChildren(ctx)
}

type PositionContext interface {
	GetStart() antlr.Token
}

func (v *AstBuilder) newPosition(ctx PositionContext) *ast.Position {
	return &ast.Position{
		FileName: v.CurrentFile,
		Column:   ctx.GetStart().GetColumn(),
		Line:     ctx.GetStart().GetLine(),
	}
}
