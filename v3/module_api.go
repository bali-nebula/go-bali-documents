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
│         This "module_api.go" file was automatically generated using:         │
│            https://github.com/craterdog/go-development-tools/wiki            │
│                                                                              │
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
	age "github.com/bali-nebula/go-bali-documents/v3/agents"
	doc "github.com/bali-nebula/go-bali-documents/v3/documents"
	not "github.com/bali-nebula/go-document-notation/v3"
	fra "github.com/craterdog/go-component-framework/v7"
	uri "net/url"
)

// TYPE ALIASES

// Agents

type (
	DeflatorClassLike  = age.DeflatorClassLike
	InflatorClassLike  = age.InflatorClassLike
	ProcessorClassLike = age.ProcessorClassLike
	ValidatorClassLike = age.ValidatorClassLike
	VisitorClassLike   = age.VisitorClassLike
)

type (
	DeflatorLike  = age.DeflatorLike
	InflatorLike  = age.InflatorLike
	ProcessorLike = age.ProcessorLike
	ValidatorLike = age.ValidatorLike
	VisitorLike   = age.VisitorLike
)

type (
	Methodical = age.Methodical
)

// Documents

type (
	Assignment = doc.Assignment
	Inverse    = doc.Inverse
	Invoke     = doc.Invoke
	Operator   = doc.Operator
)

const (
	Equals         = doc.Equals
	EqualsDefault  = doc.EqualsDefault
	EqualsPlus     = doc.EqualsPlus
	EqualsMinus    = doc.EqualsMinus
	EqualsTimes    = doc.EqualsTimes
	EqualsDivide   = doc.EqualsDivide
	Additive       = doc.Additive
	Multiplicative = doc.Multiplicative
	Conjugate      = doc.Conjugate
	Synchronous    = doc.Synchronous
	Asynchronous   = doc.Asynchronous
	Chain          = doc.Chain
	And            = doc.And
	San            = doc.San
	Ior            = doc.Ior
	Xor            = doc.Xor
	Plus           = doc.Plus
	Minus          = doc.Minus
	Times          = doc.Times
	Divide         = doc.Divide
	Remainder      = doc.Remainder
	Power          = doc.Power
	Less           = doc.Less
	Equal          = doc.Equal
	More           = doc.More
	Is             = doc.Is
	Matches        = doc.Matches
)

type (
	AcceptClauseClassLike     = doc.AcceptClauseClassLike
	AttributesClassLike       = doc.AttributesClassLike
	BreakClauseClassLike      = doc.BreakClauseClassLike
	CheckoutClauseClassLike   = doc.CheckoutClauseClassLike
	ComplementClassLike       = doc.ComplementClassLike
	ComponentClassLike        = doc.ComponentClassLike
	ConstraintClassLike       = doc.ConstraintClassLike
	ContinueClauseClassLike   = doc.ContinueClauseClassLike
	DiscardClauseClassLike    = doc.DiscardClauseClassLike
	DoClauseClassLike         = doc.DoClauseClassLike
	DocumentClassLike         = doc.DocumentClassLike
	ExpressionClassLike       = doc.ExpressionClassLike
	FunctionClassLike         = doc.FunctionClassLike
	IfClauseClassLike         = doc.IfClauseClassLike
	InversionClassLike        = doc.InversionClassLike
	ItemsClassLike            = doc.ItemsClassLike
	LetClauseClassLike        = doc.LetClauseClassLike
	MagnitudeClassLike        = doc.MagnitudeClassLike
	MatchingClauseClassLike   = doc.MatchingClauseClassLike
	MethodClassLike           = doc.MethodClassLike
	NotarizeClauseClassLike   = doc.NotarizeClauseClassLike
	CompositeClassLike        = doc.CompositeClassLike
	OnClauseClassLike         = doc.OnClauseClassLike
	ParameterizationClassLike = doc.ParameterizationClassLike
	PrecedenceClassLike       = doc.PrecedenceClassLike
	PredicateClassLike        = doc.PredicateClassLike
	ProcedureClassLike        = doc.ProcedureClassLike
	PublishClauseClassLike    = doc.PublishClauseClassLike
	RangeClassLike            = doc.RangeClassLike
	ReferentClassLike         = doc.ReferentClassLike
	RejectClauseClassLike     = doc.RejectClauseClassLike
	RetrieveClauseClassLike   = doc.RetrieveClauseClassLike
	ReturnClauseClassLike     = doc.ReturnClauseClassLike
	SaveClauseClassLike       = doc.SaveClauseClassLike
	SelectClauseClassLike     = doc.SelectClauseClassLike
	SendClauseClassLike       = doc.SendClauseClassLike
	StatementClassLike        = doc.StatementClassLike
	SubcomponentClassLike     = doc.SubcomponentClassLike
	ThrowClauseClassLike      = doc.ThrowClauseClassLike
	WhileClauseClassLike      = doc.WhileClauseClassLike
	WithClauseClassLike       = doc.WithClauseClassLike
)

type (
	AcceptClauseLike     = doc.AcceptClauseLike
	AttributesLike       = doc.AttributesLike
	BreakClauseLike      = doc.BreakClauseLike
	CheckoutClauseLike   = doc.CheckoutClauseLike
	ComplementLike       = doc.ComplementLike
	ComponentLike        = doc.ComponentLike
	ConstraintLike       = doc.ConstraintLike
	ContinueClauseLike   = doc.ContinueClauseLike
	DiscardClauseLike    = doc.DiscardClauseLike
	DoClauseLike         = doc.DoClauseLike
	DocumentLike         = doc.DocumentLike
	ExpressionLike       = doc.ExpressionLike
	FunctionLike         = doc.FunctionLike
	IfClauseLike         = doc.IfClauseLike
	InversionLike        = doc.InversionLike
	ItemsLike            = doc.ItemsLike
	LetClauseLike        = doc.LetClauseLike
	MagnitudeLike        = doc.MagnitudeLike
	MatchingClauseLike   = doc.MatchingClauseLike
	MethodLike           = doc.MethodLike
	NotarizeClauseLike   = doc.NotarizeClauseLike
	CompositeLike        = doc.CompositeLike
	OnClauseLike         = doc.OnClauseLike
	ParameterizationLike = doc.ParameterizationLike
	PrecedenceLike       = doc.PrecedenceLike
	PredicateLike        = doc.PredicateLike
	ProcedureLike        = doc.ProcedureLike
	PublishClauseLike    = doc.PublishClauseLike
	RangeLike            = doc.RangeLike
	ReferentLike         = doc.ReferentLike
	RejectClauseLike     = doc.RejectClauseLike
	RetrieveClauseLike   = doc.RetrieveClauseLike
	ReturnClauseLike     = doc.ReturnClauseLike
	SaveClauseLike       = doc.SaveClauseLike
	SelectClauseLike     = doc.SelectClauseLike
	SendClauseLike       = doc.SendClauseLike
	StatementLike        = doc.StatementLike
	SubcomponentLike     = doc.SubcomponentLike
	ThrowClauseLike      = doc.ThrowClauseLike
	WhileClauseLike      = doc.WhileClauseLike
	WithClauseLike       = doc.WithClauseLike
)

// CLASS ACCESSORS

// Agents

func DeflatorClass() DeflatorClassLike {
	return age.DeflatorClass()
}

func Deflator() DeflatorLike {
	return DeflatorClass().Deflator()
}

func InflatorClass() InflatorClassLike {
	return age.InflatorClass()
}

func Inflator() InflatorLike {
	return InflatorClass().Inflator()
}

func ProcessorClass() ProcessorClassLike {
	return age.ProcessorClass()
}

func Processor() ProcessorLike {
	return ProcessorClass().Processor()
}

func ValidatorClass() ValidatorClassLike {
	return age.ValidatorClass()
}

func Validator() ValidatorLike {
	return ValidatorClass().Validator()
}

func VisitorClass() VisitorClassLike {
	return age.VisitorClass()
}

func Visitor(
	processor age.Methodical,
) VisitorLike {
	return VisitorClass().Visitor(
		processor,
	)
}

// Documents

func AcceptClauseClass() AcceptClauseClassLike {
	return doc.AcceptClauseClass()
}

func AcceptClause(
	message doc.ExpressionLike,
) AcceptClauseLike {
	return AcceptClauseClass().AcceptClause(
		message,
	)
}

func AttributesClass() AttributesClassLike {
	return doc.AttributesClass()
}

func Attributes(
	associations fra.CatalogLike[any, doc.CompositeLike],
) AttributesLike {
	return AttributesClass().Attributes(
		associations,
	)
}

func BreakClauseClass() BreakClauseClassLike {
	return doc.BreakClauseClass()
}

func BreakClause() BreakClauseLike {
	return BreakClauseClass().BreakClause()
}

func CheckoutClauseClass() CheckoutClauseClassLike {
	return doc.CheckoutClauseClass()
}

func CheckoutClause(
	recipient any,
	optionalAtLevel doc.ExpressionLike,
	location doc.ExpressionLike,
) CheckoutClauseLike {
	return CheckoutClauseClass().CheckoutClause(
		recipient,
		optionalAtLevel,
		location,
	)
}

func ComplementClass() ComplementClassLike {
	return doc.ComplementClass()
}

func Complement(
	logical any,
) ComplementLike {
	return ComplementClass().Complement(
		logical,
	)
}

func ComponentClass() ComponentClassLike {
	return doc.ComponentClass()
}

func Component(
	entity any,
	optionalParameterization doc.ParameterizationLike,
) ComponentLike {
	return ComponentClass().Component(
		entity,
		optionalParameterization,
	)
}

func ConstraintClass() ConstraintClassLike {
	return doc.ConstraintClass()
}

func Constraint(
	metadata any,
	optionalParameterization doc.ParameterizationLike,
) ConstraintLike {
	return ConstraintClass().Constraint(
		metadata,
		optionalParameterization,
	)
}

func ContinueClauseClass() ContinueClauseClassLike {
	return doc.ContinueClauseClass()
}

func ContinueClause() ContinueClauseLike {
	return ContinueClauseClass().ContinueClause()
}

func DiscardClauseClass() DiscardClauseClassLike {
	return doc.DiscardClauseClass()
}

func DiscardClause(
	location doc.ExpressionLike,
) DiscardClauseLike {
	return DiscardClauseClass().DiscardClause(
		location,
	)
}

func DoClauseClass() DoClauseClassLike {
	return doc.DoClauseClass()
}

func DoClause(
	method doc.MethodLike,
) DoClauseLike {
	return DoClauseClass().DoClause(
		method,
	)
}

func DocumentClass() DocumentClassLike {
	return doc.DocumentClass()
}

func Document(
	annotation string,
	component doc.ComponentLike,
) DocumentLike {
	return DocumentClass().Document(
		annotation,
		component,
	)
}

func ExpressionClass() ExpressionClassLike {
	return doc.ExpressionClass()
}

func Expression(
	subject any,
	predicates fra.ListLike[doc.PredicateLike],
) ExpressionLike {
	return ExpressionClass().Expression(
		subject,
		predicates,
	)
}

func FunctionClass() FunctionClassLike {
	return doc.FunctionClass()
}

func Function(
	identifier string,
	arguments fra.ListLike[any],
) FunctionLike {
	return FunctionClass().Function(
		identifier,
		arguments,
	)
}

func IfClauseClass() IfClauseClassLike {
	return doc.IfClauseClass()
}

func IfClause(
	condition doc.ExpressionLike,
	procedure doc.ProcedureLike,
) IfClauseLike {
	return IfClauseClass().IfClause(
		condition,
		procedure,
	)
}

func InversionClass() InversionClassLike {
	return doc.InversionClass()
}

func Inversion(
	inverse doc.Inverse,
	numerical any,
) InversionLike {
	return InversionClass().Inversion(
		inverse,
		numerical,
	)
}

func ItemsClass() ItemsClassLike {
	return doc.ItemsClass()
}

func Items(
	composites fra.ListLike[doc.CompositeLike],
) ItemsLike {
	return ItemsClass().Items(
		composites,
	)
}

func LetClauseClass() LetClauseClassLike {
	return doc.LetClauseClass()
}

func LetClause(
	recipient any,
	assignment doc.Assignment,
	expression doc.ExpressionLike,
) LetClauseLike {
	return LetClauseClass().LetClause(
		recipient,
		assignment,
		expression,
	)
}

func MagnitudeClass() MagnitudeClassLike {
	return doc.MagnitudeClass()
}

func Magnitude(
	expression doc.ExpressionLike,
) MagnitudeLike {
	return MagnitudeClass().Magnitude(
		expression,
	)
}

func MatchingClauseClass() MatchingClauseClassLike {
	return doc.MatchingClauseClass()
}

func MatchingClause(
	template doc.ExpressionLike,
	procedure doc.ProcedureLike,
) MatchingClauseLike {
	return MatchingClauseClass().MatchingClause(
		template,
		procedure,
	)
}

func MethodClass() MethodClassLike {
	return doc.MethodClass()
}

func Method(
	target string,
	invoke doc.Invoke,
	message string,
	arguments fra.ListLike[any],
) MethodLike {
	return MethodClass().Method(
		target,
		invoke,
		message,
		arguments,
	)
}

func NotarizeClauseClass() NotarizeClauseClassLike {
	return doc.NotarizeClauseClass()
}

func NotarizeClause(
	draft doc.ExpressionLike,
	location doc.ExpressionLike,
) NotarizeClauseLike {
	return NotarizeClauseClass().NotarizeClause(
		draft,
		location,
	)
}

func CompositeClass() CompositeClassLike {
	return doc.CompositeClass()
}

func Composite(
	component doc.ComponentLike,
	optionalNote string,
) CompositeLike {
	return CompositeClass().Composite(
		component,
		optionalNote,
	)
}

func OnClauseClass() OnClauseClassLike {
	return doc.OnClauseClass()
}

func OnClause(
	failure fra.SymbolLike,
	matchingClauses fra.ListLike[doc.MatchingClauseLike],
) OnClauseLike {
	return OnClauseClass().OnClause(
		failure,
		matchingClauses,
	)
}

func ParameterizationClass() ParameterizationClassLike {
	return doc.ParameterizationClass()
}

func Parameterization(
	parameters fra.CatalogLike[fra.SymbolLike, doc.ConstraintLike],
) ParameterizationLike {
	return ParameterizationClass().Parameterization(
		parameters,
	)
}

func PrecedenceClass() PrecedenceClassLike {
	return doc.PrecedenceClass()
}

func Precedence(
	expression doc.ExpressionLike,
) PrecedenceLike {
	return PrecedenceClass().Precedence(
		expression,
	)
}

func PredicateClass() PredicateClassLike {
	return doc.PredicateClass()
}

func Predicate(
	operator doc.Operator,
	expression doc.ExpressionLike,
) PredicateLike {
	return PredicateClass().Predicate(
		operator,
		expression,
	)
}

func ProcedureClass() ProcedureClassLike {
	return doc.ProcedureClass()
}

func Procedure(
	lines fra.ListLike[any],
) ProcedureLike {
	return ProcedureClass().Procedure(
		lines,
	)
}

func PublishClauseClass() PublishClauseClassLike {
	return doc.PublishClauseClass()
}

func PublishClause(
	event doc.ExpressionLike,
) PublishClauseLike {
	return PublishClauseClass().PublishClause(
		event,
	)
}

func RangeClass() RangeClassLike {
	return doc.RangeClass()
}

func Range(
	left fra.Bracket,
	first any,
	last any,
	right fra.Bracket,
) RangeLike {
	return RangeClass().Range(
		left,
		first,
		last,
		right,
	)
}

func ReferentClass() ReferentClassLike {
	return doc.ReferentClass()
}

func Referent(
	reference any,
) ReferentLike {
	return ReferentClass().Referent(
		reference,
	)
}

func RejectClauseClass() RejectClauseClassLike {
	return doc.RejectClauseClass()
}

func RejectClause(
	message doc.ExpressionLike,
) RejectClauseLike {
	return RejectClauseClass().RejectClause(
		message,
	)
}

func RetrieveClauseClass() RetrieveClauseClassLike {
	return doc.RetrieveClauseClass()
}

func RetrieveClause(
	recipient any,
	bag doc.ExpressionLike,
) RetrieveClauseLike {
	return RetrieveClauseClass().RetrieveClause(
		recipient,
		bag,
	)
}

func ReturnClauseClass() ReturnClauseClassLike {
	return doc.ReturnClauseClass()
}

func ReturnClause(
	result doc.ExpressionLike,
) ReturnClauseLike {
	return ReturnClauseClass().ReturnClause(
		result,
	)
}

func SaveClauseClass() SaveClauseClassLike {
	return doc.SaveClauseClass()
}

func SaveClause(
	draft doc.ExpressionLike,
	location doc.ExpressionLike,
) SaveClauseLike {
	return SaveClauseClass().SaveClause(
		draft,
		location,
	)
}

func SelectClauseClass() SelectClauseClassLike {
	return doc.SelectClauseClass()
}

func SelectClause(
	expression doc.ExpressionLike,
	matchingClauses fra.ListLike[doc.MatchingClauseLike],
) SelectClauseLike {
	return SelectClauseClass().SelectClause(
		expression,
		matchingClauses,
	)
}

func SendClauseClass() SendClauseClassLike {
	return doc.SendClauseClass()
}

func SendClause(
	message doc.ExpressionLike,
	bag doc.ExpressionLike,
) SendClauseLike {
	return SendClauseClass().SendClause(
		message,
		bag,
	)
}

func StatementClass() StatementClassLike {
	return doc.StatementClass()
}

func Statement(
	mainClause any,
	optionalOnClause doc.OnClauseLike,
) StatementLike {
	return StatementClass().Statement(
		mainClause,
		optionalOnClause,
	)
}

func SubcomponentClass() SubcomponentClassLike {
	return doc.SubcomponentClass()
}

func Subcomponent(
	identifier string,
	indexes fra.ListLike[any],
) SubcomponentLike {
	return SubcomponentClass().Subcomponent(
		identifier,
		indexes,
	)
}

func ThrowClauseClass() ThrowClauseClassLike {
	return doc.ThrowClauseClass()
}

func ThrowClause(
	exception doc.ExpressionLike,
) ThrowClauseLike {
	return ThrowClauseClass().ThrowClause(
		exception,
	)
}

func WhileClauseClass() WhileClauseClassLike {
	return doc.WhileClauseClass()
}

func WhileClause(
	condition doc.ExpressionLike,
	procedure doc.ProcedureLike,
) WhileClauseLike {
	return WhileClauseClass().WhileClause(
		condition,
		procedure,
	)
}

func WithClauseClass() WithClauseClassLike {
	return doc.WithClauseClass()
}

func WithClause(
	variable fra.SymbolLike,
	sequence doc.ExpressionLike,
	procedure doc.ProcedureLike,
) WithClauseLike {
	return WithClauseClass().WithClause(
		variable,
		sequence,
		procedure,
	)
}

// GLOBAL FUNCTIONS

// Components

func ParseComponent(
	source string,
) ComponentLike {
	var document = ParseDocument(source)
	return document.GetComponent()
}

func ParseDocument(
	source string,
) DocumentLike {
	var inflator = Inflator()
	var parser = not.Parser()
	var document = inflator.InflateDocument(parser.ParseSource(source))
	return document
}

func FormatComponent(
	value any,
) string {
	var component ComponentLike
	switch actual := value.(type) {
	case ComponentLike:
		component = actual
	case ConstraintLike:
		var entity = actual.GetMetadata()
		var parameterization = actual.GetOptionalParameterization()
		component = Component(entity, parameterization)
	case DocumentLike:
		component = actual.GetComponent()
	case CompositeLike:
		component = actual.GetComponent()
	default:
		component = Component(value, nil)
	}
	var source = FormatDocument(Document("", component))
	return source[:len(source)-1] // Remove the trailing newline.
}

func FormatDocument(
	document DocumentLike,
) string {
	var deflator = Deflator()
	var formatter = not.Formatter()
	var source = formatter.FormatDocument(deflator.DeflateDocument(document))
	return source
}

// Agents

type (
	Event       = fra.Event
	Rank        = fra.Rank
	State       = fra.State
	Transitions = fra.Transitions
)

const (
	LesserRank  = fra.LesserRank
	EqualRank   = fra.EqualRank
	GreaterRank = fra.GreaterRank
)

type (
	RankingFunction[V any] = fra.RankingFunction[V]
)

type (
	CollatorClassLike[V any] = fra.CollatorClassLike[V]
	ControllerClassLike      = fra.ControllerClassLike
	EncoderClassLike         = fra.EncoderClassLike
	GeneratorClassLike       = fra.GeneratorClassLike
	IteratorClassLike[V any] = fra.IteratorClassLike[V]
	SorterClassLike[V any]   = fra.SorterClassLike[V]
)

type (
	CollatorLike[V any] = fra.CollatorLike[V]
	ControllerLike      = fra.ControllerLike
	EncoderLike         = fra.EncoderLike
	GeneratorLike       = fra.GeneratorLike
	IteratorLike[V any] = fra.IteratorLike[V]
	SorterLike[V any]   = fra.SorterLike[V]
)

func CollatorClass[V any]() CollatorClassLike[V] {
	return fra.CollatorClass[V]()
}

func Collator[V any](
	value ...any,
) CollatorLike[V] {
	if len(value) == 0 {
		return CollatorClass[V]().Collator()
	}
	switch actual := value[0].(type) {
	case int:
		return CollatorClass[V]().CollatorWithMaximumDepth(uint(actual))
	case uint:
		return CollatorClass[V]().CollatorWithMaximumDepth(actual)
	default:
		return CollatorClass[V]().Collator()
	}
}

func ControllerClass() ControllerClassLike {
	return fra.ControllerClass()
}

func Controller(
	events []fra.Event,
	transitions map[State]fra.Transitions,
	initialState fra.State,
) ControllerLike {
	return ControllerClass().Controller(
		events,
		transitions,
		initialState,
	)
}

func EncoderClass() EncoderClassLike {
	return fra.EncoderClass()
}

func Encoder() EncoderLike {
	return EncoderClass().Encoder()
}

func GeneratorClass() GeneratorClassLike {
	return fra.GeneratorClass()
}

func Generator() GeneratorLike {
	return GeneratorClass().Generator()
}

func IteratorClass[V any]() IteratorClassLike[V] {
	return fra.IteratorClass[V]()
}

func Iterator[V any](
	array []V,
) IteratorLike[V] {
	return IteratorClass[V]().Iterator(
		array,
	)
}

func SorterClass[V any]() SorterClassLike[V] {
	return fra.SorterClass[V]()
}

func Sorter[V any](
	value ...any,
) SorterLike[V] {
	if len(value) == 0 {
		return SorterClass[V]().Sorter()
	}
	switch actual := value[0].(type) {
	case fra.RankingFunction[V]:
		return SorterClass[V]().SorterWithRanker(actual)
	default:
		return SorterClass[V]().Sorter()
	}
}

// Elements

type (
	Units = fra.Units
)

const (
	Degrees  = fra.Degrees
	Radians  = fra.Radians
	Gradians = fra.Gradians
)

type (
	AngleClassLike       = fra.AngleClassLike
	BooleanClassLike     = fra.BooleanClassLike
	DurationClassLike    = fra.DurationClassLike
	GlyphClassLike       = fra.GlyphClassLike
	MomentClassLike      = fra.MomentClassLike
	NumberClassLike      = fra.NumberClassLike
	PercentageClassLike  = fra.PercentageClassLike
	ProbabilityClassLike = fra.ProbabilityClassLike
	ResourceClassLike    = fra.ResourceClassLike
	SymbolClassLike      = fra.SymbolClassLike
)

type (
	AngleLike       = fra.AngleLike
	BooleanLike     = fra.BooleanLike
	DurationLike    = fra.DurationLike
	GlyphLike       = fra.GlyphLike
	MomentLike      = fra.MomentLike
	NumberLike      = fra.NumberLike
	PercentageLike  = fra.PercentageLike
	ProbabilityLike = fra.ProbabilityLike
	ResourceLike    = fra.ResourceLike
	SymbolLike      = fra.SymbolLike
)

type (
	Continuous = fra.Continuous
	Discrete   = fra.Discrete
	Factored   = fra.Factored
	Polarized  = fra.Polarized
	Temporal   = fra.Temporal
)

func AngleClass() AngleClassLike {
	return fra.AngleClass()
}

func Angle(
	value ...any,
) AngleLike {
	if len(value) == 0 {
		return AngleClass().Zero()
	}
	switch actual := value[0].(type) {
	case string:
		return AngleClass().AngleFromSource(actual)
	case float64:
		return AngleClass().Angle(actual)
	default:
		return AngleClass().Zero()
	}
}

func BooleanClass() BooleanClassLike {
	return fra.BooleanClass()
}

func Boolean(
	value ...any,
) BooleanLike {
	if len(value) == 0 {
		return BooleanClass().False()
	}
	switch actual := value[0].(type) {
	case string:
		return BooleanClass().BooleanFromSource(actual)
	case bool:
		return BooleanClass().Boolean(actual)
	default:
		return BooleanClass().False()
	}
}

func DurationClass() DurationClassLike {
	return fra.DurationClass()
}

func Duration(
	value ...any,
) DurationLike {
	if len(value) == 0 {
		return DurationClass().Duration(0)
	}
	switch actual := value[0].(type) {
	case string:
		return DurationClass().DurationFromSource(actual)
	case int:
		return DurationClass().Duration(actual)
	default:
		return DurationClass().Duration(0)
	}
}

func GlyphClass() GlyphClassLike {
	return fra.GlyphClass()
}

func Glyph(
	value ...any,
) GlyphLike {
	if len(value) == 0 {
		return GlyphClass().Undefined()
	}
	switch actual := value[0].(type) {
	case string:
		return GlyphClass().GlyphFromSource(actual)
	case rune:
		return GlyphClass().Glyph(actual)
	case int:
		return GlyphClass().GlyphFromInteger(actual)
	default:
		return GlyphClass().Undefined()
	}
}

func MomentClass() MomentClassLike {
	return fra.MomentClass()
}

func Moment(
	value ...any,
) MomentLike {
	if len(value) == 0 {
		return MomentClass().Now()
	}
	switch actual := value[0].(type) {
	case string:
		return MomentClass().MomentFromSource(actual)
	case int:
		return MomentClass().Moment(actual)
	default:
		return MomentClass().Now()
	}
}

func NumberClass() NumberClassLike {
	return fra.NumberClass()
}

func Number(
	value ...any,
) NumberLike {
	if len(value) == 0 {
		return NumberClass().Undefined()
	}
	switch actual := value[0].(type) {
	case string:
		return NumberClass().NumberFromSource(actual)
	case complex128:
		return NumberClass().Number(actual)
	case int:
		return NumberClass().Number(complex(float64(actual), 0))
	case float64:
		return NumberClass().Number(complex(actual, 0))
	default:
		return NumberClass().Undefined()
	}
}

func Polar(
	magnitude float64,
	phase float64,
) NumberLike {
	return NumberClass().NumberFromPolar(magnitude, phase)
}

func Rectangular(
	x float64,
	y float64,
) NumberLike {
	return NumberClass().NumberFromRectangular(x, y)
}

func PercentageClass() PercentageClassLike {
	return fra.PercentageClass()
}

func Percentage(
	value ...any,
) PercentageLike {
	if len(value) == 0 {
		return PercentageClass().Undefined()
	}
	switch actual := value[0].(type) {
	case string:
		return PercentageClass().PercentageFromSource(actual)
	case int:
		return PercentageClass().PercentageFromInteger(actual)
	case float64:
		return PercentageClass().Percentage(actual)
	default:
		return PercentageClass().Undefined()
	}
}

func ProbabilityClass() ProbabilityClassLike {
	return fra.ProbabilityClass()
}

func Probability(
	value ...any,
) ProbabilityLike {
	if len(value) == 0 {
		return ProbabilityClass().Random()
	}
	switch actual := value[0].(type) {
	case string:
		return ProbabilityClass().ProbabilityFromSource(actual)
	case bool:
		return ProbabilityClass().ProbabilityFromBoolean(actual)
	case float64:
		return ProbabilityClass().Probability(actual)
	default:
		return ProbabilityClass().Random()
	}
}

func ResourceClass() ResourceClassLike {
	return fra.ResourceClass()
}

func Resource(
	value ...any,
) ResourceLike {
	if len(value) == 0 {
		return ResourceClass().Undefined()
	}
	switch actual := value[0].(type) {
	case string:
		if actual[0] == '<' {
			return ResourceClass().ResourceFromSource(actual)
		} else {
			return ResourceClass().Resource(actual)
		}
	case *uri.URL:
		return ResourceClass().ResourceFromUri(actual)
	default:
		return ResourceClass().Undefined()
	}
}

func SymbolClass() SymbolClassLike {
	return fra.SymbolClass()
}

func Symbol(
	value ...any,
) SymbolLike {
	if len(value) == 0 {
		return SymbolClass().Undefined()
	}
	switch actual := value[0].(type) {
	case string:
		if actual[0] == '$' {
			return SymbolClass().SymbolFromSource(actual)
		} else {
			return SymbolClass().Symbol(actual)
		}
	default:
		return SymbolClass().Undefined()
	}
}

// Strings

type (
	Folder = fra.Folder
)

type (
	BinaryClassLike    = fra.BinaryClassLike
	BytecodeClassLike  = fra.BytecodeClassLike
	NameClassLike      = fra.NameClassLike
	NarrativeClassLike = fra.NarrativeClassLike
	PatternClassLike   = fra.PatternClassLike
	QuoteClassLike     = fra.QuoteClassLike
	TagClassLike       = fra.TagClassLike
	VersionClassLike   = fra.VersionClassLike
)

type (
	BinaryLike    = fra.BinaryLike
	BytecodeLike  = fra.BytecodeLike
	NameLike      = fra.NameLike
	NarrativeLike = fra.NarrativeLike
	PatternLike   = fra.PatternLike
	QuoteLike     = fra.QuoteLike
	TagLike       = fra.TagLike
	VersionLike   = fra.VersionLike
)

type (
	Accessible[V any] = fra.Accessible[V]
	Searchable[V any] = fra.Searchable[V]
	Sequential[V any] = fra.Sequential[V]
	Spectral[V any]   = fra.Spectral[V]
)

func BinaryClass() BinaryClassLike {
	return fra.BinaryClass()
}

func Binary(
	value ...any,
) BinaryLike {
	if len(value) == 0 {
		return BinaryClass().Binary([]byte{})
	}
	switch actual := value[0].(type) {
	case string:
		return BinaryClass().BinaryFromSource(actual)
	case []byte:
		return BinaryClass().Binary(actual)
	case fra.Sequential[byte]:
		return BinaryClass().BinaryFromSequence(actual)
	default:
		return BinaryClass().Binary([]byte{})
	}
}

func BytecodeClass() BytecodeClassLike {
	return fra.BytecodeClass()
}

func Bytecode(
	value ...any,
) BytecodeLike {
	if len(value) == 0 {
		return BytecodeClass().Bytecode([]uint16{0})
	}
	switch actual := value[0].(type) {
	case string:
		return BytecodeClass().BytecodeFromSource(actual)
	case []uint16:
		return BytecodeClass().Bytecode(actual)
	default:
		return BytecodeClass().Bytecode([]uint16{0})
	}
}

func NameClass() NameClassLike {
	return fra.NameClass()
}

func Name(
	value ...any,
) NameLike {
	if len(value) == 0 {
		return NameClass().Name([]fra.Folder{})
	}
	switch actual := value[0].(type) {
	case string:
		return NameClass().NameFromSource(actual)
	case []fra.Folder:
		return NameClass().Name(actual)
	case fra.Sequential[fra.Folder]:
		return NameClass().NameFromSequence(actual)
	default:
		return NameClass().Name([]fra.Folder{})
	}
}

func NarrativeClass() NarrativeClassLike {
	return fra.NarrativeClass()
}

func Narrative(
	value ...any,
) NarrativeLike {
	if len(value) == 0 {
		return NarrativeClass().Narrative([]string{})
	}
	switch actual := value[0].(type) {
	case string:
		return NarrativeClass().NarrativeFromSource(actual)
	case []string:
		return NarrativeClass().Narrative(actual)
	case fra.Sequential[string]:
		return NarrativeClass().NarrativeFromSequence(actual)
	default:
		return NarrativeClass().Narrative([]string{})
	}
}

func PatternClass() PatternClassLike {
	return fra.PatternClass()
}

func Pattern(
	value ...any,
) PatternLike {
	if len(value) == 0 {
		return PatternClass().None()
	}
	switch actual := value[0].(type) {
	case string:
		return PatternClass().PatternFromSource(actual)
	case []rune:
		return PatternClass().Pattern(actual)
	case fra.Sequential[rune]:
		return PatternClass().PatternFromSequence(actual)
	default:
		return PatternClass().None()
	}
}

func QuoteClass() QuoteClassLike {
	return fra.QuoteClass()
}

func Quote(
	value ...any,
) QuoteLike {
	if len(value) == 0 {
		return QuoteClass().QuoteFromSource(`""`)
	}
	switch actual := value[0].(type) {
	case string:
		if actual[0] == '"' {
			return QuoteClass().QuoteFromSource(actual)
		} else {
			return QuoteClass().QuoteFromSource(`"` + actual + `"`)
		}
	case []rune:
		return QuoteClass().Quote(actual)
	case fra.Sequential[rune]:
		return QuoteClass().QuoteFromSequence(actual)
	default:
		return QuoteClass().QuoteFromSource(`""`)
	}
}

func TagClass() TagClassLike {
	return fra.TagClass()
}

func Tag(
	value ...any,
) TagLike {
	if len(value) == 0 {
		return TagClass().TagWithSize(20)
	}
	switch actual := value[0].(type) {
	case string:
		return TagClass().TagFromSource(actual)
	case []byte:
		return TagClass().Tag(actual)
	case fra.Sequential[byte]:
		return TagClass().TagFromSequence(actual)
	case int:
		return TagClass().TagWithSize(uint(actual))
	case uint:
		return TagClass().TagWithSize(actual)
	default:
		return TagClass().TagWithSize(20)
	}
}

func VersionClass() VersionClassLike {
	return fra.VersionClass()
}

func Version(
	value ...any,
) VersionLike {
	if len(value) == 0 {
		return VersionClass().Version([]uint{1})
	}
	switch actual := value[0].(type) {
	case string:
		return VersionClass().VersionFromSource(actual)
	case []uint:
		return VersionClass().Version(actual)
	case fra.Sequential[uint]:
		return VersionClass().VersionFromSequence(actual)
	default:
		return VersionClass().Version([]uint{})
	}
}

// Ranges

type (
	Bracket = fra.Bracket
)

const (
	Inclusive = fra.Inclusive
	Exclusive = fra.Exclusive
)

type (
	ContinuumClassLike[V fra.Continuous] = fra.ContinuumClassLike[V]
	IntervalClassLike[V fra.Discrete]    = fra.IntervalClassLike[V]
	SpectrumClassLike[V fra.Spectral[V]] = fra.SpectrumClassLike[V]
)

type (
	ContinuumLike[V fra.Continuous] = fra.ContinuumLike[V]
	IntervalLike[V fra.Discrete]    = fra.IntervalLike[V]
	SpectrumLike[V fra.Spectral[V]] = fra.SpectrumLike[V]
)

type (
	Bounded[V any] = fra.Bounded[V]
)

func ContinuumClass[V fra.Continuous]() ContinuumClassLike[V] {
	return fra.ContinuumClass[V]()
}

func Continuum[V fra.Continuous](
	left fra.Bracket,
	minimum V,
	maximum V,
	right fra.Bracket,
) ContinuumLike[V] {
	return ContinuumClass[V]().Continuum(
		left,
		minimum,
		maximum,
		right,
	)
}

func IntervalClass[V fra.Discrete]() IntervalClassLike[V] {
	return fra.IntervalClass[V]()
}

func Interval[V fra.Discrete](
	left fra.Bracket,
	minimum V,
	maximum V,
	right fra.Bracket,
) IntervalLike[V] {
	return IntervalClass[V]().Interval(
		left,
		minimum,
		maximum,
		right,
	)
}

func SpectrumClass[V fra.Spectral[V]]() SpectrumClassLike[V] {
	return fra.SpectrumClass[V]()
}

func Spectrum[V fra.Spectral[V]](
	left fra.Bracket,
	minimum V,
	maximum V,
	right fra.Bracket,
) SpectrumLike[V] {
	return SpectrumClass[V]().Spectrum(
		left,
		minimum,
		maximum,
		right,
	)
}

// Collections

type (
	AssociationClassLike[K comparable, V any] = fra.AssociationClassLike[K, V]
	CatalogClassLike[K comparable, V any]     = fra.CatalogClassLike[K, V]
	ListClassLike[V any]                      = fra.ListClassLike[V]
	QueueClassLike[V any]                     = fra.QueueClassLike[V]
	SetClassLike[V any]                       = fra.SetClassLike[V]
	StackClassLike[V any]                     = fra.StackClassLike[V]
)

type (
	AssociationLike[K comparable, V any] = fra.AssociationLike[K, V]
	CatalogLike[K comparable, V any]     = fra.CatalogLike[K, V]
	ListLike[V any]                      = fra.ListLike[V]
	QueueLike[V any]                     = fra.QueueLike[V]
	SetLike[V any]                       = fra.SetLike[V]
	StackLike[V any]                     = fra.StackLike[V]
)

type (
	Associative[K comparable, V any] = fra.Associative[K, V]
	Elastic[V any]                   = fra.Elastic[V]
	Fifo[V any]                      = fra.Fifo[V]
	Lifo[V any]                      = fra.Lifo[V]
	Malleable[V any]                 = fra.Malleable[V]
	Sortable[V any]                  = fra.Sortable[V]
	Synchronized                     = fra.Synchronized
	Updatable[V any]                 = fra.Updatable[V]
)

func AssociationClass[K comparable, V any]() AssociationClassLike[K, V] {
	return fra.AssociationClass[K, V]()
}

func Association[K comparable, V any](
	key K,
	value V,
) AssociationLike[K, V] {
	return AssociationClass[K, V]().Association(
		key,
		value,
	)
}

func CatalogClass[K comparable, V any]() CatalogClassLike[K, V] {
	return fra.CatalogClass[K, V]()
}

func Catalog[K comparable, V any](
	value ...any,
) CatalogLike[K, V] {
	if len(value) == 0 {
		return fra.Catalog[K, V]()
	}
	switch actual := value[0].(type) {
	case string:
		return ParseComponent(actual).GetEntity().(CatalogLike[K, V])
	case []AssociationLike[K, V]:
		return fra.CatalogFromArray(actual)
	case map[K]V:
		return fra.CatalogFromMap(actual)
	case Sequential[AssociationLike[K, V]]:
		return fra.CatalogFromSequence(actual)
	default:
		return fra.Catalog[K, V]()
	}
}

func ListClass[V any]() ListClassLike[V] {
	return fra.ListClass[V]()
}

func List[V any](
	value ...any,
) ListLike[V] {
	if len(value) == 0 {
		return ListClass[V]().List()
	}
	switch actual := value[0].(type) {
	case string:
		return ParseComponent(actual).GetEntity().(ListLike[V])
	case []V:
		return ListClass[V]().ListFromArray(actual)
	case fra.Sequential[V]:
		return ListClass[V]().ListFromSequence(actual)
	default:
		return ListClass[V]().List()
	}
}

func QueueClass[V any]() QueueClassLike[V] {
	return fra.QueueClass[V]()
}

func Queue[V any](
	value ...any,
) QueueLike[V] {
	if len(value) == 0 {
		return QueueClass[V]().Queue()
	}
	switch actual := value[0].(type) {
	case string:
		return ParseComponent(actual).GetEntity().(QueueLike[V])
	case int:
		return QueueClass[V]().QueueWithCapacity(uint(actual))
	case uint:
		return QueueClass[V]().QueueWithCapacity(actual)
	case []V:
		return QueueClass[V]().QueueFromArray(actual)
	case fra.Sequential[V]:
		return QueueClass[V]().QueueFromSequence(actual)
	default:
		return QueueClass[V]().Queue()
	}
}

func SetClass[V any]() SetClassLike[V] {
	return fra.SetClass[V]()
}

func Set[V any](
	value ...any,
) SetLike[V] {
	if len(value) == 0 {
		return SetClass[V]().Set()
	}
	switch actual := value[0].(type) {
	case string:
		return ParseComponent(actual).GetEntity().(SetLike[V])
	case fra.CollatorLike[V]:
		return SetClass[V]().SetWithCollator(actual)
	case []V:
		return SetClass[V]().SetFromArray(actual)
	case fra.Sequential[V]:
		return SetClass[V]().SetFromSequence(actual)
	default:
		return SetClass[V]().Set()
	}
}

func StackClass[V any]() StackClassLike[V] {
	return fra.StackClass[V]()
}

func Stack[V any](
	value ...any,
) StackLike[V] {
	if len(value) == 0 {
		return StackClass[V]().Stack()
	}
	switch actual := value[0].(type) {
	case string:
		return ParseComponent(actual).GetEntity().(StackLike[V])
	case int:
		return StackClass[V]().StackWithCapacity(uint(actual))
	case uint:
		return StackClass[V]().StackWithCapacity(actual)
	case []V:
		return StackClass[V]().StackFromArray(actual)
	case fra.Sequential[V]:
		return StackClass[V]().StackFromSequence(actual)
	default:
		return StackClass[V]().Stack()
	}
}
