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
	ass "github.com/bali-nebula/go-bali-documents/v3/assembly"
	doc "github.com/bali-nebula/go-bali-documents/v3/documents"
	not "github.com/bali-nebula/go-document-notation/v3"
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

// Assembly

type (
	Operation = ass.Operation
	Modifier  = ass.Modifier
	Operand   = ass.Operand
)

const (
	Jump                   = ass.Jump
	Push                   = ass.Push
	Pull                   = ass.Pull
	Load                   = ass.Load
	Save                   = ass.Save
	Drop                   = ass.Drop
	Call                   = ass.Call
	Send                   = ass.Send
	OnAny                  = ass.OnAny
	OnEmpty                = ass.OnEmpty
	OnNone                 = ass.OnNone
	OnFalse                = ass.OnFalse
	Handler                = ass.Handler
	Literal                = ass.Literal
	Constant               = ass.Constant
	Argument               = ass.Argument
	Exception              = ass.Exception
	Component              = ass.Component
	Result                 = ass.Result
	Contract               = ass.Contract
	Draft                  = ass.Draft
	Message                = ass.Message
	Variable               = ass.Variable
	With0Arguments         = ass.With0Arguments
	With1Argument          = ass.With1Argument
	With2Arguments         = ass.With2Arguments
	With3Arguments         = ass.With3Arguments
	ContractWithArguments  = ass.ContractWithArguments
	ComponentWithArguments = ass.ComponentWithArguments
)

type (
	BytecodeClassLike    = ass.BytecodeClassLike
	InstructionClassLike = ass.InstructionClassLike
)

type (
	BytecodeLike    = ass.BytecodeLike
	InstructionLike = ass.InstructionLike
)

// Documents

type (
	Assignment = doc.Assignment
	Extent     = doc.Extent
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
	Inclusive      = doc.Inclusive
	Exclusive      = doc.Exclusive
	Additive       = doc.Additive
	Multiplicative = doc.Multiplicative
	Conjugate      = doc.Conjugate
	Synchronous    = doc.Synchronous
	Asynchronous   = doc.Asynchronous
	Chain          = doc.Chain
	And            = doc.And
	San            = doc.San
	Ior            = doc.Ior
	Eor            = doc.Eor
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
	AcceptClauseClassLike   = doc.AcceptClauseClassLike
	AttributesClassLike     = doc.AttributesClassLike
	BreakClauseClassLike    = doc.BreakClauseClassLike
	CheckoutClauseClassLike = doc.CheckoutClauseClassLike
	ComplementClassLike     = doc.ComplementClassLike
	ContinueClauseClassLike = doc.ContinueClauseClassLike
	DiscardClauseClassLike  = doc.DiscardClauseClassLike
	DoClauseClassLike       = doc.DoClauseClassLike
	DocumentClassLike       = doc.DocumentClassLike
	EntitiesClassLike       = doc.EntitiesClassLike
	ExpressionClassLike     = doc.ExpressionClassLike
	FunctionClassLike       = doc.FunctionClassLike
	IfClauseClassLike       = doc.IfClauseClassLike
	InversionClassLike      = doc.InversionClassLike
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
	EntitiesLike       = doc.EntitiesLike
	ExpressionLike     = doc.ExpressionLike
	FunctionLike       = doc.FunctionLike
	IfClauseLike       = doc.IfClauseLike
	InversionLike      = doc.InversionLike
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

// Assembly

func BytecodeClass() BytecodeClassLike {
	return ass.BytecodeClass()
}

func Bytecode(
	instructions []ass.InstructionLike,
) BytecodeLike {
	return BytecodeClass().Bytecode(
		instructions,
	)
}

func BytecodeFromString(
	source string,
) BytecodeLike {
	return BytecodeClass().BytecodeFromString(
		source,
	)
}

func InstructionClass() InstructionClassLike {
	return ass.InstructionClass()
}

func Instruction(
	operation ass.Operation,
	modifier ass.Modifier,
	operand ass.Operand,
) InstructionLike {
	return InstructionClass().Instruction(
		operation,
		modifier,
		operand,
	)
}

func InstructionFromInteger(
	integer uint16,
) InstructionLike {
	return InstructionClass().InstructionFromInteger(
		integer,
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

func EntitiesClass() EntitiesClassLike {
	return doc.EntitiesClass()
}

func Entities(
	items fra.ListLike[doc.DocumentLike],
) EntitiesLike {
	return EntitiesClass().Entities(
		items,
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
	expression doc.ExpressionLike,
	matchingClauses fra.ListLike[doc.MatchingClauseLike],
) SelectClauseLike {
	return SelectClauseClass().SelectClause(
		expression,
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

func FormatDocument(
	document doc.DocumentLike,
) string {
	var deflator = Deflator()
	var formatter = not.Formatter()
	return formatter.FormatDocument(deflator.DeflateDocument(document))
}

func ParseSource(
	source string,
) doc.DocumentLike {
	var inflator = Inflator()
	var parser = not.Parser()
	return inflator.InflateDocument(parser.ParseSource(source))
}
