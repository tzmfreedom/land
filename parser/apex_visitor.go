// Code generated from apex.g4 by ANTLR 4.7.1. DO NOT EDIT.

package parser // apex

import "github.com/antlr/antlr4/runtime/Go/antlr"

// A complete Visitor for a parse tree produced by apexParser.
type apexVisitor interface {
	antlr.ParseTreeVisitor

	// Visit a parse tree produced by apexParser#compilationUnit.
	VisitCompilationUnit(ctx *CompilationUnitContext) interface{}

	// Visit a parse tree produced by apexParser#typeDeclaration.
	VisitTypeDeclaration(ctx *TypeDeclarationContext) interface{}

	// Visit a parse tree produced by apexParser#triggerDeclaration.
	VisitTriggerDeclaration(ctx *TriggerDeclarationContext) interface{}

	// Visit a parse tree produced by apexParser#triggerTimings.
	VisitTriggerTimings(ctx *TriggerTimingsContext) interface{}

	// Visit a parse tree produced by apexParser#triggerTiming.
	VisitTriggerTiming(ctx *TriggerTimingContext) interface{}

	// Visit a parse tree produced by apexParser#modifier.
	VisitModifier(ctx *ModifierContext) interface{}

	// Visit a parse tree produced by apexParser#classOrInterfaceModifier.
	VisitClassOrInterfaceModifier(ctx *ClassOrInterfaceModifierContext) interface{}

	// Visit a parse tree produced by apexParser#variableModifier.
	VisitVariableModifier(ctx *VariableModifierContext) interface{}

	// Visit a parse tree produced by apexParser#classDeclaration.
	VisitClassDeclaration(ctx *ClassDeclarationContext) interface{}

	// Visit a parse tree produced by apexParser#enumDeclaration.
	VisitEnumDeclaration(ctx *EnumDeclarationContext) interface{}

	// Visit a parse tree produced by apexParser#enumConstants.
	VisitEnumConstants(ctx *EnumConstantsContext) interface{}

	// Visit a parse tree produced by apexParser#enumConstant.
	VisitEnumConstant(ctx *EnumConstantContext) interface{}

	// Visit a parse tree produced by apexParser#enumBodyDeclarations.
	VisitEnumBodyDeclarations(ctx *EnumBodyDeclarationsContext) interface{}

	// Visit a parse tree produced by apexParser#interfaceDeclaration.
	VisitInterfaceDeclaration(ctx *InterfaceDeclarationContext) interface{}

	// Visit a parse tree produced by apexParser#typeList.
	VisitTypeList(ctx *TypeListContext) interface{}

	// Visit a parse tree produced by apexParser#classBody.
	VisitClassBody(ctx *ClassBodyContext) interface{}

	// Visit a parse tree produced by apexParser#interfaceBody.
	VisitInterfaceBody(ctx *InterfaceBodyContext) interface{}

	// Visit a parse tree produced by apexParser#classBodyDeclaration.
	VisitClassBodyDeclaration(ctx *ClassBodyDeclarationContext) interface{}

	// Visit a parse tree produced by apexParser#memberDeclaration.
	VisitMemberDeclaration(ctx *MemberDeclarationContext) interface{}

	// Visit a parse tree produced by apexParser#methodDeclaration.
	VisitMethodDeclaration(ctx *MethodDeclarationContext) interface{}

	// Visit a parse tree produced by apexParser#constructorDeclaration.
	VisitConstructorDeclaration(ctx *ConstructorDeclarationContext) interface{}

	// Visit a parse tree produced by apexParser#fieldDeclaration.
	VisitFieldDeclaration(ctx *FieldDeclarationContext) interface{}

	// Visit a parse tree produced by apexParser#propertyDeclaration.
	VisitPropertyDeclaration(ctx *PropertyDeclarationContext) interface{}

	// Visit a parse tree produced by apexParser#propertyBodyDeclaration.
	VisitPropertyBodyDeclaration(ctx *PropertyBodyDeclarationContext) interface{}

	// Visit a parse tree produced by apexParser#interfaceBodyDeclaration.
	VisitInterfaceBodyDeclaration(ctx *InterfaceBodyDeclarationContext) interface{}

	// Visit a parse tree produced by apexParser#interfaceMemberDeclaration.
	VisitInterfaceMemberDeclaration(ctx *InterfaceMemberDeclarationContext) interface{}

	// Visit a parse tree produced by apexParser#constDeclaration.
	VisitConstDeclaration(ctx *ConstDeclarationContext) interface{}

	// Visit a parse tree produced by apexParser#constantDeclarator.
	VisitConstantDeclarator(ctx *ConstantDeclaratorContext) interface{}

	// Visit a parse tree produced by apexParser#interfaceMethodDeclaration.
	VisitInterfaceMethodDeclaration(ctx *InterfaceMethodDeclarationContext) interface{}

	// Visit a parse tree produced by apexParser#variableDeclarators.
	VisitVariableDeclarators(ctx *VariableDeclaratorsContext) interface{}

	// Visit a parse tree produced by apexParser#variableDeclarator.
	VisitVariableDeclarator(ctx *VariableDeclaratorContext) interface{}

	// Visit a parse tree produced by apexParser#variableDeclaratorId.
	VisitVariableDeclaratorId(ctx *VariableDeclaratorIdContext) interface{}

	// Visit a parse tree produced by apexParser#variableInitializer.
	VisitVariableInitializer(ctx *VariableInitializerContext) interface{}

	// Visit a parse tree produced by apexParser#arrayInitializer.
	VisitArrayInitializer(ctx *ArrayInitializerContext) interface{}

	// Visit a parse tree produced by apexParser#enumConstantName.
	VisitEnumConstantName(ctx *EnumConstantNameContext) interface{}

	// Visit a parse tree produced by apexParser#apexType.
	VisitApexType(ctx *ApexTypeContext) interface{}

	// Visit a parse tree produced by apexParser#typedArray.
	VisitTypedArray(ctx *TypedArrayContext) interface{}

	// Visit a parse tree produced by apexParser#classOrInterfaceType.
	VisitClassOrInterfaceType(ctx *ClassOrInterfaceTypeContext) interface{}

	// Visit a parse tree produced by apexParser#primitiveType.
	VisitPrimitiveType(ctx *PrimitiveTypeContext) interface{}

	// Visit a parse tree produced by apexParser#typeArguments.
	VisitTypeArguments(ctx *TypeArgumentsContext) interface{}

	// Visit a parse tree produced by apexParser#typeArgument.
	VisitTypeArgument(ctx *TypeArgumentContext) interface{}

	// Visit a parse tree produced by apexParser#qualifiedNameList.
	VisitQualifiedNameList(ctx *QualifiedNameListContext) interface{}

	// Visit a parse tree produced by apexParser#formalParameters.
	VisitFormalParameters(ctx *FormalParametersContext) interface{}

	// Visit a parse tree produced by apexParser#formalParameterList.
	VisitFormalParameterList(ctx *FormalParameterListContext) interface{}

	// Visit a parse tree produced by apexParser#formalParameter.
	VisitFormalParameter(ctx *FormalParameterContext) interface{}

	// Visit a parse tree produced by apexParser#lastFormalParameter.
	VisitLastFormalParameter(ctx *LastFormalParameterContext) interface{}

	// Visit a parse tree produced by apexParser#methodBody.
	VisitMethodBody(ctx *MethodBodyContext) interface{}

	// Visit a parse tree produced by apexParser#constructorBody.
	VisitConstructorBody(ctx *ConstructorBodyContext) interface{}

	// Visit a parse tree produced by apexParser#qualifiedName.
	VisitQualifiedName(ctx *QualifiedNameContext) interface{}

	// Visit a parse tree produced by apexParser#literal.
	VisitLiteral(ctx *LiteralContext) interface{}

	// Visit a parse tree produced by apexParser#annotation.
	VisitAnnotation(ctx *AnnotationContext) interface{}

	// Visit a parse tree produced by apexParser#annotationName.
	VisitAnnotationName(ctx *AnnotationNameContext) interface{}

	// Visit a parse tree produced by apexParser#elementValuePairs.
	VisitElementValuePairs(ctx *ElementValuePairsContext) interface{}

	// Visit a parse tree produced by apexParser#elementValuePair.
	VisitElementValuePair(ctx *ElementValuePairContext) interface{}

	// Visit a parse tree produced by apexParser#elementValue.
	VisitElementValue(ctx *ElementValueContext) interface{}

	// Visit a parse tree produced by apexParser#elementValueArrayInitializer.
	VisitElementValueArrayInitializer(ctx *ElementValueArrayInitializerContext) interface{}

	// Visit a parse tree produced by apexParser#block.
	VisitBlock(ctx *BlockContext) interface{}

	// Visit a parse tree produced by apexParser#blockStatement.
	VisitBlockStatement(ctx *BlockStatementContext) interface{}

	// Visit a parse tree produced by apexParser#localVariableDeclarationStatement.
	VisitLocalVariableDeclarationStatement(ctx *LocalVariableDeclarationStatementContext) interface{}

	// Visit a parse tree produced by apexParser#localVariableDeclaration.
	VisitLocalVariableDeclaration(ctx *LocalVariableDeclarationContext) interface{}

	// Visit a parse tree produced by apexParser#statement.
	VisitStatement(ctx *StatementContext) interface{}

	// Visit a parse tree produced by apexParser#propertyBlock.
	VisitPropertyBlock(ctx *PropertyBlockContext) interface{}

	// Visit a parse tree produced by apexParser#getter.
	VisitGetter(ctx *GetterContext) interface{}

	// Visit a parse tree produced by apexParser#setter.
	VisitSetter(ctx *SetterContext) interface{}

	// Visit a parse tree produced by apexParser#catchClause.
	VisitCatchClause(ctx *CatchClauseContext) interface{}

	// Visit a parse tree produced by apexParser#catchType.
	VisitCatchType(ctx *CatchTypeContext) interface{}

	// Visit a parse tree produced by apexParser#finallyBlock.
	VisitFinallyBlock(ctx *FinallyBlockContext) interface{}

	// Visit a parse tree produced by apexParser#whenStatements.
	VisitWhenStatements(ctx *WhenStatementsContext) interface{}

	// Visit a parse tree produced by apexParser#whenStatement.
	VisitWhenStatement(ctx *WhenStatementContext) interface{}

	// Visit a parse tree produced by apexParser#whenExpression.
	VisitWhenExpression(ctx *WhenExpressionContext) interface{}

	// Visit a parse tree produced by apexParser#forControl.
	VisitForControl(ctx *ForControlContext) interface{}

	// Visit a parse tree produced by apexParser#forInit.
	VisitForInit(ctx *ForInitContext) interface{}

	// Visit a parse tree produced by apexParser#enhancedForControl.
	VisitEnhancedForControl(ctx *EnhancedForControlContext) interface{}

	// Visit a parse tree produced by apexParser#forUpdate.
	VisitForUpdate(ctx *ForUpdateContext) interface{}

	// Visit a parse tree produced by apexParser#parExpression.
	VisitParExpression(ctx *ParExpressionContext) interface{}

	// Visit a parse tree produced by apexParser#expressionList.
	VisitExpressionList(ctx *ExpressionListContext) interface{}

	// Visit a parse tree produced by apexParser#statementExpression.
	VisitStatementExpression(ctx *StatementExpressionContext) interface{}

	// Visit a parse tree produced by apexParser#constantExpression.
	VisitConstantExpression(ctx *ConstantExpressionContext) interface{}

	// Visit a parse tree produced by apexParser#apexDbExpressionShort.
	VisitApexDbExpressionShort(ctx *ApexDbExpressionShortContext) interface{}

	// Visit a parse tree produced by apexParser#apexDbExpression.
	VisitApexDbExpression(ctx *ApexDbExpressionContext) interface{}

	// Visit a parse tree produced by apexParser#PrimaryExpression.
	VisitPrimaryExpression(ctx *PrimaryExpressionContext) interface{}

	// Visit a parse tree produced by apexParser#UnaryExpression.
	VisitUnaryExpression(ctx *UnaryExpressionContext) interface{}

	// Visit a parse tree produced by apexParser#MethodInvocation.
	VisitMethodInvocation(ctx *MethodInvocationContext) interface{}

	// Visit a parse tree produced by apexParser#ShiftExpression.
	VisitShiftExpression(ctx *ShiftExpressionContext) interface{}

	// Visit a parse tree produced by apexParser#NewObjectExpression.
	VisitNewObjectExpression(ctx *NewObjectExpressionContext) interface{}

	// Visit a parse tree produced by apexParser#TernalyExpression.
	VisitTernalyExpression(ctx *TernalyExpressionContext) interface{}

	// Visit a parse tree produced by apexParser#PreUnaryExpression.
	VisitPreUnaryExpression(ctx *PreUnaryExpressionContext) interface{}

	// Visit a parse tree produced by apexParser#ArrayAccess.
	VisitArrayAccess(ctx *ArrayAccessContext) interface{}

	// Visit a parse tree produced by apexParser#PostUnaryExpression.
	VisitPostUnaryExpression(ctx *PostUnaryExpressionContext) interface{}

	// Visit a parse tree produced by apexParser#OpExpression.
	VisitOpExpression(ctx *OpExpressionContext) interface{}

	// Visit a parse tree produced by apexParser#InstanceofExpression.
	VisitInstanceofExpression(ctx *InstanceofExpressionContext) interface{}

	// Visit a parse tree produced by apexParser#CastExpression.
	VisitCastExpression(ctx *CastExpressionContext) interface{}

	// Visit a parse tree produced by apexParser#FieldAccess.
	VisitFieldAccess(ctx *FieldAccessContext) interface{}

	// Visit a parse tree produced by apexParser#primary.
	VisitPrimary(ctx *PrimaryContext) interface{}

	// Visit a parse tree produced by apexParser#creator.
	VisitCreator(ctx *CreatorContext) interface{}

	// Visit a parse tree produced by apexParser#createdName.
	VisitCreatedName(ctx *CreatedNameContext) interface{}

	// Visit a parse tree produced by apexParser#innerCreator.
	VisitInnerCreator(ctx *InnerCreatorContext) interface{}

	// Visit a parse tree produced by apexParser#arrayCreatorRest.
	VisitArrayCreatorRest(ctx *ArrayCreatorRestContext) interface{}

	// Visit a parse tree produced by apexParser#mapCreatorRest.
	VisitMapCreatorRest(ctx *MapCreatorRestContext) interface{}

	// Visit a parse tree produced by apexParser#mapKey.
	VisitMapKey(ctx *MapKeyContext) interface{}

	// Visit a parse tree produced by apexParser#mapValue.
	VisitMapValue(ctx *MapValueContext) interface{}

	// Visit a parse tree produced by apexParser#setCreatorRest.
	VisitSetCreatorRest(ctx *SetCreatorRestContext) interface{}

	// Visit a parse tree produced by apexParser#setValue.
	VisitSetValue(ctx *SetValueContext) interface{}

	// Visit a parse tree produced by apexParser#classCreatorRest.
	VisitClassCreatorRest(ctx *ClassCreatorRestContext) interface{}

	// Visit a parse tree produced by apexParser#explicitGenericInvocation.
	VisitExplicitGenericInvocation(ctx *ExplicitGenericInvocationContext) interface{}

	// Visit a parse tree produced by apexParser#nonWildcardTypeArguments.
	VisitNonWildcardTypeArguments(ctx *NonWildcardTypeArgumentsContext) interface{}

	// Visit a parse tree produced by apexParser#typeArgumentsOrDiamond.
	VisitTypeArgumentsOrDiamond(ctx *TypeArgumentsOrDiamondContext) interface{}

	// Visit a parse tree produced by apexParser#nonWildcardTypeArgumentsOrDiamond.
	VisitNonWildcardTypeArgumentsOrDiamond(ctx *NonWildcardTypeArgumentsOrDiamondContext) interface{}

	// Visit a parse tree produced by apexParser#superSuffix.
	VisitSuperSuffix(ctx *SuperSuffixContext) interface{}

	// Visit a parse tree produced by apexParser#explicitGenericInvocationSuffix.
	VisitExplicitGenericInvocationSuffix(ctx *ExplicitGenericInvocationSuffixContext) interface{}

	// Visit a parse tree produced by apexParser#arguments.
	VisitArguments(ctx *ArgumentsContext) interface{}

	// Visit a parse tree produced by apexParser#soqlLiteral.
	VisitSoqlLiteral(ctx *SoqlLiteralContext) interface{}

	// Visit a parse tree produced by apexParser#query.
	VisitQuery(ctx *QueryContext) interface{}

	// Visit a parse tree produced by apexParser#selectClause.
	VisitSelectClause(ctx *SelectClauseContext) interface{}

	// Visit a parse tree produced by apexParser#fieldList.
	VisitFieldList(ctx *FieldListContext) interface{}

	// Visit a parse tree produced by apexParser#selectField.
	VisitSelectField(ctx *SelectFieldContext) interface{}

	// Visit a parse tree produced by apexParser#fromClause.
	VisitFromClause(ctx *FromClauseContext) interface{}

	// Visit a parse tree produced by apexParser#filterScope.
	VisitFilterScope(ctx *FilterScopeContext) interface{}

	// Visit a parse tree produced by apexParser#SoqlFieldReference.
	VisitSoqlFieldReference(ctx *SoqlFieldReferenceContext) interface{}

	// Visit a parse tree produced by apexParser#SoqlFunctionCall.
	VisitSoqlFunctionCall(ctx *SoqlFunctionCallContext) interface{}

	// Visit a parse tree produced by apexParser#subquery.
	VisitSubquery(ctx *SubqueryContext) interface{}

	// Visit a parse tree produced by apexParser#whereClause.
	VisitWhereClause(ctx *WhereClauseContext) interface{}

	// Visit a parse tree produced by apexParser#whereFields.
	VisitWhereFields(ctx *WhereFieldsContext) interface{}

	// Visit a parse tree produced by apexParser#whereField.
	VisitWhereField(ctx *WhereFieldContext) interface{}

	// Visit a parse tree produced by apexParser#limitClause.
	VisitLimitClause(ctx *LimitClauseContext) interface{}

	// Visit a parse tree produced by apexParser#orderClause.
	VisitOrderClause(ctx *OrderClauseContext) interface{}

	// Visit a parse tree produced by apexParser#bindVariable.
	VisitBindVariable(ctx *BindVariableContext) interface{}

	// Visit a parse tree produced by apexParser#soqlValue.
	VisitSoqlValue(ctx *SoqlValueContext) interface{}

	// Visit a parse tree produced by apexParser#withClause.
	VisitWithClause(ctx *WithClauseContext) interface{}

	// Visit a parse tree produced by apexParser#soqlFilteringExpression.
	VisitSoqlFilteringExpression(ctx *SoqlFilteringExpressionContext) interface{}

	// Visit a parse tree produced by apexParser#groupClause.
	VisitGroupClause(ctx *GroupClauseContext) interface{}

	// Visit a parse tree produced by apexParser#havingConditionExpression.
	VisitHavingConditionExpression(ctx *HavingConditionExpressionContext) interface{}

	// Visit a parse tree produced by apexParser#offsetClause.
	VisitOffsetClause(ctx *OffsetClauseContext) interface{}

	// Visit a parse tree produced by apexParser#viewClause.
	VisitViewClause(ctx *ViewClauseContext) interface{}

	// Visit a parse tree produced by apexParser#soslLiteral.
	VisitSoslLiteral(ctx *SoslLiteralContext) interface{}

	// Visit a parse tree produced by apexParser#soslQuery.
	VisitSoslQuery(ctx *SoslQueryContext) interface{}

	// Visit a parse tree produced by apexParser#soslReturningObject.
	VisitSoslReturningObject(ctx *SoslReturningObjectContext) interface{}

	// Visit a parse tree produced by apexParser#apexIdentifier.
	VisitApexIdentifier(ctx *ApexIdentifierContext) interface{}

	// Visit a parse tree produced by apexParser#typeIdentifier.
	VisitTypeIdentifier(ctx *TypeIdentifierContext) interface{}
}
