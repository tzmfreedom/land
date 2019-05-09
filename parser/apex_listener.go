// Code generated from apex.g4 by ANTLR 4.7.1. DO NOT EDIT.

package parser // apex

import "github.com/antlr/antlr4/runtime/Go/antlr"

// apexListener is a complete listener for a parse tree produced by apexParser.
type apexListener interface {
	antlr.ParseTreeListener

	// EnterCompilationUnit is called when entering the compilationUnit production.
	EnterCompilationUnit(c *CompilationUnitContext)

	// EnterTypeDeclaration is called when entering the typeDeclaration production.
	EnterTypeDeclaration(c *TypeDeclarationContext)

	// EnterTriggerDeclaration is called when entering the triggerDeclaration production.
	EnterTriggerDeclaration(c *TriggerDeclarationContext)

	// EnterTriggerTimings is called when entering the triggerTimings production.
	EnterTriggerTimings(c *TriggerTimingsContext)

	// EnterTriggerTiming is called when entering the triggerTiming production.
	EnterTriggerTiming(c *TriggerTimingContext)

	// EnterModifier is called when entering the modifier production.
	EnterModifier(c *ModifierContext)

	// EnterClassOrInterfaceModifier is called when entering the classOrInterfaceModifier production.
	EnterClassOrInterfaceModifier(c *ClassOrInterfaceModifierContext)

	// EnterVariableModifier is called when entering the variableModifier production.
	EnterVariableModifier(c *VariableModifierContext)

	// EnterClassDeclaration is called when entering the classDeclaration production.
	EnterClassDeclaration(c *ClassDeclarationContext)

	// EnterEnumDeclaration is called when entering the enumDeclaration production.
	EnterEnumDeclaration(c *EnumDeclarationContext)

	// EnterEnumConstants is called when entering the enumConstants production.
	EnterEnumConstants(c *EnumConstantsContext)

	// EnterEnumConstant is called when entering the enumConstant production.
	EnterEnumConstant(c *EnumConstantContext)

	// EnterEnumBodyDeclarations is called when entering the enumBodyDeclarations production.
	EnterEnumBodyDeclarations(c *EnumBodyDeclarationsContext)

	// EnterInterfaceDeclaration is called when entering the interfaceDeclaration production.
	EnterInterfaceDeclaration(c *InterfaceDeclarationContext)

	// EnterTypeList is called when entering the typeList production.
	EnterTypeList(c *TypeListContext)

	// EnterClassBody is called when entering the classBody production.
	EnterClassBody(c *ClassBodyContext)

	// EnterInterfaceBody is called when entering the interfaceBody production.
	EnterInterfaceBody(c *InterfaceBodyContext)

	// EnterClassBodyDeclaration is called when entering the classBodyDeclaration production.
	EnterClassBodyDeclaration(c *ClassBodyDeclarationContext)

	// EnterMemberDeclaration is called when entering the memberDeclaration production.
	EnterMemberDeclaration(c *MemberDeclarationContext)

	// EnterMethodDeclaration is called when entering the methodDeclaration production.
	EnterMethodDeclaration(c *MethodDeclarationContext)

	// EnterConstructorDeclaration is called when entering the constructorDeclaration production.
	EnterConstructorDeclaration(c *ConstructorDeclarationContext)

	// EnterFieldDeclaration is called when entering the fieldDeclaration production.
	EnterFieldDeclaration(c *FieldDeclarationContext)

	// EnterPropertyDeclaration is called when entering the propertyDeclaration production.
	EnterPropertyDeclaration(c *PropertyDeclarationContext)

	// EnterPropertyBodyDeclaration is called when entering the propertyBodyDeclaration production.
	EnterPropertyBodyDeclaration(c *PropertyBodyDeclarationContext)

	// EnterInterfaceBodyDeclaration is called when entering the interfaceBodyDeclaration production.
	EnterInterfaceBodyDeclaration(c *InterfaceBodyDeclarationContext)

	// EnterInterfaceMemberDeclaration is called when entering the interfaceMemberDeclaration production.
	EnterInterfaceMemberDeclaration(c *InterfaceMemberDeclarationContext)

	// EnterConstDeclaration is called when entering the constDeclaration production.
	EnterConstDeclaration(c *ConstDeclarationContext)

	// EnterConstantDeclarator is called when entering the constantDeclarator production.
	EnterConstantDeclarator(c *ConstantDeclaratorContext)

	// EnterInterfaceMethodDeclaration is called when entering the interfaceMethodDeclaration production.
	EnterInterfaceMethodDeclaration(c *InterfaceMethodDeclarationContext)

	// EnterVariableDeclarators is called when entering the variableDeclarators production.
	EnterVariableDeclarators(c *VariableDeclaratorsContext)

	// EnterVariableDeclarator is called when entering the variableDeclarator production.
	EnterVariableDeclarator(c *VariableDeclaratorContext)

	// EnterVariableDeclaratorId is called when entering the variableDeclaratorId production.
	EnterVariableDeclaratorId(c *VariableDeclaratorIdContext)

	// EnterVariableInitializer is called when entering the variableInitializer production.
	EnterVariableInitializer(c *VariableInitializerContext)

	// EnterArrayInitializer is called when entering the arrayInitializer production.
	EnterArrayInitializer(c *ArrayInitializerContext)

	// EnterEnumConstantName is called when entering the enumConstantName production.
	EnterEnumConstantName(c *EnumConstantNameContext)

	// EnterApexType is called when entering the apexType production.
	EnterApexType(c *ApexTypeContext)

	// EnterTypedArray is called when entering the typedArray production.
	EnterTypedArray(c *TypedArrayContext)

	// EnterClassOrInterfaceType is called when entering the classOrInterfaceType production.
	EnterClassOrInterfaceType(c *ClassOrInterfaceTypeContext)

	// EnterPrimitiveType is called when entering the primitiveType production.
	EnterPrimitiveType(c *PrimitiveTypeContext)

	// EnterTypeArguments is called when entering the typeArguments production.
	EnterTypeArguments(c *TypeArgumentsContext)

	// EnterTypeArgument is called when entering the typeArgument production.
	EnterTypeArgument(c *TypeArgumentContext)

	// EnterQualifiedNameList is called when entering the qualifiedNameList production.
	EnterQualifiedNameList(c *QualifiedNameListContext)

	// EnterFormalParameters is called when entering the formalParameters production.
	EnterFormalParameters(c *FormalParametersContext)

	// EnterFormalParameterList is called when entering the formalParameterList production.
	EnterFormalParameterList(c *FormalParameterListContext)

	// EnterFormalParameter is called when entering the formalParameter production.
	EnterFormalParameter(c *FormalParameterContext)

	// EnterLastFormalParameter is called when entering the lastFormalParameter production.
	EnterLastFormalParameter(c *LastFormalParameterContext)

	// EnterMethodBody is called when entering the methodBody production.
	EnterMethodBody(c *MethodBodyContext)

	// EnterConstructorBody is called when entering the constructorBody production.
	EnterConstructorBody(c *ConstructorBodyContext)

	// EnterQualifiedName is called when entering the qualifiedName production.
	EnterQualifiedName(c *QualifiedNameContext)

	// EnterLiteral is called when entering the literal production.
	EnterLiteral(c *LiteralContext)

	// EnterAnnotation is called when entering the annotation production.
	EnterAnnotation(c *AnnotationContext)

	// EnterAnnotationName is called when entering the annotationName production.
	EnterAnnotationName(c *AnnotationNameContext)

	// EnterElementValuePairs is called when entering the elementValuePairs production.
	EnterElementValuePairs(c *ElementValuePairsContext)

	// EnterElementValuePair is called when entering the elementValuePair production.
	EnterElementValuePair(c *ElementValuePairContext)

	// EnterElementValue is called when entering the elementValue production.
	EnterElementValue(c *ElementValueContext)

	// EnterElementValueArrayInitializer is called when entering the elementValueArrayInitializer production.
	EnterElementValueArrayInitializer(c *ElementValueArrayInitializerContext)

	// EnterBlock is called when entering the block production.
	EnterBlock(c *BlockContext)

	// EnterBlockStatement is called when entering the blockStatement production.
	EnterBlockStatement(c *BlockStatementContext)

	// EnterLocalVariableDeclarationStatement is called when entering the localVariableDeclarationStatement production.
	EnterLocalVariableDeclarationStatement(c *LocalVariableDeclarationStatementContext)

	// EnterLocalVariableDeclaration is called when entering the localVariableDeclaration production.
	EnterLocalVariableDeclaration(c *LocalVariableDeclarationContext)

	// EnterStatement is called when entering the statement production.
	EnterStatement(c *StatementContext)

	// EnterPropertyBlock is called when entering the propertyBlock production.
	EnterPropertyBlock(c *PropertyBlockContext)

	// EnterGetter is called when entering the getter production.
	EnterGetter(c *GetterContext)

	// EnterSetter is called when entering the setter production.
	EnterSetter(c *SetterContext)

	// EnterCatchClause is called when entering the catchClause production.
	EnterCatchClause(c *CatchClauseContext)

	// EnterCatchType is called when entering the catchType production.
	EnterCatchType(c *CatchTypeContext)

	// EnterFinallyBlock is called when entering the finallyBlock production.
	EnterFinallyBlock(c *FinallyBlockContext)

	// EnterWhenStatements is called when entering the whenStatements production.
	EnterWhenStatements(c *WhenStatementsContext)

	// EnterWhenStatement is called when entering the whenStatement production.
	EnterWhenStatement(c *WhenStatementContext)

	// EnterWhenExpression is called when entering the whenExpression production.
	EnterWhenExpression(c *WhenExpressionContext)

	// EnterForControl is called when entering the forControl production.
	EnterForControl(c *ForControlContext)

	// EnterForInit is called when entering the forInit production.
	EnterForInit(c *ForInitContext)

	// EnterEnhancedForControl is called when entering the enhancedForControl production.
	EnterEnhancedForControl(c *EnhancedForControlContext)

	// EnterForUpdate is called when entering the forUpdate production.
	EnterForUpdate(c *ForUpdateContext)

	// EnterParExpression is called when entering the parExpression production.
	EnterParExpression(c *ParExpressionContext)

	// EnterExpressionList is called when entering the expressionList production.
	EnterExpressionList(c *ExpressionListContext)

	// EnterStatementExpression is called when entering the statementExpression production.
	EnterStatementExpression(c *StatementExpressionContext)

	// EnterConstantExpression is called when entering the constantExpression production.
	EnterConstantExpression(c *ConstantExpressionContext)

	// EnterApexDbExpressionShort is called when entering the apexDbExpressionShort production.
	EnterApexDbExpressionShort(c *ApexDbExpressionShortContext)

	// EnterApexDbExpression is called when entering the apexDbExpression production.
	EnterApexDbExpression(c *ApexDbExpressionContext)

	// EnterPrimaryExpression is called when entering the PrimaryExpression production.
	EnterPrimaryExpression(c *PrimaryExpressionContext)

	// EnterUnaryExpression is called when entering the UnaryExpression production.
	EnterUnaryExpression(c *UnaryExpressionContext)

	// EnterMethodInvocation is called when entering the MethodInvocation production.
	EnterMethodInvocation(c *MethodInvocationContext)

	// EnterShiftExpression is called when entering the ShiftExpression production.
	EnterShiftExpression(c *ShiftExpressionContext)

	// EnterNewObjectExpression is called when entering the NewObjectExpression production.
	EnterNewObjectExpression(c *NewObjectExpressionContext)

	// EnterTernalyExpression is called when entering the TernalyExpression production.
	EnterTernalyExpression(c *TernalyExpressionContext)

	// EnterPreUnaryExpression is called when entering the PreUnaryExpression production.
	EnterPreUnaryExpression(c *PreUnaryExpressionContext)

	// EnterArrayAccess is called when entering the ArrayAccess production.
	EnterArrayAccess(c *ArrayAccessContext)

	// EnterPostUnaryExpression is called when entering the PostUnaryExpression production.
	EnterPostUnaryExpression(c *PostUnaryExpressionContext)

	// EnterOpExpression is called when entering the OpExpression production.
	EnterOpExpression(c *OpExpressionContext)

	// EnterInstanceofExpression is called when entering the InstanceofExpression production.
	EnterInstanceofExpression(c *InstanceofExpressionContext)

	// EnterCastExpression is called when entering the CastExpression production.
	EnterCastExpression(c *CastExpressionContext)

	// EnterFieldAccess is called when entering the FieldAccess production.
	EnterFieldAccess(c *FieldAccessContext)

	// EnterPrimary is called when entering the primary production.
	EnterPrimary(c *PrimaryContext)

	// EnterCreator is called when entering the creator production.
	EnterCreator(c *CreatorContext)

	// EnterCreatedName is called when entering the createdName production.
	EnterCreatedName(c *CreatedNameContext)

	// EnterInnerCreator is called when entering the innerCreator production.
	EnterInnerCreator(c *InnerCreatorContext)

	// EnterArrayCreatorRest is called when entering the arrayCreatorRest production.
	EnterArrayCreatorRest(c *ArrayCreatorRestContext)

	// EnterMapCreatorRest is called when entering the mapCreatorRest production.
	EnterMapCreatorRest(c *MapCreatorRestContext)

	// EnterMapKey is called when entering the mapKey production.
	EnterMapKey(c *MapKeyContext)

	// EnterMapValue is called when entering the mapValue production.
	EnterMapValue(c *MapValueContext)

	// EnterSetCreatorRest is called when entering the setCreatorRest production.
	EnterSetCreatorRest(c *SetCreatorRestContext)

	// EnterSetValue is called when entering the setValue production.
	EnterSetValue(c *SetValueContext)

	// EnterClassCreatorRest is called when entering the classCreatorRest production.
	EnterClassCreatorRest(c *ClassCreatorRestContext)

	// EnterExplicitGenericInvocation is called when entering the explicitGenericInvocation production.
	EnterExplicitGenericInvocation(c *ExplicitGenericInvocationContext)

	// EnterNonWildcardTypeArguments is called when entering the nonWildcardTypeArguments production.
	EnterNonWildcardTypeArguments(c *NonWildcardTypeArgumentsContext)

	// EnterTypeArgumentsOrDiamond is called when entering the typeArgumentsOrDiamond production.
	EnterTypeArgumentsOrDiamond(c *TypeArgumentsOrDiamondContext)

	// EnterNonWildcardTypeArgumentsOrDiamond is called when entering the nonWildcardTypeArgumentsOrDiamond production.
	EnterNonWildcardTypeArgumentsOrDiamond(c *NonWildcardTypeArgumentsOrDiamondContext)

	// EnterSuperSuffix is called when entering the superSuffix production.
	EnterSuperSuffix(c *SuperSuffixContext)

	// EnterExplicitGenericInvocationSuffix is called when entering the explicitGenericInvocationSuffix production.
	EnterExplicitGenericInvocationSuffix(c *ExplicitGenericInvocationSuffixContext)

	// EnterArguments is called when entering the arguments production.
	EnterArguments(c *ArgumentsContext)

	// EnterSoqlLiteral is called when entering the soqlLiteral production.
	EnterSoqlLiteral(c *SoqlLiteralContext)

	// EnterQuery is called when entering the query production.
	EnterQuery(c *QueryContext)

	// EnterSelectClause is called when entering the selectClause production.
	EnterSelectClause(c *SelectClauseContext)

	// EnterFieldList is called when entering the fieldList production.
	EnterFieldList(c *FieldListContext)

	// EnterSelectField is called when entering the selectField production.
	EnterSelectField(c *SelectFieldContext)

	// EnterFromClause is called when entering the fromClause production.
	EnterFromClause(c *FromClauseContext)

	// EnterFilterScope is called when entering the filterScope production.
	EnterFilterScope(c *FilterScopeContext)

	// EnterSoqlFieldReference is called when entering the SoqlFieldReference production.
	EnterSoqlFieldReference(c *SoqlFieldReferenceContext)

	// EnterSoqlFunctionCall is called when entering the SoqlFunctionCall production.
	EnterSoqlFunctionCall(c *SoqlFunctionCallContext)

	// EnterSubquery is called when entering the subquery production.
	EnterSubquery(c *SubqueryContext)

	// EnterWhereClause is called when entering the whereClause production.
	EnterWhereClause(c *WhereClauseContext)

	// EnterWhereFields is called when entering the whereFields production.
	EnterWhereFields(c *WhereFieldsContext)

	// EnterWhereField is called when entering the whereField production.
	EnterWhereField(c *WhereFieldContext)

	// EnterLimitClause is called when entering the limitClause production.
	EnterLimitClause(c *LimitClauseContext)

	// EnterOrderClause is called when entering the orderClause production.
	EnterOrderClause(c *OrderClauseContext)

	// EnterBindVariable is called when entering the bindVariable production.
	EnterBindVariable(c *BindVariableContext)

	// EnterSoqlValue is called when entering the soqlValue production.
	EnterSoqlValue(c *SoqlValueContext)

	// EnterWithClause is called when entering the withClause production.
	EnterWithClause(c *WithClauseContext)

	// EnterSoqlFilteringExpression is called when entering the soqlFilteringExpression production.
	EnterSoqlFilteringExpression(c *SoqlFilteringExpressionContext)

	// EnterGroupClause is called when entering the groupClause production.
	EnterGroupClause(c *GroupClauseContext)

	// EnterHavingConditionExpression is called when entering the havingConditionExpression production.
	EnterHavingConditionExpression(c *HavingConditionExpressionContext)

	// EnterOffsetClause is called when entering the offsetClause production.
	EnterOffsetClause(c *OffsetClauseContext)

	// EnterViewClause is called when entering the viewClause production.
	EnterViewClause(c *ViewClauseContext)

	// EnterSoslLiteral is called when entering the soslLiteral production.
	EnterSoslLiteral(c *SoslLiteralContext)

	// EnterSoslQuery is called when entering the soslQuery production.
	EnterSoslQuery(c *SoslQueryContext)

	// EnterSoslReturningObject is called when entering the soslReturningObject production.
	EnterSoslReturningObject(c *SoslReturningObjectContext)

	// EnterApexIdentifier is called when entering the apexIdentifier production.
	EnterApexIdentifier(c *ApexIdentifierContext)

	// EnterTypeIdentifier is called when entering the typeIdentifier production.
	EnterTypeIdentifier(c *TypeIdentifierContext)

	// ExitCompilationUnit is called when exiting the compilationUnit production.
	ExitCompilationUnit(c *CompilationUnitContext)

	// ExitTypeDeclaration is called when exiting the typeDeclaration production.
	ExitTypeDeclaration(c *TypeDeclarationContext)

	// ExitTriggerDeclaration is called when exiting the triggerDeclaration production.
	ExitTriggerDeclaration(c *TriggerDeclarationContext)

	// ExitTriggerTimings is called when exiting the triggerTimings production.
	ExitTriggerTimings(c *TriggerTimingsContext)

	// ExitTriggerTiming is called when exiting the triggerTiming production.
	ExitTriggerTiming(c *TriggerTimingContext)

	// ExitModifier is called when exiting the modifier production.
	ExitModifier(c *ModifierContext)

	// ExitClassOrInterfaceModifier is called when exiting the classOrInterfaceModifier production.
	ExitClassOrInterfaceModifier(c *ClassOrInterfaceModifierContext)

	// ExitVariableModifier is called when exiting the variableModifier production.
	ExitVariableModifier(c *VariableModifierContext)

	// ExitClassDeclaration is called when exiting the classDeclaration production.
	ExitClassDeclaration(c *ClassDeclarationContext)

	// ExitEnumDeclaration is called when exiting the enumDeclaration production.
	ExitEnumDeclaration(c *EnumDeclarationContext)

	// ExitEnumConstants is called when exiting the enumConstants production.
	ExitEnumConstants(c *EnumConstantsContext)

	// ExitEnumConstant is called when exiting the enumConstant production.
	ExitEnumConstant(c *EnumConstantContext)

	// ExitEnumBodyDeclarations is called when exiting the enumBodyDeclarations production.
	ExitEnumBodyDeclarations(c *EnumBodyDeclarationsContext)

	// ExitInterfaceDeclaration is called when exiting the interfaceDeclaration production.
	ExitInterfaceDeclaration(c *InterfaceDeclarationContext)

	// ExitTypeList is called when exiting the typeList production.
	ExitTypeList(c *TypeListContext)

	// ExitClassBody is called when exiting the classBody production.
	ExitClassBody(c *ClassBodyContext)

	// ExitInterfaceBody is called when exiting the interfaceBody production.
	ExitInterfaceBody(c *InterfaceBodyContext)

	// ExitClassBodyDeclaration is called when exiting the classBodyDeclaration production.
	ExitClassBodyDeclaration(c *ClassBodyDeclarationContext)

	// ExitMemberDeclaration is called when exiting the memberDeclaration production.
	ExitMemberDeclaration(c *MemberDeclarationContext)

	// ExitMethodDeclaration is called when exiting the methodDeclaration production.
	ExitMethodDeclaration(c *MethodDeclarationContext)

	// ExitConstructorDeclaration is called when exiting the constructorDeclaration production.
	ExitConstructorDeclaration(c *ConstructorDeclarationContext)

	// ExitFieldDeclaration is called when exiting the fieldDeclaration production.
	ExitFieldDeclaration(c *FieldDeclarationContext)

	// ExitPropertyDeclaration is called when exiting the propertyDeclaration production.
	ExitPropertyDeclaration(c *PropertyDeclarationContext)

	// ExitPropertyBodyDeclaration is called when exiting the propertyBodyDeclaration production.
	ExitPropertyBodyDeclaration(c *PropertyBodyDeclarationContext)

	// ExitInterfaceBodyDeclaration is called when exiting the interfaceBodyDeclaration production.
	ExitInterfaceBodyDeclaration(c *InterfaceBodyDeclarationContext)

	// ExitInterfaceMemberDeclaration is called when exiting the interfaceMemberDeclaration production.
	ExitInterfaceMemberDeclaration(c *InterfaceMemberDeclarationContext)

	// ExitConstDeclaration is called when exiting the constDeclaration production.
	ExitConstDeclaration(c *ConstDeclarationContext)

	// ExitConstantDeclarator is called when exiting the constantDeclarator production.
	ExitConstantDeclarator(c *ConstantDeclaratorContext)

	// ExitInterfaceMethodDeclaration is called when exiting the interfaceMethodDeclaration production.
	ExitInterfaceMethodDeclaration(c *InterfaceMethodDeclarationContext)

	// ExitVariableDeclarators is called when exiting the variableDeclarators production.
	ExitVariableDeclarators(c *VariableDeclaratorsContext)

	// ExitVariableDeclarator is called when exiting the variableDeclarator production.
	ExitVariableDeclarator(c *VariableDeclaratorContext)

	// ExitVariableDeclaratorId is called when exiting the variableDeclaratorId production.
	ExitVariableDeclaratorId(c *VariableDeclaratorIdContext)

	// ExitVariableInitializer is called when exiting the variableInitializer production.
	ExitVariableInitializer(c *VariableInitializerContext)

	// ExitArrayInitializer is called when exiting the arrayInitializer production.
	ExitArrayInitializer(c *ArrayInitializerContext)

	// ExitEnumConstantName is called when exiting the enumConstantName production.
	ExitEnumConstantName(c *EnumConstantNameContext)

	// ExitApexType is called when exiting the apexType production.
	ExitApexType(c *ApexTypeContext)

	// ExitTypedArray is called when exiting the typedArray production.
	ExitTypedArray(c *TypedArrayContext)

	// ExitClassOrInterfaceType is called when exiting the classOrInterfaceType production.
	ExitClassOrInterfaceType(c *ClassOrInterfaceTypeContext)

	// ExitPrimitiveType is called when exiting the primitiveType production.
	ExitPrimitiveType(c *PrimitiveTypeContext)

	// ExitTypeArguments is called when exiting the typeArguments production.
	ExitTypeArguments(c *TypeArgumentsContext)

	// ExitTypeArgument is called when exiting the typeArgument production.
	ExitTypeArgument(c *TypeArgumentContext)

	// ExitQualifiedNameList is called when exiting the qualifiedNameList production.
	ExitQualifiedNameList(c *QualifiedNameListContext)

	// ExitFormalParameters is called when exiting the formalParameters production.
	ExitFormalParameters(c *FormalParametersContext)

	// ExitFormalParameterList is called when exiting the formalParameterList production.
	ExitFormalParameterList(c *FormalParameterListContext)

	// ExitFormalParameter is called when exiting the formalParameter production.
	ExitFormalParameter(c *FormalParameterContext)

	// ExitLastFormalParameter is called when exiting the lastFormalParameter production.
	ExitLastFormalParameter(c *LastFormalParameterContext)

	// ExitMethodBody is called when exiting the methodBody production.
	ExitMethodBody(c *MethodBodyContext)

	// ExitConstructorBody is called when exiting the constructorBody production.
	ExitConstructorBody(c *ConstructorBodyContext)

	// ExitQualifiedName is called when exiting the qualifiedName production.
	ExitQualifiedName(c *QualifiedNameContext)

	// ExitLiteral is called when exiting the literal production.
	ExitLiteral(c *LiteralContext)

	// ExitAnnotation is called when exiting the annotation production.
	ExitAnnotation(c *AnnotationContext)

	// ExitAnnotationName is called when exiting the annotationName production.
	ExitAnnotationName(c *AnnotationNameContext)

	// ExitElementValuePairs is called when exiting the elementValuePairs production.
	ExitElementValuePairs(c *ElementValuePairsContext)

	// ExitElementValuePair is called when exiting the elementValuePair production.
	ExitElementValuePair(c *ElementValuePairContext)

	// ExitElementValue is called when exiting the elementValue production.
	ExitElementValue(c *ElementValueContext)

	// ExitElementValueArrayInitializer is called when exiting the elementValueArrayInitializer production.
	ExitElementValueArrayInitializer(c *ElementValueArrayInitializerContext)

	// ExitBlock is called when exiting the block production.
	ExitBlock(c *BlockContext)

	// ExitBlockStatement is called when exiting the blockStatement production.
	ExitBlockStatement(c *BlockStatementContext)

	// ExitLocalVariableDeclarationStatement is called when exiting the localVariableDeclarationStatement production.
	ExitLocalVariableDeclarationStatement(c *LocalVariableDeclarationStatementContext)

	// ExitLocalVariableDeclaration is called when exiting the localVariableDeclaration production.
	ExitLocalVariableDeclaration(c *LocalVariableDeclarationContext)

	// ExitStatement is called when exiting the statement production.
	ExitStatement(c *StatementContext)

	// ExitPropertyBlock is called when exiting the propertyBlock production.
	ExitPropertyBlock(c *PropertyBlockContext)

	// ExitGetter is called when exiting the getter production.
	ExitGetter(c *GetterContext)

	// ExitSetter is called when exiting the setter production.
	ExitSetter(c *SetterContext)

	// ExitCatchClause is called when exiting the catchClause production.
	ExitCatchClause(c *CatchClauseContext)

	// ExitCatchType is called when exiting the catchType production.
	ExitCatchType(c *CatchTypeContext)

	// ExitFinallyBlock is called when exiting the finallyBlock production.
	ExitFinallyBlock(c *FinallyBlockContext)

	// ExitWhenStatements is called when exiting the whenStatements production.
	ExitWhenStatements(c *WhenStatementsContext)

	// ExitWhenStatement is called when exiting the whenStatement production.
	ExitWhenStatement(c *WhenStatementContext)

	// ExitWhenExpression is called when exiting the whenExpression production.
	ExitWhenExpression(c *WhenExpressionContext)

	// ExitForControl is called when exiting the forControl production.
	ExitForControl(c *ForControlContext)

	// ExitForInit is called when exiting the forInit production.
	ExitForInit(c *ForInitContext)

	// ExitEnhancedForControl is called when exiting the enhancedForControl production.
	ExitEnhancedForControl(c *EnhancedForControlContext)

	// ExitForUpdate is called when exiting the forUpdate production.
	ExitForUpdate(c *ForUpdateContext)

	// ExitParExpression is called when exiting the parExpression production.
	ExitParExpression(c *ParExpressionContext)

	// ExitExpressionList is called when exiting the expressionList production.
	ExitExpressionList(c *ExpressionListContext)

	// ExitStatementExpression is called when exiting the statementExpression production.
	ExitStatementExpression(c *StatementExpressionContext)

	// ExitConstantExpression is called when exiting the constantExpression production.
	ExitConstantExpression(c *ConstantExpressionContext)

	// ExitApexDbExpressionShort is called when exiting the apexDbExpressionShort production.
	ExitApexDbExpressionShort(c *ApexDbExpressionShortContext)

	// ExitApexDbExpression is called when exiting the apexDbExpression production.
	ExitApexDbExpression(c *ApexDbExpressionContext)

	// ExitPrimaryExpression is called when exiting the PrimaryExpression production.
	ExitPrimaryExpression(c *PrimaryExpressionContext)

	// ExitUnaryExpression is called when exiting the UnaryExpression production.
	ExitUnaryExpression(c *UnaryExpressionContext)

	// ExitMethodInvocation is called when exiting the MethodInvocation production.
	ExitMethodInvocation(c *MethodInvocationContext)

	// ExitShiftExpression is called when exiting the ShiftExpression production.
	ExitShiftExpression(c *ShiftExpressionContext)

	// ExitNewObjectExpression is called when exiting the NewObjectExpression production.
	ExitNewObjectExpression(c *NewObjectExpressionContext)

	// ExitTernalyExpression is called when exiting the TernalyExpression production.
	ExitTernalyExpression(c *TernalyExpressionContext)

	// ExitPreUnaryExpression is called when exiting the PreUnaryExpression production.
	ExitPreUnaryExpression(c *PreUnaryExpressionContext)

	// ExitArrayAccess is called when exiting the ArrayAccess production.
	ExitArrayAccess(c *ArrayAccessContext)

	// ExitPostUnaryExpression is called when exiting the PostUnaryExpression production.
	ExitPostUnaryExpression(c *PostUnaryExpressionContext)

	// ExitOpExpression is called when exiting the OpExpression production.
	ExitOpExpression(c *OpExpressionContext)

	// ExitInstanceofExpression is called when exiting the InstanceofExpression production.
	ExitInstanceofExpression(c *InstanceofExpressionContext)

	// ExitCastExpression is called when exiting the CastExpression production.
	ExitCastExpression(c *CastExpressionContext)

	// ExitFieldAccess is called when exiting the FieldAccess production.
	ExitFieldAccess(c *FieldAccessContext)

	// ExitPrimary is called when exiting the primary production.
	ExitPrimary(c *PrimaryContext)

	// ExitCreator is called when exiting the creator production.
	ExitCreator(c *CreatorContext)

	// ExitCreatedName is called when exiting the createdName production.
	ExitCreatedName(c *CreatedNameContext)

	// ExitInnerCreator is called when exiting the innerCreator production.
	ExitInnerCreator(c *InnerCreatorContext)

	// ExitArrayCreatorRest is called when exiting the arrayCreatorRest production.
	ExitArrayCreatorRest(c *ArrayCreatorRestContext)

	// ExitMapCreatorRest is called when exiting the mapCreatorRest production.
	ExitMapCreatorRest(c *MapCreatorRestContext)

	// ExitMapKey is called when exiting the mapKey production.
	ExitMapKey(c *MapKeyContext)

	// ExitMapValue is called when exiting the mapValue production.
	ExitMapValue(c *MapValueContext)

	// ExitSetCreatorRest is called when exiting the setCreatorRest production.
	ExitSetCreatorRest(c *SetCreatorRestContext)

	// ExitSetValue is called when exiting the setValue production.
	ExitSetValue(c *SetValueContext)

	// ExitClassCreatorRest is called when exiting the classCreatorRest production.
	ExitClassCreatorRest(c *ClassCreatorRestContext)

	// ExitExplicitGenericInvocation is called when exiting the explicitGenericInvocation production.
	ExitExplicitGenericInvocation(c *ExplicitGenericInvocationContext)

	// ExitNonWildcardTypeArguments is called when exiting the nonWildcardTypeArguments production.
	ExitNonWildcardTypeArguments(c *NonWildcardTypeArgumentsContext)

	// ExitTypeArgumentsOrDiamond is called when exiting the typeArgumentsOrDiamond production.
	ExitTypeArgumentsOrDiamond(c *TypeArgumentsOrDiamondContext)

	// ExitNonWildcardTypeArgumentsOrDiamond is called when exiting the nonWildcardTypeArgumentsOrDiamond production.
	ExitNonWildcardTypeArgumentsOrDiamond(c *NonWildcardTypeArgumentsOrDiamondContext)

	// ExitSuperSuffix is called when exiting the superSuffix production.
	ExitSuperSuffix(c *SuperSuffixContext)

	// ExitExplicitGenericInvocationSuffix is called when exiting the explicitGenericInvocationSuffix production.
	ExitExplicitGenericInvocationSuffix(c *ExplicitGenericInvocationSuffixContext)

	// ExitArguments is called when exiting the arguments production.
	ExitArguments(c *ArgumentsContext)

	// ExitSoqlLiteral is called when exiting the soqlLiteral production.
	ExitSoqlLiteral(c *SoqlLiteralContext)

	// ExitQuery is called when exiting the query production.
	ExitQuery(c *QueryContext)

	// ExitSelectClause is called when exiting the selectClause production.
	ExitSelectClause(c *SelectClauseContext)

	// ExitFieldList is called when exiting the fieldList production.
	ExitFieldList(c *FieldListContext)

	// ExitSelectField is called when exiting the selectField production.
	ExitSelectField(c *SelectFieldContext)

	// ExitFromClause is called when exiting the fromClause production.
	ExitFromClause(c *FromClauseContext)

	// ExitFilterScope is called when exiting the filterScope production.
	ExitFilterScope(c *FilterScopeContext)

	// ExitSoqlFieldReference is called when exiting the SoqlFieldReference production.
	ExitSoqlFieldReference(c *SoqlFieldReferenceContext)

	// ExitSoqlFunctionCall is called when exiting the SoqlFunctionCall production.
	ExitSoqlFunctionCall(c *SoqlFunctionCallContext)

	// ExitSubquery is called when exiting the subquery production.
	ExitSubquery(c *SubqueryContext)

	// ExitWhereClause is called when exiting the whereClause production.
	ExitWhereClause(c *WhereClauseContext)

	// ExitWhereFields is called when exiting the whereFields production.
	ExitWhereFields(c *WhereFieldsContext)

	// ExitWhereField is called when exiting the whereField production.
	ExitWhereField(c *WhereFieldContext)

	// ExitLimitClause is called when exiting the limitClause production.
	ExitLimitClause(c *LimitClauseContext)

	// ExitOrderClause is called when exiting the orderClause production.
	ExitOrderClause(c *OrderClauseContext)

	// ExitBindVariable is called when exiting the bindVariable production.
	ExitBindVariable(c *BindVariableContext)

	// ExitSoqlValue is called when exiting the soqlValue production.
	ExitSoqlValue(c *SoqlValueContext)

	// ExitWithClause is called when exiting the withClause production.
	ExitWithClause(c *WithClauseContext)

	// ExitSoqlFilteringExpression is called when exiting the soqlFilteringExpression production.
	ExitSoqlFilteringExpression(c *SoqlFilteringExpressionContext)

	// ExitGroupClause is called when exiting the groupClause production.
	ExitGroupClause(c *GroupClauseContext)

	// ExitHavingConditionExpression is called when exiting the havingConditionExpression production.
	ExitHavingConditionExpression(c *HavingConditionExpressionContext)

	// ExitOffsetClause is called when exiting the offsetClause production.
	ExitOffsetClause(c *OffsetClauseContext)

	// ExitViewClause is called when exiting the viewClause production.
	ExitViewClause(c *ViewClauseContext)

	// ExitSoslLiteral is called when exiting the soslLiteral production.
	ExitSoslLiteral(c *SoslLiteralContext)

	// ExitSoslQuery is called when exiting the soslQuery production.
	ExitSoslQuery(c *SoslQueryContext)

	// ExitSoslReturningObject is called when exiting the soslReturningObject production.
	ExitSoslReturningObject(c *SoslReturningObjectContext)

	// ExitApexIdentifier is called when exiting the apexIdentifier production.
	ExitApexIdentifier(c *ApexIdentifierContext)

	// ExitTypeIdentifier is called when exiting the typeIdentifier production.
	ExitTypeIdentifier(c *TypeIdentifierContext)
}
