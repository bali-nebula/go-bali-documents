/*
................................................................................
.    Copyright (c) 2009-2025 Crater Dog Technologies.  All Rights Reserved.    .
................................................................................
.  DO NOT ALTER OR REMOVE COPYRIGHT NOTICES OR THIS FILE HEADER.               .
.                                                                              .
.  This code is free software; you can redistribute it and/or modify it under  .
.  the terms of The MIT License (MIT), as published by the Open Source         .
.  Initiative. (See https://opensource.org/license/MIT)                        .
................................................................................
*/

/*
┌────────────────────────────────── WARNING ───────────────────────────────────┐
│             This "module_api.go" file was automatically generated.           │
│      Updates to any part of this file—other than the Module Description      │
│             and the Global Functions sections may be overwritten.            │
└──────────────────────────────────────────────────────────────────────────────┘

Package "module" declares type aliases for the commonly used types declared in
the packages contained in this module.  It also provides constructors for each
commonly used class that is exported by the module.  Each constructor delegates
the actual construction process to its corresponding concrete class declared in
the corresponding package contained within this module.

For detailed documentation on this entire module refer to the wiki:
  - https://github.com/bali-nebula/go-bali-documents/wiki
*/
package module

import (
	ast "github.com/bali-nebula/go-bali-documents/v3/ast"
	gra "github.com/bali-nebula/go-bali-documents/v3/grammar"
	col "github.com/craterdog/go-collection-framework/v7"
)

// TYPE ALIASES

// Ast

type (
	AcceptClauseClassLike          = ast.AcceptClauseClassLike
	ActionClassLike                = ast.ActionClassLike
	AdditionalArgumentClassLike    = ast.AdditionalArgumentClassLike
	AdditionalAssociationClassLike = ast.AdditionalAssociationClassLike
	AdditionalIndexClassLike       = ast.AdditionalIndexClassLike
	AdditionalStatementClassLike   = ast.AdditionalStatementClassLike
	AdditionalValueClassLike       = ast.AdditionalValueClassLike
	AndClassLike                   = ast.AndClassLike
	AnnotatedAssociationClassLike  = ast.AnnotatedAssociationClassLike
	AnnotatedStatementClassLike    = ast.AnnotatedStatementClassLike
	AnnotatedValueClassLike        = ast.AnnotatedValueClassLike
	ArgumentClassLike              = ast.ArgumentClassLike
	ArgumentsClassLike             = ast.ArgumentsClassLike
	AssignmentClassLike            = ast.AssignmentClassLike
	AssociationClassLike           = ast.AssociationClassLike
	AtLevelClassLike               = ast.AtLevelClassLike
	BagClassLike                   = ast.BagClassLike
	BlockingClassLike              = ast.BlockingClassLike
	BracketClassLike               = ast.BracketClassLike
	BreakClauseClassLike           = ast.BreakClauseClassLike
	CheckoutClauseClassLike        = ast.CheckoutClauseClassLike
	CitedClassLike                 = ast.CitedClassLike
	CollectionClassLike            = ast.CollectionClassLike
	ColonEqualClassLike            = ast.ColonEqualClassLike
	CommentLineClassLike           = ast.CommentLineClassLike
	ComplementClassLike            = ast.ComplementClassLike
	ComponentClassLike             = ast.ComponentClassLike
	ConditionClassLike             = ast.ConditionClassLike
	ContinueClauseClassLike        = ast.ContinueClauseClassLike
	DashEqualClassLike             = ast.DashEqualClassLike
	DefaultEqualClassLike          = ast.DefaultEqualClassLike
	DiscardClauseClassLike         = ast.DiscardClauseClassLike
	DoClauseClassLike              = ast.DoClauseClassLike
	DocumentClassLike              = ast.DocumentClassLike
	DraftClassLike                 = ast.DraftClassLike
	ElementClassLike               = ast.ElementClassLike
	EntityClassLike                = ast.EntityClassLike
	EventClassLike                 = ast.EventClassLike
	ExceptionClassLike             = ast.ExceptionClassLike
	ExclusiveClassLike             = ast.ExclusiveClassLike
	ExclusiveRangeClassLike        = ast.ExclusiveRangeClassLike
	ExpressionClassLike            = ast.ExpressionClassLike
	FailureClassLike               = ast.FailureClassLike
	FlowClassLike                  = ast.FlowClassLike
	FunctionClassLike              = ast.FunctionClassLike
	IfClauseClassLike              = ast.IfClauseClassLike
	InclusiveClassLike             = ast.InclusiveClassLike
	InclusiveRangeClassLike        = ast.InclusiveRangeClassLike
	IndexClassLike                 = ast.IndexClassLike
	IndicesClassLike               = ast.IndicesClassLike
	IndirectClassLike              = ast.IndirectClassLike
	InductionClassLike             = ast.InductionClassLike
	InlineAttributesClassLike      = ast.InlineAttributesClassLike
	InlineParametersClassLike      = ast.InlineParametersClassLike
	InlineStatementsClassLike      = ast.InlineStatementsClassLike
	InlineValuesClassLike          = ast.InlineValuesClassLike
	InverseClassLike               = ast.InverseClassLike
	InversionClassLike             = ast.InversionClassLike
	InvocationClassLike            = ast.InvocationClassLike
	IorClassLike                   = ast.IorClassLike
	IsClassLike                    = ast.IsClassLike
	ItemClassLike                  = ast.ItemClassLike
	LetClauseClassLike             = ast.LetClauseClassLike
	LogicalClassLike               = ast.LogicalClassLike
	MagnitudeClassLike             = ast.MagnitudeClassLike
	MainClauseClassLike            = ast.MainClauseClassLike
	MatchHandlerClassLike          = ast.MatchHandlerClassLike
	MatchesClassLike               = ast.MatchesClassLike
	MessageClassLike               = ast.MessageClassLike
	MessagingClassLike             = ast.MessagingClassLike
	MethodClassLike                = ast.MethodClassLike
	MultilineAttributesClassLike   = ast.MultilineAttributesClassLike
	MultilineParametersClassLike   = ast.MultilineParametersClassLike
	MultilineStatementsClassLike   = ast.MultilineStatementsClassLike
	MultilineValuesClassLike       = ast.MultilineValuesClassLike
	NoAttributesClassLike          = ast.NoAttributesClassLike
	NoStatementsClassLike          = ast.NoStatementsClassLike
	NoValuesClassLike              = ast.NoValuesClassLike
	NotarizeClauseClassLike        = ast.NotarizeClauseClassLike
	NoticeClassLike                = ast.NoticeClassLike
	NumericalClassLike             = ast.NumericalClassLike
	OnClauseClassLike              = ast.OnClauseClassLike
	OperationClassLike             = ast.OperationClassLike
	OperatorClassLike              = ast.OperatorClassLike
	ParametersClassLike            = ast.ParametersClassLike
	PlusEqualClassLike             = ast.PlusEqualClassLike
	PostClauseClassLike            = ast.PostClauseClassLike
	PrecedenceClassLike            = ast.PrecedenceClassLike
	PredicateClassLike             = ast.PredicateClassLike
	PrimitiveClassLike             = ast.PrimitiveClassLike
	ProcedureClassLike             = ast.ProcedureClassLike
	PublishClauseClassLike         = ast.PublishClauseClassLike
	RecipientClassLike             = ast.RecipientClassLike
	ReferentClassLike              = ast.ReferentClassLike
	RejectClauseClassLike          = ast.RejectClauseClassLike
	RepositoryClassLike            = ast.RepositoryClassLike
	ResultClassLike                = ast.ResultClassLike
	RetrieveClauseClassLike        = ast.RetrieveClauseClassLike
	ReturnClauseClassLike          = ast.ReturnClauseClassLike
	SanClassLike                   = ast.SanClassLike
	SaveClauseClassLike            = ast.SaveClauseClassLike
	SelectClauseClassLike          = ast.SelectClauseClassLike
	SequenceClassLike              = ast.SequenceClassLike
	SlashEqualClassLike            = ast.SlashEqualClassLike
	StarEqualClassLike             = ast.StarEqualClassLike
	StatementClassLike             = ast.StatementClassLike
	StatementLineClassLike         = ast.StatementLineClassLike
	StringClassLike                = ast.StringClassLike
	SubcomponentClassLike          = ast.SubcomponentClassLike
	SubjectClassLike               = ast.SubjectClassLike
	TargetClassLike                = ast.TargetClassLike
	TemplateClassLike              = ast.TemplateClassLike
	ThrowClauseClassLike           = ast.ThrowClauseClassLike
	VariableClassLike              = ast.VariableClassLike
	WhileClauseClassLike           = ast.WhileClauseClassLike
	WithClauseClassLike            = ast.WithClauseClassLike
	XorClassLike                   = ast.XorClassLike
)

type (
	AcceptClauseLike          = ast.AcceptClauseLike
	ActionLike                = ast.ActionLike
	AdditionalArgumentLike    = ast.AdditionalArgumentLike
	AdditionalAssociationLike = ast.AdditionalAssociationLike
	AdditionalIndexLike       = ast.AdditionalIndexLike
	AdditionalStatementLike   = ast.AdditionalStatementLike
	AdditionalValueLike       = ast.AdditionalValueLike
	AndLike                   = ast.AndLike
	AnnotatedAssociationLike  = ast.AnnotatedAssociationLike
	AnnotatedStatementLike    = ast.AnnotatedStatementLike
	AnnotatedValueLike        = ast.AnnotatedValueLike
	ArgumentLike              = ast.ArgumentLike
	ArgumentsLike             = ast.ArgumentsLike
	AssignmentLike            = ast.AssignmentLike
	AssociationLike           = ast.AssociationLike
	AtLevelLike               = ast.AtLevelLike
	BagLike                   = ast.BagLike
	BlockingLike              = ast.BlockingLike
	BracketLike               = ast.BracketLike
	BreakClauseLike           = ast.BreakClauseLike
	CheckoutClauseLike        = ast.CheckoutClauseLike
	CitedLike                 = ast.CitedLike
	CollectionLike            = ast.CollectionLike
	ColonEqualLike            = ast.ColonEqualLike
	CommentLineLike           = ast.CommentLineLike
	ComplementLike            = ast.ComplementLike
	ComponentLike             = ast.ComponentLike
	ConditionLike             = ast.ConditionLike
	ContinueClauseLike        = ast.ContinueClauseLike
	DashEqualLike             = ast.DashEqualLike
	DefaultEqualLike          = ast.DefaultEqualLike
	DiscardClauseLike         = ast.DiscardClauseLike
	DoClauseLike              = ast.DoClauseLike
	DocumentLike              = ast.DocumentLike
	DraftLike                 = ast.DraftLike
	ElementLike               = ast.ElementLike
	EntityLike                = ast.EntityLike
	EventLike                 = ast.EventLike
	ExceptionLike             = ast.ExceptionLike
	ExclusiveLike             = ast.ExclusiveLike
	ExclusiveRangeLike        = ast.ExclusiveRangeLike
	ExpressionLike            = ast.ExpressionLike
	FailureLike               = ast.FailureLike
	FlowLike                  = ast.FlowLike
	FunctionLike              = ast.FunctionLike
	IfClauseLike              = ast.IfClauseLike
	InclusiveLike             = ast.InclusiveLike
	InclusiveRangeLike        = ast.InclusiveRangeLike
	IndexLike                 = ast.IndexLike
	IndicesLike               = ast.IndicesLike
	IndirectLike              = ast.IndirectLike
	InductionLike             = ast.InductionLike
	InlineAttributesLike      = ast.InlineAttributesLike
	InlineParametersLike      = ast.InlineParametersLike
	InlineStatementsLike      = ast.InlineStatementsLike
	InlineValuesLike          = ast.InlineValuesLike
	InverseLike               = ast.InverseLike
	InversionLike             = ast.InversionLike
	InvocationLike            = ast.InvocationLike
	IorLike                   = ast.IorLike
	IsLike                    = ast.IsLike
	ItemLike                  = ast.ItemLike
	LetClauseLike             = ast.LetClauseLike
	LogicalLike               = ast.LogicalLike
	MagnitudeLike             = ast.MagnitudeLike
	MainClauseLike            = ast.MainClauseLike
	MatchHandlerLike          = ast.MatchHandlerLike
	MatchesLike               = ast.MatchesLike
	MessageLike               = ast.MessageLike
	MessagingLike             = ast.MessagingLike
	MethodLike                = ast.MethodLike
	MultilineAttributesLike   = ast.MultilineAttributesLike
	MultilineParametersLike   = ast.MultilineParametersLike
	MultilineStatementsLike   = ast.MultilineStatementsLike
	MultilineValuesLike       = ast.MultilineValuesLike
	NoAttributesLike          = ast.NoAttributesLike
	NoStatementsLike          = ast.NoStatementsLike
	NoValuesLike              = ast.NoValuesLike
	NotarizeClauseLike        = ast.NotarizeClauseLike
	NoticeLike                = ast.NoticeLike
	NumericalLike             = ast.NumericalLike
	OnClauseLike              = ast.OnClauseLike
	OperationLike             = ast.OperationLike
	OperatorLike              = ast.OperatorLike
	ParametersLike            = ast.ParametersLike
	PlusEqualLike             = ast.PlusEqualLike
	PostClauseLike            = ast.PostClauseLike
	PrecedenceLike            = ast.PrecedenceLike
	PredicateLike             = ast.PredicateLike
	PrimitiveLike             = ast.PrimitiveLike
	ProcedureLike             = ast.ProcedureLike
	PublishClauseLike         = ast.PublishClauseLike
	RecipientLike             = ast.RecipientLike
	ReferentLike              = ast.ReferentLike
	RejectClauseLike          = ast.RejectClauseLike
	RepositoryLike            = ast.RepositoryLike
	ResultLike                = ast.ResultLike
	RetrieveClauseLike        = ast.RetrieveClauseLike
	ReturnClauseLike          = ast.ReturnClauseLike
	SanLike                   = ast.SanLike
	SaveClauseLike            = ast.SaveClauseLike
	SelectClauseLike          = ast.SelectClauseLike
	SequenceLike              = ast.SequenceLike
	SlashEqualLike            = ast.SlashEqualLike
	StarEqualLike             = ast.StarEqualLike
	StatementLike             = ast.StatementLike
	StatementLineLike         = ast.StatementLineLike
	StringLike                = ast.StringLike
	SubcomponentLike          = ast.SubcomponentLike
	SubjectLike               = ast.SubjectLike
	TargetLike                = ast.TargetLike
	TemplateLike              = ast.TemplateLike
	ThrowClauseLike           = ast.ThrowClauseLike
	VariableLike              = ast.VariableLike
	WhileClauseLike           = ast.WhileClauseLike
	WithClauseLike            = ast.WithClauseLike
	XorLike                   = ast.XorLike
)

// Grammar

type (
	TokenType = gra.TokenType
)

const (
	ErrorToken       = gra.ErrorToken
	AmpersandToken   = gra.AmpersandToken
	AngleToken       = gra.AngleToken
	ArrowToken       = gra.ArrowToken
	BinaryToken      = gra.BinaryToken
	BooleanToken     = gra.BooleanToken
	BytecodeToken    = gra.BytecodeToken
	CaretToken       = gra.CaretToken
	CitationToken    = gra.CitationToken
	CommentToken     = gra.CommentToken
	DashToken        = gra.DashToken
	DelimiterToken   = gra.DelimiterToken
	DotToken         = gra.DotToken
	DoubleToken      = gra.DoubleToken
	DurationToken    = gra.DurationToken
	EqualToken       = gra.EqualToken
	IdentifierToken  = gra.IdentifierToken
	LessToken        = gra.LessToken
	MomentToken      = gra.MomentToken
	MoreToken        = gra.MoreToken
	NameToken        = gra.NameToken
	NarrativeToken   = gra.NarrativeToken
	NewlineToken     = gra.NewlineToken
	NoteToken        = gra.NoteToken
	NumberToken      = gra.NumberToken
	PatternToken     = gra.PatternToken
	PercentageToken  = gra.PercentageToken
	PlusToken        = gra.PlusToken
	ProbabilityToken = gra.ProbabilityToken
	QuoteToken       = gra.QuoteToken
	ResourceToken    = gra.ResourceToken
	SlashToken       = gra.SlashToken
	SpaceToken       = gra.SpaceToken
	StarToken        = gra.StarToken
	SymbolToken      = gra.SymbolToken
	TagToken         = gra.TagToken
	VersionToken     = gra.VersionToken
)

type (
	FormatterClassLike = gra.FormatterClassLike
	ParserClassLike    = gra.ParserClassLike
	ProcessorClassLike = gra.ProcessorClassLike
	ScannerClassLike   = gra.ScannerClassLike
	TokenClassLike     = gra.TokenClassLike
	ValidatorClassLike = gra.ValidatorClassLike
	VisitorClassLike   = gra.VisitorClassLike
)

type (
	FormatterLike = gra.FormatterLike
	ParserLike    = gra.ParserLike
	ProcessorLike = gra.ProcessorLike
	ScannerLike   = gra.ScannerLike
	TokenLike     = gra.TokenLike
	ValidatorLike = gra.ValidatorLike
	VisitorLike   = gra.VisitorLike
)

type (
	Methodical = gra.Methodical
)

// CLASS ACCESSORS

// Ast

func AcceptClauseClass() AcceptClauseClassLike {
	return ast.AcceptClauseClass()
}

func AcceptClause(
	message ast.MessageLike,
) AcceptClauseLike {
	return AcceptClauseClass().AcceptClause(
		message,
	)
}

func ActionClass() ActionClassLike {
	return ast.ActionClass()
}

func Action(
	any_ any,
) ActionLike {
	return ActionClass().Action(
		any_,
	)
}

func AdditionalArgumentClass() AdditionalArgumentClassLike {
	return ast.AdditionalArgumentClass()
}

func AdditionalArgument(
	argument ast.ArgumentLike,
) AdditionalArgumentLike {
	return AdditionalArgumentClass().AdditionalArgument(
		argument,
	)
}

func AdditionalAssociationClass() AdditionalAssociationClassLike {
	return ast.AdditionalAssociationClass()
}

func AdditionalAssociation(
	association ast.AssociationLike,
) AdditionalAssociationLike {
	return AdditionalAssociationClass().AdditionalAssociation(
		association,
	)
}

func AdditionalIndexClass() AdditionalIndexClassLike {
	return ast.AdditionalIndexClass()
}

func AdditionalIndex(
	index ast.IndexLike,
) AdditionalIndexLike {
	return AdditionalIndexClass().AdditionalIndex(
		index,
	)
}

func AdditionalStatementClass() AdditionalStatementClassLike {
	return ast.AdditionalStatementClass()
}

func AdditionalStatement(
	statement ast.StatementLike,
) AdditionalStatementLike {
	return AdditionalStatementClass().AdditionalStatement(
		statement,
	)
}

func AdditionalValueClass() AdditionalValueClassLike {
	return ast.AdditionalValueClass()
}

func AdditionalValue(
	component ast.ComponentLike,
) AdditionalValueLike {
	return AdditionalValueClass().AdditionalValue(
		component,
	)
}

func AndClass() AndClassLike {
	return ast.AndClass()
}

func And() AndLike {
	return AndClass().And()
}

func AnnotatedAssociationClass() AnnotatedAssociationClassLike {
	return ast.AnnotatedAssociationClass()
}

func AnnotatedAssociation(
	association ast.AssociationLike,
	optionalNote string,
) AnnotatedAssociationLike {
	return AnnotatedAssociationClass().AnnotatedAssociation(
		association,
		optionalNote,
	)
}

func AnnotatedStatementClass() AnnotatedStatementClassLike {
	return ast.AnnotatedStatementClass()
}

func AnnotatedStatement(
	any_ any,
) AnnotatedStatementLike {
	return AnnotatedStatementClass().AnnotatedStatement(
		any_,
	)
}

func AnnotatedValueClass() AnnotatedValueClassLike {
	return ast.AnnotatedValueClass()
}

func AnnotatedValue(
	component ast.ComponentLike,
	optionalNote string,
) AnnotatedValueLike {
	return AnnotatedValueClass().AnnotatedValue(
		component,
		optionalNote,
	)
}

func ArgumentClass() ArgumentClassLike {
	return ast.ArgumentClass()
}

func Argument(
	identifier string,
) ArgumentLike {
	return ArgumentClass().Argument(
		identifier,
	)
}

func ArgumentsClass() ArgumentsClassLike {
	return ast.ArgumentsClass()
}

func Arguments(
	argument ast.ArgumentLike,
	additionalArguments col.Sequential[ast.AdditionalArgumentLike],
) ArgumentsLike {
	return ArgumentsClass().Arguments(
		argument,
		additionalArguments,
	)
}

func AssignmentClass() AssignmentClassLike {
	return ast.AssignmentClass()
}

func Assignment(
	any_ any,
) AssignmentLike {
	return AssignmentClass().Assignment(
		any_,
	)
}

func AssociationClass() AssociationClassLike {
	return ast.AssociationClass()
}

func Association(
	primitive ast.PrimitiveLike,
	component ast.ComponentLike,
) AssociationLike {
	return AssociationClass().Association(
		primitive,
		component,
	)
}

func AtLevelClass() AtLevelClassLike {
	return ast.AtLevelClass()
}

func AtLevel(
	expression ast.ExpressionLike,
) AtLevelLike {
	return AtLevelClass().AtLevel(
		expression,
	)
}

func BagClass() BagClassLike {
	return ast.BagClass()
}

func Bag(
	expression ast.ExpressionLike,
) BagLike {
	return BagClass().Bag(
		expression,
	)
}

func BlockingClass() BlockingClassLike {
	return ast.BlockingClass()
}

func Blocking(
	any_ any,
) BlockingLike {
	return BlockingClass().Blocking(
		any_,
	)
}

func BracketClass() BracketClassLike {
	return ast.BracketClass()
}

func Bracket(
	any_ any,
) BracketLike {
	return BracketClass().Bracket(
		any_,
	)
}

func BreakClauseClass() BreakClauseClassLike {
	return ast.BreakClauseClass()
}

func BreakClause() BreakClauseLike {
	return BreakClauseClass().BreakClause()
}

func CheckoutClauseClass() CheckoutClauseClassLike {
	return ast.CheckoutClauseClass()
}

func CheckoutClause(
	recipient ast.RecipientLike,
	optionalAtLevel ast.AtLevelLike,
	cited ast.CitedLike,
) CheckoutClauseLike {
	return CheckoutClauseClass().CheckoutClause(
		recipient,
		optionalAtLevel,
		cited,
	)
}

func CitedClass() CitedClassLike {
	return ast.CitedClass()
}

func Cited(
	expression ast.ExpressionLike,
) CitedLike {
	return CitedClass().Cited(
		expression,
	)
}

func CollectionClass() CollectionClassLike {
	return ast.CollectionClass()
}

func Collection(
	any_ any,
) CollectionLike {
	return CollectionClass().Collection(
		any_,
	)
}

func ColonEqualClass() ColonEqualClassLike {
	return ast.ColonEqualClass()
}

func ColonEqual() ColonEqualLike {
	return ColonEqualClass().ColonEqual()
}

func CommentLineClass() CommentLineClassLike {
	return ast.CommentLineClass()
}

func CommentLine(
	comment string,
) CommentLineLike {
	return CommentLineClass().CommentLine(
		comment,
	)
}

func ComplementClass() ComplementClassLike {
	return ast.ComplementClass()
}

func Complement(
	logical ast.LogicalLike,
) ComplementLike {
	return ComplementClass().Complement(
		logical,
	)
}

func ComponentClass() ComponentClassLike {
	return ast.ComponentClass()
}

func Component(
	entity ast.EntityLike,
	optionalParameters ast.ParametersLike,
) ComponentLike {
	return ComponentClass().Component(
		entity,
		optionalParameters,
	)
}

func ConditionClass() ConditionClassLike {
	return ast.ConditionClass()
}

func Condition(
	expression ast.ExpressionLike,
) ConditionLike {
	return ConditionClass().Condition(
		expression,
	)
}

func ContinueClauseClass() ContinueClauseClassLike {
	return ast.ContinueClauseClass()
}

func ContinueClause() ContinueClauseLike {
	return ContinueClauseClass().ContinueClause()
}

func DashEqualClass() DashEqualClassLike {
	return ast.DashEqualClass()
}

func DashEqual() DashEqualLike {
	return DashEqualClass().DashEqual()
}

func DefaultEqualClass() DefaultEqualClassLike {
	return ast.DefaultEqualClass()
}

func DefaultEqual() DefaultEqualLike {
	return DefaultEqualClass().DefaultEqual()
}

func DiscardClauseClass() DiscardClauseClassLike {
	return ast.DiscardClauseClass()
}

func DiscardClause(
	draft ast.DraftLike,
) DiscardClauseLike {
	return DiscardClauseClass().DiscardClause(
		draft,
	)
}

func DoClauseClass() DoClauseClassLike {
	return ast.DoClauseClass()
}

func DoClause(
	invocation ast.InvocationLike,
) DoClauseLike {
	return DoClauseClass().DoClause(
		invocation,
	)
}

func DocumentClass() DocumentClassLike {
	return ast.DocumentClass()
}

func Document(
	optionalNotice ast.NoticeLike,
	component ast.ComponentLike,
) DocumentLike {
	return DocumentClass().Document(
		optionalNotice,
		component,
	)
}

func DraftClass() DraftClassLike {
	return ast.DraftClass()
}

func Draft(
	expression ast.ExpressionLike,
) DraftLike {
	return DraftClass().Draft(
		expression,
	)
}

func ElementClass() ElementClassLike {
	return ast.ElementClass()
}

func Element(
	any_ any,
) ElementLike {
	return ElementClass().Element(
		any_,
	)
}

func EntityClass() EntityClassLike {
	return ast.EntityClass()
}

func Entity(
	any_ any,
) EntityLike {
	return EntityClass().Entity(
		any_,
	)
}

func EventClass() EventClassLike {
	return ast.EventClass()
}

func Event(
	expression ast.ExpressionLike,
) EventLike {
	return EventClass().Event(
		expression,
	)
}

func ExceptionClass() ExceptionClassLike {
	return ast.ExceptionClass()
}

func Exception(
	expression ast.ExpressionLike,
) ExceptionLike {
	return ExceptionClass().Exception(
		expression,
	)
}

func ExclusiveClass() ExclusiveClassLike {
	return ast.ExclusiveClass()
}

func Exclusive() ExclusiveLike {
	return ExclusiveClass().Exclusive()
}

func ExclusiveRangeClass() ExclusiveRangeClassLike {
	return ast.ExclusiveRangeClass()
}

func ExclusiveRange(
	primitive1 ast.PrimitiveLike,
	primitive2 ast.PrimitiveLike,
	bracket ast.BracketLike,
) ExclusiveRangeLike {
	return ExclusiveRangeClass().ExclusiveRange(
		primitive1,
		primitive2,
		bracket,
	)
}

func ExpressionClass() ExpressionClassLike {
	return ast.ExpressionClass()
}

func Expression(
	subject ast.SubjectLike,
	predicates col.Sequential[ast.PredicateLike],
) ExpressionLike {
	return ExpressionClass().Expression(
		subject,
		predicates,
	)
}

func FailureClass() FailureClassLike {
	return ast.FailureClass()
}

func Failure(
	symbol string,
) FailureLike {
	return FailureClass().Failure(
		symbol,
	)
}

func FlowClass() FlowClassLike {
	return ast.FlowClass()
}

func Flow(
	any_ any,
) FlowLike {
	return FlowClass().Flow(
		any_,
	)
}

func FunctionClass() FunctionClassLike {
	return ast.FunctionClass()
}

func Function(
	identifier string,
	optionalArguments ast.ArgumentsLike,
) FunctionLike {
	return FunctionClass().Function(
		identifier,
		optionalArguments,
	)
}

func IfClauseClass() IfClauseClassLike {
	return ast.IfClauseClass()
}

func IfClause(
	condition ast.ConditionLike,
	procedure ast.ProcedureLike,
) IfClauseLike {
	return IfClauseClass().IfClause(
		condition,
		procedure,
	)
}

func InclusiveClass() InclusiveClassLike {
	return ast.InclusiveClass()
}

func Inclusive() InclusiveLike {
	return InclusiveClass().Inclusive()
}

func InclusiveRangeClass() InclusiveRangeClassLike {
	return ast.InclusiveRangeClass()
}

func InclusiveRange(
	primitive1 ast.PrimitiveLike,
	primitive2 ast.PrimitiveLike,
	bracket ast.BracketLike,
) InclusiveRangeLike {
	return InclusiveRangeClass().InclusiveRange(
		primitive1,
		primitive2,
		bracket,
	)
}

func IndexClass() IndexClassLike {
	return ast.IndexClass()
}

func Index(
	expression ast.ExpressionLike,
) IndexLike {
	return IndexClass().Index(
		expression,
	)
}

func IndicesClass() IndicesClassLike {
	return ast.IndicesClass()
}

func Indices(
	index ast.IndexLike,
	additionalIndexes col.Sequential[ast.AdditionalIndexLike],
) IndicesLike {
	return IndicesClass().Indices(
		index,
		additionalIndexes,
	)
}

func IndirectClass() IndirectClassLike {
	return ast.IndirectClass()
}

func Indirect(
	any_ any,
) IndirectLike {
	return IndirectClass().Indirect(
		any_,
	)
}

func InductionClass() InductionClassLike {
	return ast.InductionClass()
}

func Induction(
	any_ any,
) InductionLike {
	return InductionClass().Induction(
		any_,
	)
}

func InlineAttributesClass() InlineAttributesClassLike {
	return ast.InlineAttributesClass()
}

func InlineAttributes(
	association ast.AssociationLike,
	additionalAssociations col.Sequential[ast.AdditionalAssociationLike],
) InlineAttributesLike {
	return InlineAttributesClass().InlineAttributes(
		association,
		additionalAssociations,
	)
}

func InlineParametersClass() InlineParametersClassLike {
	return ast.InlineParametersClass()
}

func InlineParameters(
	association ast.AssociationLike,
	additionalAssociations col.Sequential[ast.AdditionalAssociationLike],
) InlineParametersLike {
	return InlineParametersClass().InlineParameters(
		association,
		additionalAssociations,
	)
}

func InlineStatementsClass() InlineStatementsClassLike {
	return ast.InlineStatementsClass()
}

func InlineStatements(
	statement ast.StatementLike,
	additionalStatements col.Sequential[ast.AdditionalStatementLike],
) InlineStatementsLike {
	return InlineStatementsClass().InlineStatements(
		statement,
		additionalStatements,
	)
}

func InlineValuesClass() InlineValuesClassLike {
	return ast.InlineValuesClass()
}

func InlineValues(
	component ast.ComponentLike,
	additionalValues col.Sequential[ast.AdditionalValueLike],
) InlineValuesLike {
	return InlineValuesClass().InlineValues(
		component,
		additionalValues,
	)
}

func InverseClass() InverseClassLike {
	return ast.InverseClass()
}

func Inverse(
	any_ any,
) InverseLike {
	return InverseClass().Inverse(
		any_,
	)
}

func InversionClass() InversionClassLike {
	return ast.InversionClass()
}

func Inversion(
	inverse ast.InverseLike,
	numerical ast.NumericalLike,
) InversionLike {
	return InversionClass().Inversion(
		inverse,
		numerical,
	)
}

func InvocationClass() InvocationClassLike {
	return ast.InvocationClass()
}

func Invocation(
	any_ any,
) InvocationLike {
	return InvocationClass().Invocation(
		any_,
	)
}

func IorClass() IorClassLike {
	return ast.IorClass()
}

func Ior() IorLike {
	return IorClass().Ior()
}

func IsClass() IsClassLike {
	return ast.IsClass()
}

func Is() IsLike {
	return IsClass().Is()
}

func ItemClass() ItemClassLike {
	return ast.ItemClass()
}

func Item(
	symbol string,
) ItemLike {
	return ItemClass().Item(
		symbol,
	)
}

func LetClauseClass() LetClauseClassLike {
	return ast.LetClauseClass()
}

func LetClause(
	recipient ast.RecipientLike,
	assignment ast.AssignmentLike,
	expression ast.ExpressionLike,
) LetClauseLike {
	return LetClauseClass().LetClause(
		recipient,
		assignment,
		expression,
	)
}

func LogicalClass() LogicalClassLike {
	return ast.LogicalClass()
}

func Logical(
	any_ any,
) LogicalLike {
	return LogicalClass().Logical(
		any_,
	)
}

func MagnitudeClass() MagnitudeClassLike {
	return ast.MagnitudeClass()
}

func Magnitude(
	numerical ast.NumericalLike,
) MagnitudeLike {
	return MagnitudeClass().Magnitude(
		numerical,
	)
}

func MainClauseClass() MainClauseClassLike {
	return ast.MainClauseClass()
}

func MainClause(
	any_ any,
) MainClauseLike {
	return MainClauseClass().MainClause(
		any_,
	)
}

func MatchHandlerClass() MatchHandlerClassLike {
	return ast.MatchHandlerClass()
}

func MatchHandler(
	template ast.TemplateLike,
	procedure ast.ProcedureLike,
) MatchHandlerLike {
	return MatchHandlerClass().MatchHandler(
		template,
		procedure,
	)
}

func MatchesClass() MatchesClassLike {
	return ast.MatchesClass()
}

func Matches() MatchesLike {
	return MatchesClass().Matches()
}

func MessageClass() MessageClassLike {
	return ast.MessageClass()
}

func Message(
	expression ast.ExpressionLike,
) MessageLike {
	return MessageClass().Message(
		expression,
	)
}

func MessagingClass() MessagingClassLike {
	return ast.MessagingClass()
}

func Messaging(
	any_ any,
) MessagingLike {
	return MessagingClass().Messaging(
		any_,
	)
}

func MethodClass() MethodClassLike {
	return ast.MethodClass()
}

func Method(
	identifier1 string,
	blocking ast.BlockingLike,
	identifier2 string,
	optionalArguments ast.ArgumentsLike,
) MethodLike {
	return MethodClass().Method(
		identifier1,
		blocking,
		identifier2,
		optionalArguments,
	)
}

func MultilineAttributesClass() MultilineAttributesClassLike {
	return ast.MultilineAttributesClass()
}

func MultilineAttributes(
	newline string,
	annotatedAssociations col.Sequential[ast.AnnotatedAssociationLike],
) MultilineAttributesLike {
	return MultilineAttributesClass().MultilineAttributes(
		newline,
		annotatedAssociations,
	)
}

func MultilineParametersClass() MultilineParametersClassLike {
	return ast.MultilineParametersClass()
}

func MultilineParameters(
	newline string,
	annotatedAssociations col.Sequential[ast.AnnotatedAssociationLike],
) MultilineParametersLike {
	return MultilineParametersClass().MultilineParameters(
		newline,
		annotatedAssociations,
	)
}

func MultilineStatementsClass() MultilineStatementsClassLike {
	return ast.MultilineStatementsClass()
}

func MultilineStatements(
	newline string,
	annotatedStatements col.Sequential[ast.AnnotatedStatementLike],
) MultilineStatementsLike {
	return MultilineStatementsClass().MultilineStatements(
		newline,
		annotatedStatements,
	)
}

func MultilineValuesClass() MultilineValuesClassLike {
	return ast.MultilineValuesClass()
}

func MultilineValues(
	newline string,
	annotatedValues col.Sequential[ast.AnnotatedValueLike],
) MultilineValuesLike {
	return MultilineValuesClass().MultilineValues(
		newline,
		annotatedValues,
	)
}

func NoAttributesClass() NoAttributesClassLike {
	return ast.NoAttributesClass()
}

func NoAttributes() NoAttributesLike {
	return NoAttributesClass().NoAttributes()
}

func NoStatementsClass() NoStatementsClassLike {
	return ast.NoStatementsClass()
}

func NoStatements() NoStatementsLike {
	return NoStatementsClass().NoStatements()
}

func NoValuesClass() NoValuesClassLike {
	return ast.NoValuesClass()
}

func NoValues() NoValuesLike {
	return NoValuesClass().NoValues()
}

func NotarizeClauseClass() NotarizeClauseClassLike {
	return ast.NotarizeClauseClass()
}

func NotarizeClause(
	draft ast.DraftLike,
	cited ast.CitedLike,
) NotarizeClauseLike {
	return NotarizeClauseClass().NotarizeClause(
		draft,
		cited,
	)
}

func NoticeClass() NoticeClassLike {
	return ast.NoticeClass()
}

func Notice(
	comment string,
	newline string,
) NoticeLike {
	return NoticeClass().Notice(
		comment,
		newline,
	)
}

func NumericalClass() NumericalClassLike {
	return ast.NumericalClass()
}

func Numerical(
	any_ any,
) NumericalLike {
	return NumericalClass().Numerical(
		any_,
	)
}

func OnClauseClass() OnClauseClassLike {
	return ast.OnClauseClass()
}

func OnClause(
	failure ast.FailureLike,
	matchHandlers col.Sequential[ast.MatchHandlerLike],
) OnClauseLike {
	return OnClauseClass().OnClause(
		failure,
		matchHandlers,
	)
}

func OperationClass() OperationClassLike {
	return ast.OperationClass()
}

func Operation(
	any_ any,
) OperationLike {
	return OperationClass().Operation(
		any_,
	)
}

func OperatorClass() OperatorClassLike {
	return ast.OperatorClass()
}

func Operator(
	any_ any,
) OperatorLike {
	return OperatorClass().Operator(
		any_,
	)
}

func ParametersClass() ParametersClassLike {
	return ast.ParametersClass()
}

func Parameters(
	any_ any,
) ParametersLike {
	return ParametersClass().Parameters(
		any_,
	)
}

func PlusEqualClass() PlusEqualClassLike {
	return ast.PlusEqualClass()
}

func PlusEqual() PlusEqualLike {
	return PlusEqualClass().PlusEqual()
}

func PostClauseClass() PostClauseClassLike {
	return ast.PostClauseClass()
}

func PostClause(
	message ast.MessageLike,
	bag ast.BagLike,
) PostClauseLike {
	return PostClauseClass().PostClause(
		message,
		bag,
	)
}

func PrecedenceClass() PrecedenceClassLike {
	return ast.PrecedenceClass()
}

func Precedence(
	expression ast.ExpressionLike,
) PrecedenceLike {
	return PrecedenceClass().Precedence(
		expression,
	)
}

func PredicateClass() PredicateClassLike {
	return ast.PredicateClass()
}

func Predicate(
	action ast.ActionLike,
	expression ast.ExpressionLike,
) PredicateLike {
	return PredicateClass().Predicate(
		action,
		expression,
	)
}

func PrimitiveClass() PrimitiveClassLike {
	return ast.PrimitiveClass()
}

func Primitive(
	any_ any,
) PrimitiveLike {
	return PrimitiveClass().Primitive(
		any_,
	)
}

func ProcedureClass() ProcedureClassLike {
	return ast.ProcedureClass()
}

func Procedure(
	any_ any,
) ProcedureLike {
	return ProcedureClass().Procedure(
		any_,
	)
}

func PublishClauseClass() PublishClauseClassLike {
	return ast.PublishClauseClass()
}

func PublishClause(
	event ast.EventLike,
) PublishClauseLike {
	return PublishClauseClass().PublishClause(
		event,
	)
}

func RecipientClass() RecipientClassLike {
	return ast.RecipientClass()
}

func Recipient(
	any_ any,
) RecipientLike {
	return RecipientClass().Recipient(
		any_,
	)
}

func ReferentClass() ReferentClassLike {
	return ast.ReferentClass()
}

func Referent(
	indirect ast.IndirectLike,
) ReferentLike {
	return ReferentClass().Referent(
		indirect,
	)
}

func RejectClauseClass() RejectClauseClassLike {
	return ast.RejectClauseClass()
}

func RejectClause(
	message ast.MessageLike,
) RejectClauseLike {
	return RejectClauseClass().RejectClause(
		message,
	)
}

func RepositoryClass() RepositoryClassLike {
	return ast.RepositoryClass()
}

func Repository(
	any_ any,
) RepositoryLike {
	return RepositoryClass().Repository(
		any_,
	)
}

func ResultClass() ResultClassLike {
	return ast.ResultClass()
}

func Result(
	expression ast.ExpressionLike,
) ResultLike {
	return ResultClass().Result(
		expression,
	)
}

func RetrieveClauseClass() RetrieveClauseClassLike {
	return ast.RetrieveClauseClass()
}

func RetrieveClause(
	recipient ast.RecipientLike,
	bag ast.BagLike,
) RetrieveClauseLike {
	return RetrieveClauseClass().RetrieveClause(
		recipient,
		bag,
	)
}

func ReturnClauseClass() ReturnClauseClassLike {
	return ast.ReturnClauseClass()
}

func ReturnClause(
	result ast.ResultLike,
) ReturnClauseLike {
	return ReturnClauseClass().ReturnClause(
		result,
	)
}

func SanClass() SanClassLike {
	return ast.SanClass()
}

func San() SanLike {
	return SanClass().San()
}

func SaveClauseClass() SaveClauseClassLike {
	return ast.SaveClauseClass()
}

func SaveClause(
	draft ast.DraftLike,
	cited ast.CitedLike,
) SaveClauseLike {
	return SaveClauseClass().SaveClause(
		draft,
		cited,
	)
}

func SelectClauseClass() SelectClauseClassLike {
	return ast.SelectClauseClass()
}

func SelectClause(
	target ast.TargetLike,
	matchHandlers col.Sequential[ast.MatchHandlerLike],
) SelectClauseLike {
	return SelectClauseClass().SelectClause(
		target,
		matchHandlers,
	)
}

func SequenceClass() SequenceClassLike {
	return ast.SequenceClass()
}

func Sequence(
	expression ast.ExpressionLike,
) SequenceLike {
	return SequenceClass().Sequence(
		expression,
	)
}

func SlashEqualClass() SlashEqualClassLike {
	return ast.SlashEqualClass()
}

func SlashEqual() SlashEqualLike {
	return SlashEqualClass().SlashEqual()
}

func StarEqualClass() StarEqualClassLike {
	return ast.StarEqualClass()
}

func StarEqual() StarEqualLike {
	return StarEqualClass().StarEqual()
}

func StatementClass() StatementClassLike {
	return ast.StatementClass()
}

func Statement(
	mainClause ast.MainClauseLike,
	optionalOnClause ast.OnClauseLike,
) StatementLike {
	return StatementClass().Statement(
		mainClause,
		optionalOnClause,
	)
}

func StatementLineClass() StatementLineClassLike {
	return ast.StatementLineClass()
}

func StatementLine(
	statement ast.StatementLike,
	optionalNote string,
) StatementLineLike {
	return StatementLineClass().StatementLine(
		statement,
		optionalNote,
	)
}

func StringClass() StringClassLike {
	return ast.StringClass()
}

func String(
	any_ any,
) StringLike {
	return StringClass().String(
		any_,
	)
}

func SubcomponentClass() SubcomponentClassLike {
	return ast.SubcomponentClass()
}

func Subcomponent(
	identifier string,
	indices ast.IndicesLike,
) SubcomponentLike {
	return SubcomponentClass().Subcomponent(
		identifier,
		indices,
	)
}

func SubjectClass() SubjectClassLike {
	return ast.SubjectClass()
}

func Subject(
	any_ any,
) SubjectLike {
	return SubjectClass().Subject(
		any_,
	)
}

func TargetClass() TargetClassLike {
	return ast.TargetClass()
}

func Target(
	any_ any,
) TargetLike {
	return TargetClass().Target(
		any_,
	)
}

func TemplateClass() TemplateClassLike {
	return ast.TemplateClass()
}

func Template(
	expression ast.ExpressionLike,
) TemplateLike {
	return TemplateClass().Template(
		expression,
	)
}

func ThrowClauseClass() ThrowClauseClassLike {
	return ast.ThrowClauseClass()
}

func ThrowClause(
	exception ast.ExceptionLike,
) ThrowClauseLike {
	return ThrowClauseClass().ThrowClause(
		exception,
	)
}

func VariableClass() VariableClassLike {
	return ast.VariableClass()
}

func Variable(
	identifier string,
) VariableLike {
	return VariableClass().Variable(
		identifier,
	)
}

func WhileClauseClass() WhileClauseClassLike {
	return ast.WhileClauseClass()
}

func WhileClause(
	condition ast.ConditionLike,
	procedure ast.ProcedureLike,
) WhileClauseLike {
	return WhileClauseClass().WhileClause(
		condition,
		procedure,
	)
}

func WithClauseClass() WithClauseClassLike {
	return ast.WithClauseClass()
}

func WithClause(
	item ast.ItemLike,
	sequence ast.SequenceLike,
	procedure ast.ProcedureLike,
) WithClauseLike {
	return WithClauseClass().WithClause(
		item,
		sequence,
		procedure,
	)
}

func XorClass() XorClassLike {
	return ast.XorClass()
}

func Xor() XorLike {
	return XorClass().Xor()
}

// Grammar

func FormatterClass() FormatterClassLike {
	return gra.FormatterClass()
}

func Formatter() FormatterLike {
	return FormatterClass().Formatter()
}

func ParserClass() ParserClassLike {
	return gra.ParserClass()
}

func Parser() ParserLike {
	return ParserClass().Parser()
}

func ProcessorClass() ProcessorClassLike {
	return gra.ProcessorClass()
}

func Processor() ProcessorLike {
	return ProcessorClass().Processor()
}

func ScannerClass() ScannerClassLike {
	return gra.ScannerClass()
}

func Scanner(
	source string,
	tokens col.QueueLike[gra.TokenLike],
) ScannerLike {
	return ScannerClass().Scanner(
		source,
		tokens,
	)
}

func TokenClass() TokenClassLike {
	return gra.TokenClass()
}

func Token(
	line uint,
	position uint,
	type_ gra.TokenType,
	value string,
) TokenLike {
	return TokenClass().Token(
		line,
		position,
		type_,
		value,
	)
}

func ValidatorClass() ValidatorClassLike {
	return gra.ValidatorClass()
}

func Validator() ValidatorLike {
	return ValidatorClass().Validator()
}

func VisitorClass() VisitorClassLike {
	return gra.VisitorClass()
}

func Visitor(
	processor gra.Methodical,
) VisitorLike {
	return VisitorClass().Visitor(
		processor,
	)
}

// GLOBAL FUNCTIONS
