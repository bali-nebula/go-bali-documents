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
│         This "package_api.go" file was automatically generated using:        │
│            https://github.com/craterdog/go-development-tools/wiki            │
│                                                                              │
│                     Any updates to it may be overwritten.                    │
└──────────────────────────────────────────────────────────────────────────────┘

Package "grammar" provides the following grammar classes that operate on the
abstract syntax tree (AST) for this module:
  - Token captures the attributes associated with a parsed token.
  - Scanner is used to scan the source byte stream and recognize matching tokens.
  - Parser is used to process the token stream and generate the AST.
  - Validator is used to validate the semantics associated with an AST.
  - Formatter is used to format an AST back into a canonical version of its source.
  - Visitor walks the AST and calls processor methods for each node in the tree.
  - Processor provides empty processor methods to be inherited by the processors.

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
package grammar

import (
	ast "github.com/bali-nebula/go-bali-documents/v3/ast"
	fra "github.com/craterdog/go-component-framework/v7"
)

// TYPE DECLARATIONS

/*
TokenType is a constrained type representing any token type recognized by a
scanner.
*/
type TokenType uint8

const (
	ErrorToken TokenType = iota
	AngleToken
	BinaryToken
	BooleanToken
	BytecodeToken
	CommentToken
	DelimiterToken
	DurationToken
	GlyphToken
	IdentifierToken
	MomentToken
	NameToken
	NarrativeToken
	NewlineToken
	NoteToken
	NumberToken
	PatternToken
	PercentageToken
	ProbabilityToken
	QuoteToken
	ResourceToken
	SpaceToken
	SymbolToken
	TagToken
	VersionToken
)

// FUNCTIONAL DECLARATIONS

// CLASS DECLARATIONS

/*
FormatterClassLike is a class interface that declares the complete set of
class constructors, constants and functions that must be supported by each
concrete formatter-like class.
*/
type FormatterClassLike interface {
	// Constructor Methods
	Formatter() FormatterLike
}

/*
ParserClassLike is a class interface that declares the complete set of
class constructors, constants and functions that must be supported by each
concrete parser-like class.
*/
type ParserClassLike interface {
	// Constructor Methods
	Parser() ParserLike
}

/*
ProcessorClassLike is a class interface that declares the complete set of
class constructors, constants and functions that must be supported by each
concrete processor-like class.
*/
type ProcessorClassLike interface {
	// Constructor Methods
	Processor() ProcessorLike
}

/*
ScannerClassLike is a class interface that declares the complete set of
class constructors, constants and functions that must be supported by each
concrete scanner-like class.  The following functions are supported:

FormatToken() returns a formatted string containing the attributes of the token.

FormatType() returns the string version of the token type.

MatchesType() determines whether or not a token value is of a specified type.
*/
type ScannerClassLike interface {
	// Constructor Methods
	Scanner(
		source string,
		tokens fra.QueueLike[TokenLike],
	) ScannerLike

	// Function Methods
	FormatToken(
		token TokenLike,
	) string
	FormatType(
		tokenType TokenType,
	) string
	MatchesType(
		tokenValue string,
		tokenType TokenType,
	) bool
}

/*
TokenClassLike is a class interface that declares the complete set of
class constructors, constants and functions that must be supported by each
concrete visitor-like class.
*/
type TokenClassLike interface {
	// Constructor Methods
	Token(
		line uint,
		position uint,
		type_ TokenType,
		value string,
	) TokenLike
}

/*
ValidatorClassLike is a class interface that declares the complete set of
class constructors, constants and functions that must be supported by each
concrete validator-like class.
*/
type ValidatorClassLike interface {
	// Constructor Methods
	Validator() ValidatorLike
}

/*
VisitorClassLike is a class interface that declares the complete set of
class constructors, constants and functions that must be supported by each
concrete visitor-like class.
*/
type VisitorClassLike interface {
	// Constructor Methods
	Visitor(
		processor Methodical,
	) VisitorLike
}

// INSTANCE DECLARATIONS

/*
FormatterLike is an instance interface that declares the complete set of
principal, attribute and aspect methods that must be supported by each
instance of a concrete formatter-like class.
*/
type FormatterLike interface {
	// Principal Methods
	GetClass() FormatterClassLike
	FormatDocument(
		document ast.DocumentLike,
	) string

	// Aspect Interfaces
	Methodical
}

/*
ParserLike is an instance interface that declares the complete set of
principal, attribute and aspect methods that must be supported by each
instance of a concrete parser-like class.
*/
type ParserLike interface {
	// Principal Methods
	GetClass() ParserClassLike
	ParseSource(
		source string,
	) ast.DocumentLike
}

/*
ProcessorLike is an instance interface that declares the complete set of
principal, attribute and aspect methods that must be supported by each
instance of a concrete processor-like class.
*/
type ProcessorLike interface {
	// Principal Methods
	GetClass() ProcessorClassLike

	// Aspect Interfaces
	Methodical
}

/*
ScannerLike is an instance interface that declares the complete set of
principal, attribute and aspect methods that must be supported by each
instance of a concrete scanner-like class.
*/
type ScannerLike interface {
	// Principal Methods
	GetClass() ScannerClassLike
}

/*
TokenLike is an instance interface that declares the complete set of
principal, attribute and aspect methods that must be supported by each
instance of a concrete token-like class.
*/
type TokenLike interface {
	// Principal Methods
	GetClass() TokenClassLike

	// Attribute Methods
	GetLine() uint
	GetPosition() uint
	GetType() TokenType
	GetValue() string
}

/*
ValidatorLike is an instance interface that declares the complete set of
principal, attribute and aspect methods that must be supported by each
instance of a concrete validator-like class.
*/
type ValidatorLike interface {
	// Principal Methods
	GetClass() ValidatorClassLike
	ValidateDocument(
		document ast.DocumentLike,
	)

	// Aspect Interfaces
	Methodical
}

/*
VisitorLike is an instance interface that declares the complete set of
principal, attribute and aspect methods that must be supported by each
instance of a concrete visitor-like class.
*/
type VisitorLike interface {
	// Principal Methods
	GetClass() VisitorClassLike
	VisitDocument(
		document ast.DocumentLike,
	)
}

// ASPECT DECLARATIONS

/*
Methodical declares the set of method signatures that must be supported by
all methodical processors.
*/
type Methodical interface {
	ProcessAngle(
		angle string,
	)
	ProcessBinary(
		binary string,
	)
	ProcessBoolean(
		boolean string,
	)
	ProcessBytecode(
		bytecode string,
	)
	ProcessComment(
		comment string,
	)
	ProcessDelimiter(
		delimiter string,
	)
	ProcessDuration(
		duration string,
	)
	ProcessGlyph(
		glyph string,
	)
	ProcessIdentifier(
		identifier string,
	)
	ProcessMoment(
		moment string,
	)
	ProcessName(
		name string,
	)
	ProcessNarrative(
		narrative string,
	)
	ProcessNewline(
		newline string,
	)
	ProcessNote(
		note string,
	)
	ProcessNumber(
		number string,
	)
	ProcessPattern(
		pattern string,
	)
	ProcessPercentage(
		percentage string,
	)
	ProcessProbability(
		probability string,
	)
	ProcessQuote(
		quote string,
	)
	ProcessResource(
		resource string,
	)
	ProcessSpace(
		space string,
	)
	ProcessSymbol(
		symbol string,
	)
	ProcessTag(
		tag string,
	)
	ProcessVersion(
		version string,
	)
	PreprocessAcceptClause(
		acceptClause ast.AcceptClauseLike,
		index_ uint,
		count_ uint,
	)
	PostprocessAcceptClause(
		acceptClause ast.AcceptClauseLike,
		index_ uint,
		count_ uint,
	)
	ProcessAcceptClauseSlot(
		acceptClause ast.AcceptClauseLike,
		slot_ uint,
	)
	PreprocessActionInduction(
		actionInduction ast.ActionInductionLike,
		index_ uint,
		count_ uint,
	)
	PostprocessActionInduction(
		actionInduction ast.ActionInductionLike,
		index_ uint,
		count_ uint,
	)
	ProcessActionInductionSlot(
		actionInduction ast.ActionInductionLike,
		slot_ uint,
	)
	PreprocessAnnotation(
		annotation ast.AnnotationLike,
		index_ uint,
		count_ uint,
	)
	PostprocessAnnotation(
		annotation ast.AnnotationLike,
		index_ uint,
		count_ uint,
	)
	ProcessAnnotationSlot(
		annotation ast.AnnotationLike,
		slot_ uint,
	)
	PreprocessArgument(
		argument ast.ArgumentLike,
		index_ uint,
		count_ uint,
	)
	PostprocessArgument(
		argument ast.ArgumentLike,
		index_ uint,
		count_ uint,
	)
	ProcessArgumentSlot(
		argument ast.ArgumentLike,
		slot_ uint,
	)
	PreprocessArithmetic(
		arithmetic ast.ArithmeticLike,
		index_ uint,
		count_ uint,
	)
	PostprocessArithmetic(
		arithmetic ast.ArithmeticLike,
		index_ uint,
		count_ uint,
	)
	ProcessArithmeticSlot(
		arithmetic ast.ArithmeticLike,
		slot_ uint,
	)
	PreprocessAssignment(
		assignment ast.AssignmentLike,
		index_ uint,
		count_ uint,
	)
	PostprocessAssignment(
		assignment ast.AssignmentLike,
		index_ uint,
		count_ uint,
	)
	ProcessAssignmentSlot(
		assignment ast.AssignmentLike,
		slot_ uint,
	)
	PreprocessAssociation(
		association ast.AssociationLike,
		index_ uint,
		count_ uint,
	)
	PostprocessAssociation(
		association ast.AssociationLike,
		index_ uint,
		count_ uint,
	)
	ProcessAssociationSlot(
		association ast.AssociationLike,
		slot_ uint,
	)
	PreprocessAtLevel(
		atLevel ast.AtLevelLike,
		index_ uint,
		count_ uint,
	)
	PostprocessAtLevel(
		atLevel ast.AtLevelLike,
		index_ uint,
		count_ uint,
	)
	ProcessAtLevelSlot(
		atLevel ast.AtLevelLike,
		slot_ uint,
	)
	PreprocessAttributes(
		attributes ast.AttributesLike,
		index_ uint,
		count_ uint,
	)
	PostprocessAttributes(
		attributes ast.AttributesLike,
		index_ uint,
		count_ uint,
	)
	ProcessAttributesSlot(
		attributes ast.AttributesLike,
		slot_ uint,
	)
	PreprocessBag(
		bag ast.BagLike,
		index_ uint,
		count_ uint,
	)
	PostprocessBag(
		bag ast.BagLike,
		index_ uint,
		count_ uint,
	)
	ProcessBagSlot(
		bag ast.BagLike,
		slot_ uint,
	)
	PreprocessBra(
		bra ast.BraLike,
		index_ uint,
		count_ uint,
	)
	PostprocessBra(
		bra ast.BraLike,
		index_ uint,
		count_ uint,
	)
	ProcessBraSlot(
		bra ast.BraLike,
		slot_ uint,
	)
	PreprocessBreakClause(
		breakClause ast.BreakClauseLike,
		index_ uint,
		count_ uint,
	)
	PostprocessBreakClause(
		breakClause ast.BreakClauseLike,
		index_ uint,
		count_ uint,
	)
	ProcessBreakClauseSlot(
		breakClause ast.BreakClauseLike,
		slot_ uint,
	)
	PreprocessCheckoutClause(
		checkoutClause ast.CheckoutClauseLike,
		index_ uint,
		count_ uint,
	)
	PostprocessCheckoutClause(
		checkoutClause ast.CheckoutClauseLike,
		index_ uint,
		count_ uint,
	)
	ProcessCheckoutClauseSlot(
		checkoutClause ast.CheckoutClauseLike,
		slot_ uint,
	)
	PreprocessCited(
		cited ast.CitedLike,
		index_ uint,
		count_ uint,
	)
	PostprocessCited(
		cited ast.CitedLike,
		index_ uint,
		count_ uint,
	)
	ProcessCitedSlot(
		cited ast.CitedLike,
		slot_ uint,
	)
	PreprocessCollection(
		collection ast.CollectionLike,
		index_ uint,
		count_ uint,
	)
	PostprocessCollection(
		collection ast.CollectionLike,
		index_ uint,
		count_ uint,
	)
	ProcessCollectionSlot(
		collection ast.CollectionLike,
		slot_ uint,
	)
	PreprocessComparison(
		comparison ast.ComparisonLike,
		index_ uint,
		count_ uint,
	)
	PostprocessComparison(
		comparison ast.ComparisonLike,
		index_ uint,
		count_ uint,
	)
	ProcessComparisonSlot(
		comparison ast.ComparisonLike,
		slot_ uint,
	)
	PreprocessComplement(
		complement ast.ComplementLike,
		index_ uint,
		count_ uint,
	)
	PostprocessComplement(
		complement ast.ComplementLike,
		index_ uint,
		count_ uint,
	)
	ProcessComplementSlot(
		complement ast.ComplementLike,
		slot_ uint,
	)
	PreprocessComponent(
		component ast.ComponentLike,
		index_ uint,
		count_ uint,
	)
	PostprocessComponent(
		component ast.ComponentLike,
		index_ uint,
		count_ uint,
	)
	ProcessComponentSlot(
		component ast.ComponentLike,
		slot_ uint,
	)
	PreprocessCondition(
		condition ast.ConditionLike,
		index_ uint,
		count_ uint,
	)
	PostprocessCondition(
		condition ast.ConditionLike,
		index_ uint,
		count_ uint,
	)
	ProcessConditionSlot(
		condition ast.ConditionLike,
		slot_ uint,
	)
	PreprocessContinueClause(
		continueClause ast.ContinueClauseLike,
		index_ uint,
		count_ uint,
	)
	PostprocessContinueClause(
		continueClause ast.ContinueClauseLike,
		index_ uint,
		count_ uint,
	)
	ProcessContinueClauseSlot(
		continueClause ast.ContinueClauseLike,
		slot_ uint,
	)
	PreprocessDiscardClause(
		discardClause ast.DiscardClauseLike,
		index_ uint,
		count_ uint,
	)
	PostprocessDiscardClause(
		discardClause ast.DiscardClauseLike,
		index_ uint,
		count_ uint,
	)
	ProcessDiscardClauseSlot(
		discardClause ast.DiscardClauseLike,
		slot_ uint,
	)
	PreprocessDoClause(
		doClause ast.DoClauseLike,
		index_ uint,
		count_ uint,
	)
	PostprocessDoClause(
		doClause ast.DoClauseLike,
		index_ uint,
		count_ uint,
	)
	ProcessDoClauseSlot(
		doClause ast.DoClauseLike,
		slot_ uint,
	)
	PreprocessDocument(
		document ast.DocumentLike,
		index_ uint,
		count_ uint,
	)
	PostprocessDocument(
		document ast.DocumentLike,
		index_ uint,
		count_ uint,
	)
	ProcessDocumentSlot(
		document ast.DocumentLike,
		slot_ uint,
	)
	PreprocessDraft(
		draft ast.DraftLike,
		index_ uint,
		count_ uint,
	)
	PostprocessDraft(
		draft ast.DraftLike,
		index_ uint,
		count_ uint,
	)
	ProcessDraftSlot(
		draft ast.DraftLike,
		slot_ uint,
	)
	PreprocessElement(
		element ast.ElementLike,
		index_ uint,
		count_ uint,
	)
	PostprocessElement(
		element ast.ElementLike,
		index_ uint,
		count_ uint,
	)
	ProcessElementSlot(
		element ast.ElementLike,
		slot_ uint,
	)
	PreprocessEmpty(
		empty ast.EmptyLike,
		index_ uint,
		count_ uint,
	)
	PostprocessEmpty(
		empty ast.EmptyLike,
		index_ uint,
		count_ uint,
	)
	ProcessEmptySlot(
		empty ast.EmptyLike,
		slot_ uint,
	)
	PreprocessEvent(
		event ast.EventLike,
		index_ uint,
		count_ uint,
	)
	PostprocessEvent(
		event ast.EventLike,
		index_ uint,
		count_ uint,
	)
	ProcessEventSlot(
		event ast.EventLike,
		slot_ uint,
	)
	PreprocessException(
		exception ast.ExceptionLike,
		index_ uint,
		count_ uint,
	)
	PostprocessException(
		exception ast.ExceptionLike,
		index_ uint,
		count_ uint,
	)
	ProcessExceptionSlot(
		exception ast.ExceptionLike,
		slot_ uint,
	)
	PreprocessExpression(
		expression ast.ExpressionLike,
		index_ uint,
		count_ uint,
	)
	PostprocessExpression(
		expression ast.ExpressionLike,
		index_ uint,
		count_ uint,
	)
	ProcessExpressionSlot(
		expression ast.ExpressionLike,
		slot_ uint,
	)
	PreprocessFailure(
		failure ast.FailureLike,
		index_ uint,
		count_ uint,
	)
	PostprocessFailure(
		failure ast.FailureLike,
		index_ uint,
		count_ uint,
	)
	ProcessFailureSlot(
		failure ast.FailureLike,
		slot_ uint,
	)
	PreprocessFlowControl(
		flowControl ast.FlowControlLike,
		index_ uint,
		count_ uint,
	)
	PostprocessFlowControl(
		flowControl ast.FlowControlLike,
		index_ uint,
		count_ uint,
	)
	ProcessFlowControlSlot(
		flowControl ast.FlowControlLike,
		slot_ uint,
	)
	PreprocessFunction(
		function ast.FunctionLike,
		index_ uint,
		count_ uint,
	)
	PostprocessFunction(
		function ast.FunctionLike,
		index_ uint,
		count_ uint,
	)
	ProcessFunctionSlot(
		function ast.FunctionLike,
		slot_ uint,
	)
	PreprocessIfClause(
		ifClause ast.IfClauseLike,
		index_ uint,
		count_ uint,
	)
	PostprocessIfClause(
		ifClause ast.IfClauseLike,
		index_ uint,
		count_ uint,
	)
	ProcessIfClauseSlot(
		ifClause ast.IfClauseLike,
		slot_ uint,
	)
	PreprocessIndex(
		index ast.IndexLike,
		index_ uint,
		count_ uint,
	)
	PostprocessIndex(
		index ast.IndexLike,
		index_ uint,
		count_ uint,
	)
	ProcessIndexSlot(
		index ast.IndexLike,
		slot_ uint,
	)
	PreprocessIndirect(
		indirect ast.IndirectLike,
		index_ uint,
		count_ uint,
	)
	PostprocessIndirect(
		indirect ast.IndirectLike,
		index_ uint,
		count_ uint,
	)
	ProcessIndirectSlot(
		indirect ast.IndirectLike,
		slot_ uint,
	)
	PreprocessInverse(
		inverse ast.InverseLike,
		index_ uint,
		count_ uint,
	)
	PostprocessInverse(
		inverse ast.InverseLike,
		index_ uint,
		count_ uint,
	)
	ProcessInverseSlot(
		inverse ast.InverseLike,
		slot_ uint,
	)
	PreprocessInversion(
		inversion ast.InversionLike,
		index_ uint,
		count_ uint,
	)
	PostprocessInversion(
		inversion ast.InversionLike,
		index_ uint,
		count_ uint,
	)
	ProcessInversionSlot(
		inversion ast.InversionLike,
		slot_ uint,
	)
	PreprocessInvocation(
		invocation ast.InvocationLike,
		index_ uint,
		count_ uint,
	)
	PostprocessInvocation(
		invocation ast.InvocationLike,
		index_ uint,
		count_ uint,
	)
	ProcessInvocationSlot(
		invocation ast.InvocationLike,
		slot_ uint,
	)
	PreprocessInvoke(
		invoke ast.InvokeLike,
		index_ uint,
		count_ uint,
	)
	PostprocessInvoke(
		invoke ast.InvokeLike,
		index_ uint,
		count_ uint,
	)
	ProcessInvokeSlot(
		invoke ast.InvokeLike,
		slot_ uint,
	)
	PreprocessItems(
		items ast.ItemsLike,
		index_ uint,
		count_ uint,
	)
	PostprocessItems(
		items ast.ItemsLike,
		index_ uint,
		count_ uint,
	)
	ProcessItemsSlot(
		items ast.ItemsLike,
		slot_ uint,
	)
	PreprocessKet(
		ket ast.KetLike,
		index_ uint,
		count_ uint,
	)
	PostprocessKet(
		ket ast.KetLike,
		index_ uint,
		count_ uint,
	)
	ProcessKetSlot(
		ket ast.KetLike,
		slot_ uint,
	)
	PreprocessLetClause(
		letClause ast.LetClauseLike,
		index_ uint,
		count_ uint,
	)
	PostprocessLetClause(
		letClause ast.LetClauseLike,
		index_ uint,
		count_ uint,
	)
	ProcessLetClauseSlot(
		letClause ast.LetClauseLike,
		slot_ uint,
	)
	PreprocessLine(
		line ast.LineLike,
		index_ uint,
		count_ uint,
	)
	PostprocessLine(
		line ast.LineLike,
		index_ uint,
		count_ uint,
	)
	ProcessLineSlot(
		line ast.LineLike,
		slot_ uint,
	)
	PreprocessLogic(
		logic ast.LogicLike,
		index_ uint,
		count_ uint,
	)
	PostprocessLogic(
		logic ast.LogicLike,
		index_ uint,
		count_ uint,
	)
	ProcessLogicSlot(
		logic ast.LogicLike,
		slot_ uint,
	)
	PreprocessLogical(
		logical ast.LogicalLike,
		index_ uint,
		count_ uint,
	)
	PostprocessLogical(
		logical ast.LogicalLike,
		index_ uint,
		count_ uint,
	)
	ProcessLogicalSlot(
		logical ast.LogicalLike,
		slot_ uint,
	)
	PreprocessMagnitude(
		magnitude ast.MagnitudeLike,
		index_ uint,
		count_ uint,
	)
	PostprocessMagnitude(
		magnitude ast.MagnitudeLike,
		index_ uint,
		count_ uint,
	)
	ProcessMagnitudeSlot(
		magnitude ast.MagnitudeLike,
		slot_ uint,
	)
	PreprocessMainClause(
		mainClause ast.MainClauseLike,
		index_ uint,
		count_ uint,
	)
	PostprocessMainClause(
		mainClause ast.MainClauseLike,
		index_ uint,
		count_ uint,
	)
	ProcessMainClauseSlot(
		mainClause ast.MainClauseLike,
		slot_ uint,
	)
	PreprocessMatchingClause(
		matchingClause ast.MatchingClauseLike,
		index_ uint,
		count_ uint,
	)
	PostprocessMatchingClause(
		matchingClause ast.MatchingClauseLike,
		index_ uint,
		count_ uint,
	)
	ProcessMatchingClauseSlot(
		matchingClause ast.MatchingClauseLike,
		slot_ uint,
	)
	PreprocessMessage(
		message ast.MessageLike,
		index_ uint,
		count_ uint,
	)
	PostprocessMessage(
		message ast.MessageLike,
		index_ uint,
		count_ uint,
	)
	ProcessMessageSlot(
		message ast.MessageLike,
		slot_ uint,
	)
	PreprocessMessageHandling(
		messageHandling ast.MessageHandlingLike,
		index_ uint,
		count_ uint,
	)
	PostprocessMessageHandling(
		messageHandling ast.MessageHandlingLike,
		index_ uint,
		count_ uint,
	)
	ProcessMessageHandlingSlot(
		messageHandling ast.MessageHandlingLike,
		slot_ uint,
	)
	PreprocessMethod(
		method ast.MethodLike,
		index_ uint,
		count_ uint,
	)
	PostprocessMethod(
		method ast.MethodLike,
		index_ uint,
		count_ uint,
	)
	ProcessMethodSlot(
		method ast.MethodLike,
		slot_ uint,
	)
	PreprocessNotarizeClause(
		notarizeClause ast.NotarizeClauseLike,
		index_ uint,
		count_ uint,
	)
	PostprocessNotarizeClause(
		notarizeClause ast.NotarizeClauseLike,
		index_ uint,
		count_ uint,
	)
	ProcessNotarizeClauseSlot(
		notarizeClause ast.NotarizeClauseLike,
		slot_ uint,
	)
	PreprocessNumerical(
		numerical ast.NumericalLike,
		index_ uint,
		count_ uint,
	)
	PostprocessNumerical(
		numerical ast.NumericalLike,
		index_ uint,
		count_ uint,
	)
	ProcessNumericalSlot(
		numerical ast.NumericalLike,
		slot_ uint,
	)
	PreprocessOnClause(
		onClause ast.OnClauseLike,
		index_ uint,
		count_ uint,
	)
	PostprocessOnClause(
		onClause ast.OnClauseLike,
		index_ uint,
		count_ uint,
	)
	ProcessOnClauseSlot(
		onClause ast.OnClauseLike,
		slot_ uint,
	)
	PreprocessOperation(
		operation ast.OperationLike,
		index_ uint,
		count_ uint,
	)
	PostprocessOperation(
		operation ast.OperationLike,
		index_ uint,
		count_ uint,
	)
	ProcessOperationSlot(
		operation ast.OperationLike,
		slot_ uint,
	)
	PreprocessParameters(
		parameters ast.ParametersLike,
		index_ uint,
		count_ uint,
	)
	PostprocessParameters(
		parameters ast.ParametersLike,
		index_ uint,
		count_ uint,
	)
	ProcessParametersSlot(
		parameters ast.ParametersLike,
		slot_ uint,
	)
	PreprocessPostClause(
		postClause ast.PostClauseLike,
		index_ uint,
		count_ uint,
	)
	PostprocessPostClause(
		postClause ast.PostClauseLike,
		index_ uint,
		count_ uint,
	)
	ProcessPostClauseSlot(
		postClause ast.PostClauseLike,
		slot_ uint,
	)
	PreprocessPrecedence(
		precedence ast.PrecedenceLike,
		index_ uint,
		count_ uint,
	)
	PostprocessPrecedence(
		precedence ast.PrecedenceLike,
		index_ uint,
		count_ uint,
	)
	ProcessPrecedenceSlot(
		precedence ast.PrecedenceLike,
		slot_ uint,
	)
	PreprocessPredicate(
		predicate ast.PredicateLike,
		index_ uint,
		count_ uint,
	)
	PostprocessPredicate(
		predicate ast.PredicateLike,
		index_ uint,
		count_ uint,
	)
	ProcessPredicateSlot(
		predicate ast.PredicateLike,
		slot_ uint,
	)
	PreprocessPrimitive(
		primitive ast.PrimitiveLike,
		index_ uint,
		count_ uint,
	)
	PostprocessPrimitive(
		primitive ast.PrimitiveLike,
		index_ uint,
		count_ uint,
	)
	ProcessPrimitiveSlot(
		primitive ast.PrimitiveLike,
		slot_ uint,
	)
	PreprocessProcedure(
		procedure ast.ProcedureLike,
		index_ uint,
		count_ uint,
	)
	PostprocessProcedure(
		procedure ast.ProcedureLike,
		index_ uint,
		count_ uint,
	)
	ProcessProcedureSlot(
		procedure ast.ProcedureLike,
		slot_ uint,
	)
	PreprocessPublishClause(
		publishClause ast.PublishClauseLike,
		index_ uint,
		count_ uint,
	)
	PostprocessPublishClause(
		publishClause ast.PublishClauseLike,
		index_ uint,
		count_ uint,
	)
	ProcessPublishClauseSlot(
		publishClause ast.PublishClauseLike,
		slot_ uint,
	)
	PreprocessRange(
		range_ ast.RangeLike,
		index_ uint,
		count_ uint,
	)
	PostprocessRange(
		range_ ast.RangeLike,
		index_ uint,
		count_ uint,
	)
	ProcessRangeSlot(
		range_ ast.RangeLike,
		slot_ uint,
	)
	PreprocessRecipient(
		recipient ast.RecipientLike,
		index_ uint,
		count_ uint,
	)
	PostprocessRecipient(
		recipient ast.RecipientLike,
		index_ uint,
		count_ uint,
	)
	ProcessRecipientSlot(
		recipient ast.RecipientLike,
		slot_ uint,
	)
	PreprocessReferent(
		referent ast.ReferentLike,
		index_ uint,
		count_ uint,
	)
	PostprocessReferent(
		referent ast.ReferentLike,
		index_ uint,
		count_ uint,
	)
	ProcessReferentSlot(
		referent ast.ReferentLike,
		slot_ uint,
	)
	PreprocessRejectClause(
		rejectClause ast.RejectClauseLike,
		index_ uint,
		count_ uint,
	)
	PostprocessRejectClause(
		rejectClause ast.RejectClauseLike,
		index_ uint,
		count_ uint,
	)
	ProcessRejectClauseSlot(
		rejectClause ast.RejectClauseLike,
		slot_ uint,
	)
	PreprocessRepositoryAccess(
		repositoryAccess ast.RepositoryAccessLike,
		index_ uint,
		count_ uint,
	)
	PostprocessRepositoryAccess(
		repositoryAccess ast.RepositoryAccessLike,
		index_ uint,
		count_ uint,
	)
	ProcessRepositoryAccessSlot(
		repositoryAccess ast.RepositoryAccessLike,
		slot_ uint,
	)
	PreprocessResult(
		result ast.ResultLike,
		index_ uint,
		count_ uint,
	)
	PostprocessResult(
		result ast.ResultLike,
		index_ uint,
		count_ uint,
	)
	ProcessResultSlot(
		result ast.ResultLike,
		slot_ uint,
	)
	PreprocessRetrieveClause(
		retrieveClause ast.RetrieveClauseLike,
		index_ uint,
		count_ uint,
	)
	PostprocessRetrieveClause(
		retrieveClause ast.RetrieveClauseLike,
		index_ uint,
		count_ uint,
	)
	ProcessRetrieveClauseSlot(
		retrieveClause ast.RetrieveClauseLike,
		slot_ uint,
	)
	PreprocessReturnClause(
		returnClause ast.ReturnClauseLike,
		index_ uint,
		count_ uint,
	)
	PostprocessReturnClause(
		returnClause ast.ReturnClauseLike,
		index_ uint,
		count_ uint,
	)
	ProcessReturnClauseSlot(
		returnClause ast.ReturnClauseLike,
		slot_ uint,
	)
	PreprocessSaveClause(
		saveClause ast.SaveClauseLike,
		index_ uint,
		count_ uint,
	)
	PostprocessSaveClause(
		saveClause ast.SaveClauseLike,
		index_ uint,
		count_ uint,
	)
	ProcessSaveClauseSlot(
		saveClause ast.SaveClauseLike,
		slot_ uint,
	)
	PreprocessSelectClause(
		selectClause ast.SelectClauseLike,
		index_ uint,
		count_ uint,
	)
	PostprocessSelectClause(
		selectClause ast.SelectClauseLike,
		index_ uint,
		count_ uint,
	)
	ProcessSelectClauseSlot(
		selectClause ast.SelectClauseLike,
		slot_ uint,
	)
	PreprocessSequence(
		sequence ast.SequenceLike,
		index_ uint,
		count_ uint,
	)
	PostprocessSequence(
		sequence ast.SequenceLike,
		index_ uint,
		count_ uint,
	)
	ProcessSequenceSlot(
		sequence ast.SequenceLike,
		slot_ uint,
	)
	PreprocessStatement(
		statement ast.StatementLike,
		index_ uint,
		count_ uint,
	)
	PostprocessStatement(
		statement ast.StatementLike,
		index_ uint,
		count_ uint,
	)
	ProcessStatementSlot(
		statement ast.StatementLike,
		slot_ uint,
	)
	PreprocessString(
		string_ ast.StringLike,
		index_ uint,
		count_ uint,
	)
	PostprocessString(
		string_ ast.StringLike,
		index_ uint,
		count_ uint,
	)
	ProcessStringSlot(
		string_ ast.StringLike,
		slot_ uint,
	)
	PreprocessSubcomponent(
		subcomponent ast.SubcomponentLike,
		index_ uint,
		count_ uint,
	)
	PostprocessSubcomponent(
		subcomponent ast.SubcomponentLike,
		index_ uint,
		count_ uint,
	)
	ProcessSubcomponentSlot(
		subcomponent ast.SubcomponentLike,
		slot_ uint,
	)
	PreprocessSubject(
		subject ast.SubjectLike,
		index_ uint,
		count_ uint,
	)
	PostprocessSubject(
		subject ast.SubjectLike,
		index_ uint,
		count_ uint,
	)
	ProcessSubjectSlot(
		subject ast.SubjectLike,
		slot_ uint,
	)
	PreprocessTarget(
		target ast.TargetLike,
		index_ uint,
		count_ uint,
	)
	PostprocessTarget(
		target ast.TargetLike,
		index_ uint,
		count_ uint,
	)
	ProcessTargetSlot(
		target ast.TargetLike,
		slot_ uint,
	)
	PreprocessTemplate(
		template ast.TemplateLike,
		index_ uint,
		count_ uint,
	)
	PostprocessTemplate(
		template ast.TemplateLike,
		index_ uint,
		count_ uint,
	)
	ProcessTemplateSlot(
		template ast.TemplateLike,
		slot_ uint,
	)
	PreprocessTextual(
		textual ast.TextualLike,
		index_ uint,
		count_ uint,
	)
	PostprocessTextual(
		textual ast.TextualLike,
		index_ uint,
		count_ uint,
	)
	ProcessTextualSlot(
		textual ast.TextualLike,
		slot_ uint,
	)
	PreprocessThrowClause(
		throwClause ast.ThrowClauseLike,
		index_ uint,
		count_ uint,
	)
	PostprocessThrowClause(
		throwClause ast.ThrowClauseLike,
		index_ uint,
		count_ uint,
	)
	ProcessThrowClauseSlot(
		throwClause ast.ThrowClauseLike,
		slot_ uint,
	)
	PreprocessValue(
		value ast.ValueLike,
		index_ uint,
		count_ uint,
	)
	PostprocessValue(
		value ast.ValueLike,
		index_ uint,
		count_ uint,
	)
	ProcessValueSlot(
		value ast.ValueLike,
		slot_ uint,
	)
	PreprocessVariable(
		variable ast.VariableLike,
		index_ uint,
		count_ uint,
	)
	PostprocessVariable(
		variable ast.VariableLike,
		index_ uint,
		count_ uint,
	)
	ProcessVariableSlot(
		variable ast.VariableLike,
		slot_ uint,
	)
	PreprocessWhileClause(
		whileClause ast.WhileClauseLike,
		index_ uint,
		count_ uint,
	)
	PostprocessWhileClause(
		whileClause ast.WhileClauseLike,
		index_ uint,
		count_ uint,
	)
	ProcessWhileClauseSlot(
		whileClause ast.WhileClauseLike,
		slot_ uint,
	)
	PreprocessWithClause(
		withClause ast.WithClauseLike,
		index_ uint,
		count_ uint,
	)
	PostprocessWithClause(
		withClause ast.WithClauseLike,
		index_ uint,
		count_ uint,
	)
	ProcessWithClauseSlot(
		withClause ast.WithClauseLike,
		slot_ uint,
	)
}
