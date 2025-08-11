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
	ass "github.com/bali-nebula/go-bali-documents/v3/assembly"
	doc "github.com/bali-nebula/go-bali-documents/v3/documents"
	fra "github.com/craterdog/go-component-framework/v7"
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
	v.visitDocument(document)
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
	switch actual := argument.(type) {
	case string:
		v.processor_.ProcessIdentifier(actual)
	default:
		v.visitPrimitive(argument)
	}
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
	switch actual := component.(type) {
	case doc.RangeLike:
		v.visitRange(actual)
	case doc.AttributesLike:
		v.visitAttributes(actual)
	case doc.EntitiesLike:
		v.visitEntities(actual)
	case doc.ProcedureLike:
		v.visitProcedure(actual)
	default:
		v.visitPrimitive(component)
	}
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
	var method = doClause.GetMethod()
	v.visitMethod(method)
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

func (v *visitor_) visitEntities(
	entities doc.EntitiesLike,
) {
	var itemsIndex uint
	var items = entities.GetItems().GetIterator()
	var itemsCount = uint(items.GetSize())
	for items.HasNext() {
		itemsIndex++
		var document = items.GetNext()
		v.processor_.PreprocessDocument(
			document,
			itemsIndex,
			itemsCount,
		)
		v.visitDocument(document)
		v.processor_.PostprocessDocument(
			document,
			itemsIndex,
			itemsCount,
		)
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
	switch actual := index.(type) {
	case string:
		v.processor_.ProcessIdentifier(actual)
	default:
		v.visitPrimitive(index)
	}
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
	switch actual := line.(type) {
	case doc.StatementLike:
		v.visitStatement(actual)
	case string:
		v.processor_.ProcessAnnotation(actual)
	}
}

func (v *visitor_) visitLogical(
	logical any,
) {
	switch actual := logical.(type) {
	case doc.DocumentLike:
		v.visitDocument(actual)
	case doc.SubcomponentLike:
		v.visitSubcomponent(actual)
	case doc.PrecedenceLike:
		v.visitPrecedence(actual)
	case doc.ReferentLike:
		v.visitReferent(actual)
	case doc.ComplementLike:
		v.visitComplement(actual)
	case doc.FunctionLike:
		v.visitFunction(actual)
	case doc.MethodLike:
		v.visitMethod(actual)
	case string:
		v.processor_.ProcessIdentifier(actual)
	}
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
	switch actual := mainClause.(type) {
	case doc.IfClauseLike:
		v.visitIfClause(actual)
	case doc.SelectClauseLike:
		v.visitSelectClause(actual)
	case doc.WhileClauseLike:
		v.visitWhileClause(actual)
	case doc.WithClauseLike:
		v.visitWithClause(actual)
	case doc.ContinueClauseLike:
		v.visitContinueClause(actual)
	case doc.BreakClauseLike:
		v.visitBreakClause(actual)
	case doc.ReturnClauseLike:
		v.visitReturnClause(actual)
	case doc.ThrowClauseLike:
		v.visitThrowClause(actual)
	case doc.DoClauseLike:
		v.visitDoClause(actual)
	case doc.LetClauseLike:
		v.visitLetClause(actual)
	case doc.PostClauseLike:
		v.visitPostClause(actual)
	case doc.RetrieveClauseLike:
		v.visitRetrieveClause(actual)
	case doc.AcceptClauseLike:
		v.visitAcceptClause(actual)
	case doc.RejectClauseLike:
		v.visitRejectClause(actual)
	case doc.PublishClauseLike:
		v.visitPublishClause(actual)
	case doc.CheckoutClauseLike:
		v.visitCheckoutClause(actual)
	case doc.SaveClauseLike:
		v.visitSaveClause(actual)
	case doc.DiscardClauseLike:
		v.visitDiscardClause(actual)
	case doc.NotarizeClauseLike:
		v.visitNotarizeClause(actual)
	}
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
	switch actual := numerical.(type) {
	case doc.DocumentLike:
		v.visitDocument(actual)
	case doc.SubcomponentLike:
		v.visitSubcomponent(actual)
	case doc.PrecedenceLike:
		v.visitPrecedence(actual)
	case doc.ReferentLike:
		v.visitReferent(actual)
	case doc.InversionLike:
		v.visitInversion(actual)
	case doc.MagnitudeLike:
		v.visitMagnitude(actual)
	case doc.FunctionLike:
		v.visitFunction(actual)
	case doc.MethodLike:
		v.visitMethod(actual)
	case string:
		v.processor_.ProcessIdentifier(actual)
	}
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
	switch actual := primitive.(type) {
	case fra.AngleLike:
		v.processor_.ProcessAngle(actual)
	case fra.BooleanLike:
		v.processor_.ProcessBoolean(actual)
	case fra.DurationLike:
		v.processor_.ProcessDuration(actual)
	case fra.GlyphLike:
		v.processor_.ProcessGlyph(actual)
	case fra.MomentLike:
		v.processor_.ProcessMoment(actual)
	case fra.NumberLike:
		v.processor_.ProcessNumber(actual)
	case fra.PercentageLike:
		v.processor_.ProcessPercentage(actual)
	case fra.ProbabilityLike:
		v.processor_.ProcessProbability(actual)
	case fra.ResourceLike:
		v.processor_.ProcessResource(actual)
	case fra.SymbolLike:
		v.processor_.ProcessSymbol(actual)
	case fra.BinaryLike:
		v.processor_.ProcessBinary(actual)
	case ass.BytecodeLike:
		v.processor_.ProcessBytecode(actual)
	case fra.NameLike:
		v.processor_.ProcessName(actual)
	case fra.NarrativeLike:
		v.processor_.ProcessNarrative(actual)
	case fra.PatternLike:
		v.processor_.ProcessPattern(actual)
	case fra.QuoteLike:
		v.processor_.ProcessQuote(actual)
	case fra.TagLike:
		v.processor_.ProcessTag(actual)
	case fra.VersionLike:
		v.processor_.ProcessVersion(actual)
	}
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
	switch actual := recipient.(type) {
	case doc.SubcomponentLike:
		v.visitSubcomponent(actual)
	case fra.SymbolLike:
		v.processor_.ProcessSymbol(actual)
	}
}

func (v *visitor_) visitReference(
	reference any,
) {
	switch actual := reference.(type) {
	case doc.DocumentLike:
		v.visitDocument(actual)
	case doc.SubcomponentLike:
		v.visitSubcomponent(actual)
	case doc.ReferentLike:
		v.visitReferent(actual)
	case doc.FunctionLike:
		v.visitFunction(actual)
	case doc.MethodLike:
		v.visitMethod(actual)
	case string:
		v.processor_.ProcessIdentifier(actual)
	}
}

func (v *visitor_) visitReferent(
	referent doc.ReferentLike,
) {
	var reference = referent.GetReference()
	v.visitReference(reference)
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
	var expression = selectClause.GetExpression()
	v.visitExpression(expression)

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
	switch actual := subject.(type) {
	case doc.DocumentLike:
		v.visitDocument(actual)
	case doc.SubcomponentLike:
		v.visitSubcomponent(actual)
	case doc.PrecedenceLike:
		v.visitPrecedence(actual)
	case doc.ReferentLike:
		v.visitReferent(actual)
	case doc.ComplementLike:
		v.visitComplement(actual)
	case doc.InversionLike:
		v.visitInversion(actual)
	case doc.MagnitudeLike:
		v.visitMagnitude(actual)
	case doc.FunctionLike:
		v.visitFunction(actual)
	case doc.MethodLike:
		v.visitMethod(actual)
	case string:
		v.processor_.ProcessIdentifier(actual)
	}
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

	// Visit slot 1 between terms.
	v.processor_.ProcessWhileClauseSlot(
		whileClause,
		1,
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
	var variable = withClause.GetVariable()
	v.processor_.ProcessSymbol(variable)

	// Visit slot 1 between terms.
	v.processor_.ProcessWithClauseSlot(
		withClause,
		1,
	)

	var sequence = withClause.GetSequence()
	v.processor_.PreprocessExpression(
		sequence,
		1,
		1,
	)
	v.visitExpression(sequence)
	v.processor_.PostprocessExpression(
		sequence,
		1,
		1,
	)
	// Visit slot 2 between terms.
	v.processor_.ProcessWithClauseSlot(
		withClause,
		2,
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
