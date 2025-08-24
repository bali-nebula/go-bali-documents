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

func DeflatorClass() DeflatorClassLike {
	return deflatorClass()
}

// Constructor Methods

func (c *deflatorClass_) Deflator() DeflatorLike {
	var instance = &deflator_{
		// Initialize the instance attributes.
		stack_: fra.StackWithCapacity[any](256),

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
	if v.stack_.GetSize() != 1 {
		var message = fmt.Sprintf(
			"Internal Error: the deflator stack is corrupted: %v",
			v.stack_,
		)
		panic(message)
	}
	return v.stack_.RemoveLast().(not.DocumentLike)
}

// Attribute Methods

// Methodical Methods

func (v *deflator_) ProcessAngle(
	angle fra.AngleLike,
) {
	v.stack_.AddValue(not.Element(angle.AsString()))
}

func (v *deflator_) ProcessAnnotation(
	annotation string,
) {
	v.stack_.AddValue(not.Line(not.Annotation(annotation)))
}

func (v *deflator_) ProcessAssignment(
	assignment doc.Assignment,
) {
	var operation string
	switch assignment {
	case doc.Equals:
		operation = ":="
	case doc.EqualsDefault:
		operation = "?="
	case doc.EqualsPlus:
		operation = "+="
	case doc.EqualsMinus:
		operation = "-="
	case doc.EqualsTimes:
		operation = "*="
	case doc.EqualsDivide:
		operation = "/="
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
	binary fra.BinaryLike,
) {
	v.stack_.AddValue(not.String(binary.AsString()))
}

func (v *deflator_) ProcessBoolean(
	boolean fra.BooleanLike,
) {
	v.stack_.AddValue(not.Element(boolean.AsString()))
}

func (v *deflator_) ProcessBytecode(
	bytecode ass.BytecodeLike,
) {
	v.stack_.AddValue(not.String(bytecode.AsString()))
}

func (v *deflator_) ProcessComment(
	comment string,
) {
	v.stack_.AddValue(comment)
}

func (v *deflator_) ProcessDuration(
	duration fra.DurationLike,
) {
	v.stack_.AddValue(not.Element(duration.AsString()))
}

func (v *deflator_) ProcessGlyph(
	glyph fra.GlyphLike,
) {
	v.stack_.AddValue(not.Element(glyph.AsString()))
}

func (v *deflator_) ProcessIdentifier(
	identifier string,
) {
	v.stack_.AddValue(identifier)
}

func (v *deflator_) ProcessInverse(
	inverse doc.Inverse,
) {
	var operation string
	switch inverse {
	case doc.Additive:
		operation = "-"
	case doc.Multiplicative:
		operation = "/"
	case doc.Conjugate:
		operation = "*"
	default:
		var message = fmt.Sprintf(
			"Found an unexpected value in a switch statement: %v(%T)",
			inverse,
			inverse,
		)
		panic(message)
	}
	v.stack_.AddValue(not.Inverse(operation))
}

func (v *deflator_) ProcessInvoke(
	invoke doc.Invoke,
) {
	var operation string
	switch invoke {
	case doc.Synchronous:
		operation = "<-"
	case doc.Asynchronous:
		operation = "<~"
	default:
		var message = fmt.Sprintf(
			"Found an unexpected value in a switch statement: %v(%T)",
			invoke,
			invoke,
		)
		panic(message)
	}
	v.stack_.AddValue(not.Invoke(operation))
}

func (v *deflator_) ProcessMoment(
	moment fra.MomentLike,
) {
	v.stack_.AddValue(not.Element(moment.AsString()))
}

func (v *deflator_) ProcessName(
	name fra.NameLike,
) {
	v.stack_.AddValue(not.String(name.AsString()))
}

func (v *deflator_) ProcessNarrative(
	narrative fra.NarrativeLike,
) {
	v.stack_.AddValue(not.String(narrative.AsString()))
}

func (v *deflator_) ProcessNote(
	note string,
) {
	v.stack_.AddValue(note)
}

func (v *deflator_) ProcessNumber(
	number fra.NumberLike,
) {
	v.stack_.AddValue(not.Element(number.AsString()))
}

func (v *deflator_) ProcessOperator(
	operator doc.Operator,
) {
	var operation any
	switch operator {
	case doc.Chain:
		operation = not.LexicalOperator("&")
	case doc.And:
		operation = not.LogicalOperator("and")
	case doc.San:
		operation = not.LogicalOperator("san")
	case doc.Ior:
		operation = not.LogicalOperator("ior")
	case doc.Xor:
		operation = not.LogicalOperator("xor")
	case doc.Plus:
		operation = not.ArithmeticOperator("+")
	case doc.Minus:
		operation = not.ArithmeticOperator("-")
	case doc.Times:
		operation = not.ArithmeticOperator("*")
	case doc.Divide:
		operation = not.ArithmeticOperator("/")
	case doc.Remainder:
		operation = not.ArithmeticOperator("%")
	case doc.Power:
		operation = not.ArithmeticOperator("^")
	case doc.Less:
		operation = not.ComparisonOperator("<")
	case doc.Equal:
		operation = not.ComparisonOperator("=")
	case doc.More:
		operation = not.ComparisonOperator(">")
	case doc.Is:
		operation = not.ComparisonOperator("is")
	case doc.Matches:
		operation = not.ComparisonOperator("matches")
	default:
		var message = fmt.Sprintf(
			"Found an unexpected value in a switch statement: %v(%T)",
			operator,
			operator,
		)
		panic(message)
	}
	v.stack_.AddValue(not.Operator(operation))
}

func (v *deflator_) ProcessPattern(
	pattern fra.PatternLike,
) {
	v.stack_.AddValue(not.String(pattern.AsString()))
}

func (v *deflator_) ProcessPercentage(
	percentage fra.PercentageLike,
) {
	v.stack_.AddValue(not.Element(percentage.AsString()))
}

func (v *deflator_) ProcessProbability(
	probability fra.ProbabilityLike,
) {
	v.stack_.AddValue(not.Element(probability.AsString()))
}

func (v *deflator_) ProcessQuote(
	quote fra.QuoteLike,
) {
	v.stack_.AddValue(not.String(quote.AsString()))
}

func (v *deflator_) ProcessResource(
	resource fra.ResourceLike,
) {
	v.stack_.AddValue(not.Element(resource.AsString()))
}

func (v *deflator_) ProcessSymbol(
	symbol fra.SymbolLike,
) {
	v.stack_.AddValue(not.Element(symbol.AsString()))
}

func (v *deflator_) ProcessTag(
	tag fra.TagLike,
) {
	v.stack_.AddValue(not.String(tag.AsString()))
}

func (v *deflator_) ProcessVersion(
	version fra.VersionLike,
) {
	v.stack_.AddValue(not.String(version.AsString()))
}

func (v *deflator_) PostprocessAcceptClause(
	acceptClause doc.AcceptClauseLike,
	index_ uint,
	count_ uint,
) {
	var expression = v.stack_.RemoveLast().(not.ExpressionLike)
	var message = not.Message(expression)
	v.stack_.AddValue(
		not.MessageHandling(
			not.AcceptClause("accept", message),
		),
	)
}

func (v *deflator_) PostprocessAttributes(
	attributes doc.AttributesLike,
	index_ uint,
	count_ uint,
) {
	var associations = fra.List[not.AssociationLike]()
	var iterator = attributes.GetAssociations().GetIterator()
	for iterator.HasNext() {
		var member = v.stack_.RemoveLast().(not.MemberLike)
		var primitive = not.Primitive(v.stack_.RemoveLast())
		var association = not.Association(primitive, ":", member)
		associations.AppendValue(association)
		iterator.GetNext()
	}
	associations.ReverseValues() // They were pulled off the stack in reverse order.
	v.stack_.AddValue(not.Collection(not.Attributes("[", associations, "]")))
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

func (v *deflator_) ProcessCheckoutClauseSlot(
	checkoutClause doc.CheckoutClauseLike,
	slot_ uint,
) {
	switch slot_ {
	case 2:
		if uti.IsUndefined(checkoutClause.GetOptionalAtLevel()) {
			v.stack_.AddValue(nil)
		}
	}
}

func (v *deflator_) PostprocessCheckoutClause(
	checkoutClause doc.CheckoutClauseLike,
	index_ uint,
	count_ uint,
) {
	var expression = v.stack_.RemoveLast().(not.ExpressionLike)
	var cited = not.Cited(expression)
	var atLevel not.AtLevelLike
	var optional = v.stack_.RemoveLast()
	if uti.IsDefined(optional) {
		var expression = optional.(not.ExpressionLike)
		atLevel = not.AtLevel("at", "level", expression)
	}
	var recipient not.RecipientLike
	switch actual := v.stack_.RemoveLast().(type) {
	case not.ElementLike:
		recipient = not.Recipient(not.Variable(actual.GetAny().(string)))
	case not.SubcomponentLike:
		recipient = not.Recipient(actual)
	}
	v.stack_.AddValue(
		not.RepositoryAccess(
			not.CheckoutClause("checkout", recipient, atLevel, "from", cited),
		),
	)
}

func (v *deflator_) PostprocessComplement(
	complement doc.ComplementLike,
	index_ uint,
	count_ uint,
) {
	var logical not.LogicalLike
	switch actual := v.stack_.RemoveLast().(type) {
	case string:
		logical = not.Logical(not.Value(actual))
	default:
		logical = not.Logical(actual)
	}
	v.stack_.AddValue(
		not.Complement(
			"not",
			logical,
		),
	)
}

func (v *deflator_) ProcessComponentSlot(
	component doc.ComponentLike,
	slot_ uint,
) {
	switch slot_ {
	case 1:
		if uti.IsUndefined(component.GetOptionalParameterization()) {
			var parameterization not.ParameterizationLike
			v.stack_.AddValue(parameterization)
		}
	}
}

func (v *deflator_) PostprocessComponent(
	component doc.ComponentLike,
	index_ uint,
	count_ uint,
) {
	var parameterization not.ParameterizationLike
	var optional = v.stack_.RemoveLast()
	if uti.IsDefined(optional) {
		parameterization = optional.(not.ParameterizationLike)
	}
	var entity = not.Entity(v.stack_.RemoveLast())
	v.stack_.AddValue(
		not.Component(
			entity,
			parameterization,
		),
	)
}

func (v *deflator_) ProcessConstraintSlot(
	constraint doc.ConstraintLike,
	slot_ uint,
) {
	switch slot_ {
	case 1:
		if uti.IsUndefined(constraint.GetOptionalParameterization()) {
			var parameterization not.ParameterizationLike
			v.stack_.AddValue(parameterization)
		}
	}
}

func (v *deflator_) PostprocessConstraint(
	constraint doc.ConstraintLike,
	index_ uint,
	count_ uint,
) {
	var parameterization not.ParameterizationLike
	var optional = v.stack_.RemoveLast()
	if uti.IsDefined(optional) {
		parameterization = optional.(not.ParameterizationLike)
	}
	var type_ = not.Type(v.stack_.RemoveLast())
	v.stack_.AddValue(
		not.Constraint(
			type_,
			parameterization,
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

func (v *deflator_) PostprocessDiscardClause(
	discardClause doc.DiscardClauseLike,
	index_ uint,
	count_ uint,
) {
	var expression = v.stack_.RemoveLast().(not.ExpressionLike)
	var draft = not.Draft(expression)
	v.stack_.AddValue(
		not.RepositoryAccess(
			not.DiscardClause("discard", draft),
		),
	)
}

func (v *deflator_) PostprocessDoClause(
	doClause doc.DoClauseLike,
	index_ uint,
	count_ uint,
) {
	var method = v.stack_.RemoveLast().(not.MethodLike)
	v.stack_.AddValue(
		not.ActionInduction(
			not.DoClause("do", method),
		),
	)
}

func (v *deflator_) PostprocessDocument(
	document doc.DocumentLike,
	index_ uint,
	count_ uint,
) {
	var component = v.stack_.RemoveLast().(not.ComponentLike)
	v.stack_.AddValue(
		not.Document(
			component,
		),
	)
}

func (v *deflator_) PostprocessExpression(
	expression doc.ExpressionLike,
	index_ uint,
	count_ uint,
) {
	var predicates = fra.List[not.PredicateLike]()
	var iterator = expression.GetPredicates().GetIterator()
	for iterator.HasNext() {
		var predicate = v.stack_.RemoveLast().(not.PredicateLike)
		predicates.AppendValue(predicate)
		iterator.GetNext()
	}
	predicates.ReverseValues() // They were pulled off the stack in reverse order.
	var subject not.SubjectLike
	switch actual := v.stack_.RemoveLast().(type) {
	case string:
		subject = not.Subject(not.Value(actual))
	default:
		subject = not.Subject(actual)
	}
	v.stack_.AddValue(not.Expression(subject, predicates))
}

func (v *deflator_) PostprocessFunction(
	function doc.FunctionLike,
	index_ uint,
	count_ uint,
) {
	var arguments = fra.List[not.ArgumentLike]()
	var iterator = function.GetArguments().GetIterator()
	for iterator.HasNext() {
		var argument not.ArgumentLike
		switch actual := v.stack_.RemoveLast().(type) {
		case string:
			argument = not.Argument(not.Value(actual))
		default:
			argument = not.Argument(not.Primitive(actual))
		}
		arguments.AppendValue(argument)
		iterator.GetNext()
	}
	arguments.ReverseValues() // They were pulled off the stack in reverse order.
	var identifier = v.stack_.RemoveLast().(string)
	v.stack_.AddValue(not.Function(identifier, "(", arguments, ")"))
}

func (v *deflator_) PostprocessIfClause(
	ifClause doc.IfClauseLike,
	index_ uint,
	count_ uint,
) {
	var procedure = v.stack_.RemoveLast().(not.ProcedureLike)
	var expression = v.stack_.RemoveLast().(not.ExpressionLike)
	var condition = not.Condition(expression)
	v.stack_.AddValue(
		not.FlowControl(
			not.IfClause("if", condition, "do", procedure),
		),
	)
}

func (v *deflator_) PostprocessInversion(
	inversion doc.InversionLike,
	index_ uint,
	count_ uint,
) {
	var numerical not.NumericalLike
	switch actual := v.stack_.RemoveLast().(type) {
	case string:
		numerical = not.Numerical(not.Value(actual))
	default:
		numerical = not.Numerical(actual)
	}
	var inverse = v.stack_.RemoveLast().(not.InverseLike)
	v.stack_.AddValue(
		not.Inversion(
			inverse,
			numerical,
		),
	)
}

func (v *deflator_) PostprocessItems(
	items doc.ItemsLike,
	index_ uint,
	count_ uint,
) {
	var members = fra.List[not.MemberLike]()
	var iterator = items.GetMembers().GetIterator()
	for iterator.HasNext() {
		var member = v.stack_.RemoveLast().(not.MemberLike)
		members.AppendValue(member)
		iterator.GetNext()
	}
	members.ReverseValues() // They were pulled off the stack in reverse order.
	v.stack_.AddValue(not.Collection(not.Items("[", members, "]")))
}

func (v *deflator_) PostprocessLetClause(
	letClause doc.LetClauseLike,
	index_ uint,
	count_ uint,
) {
	var expression = v.stack_.RemoveLast().(not.ExpressionLike)
	var assignment = v.stack_.RemoveLast().(not.AssignmentLike)
	var recipient not.RecipientLike
	switch actual := v.stack_.RemoveLast().(type) {
	case not.ElementLike:
		recipient = not.Recipient(not.Variable(actual.GetAny().(string)))
	case not.SubcomponentLike:
		recipient = not.Recipient(actual)
	}
	v.stack_.AddValue(
		not.ActionInduction(
			not.LetClause("let", recipient, assignment, expression),
		),
	)
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
	var template = not.Template(expression)
	v.stack_.AddValue(
		not.MatchingClause("matching", template, "do", procedure),
	)
}

func (v *deflator_) ProcessMemberSlot(
	member doc.MemberLike,
	slot_ uint,
) {
	switch slot_ {
	case 1:
		if uti.IsUndefined(member.GetOptionalNote()) {
			var note string
			v.stack_.AddValue(note)
		}
	}
}

func (v *deflator_) PostprocessMember(
	member doc.MemberLike,
	index_ uint,
	count_ uint,
) {
	var note = v.stack_.RemoveLast().(string)
	var component = v.stack_.RemoveLast().(not.ComponentLike)
	v.stack_.AddValue(
		not.Member(
			component,
			note,
		),
	)
}

func (v *deflator_) PostprocessMethod(
	method doc.MethodLike,
	index_ uint,
	count_ uint,
) {
	var arguments = fra.List[not.ArgumentLike]()
	var iterator = method.GetArguments().GetIterator()
	for iterator.HasNext() {
		var argument not.ArgumentLike
		switch actual := v.stack_.RemoveLast().(type) {
		case string:
			argument = not.Argument(not.Value(actual))
		default:
			argument = not.Argument(not.Primitive(actual))
		}
		arguments.AppendValue(argument)
		iterator.GetNext()
	}
	arguments.ReverseValues() // They were pulled off the stack in reverse order.
	var identifier = v.stack_.RemoveLast().(string)
	var invoke = v.stack_.RemoveLast().(not.InvokeLike)
	var target = v.stack_.RemoveLast().(string)
	v.stack_.AddValue(not.Method(target, invoke, identifier, "(", arguments, ")"))
}

func (v *deflator_) PostprocessNotarizeClause(
	notarizeClause doc.NotarizeClauseLike,
	index_ uint,
	count_ uint,
) {
	var expression = v.stack_.RemoveLast().(not.ExpressionLike)
	var cited = not.Cited(expression)
	expression = v.stack_.RemoveLast().(not.ExpressionLike)
	var draft = not.Draft(expression)
	v.stack_.AddValue(
		not.RepositoryAccess(
			not.NotarizeClause(
				"notarize",
				draft,
				"as",
				cited,
			),
		),
	)
}

func (v *deflator_) PostprocessOnClause(
	onClause doc.OnClauseLike,
	index_ uint,
	count_ uint,
) {
	var matchingClauses = fra.List[not.MatchingClauseLike]()
	var iterator = onClause.GetMatchingClauses().GetIterator()
	for iterator.HasNext() {
		var matchingClause = v.stack_.RemoveLast().(not.MatchingClauseLike)
		matchingClauses.AppendValue(matchingClause)
		iterator.GetNext()
	}
	matchingClauses.ReverseValues() // They were pulled off the stack in reverse order.
	var element = v.stack_.RemoveLast().(not.ElementLike)
	var failure = not.Failure(element.GetAny().(string))
	v.stack_.AddValue(
		not.OnClause(
			"on",
			failure,
			matchingClauses,
		),
	)
}

func (v *deflator_) PostprocessParameterization(
	parameterization doc.ParameterizationLike,
	index_ uint,
	count_ uint,
) {
	var parameters = fra.List[not.ParameterLike]()
	var iterator = parameterization.GetParameters().GetIterator()
	for iterator.HasNext() {
		var constraint = v.stack_.RemoveLast().(not.ConstraintLike)
		var element = v.stack_.RemoveLast().(not.ElementLike)
		var symbol = element.GetAny().(string)
		var parameter = not.Parameter(symbol, ":", constraint)
		parameters.AppendValue(parameter)
		iterator.GetNext()
	}
	parameters.ReverseValues() // They were pulled off the stack in reverse order.
	v.stack_.AddValue(not.Parameterization("(", parameters, ")"))
}

func (v *deflator_) PostprocessPostClause(
	postClause doc.PostClauseLike,
	index_ uint,
	count_ uint,
) {
	var expression = v.stack_.RemoveLast().(not.ExpressionLike)
	var bag = not.Bag(expression)
	expression = v.stack_.RemoveLast().(not.ExpressionLike)
	var message = not.Message(expression)
	v.stack_.AddValue(
		not.MessageHandling(
			not.PostClause("post", message, "to", bag),
		),
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
	var operator = v.stack_.RemoveLast().(not.OperatorLike)
	v.stack_.AddValue(
		not.Predicate(operator, expression),
	)
}

func (v *deflator_) PostprocessProcedure(
	procedure doc.ProcedureLike,
	index_ uint,
	count_ uint,
) {
	var lines = fra.List[not.LineLike]()
	var iterator = procedure.GetLines().GetIterator()
	for iterator.HasNext() {
		var line = v.stack_.RemoveLast().(not.LineLike)
		lines.AppendValue(line)
		iterator.GetNext()
	}
	lines.ReverseValues() // They were pulled off the stack in reverse order.
	v.stack_.AddValue(not.Procedure("{", lines, "}"))
}

func (v *deflator_) PostprocessPublishClause(
	publishClause doc.PublishClauseLike,
	index_ uint,
	count_ uint,
) {
	var expression = v.stack_.RemoveLast().(not.ExpressionLike)
	var event = not.Event(expression)
	v.stack_.AddValue(
		not.MessageHandling(
			not.PublishClause("publish", event),
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
	case doc.Inclusive:
		right = not.Right("]")
	case doc.Exclusive:
		right = not.Right(")")
	default:
		var message = fmt.Sprintf(
			"Found an unexpected value in a switch statement: %v(%T)",
			extent,
			extent,
		)
		panic(message)
	}
	var primitive2 = not.Primitive(v.stack_.RemoveLast())
	var primitive1 = not.Primitive(v.stack_.RemoveLast())
	var left not.LeftLike
	extent = range_.GetLeft()
	switch extent {
	case doc.Inclusive:
		left = not.Left("[")
	case doc.Exclusive:
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
		not.Range(
			left,
			primitive1,
			"..",
			primitive2,
			right,
		),
	)
}

func (v *deflator_) PostprocessReferent(
	referent doc.ReferentLike,
	index_ uint,
	count_ uint,
) {
	var reference not.ReferenceLike
	switch actual := v.stack_.RemoveLast().(type) {
	case string:
		reference = not.Reference(not.Value(actual))
	default:
		reference = not.Reference(actual)
	}
	v.stack_.AddValue(not.Referent("@", reference))
}

func (v *deflator_) PostprocessRejectClause(
	rejectClause doc.RejectClauseLike,
	index_ uint,
	count_ uint,
) {
	var expression = v.stack_.RemoveLast().(not.ExpressionLike)
	var message = not.Message(expression)
	v.stack_.AddValue(
		not.MessageHandling(
			not.RejectClause(
				"reject",
				message,
			),
		),
	)
}

func (v *deflator_) PostprocessRetrieveClause(
	retrieveClause doc.RetrieveClauseLike,
	index_ uint,
	count_ uint,
) {
	var expression = v.stack_.RemoveLast().(not.ExpressionLike)
	var bag = not.Bag(expression)
	var recipient not.RecipientLike
	switch actual := v.stack_.RemoveLast().(type) {
	case not.ElementLike:
		recipient = not.Recipient(not.Variable(actual.GetAny().(string)))
	case not.SubcomponentLike:
		recipient = not.Recipient(actual)
	}
	v.stack_.AddValue(
		not.MessageHandling(
			not.RetrieveClause(
				"retrieve",
				recipient,
				"from",
				bag,
			),
		),
	)
}

func (v *deflator_) PostprocessReturnClause(
	returnClause doc.ReturnClauseLike,
	index_ uint,
	count_ uint,
) {
	var expression = v.stack_.RemoveLast().(not.ExpressionLike)
	var result = not.Result(expression)
	v.stack_.AddValue(
		not.FlowControl(
			not.ReturnClause("return", result),
		),
	)
}

func (v *deflator_) PostprocessSaveClause(
	saveClause doc.SaveClauseLike,
	index_ uint,
	count_ uint,
) {
	var expression = v.stack_.RemoveLast().(not.ExpressionLike)
	var cited = not.Cited(expression)
	expression = v.stack_.RemoveLast().(not.ExpressionLike)
	var draft = not.Draft(expression)
	v.stack_.AddValue(
		not.RepositoryAccess(
			not.SaveClause("save", draft, "as", cited),
		),
	)
}

func (v *deflator_) PostprocessSelectClause(
	selectClause doc.SelectClauseLike,
	index_ uint,
	count_ uint,
) {
	var matchingClauses = fra.List[not.MatchingClauseLike]()
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

func (v *deflator_) ProcessStatementSlot(
	statement doc.StatementLike,
	slot_ uint,
) {
	switch slot_ {
	case 1:
		if uti.IsUndefined(statement.GetOptionalOnClause()) {
			v.stack_.AddValue(nil)
		}
	}
}

func (v *deflator_) PostprocessStatement(
	statement doc.StatementLike,
	index_ uint,
	count_ uint,
) {
	var onClause not.OnClauseLike
	var optional = v.stack_.RemoveLast()
	if uti.IsDefined(optional) {
		onClause = optional.(not.OnClauseLike)
	}
	var mainClause = not.MainClause(v.stack_.RemoveLast())
	v.stack_.AddValue(not.Line(not.Statement(mainClause, onClause)))
}

func (v *deflator_) PostprocessSubcomponent(
	subcomponent doc.SubcomponentLike,
	index_ uint,
	count_ uint,
) {
	var indexes = fra.List[not.IndexLike]()
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
	var exception = not.Exception(expression)
	v.stack_.AddValue(
		not.FlowControl(
			not.ThrowClause("throw", exception),
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
	var condition = not.Condition(expression)
	v.stack_.AddValue(
		not.FlowControl(
			not.WhileClause(
				"while",
				condition,
				"do",
				procedure,
			),
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
	var sequence = not.Sequence(expression)
	var element = v.stack_.RemoveLast().(not.ElementLike)
	var variable = not.Variable(element.GetAny().(string))
	v.stack_.AddValue(
		not.FlowControl(
			not.WithClause(
				"with",
				"each",
				variable,
				"in",
				sequence,
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
	stack_ fra.StackLike[any]

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
