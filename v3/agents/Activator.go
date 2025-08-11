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
	fmt "fmt"
	ass "github.com/bali-nebula/go-bali-documents/v3/assembly"
	doc "github.com/bali-nebula/go-bali-documents/v3/documents"
	not "github.com/bali-nebula/go-document-notation/v3"
	fra "github.com/craterdog/go-component-framework/v7"
	uti "github.com/craterdog/go-missing-utilities/v7"
)

// CLASS INTERFACE

// Access Function

func ActivatorClass() ActivatorClassLike {
	return activatorClass()
}

// Constructor Methods

func (c *activatorClass_) Activator() ActivatorLike {
	var instance = &activator_{
		// Initialize the instance attributes.
		stack_: fra.Stack[any](),

		// Initialize the inherited aspects.
		Methodical: not.ProcessorClass().Processor(),
	}
	return instance
}

// Constant Methods

// Function Methods

// INSTANCE INTERFACE

// Principal Methods

func (v *activator_) GetClass() ActivatorClassLike {
	return activatorClass()
}

func (v *activator_) ActivateDocument(
	document not.DocumentLike,
) doc.DocumentLike {
	not.VisitorClass().Visitor(v).VisitDocument(document)
	if v.stack_.GetSize() != 1 {
		var message = fmt.Sprintf(
			"Internal Error: the activator stack is corrupted: %v",
			v.stack_,
		)
		panic(message)
	}
	return v.stack_.RemoveLast().(doc.DocumentLike)
}

// Attribute Methods

// not.Methodical Methods

func (v *activator_) ProcessAngle(
	angle string,
) {
	v.stack_.AddValue(fra.AngleFromString(angle))
}

func (v *activator_) ProcessBinary(
	binary string,
) {
	v.stack_.AddValue(fra.BinaryFromString(binary))
}

func (v *activator_) ProcessBoolean(
	boolean string,
) {
	v.stack_.AddValue(fra.BooleanFromString(boolean))
}

func (v *activator_) ProcessBytecode(
	bytecode string,
) {
	v.stack_.AddValue(ass.BytecodeClass().BytecodeFromString(bytecode))
}

func (v *activator_) ProcessComment(
	comment string,
) {
	v.stack_.AddValue(comment)
}

func (v *activator_) ProcessDuration(
	duration string,
) {
	v.stack_.AddValue(fra.DurationFromString(duration))
}

func (v *activator_) ProcessGlyph(
	glyph string,
) {
	v.stack_.AddValue(fra.GlyphFromString(glyph))
}

func (v *activator_) ProcessIdentifier(
	identifier string,
) {
	v.stack_.AddValue(identifier)
}

func (v *activator_) ProcessMoment(
	moment string,
) {
	v.stack_.AddValue(fra.MomentFromString(moment))
}

func (v *activator_) ProcessName(
	name string,
) {
	v.stack_.AddValue(fra.NameFromString(name))
}

func (v *activator_) ProcessNarrative(
	narrative string,
) {
	v.stack_.AddValue(fra.NarrativeFromString(narrative))
}

func (v *activator_) ProcessNote(
	note string,
) {
	v.stack_.AddValue(note)
}

func (v *activator_) ProcessNumber(
	number string,
) {
	v.stack_.AddValue(fra.NumberFromString(number))
}

func (v *activator_) ProcessPattern(
	pattern string,
) {
	v.stack_.AddValue(fra.PatternFromString(pattern))
}

func (v *activator_) ProcessPercentage(
	percentage string,
) {
	v.stack_.AddValue(fra.PercentageFromString(percentage))
}

func (v *activator_) ProcessProbability(
	probability string,
) {
	v.stack_.AddValue(fra.ProbabilityFromString(probability))
}

func (v *activator_) ProcessQuote(
	quote string,
) {
	v.stack_.AddValue(fra.QuoteFromString(quote))
}

func (v *activator_) ProcessResource(
	resource string,
) {
	v.stack_.AddValue(fra.ResourceFromString(resource))
}

func (v *activator_) ProcessSymbol(
	symbol string,
) {
	v.stack_.AddValue(fra.SymbolFromString(symbol))
}

func (v *activator_) ProcessTag(
	tag string,
) {
	v.stack_.AddValue(fra.TagFromString(tag))
}

func (v *activator_) ProcessVersion(
	version string,
) {
	v.stack_.AddValue(fra.VersionFromString(version))
}

func (v *activator_) PostprocessAcceptClause(
	acceptClause not.AcceptClauseLike,
	index_ uint,
	count_ uint,
) {
	var message = v.stack_.RemoveLast().(doc.ExpressionLike)
	v.stack_.AddValue(doc.AcceptClauseClass().AcceptClause(message))
}

func (v *activator_) PostprocessAssignment(
	assignment not.AssignmentLike,
	index_ uint,
	count_ uint,
) {
	v.stack_.AddValue(assignment.GetAny().(string))
}

func (v *activator_) PostprocessAttributes(
	attributes not.AttributesLike,
	index_ uint,
	count_ uint,
) {
	var catalog = fra.Catalog[any, doc.DocumentLike]()
	var associations = attributes.GetAssociations()
	var iterator = associations.GetIterator()
	for iterator.HasNext() {
		var document = v.stack_.RemoveLast().(doc.DocumentLike)
		var primitive = v.stack_.RemoveLast()
		catalog.SetValue(primitive, document)
	}
	catalog.ReverseValues() // They were pulled off the stack in reverse order.
	v.stack_.AddValue(catalog)
}

func (v *activator_) PostprocessBreakClause(
	breakClause not.BreakClauseLike,
	index_ uint,
	count_ uint,
) {
	v.stack_.AddValue(doc.BreakClauseClass().BreakClause())
}

func (v *activator_) ProcessCheckoutClauseSlot(
	checkoutClause not.CheckoutClauseLike,
	slot_ uint,
) {
	switch slot_ {
	case 3:
		if uti.IsUndefined(checkoutClause.GetOptionalAtLevel()) {
			v.stack_.AddValue(nil)
		}
	}
}

func (v *activator_) PostprocessCheckoutClause(
	checkoutClause not.CheckoutClauseLike,
	index_ uint,
	count_ uint,
) {
	var cited = v.stack_.RemoveLast().(doc.ExpressionLike)
	var atLevel = v.stack_.RemoveLast().(doc.ExpressionLike)
	var recipient = v.stack_.RemoveLast()
	v.stack_.AddValue(
		doc.CheckoutClauseClass().CheckoutClause(recipient, atLevel, cited),
	)
}

func (v *activator_) PostprocessComplement(
	complement not.ComplementLike,
	index_ uint,
	count_ uint,
) {
	var logical = v.stack_.RemoveLast()
	v.stack_.AddValue(doc.ComplementClass().Complement(logical))
}

func (v *activator_) PostprocessContinueClause(
	continueClause not.ContinueClauseLike,
	index_ uint,
	count_ uint,
) {
	v.stack_.AddValue(doc.ContinueClauseClass().ContinueClause())
}

func (v *activator_) PostprocessDiscardClause(
	discardClause not.DiscardClauseLike,
	index_ uint,
	count_ uint,
) {
	var draft = v.stack_.RemoveLast().(doc.ExpressionLike)
	v.stack_.AddValue(doc.DiscardClauseClass().DiscardClause(draft))
}

func (v *activator_) PostprocessDoClause(
	doClause not.DoClauseLike,
	index_ uint,
	count_ uint,
) {
	var method = v.stack_.RemoveLast().(doc.MethodLike)
	v.stack_.AddValue(doc.DoClauseClass().DoClause(method))
}

func (v *activator_) ProcessDocumentSlot(
	document not.DocumentLike,
	slot_ uint,
) {
	switch slot_ {
	case 2:
		if uti.IsUndefined(document.GetOptionalParameters()) {
			v.stack_.AddValue(nil)
		}
	}
}

func (v *activator_) PostprocessDocument(
	document not.DocumentLike,
	index_ uint,
	count_ uint,
) {
	if uti.IsUndefined(document.GetOptionalNote()) {
		v.stack_.AddValue("")
	}
	var note = v.stack_.RemoveLast().(string)
	var parameters = v.stack_.RemoveLast().(doc.ParametersLike)
	var component = v.stack_.RemoveLast()
	v.stack_.AddValue(doc.DocumentClass().Document(component, parameters, note))
}

func (v *activator_) PostprocessEntities(
	entities not.EntitiesLike,
	index_ uint,
	count_ uint,
) {
	var list = fra.List[doc.DocumentLike]()
	var items = entities.GetItems()
	var iterator = items.GetIterator()
	for iterator.HasNext() {
		var document = v.stack_.RemoveLast().(doc.DocumentLike)
		list.AppendValue(document)
	}
	list.ReverseValues() // They were pulled off the stack in reverse order.
	v.stack_.AddValue(list)
}

func (v *activator_) PostprocessExpression(
	expression not.ExpressionLike,
	index_ uint,
	count_ uint,
) {
	var predicates = v.stack_.RemoveLast().(fra.ListLike[doc.PredicateLike])
	var subject = v.stack_.RemoveLast()
	v.stack_.AddValue(doc.ExpressionClass().Expression(subject, predicates))
}

func (v *activator_) PostprocessFunction(
	function not.FunctionLike,
	index_ uint,
	count_ uint,
) {
	var arguments = v.stack_.RemoveLast().(fra.ListLike[any])
	var identifier = v.stack_.RemoveLast().(string)
	v.stack_.AddValue(doc.FunctionClass().Function(identifier, arguments))
}

func (v *activator_) PostprocessIfClause(
	ifClause not.IfClauseLike,
	index_ uint,
	count_ uint,
) {
	var procedure = v.stack_.RemoveLast().(doc.ProcedureLike)
	var condition = v.stack_.RemoveLast().(doc.ExpressionLike)
	v.stack_.AddValue(doc.IfClauseClass().IfClause(condition, procedure))
}

func (v *activator_) PostprocessInversion(
	inversion not.InversionLike,
	index_ uint,
	count_ uint,
) {
	var numerical = v.stack_.RemoveLast()
	var inverse = v.stack_.RemoveLast().(doc.Inverse)
	v.stack_.AddValue(doc.InversionClass().Inversion(inverse, numerical))
}

func (v *activator_) PostprocessLeft(
	left not.LeftLike,
	index_ uint,
	count_ uint,
) {
	var extent doc.Extent
	switch left.GetAny().(string) {
	case "[":
		extent = doc.Inclusive
	case "(":
		extent = doc.Exclusive
	}
	v.stack_.AddValue(extent)
}

func (v *activator_) PostprocessLetClause(
	letClause not.LetClauseLike,
	index_ uint,
	count_ uint,
) {
	var expression = v.stack_.RemoveLast().(doc.ExpressionLike)
	var assignment = v.stack_.RemoveLast().(doc.Assignment)
	var recipient = v.stack_.RemoveLast()
	v.stack_.AddValue(
		doc.LetClauseClass().LetClause(recipient, assignment, expression),
	)
}

func (v *activator_) PostprocessMagnitude(
	magnitude not.MagnitudeLike,
	index_ uint,
	count_ uint,
) {
	var expression = v.stack_.RemoveLast().(doc.ExpressionLike)
	v.stack_.AddValue(doc.MagnitudeClass().Magnitude(expression))
}

func (v *activator_) PostprocessMatchingClause(
	matchingClause not.MatchingClauseLike,
	index_ uint,
	count_ uint,
) {
	var procedure = v.stack_.RemoveLast().(doc.ProcedureLike)
	var template = v.stack_.RemoveLast().(doc.ExpressionLike)
	v.stack_.AddValue(
		doc.MatchingClauseClass().MatchingClause(template, procedure),
	)
}

func (v *activator_) PostprocessMethod(
	method not.MethodLike,
	index_ uint,
	count_ uint,
) {
	var arguments = v.stack_.RemoveLast().(fra.ListLike[any])
	var identifier = v.stack_.RemoveLast().(string)
	var invoke = v.stack_.RemoveLast().(doc.Invoke)
	var target = v.stack_.RemoveLast().(string)
	v.stack_.AddValue(
		doc.MethodClass().Method(target, invoke, identifier, arguments),
	)
}

func (v *activator_) PostprocessNotarizeClause(
	notarizeClause not.NotarizeClauseLike,
	index_ uint,
	count_ uint,
) {
	var cited = v.stack_.RemoveLast().(doc.ExpressionLike)
	var draft = v.stack_.RemoveLast().(doc.ExpressionLike)
	v.stack_.AddValue(
		doc.NotarizeClauseClass().NotarizeClause(draft, cited),
	)
}

func (v *activator_) PostprocessOnClause(
	onClause not.OnClauseLike,
	index_ uint,
	count_ uint,
) {
	var list = fra.List[doc.MatchingClauseLike]()
	var matchingClauses = onClause.GetMatchingClauses()
	var iterator = matchingClauses.GetIterator()
	for iterator.HasNext() {
		var matchingClause = v.stack_.RemoveLast().(doc.MatchingClauseLike)
		list.AppendValue(matchingClause)
	}
	list.ReverseValues() // They were pulled off the stack in reverse order.
	var symbol = v.stack_.RemoveLast().(fra.SymbolLike)
	v.stack_.AddValue(
		doc.OnClauseClass().OnClause(symbol, list),
	)
}

func (v *activator_) PostprocessParameters(
	parameters not.ParametersLike,
	index_ uint,
	count_ uint,
) {
	var catalog = fra.Catalog[fra.SymbolLike, doc.DocumentLike]()
	var associations = parameters.GetAssociations()
	var iterator = associations.GetIterator()
	for iterator.HasNext() {
		var document = v.stack_.RemoveLast().(doc.DocumentLike)
		var symbol = v.stack_.RemoveLast().(fra.SymbolLike)
		catalog.SetValue(symbol, document)
	}
	catalog.ReverseValues() // They were pulled off the stack in reverse order.
	v.stack_.AddValue(catalog)
}

func (v *activator_) PostprocessPostClause(
	postClause not.PostClauseLike,
	index_ uint,
	count_ uint,
) {
	var bag = v.stack_.RemoveLast().(doc.ExpressionLike)
	var message = v.stack_.RemoveLast().(doc.ExpressionLike)
	v.stack_.AddValue(
		doc.PostClauseClass().PostClause(message, bag),
	)
}

func (v *activator_) PostprocessPrecedence(
	precedence not.PrecedenceLike,
	index_ uint,
	count_ uint,
) {
	var expression = v.stack_.RemoveLast().(doc.ExpressionLike)
	v.stack_.AddValue(doc.PrecedenceClass().Precedence(expression))
}

func (v *activator_) PostprocessPredicate(
	predicate not.PredicateLike,
	index_ uint,
	count_ uint,
) {
	var expression = v.stack_.RemoveLast().(doc.ExpressionLike)
	var operator = v.stack_.RemoveLast().(doc.Operator)
	v.stack_.AddValue(doc.PredicateClass().Predicate(operator, expression))
}

func (v *activator_) PostprocessProcedure(
	procedure not.ProcedureLike,
	index_ uint,
	count_ uint,
) {
	var list = fra.List[any]()
	var lines = procedure.GetLines()
	var iterator = lines.GetIterator()
	for iterator.HasNext() {
		var line = v.stack_.RemoveLast()
		list.AppendValue(line)
	}
	list.ReverseValues() // They were pulled off the stack in reverse order.
	v.stack_.AddValue(list)
}

func (v *activator_) PostprocessPublishClause(
	publishClause not.PublishClauseLike,
	index_ uint,
	count_ uint,
) {
	var event = v.stack_.RemoveLast().(doc.ExpressionLike)
	v.stack_.AddValue(doc.PublishClauseClass().PublishClause(event))
}

func (v *activator_) PostprocessRange(
	range_ not.RangeLike,
	index_ uint,
	count_ uint,
) {
	var right = v.stack_.RemoveLast().(doc.Extent)
	var last = v.stack_.RemoveLast()
	var first = v.stack_.RemoveLast()
	var left = v.stack_.RemoveLast().(doc.Extent)
	v.stack_.AddValue(doc.RangeClass().Range(left, first, last, right))
}

func (v *activator_) PostprocessReferent(
	referent not.ReferentLike,
	index_ uint,
	count_ uint,
) {
	var indirect = v.stack_.RemoveLast()
	v.stack_.AddValue(doc.ReferentClass().Referent(indirect))
}

func (v *activator_) PostprocessRejectClause(
	rejectClause not.RejectClauseLike,
	index_ uint,
	count_ uint,
) {
	var message = v.stack_.RemoveLast().(doc.ExpressionLike)
	v.stack_.AddValue(doc.RejectClauseClass().RejectClause(message))
}

func (v *activator_) PostprocessRetrieveClause(
	retrieveClause not.RetrieveClauseLike,
	index_ uint,
	count_ uint,
) {
	var bag = v.stack_.RemoveLast().(doc.ExpressionLike)
	var recipient = v.stack_.RemoveLast()
	v.stack_.AddValue(
		doc.RetrieveClauseClass().RetrieveClause(recipient, bag),
	)
}

func (v *activator_) PostprocessReturnClause(
	returnClause not.ReturnClauseLike,
	index_ uint,
	count_ uint,
) {
	var result = v.stack_.RemoveLast().(doc.ExpressionLike)
	v.stack_.AddValue(doc.ReturnClauseClass().ReturnClause(result))
}

func (v *activator_) PostprocessRight(
	right not.RightLike,
	index_ uint,
	count_ uint,
) {
	var extent doc.Extent
	switch right.GetAny().(string) {
	case "]":
		extent = doc.Inclusive
	case ")":
		extent = doc.Exclusive
	}
	v.stack_.AddValue(extent)
}

func (v *activator_) PostprocessSaveClause(
	saveClause not.SaveClauseLike,
	index_ uint,
	count_ uint,
) {
	var cited = v.stack_.RemoveLast().(doc.ExpressionLike)
	var draft = v.stack_.RemoveLast().(doc.ExpressionLike)
	v.stack_.AddValue(
		doc.SaveClauseClass().SaveClause(draft, cited),
	)
}

func (v *activator_) PostprocessSelectClause(
	selectClause not.SelectClauseLike,
	index_ uint,
	count_ uint,
) {
	var list = fra.List[doc.MatchingClauseLike]()
	var matchingClauses = selectClause.GetMatchingClauses()
	var iterator = matchingClauses.GetIterator()
	for iterator.HasNext() {
		var matchingClause = v.stack_.RemoveLast().(doc.MatchingClauseLike)
		list.AppendValue(matchingClause)
	}
	list.ReverseValues() // They were pulled off the stack in reverse order.
	var expression = v.stack_.RemoveLast().(doc.ExpressionLike)
	v.stack_.AddValue(
		doc.SelectClauseClass().SelectClause(expression, list),
	)
}

func (v *activator_) PostprocessStatement(
	statement not.StatementLike,
	index_ uint,
	count_ uint,
) {
	if uti.IsUndefined(statement.GetOptionalOnClause()) {
		v.stack_.AddValue(nil)
	}
	var onClause = v.stack_.RemoveLast().(doc.OnClauseLike)
	var mainClause = v.stack_.RemoveLast()
	v.stack_.AddValue(doc.StatementClass().Statement(mainClause, onClause))
}

func (v *activator_) PostprocessSubcomponent(
	subcomponent not.SubcomponentLike,
	index_ uint,
	count_ uint,
) {
	var indexes = v.stack_.RemoveLast().(fra.ListLike[any])
	var identifier = v.stack_.RemoveLast().(string)
	v.stack_.AddValue(doc.SubcomponentClass().Subcomponent(identifier, indexes))
}

func (v *activator_) PostprocessThrowClause(
	throwClause not.ThrowClauseLike,
	index_ uint,
	count_ uint,
) {
	var exception = v.stack_.RemoveLast().(doc.ExpressionLike)
	v.stack_.AddValue(doc.ThrowClauseClass().ThrowClause(exception))
}

func (v *activator_) PostprocessWhileClause(
	whileClause not.WhileClauseLike,
	index_ uint,
	count_ uint,
) {
	var procedure = v.stack_.RemoveLast().(doc.ProcedureLike)
	var condition = v.stack_.RemoveLast().(doc.ExpressionLike)
	v.stack_.AddValue(doc.WhileClauseClass().WhileClause(condition, procedure))
}

func (v *activator_) PostprocessWithClause(
	withClause not.WithClauseLike,
	index_ uint,
	count_ uint,
) {
	var procedure = v.stack_.RemoveLast().(doc.ProcedureLike)
	var sequence = v.stack_.RemoveLast().(doc.ExpressionLike)
	var variable = v.stack_.RemoveLast().(fra.SymbolLike)
	v.stack_.AddValue(
		doc.WithClauseClass().WithClause(variable, sequence, procedure),
	)
}

// PROTECTED INTERFACE

// Private Methods

// Instance Structure

type activator_ struct {
	// Declare the instance attributes.
	stack_ fra.StackLike[any]

	// Declare the inherited aspects.
	not.Methodical
}

// Class Structure

type activatorClass_ struct {
	// Declare the class constants.
}

// Class Reference

func activatorClass() *activatorClass_ {
	return activatorClassReference_
}

var activatorClassReference_ = &activatorClass_{
	// Initialize the class constants.
}
