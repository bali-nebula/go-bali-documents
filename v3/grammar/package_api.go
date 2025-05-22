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
	col "github.com/craterdog/go-collection-framework/v7"
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
	ArrowToken
	BinaryToken
	BooleanToken
	BytecodeToken
	CaretToken
	CitationToken
	CommentToken
	DelimiterToken
	DotToken
	DurationToken
	EqualToken
	IdentifierToken
	LessToken
	MinusToken
	MomentToken
	MoreToken
	NameToken
	NarrativeToken
	NewlineToken
	NoteToken
	NumberToken
	PatternToken
	PercentToken
	PercentageToken
	PlusToken
	ProbabilityToken
	QuoteToken
	ResourceToken
	SlashToken
	SpaceToken
	StarToken
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
		tokens col.QueueLike[TokenLike],
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
	ProcessArrow(
		arrow string,
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
	ProcessCaret(
		caret string,
	)
	ProcessCitation(
		citation string,
	)
	ProcessComment(
		comment string,
	)
	ProcessDelimiter(
		delimiter string,
	)
	ProcessDot(
		dot string,
	)
	ProcessDuration(
		duration string,
	)
	ProcessEqual(
		equal string,
	)
	ProcessIdentifier(
		identifier string,
	)
	ProcessLess(
		less string,
	)
	ProcessMinus(
		minus string,
	)
	ProcessMoment(
		moment string,
	)
	ProcessMore(
		more string,
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
	ProcessPercent(
		percent string,
	)
	ProcessPercentage(
		percentage string,
	)
	ProcessPlus(
		plus string,
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
	ProcessSlash(
		slash string,
	)
	ProcessSpace(
		space string,
	)
	ProcessStar(
		star string,
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
	ProcessAcceptClauseSlot(
		slot uint,
	)
	PostprocessAcceptClause(
		acceptClause ast.AcceptClauseLike,
		index_ uint,
		count_ uint,
	)
	PreprocessAdditionalArgument(
		additionalArgument ast.AdditionalArgumentLike,
		index_ uint,
		count_ uint,
	)
	ProcessAdditionalArgumentSlot(
		slot uint,
	)
	PostprocessAdditionalArgument(
		additionalArgument ast.AdditionalArgumentLike,
		index_ uint,
		count_ uint,
	)
	PreprocessAdditionalAssociation(
		additionalAssociation ast.AdditionalAssociationLike,
		index_ uint,
		count_ uint,
	)
	ProcessAdditionalAssociationSlot(
		slot uint,
	)
	PostprocessAdditionalAssociation(
		additionalAssociation ast.AdditionalAssociationLike,
		index_ uint,
		count_ uint,
	)
	PreprocessAdditionalIndex(
		additionalIndex ast.AdditionalIndexLike,
		index_ uint,
		count_ uint,
	)
	ProcessAdditionalIndexSlot(
		slot uint,
	)
	PostprocessAdditionalIndex(
		additionalIndex ast.AdditionalIndexLike,
		index_ uint,
		count_ uint,
	)
	PreprocessAdditionalStatement(
		additionalStatement ast.AdditionalStatementLike,
		index_ uint,
		count_ uint,
	)
	ProcessAdditionalStatementSlot(
		slot uint,
	)
	PostprocessAdditionalStatement(
		additionalStatement ast.AdditionalStatementLike,
		index_ uint,
		count_ uint,
	)
	PreprocessAdditionalValue(
		additionalValue ast.AdditionalValueLike,
		index_ uint,
		count_ uint,
	)
	ProcessAdditionalValueSlot(
		slot uint,
	)
	PostprocessAdditionalValue(
		additionalValue ast.AdditionalValueLike,
		index_ uint,
		count_ uint,
	)
	PreprocessAnnotatedAssociation(
		annotatedAssociation ast.AnnotatedAssociationLike,
		index_ uint,
		count_ uint,
	)
	ProcessAnnotatedAssociationSlot(
		slot uint,
	)
	PostprocessAnnotatedAssociation(
		annotatedAssociation ast.AnnotatedAssociationLike,
		index_ uint,
		count_ uint,
	)
	PreprocessAnnotatedStatement(
		annotatedStatement ast.AnnotatedStatementLike,
		index_ uint,
		count_ uint,
	)
	ProcessAnnotatedStatementSlot(
		slot uint,
	)
	PostprocessAnnotatedStatement(
		annotatedStatement ast.AnnotatedStatementLike,
		index_ uint,
		count_ uint,
	)
	PreprocessAnnotatedValue(
		annotatedValue ast.AnnotatedValueLike,
		index_ uint,
		count_ uint,
	)
	ProcessAnnotatedValueSlot(
		slot uint,
	)
	PostprocessAnnotatedValue(
		annotatedValue ast.AnnotatedValueLike,
		index_ uint,
		count_ uint,
	)
	PreprocessArgument(
		argument ast.ArgumentLike,
		index_ uint,
		count_ uint,
	)
	ProcessArgumentSlot(
		slot uint,
	)
	PostprocessArgument(
		argument ast.ArgumentLike,
		index_ uint,
		count_ uint,
	)
	PreprocessArguments(
		arguments ast.ArgumentsLike,
		index_ uint,
		count_ uint,
	)
	ProcessArgumentsSlot(
		slot uint,
	)
	PostprocessArguments(
		arguments ast.ArgumentsLike,
		index_ uint,
		count_ uint,
	)
	PreprocessArithmetic(
		arithmetic ast.ArithmeticLike,
		index_ uint,
		count_ uint,
	)
	ProcessArithmeticSlot(
		slot uint,
	)
	PostprocessArithmetic(
		arithmetic ast.ArithmeticLike,
		index_ uint,
		count_ uint,
	)
	PreprocessAssignment(
		assignment ast.AssignmentLike,
		index_ uint,
		count_ uint,
	)
	ProcessAssignmentSlot(
		slot uint,
	)
	PostprocessAssignment(
		assignment ast.AssignmentLike,
		index_ uint,
		count_ uint,
	)
	PreprocessAssociation(
		association ast.AssociationLike,
		index_ uint,
		count_ uint,
	)
	ProcessAssociationSlot(
		slot uint,
	)
	PostprocessAssociation(
		association ast.AssociationLike,
		index_ uint,
		count_ uint,
	)
	PreprocessAtLevel(
		atLevel ast.AtLevelLike,
		index_ uint,
		count_ uint,
	)
	ProcessAtLevelSlot(
		slot uint,
	)
	PostprocessAtLevel(
		atLevel ast.AtLevelLike,
		index_ uint,
		count_ uint,
	)
	PreprocessBag(
		bag ast.BagLike,
		index_ uint,
		count_ uint,
	)
	ProcessBagSlot(
		slot uint,
	)
	PostprocessBag(
		bag ast.BagLike,
		index_ uint,
		count_ uint,
	)
	PreprocessBlocking(
		blocking ast.BlockingLike,
		index_ uint,
		count_ uint,
	)
	ProcessBlockingSlot(
		slot uint,
	)
	PostprocessBlocking(
		blocking ast.BlockingLike,
		index_ uint,
		count_ uint,
	)
	PreprocessBracket(
		bracket ast.BracketLike,
		index_ uint,
		count_ uint,
	)
	ProcessBracketSlot(
		slot uint,
	)
	PostprocessBracket(
		bracket ast.BracketLike,
		index_ uint,
		count_ uint,
	)
	PreprocessBreakClause(
		breakClause ast.BreakClauseLike,
		index_ uint,
		count_ uint,
	)
	ProcessBreakClauseSlot(
		slot uint,
	)
	PostprocessBreakClause(
		breakClause ast.BreakClauseLike,
		index_ uint,
		count_ uint,
	)
	PreprocessCheckoutClause(
		checkoutClause ast.CheckoutClauseLike,
		index_ uint,
		count_ uint,
	)
	ProcessCheckoutClauseSlot(
		slot uint,
	)
	PostprocessCheckoutClause(
		checkoutClause ast.CheckoutClauseLike,
		index_ uint,
		count_ uint,
	)
	PreprocessCited(
		cited ast.CitedLike,
		index_ uint,
		count_ uint,
	)
	ProcessCitedSlot(
		slot uint,
	)
	PostprocessCited(
		cited ast.CitedLike,
		index_ uint,
		count_ uint,
	)
	PreprocessCollection(
		collection ast.CollectionLike,
		index_ uint,
		count_ uint,
	)
	ProcessCollectionSlot(
		slot uint,
	)
	PostprocessCollection(
		collection ast.CollectionLike,
		index_ uint,
		count_ uint,
	)
	PreprocessCommentLine(
		commentLine ast.CommentLineLike,
		index_ uint,
		count_ uint,
	)
	ProcessCommentLineSlot(
		slot uint,
	)
	PostprocessCommentLine(
		commentLine ast.CommentLineLike,
		index_ uint,
		count_ uint,
	)
	PreprocessComparison(
		comparison ast.ComparisonLike,
		index_ uint,
		count_ uint,
	)
	ProcessComparisonSlot(
		slot uint,
	)
	PostprocessComparison(
		comparison ast.ComparisonLike,
		index_ uint,
		count_ uint,
	)
	PreprocessComplement(
		complement ast.ComplementLike,
		index_ uint,
		count_ uint,
	)
	ProcessComplementSlot(
		slot uint,
	)
	PostprocessComplement(
		complement ast.ComplementLike,
		index_ uint,
		count_ uint,
	)
	PreprocessComponent(
		component ast.ComponentLike,
		index_ uint,
		count_ uint,
	)
	ProcessComponentSlot(
		slot uint,
	)
	PostprocessComponent(
		component ast.ComponentLike,
		index_ uint,
		count_ uint,
	)
	PreprocessCondition(
		condition ast.ConditionLike,
		index_ uint,
		count_ uint,
	)
	ProcessConditionSlot(
		slot uint,
	)
	PostprocessCondition(
		condition ast.ConditionLike,
		index_ uint,
		count_ uint,
	)
	PreprocessContinueClause(
		continueClause ast.ContinueClauseLike,
		index_ uint,
		count_ uint,
	)
	ProcessContinueClauseSlot(
		slot uint,
	)
	PostprocessContinueClause(
		continueClause ast.ContinueClauseLike,
		index_ uint,
		count_ uint,
	)
	PreprocessDiscardClause(
		discardClause ast.DiscardClauseLike,
		index_ uint,
		count_ uint,
	)
	ProcessDiscardClauseSlot(
		slot uint,
	)
	PostprocessDiscardClause(
		discardClause ast.DiscardClauseLike,
		index_ uint,
		count_ uint,
	)
	PreprocessDoClause(
		doClause ast.DoClauseLike,
		index_ uint,
		count_ uint,
	)
	ProcessDoClauseSlot(
		slot uint,
	)
	PostprocessDoClause(
		doClause ast.DoClauseLike,
		index_ uint,
		count_ uint,
	)
	PreprocessDocument(
		document ast.DocumentLike,
		index_ uint,
		count_ uint,
	)
	ProcessDocumentSlot(
		slot uint,
	)
	PostprocessDocument(
		document ast.DocumentLike,
		index_ uint,
		count_ uint,
	)
	PreprocessDraft(
		draft ast.DraftLike,
		index_ uint,
		count_ uint,
	)
	ProcessDraftSlot(
		slot uint,
	)
	PostprocessDraft(
		draft ast.DraftLike,
		index_ uint,
		count_ uint,
	)
	PreprocessElement(
		element ast.ElementLike,
		index_ uint,
		count_ uint,
	)
	ProcessElementSlot(
		slot uint,
	)
	PostprocessElement(
		element ast.ElementLike,
		index_ uint,
		count_ uint,
	)
	PreprocessEntity(
		entity ast.EntityLike,
		index_ uint,
		count_ uint,
	)
	ProcessEntitySlot(
		slot uint,
	)
	PostprocessEntity(
		entity ast.EntityLike,
		index_ uint,
		count_ uint,
	)
	PreprocessEvent(
		event ast.EventLike,
		index_ uint,
		count_ uint,
	)
	ProcessEventSlot(
		slot uint,
	)
	PostprocessEvent(
		event ast.EventLike,
		index_ uint,
		count_ uint,
	)
	PreprocessException(
		exception ast.ExceptionLike,
		index_ uint,
		count_ uint,
	)
	ProcessExceptionSlot(
		slot uint,
	)
	PostprocessException(
		exception ast.ExceptionLike,
		index_ uint,
		count_ uint,
	)
	PreprocessExclusive(
		exclusive ast.ExclusiveLike,
		index_ uint,
		count_ uint,
	)
	ProcessExclusiveSlot(
		slot uint,
	)
	PostprocessExclusive(
		exclusive ast.ExclusiveLike,
		index_ uint,
		count_ uint,
	)
	PreprocessExclusiveRange(
		exclusiveRange ast.ExclusiveRangeLike,
		index_ uint,
		count_ uint,
	)
	ProcessExclusiveRangeSlot(
		slot uint,
	)
	PostprocessExclusiveRange(
		exclusiveRange ast.ExclusiveRangeLike,
		index_ uint,
		count_ uint,
	)
	PreprocessExpression(
		expression ast.ExpressionLike,
		index_ uint,
		count_ uint,
	)
	ProcessExpressionSlot(
		slot uint,
	)
	PostprocessExpression(
		expression ast.ExpressionLike,
		index_ uint,
		count_ uint,
	)
	PreprocessFailure(
		failure ast.FailureLike,
		index_ uint,
		count_ uint,
	)
	ProcessFailureSlot(
		slot uint,
	)
	PostprocessFailure(
		failure ast.FailureLike,
		index_ uint,
		count_ uint,
	)
	PreprocessFlow(
		flow ast.FlowLike,
		index_ uint,
		count_ uint,
	)
	ProcessFlowSlot(
		slot uint,
	)
	PostprocessFlow(
		flow ast.FlowLike,
		index_ uint,
		count_ uint,
	)
	PreprocessFunction(
		function ast.FunctionLike,
		index_ uint,
		count_ uint,
	)
	ProcessFunctionSlot(
		slot uint,
	)
	PostprocessFunction(
		function ast.FunctionLike,
		index_ uint,
		count_ uint,
	)
	PreprocessIfClause(
		ifClause ast.IfClauseLike,
		index_ uint,
		count_ uint,
	)
	ProcessIfClauseSlot(
		slot uint,
	)
	PostprocessIfClause(
		ifClause ast.IfClauseLike,
		index_ uint,
		count_ uint,
	)
	PreprocessInclusive(
		inclusive ast.InclusiveLike,
		index_ uint,
		count_ uint,
	)
	ProcessInclusiveSlot(
		slot uint,
	)
	PostprocessInclusive(
		inclusive ast.InclusiveLike,
		index_ uint,
		count_ uint,
	)
	PreprocessInclusiveRange(
		inclusiveRange ast.InclusiveRangeLike,
		index_ uint,
		count_ uint,
	)
	ProcessInclusiveRangeSlot(
		slot uint,
	)
	PostprocessInclusiveRange(
		inclusiveRange ast.InclusiveRangeLike,
		index_ uint,
		count_ uint,
	)
	PreprocessIndex(
		index ast.IndexLike,
		index_ uint,
		count_ uint,
	)
	ProcessIndexSlot(
		slot uint,
	)
	PostprocessIndex(
		index ast.IndexLike,
		index_ uint,
		count_ uint,
	)
	PreprocessIndices(
		indices ast.IndicesLike,
		index_ uint,
		count_ uint,
	)
	ProcessIndicesSlot(
		slot uint,
	)
	PostprocessIndices(
		indices ast.IndicesLike,
		index_ uint,
		count_ uint,
	)
	PreprocessIndirect(
		indirect ast.IndirectLike,
		index_ uint,
		count_ uint,
	)
	ProcessIndirectSlot(
		slot uint,
	)
	PostprocessIndirect(
		indirect ast.IndirectLike,
		index_ uint,
		count_ uint,
	)
	PreprocessInduction(
		induction ast.InductionLike,
		index_ uint,
		count_ uint,
	)
	ProcessInductionSlot(
		slot uint,
	)
	PostprocessInduction(
		induction ast.InductionLike,
		index_ uint,
		count_ uint,
	)
	PreprocessInlineAttributes(
		inlineAttributes ast.InlineAttributesLike,
		index_ uint,
		count_ uint,
	)
	ProcessInlineAttributesSlot(
		slot uint,
	)
	PostprocessInlineAttributes(
		inlineAttributes ast.InlineAttributesLike,
		index_ uint,
		count_ uint,
	)
	PreprocessInlineParameters(
		inlineParameters ast.InlineParametersLike,
		index_ uint,
		count_ uint,
	)
	ProcessInlineParametersSlot(
		slot uint,
	)
	PostprocessInlineParameters(
		inlineParameters ast.InlineParametersLike,
		index_ uint,
		count_ uint,
	)
	PreprocessInlineStatements(
		inlineStatements ast.InlineStatementsLike,
		index_ uint,
		count_ uint,
	)
	ProcessInlineStatementsSlot(
		slot uint,
	)
	PostprocessInlineStatements(
		inlineStatements ast.InlineStatementsLike,
		index_ uint,
		count_ uint,
	)
	PreprocessInlineValues(
		inlineValues ast.InlineValuesLike,
		index_ uint,
		count_ uint,
	)
	ProcessInlineValuesSlot(
		slot uint,
	)
	PostprocessInlineValues(
		inlineValues ast.InlineValuesLike,
		index_ uint,
		count_ uint,
	)
	PreprocessInverse(
		inverse ast.InverseLike,
		index_ uint,
		count_ uint,
	)
	ProcessInverseSlot(
		slot uint,
	)
	PostprocessInverse(
		inverse ast.InverseLike,
		index_ uint,
		count_ uint,
	)
	PreprocessInversion(
		inversion ast.InversionLike,
		index_ uint,
		count_ uint,
	)
	ProcessInversionSlot(
		slot uint,
	)
	PostprocessInversion(
		inversion ast.InversionLike,
		index_ uint,
		count_ uint,
	)
	PreprocessInvocation(
		invocation ast.InvocationLike,
		index_ uint,
		count_ uint,
	)
	ProcessInvocationSlot(
		slot uint,
	)
	PostprocessInvocation(
		invocation ast.InvocationLike,
		index_ uint,
		count_ uint,
	)
	PreprocessItem(
		item ast.ItemLike,
		index_ uint,
		count_ uint,
	)
	ProcessItemSlot(
		slot uint,
	)
	PostprocessItem(
		item ast.ItemLike,
		index_ uint,
		count_ uint,
	)
	PreprocessLegalNotice(
		legalNotice ast.LegalNoticeLike,
		index_ uint,
		count_ uint,
	)
	ProcessLegalNoticeSlot(
		slot uint,
	)
	PostprocessLegalNotice(
		legalNotice ast.LegalNoticeLike,
		index_ uint,
		count_ uint,
	)
	PreprocessLetClause(
		letClause ast.LetClauseLike,
		index_ uint,
		count_ uint,
	)
	ProcessLetClauseSlot(
		slot uint,
	)
	PostprocessLetClause(
		letClause ast.LetClauseLike,
		index_ uint,
		count_ uint,
	)
	PreprocessLogic(
		logic ast.LogicLike,
		index_ uint,
		count_ uint,
	)
	ProcessLogicSlot(
		slot uint,
	)
	PostprocessLogic(
		logic ast.LogicLike,
		index_ uint,
		count_ uint,
	)
	PreprocessLogical(
		logical ast.LogicalLike,
		index_ uint,
		count_ uint,
	)
	ProcessLogicalSlot(
		slot uint,
	)
	PostprocessLogical(
		logical ast.LogicalLike,
		index_ uint,
		count_ uint,
	)
	PreprocessMagnitude(
		magnitude ast.MagnitudeLike,
		index_ uint,
		count_ uint,
	)
	ProcessMagnitudeSlot(
		slot uint,
	)
	PostprocessMagnitude(
		magnitude ast.MagnitudeLike,
		index_ uint,
		count_ uint,
	)
	PreprocessMainClause(
		mainClause ast.MainClauseLike,
		index_ uint,
		count_ uint,
	)
	ProcessMainClauseSlot(
		slot uint,
	)
	PostprocessMainClause(
		mainClause ast.MainClauseLike,
		index_ uint,
		count_ uint,
	)
	PreprocessMatchHandler(
		matchHandler ast.MatchHandlerLike,
		index_ uint,
		count_ uint,
	)
	ProcessMatchHandlerSlot(
		slot uint,
	)
	PostprocessMatchHandler(
		matchHandler ast.MatchHandlerLike,
		index_ uint,
		count_ uint,
	)
	PreprocessMessage(
		message ast.MessageLike,
		index_ uint,
		count_ uint,
	)
	ProcessMessageSlot(
		slot uint,
	)
	PostprocessMessage(
		message ast.MessageLike,
		index_ uint,
		count_ uint,
	)
	PreprocessMessaging(
		messaging ast.MessagingLike,
		index_ uint,
		count_ uint,
	)
	ProcessMessagingSlot(
		slot uint,
	)
	PostprocessMessaging(
		messaging ast.MessagingLike,
		index_ uint,
		count_ uint,
	)
	PreprocessMethod(
		method ast.MethodLike,
		index_ uint,
		count_ uint,
	)
	ProcessMethodSlot(
		slot uint,
	)
	PostprocessMethod(
		method ast.MethodLike,
		index_ uint,
		count_ uint,
	)
	PreprocessMultilineAttributes(
		multilineAttributes ast.MultilineAttributesLike,
		index_ uint,
		count_ uint,
	)
	ProcessMultilineAttributesSlot(
		slot uint,
	)
	PostprocessMultilineAttributes(
		multilineAttributes ast.MultilineAttributesLike,
		index_ uint,
		count_ uint,
	)
	PreprocessMultilineParameters(
		multilineParameters ast.MultilineParametersLike,
		index_ uint,
		count_ uint,
	)
	ProcessMultilineParametersSlot(
		slot uint,
	)
	PostprocessMultilineParameters(
		multilineParameters ast.MultilineParametersLike,
		index_ uint,
		count_ uint,
	)
	PreprocessMultilineStatements(
		multilineStatements ast.MultilineStatementsLike,
		index_ uint,
		count_ uint,
	)
	ProcessMultilineStatementsSlot(
		slot uint,
	)
	PostprocessMultilineStatements(
		multilineStatements ast.MultilineStatementsLike,
		index_ uint,
		count_ uint,
	)
	PreprocessMultilineValues(
		multilineValues ast.MultilineValuesLike,
		index_ uint,
		count_ uint,
	)
	ProcessMultilineValuesSlot(
		slot uint,
	)
	PostprocessMultilineValues(
		multilineValues ast.MultilineValuesLike,
		index_ uint,
		count_ uint,
	)
	PreprocessNoAttributes(
		noAttributes ast.NoAttributesLike,
		index_ uint,
		count_ uint,
	)
	ProcessNoAttributesSlot(
		slot uint,
	)
	PostprocessNoAttributes(
		noAttributes ast.NoAttributesLike,
		index_ uint,
		count_ uint,
	)
	PreprocessNoStatements(
		noStatements ast.NoStatementsLike,
		index_ uint,
		count_ uint,
	)
	ProcessNoStatementsSlot(
		slot uint,
	)
	PostprocessNoStatements(
		noStatements ast.NoStatementsLike,
		index_ uint,
		count_ uint,
	)
	PreprocessNoValues(
		noValues ast.NoValuesLike,
		index_ uint,
		count_ uint,
	)
	ProcessNoValuesSlot(
		slot uint,
	)
	PostprocessNoValues(
		noValues ast.NoValuesLike,
		index_ uint,
		count_ uint,
	)
	PreprocessNotarizeClause(
		notarizeClause ast.NotarizeClauseLike,
		index_ uint,
		count_ uint,
	)
	ProcessNotarizeClauseSlot(
		slot uint,
	)
	PostprocessNotarizeClause(
		notarizeClause ast.NotarizeClauseLike,
		index_ uint,
		count_ uint,
	)
	PreprocessNumerical(
		numerical ast.NumericalLike,
		index_ uint,
		count_ uint,
	)
	ProcessNumericalSlot(
		slot uint,
	)
	PostprocessNumerical(
		numerical ast.NumericalLike,
		index_ uint,
		count_ uint,
	)
	PreprocessOnClause(
		onClause ast.OnClauseLike,
		index_ uint,
		count_ uint,
	)
	ProcessOnClauseSlot(
		slot uint,
	)
	PostprocessOnClause(
		onClause ast.OnClauseLike,
		index_ uint,
		count_ uint,
	)
	PreprocessOperation(
		operation ast.OperationLike,
		index_ uint,
		count_ uint,
	)
	ProcessOperationSlot(
		slot uint,
	)
	PostprocessOperation(
		operation ast.OperationLike,
		index_ uint,
		count_ uint,
	)
	PreprocessParameters(
		parameters ast.ParametersLike,
		index_ uint,
		count_ uint,
	)
	ProcessParametersSlot(
		slot uint,
	)
	PostprocessParameters(
		parameters ast.ParametersLike,
		index_ uint,
		count_ uint,
	)
	PreprocessPostClause(
		postClause ast.PostClauseLike,
		index_ uint,
		count_ uint,
	)
	ProcessPostClauseSlot(
		slot uint,
	)
	PostprocessPostClause(
		postClause ast.PostClauseLike,
		index_ uint,
		count_ uint,
	)
	PreprocessPrecedence(
		precedence ast.PrecedenceLike,
		index_ uint,
		count_ uint,
	)
	ProcessPrecedenceSlot(
		slot uint,
	)
	PostprocessPrecedence(
		precedence ast.PrecedenceLike,
		index_ uint,
		count_ uint,
	)
	PreprocessPredicate(
		predicate ast.PredicateLike,
		index_ uint,
		count_ uint,
	)
	ProcessPredicateSlot(
		slot uint,
	)
	PostprocessPredicate(
		predicate ast.PredicateLike,
		index_ uint,
		count_ uint,
	)
	PreprocessPrimitive(
		primitive ast.PrimitiveLike,
		index_ uint,
		count_ uint,
	)
	ProcessPrimitiveSlot(
		slot uint,
	)
	PostprocessPrimitive(
		primitive ast.PrimitiveLike,
		index_ uint,
		count_ uint,
	)
	PreprocessProcedure(
		procedure ast.ProcedureLike,
		index_ uint,
		count_ uint,
	)
	ProcessProcedureSlot(
		slot uint,
	)
	PostprocessProcedure(
		procedure ast.ProcedureLike,
		index_ uint,
		count_ uint,
	)
	PreprocessPublishClause(
		publishClause ast.PublishClauseLike,
		index_ uint,
		count_ uint,
	)
	ProcessPublishClauseSlot(
		slot uint,
	)
	PostprocessPublishClause(
		publishClause ast.PublishClauseLike,
		index_ uint,
		count_ uint,
	)
	PreprocessRecipient(
		recipient ast.RecipientLike,
		index_ uint,
		count_ uint,
	)
	ProcessRecipientSlot(
		slot uint,
	)
	PostprocessRecipient(
		recipient ast.RecipientLike,
		index_ uint,
		count_ uint,
	)
	PreprocessReferent(
		referent ast.ReferentLike,
		index_ uint,
		count_ uint,
	)
	ProcessReferentSlot(
		slot uint,
	)
	PostprocessReferent(
		referent ast.ReferentLike,
		index_ uint,
		count_ uint,
	)
	PreprocessRejectClause(
		rejectClause ast.RejectClauseLike,
		index_ uint,
		count_ uint,
	)
	ProcessRejectClauseSlot(
		slot uint,
	)
	PostprocessRejectClause(
		rejectClause ast.RejectClauseLike,
		index_ uint,
		count_ uint,
	)
	PreprocessRepository(
		repository ast.RepositoryLike,
		index_ uint,
		count_ uint,
	)
	ProcessRepositorySlot(
		slot uint,
	)
	PostprocessRepository(
		repository ast.RepositoryLike,
		index_ uint,
		count_ uint,
	)
	PreprocessResult(
		result ast.ResultLike,
		index_ uint,
		count_ uint,
	)
	ProcessResultSlot(
		slot uint,
	)
	PostprocessResult(
		result ast.ResultLike,
		index_ uint,
		count_ uint,
	)
	PreprocessRetrieveClause(
		retrieveClause ast.RetrieveClauseLike,
		index_ uint,
		count_ uint,
	)
	ProcessRetrieveClauseSlot(
		slot uint,
	)
	PostprocessRetrieveClause(
		retrieveClause ast.RetrieveClauseLike,
		index_ uint,
		count_ uint,
	)
	PreprocessReturnClause(
		returnClause ast.ReturnClauseLike,
		index_ uint,
		count_ uint,
	)
	ProcessReturnClauseSlot(
		slot uint,
	)
	PostprocessReturnClause(
		returnClause ast.ReturnClauseLike,
		index_ uint,
		count_ uint,
	)
	PreprocessSaveClause(
		saveClause ast.SaveClauseLike,
		index_ uint,
		count_ uint,
	)
	ProcessSaveClauseSlot(
		slot uint,
	)
	PostprocessSaveClause(
		saveClause ast.SaveClauseLike,
		index_ uint,
		count_ uint,
	)
	PreprocessSelectClause(
		selectClause ast.SelectClauseLike,
		index_ uint,
		count_ uint,
	)
	ProcessSelectClauseSlot(
		slot uint,
	)
	PostprocessSelectClause(
		selectClause ast.SelectClauseLike,
		index_ uint,
		count_ uint,
	)
	PreprocessSequence(
		sequence ast.SequenceLike,
		index_ uint,
		count_ uint,
	)
	ProcessSequenceSlot(
		slot uint,
	)
	PostprocessSequence(
		sequence ast.SequenceLike,
		index_ uint,
		count_ uint,
	)
	PreprocessStatement(
		statement ast.StatementLike,
		index_ uint,
		count_ uint,
	)
	ProcessStatementSlot(
		slot uint,
	)
	PostprocessStatement(
		statement ast.StatementLike,
		index_ uint,
		count_ uint,
	)
	PreprocessStatementLine(
		statementLine ast.StatementLineLike,
		index_ uint,
		count_ uint,
	)
	ProcessStatementLineSlot(
		slot uint,
	)
	PostprocessStatementLine(
		statementLine ast.StatementLineLike,
		index_ uint,
		count_ uint,
	)
	PreprocessString(
		string_ ast.StringLike,
		index_ uint,
		count_ uint,
	)
	ProcessStringSlot(
		slot uint,
	)
	PostprocessString(
		string_ ast.StringLike,
		index_ uint,
		count_ uint,
	)
	PreprocessSubcomponent(
		subcomponent ast.SubcomponentLike,
		index_ uint,
		count_ uint,
	)
	ProcessSubcomponentSlot(
		slot uint,
	)
	PostprocessSubcomponent(
		subcomponent ast.SubcomponentLike,
		index_ uint,
		count_ uint,
	)
	PreprocessSubject(
		subject ast.SubjectLike,
		index_ uint,
		count_ uint,
	)
	ProcessSubjectSlot(
		slot uint,
	)
	PostprocessSubject(
		subject ast.SubjectLike,
		index_ uint,
		count_ uint,
	)
	PreprocessTarget(
		target ast.TargetLike,
		index_ uint,
		count_ uint,
	)
	ProcessTargetSlot(
		slot uint,
	)
	PostprocessTarget(
		target ast.TargetLike,
		index_ uint,
		count_ uint,
	)
	PreprocessTemplate(
		template ast.TemplateLike,
		index_ uint,
		count_ uint,
	)
	ProcessTemplateSlot(
		slot uint,
	)
	PostprocessTemplate(
		template ast.TemplateLike,
		index_ uint,
		count_ uint,
	)
	PreprocessTextual(
		textual ast.TextualLike,
		index_ uint,
		count_ uint,
	)
	ProcessTextualSlot(
		slot uint,
	)
	PostprocessTextual(
		textual ast.TextualLike,
		index_ uint,
		count_ uint,
	)
	PreprocessThrowClause(
		throwClause ast.ThrowClauseLike,
		index_ uint,
		count_ uint,
	)
	ProcessThrowClauseSlot(
		slot uint,
	)
	PostprocessThrowClause(
		throwClause ast.ThrowClauseLike,
		index_ uint,
		count_ uint,
	)
	PreprocessVariable(
		variable ast.VariableLike,
		index_ uint,
		count_ uint,
	)
	ProcessVariableSlot(
		slot uint,
	)
	PostprocessVariable(
		variable ast.VariableLike,
		index_ uint,
		count_ uint,
	)
	PreprocessWhileClause(
		whileClause ast.WhileClauseLike,
		index_ uint,
		count_ uint,
	)
	ProcessWhileClauseSlot(
		slot uint,
	)
	PostprocessWhileClause(
		whileClause ast.WhileClauseLike,
		index_ uint,
		count_ uint,
	)
	PreprocessWithClause(
		withClause ast.WithClauseLike,
		index_ uint,
		count_ uint,
	)
	ProcessWithClauseSlot(
		slot uint,
	)
	PostprocessWithClause(
		withClause ast.WithClauseLike,
		index_ uint,
		count_ uint,
	)
}
