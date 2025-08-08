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
	fra "github.com/craterdog/go-component-framework/v7"
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
	Extent = doc.Extent
)

const (
	Inclusive = doc.Inclusive
	Exclusive = doc.Exclusive
)

type (
	AcceptClauseClassLike   = doc.AcceptClauseClassLike
	AttributesClassLike     = doc.AttributesClassLike
	BreakClauseClassLike    = doc.BreakClauseClassLike
	CheckoutClauseClassLike = doc.CheckoutClauseClassLike
	ComplementClassLike     = doc.ComplementClassLike
	ContinueClauseClassLike = doc.ContinueClauseClassLike
	DiscardClauseClassLike  = doc.DiscardClauseClassLike
	DoClauseClassLike       = doc.DoClauseClassLike
	DocumentClassLike       = doc.DocumentClassLike
	ExpressionClassLike     = doc.ExpressionClassLike
	FunctionClassLike       = doc.FunctionClassLike
	IfClauseClassLike       = doc.IfClauseClassLike
	InversionClassLike      = doc.InversionClassLike
	ItemsClassLike          = doc.ItemsClassLike
	LetClauseClassLike      = doc.LetClauseClassLike
	MagnitudeClassLike      = doc.MagnitudeClassLike
	MatchingClauseClassLike = doc.MatchingClauseClassLike
	MethodClassLike         = doc.MethodClassLike
	NotarizeClauseClassLike = doc.NotarizeClauseClassLike
	OnClauseClassLike       = doc.OnClauseClassLike
	ParametersClassLike     = doc.ParametersClassLike
	PostClauseClassLike     = doc.PostClauseClassLike
	PrecedenceClassLike     = doc.PrecedenceClassLike
	PredicateClassLike      = doc.PredicateClassLike
	ProcedureClassLike      = doc.ProcedureClassLike
	PublishClauseClassLike  = doc.PublishClauseClassLike
	RangeClassLike          = doc.RangeClassLike
	ReferentClassLike       = doc.ReferentClassLike
	RejectClauseClassLike   = doc.RejectClauseClassLike
	RetrieveClauseClassLike = doc.RetrieveClauseClassLike
	ReturnClauseClassLike   = doc.ReturnClauseClassLike
	SaveClauseClassLike     = doc.SaveClauseClassLike
	SelectClauseClassLike   = doc.SelectClauseClassLike
	StatementClassLike      = doc.StatementClassLike
	SubcomponentClassLike   = doc.SubcomponentClassLike
	ThrowClauseClassLike    = doc.ThrowClauseClassLike
	WhileClauseClassLike    = doc.WhileClauseClassLike
	WithClauseClassLike     = doc.WithClauseClassLike
)

type (
	AcceptClauseLike   = doc.AcceptClauseLike
	AttributesLike     = doc.AttributesLike
	BreakClauseLike    = doc.BreakClauseLike
	CheckoutClauseLike = doc.CheckoutClauseLike
	ComplementLike     = doc.ComplementLike
	ContinueClauseLike = doc.ContinueClauseLike
	DiscardClauseLike  = doc.DiscardClauseLike
	DoClauseLike       = doc.DoClauseLike
	DocumentLike       = doc.DocumentLike
	ExpressionLike     = doc.ExpressionLike
	FunctionLike       = doc.FunctionLike
	IfClauseLike       = doc.IfClauseLike
	InversionLike      = doc.InversionLike
	ItemsLike          = doc.ItemsLike
	LetClauseLike      = doc.LetClauseLike
	MagnitudeLike      = doc.MagnitudeLike
	MatchingClauseLike = doc.MatchingClauseLike
	MethodLike         = doc.MethodLike
	NotarizeClauseLike = doc.NotarizeClauseLike
	OnClauseLike       = doc.OnClauseLike
	ParametersLike     = doc.ParametersLike
	PostClauseLike     = doc.PostClauseLike
	PrecedenceLike     = doc.PrecedenceLike
	PredicateLike      = doc.PredicateLike
	ProcedureLike      = doc.ProcedureLike
	PublishClauseLike  = doc.PublishClauseLike
	RangeLike          = doc.RangeLike
	ReferentLike       = doc.ReferentLike
	RejectClauseLike   = doc.RejectClauseLike
	RetrieveClauseLike = doc.RetrieveClauseLike
	ReturnClauseLike   = doc.ReturnClauseLike
	SaveClauseLike     = doc.SaveClauseLike
	SelectClauseLike   = doc.SelectClauseLike
	StatementLike      = doc.StatementLike
	SubcomponentLike   = doc.SubcomponentLike
	ThrowClauseLike    = doc.ThrowClauseLike
	WhileClauseLike    = doc.WhileClauseLike
	WithClauseLike     = doc.WithClauseLike
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
	associations fra.CatalogLike[any, doc.DocumentLike],
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
	cited doc.ExpressionLike,
) CheckoutClauseLike {
	return CheckoutClauseClass().CheckoutClause(
		recipient,
		optionalAtLevel,
		cited,
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
	draft doc.ExpressionLike,
) DiscardClauseLike {
	return DiscardClauseClass().DiscardClause(
		draft,
	)
}

func DoClauseClass() DoClauseClassLike {
	return doc.DoClauseClass()
}

func DoClause(
	invocation any,
) DoClauseLike {
	return DoClauseClass().DoClause(
		invocation,
	)
}

func DocumentClass() DocumentClassLike {
	return doc.DocumentClass()
}

func Document(
	component any,
	optionalParameters doc.ParametersLike,
	optionalNote string,
) DocumentLike {
	return DocumentClass().Document(
		component,
		optionalParameters,
		optionalNote,
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
	inverse string,
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
	entities fra.ListLike[doc.DocumentLike],
) ItemsLike {
	return ItemsClass().Items(
		entities,
	)
}

func LetClauseClass() LetClauseClassLike {
	return doc.LetClauseClass()
}

func LetClause(
	recipient any,
	assignment string,
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
	invoke string,
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
	cited doc.ExpressionLike,
) NotarizeClauseLike {
	return NotarizeClauseClass().NotarizeClause(
		draft,
		cited,
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

func ParametersClass() ParametersClassLike {
	return doc.ParametersClass()
}

func Parameters(
	associations fra.CatalogLike[fra.SymbolLike, doc.DocumentLike],
) ParametersLike {
	return ParametersClass().Parameters(
		associations,
	)
}

func PostClauseClass() PostClauseClassLike {
	return doc.PostClauseClass()
}

func PostClause(
	message doc.ExpressionLike,
	bag doc.ExpressionLike,
) PostClauseLike {
	return PostClauseClass().PostClause(
		message,
		bag,
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
	operation string,
	expression doc.ExpressionLike,
) PredicateLike {
	return PredicateClass().Predicate(
		operation,
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
	left doc.Extent,
	first any,
	last any,
	right doc.Extent,
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
	indirect any,
) ReferentLike {
	return ReferentClass().Referent(
		indirect,
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
	cited doc.ExpressionLike,
) SaveClauseLike {
	return SaveClauseClass().SaveClause(
		draft,
		cited,
	)
}

func SelectClauseClass() SelectClauseClassLike {
	return doc.SelectClauseClass()
}

func SelectClause(
	target any,
	matchingClauses fra.ListLike[doc.MatchingClauseLike],
) SelectClauseLike {
	return SelectClauseClass().SelectClause(
		target,
		matchingClauses,
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
