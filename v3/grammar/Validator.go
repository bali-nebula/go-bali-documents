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
│                 This class file was automatically generated.                 │
│ Updates to any section other than the Methodical Methods may be overwritten. │
└──────────────────────────────────────────────────────────────────────────────┘
*/

package grammar

import (
	fmt "fmt"
	ast "github.com/bali-nebula/go-bali-documents/v3/ast"
)

// CLASS INTERFACE

// Access Function

func ValidatorClass() ValidatorClassLike {
	return validatorClass()
}

// Constructor Methods

func (c *validatorClass_) Validator() ValidatorLike {
	var instance = &validator_{
		// Initialize the instance attributes.

		// Initialize the inherited aspects.
		Methodical: ProcessorClass().Processor(),
	}
	instance.visitor_ = VisitorClass().Visitor(instance)
	return instance
}

// INSTANCE INTERFACE

// Principal Methods

func (v *validator_) GetClass() ValidatorClassLike {
	return validatorClass()
}

func (v *validator_) ValidateDocument(
	document ast.DocumentLike,
) {
	v.visitor_.VisitDocument(document)
}

// Methodical Methods

func (v *validator_) ProcessAngle(
	angle string,
) {
	v.validateToken(angle, AngleToken)
}

func (v *validator_) ProcessArrow(
	arrow string,
) {
	v.validateToken(arrow, ArrowToken)
}

func (v *validator_) ProcessBinary(
	binary string,
) {
	v.validateToken(binary, BinaryToken)
}

func (v *validator_) ProcessBoolean(
	boolean string,
) {
	v.validateToken(boolean, BooleanToken)
}

func (v *validator_) ProcessBytecode(
	bytecode string,
) {
	v.validateToken(bytecode, BytecodeToken)
}

func (v *validator_) ProcessCaret(
	caret string,
) {
	v.validateToken(caret, CaretToken)
}

func (v *validator_) ProcessCitation(
	citation string,
) {
	v.validateToken(citation, CitationToken)
}

func (v *validator_) ProcessComment(
	comment string,
) {
	v.validateToken(comment, CommentToken)
}

func (v *validator_) ProcessDot(
	dot string,
) {
	v.validateToken(dot, DotToken)
}

func (v *validator_) ProcessDuration(
	duration string,
) {
	v.validateToken(duration, DurationToken)
}

func (v *validator_) ProcessEqual(
	equal string,
) {
	v.validateToken(equal, EqualToken)
}

func (v *validator_) ProcessIdentifier(
	identifier string,
) {
	v.validateToken(identifier, IdentifierToken)
}

func (v *validator_) ProcessLess(
	less string,
) {
	v.validateToken(less, LessToken)
}

func (v *validator_) ProcessMinus(
	minus string,
) {
	v.validateToken(minus, MinusToken)
}

func (v *validator_) ProcessMoment(
	moment string,
) {
	v.validateToken(moment, MomentToken)
}

func (v *validator_) ProcessMore(
	more string,
) {
	v.validateToken(more, MoreToken)
}

func (v *validator_) ProcessName(
	name string,
) {
	v.validateToken(name, NameToken)
}

func (v *validator_) ProcessNarrative(
	narrative string,
) {
	v.validateToken(narrative, NarrativeToken)
}

func (v *validator_) ProcessNewline(
	newline string,
) {
	v.validateToken(newline, NewlineToken)
}

func (v *validator_) ProcessNote(
	note string,
) {
	v.validateToken(note, NoteToken)
}

func (v *validator_) ProcessNumber(
	number string,
) {
	v.validateToken(number, NumberToken)
}

func (v *validator_) ProcessPattern(
	pattern string,
) {
	v.validateToken(pattern, PatternToken)
}

func (v *validator_) ProcessPercent(
	percent string,
) {
	v.validateToken(percent, PercentToken)
}

func (v *validator_) ProcessPercentage(
	percentage string,
) {
	v.validateToken(percentage, PercentageToken)
}

func (v *validator_) ProcessPlus(
	plus string,
) {
	v.validateToken(plus, PlusToken)
}

func (v *validator_) ProcessProbability(
	probability string,
) {
	v.validateToken(probability, ProbabilityToken)
}

func (v *validator_) ProcessQuote(
	quote string,
) {
	v.validateToken(quote, QuoteToken)
}

func (v *validator_) ProcessResource(
	resource string,
) {
	v.validateToken(resource, ResourceToken)
}

func (v *validator_) ProcessSlash(
	slash string,
) {
	v.validateToken(slash, SlashToken)
}

func (v *validator_) ProcessSpace(
	space string,
) {
	v.validateToken(space, SpaceToken)
}

func (v *validator_) ProcessStar(
	star string,
) {
	v.validateToken(star, StarToken)
}

func (v *validator_) ProcessSymbol(
	symbol string,
) {
	v.validateToken(symbol, SymbolToken)
}

func (v *validator_) ProcessTag(
	tag string,
) {
	v.validateToken(tag, TagToken)
}

func (v *validator_) ProcessVersion(
	version string,
) {
	v.validateToken(version, VersionToken)
}

func (v *validator_) PreprocessAcceptClause(
	acceptClause ast.AcceptClauseLike,
	index_ uint,
	count_ uint,
) {
	// TBD - Add any validation checks.
}

func (v *validator_) PreprocessAdditionalArgument(
	additionalArgument ast.AdditionalArgumentLike,
	index_ uint,
	count_ uint,
) {
	// TBD - Add any validation checks.
}

func (v *validator_) PreprocessAdditionalAssociation(
	additionalAssociation ast.AdditionalAssociationLike,
	index_ uint,
	count_ uint,
) {
	// TBD - Add any validation checks.
}

func (v *validator_) PreprocessAdditionalIndex(
	additionalIndex ast.AdditionalIndexLike,
	index_ uint,
	count_ uint,
) {
	// TBD - Add any validation checks.
}

func (v *validator_) PreprocessAdditionalStatement(
	additionalStatement ast.AdditionalStatementLike,
	index_ uint,
	count_ uint,
) {
	// TBD - Add any validation checks.
}

func (v *validator_) PreprocessAdditionalValue(
	additionalValue ast.AdditionalValueLike,
	index_ uint,
	count_ uint,
) {
	// TBD - Add any validation checks.
}

func (v *validator_) PreprocessAnnotatedAssociation(
	annotatedAssociation ast.AnnotatedAssociationLike,
	index_ uint,
	count_ uint,
) {
	// TBD - Add any validation checks.
}

func (v *validator_) PreprocessAnnotatedStatement(
	annotatedStatement ast.AnnotatedStatementLike,
	index_ uint,
	count_ uint,
) {
	// TBD - Add any validation checks.
}

func (v *validator_) PreprocessAnnotatedValue(
	annotatedValue ast.AnnotatedValueLike,
	index_ uint,
	count_ uint,
) {
	// TBD - Add any validation checks.
}

func (v *validator_) PreprocessArgument(
	argument ast.ArgumentLike,
	index_ uint,
	count_ uint,
) {
	// TBD - Add any validation checks.
}

func (v *validator_) PreprocessArguments(
	arguments ast.ArgumentsLike,
	index_ uint,
	count_ uint,
) {
	// TBD - Add any validation checks.
}

func (v *validator_) PreprocessArithmetic(
	arithmetic ast.ArithmeticLike,
	index_ uint,
	count_ uint,
) {
	// TBD - Add any validation checks.
}

func (v *validator_) PreprocessAssignment(
	assignment ast.AssignmentLike,
	index_ uint,
	count_ uint,
) {
	// TBD - Add any validation checks.
}

func (v *validator_) PreprocessAssociation(
	association ast.AssociationLike,
	index_ uint,
	count_ uint,
) {
	// TBD - Add any validation checks.
}

func (v *validator_) PreprocessAtLevel(
	atLevel ast.AtLevelLike,
	index_ uint,
	count_ uint,
) {
	// TBD - Add any validation checks.
}

func (v *validator_) PreprocessBag(
	bag ast.BagLike,
	index_ uint,
	count_ uint,
) {
	// TBD - Add any validation checks.
}

func (v *validator_) PreprocessBlocking(
	blocking ast.BlockingLike,
	index_ uint,
	count_ uint,
) {
	// TBD - Add any validation checks.
}

func (v *validator_) PreprocessBracket(
	bracket ast.BracketLike,
	index_ uint,
	count_ uint,
) {
	// TBD - Add any validation checks.
}

func (v *validator_) PreprocessBreakClause(
	breakClause ast.BreakClauseLike,
	index_ uint,
	count_ uint,
) {
	// TBD - Add any validation checks.
}

func (v *validator_) PreprocessCheckoutClause(
	checkoutClause ast.CheckoutClauseLike,
	index_ uint,
	count_ uint,
) {
	// TBD - Add any validation checks.
}

func (v *validator_) PreprocessCited(
	cited ast.CitedLike,
	index_ uint,
	count_ uint,
) {
	// TBD - Add any validation checks.
}

func (v *validator_) PreprocessCollection(
	collection ast.CollectionLike,
	index_ uint,
	count_ uint,
) {
	// TBD - Add any validation checks.
}

func (v *validator_) PreprocessCommentLine(
	commentLine ast.CommentLineLike,
	index_ uint,
	count_ uint,
) {
	// TBD - Add any validation checks.
}

func (v *validator_) PreprocessComparison(
	comparison ast.ComparisonLike,
	index_ uint,
	count_ uint,
) {
	// TBD - Add any validation checks.
}

func (v *validator_) PreprocessComplement(
	complement ast.ComplementLike,
	index_ uint,
	count_ uint,
) {
	// TBD - Add any validation checks.
}

func (v *validator_) PreprocessComponent(
	component ast.ComponentLike,
	index_ uint,
	count_ uint,
) {
	// TBD - Add any validation checks.
}

func (v *validator_) PreprocessCondition(
	condition ast.ConditionLike,
	index_ uint,
	count_ uint,
) {
	// TBD - Add any validation checks.
}

func (v *validator_) PreprocessContinueClause(
	continueClause ast.ContinueClauseLike,
	index_ uint,
	count_ uint,
) {
	// TBD - Add any validation checks.
}

func (v *validator_) PreprocessDiscardClause(
	discardClause ast.DiscardClauseLike,
	index_ uint,
	count_ uint,
) {
	// TBD - Add any validation checks.
}

func (v *validator_) PreprocessDoClause(
	doClause ast.DoClauseLike,
	index_ uint,
	count_ uint,
) {
	// TBD - Add any validation checks.
}

func (v *validator_) PreprocessDocument(
	document ast.DocumentLike,
	index_ uint,
	count_ uint,
) {
	// TBD - Add any validation checks.
}

func (v *validator_) PreprocessDraft(
	draft ast.DraftLike,
	index_ uint,
	count_ uint,
) {
	// TBD - Add any validation checks.
}

func (v *validator_) PreprocessElement(
	element ast.ElementLike,
	index_ uint,
	count_ uint,
) {
	// TBD - Add any validation checks.
}

func (v *validator_) PreprocessEntity(
	entity ast.EntityLike,
	index_ uint,
	count_ uint,
) {
	// TBD - Add any validation checks.
}

func (v *validator_) PreprocessEvent(
	event ast.EventLike,
	index_ uint,
	count_ uint,
) {
	// TBD - Add any validation checks.
}

func (v *validator_) PreprocessException(
	exception ast.ExceptionLike,
	index_ uint,
	count_ uint,
) {
	// TBD - Add any validation checks.
}

func (v *validator_) PreprocessExclusive(
	exclusive ast.ExclusiveLike,
	index_ uint,
	count_ uint,
) {
	// TBD - Add any validation checks.
}

func (v *validator_) PreprocessExclusiveRange(
	exclusiveRange ast.ExclusiveRangeLike,
	index_ uint,
	count_ uint,
) {
	// TBD - Add any validation checks.
}

func (v *validator_) PreprocessExpression(
	expression ast.ExpressionLike,
	index_ uint,
	count_ uint,
) {
	// TBD - Add any validation checks.
}

func (v *validator_) PreprocessFailure(
	failure ast.FailureLike,
	index_ uint,
	count_ uint,
) {
	// TBD - Add any validation checks.
}

func (v *validator_) PreprocessFlow(
	flow ast.FlowLike,
	index_ uint,
	count_ uint,
) {
	// TBD - Add any validation checks.
}

func (v *validator_) PreprocessFunction(
	function ast.FunctionLike,
	index_ uint,
	count_ uint,
) {
	// TBD - Add any validation checks.
}

func (v *validator_) PreprocessIfClause(
	ifClause ast.IfClauseLike,
	index_ uint,
	count_ uint,
) {
	// TBD - Add any validation checks.
}

func (v *validator_) PreprocessInclusive(
	inclusive ast.InclusiveLike,
	index_ uint,
	count_ uint,
) {
	// TBD - Add any validation checks.
}

func (v *validator_) PreprocessInclusiveRange(
	inclusiveRange ast.InclusiveRangeLike,
	index_ uint,
	count_ uint,
) {
	// TBD - Add any validation checks.
}

func (v *validator_) PreprocessIndex(
	index ast.IndexLike,
	index_ uint,
	count_ uint,
) {
	// TBD - Add any validation checks.
}

func (v *validator_) PreprocessIndices(
	indices ast.IndicesLike,
	index_ uint,
	count_ uint,
) {
	// TBD - Add any validation checks.
}

func (v *validator_) PreprocessIndirect(
	indirect ast.IndirectLike,
	index_ uint,
	count_ uint,
) {
	// TBD - Add any validation checks.
}

func (v *validator_) PreprocessInduction(
	induction ast.InductionLike,
	index_ uint,
	count_ uint,
) {
	// TBD - Add any validation checks.
}

func (v *validator_) PreprocessInlineAttributes(
	inlineAttributes ast.InlineAttributesLike,
	index_ uint,
	count_ uint,
) {
	// TBD - Add any validation checks.
}

func (v *validator_) PreprocessInlineParameters(
	inlineParameters ast.InlineParametersLike,
	index_ uint,
	count_ uint,
) {
	// TBD - Add any validation checks.
}

func (v *validator_) PreprocessInlineStatements(
	inlineStatements ast.InlineStatementsLike,
	index_ uint,
	count_ uint,
) {
	// TBD - Add any validation checks.
}

func (v *validator_) PreprocessInlineValues(
	inlineValues ast.InlineValuesLike,
	index_ uint,
	count_ uint,
) {
	// TBD - Add any validation checks.
}

func (v *validator_) PreprocessInverse(
	inverse ast.InverseLike,
	index_ uint,
	count_ uint,
) {
	// TBD - Add any validation checks.
}

func (v *validator_) PreprocessInversion(
	inversion ast.InversionLike,
	index_ uint,
	count_ uint,
) {
	// TBD - Add any validation checks.
}

func (v *validator_) PreprocessInvocation(
	invocation ast.InvocationLike,
	index_ uint,
	count_ uint,
) {
	// TBD - Add any validation checks.
}

func (v *validator_) PreprocessItem(
	item ast.ItemLike,
	index_ uint,
	count_ uint,
) {
	// TBD - Add any validation checks.
}

func (v *validator_) PreprocessLegalNotice(
	legalNotice ast.LegalNoticeLike,
	index_ uint,
	count_ uint,
) {
	// TBD - Add any validation checks.
}

func (v *validator_) PreprocessLetClause(
	letClause ast.LetClauseLike,
	index_ uint,
	count_ uint,
) {
	// TBD - Add any validation checks.
}

func (v *validator_) PreprocessLogic(
	logic ast.LogicLike,
	index_ uint,
	count_ uint,
) {
	// TBD - Add any validation checks.
}

func (v *validator_) PreprocessLogical(
	logical ast.LogicalLike,
	index_ uint,
	count_ uint,
) {
	// TBD - Add any validation checks.
}

func (v *validator_) PreprocessMagnitude(
	magnitude ast.MagnitudeLike,
	index_ uint,
	count_ uint,
) {
	// TBD - Add any validation checks.
}

func (v *validator_) PreprocessMainClause(
	mainClause ast.MainClauseLike,
	index_ uint,
	count_ uint,
) {
	// TBD - Add any validation checks.
}

func (v *validator_) PreprocessMatchHandler(
	matchHandler ast.MatchHandlerLike,
	index_ uint,
	count_ uint,
) {
	// TBD - Add any validation checks.
}

func (v *validator_) PreprocessMessage(
	message ast.MessageLike,
	index_ uint,
	count_ uint,
) {
	// TBD - Add any validation checks.
}

func (v *validator_) PreprocessMessaging(
	messaging ast.MessagingLike,
	index_ uint,
	count_ uint,
) {
	// TBD - Add any validation checks.
}

func (v *validator_) PreprocessMethod(
	method ast.MethodLike,
	index_ uint,
	count_ uint,
) {
	// TBD - Add any validation checks.
}

func (v *validator_) PreprocessMultilineAttributes(
	multilineAttributes ast.MultilineAttributesLike,
	index_ uint,
	count_ uint,
) {
	// TBD - Add any validation checks.
}

func (v *validator_) PreprocessMultilineParameters(
	multilineParameters ast.MultilineParametersLike,
	index_ uint,
	count_ uint,
) {
	// TBD - Add any validation checks.
}

func (v *validator_) PreprocessMultilineStatements(
	multilineStatements ast.MultilineStatementsLike,
	index_ uint,
	count_ uint,
) {
	// TBD - Add any validation checks.
}

func (v *validator_) PreprocessMultilineValues(
	multilineValues ast.MultilineValuesLike,
	index_ uint,
	count_ uint,
) {
	// TBD - Add any validation checks.
}

func (v *validator_) PreprocessNoAttributes(
	noAttributes ast.NoAttributesLike,
	index_ uint,
	count_ uint,
) {
	// TBD - Add any validation checks.
}

func (v *validator_) PreprocessNoStatements(
	noStatements ast.NoStatementsLike,
	index_ uint,
	count_ uint,
) {
	// TBD - Add any validation checks.
}

func (v *validator_) PreprocessNoValues(
	noValues ast.NoValuesLike,
	index_ uint,
	count_ uint,
) {
	// TBD - Add any validation checks.
}

func (v *validator_) PreprocessNotarizeClause(
	notarizeClause ast.NotarizeClauseLike,
	index_ uint,
	count_ uint,
) {
	// TBD - Add any validation checks.
}

func (v *validator_) PreprocessNumerical(
	numerical ast.NumericalLike,
	index_ uint,
	count_ uint,
) {
	// TBD - Add any validation checks.
}

func (v *validator_) PreprocessOnClause(
	onClause ast.OnClauseLike,
	index_ uint,
	count_ uint,
) {
	// TBD - Add any validation checks.
}

func (v *validator_) PreprocessOperation(
	operation ast.OperationLike,
	index_ uint,
	count_ uint,
) {
	// TBD - Add any validation checks.
}

func (v *validator_) PreprocessParameters(
	parameters ast.ParametersLike,
	index_ uint,
	count_ uint,
) {
	// TBD - Add any validation checks.
}

func (v *validator_) PreprocessPostClause(
	postClause ast.PostClauseLike,
	index_ uint,
	count_ uint,
) {
	// TBD - Add any validation checks.
}

func (v *validator_) PreprocessPrecedence(
	precedence ast.PrecedenceLike,
	index_ uint,
	count_ uint,
) {
	// TBD - Add any validation checks.
}

func (v *validator_) PreprocessPredicate(
	predicate ast.PredicateLike,
	index_ uint,
	count_ uint,
) {
	// TBD - Add any validation checks.
}

func (v *validator_) PreprocessPrimitive(
	primitive ast.PrimitiveLike,
	index_ uint,
	count_ uint,
) {
	// TBD - Add any validation checks.
}

func (v *validator_) PreprocessProcedure(
	procedure ast.ProcedureLike,
	index_ uint,
	count_ uint,
) {
	// TBD - Add any validation checks.
}

func (v *validator_) PreprocessPublishClause(
	publishClause ast.PublishClauseLike,
	index_ uint,
	count_ uint,
) {
	// TBD - Add any validation checks.
}

func (v *validator_) PreprocessRecipient(
	recipient ast.RecipientLike,
	index_ uint,
	count_ uint,
) {
	// TBD - Add any validation checks.
}

func (v *validator_) PreprocessReferent(
	referent ast.ReferentLike,
	index_ uint,
	count_ uint,
) {
	// TBD - Add any validation checks.
}

func (v *validator_) PreprocessRejectClause(
	rejectClause ast.RejectClauseLike,
	index_ uint,
	count_ uint,
) {
	// TBD - Add any validation checks.
}

func (v *validator_) PreprocessRepository(
	repository ast.RepositoryLike,
	index_ uint,
	count_ uint,
) {
	// TBD - Add any validation checks.
}

func (v *validator_) PreprocessResult(
	result ast.ResultLike,
	index_ uint,
	count_ uint,
) {
	// TBD - Add any validation checks.
}

func (v *validator_) PreprocessRetrieveClause(
	retrieveClause ast.RetrieveClauseLike,
	index_ uint,
	count_ uint,
) {
	// TBD - Add any validation checks.
}

func (v *validator_) PreprocessReturnClause(
	returnClause ast.ReturnClauseLike,
	index_ uint,
	count_ uint,
) {
	// TBD - Add any validation checks.
}

func (v *validator_) PreprocessSaveClause(
	saveClause ast.SaveClauseLike,
	index_ uint,
	count_ uint,
) {
	// TBD - Add any validation checks.
}

func (v *validator_) PreprocessSelectClause(
	selectClause ast.SelectClauseLike,
	index_ uint,
	count_ uint,
) {
	// TBD - Add any validation checks.
}

func (v *validator_) PreprocessSequence(
	sequence ast.SequenceLike,
	index_ uint,
	count_ uint,
) {
	// TBD - Add any validation checks.
}

func (v *validator_) PreprocessStatement(
	statement ast.StatementLike,
	index_ uint,
	count_ uint,
) {
	// TBD - Add any validation checks.
}

func (v *validator_) PreprocessStatementLine(
	statementLine ast.StatementLineLike,
	index_ uint,
	count_ uint,
) {
	// TBD - Add any validation checks.
}

func (v *validator_) PreprocessString(
	string_ ast.StringLike,
	index_ uint,
	count_ uint,
) {
	// TBD - Add any validation checks.
}

func (v *validator_) PreprocessSubcomponent(
	subcomponent ast.SubcomponentLike,
	index_ uint,
	count_ uint,
) {
	// TBD - Add any validation checks.
}

func (v *validator_) PreprocessSubject(
	subject ast.SubjectLike,
	index_ uint,
	count_ uint,
) {
	// TBD - Add any validation checks.
}

func (v *validator_) PreprocessTarget(
	target ast.TargetLike,
	index_ uint,
	count_ uint,
) {
	// TBD - Add any validation checks.
}

func (v *validator_) PreprocessTemplate(
	template ast.TemplateLike,
	index_ uint,
	count_ uint,
) {
	// TBD - Add any validation checks.
}

func (v *validator_) PreprocessTextual(
	textual ast.TextualLike,
	index_ uint,
	count_ uint,
) {
	// TBD - Add any validation checks.
}

func (v *validator_) PreprocessThrowClause(
	throwClause ast.ThrowClauseLike,
	index_ uint,
	count_ uint,
) {
	// TBD - Add any validation checks.
}

func (v *validator_) PreprocessVariable(
	variable ast.VariableLike,
	index_ uint,
	count_ uint,
) {
	// TBD - Add any validation checks.
}

func (v *validator_) PreprocessWhileClause(
	whileClause ast.WhileClauseLike,
	index_ uint,
	count_ uint,
) {
	// TBD - Add any validation checks.
}

func (v *validator_) PreprocessWithClause(
	withClause ast.WithClauseLike,
	index_ uint,
	count_ uint,
) {
	// TBD - Add any validation checks.
}

// PROTECTED INTERFACE

// Private Methods

func (v *validator_) validateToken(
	tokenValue string,
	tokenType TokenType,
) {
	var scannerClass = ScannerClass()
	if !scannerClass.MatchesType(tokenValue, tokenType) {
		var message = fmt.Sprintf(
			"The following token value is not of type %v: %v",
			scannerClass.FormatType(tokenType),
			tokenValue,
		)
		panic(message)
	}
}

// Instance Structure

type validator_ struct {
	// Declare the instance attributes.
	visitor_ VisitorLike

	// Declare the inherited aspects.
	Methodical
}

// Class Structure

type validatorClass_ struct {
	// Declare the class constants.
}

// Class Reference

func validatorClass() *validatorClass_ {
	return validatorClassReference_
}

var validatorClassReference_ = &validatorClass_{
	// Initialize the class constants.
}
