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
		stack_: fra.Stack[any](),

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
	v.stack_.AddValue(not.Annotation(annotation))
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
	v.stack_.AddValue(not.Annotation(comment))
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
	v.stack_.AddValue(not.Annotation(note))
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
	case doc.Eor:
		operation = not.LogicalOperator("eor")
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
	var message = v.stack_.RemoveLast().(not.MessageLike)
	v.stack_.AddValue(
		not.MainClause(not.MessageHandling(not.AcceptClause("accept", message))),
	)
}

func (v *deflator_) PostprocessAttributes(
	attributes doc.AttributesLike,
	index_ uint,
	count_ uint,
) {
	var list = fra.List[not.AssociationLike]()
	var associations = attributes.GetAssociations()
	var iterator = associations.GetIterator()
	for iterator.HasNext() {
		iterator.GetNext()
		var document = v.stack_.RemoveLast().(not.DocumentLike)
		var primitive = v.stack_.RemoveLast().(not.PrimitiveLike)
		var association = not.Association(primitive, ":", document)
		list.AppendValue(association)
	}
	list.ReverseValues() // They were pulled off the stack in reverse order.
	v.stack_.AddValue(not.Collection(not.Attributes("[", list, "]")))
}

func (v *deflator_) PostprocessBreakClause(
	breakClause doc.BreakClauseLike,
	index_ uint,
	count_ uint,
) {
	v.stack_.AddValue(
		not.MainClause(not.FlowControl(not.BreakClause("break", "loop"))),
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
	var cited = v.stack_.RemoveLast().(not.CitedLike)
	var atLevel = v.stack_.RemoveLast().(not.AtLevelLike)
	var recipient not.RecipientLike
	switch actual := v.stack_.RemoveLast().(type) {
	case not.VariableLike:
		recipient = not.Recipient(actual)
	case not.SubcomponentLike:
		recipient = not.Recipient(actual)
	}
	v.stack_.AddValue(
		not.MainClause(
			not.RepositoryAccess(
				not.CheckoutClause(
					"checkout",
					recipient,
					atLevel,
					"from",
					cited,
				),
			),
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
	case not.DocumentLike:
		logical = not.Logical(actual)
	case not.SubcomponentLike:
		logical = not.Logical(actual)
	case not.PrecedenceLike:
		logical = not.Logical(actual)
	case not.ReferentLike:
		logical = not.Logical(actual)
	case not.ComplementLike:
		logical = not.Logical(actual)
	case not.FunctionLike:
		logical = not.Logical(actual)
	case not.MethodLike:
		logical = not.Logical(actual)
	case not.ValueLike:
		logical = not.Logical(actual)
	}
	v.stack_.AddValue(
		not.Complement(
			"not",
			logical,
		),
	)
}

func (v *deflator_) PostprocessContinueClause(
	continueClause doc.ContinueClauseLike,
	index_ uint,
	count_ uint,
) {
	v.stack_.AddValue(
		not.MainClause(not.FlowControl(not.ContinueClause("continue", "loop"))),
	)
}

func (v *deflator_) PostprocessDiscardClause(
	discardClause doc.DiscardClauseLike,
	index_ uint,
	count_ uint,
) {
	var draft = v.stack_.RemoveLast().(not.DraftLike)
	v.stack_.AddValue(
		not.MainClause(not.RepositoryAccess(not.DiscardClause("discard", draft))),
	)
}

func (v *deflator_) PostprocessDoClause(
	doClause doc.DoClauseLike,
	index_ uint,
	count_ uint,
) {
	var method = v.stack_.RemoveLast().(not.MethodLike)
	v.stack_.AddValue(
		not.MainClause(not.ActionInduction(not.DoClause("do", method))),
	)
}

func (v *deflator_) ProcessDocumentSlot(
	document doc.DocumentLike,
	slot_ uint,
) {
	switch slot_ {
	case 2:
		if uti.IsUndefined(document.GetOptionalParameters()) {
			var parameters not.ParametersLike
			v.stack_.AddValue(parameters)
		}
		if uti.IsUndefined(document.GetOptionalNote()) {
			var note string
			v.stack_.AddValue(note)
		}
	}
}

func (v *deflator_) PostprocessDocument(
	document doc.DocumentLike,
	index_ uint,
	count_ uint,
) {
	var note = v.stack_.RemoveLast().(string)
	var parameters not.ParametersLike
	var optional = v.stack_.RemoveLast()
	if uti.IsDefined(optional) {
		parameters = optional.(not.ParametersLike)
	}
	var component not.ComponentLike
	switch actual := v.stack_.RemoveLast().(type) {
	case not.ElementLike:
		component = not.Component(actual)
	case not.StringLike:
		component = not.Component(actual)
	case not.RangeLike:
		component = not.Component(actual)
	case not.AttributesLike:
		component = not.Component(not.Collection(actual))
	case not.EntitiesLike:
		component = not.Component(not.Collection(actual))
	case not.ProcedureLike:
		component = not.Component(actual)
	}
	v.stack_.AddValue(
		not.Document(
			component,
			parameters,
			note,
		),
	)
}

func (v *deflator_) PostprocessEntities(
	entities doc.EntitiesLike,
	index_ uint,
	count_ uint,
) {
	var list = fra.List[not.ItemLike]()
	var items = entities.GetItems()
	var iterator = items.GetIterator()
	for iterator.HasNext() {
		iterator.GetNext()
		var document = v.stack_.RemoveLast().(not.DocumentLike)
		var item = not.Item(document)
		list.AppendValue(item)
	}
	list.ReverseValues() // They were pulled off the stack in reverse order.
	v.stack_.AddValue(not.Collection(not.Entities("[", list, "]")))
}

func (v *deflator_) PostprocessExpression(
	expression doc.ExpressionLike,
	index_ uint,
	count_ uint,
) {
	var predicates = v.stack_.RemoveLast().(fra.ListLike[not.PredicateLike])
	var subject not.SubjectLike
	switch actual := v.stack_.RemoveLast().(type) {
	case not.DocumentLike:
		subject = not.Subject(actual)
	case not.SubcomponentLike:
		subject = not.Subject(actual)
	case not.PrecedenceLike:
		subject = not.Subject(actual)
	case not.ReferentLike:
		subject = not.Subject(actual)
	case not.ComplementLike:
		subject = not.Subject(actual)
	case not.InversionLike:
		subject = not.Subject(actual)
	case not.MagnitudeLike:
		subject = not.Subject(actual)
	case not.FunctionLike:
		subject = not.Subject(actual)
	case not.MethodLike:
		subject = not.Subject(actual)
	case not.ValueLike:
		subject = not.Subject(actual)
	}
	v.stack_.AddValue(
		not.Expression(
			subject,
			predicates,
		),
	)
}

func (v *deflator_) PostprocessFunction(
	function doc.FunctionLike,
	index_ uint,
	count_ uint,
) {
	var arguments = v.stack_.RemoveLast().(fra.ListLike[not.ArgumentLike])
	var identifier = v.stack_.RemoveLast().(string)
	v.stack_.AddValue(
		not.Function(
			identifier,
			"(",
			arguments,
			")",
		),
	)
}

func (v *deflator_) PostprocessIfClause(
	ifClause doc.IfClauseLike,
	index_ uint,
	count_ uint,
) {
	var procedure = v.stack_.RemoveLast().(not.ProcedureLike)
	var condition = v.stack_.RemoveLast().(not.ConditionLike)
	v.stack_.AddValue(
		not.MainClause(
			not.FlowControl(
				not.IfClause(
					"if",
					condition,
					"do",
					procedure,
				),
			),
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
	case not.DocumentLike:
		numerical = not.Numerical(actual)
	case not.SubcomponentLike:
		numerical = not.Numerical(actual)
	case not.PrecedenceLike:
		numerical = not.Numerical(actual)
	case not.ReferentLike:
		numerical = not.Numerical(actual)
	case not.InversionLike:
		numerical = not.Numerical(actual)
	case not.MagnitudeLike:
		numerical = not.Numerical(actual)
	case not.FunctionLike:
		numerical = not.Numerical(actual)
	case not.MethodLike:
		numerical = not.Numerical(actual)
	case not.ValueLike:
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

func (v *deflator_) PostprocessLetClause(
	letClause doc.LetClauseLike,
	index_ uint,
	count_ uint,
) {
	var expression = v.stack_.RemoveLast().(not.ExpressionLike)
	var assignment = v.stack_.RemoveLast().(not.AssignmentLike)
	var recipient not.RecipientLike
	switch actual := v.stack_.RemoveLast().(type) {
	case not.VariableLike:
		recipient = not.Recipient(actual)
	case not.SubcomponentLike:
		recipient = not.Recipient(actual)
	}
	v.stack_.AddValue(
		not.MainClause(
			not.ActionInduction(
				not.LetClause(
					"let",
					recipient,
					assignment,
					expression,
				),
			),
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
		not.Magnitude(
			"|",
			expression,
			"|",
		),
	)
}

func (v *deflator_) PostprocessMatchingClause(
	matchingClause doc.MatchingClauseLike,
	index_ uint,
	count_ uint,
) {
	var procedure = v.stack_.RemoveLast().(not.ProcedureLike)
	var template = v.stack_.RemoveLast().(not.TemplateLike)
	v.stack_.AddValue(
		not.MatchingClause(
			"matching",
			template,
			"do",
			procedure,
		),
	)
}

func (v *deflator_) PostprocessMethod(
	method doc.MethodLike,
	index_ uint,
	count_ uint,
) {
	var arguments = v.stack_.RemoveLast().(fra.ListLike[not.ArgumentLike])
	var identifier = v.stack_.RemoveLast().(string)
	var invoke = v.stack_.RemoveLast().(not.InvokeLike)
	var target = v.stack_.RemoveLast().(string)
	v.stack_.AddValue(
		not.Method(
			target,
			invoke,
			identifier,
			"(",
			arguments,
			")",
		),
	)
}

func (v *deflator_) PostprocessNotarizeClause(
	notarizeClause doc.NotarizeClauseLike,
	index_ uint,
	count_ uint,
) {
	var cited = v.stack_.RemoveLast().(not.CitedLike)
	var draft = v.stack_.RemoveLast().(not.DraftLike)
	v.stack_.AddValue(
		not.MainClause(
			not.RepositoryAccess(
				not.NotarizeClause(
					"notarize",
					draft,
					"as",
					cited,
				),
			),
		),
	)
}

func (v *deflator_) PostprocessOnClause(
	onClause doc.OnClauseLike,
	index_ uint,
	count_ uint,
) {
	var matchingClauses = v.stack_.RemoveLast().(fra.ListLike[not.MatchingClauseLike])
	var failure = v.stack_.RemoveLast().(not.FailureLike)
	v.stack_.AddValue(
		not.OnClause(
			"on",
			failure,
			matchingClauses,
		),
	)
}

func (v *deflator_) PostprocessParameters(
	parameters doc.ParametersLike,
	index_ uint,
	count_ uint,
) {
	var associations = v.stack_.RemoveLast().(fra.ListLike[not.AssociationLike])
	v.stack_.AddValue(
		not.Parameters(
			"[",
			associations,
			"]",
		),
	)
}

func (v *deflator_) PostprocessPostClause(
	postClause doc.PostClauseLike,
	index_ uint,
	count_ uint,
) {
	var bag = v.stack_.RemoveLast().(not.BagLike)
	var message = v.stack_.RemoveLast().(not.MessageLike)
	v.stack_.AddValue(
		not.MainClause(
			not.MessageHandling(
				not.PostClause(
					"post",
					message,
					"to",
					bag,
				),
			),
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
		not.Precedence(
			"(",
			expression,
			")",
		),
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
		not.Predicate(
			operator,
			expression,
		),
	)
}

func (v *deflator_) PostprocessProcedure(
	procedure doc.ProcedureLike,
	index_ uint,
	count_ uint,
) {
	var lines = v.stack_.RemoveLast().(fra.ListLike[not.LineLike])
	v.stack_.AddValue(
		not.Procedure(
			"{",
			lines,
			"}",
		),
	)
}

func (v *deflator_) PostprocessPublishClause(
	publishClause doc.PublishClauseLike,
	index_ uint,
	count_ uint,
) {
	var event = v.stack_.RemoveLast().(not.EventLike)
	v.stack_.AddValue(
		not.MainClause(
			not.MessageHandling(
				not.PublishClause(
					"publish",
					event,
				),
			),
		),
	)
}

func (v *deflator_) PostprocessRange(
	range_ doc.RangeLike,
	index_ uint,
	count_ uint,
) {
	var right not.RightLike
	switch range_.GetRight() {
	case doc.Inclusive:
		right = not.Right("]")
	case doc.Exclusive:
		right = not.Right(")")
	}
	var primitive2 = v.stack_.RemoveLast().(not.PrimitiveLike)
	var primitive1 = v.stack_.RemoveLast().(not.PrimitiveLike)
	var left not.LeftLike
	switch range_.GetLeft() {
	case doc.Inclusive:
		left = not.Left("[")
	case doc.Exclusive:
		left = not.Left("(")
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
	case not.DocumentLike:
		reference = not.Reference(actual)
	case not.SubcomponentLike:
		reference = not.Reference(actual)
	case not.ReferentLike:
		reference = not.Reference(actual)
	case not.FunctionLike:
		reference = not.Reference(actual)
	case not.MethodLike:
		reference = not.Reference(actual)
	case not.ValueLike:
		reference = not.Reference(actual)
	}
	v.stack_.AddValue( not.Referent("@", reference))
}

func (v *deflator_) PostprocessRejectClause(
	rejectClause doc.RejectClauseLike,
	index_ uint,
	count_ uint,
) {
	var message = v.stack_.RemoveLast().(not.MessageLike)
	v.stack_.AddValue(
		not.MainClause(
			not.MessageHandling(
				not.RejectClause(
					"reject",
					message,
				),
			),
		),
	)
}

func (v *deflator_) PostprocessRetrieveClause(
	retrieveClause doc.RetrieveClauseLike,
	index_ uint,
	count_ uint,
) {
	var bag = v.stack_.RemoveLast().(not.BagLike)
	var recipient not.RecipientLike
	switch actual := v.stack_.RemoveLast().(type) {
	case not.VariableLike:
		recipient = not.Recipient(actual)
	case not.SubcomponentLike:
		recipient = not.Recipient(actual)
	}
	v.stack_.AddValue(
		not.MainClause(
			not.MessageHandling(
				not.RetrieveClause(
					"retrieve",
					recipient,
					"from",
					bag,
				),
			),
		),
	)
}

func (v *deflator_) PostprocessReturnClause(
	returnClause doc.ReturnClauseLike,
	index_ uint,
	count_ uint,
) {
	var result = v.stack_.RemoveLast().(not.ResultLike)
	v.stack_.AddValue(
		not.MainClause(
			not.FlowControl(
				not.ReturnClause(
					"return",
					result,
				),
			),
		),
	)
}

func (v *deflator_) PostprocessSaveClause(
	saveClause doc.SaveClauseLike,
	index_ uint,
	count_ uint,
) {
	var cited = v.stack_.RemoveLast().(not.CitedLike)
	var draft = v.stack_.RemoveLast().(not.DraftLike)
	v.stack_.AddValue(
		not.MainClause(
			not.RepositoryAccess(
				not.SaveClause(
					"save",
					draft,
					"as",
					cited,
				),
			),
		),
	)
}

func (v *deflator_) PostprocessSelectClause(
	selectClause doc.SelectClauseLike,
	index_ uint,
	count_ uint,
) {
	var matchingClauses = v.stack_.RemoveLast().(fra.ListLike[not.MatchingClauseLike])
	var expression = v.stack_.RemoveLast().(not.ExpressionLike)
	v.stack_.AddValue(
		not.MainClause(
			not.FlowControl(
				not.SelectClause(
					"select",
					expression,
					matchingClauses,
				),
			),
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
	var onClause = v.stack_.RemoveLast().(not.OnClauseLike)
	var mainClause = v.stack_.RemoveLast().(not.MainClauseLike)
	v.stack_.AddValue(
		not.Statement(
			mainClause,
			onClause,
		),
	)
}

func (v *deflator_) PostprocessSubcomponent(
	subcomponent doc.SubcomponentLike,
	index_ uint,
	count_ uint,
) {
	var indexes = v.stack_.RemoveLast().(fra.ListLike[not.IndexLike])
	var identifier = v.stack_.RemoveLast().(string)
	v.stack_.AddValue(
		not.Subcomponent(
			identifier,
			"[",
			indexes,
			"]",
		),
	)
}

func (v *deflator_) PostprocessThrowClause(
	throwClause doc.ThrowClauseLike,
	index_ uint,
	count_ uint,
) {
	var exception = v.stack_.RemoveLast().(not.ExceptionLike)
	v.stack_.AddValue(
		not.ThrowClause(
			"throw",
			exception,
		),
	)
}

func (v *deflator_) PostprocessWhileClause(
	whileClause doc.WhileClauseLike,
	index_ uint,
	count_ uint,
) {
	var procedure = v.stack_.RemoveLast().(not.ProcedureLike)
	var condition = v.stack_.RemoveLast().(not.ConditionLike)
	v.stack_.AddValue(
		not.WhileClause(
			"while",
			condition,
			"do",
			procedure,
		),
	)
}

func (v *deflator_) PostprocessWithClause(
	withClause doc.WithClauseLike,
	index_ uint,
	count_ uint,
) {
	var procedure = v.stack_.RemoveLast().(not.ProcedureLike)
	var sequence = v.stack_.RemoveLast().(not.SequenceLike)
	var variable = v.stack_.RemoveLast().(not.VariableLike)
	v.stack_.AddValue(
		not.WithClause(
			"with",
			"each",
			variable,
			"in",
			sequence,
			"do",
			procedure,
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
