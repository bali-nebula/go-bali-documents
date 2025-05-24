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
	AcceptClauseClassLike   = ast.AcceptClauseClassLike
	AnnotationClassLike     = ast.AnnotationClassLike
	ArgumentClassLike       = ast.ArgumentClassLike
	ArithmeticClassLike     = ast.ArithmeticClassLike
	AssignmentClassLike     = ast.AssignmentClassLike
	AssociationClassLike    = ast.AssociationClassLike
	AtLevelClassLike        = ast.AtLevelClassLike
	AttributesClassLike     = ast.AttributesClassLike
	BagClassLike            = ast.BagClassLike
	BlockingClassLike       = ast.BlockingClassLike
	BreakClauseClassLike    = ast.BreakClauseClassLike
	CheckoutClauseClassLike = ast.CheckoutClauseClassLike
	CitedClassLike          = ast.CitedClassLike
	CollectionClassLike     = ast.CollectionClassLike
	ComparisonClassLike     = ast.ComparisonClassLike
	ComplementClassLike     = ast.ComplementClassLike
	ComponentClassLike      = ast.ComponentClassLike
	ConditionClassLike      = ast.ConditionClassLike
	ContinueClauseClassLike = ast.ContinueClauseClassLike
	DiscardClauseClassLike  = ast.DiscardClauseClassLike
	DoClauseClassLike       = ast.DoClauseClassLike
	DocumentClassLike       = ast.DocumentClassLike
	DraftClassLike          = ast.DraftClassLike
	ElementClassLike        = ast.ElementClassLike
	EmptyClassLike          = ast.EmptyClassLike
	EntityClassLike         = ast.EntityClassLike
	EventClassLike          = ast.EventClassLike
	ExceptionClassLike      = ast.ExceptionClassLike
	ExpressionClassLike     = ast.ExpressionClassLike
	FailureClassLike        = ast.FailureClassLike
	FlowClassLike           = ast.FlowClassLike
	FunctionClassLike       = ast.FunctionClassLike
	HeaderClassLike         = ast.HeaderClassLike
	IfClauseClassLike       = ast.IfClauseClassLike
	IndexClassLike          = ast.IndexClassLike
	IndirectClassLike       = ast.IndirectClassLike
	InductionClassLike      = ast.InductionClassLike
	InverseClassLike        = ast.InverseClassLike
	InversionClassLike      = ast.InversionClassLike
	InvocationClassLike     = ast.InvocationClassLike
	ItemsClassLike          = ast.ItemsClassLike
	LeftBracketClassLike    = ast.LeftBracketClassLike
	LetClauseClassLike      = ast.LetClauseClassLike
	LineClassLike           = ast.LineClassLike
	LogicClassLike          = ast.LogicClassLike
	LogicalClassLike        = ast.LogicalClassLike
	MagnitudeClassLike      = ast.MagnitudeClassLike
	MainClauseClassLike     = ast.MainClauseClassLike
	MatchingClauseClassLike = ast.MatchingClauseClassLike
	MessageClassLike        = ast.MessageClassLike
	MessagingClassLike      = ast.MessagingClassLike
	MethodClassLike         = ast.MethodClassLike
	NotarizeClauseClassLike = ast.NotarizeClauseClassLike
	NumericalClassLike      = ast.NumericalClassLike
	OnClauseClassLike       = ast.OnClauseClassLike
	OperationClassLike      = ast.OperationClassLike
	ParametersClassLike     = ast.ParametersClassLike
	PostClauseClassLike     = ast.PostClauseClassLike
	PrecedenceClassLike     = ast.PrecedenceClassLike
	PredicateClassLike      = ast.PredicateClassLike
	PrimitiveClassLike      = ast.PrimitiveClassLike
	ProcedureClassLike      = ast.ProcedureClassLike
	PublishClauseClassLike  = ast.PublishClauseClassLike
	RangeClassLike          = ast.RangeClassLike
	RecipientClassLike      = ast.RecipientClassLike
	ReferentClassLike       = ast.ReferentClassLike
	RejectClauseClassLike   = ast.RejectClauseClassLike
	RepositoryClassLike     = ast.RepositoryClassLike
	ResultClassLike         = ast.ResultClassLike
	RetrieveClauseClassLike = ast.RetrieveClauseClassLike
	ReturnClauseClassLike   = ast.ReturnClauseClassLike
	RightBracketClassLike   = ast.RightBracketClassLike
	SaveClauseClassLike     = ast.SaveClauseClassLike
	SelectClauseClassLike   = ast.SelectClauseClassLike
	SequenceClassLike       = ast.SequenceClassLike
	StatementClassLike      = ast.StatementClassLike
	StringClassLike         = ast.StringClassLike
	SubcomponentClassLike   = ast.SubcomponentClassLike
	SubjectClassLike        = ast.SubjectClassLike
	TargetClassLike         = ast.TargetClassLike
	TemplateClassLike       = ast.TemplateClassLike
	TextualClassLike        = ast.TextualClassLike
	ThrowClauseClassLike    = ast.ThrowClauseClassLike
	ValueClassLike          = ast.ValueClassLike
	VariableClassLike       = ast.VariableClassLike
	WhileClauseClassLike    = ast.WhileClauseClassLike
	WithClauseClassLike     = ast.WithClauseClassLike
)

type (
	AcceptClauseLike   = ast.AcceptClauseLike
	AnnotationLike     = ast.AnnotationLike
	ArgumentLike       = ast.ArgumentLike
	ArithmeticLike     = ast.ArithmeticLike
	AssignmentLike     = ast.AssignmentLike
	AssociationLike    = ast.AssociationLike
	AtLevelLike        = ast.AtLevelLike
	AttributesLike     = ast.AttributesLike
	BagLike            = ast.BagLike
	BlockingLike       = ast.BlockingLike
	BreakClauseLike    = ast.BreakClauseLike
	CheckoutClauseLike = ast.CheckoutClauseLike
	CitedLike          = ast.CitedLike
	CollectionLike     = ast.CollectionLike
	ComparisonLike     = ast.ComparisonLike
	ComplementLike     = ast.ComplementLike
	ComponentLike      = ast.ComponentLike
	ConditionLike      = ast.ConditionLike
	ContinueClauseLike = ast.ContinueClauseLike
	DiscardClauseLike  = ast.DiscardClauseLike
	DoClauseLike       = ast.DoClauseLike
	DocumentLike       = ast.DocumentLike
	DraftLike          = ast.DraftLike
	ElementLike        = ast.ElementLike
	EmptyLike          = ast.EmptyLike
	EntityLike         = ast.EntityLike
	EventLike          = ast.EventLike
	ExceptionLike      = ast.ExceptionLike
	ExpressionLike     = ast.ExpressionLike
	FailureLike        = ast.FailureLike
	FlowLike           = ast.FlowLike
	FunctionLike       = ast.FunctionLike
	HeaderLike         = ast.HeaderLike
	IfClauseLike       = ast.IfClauseLike
	IndexLike          = ast.IndexLike
	IndirectLike       = ast.IndirectLike
	InductionLike      = ast.InductionLike
	InverseLike        = ast.InverseLike
	InversionLike      = ast.InversionLike
	InvocationLike     = ast.InvocationLike
	ItemsLike          = ast.ItemsLike
	LeftBracketLike    = ast.LeftBracketLike
	LetClauseLike      = ast.LetClauseLike
	LineLike           = ast.LineLike
	LogicLike          = ast.LogicLike
	LogicalLike        = ast.LogicalLike
	MagnitudeLike      = ast.MagnitudeLike
	MainClauseLike     = ast.MainClauseLike
	MatchingClauseLike = ast.MatchingClauseLike
	MessageLike        = ast.MessageLike
	MessagingLike      = ast.MessagingLike
	MethodLike         = ast.MethodLike
	NotarizeClauseLike = ast.NotarizeClauseLike
	NumericalLike      = ast.NumericalLike
	OnClauseLike       = ast.OnClauseLike
	OperationLike      = ast.OperationLike
	ParametersLike     = ast.ParametersLike
	PostClauseLike     = ast.PostClauseLike
	PrecedenceLike     = ast.PrecedenceLike
	PredicateLike      = ast.PredicateLike
	PrimitiveLike      = ast.PrimitiveLike
	ProcedureLike      = ast.ProcedureLike
	PublishClauseLike  = ast.PublishClauseLike
	RangeLike          = ast.RangeLike
	RecipientLike      = ast.RecipientLike
	ReferentLike       = ast.ReferentLike
	RejectClauseLike   = ast.RejectClauseLike
	RepositoryLike     = ast.RepositoryLike
	ResultLike         = ast.ResultLike
	RetrieveClauseLike = ast.RetrieveClauseLike
	ReturnClauseLike   = ast.ReturnClauseLike
	RightBracketLike   = ast.RightBracketLike
	SaveClauseLike     = ast.SaveClauseLike
	SelectClauseLike   = ast.SelectClauseLike
	SequenceLike       = ast.SequenceLike
	StatementLike      = ast.StatementLike
	StringLike         = ast.StringLike
	SubcomponentLike   = ast.SubcomponentLike
	SubjectLike        = ast.SubjectLike
	TargetLike         = ast.TargetLike
	TemplateLike       = ast.TemplateLike
	TextualLike        = ast.TextualLike
	ThrowClauseLike    = ast.ThrowClauseLike
	ValueLike          = ast.ValueLike
	VariableLike       = ast.VariableLike
	WhileClauseLike    = ast.WhileClauseLike
	WithClauseLike     = ast.WithClauseLike
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
	GlyphToken       = gra.GlyphToken
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

func AnnotationClass() AnnotationClassLike {
	return ast.AnnotationClass()
}

func Annotation(
	comment string,
) AnnotationLike {
	return AnnotationClass().Annotation(
		comment,
	)
}

func ArgumentClass() ArgumentClassLike {
	return ast.ArgumentClass()
}

func Argument(
	any_ any,
) ArgumentLike {
	return ArgumentClass().Argument(
		any_,
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

func AttributesClass() AttributesClassLike {
	return ast.AttributesClass()
}

func Attributes(
	delimiter1 string,
	associations col.Sequential[ast.AssociationLike],
	delimiter2 string,
) AttributesLike {
	return AttributesClass().Attributes(
		delimiter1,
		associations,
		delimiter2,
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
	optionalNote string,
) ComponentLike {
	return ComponentClass().Component(
		entity,
		optionalParameters,
		optionalNote,
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
	optionalHeader ast.HeaderLike,
	component ast.ComponentLike,
) DocumentLike {
	return DocumentClass().Document(
		optionalHeader,
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

func EmptyClass() EmptyClassLike {
	return ast.EmptyClass()
}

func Empty(
	delimiter1 string,
	delimiter2 string,
	delimiter3 string,
) EmptyLike {
	return EmptyClass().Empty(
		delimiter1,
		delimiter2,
		delimiter3,
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
	arguments col.Sequential[ast.ArgumentLike],
	delimiter2 string,
) FunctionLike {
	return FunctionClass().Function(
		identifier,
		delimiter1,
		arguments,
		delimiter2,
	)
}

func HeaderClass() HeaderClassLike {
	return ast.HeaderClass()
}

func Header(
	comment string,
) HeaderLike {
	return HeaderClass().Header(
		comment,
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

func IndexClass() IndexClassLike {
	return ast.IndexClass()
}

func Index(
	any_ any,
) IndexLike {
	return IndexClass().Index(
		any_,
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

func ItemsClass() ItemsClassLike {
	return ast.ItemsClass()
}

func Items(
	delimiter1 string,
	components col.Sequential[ast.ComponentLike],
	delimiter2 string,
) ItemsLike {
	return ItemsClass().Items(
		delimiter1,
		components,
		delimiter2,
	)
}

func LeftBracketClass() LeftBracketClassLike {
	return ast.LeftBracketClass()
}

func LeftBracket(
	any_ any,
) LeftBracketLike {
	return LeftBracketClass().LeftBracket(
		any_,
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

func LineClass() LineClassLike {
	return ast.LineClass()
}

func Line(
	any_ any,
) LineLike {
	return LineClass().Line(
		any_,
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

func MatchingClauseClass() MatchingClauseClassLike {
	return ast.MatchingClauseClass()
}

func MatchingClause(
	delimiter1 string,
	template ast.TemplateLike,
	delimiter2 string,
	procedure ast.ProcedureLike,
) MatchingClauseLike {
	return MatchingClauseClass().MatchingClause(
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
	arguments col.Sequential[ast.ArgumentLike],
	delimiter2 string,
) MethodLike {
	return MethodClass().Method(
		identifier1,
		blocking,
		identifier2,
		delimiter1,
		arguments,
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
	matchingClauses col.Sequential[ast.MatchingClauseLike],
) OnClauseLike {
	return OnClauseClass().OnClause(
		delimiter,
		failure,
		matchingClauses,
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
	delimiter1 string,
	associations col.Sequential[ast.AssociationLike],
	delimiter2 string,
) ParametersLike {
	return ParametersClass().Parameters(
		delimiter1,
		associations,
		delimiter2,
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
	delimiter1 string,
	lines col.Sequential[ast.LineLike],
	delimiter2 string,
) ProcedureLike {
	return ProcedureClass().Procedure(
		delimiter1,
		lines,
		delimiter2,
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

func RangeClass() RangeClassLike {
	return ast.RangeClass()
}

func Range(
	leftBracket ast.LeftBracketLike,
	primitive1 ast.PrimitiveLike,
	delimiter string,
	primitive2 ast.PrimitiveLike,
	rightBracket ast.RightBracketLike,
) RangeLike {
	return RangeClass().Range(
		leftBracket,
		primitive1,
		delimiter,
		primitive2,
		rightBracket,
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

func RightBracketClass() RightBracketClassLike {
	return ast.RightBracketClass()
}

func RightBracket(
	any_ any,
) RightBracketLike {
	return RightBracketClass().RightBracket(
		any_,
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
	matchingClauses col.Sequential[ast.MatchingClauseLike],
) SelectClauseLike {
	return SelectClauseClass().SelectClause(
		delimiter,
		target,
		matchingClauses,
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
	optionalNote string,
) StatementLike {
	return StatementClass().Statement(
		mainClause,
		optionalOnClause,
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
	indexes col.Sequential[ast.IndexLike],
	delimiter2 string,
) SubcomponentLike {
	return SubcomponentClass().Subcomponent(
		identifier,
		delimiter1,
		indexes,
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

func ValueClass() ValueClassLike {
	return ast.ValueClass()
}

func Value(
	identifier string,
) ValueLike {
	return ValueClass().Value(
		identifier,
	)
}

func VariableClass() VariableClassLike {
	return ast.VariableClass()
}

func Variable(
	symbol string,
) VariableLike {
	return VariableClass().Variable(
		symbol,
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
	variable ast.VariableLike,
	delimiter3 string,
	sequence ast.SequenceLike,
	delimiter4 string,
	procedure ast.ProcedureLike,
) WithClauseLike {
	return WithClauseClass().WithClause(
		delimiter1,
		delimiter2,
		variable,
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
