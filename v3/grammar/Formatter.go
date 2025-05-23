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
	ast "github.com/bali-nebula/go-bali-documents/v3/ast"
	sts "strings"
)

// CLASS INTERFACE

// Access Function

func FormatterClass() FormatterClassLike {
	return formatterClass()
}

// Constructor Methods

func (c *formatterClass_) Formatter() FormatterLike {
	var instance = &formatter_{
		// Initialize the instance attributes.

		// Initialize the inherited aspects.
		Methodical: ProcessorClass().Processor(),
	}
	instance.visitor_ = VisitorClass().Visitor(instance)
	return instance
}

// INSTANCE INTERFACE

// Principal Methods

func (v *formatter_) GetClass() FormatterClassLike {
	return formatterClass()
}

func (v *formatter_) FormatDocument(document ast.DocumentLike) string {
	v.visitor_.VisitDocument(document)
	return v.getResult()
}

// Methodical Methods

func (v *formatter_) ProcessAngle(
	angle string,
) {
	v.appendString(angle)
}

func (v *formatter_) ProcessArrow(
	arrow string,
) {
	v.appendString(arrow)
}

func (v *formatter_) ProcessBinary(
	binary string,
) {
	v.appendString(binary)
}

func (v *formatter_) ProcessBoolean(
	boolean string,
) {
	v.appendString(boolean)
}

func (v *formatter_) ProcessBytecode(
	bytecode string,
) {
	v.appendString(bytecode)
}

func (v *formatter_) ProcessCaret(
	caret string,
) {
	v.appendString(caret)
}

func (v *formatter_) ProcessCitation(
	citation string,
) {
	v.appendString(citation)
}

func (v *formatter_) ProcessComment(
	comment string,
) {
	v.appendString(comment)
}

func (v *formatter_) ProcessDelimiter(
	delimiter string,
) {
	v.appendString(delimiter)
}

func (v *formatter_) ProcessDot(
	dot string,
) {
	v.appendString(dot)
}

func (v *formatter_) ProcessDuration(
	duration string,
) {
	v.appendString(duration)
}

func (v *formatter_) ProcessEqual(
	equal string,
) {
	v.appendString(equal)
}

func (v *formatter_) ProcessIdentifier(
	identifier string,
) {
	v.appendString(identifier)
}

func (v *formatter_) ProcessLess(
	less string,
) {
	v.appendString(less)
}

func (v *formatter_) ProcessMinus(
	minus string,
) {
	v.appendString(minus)
}

func (v *formatter_) ProcessMoment(
	moment string,
) {
	v.appendString(moment)
}

func (v *formatter_) ProcessMore(
	more string,
) {
	v.appendString(more)
}

func (v *formatter_) ProcessName(
	name string,
) {
	v.appendString(name)
}

func (v *formatter_) ProcessNarrative(
	narrative string,
) {
	v.appendString(narrative)
}

func (v *formatter_) ProcessNewline(
	newline string,
) {
	v.appendNewline()
}

func (v *formatter_) ProcessNote(
	note string,
) {
	v.appendString("  ")
	v.appendString(note)
}

func (v *formatter_) ProcessNumber(
	number string,
) {
	v.appendString(number)
}

func (v *formatter_) ProcessPattern(
	pattern string,
) {
	v.appendString(pattern)
}

func (v *formatter_) ProcessPercent(
	percent string,
) {
	v.appendString(percent)
}

func (v *formatter_) ProcessPercentage(
	percentage string,
) {
	v.appendString(percentage)
}

func (v *formatter_) ProcessPlus(
	plus string,
) {
	v.appendString(plus)
}

func (v *formatter_) ProcessProbability(
	probability string,
) {
	v.appendString(probability)
}

func (v *formatter_) ProcessQuote(
	quote string,
) {
	v.appendString(quote)
}

func (v *formatter_) ProcessResource(
	resource string,
) {
	v.appendString(resource)
}

func (v *formatter_) ProcessSlash(
	slash string,
) {
	v.appendString(slash)
}

func (v *formatter_) ProcessSpace(
	space string,
) {
	v.appendString(space)
}

func (v *formatter_) ProcessStar(
	star string,
) {
	v.appendString(star)
}

func (v *formatter_) ProcessSymbol(
	symbol string,
) {
	v.appendString(symbol)
}

func (v *formatter_) ProcessTag(
	tag string,
) {
	v.appendString(tag)
}

func (v *formatter_) ProcessVersion(
	version string,
) {
	v.appendString(version)
}

func (v *formatter_) PreprocessAcceptClause(
	acceptClause ast.AcceptClauseLike,
	index_ uint,
	count_ uint,
) {
	// TBD - Add formatting of the rule.
}

func (v *formatter_) PostprocessAnnotation(
	annotation ast.AnnotationLike,
	index_ uint,
	count_ uint,
) {
	v.appendNewline()
}

func (v *formatter_) PreprocessArgument(
	argument ast.ArgumentLike,
	index_ uint,
	count_ uint,
) {
	// TBD - Add formatting of the rule.
}

func (v *formatter_) PreprocessArithmetic(
	arithmetic ast.ArithmeticLike,
	index_ uint,
	count_ uint,
) {
	// TBD - Add formatting of the rule.
}

func (v *formatter_) ProcessArithmeticSlot(
	arithmetic ast.ArithmeticLike,
	slot_ uint,
) {
	switch slot_ {
	default:
		v.appendString(" ")
	}
}

func (v *formatter_) PreprocessAssignment(
	assignment ast.AssignmentLike,
	index_ uint,
	count_ uint,
) {
	// TBD - Add formatting of the rule.
}

func (v *formatter_) ProcessAssignmentSlot(
	assignment ast.AssignmentLike,
	slot_ uint,
) {
	switch slot_ {
	default:
		v.appendString(" ")
	}
}

func (v *formatter_) ProcessAssociationSlot(
	association ast.AssociationLike,
	slot_ uint,
) {
	if slot_ == 2 {
		v.appendString(" ")
	}
}

func (v *formatter_) PreprocessAtLevel(
	atLevel ast.AtLevelLike,
	index_ uint,
	count_ uint,
) {
	// TBD - Add formatting of the rule.
}

func (v *formatter_) ProcessAttributesSlot(
	attributes ast.AttributesLike,
	slot_ uint,
) {
	switch slot_ {
	case 1:
		v.depth_++
	case 2:
		v.depth_--
		var associations = attributes.GetAssociations()
		if associations.GetSize() > 1 {
			v.appendNewline()
		}
	}
}

func (v *formatter_) PreprocessBag(
	bag ast.BagLike,
	index_ uint,
	count_ uint,
) {
	// TBD - Add formatting of the rule.
}

func (v *formatter_) PreprocessBlocking(
	blocking ast.BlockingLike,
	index_ uint,
	count_ uint,
) {
	// TBD - Add formatting of the rule.
}

func (v *formatter_) PreprocessBreakClause(
	breakClause ast.BreakClauseLike,
	index_ uint,
	count_ uint,
) {
	// TBD - Add formatting of the rule.
}

func (v *formatter_) PreprocessCheckoutClause(
	checkoutClause ast.CheckoutClauseLike,
	index_ uint,
	count_ uint,
) {
	// TBD - Add formatting of the rule.
}

func (v *formatter_) PreprocessCited(
	cited ast.CitedLike,
	index_ uint,
	count_ uint,
) {
	// TBD - Add formatting of the rule.
}

func (v *formatter_) PreprocessLine(
	line ast.LineLike,
	index_ uint,
	count_ uint,
) {
	v.appendNewline()
}

func (v *formatter_) PreprocessComparison(
	comparison ast.ComparisonLike,
	index_ uint,
	count_ uint,
) {
	// TBD - Add formatting of the rule.
}

func (v *formatter_) PreprocessComplement(
	complement ast.ComplementLike,
	index_ uint,
	count_ uint,
) {
	// TBD - Add formatting of the rule.
}

func (v *formatter_) PreprocessCondition(
	condition ast.ConditionLike,
	index_ uint,
	count_ uint,
) {
	// TBD - Add formatting of the rule.
}

func (v *formatter_) PreprocessContinueClause(
	continueClause ast.ContinueClauseLike,
	index_ uint,
	count_ uint,
) {
	// TBD - Add formatting of the rule.
}

func (v *formatter_) PreprocessDiscardClause(
	discardClause ast.DiscardClauseLike,
	index_ uint,
	count_ uint,
) {
	// TBD - Add formatting of the rule.
}

func (v *formatter_) PreprocessDoClause(
	doClause ast.DoClauseLike,
	index_ uint,
	count_ uint,
) {
	// TBD - Add formatting of the rule.
}

func (v *formatter_) PostprocessDocument(
	document ast.DocumentLike,
	index_ uint,
	count_ uint,
) {
	v.appendNewline()
}

func (v *formatter_) PreprocessDraft(
	draft ast.DraftLike,
	index_ uint,
	count_ uint,
) {
	// TBD - Add formatting of the rule.
}

func (v *formatter_) PreprocessEvent(
	event ast.EventLike,
	index_ uint,
	count_ uint,
) {
	// TBD - Add formatting of the rule.
}

func (v *formatter_) PreprocessException(
	exception ast.ExceptionLike,
	index_ uint,
	count_ uint,
) {
	// TBD - Add formatting of the rule.
}

func (v *formatter_) PreprocessExpression(
	expression ast.ExpressionLike,
	index_ uint,
	count_ uint,
) {
	// TBD - Add formatting of the rule.
}

func (v *formatter_) PreprocessFailure(
	failure ast.FailureLike,
	index_ uint,
	count_ uint,
) {
	// TBD - Add formatting of the rule.
}

func (v *formatter_) PreprocessFlow(
	flow ast.FlowLike,
	index_ uint,
	count_ uint,
) {
	// TBD - Add formatting of the rule.
}

func (v *formatter_) PreprocessFunction(
	function ast.FunctionLike,
	index_ uint,
	count_ uint,
) {
	// TBD - Add formatting of the rule.
}

func (v *formatter_) PreprocessHandler(
	handler ast.HandlerLike,
	index_ uint,
	count_ uint,
) {
	// TBD - Add formatting of the rule.
}

func (v *formatter_) PreprocessIfClause(
	ifClause ast.IfClauseLike,
	index_ uint,
	count_ uint,
) {
	// TBD - Add formatting of the rule.
}

func (v *formatter_) PreprocessIndex(
	index ast.IndexLike,
	index_ uint,
	count_ uint,
) {
	// TBD - Add formatting of the rule.
}

func (v *formatter_) PreprocessIndirect(
	indirect ast.IndirectLike,
	index_ uint,
	count_ uint,
) {
	// TBD - Add formatting of the rule.
}

func (v *formatter_) PreprocessInduction(
	induction ast.InductionLike,
	index_ uint,
	count_ uint,
) {
	// TBD - Add formatting of the rule.
}

func (v *formatter_) PreprocessInverse(
	inverse ast.InverseLike,
	index_ uint,
	count_ uint,
) {
	// TBD - Add formatting of the rule.
}

func (v *formatter_) PreprocessInversion(
	inversion ast.InversionLike,
	index_ uint,
	count_ uint,
) {
	// TBD - Add formatting of the rule.
}

func (v *formatter_) PreprocessInvocation(
	invocation ast.InvocationLike,
	index_ uint,
	count_ uint,
) {
	// TBD - Add formatting of the rule.
}

func (v *formatter_) ProcessItemsSlot(
	items ast.ItemsLike,
	slot_ uint,
) {
	switch slot_ {
	case 1:
		v.depth_++
	case 2:
		v.depth_--
		var components = items.GetComponents()
		if components.GetSize() > 1 {
			v.appendNewline()
		}
	}
}

func (v *formatter_) PreprocessLetClause(
	letClause ast.LetClauseLike,
	index_ uint,
	count_ uint,
) {
	// TBD - Add formatting of the rule.
}

func (v *formatter_) PreprocessLogic(
	logic ast.LogicLike,
	index_ uint,
	count_ uint,
) {
	// TBD - Add formatting of the rule.
}

func (v *formatter_) PreprocessLogical(
	logical ast.LogicalLike,
	index_ uint,
	count_ uint,
) {
	// TBD - Add formatting of the rule.
}

func (v *formatter_) PreprocessMagnitude(
	magnitude ast.MagnitudeLike,
	index_ uint,
	count_ uint,
) {
	// TBD - Add formatting of the rule.
}

func (v *formatter_) PreprocessMainClause(
	mainClause ast.MainClauseLike,
	index_ uint,
	count_ uint,
) {
	// TBD - Add formatting of the rule.
}

func (v *formatter_) PreprocessMessage(
	message ast.MessageLike,
	index_ uint,
	count_ uint,
) {
	// TBD - Add formatting of the rule.
}

func (v *formatter_) PreprocessMessaging(
	messaging ast.MessagingLike,
	index_ uint,
	count_ uint,
) {
	// TBD - Add formatting of the rule.
}

func (v *formatter_) PreprocessMethod(
	method ast.MethodLike,
	index_ uint,
	count_ uint,
) {
	// TBD - Add formatting of the rule.
}

func (v *formatter_) PreprocessNotarizeClause(
	notarizeClause ast.NotarizeClauseLike,
	index_ uint,
	count_ uint,
) {
	// TBD - Add formatting of the rule.
}

func (v *formatter_) PreprocessNumerical(
	numerical ast.NumericalLike,
	index_ uint,
	count_ uint,
) {
	// TBD - Add formatting of the rule.
}

func (v *formatter_) PreprocessOnClause(
	onClause ast.OnClauseLike,
	index_ uint,
	count_ uint,
) {
	v.appendString(" ")
}

func (v *formatter_) ProcessOnClauseSlot(
	onClause ast.OnClauseLike,
	slot_ uint,
) {
	switch slot_ {
	case 1:
		v.appendString(" ")
	}
}

func (v *formatter_) PreprocessOperation(
	operation ast.OperationLike,
	index_ uint,
	count_ uint,
) {
	// TBD - Add formatting of the rule.
}

func (v *formatter_) ProcessParametersSlot(
	parameters ast.ParametersLike,
	slot_ uint,
) {
	switch slot_ {
	case 1:
		v.depth_++
	case 2:
		v.depth_--
		var associations = parameters.GetAssociations()
		if associations.GetSize() > 1 {
			v.appendNewline()
		}
	}
}

func (v *formatter_) PreprocessPostClause(
	postClause ast.PostClauseLike,
	index_ uint,
	count_ uint,
) {
	// TBD - Add formatting of the rule.
}

func (v *formatter_) PreprocessPrecedence(
	precedence ast.PrecedenceLike,
	index_ uint,
	count_ uint,
) {
	// TBD - Add formatting of the rule.
}

func (v *formatter_) PreprocessPredicate(
	predicate ast.PredicateLike,
	index_ uint,
	count_ uint,
) {
	// TBD - Add formatting of the rule.
}

func (v *formatter_) ProcessProcedureSlot(
	procedure ast.ProcedureLike,
	slot_ uint,
) {
	switch slot_ {
	case 1:
		v.depth_++
	case 2:
		v.depth_--
		var lines = procedure.GetLines()
		if lines.GetSize() > 1 {
			v.appendNewline()
		}
	}
}

func (v *formatter_) PreprocessPublishClause(
	publishClause ast.PublishClauseLike,
	index_ uint,
	count_ uint,
) {
	// TBD - Add formatting of the rule.
}

func (v *formatter_) PreprocessRecipient(
	recipient ast.RecipientLike,
	index_ uint,
	count_ uint,
) {
	// TBD - Add formatting of the rule.
}

func (v *formatter_) PreprocessReferent(
	referent ast.ReferentLike,
	index_ uint,
	count_ uint,
) {
	// TBD - Add formatting of the rule.
}

func (v *formatter_) PreprocessRejectClause(
	rejectClause ast.RejectClauseLike,
	index_ uint,
	count_ uint,
) {
	// TBD - Add formatting of the rule.
}

func (v *formatter_) PreprocessRepository(
	repository ast.RepositoryLike,
	index_ uint,
	count_ uint,
) {
	// TBD - Add formatting of the rule.
}

func (v *formatter_) PreprocessResult(
	result ast.ResultLike,
	index_ uint,
	count_ uint,
) {
	// TBD - Add formatting of the rule.
}

func (v *formatter_) PreprocessRetrieveClause(
	retrieveClause ast.RetrieveClauseLike,
	index_ uint,
	count_ uint,
) {
	// TBD - Add formatting of the rule.
}

func (v *formatter_) PreprocessReturnClause(
	returnClause ast.ReturnClauseLike,
	index_ uint,
	count_ uint,
) {
	// TBD - Add formatting of the rule.
}

func (v *formatter_) PreprocessSaveClause(
	saveClause ast.SaveClauseLike,
	index_ uint,
	count_ uint,
) {
	// TBD - Add formatting of the rule.
}

func (v *formatter_) PreprocessSelectClause(
	selectClause ast.SelectClauseLike,
	index_ uint,
	count_ uint,
) {
	// TBD - Add formatting of the rule.
}

func (v *formatter_) PreprocessSequence(
	sequence ast.SequenceLike,
	index_ uint,
	count_ uint,
) {
	// TBD - Add formatting of the rule.
}

func (v *formatter_) PreprocessStatement(
	statement ast.StatementLike,
	index_ uint,
	count_ uint,
) {
	// TBD - Add formatting of the rule.
}

func (v *formatter_) PreprocessSubcomponent(
	subcomponent ast.SubcomponentLike,
	index_ uint,
	count_ uint,
) {
	// TBD - Add formatting of the rule.
}

func (v *formatter_) PreprocessSubject(
	subject ast.SubjectLike,
	index_ uint,
	count_ uint,
) {
	// TBD - Add formatting of the rule.
}

func (v *formatter_) PreprocessTarget(
	target ast.TargetLike,
	index_ uint,
	count_ uint,
) {
	// TBD - Add formatting of the rule.
}

func (v *formatter_) PreprocessTemplate(
	template ast.TemplateLike,
	index_ uint,
	count_ uint,
) {
	// TBD - Add formatting of the rule.
}

func (v *formatter_) PreprocessTextual(
	textual ast.TextualLike,
	index_ uint,
	count_ uint,
) {
	// TBD - Add formatting of the rule.
}

func (v *formatter_) PreprocessThrowClause(
	throwClause ast.ThrowClauseLike,
	index_ uint,
	count_ uint,
) {
	// TBD - Add formatting of the rule.
}

func (v *formatter_) PreprocessVariable(
	variable ast.VariableLike,
	index_ uint,
	count_ uint,
) {
	// TBD - Add formatting of the rule.
}

func (v *formatter_) PreprocessWhileClause(
	whileClause ast.WhileClauseLike,
	index_ uint,
	count_ uint,
) {
	// TBD - Add formatting of the rule.
}

func (v *formatter_) PreprocessWithClause(
	withClause ast.WithClauseLike,
	index_ uint,
	count_ uint,
) {
	// TBD - Add formatting of the rule.
}

const _indentation = "    "

// PROTECTED INTERFACE

// Private Methods

func (v *formatter_) appendNewline() {
	var newline = "\n"
	var level uint
	for ; level < v.depth_; level++ {
		newline += _indentation
	}
	v.appendString(newline)
}

func (v *formatter_) appendString(
	string_ string,
) {
	v.result_.WriteString(string_)
}

func (v *formatter_) getResult() string {
	var result = v.result_.String()
	v.result_.Reset()
	return result
}

// Instance Structure

type formatter_ struct {
	// Declare the instance attributes.
	visitor_ VisitorLike
	depth_   uint
	result_  sts.Builder

	// Declare the inherited aspects.
	Methodical
}

// Class Structure

type formatterClass_ struct {
	// Declare the class constants.
}

// Class Reference

func formatterClass() *formatterClass_ {
	return formatterClassReference_
}

var formatterClassReference_ = &formatterClass_{
	// Initialize the class constants.
}
