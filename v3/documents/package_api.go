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
Package "documents" provides the class definitions for each part of in an
"inflated" Bali Document™.  An inflated document can be manipulated by the
agents defined in this project.

For detailed documentation on this package refer to the wiki:
  - https://github.com/bali-nebula/go-bali-documents/wiki

This package follows the Crater Dog Technologies™ Go Coding Conventions located
here:
  - https://github.com/craterdog/go-development-tools/wiki/Coding-Conventions

Additional concrete implementations of the classes declared by this package can
be developed and used seamlessly since the interface declarations only depend on
other interfaces and intrinsic types—and the class implementations only depend
on interfaces, not on each other.
*/
package documents

import (
	fra "github.com/craterdog/go-component-framework/v7"
)

// TYPE DECLARATIONS

/*
Assignment is a constrained type specifying a type of mathematical assignment.
*/
type Assignment uint8

const (
	Equals Assignment = iota
	EqualsDefault
	EqualsPlus
	EqualsMinus
	EqualsTimes
	EqualsDivide
)

/*
Inverse is a constrained type specifying a type of mathematical inverse.
*/
type Inverse uint8

const (
	Additive Inverse = iota
	Multiplicative
	Conjugate
)

/*
Invoke is a constrained type specifying whether or not an invocation is
synchronous or asynchronous.
*/
type Invoke uint8

const (
	Synchronous Invoke = iota
	Asynchronous
)

/*
Operator is a constrained type representing a mathematical operator.
*/
type Operator uint16

const (
	Chain Operator = iota
	And
	San
	Ior
	Xor
	Plus
	Minus
	Times
	Divide
	Remainder
	Power
	Less
	Equal
	More
	Is
	Matches
)

// FUNCTIONAL DECLARATIONS

// CLASS DECLARATIONS

/*
AcceptClauseClassLike is a class interface that declares the complete set of
class constructors, constants and functions that must be supported by each
concrete accept-clause-like class.
*/
type AcceptClauseClassLike interface {
	// Constructor Methods
	AcceptClause(
		message ExpressionLike,
	) AcceptClauseLike
}

/*
AttributesClassLike is a class interface that declares the complete set of class
constructors, constants and functions that must be supported by each concrete
attributes-like class.
*/
type AttributesClassLike interface {
	// Constructor Methods
	Attributes(
		associations fra.CatalogLike[any, ObjectLike],
	) AttributesLike
}

/*
BreakClauseClassLike is a class interface that declares the complete set of
class constructors, constants and functions that must be supported by each
concrete break-clause-like class.
*/
type BreakClauseClassLike interface {
	// Constructor Methods
	BreakClause() BreakClauseLike
}

/*
BytecodeClassLike is a class interface that declares the complete set of
class constructors, constants and functions that must be supported by each
concrete bytecode-like class.
*/
type BytecodeClassLike interface {
	// Constructor Methods
	Bytecode(
		instructions []uint16,
	) BytecodeLike
	BytecodeFromString(
		source string,
	) BytecodeLike
}

/*
CheckoutClauseClassLike is a class interface that declares the complete set of
class constructors, constants and functions that must be supported by each
concrete checkout-clause-like class.
*/
type CheckoutClauseClassLike interface {
	// Constructor Methods
	CheckoutClause(
		recipient any,
		optionalAtLevel ExpressionLike,
		location ExpressionLike,
	) CheckoutClauseLike
}

/*
ComplementClassLike is a class interface that declares the complete set of class
constructors, constants and functions that must be supported by each concrete
complement-like class.
*/
type ComplementClassLike interface {
	// Constructor Methods
	Complement(
		logical any,
	) ComplementLike
}

/*
ComponentClassLike is a class interface that declares the complete set of class
constructors, constants and functions that must be supported by each concrete
component-like class.
*/
type ComponentClassLike interface {
	// Constructor Methods
	Component(
		entity any,
		optionalParameterization ParameterizationLike,
	) ComponentLike
}

/*
ConstraintClassLike is a class interface that declares the complete set of class
constructors, constants and functions that must be supported by each concrete
constraint-like class.
*/
type ConstraintClassLike interface {
	// Constructor Methods
	Constraint(
		metadata any,
		optionalParameterization ParameterizationLike,
	) ConstraintLike
}

/*
ContinueClauseClassLike is a class interface that declares the complete set of
class constructors, constants and functions that must be supported by each
concrete continue-clause-like class.
*/
type ContinueClauseClassLike interface {
	// Constructor Methods
	ContinueClause() ContinueClauseLike
}

/*
DiscardClauseClassLike is a class interface that declares the complete set of
class constructors, constants and functions that must be supported by each
concrete discard-clause-like class.
*/
type DiscardClauseClassLike interface {
	// Constructor Methods
	DiscardClause(
		location ExpressionLike,
	) DiscardClauseLike
}

/*
DoClauseClassLike is a class interface that declares the complete set of class
constructors, constants and functions that must be supported by each concrete
do-clause-like class.
*/
type DoClauseClassLike interface {
	// Constructor Methods
	DoClause(
		method MethodLike,
	) DoClauseLike
}

/*
DocumentClassLike is a class interface that declares the complete set of class
constructors, constants and functions that must be supported by each concrete
document-like class.
*/
type DocumentClassLike interface {
	// Constructor Methods
	Document(
		component ComponentLike,
	) DocumentLike
}

/*
ExpressionClassLike is a class interface that declares the complete set of class
constructors, constants and functions that must be supported by each concrete
expression-like class.
*/
type ExpressionClassLike interface {
	// Constructor Methods
	Expression(
		subject any,
		predicates fra.Sequential[PredicateLike],
	) ExpressionLike
}

/*
FunctionClassLike is a class interface that declares the complete set of class
constructors, constants and functions that must be supported by each concrete
function-like class.
*/
type FunctionClassLike interface {
	// Constructor Methods
	Function(
		identifier string,
		arguments fra.Sequential[any],
	) FunctionLike
}

/*
IfClauseClassLike is a class interface that declares the complete set of class
constructors, constants and functions that must be supported by each concrete
if-clause-like class.
*/
type IfClauseClassLike interface {
	// Constructor Methods
	IfClause(
		condition ExpressionLike,
		procedure ProcedureLike,
	) IfClauseLike
}

/*
InversionClassLike is a class interface that declares the complete set of class
constructors, constants and functions that must be supported by each concrete
inversion-like class.
*/
type InversionClassLike interface {
	// Constructor Methods
	Inversion(
		inverse Inverse,
		numerical any,
	) InversionLike
}

/*
ItemsClassLike is a class interface that declares the complete set of class
constructors, constants and functions that must be supported by each concrete
items-like class.
*/
type ItemsClassLike interface {
	// Constructor Methods
	Items(
		objects fra.Sequential[ObjectLike],
	) ItemsLike
}

/*
LetClauseClassLike is a class interface that declares the complete set of class
constructors, constants and functions that must be supported by each concrete
let-clause-like class.
*/
type LetClauseClassLike interface {
	// Constructor Methods
	LetClause(
		recipient any,
		assignment Assignment,
		expression ExpressionLike,
	) LetClauseLike
}

/*
MagnitudeClassLike is a class interface that declares the complete set of class
constructors, constants and functions that must be supported by each concrete
magnitude-like class.
*/
type MagnitudeClassLike interface {
	// Constructor Methods
	Magnitude(
		expression ExpressionLike,
	) MagnitudeLike
}

/*
MatchingClauseClassLike is a class interface that declares the complete set of
class constructors, constants and functions that must be supported by each
concrete matching-clause-like class.
*/
type MatchingClauseClassLike interface {
	// Constructor Methods
	MatchingClause(
		template ExpressionLike,
		procedure ProcedureLike,
	) MatchingClauseLike
}

/*
MethodClassLike is a class interface that declares the complete set of class
constructors, constants and functions that must be supported by each concrete
method-like class.
*/
type MethodClassLike interface {
	// Constructor Methods
	Method(
		target string,
		invoke Invoke,
		message string,
		arguments fra.Sequential[any],
	) MethodLike
}

/*
NotarizeClauseClassLike is a class interface that declares the complete set of
class constructors, constants and functions that must be supported by each
concrete notarize-clause-like class.
*/
type NotarizeClauseClassLike interface {
	// Constructor Methods
	NotarizeClause(
		draft ExpressionLike,
		location ExpressionLike,
	) NotarizeClauseLike
}

/*
ObjectClassLike is a class interface that declares the complete set of class
constructors, constants and functions that must be supported by each concrete
component-like class.
*/
type ObjectClassLike interface {
	// Constructor Methods
	Object(
		component ComponentLike,
		optionalNote string,
	) ObjectLike
}

/*
OnClauseClassLike is a class interface that declares the complete set of class
constructors, constants and functions that must be supported by each concrete
on-clause-like class.
*/
type OnClauseClassLike interface {
	// Constructor Methods
	OnClause(
		failure fra.SymbolLike,
		matchingClauses fra.Sequential[MatchingClauseLike],
	) OnClauseLike
}

/*
ParameterizationClassLike is a class interface that declares the complete set of
class constructors, constants and functions that must be supported by each
concrete parameterization-like class.
*/
type ParameterizationClassLike interface {
	// Constructor Methods
	Parameterization(
		parameters fra.CatalogLike[fra.SymbolLike, ConstraintLike],
	) ParameterizationLike
}

/*
PrecedenceClassLike is a class interface that declares the complete set of class
constructors, constants and functions that must be supported by each concrete
precedence-like class.
*/
type PrecedenceClassLike interface {
	// Constructor Methods
	Precedence(
		expression ExpressionLike,
	) PrecedenceLike
}

/*
PredicateClassLike is a class interface that declares the complete set of class
constructors, constants and functions that must be supported by each concrete
predicate-like class.
*/
type PredicateClassLike interface {
	// Constructor Methods
	Predicate(
		operator Operator,
		expression ExpressionLike,
	) PredicateLike
}

/*
ProcedureClassLike is a class interface that declares the complete set of class
constructors, constants and functions that must be supported by each concrete
procedure-like class.
*/
type ProcedureClassLike interface {
	// Constructor Methods
	Procedure(
		lines fra.Sequential[any],
	) ProcedureLike
}

/*
PublishClauseClassLike is a class interface that declares the complete set of
class constructors, constants and functions that must be supported by each
concrete publish-clause-like class.
*/
type PublishClauseClassLike interface {
	// Constructor Methods
	PublishClause(
		event ExpressionLike,
	) PublishClauseLike
}

/*
RangeClassLike is a class interface that declares the complete set of class
constructors, constants and functions that must be supported by each concrete
range-like class.
*/
type RangeClassLike interface {
	// Constructor Methods
	Range(
		left fra.Bracket,
		first any,
		last any,
		right fra.Bracket,
	) RangeLike
}

/*
ReferentClassLike is a class interface that declares the complete set of class
constructors, constants and functions that must be supported by each concrete
referent-like class.
*/
type ReferentClassLike interface {
	// Constructor Methods
	Referent(
		reference any,
	) ReferentLike
}

/*
RejectClauseClassLike is a class interface that declares the complete set of
class constructors, constants and functions that must be supported by each
concrete reject-clause-like class.
*/
type RejectClauseClassLike interface {
	// Constructor Methods
	RejectClause(
		message ExpressionLike,
	) RejectClauseLike
}

/*
RetrieveClauseClassLike is a class interface that declares the complete set of
class constructors, constants and functions that must be supported by each
concrete retrieve-clause-like class.
*/
type RetrieveClauseClassLike interface {
	// Constructor Methods
	RetrieveClause(
		recipient any,
		bag ExpressionLike,
	) RetrieveClauseLike
}

/*
ReturnClauseClassLike is a class interface that declares the complete set of
class constructors, constants and functions that must be supported by each
concrete return-clause-like class.
*/
type ReturnClauseClassLike interface {
	// Constructor Methods
	ReturnClause(
		result ExpressionLike,
	) ReturnClauseLike
}

/*
SaveClauseClassLike is a class interface that declares the complete set of class
constructors, constants and functions that must be supported by each concrete
save-clause-like class.
*/
type SaveClauseClassLike interface {
	// Constructor Methods
	SaveClause(
		draft ExpressionLike,
		location ExpressionLike,
	) SaveClauseLike
}

/*
SelectClauseClassLike is a class interface that declares the complete set of
class constructors, constants and functions that must be supported by each
concrete select-clause-like class.
*/
type SelectClauseClassLike interface {
	// Constructor Methods
	SelectClause(
		expression ExpressionLike,
		matchingClauses fra.Sequential[MatchingClauseLike],
	) SelectClauseLike
}

/*
SendClauseClassLike is a class interface that declares the complete set of class
constructors, constants and functions that must be supported by each concrete
send-clause-like class.
*/
type SendClauseClassLike interface {
	// Constructor Methods
	SendClause(
		message ExpressionLike,
		bag ExpressionLike,
	) SendClauseLike
}

/*
StatementClassLike is a class interface that declares the complete set of class
constructors, constants and functions that must be supported by each concrete
statement-like class.
*/
type StatementClassLike interface {
	// Constructor Methods
	Statement(
		mainClause any,
		optionalOnClause OnClauseLike,
	) StatementLike
}

/*
SubcomponentClassLike is a class interface that declares the complete set of
class constructors, constants and functions that must be supported by each
concrete subcomponent-like class.
*/
type SubcomponentClassLike interface {
	// Constructor Methods
	Subcomponent(
		identifier string,
		indexes fra.Sequential[any],
	) SubcomponentLike
}

/*
ThrowClauseClassLike is a class interface that declares the complete set of
class constructors, constants and functions that must be supported by each
concrete throw-clause-like class.
*/
type ThrowClauseClassLike interface {
	// Constructor Methods
	ThrowClause(
		exception ExpressionLike,
	) ThrowClauseLike
}

/*
WhileClauseClassLike is a class interface that declares the complete set of
class constructors, constants and functions that must be supported by each
concrete while-clause-like class.
*/
type WhileClauseClassLike interface {
	// Constructor Methods
	WhileClause(
		condition ExpressionLike,
		procedure ProcedureLike,
	) WhileClauseLike
}

/*
WithClauseClassLike is a class interface that declares the complete set of class
constructors, constants and functions that must be supported by each concrete
with-clause-like class.
*/
type WithClauseClassLike interface {
	// Constructor Methods
	WithClause(
		variable fra.SymbolLike,
		sequence ExpressionLike,
		procedure ProcedureLike,
	) WithClauseLike
}

// INSTANCE DECLARATIONS

/*
AcceptClauseLike is an instance interface that declares the complete set of
principal, attribute and aspect methods that must be supported by each instance
of a concrete accept-clause-like class.
*/
type AcceptClauseLike interface {
	// Principal Methods
	GetClass() AcceptClauseClassLike

	// Attribute Methods
	GetMessage() ExpressionLike
}

/*
AttributesLike is an instance interface that declares the complete set of
principal, attribute and aspect methods that must be supported by each instance
of a concrete attributes-like class.
*/
type AttributesLike interface {
	// Principal Methods
	GetClass() AttributesClassLike

	// Attribute Methods
	GetAssociations() fra.CatalogLike[any, ObjectLike]
}

/*
BreakClauseLike is an instance interface that declares the complete set of
principal, attribute and aspect methods that must be supported by each instance
of a concrete break-clause-like class.
*/
type BreakClauseLike interface {
	// Principal Methods
	GetClass() BreakClauseClassLike
}

/*
BytecodeLike is an instance interface that declares the complete set of
principal, attribute and aspect methods that must be supported by each instance
of a concrete bytecode-like class.
*/
type BytecodeLike interface {
	// Principal Methods
	GetClass() BytecodeClassLike
	AsString() string
	AsIntrinsic() []uint16

	// Aspect Interfaces
	fra.Sequential[uint16]
}

/*
CheckoutClauseLike is an instance interface that declares the complete set of
principal, attribute and aspect methods that must be supported by each instance
of a concrete checkout-clause-like class.
*/
type CheckoutClauseLike interface {
	// Principal Methods
	GetClass() CheckoutClauseClassLike

	// Attribute Methods
	GetRecipient() any
	GetOptionalAtLevel() ExpressionLike
	GetLocation() ExpressionLike
}

/*
ComplementLike is an instance interface that declares the complete set of
principal, attribute and aspect methods that must be supported by each instance
of a concrete complement-like class.
*/
type ComplementLike interface {
	// Principal Methods
	GetClass() ComplementClassLike

	// Attribute Methods
	GetLogical() any
}

/*
ComponentLike is an instance interface that declares the complete set of
principal, attribute and aspect methods that must be supported by each instance
of a concrete component-like class.
*/
type ComponentLike interface {
	// Principal Methods
	GetClass() ComponentClassLike
	GetParameter(
		symbol fra.SymbolLike,
	) ComponentLike
	SetObject(
		value any,
		indices ...any,
	)
	GetObject(
		indices ...any,
	) ObjectLike
	RemoveObject(
		indices ...any,
	) ObjectLike

	// Attribute Methods
	GetEntity() any
	GetOptionalParameterization() ParameterizationLike
}

/*
ConstraintLike is an instance interface that declares the complete set of
principal, attribute and aspect methods that must be supported by each instance
of a concrete constraint-like class.
*/
type ConstraintLike interface {
	// Principal Methods
	GetClass() ConstraintClassLike

	// Attribute Methods
	GetMetadata() any
	GetOptionalParameterization() ParameterizationLike
}

/*
ContinueClauseLike is an instance interface that declares the complete set of
principal, attribute and aspect methods that must be supported by each instance
of a concrete continue-clause-like class.
*/
type ContinueClauseLike interface {
	// Principal Methods
	GetClass() ContinueClauseClassLike
}

/*
DiscardClauseLike is an instance interface that declares the complete set of
principal, attribute and aspect methods that must be supported by each instance
of a concrete discard-clause-like class.
*/
type DiscardClauseLike interface {
	// Principal Methods
	GetClass() DiscardClauseClassLike

	// Attribute Methods
	GetLocation() ExpressionLike
}

/*
DoClauseLike is an instance interface that declares the complete set of
principal, attribute and aspect methods that must be supported by each instance
of a concrete do-clause-like class.
*/
type DoClauseLike interface {
	// Principal Methods
	GetClass() DoClauseClassLike

	// Attribute Methods
	GetMethod() MethodLike
}

/*
DocumentLike is an instance interface that declares the complete set of
principal, attribute and aspect methods that must be supported by each instance
of a concrete document-like class.
*/
type DocumentLike interface {
	// Principal Methods
	GetClass() DocumentClassLike

	// Attribute Methods
	GetComponent() ComponentLike
}

/*
ExpressionLike is an instance interface that declares the complete set of
principal, attribute and aspect methods that must be supported by each instance
of a concrete expression-like class.
*/
type ExpressionLike interface {
	// Principal Methods
	GetClass() ExpressionClassLike

	// Attribute Methods
	GetSubject() any
	GetPredicates() fra.Sequential[PredicateLike]
}

/*
FunctionLike is an instance interface that declares the complete set of
principal, attribute and aspect methods that must be supported by each instance
of a concrete function-like class.
*/
type FunctionLike interface {
	// Principal Methods
	GetClass() FunctionClassLike

	// Attribute Methods
	GetIdentifier() string
	GetArguments() fra.Sequential[any]
}

/*
IfClauseLike is an instance interface that declares the complete set of
principal, attribute and aspect methods that must be supported by each instance
of a concrete if-clause-like class.
*/
type IfClauseLike interface {
	// Principal Methods
	GetClass() IfClauseClassLike

	// Attribute Methods
	GetCondition() ExpressionLike
	GetProcedure() ProcedureLike
}

/*
InversionLike is an instance interface that declares the complete set of
principal, attribute and aspect methods that must be supported by each instance
of a concrete inversion-like class.
*/
type InversionLike interface {
	// Principal Methods
	GetClass() InversionClassLike

	// Attribute Methods
	GetInverse() Inverse
	GetNumerical() any
}

/*
ItemsLike is an instance interface that declares the complete set of principal,
attribute and aspect methods that must be supported by each instance of a
concrete items-like class.
*/
type ItemsLike interface {
	// Principal Methods
	GetClass() ItemsClassLike

	// Attribute Methods
	GetObjects() fra.Sequential[ObjectLike]
}

/*
LetClauseLike is an instance interface that declares the complete set of
principal, attribute and aspect methods that must be supported by each instance
of a concrete let-clause-like class.
*/
type LetClauseLike interface {
	// Principal Methods
	GetClass() LetClauseClassLike

	// Attribute Methods
	GetRecipient() any
	GetAssignment() Assignment
	GetExpression() ExpressionLike
}

/*
MagnitudeLike is an instance interface that declares the complete set of
principal, attribute and aspect methods that must be supported by each instance
of a concrete magnitude-like class.
*/
type MagnitudeLike interface {
	// Principal Methods
	GetClass() MagnitudeClassLike

	// Attribute Methods
	GetExpression() ExpressionLike
}

/*
MatchingClauseLike is an instance interface that declares the complete set of
principal, attribute and aspect methods that must be supported by each instance
of a concrete matching-clause-like class.
*/
type MatchingClauseLike interface {
	// Principal Methods
	GetClass() MatchingClauseClassLike

	// Attribute Methods
	GetTemplate() ExpressionLike
	GetProcedure() ProcedureLike
}

/*
MethodLike is an instance interface that declares the complete set of principal,
attribute and aspect methods that must be supported by each instance of a
concrete method-like class.
*/
type MethodLike interface {
	// Principal Methods
	GetClass() MethodClassLike

	// Attribute Methods
	GetTarget() string
	GetInvoke() Invoke
	GetMessage() string
	GetArguments() fra.Sequential[any]
}

/*
NotarizeClauseLike is an instance interface that declares the complete set of
principal, attribute and aspect methods that must be supported by each instance
of a concrete notarize-clause-like class.
*/
type NotarizeClauseLike interface {
	// Principal Methods
	GetClass() NotarizeClauseClassLike

	// Attribute Methods
	GetDraft() ExpressionLike
	GetLocation() ExpressionLike
}

/*
ObjectLike is an instance interface that declares the complete set of
principal, attribute and aspect methods that must be supported by each instance
of a concrete component-like class.
*/
type ObjectLike interface {
	// Principal Methods
	GetClass() ObjectClassLike

	// Attribute Methods
	GetComponent() ComponentLike
	GetOptionalNote() string
}

/*
OnClauseLike is an instance interface that declares the complete set of
principal, attribute and aspect methods that must be supported by each instance
of a concrete on-clause-like class.
*/
type OnClauseLike interface {
	// Principal Methods
	GetClass() OnClauseClassLike

	// Attribute Methods
	GetFailure() fra.SymbolLike
	GetMatchingClauses() fra.Sequential[MatchingClauseLike]
}

/*
ParameterizationLike is an instance interface that declares the complete set of
principal, attribute and aspect methods that must be supported by each instance
of a concrete parameterization-like class.
*/
type ParameterizationLike interface {
	// Principal Methods
	GetClass() ParameterizationClassLike

	// Attribute Methods
	GetParameters() fra.CatalogLike[fra.SymbolLike, ConstraintLike]
}

/*
PrecedenceLike is an instance interface that declares the complete set of
principal, attribute and aspect methods that must be supported by each instance
of a concrete precedence-like class.
*/
type PrecedenceLike interface {
	// Principal Methods
	GetClass() PrecedenceClassLike

	// Attribute Methods
	GetExpression() ExpressionLike
}

/*
PredicateLike is an instance interface that declares the complete set of
principal, attribute and aspect methods that must be supported by each instance
of a concrete predicate-like class.
*/
type PredicateLike interface {
	// Principal Methods
	GetClass() PredicateClassLike

	// Attribute Methods
	GetOperator() Operator
	GetExpression() ExpressionLike
}

/*
ProcedureLike is an instance interface that declares the complete set of
principal, attribute and aspect methods that must be supported by each instance
of a concrete procedure-like class.
*/
type ProcedureLike interface {
	// Principal Methods
	GetClass() ProcedureClassLike

	// Attribute Methods
	GetLines() fra.Sequential[any]
}

/*
PublishClauseLike is an instance interface that declares the complete set of
principal, attribute and aspect methods that must be supported by each instance
of a concrete publish-clause-like class.
*/
type PublishClauseLike interface {
	// Principal Methods
	GetClass() PublishClauseClassLike

	// Attribute Methods
	GetEvent() ExpressionLike
}

/*
RangeLike is an instance interface that declares the complete set of principal,
attribute and aspect methods that must be supported by each instance of a
concrete range-like class.
*/
type RangeLike interface {
	// Principal Methods
	GetClass() RangeClassLike

	// Attribute Methods
	GetLeft() fra.Bracket
	GetFirst() any
	GetLast() any
	GetRight() fra.Bracket
}

/*
ReferentLike is an instance interface that declares the complete set of
principal, attribute and aspect methods that must be supported by each instance
of a concrete referent-like class.
*/
type ReferentLike interface {
	// Principal Methods
	GetClass() ReferentClassLike

	// Attribute Methods
	GetReference() any
}

/*
RejectClauseLike is an instance interface that declares the complete set of
principal, attribute and aspect methods that must be supported by each instance
of a concrete reject-clause-like class.
*/
type RejectClauseLike interface {
	// Principal Methods
	GetClass() RejectClauseClassLike

	// Attribute Methods
	GetMessage() ExpressionLike
}

/*
RetrieveClauseLike is an instance interface that declares the complete set of
principal, attribute and aspect methods that must be supported by each instance
of a concrete retrieve-clause-like class.
*/
type RetrieveClauseLike interface {
	// Principal Methods
	GetClass() RetrieveClauseClassLike

	// Attribute Methods
	GetRecipient() any
	GetBag() ExpressionLike
}

/*
ReturnClauseLike is an instance interface that declares the complete set of
principal, attribute and aspect methods that must be supported by each instance
of a concrete return-clause-like class.
*/
type ReturnClauseLike interface {
	// Principal Methods
	GetClass() ReturnClauseClassLike

	// Attribute Methods
	GetResult() ExpressionLike
}

/*
SaveClauseLike is an instance interface that declares the complete set of
principal, attribute and aspect methods that must be supported by each instance
of a concrete save-clause-like class.
*/
type SaveClauseLike interface {
	// Principal Methods
	GetClass() SaveClauseClassLike

	// Attribute Methods
	GetDraft() ExpressionLike
	GetLocation() ExpressionLike
}

/*
SelectClauseLike is an instance interface that declares the complete set of
principal, attribute and aspect methods that must be supported by each instance
of a concrete select-clause-like class.
*/
type SelectClauseLike interface {
	// Principal Methods
	GetClass() SelectClauseClassLike

	// Attribute Methods
	GetExpression() ExpressionLike
	GetMatchingClauses() fra.Sequential[MatchingClauseLike]
}

/*
SendClauseLike is an instance interface that declares the complete set of
principal, attribute and aspect methods that must be supported by each instance
of a concrete send-clause-like class.
*/
type SendClauseLike interface {
	// Principal Methods
	GetClass() SendClauseClassLike

	// Attribute Methods
	GetMessage() ExpressionLike
	GetBag() ExpressionLike
}

/*
StatementLike is an instance interface that declares the complete set of
principal, attribute and aspect methods that must be supported by each instance
of a concrete statement-like class.
*/
type StatementLike interface {
	// Principal Methods
	GetClass() StatementClassLike

	// Attribute Methods
	GetMainClause() any
	GetOptionalOnClause() OnClauseLike
}

/*
SubcomponentLike is an instance interface that declares the complete set of
principal, attribute and aspect methods that must be supported by each instance
of a concrete subcomponent-like class.
*/
type SubcomponentLike interface {
	// Principal Methods
	GetClass() SubcomponentClassLike

	// Attribute Methods
	GetIdentifier() string
	GetIndexes() fra.Sequential[any]
}

/*
ThrowClauseLike is an instance interface that declares the complete set of
principal, attribute and aspect methods that must be supported by each instance
of a concrete throw-clause-like class.
*/
type ThrowClauseLike interface {
	// Principal Methods
	GetClass() ThrowClauseClassLike

	// Attribute Methods
	GetException() ExpressionLike
}

/*
WhileClauseLike is an instance interface that declares the complete set of
principal, attribute and aspect methods that must be supported by each instance
of a concrete while-clause-like class.
*/
type WhileClauseLike interface {
	// Principal Methods
	GetClass() WhileClauseClassLike

	// Attribute Methods
	GetCondition() ExpressionLike
	GetProcedure() ProcedureLike
}

/*
WithClauseLike is an instance interface that declares the complete set of
principal, attribute and aspect methods that must be supported by each instance
of a concrete with-clause-like class.
*/
type WithClauseLike interface {
	// Principal Methods
	GetClass() WithClauseClassLike

	// Attribute Methods
	GetVariable() fra.SymbolLike
	GetSequence() ExpressionLike
	GetProcedure() ProcedureLike
}

// ASPECT DECLARATIONS
