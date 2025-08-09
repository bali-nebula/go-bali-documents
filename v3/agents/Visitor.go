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

package agents

import (
	doc "github.com/bali-nebula/go-bali-documents/v3/documents"
	uti "github.com/craterdog/go-missing-utilities/v7"
)

// CLASS INTERFACE

// Access Function

func VisitorClass() VisitorClassLike {
	return visitorClass()
}

// Constructor Methods

func (c *visitorClass_) Visitor(
	processor Methodical,
) VisitorLike {
	if uti.IsUndefined(processor) {
		panic("The \"processor\" attribute is required by this class.")
	}
	var instance = &visitor_{
		// Initialize the instance attributes.
		processor_: processor,
	}
	return instance
}

// Constant Methods

// Function Methods

// INSTANCE INTERFACE

// Principal Methods

func (v *visitor_) GetClass() VisitorClassLike {
	return visitorClass()
}

func (v *visitor_) VisitDocument(
	document doc.DocumentLike,
) {
	// TBD - Add the method implementation.
}

// Attribute Methods

// PROTECTED INTERFACE

// Private Methods

func (v *visitor_) visitAcceptClause(
	acceptClause doc.AcceptClauseLike,
) {
	var message = acceptClause.GetMessage()
	v.processor_.PreprocessExpression(
		message,
		1,
		1,
	)
	v.visitExpression(message)
	v.processor_.PostprocessExpression(
		message,
		1,
		1,
	)
}

func (v *visitor_) visitArgument(
	argument any,
) {
	panic("Not yet implemented.")
}

func (v *visitor_) visitAttributes(
	attributes doc.AttributesLike,
) {
	var associationsIndex uint
	var associations = attributes.GetAssociations().GetIterator()
	var associationsCount = uint(associations.GetSize())
	for associations.HasNext() {
		associationsIndex++
		var association = associations.GetNext()
		var primitive = association.GetKey()
		v.visitPrimitive(primitive)
		var document = association.GetValue()
		v.processor_.PreprocessDocument(
			document,
			associationsIndex,
			associationsCount,
		)
		v.visitDocument(document)
		v.processor_.PostprocessDocument(
			document,
			associationsIndex,
			associationsCount,
		)
	}
}

func (v *visitor_) visitBreakClause(
	breakClause doc.BreakClauseLike,
) {
}

func (v *visitor_) visitCheckoutClause(
	checkoutClause doc.CheckoutClauseLike,
) {
	var recipient = checkoutClause.GetRecipient()
	v.visitRecipient(recipient)

	// Visit slot 1 between terms.
	v.processor_.ProcessCheckoutClauseSlot(
		checkoutClause,
		1,
	)

	var optionalAtLevel = checkoutClause.GetOptionalAtLevel()
	if uti.IsDefined(optionalAtLevel) {
		v.processor_.PreprocessExpression(
			optionalAtLevel,
			1,
			1,
		)
		v.visitExpression(optionalAtLevel)
		v.processor_.PostprocessExpression(
			optionalAtLevel,
			1,
			1,
		)
	}

	// Visit slot 2 between terms.
	v.processor_.ProcessCheckoutClauseSlot(
		checkoutClause,
		2,
	)

	var cited = checkoutClause.GetCited()
	v.processor_.PreprocessExpression(
		cited,
		1,
		1,
	)
	v.visitExpression(cited)
	v.processor_.PostprocessExpression(
		cited,
		1,
		1,
	)
}

func (v *visitor_) visitComplement(
	complement doc.ComplementLike,
) {
	var logical = complement.GetLogical()
	v.visitLogical(logical)
}

func (v *visitor_) visitComponent(
	component any,
) {
	panic("Not yet implemented.")
}

func (v *visitor_) visitContinueClause(
	continueClause doc.ContinueClauseLike,
) {
}

func (v *visitor_) visitDiscardClause(
	discardClause doc.DiscardClauseLike,
) {
	var draft = discardClause.GetDraft()
	v.processor_.PreprocessExpression(
		draft,
		1,
		1,
	)
	v.visitExpression(draft)
	v.processor_.PostprocessExpression(
		draft,
		1,
		1,
	)
}

func (v *visitor_) visitDoClause(
	doClause doc.DoClauseLike,
) {
	var invocation = doClause.GetInvocation()
	v.visitInvocation(invocation)
}

func (v *visitor_) visitDocument(
	document doc.DocumentLike,
) {
	var component = document.GetComponent()
	v.visitComponent(component)

	// Visit slot 1 between terms.
	v.processor_.ProcessDocumentSlot(
		document,
		1,
	)

	var optionalParameters = document.GetOptionalParameters()
	if uti.IsDefined(optionalParameters) {
		v.processor_.PreprocessParameters(
			optionalParameters,
			1,
			1,
		)
		v.visitParameters(optionalParameters)
		v.processor_.PostprocessParameters(
			optionalParameters,
			1,
			1,
		)
	}
	// Visit slot 2 between terms.
	v.processor_.ProcessDocumentSlot(
		document,
		2,
	)

	var optionalNote = document.GetOptionalNote()
	if uti.IsDefined(optionalNote) {
		v.processor_.ProcessNote(optionalNote)
	}
}

func (v *visitor_) visitExpression(
	expression doc.ExpressionLike,
) {
	var subject = expression.GetSubject()
	v.visitSubject(subject)

	// Visit slot 1 between terms.
	v.processor_.ProcessExpressionSlot(
		expression,
		1,
	)

	var predicatesIndex uint
	var predicates = expression.GetPredicates().GetIterator()
	var predicatesCount = uint(predicates.GetSize())
	for predicates.HasNext() {
		predicatesIndex++
		var predicate = predicates.GetNext()
		v.processor_.PreprocessPredicate(
			predicate,
			predicatesIndex,
			predicatesCount,
		)
		v.visitPredicate(predicate)
		v.processor_.PostprocessPredicate(
			predicate,
			predicatesIndex,
			predicatesCount,
		)
	}
}

func (v *visitor_) visitExtent(
	extent doc.Extent,
) {
	v.processor_.ProcessExtent(extent)
}

func (v *visitor_) visitFunction(
	function doc.FunctionLike,
) {
	var identifier = function.GetIdentifier()
	v.processor_.ProcessIdentifier(identifier)
	// Visit slot 1 between terms.
	v.processor_.ProcessFunctionSlot(
		function,
		1,
	)

	var arguments = function.GetArguments().GetIterator()
	for arguments.HasNext() {
		var argument = arguments.GetNext()
		v.visitArgument(argument)
	}
}

func (v *visitor_) visitIfClause(
	ifClause doc.IfClauseLike,
) {
	var condition = ifClause.GetCondition()
	v.processor_.PreprocessExpression(
		condition,
		1,
		1,
	)
	v.visitExpression(condition)
	v.processor_.PostprocessExpression(
		condition,
		1,
		1,
	)
	// Visit slot 1 between terms.
	v.processor_.ProcessIfClauseSlot(
		ifClause,
		1,
	)

	var procedure = ifClause.GetProcedure()
	v.processor_.PreprocessProcedure(
		procedure,
		1,
		1,
	)
	v.visitProcedure(procedure)
	v.processor_.PostprocessProcedure(
		procedure,
		1,
		1,
	)
}

func (v *visitor_) visitIndex(
	index any,
) {
	panic("Not yet implemented.")
}

func (v *visitor_) visitIndirect(
	indirect any,
) {
	panic("Not yet implemented.")
}

func (v *visitor_) visitInversion(
	inversion doc.InversionLike,
) {
	var inverse = inversion.GetInverse()
	v.processor_.ProcessInverse(inverse)

	// Visit slot 1 between terms.
	v.processor_.ProcessInversionSlot(
		inversion,
		1,
	)

	var numerical = inversion.GetNumerical()
	v.visitNumerical(numerical)
}

func (v *visitor_) visitInvocation(
	invocation any,
) {
	panic("Not yet implemented.")
}

func (v *visitor_) visitItems(
	items doc.ItemsLike,
) {
	var entitiesIndex uint
	var entities = items.GetEntities().GetIterator()
	var entitiesCount = uint(entities.GetSize())
	for entities.HasNext() {
		entitiesIndex++
		var document = entities.GetNext()
		v.processor_.PreprocessDocument(
			document,
			entitiesIndex,
			entitiesCount,
		)
		v.visitDocument(document)
		v.processor_.PostprocessDocument(
			document,
			entitiesIndex,
			entitiesCount,
		)
	}
}

func (v *visitor_) visitLetClause(
	letClause doc.LetClauseLike,
) {
	var recipient = letClause.GetRecipient()
	v.visitRecipient(recipient)

	// Visit slot 1 between terms.
	v.processor_.ProcessLetClauseSlot(
		letClause,
		1,
	)

	var assignment = letClause.GetAssignment()
	v.processor_.ProcessAssignment(assignment)

	// Visit slot 2 between terms.
	v.processor_.ProcessLetClauseSlot(
		letClause,
		2,
	)

	var expression = letClause.GetExpression()
	v.processor_.PreprocessExpression(
		expression,
		1,
		1,
	)
	v.visitExpression(expression)
	v.processor_.PostprocessExpression(
		expression,
		1,
		1,
	)
}

func (v *visitor_) visitLine(
	line any,
) {
	panic("Not yet implemented.")
}

func (v *visitor_) visitLogical(
	logical any,
) {
	panic("Not yet implemented.")
}

func (v *visitor_) visitMagnitude(
	magnitude doc.MagnitudeLike,
) {
	var expression = magnitude.GetExpression()
	v.processor_.PreprocessExpression(
		expression,
		1,
		1,
	)
	v.visitExpression(expression)
	v.processor_.PostprocessExpression(
		expression,
		1,
		1,
	)
}

func (v *visitor_) visitMainClause(
	mainClause any,
) {
	panic("Not yet implemented.")
}

func (v *visitor_) visitMatchingClause(
	matchingClause doc.MatchingClauseLike,
) {
	var template = matchingClause.GetTemplate()
	v.processor_.PreprocessExpression(
		template,
		1,
		1,
	)
	v.visitExpression(template)
	v.processor_.PostprocessExpression(
		template,
		1,
		1,
	)
	// Visit slot 1 between terms.
	v.processor_.ProcessMatchingClauseSlot(
		matchingClause,
		1,
	)

	var procedure = matchingClause.GetProcedure()
	v.processor_.PreprocessProcedure(
		procedure,
		1,
		1,
	)
	v.visitProcedure(procedure)
	v.processor_.PostprocessProcedure(
		procedure,
		1,
		1,
	)
}

func (v *visitor_) visitMethod(
	method doc.MethodLike,
) {
	var target = method.GetTarget()
	v.processor_.ProcessIdentifier(target)

	// Visit slot 1 between terms.
	v.processor_.ProcessMethodSlot(
		method,
		1,
	)

	var invoke = method.GetInvoke()
	v.processor_.ProcessInvoke(invoke)

	// Visit slot 2 between terms.
	v.processor_.ProcessMethodSlot(
		method,
		2,
	)

	var message = method.GetMessage()
	v.processor_.ProcessIdentifier(message)

	// Visit slot 3 between terms.
	v.processor_.ProcessMethodSlot(
		method,
		3,
	)

	var arguments = method.GetArguments().GetIterator()
	for arguments.HasNext() {
		var argument = arguments.GetNext()
		v.visitArgument(argument)
	}
}

func (v *visitor_) visitNotarizeClause(
	notarizeClause doc.NotarizeClauseLike,
) {
	var draft = notarizeClause.GetDraft()
	v.processor_.PreprocessExpression(
		draft,
		1,
		1,
	)
	v.visitExpression(draft)
	v.processor_.PostprocessExpression(
		draft,
		1,
		1,
	)

	// Visit slot 1 between terms.
	v.processor_.ProcessNotarizeClauseSlot(
		notarizeClause,
		1,
	)

	var cited = notarizeClause.GetCited()
	v.processor_.PreprocessExpression(
		cited,
		1,
		1,
	)
	v.visitExpression(cited)
	v.processor_.PostprocessExpression(
		cited,
		1,
		1,
	)
}

func (v *visitor_) visitNumerical(
	numerical any,
) {
	panic("Not yet implemented.")
}

func (v *visitor_) visitOnClause(
	onClause doc.OnClauseLike,
) {
	var failure = onClause.GetFailure()
	v.processor_.ProcessSymbol(failure)

	// Visit slot 1 between terms.
	v.processor_.ProcessOnClauseSlot(
		onClause,
		1,
	)

	var matchingClausesIndex uint
	var matchingClauses = onClause.GetMatchingClauses().GetIterator()
	var matchingClausesCount = uint(matchingClauses.GetSize())
	for matchingClauses.HasNext() {
		matchingClausesIndex++
		var clause = matchingClauses.GetNext()
		v.processor_.PreprocessMatchingClause(
			clause,
			matchingClausesIndex,
			matchingClausesCount,
		)
		v.visitMatchingClause(clause)
		v.processor_.PostprocessMatchingClause(
			clause,
			matchingClausesIndex,
			matchingClausesCount,
		)
	}
}

func (v *visitor_) visitParameters(
	parameters doc.ParametersLike,
) {
	var associationsIndex uint
	var associations = parameters.GetAssociations().GetIterator()
	var associationsCount = uint(associations.GetSize())
	for associations.HasNext() {
		associationsIndex++
		var association = associations.GetNext()
		var symbol = association.GetKey()
		v.processor_.ProcessSymbol(symbol)
		var document = association.GetValue()
		v.processor_.PreprocessDocument(
			document,
			associationsIndex,
			associationsCount,
		)
		v.visitDocument(document)
		v.processor_.PostprocessDocument(
			document,
			associationsIndex,
			associationsCount,
		)
	}
}

func (v *visitor_) visitPostClause(
	postClause doc.PostClauseLike,
) {
	var message = postClause.GetMessage()
	v.processor_.PreprocessExpression(
		message,
		1,
		1,
	)
	v.visitExpression(message)
	v.processor_.PostprocessExpression(
		message,
		1,
		1,
	)

	// Visit slot 1 between terms.
	v.processor_.ProcessPostClauseSlot(
		postClause,
		1,
	)

	var bag = postClause.GetBag()
	v.processor_.PreprocessExpression(
		bag,
		1,
		1,
	)
	v.visitExpression(bag)
	v.processor_.PostprocessExpression(
		bag,
		1,
		1,
	)
}

func (v *visitor_) visitPrecedence(
	precedence doc.PrecedenceLike,
) {
	var expression = precedence.GetExpression()
	v.processor_.PreprocessExpression(
		expression,
		1,
		1,
	)
	v.visitExpression(expression)
	v.processor_.PostprocessExpression(
		expression,
		1,
		1,
	)
}

func (v *visitor_) visitPredicate(
	predicate doc.PredicateLike,
) {
	var operator = predicate.GetOperator()
	v.processor_.ProcessOperator(operator)

	// Visit slot 1 between terms.
	v.processor_.ProcessPredicateSlot(
		predicate,
		1,
	)

	var expression = predicate.GetExpression()
	v.processor_.PreprocessExpression(
		expression,
		1,
		1,
	)
	v.visitExpression(expression)
	v.processor_.PostprocessExpression(
		expression,
		1,
		1,
	)
}

func (v *visitor_) visitPrimitive(
	primitive any,
) {
	panic("Not yet implemented.")
}

func (v *visitor_) visitProcedure(
	procedure doc.ProcedureLike,
) {
	var lines = procedure.GetLines().GetIterator()
	for lines.HasNext() {
		var line = lines.GetNext()
		v.visitLine(line)
	}
}

func (v *visitor_) visitPublishClause(
	publishClause doc.PublishClauseLike,
) {
	var event = publishClause.GetEvent()
	v.processor_.PreprocessExpression(
		event,
		1,
		1,
	)
	v.visitExpression(event)
	v.processor_.PostprocessExpression(
		event,
		1,
		1,
	)
}

func (v *visitor_) visitRange(
	range_ doc.RangeLike,
) {
	var left = range_.GetLeft()
	v.visitExtent(left)

	// Visit slot 1 between terms.
	v.processor_.ProcessRangeSlot(
		range_,
		1,
	)

	var first = range_.GetFirst()
	v.visitPrimitive(first)

	// Visit slot 2 between terms.
	v.processor_.ProcessRangeSlot(
		range_,
		2,
	)

	var last = range_.GetLast()
	v.visitPrimitive(last)

	// Visit slot 3 between terms.
	v.processor_.ProcessRangeSlot(
		range_,
		3,
	)

	var right = range_.GetRight()
	v.visitExtent(right)
}

func (v *visitor_) visitRecipient(
	recipient any,
) {
	panic("Not yet implemented.")
}

func (v *visitor_) visitReferent(
	referent doc.ReferentLike,
) {
	var indirect = referent.GetIndirect()
	v.visitIndirect(indirect)
}

func (v *visitor_) visitRejectClause(
	rejectClause doc.RejectClauseLike,
) {
	var message = rejectClause.GetMessage()
	v.processor_.PreprocessExpression(
		message,
		1,
		1,
	)
	v.visitExpression(message)
	v.processor_.PostprocessExpression(
		message,
		1,
		1,
	)
}

func (v *visitor_) visitRetrieveClause(
	retrieveClause doc.RetrieveClauseLike,
) {
	var recipient = retrieveClause.GetRecipient()
	v.visitRecipient(recipient)

	// Visit slot 1 between terms.
	v.processor_.ProcessRetrieveClauseSlot(
		retrieveClause,
		1,
	)

	var bag = retrieveClause.GetBag()
	v.processor_.PreprocessExpression(
		bag,
		1,
		1,
	)
	v.visitExpression(bag)
	v.processor_.PostprocessExpression(
		bag,
		1,
		1,
	)
}

func (v *visitor_) visitReturnClause(
	returnClause doc.ReturnClauseLike,
) {
	var result = returnClause.GetResult()
	v.processor_.PreprocessExpression(
		result,
		1,
		1,
	)
	v.visitExpression(result)
	v.processor_.PostprocessExpression(
		result,
		1,
		1,
	)
}

func (v *visitor_) visitSaveClause(
	saveClause doc.SaveClauseLike,
) {
	var draft = saveClause.GetDraft()
	v.processor_.PreprocessExpression(
		draft,
		1,
		1,
	)
	v.visitExpression(draft)
	v.processor_.PostprocessExpression(
		draft,
		1,
		1,
	)
	// Visit slot 1 between terms.
	v.processor_.ProcessSaveClauseSlot(
		saveClause,
		1,
	)

	var cited = saveClause.GetCited()
	v.processor_.PreprocessExpression(
		cited,
		1,
		1,
	)
	v.visitExpression(cited)
	v.processor_.PostprocessExpression(
		cited,
		1,
		1,
	)
}

func (v *visitor_) visitSelectClause(
	selectClause doc.SelectClauseLike,
) {
	var target = selectClause.GetTarget()
	v.visitTarget(target)

	// Visit slot 1 between terms.
	v.processor_.ProcessSelectClauseSlot(
		selectClause,
		1,
	)

	var matchingClausesIndex uint
	var matchingClauses = selectClause.GetMatchingClauses().GetIterator()
	var matchingClausesCount = uint(matchingClauses.GetSize())
	for matchingClauses.HasNext() {
		matchingClausesIndex++
		var clause = matchingClauses.GetNext()
		v.processor_.PreprocessMatchingClause(
			clause,
			matchingClausesIndex,
			matchingClausesCount,
		)
		v.visitMatchingClause(clause)
		v.processor_.PostprocessMatchingClause(
			clause,
			matchingClausesIndex,
			matchingClausesCount,
		)
	}
}

func (v *visitor_) visitStatement(
	statement doc.StatementLike,
) {
	var mainClause = statement.GetMainClause()
	v.visitMainClause(mainClause)

	// Visit slot 1 between terms.
	v.processor_.ProcessStatementSlot(
		statement,
		1,
	)

	var optionalOnClause = statement.GetOptionalOnClause()
	if uti.IsDefined(optionalOnClause) {
		v.processor_.PreprocessOnClause(
			optionalOnClause,
			1,
			1,
		)
		v.visitOnClause(optionalOnClause)
		v.processor_.PostprocessOnClause(
			optionalOnClause,
			1,
			1,
		)
	}
}

func (v *visitor_) visitSubcomponent(
	subcomponent doc.SubcomponentLike,
) {
	var identifier = subcomponent.GetIdentifier()
	v.processor_.ProcessIdentifier(identifier)

	// Visit slot 1 between terms.
	v.processor_.ProcessSubcomponentSlot(
		subcomponent,
		1,
	)

	var indexes = subcomponent.GetIndexes().GetIterator()
	for indexes.HasNext() {
		var index = indexes.GetNext()
		v.visitIndex(index)
	}
}

func (v *visitor_) visitSubject(
	subject any,
) {
	panic("Not yet implemented.")
}

func (v *visitor_) visitTarget(
	target any,
) {
	panic("Not yet implemented.")
}

func (v *visitor_) visitThrowClause(
	throwClause doc.ThrowClauseLike,
) {
	var exception = throwClause.GetException()
	v.processor_.PreprocessExpression(
		exception,
		1,
		1,
	)
	v.visitExpression(exception)
	v.processor_.PostprocessExpression(
		exception,
		1,
		1,
	)
}

func (v *visitor_) visitWhileClause(
	whileClause doc.WhileClauseLike,
) {
	var delimiter1 = whileClause.GetDelimiter1()
	v.processor_.ProcessDelimiter(delimiter1)
	// Visit slot 1 between terms.
	v.processor_.ProcessWhileClauseSlot(
		whileClause,
		1,
	)

	var condition = whileClause.GetCondition()
	v.processor_.PreprocessExpression(
		condition,
		1,
		1,
	)
	v.visitExpression(condition)
	v.processor_.PostprocessExpression(
		condition,
		1,
		1,
	)
	// Visit slot 2 between terms.
	v.processor_.ProcessWhileClauseSlot(
		whileClause,
		2,
	)

	var delimiter2 = whileClause.GetDelimiter2()
	v.processor_.ProcessDelimiter(delimiter2)
	// Visit slot 3 between terms.
	v.processor_.ProcessWhileClauseSlot(
		whileClause,
		3,
	)

	var procedure = whileClause.GetProcedure()
	v.processor_.PreprocessProcedure(
		procedure,
		1,
		1,
	)
	v.visitProcedure(procedure)
	v.processor_.PostprocessProcedure(
		procedure,
		1,
		1,
	)
}

func (v *visitor_) visitWithClause(
	withClause doc.WithClauseLike,
) {
	var delimiter1 = withClause.GetDelimiter1()
	v.processor_.ProcessDelimiter(delimiter1)
	// Visit slot 1 between terms.
	v.processor_.ProcessWithClauseSlot(
		withClause,
		1,
	)

	var delimiter2 = withClause.GetDelimiter2()
	v.processor_.ProcessDelimiter(delimiter2)
	// Visit slot 2 between terms.
	v.processor_.ProcessWithClauseSlot(
		withClause,
		2,
	)

	var variable = withClause.GetVariable()
	v.processor_.PreprocessVariable(
		variable,
		1,
		1,
	)
	v.visitVariable(variable)
	v.processor_.PostprocessVariable(
		variable,
		1,
		1,
	)
	// Visit slot 3 between terms.
	v.processor_.ProcessWithClauseSlot(
		withClause,
		3,
	)

	var delimiter3 = withClause.GetDelimiter3()
	v.processor_.ProcessDelimiter(delimiter3)
	// Visit slot 4 between terms.
	v.processor_.ProcessWithClauseSlot(
		withClause,
		4,
	)

	var sequence = withClause.GetSequence()
	v.processor_.PreprocessSequence(
		sequence,
		1,
		1,
	)
	v.visitSequence(sequence)
	v.processor_.PostprocessSequence(
		sequence,
		1,
		1,
	)
	// Visit slot 5 between terms.
	v.processor_.ProcessWithClauseSlot(
		withClause,
		5,
	)

	var delimiter4 = withClause.GetDelimiter4()
	v.processor_.ProcessDelimiter(delimiter4)
	// Visit slot 6 between terms.
	v.processor_.ProcessWithClauseSlot(
		withClause,
		6,
	)

	var procedure = withClause.GetProcedure()
	v.processor_.PreprocessProcedure(
		procedure,
		1,
		1,
	)
	v.visitProcedure(procedure)
	v.processor_.PostprocessProcedure(
		procedure,
		1,
		1,
	)
}

// Instance Structure

type visitor_ struct {
	// Declare the instance attributes.
	processor_ Methodical
}

// Class Structure

type visitorClass_ struct {
	// Declare the class constants.
}

// Class Reference

func visitorClass() *visitorClass_ {
	return visitorClassReference_
}

var visitorClassReference_ = &visitorClass_{
	// Initialize the class constants.
}
