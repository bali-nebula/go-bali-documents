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
│            This "package_api.go" file was automatically generated.           │
│                     Any updates to it may be overwritten.                    │
└──────────────────────────────────────────────────────────────────────────────┘

Package "ast" provides the abstract syntax tree (AST) classes for this module
based on the "syntax.cdsn" grammar for the module.  Each AST class manages the
attributes associated with its corresponding rule definition found in the
grammar.

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
package ast

import (
	col "github.com/craterdog/go-collection-framework/v7"
)

// TYPE DECLARATIONS

// FUNCTIONAL DECLARATIONS

// CLASS DECLARATIONS

/*
AcceptClauseClassLike is a class interface that declares the
complete set of class constructors, constants and functions that must be
supported by each concrete accept-clause-like class.
*/
type AcceptClauseClassLike interface {
	// Constructor Methods
	AcceptClause(
		delimiter string,
		message MessageLike,
	) AcceptClauseLike
}

/*
AdditionalArgumentClassLike is a class interface that declares the
complete set of class constructors, constants and functions that must be
supported by each concrete additional-argument-like class.
*/
type AdditionalArgumentClassLike interface {
	// Constructor Methods
	AdditionalArgument(
		delimiter string,
		argument ArgumentLike,
	) AdditionalArgumentLike
}

/*
AdditionalAssociationClassLike is a class interface that declares the
complete set of class constructors, constants and functions that must be
supported by each concrete additional-association-like class.
*/
type AdditionalAssociationClassLike interface {
	// Constructor Methods
	AdditionalAssociation(
		delimiter string,
		association AssociationLike,
	) AdditionalAssociationLike
}

/*
AdditionalIndexClassLike is a class interface that declares the
complete set of class constructors, constants and functions that must be
supported by each concrete additional-index-like class.
*/
type AdditionalIndexClassLike interface {
	// Constructor Methods
	AdditionalIndex(
		delimiter string,
		index IndexLike,
	) AdditionalIndexLike
}

/*
AdditionalStatementClassLike is a class interface that declares the
complete set of class constructors, constants and functions that must be
supported by each concrete additional-statement-like class.
*/
type AdditionalStatementClassLike interface {
	// Constructor Methods
	AdditionalStatement(
		delimiter string,
		statement StatementLike,
	) AdditionalStatementLike
}

/*
AdditionalValueClassLike is a class interface that declares the
complete set of class constructors, constants and functions that must be
supported by each concrete additional-value-like class.
*/
type AdditionalValueClassLike interface {
	// Constructor Methods
	AdditionalValue(
		delimiter string,
		component ComponentLike,
	) AdditionalValueLike
}

/*
AnnotatedAssociationClassLike is a class interface that declares the
complete set of class constructors, constants and functions that must be
supported by each concrete annotated-association-like class.
*/
type AnnotatedAssociationClassLike interface {
	// Constructor Methods
	AnnotatedAssociation(
		association AssociationLike,
		optionalNote string,
	) AnnotatedAssociationLike
}

/*
AnnotatedStatementClassLike is a class interface that declares the
complete set of class constructors, constants and functions that must be
supported by each concrete annotated-statement-like class.
*/
type AnnotatedStatementClassLike interface {
	// Constructor Methods
	AnnotatedStatement(
		any_ any,
	) AnnotatedStatementLike
}

/*
AnnotatedValueClassLike is a class interface that declares the
complete set of class constructors, constants and functions that must be
supported by each concrete annotated-value-like class.
*/
type AnnotatedValueClassLike interface {
	// Constructor Methods
	AnnotatedValue(
		component ComponentLike,
		optionalNote string,
	) AnnotatedValueLike
}

/*
ArgumentClassLike is a class interface that declares the
complete set of class constructors, constants and functions that must be
supported by each concrete argument-like class.
*/
type ArgumentClassLike interface {
	// Constructor Methods
	Argument(
		identifier string,
	) ArgumentLike
}

/*
ArgumentsClassLike is a class interface that declares the
complete set of class constructors, constants and functions that must be
supported by each concrete arguments-like class.
*/
type ArgumentsClassLike interface {
	// Constructor Methods
	Arguments(
		argument ArgumentLike,
		additionalArguments col.Sequential[AdditionalArgumentLike],
	) ArgumentsLike
}

/*
AssignmentClassLike is a class interface that declares the
complete set of class constructors, constants and functions that must be
supported by each concrete assignment-like class.
*/
type AssignmentClassLike interface {
	// Constructor Methods
	Assignment(
		any_ any,
	) AssignmentLike
}

/*
AssociationClassLike is a class interface that declares the
complete set of class constructors, constants and functions that must be
supported by each concrete association-like class.
*/
type AssociationClassLike interface {
	// Constructor Methods
	Association(
		primitive PrimitiveLike,
		delimiter string,
		component ComponentLike,
	) AssociationLike
}

/*
AtLevelClassLike is a class interface that declares the
complete set of class constructors, constants and functions that must be
supported by each concrete at-level-like class.
*/
type AtLevelClassLike interface {
	// Constructor Methods
	AtLevel(
		delimiter1 string,
		delimiter2 string,
		expression ExpressionLike,
	) AtLevelLike
}

/*
BagClassLike is a class interface that declares the
complete set of class constructors, constants and functions that must be
supported by each concrete bag-like class.
*/
type BagClassLike interface {
	// Constructor Methods
	Bag(
		expression ExpressionLike,
	) BagLike
}

/*
BlockingClassLike is a class interface that declares the
complete set of class constructors, constants and functions that must be
supported by each concrete blocking-like class.
*/
type BlockingClassLike interface {
	// Constructor Methods
	Blocking(
		any_ any,
	) BlockingLike
}

/*
BracketClassLike is a class interface that declares the
complete set of class constructors, constants and functions that must be
supported by each concrete bracket-like class.
*/
type BracketClassLike interface {
	// Constructor Methods
	Bracket(
		any_ any,
	) BracketLike
}

/*
BreakClauseClassLike is a class interface that declares the
complete set of class constructors, constants and functions that must be
supported by each concrete break-clause-like class.
*/
type BreakClauseClassLike interface {
	// Constructor Methods
	BreakClause(
		delimiter1 string,
		delimiter2 string,
	) BreakClauseLike
}

/*
CheckoutClauseClassLike is a class interface that declares the
complete set of class constructors, constants and functions that must be
supported by each concrete checkout-clause-like class.
*/
type CheckoutClauseClassLike interface {
	// Constructor Methods
	CheckoutClause(
		delimiter1 string,
		recipient RecipientLike,
		optionalAtLevel AtLevelLike,
		delimiter2 string,
		cited CitedLike,
	) CheckoutClauseLike
}

/*
CitedClassLike is a class interface that declares the
complete set of class constructors, constants and functions that must be
supported by each concrete cited-like class.
*/
type CitedClassLike interface {
	// Constructor Methods
	Cited(
		expression ExpressionLike,
	) CitedLike
}

/*
CollectionClassLike is a class interface that declares the
complete set of class constructors, constants and functions that must be
supported by each concrete collection-like class.
*/
type CollectionClassLike interface {
	// Constructor Methods
	Collection(
		any_ any,
	) CollectionLike
}

/*
ColonEqualClassLike is a class interface that declares the
complete set of class constructors, constants and functions that must be
supported by each concrete colon-equal-like class.
*/
type ColonEqualClassLike interface {
	// Constructor Methods
	ColonEqual(
		delimiter string,
	) ColonEqualLike
}

/*
CommentLineClassLike is a class interface that declares the
complete set of class constructors, constants and functions that must be
supported by each concrete comment-line-like class.
*/
type CommentLineClassLike interface {
	// Constructor Methods
	CommentLine(
		comment string,
	) CommentLineLike
}

/*
ComplementClassLike is a class interface that declares the
complete set of class constructors, constants and functions that must be
supported by each concrete complement-like class.
*/
type ComplementClassLike interface {
	// Constructor Methods
	Complement(
		delimiter string,
		logical LogicalLike,
	) ComplementLike
}

/*
ComponentClassLike is a class interface that declares the
complete set of class constructors, constants and functions that must be
supported by each concrete component-like class.
*/
type ComponentClassLike interface {
	// Constructor Methods
	Component(
		entity EntityLike,
		optionalParameters ParametersLike,
	) ComponentLike
}

/*
ConditionClassLike is a class interface that declares the
complete set of class constructors, constants and functions that must be
supported by each concrete condition-like class.
*/
type ConditionClassLike interface {
	// Constructor Methods
	Condition(
		expression ExpressionLike,
	) ConditionLike
}

/*
ContinueClauseClassLike is a class interface that declares the
complete set of class constructors, constants and functions that must be
supported by each concrete continue-clause-like class.
*/
type ContinueClauseClassLike interface {
	// Constructor Methods
	ContinueClause(
		delimiter1 string,
		delimiter2 string,
	) ContinueClauseLike
}

/*
DashEqualClassLike is a class interface that declares the
complete set of class constructors, constants and functions that must be
supported by each concrete dash-equal-like class.
*/
type DashEqualClassLike interface {
	// Constructor Methods
	DashEqual(
		delimiter string,
	) DashEqualLike
}

/*
DefaultEqualClassLike is a class interface that declares the
complete set of class constructors, constants and functions that must be
supported by each concrete default-equal-like class.
*/
type DefaultEqualClassLike interface {
	// Constructor Methods
	DefaultEqual(
		delimiter string,
	) DefaultEqualLike
}

/*
DiscardClauseClassLike is a class interface that declares the
complete set of class constructors, constants and functions that must be
supported by each concrete discard-clause-like class.
*/
type DiscardClauseClassLike interface {
	// Constructor Methods
	DiscardClause(
		delimiter string,
		draft DraftLike,
	) DiscardClauseLike
}

/*
DoClauseClassLike is a class interface that declares the
complete set of class constructors, constants and functions that must be
supported by each concrete do-clause-like class.
*/
type DoClauseClassLike interface {
	// Constructor Methods
	DoClause(
		delimiter string,
		invocation InvocationLike,
	) DoClauseLike
}

/*
DocumentClassLike is a class interface that declares the
complete set of class constructors, constants and functions that must be
supported by each concrete document-like class.
*/
type DocumentClassLike interface {
	// Constructor Methods
	Document(
		optionalLegalNotice LegalNoticeLike,
		component ComponentLike,
	) DocumentLike
}

/*
DraftClassLike is a class interface that declares the
complete set of class constructors, constants and functions that must be
supported by each concrete draft-like class.
*/
type DraftClassLike interface {
	// Constructor Methods
	Draft(
		expression ExpressionLike,
	) DraftLike
}

/*
ElementClassLike is a class interface that declares the
complete set of class constructors, constants and functions that must be
supported by each concrete element-like class.
*/
type ElementClassLike interface {
	// Constructor Methods
	Element(
		any_ any,
	) ElementLike
}

/*
EntityClassLike is a class interface that declares the
complete set of class constructors, constants and functions that must be
supported by each concrete entity-like class.
*/
type EntityClassLike interface {
	// Constructor Methods
	Entity(
		any_ any,
	) EntityLike
}

/*
EventClassLike is a class interface that declares the
complete set of class constructors, constants and functions that must be
supported by each concrete event-like class.
*/
type EventClassLike interface {
	// Constructor Methods
	Event(
		expression ExpressionLike,
	) EventLike
}

/*
ExceptionClassLike is a class interface that declares the
complete set of class constructors, constants and functions that must be
supported by each concrete exception-like class.
*/
type ExceptionClassLike interface {
	// Constructor Methods
	Exception(
		expression ExpressionLike,
	) ExceptionLike
}

/*
ExclusiveClassLike is a class interface that declares the
complete set of class constructors, constants and functions that must be
supported by each concrete exclusive-like class.
*/
type ExclusiveClassLike interface {
	// Constructor Methods
	Exclusive(
		delimiter string,
	) ExclusiveLike
}

/*
ExclusiveRangeClassLike is a class interface that declares the
complete set of class constructors, constants and functions that must be
supported by each concrete exclusive-range-like class.
*/
type ExclusiveRangeClassLike interface {
	// Constructor Methods
	ExclusiveRange(
		delimiter1 string,
		primitive1 PrimitiveLike,
		delimiter2 string,
		primitive2 PrimitiveLike,
		bracket BracketLike,
	) ExclusiveRangeLike
}

/*
ExpressionClassLike is a class interface that declares the
complete set of class constructors, constants and functions that must be
supported by each concrete expression-like class.
*/
type ExpressionClassLike interface {
	// Constructor Methods
	Expression(
		subject SubjectLike,
		predicates col.Sequential[PredicateLike],
	) ExpressionLike
}

/*
FailureClassLike is a class interface that declares the
complete set of class constructors, constants and functions that must be
supported by each concrete failure-like class.
*/
type FailureClassLike interface {
	// Constructor Methods
	Failure(
		symbol string,
	) FailureLike
}

/*
FlowClassLike is a class interface that declares the
complete set of class constructors, constants and functions that must be
supported by each concrete flow-like class.
*/
type FlowClassLike interface {
	// Constructor Methods
	Flow(
		any_ any,
	) FlowLike
}

/*
FunctionClassLike is a class interface that declares the
complete set of class constructors, constants and functions that must be
supported by each concrete function-like class.
*/
type FunctionClassLike interface {
	// Constructor Methods
	Function(
		identifier string,
		delimiter1 string,
		optionalArguments ArgumentsLike,
		delimiter2 string,
	) FunctionLike
}

/*
IfClauseClassLike is a class interface that declares the
complete set of class constructors, constants and functions that must be
supported by each concrete if-clause-like class.
*/
type IfClauseClassLike interface {
	// Constructor Methods
	IfClause(
		delimiter1 string,
		condition ConditionLike,
		delimiter2 string,
		procedure ProcedureLike,
	) IfClauseLike
}

/*
InclusiveClassLike is a class interface that declares the
complete set of class constructors, constants and functions that must be
supported by each concrete inclusive-like class.
*/
type InclusiveClassLike interface {
	// Constructor Methods
	Inclusive(
		delimiter string,
	) InclusiveLike
}

/*
InclusiveRangeClassLike is a class interface that declares the
complete set of class constructors, constants and functions that must be
supported by each concrete inclusive-range-like class.
*/
type InclusiveRangeClassLike interface {
	// Constructor Methods
	InclusiveRange(
		delimiter1 string,
		primitive1 PrimitiveLike,
		delimiter2 string,
		primitive2 PrimitiveLike,
		bracket BracketLike,
	) InclusiveRangeLike
}

/*
IndexClassLike is a class interface that declares the
complete set of class constructors, constants and functions that must be
supported by each concrete index-like class.
*/
type IndexClassLike interface {
	// Constructor Methods
	Index(
		expression ExpressionLike,
	) IndexLike
}

/*
IndicesClassLike is a class interface that declares the
complete set of class constructors, constants and functions that must be
supported by each concrete indices-like class.
*/
type IndicesClassLike interface {
	// Constructor Methods
	Indices(
		index IndexLike,
		additionalIndexes col.Sequential[AdditionalIndexLike],
	) IndicesLike
}

/*
IndirectClassLike is a class interface that declares the
complete set of class constructors, constants and functions that must be
supported by each concrete indirect-like class.
*/
type IndirectClassLike interface {
	// Constructor Methods
	Indirect(
		any_ any,
	) IndirectLike
}

/*
InductionClassLike is a class interface that declares the
complete set of class constructors, constants and functions that must be
supported by each concrete induction-like class.
*/
type InductionClassLike interface {
	// Constructor Methods
	Induction(
		any_ any,
	) InductionLike
}

/*
InlineAttributesClassLike is a class interface that declares the
complete set of class constructors, constants and functions that must be
supported by each concrete inline-attributes-like class.
*/
type InlineAttributesClassLike interface {
	// Constructor Methods
	InlineAttributes(
		delimiter1 string,
		association AssociationLike,
		additionalAssociations col.Sequential[AdditionalAssociationLike],
		delimiter2 string,
	) InlineAttributesLike
}

/*
InlineParametersClassLike is a class interface that declares the
complete set of class constructors, constants and functions that must be
supported by each concrete inline-parameters-like class.
*/
type InlineParametersClassLike interface {
	// Constructor Methods
	InlineParameters(
		delimiter1 string,
		association AssociationLike,
		additionalAssociations col.Sequential[AdditionalAssociationLike],
		delimiter2 string,
	) InlineParametersLike
}

/*
InlineStatementsClassLike is a class interface that declares the
complete set of class constructors, constants and functions that must be
supported by each concrete inline-statements-like class.
*/
type InlineStatementsClassLike interface {
	// Constructor Methods
	InlineStatements(
		delimiter1 string,
		statement StatementLike,
		additionalStatements col.Sequential[AdditionalStatementLike],
		delimiter2 string,
	) InlineStatementsLike
}

/*
InlineValuesClassLike is a class interface that declares the
complete set of class constructors, constants and functions that must be
supported by each concrete inline-values-like class.
*/
type InlineValuesClassLike interface {
	// Constructor Methods
	InlineValues(
		delimiter1 string,
		component ComponentLike,
		additionalValues col.Sequential[AdditionalValueLike],
		delimiter2 string,
	) InlineValuesLike
}

/*
InverseClassLike is a class interface that declares the
complete set of class constructors, constants and functions that must be
supported by each concrete inverse-like class.
*/
type InverseClassLike interface {
	// Constructor Methods
	Inverse(
		any_ any,
	) InverseLike
}

/*
InversionClassLike is a class interface that declares the
complete set of class constructors, constants and functions that must be
supported by each concrete inversion-like class.
*/
type InversionClassLike interface {
	// Constructor Methods
	Inversion(
		inverse InverseLike,
		numerical NumericalLike,
	) InversionLike
}

/*
InvocationClassLike is a class interface that declares the
complete set of class constructors, constants and functions that must be
supported by each concrete invocation-like class.
*/
type InvocationClassLike interface {
	// Constructor Methods
	Invocation(
		any_ any,
	) InvocationLike
}

/*
ItemClassLike is a class interface that declares the
complete set of class constructors, constants and functions that must be
supported by each concrete item-like class.
*/
type ItemClassLike interface {
	// Constructor Methods
	Item(
		symbol string,
	) ItemLike
}

/*
LegalNoticeClassLike is a class interface that declares the
complete set of class constructors, constants and functions that must be
supported by each concrete legal-notice-like class.
*/
type LegalNoticeClassLike interface {
	// Constructor Methods
	LegalNotice(
		comment string,
	) LegalNoticeLike
}

/*
LetClauseClassLike is a class interface that declares the
complete set of class constructors, constants and functions that must be
supported by each concrete let-clause-like class.
*/
type LetClauseClassLike interface {
	// Constructor Methods
	LetClause(
		delimiter string,
		recipient RecipientLike,
		assignment AssignmentLike,
		expression ExpressionLike,
	) LetClauseLike
}

/*
LogicalClassLike is a class interface that declares the
complete set of class constructors, constants and functions that must be
supported by each concrete logical-like class.
*/
type LogicalClassLike interface {
	// Constructor Methods
	Logical(
		any_ any,
	) LogicalLike
}

/*
MagnitudeClassLike is a class interface that declares the
complete set of class constructors, constants and functions that must be
supported by each concrete magnitude-like class.
*/
type MagnitudeClassLike interface {
	// Constructor Methods
	Magnitude(
		delimiter1 string,
		numerical NumericalLike,
		delimiter2 string,
	) MagnitudeLike
}

/*
MainClauseClassLike is a class interface that declares the
complete set of class constructors, constants and functions that must be
supported by each concrete main-clause-like class.
*/
type MainClauseClassLike interface {
	// Constructor Methods
	MainClause(
		any_ any,
	) MainClauseLike
}

/*
MatchHandlerClassLike is a class interface that declares the
complete set of class constructors, constants and functions that must be
supported by each concrete match-handler-like class.
*/
type MatchHandlerClassLike interface {
	// Constructor Methods
	MatchHandler(
		delimiter1 string,
		template TemplateLike,
		delimiter2 string,
		procedure ProcedureLike,
	) MatchHandlerLike
}

/*
MessageClassLike is a class interface that declares the
complete set of class constructors, constants and functions that must be
supported by each concrete message-like class.
*/
type MessageClassLike interface {
	// Constructor Methods
	Message(
		expression ExpressionLike,
	) MessageLike
}

/*
MessagingClassLike is a class interface that declares the
complete set of class constructors, constants and functions that must be
supported by each concrete messaging-like class.
*/
type MessagingClassLike interface {
	// Constructor Methods
	Messaging(
		any_ any,
	) MessagingLike
}

/*
MethodClassLike is a class interface that declares the
complete set of class constructors, constants and functions that must be
supported by each concrete method-like class.
*/
type MethodClassLike interface {
	// Constructor Methods
	Method(
		identifier1 string,
		blocking BlockingLike,
		identifier2 string,
		delimiter1 string,
		optionalArguments ArgumentsLike,
		delimiter2 string,
	) MethodLike
}

/*
MultilineAttributesClassLike is a class interface that declares the
complete set of class constructors, constants and functions that must be
supported by each concrete multiline-attributes-like class.
*/
type MultilineAttributesClassLike interface {
	// Constructor Methods
	MultilineAttributes(
		delimiter1 string,
		newline string,
		annotatedAssociations col.Sequential[AnnotatedAssociationLike],
		delimiter2 string,
	) MultilineAttributesLike
}

/*
MultilineParametersClassLike is a class interface that declares the
complete set of class constructors, constants and functions that must be
supported by each concrete multiline-parameters-like class.
*/
type MultilineParametersClassLike interface {
	// Constructor Methods
	MultilineParameters(
		delimiter1 string,
		newline string,
		annotatedAssociations col.Sequential[AnnotatedAssociationLike],
		delimiter2 string,
	) MultilineParametersLike
}

/*
MultilineStatementsClassLike is a class interface that declares the
complete set of class constructors, constants and functions that must be
supported by each concrete multiline-statements-like class.
*/
type MultilineStatementsClassLike interface {
	// Constructor Methods
	MultilineStatements(
		delimiter1 string,
		newline string,
		annotatedStatements col.Sequential[AnnotatedStatementLike],
		delimiter2 string,
	) MultilineStatementsLike
}

/*
MultilineValuesClassLike is a class interface that declares the
complete set of class constructors, constants and functions that must be
supported by each concrete multiline-values-like class.
*/
type MultilineValuesClassLike interface {
	// Constructor Methods
	MultilineValues(
		delimiter1 string,
		newline string,
		annotatedValues col.Sequential[AnnotatedValueLike],
		delimiter2 string,
	) MultilineValuesLike
}

/*
NoAttributesClassLike is a class interface that declares the
complete set of class constructors, constants and functions that must be
supported by each concrete no-attributes-like class.
*/
type NoAttributesClassLike interface {
	// Constructor Methods
	NoAttributes(
		delimiter1 string,
		delimiter2 string,
		delimiter3 string,
	) NoAttributesLike
}

/*
NoStatementsClassLike is a class interface that declares the
complete set of class constructors, constants and functions that must be
supported by each concrete no-statements-like class.
*/
type NoStatementsClassLike interface {
	// Constructor Methods
	NoStatements(
		delimiter1 string,
		delimiter2 string,
	) NoStatementsLike
}

/*
NoValuesClassLike is a class interface that declares the
complete set of class constructors, constants and functions that must be
supported by each concrete no-values-like class.
*/
type NoValuesClassLike interface {
	// Constructor Methods
	NoValues(
		delimiter1 string,
		delimiter2 string,
	) NoValuesLike
}

/*
NotarizeClauseClassLike is a class interface that declares the
complete set of class constructors, constants and functions that must be
supported by each concrete notarize-clause-like class.
*/
type NotarizeClauseClassLike interface {
	// Constructor Methods
	NotarizeClause(
		delimiter1 string,
		draft DraftLike,
		delimiter2 string,
		cited CitedLike,
	) NotarizeClauseLike
}

/*
NumericalClassLike is a class interface that declares the
complete set of class constructors, constants and functions that must be
supported by each concrete numerical-like class.
*/
type NumericalClassLike interface {
	// Constructor Methods
	Numerical(
		any_ any,
	) NumericalLike
}

/*
OnClauseClassLike is a class interface that declares the
complete set of class constructors, constants and functions that must be
supported by each concrete on-clause-like class.
*/
type OnClauseClassLike interface {
	// Constructor Methods
	OnClause(
		delimiter string,
		failure FailureLike,
		matchHandlers col.Sequential[MatchHandlerLike],
	) OnClauseLike
}

/*
OperationClassLike is a class interface that declares the
complete set of class constructors, constants and functions that must be
supported by each concrete operation-like class.
*/
type OperationClassLike interface {
	// Constructor Methods
	Operation(
		any_ any,
	) OperationLike
}

/*
ParametersClassLike is a class interface that declares the
complete set of class constructors, constants and functions that must be
supported by each concrete parameters-like class.
*/
type ParametersClassLike interface {
	// Constructor Methods
	Parameters(
		any_ any,
	) ParametersLike
}

/*
PlusEqualClassLike is a class interface that declares the
complete set of class constructors, constants and functions that must be
supported by each concrete plus-equal-like class.
*/
type PlusEqualClassLike interface {
	// Constructor Methods
	PlusEqual(
		delimiter string,
	) PlusEqualLike
}

/*
PostClauseClassLike is a class interface that declares the
complete set of class constructors, constants and functions that must be
supported by each concrete post-clause-like class.
*/
type PostClauseClassLike interface {
	// Constructor Methods
	PostClause(
		delimiter1 string,
		message MessageLike,
		delimiter2 string,
		bag BagLike,
	) PostClauseLike
}

/*
PrecedenceClassLike is a class interface that declares the
complete set of class constructors, constants and functions that must be
supported by each concrete precedence-like class.
*/
type PrecedenceClassLike interface {
	// Constructor Methods
	Precedence(
		delimiter1 string,
		expression ExpressionLike,
		delimiter2 string,
	) PrecedenceLike
}

/*
PredicateClassLike is a class interface that declares the
complete set of class constructors, constants and functions that must be
supported by each concrete predicate-like class.
*/
type PredicateClassLike interface {
	// Constructor Methods
	Predicate(
		operation OperationLike,
		expression ExpressionLike,
	) PredicateLike
}

/*
PrimitiveClassLike is a class interface that declares the
complete set of class constructors, constants and functions that must be
supported by each concrete primitive-like class.
*/
type PrimitiveClassLike interface {
	// Constructor Methods
	Primitive(
		any_ any,
	) PrimitiveLike
}

/*
ProcedureClassLike is a class interface that declares the
complete set of class constructors, constants and functions that must be
supported by each concrete procedure-like class.
*/
type ProcedureClassLike interface {
	// Constructor Methods
	Procedure(
		any_ any,
	) ProcedureLike
}

/*
PublishClauseClassLike is a class interface that declares the
complete set of class constructors, constants and functions that must be
supported by each concrete publish-clause-like class.
*/
type PublishClauseClassLike interface {
	// Constructor Methods
	PublishClause(
		delimiter string,
		event EventLike,
	) PublishClauseLike
}

/*
RecipientClassLike is a class interface that declares the
complete set of class constructors, constants and functions that must be
supported by each concrete recipient-like class.
*/
type RecipientClassLike interface {
	// Constructor Methods
	Recipient(
		any_ any,
	) RecipientLike
}

/*
ReferentClassLike is a class interface that declares the
complete set of class constructors, constants and functions that must be
supported by each concrete referent-like class.
*/
type ReferentClassLike interface {
	// Constructor Methods
	Referent(
		delimiter string,
		indirect IndirectLike,
	) ReferentLike
}

/*
RejectClauseClassLike is a class interface that declares the
complete set of class constructors, constants and functions that must be
supported by each concrete reject-clause-like class.
*/
type RejectClauseClassLike interface {
	// Constructor Methods
	RejectClause(
		delimiter string,
		message MessageLike,
	) RejectClauseLike
}

/*
RepositoryClassLike is a class interface that declares the
complete set of class constructors, constants and functions that must be
supported by each concrete repository-like class.
*/
type RepositoryClassLike interface {
	// Constructor Methods
	Repository(
		any_ any,
	) RepositoryLike
}

/*
ResultClassLike is a class interface that declares the
complete set of class constructors, constants and functions that must be
supported by each concrete result-like class.
*/
type ResultClassLike interface {
	// Constructor Methods
	Result(
		expression ExpressionLike,
	) ResultLike
}

/*
RetrieveClauseClassLike is a class interface that declares the
complete set of class constructors, constants and functions that must be
supported by each concrete retrieve-clause-like class.
*/
type RetrieveClauseClassLike interface {
	// Constructor Methods
	RetrieveClause(
		delimiter1 string,
		recipient RecipientLike,
		delimiter2 string,
		bag BagLike,
	) RetrieveClauseLike
}

/*
ReturnClauseClassLike is a class interface that declares the
complete set of class constructors, constants and functions that must be
supported by each concrete return-clause-like class.
*/
type ReturnClauseClassLike interface {
	// Constructor Methods
	ReturnClause(
		delimiter string,
		result ResultLike,
	) ReturnClauseLike
}

/*
SaveClauseClassLike is a class interface that declares the
complete set of class constructors, constants and functions that must be
supported by each concrete save-clause-like class.
*/
type SaveClauseClassLike interface {
	// Constructor Methods
	SaveClause(
		delimiter1 string,
		draft DraftLike,
		delimiter2 string,
		cited CitedLike,
	) SaveClauseLike
}

/*
SelectClauseClassLike is a class interface that declares the
complete set of class constructors, constants and functions that must be
supported by each concrete select-clause-like class.
*/
type SelectClauseClassLike interface {
	// Constructor Methods
	SelectClause(
		delimiter string,
		target TargetLike,
		matchHandlers col.Sequential[MatchHandlerLike],
	) SelectClauseLike
}

/*
SequenceClassLike is a class interface that declares the
complete set of class constructors, constants and functions that must be
supported by each concrete sequence-like class.
*/
type SequenceClassLike interface {
	// Constructor Methods
	Sequence(
		expression ExpressionLike,
	) SequenceLike
}

/*
SlashEqualClassLike is a class interface that declares the
complete set of class constructors, constants and functions that must be
supported by each concrete slash-equal-like class.
*/
type SlashEqualClassLike interface {
	// Constructor Methods
	SlashEqual(
		delimiter string,
	) SlashEqualLike
}

/*
StarEqualClassLike is a class interface that declares the
complete set of class constructors, constants and functions that must be
supported by each concrete star-equal-like class.
*/
type StarEqualClassLike interface {
	// Constructor Methods
	StarEqual(
		delimiter string,
	) StarEqualLike
}

/*
StatementClassLike is a class interface that declares the
complete set of class constructors, constants and functions that must be
supported by each concrete statement-like class.
*/
type StatementClassLike interface {
	// Constructor Methods
	Statement(
		mainClause MainClauseLike,
		optionalOnClause OnClauseLike,
	) StatementLike
}

/*
StatementLineClassLike is a class interface that declares the
complete set of class constructors, constants and functions that must be
supported by each concrete statement-line-like class.
*/
type StatementLineClassLike interface {
	// Constructor Methods
	StatementLine(
		statement StatementLike,
		optionalNote string,
	) StatementLineLike
}

/*
StringClassLike is a class interface that declares the
complete set of class constructors, constants and functions that must be
supported by each concrete string-like class.
*/
type StringClassLike interface {
	// Constructor Methods
	String(
		any_ any,
	) StringLike
}

/*
SubcomponentClassLike is a class interface that declares the
complete set of class constructors, constants and functions that must be
supported by each concrete subcomponent-like class.
*/
type SubcomponentClassLike interface {
	// Constructor Methods
	Subcomponent(
		identifier string,
		delimiter1 string,
		indices IndicesLike,
		delimiter2 string,
	) SubcomponentLike
}

/*
SubjectClassLike is a class interface that declares the
complete set of class constructors, constants and functions that must be
supported by each concrete subject-like class.
*/
type SubjectClassLike interface {
	// Constructor Methods
	Subject(
		any_ any,
	) SubjectLike
}

/*
TargetClassLike is a class interface that declares the
complete set of class constructors, constants and functions that must be
supported by each concrete target-like class.
*/
type TargetClassLike interface {
	// Constructor Methods
	Target(
		any_ any,
	) TargetLike
}

/*
TemplateClassLike is a class interface that declares the
complete set of class constructors, constants and functions that must be
supported by each concrete template-like class.
*/
type TemplateClassLike interface {
	// Constructor Methods
	Template(
		expression ExpressionLike,
	) TemplateLike
}

/*
ThrowClauseClassLike is a class interface that declares the
complete set of class constructors, constants and functions that must be
supported by each concrete throw-clause-like class.
*/
type ThrowClauseClassLike interface {
	// Constructor Methods
	ThrowClause(
		delimiter string,
		exception ExceptionLike,
	) ThrowClauseLike
}

/*
VariableClassLike is a class interface that declares the
complete set of class constructors, constants and functions that must be
supported by each concrete variable-like class.
*/
type VariableClassLike interface {
	// Constructor Methods
	Variable(
		identifier string,
	) VariableLike
}

/*
WhileClauseClassLike is a class interface that declares the
complete set of class constructors, constants and functions that must be
supported by each concrete while-clause-like class.
*/
type WhileClauseClassLike interface {
	// Constructor Methods
	WhileClause(
		delimiter1 string,
		condition ConditionLike,
		delimiter2 string,
		procedure ProcedureLike,
	) WhileClauseLike
}

/*
WithClauseClassLike is a class interface that declares the
complete set of class constructors, constants and functions that must be
supported by each concrete with-clause-like class.
*/
type WithClauseClassLike interface {
	// Constructor Methods
	WithClause(
		delimiter1 string,
		delimiter2 string,
		item ItemLike,
		delimiter3 string,
		sequence SequenceLike,
		delimiter4 string,
		procedure ProcedureLike,
	) WithClauseLike
}

// INSTANCE DECLARATIONS

/*
AcceptClauseLike is an instance interface that declares the
complete set of principal, attribute and aspect methods that must be supported
by each instance of a concrete accept-clause-like class.
*/
type AcceptClauseLike interface {
	// Principal Methods
	GetClass() AcceptClauseClassLike

	// Attribute Methods
	GetDelimiter() string
	GetMessage() MessageLike
}

/*
AdditionalArgumentLike is an instance interface that declares the
complete set of principal, attribute and aspect methods that must be supported
by each instance of a concrete additional-argument-like class.
*/
type AdditionalArgumentLike interface {
	// Principal Methods
	GetClass() AdditionalArgumentClassLike

	// Attribute Methods
	GetDelimiter() string
	GetArgument() ArgumentLike
}

/*
AdditionalAssociationLike is an instance interface that declares the
complete set of principal, attribute and aspect methods that must be supported
by each instance of a concrete additional-association-like class.
*/
type AdditionalAssociationLike interface {
	// Principal Methods
	GetClass() AdditionalAssociationClassLike

	// Attribute Methods
	GetDelimiter() string
	GetAssociation() AssociationLike
}

/*
AdditionalIndexLike is an instance interface that declares the
complete set of principal, attribute and aspect methods that must be supported
by each instance of a concrete additional-index-like class.
*/
type AdditionalIndexLike interface {
	// Principal Methods
	GetClass() AdditionalIndexClassLike

	// Attribute Methods
	GetDelimiter() string
	GetIndex() IndexLike
}

/*
AdditionalStatementLike is an instance interface that declares the
complete set of principal, attribute and aspect methods that must be supported
by each instance of a concrete additional-statement-like class.
*/
type AdditionalStatementLike interface {
	// Principal Methods
	GetClass() AdditionalStatementClassLike

	// Attribute Methods
	GetDelimiter() string
	GetStatement() StatementLike
}

/*
AdditionalValueLike is an instance interface that declares the
complete set of principal, attribute and aspect methods that must be supported
by each instance of a concrete additional-value-like class.
*/
type AdditionalValueLike interface {
	// Principal Methods
	GetClass() AdditionalValueClassLike

	// Attribute Methods
	GetDelimiter() string
	GetComponent() ComponentLike
}

/*
AnnotatedAssociationLike is an instance interface that declares the
complete set of principal, attribute and aspect methods that must be supported
by each instance of a concrete annotated-association-like class.
*/
type AnnotatedAssociationLike interface {
	// Principal Methods
	GetClass() AnnotatedAssociationClassLike

	// Attribute Methods
	GetAssociation() AssociationLike
	GetOptionalNote() string
}

/*
AnnotatedStatementLike is an instance interface that declares the
complete set of principal, attribute and aspect methods that must be supported
by each instance of a concrete annotated-statement-like class.
*/
type AnnotatedStatementLike interface {
	// Principal Methods
	GetClass() AnnotatedStatementClassLike

	// Attribute Methods
	GetAny() any
}

/*
AnnotatedValueLike is an instance interface that declares the
complete set of principal, attribute and aspect methods that must be supported
by each instance of a concrete annotated-value-like class.
*/
type AnnotatedValueLike interface {
	// Principal Methods
	GetClass() AnnotatedValueClassLike

	// Attribute Methods
	GetComponent() ComponentLike
	GetOptionalNote() string
}

/*
ArgumentLike is an instance interface that declares the
complete set of principal, attribute and aspect methods that must be supported
by each instance of a concrete argument-like class.
*/
type ArgumentLike interface {
	// Principal Methods
	GetClass() ArgumentClassLike

	// Attribute Methods
	GetIdentifier() string
}

/*
ArgumentsLike is an instance interface that declares the
complete set of principal, attribute and aspect methods that must be supported
by each instance of a concrete arguments-like class.
*/
type ArgumentsLike interface {
	// Principal Methods
	GetClass() ArgumentsClassLike

	// Attribute Methods
	GetArgument() ArgumentLike
	GetAdditionalArguments() col.Sequential[AdditionalArgumentLike]
}

/*
AssignmentLike is an instance interface that declares the
complete set of principal, attribute and aspect methods that must be supported
by each instance of a concrete assignment-like class.
*/
type AssignmentLike interface {
	// Principal Methods
	GetClass() AssignmentClassLike

	// Attribute Methods
	GetAny() any
}

/*
AssociationLike is an instance interface that declares the
complete set of principal, attribute and aspect methods that must be supported
by each instance of a concrete association-like class.
*/
type AssociationLike interface {
	// Principal Methods
	GetClass() AssociationClassLike

	// Attribute Methods
	GetPrimitive() PrimitiveLike
	GetDelimiter() string
	GetComponent() ComponentLike
}

/*
AtLevelLike is an instance interface that declares the
complete set of principal, attribute and aspect methods that must be supported
by each instance of a concrete at-level-like class.
*/
type AtLevelLike interface {
	// Principal Methods
	GetClass() AtLevelClassLike

	// Attribute Methods
	GetDelimiter1() string
	GetDelimiter2() string
	GetExpression() ExpressionLike
}

/*
BagLike is an instance interface that declares the
complete set of principal, attribute and aspect methods that must be supported
by each instance of a concrete bag-like class.
*/
type BagLike interface {
	// Principal Methods
	GetClass() BagClassLike

	// Attribute Methods
	GetExpression() ExpressionLike
}

/*
BlockingLike is an instance interface that declares the
complete set of principal, attribute and aspect methods that must be supported
by each instance of a concrete blocking-like class.
*/
type BlockingLike interface {
	// Principal Methods
	GetClass() BlockingClassLike

	// Attribute Methods
	GetAny() any
}

/*
BracketLike is an instance interface that declares the
complete set of principal, attribute and aspect methods that must be supported
by each instance of a concrete bracket-like class.
*/
type BracketLike interface {
	// Principal Methods
	GetClass() BracketClassLike

	// Attribute Methods
	GetAny() any
}

/*
BreakClauseLike is an instance interface that declares the
complete set of principal, attribute and aspect methods that must be supported
by each instance of a concrete break-clause-like class.
*/
type BreakClauseLike interface {
	// Principal Methods
	GetClass() BreakClauseClassLike

	// Attribute Methods
	GetDelimiter1() string
	GetDelimiter2() string
}

/*
CheckoutClauseLike is an instance interface that declares the
complete set of principal, attribute and aspect methods that must be supported
by each instance of a concrete checkout-clause-like class.
*/
type CheckoutClauseLike interface {
	// Principal Methods
	GetClass() CheckoutClauseClassLike

	// Attribute Methods
	GetDelimiter1() string
	GetRecipient() RecipientLike
	GetOptionalAtLevel() AtLevelLike
	GetDelimiter2() string
	GetCited() CitedLike
}

/*
CitedLike is an instance interface that declares the
complete set of principal, attribute and aspect methods that must be supported
by each instance of a concrete cited-like class.
*/
type CitedLike interface {
	// Principal Methods
	GetClass() CitedClassLike

	// Attribute Methods
	GetExpression() ExpressionLike
}

/*
CollectionLike is an instance interface that declares the
complete set of principal, attribute and aspect methods that must be supported
by each instance of a concrete collection-like class.
*/
type CollectionLike interface {
	// Principal Methods
	GetClass() CollectionClassLike

	// Attribute Methods
	GetAny() any
}

/*
ColonEqualLike is an instance interface that declares the
complete set of principal, attribute and aspect methods that must be supported
by each instance of a concrete colon-equal-like class.
*/
type ColonEqualLike interface {
	// Principal Methods
	GetClass() ColonEqualClassLike

	// Attribute Methods
	GetDelimiter() string
}

/*
CommentLineLike is an instance interface that declares the
complete set of principal, attribute and aspect methods that must be supported
by each instance of a concrete comment-line-like class.
*/
type CommentLineLike interface {
	// Principal Methods
	GetClass() CommentLineClassLike

	// Attribute Methods
	GetComment() string
}

/*
ComplementLike is an instance interface that declares the
complete set of principal, attribute and aspect methods that must be supported
by each instance of a concrete complement-like class.
*/
type ComplementLike interface {
	// Principal Methods
	GetClass() ComplementClassLike

	// Attribute Methods
	GetDelimiter() string
	GetLogical() LogicalLike
}

/*
ComponentLike is an instance interface that declares the
complete set of principal, attribute and aspect methods that must be supported
by each instance of a concrete component-like class.
*/
type ComponentLike interface {
	// Principal Methods
	GetClass() ComponentClassLike

	// Attribute Methods
	GetEntity() EntityLike
	GetOptionalParameters() ParametersLike
}

/*
ConditionLike is an instance interface that declares the
complete set of principal, attribute and aspect methods that must be supported
by each instance of a concrete condition-like class.
*/
type ConditionLike interface {
	// Principal Methods
	GetClass() ConditionClassLike

	// Attribute Methods
	GetExpression() ExpressionLike
}

/*
ContinueClauseLike is an instance interface that declares the
complete set of principal, attribute and aspect methods that must be supported
by each instance of a concrete continue-clause-like class.
*/
type ContinueClauseLike interface {
	// Principal Methods
	GetClass() ContinueClauseClassLike

	// Attribute Methods
	GetDelimiter1() string
	GetDelimiter2() string
}

/*
DashEqualLike is an instance interface that declares the
complete set of principal, attribute and aspect methods that must be supported
by each instance of a concrete dash-equal-like class.
*/
type DashEqualLike interface {
	// Principal Methods
	GetClass() DashEqualClassLike

	// Attribute Methods
	GetDelimiter() string
}

/*
DefaultEqualLike is an instance interface that declares the
complete set of principal, attribute and aspect methods that must be supported
by each instance of a concrete default-equal-like class.
*/
type DefaultEqualLike interface {
	// Principal Methods
	GetClass() DefaultEqualClassLike

	// Attribute Methods
	GetDelimiter() string
}

/*
DiscardClauseLike is an instance interface that declares the
complete set of principal, attribute and aspect methods that must be supported
by each instance of a concrete discard-clause-like class.
*/
type DiscardClauseLike interface {
	// Principal Methods
	GetClass() DiscardClauseClassLike

	// Attribute Methods
	GetDelimiter() string
	GetDraft() DraftLike
}

/*
DoClauseLike is an instance interface that declares the
complete set of principal, attribute and aspect methods that must be supported
by each instance of a concrete do-clause-like class.
*/
type DoClauseLike interface {
	// Principal Methods
	GetClass() DoClauseClassLike

	// Attribute Methods
	GetDelimiter() string
	GetInvocation() InvocationLike
}

/*
DocumentLike is an instance interface that declares the
complete set of principal, attribute and aspect methods that must be supported
by each instance of a concrete document-like class.
*/
type DocumentLike interface {
	// Principal Methods
	GetClass() DocumentClassLike

	// Attribute Methods
	GetOptionalLegalNotice() LegalNoticeLike
	GetComponent() ComponentLike
}

/*
DraftLike is an instance interface that declares the
complete set of principal, attribute and aspect methods that must be supported
by each instance of a concrete draft-like class.
*/
type DraftLike interface {
	// Principal Methods
	GetClass() DraftClassLike

	// Attribute Methods
	GetExpression() ExpressionLike
}

/*
ElementLike is an instance interface that declares the
complete set of principal, attribute and aspect methods that must be supported
by each instance of a concrete element-like class.
*/
type ElementLike interface {
	// Principal Methods
	GetClass() ElementClassLike

	// Attribute Methods
	GetAny() any
}

/*
EntityLike is an instance interface that declares the
complete set of principal, attribute and aspect methods that must be supported
by each instance of a concrete entity-like class.
*/
type EntityLike interface {
	// Principal Methods
	GetClass() EntityClassLike

	// Attribute Methods
	GetAny() any
}

/*
EventLike is an instance interface that declares the
complete set of principal, attribute and aspect methods that must be supported
by each instance of a concrete event-like class.
*/
type EventLike interface {
	// Principal Methods
	GetClass() EventClassLike

	// Attribute Methods
	GetExpression() ExpressionLike
}

/*
ExceptionLike is an instance interface that declares the
complete set of principal, attribute and aspect methods that must be supported
by each instance of a concrete exception-like class.
*/
type ExceptionLike interface {
	// Principal Methods
	GetClass() ExceptionClassLike

	// Attribute Methods
	GetExpression() ExpressionLike
}

/*
ExclusiveLike is an instance interface that declares the
complete set of principal, attribute and aspect methods that must be supported
by each instance of a concrete exclusive-like class.
*/
type ExclusiveLike interface {
	// Principal Methods
	GetClass() ExclusiveClassLike

	// Attribute Methods
	GetDelimiter() string
}

/*
ExclusiveRangeLike is an instance interface that declares the
complete set of principal, attribute and aspect methods that must be supported
by each instance of a concrete exclusive-range-like class.
*/
type ExclusiveRangeLike interface {
	// Principal Methods
	GetClass() ExclusiveRangeClassLike

	// Attribute Methods
	GetDelimiter1() string
	GetPrimitive1() PrimitiveLike
	GetDelimiter2() string
	GetPrimitive2() PrimitiveLike
	GetBracket() BracketLike
}

/*
ExpressionLike is an instance interface that declares the
complete set of principal, attribute and aspect methods that must be supported
by each instance of a concrete expression-like class.
*/
type ExpressionLike interface {
	// Principal Methods
	GetClass() ExpressionClassLike

	// Attribute Methods
	GetSubject() SubjectLike
	GetPredicates() col.Sequential[PredicateLike]
}

/*
FailureLike is an instance interface that declares the
complete set of principal, attribute and aspect methods that must be supported
by each instance of a concrete failure-like class.
*/
type FailureLike interface {
	// Principal Methods
	GetClass() FailureClassLike

	// Attribute Methods
	GetSymbol() string
}

/*
FlowLike is an instance interface that declares the
complete set of principal, attribute and aspect methods that must be supported
by each instance of a concrete flow-like class.
*/
type FlowLike interface {
	// Principal Methods
	GetClass() FlowClassLike

	// Attribute Methods
	GetAny() any
}

/*
FunctionLike is an instance interface that declares the
complete set of principal, attribute and aspect methods that must be supported
by each instance of a concrete function-like class.
*/
type FunctionLike interface {
	// Principal Methods
	GetClass() FunctionClassLike

	// Attribute Methods
	GetIdentifier() string
	GetDelimiter1() string
	GetOptionalArguments() ArgumentsLike
	GetDelimiter2() string
}

/*
IfClauseLike is an instance interface that declares the
complete set of principal, attribute and aspect methods that must be supported
by each instance of a concrete if-clause-like class.
*/
type IfClauseLike interface {
	// Principal Methods
	GetClass() IfClauseClassLike

	// Attribute Methods
	GetDelimiter1() string
	GetCondition() ConditionLike
	GetDelimiter2() string
	GetProcedure() ProcedureLike
}

/*
InclusiveLike is an instance interface that declares the
complete set of principal, attribute and aspect methods that must be supported
by each instance of a concrete inclusive-like class.
*/
type InclusiveLike interface {
	// Principal Methods
	GetClass() InclusiveClassLike

	// Attribute Methods
	GetDelimiter() string
}

/*
InclusiveRangeLike is an instance interface that declares the
complete set of principal, attribute and aspect methods that must be supported
by each instance of a concrete inclusive-range-like class.
*/
type InclusiveRangeLike interface {
	// Principal Methods
	GetClass() InclusiveRangeClassLike

	// Attribute Methods
	GetDelimiter1() string
	GetPrimitive1() PrimitiveLike
	GetDelimiter2() string
	GetPrimitive2() PrimitiveLike
	GetBracket() BracketLike
}

/*
IndexLike is an instance interface that declares the
complete set of principal, attribute and aspect methods that must be supported
by each instance of a concrete index-like class.
*/
type IndexLike interface {
	// Principal Methods
	GetClass() IndexClassLike

	// Attribute Methods
	GetExpression() ExpressionLike
}

/*
IndicesLike is an instance interface that declares the
complete set of principal, attribute and aspect methods that must be supported
by each instance of a concrete indices-like class.
*/
type IndicesLike interface {
	// Principal Methods
	GetClass() IndicesClassLike

	// Attribute Methods
	GetIndex() IndexLike
	GetAdditionalIndexes() col.Sequential[AdditionalIndexLike]
}

/*
IndirectLike is an instance interface that declares the
complete set of principal, attribute and aspect methods that must be supported
by each instance of a concrete indirect-like class.
*/
type IndirectLike interface {
	// Principal Methods
	GetClass() IndirectClassLike

	// Attribute Methods
	GetAny() any
}

/*
InductionLike is an instance interface that declares the
complete set of principal, attribute and aspect methods that must be supported
by each instance of a concrete induction-like class.
*/
type InductionLike interface {
	// Principal Methods
	GetClass() InductionClassLike

	// Attribute Methods
	GetAny() any
}

/*
InlineAttributesLike is an instance interface that declares the
complete set of principal, attribute and aspect methods that must be supported
by each instance of a concrete inline-attributes-like class.
*/
type InlineAttributesLike interface {
	// Principal Methods
	GetClass() InlineAttributesClassLike

	// Attribute Methods
	GetDelimiter1() string
	GetAssociation() AssociationLike
	GetAdditionalAssociations() col.Sequential[AdditionalAssociationLike]
	GetDelimiter2() string
}

/*
InlineParametersLike is an instance interface that declares the
complete set of principal, attribute and aspect methods that must be supported
by each instance of a concrete inline-parameters-like class.
*/
type InlineParametersLike interface {
	// Principal Methods
	GetClass() InlineParametersClassLike

	// Attribute Methods
	GetDelimiter1() string
	GetAssociation() AssociationLike
	GetAdditionalAssociations() col.Sequential[AdditionalAssociationLike]
	GetDelimiter2() string
}

/*
InlineStatementsLike is an instance interface that declares the
complete set of principal, attribute and aspect methods that must be supported
by each instance of a concrete inline-statements-like class.
*/
type InlineStatementsLike interface {
	// Principal Methods
	GetClass() InlineStatementsClassLike

	// Attribute Methods
	GetDelimiter1() string
	GetStatement() StatementLike
	GetAdditionalStatements() col.Sequential[AdditionalStatementLike]
	GetDelimiter2() string
}

/*
InlineValuesLike is an instance interface that declares the
complete set of principal, attribute and aspect methods that must be supported
by each instance of a concrete inline-values-like class.
*/
type InlineValuesLike interface {
	// Principal Methods
	GetClass() InlineValuesClassLike

	// Attribute Methods
	GetDelimiter1() string
	GetComponent() ComponentLike
	GetAdditionalValues() col.Sequential[AdditionalValueLike]
	GetDelimiter2() string
}

/*
InverseLike is an instance interface that declares the
complete set of principal, attribute and aspect methods that must be supported
by each instance of a concrete inverse-like class.
*/
type InverseLike interface {
	// Principal Methods
	GetClass() InverseClassLike

	// Attribute Methods
	GetAny() any
}

/*
InversionLike is an instance interface that declares the
complete set of principal, attribute and aspect methods that must be supported
by each instance of a concrete inversion-like class.
*/
type InversionLike interface {
	// Principal Methods
	GetClass() InversionClassLike

	// Attribute Methods
	GetInverse() InverseLike
	GetNumerical() NumericalLike
}

/*
InvocationLike is an instance interface that declares the
complete set of principal, attribute and aspect methods that must be supported
by each instance of a concrete invocation-like class.
*/
type InvocationLike interface {
	// Principal Methods
	GetClass() InvocationClassLike

	// Attribute Methods
	GetAny() any
}

/*
ItemLike is an instance interface that declares the
complete set of principal, attribute and aspect methods that must be supported
by each instance of a concrete item-like class.
*/
type ItemLike interface {
	// Principal Methods
	GetClass() ItemClassLike

	// Attribute Methods
	GetSymbol() string
}

/*
LegalNoticeLike is an instance interface that declares the
complete set of principal, attribute and aspect methods that must be supported
by each instance of a concrete legal-notice-like class.
*/
type LegalNoticeLike interface {
	// Principal Methods
	GetClass() LegalNoticeClassLike

	// Attribute Methods
	GetComment() string
}

/*
LetClauseLike is an instance interface that declares the
complete set of principal, attribute and aspect methods that must be supported
by each instance of a concrete let-clause-like class.
*/
type LetClauseLike interface {
	// Principal Methods
	GetClass() LetClauseClassLike

	// Attribute Methods
	GetDelimiter() string
	GetRecipient() RecipientLike
	GetAssignment() AssignmentLike
	GetExpression() ExpressionLike
}

/*
LogicalLike is an instance interface that declares the
complete set of principal, attribute and aspect methods that must be supported
by each instance of a concrete logical-like class.
*/
type LogicalLike interface {
	// Principal Methods
	GetClass() LogicalClassLike

	// Attribute Methods
	GetAny() any
}

/*
MagnitudeLike is an instance interface that declares the
complete set of principal, attribute and aspect methods that must be supported
by each instance of a concrete magnitude-like class.
*/
type MagnitudeLike interface {
	// Principal Methods
	GetClass() MagnitudeClassLike

	// Attribute Methods
	GetDelimiter1() string
	GetNumerical() NumericalLike
	GetDelimiter2() string
}

/*
MainClauseLike is an instance interface that declares the
complete set of principal, attribute and aspect methods that must be supported
by each instance of a concrete main-clause-like class.
*/
type MainClauseLike interface {
	// Principal Methods
	GetClass() MainClauseClassLike

	// Attribute Methods
	GetAny() any
}

/*
MatchHandlerLike is an instance interface that declares the
complete set of principal, attribute and aspect methods that must be supported
by each instance of a concrete match-handler-like class.
*/
type MatchHandlerLike interface {
	// Principal Methods
	GetClass() MatchHandlerClassLike

	// Attribute Methods
	GetDelimiter1() string
	GetTemplate() TemplateLike
	GetDelimiter2() string
	GetProcedure() ProcedureLike
}

/*
MessageLike is an instance interface that declares the
complete set of principal, attribute and aspect methods that must be supported
by each instance of a concrete message-like class.
*/
type MessageLike interface {
	// Principal Methods
	GetClass() MessageClassLike

	// Attribute Methods
	GetExpression() ExpressionLike
}

/*
MessagingLike is an instance interface that declares the
complete set of principal, attribute and aspect methods that must be supported
by each instance of a concrete messaging-like class.
*/
type MessagingLike interface {
	// Principal Methods
	GetClass() MessagingClassLike

	// Attribute Methods
	GetAny() any
}

/*
MethodLike is an instance interface that declares the
complete set of principal, attribute and aspect methods that must be supported
by each instance of a concrete method-like class.
*/
type MethodLike interface {
	// Principal Methods
	GetClass() MethodClassLike

	// Attribute Methods
	GetIdentifier1() string
	GetBlocking() BlockingLike
	GetIdentifier2() string
	GetDelimiter1() string
	GetOptionalArguments() ArgumentsLike
	GetDelimiter2() string
}

/*
MultilineAttributesLike is an instance interface that declares the
complete set of principal, attribute and aspect methods that must be supported
by each instance of a concrete multiline-attributes-like class.
*/
type MultilineAttributesLike interface {
	// Principal Methods
	GetClass() MultilineAttributesClassLike

	// Attribute Methods
	GetDelimiter1() string
	GetNewline() string
	GetAnnotatedAssociations() col.Sequential[AnnotatedAssociationLike]
	GetDelimiter2() string
}

/*
MultilineParametersLike is an instance interface that declares the
complete set of principal, attribute and aspect methods that must be supported
by each instance of a concrete multiline-parameters-like class.
*/
type MultilineParametersLike interface {
	// Principal Methods
	GetClass() MultilineParametersClassLike

	// Attribute Methods
	GetDelimiter1() string
	GetNewline() string
	GetAnnotatedAssociations() col.Sequential[AnnotatedAssociationLike]
	GetDelimiter2() string
}

/*
MultilineStatementsLike is an instance interface that declares the
complete set of principal, attribute and aspect methods that must be supported
by each instance of a concrete multiline-statements-like class.
*/
type MultilineStatementsLike interface {
	// Principal Methods
	GetClass() MultilineStatementsClassLike

	// Attribute Methods
	GetDelimiter1() string
	GetNewline() string
	GetAnnotatedStatements() col.Sequential[AnnotatedStatementLike]
	GetDelimiter2() string
}

/*
MultilineValuesLike is an instance interface that declares the
complete set of principal, attribute and aspect methods that must be supported
by each instance of a concrete multiline-values-like class.
*/
type MultilineValuesLike interface {
	// Principal Methods
	GetClass() MultilineValuesClassLike

	// Attribute Methods
	GetDelimiter1() string
	GetNewline() string
	GetAnnotatedValues() col.Sequential[AnnotatedValueLike]
	GetDelimiter2() string
}

/*
NoAttributesLike is an instance interface that declares the
complete set of principal, attribute and aspect methods that must be supported
by each instance of a concrete no-attributes-like class.
*/
type NoAttributesLike interface {
	// Principal Methods
	GetClass() NoAttributesClassLike

	// Attribute Methods
	GetDelimiter1() string
	GetDelimiter2() string
	GetDelimiter3() string
}

/*
NoStatementsLike is an instance interface that declares the
complete set of principal, attribute and aspect methods that must be supported
by each instance of a concrete no-statements-like class.
*/
type NoStatementsLike interface {
	// Principal Methods
	GetClass() NoStatementsClassLike

	// Attribute Methods
	GetDelimiter1() string
	GetDelimiter2() string
}

/*
NoValuesLike is an instance interface that declares the
complete set of principal, attribute and aspect methods that must be supported
by each instance of a concrete no-values-like class.
*/
type NoValuesLike interface {
	// Principal Methods
	GetClass() NoValuesClassLike

	// Attribute Methods
	GetDelimiter1() string
	GetDelimiter2() string
}

/*
NotarizeClauseLike is an instance interface that declares the
complete set of principal, attribute and aspect methods that must be supported
by each instance of a concrete notarize-clause-like class.
*/
type NotarizeClauseLike interface {
	// Principal Methods
	GetClass() NotarizeClauseClassLike

	// Attribute Methods
	GetDelimiter1() string
	GetDraft() DraftLike
	GetDelimiter2() string
	GetCited() CitedLike
}

/*
NumericalLike is an instance interface that declares the
complete set of principal, attribute and aspect methods that must be supported
by each instance of a concrete numerical-like class.
*/
type NumericalLike interface {
	// Principal Methods
	GetClass() NumericalClassLike

	// Attribute Methods
	GetAny() any
}

/*
OnClauseLike is an instance interface that declares the
complete set of principal, attribute and aspect methods that must be supported
by each instance of a concrete on-clause-like class.
*/
type OnClauseLike interface {
	// Principal Methods
	GetClass() OnClauseClassLike

	// Attribute Methods
	GetDelimiter() string
	GetFailure() FailureLike
	GetMatchHandlers() col.Sequential[MatchHandlerLike]
}

/*
OperationLike is an instance interface that declares the
complete set of principal, attribute and aspect methods that must be supported
by each instance of a concrete operation-like class.
*/
type OperationLike interface {
	// Principal Methods
	GetClass() OperationClassLike

	// Attribute Methods
	GetAny() any
}

/*
ParametersLike is an instance interface that declares the
complete set of principal, attribute and aspect methods that must be supported
by each instance of a concrete parameters-like class.
*/
type ParametersLike interface {
	// Principal Methods
	GetClass() ParametersClassLike

	// Attribute Methods
	GetAny() any
}

/*
PlusEqualLike is an instance interface that declares the
complete set of principal, attribute and aspect methods that must be supported
by each instance of a concrete plus-equal-like class.
*/
type PlusEqualLike interface {
	// Principal Methods
	GetClass() PlusEqualClassLike

	// Attribute Methods
	GetDelimiter() string
}

/*
PostClauseLike is an instance interface that declares the
complete set of principal, attribute and aspect methods that must be supported
by each instance of a concrete post-clause-like class.
*/
type PostClauseLike interface {
	// Principal Methods
	GetClass() PostClauseClassLike

	// Attribute Methods
	GetDelimiter1() string
	GetMessage() MessageLike
	GetDelimiter2() string
	GetBag() BagLike
}

/*
PrecedenceLike is an instance interface that declares the
complete set of principal, attribute and aspect methods that must be supported
by each instance of a concrete precedence-like class.
*/
type PrecedenceLike interface {
	// Principal Methods
	GetClass() PrecedenceClassLike

	// Attribute Methods
	GetDelimiter1() string
	GetExpression() ExpressionLike
	GetDelimiter2() string
}

/*
PredicateLike is an instance interface that declares the
complete set of principal, attribute and aspect methods that must be supported
by each instance of a concrete predicate-like class.
*/
type PredicateLike interface {
	// Principal Methods
	GetClass() PredicateClassLike

	// Attribute Methods
	GetOperation() OperationLike
	GetExpression() ExpressionLike
}

/*
PrimitiveLike is an instance interface that declares the
complete set of principal, attribute and aspect methods that must be supported
by each instance of a concrete primitive-like class.
*/
type PrimitiveLike interface {
	// Principal Methods
	GetClass() PrimitiveClassLike

	// Attribute Methods
	GetAny() any
}

/*
ProcedureLike is an instance interface that declares the
complete set of principal, attribute and aspect methods that must be supported
by each instance of a concrete procedure-like class.
*/
type ProcedureLike interface {
	// Principal Methods
	GetClass() ProcedureClassLike

	// Attribute Methods
	GetAny() any
}

/*
PublishClauseLike is an instance interface that declares the
complete set of principal, attribute and aspect methods that must be supported
by each instance of a concrete publish-clause-like class.
*/
type PublishClauseLike interface {
	// Principal Methods
	GetClass() PublishClauseClassLike

	// Attribute Methods
	GetDelimiter() string
	GetEvent() EventLike
}

/*
RecipientLike is an instance interface that declares the
complete set of principal, attribute and aspect methods that must be supported
by each instance of a concrete recipient-like class.
*/
type RecipientLike interface {
	// Principal Methods
	GetClass() RecipientClassLike

	// Attribute Methods
	GetAny() any
}

/*
ReferentLike is an instance interface that declares the
complete set of principal, attribute and aspect methods that must be supported
by each instance of a concrete referent-like class.
*/
type ReferentLike interface {
	// Principal Methods
	GetClass() ReferentClassLike

	// Attribute Methods
	GetDelimiter() string
	GetIndirect() IndirectLike
}

/*
RejectClauseLike is an instance interface that declares the
complete set of principal, attribute and aspect methods that must be supported
by each instance of a concrete reject-clause-like class.
*/
type RejectClauseLike interface {
	// Principal Methods
	GetClass() RejectClauseClassLike

	// Attribute Methods
	GetDelimiter() string
	GetMessage() MessageLike
}

/*
RepositoryLike is an instance interface that declares the
complete set of principal, attribute and aspect methods that must be supported
by each instance of a concrete repository-like class.
*/
type RepositoryLike interface {
	// Principal Methods
	GetClass() RepositoryClassLike

	// Attribute Methods
	GetAny() any
}

/*
ResultLike is an instance interface that declares the
complete set of principal, attribute and aspect methods that must be supported
by each instance of a concrete result-like class.
*/
type ResultLike interface {
	// Principal Methods
	GetClass() ResultClassLike

	// Attribute Methods
	GetExpression() ExpressionLike
}

/*
RetrieveClauseLike is an instance interface that declares the
complete set of principal, attribute and aspect methods that must be supported
by each instance of a concrete retrieve-clause-like class.
*/
type RetrieveClauseLike interface {
	// Principal Methods
	GetClass() RetrieveClauseClassLike

	// Attribute Methods
	GetDelimiter1() string
	GetRecipient() RecipientLike
	GetDelimiter2() string
	GetBag() BagLike
}

/*
ReturnClauseLike is an instance interface that declares the
complete set of principal, attribute and aspect methods that must be supported
by each instance of a concrete return-clause-like class.
*/
type ReturnClauseLike interface {
	// Principal Methods
	GetClass() ReturnClauseClassLike

	// Attribute Methods
	GetDelimiter() string
	GetResult() ResultLike
}

/*
SaveClauseLike is an instance interface that declares the
complete set of principal, attribute and aspect methods that must be supported
by each instance of a concrete save-clause-like class.
*/
type SaveClauseLike interface {
	// Principal Methods
	GetClass() SaveClauseClassLike

	// Attribute Methods
	GetDelimiter1() string
	GetDraft() DraftLike
	GetDelimiter2() string
	GetCited() CitedLike
}

/*
SelectClauseLike is an instance interface that declares the
complete set of principal, attribute and aspect methods that must be supported
by each instance of a concrete select-clause-like class.
*/
type SelectClauseLike interface {
	// Principal Methods
	GetClass() SelectClauseClassLike

	// Attribute Methods
	GetDelimiter() string
	GetTarget() TargetLike
	GetMatchHandlers() col.Sequential[MatchHandlerLike]
}

/*
SequenceLike is an instance interface that declares the
complete set of principal, attribute and aspect methods that must be supported
by each instance of a concrete sequence-like class.
*/
type SequenceLike interface {
	// Principal Methods
	GetClass() SequenceClassLike

	// Attribute Methods
	GetExpression() ExpressionLike
}

/*
SlashEqualLike is an instance interface that declares the
complete set of principal, attribute and aspect methods that must be supported
by each instance of a concrete slash-equal-like class.
*/
type SlashEqualLike interface {
	// Principal Methods
	GetClass() SlashEqualClassLike

	// Attribute Methods
	GetDelimiter() string
}

/*
StarEqualLike is an instance interface that declares the
complete set of principal, attribute and aspect methods that must be supported
by each instance of a concrete star-equal-like class.
*/
type StarEqualLike interface {
	// Principal Methods
	GetClass() StarEqualClassLike

	// Attribute Methods
	GetDelimiter() string
}

/*
StatementLike is an instance interface that declares the
complete set of principal, attribute and aspect methods that must be supported
by each instance of a concrete statement-like class.
*/
type StatementLike interface {
	// Principal Methods
	GetClass() StatementClassLike

	// Attribute Methods
	GetMainClause() MainClauseLike
	GetOptionalOnClause() OnClauseLike
}

/*
StatementLineLike is an instance interface that declares the
complete set of principal, attribute and aspect methods that must be supported
by each instance of a concrete statement-line-like class.
*/
type StatementLineLike interface {
	// Principal Methods
	GetClass() StatementLineClassLike

	// Attribute Methods
	GetStatement() StatementLike
	GetOptionalNote() string
}

/*
StringLike is an instance interface that declares the
complete set of principal, attribute and aspect methods that must be supported
by each instance of a concrete string-like class.
*/
type StringLike interface {
	// Principal Methods
	GetClass() StringClassLike

	// Attribute Methods
	GetAny() any
}

/*
SubcomponentLike is an instance interface that declares the
complete set of principal, attribute and aspect methods that must be supported
by each instance of a concrete subcomponent-like class.
*/
type SubcomponentLike interface {
	// Principal Methods
	GetClass() SubcomponentClassLike

	// Attribute Methods
	GetIdentifier() string
	GetDelimiter1() string
	GetIndices() IndicesLike
	GetDelimiter2() string
}

/*
SubjectLike is an instance interface that declares the
complete set of principal, attribute and aspect methods that must be supported
by each instance of a concrete subject-like class.
*/
type SubjectLike interface {
	// Principal Methods
	GetClass() SubjectClassLike

	// Attribute Methods
	GetAny() any
}

/*
TargetLike is an instance interface that declares the
complete set of principal, attribute and aspect methods that must be supported
by each instance of a concrete target-like class.
*/
type TargetLike interface {
	// Principal Methods
	GetClass() TargetClassLike

	// Attribute Methods
	GetAny() any
}

/*
TemplateLike is an instance interface that declares the
complete set of principal, attribute and aspect methods that must be supported
by each instance of a concrete template-like class.
*/
type TemplateLike interface {
	// Principal Methods
	GetClass() TemplateClassLike

	// Attribute Methods
	GetExpression() ExpressionLike
}

/*
ThrowClauseLike is an instance interface that declares the
complete set of principal, attribute and aspect methods that must be supported
by each instance of a concrete throw-clause-like class.
*/
type ThrowClauseLike interface {
	// Principal Methods
	GetClass() ThrowClauseClassLike

	// Attribute Methods
	GetDelimiter() string
	GetException() ExceptionLike
}

/*
VariableLike is an instance interface that declares the
complete set of principal, attribute and aspect methods that must be supported
by each instance of a concrete variable-like class.
*/
type VariableLike interface {
	// Principal Methods
	GetClass() VariableClassLike

	// Attribute Methods
	GetIdentifier() string
}

/*
WhileClauseLike is an instance interface that declares the
complete set of principal, attribute and aspect methods that must be supported
by each instance of a concrete while-clause-like class.
*/
type WhileClauseLike interface {
	// Principal Methods
	GetClass() WhileClauseClassLike

	// Attribute Methods
	GetDelimiter1() string
	GetCondition() ConditionLike
	GetDelimiter2() string
	GetProcedure() ProcedureLike
}

/*
WithClauseLike is an instance interface that declares the
complete set of principal, attribute and aspect methods that must be supported
by each instance of a concrete with-clause-like class.
*/
type WithClauseLike interface {
	// Principal Methods
	GetClass() WithClauseClassLike

	// Attribute Methods
	GetDelimiter1() string
	GetDelimiter2() string
	GetItem() ItemLike
	GetDelimiter3() string
	GetSequence() SequenceLike
	GetDelimiter4() string
	GetProcedure() ProcedureLike
}

// ASPECT DECLARATIONS
