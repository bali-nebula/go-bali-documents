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
│                     Any updates to it may be overwritten.                    │
└──────────────────────────────────────────────────────────────────────────────┘
*/

package grammar

import (
	ast "github.com/bali-nebula/go-bali-documents/v3/ast"
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

// INSTANCE INTERFACE

// Principal Methods

func (v *visitor_) GetClass() VisitorClassLike {
	return visitorClass()
}

func (v *visitor_) VisitDocument(
	document ast.DocumentLike,
) {
	v.processor_.PreprocessDocument(
		document,
		1,
		1,
	)
	v.visitDocument(document)
	v.processor_.PostprocessDocument(
		document,
		1,
		1,
	)
}

// PROTECTED INTERFACE

// Private Methods

func (v *visitor_) visitAcceptClause(
	acceptClause ast.AcceptClauseLike,
) {
	var delimiter = acceptClause.GetDelimiter()
	v.processor_.ProcessDelimiter(delimiter)
	// Visit slot 1 between terms.
	v.processor_.ProcessAcceptClauseSlot(1)

	var message = acceptClause.GetMessage()
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

func (v *visitor_) visitAdditionalArgument(
	additionalArgument ast.AdditionalArgumentLike,
) {
	var delimiter = additionalArgument.GetDelimiter()
	v.processor_.ProcessDelimiter(delimiter)
	// Visit slot 1 between terms.
	v.processor_.ProcessAdditionalArgumentSlot(1)

	var argument = additionalArgument.GetArgument()
	v.processor_.PreprocessArgument(
		argument,
		1,
		1,
	)
	v.visitArgument(argument)
	v.processor_.PostprocessArgument(
		argument,
		1,
		1,
	)
}

func (v *visitor_) visitAdditionalAssociation(
	additionalAssociation ast.AdditionalAssociationLike,
) {
	var delimiter = additionalAssociation.GetDelimiter()
	v.processor_.ProcessDelimiter(delimiter)
	// Visit slot 1 between terms.
	v.processor_.ProcessAdditionalAssociationSlot(1)

	var association = additionalAssociation.GetAssociation()
	v.processor_.PreprocessAssociation(
		association,
		1,
		1,
	)
	v.visitAssociation(association)
	v.processor_.PostprocessAssociation(
		association,
		1,
		1,
	)
}

func (v *visitor_) visitAdditionalIndex(
	additionalIndex ast.AdditionalIndexLike,
) {
	var delimiter = additionalIndex.GetDelimiter()
	v.processor_.ProcessDelimiter(delimiter)
	// Visit slot 1 between terms.
	v.processor_.ProcessAdditionalIndexSlot(1)

	var index = additionalIndex.GetIndex()
	v.processor_.PreprocessIndex(
		index,
		1,
		1,
	)
	v.visitIndex(index)
	v.processor_.PostprocessIndex(
		index,
		1,
		1,
	)
}

func (v *visitor_) visitAdditionalStatement(
	additionalStatement ast.AdditionalStatementLike,
) {
	var delimiter = additionalStatement.GetDelimiter()
	v.processor_.ProcessDelimiter(delimiter)
	// Visit slot 1 between terms.
	v.processor_.ProcessAdditionalStatementSlot(1)

	var statement = additionalStatement.GetStatement()
	v.processor_.PreprocessStatement(
		statement,
		1,
		1,
	)
	v.visitStatement(statement)
	v.processor_.PostprocessStatement(
		statement,
		1,
		1,
	)
}

func (v *visitor_) visitAdditionalValue(
	additionalValue ast.AdditionalValueLike,
) {
	var delimiter = additionalValue.GetDelimiter()
	v.processor_.ProcessDelimiter(delimiter)
	// Visit slot 1 between terms.
	v.processor_.ProcessAdditionalValueSlot(1)

	var component = additionalValue.GetComponent()
	v.processor_.PreprocessComponent(
		component,
		1,
		1,
	)
	v.visitComponent(component)
	v.processor_.PostprocessComponent(
		component,
		1,
		1,
	)
}

func (v *visitor_) visitAnnotatedAssociation(
	annotatedAssociation ast.AnnotatedAssociationLike,
) {
	var association = annotatedAssociation.GetAssociation()
	v.processor_.PreprocessAssociation(
		association,
		1,
		1,
	)
	v.visitAssociation(association)
	v.processor_.PostprocessAssociation(
		association,
		1,
		1,
	)
	// Visit slot 1 between terms.
	v.processor_.ProcessAnnotatedAssociationSlot(1)

	var optionalNote = annotatedAssociation.GetOptionalNote()
	if uti.IsDefined(optionalNote) {
		v.processor_.ProcessNote(optionalNote)
	}
}

func (v *visitor_) visitAnnotatedStatement(
	annotatedStatement ast.AnnotatedStatementLike,
) {
	// Visit the possible annotatedStatement rule types.
	switch actual := annotatedStatement.GetAny().(type) {
	case ast.CommentLineLike:
		v.processor_.PreprocessCommentLine(
			actual,
			1,
			1,
		)
		v.visitCommentLine(actual)
		v.processor_.PostprocessCommentLine(
			actual,
			1,
			1,
		)
	case ast.StatementLineLike:
		v.processor_.PreprocessStatementLine(
			actual,
			1,
			1,
		)
		v.visitStatementLine(actual)
		v.processor_.PostprocessStatementLine(
			actual,
			1,
			1,
		)
	}
}

func (v *visitor_) visitAnnotatedValue(
	annotatedValue ast.AnnotatedValueLike,
) {
	var component = annotatedValue.GetComponent()
	v.processor_.PreprocessComponent(
		component,
		1,
		1,
	)
	v.visitComponent(component)
	v.processor_.PostprocessComponent(
		component,
		1,
		1,
	)
	// Visit slot 1 between terms.
	v.processor_.ProcessAnnotatedValueSlot(1)

	var optionalNote = annotatedValue.GetOptionalNote()
	if uti.IsDefined(optionalNote) {
		v.processor_.ProcessNote(optionalNote)
	}
}

func (v *visitor_) visitArgument(
	argument ast.ArgumentLike,
) {
	var identifier = argument.GetIdentifier()
	v.processor_.ProcessIdentifier(identifier)
}

func (v *visitor_) visitArguments(
	arguments ast.ArgumentsLike,
) {
	var argument = arguments.GetArgument()
	v.processor_.PreprocessArgument(
		argument,
		1,
		1,
	)
	v.visitArgument(argument)
	v.processor_.PostprocessArgument(
		argument,
		1,
		1,
	)
	// Visit slot 1 between terms.
	v.processor_.ProcessArgumentsSlot(1)

	var additionalArgumentsIndex uint
	var additionalArguments = arguments.GetAdditionalArguments().GetIterator()
	var additionalArgumentsCount = uint(additionalArguments.GetSize())
	for additionalArguments.HasNext() {
		additionalArgumentsIndex++
		var rule = additionalArguments.GetNext()
		v.processor_.PreprocessAdditionalArgument(
			rule,
			additionalArgumentsIndex,
			additionalArgumentsCount,
		)
		v.visitAdditionalArgument(rule)
		v.processor_.PostprocessAdditionalArgument(
			rule,
			additionalArgumentsIndex,
			additionalArgumentsCount,
		)
	}
}

func (v *visitor_) visitArithmetic(
	arithmetic ast.ArithmeticLike,
) {
	// Visit the possible arithmetic expression types.
	var actual = arithmetic.GetAny().(string)
	switch {
	case ScannerClass().MatchesType(actual, PlusToken):
		v.processor_.ProcessPlus(actual)
	case ScannerClass().MatchesType(actual, MinusToken):
		v.processor_.ProcessMinus(actual)
	case ScannerClass().MatchesType(actual, StarToken):
		v.processor_.ProcessStar(actual)
	case ScannerClass().MatchesType(actual, SlashToken):
		v.processor_.ProcessSlash(actual)
	case ScannerClass().MatchesType(actual, PercentToken):
		v.processor_.ProcessPercent(actual)
	case ScannerClass().MatchesType(actual, CaretToken):
		v.processor_.ProcessCaret(actual)
	}
}

func (v *visitor_) visitAssignment(
	assignment ast.AssignmentLike,
) {
	// Visit the possible assignment literal values.
	var actual = assignment.GetAny().(string)
	switch actual {
	case ":=":
		v.processor_.ProcessDelimiter(":=")
	case "?=":
		v.processor_.ProcessDelimiter("?=")
	case "+=":
		v.processor_.ProcessDelimiter("+=")
	case "-=":
		v.processor_.ProcessDelimiter("-=")
	case "*=":
		v.processor_.ProcessDelimiter("*=")
	case "/=":
		v.processor_.ProcessDelimiter("/=")
	}
}

func (v *visitor_) visitAssociation(
	association ast.AssociationLike,
) {
	var primitive = association.GetPrimitive()
	v.processor_.PreprocessPrimitive(
		primitive,
		1,
		1,
	)
	v.visitPrimitive(primitive)
	v.processor_.PostprocessPrimitive(
		primitive,
		1,
		1,
	)
	// Visit slot 1 between terms.
	v.processor_.ProcessAssociationSlot(1)

	var delimiter = association.GetDelimiter()
	v.processor_.ProcessDelimiter(delimiter)
	// Visit slot 2 between terms.
	v.processor_.ProcessAssociationSlot(2)

	var component = association.GetComponent()
	v.processor_.PreprocessComponent(
		component,
		1,
		1,
	)
	v.visitComponent(component)
	v.processor_.PostprocessComponent(
		component,
		1,
		1,
	)
}

func (v *visitor_) visitAtLevel(
	atLevel ast.AtLevelLike,
) {
	var delimiter1 = atLevel.GetDelimiter1()
	v.processor_.ProcessDelimiter(delimiter1)
	// Visit slot 1 between terms.
	v.processor_.ProcessAtLevelSlot(1)

	var delimiter2 = atLevel.GetDelimiter2()
	v.processor_.ProcessDelimiter(delimiter2)
	// Visit slot 2 between terms.
	v.processor_.ProcessAtLevelSlot(2)

	var expression = atLevel.GetExpression()
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

func (v *visitor_) visitBag(
	bag ast.BagLike,
) {
	var expression = bag.GetExpression()
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

func (v *visitor_) visitBlocking(
	blocking ast.BlockingLike,
) {
	// Visit the possible blocking expression types.
	var actual = blocking.GetAny().(string)
	switch {
	case ScannerClass().MatchesType(actual, DotToken):
		v.processor_.ProcessDot(actual)
	case ScannerClass().MatchesType(actual, ArrowToken):
		v.processor_.ProcessArrow(actual)
	}
}

func (v *visitor_) visitBracket(
	bracket ast.BracketLike,
) {
	// Visit the possible bracket rule types.
	switch actual := bracket.GetAny().(type) {
	case ast.InclusiveLike:
		v.processor_.PreprocessInclusive(
			actual,
			1,
			1,
		)
		v.visitInclusive(actual)
		v.processor_.PostprocessInclusive(
			actual,
			1,
			1,
		)
	case ast.ExclusiveLike:
		v.processor_.PreprocessExclusive(
			actual,
			1,
			1,
		)
		v.visitExclusive(actual)
		v.processor_.PostprocessExclusive(
			actual,
			1,
			1,
		)
	}
}

func (v *visitor_) visitBreakClause(
	breakClause ast.BreakClauseLike,
) {
	var delimiter1 = breakClause.GetDelimiter1()
	v.processor_.ProcessDelimiter(delimiter1)
	// Visit slot 1 between terms.
	v.processor_.ProcessBreakClauseSlot(1)

	var delimiter2 = breakClause.GetDelimiter2()
	v.processor_.ProcessDelimiter(delimiter2)
}

func (v *visitor_) visitCheckoutClause(
	checkoutClause ast.CheckoutClauseLike,
) {
	var delimiter1 = checkoutClause.GetDelimiter1()
	v.processor_.ProcessDelimiter(delimiter1)
	// Visit slot 1 between terms.
	v.processor_.ProcessCheckoutClauseSlot(1)

	var recipient = checkoutClause.GetRecipient()
	v.processor_.PreprocessRecipient(
		recipient,
		1,
		1,
	)
	v.visitRecipient(recipient)
	v.processor_.PostprocessRecipient(
		recipient,
		1,
		1,
	)
	// Visit slot 2 between terms.
	v.processor_.ProcessCheckoutClauseSlot(2)

	var optionalAtLevel = checkoutClause.GetOptionalAtLevel()
	if uti.IsDefined(optionalAtLevel) {
		v.processor_.PreprocessAtLevel(
			optionalAtLevel,
			1,
			1,
		)
		v.visitAtLevel(optionalAtLevel)
		v.processor_.PostprocessAtLevel(
			optionalAtLevel,
			1,
			1,
		)
	}
	// Visit slot 3 between terms.
	v.processor_.ProcessCheckoutClauseSlot(3)

	var delimiter2 = checkoutClause.GetDelimiter2()
	v.processor_.ProcessDelimiter(delimiter2)
	// Visit slot 4 between terms.
	v.processor_.ProcessCheckoutClauseSlot(4)

	var cited = checkoutClause.GetCited()
	v.processor_.PreprocessCited(
		cited,
		1,
		1,
	)
	v.visitCited(cited)
	v.processor_.PostprocessCited(
		cited,
		1,
		1,
	)
}

func (v *visitor_) visitCited(
	cited ast.CitedLike,
) {
	var expression = cited.GetExpression()
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

func (v *visitor_) visitCollection(
	collection ast.CollectionLike,
) {
	// Visit the possible collection rule types.
	switch actual := collection.GetAny().(type) {
	case ast.MultilineAttributesLike:
		v.processor_.PreprocessMultilineAttributes(
			actual,
			1,
			1,
		)
		v.visitMultilineAttributes(actual)
		v.processor_.PostprocessMultilineAttributes(
			actual,
			1,
			1,
		)
	case ast.MultilineValuesLike:
		v.processor_.PreprocessMultilineValues(
			actual,
			1,
			1,
		)
		v.visitMultilineValues(actual)
		v.processor_.PostprocessMultilineValues(
			actual,
			1,
			1,
		)
	case ast.InclusiveRangeLike:
		v.processor_.PreprocessInclusiveRange(
			actual,
			1,
			1,
		)
		v.visitInclusiveRange(actual)
		v.processor_.PostprocessInclusiveRange(
			actual,
			1,
			1,
		)
	case ast.ExclusiveRangeLike:
		v.processor_.PreprocessExclusiveRange(
			actual,
			1,
			1,
		)
		v.visitExclusiveRange(actual)
		v.processor_.PostprocessExclusiveRange(
			actual,
			1,
			1,
		)
	case ast.InlineAttributesLike:
		v.processor_.PreprocessInlineAttributes(
			actual,
			1,
			1,
		)
		v.visitInlineAttributes(actual)
		v.processor_.PostprocessInlineAttributes(
			actual,
			1,
			1,
		)
	case ast.InlineValuesLike:
		v.processor_.PreprocessInlineValues(
			actual,
			1,
			1,
		)
		v.visitInlineValues(actual)
		v.processor_.PostprocessInlineValues(
			actual,
			1,
			1,
		)
	case ast.NoAttributesLike:
		v.processor_.PreprocessNoAttributes(
			actual,
			1,
			1,
		)
		v.visitNoAttributes(actual)
		v.processor_.PostprocessNoAttributes(
			actual,
			1,
			1,
		)
	case ast.NoValuesLike:
		v.processor_.PreprocessNoValues(
			actual,
			1,
			1,
		)
		v.visitNoValues(actual)
		v.processor_.PostprocessNoValues(
			actual,
			1,
			1,
		)
	}
}

func (v *visitor_) visitCommentLine(
	commentLine ast.CommentLineLike,
) {
	var comment = commentLine.GetComment()
	v.processor_.ProcessComment(comment)
}

func (v *visitor_) visitComparison(
	comparison ast.ComparisonLike,
) {
	// Visit the possible comparison expression types.
	var actual = comparison.GetAny().(string)
	switch {
	case ScannerClass().MatchesType(actual, LessToken):
		v.processor_.ProcessLess(actual)
	case ScannerClass().MatchesType(actual, EqualToken):
		v.processor_.ProcessEqual(actual)
	case ScannerClass().MatchesType(actual, MoreToken):
		v.processor_.ProcessMore(actual)
	}
}

func (v *visitor_) visitComplement(
	complement ast.ComplementLike,
) {
	var delimiter = complement.GetDelimiter()
	v.processor_.ProcessDelimiter(delimiter)
	// Visit slot 1 between terms.
	v.processor_.ProcessComplementSlot(1)

	var logical = complement.GetLogical()
	v.processor_.PreprocessLogical(
		logical,
		1,
		1,
	)
	v.visitLogical(logical)
	v.processor_.PostprocessLogical(
		logical,
		1,
		1,
	)
}

func (v *visitor_) visitComponent(
	component ast.ComponentLike,
) {
	var entity = component.GetEntity()
	v.processor_.PreprocessEntity(
		entity,
		1,
		1,
	)
	v.visitEntity(entity)
	v.processor_.PostprocessEntity(
		entity,
		1,
		1,
	)
	// Visit slot 1 between terms.
	v.processor_.ProcessComponentSlot(1)

	var optionalParameters = component.GetOptionalParameters()
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
}

func (v *visitor_) visitCondition(
	condition ast.ConditionLike,
) {
	var expression = condition.GetExpression()
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

func (v *visitor_) visitContinueClause(
	continueClause ast.ContinueClauseLike,
) {
	var delimiter1 = continueClause.GetDelimiter1()
	v.processor_.ProcessDelimiter(delimiter1)
	// Visit slot 1 between terms.
	v.processor_.ProcessContinueClauseSlot(1)

	var delimiter2 = continueClause.GetDelimiter2()
	v.processor_.ProcessDelimiter(delimiter2)
}

func (v *visitor_) visitDiscardClause(
	discardClause ast.DiscardClauseLike,
) {
	var delimiter = discardClause.GetDelimiter()
	v.processor_.ProcessDelimiter(delimiter)
	// Visit slot 1 between terms.
	v.processor_.ProcessDiscardClauseSlot(1)

	var draft = discardClause.GetDraft()
	v.processor_.PreprocessDraft(
		draft,
		1,
		1,
	)
	v.visitDraft(draft)
	v.processor_.PostprocessDraft(
		draft,
		1,
		1,
	)
}

func (v *visitor_) visitDoClause(
	doClause ast.DoClauseLike,
) {
	var delimiter = doClause.GetDelimiter()
	v.processor_.ProcessDelimiter(delimiter)
	// Visit slot 1 between terms.
	v.processor_.ProcessDoClauseSlot(1)

	var invocation = doClause.GetInvocation()
	v.processor_.PreprocessInvocation(
		invocation,
		1,
		1,
	)
	v.visitInvocation(invocation)
	v.processor_.PostprocessInvocation(
		invocation,
		1,
		1,
	)
}

func (v *visitor_) visitDocument(
	document ast.DocumentLike,
) {
	var optionalLegalNotice = document.GetOptionalLegalNotice()
	if uti.IsDefined(optionalLegalNotice) {
		v.processor_.PreprocessLegalNotice(
			optionalLegalNotice,
			1,
			1,
		)
		v.visitLegalNotice(optionalLegalNotice)
		v.processor_.PostprocessLegalNotice(
			optionalLegalNotice,
			1,
			1,
		)
	}
	// Visit slot 1 between terms.
	v.processor_.ProcessDocumentSlot(1)

	var component = document.GetComponent()
	v.processor_.PreprocessComponent(
		component,
		1,
		1,
	)
	v.visitComponent(component)
	v.processor_.PostprocessComponent(
		component,
		1,
		1,
	)
}

func (v *visitor_) visitDraft(
	draft ast.DraftLike,
) {
	var expression = draft.GetExpression()
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

func (v *visitor_) visitElement(
	element ast.ElementLike,
) {
	// Visit the possible element expression types.
	var actual = element.GetAny().(string)
	switch {
	case ScannerClass().MatchesType(actual, AngleToken):
		v.processor_.ProcessAngle(actual)
	case ScannerClass().MatchesType(actual, BooleanToken):
		v.processor_.ProcessBoolean(actual)
	case ScannerClass().MatchesType(actual, CitationToken):
		v.processor_.ProcessCitation(actual)
	case ScannerClass().MatchesType(actual, DurationToken):
		v.processor_.ProcessDuration(actual)
	case ScannerClass().MatchesType(actual, MomentToken):
		v.processor_.ProcessMoment(actual)
	case ScannerClass().MatchesType(actual, NumberToken):
		v.processor_.ProcessNumber(actual)
	case ScannerClass().MatchesType(actual, PatternToken):
		v.processor_.ProcessPattern(actual)
	case ScannerClass().MatchesType(actual, PercentageToken):
		v.processor_.ProcessPercentage(actual)
	case ScannerClass().MatchesType(actual, ProbabilityToken):
		v.processor_.ProcessProbability(actual)
	case ScannerClass().MatchesType(actual, ResourceToken):
		v.processor_.ProcessResource(actual)
	}
}

func (v *visitor_) visitEntity(
	entity ast.EntityLike,
) {
	// Visit the possible entity rule types.
	switch actual := entity.GetAny().(type) {
	case ast.ElementLike:
		v.processor_.PreprocessElement(
			actual,
			1,
			1,
		)
		v.visitElement(actual)
		v.processor_.PostprocessElement(
			actual,
			1,
			1,
		)
	case ast.StringLike:
		v.processor_.PreprocessString(
			actual,
			1,
			1,
		)
		v.visitString(actual)
		v.processor_.PostprocessString(
			actual,
			1,
			1,
		)
	case ast.CollectionLike:
		v.processor_.PreprocessCollection(
			actual,
			1,
			1,
		)
		v.visitCollection(actual)
		v.processor_.PostprocessCollection(
			actual,
			1,
			1,
		)
	case ast.ProcedureLike:
		v.processor_.PreprocessProcedure(
			actual,
			1,
			1,
		)
		v.visitProcedure(actual)
		v.processor_.PostprocessProcedure(
			actual,
			1,
			1,
		)
	}
}

func (v *visitor_) visitEvent(
	event ast.EventLike,
) {
	var expression = event.GetExpression()
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

func (v *visitor_) visitException(
	exception ast.ExceptionLike,
) {
	var expression = exception.GetExpression()
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

func (v *visitor_) visitExclusive(
	exclusive ast.ExclusiveLike,
) {
	var delimiter = exclusive.GetDelimiter()
	v.processor_.ProcessDelimiter(delimiter)
}

func (v *visitor_) visitExclusiveRange(
	exclusiveRange ast.ExclusiveRangeLike,
) {
	var delimiter1 = exclusiveRange.GetDelimiter1()
	v.processor_.ProcessDelimiter(delimiter1)
	// Visit slot 1 between terms.
	v.processor_.ProcessExclusiveRangeSlot(1)

	var primitive1 = exclusiveRange.GetPrimitive1()
	v.processor_.PreprocessPrimitive(
		primitive1,
		1,
		1,
	)
	v.visitPrimitive(primitive1)
	v.processor_.PostprocessPrimitive(
		primitive1,
		1,
		1,
	)
	// Visit slot 2 between terms.
	v.processor_.ProcessExclusiveRangeSlot(2)

	var delimiter2 = exclusiveRange.GetDelimiter2()
	v.processor_.ProcessDelimiter(delimiter2)
	// Visit slot 3 between terms.
	v.processor_.ProcessExclusiveRangeSlot(3)

	var primitive2 = exclusiveRange.GetPrimitive2()
	v.processor_.PreprocessPrimitive(
		primitive2,
		1,
		1,
	)
	v.visitPrimitive(primitive2)
	v.processor_.PostprocessPrimitive(
		primitive2,
		1,
		1,
	)
	// Visit slot 4 between terms.
	v.processor_.ProcessExclusiveRangeSlot(4)

	var bracket = exclusiveRange.GetBracket()
	v.processor_.PreprocessBracket(
		bracket,
		1,
		1,
	)
	v.visitBracket(bracket)
	v.processor_.PostprocessBracket(
		bracket,
		1,
		1,
	)
}

func (v *visitor_) visitExpression(
	expression ast.ExpressionLike,
) {
	var subject = expression.GetSubject()
	v.processor_.PreprocessSubject(
		subject,
		1,
		1,
	)
	v.visitSubject(subject)
	v.processor_.PostprocessSubject(
		subject,
		1,
		1,
	)
	// Visit slot 1 between terms.
	v.processor_.ProcessExpressionSlot(1)

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

func (v *visitor_) visitFailure(
	failure ast.FailureLike,
) {
	var symbol = failure.GetSymbol()
	v.processor_.ProcessSymbol(symbol)
}

func (v *visitor_) visitFlow(
	flow ast.FlowLike,
) {
	// Visit the possible flow rule types.
	switch actual := flow.GetAny().(type) {
	case ast.IfClauseLike:
		v.processor_.PreprocessIfClause(
			actual,
			1,
			1,
		)
		v.visitIfClause(actual)
		v.processor_.PostprocessIfClause(
			actual,
			1,
			1,
		)
	case ast.SelectClauseLike:
		v.processor_.PreprocessSelectClause(
			actual,
			1,
			1,
		)
		v.visitSelectClause(actual)
		v.processor_.PostprocessSelectClause(
			actual,
			1,
			1,
		)
	case ast.WhileClauseLike:
		v.processor_.PreprocessWhileClause(
			actual,
			1,
			1,
		)
		v.visitWhileClause(actual)
		v.processor_.PostprocessWhileClause(
			actual,
			1,
			1,
		)
	case ast.WithClauseLike:
		v.processor_.PreprocessWithClause(
			actual,
			1,
			1,
		)
		v.visitWithClause(actual)
		v.processor_.PostprocessWithClause(
			actual,
			1,
			1,
		)
	case ast.ContinueClauseLike:
		v.processor_.PreprocessContinueClause(
			actual,
			1,
			1,
		)
		v.visitContinueClause(actual)
		v.processor_.PostprocessContinueClause(
			actual,
			1,
			1,
		)
	case ast.BreakClauseLike:
		v.processor_.PreprocessBreakClause(
			actual,
			1,
			1,
		)
		v.visitBreakClause(actual)
		v.processor_.PostprocessBreakClause(
			actual,
			1,
			1,
		)
	case ast.ReturnClauseLike:
		v.processor_.PreprocessReturnClause(
			actual,
			1,
			1,
		)
		v.visitReturnClause(actual)
		v.processor_.PostprocessReturnClause(
			actual,
			1,
			1,
		)
	case ast.ThrowClauseLike:
		v.processor_.PreprocessThrowClause(
			actual,
			1,
			1,
		)
		v.visitThrowClause(actual)
		v.processor_.PostprocessThrowClause(
			actual,
			1,
			1,
		)
	}
}

func (v *visitor_) visitFunction(
	function ast.FunctionLike,
) {
	var identifier = function.GetIdentifier()
	v.processor_.ProcessIdentifier(identifier)
	// Visit slot 1 between terms.
	v.processor_.ProcessFunctionSlot(1)

	var delimiter1 = function.GetDelimiter1()
	v.processor_.ProcessDelimiter(delimiter1)
	// Visit slot 2 between terms.
	v.processor_.ProcessFunctionSlot(2)

	var optionalArguments = function.GetOptionalArguments()
	if uti.IsDefined(optionalArguments) {
		v.processor_.PreprocessArguments(
			optionalArguments,
			1,
			1,
		)
		v.visitArguments(optionalArguments)
		v.processor_.PostprocessArguments(
			optionalArguments,
			1,
			1,
		)
	}
	// Visit slot 3 between terms.
	v.processor_.ProcessFunctionSlot(3)

	var delimiter2 = function.GetDelimiter2()
	v.processor_.ProcessDelimiter(delimiter2)
}

func (v *visitor_) visitIfClause(
	ifClause ast.IfClauseLike,
) {
	var delimiter1 = ifClause.GetDelimiter1()
	v.processor_.ProcessDelimiter(delimiter1)
	// Visit slot 1 between terms.
	v.processor_.ProcessIfClauseSlot(1)

	var condition = ifClause.GetCondition()
	v.processor_.PreprocessCondition(
		condition,
		1,
		1,
	)
	v.visitCondition(condition)
	v.processor_.PostprocessCondition(
		condition,
		1,
		1,
	)
	// Visit slot 2 between terms.
	v.processor_.ProcessIfClauseSlot(2)

	var delimiter2 = ifClause.GetDelimiter2()
	v.processor_.ProcessDelimiter(delimiter2)
	// Visit slot 3 between terms.
	v.processor_.ProcessIfClauseSlot(3)

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

func (v *visitor_) visitInclusive(
	inclusive ast.InclusiveLike,
) {
	var delimiter = inclusive.GetDelimiter()
	v.processor_.ProcessDelimiter(delimiter)
}

func (v *visitor_) visitInclusiveRange(
	inclusiveRange ast.InclusiveRangeLike,
) {
	var delimiter1 = inclusiveRange.GetDelimiter1()
	v.processor_.ProcessDelimiter(delimiter1)
	// Visit slot 1 between terms.
	v.processor_.ProcessInclusiveRangeSlot(1)

	var primitive1 = inclusiveRange.GetPrimitive1()
	v.processor_.PreprocessPrimitive(
		primitive1,
		1,
		1,
	)
	v.visitPrimitive(primitive1)
	v.processor_.PostprocessPrimitive(
		primitive1,
		1,
		1,
	)
	// Visit slot 2 between terms.
	v.processor_.ProcessInclusiveRangeSlot(2)

	var delimiter2 = inclusiveRange.GetDelimiter2()
	v.processor_.ProcessDelimiter(delimiter2)
	// Visit slot 3 between terms.
	v.processor_.ProcessInclusiveRangeSlot(3)

	var primitive2 = inclusiveRange.GetPrimitive2()
	v.processor_.PreprocessPrimitive(
		primitive2,
		1,
		1,
	)
	v.visitPrimitive(primitive2)
	v.processor_.PostprocessPrimitive(
		primitive2,
		1,
		1,
	)
	// Visit slot 4 between terms.
	v.processor_.ProcessInclusiveRangeSlot(4)

	var bracket = inclusiveRange.GetBracket()
	v.processor_.PreprocessBracket(
		bracket,
		1,
		1,
	)
	v.visitBracket(bracket)
	v.processor_.PostprocessBracket(
		bracket,
		1,
		1,
	)
}

func (v *visitor_) visitIndex(
	index ast.IndexLike,
) {
	var expression = index.GetExpression()
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

func (v *visitor_) visitIndices(
	indices ast.IndicesLike,
) {
	var index = indices.GetIndex()
	v.processor_.PreprocessIndex(
		index,
		1,
		1,
	)
	v.visitIndex(index)
	v.processor_.PostprocessIndex(
		index,
		1,
		1,
	)
	// Visit slot 1 between terms.
	v.processor_.ProcessIndicesSlot(1)

	var additionalIndexesIndex uint
	var additionalIndexes = indices.GetAdditionalIndexes().GetIterator()
	var additionalIndexesCount = uint(additionalIndexes.GetSize())
	for additionalIndexes.HasNext() {
		additionalIndexesIndex++
		var rule = additionalIndexes.GetNext()
		v.processor_.PreprocessAdditionalIndex(
			rule,
			additionalIndexesIndex,
			additionalIndexesCount,
		)
		v.visitAdditionalIndex(rule)
		v.processor_.PostprocessAdditionalIndex(
			rule,
			additionalIndexesIndex,
			additionalIndexesCount,
		)
	}
}

func (v *visitor_) visitIndirect(
	indirect ast.IndirectLike,
) {
	// Visit the possible indirect rule types.
	switch actual := indirect.GetAny().(type) {
	case ast.ComponentLike:
		v.processor_.PreprocessComponent(
			actual,
			1,
			1,
		)
		v.visitComponent(actual)
		v.processor_.PostprocessComponent(
			actual,
			1,
			1,
		)
	case ast.SubcomponentLike:
		v.processor_.PreprocessSubcomponent(
			actual,
			1,
			1,
		)
		v.visitSubcomponent(actual)
		v.processor_.PostprocessSubcomponent(
			actual,
			1,
			1,
		)
	case ast.ReferentLike:
		v.processor_.PreprocessReferent(
			actual,
			1,
			1,
		)
		v.visitReferent(actual)
		v.processor_.PostprocessReferent(
			actual,
			1,
			1,
		)
	case ast.FunctionLike:
		v.processor_.PreprocessFunction(
			actual,
			1,
			1,
		)
		v.visitFunction(actual)
		v.processor_.PostprocessFunction(
			actual,
			1,
			1,
		)
	case ast.MethodLike:
		v.processor_.PreprocessMethod(
			actual,
			1,
			1,
		)
		v.visitMethod(actual)
		v.processor_.PostprocessMethod(
			actual,
			1,
			1,
		)
	case ast.VariableLike:
		v.processor_.PreprocessVariable(
			actual,
			1,
			1,
		)
		v.visitVariable(actual)
		v.processor_.PostprocessVariable(
			actual,
			1,
			1,
		)
	}
}

func (v *visitor_) visitInduction(
	induction ast.InductionLike,
) {
	// Visit the possible induction rule types.
	switch actual := induction.GetAny().(type) {
	case ast.DoClauseLike:
		v.processor_.PreprocessDoClause(
			actual,
			1,
			1,
		)
		v.visitDoClause(actual)
		v.processor_.PostprocessDoClause(
			actual,
			1,
			1,
		)
	case ast.LetClauseLike:
		v.processor_.PreprocessLetClause(
			actual,
			1,
			1,
		)
		v.visitLetClause(actual)
		v.processor_.PostprocessLetClause(
			actual,
			1,
			1,
		)
	}
}

func (v *visitor_) visitInlineAttributes(
	inlineAttributes ast.InlineAttributesLike,
) {
	var delimiter1 = inlineAttributes.GetDelimiter1()
	v.processor_.ProcessDelimiter(delimiter1)
	// Visit slot 1 between terms.
	v.processor_.ProcessInlineAttributesSlot(1)

	var association = inlineAttributes.GetAssociation()
	v.processor_.PreprocessAssociation(
		association,
		1,
		1,
	)
	v.visitAssociation(association)
	v.processor_.PostprocessAssociation(
		association,
		1,
		1,
	)
	// Visit slot 2 between terms.
	v.processor_.ProcessInlineAttributesSlot(2)

	var additionalAssociationsIndex uint
	var additionalAssociations = inlineAttributes.GetAdditionalAssociations().GetIterator()
	var additionalAssociationsCount = uint(additionalAssociations.GetSize())
	for additionalAssociations.HasNext() {
		additionalAssociationsIndex++
		var rule = additionalAssociations.GetNext()
		v.processor_.PreprocessAdditionalAssociation(
			rule,
			additionalAssociationsIndex,
			additionalAssociationsCount,
		)
		v.visitAdditionalAssociation(rule)
		v.processor_.PostprocessAdditionalAssociation(
			rule,
			additionalAssociationsIndex,
			additionalAssociationsCount,
		)
	}
	// Visit slot 3 between terms.
	v.processor_.ProcessInlineAttributesSlot(3)

	var delimiter2 = inlineAttributes.GetDelimiter2()
	v.processor_.ProcessDelimiter(delimiter2)
}

func (v *visitor_) visitInlineParameters(
	inlineParameters ast.InlineParametersLike,
) {
	var delimiter1 = inlineParameters.GetDelimiter1()
	v.processor_.ProcessDelimiter(delimiter1)
	// Visit slot 1 between terms.
	v.processor_.ProcessInlineParametersSlot(1)

	var association = inlineParameters.GetAssociation()
	v.processor_.PreprocessAssociation(
		association,
		1,
		1,
	)
	v.visitAssociation(association)
	v.processor_.PostprocessAssociation(
		association,
		1,
		1,
	)
	// Visit slot 2 between terms.
	v.processor_.ProcessInlineParametersSlot(2)

	var additionalAssociationsIndex uint
	var additionalAssociations = inlineParameters.GetAdditionalAssociations().GetIterator()
	var additionalAssociationsCount = uint(additionalAssociations.GetSize())
	for additionalAssociations.HasNext() {
		additionalAssociationsIndex++
		var rule = additionalAssociations.GetNext()
		v.processor_.PreprocessAdditionalAssociation(
			rule,
			additionalAssociationsIndex,
			additionalAssociationsCount,
		)
		v.visitAdditionalAssociation(rule)
		v.processor_.PostprocessAdditionalAssociation(
			rule,
			additionalAssociationsIndex,
			additionalAssociationsCount,
		)
	}
	// Visit slot 3 between terms.
	v.processor_.ProcessInlineParametersSlot(3)

	var delimiter2 = inlineParameters.GetDelimiter2()
	v.processor_.ProcessDelimiter(delimiter2)
}

func (v *visitor_) visitInlineStatements(
	inlineStatements ast.InlineStatementsLike,
) {
	var delimiter1 = inlineStatements.GetDelimiter1()
	v.processor_.ProcessDelimiter(delimiter1)
	// Visit slot 1 between terms.
	v.processor_.ProcessInlineStatementsSlot(1)

	var statement = inlineStatements.GetStatement()
	v.processor_.PreprocessStatement(
		statement,
		1,
		1,
	)
	v.visitStatement(statement)
	v.processor_.PostprocessStatement(
		statement,
		1,
		1,
	)
	// Visit slot 2 between terms.
	v.processor_.ProcessInlineStatementsSlot(2)

	var additionalStatementsIndex uint
	var additionalStatements = inlineStatements.GetAdditionalStatements().GetIterator()
	var additionalStatementsCount = uint(additionalStatements.GetSize())
	for additionalStatements.HasNext() {
		additionalStatementsIndex++
		var rule = additionalStatements.GetNext()
		v.processor_.PreprocessAdditionalStatement(
			rule,
			additionalStatementsIndex,
			additionalStatementsCount,
		)
		v.visitAdditionalStatement(rule)
		v.processor_.PostprocessAdditionalStatement(
			rule,
			additionalStatementsIndex,
			additionalStatementsCount,
		)
	}
	// Visit slot 3 between terms.
	v.processor_.ProcessInlineStatementsSlot(3)

	var delimiter2 = inlineStatements.GetDelimiter2()
	v.processor_.ProcessDelimiter(delimiter2)
}

func (v *visitor_) visitInlineValues(
	inlineValues ast.InlineValuesLike,
) {
	var delimiter1 = inlineValues.GetDelimiter1()
	v.processor_.ProcessDelimiter(delimiter1)
	// Visit slot 1 between terms.
	v.processor_.ProcessInlineValuesSlot(1)

	var component = inlineValues.GetComponent()
	v.processor_.PreprocessComponent(
		component,
		1,
		1,
	)
	v.visitComponent(component)
	v.processor_.PostprocessComponent(
		component,
		1,
		1,
	)
	// Visit slot 2 between terms.
	v.processor_.ProcessInlineValuesSlot(2)

	var additionalValuesIndex uint
	var additionalValues = inlineValues.GetAdditionalValues().GetIterator()
	var additionalValuesCount = uint(additionalValues.GetSize())
	for additionalValues.HasNext() {
		additionalValuesIndex++
		var rule = additionalValues.GetNext()
		v.processor_.PreprocessAdditionalValue(
			rule,
			additionalValuesIndex,
			additionalValuesCount,
		)
		v.visitAdditionalValue(rule)
		v.processor_.PostprocessAdditionalValue(
			rule,
			additionalValuesIndex,
			additionalValuesCount,
		)
	}
	// Visit slot 3 between terms.
	v.processor_.ProcessInlineValuesSlot(3)

	var delimiter2 = inlineValues.GetDelimiter2()
	v.processor_.ProcessDelimiter(delimiter2)
}

func (v *visitor_) visitInverse(
	inverse ast.InverseLike,
) {
	// Visit the possible inverse expression types.
	var actual = inverse.GetAny().(string)
	switch {
	case ScannerClass().MatchesType(actual, MinusToken):
		v.processor_.ProcessMinus(actual)
	case ScannerClass().MatchesType(actual, SlashToken):
		v.processor_.ProcessSlash(actual)
	case ScannerClass().MatchesType(actual, StarToken):
		v.processor_.ProcessStar(actual)
	}
}

func (v *visitor_) visitInversion(
	inversion ast.InversionLike,
) {
	var inverse = inversion.GetInverse()
	v.processor_.PreprocessInverse(
		inverse,
		1,
		1,
	)
	v.visitInverse(inverse)
	v.processor_.PostprocessInverse(
		inverse,
		1,
		1,
	)
	// Visit slot 1 between terms.
	v.processor_.ProcessInversionSlot(1)

	var numerical = inversion.GetNumerical()
	v.processor_.PreprocessNumerical(
		numerical,
		1,
		1,
	)
	v.visitNumerical(numerical)
	v.processor_.PostprocessNumerical(
		numerical,
		1,
		1,
	)
}

func (v *visitor_) visitInvocation(
	invocation ast.InvocationLike,
) {
	// Visit the possible invocation rule types.
	switch actual := invocation.GetAny().(type) {
	case ast.FunctionLike:
		v.processor_.PreprocessFunction(
			actual,
			1,
			1,
		)
		v.visitFunction(actual)
		v.processor_.PostprocessFunction(
			actual,
			1,
			1,
		)
	case ast.MethodLike:
		v.processor_.PreprocessMethod(
			actual,
			1,
			1,
		)
		v.visitMethod(actual)
		v.processor_.PostprocessMethod(
			actual,
			1,
			1,
		)
	}
}

func (v *visitor_) visitItem(
	item ast.ItemLike,
) {
	var symbol = item.GetSymbol()
	v.processor_.ProcessSymbol(symbol)
}

func (v *visitor_) visitLegalNotice(
	legalNotice ast.LegalNoticeLike,
) {
	var comment = legalNotice.GetComment()
	v.processor_.ProcessComment(comment)
}

func (v *visitor_) visitLetClause(
	letClause ast.LetClauseLike,
) {
	var delimiter = letClause.GetDelimiter()
	v.processor_.ProcessDelimiter(delimiter)
	// Visit slot 1 between terms.
	v.processor_.ProcessLetClauseSlot(1)

	var recipient = letClause.GetRecipient()
	v.processor_.PreprocessRecipient(
		recipient,
		1,
		1,
	)
	v.visitRecipient(recipient)
	v.processor_.PostprocessRecipient(
		recipient,
		1,
		1,
	)
	// Visit slot 2 between terms.
	v.processor_.ProcessLetClauseSlot(2)

	var assignment = letClause.GetAssignment()
	v.processor_.PreprocessAssignment(
		assignment,
		1,
		1,
	)
	v.visitAssignment(assignment)
	v.processor_.PostprocessAssignment(
		assignment,
		1,
		1,
	)
	// Visit slot 3 between terms.
	v.processor_.ProcessLetClauseSlot(3)

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

func (v *visitor_) visitLogic(
	logic ast.LogicLike,
) {
	// Visit the possible logic literal values.
	var actual = logic.GetAny().(string)
	switch actual {
	case "and":
		v.processor_.ProcessDelimiter("and")
	case "san":
		v.processor_.ProcessDelimiter("san")
	case "ior":
		v.processor_.ProcessDelimiter("ior")
	case "xor":
		v.processor_.ProcessDelimiter("xor")
	}
}

func (v *visitor_) visitLogical(
	logical ast.LogicalLike,
) {
	// Visit the possible logical rule types.
	switch actual := logical.GetAny().(type) {
	case ast.ComponentLike:
		v.processor_.PreprocessComponent(
			actual,
			1,
			1,
		)
		v.visitComponent(actual)
		v.processor_.PostprocessComponent(
			actual,
			1,
			1,
		)
	case ast.SubcomponentLike:
		v.processor_.PreprocessSubcomponent(
			actual,
			1,
			1,
		)
		v.visitSubcomponent(actual)
		v.processor_.PostprocessSubcomponent(
			actual,
			1,
			1,
		)
	case ast.PrecedenceLike:
		v.processor_.PreprocessPrecedence(
			actual,
			1,
			1,
		)
		v.visitPrecedence(actual)
		v.processor_.PostprocessPrecedence(
			actual,
			1,
			1,
		)
	case ast.ReferentLike:
		v.processor_.PreprocessReferent(
			actual,
			1,
			1,
		)
		v.visitReferent(actual)
		v.processor_.PostprocessReferent(
			actual,
			1,
			1,
		)
	case ast.ComplementLike:
		v.processor_.PreprocessComplement(
			actual,
			1,
			1,
		)
		v.visitComplement(actual)
		v.processor_.PostprocessComplement(
			actual,
			1,
			1,
		)
	case ast.FunctionLike:
		v.processor_.PreprocessFunction(
			actual,
			1,
			1,
		)
		v.visitFunction(actual)
		v.processor_.PostprocessFunction(
			actual,
			1,
			1,
		)
	case ast.MethodLike:
		v.processor_.PreprocessMethod(
			actual,
			1,
			1,
		)
		v.visitMethod(actual)
		v.processor_.PostprocessMethod(
			actual,
			1,
			1,
		)
	case ast.VariableLike:
		v.processor_.PreprocessVariable(
			actual,
			1,
			1,
		)
		v.visitVariable(actual)
		v.processor_.PostprocessVariable(
			actual,
			1,
			1,
		)
	}
}

func (v *visitor_) visitMagnitude(
	magnitude ast.MagnitudeLike,
) {
	var delimiter1 = magnitude.GetDelimiter1()
	v.processor_.ProcessDelimiter(delimiter1)
	// Visit slot 1 between terms.
	v.processor_.ProcessMagnitudeSlot(1)

	var numerical = magnitude.GetNumerical()
	v.processor_.PreprocessNumerical(
		numerical,
		1,
		1,
	)
	v.visitNumerical(numerical)
	v.processor_.PostprocessNumerical(
		numerical,
		1,
		1,
	)
	// Visit slot 2 between terms.
	v.processor_.ProcessMagnitudeSlot(2)

	var delimiter2 = magnitude.GetDelimiter2()
	v.processor_.ProcessDelimiter(delimiter2)
}

func (v *visitor_) visitMainClause(
	mainClause ast.MainClauseLike,
) {
	// Visit the possible mainClause rule types.
	switch actual := mainClause.GetAny().(type) {
	case ast.FlowLike:
		v.processor_.PreprocessFlow(
			actual,
			1,
			1,
		)
		v.visitFlow(actual)
		v.processor_.PostprocessFlow(
			actual,
			1,
			1,
		)
	case ast.InductionLike:
		v.processor_.PreprocessInduction(
			actual,
			1,
			1,
		)
		v.visitInduction(actual)
		v.processor_.PostprocessInduction(
			actual,
			1,
			1,
		)
	case ast.MessagingLike:
		v.processor_.PreprocessMessaging(
			actual,
			1,
			1,
		)
		v.visitMessaging(actual)
		v.processor_.PostprocessMessaging(
			actual,
			1,
			1,
		)
	case ast.RepositoryLike:
		v.processor_.PreprocessRepository(
			actual,
			1,
			1,
		)
		v.visitRepository(actual)
		v.processor_.PostprocessRepository(
			actual,
			1,
			1,
		)
	}
}

func (v *visitor_) visitMatchHandler(
	matchHandler ast.MatchHandlerLike,
) {
	var delimiter1 = matchHandler.GetDelimiter1()
	v.processor_.ProcessDelimiter(delimiter1)
	// Visit slot 1 between terms.
	v.processor_.ProcessMatchHandlerSlot(1)

	var template = matchHandler.GetTemplate()
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
	v.processor_.ProcessMatchHandlerSlot(2)

	var delimiter2 = matchHandler.GetDelimiter2()
	v.processor_.ProcessDelimiter(delimiter2)
	// Visit slot 3 between terms.
	v.processor_.ProcessMatchHandlerSlot(3)

	var procedure = matchHandler.GetProcedure()
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

func (v *visitor_) visitMessage(
	message ast.MessageLike,
) {
	var expression = message.GetExpression()
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

func (v *visitor_) visitMessaging(
	messaging ast.MessagingLike,
) {
	// Visit the possible messaging rule types.
	switch actual := messaging.GetAny().(type) {
	case ast.PostClauseLike:
		v.processor_.PreprocessPostClause(
			actual,
			1,
			1,
		)
		v.visitPostClause(actual)
		v.processor_.PostprocessPostClause(
			actual,
			1,
			1,
		)
	case ast.RetrieveClauseLike:
		v.processor_.PreprocessRetrieveClause(
			actual,
			1,
			1,
		)
		v.visitRetrieveClause(actual)
		v.processor_.PostprocessRetrieveClause(
			actual,
			1,
			1,
		)
	case ast.AcceptClauseLike:
		v.processor_.PreprocessAcceptClause(
			actual,
			1,
			1,
		)
		v.visitAcceptClause(actual)
		v.processor_.PostprocessAcceptClause(
			actual,
			1,
			1,
		)
	case ast.RejectClauseLike:
		v.processor_.PreprocessRejectClause(
			actual,
			1,
			1,
		)
		v.visitRejectClause(actual)
		v.processor_.PostprocessRejectClause(
			actual,
			1,
			1,
		)
	case ast.PublishClauseLike:
		v.processor_.PreprocessPublishClause(
			actual,
			1,
			1,
		)
		v.visitPublishClause(actual)
		v.processor_.PostprocessPublishClause(
			actual,
			1,
			1,
		)
	}
}

func (v *visitor_) visitMethod(
	method ast.MethodLike,
) {
	var identifier1 = method.GetIdentifier1()
	v.processor_.ProcessIdentifier(identifier1)
	// Visit slot 1 between terms.
	v.processor_.ProcessMethodSlot(1)

	var blocking = method.GetBlocking()
	v.processor_.PreprocessBlocking(
		blocking,
		1,
		1,
	)
	v.visitBlocking(blocking)
	v.processor_.PostprocessBlocking(
		blocking,
		1,
		1,
	)
	// Visit slot 2 between terms.
	v.processor_.ProcessMethodSlot(2)

	var identifier2 = method.GetIdentifier2()
	v.processor_.ProcessIdentifier(identifier2)
	// Visit slot 3 between terms.
	v.processor_.ProcessMethodSlot(3)

	var delimiter1 = method.GetDelimiter1()
	v.processor_.ProcessDelimiter(delimiter1)
	// Visit slot 4 between terms.
	v.processor_.ProcessMethodSlot(4)

	var optionalArguments = method.GetOptionalArguments()
	if uti.IsDefined(optionalArguments) {
		v.processor_.PreprocessArguments(
			optionalArguments,
			1,
			1,
		)
		v.visitArguments(optionalArguments)
		v.processor_.PostprocessArguments(
			optionalArguments,
			1,
			1,
		)
	}
	// Visit slot 5 between terms.
	v.processor_.ProcessMethodSlot(5)

	var delimiter2 = method.GetDelimiter2()
	v.processor_.ProcessDelimiter(delimiter2)
}

func (v *visitor_) visitMultilineAttributes(
	multilineAttributes ast.MultilineAttributesLike,
) {
	var delimiter1 = multilineAttributes.GetDelimiter1()
	v.processor_.ProcessDelimiter(delimiter1)
	// Visit slot 1 between terms.
	v.processor_.ProcessMultilineAttributesSlot(1)

	var newline = multilineAttributes.GetNewline()
	v.processor_.ProcessNewline(newline)
	// Visit slot 2 between terms.
	v.processor_.ProcessMultilineAttributesSlot(2)

	var annotatedAssociationsIndex uint
	var annotatedAssociations = multilineAttributes.GetAnnotatedAssociations().GetIterator()
	var annotatedAssociationsCount = uint(annotatedAssociations.GetSize())
	for annotatedAssociations.HasNext() {
		annotatedAssociationsIndex++
		var rule = annotatedAssociations.GetNext()
		v.processor_.PreprocessAnnotatedAssociation(
			rule,
			annotatedAssociationsIndex,
			annotatedAssociationsCount,
		)
		v.visitAnnotatedAssociation(rule)
		v.processor_.PostprocessAnnotatedAssociation(
			rule,
			annotatedAssociationsIndex,
			annotatedAssociationsCount,
		)
	}
	// Visit slot 3 between terms.
	v.processor_.ProcessMultilineAttributesSlot(3)

	var delimiter2 = multilineAttributes.GetDelimiter2()
	v.processor_.ProcessDelimiter(delimiter2)
}

func (v *visitor_) visitMultilineParameters(
	multilineParameters ast.MultilineParametersLike,
) {
	var delimiter1 = multilineParameters.GetDelimiter1()
	v.processor_.ProcessDelimiter(delimiter1)
	// Visit slot 1 between terms.
	v.processor_.ProcessMultilineParametersSlot(1)

	var newline = multilineParameters.GetNewline()
	v.processor_.ProcessNewline(newline)
	// Visit slot 2 between terms.
	v.processor_.ProcessMultilineParametersSlot(2)

	var annotatedAssociationsIndex uint
	var annotatedAssociations = multilineParameters.GetAnnotatedAssociations().GetIterator()
	var annotatedAssociationsCount = uint(annotatedAssociations.GetSize())
	for annotatedAssociations.HasNext() {
		annotatedAssociationsIndex++
		var rule = annotatedAssociations.GetNext()
		v.processor_.PreprocessAnnotatedAssociation(
			rule,
			annotatedAssociationsIndex,
			annotatedAssociationsCount,
		)
		v.visitAnnotatedAssociation(rule)
		v.processor_.PostprocessAnnotatedAssociation(
			rule,
			annotatedAssociationsIndex,
			annotatedAssociationsCount,
		)
	}
	// Visit slot 3 between terms.
	v.processor_.ProcessMultilineParametersSlot(3)

	var delimiter2 = multilineParameters.GetDelimiter2()
	v.processor_.ProcessDelimiter(delimiter2)
}

func (v *visitor_) visitMultilineStatements(
	multilineStatements ast.MultilineStatementsLike,
) {
	var delimiter1 = multilineStatements.GetDelimiter1()
	v.processor_.ProcessDelimiter(delimiter1)
	// Visit slot 1 between terms.
	v.processor_.ProcessMultilineStatementsSlot(1)

	var newline = multilineStatements.GetNewline()
	v.processor_.ProcessNewline(newline)
	// Visit slot 2 between terms.
	v.processor_.ProcessMultilineStatementsSlot(2)

	var annotatedStatementsIndex uint
	var annotatedStatements = multilineStatements.GetAnnotatedStatements().GetIterator()
	var annotatedStatementsCount = uint(annotatedStatements.GetSize())
	for annotatedStatements.HasNext() {
		annotatedStatementsIndex++
		var rule = annotatedStatements.GetNext()
		v.processor_.PreprocessAnnotatedStatement(
			rule,
			annotatedStatementsIndex,
			annotatedStatementsCount,
		)
		v.visitAnnotatedStatement(rule)
		v.processor_.PostprocessAnnotatedStatement(
			rule,
			annotatedStatementsIndex,
			annotatedStatementsCount,
		)
	}
	// Visit slot 3 between terms.
	v.processor_.ProcessMultilineStatementsSlot(3)

	var delimiter2 = multilineStatements.GetDelimiter2()
	v.processor_.ProcessDelimiter(delimiter2)
}

func (v *visitor_) visitMultilineValues(
	multilineValues ast.MultilineValuesLike,
) {
	var delimiter1 = multilineValues.GetDelimiter1()
	v.processor_.ProcessDelimiter(delimiter1)
	// Visit slot 1 between terms.
	v.processor_.ProcessMultilineValuesSlot(1)

	var newline = multilineValues.GetNewline()
	v.processor_.ProcessNewline(newline)
	// Visit slot 2 between terms.
	v.processor_.ProcessMultilineValuesSlot(2)

	var annotatedValuesIndex uint
	var annotatedValues = multilineValues.GetAnnotatedValues().GetIterator()
	var annotatedValuesCount = uint(annotatedValues.GetSize())
	for annotatedValues.HasNext() {
		annotatedValuesIndex++
		var rule = annotatedValues.GetNext()
		v.processor_.PreprocessAnnotatedValue(
			rule,
			annotatedValuesIndex,
			annotatedValuesCount,
		)
		v.visitAnnotatedValue(rule)
		v.processor_.PostprocessAnnotatedValue(
			rule,
			annotatedValuesIndex,
			annotatedValuesCount,
		)
	}
	// Visit slot 3 between terms.
	v.processor_.ProcessMultilineValuesSlot(3)

	var delimiter2 = multilineValues.GetDelimiter2()
	v.processor_.ProcessDelimiter(delimiter2)
}

func (v *visitor_) visitNoAttributes(
	noAttributes ast.NoAttributesLike,
) {
	var delimiter1 = noAttributes.GetDelimiter1()
	v.processor_.ProcessDelimiter(delimiter1)
	// Visit slot 1 between terms.
	v.processor_.ProcessNoAttributesSlot(1)

	var delimiter2 = noAttributes.GetDelimiter2()
	v.processor_.ProcessDelimiter(delimiter2)
	// Visit slot 2 between terms.
	v.processor_.ProcessNoAttributesSlot(2)

	var delimiter3 = noAttributes.GetDelimiter3()
	v.processor_.ProcessDelimiter(delimiter3)
}

func (v *visitor_) visitNoStatements(
	noStatements ast.NoStatementsLike,
) {
	var delimiter1 = noStatements.GetDelimiter1()
	v.processor_.ProcessDelimiter(delimiter1)
	// Visit slot 1 between terms.
	v.processor_.ProcessNoStatementsSlot(1)

	var delimiter2 = noStatements.GetDelimiter2()
	v.processor_.ProcessDelimiter(delimiter2)
}

func (v *visitor_) visitNoValues(
	noValues ast.NoValuesLike,
) {
	var delimiter1 = noValues.GetDelimiter1()
	v.processor_.ProcessDelimiter(delimiter1)
	// Visit slot 1 between terms.
	v.processor_.ProcessNoValuesSlot(1)

	var delimiter2 = noValues.GetDelimiter2()
	v.processor_.ProcessDelimiter(delimiter2)
}

func (v *visitor_) visitNotarizeClause(
	notarizeClause ast.NotarizeClauseLike,
) {
	var delimiter1 = notarizeClause.GetDelimiter1()
	v.processor_.ProcessDelimiter(delimiter1)
	// Visit slot 1 between terms.
	v.processor_.ProcessNotarizeClauseSlot(1)

	var draft = notarizeClause.GetDraft()
	v.processor_.PreprocessDraft(
		draft,
		1,
		1,
	)
	v.visitDraft(draft)
	v.processor_.PostprocessDraft(
		draft,
		1,
		1,
	)
	// Visit slot 2 between terms.
	v.processor_.ProcessNotarizeClauseSlot(2)

	var delimiter2 = notarizeClause.GetDelimiter2()
	v.processor_.ProcessDelimiter(delimiter2)
	// Visit slot 3 between terms.
	v.processor_.ProcessNotarizeClauseSlot(3)

	var cited = notarizeClause.GetCited()
	v.processor_.PreprocessCited(
		cited,
		1,
		1,
	)
	v.visitCited(cited)
	v.processor_.PostprocessCited(
		cited,
		1,
		1,
	)
}

func (v *visitor_) visitNumerical(
	numerical ast.NumericalLike,
) {
	// Visit the possible numerical rule types.
	switch actual := numerical.GetAny().(type) {
	case ast.ComponentLike:
		v.processor_.PreprocessComponent(
			actual,
			1,
			1,
		)
		v.visitComponent(actual)
		v.processor_.PostprocessComponent(
			actual,
			1,
			1,
		)
	case ast.SubcomponentLike:
		v.processor_.PreprocessSubcomponent(
			actual,
			1,
			1,
		)
		v.visitSubcomponent(actual)
		v.processor_.PostprocessSubcomponent(
			actual,
			1,
			1,
		)
	case ast.PrecedenceLike:
		v.processor_.PreprocessPrecedence(
			actual,
			1,
			1,
		)
		v.visitPrecedence(actual)
		v.processor_.PostprocessPrecedence(
			actual,
			1,
			1,
		)
	case ast.ReferentLike:
		v.processor_.PreprocessReferent(
			actual,
			1,
			1,
		)
		v.visitReferent(actual)
		v.processor_.PostprocessReferent(
			actual,
			1,
			1,
		)
	case ast.InversionLike:
		v.processor_.PreprocessInversion(
			actual,
			1,
			1,
		)
		v.visitInversion(actual)
		v.processor_.PostprocessInversion(
			actual,
			1,
			1,
		)
	case ast.MagnitudeLike:
		v.processor_.PreprocessMagnitude(
			actual,
			1,
			1,
		)
		v.visitMagnitude(actual)
		v.processor_.PostprocessMagnitude(
			actual,
			1,
			1,
		)
	case ast.FunctionLike:
		v.processor_.PreprocessFunction(
			actual,
			1,
			1,
		)
		v.visitFunction(actual)
		v.processor_.PostprocessFunction(
			actual,
			1,
			1,
		)
	case ast.MethodLike:
		v.processor_.PreprocessMethod(
			actual,
			1,
			1,
		)
		v.visitMethod(actual)
		v.processor_.PostprocessMethod(
			actual,
			1,
			1,
		)
	case ast.VariableLike:
		v.processor_.PreprocessVariable(
			actual,
			1,
			1,
		)
		v.visitVariable(actual)
		v.processor_.PostprocessVariable(
			actual,
			1,
			1,
		)
	}
}

func (v *visitor_) visitOnClause(
	onClause ast.OnClauseLike,
) {
	var delimiter = onClause.GetDelimiter()
	v.processor_.ProcessDelimiter(delimiter)
	// Visit slot 1 between terms.
	v.processor_.ProcessOnClauseSlot(1)

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
	v.processor_.ProcessOnClauseSlot(2)

	var matchHandlersIndex uint
	var matchHandlers = onClause.GetMatchHandlers().GetIterator()
	var matchHandlersCount = uint(matchHandlers.GetSize())
	for matchHandlers.HasNext() {
		matchHandlersIndex++
		var rule = matchHandlers.GetNext()
		v.processor_.PreprocessMatchHandler(
			rule,
			matchHandlersIndex,
			matchHandlersCount,
		)
		v.visitMatchHandler(rule)
		v.processor_.PostprocessMatchHandler(
			rule,
			matchHandlersIndex,
			matchHandlersCount,
		)
	}
}

func (v *visitor_) visitOperation(
	operation ast.OperationLike,
) {
	// Visit the possible operation rule types.
	switch actual := operation.GetAny().(type) {
	case ast.TextualLike:
		v.processor_.PreprocessTextual(
			actual,
			1,
			1,
		)
		v.visitTextual(actual)
		v.processor_.PostprocessTextual(
			actual,
			1,
			1,
		)
	case ast.LogicLike:
		v.processor_.PreprocessLogic(
			actual,
			1,
			1,
		)
		v.visitLogic(actual)
		v.processor_.PostprocessLogic(
			actual,
			1,
			1,
		)
	case ast.ArithmeticLike:
		v.processor_.PreprocessArithmetic(
			actual,
			1,
			1,
		)
		v.visitArithmetic(actual)
		v.processor_.PostprocessArithmetic(
			actual,
			1,
			1,
		)
	case ast.ComparisonLike:
		v.processor_.PreprocessComparison(
			actual,
			1,
			1,
		)
		v.visitComparison(actual)
		v.processor_.PostprocessComparison(
			actual,
			1,
			1,
		)
	}
}

func (v *visitor_) visitParameters(
	parameters ast.ParametersLike,
) {
	// Visit the possible parameters rule types.
	switch actual := parameters.GetAny().(type) {
	case ast.MultilineParametersLike:
		v.processor_.PreprocessMultilineParameters(
			actual,
			1,
			1,
		)
		v.visitMultilineParameters(actual)
		v.processor_.PostprocessMultilineParameters(
			actual,
			1,
			1,
		)
	case ast.InlineParametersLike:
		v.processor_.PreprocessInlineParameters(
			actual,
			1,
			1,
		)
		v.visitInlineParameters(actual)
		v.processor_.PostprocessInlineParameters(
			actual,
			1,
			1,
		)
	}
}

func (v *visitor_) visitPostClause(
	postClause ast.PostClauseLike,
) {
	var delimiter1 = postClause.GetDelimiter1()
	v.processor_.ProcessDelimiter(delimiter1)
	// Visit slot 1 between terms.
	v.processor_.ProcessPostClauseSlot(1)

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
	v.processor_.ProcessPostClauseSlot(2)

	var delimiter2 = postClause.GetDelimiter2()
	v.processor_.ProcessDelimiter(delimiter2)
	// Visit slot 3 between terms.
	v.processor_.ProcessPostClauseSlot(3)

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
	precedence ast.PrecedenceLike,
) {
	var delimiter1 = precedence.GetDelimiter1()
	v.processor_.ProcessDelimiter(delimiter1)
	// Visit slot 1 between terms.
	v.processor_.ProcessPrecedenceSlot(1)

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
	v.processor_.ProcessPrecedenceSlot(2)

	var delimiter2 = precedence.GetDelimiter2()
	v.processor_.ProcessDelimiter(delimiter2)
}

func (v *visitor_) visitPredicate(
	predicate ast.PredicateLike,
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
	v.processor_.ProcessPredicateSlot(1)

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
	primitive ast.PrimitiveLike,
) {
	// Visit the possible primitive rule types.
	switch actual := primitive.GetAny().(type) {
	case ast.ElementLike:
		v.processor_.PreprocessElement(
			actual,
			1,
			1,
		)
		v.visitElement(actual)
		v.processor_.PostprocessElement(
			actual,
			1,
			1,
		)
	case ast.StringLike:
		v.processor_.PreprocessString(
			actual,
			1,
			1,
		)
		v.visitString(actual)
		v.processor_.PostprocessString(
			actual,
			1,
			1,
		)
	}
}

func (v *visitor_) visitProcedure(
	procedure ast.ProcedureLike,
) {
	// Visit the possible procedure rule types.
	switch actual := procedure.GetAny().(type) {
	case ast.MultilineStatementsLike:
		v.processor_.PreprocessMultilineStatements(
			actual,
			1,
			1,
		)
		v.visitMultilineStatements(actual)
		v.processor_.PostprocessMultilineStatements(
			actual,
			1,
			1,
		)
	case ast.InlineStatementsLike:
		v.processor_.PreprocessInlineStatements(
			actual,
			1,
			1,
		)
		v.visitInlineStatements(actual)
		v.processor_.PostprocessInlineStatements(
			actual,
			1,
			1,
		)
	case ast.NoStatementsLike:
		v.processor_.PreprocessNoStatements(
			actual,
			1,
			1,
		)
		v.visitNoStatements(actual)
		v.processor_.PostprocessNoStatements(
			actual,
			1,
			1,
		)
	}
}

func (v *visitor_) visitPublishClause(
	publishClause ast.PublishClauseLike,
) {
	var delimiter = publishClause.GetDelimiter()
	v.processor_.ProcessDelimiter(delimiter)
	// Visit slot 1 between terms.
	v.processor_.ProcessPublishClauseSlot(1)

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

func (v *visitor_) visitRecipient(
	recipient ast.RecipientLike,
) {
	// Visit the possible recipient rule types.
	switch actual := recipient.GetAny().(type) {
	case ast.ItemLike:
		v.processor_.PreprocessItem(
			actual,
			1,
			1,
		)
		v.visitItem(actual)
		v.processor_.PostprocessItem(
			actual,
			1,
			1,
		)
	case ast.SubcomponentLike:
		v.processor_.PreprocessSubcomponent(
			actual,
			1,
			1,
		)
		v.visitSubcomponent(actual)
		v.processor_.PostprocessSubcomponent(
			actual,
			1,
			1,
		)
	}
}

func (v *visitor_) visitReferent(
	referent ast.ReferentLike,
) {
	var delimiter = referent.GetDelimiter()
	v.processor_.ProcessDelimiter(delimiter)
	// Visit slot 1 between terms.
	v.processor_.ProcessReferentSlot(1)

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
	rejectClause ast.RejectClauseLike,
) {
	var delimiter = rejectClause.GetDelimiter()
	v.processor_.ProcessDelimiter(delimiter)
	// Visit slot 1 between terms.
	v.processor_.ProcessRejectClauseSlot(1)

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

func (v *visitor_) visitRepository(
	repository ast.RepositoryLike,
) {
	// Visit the possible repository rule types.
	switch actual := repository.GetAny().(type) {
	case ast.CheckoutClauseLike:
		v.processor_.PreprocessCheckoutClause(
			actual,
			1,
			1,
		)
		v.visitCheckoutClause(actual)
		v.processor_.PostprocessCheckoutClause(
			actual,
			1,
			1,
		)
	case ast.SaveClauseLike:
		v.processor_.PreprocessSaveClause(
			actual,
			1,
			1,
		)
		v.visitSaveClause(actual)
		v.processor_.PostprocessSaveClause(
			actual,
			1,
			1,
		)
	case ast.DiscardClauseLike:
		v.processor_.PreprocessDiscardClause(
			actual,
			1,
			1,
		)
		v.visitDiscardClause(actual)
		v.processor_.PostprocessDiscardClause(
			actual,
			1,
			1,
		)
	case ast.NotarizeClauseLike:
		v.processor_.PreprocessNotarizeClause(
			actual,
			1,
			1,
		)
		v.visitNotarizeClause(actual)
		v.processor_.PostprocessNotarizeClause(
			actual,
			1,
			1,
		)
	}
}

func (v *visitor_) visitResult(
	result ast.ResultLike,
) {
	var expression = result.GetExpression()
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

func (v *visitor_) visitRetrieveClause(
	retrieveClause ast.RetrieveClauseLike,
) {
	var delimiter1 = retrieveClause.GetDelimiter1()
	v.processor_.ProcessDelimiter(delimiter1)
	// Visit slot 1 between terms.
	v.processor_.ProcessRetrieveClauseSlot(1)

	var recipient = retrieveClause.GetRecipient()
	v.processor_.PreprocessRecipient(
		recipient,
		1,
		1,
	)
	v.visitRecipient(recipient)
	v.processor_.PostprocessRecipient(
		recipient,
		1,
		1,
	)
	// Visit slot 2 between terms.
	v.processor_.ProcessRetrieveClauseSlot(2)

	var delimiter2 = retrieveClause.GetDelimiter2()
	v.processor_.ProcessDelimiter(delimiter2)
	// Visit slot 3 between terms.
	v.processor_.ProcessRetrieveClauseSlot(3)

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
	returnClause ast.ReturnClauseLike,
) {
	var delimiter = returnClause.GetDelimiter()
	v.processor_.ProcessDelimiter(delimiter)
	// Visit slot 1 between terms.
	v.processor_.ProcessReturnClauseSlot(1)

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
	saveClause ast.SaveClauseLike,
) {
	var delimiter1 = saveClause.GetDelimiter1()
	v.processor_.ProcessDelimiter(delimiter1)
	// Visit slot 1 between terms.
	v.processor_.ProcessSaveClauseSlot(1)

	var draft = saveClause.GetDraft()
	v.processor_.PreprocessDraft(
		draft,
		1,
		1,
	)
	v.visitDraft(draft)
	v.processor_.PostprocessDraft(
		draft,
		1,
		1,
	)
	// Visit slot 2 between terms.
	v.processor_.ProcessSaveClauseSlot(2)

	var delimiter2 = saveClause.GetDelimiter2()
	v.processor_.ProcessDelimiter(delimiter2)
	// Visit slot 3 between terms.
	v.processor_.ProcessSaveClauseSlot(3)

	var cited = saveClause.GetCited()
	v.processor_.PreprocessCited(
		cited,
		1,
		1,
	)
	v.visitCited(cited)
	v.processor_.PostprocessCited(
		cited,
		1,
		1,
	)
}

func (v *visitor_) visitSelectClause(
	selectClause ast.SelectClauseLike,
) {
	var delimiter = selectClause.GetDelimiter()
	v.processor_.ProcessDelimiter(delimiter)
	// Visit slot 1 between terms.
	v.processor_.ProcessSelectClauseSlot(1)

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
	v.processor_.ProcessSelectClauseSlot(2)

	var matchHandlersIndex uint
	var matchHandlers = selectClause.GetMatchHandlers().GetIterator()
	var matchHandlersCount = uint(matchHandlers.GetSize())
	for matchHandlers.HasNext() {
		matchHandlersIndex++
		var rule = matchHandlers.GetNext()
		v.processor_.PreprocessMatchHandler(
			rule,
			matchHandlersIndex,
			matchHandlersCount,
		)
		v.visitMatchHandler(rule)
		v.processor_.PostprocessMatchHandler(
			rule,
			matchHandlersIndex,
			matchHandlersCount,
		)
	}
}

func (v *visitor_) visitSequence(
	sequence ast.SequenceLike,
) {
	var expression = sequence.GetExpression()
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

func (v *visitor_) visitStatement(
	statement ast.StatementLike,
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
	v.processor_.ProcessStatementSlot(1)

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

func (v *visitor_) visitStatementLine(
	statementLine ast.StatementLineLike,
) {
	var statement = statementLine.GetStatement()
	v.processor_.PreprocessStatement(
		statement,
		1,
		1,
	)
	v.visitStatement(statement)
	v.processor_.PostprocessStatement(
		statement,
		1,
		1,
	)
	// Visit slot 1 between terms.
	v.processor_.ProcessStatementLineSlot(1)

	var optionalNote = statementLine.GetOptionalNote()
	if uti.IsDefined(optionalNote) {
		v.processor_.ProcessNote(optionalNote)
	}
}

func (v *visitor_) visitString(
	string_ ast.StringLike,
) {
	// Visit the possible string expression types.
	var actual = string_.GetAny().(string)
	switch {
	case ScannerClass().MatchesType(actual, BinaryToken):
		v.processor_.ProcessBinary(actual)
	case ScannerClass().MatchesType(actual, BytecodeToken):
		v.processor_.ProcessBytecode(actual)
	case ScannerClass().MatchesType(actual, NameToken):
		v.processor_.ProcessName(actual)
	case ScannerClass().MatchesType(actual, NarrativeToken):
		v.processor_.ProcessNarrative(actual)
	case ScannerClass().MatchesType(actual, QuoteToken):
		v.processor_.ProcessQuote(actual)
	case ScannerClass().MatchesType(actual, SymbolToken):
		v.processor_.ProcessSymbol(actual)
	case ScannerClass().MatchesType(actual, TagToken):
		v.processor_.ProcessTag(actual)
	case ScannerClass().MatchesType(actual, VersionToken):
		v.processor_.ProcessVersion(actual)
	}
}

func (v *visitor_) visitSubcomponent(
	subcomponent ast.SubcomponentLike,
) {
	var identifier = subcomponent.GetIdentifier()
	v.processor_.ProcessIdentifier(identifier)
	// Visit slot 1 between terms.
	v.processor_.ProcessSubcomponentSlot(1)

	var delimiter1 = subcomponent.GetDelimiter1()
	v.processor_.ProcessDelimiter(delimiter1)
	// Visit slot 2 between terms.
	v.processor_.ProcessSubcomponentSlot(2)

	var indices = subcomponent.GetIndices()
	v.processor_.PreprocessIndices(
		indices,
		1,
		1,
	)
	v.visitIndices(indices)
	v.processor_.PostprocessIndices(
		indices,
		1,
		1,
	)
	// Visit slot 3 between terms.
	v.processor_.ProcessSubcomponentSlot(3)

	var delimiter2 = subcomponent.GetDelimiter2()
	v.processor_.ProcessDelimiter(delimiter2)
}

func (v *visitor_) visitSubject(
	subject ast.SubjectLike,
) {
	// Visit the possible subject rule types.
	switch actual := subject.GetAny().(type) {
	case ast.ComponentLike:
		v.processor_.PreprocessComponent(
			actual,
			1,
			1,
		)
		v.visitComponent(actual)
		v.processor_.PostprocessComponent(
			actual,
			1,
			1,
		)
	case ast.SubcomponentLike:
		v.processor_.PreprocessSubcomponent(
			actual,
			1,
			1,
		)
		v.visitSubcomponent(actual)
		v.processor_.PostprocessSubcomponent(
			actual,
			1,
			1,
		)
	case ast.PrecedenceLike:
		v.processor_.PreprocessPrecedence(
			actual,
			1,
			1,
		)
		v.visitPrecedence(actual)
		v.processor_.PostprocessPrecedence(
			actual,
			1,
			1,
		)
	case ast.ReferentLike:
		v.processor_.PreprocessReferent(
			actual,
			1,
			1,
		)
		v.visitReferent(actual)
		v.processor_.PostprocessReferent(
			actual,
			1,
			1,
		)
	case ast.ComplementLike:
		v.processor_.PreprocessComplement(
			actual,
			1,
			1,
		)
		v.visitComplement(actual)
		v.processor_.PostprocessComplement(
			actual,
			1,
			1,
		)
	case ast.InversionLike:
		v.processor_.PreprocessInversion(
			actual,
			1,
			1,
		)
		v.visitInversion(actual)
		v.processor_.PostprocessInversion(
			actual,
			1,
			1,
		)
	case ast.MagnitudeLike:
		v.processor_.PreprocessMagnitude(
			actual,
			1,
			1,
		)
		v.visitMagnitude(actual)
		v.processor_.PostprocessMagnitude(
			actual,
			1,
			1,
		)
	case ast.FunctionLike:
		v.processor_.PreprocessFunction(
			actual,
			1,
			1,
		)
		v.visitFunction(actual)
		v.processor_.PostprocessFunction(
			actual,
			1,
			1,
		)
	case ast.MethodLike:
		v.processor_.PreprocessMethod(
			actual,
			1,
			1,
		)
		v.visitMethod(actual)
		v.processor_.PostprocessMethod(
			actual,
			1,
			1,
		)
	case ast.VariableLike:
		v.processor_.PreprocessVariable(
			actual,
			1,
			1,
		)
		v.visitVariable(actual)
		v.processor_.PostprocessVariable(
			actual,
			1,
			1,
		)
	}
}

func (v *visitor_) visitTarget(
	target ast.TargetLike,
) {
	// Visit the possible target rule types.
	switch actual := target.GetAny().(type) {
	case ast.FunctionLike:
		v.processor_.PreprocessFunction(
			actual,
			1,
			1,
		)
		v.visitFunction(actual)
		v.processor_.PostprocessFunction(
			actual,
			1,
			1,
		)
	case ast.MethodLike:
		v.processor_.PreprocessMethod(
			actual,
			1,
			1,
		)
		v.visitMethod(actual)
		v.processor_.PostprocessMethod(
			actual,
			1,
			1,
		)
	case ast.SubcomponentLike:
		v.processor_.PreprocessSubcomponent(
			actual,
			1,
			1,
		)
		v.visitSubcomponent(actual)
		v.processor_.PostprocessSubcomponent(
			actual,
			1,
			1,
		)
	case ast.VariableLike:
		v.processor_.PreprocessVariable(
			actual,
			1,
			1,
		)
		v.visitVariable(actual)
		v.processor_.PostprocessVariable(
			actual,
			1,
			1,
		)
	}
}

func (v *visitor_) visitTemplate(
	template ast.TemplateLike,
) {
	var expression = template.GetExpression()
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

func (v *visitor_) visitTextual(
	textual ast.TextualLike,
) {
	// Visit the possible textual literal values.
	var actual = textual.GetAny().(string)
	switch actual {
	case "&":
		v.processor_.ProcessDelimiter("&")
	case "is":
		v.processor_.ProcessDelimiter("is")
	case "matches":
		v.processor_.ProcessDelimiter("matches")
	}
}

func (v *visitor_) visitThrowClause(
	throwClause ast.ThrowClauseLike,
) {
	var delimiter = throwClause.GetDelimiter()
	v.processor_.ProcessDelimiter(delimiter)
	// Visit slot 1 between terms.
	v.processor_.ProcessThrowClauseSlot(1)

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

func (v *visitor_) visitVariable(
	variable ast.VariableLike,
) {
	var identifier = variable.GetIdentifier()
	v.processor_.ProcessIdentifier(identifier)
}

func (v *visitor_) visitWhileClause(
	whileClause ast.WhileClauseLike,
) {
	var delimiter1 = whileClause.GetDelimiter1()
	v.processor_.ProcessDelimiter(delimiter1)
	// Visit slot 1 between terms.
	v.processor_.ProcessWhileClauseSlot(1)

	var condition = whileClause.GetCondition()
	v.processor_.PreprocessCondition(
		condition,
		1,
		1,
	)
	v.visitCondition(condition)
	v.processor_.PostprocessCondition(
		condition,
		1,
		1,
	)
	// Visit slot 2 between terms.
	v.processor_.ProcessWhileClauseSlot(2)

	var delimiter2 = whileClause.GetDelimiter2()
	v.processor_.ProcessDelimiter(delimiter2)
	// Visit slot 3 between terms.
	v.processor_.ProcessWhileClauseSlot(3)

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
	withClause ast.WithClauseLike,
) {
	var delimiter1 = withClause.GetDelimiter1()
	v.processor_.ProcessDelimiter(delimiter1)
	// Visit slot 1 between terms.
	v.processor_.ProcessWithClauseSlot(1)

	var delimiter2 = withClause.GetDelimiter2()
	v.processor_.ProcessDelimiter(delimiter2)
	// Visit slot 2 between terms.
	v.processor_.ProcessWithClauseSlot(2)

	var item = withClause.GetItem()
	v.processor_.PreprocessItem(
		item,
		1,
		1,
	)
	v.visitItem(item)
	v.processor_.PostprocessItem(
		item,
		1,
		1,
	)
	// Visit slot 3 between terms.
	v.processor_.ProcessWithClauseSlot(3)

	var delimiter3 = withClause.GetDelimiter3()
	v.processor_.ProcessDelimiter(delimiter3)
	// Visit slot 4 between terms.
	v.processor_.ProcessWithClauseSlot(4)

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
	v.processor_.ProcessWithClauseSlot(5)

	var delimiter4 = withClause.GetDelimiter4()
	v.processor_.ProcessDelimiter(delimiter4)
	// Visit slot 6 between terms.
	v.processor_.ProcessWithClauseSlot(6)

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
