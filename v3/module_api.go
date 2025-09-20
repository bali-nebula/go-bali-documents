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
	uti "github.com/craterdog/go-missing-utilities/v7"
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

// Assembly

type (
	Operation = ass.Operation
	Modifier  = ass.Modifier
	Operand   = ass.Operand
)

const (
	JumpOperation         = ass.JumpOperation
	PushOperation         = ass.PushOperation
	PullOperation         = ass.PullOperation
	LoadOperation         = ass.LoadOperation
	SaveOperation         = ass.SaveOperation
	DropOperation         = ass.DropOperation
	CallOperation         = ass.CallOperation
	SendOperation         = ass.SendOperation
	OnAnyModifier         = ass.OnAnyModifier
	OnNoneModifier        = ass.OnNoneModifier
	OnFalseModifier       = ass.OnFalseModifier
	OnEmptyModifier       = ass.OnEmptyModifier
	LiteralModifier       = ass.LiteralModifier
	ConstantModifier      = ass.ConstantModifier
	ArgumentModifier      = ass.ArgumentModifier
	HandlerModifier       = ass.HandlerModifier
	ComponentModifier     = ass.ComponentModifier
	ResultModifier        = ass.ResultModifier
	ExceptionModifier     = ass.ExceptionModifier
	DraftModifier         = ass.DraftModifier
	ContractModifier      = ass.ContractModifier
	VariableModifier      = ass.VariableModifier
	MessageModifier       = ass.MessageModifier
	With0Modifier         = ass.With0Modifier
	With1Modifier         = ass.With1Modifier
	With2Modifier         = ass.With2Modifier
	With3Modifier         = ass.With3Modifier
	ComponentWithModifier = ass.ComponentWithModifier
	ContractWithModifier  = ass.ContractWithModifier
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
	ObjectClassLike           = doc.ObjectClassLike
	OnClauseClassLike         = doc.OnClauseClassLike
	ParameterizationClassLike = doc.ParameterizationClassLike
	PostClauseClassLike       = doc.PostClauseClassLike
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
	ObjectLike           = doc.ObjectLike
	OnClauseLike         = doc.OnClauseLike
	ParameterizationLike = doc.ParameterizationLike
	PostClauseLike       = doc.PostClauseLike
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
	StatementLike        = doc.StatementLike
	SubcomponentLike     = doc.SubcomponentLike
	ThrowClauseLike      = doc.ThrowClauseLike
	WhileClauseLike      = doc.WhileClauseLike
	WithClauseLike       = doc.WithClauseLike
)

type (
	Declarative = doc.Declarative
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
	associations fra.CatalogLike[any, doc.ObjectLike],
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
	component doc.ComponentLike,
) DocumentLike {
	return DocumentClass().Document(
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
	objects fra.ListLike[doc.ObjectLike],
) ItemsLike {
	return ItemsClass().Items(
		objects,
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

func ObjectClass() ObjectClassLike {
	return doc.ObjectClass()
}

func Object(
	component doc.ComponentLike,
	optionalNote string,
) ObjectLike {
	return ObjectClass().Object(
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

func ParseSource(
	source string,
) ComponentLike {
	var inflator = Inflator()
	var parser = not.Parser()
	var document = inflator.InflateDocument(parser.ParseSource(source))
	return document.GetComponent()
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
	case ObjectLike:
		component = actual.GetComponent()
	default:
		component = Component(value, nil)
	}
	var source = FormatDocument(component)
	return source[:len(source)-1] // Remove the trailing newline.
}

func FormatDocument(
	component ComponentLike,
) string {
	var deflator = Deflator()
	var formatter = not.Formatter()
	var document = deflator.DeflateDocument(Document(component))
	var source = formatter.FormatDocument(document)
	return source
}

// Agents

type (
	Event       = fra.Event
	Rank        = fra.Rank
	Slot        = fra.Slot
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

func Collator[V any]() CollatorLike[V] {
	return CollatorClass[V]().Collator()
}

func CollatorWithMaximumDepth[V any](
	maximumDepth uti.Cardinal,
) CollatorLike[V] {
	return CollatorClass[V]().CollatorWithMaximumDepth(
		maximumDepth,
	)
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

func Sorter[V any]() SorterLike[V] {
	return SorterClass[V]().Sorter()
}

func SorterWithRanker[V any](
	ranker fra.RankingFunction[V],
) SorterLike[V] {
	return SorterClass[V]().SorterWithRanker(
		ranker,
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
	value any,
) CatalogLike[K, V] {
	switch actual := value.(type) {
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

func CatalogFromArray[K comparable, V any](
	associations []fra.AssociationLike[K, V],
) CatalogLike[K, V] {
	return CatalogClass[K, V]().CatalogFromArray(
		associations,
	)
}

func CatalogFromMap[K comparable, V any](
	associations map[K]V,
) CatalogLike[K, V] {
	return CatalogClass[K, V]().CatalogFromMap(
		associations,
	)
}

func CatalogFromSequence[K comparable, V any](
	associations fra.Sequential[fra.AssociationLike[K, V]],
) CatalogLike[K, V] {
	return CatalogClass[K, V]().CatalogFromSequence(
		associations,
	)
}

func ListClass[V any]() ListClassLike[V] {
	return fra.ListClass[V]()
}

func List[V any]() ListLike[V] {
	return ListClass[V]().List()
}

func ListFromArray[V any](
	values []V,
) ListLike[V] {
	return ListClass[V]().ListFromArray(
		values,
	)
}

func ListFromSequence[V any](
	values fra.Sequential[V],
) ListLike[V] {
	return ListClass[V]().ListFromSequence(
		values,
	)
}

func QueueClass[V any]() QueueClassLike[V] {
	return fra.QueueClass[V]()
}

func Queue[V any]() QueueLike[V] {
	return QueueClass[V]().Queue()
}

func QueueWithCapacity[V any](
	capacity uti.Cardinal,
) QueueLike[V] {
	return QueueClass[V]().QueueWithCapacity(
		capacity,
	)
}

func QueueFromArray[V any](
	values []V,
) QueueLike[V] {
	return QueueClass[V]().QueueFromArray(
		values,
	)
}

func QueueFromSequence[V any](
	values fra.Sequential[V],
) QueueLike[V] {
	return QueueClass[V]().QueueFromSequence(
		values,
	)
}

func SetClass[V any]() SetClassLike[V] {
	return fra.SetClass[V]()
}

func Set[V any]() SetLike[V] {
	return SetClass[V]().Set()
}

func SetWithCollator[V any](
	collator fra.CollatorLike[V],
) SetLike[V] {
	return SetClass[V]().SetWithCollator(
		collator,
	)
}

func SetFromArray[V any](
	values []V,
) SetLike[V] {
	return SetClass[V]().SetFromArray(
		values,
	)
}

func SetFromSequence[V any](
	values fra.Sequential[V],
) SetLike[V] {
	return SetClass[V]().SetFromSequence(
		values,
	)
}

func StackClass[V any]() StackClassLike[V] {
	return fra.StackClass[V]()
}

func Stack[V any]() StackLike[V] {
	return StackClass[V]().Stack()
}

func StackWithCapacity[V any](
	capacity uti.Cardinal,
) StackLike[V] {
	return StackClass[V]().StackWithCapacity(
		capacity,
	)
}

func StackFromArray[V any](
	values []V,
) StackLike[V] {
	return StackClass[V]().StackFromArray(
		values,
	)
}

func StackFromSequence[V any](
	values fra.Sequential[V],
) StackLike[V] {
	return StackClass[V]().StackFromSequence(
		values,
	)
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
	radians float64,
) AngleLike {
	return AngleClass().Angle(
		radians,
	)
}

func AngleFromString(
	source string,
) AngleLike {
	return AngleClass().AngleFromString(
		source,
	)
}

func BooleanClass() BooleanClassLike {
	return fra.BooleanClass()
}

func Boolean(
	boolean bool,
) BooleanLike {
	return BooleanClass().Boolean(
		boolean,
	)
}

func BooleanFromString(
	source string,
) BooleanLike {
	return BooleanClass().BooleanFromString(
		source,
	)
}

func DurationClass() DurationClassLike {
	return fra.DurationClass()
}

func Duration(
	milliseconds int,
) DurationLike {
	return DurationClass().Duration(
		milliseconds,
	)
}

func DurationFromString(
	source string,
) DurationLike {
	return DurationClass().DurationFromString(
		source,
	)
}

func GlyphClass() GlyphClassLike {
	return fra.GlyphClass()
}

func Glyph(
	rune_ rune,
) GlyphLike {
	return GlyphClass().Glyph(
		rune_,
	)
}

func GlyphFromInteger(
	integer int,
) GlyphLike {
	return GlyphClass().GlyphFromInteger(
		integer,
	)
}

func GlyphFromString(
	source string,
) GlyphLike {
	return GlyphClass().GlyphFromString(
		source,
	)
}

func MomentClass() MomentClassLike {
	return fra.MomentClass()
}

func Now() MomentLike {
	return fra.MomentClass().Now()
}

func Moment(
	milliseconds int,
) MomentLike {
	return MomentClass().Moment(
		milliseconds,
	)
}

func MomentFromString(
	source string,
) MomentLike {
	return MomentClass().MomentFromString(
		source,
	)
}

func NumberClass() NumberClassLike {
	return fra.NumberClass()
}

func Number(
	complex_ complex128,
) NumberLike {
	return NumberClass().Number(
		complex_,
	)
}

func NumberFromPolar(
	magnitude float64,
	phase float64,
) NumberLike {
	return NumberClass().NumberFromPolar(
		magnitude,
		phase,
	)
}

func NumberFromRectangular(
	real_ float64,
	imaginary float64,
) NumberLike {
	return NumberClass().NumberFromRectangular(
		real_,
		imaginary,
	)
}

func NumberFromString(
	source string,
) NumberLike {
	return NumberClass().NumberFromString(
		source,
	)
}

func PercentageClass() PercentageClassLike {
	return fra.PercentageClass()
}

func Percentage(
	float float64,
) PercentageLike {
	return PercentageClass().Percentage(
		float,
	)
}

func PercentageFromInteger(
	integer int,
) PercentageLike {
	return PercentageClass().PercentageFromInteger(
		integer,
	)
}

func PercentageFromString(
	source string,
) PercentageLike {
	return PercentageClass().PercentageFromString(
		source,
	)
}

func ProbabilityClass() ProbabilityClassLike {
	return fra.ProbabilityClass()
}

func Random() ProbabilityLike {
	return fra.ProbabilityClass().Random()
}

func Probability(
	float float64,
) ProbabilityLike {
	return ProbabilityClass().Probability(
		float,
	)
}

func ProbabilityFromBoolean(
	boolean bool,
) ProbabilityLike {
	return ProbabilityClass().ProbabilityFromBoolean(
		boolean,
	)
}

func ProbabilityFromString(
	source string,
) ProbabilityLike {
	return ProbabilityClass().ProbabilityFromString(
		source,
	)
}

func ResourceClass() ResourceClassLike {
	return fra.ResourceClass()
}

func Resource(
	string_ string,
) ResourceLike {
	return ResourceClass().Resource(
		string_,
	)
}

func ResourceFromString(
	source string,
) ResourceLike {
	return ResourceClass().ResourceFromString(
		source,
	)
}

func ResourceFromUri(
	url *uri.URL,
) ResourceLike {
	return ResourceClass().ResourceFromUri(
		url,
	)
}

func SymbolClass() SymbolClassLike {
	return fra.SymbolClass()
}

func Symbol(
	string_ string,
) SymbolLike {
	return SymbolClass().Symbol(
		string_,
	)
}

func SymbolFromString(
	source string,
) SymbolLike {
	return SymbolClass().SymbolFromString(
		source,
	)
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

// Strings

type (
	Identifier = fra.Identifier
	Line       = fra.Line
	Character  = fra.Character
)

type (
	BinaryClassLike    = fra.BinaryClassLike
	NameClassLike      = fra.NameClassLike
	NarrativeClassLike = fra.NarrativeClassLike
	PatternClassLike   = fra.PatternClassLike
	QuoteClassLike     = fra.QuoteClassLike
	TagClassLike       = fra.TagClassLike
	VersionClassLike   = fra.VersionClassLike
)

type (
	BinaryLike    = fra.BinaryLike
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
	bytes []byte,
) BinaryLike {
	return BinaryClass().Binary(
		bytes,
	)
}

func BinaryFromSequence(
	sequence fra.Sequential[byte],
) BinaryLike {
	return BinaryClass().BinaryFromSequence(
		sequence,
	)
}

func BinaryFromString(
	source string,
) BinaryLike {
	return BinaryClass().BinaryFromString(
		source,
	)
}

func NameClass() NameClassLike {
	return fra.NameClass()
}

func Name(
	identifiers []fra.Identifier,
) NameLike {
	return NameClass().Name(
		identifiers,
	)
}

func NameFromSequence(
	sequence fra.Sequential[fra.Identifier],
) NameLike {
	return NameClass().NameFromSequence(
		sequence,
	)
}

func NameFromString(
	source string,
) NameLike {
	return NameClass().NameFromString(
		source,
	)
}

func NarrativeClass() NarrativeClassLike {
	return fra.NarrativeClass()
}

func Narrative(
	lines []fra.Line,
) NarrativeLike {
	return NarrativeClass().Narrative(
		lines,
	)
}

func NarrativeFromSequence(
	sequence fra.Sequential[fra.Line],
) NarrativeLike {
	return NarrativeClass().NarrativeFromSequence(
		sequence,
	)
}

func NarrativeFromString(
	source string,
) NarrativeLike {
	return NarrativeClass().NarrativeFromString(
		source,
	)
}

func PatternClass() PatternClassLike {
	return fra.PatternClass()
}

func Pattern(
	characters []fra.Character,
) PatternLike {
	return PatternClass().Pattern(
		characters,
	)
}

func PatternFromSequence(
	sequence fra.Sequential[fra.Character],
) PatternLike {
	return PatternClass().PatternFromSequence(
		sequence,
	)
}

func PatternFromString(
	source string,
) PatternLike {
	return PatternClass().PatternFromString(
		source,
	)
}

func QuoteClass() QuoteClassLike {
	return fra.QuoteClass()
}

func Quote(
	characters []fra.Character,
) QuoteLike {
	return QuoteClass().Quote(
		characters,
	)
}

func QuoteFromSequence(
	sequence fra.Sequential[fra.Character],
) QuoteLike {
	return QuoteClass().QuoteFromSequence(
		sequence,
	)
}

func QuoteFromString(
	source string,
) QuoteLike {
	return QuoteClass().QuoteFromString(
		source,
	)
}

func TagClass() TagClassLike {
	return fra.TagClass()
}

func Tag(
	bytes []byte,
) TagLike {
	return TagClass().Tag(
		bytes,
	)
}

func TagWithSize(
	size uti.Cardinal,
) TagLike {
	return TagClass().TagWithSize(
		size,
	)
}

func TagFromSequence(
	sequence fra.Sequential[byte],
) TagLike {
	return TagClass().TagFromSequence(
		sequence,
	)
}

func TagFromString(
	source string,
) TagLike {
	return TagClass().TagFromString(
		source,
	)
}

func VersionClass() VersionClassLike {
	return fra.VersionClass()
}

func Version(
	ordinals []uti.Ordinal,
) VersionLike {
	return VersionClass().Version(
		ordinals,
	)
}

func VersionFromSequence(
	sequence fra.Sequential[uti.Ordinal],
) VersionLike {
	return VersionClass().VersionFromSequence(
		sequence,
	)
}

func VersionFromString(
	source string,
) VersionLike {
	return VersionClass().VersionFromString(
		source,
	)
}
