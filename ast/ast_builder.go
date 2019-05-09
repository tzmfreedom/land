package ast

import (
	"strings"

	"strconv"

	"github.com/antlr/antlr4/runtime/Go/antlr"
	"github.com/tzmfreedom/land/parser"
)

type Builder struct {
	*parser.BaseapexVisitor
	Source string
}

func (v *Builder) VisitCompilationUnit(ctx *parser.CompilationUnitContext) interface{} {
	return ctx.TypeDeclaration().Accept(v)
}

func (v *Builder) VisitTypeDeclaration(ctx *parser.TypeDeclarationContext) interface{} {
	classOrInterfaceModifiers := ctx.AllClassOrInterfaceModifier()
	modifiers := []*Modifier{}
	annotations := []*Annotation{}
	for _, classOrInterfaceModifier := range classOrInterfaceModifiers {
		r := classOrInterfaceModifier.Accept(v)
		switch n := r.(type) {
		case *Modifier:
			modifiers = append(modifiers, n)
		case *Annotation:
			annotations = append(annotations, n)
		}
	}

	if n := ctx.ClassDeclaration(); n != nil {
		cd := n.Accept(v)
		decl := cd.(*ClassDeclaration)
		decl.Modifiers = setParentNodeToModifiers(modifiers, decl)
		decl.Annotations = setParentNodeToAnnotations(annotations, decl)
		return decl
	} else if n := ctx.TriggerDeclaration(); n != nil {
		return n.Accept(v)
	} else if n := ctx.InterfaceDeclaration(); n != nil {
		cd := n.Accept(v)
		decl := cd.(*InterfaceDeclaration)
		decl.Modifiers = setParentNodeToModifiers(modifiers, decl)
		decl.Annotations = setParentNodeToAnnotations(annotations, decl)
		return decl
	} else if n := ctx.EnumDeclaration(); n != nil {
		return n.Accept(v)
	}
	panic("not pass")
}

func (v *Builder) VisitTriggerDeclaration(ctx *parser.TriggerDeclarationContext) interface{} {
	n := &Trigger{Location: v.newLocation(ctx)}
	n.Name = ctx.ApexIdentifier(0).GetText()
	n.Object = ctx.ApexIdentifier(1).GetText()
	n.TriggerTimings = ctx.TriggerTimings().Accept(v).([]Node)
	n.Statements = ctx.Block().Accept(v).(*Block)
	return n
}

func (v *Builder) VisitTriggerTimings(ctx *parser.TriggerTimingsContext) interface{} {
	allTimings := ctx.AllTriggerTiming()
	timings := make([]Node, len(allTimings))
	for i, timing := range allTimings {
		timings[i] = timing.Accept(v).(Node)
	}
	return timings
}

func (v *Builder) VisitTriggerTiming(ctx *parser.TriggerTimingContext) interface{} {
	return &TriggerTiming{
		Timing:   ctx.GetTiming().GetText(),
		Dml:      ctx.GetDml().GetText(),
		Location: v.newLocation(ctx),
	}
}

func (v *Builder) VisitModifier(ctx *parser.ModifierContext) interface{} {
	m := ctx.ClassOrInterfaceModifier()
	if m != nil {
		return m.Accept(v)
	}
	return &Modifier{
		Name:     ctx.GetText(),
		Location: v.newLocation(ctx),
	}
}

func (v *Builder) VisitClassOrInterfaceModifier(ctx *parser.ClassOrInterfaceModifierContext) interface{} {
	annotation := ctx.Annotation()
	if annotation != nil {
		return ctx.Annotation().Accept(v)
	}
	return &Modifier{
		Name:     ctx.GetText(),
		Location: v.newLocation(ctx),
	}
}

func (v *Builder) VisitVariableModifier(ctx *parser.VariableModifierContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *Builder) VisitClassDeclaration(ctx *parser.ClassDeclarationContext) interface{} {
	declarations := ctx.ClassBody().Accept(v).([]Node)

	n := &ClassDeclaration{
		Name:     ctx.ApexIdentifier().GetText(),
		Location: v.newLocation(ctx),
	}
	if t := ctx.ApexType(); t != nil {
		n.SuperClassRef = t.Accept(v).(*TypeRef)
	}
	if tl := ctx.TypeList(); tl != nil {
		n.ImplementClassRefs = tl.Accept(v).([]*TypeRef)
	}
	n.Declarations = make([]Node, len(declarations))
	for i, d := range declarations {
		d.SetParent(n)
		n.Declarations[i] = d
	}
	return n
}

func (v *Builder) VisitEnumDeclaration(ctx *parser.EnumDeclarationContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *Builder) VisitEnumConstants(ctx *parser.EnumConstantsContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *Builder) VisitEnumConstant(ctx *parser.EnumConstantContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *Builder) VisitEnumBodyDeclarations(ctx *parser.EnumBodyDeclarationsContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *Builder) VisitInterfaceDeclaration(ctx *parser.InterfaceDeclarationContext) interface{} {
	n := &InterfaceDeclaration{
		Name:     ctx.ApexIdentifier().GetText(),
		Location: v.newLocation(ctx),
	}
	n.Methods = ctx.InterfaceBody().Accept(v).([]*MethodDeclaration)
	setParentNodeToMethods(n.Methods, n)
	return n
}

func (v *Builder) VisitTypeList(ctx *parser.TypeListContext) interface{} {
	apexTypes := ctx.AllApexType()
	types := make([]*TypeRef, len(apexTypes))
	for i, t := range apexTypes {
		types[i] = t.Accept(v).(*TypeRef)
	}
	return types
}

func (v *Builder) VisitClassBody(ctx *parser.ClassBodyContext) interface{} {
	bodyDeclarations := ctx.AllClassBodyDeclaration()
	declarations := make([]Node, len(bodyDeclarations))
	for i, d := range bodyDeclarations {
		declarations[i] = d.Accept(v).(Node)
	}
	return declarations
}

func (v *Builder) VisitInterfaceBody(ctx *parser.InterfaceBodyContext) interface{} {
	bodyDeclarations := ctx.AllInterfaceBodyDeclaration()
	declarations := make([]*MethodDeclaration, len(bodyDeclarations))
	for i, d := range bodyDeclarations {
		declarations[i] = d.Accept(v).(*MethodDeclaration)
	}
	return declarations
}

func (v *Builder) VisitClassBodyDeclaration(ctx *parser.ClassBodyDeclarationContext) interface{} {
	memberDeclaration := ctx.MemberDeclaration()
	if memberDeclaration != nil {
		declaration := memberDeclaration.Accept(v)

		modifiers := ctx.AllModifier()
		declarationModifiers := []*Modifier{}
		declarationAnnotations := []*Annotation{}
		for _, m := range modifiers {
			n := m.Accept(v)
			switch modifier := n.(type) {
			case *Modifier:
				declarationModifiers = append(declarationModifiers, modifier)
			case *Annotation:
				declarationAnnotations = append(declarationAnnotations, modifier)
			}
		}
		switch decl := declaration.(type) {
		case *MethodDeclaration:
			decl.Modifiers = declarationModifiers
			decl.Annotations = declarationAnnotations
			return decl
		case *FieldDeclaration:
			decl.Modifiers = declarationModifiers
			decl.Annotations = declarationAnnotations
			return decl
		case *ConstructorDeclaration:
			decl.Modifiers = declarationModifiers
			decl.Annotations = declarationAnnotations
			return decl
		case *InterfaceDeclaration:
			decl.Modifiers = declarationModifiers
			decl.Annotations = declarationAnnotations
			return decl
		case *ClassDeclaration:
			decl.Modifiers = declarationModifiers
			decl.Annotations = declarationAnnotations
			return decl
		//case *EnumDeclaration:
		//	decl.Modifiers = declarationModifiers
		//	return decl
		case *PropertyDeclaration:
			decl.Modifiers = declarationModifiers
			decl.Annotations = declarationAnnotations
			return decl
		}
	}
	return nil
}

func (v *Builder) VisitMemberDeclaration(ctx *parser.MemberDeclarationContext) interface{} {
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

func (v *Builder) VisitMethodDeclaration(ctx *parser.MethodDeclarationContext) interface{} {
	n := &MethodDeclaration{Location: v.newLocation(ctx)}
	n.Name = ctx.ApexIdentifier().GetText()
	if ctx.ApexType() != nil {
		n.ReturnType = ctx.ApexType().Accept(v).(*TypeRef)
	} else {
		n.ReturnType = nil
	}
	n.Parameters = ctx.FormalParameters().Accept(v).([]*Parameter)
	if ctx.QualifiedNameList() != nil {
		n.Throws = ctx.QualifiedNameList().Accept(v).([]Node)
		setParentNodeToNodes(n.Throws, n)
	} else {
		n.Throws = []Node{}
	}
	if ctx.MethodBody() != nil {
		n.Statements = ctx.MethodBody().Accept(v).(*Block)
		n.Statements.SetParent(n)
	} else {
		n.Statements = nil
	}
	return n
}

func (v *Builder) VisitConstructorDeclaration(ctx *parser.ConstructorDeclarationContext) interface{} {
	parameters := ctx.FormalParameters().Accept(v).([]*Parameter)
	var throws []Node
	if q := ctx.QualifiedNameList(); q != nil {
		throws = q.Accept(v).([]Node)
	} else {
		throws = []Node{}
	}
	name := ctx.ApexIdentifier().Accept(v).(string)
	body := ctx.ConstructorBody().Accept(v).(*Block)
	return &ConstructorDeclaration{
		Name:       name,
		Parameters: parameters,
		Throws:     throws,
		Statements: body,
		Location:   v.newLocation(ctx),
	}
}

func (v *Builder) VisitFieldDeclaration(ctx *parser.FieldDeclarationContext) interface{} {
	t := ctx.ApexType().Accept(v).(*TypeRef)
	d := ctx.VariableDeclarators().Accept(v).([]*VariableDeclarator)
	return &FieldDeclaration{
		TypeRef:     t,
		Declarators: d,
	}
}

func (v *Builder) VisitPropertyDeclaration(ctx *parser.PropertyDeclarationContext) interface{} {
	t := ctx.ApexType().Accept(v).(*TypeRef)
	d := ctx.VariableDeclaratorId().Accept(v).(string)
	getterSetters := ctx.PropertyBodyDeclaration().Accept(v).([]*GetterSetter)
	return &PropertyDeclaration{
		TypeRef:       t,
		Identifier:    d,
		GetterSetters: getterSetters,
	}
}

func (v *Builder) VisitPropertyBodyDeclaration(ctx *parser.PropertyBodyDeclarationContext) interface{} {
	blocks := ctx.AllPropertyBlock()
	declarations := make([]*GetterSetter, len(blocks))
	for i, b := range blocks {
		declarations[i] = b.Accept(v).(*GetterSetter)
	}
	return declarations
}

func (v *Builder) VisitInterfaceBodyDeclaration(ctx *parser.InterfaceBodyDeclarationContext) interface{} {
	memberDeclaration := ctx.InterfaceMemberDeclaration()
	if memberDeclaration != nil {
		declaration := memberDeclaration.Accept(v)

		modifiers := ctx.AllModifier()
		declarationModifiers := []*Modifier{}
		declarationAnnotations := []*Annotation{}
		for _, m := range modifiers {
			n := m.Accept(v)
			switch modifier := n.(type) {
			case *Modifier:
				declarationModifiers = append(declarationModifiers, modifier)
			case *Annotation:
				declarationAnnotations = append(declarationAnnotations, modifier)
			}
		}
		switch decl := declaration.(type) {
		case *MethodDeclaration:
			decl.Modifiers = declarationModifiers
			decl.Annotations = declarationAnnotations
			return decl
		case *InterfaceDeclaration:
			decl.Modifiers = declarationModifiers
			decl.Annotations = declarationAnnotations
			return decl
		case *ClassDeclaration:
			decl.Modifiers = declarationModifiers
			decl.Annotations = declarationAnnotations
			return decl
		}
	}
	return nil
}

func (v *Builder) VisitInterfaceMemberDeclaration(ctx *parser.InterfaceMemberDeclarationContext) interface{} {
	if decl := ctx.ConstDeclaration(); decl != nil {
		return decl.Accept(v)
	}
	if decl := ctx.InterfaceDeclaration(); decl != nil {
		return decl.Accept(v)
	}
	if decl := ctx.InterfaceMethodDeclaration(); decl != nil {
		return decl.Accept(v)
	}
	if decl := ctx.ClassDeclaration(); decl != nil {
		return decl.Accept(v)
	}
	if decl := ctx.EnumDeclaration(); decl != nil {
		return decl.Accept(v)
	}
	panic("not pass")
}

func (v *Builder) VisitConstDeclaration(ctx *parser.ConstDeclarationContext) interface{} {
	_ = ctx.ApexType().Accept(v)
	_ = ctx.AllConstantDeclarator()

	// TODO: implement
	return nil
}

func (v *Builder) VisitConstantDeclarator(ctx *parser.ConstantDeclaratorContext) interface{} {
	_ = ctx.ApexIdentifier().Accept(v)
	_ = ctx.VariableInitializer().Accept(v)

	// TODO: implement
	return nil
}

func (v *Builder) VisitInterfaceMethodDeclaration(ctx *parser.InterfaceMethodDeclarationContext) interface{} {
	decl := &MethodDeclaration{Location: v.newLocation(ctx)}
	decl.Name = ctx.ApexIdentifier().Accept(v).(string)

	if t := ctx.ApexType(); t != nil {
		decl.ReturnType = t.Accept(v).(*TypeRef)
	} else {
		// TODO: implement void
	}
	decl.Parameters = ctx.FormalParameters().Accept(v).([]*Parameter)
	if q := ctx.QualifiedNameList(); q != nil {
		decl.Throws = q.Accept(v).([]Node)
	} else {
		decl.Throws = []Node{}
	}
	return decl
}

func (v *Builder) VisitVariableDeclarators(ctx *parser.VariableDeclaratorsContext) interface{} {
	variableDeclarators := ctx.AllVariableDeclarator()
	declarators := make([]*VariableDeclarator, len(variableDeclarators))
	for i, d := range variableDeclarators {
		declarators[i] = d.Accept(v).(*VariableDeclarator)
	}
	return declarators
}

func (v *Builder) VisitVariableDeclarator(ctx *parser.VariableDeclaratorContext) interface{} {
	decl := &VariableDeclarator{Location: v.newLocation(ctx)}
	decl.Name = ctx.VariableDeclaratorId().Accept(v).(string)
	if init := ctx.VariableInitializer(); init != nil {
		decl.Expression = init.Accept(v).(Node)
		decl.Expression.SetParent(decl)
	}
	return decl
}

func (v *Builder) VisitVariableDeclaratorId(ctx *parser.VariableDeclaratorIdContext) interface{} {
	return ctx.ApexIdentifier().GetText()
}

func (v *Builder) VisitVariableInitializer(ctx *parser.VariableInitializerContext) interface{} {
	if init := ctx.ArrayInitializer(); init != nil {
		return init.Accept(v)
	}
	return ctx.Expression().Accept(v)
}

func (v *Builder) VisitArrayInitializer(ctx *parser.ArrayInitializerContext) interface{} {
	if inits := ctx.AllVariableInitializer(); len(inits) != 0 {
		records := make([]Node, len(inits))
		for i, init := range inits {
			records[i] = init.Accept(v).(Node)
		}
		return &Init{Records: records}
	}
	return &Init{}
}

func (v *Builder) VisitEnumConstantName(ctx *parser.EnumConstantNameContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *Builder) VisitApexType(ctx *parser.ApexTypeContext) interface{} {
	if interfaceType := ctx.ClassOrInterfaceType(); interfaceType != nil {
		t := interfaceType.Accept(v).(*TypeRef)
		t.Dimmension = len(ctx.AllTypedArray())
		return t
	} else if primitiveType := ctx.PrimitiveType(); primitiveType != nil {
		t := primitiveType.Accept(v).(*TypeRef)
		t.Dimmension = len(ctx.AllTypedArray())
		return t
	}
	return nil
}

func (v *Builder) VisitTypedArray(ctx *parser.TypedArrayContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *Builder) VisitClassOrInterfaceType(ctx *parser.ClassOrInterfaceTypeContext) interface{} {
	t := &TypeRef{Location: v.newLocation(ctx)}
	if len(ctx.AllTypeArguments()) != 0 {
		t.Parameters = ctx.TypeArguments(0).Accept(v).([]*TypeRef)
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

func (v *Builder) VisitPrimitiveType(ctx *parser.PrimitiveTypeContext) interface{} {
	return &TypeRef{
		Name:       []string{ctx.GetText()},
		Parameters: []*TypeRef{},
		Location:   v.newLocation(ctx),
	}
}

func (v *Builder) VisitTypeArguments(ctx *parser.TypeArgumentsContext) interface{} {
	arguments := ctx.AllTypeArgument()
	typeArguments := make([]*TypeRef, len(arguments))
	for i, a := range arguments {
		typeArguments[i] = a.Accept(v).(*TypeRef)
	}
	return typeArguments
}

func (v *Builder) VisitTypeArgument(ctx *parser.TypeArgumentContext) interface{} {
	return ctx.ApexType().Accept(v)
}

func (v *Builder) VisitQualifiedNameList(ctx *parser.QualifiedNameListContext) interface{} {
	qualifiedNames := ctx.AllQualifiedName()
	names := make([]Node, len(qualifiedNames))
	for i, qn := range qualifiedNames {
		names[i] = qn.Accept(v).(Node)
	}
	return names
}

func (v *Builder) VisitFormalParameters(ctx *parser.FormalParametersContext) interface{} {
	if p := ctx.FormalParameterList(); p != nil {
		return p.Accept(v)
	}
	return []*Parameter{}
}

func (v *Builder) VisitFormalParameterList(ctx *parser.FormalParameterListContext) interface{} {
	formalParameters := ctx.AllFormalParameter()
	parameters := make([]*Parameter, len(formalParameters))
	for i, p := range formalParameters {
		parameters[i] = p.Accept(v).(*Parameter)
	}
	return parameters
}

func (v *Builder) VisitFormalParameter(ctx *parser.FormalParameterContext) interface{} {
	p := &Parameter{Location: v.newLocation(ctx)}
	modifiers := ctx.AllVariableModifier()
	p.Modifiers = make([]*Modifier, len(modifiers))
	for i, m := range modifiers {
		p.Modifiers[i] = m.Accept(v).(*Modifier)
	}
	p.TypeRef = ctx.ApexType().Accept(v).(*TypeRef)
	p.Name = ctx.VariableDeclaratorId().Accept(v).(string)
	return p
}

func (v *Builder) VisitLastFormalParameter(ctx *parser.LastFormalParameterContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *Builder) VisitMethodBody(ctx *parser.MethodBodyContext) interface{} {
	return ctx.Block().Accept(v)
}

func (v *Builder) VisitConstructorBody(ctx *parser.ConstructorBodyContext) interface{} {
	return ctx.Block().Accept(v)
}

func (v *Builder) VisitQualifiedName(ctx *parser.QualifiedNameContext) interface{} {
	allIdentifiers := ctx.AllApexIdentifier()
	identifiers := make([]string, len(allIdentifiers))
	for i, identifier := range allIdentifiers {
		ident := identifier.Accept(v)
		identifiers[i] = ident.(string)
	}
	n := &TypeRef{Location: v.newLocation(ctx)}
	n.Name = identifiers
	return n
}

func (v *Builder) VisitLiteral(ctx *parser.LiteralContext) interface{} {
	if lit := ctx.IntegerLiteral(); lit != nil {
		val, err := strconv.Atoi(lit.GetText())
		if err != nil {
			panic(err)
		}
		return &IntegerLiteral{Value: val, Location: v.newLocation(ctx)}
	} else if lit := ctx.FloatingPointLiteral(); lit != nil {
		val, err := strconv.ParseFloat(lit.GetText(), 64)
		if err != nil {
			panic(err)
		}
		return &DoubleLiteral{Value: val, Location: v.newLocation(ctx)}
	} else if lit := ctx.StringLiteral(); lit != nil {
		str := lit.GetText()
		return &StringLiteral{Value: str[1 : len(str)-1], Location: v.newLocation(ctx)}
	} else if lit := ctx.BooleanLiteral(); lit != nil {
		return &BooleanLiteral{Value: strings.ToLower(lit.GetText()) == "true", Location: v.newLocation(ctx)}
	} else if lit := ctx.NullLiteral(); lit != nil {
		return &NullLiteral{Location: v.newLocation(ctx)}
	}
	return nil
}

func (v *Builder) VisitAnnotation(ctx *parser.AnnotationContext) interface{} {
	name := ctx.AnnotationName().Accept(v).(*TypeRef)
	annotation := &Annotation{}
	// TODO: implement annotationName
	annotation.Name = name.Name[0]
	annotation.Location = v.newLocation(ctx)
	return annotation
}

func (v *Builder) VisitAnnotationName(ctx *parser.AnnotationNameContext) interface{} {
	return ctx.QualifiedName().Accept(v)
}

func (v *Builder) VisitElementValuePairs(ctx *parser.ElementValuePairsContext) interface{} {
	elementValuePairs := ctx.AllElementValuePair()
	pairs := make([]Node, len(elementValuePairs))
	for i, p := range elementValuePairs {
		pairs[i] = p.Accept(v).(Node)
	}
	return pairs
}

func (v *Builder) VisitElementValuePair(ctx *parser.ElementValuePairContext) interface{} {
	ctx.ApexIdentifier().GetText()
	ctx.ElementValue().Accept(v)
	// TODO: implement
	return nil
}

func (v *Builder) VisitElementValue(ctx *parser.ElementValueContext) interface{} {
	if e := ctx.Expression(); e != nil {
		return e.Accept(v)
	} else if a := ctx.Annotation(); a != nil {
		return a.Accept(v)
	} else if init := ctx.ElementValueArrayInitializer(); init != nil {
		return init.Accept(v)
	}
	return nil
}

func (v *Builder) VisitElementValueArrayInitializer(ctx *parser.ElementValueArrayInitializerContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *Builder) VisitBlock(ctx *parser.BlockContext) interface{} {
	blk := &Block{Location: v.newLocation(ctx)}
	statements := ctx.AllBlockStatement()
	blk.Statements = make([]Node, len(statements))
	for i, statement := range statements {
		s := statement.Accept(v)
		blk.Statements[i] = s.(Node)
	}
	setParentNodeToNodes(blk.Statements, blk)
	return blk
}

func (v *Builder) VisitBlockStatement(ctx *parser.BlockStatementContext) interface{} {
	if s := ctx.Statement(); s != nil {
		return s.Accept(v)
	} else if s := ctx.LocalVariableDeclarationStatement(); s != nil {
		return s.Accept(v)
	} else if t := ctx.TypeDeclaration(); t != nil {
		return t.Accept(v)
	}
	return nil
}

func (v *Builder) VisitLocalVariableDeclarationStatement(ctx *parser.LocalVariableDeclarationStatementContext) interface{} {
	return ctx.LocalVariableDeclaration().Accept(v)
}

func (v *Builder) VisitLocalVariableDeclaration(ctx *parser.LocalVariableDeclarationContext) interface{} {
	decl := &VariableDeclaration{Location: v.newLocation(ctx)}
	modifiers := ctx.AllVariableModifier()
	decl.Modifiers = make([]*Modifier, len(modifiers))
	for i, m := range modifiers {
		decl.Modifiers[i] = m.Accept(v).(*Modifier)
	}
	decl.TypeRef = ctx.ApexType().Accept(v).(*TypeRef)
	decl.Declarators = ctx.VariableDeclarators().Accept(v).([]*VariableDeclarator)
	for _, d := range decl.Declarators {
		d.SetParent(decl)
	}
	return decl
}

func (v *Builder) VisitStatement(ctx *parser.StatementContext) interface{} {
	if t := ctx.TRY(); t != nil {
		try := &Try{Location: v.newLocation(ctx)}
		try.Block = ctx.Block().Accept(v).(*Block)
		try.Block.SetParent(try)
		if clauses := ctx.AllCatchClause(); len(clauses) != 0 {
			try.CatchClause = make([]*Catch, len(clauses))
			for i, c := range clauses {
				try.CatchClause[i] = c.Accept(v).(*Catch)
				try.CatchClause[i].SetParent(try)
			}
		} else {
			try.CatchClause = []*Catch{}
		}
		if b := ctx.FinallyBlock(); b != nil {
			try.FinallyBlock = b.Accept(v).(*Block)
			try.FinallyBlock.SetParent(try)
		}
		return try
	} else if t := ctx.IF(); t != nil {
		n := &If{Location: v.newLocation(ctx)}
		n.Condition = ctx.ParExpression().Accept(v).(Node)
		n.Condition.SetParent(n)
		n.IfStatement = ctx.Statement(0).Accept(v).(Node)
		if ctx.Statement(1) != nil {
			n.ElseStatement = ctx.Statement(1).Accept(v).(Node)
			n.ElseStatement.SetParent(n)
		}
		n.IfStatement.SetParent(n)
		return n
	} else if s := ctx.SWITCH(); s != nil {
		n := &Switch{Location: v.newLocation(ctx)}
		n.Expression = ctx.Expression().Accept(v).(Node)
		n.Expression.SetParent(n)
		n.WhenStatements = ctx.WhenStatements().Accept(v).([]*When)
		for _, when := range n.WhenStatements {
			when.SetParent(n)
		}
		if b := ctx.Block(); b != nil {
			n.ElseStatement = b.Accept(v).(*Block)
			n.ElseStatement.SetParent(n)
		}
		return n
	} else if s := ctx.FOR(); s != nil {
		n := &For{Location: v.newLocation(ctx)}
		n.Control = ctx.ForControl().Accept(v).(Node)
		n.Control.SetParent(n)
		n.Statements = ctx.Statement(0).Accept(v).(*Block)
		n.Statements.SetParent(n)
		return n
	} else if s := ctx.WHILE(); s != nil {
		n := &While{Location: v.newLocation(ctx)}
		n.Condition = ctx.ParExpression().Accept(v).(Node)
		n.Condition.SetParent(n)
		n.Statements = ctx.Statement(0).Accept(v).(*Block)
		n.Statements.SetParent(n)
		n.IsDo = ctx.DO() != nil
		return n
	} else if s := ctx.RETURN(); s != nil {
		n := &Return{Location: v.newLocation(ctx)}
		if e := ctx.Expression(); e != nil {
			n.Expression = e.Accept(v).(Node)
			n.Expression.SetParent(n)
		}
		return n
	} else if s := ctx.THROW(); s != nil {
		n := &Throw{Location: v.newLocation(ctx)}
		n.Expression = ctx.Expression().Accept(v).(Node)
		n.Expression.SetParent(n)
		return n
	} else if s := ctx.BREAK(); s != nil {
		return &Break{Location: v.newLocation(ctx)}
	} else if s := ctx.CONTINUE(); s != nil {
		return &Continue{Location: v.newLocation(ctx)}
	} else if s := ctx.BREAK(); s != nil {
		return &Break{Location: v.newLocation(ctx)}
	} else if s := ctx.StatementExpression(); s != nil {
		return s.Accept(v)
	} else if s := ctx.ApexDbExpression(); s != nil {
		return s.Accept(v)
	} else if s := ctx.Block(); s != nil {
		return s.Accept(v)
	}
	return &NothingStatement{Location: v.newLocation(ctx)}
}

// goal

func (v *Builder) VisitPropertyBlock(ctx *parser.PropertyBlockContext) interface{} {
	n := &GetterSetter{Location: v.newLocation(ctx)}
	if ctx.Getter() != nil {
		n.Type = ctx.Getter().Accept(v).(string)
	} else {
		n.Type = ctx.Setter().Accept(v).(string)
	}
	modifiers := ctx.AllModifier()
	n.Modifiers = make([]*Modifier, len(modifiers))
	for i, m := range modifiers {
		n.Modifiers[i] = m.Accept(v).(*Modifier)
	}
	return n
}

func (v *Builder) VisitGetter(ctx *parser.GetterContext) interface{} {
	return ctx.GET().GetText()
}

func (v *Builder) VisitSetter(ctx *parser.SetterContext) interface{} {
	return ctx.SET().GetText()
}

func (v *Builder) VisitCatchClause(ctx *parser.CatchClauseContext) interface{} {
	c := &Catch{Location: v.newLocation(ctx)}
	c.TypeRef = ctx.CatchType().Accept(v).(*TypeRef)
	c.Identifier = ctx.ApexIdentifier().GetText()
	modifiers := ctx.AllVariableModifier()
	c.Modifiers = make([]*Modifier, len(modifiers))
	for i, m := range modifiers {
		c.Modifiers[i] = m.Accept(v).(*Modifier)
	}
	c.Block = ctx.Block().Accept(v).(*Block)
	return c
}

func (v *Builder) VisitCatchType(ctx *parser.CatchTypeContext) interface{} {
	names := ctx.AllQualifiedName()
	return names[0].Accept(v)
}

func (v *Builder) VisitFinallyBlock(ctx *parser.FinallyBlockContext) interface{} {
	return ctx.Block().Accept(v)
}

func (v *Builder) VisitWhenStatements(ctx *parser.WhenStatementsContext) interface{} {
	whenStatements := ctx.AllWhenStatement()
	statements := make([]*When, len(whenStatements))
	for i, s := range whenStatements {
		statements[i] = s.Accept(v).(*When)
	}
	return statements
}

func (v *Builder) VisitWhenStatement(ctx *parser.WhenStatementContext) interface{} {
	n := &When{Location: v.newLocation(ctx)}
	n.Condition = ctx.WhenExpression().Accept(v).([]Node)
	n.Statements = ctx.Block().Accept(v).(*Block)
	return n
}

func (v *Builder) VisitWhenExpression(ctx *parser.WhenExpressionContext) interface{} {
	if literals := ctx.AllLiteral(); len(literals) != 0 {
		expressions := make([]Node, len(literals))
		for i, l := range literals {
			expressions[i] = l.Accept(v).(Node)
		}
		return expressions
	}
	n := &WhenType{Location: v.newLocation(ctx)}
	n.TypeRef = ctx.ApexType().Accept(v).(*TypeRef)
	n.Identifier = ctx.ApexIdentifier().GetText()
	return []Node{n}
}

func (v *Builder) VisitForControl(ctx *parser.ForControlContext) interface{} {
	if c := ctx.EnhancedForControl(); c != nil {
		return c.Accept(v)
	}
	c := &ForControl{Location: v.newLocation(ctx)}
	if f := ctx.ForInit(); f != nil {
		r := f.Accept(v)
		switch forInit := r.(type) {
		case Node:
			c.ForInit = []Node{forInit}
		case []Node:
			c.ForInit = forInit
		}
	}
	if e := ctx.Expression(); e != nil {
		c.Expression = e.Accept(v).(Node)
	}
	if u := ctx.ForUpdate(); u != nil {
		c.ForUpdate = u.Accept(v).([]Node)
	}
	return c
}

func (v *Builder) VisitForInit(ctx *parser.ForInitContext) interface{} {
	if d := ctx.LocalVariableDeclaration(); d != nil {
		return d.Accept(v)
	}
	return ctx.ExpressionList().Accept(v)
}

func (v *Builder) VisitEnhancedForControl(ctx *parser.EnhancedForControlContext) interface{} {
	n := &EnhancedForControl{Location: v.newLocation(ctx)}
	n.TypeRef = ctx.ApexType().Accept(v).(*TypeRef)
	n.VariableDeclaratorId = ctx.VariableDeclaratorId().Accept(v).(string)
	n.Expression = ctx.Expression().Accept(v).(Node)
	modifiers := ctx.AllVariableModifier()
	n.Modifiers = make([]*Modifier, len(modifiers))
	for i, m := range modifiers {
		n.Modifiers[i] = m.Accept(v).(*Modifier)
	}
	return n
}

func (v *Builder) VisitForUpdate(ctx *parser.ForUpdateContext) interface{} {
	return ctx.ExpressionList().Accept(v)
}

func (v *Builder) VisitParExpression(ctx *parser.ParExpressionContext) interface{} {
	return ctx.Expression().Accept(v)
}

func (v *Builder) VisitExpressionList(ctx *parser.ExpressionListContext) interface{} {
	expressions := ctx.AllExpression()
	nodes := make([]Node, len(expressions))
	for i, e := range expressions {
		nodes[i] = e.Accept(v).(Node)
	}
	return nodes
}

func (v *Builder) VisitStatementExpression(ctx *parser.StatementExpressionContext) interface{} {
	return ctx.Expression().Accept(v)
}

func (v *Builder) VisitConstantExpression(ctx *parser.ConstantExpressionContext) interface{} {
	return ctx.Expression().Accept(v)
}

func (v *Builder) VisitApexDbExpressionShort(ctx *parser.ApexDbExpressionShortContext) interface{} {
	n := &Dml{Location: v.newLocation(ctx)}
	n.Type = ctx.GetDml().GetText()
	if ident := ctx.ApexIdentifier(); ident != nil {
		n.UpsertKey = ident.Accept(v).(string)
	}
	n.Expression = ctx.Expression().Accept(v).(Node)
	return n
}

func (v *Builder) VisitApexDbExpression(ctx *parser.ApexDbExpressionContext) interface{} {
	return ctx.ApexDbExpressionShort().Accept(v)
}

func (v *Builder) VisitTernalyExpression(ctx *parser.TernalyExpressionContext) interface{} {
	n := &TernalyExpression{Location: v.newLocation(ctx)}
	n.Condition = ctx.Expression(0).Accept(v).(Node)
	n.TrueExpression = ctx.Expression(1).Accept(v).(Node)
	n.FalseExpression = ctx.Expression(2).Accept(v).(Node)
	return n
}

func (v *Builder) VisitPreUnaryExpression(ctx *parser.PreUnaryExpressionContext) interface{} {
	n := &UnaryOperator{Location: v.newLocation(ctx)}
	n.Op = ctx.GetOp().GetText()
	n.Expression = ctx.Expression().Accept(v).(Node)
	n.IsPrefix = true
	return n
}

func (v *Builder) VisitArrayAccess(ctx *parser.ArrayAccessContext) interface{} {
	n := &ArrayAccess{Location: v.newLocation(ctx)}
	n.Receiver = ctx.Expression(0).Accept(v).(Node)
	n.Key = ctx.Expression(1).Accept(v).(Node)
	return n
}

func (v *Builder) VisitPostUnaryExpression(ctx *parser.PostUnaryExpressionContext) interface{} {
	n := &UnaryOperator{Location: v.newLocation(ctx)}
	n.Op = ctx.GetOp().GetText()
	n.Expression = ctx.Expression().Accept(v).(Node)
	n.IsPrefix = false
	return n
}

func (v *Builder) VisitPrimaryExpression(ctx *parser.PrimaryExpressionContext) interface{} {
	return ctx.Primary().Accept(v)
}

func (v *Builder) VisitOpExpression(ctx *parser.OpExpressionContext) interface{} {
	n := &BinaryOperator{Location: v.newLocation(ctx)}
	n.Op = ctx.GetOp().GetText()
	n.Left = ctx.Expression(0).Accept(v).(Node)
	n.Right = ctx.Expression(1).Accept(v).(Node)
	return n
}

func (v *Builder) VisitInstanceofExpression(ctx *parser.InstanceofExpressionContext) interface{} {
	n := &InstanceofOperator{Location: v.newLocation(ctx)}
	n.Op = ctx.GetOp().GetText()
	n.Expression = ctx.Expression().Accept(v).(Node)
	n.TypeRef = ctx.ApexType().Accept(v).(*TypeRef)
	return n
}

func (v *Builder) VisitNewObjectExpression(ctx *parser.NewObjectExpressionContext) interface{} {
	return ctx.Creator().Accept(v)
}

func (v *Builder) VisitUnaryExpression(ctx *parser.UnaryExpressionContext) interface{} {
	n := &UnaryOperator{Location: v.newLocation(ctx)}
	n.Op = ctx.GetOp().GetText()
	n.Expression = ctx.Expression().Accept(v).(Node)
	return n
}

func (v *Builder) VisitMethodInvocation(ctx *parser.MethodInvocationContext) interface{} {
	n := &MethodInvocation{Location: v.newLocation(ctx)}
	n.NameOrExpression = ctx.Expression().Accept(v).(Node)
	if list := ctx.ExpressionList(); list != nil {
		n.Parameters = list.Accept(v).([]Node)
	} else {
		n.Parameters = []Node{}
	}
	return n
}

func (v *Builder) VisitCastExpression(ctx *parser.CastExpressionContext) interface{} {
	n := &CastExpression{Location: v.newLocation(ctx)}
	n.CastTypeRef = ctx.ApexType().Accept(v).(*TypeRef)
	n.Expression = ctx.Expression().Accept(v).(Node)
	return n
}

func (v *Builder) VisitShiftExpression(ctx *parser.ShiftExpressionContext) interface{} {
	n := &BinaryOperator{Location: v.newLocation(ctx)}
	ops := []string{}
	for i, o := range ctx.GetOp() {
		ops[i] = o.GetText()
	}
	n.Op = strings.Join(ops, "")
	n.Left = ctx.Expression(0).Accept(v).(Node)
	n.Right = ctx.Expression(1).Accept(v).(Node)
	return n
}

func (v *Builder) VisitFieldAccess(ctx *parser.FieldAccessContext) interface{} {
	expression := ctx.Expression().Accept(v).(Node)
	fieldName := ctx.ApexIdentifier().Accept(v).(string)
	switch n := expression.(type) {
	case *Name:
		value := append(n.Value, fieldName)
		return &Name{
			Value:    value,
			Location: n.Location,
			Parent:   n.Parent,
		}
	default:
		return &FieldAccess{
			Expression: expression,
			FieldName:  fieldName,
			Location:   n.GetLocation(),
			Parent:     n.GetParent(),
		}
	}
	return &FieldAccess{
		Expression: expression,
		FieldName:  fieldName,
		Location:   v.newLocation(ctx),
	}
}

func (v *Builder) VisitPrimary(ctx *parser.PrimaryContext) interface{} {
	if e := ctx.Expression(); e != nil {
		return e.Accept(v)
	} else if t := ctx.THIS(); t != nil {
		return &Name{
			Value:    []string{t.GetText()},
			Location: v.newLocation(ctx),
		}
	} else if s := ctx.SUPER(); s != nil {
		return &Name{
			Value:    []string{t.GetText()},
			Location: v.newLocation(ctx),
		}
	} else if l := ctx.Literal(); l != nil {
		return l.Accept(v)
	} else if i := ctx.ApexIdentifier(); i != nil {
		n := &Name{Location: v.newLocation(ctx)}
		res := i.Accept(v)
		var value string
		switch r := res.(type) {
		case string:
			value = r
		case *TypeRef:
			value = strings.Join(r.Name, ".")
		}
		n.Value = []string{value}
		return n
	} else if l := ctx.SoqlLiteral(); l != nil {
		return l.Accept(v)
	} else if l := ctx.SoslLiteral(); l != nil {
		return l.Accept(v)
	} else if t := ctx.ApexType(); t != nil {
		return t.Accept(v)
	}
	n := &Name{Location: v.newLocation(ctx)}
	value := ctx.PrimitiveType().GetText()
	n.Value = []string{value}
	return n
}

func (v *Builder) VisitCreator(ctx *parser.CreatorContext) interface{} {
	classType := ctx.CreatedName().Accept(v).(*TypeRef)
	if r := ctx.ArrayCreatorRest(); r != nil {
		init := r.Accept(v).(*Init)
		return &New{
			TypeRef: &TypeRef{
				Name:       []string{"List"},
				Parameters: []*TypeRef{classType},
			},
			Parameters: []Node{},
			Init:       init,
			Location:   v.newLocation(ctx),
		}
	}

	if r := ctx.ClassCreatorRest(); r != nil {
		parameters := r.Accept(v)
		return &New{
			TypeRef:    classType,
			Parameters: parameters.([]Node),
			Init:       nil,
			Location:   v.newLocation(ctx),
		}
	}

	if r := ctx.MapCreatorRest(); r != nil {
		init := r.Accept(v)
		return &New{
			TypeRef:    classType,
			Parameters: []Node{},
			Init:       init.(*Init),
			Location:   v.newLocation(ctx),
		}
	}

	init := ctx.SetCreatorRest().Accept(v)
	return &New{
		TypeRef:    classType,
		Parameters: []Node{},
		Init:       init.(*Init),
		Location:   v.newLocation(ctx),
	}
}

func (v *Builder) VisitCreatedName(ctx *parser.CreatedNameContext) interface{} {
	if identifiers := ctx.AllApexIdentifier(); len(identifiers) != 0 {
		n := &TypeRef{Location: v.newLocation(ctx)}
		if len(ctx.AllTypeArgumentsOrDiamond()) != 0 {
			n.Parameters = ctx.TypeArgumentsOrDiamond(0).Accept(v).([]*TypeRef)
		}
		names := make([]string, len(identifiers))
		for i, ident := range identifiers {
			// TODO
			name := ident.Accept(v)
			switch val := name.(type) {
			case *TypeRef:
				names[i] = val.Name[0]
			case string:
				names[i] = name.(string)
			}
		}
		n.Name = names
		// TODO: implement
		return n
	}
	return nil
}

func (v *Builder) VisitInnerCreator(ctx *parser.InnerCreatorContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *Builder) VisitArrayCreatorRest(ctx *parser.ArrayCreatorRestContext) interface{} {
	if expressions := ctx.AllExpression(); len(expressions) != 0 {
		init := &Init{}
		init.Sizes = make([]Node, len(expressions))
		for i, exp := range expressions {
			init.Sizes[i] = exp.Accept(v).(Node)
		}
		for range ctx.AllTypedArray() {
			init.Sizes = append(init.Sizes, &IntegerLiteral{Value: 0})
		}
		return init
	}
	init := ctx.ArrayInitializer().Accept(v).(*Init)
	init.Sizes = make([]Node, len(ctx.AllTypedArray()))
	for i, _ := range init.Sizes {
		init.Sizes[i] = &IntegerLiteral{Value: 0}
	}
	return init
}

func (v *Builder) VisitMapCreatorRest(ctx *parser.MapCreatorRestContext) interface{} {
	keys := ctx.AllMapKey()
	values := ctx.AllMapValue()
	if len(keys) == 0 {
		return &Init{
			Values: map[Node]Node{},
		}
	}
	initValues := map[Node]Node{}
	for i, key := range keys {
		keyNode := key.Accept(v).(Node)
		valueNode := values[i].Accept(v).(Node)
		initValues[keyNode] = valueNode
	}
	return &Init{Values: initValues}
}

func (v *Builder) VisitMapKey(ctx *parser.MapKeyContext) interface{} {
	if r := ctx.ApexIdentifier(); r != nil {
		return r.Accept(v)
	}
	return ctx.Expression().Accept(v)
}

func (v *Builder) VisitMapValue(ctx *parser.MapValueContext) interface{} {
	if r := ctx.Literal(); r != nil {
		return r.Accept(v)
	}
	return ctx.Expression().Accept(v)
}

func (v *Builder) VisitSetCreatorRest(ctx *parser.SetCreatorRestContext) interface{} {
	setValues := ctx.AllSetValue()
	initValues := make([]Node, len(setValues))
	for i, setValue := range setValues {
		initValues[i] = setValue.Accept(v).(Node)
	}
	return &Init{Records: initValues}
}

func (v *Builder) VisitSetValue(ctx *parser.SetValueContext) interface{} {
	if r := ctx.Literal(); r != nil {
		return r.Accept(v)
	}
	return ctx.Expression().Accept(v)
}

func (v *Builder) VisitClassCreatorRest(ctx *parser.ClassCreatorRestContext) interface{} {
	return ctx.Arguments().Accept(v)
}

func (v *Builder) VisitExplicitGenericInvocation(ctx *parser.ExplicitGenericInvocationContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *Builder) VisitNonWildcardTypeArguments(ctx *parser.NonWildcardTypeArgumentsContext) interface{} {
	return ctx.TypeList().Accept(v)
}

func (v *Builder) VisitTypeArgumentsOrDiamond(ctx *parser.TypeArgumentsOrDiamondContext) interface{} {
	return ctx.TypeArguments().Accept(v)
}

func (v *Builder) VisitNonWildcardTypeArgumentsOrDiamond(ctx *parser.NonWildcardTypeArgumentsOrDiamondContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *Builder) VisitSuperSuffix(ctx *parser.SuperSuffixContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *Builder) VisitExplicitGenericInvocationSuffix(ctx *parser.ExplicitGenericInvocationSuffixContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *Builder) VisitArguments(ctx *parser.ArgumentsContext) interface{} {
	if list := ctx.ExpressionList(); list != nil {
		return list.Accept(v)
	}
	return []Node{}
}

func (v *Builder) VisitApexIdentifier(ctx *parser.ApexIdentifierContext) interface{} {
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

func (v *Builder) VisitTypeIdentifier(ctx *parser.TypeIdentifierContext) interface{} {
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
func (v *Builder) VisitSoqlLiteral(ctx *parser.SoqlLiteralContext) interface{} {
	return ctx.Query().Accept(v)
}

func (v *Builder) VisitQuery(ctx *parser.QueryContext) interface{} {
	n := &Soql{Location: v.newLocation(ctx)}
	n.SelectFields = ctx.SelectClause().Accept(v).([]Node)
	n.FromObject = ctx.FromClause().Accept(v).(string)
	if where := ctx.WhereClause(); where != nil {
		n.Where = where.Accept(v).(Node)
	}
	if group := ctx.GroupClause(); group != nil {
		n.Group = group.Accept(v).(*Group)
	}
	if order := ctx.OrderClause(); order != nil {
		n.Order = order.Accept(v).(Node)
	}
	if limit := ctx.LimitClause(); limit != nil {
		n.Limit = limit.Accept(v).(Node)
	}
	if offset := ctx.OffsetClause(); offset != nil {
		n.Offset = offset.Accept(v).(Node)
	}
	return n
}

func (v *Builder) VisitSelectClause(ctx *parser.SelectClauseContext) interface{} {
	return ctx.FieldList().Accept(v)
}

func (v *Builder) VisitFieldList(ctx *parser.FieldListContext) interface{} {
	selectFields := ctx.AllSelectField()
	fields := make([]Node, len(selectFields))
	for i, f := range selectFields {
		fields[i] = f.Accept(v).(Node)
	}
	return fields
}

func (v *Builder) VisitSelectField(ctx *parser.SelectFieldContext) interface{} {
	if t := ctx.SoqlField(); t != nil {
		return t.Accept(v)
	}
	if t := ctx.Subquery(); t != nil {
		return t.Accept(v)
	}
	return nil
}

func (v *Builder) VisitFromClause(ctx *parser.FromClauseContext) interface{} {
	return ctx.ApexIdentifier().GetText()
}

func (v *Builder) VisitFilterScope(ctx *parser.FilterScopeContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *Builder) VisitSoqlFieldReference(ctx *parser.SoqlFieldReferenceContext) interface{} {
	n := &SelectField{Location: v.newLocation(ctx)}
	identifiers := ctx.AllApexIdentifier()
	n.Value = make([]string, len(identifiers))
	for i, ident := range identifiers {
		n.Value[i] = ident.GetText()
	}
	return n
}

func (v *Builder) VisitSoqlFunctionCall(ctx *parser.SoqlFunctionCallContext) interface{} {
	n := &SoqlFunction{}
	n.Name = ctx.ApexIdentifier().GetText()
	for _, f := range ctx.AllSoqlField() {
		f.Accept(v)
	}
	return n
}

func (v *Builder) VisitSubquery(ctx *parser.SubqueryContext) interface{} {
	return ctx.Query().Accept(v)
}

func (v *Builder) VisitWhereClause(ctx *parser.WhereClauseContext) interface{} {
	return ctx.WhereFields().Accept(v)
}

func (v *Builder) VisitWhereFields(ctx *parser.WhereFieldsContext) interface{} {
	if f := ctx.WhereField(); f != nil {
		return f.Accept(v)
	}
	n := &WhereBinaryOperator{Location: v.newLocation(ctx)}
	n.Left = ctx.WhereFields(0).Accept(v).(Node)
	n.Right = ctx.WhereFields(1).Accept(v).(Node)
	n.Op = ctx.GetAnd_or().GetText()
	return n
}

func (v *Builder) VisitWhereField(ctx *parser.WhereFieldContext) interface{} {
	n := &WhereCondition{Location: v.newLocation(ctx)}
	n.Field = ctx.SoqlField().Accept(v).(Node)
	n.Not = ctx.SOQL_NOT() != nil
	n.Op = ctx.GetOp().GetText()
	n.Expression = ctx.SoqlValue().Accept(v).(Node)
	return n
}

func (v *Builder) VisitLimitClause(ctx *parser.LimitClauseContext) interface{} {
	if l := ctx.IntegerLiteral(); l != nil {
		val, err := strconv.Atoi(l.GetText())
		if err != nil {
			panic(err)
		}
		return &IntegerLiteral{Value: val, Location: v.newLocation(ctx)}
	}
	return ctx.BindVariable().Accept(v)
}

func (v *Builder) VisitOrderClause(ctx *parser.OrderClauseContext) interface{} {
	n := &Order{Location: v.newLocation(ctx)}
	fields := make([]Node, len(ctx.AllSoqlField()))
	for i, f := range ctx.AllSoqlField() {
		fields[i] = f.Accept(v).(Node)
	}
	n.Asc = ctx.GetAsc_desc().GetText() == "asc"
	if nulls := ctx.GetNulls(); nulls != nil {
		n.Nulls = nulls.GetText()
	}
	return n
}

func (v *Builder) VisitBindVariable(ctx *parser.BindVariableContext) interface{} {
	return ctx.Expression().Accept(v)
}

func (v *Builder) VisitSoqlValue(ctx *parser.SoqlValueContext) interface{} {
	if l := ctx.Literal(); l != nil {
		return l.Accept(v)
	}
	if b := ctx.BindVariable(); b != nil {
		return b.Accept(v)
	}
	return nil
}

func (v *Builder) VisitWithClause(ctx *parser.WithClauseContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *Builder) VisitSoqlFilteringExpression(ctx *parser.SoqlFilteringExpressionContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *Builder) VisitGroupClause(ctx *parser.GroupClauseContext) interface{} {
	fields := make([]Node, len(ctx.AllSoqlField()))
	for i, f := range ctx.AllSoqlField() {
		fields[i] = f.Accept(v).(Node)
	}
	var having Node
	if ctx.HavingConditionExpression() != nil {
		having = ctx.HavingConditionExpression().Accept(v).(Node)
	}
	return &Group{
		Fields: fields,
		Having: having,
	}
}

func (v *Builder) VisitHavingConditionExpression(ctx *parser.HavingConditionExpressionContext) interface{} {
	return ctx.WhereFields().Accept(v)
}

func (v *Builder) VisitOffsetClause(ctx *parser.OffsetClauseContext) interface{} {
	if l := ctx.IntegerLiteral(); l != nil {
		val, err := strconv.Atoi(l.GetText())
		if err != nil {
			panic(err)
		}
		return &IntegerLiteral{Value: val, Location: v.newLocation(ctx)}
	}
	return ctx.BindVariable().Accept(v)
}

func (v *Builder) VisitViewClause(ctx *parser.ViewClauseContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *Builder) VisitSoslLiteral(ctx *parser.SoslLiteralContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *Builder) VisitSoslQuery(ctx *parser.SoslQueryContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *Builder) VisitSoslReturningObject(ctx *parser.SoslReturningObjectContext) interface{} {
	return v.VisitChildren(ctx)
}

type LocationContext interface {
	GetStart() antlr.Token
}

func (v *Builder) newLocation(ctx LocationContext) *Location {
	return &Location{
		FileName: v.Source,
		Column:   ctx.GetStart().GetColumn(),
		Line:     ctx.GetStart().GetLine(),
	}
}

func setParentNodeToModifiers(modifiers []*Modifier, parent Node) []*Modifier {
	for _, modifier := range modifiers {
		modifier.SetParent(parent)
	}
	return modifiers
}

func setParentNodeToAnnotations(annotations []*Annotation, parent Node) []*Annotation {
	for _, annotation := range annotations {
		annotation.SetParent(parent)
	}
	return annotations
}

func setParentNodeToMethods(methods []*MethodDeclaration, parent Node) []*MethodDeclaration {
	for _, method := range methods {
		method.SetParent(parent)
	}
	return methods
}

func setParentNodeToThrows(throws []*Throw, parent Node) []*Throw {
	for _, throw := range throws {
		throw.SetParent(parent)
	}
	return throws
}
func setParentNodeToNodes(nodes []Node, parent Node) []Node {
	for _, node := range nodes {
		node.SetParent(parent)
	}
	return nodes
}
