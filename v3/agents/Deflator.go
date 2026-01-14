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

func DeflatorClass() DeflatorClassLike {
	return deflatorClass()
}

// Constructor Methods

func (c *deflatorClass_) Deflator() DeflatorLike {
	var instance = &deflator_{
		// Initialize the instance attributes.
		stack_: com.StackWithCapacity[any](256),

		// Initialize the inherited aspects.
		Methodical: ProcessorClass().Processor(),
	}
	return instance
}

// Constant Methods

// Function Methods

// INSTANCE INTERFACE

// Principal Methods

func (v *deflator_) GetClass() DeflatorClassLike {
	return deflatorClass()
}

func (v *deflator_) DeflateDocument(
	document doc.DocumentLike,
) not.DocumentLike {
	VisitorClass().Visitor(v).VisitDocument(document)
	return v.stack_.RemoveLast().(not.DocumentLike)
}

// Attribute Methods

// Methodical Methods

func (v *deflator_) ProcessAngle(
	angle pri.AngleLike,
) {
	v.stack_.AddValue(not.Element(angle.AsSource()))
}

func (v *deflator_) ProcessAssignment(
	assignment doc.Assignment,
) {
	var operation string
	switch assignment {
	case doc.Default:
		operation = "?="
	case doc.Assign:
		operation = ":="
	case doc.Add:
		operation = "+="
	case doc.Subtract:
		operation = "-="
	case doc.Multiply:
		operation = "*="
	case doc.Divide:
		operation = "/="
	case doc.Join:
		operation = "&="
	default:
		var message = fmt.Sprintf(
			"Found an unexpected value in a switch statement: %v(%T)",
			assignment,
			assignment,
		)
		panic(message)
	}
	v.stack_.AddValue(not.Assignment(operation))
}

func (v *deflator_) ProcessBinary(
	binary pri.BinaryLike,
) {
	v.stack_.AddValue(not.Sequence(binary.AsSource()))
}

func (v *deflator_) ProcessBoolean(
	boolean pri.BooleanLike,
) {
	v.stack_.AddValue(not.Element(boolean.AsSource()))
}

func (v *deflator_) ProcessBytecode(
	bytecode pri.BytecodeLike,
) {
	v.stack_.AddValue(not.Sequence(bytecode.AsSource()))
}

func (v *deflator_) ProcessComment(
	comment string,
) {
	v.stack_.AddValue(comment)
}

func (v *deflator_) ProcessDuration(
	duration pri.DurationLike,
) {
	v.stack_.AddValue(not.Element(duration.AsSource()))
}

func (v *deflator_) ProcessGlyph(
	glyph pri.GlyphLike,
) {
	v.stack_.AddValue(not.Element(glyph.AsSource()))
}

func (v *deflator_) ProcessIdentifier(
	identifier pri.IdentifierLike,
) {
	v.stack_.AddValue(identifier.AsSource())
}

func (v *deflator_) ProcessIsSynchronous(
	isSynchronous bool,
) {
	var operation string
	if isSynchronous {
		operation = "<-"
	} else {
		operation = "<~"
	}
	v.stack_.AddValue(not.Invocation(operation))
}

func (v *deflator_) ProcessModifier(
	modifier doc.Modifier,
) {
	var operation string
	switch modifier {
	case doc.Complement:
		operation = "not"
	case doc.Additive:
		operation = "-"
	case doc.Multiplicative:
		operation = "/"
	case doc.Conjugate:
		operation = "*"
	case doc.Referent:
		operation = "@"
	default:
		var message = fmt.Sprintf(
			"Found an unexpected value in a switch statement: %v(%T)",
			modifier,
			modifier,
		)
		panic(message)
	}
	v.stack_.AddValue(not.Modifier(operation))
}

func (v *deflator_) ProcessMoment(
	moment pri.MomentLike,
) {
	v.stack_.AddValue(not.Element(moment.AsSource()))
}

func (v *deflator_) ProcessName(
	name pri.NameLike,
) {
	v.stack_.AddValue(not.Sequence(name.AsSource()))
}

func (v *deflator_) ProcessNarrative(
	narrative pri.NarrativeLike,
) {
	v.stack_.AddValue(not.Sequence(narrative.AsSource()))
}

func (v *deflator_) ProcessNote(
	note string,
) {
	v.stack_.AddValue(note)
}

func (v *deflator_) ProcessNumber(
	number pri.NumberLike,
) {
	v.stack_.AddValue(not.Element(number.AsSource()))
}

func (v *deflator_) ProcessOperation(
	operation doc.Operation,
) {
	switch operation {
	case doc.Chain:
		v.stack_.AddValue(not.Operation("&"))
	case doc.Sum:
		v.stack_.AddValue(not.Operation("+"))
	case doc.Difference:
		v.stack_.AddValue(not.Operation("-"))
	case doc.Product:
		v.stack_.AddValue(not.Operation("*"))
	case doc.Quotient:
		v.stack_.AddValue(not.Operation("/"))
	case doc.Remainder:
		v.stack_.AddValue(not.Operation("%"))
	case doc.Power:
		v.stack_.AddValue(not.Operation("^"))
	case doc.Less:
		v.stack_.AddValue(not.Operation("<"))
	case doc.Equal:
		v.stack_.AddValue(not.Operation("="))
	case doc.More:
		v.stack_.AddValue(not.Operation(">"))
	case doc.Is:
		v.stack_.AddValue(not.Operation("is"))
	case doc.Matches:
		v.stack_.AddValue(not.Operation("matches"))
	case doc.And:
		v.stack_.AddValue(not.Operation("and"))
	case doc.San:
		v.stack_.AddValue(not.Operation("san"))
	case doc.Ior:
		v.stack_.AddValue(not.Operation("ior"))
	case doc.Xor:
		v.stack_.AddValue(not.Operation("xor"))
	default:
		var message = fmt.Sprintf(
			"Found an unexpected value in a switch statement: %v(%T)",
			operation,
			operation,
		)
		panic(message)
	}
}

func (v *deflator_) ProcessPattern(
	pattern pri.PatternLike,
) {
	v.stack_.AddValue(not.Sequence(pattern.AsSource()))
}

func (v *deflator_) ProcessPercentage(
	percentage pri.PercentageLike,
) {
	v.stack_.AddValue(not.Element(percentage.AsSource()))
}

func (v *deflator_) ProcessProbability(
	probability pri.ProbabilityLike,
) {
	v.stack_.AddValue(not.Element(probability.AsSource()))
}

func (v *deflator_) ProcessQuote(
	quote pri.QuoteLike,
) {
	v.stack_.AddValue(not.Sequence(quote.AsSource()))
}

func (v *deflator_) ProcessResource(
	resource pri.ResourceLike,
) {
	v.stack_.AddValue(not.Element(resource.AsSource()))
}

func (v *deflator_) ProcessSymbol(
	symbol pri.SymbolLike,
) {
	v.stack_.AddValue(not.Sequence(symbol.AsSource()))
}

func (v *deflator_) ProcessTag(
	tag pri.TagLike,
) {
	v.stack_.AddValue(not.Sequence(tag.AsSource()))
}

func (v *deflator_) ProcessVersion(
	version pri.VersionLike,
) {
	v.stack_.AddValue(not.Sequence(version.AsSource()))
}

func (v *deflator_) PostprocessAcceptClause(
	acceptClause doc.AcceptClauseLike,
	index_ uint,
	count_ uint,
) {
	var expression = v.stack_.RemoveLast().(not.ExpressionLike)
	var bag = not.Bag(expression)
	expression = v.stack_.RemoveLast().(not.ExpressionLike)
	var message = not.Message(expression)
	v.stack_.AddValue(
		not.MessageHandling(
			not.AcceptClause("accept", message, "from", bag),
		),
	)
}

func (v *deflator_) PostprocessAssignClause(
	assignClause doc.AssignClauseLike,
	index_ uint,
	count_ uint,
) {
	var expression = v.stack_.RemoveLast().(not.ExpressionLike)
	var assignment = v.stack_.RemoveLast().(not.AssignmentLike)
	var recipient not.RecipientLike
	switch actual := v.stack_.RemoveLast().(type) {
	case not.SequenceLike:
		recipient = not.Recipient(not.Variable(actual.GetAny().(string)))
	case not.SubcomponentLike:
		recipient = not.Recipient(actual)
	}
	v.stack_.AddValue(
		not.LocalTransformation(
			not.AssignClause("assign", recipient, assignment, expression),
		),
	)
}

func (v *deflator_) PostprocessAttributes(
	attributes doc.AttributesLike,
	index_ uint,
	count_ uint,
) {
	var collection not.CollectionLike
	if attributes.GetAssociations().IsEmpty() {
		collection = not.Collection(not.Empty("[", ":", "]"))
	} else {
		var associations = com.List[not.AssociationLike]()
		var iterator = attributes.GetAssociations().GetIterator()
		for iterator.HasNext() {
			var component = v.stack_.RemoveLast().(not.ComponentLike)
			var primitive = not.Primitive(v.stack_.RemoveLast())
			var association = not.Association(primitive, ":", component)
			associations.AppendValue(association)
			iterator.GetNext()
		}
		associations.ReverseValues() // Pulled off the stack in reverse order.
		collection = not.Collection(not.Attributes("[", associations, "]"))
	}
	v.stack_.AddValue(collection)
}

func (v *deflator_) PostprocessBreakClause(
	breakClause doc.BreakClauseLike,
	index_ uint,
	count_ uint,
) {
	v.stack_.AddValue(
		not.FlowControl(
			not.BreakClause("break", "loop"),
		),
	)
}

func (v *deflator_) PostprocessCheckoutClause(
	checkoutClause doc.CheckoutClauseLike,
	index_ uint,
	count_ uint,
) {
	var expression = v.stack_.RemoveLast().(not.ExpressionLike)
	var location = not.Location(expression)
	var atLevel not.AtLevelLike
	if uti.IsDefined(checkoutClause.GetOptionalAtLevel()) {
		expression = v.stack_.RemoveLast().(not.ExpressionLike)
		atLevel = not.AtLevel("at", "level", expression)
	}
	var recipient not.RecipientLike
	switch actual := v.stack_.RemoveLast().(type) {
	case not.SequenceLike:
		recipient = not.Recipient(not.Variable(actual.GetAny().(string)))
	case not.SubcomponentLike:
		recipient = not.Recipient(actual)
	}
	v.stack_.AddValue(
		not.RepositoryAccess(
			not.CheckoutClause("checkout", recipient, atLevel, "from", location),
		),
	)
}

func (v *deflator_) PostprocessComponent(
	component doc.ComponentLike,
	index_ uint,
	count_ uint,
) {
	var note string
	if uti.IsDefined(component.GetOptionalNote()) {
		note = v.stack_.RemoveLast().(string)
	}
	var parameters not.ParametersLike
	if uti.IsDefined(component.GetOptionalParameters()) {
		parameters = v.stack_.RemoveLast().(not.ParametersLike)
	}
	var literal = not.Literal(v.stack_.RemoveLast())
	v.stack_.AddValue(
		not.Component(
			literal,
			parameters,
			note,
		),
	)
}

func (v *deflator_) PostprocessContinueClause(
	continueClause doc.ContinueClauseLike,
	index_ uint,
	count_ uint,
) {
	v.stack_.AddValue(
		not.FlowControl(
			not.ContinueClause("continue", "loop"),
		),
	)
}

func (v *deflator_) PostprocessDefineClause(
	assignClause doc.DefineClauseLike,
	index_ uint,
	count_ uint,
) {
	var expression = v.stack_.RemoveLast().(not.ExpressionLike)
	var sequence = v.stack_.RemoveLast().(not.SequenceLike)
	var symbol = sequence.GetAny().(string)
	var constant = not.Constant(symbol)
	v.stack_.AddValue(
		not.LocalTransformation(
			not.DefineClause("define", constant, "as", expression),
		),
	)
}

func (v *deflator_) PostprocessDiscardClause(
	discardClause doc.DiscardClauseLike,
	index_ uint,
	count_ uint,
) {
	var expression = v.stack_.RemoveLast().(not.ExpressionLike)
	var citation = not.Citation(expression)
	v.stack_.AddValue(
		not.RepositoryAccess(
			not.DiscardClause("discard", citation),
		),
	)
}

func (v *deflator_) PostprocessDocument(
	document doc.DocumentLike,
	index_ uint,
	count_ uint,
) {
	var component = v.stack_.RemoveLast().(not.ComponentLike)
	var comment string
	if uti.IsDefined(document.GetOptionalComment()) {
		comment = v.stack_.RemoveLast().(string)
	}
	v.stack_.AddValue(
		not.Document(
			comment,
			component,
		),
	)
	if v.stack_.GetSize() != 1 {
		var message = fmt.Sprintf(
			"Internal Error: the deflator stack is corrupted: %v",
			v.stack_,
		)
		panic(message)
	}
}

func (v *deflator_) PostprocessExpression(
	expression doc.ExpressionLike,
	index_ uint,
	count_ uint,
) {
	var predicate not.PredicateLike
	if uti.IsDefined(expression.GetOptionalPredicate()) {
		predicate = v.stack_.RemoveLast().(not.PredicateLike)
	}
	var subject not.SubjectLike
	switch actual := v.stack_.RemoveLast().(type) {
	case string:
		subject = not.Subject(not.Value(actual))
	default:
		subject = not.Subject(actual)
	}
	v.stack_.AddValue(not.Expression(subject, predicate))
}

func (v *deflator_) PostprocessFunction(
	function doc.FunctionLike,
	index_ uint,
	count_ uint,
) {
	var arguments = com.List[not.ArgumentLike]()
	var iterator = function.GetArguments().GetIterator()
	for iterator.HasNext() {
		var argument not.ArgumentLike
		switch actual := v.stack_.RemoveLast().(type) {
		case string:
			argument = not.Argument(not.Value(actual))
		default:
			argument = not.Argument(actual.(not.ComponentLike))
		}
		arguments.AppendValue(argument)
		iterator.GetNext()
	}
	arguments.ReverseValues() // They were pulled off the stack in reverse order.
	var identifier = v.stack_.RemoveLast().(string)
	v.stack_.AddValue(not.Function(identifier, "(", arguments, ")"))
}

func (v *deflator_) PostprocessParameters(
	parameters doc.ParametersLike,
	index_ uint,
	count_ uint,
) {
	var constraints = com.List[not.ConstraintLike]()
	var iterator = parameters.GetConstraints().GetIterator()
	for iterator.HasNext() {
		var component = v.stack_.RemoveLast().(not.ComponentLike)
		var sequence = v.stack_.RemoveLast().(not.SequenceLike)
		var symbol = sequence.GetAny().(string)
		var constraint = not.Constraint(symbol, ":", component)
		constraints.AppendValue(constraint)
		iterator.GetNext()
	}
	constraints.ReverseValues() // They were pulled off the stack in reverse order.
	v.stack_.AddValue(not.Parameters("(", constraints, ")"))
}

func (v *deflator_) PostprocessIfClause(
	ifClause doc.IfClauseLike,
	index_ uint,
	count_ uint,
) {
	var procedure = v.stack_.RemoveLast().(not.ProcedureLike)
	var expression = v.stack_.RemoveLast().(not.ExpressionLike)
	v.stack_.AddValue(
		not.FlowControl(
			not.IfClause("if", expression, "do", procedure),
		),
	)
}

func (v *deflator_) PostprocessInvokeClause(
	invokeClause doc.InvokeClauseLike,
	index_ uint,
	count_ uint,
) {
	var method = v.stack_.RemoveLast().(not.MethodLike)
	v.stack_.AddValue(
		not.LocalTransformation(
			not.InvokeClause("invoke", method),
		),
	)
}

func (v *deflator_) PostprocessItems(
	items doc.ItemsLike,
	index_ uint,
	count_ uint,
) {
	var collection not.CollectionLike
	if items.GetComponents().IsEmpty() {
		collection = not.Collection(not.Empty("[", "", "]"))
	} else {
		var components = com.List[not.ComponentLike]()
		var iterator = items.GetComponents().GetIterator()
		for iterator.HasNext() {
			var component = v.stack_.RemoveLast().(not.ComponentLike)
			components.AppendValue(component)
			iterator.GetNext()
		}
		components.ReverseValues() // Pulled off the stack in reverse order.
		collection = not.Collection(not.Items("[", components, "]"))
	}
	v.stack_.AddValue(collection)
}

func (v *deflator_) PostprocessMagnitude(
	magnitude doc.MagnitudeLike,
	index_ uint,
	count_ uint,
) {
	var expression = v.stack_.RemoveLast().(not.ExpressionLike)
	v.stack_.AddValue(
		not.Magnitude("|", expression, "|"),
	)
}

func (v *deflator_) PostprocessMatchingClause(
	matchingClause doc.MatchingClauseLike,
	index_ uint,
	count_ uint,
) {
	var procedure = v.stack_.RemoveLast().(not.ProcedureLike)
	var expression = v.stack_.RemoveLast().(not.ExpressionLike)
	v.stack_.AddValue(
		not.MatchingClause("matching", expression, "do", procedure),
	)
}

func (v *deflator_) PostprocessMethod(
	method doc.MethodLike,
	index_ uint,
	count_ uint,
) {
	var operation = v.stack_.RemoveLast().(not.InvocationLike)
	var arguments = com.List[not.ArgumentLike]()
	var iterator = method.GetArguments().GetIterator()
	for iterator.HasNext() {
		var argument not.ArgumentLike
		switch actual := v.stack_.RemoveLast().(type) {
		case string:
			argument = not.Argument(not.Value(actual))
		default:
			argument = not.Argument(actual.(not.ComponentLike))
		}
		arguments.AppendValue(argument)
		iterator.GetNext()
	}
	arguments.ReverseValues() // They were pulled off the stack in reverse order.
	var identifier = v.stack_.RemoveLast().(string)
	var target = v.stack_.RemoveLast().(string)
	v.stack_.AddValue(
		not.Method(
			target, operation, identifier, "(", arguments, ")",
		),
	)
}

func (v *deflator_) PostprocessNotarizeClause(
	notarizeClause doc.NotarizeClauseLike,
	index_ uint,
	count_ uint,
) {
	var expression = v.stack_.RemoveLast().(not.ExpressionLike)
	var location = not.Location(expression)
	expression = v.stack_.RemoveLast().(not.ExpressionLike)
	var draft = not.Draft(expression)
	v.stack_.AddValue(
		not.RepositoryAccess(
			not.NotarizeClause("notarize", draft, "as", location),
		),
	)
}

func (v *deflator_) PostprocessOnClause(
	onClause doc.OnClauseLike,
	index_ uint,
	count_ uint,
) {
	var matchingClauses = com.List[not.MatchingClauseLike]()
	var iterator = onClause.GetMatchingClauses().GetIterator()
	for iterator.HasNext() {
		var matchingClause = v.stack_.RemoveLast().(not.MatchingClauseLike)
		matchingClauses.AppendValue(matchingClause)
		iterator.GetNext()
	}
	matchingClauses.ReverseValues() // They were pulled off the stack in reverse order.
	var sequence = v.stack_.RemoveLast().(not.SequenceLike)
	var symbol = sequence.GetAny().(string)
	v.stack_.AddValue(
		not.OnClause("on", symbol, matchingClauses),
	)
}

func (v *deflator_) PostprocessPrecedence(
	precedence doc.PrecedenceLike,
	index_ uint,
	count_ uint,
) {
	var expression = v.stack_.RemoveLast().(not.ExpressionLike)
	v.stack_.AddValue(
		not.Precedence("(", expression, ")"),
	)
}

func (v *deflator_) PostprocessPredicate(
	predicate doc.PredicateLike,
	index_ uint,
	count_ uint,
) {
	var expression = v.stack_.RemoveLast().(not.ExpressionLike)
	var operation = v.stack_.RemoveLast().(not.OperationLike)
	v.stack_.AddValue(
		not.Predicate(operation, expression),
	)
}

func (v *deflator_) PostprocessProcedure(
	procedure doc.ProcedureLike,
	index_ uint,
	count_ uint,
) {
	var statements = com.List[not.StatementLike]()
	var iterator = procedure.GetStatements().GetIterator()
	for iterator.HasNext() {
		var statement = v.stack_.RemoveLast().(not.StatementLike)
		statements.AppendValue(statement)
		iterator.GetNext()
	}
	statements.ReverseValues() // They were pulled off the stack in reverse order.
	v.stack_.AddValue(not.Procedure("{", statements, "}"))
}

func (v *deflator_) PostprocessPublishClause(
	publishClause doc.PublishClauseLike,
	index_ uint,
	count_ uint,
) {
	var expression = v.stack_.RemoveLast().(not.ExpressionLike)
	var message = not.Message(expression)
	v.stack_.AddValue(
		not.MessageHandling(
			not.PublishClause("publish", message),
		),
	)
}

func (v *deflator_) PostprocessRange(
	range_ doc.RangeLike,
	index_ uint,
	count_ uint,
) {
	var right not.RightLike
	var extent = range_.GetRight()
	switch extent {
	case com.Inclusive:
		right = not.Right("]")
	case com.Exclusive:
		right = not.Right(")")
	default:
		var message = fmt.Sprintf(
			"Found an unexpected value in a switch statement: %v(%T)",
			extent,
			extent,
		)
		panic(message)
	}
	var primitive2 not.PrimitiveLike
	if uti.IsDefined(range_.GetOptionalLast()) {
		primitive2 = not.Primitive(v.stack_.RemoveLast())
	}
	var primitive1 not.PrimitiveLike
	if uti.IsDefined(range_.GetOptionalFirst()) {
		primitive1 = not.Primitive(v.stack_.RemoveLast())
	}
	var left not.LeftLike
	extent = range_.GetLeft()
	switch extent {
	case com.Inclusive:
		left = not.Left("[")
	case com.Exclusive:
		left = not.Left("(")
	default:
		var message = fmt.Sprintf(
			"Found an unexpected value in a switch statement: %v(%T)",
			extent,
			extent,
		)
		panic(message)
	}
	v.stack_.AddValue(
		not.Range(left, primitive1, "..", primitive2, right),
	)
}

func (v *deflator_) PostprocessReceiveClause(
	receiveClause doc.ReceiveClauseLike,
	index_ uint,
	count_ uint,
) {
	var expression = v.stack_.RemoveLast().(not.ExpressionLike)
	var bag = not.Bag(expression)
	var recipient not.RecipientLike
	switch actual := v.stack_.RemoveLast().(type) {
	case not.SequenceLike:
		recipient = not.Recipient(not.Variable(actual.GetAny().(string)))
	case not.SubcomponentLike:
		recipient = not.Recipient(actual)
	}
	v.stack_.AddValue(
		not.MessageHandling(
			not.ReceiveClause("receive", recipient, "from", bag),
		),
	)
}

func (v *deflator_) PostprocessRefinement(
	refinement doc.RefinementLike,
	index_ uint,
	count_ uint,
) {
	var subject not.SubjectLike
	switch actual := v.stack_.RemoveLast().(type) {
	case string:
		subject = not.Subject(not.Value(actual))
	default:
		subject = not.Subject(actual)
	}
	var modifier = v.stack_.RemoveLast().(not.ModifierLike)
	v.stack_.AddValue(
		not.Refinement(
			modifier,
			subject,
		),
	)
}

func (v *deflator_) PostprocessRejectClause(
	rejectClause doc.RejectClauseLike,
	index_ uint,
	count_ uint,
) {
	var expression = v.stack_.RemoveLast().(not.ExpressionLike)
	var bag = not.Bag(expression)
	expression = v.stack_.RemoveLast().(not.ExpressionLike)
	var message = not.Message(expression)
	v.stack_.AddValue(
		not.MessageHandling(
			not.RejectClause("reject", message, "from", bag),
		),
	)
}

func (v *deflator_) PostprocessReturnClause(
	returnClause doc.ReturnClauseLike,
	index_ uint,
	count_ uint,
) {
	var expression = v.stack_.RemoveLast().(not.ExpressionLike)
	v.stack_.AddValue(
		not.FlowControl(
			not.ReturnClause("return", expression),
		),
	)
}

func (v *deflator_) PostprocessSaveClause(
	saveClause doc.SaveClauseLike,
	index_ uint,
	count_ uint,
) {
	var recipient not.RecipientLike
	switch actual := v.stack_.RemoveLast().(type) {
	case not.SequenceLike:
		recipient = not.Recipient(not.Variable(actual.GetAny().(string)))
	case not.SubcomponentLike:
		recipient = not.Recipient(actual)
	}
	var expression = v.stack_.RemoveLast().(not.ExpressionLike)
	var draft = not.Draft(expression)
	v.stack_.AddValue(
		not.RepositoryAccess(
			not.SaveClause("save", draft, "to", recipient),
		),
	)
}

func (v *deflator_) PostprocessSelectClause(
	selectClause doc.SelectClauseLike,
	index_ uint,
	count_ uint,
) {
	var matchingClauses = com.List[not.MatchingClauseLike]()
	var iterator = selectClause.GetMatchingClauses().GetIterator()
	for iterator.HasNext() {
		var matchingClause = v.stack_.RemoveLast().(not.MatchingClauseLike)
		matchingClauses.AppendValue(matchingClause)
		iterator.GetNext()
	}
	matchingClauses.ReverseValues() // They were pulled off the stack in reverse order.
	var expression = v.stack_.RemoveLast().(not.ExpressionLike)
	v.stack_.AddValue(
		not.FlowControl(
			not.SelectClause("select", expression, matchingClauses),
		),
	)
}

func (v *deflator_) PostprocessSendClause(
	sendClause doc.SendClauseLike,
	index_ uint,
	count_ uint,
) {
	var expression = v.stack_.RemoveLast().(not.ExpressionLike)
	var bag = not.Bag(expression)
	expression = v.stack_.RemoveLast().(not.ExpressionLike)
	var message = not.Message(expression)
	v.stack_.AddValue(
		not.MessageHandling(
			not.SendClause("send", message, "to", bag),
		),
	)
}

func (v *deflator_) PostprocessStatement(
	statement doc.StatementLike,
	index_ uint,
	count_ uint,
) {
	var onClause not.OnClauseLike
	if uti.IsDefined(statement.GetOptionalOnClause()) {
		onClause = v.stack_.RemoveLast().(not.OnClauseLike)
	}
	var mainClause = not.MainClause(v.stack_.RemoveLast())
	var comment string
	if uti.IsDefined(statement.GetOptionalComment()) {
		comment = v.stack_.RemoveLast().(string)
	}
	v.stack_.AddValue(not.Statement(comment, mainClause, onClause))
}

func (v *deflator_) PostprocessSubcomponent(
	subcomponent doc.SubcomponentLike,
	index_ uint,
	count_ uint,
) {
	var indexes = com.List[not.IndexLike]()
	var iterator = subcomponent.GetIndexes().GetIterator()
	for iterator.HasNext() {
		var index not.IndexLike
		switch actual := v.stack_.RemoveLast().(type) {
		case string:
			index = not.Index(not.Value(actual))
		default:
			index = not.Index(not.Primitive(actual))
		}
		indexes.AppendValue(index)
		iterator.GetNext()
	}
	indexes.ReverseValues() // They were pulled off the stack in reverse order.
	var identifier = v.stack_.RemoveLast().(string)
	v.stack_.AddValue(not.Subcomponent(identifier, "[", indexes, "]"))
}

func (v *deflator_) PostprocessThrowClause(
	throwClause doc.ThrowClauseLike,
	index_ uint,
	count_ uint,
) {
	var expression = v.stack_.RemoveLast().(not.ExpressionLike)
	v.stack_.AddValue(
		not.FlowControl(
			not.ThrowClause("throw", expression),
		),
	)
}

func (v *deflator_) PostprocessWhileClause(
	whileClause doc.WhileClauseLike,
	index_ uint,
	count_ uint,
) {
	var procedure = v.stack_.RemoveLast().(not.ProcedureLike)
	var expression = v.stack_.RemoveLast().(not.ExpressionLike)
	v.stack_.AddValue(
		not.FlowControl(
			not.WhileClause("while", expression, "do", procedure),
		),
	)
}

func (v *deflator_) PostprocessWithClause(
	withClause doc.WithClauseLike,
	index_ uint,
	count_ uint,
) {
	var procedure = v.stack_.RemoveLast().(not.ProcedureLike)
	var expression = v.stack_.RemoveLast().(not.ExpressionLike)
	var sequence = v.stack_.RemoveLast().(not.SequenceLike)
	var symbol = sequence.GetAny().(string)
	v.stack_.AddValue(
		not.FlowControl(
			not.WithClause(
				"with",
				"each",
				symbol,
				"in",
				expression,
				"do",
				procedure,
			),
		),
	)
}

// PROTECTED INTERFACE

// Private Methods

// Instance Structure

type deflator_ struct {
	// Declare the instance attributes.
	stack_ com.StackLike[any]

	// Declare the inherited aspects.
	Methodical
}

// Class Structure

type deflatorClass_ struct {
	// Declare the class constants.
}

// Class Reference

func deflatorClass() *deflatorClass_ {
	return deflatorClassReference_
}

var deflatorClassReference_ = &deflatorClass_{
	// Initialize the class constants.
}
