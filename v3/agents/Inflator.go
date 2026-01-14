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

package agents

import (
	fmt "fmt"
	doc "github.com/bali-nebula/go-bali-documents/v3/documents"
	not "github.com/bali-nebula/go-document-notation/v3"
	com "github.com/craterdog/go-essential-composites/v8"
	pri "github.com/craterdog/go-essential-primitives/v8"
	uti "github.com/craterdog/go-essential-utilities/v8"
)

// CLASS INTERFACE

// Access Function

func InflatorClass() InflatorClassLike {
	return inflatorClass()
}

// Constructor Methods

func (c *inflatorClass_) Inflator() InflatorLike {
	var instance = &inflator_{
		// Initialize the instance attributes.
		stack_: com.StackWithCapacity[any](256),

		// Initialize the inherited aspects.
		Methodical: not.ProcessorClass().Processor(),
	}
	return instance
}

// Constant Methods

// Function Methods

// INSTANCE INTERFACE

// Principal Methods

func (v *inflator_) GetClass() InflatorClassLike {
	return inflatorClass()
}

func (v *inflator_) InflateDocument(
	document not.DocumentLike,
) doc.DocumentLike {
	not.VisitorClass().Visitor(v).VisitDocument(document)
	return v.stack_.RemoveLast().(doc.DocumentLike)
}

// Attribute Methods

// not.Methodical Methods

func (v *inflator_) ProcessAngle(
	angle string,
) {
	v.stack_.AddValue(pri.AngleClass().AngleFromSource(angle))
}

func (v *inflator_) ProcessBinary(
	binary string,
) {
	v.stack_.AddValue(pri.BinaryClass().BinaryFromSource(binary))
}

func (v *inflator_) ProcessBoolean(
	boolean string,
) {
	v.stack_.AddValue(pri.BooleanClass().BooleanFromSource(boolean))
}

func (v *inflator_) ProcessBytecode(
	bytecode string,
) {
	v.stack_.AddValue(pri.BytecodeClass().BytecodeFromSource(bytecode))
}

func (v *inflator_) ProcessComment(
	comment string,
) {
	v.stack_.AddValue(comment)
}

func (v *inflator_) ProcessDuration(
	duration string,
) {
	v.stack_.AddValue(pri.DurationClass().DurationFromSource(duration))
}

func (v *inflator_) ProcessGlyph(
	glyph string,
) {
	v.stack_.AddValue(pri.GlyphClass().GlyphFromSource(glyph))
}

func (v *inflator_) ProcessIdentifier(
	identifier string,
) {
	v.stack_.AddValue(pri.IdentifierClass().IdentifierFromSource(identifier))
}

func (v *inflator_) ProcessMoment(
	moment string,
) {
	v.stack_.AddValue(pri.MomentClass().MomentFromSource(moment))
}

func (v *inflator_) ProcessName(
	name string,
) {
	v.stack_.AddValue(pri.NameClass().NameFromSource(name))
}

func (v *inflator_) ProcessNarrative(
	narrative string,
) {
	v.stack_.AddValue(pri.NarrativeClass().NarrativeFromSource(narrative))
}

func (v *inflator_) ProcessNote(
	note string,
) {
	v.stack_.AddValue(note)
}

func (v *inflator_) ProcessNumber(
	number string,
) {
	v.stack_.AddValue(pri.NumberClass().NumberFromSource(number))
}

func (v *inflator_) ProcessPattern(
	pattern string,
) {
	v.stack_.AddValue(pri.PatternClass().PatternFromSource(pattern))
}

func (v *inflator_) ProcessPercentage(
	percentage string,
) {
	v.stack_.AddValue(pri.PercentageClass().PercentageFromSource(percentage))
}

func (v *inflator_) ProcessProbability(
	probability string,
) {
	v.stack_.AddValue(pri.ProbabilityClass().ProbabilityFromSource(probability))
}

func (v *inflator_) ProcessQuote(
	quote string,
) {
	v.stack_.AddValue(pri.QuoteClass().QuoteFromSource(quote))
}

func (v *inflator_) ProcessResource(
	resource string,
) {
	v.stack_.AddValue(pri.ResourceClass().ResourceFromSource(resource))
}

func (v *inflator_) ProcessSymbol(
	symbol string,
) {
	v.stack_.AddValue(pri.SymbolClass().SymbolFromSource(symbol))
}

func (v *inflator_) ProcessTag(
	tag string,
) {
	v.stack_.AddValue(pri.TagClass().TagFromSource(tag))
}

func (v *inflator_) ProcessVersion(
	version string,
) {
	v.stack_.AddValue(pri.VersionClass().VersionFromSource(version))
}

func (v *inflator_) PostprocessAcceptClause(
	acceptClause not.AcceptClauseLike,
	index_ uint,
	count_ uint,
) {
	var bag = v.stack_.RemoveLast().(doc.ExpressionLike)
	var message = v.stack_.RemoveLast().(doc.ExpressionLike)
	v.stack_.AddValue(doc.AcceptClauseClass().AcceptClause(message, bag))
}

func (v *inflator_) PostprocessAssignClause(
	assignClause not.AssignClauseLike,
	index_ uint,
	count_ uint,
) {
	var expression = v.stack_.RemoveLast().(doc.ExpressionLike)
	var assignment = v.stack_.RemoveLast().(doc.Assignment)
	var recipient = v.stack_.RemoveLast()
	v.stack_.AddValue(
		doc.AssignClauseClass().AssignClause(recipient, assignment, expression),
	)
}

func (v *inflator_) PostprocessAssignment(
	assignment not.AssignmentLike,
	index_ uint,
	count_ uint,
) {
	var operation = assignment.GetAny().(string)
	switch operation {
	case "?=":
		v.stack_.AddValue(doc.Default)
	case ":=":
		v.stack_.AddValue(doc.Assign)
	case "+=":
		v.stack_.AddValue(doc.Add)
	case "-=":
		v.stack_.AddValue(doc.Subtract)
	case "*=":
		v.stack_.AddValue(doc.Multiply)
	case "/=":
		v.stack_.AddValue(doc.Divide)
	case "&=":
		v.stack_.AddValue(doc.Join)
	default:
		var message = fmt.Sprintf(
			"Found an unexpected string value in a switch statement: %v",
			operation,
		)
		panic(message)
	}
}

func (v *inflator_) PostprocessAttributes(
	attributes not.AttributesLike,
	index_ uint,
	count_ uint,
) {
	var components = com.Catalog[any, doc.ComponentLike]()
	var associations = attributes.GetAssociations().GetIterator()
	for associations.HasNext() {
		var component = v.stack_.RemoveLast().(doc.ComponentLike)
		var primitive = v.stack_.RemoveLast()
		components.SetValue(primitive, component)
		associations.GetNext()
	}
	components.ReverseValues() // They were pulled off the stack in reverse order.
	v.stack_.AddValue(doc.AttributesClass().Attributes(components))
}

func (v *inflator_) PostprocessBreakClause(
	breakClause not.BreakClauseLike,
	index_ uint,
	count_ uint,
) {
	v.stack_.AddValue(doc.BreakClauseClass().BreakClause())
}

func (v *inflator_) PostprocessCheckoutClause(
	checkoutClause not.CheckoutClauseLike,
	index_ uint,
	count_ uint,
) {
	var location = v.stack_.RemoveLast().(doc.ExpressionLike)
	var atLevel doc.ExpressionLike
	if uti.IsDefined(checkoutClause.GetOptionalAtLevel()) {
		atLevel = v.stack_.RemoveLast().(doc.ExpressionLike)
	}
	var recipient = v.stack_.RemoveLast()
	v.stack_.AddValue(
		doc.CheckoutClauseClass().CheckoutClause(recipient, atLevel, location),
	)
}

func (v *inflator_) PostprocessComponent(
	component not.ComponentLike,
	index_ uint,
	count_ uint,
) {
	var note string
	if uti.IsDefined(component.GetOptionalNote()) {
		note = v.stack_.RemoveLast().(string)
	}
	var parameters doc.ParametersLike
	if uti.IsDefined(component.GetOptionalParameters()) {
		parameters = v.stack_.RemoveLast().(doc.ParametersLike)
	}
	var literal = v.stack_.RemoveLast()
	if uti.IsDefined(parameters) {
		switch actual := literal.(type) {
		case doc.ItemsLike:
			var components = actual.GetComponents()
			var constraints = parameters.GetConstraints()
			var component = constraints.GetValue(
				pri.SymbolClass().SymbolFromSource("$type"),
			)
			if uti.IsDefined(component) {
				switch actual := component.GetLiteral().(type) {
				case pri.NameLike:
					switch actual.AsSource() {
					case "/types/collections/Queue/v3":
						components = com.QueueFromSequence(components)
					case "/types/collections/Set/v3":
						components = com.SetFromSequence(components)
					case "/types/collections/Stack/v3":
						components = com.StackFromSequence(components)
					}
					literal = doc.ItemsClass().Items(components)
				}
			}
		}
	}
	v.stack_.AddValue(doc.ComponentClass().Component(literal, parameters, note))
}

func (v *inflator_) PostprocessContinueClause(
	continueClause not.ContinueClauseLike,
	index_ uint,
	count_ uint,
) {
	v.stack_.AddValue(doc.ContinueClauseClass().ContinueClause())
}

func (v *inflator_) PostprocessDefineClause(
	defineClause not.DefineClauseLike,
	index_ uint,
	count_ uint,
) {
	var expression = v.stack_.RemoveLast().(doc.ExpressionLike)
	var constant = v.stack_.RemoveLast().(pri.SymbolLike)
	v.stack_.AddValue(
		doc.DefineClauseClass().DefineClause(constant, expression),
	)
}

func (v *inflator_) PostprocessDiscardClause(
	discardClause not.DiscardClauseLike,
	index_ uint,
	count_ uint,
) {
	var location = v.stack_.RemoveLast().(doc.ExpressionLike)
	v.stack_.AddValue(doc.DiscardClauseClass().DiscardClause(location))
}

func (v *inflator_) PostprocessDocument(
	document not.DocumentLike,
	index_ uint,
	count_ uint,
) {
	var component = v.stack_.RemoveLast().(doc.ComponentLike)
	var comment string
	if uti.IsDefined(document.GetOptionalComment()) {
		comment = v.stack_.RemoveLast().(string)
	}
	v.stack_.AddValue(doc.DocumentClass().Document(comment, component))
	if v.stack_.GetSize() != 1 {
		var message = fmt.Sprintf(
			"Internal Error: the inflator stack is corrupted: %v",
			v.stack_,
		)
		panic(message)
	}
}

func (v *inflator_) PostprocessEmpty(
	empty not.EmptyLike,
	index_ uint,
	count_ uint,
) {
	if uti.IsDefined(empty.GetOptionalDelimiter()) {
		var catalog = com.Catalog[any, doc.ComponentLike]()
		v.stack_.AddValue(doc.AttributesClass().Attributes(catalog))
	} else {
		var list = com.List[doc.ComponentLike]()
		v.stack_.AddValue(doc.ItemsClass().Items(list))
	}
}

func (v *inflator_) PostprocessExpression(
	expression not.ExpressionLike,
	index_ uint,
	count_ uint,
) {
	var predicate doc.PredicateLike
	if uti.IsDefined(expression.GetOptionalPredicate()) {
		predicate = v.stack_.RemoveLast().(doc.PredicateLike)
	}
	var subject = v.stack_.RemoveLast()
	v.stack_.AddValue(doc.ExpressionClass().Expression(subject, predicate))
}

func (v *inflator_) PostprocessFunction(
	function not.FunctionLike,
	index_ uint,
	count_ uint,
) {
	var list = com.List[any]()
	var arguments = function.GetArguments()
	var iterator = arguments.GetIterator()
	for iterator.HasNext() {
		var argument = v.stack_.RemoveLast()
		list.AppendValue(argument)
		iterator.GetNext()
	}
	list.ReverseValues() // They were pulled off the stack in reverse order.
	var identifier = v.stack_.RemoveLast().(pri.IdentifierLike)
	v.stack_.AddValue(doc.FunctionClass().Function(identifier, list))
}

func (v *inflator_) PostprocessIfClause(
	ifClause not.IfClauseLike,
	index_ uint,
	count_ uint,
) {
	var procedure = v.stack_.RemoveLast().(doc.ProcedureLike)
	var condition = v.stack_.RemoveLast().(doc.ExpressionLike)
	v.stack_.AddValue(doc.IfClauseClass().IfClause(condition, procedure))
}

func (v *inflator_) PostprocessInvocation(
	invocation not.InvocationLike,
	index_ uint,
	count_ uint,
) {
	var isSynchronous bool
	var operation = invocation.GetAny().(string)
	switch operation {
	case "<-":
		isSynchronous = true
	case "<~":
		isSynchronous = false
	default:
		var message = fmt.Sprintf(
			"Found an unexpected string value in a switch statement: %v",
			operation,
		)
		panic(message)
	}
	v.stack_.AddValue(isSynchronous)
}

func (v *inflator_) PostprocessInvokeClause(
	invokeClause not.InvokeClauseLike,
	index_ uint,
	count_ uint,
) {
	var method = v.stack_.RemoveLast().(doc.MethodLike)
	v.stack_.AddValue(doc.InvokeClauseClass().InvokeClause(method))
}

func (v *inflator_) PostprocessItems(
	items not.ItemsLike,
	index_ uint,
	count_ uint,
) {
	var components = com.List[doc.ComponentLike]()
	var iterator = items.GetComponents().GetIterator()
	for iterator.HasNext() {
		var component = v.stack_.RemoveLast().(doc.ComponentLike)
		components.AppendValue(component)
		iterator.GetNext()
	}
	components.ReverseValues() // They were pulled off the stack in reverse order.
	v.stack_.AddValue(doc.ItemsClass().Items(components))
}

func (v *inflator_) PostprocessLeft(
	left not.LeftLike,
	index_ uint,
	count_ uint,
) {
	var bracket = left.GetAny().(string)
	switch bracket {
	case "[":
		v.stack_.AddValue(com.Inclusive)
	case "(":
		v.stack_.AddValue(com.Exclusive)
	default:
		var message = fmt.Sprintf(
			"Found an unexpected string value in a switch statement: %v",
			bracket,
		)
		panic(message)
	}
}

func (v *inflator_) PostprocessMagnitude(
	magnitude not.MagnitudeLike,
	index_ uint,
	count_ uint,
) {
	var expression = v.stack_.RemoveLast().(doc.ExpressionLike)
	v.stack_.AddValue(doc.MagnitudeClass().Magnitude(expression))
}

func (v *inflator_) PostprocessMatchingClause(
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

func (v *inflator_) PostprocessMethod(
	method not.MethodLike,
	index_ uint,
	count_ uint,
) {
	var list = com.List[any]()
	var arguments = method.GetArguments()
	var iterator = arguments.GetIterator()
	for iterator.HasNext() {
		var argument = v.stack_.RemoveLast()
		list.AppendValue(argument)
		iterator.GetNext()
	}
	list.ReverseValues() // They were pulled off the stack in reverse order.
	var identifier = v.stack_.RemoveLast().(pri.IdentifierLike)
	var isSynchronous = v.stack_.RemoveLast().(bool)
	var target = v.stack_.RemoveLast().(pri.IdentifierLike)
	v.stack_.AddValue(
		doc.MethodClass().Method(target, identifier, list, isSynchronous),
	)
}

func (v *inflator_) PostprocessModifier(
	modifier not.ModifierLike,
	index_ uint,
	count_ uint,
) {
	var operation = modifier.GetAny().(string)
	switch operation {
	case "not":
		v.stack_.AddValue(doc.Complement)
	case "-":
		v.stack_.AddValue(doc.Additive)
	case "/":
		v.stack_.AddValue(doc.Multiplicative)
	case "*":
		v.stack_.AddValue(doc.Conjugate)
	case "@":
		v.stack_.AddValue(doc.Referent)
	default:
		var message = fmt.Sprintf(
			"Found an unexpected string value in a switch statement: %v",
			operation,
		)
		panic(message)
	}
}

func (v *inflator_) PostprocessNotarizeClause(
	notarizeClause not.NotarizeClauseLike,
	index_ uint,
	count_ uint,
) {
	var location = v.stack_.RemoveLast().(doc.ExpressionLike)
	var draft = v.stack_.RemoveLast().(doc.ExpressionLike)
	v.stack_.AddValue(
		doc.NotarizeClauseClass().NotarizeClause(draft, location),
	)
}

func (v *inflator_) PostprocessOnClause(
	onClause not.OnClauseLike,
	index_ uint,
	count_ uint,
) {
	var list = com.List[doc.MatchingClauseLike]()
	var matchingClauses = onClause.GetMatchingClauses()
	var iterator = matchingClauses.GetIterator()
	for iterator.HasNext() {
		var matchingClause = v.stack_.RemoveLast().(doc.MatchingClauseLike)
		list.AppendValue(matchingClause)
		iterator.GetNext()
	}
	list.ReverseValues() // They were pulled off the stack in reverse order.
	var symbol = v.stack_.RemoveLast().(pri.SymbolLike)
	v.stack_.AddValue(
		doc.OnClauseClass().OnClause(symbol, list),
	)
}

func (v *inflator_) PostprocessOperation(
	operation not.OperationLike,
	index_ uint,
	count_ uint,
) {
	var operator = operation.GetAny().(string)
	switch operator {
	case "&":
		v.stack_.AddValue(doc.Chain)
	case "+":
		v.stack_.AddValue(doc.Sum)
	case "-":
		v.stack_.AddValue(doc.Difference)
	case "*":
		v.stack_.AddValue(doc.Product)
	case "/":
		v.stack_.AddValue(doc.Quotient)
	case "%":
		v.stack_.AddValue(doc.Remainder)
	case "^":
		v.stack_.AddValue(doc.Power)
	case "<":
		v.stack_.AddValue(doc.Less)
	case "=":
		v.stack_.AddValue(doc.Equal)
	case ">":
		v.stack_.AddValue(doc.More)
	case "is":
		v.stack_.AddValue(doc.Is)
	case "matches":
		v.stack_.AddValue(doc.Matches)
	case "and":
		v.stack_.AddValue(doc.And)
	case "san":
		v.stack_.AddValue(doc.San)
	case "ior":
		v.stack_.AddValue(doc.Ior)
	case "xor":
		v.stack_.AddValue(doc.Xor)
	default:
		var message = fmt.Sprintf(
			"Found an unexpected string value in a switch statement: %v",
			operator,
		)
		panic(message)
	}
}

func (v *inflator_) PostprocessParameters(
	parameters not.ParametersLike,
	index_ uint,
	count_ uint,
) {
	var catalog = com.Catalog[pri.SymbolLike, doc.ComponentLike]()
	var constraints = parameters.GetConstraints()
	var iterator = constraints.GetIterator()
	for iterator.HasNext() {
		var component = v.stack_.RemoveLast().(doc.ComponentLike)
		var symbol = v.stack_.RemoveLast().(pri.SymbolLike)
		catalog.SetValue(symbol, component)
		iterator.GetNext()
	}
	catalog.ReverseValues() // They were pulled off the stack in reverse order.
	v.stack_.AddValue(doc.ParametersClass().Parameters(catalog))
}

func (v *inflator_) PostprocessPrecedence(
	precedence not.PrecedenceLike,
	index_ uint,
	count_ uint,
) {
	var expression = v.stack_.RemoveLast().(doc.ExpressionLike)
	v.stack_.AddValue(doc.PrecedenceClass().Precedence(expression))
}

func (v *inflator_) PostprocessPredicate(
	predicate not.PredicateLike,
	index_ uint,
	count_ uint,
) {
	var expression = v.stack_.RemoveLast().(doc.ExpressionLike)
	var operation = v.stack_.RemoveLast().(doc.Operation)
	v.stack_.AddValue(doc.PredicateClass().Predicate(operation, expression))
}

func (v *inflator_) PostprocessProcedure(
	procedure not.ProcedureLike,
	index_ uint,
	count_ uint,
) {
	var statements = com.List[doc.StatementLike]()
	var iterator = procedure.GetStatements().GetIterator()
	for iterator.HasNext() {
		var statement = v.stack_.RemoveLast().(doc.StatementLike)
		statements.AppendValue(statement)
		iterator.GetNext()
	}
	statements.ReverseValues() // They were pulled off the stack in reverse order.
	v.stack_.AddValue(doc.ProcedureClass().Procedure(statements))
}

func (v *inflator_) PostprocessPublishClause(
	publishClause not.PublishClauseLike,
	index_ uint,
	count_ uint,
) {
	var event = v.stack_.RemoveLast().(doc.ExpressionLike)
	v.stack_.AddValue(doc.PublishClauseClass().PublishClause(event))
}

func (v *inflator_) PostprocessRange(
	range_ not.RangeLike,
	index_ uint,
	count_ uint,
) {
	var right = v.stack_.RemoveLast().(com.Bracket)
	var last any
	if uti.IsDefined(range_.GetOptionalPrimitive2()) {
		last = v.stack_.RemoveLast()
	}
	var first any
	if uti.IsDefined(range_.GetOptionalPrimitive1()) {
		first = v.stack_.RemoveLast()
	}
	var left = v.stack_.RemoveLast().(com.Bracket)
	v.stack_.AddValue(doc.RangeClass().Range(left, first, last, right))
}

func (v *inflator_) PostprocessReceiveClause(
	receiveClause not.ReceiveClauseLike,
	index_ uint,
	count_ uint,
) {
	var bag = v.stack_.RemoveLast().(doc.ExpressionLike)
	var recipient = v.stack_.RemoveLast()
	v.stack_.AddValue(doc.ReceiveClauseClass().ReceiveClause(recipient, bag))
}

func (v *inflator_) PostprocessRefinement(
	refinement not.RefinementLike,
	index_ uint,
	count_ uint,
) {
	var subject = v.stack_.RemoveLast()
	var modifier = v.stack_.RemoveLast().(doc.Modifier)
	v.stack_.AddValue(doc.RefinementClass().Refinement(modifier, subject))
}

func (v *inflator_) PostprocessRejectClause(
	rejectClause not.RejectClauseLike,
	index_ uint,
	count_ uint,
) {
	var bag = v.stack_.RemoveLast().(doc.ExpressionLike)
	var message = v.stack_.RemoveLast().(doc.ExpressionLike)
	v.stack_.AddValue(doc.RejectClauseClass().RejectClause(message, bag))
}

func (v *inflator_) PostprocessReturnClause(
	returnClause not.ReturnClauseLike,
	index_ uint,
	count_ uint,
) {
	var result = v.stack_.RemoveLast().(doc.ExpressionLike)
	v.stack_.AddValue(doc.ReturnClauseClass().ReturnClause(result))
}

func (v *inflator_) PostprocessRight(
	right not.RightLike,
	index_ uint,
	count_ uint,
) {
	var bracket = right.GetAny().(string)
	switch bracket {
	case "]":
		v.stack_.AddValue(com.Inclusive)
	case ")":
		v.stack_.AddValue(com.Exclusive)
	default:
		var message = fmt.Sprintf(
			"Found an unexpected string value in a switch statement: %v",
			bracket,
		)
		panic(message)
	}
}

func (v *inflator_) PostprocessSaveClause(
	saveClause not.SaveClauseLike,
	index_ uint,
	count_ uint,
) {
	var recipient = v.stack_.RemoveLast()
	var draft = v.stack_.RemoveLast().(doc.ExpressionLike)
	v.stack_.AddValue(doc.SaveClauseClass().SaveClause(draft, recipient))
}

func (v *inflator_) PostprocessSelectClause(
	selectClause not.SelectClauseLike,
	index_ uint,
	count_ uint,
) {
	var list = com.List[doc.MatchingClauseLike]()
	var matchingClauses = selectClause.GetMatchingClauses()
	var iterator = matchingClauses.GetIterator()
	for iterator.HasNext() {
		var matchingClause = v.stack_.RemoveLast().(doc.MatchingClauseLike)
		list.AppendValue(matchingClause)
		iterator.GetNext()
	}
	list.ReverseValues() // They were pulled off the stack in reverse order.
	var expression = v.stack_.RemoveLast().(doc.ExpressionLike)
	v.stack_.AddValue(doc.SelectClauseClass().SelectClause(expression, list))
}

func (v *inflator_) PostprocessSendClause(
	sendClause not.SendClauseLike,
	index_ uint,
	count_ uint,
) {
	var bag = v.stack_.RemoveLast().(doc.ExpressionLike)
	var message = v.stack_.RemoveLast().(doc.ExpressionLike)
	v.stack_.AddValue(doc.SendClauseClass().SendClause(message, bag))
}

func (v *inflator_) PostprocessStatement(
	statement not.StatementLike,
	index_ uint,
	count_ uint,
) {
	var onClause doc.OnClauseLike
	if uti.IsDefined(statement.GetOptionalOnClause()) {
		onClause = v.stack_.RemoveLast().(doc.OnClauseLike)
	}
	var mainClause = v.stack_.RemoveLast()
	var comment string
	if uti.IsDefined(statement.GetOptionalComment()) {
		comment = v.stack_.RemoveLast().(string)
	}
	v.stack_.AddValue(doc.StatementClass().Statement(comment, mainClause, onClause))
}

func (v *inflator_) PostprocessSubcomponent(
	subcomponent not.SubcomponentLike,
	index_ uint,
	count_ uint,
) {
	var list = com.List[any]()
	var indexes = subcomponent.GetIndexes()
	var iterator = indexes.GetIterator()
	for iterator.HasNext() {
		var index = v.stack_.RemoveLast()
		list.AppendValue(index)
		iterator.GetNext()
	}
	list.ReverseValues() // They were pulled off the stack in reverse order.
	var identifier = v.stack_.RemoveLast().(pri.IdentifierLike)
	v.stack_.AddValue(doc.SubcomponentClass().Subcomponent(identifier, list))
}

func (v *inflator_) PostprocessThrowClause(
	throwClause not.ThrowClauseLike,
	index_ uint,
	count_ uint,
) {
	var exception = v.stack_.RemoveLast().(doc.ExpressionLike)
	v.stack_.AddValue(doc.ThrowClauseClass().ThrowClause(exception))
}

func (v *inflator_) PostprocessWhileClause(
	whileClause not.WhileClauseLike,
	index_ uint,
	count_ uint,
) {
	var procedure = v.stack_.RemoveLast().(doc.ProcedureLike)
	var condition = v.stack_.RemoveLast().(doc.ExpressionLike)
	v.stack_.AddValue(doc.WhileClauseClass().WhileClause(condition, procedure))
}

func (v *inflator_) PostprocessWithClause(
	withClause not.WithClauseLike,
	index_ uint,
	count_ uint,
) {
	var procedure = v.stack_.RemoveLast().(doc.ProcedureLike)
	var sequence = v.stack_.RemoveLast().(doc.ExpressionLike)
	var variable = v.stack_.RemoveLast().(pri.SymbolLike)
	v.stack_.AddValue(
		doc.WithClauseClass().WithClause(variable, sequence, procedure),
	)
}

// PROTECTED INTERFACE

// Private Methods

// Instance Structure

type inflator_ struct {
	// Declare the instance attributes.
	stack_ com.StackLike[any]

	// Declare the inherited aspects.
	not.Methodical
}

// Class Structure

type inflatorClass_ struct {
	// Declare the class constants.
}

// Class Reference

func inflatorClass() *inflatorClass_ {
	return inflatorClassReference_
}

var inflatorClassReference_ = &inflatorClass_{
	// Initialize the class constants.
}
