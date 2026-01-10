/*
................................................................................
.    Copyright (c) 2009-2026 Crater Dog Technologies™.  All Rights Reserved.   .
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
	com "github.com/craterdog/go-essential-composites/v8"
	pri "github.com/craterdog/go-essential-primitives/v8"
)

// TYPE DECLARATIONS

/*
Assignment is a constrained type specifying a specific type of assignment.
*/
type Assignment uint8

const (
	Default Assignment = iota
	Assign
	Add
	Subtract
	Multiply
	Divide
	Join
)

/*
Inverse is a constrained type specifying a type of numerical inverse.
*/
type Inverse uint8

const (
	Additive Inverse = iota
	Multiplicative
	Conjugate
)

/*
Operation is a constrained type representing an expression operation.
*/
type Operation uint16

const (
	Less Operation = iota
	Equal
	More
	Is
	Matches
	And
	San
	Ior
	Xor
	Sum
	Difference
	Product
	Quotient
	Remainder
	Power
	Chain
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
		bag ExpressionLike,
	) AcceptClauseLike
}

/*
AssignClauseClassLike is a class interface that declares the complete set of class
constructors, constants and functions that must be supported by each concrete
let-clause-like class.
*/
type AssignClauseClassLike interface {
	// Constructor Methods
	AssignClause(
		recipient any,
		assignment Assignment,
		expression ExpressionLike,
	) AssignClauseLike
}

/*
AttributesClassLike is a class interface that declares the complete set of class
constructors, constants and functions that must be supported by each concrete
attributes-like class.
*/
type AttributesClassLike interface {
	// Constructor Methods
	Attributes(
		associations com.CatalogLike[any, ComponentLike],
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
		reversible any,
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
		literal any,
		optionalParameters ParametersLike,
		optionalNote string,
	) ComponentLike
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
DefineClauseClassLike is a class interface that declares the complete set of class
constructors, constants and functions that must be supported by each concrete
define-clause-like class.
*/
type DefineClauseClassLike interface {
	// Constructor Methods
	DefineClause(
		constant pri.SymbolLike,
		expression ExpressionLike,
	) DefineClauseLike
}

/*
DiscardClauseClassLike is a class interface that declares the complete set of
class constructors, constants and functions that must be supported by each
concrete discard-clause-like class.
*/
type DiscardClauseClassLike interface {
	// Constructor Methods
	DiscardClause(
		citation ExpressionLike,
	) DiscardClauseLike
}

/*
DocumentClassLike is a class interface that declares the complete set of class
constructors, constants and functions that must be supported by each concrete
document-like class.
*/
type DocumentClassLike interface {
	// Constructor Methods
	Document(
		optionalComment string,
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
		predicates com.Sequential[PredicateLike],
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
		arguments com.Sequential[any],
	) FunctionLike
}

/*
ParametersClassLike is a class interface that declares the complete set of
class constructors, constants and functions that must be supported by each
concrete parameters-like class.
*/
type ParametersClassLike interface {
	// Constructor Methods
	Parameters(
		constraints com.CatalogLike[pri.SymbolLike, ComponentLike],
	) ParametersLike
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
InspectClauseClassLike is a class interface that declares the complete set of
class constructors, constants and functions that must be supported by each
concrete inspect-clause-like class.
*/
type InspectClauseClassLike interface {
	// Constructor Methods
	InspectClause(
		recipient any,
		location ExpressionLike,
	) InspectClauseLike
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
InvokeClauseClassLike is a class interface that declares the complete set of class
constructors, constants and functions that must be supported by each concrete
invoke-clause-like class.
*/
type InvokeClauseClassLike interface {
	// Constructor Methods
	InvokeClause(
		method MethodLike,
	) InvokeClauseLike
}

/*
ItemsClassLike is a class interface that declares the complete set of class
constructors, constants and functions that must be supported by each concrete
items-like class.
*/
type ItemsClassLike interface {
	// Constructor Methods
	Items(
		components com.Sequential[ComponentLike],
	) ItemsLike
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
		identifier string,
		arguments com.Sequential[any],
		isSynchronous bool,
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
OnClauseClassLike is a class interface that declares the complete set of class
constructors, constants and functions that must be supported by each concrete
on-clause-like class.
*/
type OnClauseClassLike interface {
	// Constructor Methods
	OnClause(
		symbol pri.SymbolLike,
		matchingClauses com.Sequential[MatchingClauseLike],
	) OnClauseLike
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
		operation Operation,
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
		statements com.Sequential[StatementLike],
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
		message ExpressionLike,
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
		left com.Bracket,
		first any,
		last any,
		right com.Bracket,
	) RangeLike
}

/*
ReceiveClauseClassLike is a class interface that declares the complete set of
class constructors, constants and functions that must be supported by each
concrete receive-clause-like class.
*/
type ReceiveClauseClassLike interface {
	// Constructor Methods
	ReceiveClause(
		recipient any,
		bag ExpressionLike,
	) ReceiveClauseLike
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
		bag ExpressionLike,
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
		citation ExpressionLike,
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
		recipient any,
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
		template ExpressionLike,
		matchingClauses com.Sequential[MatchingClauseLike],
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
		optionalComment string,
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
		indexes com.Sequential[any],
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
		symbol pri.SymbolLike,
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
	GetBag() ExpressionLike
}

/*
AssignClauseLike is an instance interface that declares the complete set of
principal, attribute and aspect methods that must be supported by each instance
of a concrete let-clause-like class.
*/
type AssignClauseLike interface {
	// Principal Methods
	GetClass() AssignClauseClassLike

	// Attribute Methods
	GetRecipient() any
	GetAssignment() Assignment
	GetExpression() ExpressionLike
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
	GetAssociations() com.CatalogLike[any, ComponentLike]
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
	GetReversible() any
}

/*
ComponentLike is an instance interface that declares the complete set of
principal, attribute and aspect methods that must be supported by each instance
of a concrete component-like class.
*/
type ComponentLike interface {
	// Principal Methods
	GetClass() ComponentClassLike

	// Aspect Interfaces
	Composite
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
DefineClauseLike is an instance interface that declares the complete set of
principal, attribute and aspect methods that must be supported by each instance
of a concrete define-clause-like class.
*/
type DefineClauseLike interface {
	// Principal Methods
	GetClass() DefineClauseClassLike

	// Attribute Methods
	GetConstant() pri.SymbolLike
	GetExpression() ExpressionLike
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
	GetCitation() ExpressionLike
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
	GetOptionalComment() string
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
	GetPredicates() com.Sequential[PredicateLike]
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
	GetArguments() com.Sequential[any]
}

/*
ParametersLike is an instance interface that declares the complete set of
principal, attribute and aspect methods that must be supported by each instance
of a concrete parameters-like class.
*/
type ParametersLike interface {
	// Principal Methods
	GetClass() ParametersClassLike

	// Attribute Methods
	GetConstraints() com.CatalogLike[pri.SymbolLike, ComponentLike]
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
InspectClauseLike is an instance interface that declares the complete set of
principal, attribute and aspect methods that must be supported by each instance
of a concrete inspect-clause-like class.
*/
type InspectClauseLike interface {
	// Principal Methods
	GetClass() InspectClauseClassLike

	// Attribute Methods
	GetRecipient() any
	GetLocation() ExpressionLike
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
InvokeClauseLike is an instance interface that declares the complete set of
principal, attribute and aspect methods that must be supported by each instance
of a concrete invoke-clause-like class.
*/
type InvokeClauseLike interface {
	// Principal Methods
	GetClass() InvokeClauseClassLike

	// Attribute Methods
	GetMethod() MethodLike
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
	GetComponents() com.Sequential[ComponentLike]
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
	GetIdentifier() string
	GetArguments() com.Sequential[any]
	IsSynchronous() bool
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
OnClauseLike is an instance interface that declares the complete set of
principal, attribute and aspect methods that must be supported by each instance
of a concrete on-clause-like class.
*/
type OnClauseLike interface {
	// Principal Methods
	GetClass() OnClauseClassLike

	// Attribute Methods
	GetSymbol() pri.SymbolLike
	GetMatchingClauses() com.Sequential[MatchingClauseLike]
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
	GetOperation() Operation
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
	GetStatements() com.Sequential[StatementLike]
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
	GetMessage() ExpressionLike
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
	GetLeft() com.Bracket
	GetOptionalFirst() any
	GetOptionalLast() any
	GetRight() com.Bracket
}

/*
ReceiveClauseLike is an instance interface that declares the complete set of
principal, attribute and aspect methods that must be supported by each instance
of a concrete receive-clause-like class.
*/
type ReceiveClauseLike interface {
	// Principal Methods
	GetClass() ReceiveClauseClassLike

	// Attribute Methods
	GetRecipient() any
	GetBag() ExpressionLike
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
	GetBag() ExpressionLike
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
	GetCitation() ExpressionLike
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
	GetRecipient() any
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
	GetTemplate() ExpressionLike
	GetMatchingClauses() com.Sequential[MatchingClauseLike]
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
	GetOptionalComment() string
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
	GetIndexes() com.Sequential[any]
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
	GetSymbol() pri.SymbolLike
	GetSequence() ExpressionLike
	GetProcedure() ProcedureLike
}

// ASPECT DECLARATIONS

/*
Composite declares the set of method signatures that must be supported by
all composite documents.
*/
type Composite interface {
	GetLiteral() any
	GetOptionalParameters() ParametersLike
	GetOptionalNote() string
	GetConstraint(
		symbol pri.SymbolLike,
	) ComponentLike
	SetSubcomponent(
		value any,
		indices ...any,
	)
	GetSubcomponent(
		indices ...any,
	) ComponentLike
	RemoveSubcomponent(
		indices ...any,
	) ComponentLike
}
