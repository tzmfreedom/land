package main

import (
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
	modifiers := make([]ast.Modifier, len(classOrInterfaceModifiers))
	for i, modifier := range classOrInterfaceModifiers {
		m := modifier.Accept(v)
		modifiers[i], _ = m.(ast.Modifier)
	}

	if ctx.ClassDeclaration() != nil {
		cd := ctx.ClassDeclaration().Accept(v)
		classDeclaration, _ := cd.(ast.ClassDeclaration)
		classDeclaration.Modifiers = modifiers
		return classDeclaration
	}
	return nil
}

func (v *AstBuilder) VisitTriggerDeclaration(ctx *parser.TriggerDeclarationContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *AstBuilder) VisitTriggerTimings(ctx *parser.TriggerTimingsContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *AstBuilder) VisitTriggerTiming(ctx *parser.TriggerTimingContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *AstBuilder) VisitModifier(ctx *parser.ModifierContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *AstBuilder) VisitClassOrInterfaceModifier(ctx *parser.ClassOrInterfaceModifierContext) interface{} {
	if ctx.Annotation() != nil {
		return ctx.Annotation().Accept(v)
	}
	return ast.Modifier{
		Name:     ctx.GetText(),
		Position: v.newPosition(ctx),
	}
}

func (v *AstBuilder) VisitVariableModifier(ctx *parser.VariableModifierContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *AstBuilder) VisitClassDeclaration(ctx *parser.ClassDeclarationContext) interface{} {
	return ast.ClassDeclaration{
		Name:     ctx.ApexIdentifier().GetText(),
		Position: v.newPosition(ctx),
	}
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
	return v.VisitChildren(ctx)
}

func (v *AstBuilder) VisitClassBody(ctx *parser.ClassBodyContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *AstBuilder) VisitInterfaceBody(ctx *parser.InterfaceBodyContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *AstBuilder) VisitClassBodyDeclaration(ctx *parser.ClassBodyDeclarationContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *AstBuilder) VisitMemberDeclaration(ctx *parser.MemberDeclarationContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *AstBuilder) VisitMethodDeclaration(ctx *parser.MethodDeclarationContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *AstBuilder) VisitConstructorDeclaration(ctx *parser.ConstructorDeclarationContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *AstBuilder) VisitFieldDeclaration(ctx *parser.FieldDeclarationContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *AstBuilder) VisitPropertyDeclaration(ctx *parser.PropertyDeclarationContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *AstBuilder) VisitPropertyBodyDeclaration(ctx *parser.PropertyBodyDeclarationContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *AstBuilder) VisitInterfaceBodyDeclaration(ctx *parser.InterfaceBodyDeclarationContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *AstBuilder) VisitInterfaceMemberDeclaration(ctx *parser.InterfaceMemberDeclarationContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *AstBuilder) VisitConstDeclaration(ctx *parser.ConstDeclarationContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *AstBuilder) VisitConstantDeclarator(ctx *parser.ConstantDeclaratorContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *AstBuilder) VisitInterfaceMethodDeclaration(ctx *parser.InterfaceMethodDeclarationContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *AstBuilder) VisitVariableDeclarators(ctx *parser.VariableDeclaratorsContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *AstBuilder) VisitVariableDeclarator(ctx *parser.VariableDeclaratorContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *AstBuilder) VisitVariableDeclaratorId(ctx *parser.VariableDeclaratorIdContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *AstBuilder) VisitVariableInitializer(ctx *parser.VariableInitializerContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *AstBuilder) VisitArrayInitializer(ctx *parser.ArrayInitializerContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *AstBuilder) VisitEnumConstantName(ctx *parser.EnumConstantNameContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *AstBuilder) VisitApexType(ctx *parser.ApexTypeContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *AstBuilder) VisitTypedArray(ctx *parser.TypedArrayContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *AstBuilder) VisitClassOrInterfaceType(ctx *parser.ClassOrInterfaceTypeContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *AstBuilder) VisitPrimitiveType(ctx *parser.PrimitiveTypeContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *AstBuilder) VisitTypeArguments(ctx *parser.TypeArgumentsContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *AstBuilder) VisitTypeArgument(ctx *parser.TypeArgumentContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *AstBuilder) VisitQualifiedNameList(ctx *parser.QualifiedNameListContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *AstBuilder) VisitFormalParameters(ctx *parser.FormalParametersContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *AstBuilder) VisitFormalParameterList(ctx *parser.FormalParameterListContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *AstBuilder) VisitFormalParameter(ctx *parser.FormalParameterContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *AstBuilder) VisitLastFormalParameter(ctx *parser.LastFormalParameterContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *AstBuilder) VisitMethodBody(ctx *parser.MethodBodyContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *AstBuilder) VisitConstructorBody(ctx *parser.ConstructorBodyContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *AstBuilder) VisitQualifiedName(ctx *parser.QualifiedNameContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *AstBuilder) VisitLiteral(ctx *parser.LiteralContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *AstBuilder) VisitAnnotation(ctx *parser.AnnotationContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *AstBuilder) VisitAnnotationName(ctx *parser.AnnotationNameContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *AstBuilder) VisitElementValuePairs(ctx *parser.ElementValuePairsContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *AstBuilder) VisitElementValuePair(ctx *parser.ElementValuePairContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *AstBuilder) VisitElementValue(ctx *parser.ElementValueContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *AstBuilder) VisitElementValueArrayInitializer(ctx *parser.ElementValueArrayInitializerContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *AstBuilder) VisitBlock(ctx *parser.BlockContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *AstBuilder) VisitBlockStatement(ctx *parser.BlockStatementContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *AstBuilder) VisitLocalVariableDeclarationStatement(ctx *parser.LocalVariableDeclarationStatementContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *AstBuilder) VisitLocalVariableDeclaration(ctx *parser.LocalVariableDeclarationContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *AstBuilder) VisitStatement(ctx *parser.StatementContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *AstBuilder) VisitPropertyBlock(ctx *parser.PropertyBlockContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *AstBuilder) VisitGetter(ctx *parser.GetterContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *AstBuilder) VisitSetter(ctx *parser.SetterContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *AstBuilder) VisitCatchClause(ctx *parser.CatchClauseContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *AstBuilder) VisitCatchType(ctx *parser.CatchTypeContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *AstBuilder) VisitFinallyBlock(ctx *parser.FinallyBlockContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *AstBuilder) VisitWhenStatements(ctx *parser.WhenStatementsContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *AstBuilder) VisitWhenStatement(ctx *parser.WhenStatementContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *AstBuilder) VisitWhenExpression(ctx *parser.WhenExpressionContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *AstBuilder) VisitForControl(ctx *parser.ForControlContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *AstBuilder) VisitForInit(ctx *parser.ForInitContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *AstBuilder) VisitEnhancedForControl(ctx *parser.EnhancedForControlContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *AstBuilder) VisitForUpdate(ctx *parser.ForUpdateContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *AstBuilder) VisitParExpression(ctx *parser.ParExpressionContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *AstBuilder) VisitExpressionList(ctx *parser.ExpressionListContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *AstBuilder) VisitStatementExpression(ctx *parser.StatementExpressionContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *AstBuilder) VisitConstantExpression(ctx *parser.ConstantExpressionContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *AstBuilder) VisitApexDbExpressionShort(ctx *parser.ApexDbExpressionShortContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *AstBuilder) VisitApexDbExpression(ctx *parser.ApexDbExpressionContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *AstBuilder) VisitTernalyExpression(ctx *parser.TernalyExpressionContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *AstBuilder) VisitPreUnaryExpression(ctx *parser.PreUnaryExpressionContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *AstBuilder) VisitArrayAccess(ctx *parser.ArrayAccessContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *AstBuilder) VisitPostUnaryExpression(ctx *parser.PostUnaryExpressionContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *AstBuilder) VisitPrimaryExpression(ctx *parser.PrimaryExpressionContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *AstBuilder) VisitOpExpression(ctx *parser.OpExpressionContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *AstBuilder) VisitNewExpression(ctx *parser.NewObjectExpressionContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *AstBuilder) VisitUnaryExpression(ctx *parser.UnaryExpressionContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *AstBuilder) VisitMethodInvocation(ctx *parser.MethodInvocationContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *AstBuilder) VisitCastExpression(ctx *parser.CastExpressionContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *AstBuilder) VisitShiftExpression(ctx *parser.ShiftExpressionContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *AstBuilder) VisitFieldAccess(ctx *parser.FieldAccessContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *AstBuilder) VisitPrimary(ctx *parser.PrimaryContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *AstBuilder) VisitCreator(ctx *parser.CreatorContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *AstBuilder) VisitCreatedName(ctx *parser.CreatedNameContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *AstBuilder) VisitInnerCreator(ctx *parser.InnerCreatorContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *AstBuilder) VisitArrayCreatorRest(ctx *parser.ArrayCreatorRestContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *AstBuilder) VisitMapCreatorRest(ctx *parser.MapCreatorRestContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *AstBuilder) VisitSetCreatorRest(ctx *parser.SetCreatorRestContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *AstBuilder) VisitClassCreatorRest(ctx *parser.ClassCreatorRestContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *AstBuilder) VisitExplicitGenericInvocation(ctx *parser.ExplicitGenericInvocationContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *AstBuilder) VisitNonWildcardTypeArguments(ctx *parser.NonWildcardTypeArgumentsContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *AstBuilder) VisitTypeArgumentsOrDiamond(ctx *parser.TypeArgumentsOrDiamondContext) interface{} {
	return v.VisitChildren(ctx)
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
	return v.VisitChildren(ctx)
}

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

func (v *AstBuilder) VisitApexIdentifier(ctx *parser.ApexIdentifierContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *AstBuilder) VisitTypeIdentifier(ctx *parser.TypeIdentifierContext) interface{} {
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
