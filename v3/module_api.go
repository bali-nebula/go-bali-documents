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
	AdditionalArgumentClassLike    = ast.AdditionalArgumentClassLike
	AdditionalAssociationClassLike = ast.AdditionalAssociationClassLike
	AdditionalIndexClassLike       = ast.AdditionalIndexClassLike
	AdditionalStatementClassLike   = ast.AdditionalStatementClassLike
	AdditionalValueClassLike       = ast.AdditionalValueClassLike
	AnnotatedAssociationClassLike  = ast.AnnotatedAssociationClassLike
	AnnotatedStatementClassLike    = ast.AnnotatedStatementClassLike
	AnnotatedValueClassLike        = ast.AnnotatedValueClassLike
	ArgumentClassLike              = ast.ArgumentClassLike
	ArgumentsClassLike             = ast.ArgumentsClassLike
	ArithmeticClassLike            = ast.ArithmeticClassLike
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
	CommentLineClassLike           = ast.CommentLineClassLike
	ComparisonClassLike            = ast.ComparisonClassLike
	ComplementClassLike            = ast.ComplementClassLike
	ComponentClassLike             = ast.ComponentClassLike
	ConditionClassLike             = ast.ConditionClassLike
	ContinueClauseClassLike        = ast.ContinueClauseClassLike
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
	ItemClassLike                  = ast.ItemClassLike
	LegalNoticeClassLike           = ast.LegalNoticeClassLike
	LetClauseClassLike             = ast.LetClauseClassLike
	LogicClassLike                 = ast.LogicClassLike
	LogicalClassLike               = ast.LogicalClassLike
	MagnitudeClassLike             = ast.MagnitudeClassLike
	MainClauseClassLike            = ast.MainClauseClassLike
	MatchHandlerClassLike          = ast.MatchHandlerClassLike
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
	NumericalClassLike             = ast.NumericalClassLike
	OnClauseClassLike              = ast.OnClauseClassLike
	OperationClassLike             = ast.OperationClassLike
	ParametersClassLike            = ast.ParametersClassLike
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
	SaveClauseClassLike            = ast.SaveClauseClassLike
	SelectClauseClassLike          = ast.SelectClauseClassLike
	SequenceClassLike              = ast.SequenceClassLike
	StatementClassLike             = ast.StatementClassLike
	StatementLineClassLike         = ast.StatementLineClassLike
	StringClassLike                = ast.StringClassLike
	SubcomponentClassLike          = ast.SubcomponentClassLike
	SubjectClassLike               = ast.SubjectClassLike
	TargetClassLike                = ast.TargetClassLike
	TemplateClassLike              = ast.TemplateClassLike
	TextualClassLike               = ast.TextualClassLike
	ThrowClauseClassLike           = ast.ThrowClauseClassLike
	VariableClassLike              = ast.VariableClassLike
	WhileClauseClassLike           = ast.WhileClauseClassLike
	WithClauseClassLike            = ast.WithClauseClassLike
)

type (
	AcceptClauseLike          = ast.AcceptClauseLike
	AdditionalArgumentLike    = ast.AdditionalArgumentLike
	AdditionalAssociationLike = ast.AdditionalAssociationLike
	AdditionalIndexLike       = ast.AdditionalIndexLike
	AdditionalStatementLike   = ast.AdditionalStatementLike
	AdditionalValueLike       = ast.AdditionalValueLike
	AnnotatedAssociationLike  = ast.AnnotatedAssociationLike
	AnnotatedStatementLike    = ast.AnnotatedStatementLike
	AnnotatedValueLike        = ast.AnnotatedValueLike
	ArgumentLike              = ast.ArgumentLike
	ArgumentsLike             = ast.ArgumentsLike
	ArithmeticLike            = ast.ArithmeticLike
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
	CommentLineLike           = ast.CommentLineLike
	ComparisonLike            = ast.ComparisonLike
	ComplementLike            = ast.ComplementLike
	ComponentLike             = ast.ComponentLike
	ConditionLike             = ast.ConditionLike
	ContinueClauseLike        = ast.ContinueClauseLike
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
	ItemLike                  = ast.ItemLike
	LegalNoticeLike           = ast.LegalNoticeLike
	LetClauseLike             = ast.LetClauseLike
	LogicLike                 = ast.LogicLike
	LogicalLike               = ast.LogicalLike
	MagnitudeLike             = ast.MagnitudeLike
	MainClauseLike            = ast.MainClauseLike
	MatchHandlerLike          = ast.MatchHandlerLike
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
	NumericalLike             = ast.NumericalLike
	OnClauseLike              = ast.OnClauseLike
	OperationLike             = ast.OperationLike
	ParametersLike            = ast.ParametersLike
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
	SaveClauseLike            = ast.SaveClauseLike
	SelectClauseLike          = ast.SelectClauseLike
	SequenceLike              = ast.SequenceLike
	StatementLike             = ast.StatementLike
	StatementLineLike         = ast.StatementLineLike
	StringLike                = ast.StringLike
	SubcomponentLike          = ast.SubcomponentLike
	SubjectLike               = ast.SubjectLike
	TargetLike                = ast.TargetLike
	TemplateLike              = ast.TemplateLike
	TextualLike               = ast.TextualLike
	ThrowClauseLike           = ast.ThrowClauseLike
	VariableLike              = ast.VariableLike
	WhileClauseLike           = ast.WhileClauseLike
	WithClauseLike            = ast.WithClauseLike
)

// Grammar

type (
	TokenType = gra.TokenType
)

const (
	ErrorToken       = gra.ErrorToken
	AngleToken       = gra.AngleToken
	ArrowToken       = gra.ArrowToken
	BinaryToken      = gra.BinaryToken
	BooleanToken     = gra.BooleanToken
	BytecodeToken    = gra.BytecodeToken
	CaretToken       = gra.CaretToken
	CitationToken    = gra.CitationToken
	CommentToken     = gra.CommentToken
	DelimiterToken   = gra.DelimiterToken
	DotToken         = gra.DotToken
	DurationToken    = gra.DurationToken
	EqualToken       = gra.EqualToken
	IdentifierToken  = gra.IdentifierToken
	LessToken        = gra.LessToken
	MinusToken       = gra.MinusToken
	MomentToken      = gra.MomentToken
	MoreToken        = gra.MoreToken
	NameToken        = gra.NameToken
	NarrativeToken   = gra.NarrativeToken
	NewlineToken     = gra.NewlineToken
	NoteToken        = gra.NoteToken
	NumberToken      = gra.NumberToken
	PatternToken     = gra.PatternToken
	PercentToken     = gra.PercentToken
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
	delimiter string,
	message ast.MessageLike,
) AcceptClauseLike {
	return AcceptClauseClass().AcceptClause(
		delimiter,
		message,
	)
}

func AdditionalArgumentClass() AdditionalArgumentClassLike {
	return ast.AdditionalArgumentClass()
}

func AdditionalArgument(
	delimiter string,
	argument ast.ArgumentLike,
) AdditionalArgumentLike {
	return AdditionalArgumentClass().AdditionalArgument(
		delimiter,
		argument,
	)
}

func AdditionalAssociationClass() AdditionalAssociationClassLike {
	return ast.AdditionalAssociationClass()
}

func AdditionalAssociation(
	delimiter string,
	association ast.AssociationLike,
) AdditionalAssociationLike {
	return AdditionalAssociationClass().AdditionalAssociation(
		delimiter,
		association,
	)
}

func AdditionalIndexClass() AdditionalIndexClassLike {
	return ast.AdditionalIndexClass()
}

func AdditionalIndex(
	delimiter string,
	index ast.IndexLike,
) AdditionalIndexLike {
	return AdditionalIndexClass().AdditionalIndex(
		delimiter,
		index,
	)
}

func AdditionalStatementClass() AdditionalStatementClassLike {
	return ast.AdditionalStatementClass()
}

func AdditionalStatement(
	delimiter string,
	statement ast.StatementLike,
) AdditionalStatementLike {
	return AdditionalStatementClass().AdditionalStatement(
		delimiter,
		statement,
	)
}

func AdditionalValueClass() AdditionalValueClassLike {
	return ast.AdditionalValueClass()
}

func AdditionalValue(
	delimiter string,
	component ast.ComponentLike,
) AdditionalValueLike {
	return AdditionalValueClass().AdditionalValue(
		delimiter,
		component,
	)
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

func ArithmeticClass() ArithmeticClassLike {
	return ast.ArithmeticClass()
}

func Arithmetic(
	any_ any,
) ArithmeticLike {
	return ArithmeticClass().Arithmetic(
		any_,
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
	delimiter string,
	component ast.ComponentLike,
) AssociationLike {
	return AssociationClass().Association(
		primitive,
		delimiter,
		component,
	)
}

func AtLevelClass() AtLevelClassLike {
	return ast.AtLevelClass()
}

func AtLevel(
	delimiter1 string,
	delimiter2 string,
	expression ast.ExpressionLike,
) AtLevelLike {
	return AtLevelClass().AtLevel(
		delimiter1,
		delimiter2,
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

func BreakClause(
	delimiter1 string,
	delimiter2 string,
) BreakClauseLike {
	return BreakClauseClass().BreakClause(
		delimiter1,
		delimiter2,
	)
}

func CheckoutClauseClass() CheckoutClauseClassLike {
	return ast.CheckoutClauseClass()
}

func CheckoutClause(
	delimiter1 string,
	recipient ast.RecipientLike,
	optionalAtLevel ast.AtLevelLike,
	delimiter2 string,
	cited ast.CitedLike,
) CheckoutClauseLike {
	return CheckoutClauseClass().CheckoutClause(
		delimiter1,
		recipient,
		optionalAtLevel,
		delimiter2,
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

func ComparisonClass() ComparisonClassLike {
	return ast.ComparisonClass()
}

func Comparison(
	any_ any,
) ComparisonLike {
	return ComparisonClass().Comparison(
		any_,
	)
}

func ComplementClass() ComplementClassLike {
	return ast.ComplementClass()
}

func Complement(
	delimiter string,
	logical ast.LogicalLike,
) ComplementLike {
	return ComplementClass().Complement(
		delimiter,
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

func ContinueClause(
	delimiter1 string,
	delimiter2 string,
) ContinueClauseLike {
	return ContinueClauseClass().ContinueClause(
		delimiter1,
		delimiter2,
	)
}

func DiscardClauseClass() DiscardClauseClassLike {
	return ast.DiscardClauseClass()
}

func DiscardClause(
	delimiter string,
	draft ast.DraftLike,
) DiscardClauseLike {
	return DiscardClauseClass().DiscardClause(
		delimiter,
		draft,
	)
}

func DoClauseClass() DoClauseClassLike {
	return ast.DoClauseClass()
}

func DoClause(
	delimiter string,
	invocation ast.InvocationLike,
) DoClauseLike {
	return DoClauseClass().DoClause(
		delimiter,
		invocation,
	)
}

func DocumentClass() DocumentClassLike {
	return ast.DocumentClass()
}

func Document(
	optionalLegalNotice ast.LegalNoticeLike,
	component ast.ComponentLike,
) DocumentLike {
	return DocumentClass().Document(
		optionalLegalNotice,
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

func Exclusive(
	delimiter string,
) ExclusiveLike {
	return ExclusiveClass().Exclusive(
		delimiter,
	)
}

func ExclusiveRangeClass() ExclusiveRangeClassLike {
	return ast.ExclusiveRangeClass()
}

func ExclusiveRange(
	delimiter1 string,
	primitive1 ast.PrimitiveLike,
	delimiter2 string,
	primitive2 ast.PrimitiveLike,
	bracket ast.BracketLike,
) ExclusiveRangeLike {
	return ExclusiveRangeClass().ExclusiveRange(
		delimiter1,
		primitive1,
		delimiter2,
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
	delimiter1 string,
	optionalArguments ast.ArgumentsLike,
	delimiter2 string,
) FunctionLike {
	return FunctionClass().Function(
		identifier,
		delimiter1,
		optionalArguments,
		delimiter2,
	)
}

func IfClauseClass() IfClauseClassLike {
	return ast.IfClauseClass()
}

func IfClause(
	delimiter1 string,
	condition ast.ConditionLike,
	delimiter2 string,
	procedure ast.ProcedureLike,
) IfClauseLike {
	return IfClauseClass().IfClause(
		delimiter1,
		condition,
		delimiter2,
		procedure,
	)
}

func InclusiveClass() InclusiveClassLike {
	return ast.InclusiveClass()
}

func Inclusive(
	delimiter string,
) InclusiveLike {
	return InclusiveClass().Inclusive(
		delimiter,
	)
}

func InclusiveRangeClass() InclusiveRangeClassLike {
	return ast.InclusiveRangeClass()
}

func InclusiveRange(
	delimiter1 string,
	primitive1 ast.PrimitiveLike,
	delimiter2 string,
	primitive2 ast.PrimitiveLike,
	bracket ast.BracketLike,
) InclusiveRangeLike {
	return InclusiveRangeClass().InclusiveRange(
		delimiter1,
		primitive1,
		delimiter2,
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
	delimiter1 string,
	association ast.AssociationLike,
	additionalAssociations col.Sequential[ast.AdditionalAssociationLike],
	delimiter2 string,
) InlineAttributesLike {
	return InlineAttributesClass().InlineAttributes(
		delimiter1,
		association,
		additionalAssociations,
		delimiter2,
	)
}

func InlineParametersClass() InlineParametersClassLike {
	return ast.InlineParametersClass()
}

func InlineParameters(
	delimiter1 string,
	association ast.AssociationLike,
	additionalAssociations col.Sequential[ast.AdditionalAssociationLike],
	delimiter2 string,
) InlineParametersLike {
	return InlineParametersClass().InlineParameters(
		delimiter1,
		association,
		additionalAssociations,
		delimiter2,
	)
}

func InlineStatementsClass() InlineStatementsClassLike {
	return ast.InlineStatementsClass()
}

func InlineStatements(
	delimiter1 string,
	statement ast.StatementLike,
	additionalStatements col.Sequential[ast.AdditionalStatementLike],
	delimiter2 string,
) InlineStatementsLike {
	return InlineStatementsClass().InlineStatements(
		delimiter1,
		statement,
		additionalStatements,
		delimiter2,
	)
}

func InlineValuesClass() InlineValuesClassLike {
	return ast.InlineValuesClass()
}

func InlineValues(
	delimiter1 string,
	component ast.ComponentLike,
	additionalValues col.Sequential[ast.AdditionalValueLike],
	delimiter2 string,
) InlineValuesLike {
	return InlineValuesClass().InlineValues(
		delimiter1,
		component,
		additionalValues,
		delimiter2,
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

func LegalNoticeClass() LegalNoticeClassLike {
	return ast.LegalNoticeClass()
}

func LegalNotice(
	comment string,
) LegalNoticeLike {
	return LegalNoticeClass().LegalNotice(
		comment,
	)
}

func LetClauseClass() LetClauseClassLike {
	return ast.LetClauseClass()
}

func LetClause(
	delimiter string,
	recipient ast.RecipientLike,
	assignment ast.AssignmentLike,
	expression ast.ExpressionLike,
) LetClauseLike {
	return LetClauseClass().LetClause(
		delimiter,
		recipient,
		assignment,
		expression,
	)
}

func LogicClass() LogicClassLike {
	return ast.LogicClass()
}

func Logic(
	any_ any,
) LogicLike {
	return LogicClass().Logic(
		any_,
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
	delimiter1 string,
	numerical ast.NumericalLike,
	delimiter2 string,
) MagnitudeLike {
	return MagnitudeClass().Magnitude(
		delimiter1,
		numerical,
		delimiter2,
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
	delimiter1 string,
	template ast.TemplateLike,
	delimiter2 string,
	procedure ast.ProcedureLike,
) MatchHandlerLike {
	return MatchHandlerClass().MatchHandler(
		delimiter1,
		template,
		delimiter2,
		procedure,
	)
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
	delimiter1 string,
	optionalArguments ast.ArgumentsLike,
	delimiter2 string,
) MethodLike {
	return MethodClass().Method(
		identifier1,
		blocking,
		identifier2,
		delimiter1,
		optionalArguments,
		delimiter2,
	)
}

func MultilineAttributesClass() MultilineAttributesClassLike {
	return ast.MultilineAttributesClass()
}

func MultilineAttributes(
	delimiter1 string,
	newline string,
	annotatedAssociations col.Sequential[ast.AnnotatedAssociationLike],
	delimiter2 string,
) MultilineAttributesLike {
	return MultilineAttributesClass().MultilineAttributes(
		delimiter1,
		newline,
		annotatedAssociations,
		delimiter2,
	)
}

func MultilineParametersClass() MultilineParametersClassLike {
	return ast.MultilineParametersClass()
}

func MultilineParameters(
	delimiter1 string,
	newline string,
	annotatedAssociations col.Sequential[ast.AnnotatedAssociationLike],
	delimiter2 string,
) MultilineParametersLike {
	return MultilineParametersClass().MultilineParameters(
		delimiter1,
		newline,
		annotatedAssociations,
		delimiter2,
	)
}

func MultilineStatementsClass() MultilineStatementsClassLike {
	return ast.MultilineStatementsClass()
}

func MultilineStatements(
	delimiter1 string,
	newline string,
	annotatedStatements col.Sequential[ast.AnnotatedStatementLike],
	delimiter2 string,
) MultilineStatementsLike {
	return MultilineStatementsClass().MultilineStatements(
		delimiter1,
		newline,
		annotatedStatements,
		delimiter2,
	)
}

func MultilineValuesClass() MultilineValuesClassLike {
	return ast.MultilineValuesClass()
}

func MultilineValues(
	delimiter1 string,
	newline string,
	annotatedValues col.Sequential[ast.AnnotatedValueLike],
	delimiter2 string,
) MultilineValuesLike {
	return MultilineValuesClass().MultilineValues(
		delimiter1,
		newline,
		annotatedValues,
		delimiter2,
	)
}

func NoAttributesClass() NoAttributesClassLike {
	return ast.NoAttributesClass()
}

func NoAttributes(
	delimiter1 string,
	delimiter2 string,
	delimiter3 string,
) NoAttributesLike {
	return NoAttributesClass().NoAttributes(
		delimiter1,
		delimiter2,
		delimiter3,
	)
}

func NoStatementsClass() NoStatementsClassLike {
	return ast.NoStatementsClass()
}

func NoStatements(
	delimiter1 string,
	delimiter2 string,
) NoStatementsLike {
	return NoStatementsClass().NoStatements(
		delimiter1,
		delimiter2,
	)
}

func NoValuesClass() NoValuesClassLike {
	return ast.NoValuesClass()
}

func NoValues(
	delimiter1 string,
	delimiter2 string,
) NoValuesLike {
	return NoValuesClass().NoValues(
		delimiter1,
		delimiter2,
	)
}

func NotarizeClauseClass() NotarizeClauseClassLike {
	return ast.NotarizeClauseClass()
}

func NotarizeClause(
	delimiter1 string,
	draft ast.DraftLike,
	delimiter2 string,
	cited ast.CitedLike,
) NotarizeClauseLike {
	return NotarizeClauseClass().NotarizeClause(
		delimiter1,
		draft,
		delimiter2,
		cited,
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
	delimiter string,
	failure ast.FailureLike,
	matchHandlers col.Sequential[ast.MatchHandlerLike],
) OnClauseLike {
	return OnClauseClass().OnClause(
		delimiter,
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

func PostClauseClass() PostClauseClassLike {
	return ast.PostClauseClass()
}

func PostClause(
	delimiter1 string,
	message ast.MessageLike,
	delimiter2 string,
	bag ast.BagLike,
) PostClauseLike {
	return PostClauseClass().PostClause(
		delimiter1,
		message,
		delimiter2,
		bag,
	)
}

func PrecedenceClass() PrecedenceClassLike {
	return ast.PrecedenceClass()
}

func Precedence(
	delimiter1 string,
	expression ast.ExpressionLike,
	delimiter2 string,
) PrecedenceLike {
	return PrecedenceClass().Precedence(
		delimiter1,
		expression,
		delimiter2,
	)
}

func PredicateClass() PredicateClassLike {
	return ast.PredicateClass()
}

func Predicate(
	operation ast.OperationLike,
	expression ast.ExpressionLike,
) PredicateLike {
	return PredicateClass().Predicate(
		operation,
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
	delimiter string,
	event ast.EventLike,
) PublishClauseLike {
	return PublishClauseClass().PublishClause(
		delimiter,
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
	delimiter string,
	indirect ast.IndirectLike,
) ReferentLike {
	return ReferentClass().Referent(
		delimiter,
		indirect,
	)
}

func RejectClauseClass() RejectClauseClassLike {
	return ast.RejectClauseClass()
}

func RejectClause(
	delimiter string,
	message ast.MessageLike,
) RejectClauseLike {
	return RejectClauseClass().RejectClause(
		delimiter,
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
	delimiter1 string,
	recipient ast.RecipientLike,
	delimiter2 string,
	bag ast.BagLike,
) RetrieveClauseLike {
	return RetrieveClauseClass().RetrieveClause(
		delimiter1,
		recipient,
		delimiter2,
		bag,
	)
}

func ReturnClauseClass() ReturnClauseClassLike {
	return ast.ReturnClauseClass()
}

func ReturnClause(
	delimiter string,
	result ast.ResultLike,
) ReturnClauseLike {
	return ReturnClauseClass().ReturnClause(
		delimiter,
		result,
	)
}

func SaveClauseClass() SaveClauseClassLike {
	return ast.SaveClauseClass()
}

func SaveClause(
	delimiter1 string,
	draft ast.DraftLike,
	delimiter2 string,
	cited ast.CitedLike,
) SaveClauseLike {
	return SaveClauseClass().SaveClause(
		delimiter1,
		draft,
		delimiter2,
		cited,
	)
}

func SelectClauseClass() SelectClauseClassLike {
	return ast.SelectClauseClass()
}

func SelectClause(
	delimiter string,
	target ast.TargetLike,
	matchHandlers col.Sequential[ast.MatchHandlerLike],
) SelectClauseLike {
	return SelectClauseClass().SelectClause(
		delimiter,
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
	delimiter1 string,
	indices ast.IndicesLike,
	delimiter2 string,
) SubcomponentLike {
	return SubcomponentClass().Subcomponent(
		identifier,
		delimiter1,
		indices,
		delimiter2,
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

func TextualClass() TextualClassLike {
	return ast.TextualClass()
}

func Textual(
	any_ any,
) TextualLike {
	return TextualClass().Textual(
		any_,
	)
}

func ThrowClauseClass() ThrowClauseClassLike {
	return ast.ThrowClauseClass()
}

func ThrowClause(
	delimiter string,
	exception ast.ExceptionLike,
) ThrowClauseLike {
	return ThrowClauseClass().ThrowClause(
		delimiter,
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
	delimiter1 string,
	condition ast.ConditionLike,
	delimiter2 string,
	procedure ast.ProcedureLike,
) WhileClauseLike {
	return WhileClauseClass().WhileClause(
		delimiter1,
		condition,
		delimiter2,
		procedure,
	)
}

func WithClauseClass() WithClauseClassLike {
	return ast.WithClauseClass()
}

func WithClause(
	delimiter1 string,
	delimiter2 string,
	item ast.ItemLike,
	delimiter3 string,
	sequence ast.SequenceLike,
	delimiter4 string,
	procedure ast.ProcedureLike,
) WithClauseLike {
	return WithClauseClass().WithClause(
		delimiter1,
		delimiter2,
		item,
		delimiter3,
		sequence,
		delimiter4,
		procedure,
	)
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
