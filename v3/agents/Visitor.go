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
		var rule = predicates.GetNext()
		v.processor_.PreprocessPredicate(
			rule,
			predicatesIndex,
			predicatesCount,
		)
		v.visitPredicate(rule)
		v.processor_.PostprocessPredicate(
			rule,
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

	var argumentsIndex uint
	var arguments = function.GetArguments().GetIterator()
	for arguments.HasNext() {
		argumentsIndex++
		var rule = arguments.GetNext()
		v.visitArgument(rule)
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

func (v *visitor_) visitInverse(
	inverse doc.Inverse,
) {
	v.processor_.ProcessInverse(inverse)
}

func (v *visitor_) visitInversion(
	inversion doc.InversionLike,
) {
	var inverse = inversion.GetInverse()
	v.visitInverse(inverse)

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
	v.visitAssignment(assignment)

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

func (v *visitor_) visitLogical(
	logical any,
) {
	panic("Not yet implemented.")
}

func (v *visitor_) visitMagnitude(
	magnitude doc.MagnitudeLike,
) {
	var delimiter1 = magnitude.GetDelimiter1()
	v.processor_.ProcessDelimiter(delimiter1)
	// Visit slot 1 between terms.
	v.processor_.ProcessMagnitudeSlot(
		magnitude,
		1,
	)

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
	// Visit slot 2 between terms.
	v.processor_.ProcessMagnitudeSlot(
		magnitude,
		2,
	)

	var delimiter2 = magnitude.GetDelimiter2()
	v.processor_.ProcessDelimiter(delimiter2)
}

func (v *visitor_) visitMatchingClause(
	matchingClause doc.MatchingClauseLike,
) {
	var delimiter1 = matchingClause.GetDelimiter1()
	v.processor_.ProcessDelimiter(delimiter1)
	// Visit slot 1 between terms.
	v.processor_.ProcessMatchingClauseSlot(
		matchingClause,
		1,
	)

	var template = matchingClause.GetTemplate()
	v.processor_.PreprocessTemplate(
		template,
		1,
		1,
	)
	v.visitTemplate(template)
	v.processor_.PostprocessTemplate(
		template,
		1,
		1,
	)
	// Visit slot 2 between terms.
	v.processor_.ProcessMatchingClauseSlot(
		matchingClause,
		2,
	)

	var delimiter2 = matchingClause.GetDelimiter2()
	v.processor_.ProcessDelimiter(delimiter2)
	// Visit slot 3 between terms.
	v.processor_.ProcessMatchingClauseSlot(
		matchingClause,
		3,
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
	var identifier1 = method.GetIdentifier1()
	v.processor_.ProcessIdentifier(identifier1)
	// Visit slot 1 between terms.
	v.processor_.ProcessMethodSlot(
		method,
		1,
	)

	var invoke = method.GetInvoke()
	v.processor_.PreprocessInvoke(
		invoke,
		1,
		1,
	)
	v.visitInvoke(invoke)
	v.processor_.PostprocessInvoke(
		invoke,
		1,
		1,
	)
	// Visit slot 2 between terms.
	v.processor_.ProcessMethodSlot(
		method,
		2,
	)

	var identifier2 = method.GetIdentifier2()
	v.processor_.ProcessIdentifier(identifier2)
	// Visit slot 3 between terms.
	v.processor_.ProcessMethodSlot(
		method,
		3,
	)

	var delimiter1 = method.GetDelimiter1()
	v.processor_.ProcessDelimiter(delimiter1)
	// Visit slot 4 between terms.
	v.processor_.ProcessMethodSlot(
		method,
		4,
	)

	var argumentsIndex uint
	var arguments = method.GetArguments().GetIterator()
	var argumentsCount = uint(arguments.GetSize())
	for arguments.HasNext() {
		argumentsIndex++
		var rule = arguments.GetNext()
		v.processor_.PreprocessArgument(
			rule,
			argumentsIndex,
			argumentsCount,
		)
		v.visitArgument(rule)
		v.processor_.PostprocessArgument(
			rule,
			argumentsIndex,
			argumentsCount,
		)
	}
	// Visit slot 5 between terms.
	v.processor_.ProcessMethodSlot(
		method,
		5,
	)

	var delimiter2 = method.GetDelimiter2()
	v.processor_.ProcessDelimiter(delimiter2)
}

func (v *visitor_) visitNotarizeClause(
	notarizeClause doc.NotarizeClauseLike,
) {
	var delimiter1 = notarizeClause.GetDelimiter1()
	v.processor_.ProcessDelimiter(delimiter1)
	// Visit slot 1 between terms.
	v.processor_.ProcessNotarizeClauseSlot(
		notarizeClause,
		1,
	)

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
	// Visit slot 2 between terms.
	v.processor_.ProcessNotarizeClauseSlot(
		notarizeClause,
		2,
	)

	var delimiter2 = notarizeClause.GetDelimiter2()
	v.processor_.ProcessDelimiter(delimiter2)
	// Visit slot 3 between terms.
	v.processor_.ProcessNotarizeClauseSlot(
		notarizeClause,
		3,
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
	var delimiter = onClause.GetDelimiter()
	v.processor_.ProcessDelimiter(delimiter)
	// Visit slot 1 between terms.
	v.processor_.ProcessOnClauseSlot(
		onClause,
		1,
	)

	var failure = onClause.GetFailure()
	v.processor_.PreprocessFailure(
		failure,
		1,
		1,
	)
	v.visitFailure(failure)
	v.processor_.PostprocessFailure(
		failure,
		1,
		1,
	)
	// Visit slot 2 between terms.
	v.processor_.ProcessOnClauseSlot(
		onClause,
		2,
	)

	var matchingClausesIndex uint
	var matchingClauses = onClause.GetMatchingClauses().GetIterator()
	var matchingClausesCount = uint(matchingClauses.GetSize())
	for matchingClauses.HasNext() {
		matchingClausesIndex++
		var rule = matchingClauses.GetNext()
		v.processor_.PreprocessMatchingClause(
			rule,
			matchingClausesIndex,
			matchingClausesCount,
		)
		v.visitMatchingClause(rule)
		v.processor_.PostprocessMatchingClause(
			rule,
			matchingClausesIndex,
			matchingClausesCount,
		)
	}
}

func (v *visitor_) visitParameters(
	parameters doc.ParametersLike,
) {
	var delimiter1 = parameters.GetDelimiter1()
	v.processor_.ProcessDelimiter(delimiter1)
	// Visit slot 1 between terms.
	v.processor_.ProcessParametersSlot(
		parameters,
		1,
	)

	var associationsIndex uint
	var associations = parameters.GetAssociations().GetIterator()
	var associationsCount = uint(associations.GetSize())
	for associations.HasNext() {
		associationsIndex++
		var rule = associations.GetNext()
		v.processor_.PreprocessAssociation(
			rule,
			associationsIndex,
			associationsCount,
		)
		v.visitAssociation(rule)
		v.processor_.PostprocessAssociation(
			rule,
			associationsIndex,
			associationsCount,
		)
	}
	// Visit slot 2 between terms.
	v.processor_.ProcessParametersSlot(
		parameters,
		2,
	)

	var delimiter2 = parameters.GetDelimiter2()
	v.processor_.ProcessDelimiter(delimiter2)
}

func (v *visitor_) visitPostClause(
	postClause doc.PostClauseLike,
) {
	var delimiter1 = postClause.GetDelimiter1()
	v.processor_.ProcessDelimiter(delimiter1)
	// Visit slot 1 between terms.
	v.processor_.ProcessPostClauseSlot(
		postClause,
		1,
	)

	var message = postClause.GetMessage()
	v.processor_.PreprocessMessage(
		message,
		1,
		1,
	)
	v.visitMessage(message)
	v.processor_.PostprocessMessage(
		message,
		1,
		1,
	)
	// Visit slot 2 between terms.
	v.processor_.ProcessPostClauseSlot(
		postClause,
		2,
	)

	var delimiter2 = postClause.GetDelimiter2()
	v.processor_.ProcessDelimiter(delimiter2)
	// Visit slot 3 between terms.
	v.processor_.ProcessPostClauseSlot(
		postClause,
		3,
	)

	var bag = postClause.GetBag()
	v.processor_.PreprocessBag(
		bag,
		1,
		1,
	)
	v.visitBag(bag)
	v.processor_.PostprocessBag(
		bag,
		1,
		1,
	)
}

func (v *visitor_) visitPrecedence(
	precedence doc.PrecedenceLike,
) {
	var delimiter1 = precedence.GetDelimiter1()
	v.processor_.ProcessDelimiter(delimiter1)
	// Visit slot 1 between terms.
	v.processor_.ProcessPrecedenceSlot(
		precedence,
		1,
	)

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
	// Visit slot 2 between terms.
	v.processor_.ProcessPrecedenceSlot(
		precedence,
		2,
	)

	var delimiter2 = precedence.GetDelimiter2()
	v.processor_.ProcessDelimiter(delimiter2)
}

func (v *visitor_) visitPredicate(
	predicate doc.PredicateLike,
) {
	var operation = predicate.GetOperation()
	v.processor_.PreprocessOperation(
		operation,
		1,
		1,
	)
	v.visitOperation(operation)
	v.processor_.PostprocessOperation(
		operation,
		1,
		1,
	)
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
	var delimiter1 = procedure.GetDelimiter1()
	v.processor_.ProcessDelimiter(delimiter1)
	// Visit slot 1 between terms.
	v.processor_.ProcessProcedureSlot(
		procedure,
		1,
	)

	var linesIndex uint
	var lines = procedure.GetLines().GetIterator()
	var linesCount = uint(lines.GetSize())
	for lines.HasNext() {
		linesIndex++
		var rule = lines.GetNext()
		v.processor_.PreprocessLine(
			rule,
			linesIndex,
			linesCount,
		)
		v.visitLine(rule)
		v.processor_.PostprocessLine(
			rule,
			linesIndex,
			linesCount,
		)
	}
	// Visit slot 2 between terms.
	v.processor_.ProcessProcedureSlot(
		procedure,
		2,
	)

	var delimiter2 = procedure.GetDelimiter2()
	v.processor_.ProcessDelimiter(delimiter2)
}

func (v *visitor_) visitPublishClause(
	publishClause doc.PublishClauseLike,
) {
	var delimiter = publishClause.GetDelimiter()
	v.processor_.ProcessDelimiter(delimiter)
	// Visit slot 1 between terms.
	v.processor_.ProcessPublishClauseSlot(
		publishClause,
		1,
	)

	var event = publishClause.GetEvent()
	v.processor_.PreprocessEvent(
		event,
		1,
		1,
	)
	v.visitEvent(event)
	v.processor_.PostprocessEvent(
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
	var delimiter = referent.GetDelimiter()
	v.processor_.ProcessDelimiter(delimiter)
	// Visit slot 1 between terms.
	v.processor_.ProcessReferentSlot(
		referent,
		1,
	)

	var indirect = referent.GetIndirect()
	v.processor_.PreprocessIndirect(
		indirect,
		1,
		1,
	)
	v.visitIndirect(indirect)
	v.processor_.PostprocessIndirect(
		indirect,
		1,
		1,
	)
}

func (v *visitor_) visitRejectClause(
	rejectClause doc.RejectClauseLike,
) {
	var delimiter = rejectClause.GetDelimiter()
	v.processor_.ProcessDelimiter(delimiter)
	// Visit slot 1 between terms.
	v.processor_.ProcessRejectClauseSlot(
		rejectClause,
		1,
	)

	var message = rejectClause.GetMessage()
	v.processor_.PreprocessMessage(
		message,
		1,
		1,
	)
	v.visitMessage(message)
	v.processor_.PostprocessMessage(
		message,
		1,
		1,
	)
}

func (v *visitor_) visitRetrieveClause(
	retrieveClause doc.RetrieveClauseLike,
) {
	var delimiter1 = retrieveClause.GetDelimiter1()
	v.processor_.ProcessDelimiter(delimiter1)
	// Visit slot 1 between terms.
	v.processor_.ProcessRetrieveClauseSlot(
		retrieveClause,
		1,
	)

	v.visitRecipient(recipient)

	// Visit slot 2 between terms.
	v.processor_.ProcessRetrieveClauseSlot(
		retrieveClause,
		2,
	)

	var delimiter2 = retrieveClause.GetDelimiter2()
	v.processor_.ProcessDelimiter(delimiter2)
	// Visit slot 3 between terms.
	v.processor_.ProcessRetrieveClauseSlot(
		retrieveClause,
		3,
	)

	var bag = retrieveClause.GetBag()
	v.processor_.PreprocessBag(
		bag,
		1,
		1,
	)
	v.visitBag(bag)
	v.processor_.PostprocessBag(
		bag,
		1,
		1,
	)
}

func (v *visitor_) visitReturnClause(
	returnClause doc.ReturnClauseLike,
) {
	var delimiter = returnClause.GetDelimiter()
	v.processor_.ProcessDelimiter(delimiter)
	// Visit slot 1 between terms.
	v.processor_.ProcessReturnClauseSlot(
		returnClause,
		1,
	)

	var result = returnClause.GetResult()
	v.processor_.PreprocessResult(
		result,
		1,
		1,
	)
	v.visitResult(result)
	v.processor_.PostprocessResult(
		result,
		1,
		1,
	)
}

func (v *visitor_) visitSaveClause(
	saveClause doc.SaveClauseLike,
) {
	var delimiter1 = saveClause.GetDelimiter1()
	v.processor_.ProcessDelimiter(delimiter1)
	// Visit slot 1 between terms.
	v.processor_.ProcessSaveClauseSlot(
		saveClause,
		1,
	)

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
	// Visit slot 2 between terms.
	v.processor_.ProcessSaveClauseSlot(
		saveClause,
		2,
	)

	var delimiter2 = saveClause.GetDelimiter2()
	v.processor_.ProcessDelimiter(delimiter2)
	// Visit slot 3 between terms.
	v.processor_.ProcessSaveClauseSlot(
		saveClause,
		3,
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
	var delimiter = selectClause.GetDelimiter()
	v.processor_.ProcessDelimiter(delimiter)
	// Visit slot 1 between terms.
	v.processor_.ProcessSelectClauseSlot(
		selectClause,
		1,
	)

	var target = selectClause.GetTarget()
	v.processor_.PreprocessTarget(
		target,
		1,
		1,
	)
	v.visitTarget(target)
	v.processor_.PostprocessTarget(
		target,
		1,
		1,
	)
	// Visit slot 2 between terms.
	v.processor_.ProcessSelectClauseSlot(
		selectClause,
		2,
	)

	var matchingClausesIndex uint
	var matchingClauses = selectClause.GetMatchingClauses().GetIterator()
	var matchingClausesCount = uint(matchingClauses.GetSize())
	for matchingClauses.HasNext() {
		matchingClausesIndex++
		var rule = matchingClauses.GetNext()
		v.processor_.PreprocessMatchingClause(
			rule,
			matchingClausesIndex,
			matchingClausesCount,
		)
		v.visitMatchingClause(rule)
		v.processor_.PostprocessMatchingClause(
			rule,
			matchingClausesIndex,
			matchingClausesCount,
		)
	}
}

func (v *visitor_) visitStatement(
	statement doc.StatementLike,
) {
	var mainClause = statement.GetMainClause()
	v.processor_.PreprocessMainClause(
		mainClause,
		1,
		1,
	)
	v.visitMainClause(mainClause)
	v.processor_.PostprocessMainClause(
		mainClause,
		1,
		1,
	)
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

	var delimiter1 = subcomponent.GetDelimiter1()
	v.processor_.ProcessDelimiter(delimiter1)
	// Visit slot 2 between terms.
	v.processor_.ProcessSubcomponentSlot(
		subcomponent,
		2,
	)

	var indexesIndex uint
	var indexes = subcomponent.GetIndexes().GetIterator()
	var indexesCount = uint(indexes.GetSize())
	for indexes.HasNext() {
		indexesIndex++
		var rule = indexes.GetNext()
		v.processor_.PreprocessIndex(
			rule,
			indexesIndex,
			indexesCount,
		)
		v.visitIndex(rule)
		v.processor_.PostprocessIndex(
			rule,
			indexesIndex,
			indexesCount,
		)
	}
	// Visit slot 3 between terms.
	v.processor_.ProcessSubcomponentSlot(
		subcomponent,
		3,
	)

	var delimiter2 = subcomponent.GetDelimiter2()
	v.processor_.ProcessDelimiter(delimiter2)
}

func (v *visitor_) visitSubject(
	subject any,
) {
	panic("Not yet implemented.")
}

func (v *visitor_) visitThrowClause(
	throwClause doc.ThrowClauseLike,
) {
	var delimiter = throwClause.GetDelimiter()
	v.processor_.ProcessDelimiter(delimiter)
	// Visit slot 1 between terms.
	v.processor_.ProcessThrowClauseSlot(
		throwClause,
		1,
	)

	var exception = throwClause.GetException()
	v.processor_.PreprocessException(
		exception,
		1,
		1,
	)
	v.visitException(exception)
	v.processor_.PostprocessException(
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
