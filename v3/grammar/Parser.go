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
	fmt "fmt"
	ast "github.com/bali-nebula/go-bali-documents/v3/ast"
	col "github.com/craterdog/go-collection-framework/v7"
	uti "github.com/craterdog/go-missing-utilities/v7"
	mat "math"
	sts "strings"
)

// CLASS INTERFACE

// Access Function

func ParserClass() ParserClassLike {
	return parserClass()
}

// Constructor Methods

func (c *parserClass_) Parser() ParserLike {
	var instance = &parser_{
		// Initialize the instance attributes.
	}
	return instance
}

// INSTANCE INTERFACE

// Principal Methods

func (v *parser_) GetClass() ParserClassLike {
	return parserClass()
}

func (v *parser_) ParseSource(
	source string,
) ast.DocumentLike {
	v.source_ = sts.ReplaceAll(source, "\t", "    ")
	v.tokens_ = col.Queue[TokenLike]()
	v.next_ = col.Stack[TokenLike]()

	// The scanner runs in a separate Go routine.
	ScannerClass().Scanner(v.source_, v.tokens_)

	// Attempt to parse the document.
	var document, token, ok = v.parseDocument()
	if !ok || !v.tokens_.IsEmpty() {
		var message = v.formatError("$Document", token)
		panic(message)
	}
	return document
}

// PROTECTED INTERFACE

// Private Methods

func (v *parser_) parseAcceptClause() (
	acceptClause ast.AcceptClauseLike,
	token TokenLike,
	ok bool,
) {
	var tokens = col.List[TokenLike]()

	// Attempt to parse a single "accept" delimiter.
	_, token, ok = v.parseDelimiter("accept")
	if !ok {
		if uti.IsDefined(tokens) {
			// This is not a single AcceptClause rule.
			v.putBack(tokens)
			return
		} else {
			// Found a syntax error.
			var message = v.formatError("$AcceptClause", token)
			panic(message)
		}
	}
	if uti.IsDefined(tokens) {
		tokens.AppendValue(token)
	}

	// Attempt to parse a single Message rule.
	var message ast.MessageLike
	message, token, ok = v.parseMessage()
	switch {
	case ok:
		// No additional put backs allowed at this point.
		tokens = nil
	case uti.IsDefined(tokens):
		// This is not a single Message rule.
		v.putBack(tokens)
		return
	default:
		// Found a syntax error.
		var message = v.formatError("$AcceptClause", token)
		panic(message)
	}

	// Found a single AcceptClause rule.
	ok = true
	v.remove(tokens)
	acceptClause = ast.AcceptClauseClass().AcceptClause(message)
	return
}

func (v *parser_) parseAction() (
	action ast.ActionLike,
	token TokenLike,
	ok bool,
) {
	// Attempt to parse a single Operator Action.
	var operator ast.OperatorLike
	operator, token, ok = v.parseOperator()
	if ok {
		// Found a single Operator Action.
		action = ast.ActionClass().Action(operator)
		return
	}

	// Attempt to parse a single Operation Action.
	var operation ast.OperationLike
	operation, token, ok = v.parseOperation()
	if ok {
		// Found a single Operation Action.
		action = ast.ActionClass().Action(operation)
		return
	}

	// This is not a single Action rule.
	return
}

func (v *parser_) parseAdditionalArgument() (
	additionalArgument ast.AdditionalArgumentLike,
	token TokenLike,
	ok bool,
) {
	var tokens = col.List[TokenLike]()

	// Attempt to parse a single "," delimiter.
	_, token, ok = v.parseDelimiter(",")
	if !ok {
		if uti.IsDefined(tokens) {
			// This is not a single AdditionalArgument rule.
			v.putBack(tokens)
			return
		} else {
			// Found a syntax error.
			var message = v.formatError("$AdditionalArgument", token)
			panic(message)
		}
	}
	if uti.IsDefined(tokens) {
		tokens.AppendValue(token)
	}

	// Attempt to parse a single Argument rule.
	var argument ast.ArgumentLike
	argument, token, ok = v.parseArgument()
	switch {
	case ok:
		// No additional put backs allowed at this point.
		tokens = nil
	case uti.IsDefined(tokens):
		// This is not a single Argument rule.
		v.putBack(tokens)
		return
	default:
		// Found a syntax error.
		var message = v.formatError("$AdditionalArgument", token)
		panic(message)
	}

	// Found a single AdditionalArgument rule.
	ok = true
	v.remove(tokens)
	additionalArgument = ast.AdditionalArgumentClass().AdditionalArgument(argument)
	return
}

func (v *parser_) parseAdditionalAssociation() (
	additionalAssociation ast.AdditionalAssociationLike,
	token TokenLike,
	ok bool,
) {
	var tokens = col.List[TokenLike]()

	// Attempt to parse a single "," delimiter.
	_, token, ok = v.parseDelimiter(",")
	if !ok {
		if uti.IsDefined(tokens) {
			// This is not a single AdditionalAssociation rule.
			v.putBack(tokens)
			return
		} else {
			// Found a syntax error.
			var message = v.formatError("$AdditionalAssociation", token)
			panic(message)
		}
	}
	if uti.IsDefined(tokens) {
		tokens.AppendValue(token)
	}

	// Attempt to parse a single Association rule.
	var association ast.AssociationLike
	association, token, ok = v.parseAssociation()
	switch {
	case ok:
		// No additional put backs allowed at this point.
		tokens = nil
	case uti.IsDefined(tokens):
		// This is not a single Association rule.
		v.putBack(tokens)
		return
	default:
		// Found a syntax error.
		var message = v.formatError("$AdditionalAssociation", token)
		panic(message)
	}

	// Found a single AdditionalAssociation rule.
	ok = true
	v.remove(tokens)
	additionalAssociation = ast.AdditionalAssociationClass().AdditionalAssociation(association)
	return
}

func (v *parser_) parseAdditionalIndex() (
	additionalIndex ast.AdditionalIndexLike,
	token TokenLike,
	ok bool,
) {
	var tokens = col.List[TokenLike]()

	// Attempt to parse a single "," delimiter.
	_, token, ok = v.parseDelimiter(",")
	if !ok {
		if uti.IsDefined(tokens) {
			// This is not a single AdditionalIndex rule.
			v.putBack(tokens)
			return
		} else {
			// Found a syntax error.
			var message = v.formatError("$AdditionalIndex", token)
			panic(message)
		}
	}
	if uti.IsDefined(tokens) {
		tokens.AppendValue(token)
	}

	// Attempt to parse a single Index rule.
	var index ast.IndexLike
	index, token, ok = v.parseIndex()
	switch {
	case ok:
		// No additional put backs allowed at this point.
		tokens = nil
	case uti.IsDefined(tokens):
		// This is not a single Index rule.
		v.putBack(tokens)
		return
	default:
		// Found a syntax error.
		var message = v.formatError("$AdditionalIndex", token)
		panic(message)
	}

	// Found a single AdditionalIndex rule.
	ok = true
	v.remove(tokens)
	additionalIndex = ast.AdditionalIndexClass().AdditionalIndex(index)
	return
}

func (v *parser_) parseAdditionalStatement() (
	additionalStatement ast.AdditionalStatementLike,
	token TokenLike,
	ok bool,
) {
	var tokens = col.List[TokenLike]()

	// Attempt to parse a single ";" delimiter.
	_, token, ok = v.parseDelimiter(";")
	if !ok {
		if uti.IsDefined(tokens) {
			// This is not a single AdditionalStatement rule.
			v.putBack(tokens)
			return
		} else {
			// Found a syntax error.
			var message = v.formatError("$AdditionalStatement", token)
			panic(message)
		}
	}
	if uti.IsDefined(tokens) {
		tokens.AppendValue(token)
	}

	// Attempt to parse a single Statement rule.
	var statement ast.StatementLike
	statement, token, ok = v.parseStatement()
	switch {
	case ok:
		// No additional put backs allowed at this point.
		tokens = nil
	case uti.IsDefined(tokens):
		// This is not a single Statement rule.
		v.putBack(tokens)
		return
	default:
		// Found a syntax error.
		var message = v.formatError("$AdditionalStatement", token)
		panic(message)
	}

	// Found a single AdditionalStatement rule.
	ok = true
	v.remove(tokens)
	additionalStatement = ast.AdditionalStatementClass().AdditionalStatement(statement)
	return
}

func (v *parser_) parseAdditionalValue() (
	additionalValue ast.AdditionalValueLike,
	token TokenLike,
	ok bool,
) {
	var tokens = col.List[TokenLike]()

	// Attempt to parse a single "," delimiter.
	_, token, ok = v.parseDelimiter(",")
	if !ok {
		if uti.IsDefined(tokens) {
			// This is not a single AdditionalValue rule.
			v.putBack(tokens)
			return
		} else {
			// Found a syntax error.
			var message = v.formatError("$AdditionalValue", token)
			panic(message)
		}
	}
	if uti.IsDefined(tokens) {
		tokens.AppendValue(token)
	}

	// Attempt to parse a single Component rule.
	var component ast.ComponentLike
	component, token, ok = v.parseComponent()
	switch {
	case ok:
		// No additional put backs allowed at this point.
		tokens = nil
	case uti.IsDefined(tokens):
		// This is not a single Component rule.
		v.putBack(tokens)
		return
	default:
		// Found a syntax error.
		var message = v.formatError("$AdditionalValue", token)
		panic(message)
	}

	// Found a single AdditionalValue rule.
	ok = true
	v.remove(tokens)
	additionalValue = ast.AdditionalValueClass().AdditionalValue(component)
	return
}

func (v *parser_) parseAnd() (
	and ast.AndLike,
	token TokenLike,
	ok bool,
) {
	var tokens = col.List[TokenLike]()

	// Attempt to parse a single "and" delimiter.
	_, token, ok = v.parseDelimiter("and")
	if !ok {
		if uti.IsDefined(tokens) {
			// This is not a single And rule.
			v.putBack(tokens)
			return
		} else {
			// Found a syntax error.
			var message = v.formatError("$And", token)
			panic(message)
		}
	}
	if uti.IsDefined(tokens) {
		tokens.AppendValue(token)
	}

	// Found a single And rule.
	ok = true
	v.remove(tokens)
	and = ast.AndClass().And()
	return
}

func (v *parser_) parseAnnotatedAssociation() (
	annotatedAssociation ast.AnnotatedAssociationLike,
	token TokenLike,
	ok bool,
) {
	var tokens = col.List[TokenLike]()

	// Attempt to parse a single Association rule.
	var association ast.AssociationLike
	association, token, ok = v.parseAssociation()
	switch {
	case ok:
		// No additional put backs allowed at this point.
		tokens = nil
	case uti.IsDefined(tokens):
		// This is not a single Association rule.
		v.putBack(tokens)
		return
	default:
		// Found a syntax error.
		var message = v.formatError("$AnnotatedAssociation", token)
		panic(message)
	}

	// Attempt to parse an optional note token.
	var optionalNote string
	optionalNote, token, ok = v.parseToken(NoteToken)
	if ok && uti.IsDefined(tokens) {
		tokens.AppendValue(token)
	}

	// Found a single AnnotatedAssociation rule.
	ok = true
	v.remove(tokens)
	annotatedAssociation = ast.AnnotatedAssociationClass().AnnotatedAssociation(
		association,
		optionalNote,
	)
	return
}

func (v *parser_) parseAnnotatedStatement() (
	annotatedStatement ast.AnnotatedStatementLike,
	token TokenLike,
	ok bool,
) {
	// Attempt to parse a single CommentLine AnnotatedStatement.
	var commentLine ast.CommentLineLike
	commentLine, token, ok = v.parseCommentLine()
	if ok {
		// Found a single CommentLine AnnotatedStatement.
		annotatedStatement = ast.AnnotatedStatementClass().AnnotatedStatement(commentLine)
		return
	}

	// Attempt to parse a single StatementLine AnnotatedStatement.
	var statementLine ast.StatementLineLike
	statementLine, token, ok = v.parseStatementLine()
	if ok {
		// Found a single StatementLine AnnotatedStatement.
		annotatedStatement = ast.AnnotatedStatementClass().AnnotatedStatement(statementLine)
		return
	}

	// This is not a single AnnotatedStatement rule.
	return
}

func (v *parser_) parseAnnotatedValue() (
	annotatedValue ast.AnnotatedValueLike,
	token TokenLike,
	ok bool,
) {
	var tokens = col.List[TokenLike]()

	// Attempt to parse a single Component rule.
	var component ast.ComponentLike
	component, token, ok = v.parseComponent()
	switch {
	case ok:
		// No additional put backs allowed at this point.
		tokens = nil
	case uti.IsDefined(tokens):
		// This is not a single Component rule.
		v.putBack(tokens)
		return
	default:
		// Found a syntax error.
		var message = v.formatError("$AnnotatedValue", token)
		panic(message)
	}

	// Attempt to parse an optional note token.
	var optionalNote string
	optionalNote, token, ok = v.parseToken(NoteToken)
	if ok && uti.IsDefined(tokens) {
		tokens.AppendValue(token)
	}

	// Found a single AnnotatedValue rule.
	ok = true
	v.remove(tokens)
	annotatedValue = ast.AnnotatedValueClass().AnnotatedValue(
		component,
		optionalNote,
	)
	return
}

func (v *parser_) parseArgument() (
	argument ast.ArgumentLike,
	token TokenLike,
	ok bool,
) {
	var tokens = col.List[TokenLike]()

	// Attempt to parse a single identifier token.
	var identifier string
	identifier, token, ok = v.parseToken(IdentifierToken)
	if !ok {
		if uti.IsDefined(tokens) {
			// This is not a single identifier token.
			v.putBack(tokens)
			return
		} else {
			// Found a syntax error.
			var message = v.formatError("$Argument", token)
			panic(message)
		}
	}
	if uti.IsDefined(tokens) {
		tokens.AppendValue(token)
	}

	// Found a single Argument rule.
	ok = true
	v.remove(tokens)
	argument = ast.ArgumentClass().Argument(identifier)
	return
}

func (v *parser_) parseArguments() (
	arguments ast.ArgumentsLike,
	token TokenLike,
	ok bool,
) {
	var tokens = col.List[TokenLike]()

	// Attempt to parse a single Argument rule.
	var argument ast.ArgumentLike
	argument, token, ok = v.parseArgument()
	switch {
	case ok:
		// No additional put backs allowed at this point.
		tokens = nil
	case uti.IsDefined(tokens):
		// This is not a single Argument rule.
		v.putBack(tokens)
		return
	default:
		// Found a syntax error.
		var message = v.formatError("$Arguments", token)
		panic(message)
	}

	// Attempt to parse multiple AdditionalArgument rules.
	var additionalArguments = col.List[ast.AdditionalArgumentLike]()
additionalArgumentsLoop:
	for count := 0; count < mat.MaxInt; count++ {
		var additionalArgument ast.AdditionalArgumentLike
		additionalArgument, token, ok = v.parseAdditionalArgument()
		if !ok {
			switch {
			case count >= 0:
				break additionalArgumentsLoop
			case uti.IsDefined(tokens):
				// This is not multiple AdditionalArgument rules.
				v.putBack(tokens)
				return
			default:
				// Found a syntax error.
				var message = v.formatError("$Arguments", token)
				message += "0 or more AdditionalArgument rules are required."
				panic(message)
			}
		}
		// No additional put backs allowed at this point.
		tokens = nil
		additionalArguments.AppendValue(additionalArgument)
	}

	// Found a single Arguments rule.
	ok = true
	v.remove(tokens)
	arguments = ast.ArgumentsClass().Arguments(
		argument,
		additionalArguments,
	)
	return
}

func (v *parser_) parseAssignment() (
	assignment ast.AssignmentLike,
	token TokenLike,
	ok bool,
) {
	// Attempt to parse a single ColonEqual Assignment.
	var colonEqual ast.ColonEqualLike
	colonEqual, token, ok = v.parseColonEqual()
	if ok {
		// Found a single ColonEqual Assignment.
		assignment = ast.AssignmentClass().Assignment(colonEqual)
		return
	}

	// Attempt to parse a single DefaultEqual Assignment.
	var defaultEqual ast.DefaultEqualLike
	defaultEqual, token, ok = v.parseDefaultEqual()
	if ok {
		// Found a single DefaultEqual Assignment.
		assignment = ast.AssignmentClass().Assignment(defaultEqual)
		return
	}

	// Attempt to parse a single PlusEqual Assignment.
	var plusEqual ast.PlusEqualLike
	plusEqual, token, ok = v.parsePlusEqual()
	if ok {
		// Found a single PlusEqual Assignment.
		assignment = ast.AssignmentClass().Assignment(plusEqual)
		return
	}

	// Attempt to parse a single DashEqual Assignment.
	var dashEqual ast.DashEqualLike
	dashEqual, token, ok = v.parseDashEqual()
	if ok {
		// Found a single DashEqual Assignment.
		assignment = ast.AssignmentClass().Assignment(dashEqual)
		return
	}

	// Attempt to parse a single StarEqual Assignment.
	var starEqual ast.StarEqualLike
	starEqual, token, ok = v.parseStarEqual()
	if ok {
		// Found a single StarEqual Assignment.
		assignment = ast.AssignmentClass().Assignment(starEqual)
		return
	}

	// Attempt to parse a single SlashEqual Assignment.
	var slashEqual ast.SlashEqualLike
	slashEqual, token, ok = v.parseSlashEqual()
	if ok {
		// Found a single SlashEqual Assignment.
		assignment = ast.AssignmentClass().Assignment(slashEqual)
		return
	}

	// This is not a single Assignment rule.
	return
}

func (v *parser_) parseAssociation() (
	association ast.AssociationLike,
	token TokenLike,
	ok bool,
) {
	var tokens = col.List[TokenLike]()

	// Attempt to parse a single Primitive rule.
	var primitive ast.PrimitiveLike
	primitive, token, ok = v.parsePrimitive()
	switch {
	case ok:
		// Found a multiexpression token.
		if uti.IsDefined(tokens) {
			tokens.AppendValue(token)
		}
	case uti.IsDefined(tokens):
		// This is not a single Primitive rule.
		v.putBack(tokens)
		return
	default:
		// Found a syntax error.
		var message = v.formatError("$Association", token)
		panic(message)
	}

	// Attempt to parse a single ":" delimiter.
	_, token, ok = v.parseDelimiter(":")
	if !ok {
		if uti.IsDefined(tokens) {
			// This is not a single Association rule.
			v.putBack(tokens)
			return
		} else {
			// Found a syntax error.
			var message = v.formatError("$Association", token)
			panic(message)
		}
	}
	if uti.IsDefined(tokens) {
		tokens.AppendValue(token)
	}

	// Attempt to parse a single Component rule.
	var component ast.ComponentLike
	component, token, ok = v.parseComponent()
	switch {
	case ok:
		// No additional put backs allowed at this point.
		tokens = nil
	case uti.IsDefined(tokens):
		// This is not a single Component rule.
		v.putBack(tokens)
		return
	default:
		// Found a syntax error.
		var message = v.formatError("$Association", token)
		panic(message)
	}

	// Found a single Association rule.
	ok = true
	v.remove(tokens)
	association = ast.AssociationClass().Association(
		primitive,
		component,
	)
	return
}

func (v *parser_) parseAtLevel() (
	atLevel ast.AtLevelLike,
	token TokenLike,
	ok bool,
) {
	var tokens = col.List[TokenLike]()

	// Attempt to parse a single "at" delimiter.
	_, token, ok = v.parseDelimiter("at")
	if !ok {
		if uti.IsDefined(tokens) {
			// This is not a single AtLevel rule.
			v.putBack(tokens)
			return
		} else {
			// Found a syntax error.
			var message = v.formatError("$AtLevel", token)
			panic(message)
		}
	}
	if uti.IsDefined(tokens) {
		tokens.AppendValue(token)
	}

	// Attempt to parse a single "level" delimiter.
	_, token, ok = v.parseDelimiter("level")
	if !ok {
		if uti.IsDefined(tokens) {
			// This is not a single AtLevel rule.
			v.putBack(tokens)
			return
		} else {
			// Found a syntax error.
			var message = v.formatError("$AtLevel", token)
			panic(message)
		}
	}
	if uti.IsDefined(tokens) {
		tokens.AppendValue(token)
	}

	// Attempt to parse a single Expression rule.
	var expression ast.ExpressionLike
	expression, token, ok = v.parseExpression()
	switch {
	case ok:
		// No additional put backs allowed at this point.
		tokens = nil
	case uti.IsDefined(tokens):
		// This is not a single Expression rule.
		v.putBack(tokens)
		return
	default:
		// Found a syntax error.
		var message = v.formatError("$AtLevel", token)
		panic(message)
	}

	// Found a single AtLevel rule.
	ok = true
	v.remove(tokens)
	atLevel = ast.AtLevelClass().AtLevel(expression)
	return
}

func (v *parser_) parseBag() (
	bag ast.BagLike,
	token TokenLike,
	ok bool,
) {
	var tokens = col.List[TokenLike]()

	// Attempt to parse a single Expression rule.
	var expression ast.ExpressionLike
	expression, token, ok = v.parseExpression()
	switch {
	case ok:
		// No additional put backs allowed at this point.
		tokens = nil
	case uti.IsDefined(tokens):
		// This is not a single Expression rule.
		v.putBack(tokens)
		return
	default:
		// Found a syntax error.
		var message = v.formatError("$Bag", token)
		panic(message)
	}

	// Found a single Bag rule.
	ok = true
	v.remove(tokens)
	bag = ast.BagClass().Bag(expression)
	return
}

func (v *parser_) parseBlocking() (
	blocking ast.BlockingLike,
	token TokenLike,
	ok bool,
) {
	// Attempt to parse a single dot Blocking.
	var dot string
	dot, token, ok = v.parseToken(DotToken)
	if ok {
		// Found a single dot Blocking.
		blocking = ast.BlockingClass().Blocking(dot)
		return
	}

	// Attempt to parse a single arrow Blocking.
	var arrow string
	arrow, token, ok = v.parseToken(ArrowToken)
	if ok {
		// Found a single arrow Blocking.
		blocking = ast.BlockingClass().Blocking(arrow)
		return
	}

	// This is not a single Blocking rule.
	return
}

func (v *parser_) parseBracket() (
	bracket ast.BracketLike,
	token TokenLike,
	ok bool,
) {
	// Attempt to parse a single Inclusive Bracket.
	var inclusive ast.InclusiveLike
	inclusive, token, ok = v.parseInclusive()
	if ok {
		// Found a single Inclusive Bracket.
		bracket = ast.BracketClass().Bracket(inclusive)
		return
	}

	// Attempt to parse a single Exclusive Bracket.
	var exclusive ast.ExclusiveLike
	exclusive, token, ok = v.parseExclusive()
	if ok {
		// Found a single Exclusive Bracket.
		bracket = ast.BracketClass().Bracket(exclusive)
		return
	}

	// This is not a single Bracket rule.
	return
}

func (v *parser_) parseBreakClause() (
	breakClause ast.BreakClauseLike,
	token TokenLike,
	ok bool,
) {
	var tokens = col.List[TokenLike]()

	// Attempt to parse a single "break" delimiter.
	_, token, ok = v.parseDelimiter("break")
	if !ok {
		if uti.IsDefined(tokens) {
			// This is not a single BreakClause rule.
			v.putBack(tokens)
			return
		} else {
			// Found a syntax error.
			var message = v.formatError("$BreakClause", token)
			panic(message)
		}
	}
	if uti.IsDefined(tokens) {
		tokens.AppendValue(token)
	}

	// Attempt to parse a single "loop" delimiter.
	_, token, ok = v.parseDelimiter("loop")
	if !ok {
		if uti.IsDefined(tokens) {
			// This is not a single BreakClause rule.
			v.putBack(tokens)
			return
		} else {
			// Found a syntax error.
			var message = v.formatError("$BreakClause", token)
			panic(message)
		}
	}
	if uti.IsDefined(tokens) {
		tokens.AppendValue(token)
	}

	// Found a single BreakClause rule.
	ok = true
	v.remove(tokens)
	breakClause = ast.BreakClauseClass().BreakClause()
	return
}

func (v *parser_) parseCheckoutClause() (
	checkoutClause ast.CheckoutClauseLike,
	token TokenLike,
	ok bool,
) {
	var tokens = col.List[TokenLike]()

	// Attempt to parse a single "checkout" delimiter.
	_, token, ok = v.parseDelimiter("checkout")
	if !ok {
		if uti.IsDefined(tokens) {
			// This is not a single CheckoutClause rule.
			v.putBack(tokens)
			return
		} else {
			// Found a syntax error.
			var message = v.formatError("$CheckoutClause", token)
			panic(message)
		}
	}
	if uti.IsDefined(tokens) {
		tokens.AppendValue(token)
	}

	// Attempt to parse a single Recipient rule.
	var recipient ast.RecipientLike
	recipient, token, ok = v.parseRecipient()
	switch {
	case ok:
		// No additional put backs allowed at this point.
		tokens = nil
	case uti.IsDefined(tokens):
		// This is not a single Recipient rule.
		v.putBack(tokens)
		return
	default:
		// Found a syntax error.
		var message = v.formatError("$CheckoutClause", token)
		panic(message)
	}

	// Attempt to parse an optional AtLevel rule.
	var optionalAtLevel ast.AtLevelLike
	optionalAtLevel, _, ok = v.parseAtLevel()
	if ok {
		// No additional put backs allowed at this point.
		tokens = nil
	}

	// Attempt to parse a single "from" delimiter.
	_, token, ok = v.parseDelimiter("from")
	if !ok {
		if uti.IsDefined(tokens) {
			// This is not a single CheckoutClause rule.
			v.putBack(tokens)
			return
		} else {
			// Found a syntax error.
			var message = v.formatError("$CheckoutClause", token)
			panic(message)
		}
	}
	if uti.IsDefined(tokens) {
		tokens.AppendValue(token)
	}

	// Attempt to parse a single Cited rule.
	var cited ast.CitedLike
	cited, token, ok = v.parseCited()
	switch {
	case ok:
		// No additional put backs allowed at this point.
		tokens = nil
	case uti.IsDefined(tokens):
		// This is not a single Cited rule.
		v.putBack(tokens)
		return
	default:
		// Found a syntax error.
		var message = v.formatError("$CheckoutClause", token)
		panic(message)
	}

	// Found a single CheckoutClause rule.
	ok = true
	v.remove(tokens)
	checkoutClause = ast.CheckoutClauseClass().CheckoutClause(
		recipient,
		optionalAtLevel,
		cited,
	)
	return
}

func (v *parser_) parseCited() (
	cited ast.CitedLike,
	token TokenLike,
	ok bool,
) {
	var tokens = col.List[TokenLike]()

	// Attempt to parse a single Expression rule.
	var expression ast.ExpressionLike
	expression, token, ok = v.parseExpression()
	switch {
	case ok:
		// No additional put backs allowed at this point.
		tokens = nil
	case uti.IsDefined(tokens):
		// This is not a single Expression rule.
		v.putBack(tokens)
		return
	default:
		// Found a syntax error.
		var message = v.formatError("$Cited", token)
		panic(message)
	}

	// Found a single Cited rule.
	ok = true
	v.remove(tokens)
	cited = ast.CitedClass().Cited(expression)
	return
}

func (v *parser_) parseCollection() (
	collection ast.CollectionLike,
	token TokenLike,
	ok bool,
) {
	// Attempt to parse a single MultilineAttributes Collection.
	var multilineAttributes ast.MultilineAttributesLike
	multilineAttributes, token, ok = v.parseMultilineAttributes()
	if ok {
		// Found a single MultilineAttributes Collection.
		collection = ast.CollectionClass().Collection(multilineAttributes)
		return
	}

	// Attempt to parse a single MultilineValues Collection.
	var multilineValues ast.MultilineValuesLike
	multilineValues, token, ok = v.parseMultilineValues()
	if ok {
		// Found a single MultilineValues Collection.
		collection = ast.CollectionClass().Collection(multilineValues)
		return
	}

	// Attempt to parse a single InclusiveRange Collection.
	var inclusiveRange ast.InclusiveRangeLike
	inclusiveRange, token, ok = v.parseInclusiveRange()
	if ok {
		// Found a single InclusiveRange Collection.
		collection = ast.CollectionClass().Collection(inclusiveRange)
		return
	}

	// Attempt to parse a single ExclusiveRange Collection.
	var exclusiveRange ast.ExclusiveRangeLike
	exclusiveRange, token, ok = v.parseExclusiveRange()
	if ok {
		// Found a single ExclusiveRange Collection.
		collection = ast.CollectionClass().Collection(exclusiveRange)
		return
	}

	// Attempt to parse a single InlineAttributes Collection.
	var inlineAttributes ast.InlineAttributesLike
	inlineAttributes, token, ok = v.parseInlineAttributes()
	if ok {
		// Found a single InlineAttributes Collection.
		collection = ast.CollectionClass().Collection(inlineAttributes)
		return
	}

	// Attempt to parse a single InlineValues Collection.
	var inlineValues ast.InlineValuesLike
	inlineValues, token, ok = v.parseInlineValues()
	if ok {
		// Found a single InlineValues Collection.
		collection = ast.CollectionClass().Collection(inlineValues)
		return
	}

	// Attempt to parse a single NoAttributes Collection.
	var noAttributes ast.NoAttributesLike
	noAttributes, token, ok = v.parseNoAttributes()
	if ok {
		// Found a single NoAttributes Collection.
		collection = ast.CollectionClass().Collection(noAttributes)
		return
	}

	// Attempt to parse a single NoValues Collection.
	var noValues ast.NoValuesLike
	noValues, token, ok = v.parseNoValues()
	if ok {
		// Found a single NoValues Collection.
		collection = ast.CollectionClass().Collection(noValues)
		return
	}

	// This is not a single Collection rule.
	return
}

func (v *parser_) parseColonEqual() (
	colonEqual ast.ColonEqualLike,
	token TokenLike,
	ok bool,
) {
	var tokens = col.List[TokenLike]()

	// Attempt to parse a single ":=" delimiter.
	_, token, ok = v.parseDelimiter(":=")
	if !ok {
		if uti.IsDefined(tokens) {
			// This is not a single ColonEqual rule.
			v.putBack(tokens)
			return
		} else {
			// Found a syntax error.
			var message = v.formatError("$ColonEqual", token)
			panic(message)
		}
	}
	if uti.IsDefined(tokens) {
		tokens.AppendValue(token)
	}

	// Found a single ColonEqual rule.
	ok = true
	v.remove(tokens)
	colonEqual = ast.ColonEqualClass().ColonEqual()
	return
}

func (v *parser_) parseCommentLine() (
	commentLine ast.CommentLineLike,
	token TokenLike,
	ok bool,
) {
	var tokens = col.List[TokenLike]()

	// Attempt to parse a single comment token.
	var comment string
	comment, token, ok = v.parseToken(CommentToken)
	if !ok {
		if uti.IsDefined(tokens) {
			// This is not a single comment token.
			v.putBack(tokens)
			return
		} else {
			// Found a syntax error.
			var message = v.formatError("$CommentLine", token)
			panic(message)
		}
	}
	if uti.IsDefined(tokens) {
		tokens.AppendValue(token)
	}

	// Found a single CommentLine rule.
	ok = true
	v.remove(tokens)
	commentLine = ast.CommentLineClass().CommentLine(comment)
	return
}

func (v *parser_) parseComplement() (
	complement ast.ComplementLike,
	token TokenLike,
	ok bool,
) {
	var tokens = col.List[TokenLike]()

	// Attempt to parse a single "not" delimiter.
	_, token, ok = v.parseDelimiter("not")
	if !ok {
		if uti.IsDefined(tokens) {
			// This is not a single Complement rule.
			v.putBack(tokens)
			return
		} else {
			// Found a syntax error.
			var message = v.formatError("$Complement", token)
			panic(message)
		}
	}
	if uti.IsDefined(tokens) {
		tokens.AppendValue(token)
	}

	// Attempt to parse a single Logical rule.
	var logical ast.LogicalLike
	logical, token, ok = v.parseLogical()
	switch {
	case ok:
		// No additional put backs allowed at this point.
		tokens = nil
	case uti.IsDefined(tokens):
		// This is not a single Logical rule.
		v.putBack(tokens)
		return
	default:
		// Found a syntax error.
		var message = v.formatError("$Complement", token)
		panic(message)
	}

	// Found a single Complement rule.
	ok = true
	v.remove(tokens)
	complement = ast.ComplementClass().Complement(logical)
	return
}

func (v *parser_) parseComponent() (
	component ast.ComponentLike,
	token TokenLike,
	ok bool,
) {
	var tokens = col.List[TokenLike]()

	// Attempt to parse a single Entity rule.
	var entity ast.EntityLike
	entity, token, ok = v.parseEntity()
	switch {
	case ok:
		// No additional put backs allowed at this point.
		tokens = nil
	case uti.IsDefined(tokens):
		// This is not a single Entity rule.
		v.putBack(tokens)
		return
	default:
		// Found a syntax error.
		var message = v.formatError("$Component", token)
		panic(message)
	}

	// Attempt to parse an optional Parameters rule.
	var optionalParameters ast.ParametersLike
	optionalParameters, _, ok = v.parseParameters()
	if ok {
		// No additional put backs allowed at this point.
		tokens = nil
	}

	// Found a single Component rule.
	ok = true
	v.remove(tokens)
	component = ast.ComponentClass().Component(
		entity,
		optionalParameters,
	)
	return
}

func (v *parser_) parseCondition() (
	condition ast.ConditionLike,
	token TokenLike,
	ok bool,
) {
	var tokens = col.List[TokenLike]()

	// Attempt to parse a single Expression rule.
	var expression ast.ExpressionLike
	expression, token, ok = v.parseExpression()
	switch {
	case ok:
		// No additional put backs allowed at this point.
		tokens = nil
	case uti.IsDefined(tokens):
		// This is not a single Expression rule.
		v.putBack(tokens)
		return
	default:
		// Found a syntax error.
		var message = v.formatError("$Condition", token)
		panic(message)
	}

	// Found a single Condition rule.
	ok = true
	v.remove(tokens)
	condition = ast.ConditionClass().Condition(expression)
	return
}

func (v *parser_) parseContinueClause() (
	continueClause ast.ContinueClauseLike,
	token TokenLike,
	ok bool,
) {
	var tokens = col.List[TokenLike]()

	// Attempt to parse a single "continue" delimiter.
	_, token, ok = v.parseDelimiter("continue")
	if !ok {
		if uti.IsDefined(tokens) {
			// This is not a single ContinueClause rule.
			v.putBack(tokens)
			return
		} else {
			// Found a syntax error.
			var message = v.formatError("$ContinueClause", token)
			panic(message)
		}
	}
	if uti.IsDefined(tokens) {
		tokens.AppendValue(token)
	}

	// Attempt to parse a single "loop" delimiter.
	_, token, ok = v.parseDelimiter("loop")
	if !ok {
		if uti.IsDefined(tokens) {
			// This is not a single ContinueClause rule.
			v.putBack(tokens)
			return
		} else {
			// Found a syntax error.
			var message = v.formatError("$ContinueClause", token)
			panic(message)
		}
	}
	if uti.IsDefined(tokens) {
		tokens.AppendValue(token)
	}

	// Found a single ContinueClause rule.
	ok = true
	v.remove(tokens)
	continueClause = ast.ContinueClauseClass().ContinueClause()
	return
}

func (v *parser_) parseDashEqual() (
	dashEqual ast.DashEqualLike,
	token TokenLike,
	ok bool,
) {
	var tokens = col.List[TokenLike]()

	// Attempt to parse a single "-=" delimiter.
	_, token, ok = v.parseDelimiter("-=")
	if !ok {
		if uti.IsDefined(tokens) {
			// This is not a single DashEqual rule.
			v.putBack(tokens)
			return
		} else {
			// Found a syntax error.
			var message = v.formatError("$DashEqual", token)
			panic(message)
		}
	}
	if uti.IsDefined(tokens) {
		tokens.AppendValue(token)
	}

	// Found a single DashEqual rule.
	ok = true
	v.remove(tokens)
	dashEqual = ast.DashEqualClass().DashEqual()
	return
}

func (v *parser_) parseDefaultEqual() (
	defaultEqual ast.DefaultEqualLike,
	token TokenLike,
	ok bool,
) {
	var tokens = col.List[TokenLike]()

	// Attempt to parse a single "?=" delimiter.
	_, token, ok = v.parseDelimiter("?=")
	if !ok {
		if uti.IsDefined(tokens) {
			// This is not a single DefaultEqual rule.
			v.putBack(tokens)
			return
		} else {
			// Found a syntax error.
			var message = v.formatError("$DefaultEqual", token)
			panic(message)
		}
	}
	if uti.IsDefined(tokens) {
		tokens.AppendValue(token)
	}

	// Found a single DefaultEqual rule.
	ok = true
	v.remove(tokens)
	defaultEqual = ast.DefaultEqualClass().DefaultEqual()
	return
}

func (v *parser_) parseDiscardClause() (
	discardClause ast.DiscardClauseLike,
	token TokenLike,
	ok bool,
) {
	var tokens = col.List[TokenLike]()

	// Attempt to parse a single "discard" delimiter.
	_, token, ok = v.parseDelimiter("discard")
	if !ok {
		if uti.IsDefined(tokens) {
			// This is not a single DiscardClause rule.
			v.putBack(tokens)
			return
		} else {
			// Found a syntax error.
			var message = v.formatError("$DiscardClause", token)
			panic(message)
		}
	}
	if uti.IsDefined(tokens) {
		tokens.AppendValue(token)
	}

	// Attempt to parse a single Draft rule.
	var draft ast.DraftLike
	draft, token, ok = v.parseDraft()
	switch {
	case ok:
		// No additional put backs allowed at this point.
		tokens = nil
	case uti.IsDefined(tokens):
		// This is not a single Draft rule.
		v.putBack(tokens)
		return
	default:
		// Found a syntax error.
		var message = v.formatError("$DiscardClause", token)
		panic(message)
	}

	// Found a single DiscardClause rule.
	ok = true
	v.remove(tokens)
	discardClause = ast.DiscardClauseClass().DiscardClause(draft)
	return
}

func (v *parser_) parseDoClause() (
	doClause ast.DoClauseLike,
	token TokenLike,
	ok bool,
) {
	var tokens = col.List[TokenLike]()

	// Attempt to parse a single "do" delimiter.
	_, token, ok = v.parseDelimiter("do")
	if !ok {
		if uti.IsDefined(tokens) {
			// This is not a single DoClause rule.
			v.putBack(tokens)
			return
		} else {
			// Found a syntax error.
			var message = v.formatError("$DoClause", token)
			panic(message)
		}
	}
	if uti.IsDefined(tokens) {
		tokens.AppendValue(token)
	}

	// Attempt to parse a single Invocation rule.
	var invocation ast.InvocationLike
	invocation, token, ok = v.parseInvocation()
	switch {
	case ok:
		// No additional put backs allowed at this point.
		tokens = nil
	case uti.IsDefined(tokens):
		// This is not a single Invocation rule.
		v.putBack(tokens)
		return
	default:
		// Found a syntax error.
		var message = v.formatError("$DoClause", token)
		panic(message)
	}

	// Found a single DoClause rule.
	ok = true
	v.remove(tokens)
	doClause = ast.DoClauseClass().DoClause(invocation)
	return
}

func (v *parser_) parseDocument() (
	document ast.DocumentLike,
	token TokenLike,
	ok bool,
) {
	var tokens = col.List[TokenLike]()

	// Attempt to parse an optional Notice rule.
	var optionalNotice ast.NoticeLike
	optionalNotice, _, ok = v.parseNotice()
	if ok {
		// No additional put backs allowed at this point.
		tokens = nil
	}

	// Attempt to parse a single Component rule.
	var component ast.ComponentLike
	component, token, ok = v.parseComponent()
	switch {
	case ok:
		// No additional put backs allowed at this point.
		tokens = nil
	case uti.IsDefined(tokens):
		// This is not a single Component rule.
		v.putBack(tokens)
		return
	default:
		// Found a syntax error.
		var message = v.formatError("$Document", token)
		panic(message)
	}

	// Found a single Document rule.
	ok = true
	v.remove(tokens)
	document = ast.DocumentClass().Document(
		optionalNotice,
		component,
	)
	return
}

func (v *parser_) parseDraft() (
	draft ast.DraftLike,
	token TokenLike,
	ok bool,
) {
	var tokens = col.List[TokenLike]()

	// Attempt to parse a single Expression rule.
	var expression ast.ExpressionLike
	expression, token, ok = v.parseExpression()
	switch {
	case ok:
		// No additional put backs allowed at this point.
		tokens = nil
	case uti.IsDefined(tokens):
		// This is not a single Expression rule.
		v.putBack(tokens)
		return
	default:
		// Found a syntax error.
		var message = v.formatError("$Draft", token)
		panic(message)
	}

	// Found a single Draft rule.
	ok = true
	v.remove(tokens)
	draft = ast.DraftClass().Draft(expression)
	return
}

func (v *parser_) parseElement() (
	element ast.ElementLike,
	token TokenLike,
	ok bool,
) {
	// Attempt to parse a single angle Element.
	var angle string
	angle, token, ok = v.parseToken(AngleToken)
	if ok {
		// Found a single angle Element.
		element = ast.ElementClass().Element(angle)
		return
	}

	// Attempt to parse a single boolean Element.
	var boolean string
	boolean, token, ok = v.parseToken(BooleanToken)
	if ok {
		// Found a single boolean Element.
		element = ast.ElementClass().Element(boolean)
		return
	}

	// Attempt to parse a single citation Element.
	var citation string
	citation, token, ok = v.parseToken(CitationToken)
	if ok {
		// Found a single citation Element.
		element = ast.ElementClass().Element(citation)
		return
	}

	// Attempt to parse a single duration Element.
	var duration string
	duration, token, ok = v.parseToken(DurationToken)
	if ok {
		// Found a single duration Element.
		element = ast.ElementClass().Element(duration)
		return
	}

	// Attempt to parse a single moment Element.
	var moment string
	moment, token, ok = v.parseToken(MomentToken)
	if ok {
		// Found a single moment Element.
		element = ast.ElementClass().Element(moment)
		return
	}

	// Attempt to parse a single number Element.
	var number string
	number, token, ok = v.parseToken(NumberToken)
	if ok {
		// Found a single number Element.
		element = ast.ElementClass().Element(number)
		return
	}

	// Attempt to parse a single pattern Element.
	var pattern string
	pattern, token, ok = v.parseToken(PatternToken)
	if ok {
		// Found a single pattern Element.
		element = ast.ElementClass().Element(pattern)
		return
	}

	// Attempt to parse a single percentage Element.
	var percentage string
	percentage, token, ok = v.parseToken(PercentageToken)
	if ok {
		// Found a single percentage Element.
		element = ast.ElementClass().Element(percentage)
		return
	}

	// Attempt to parse a single probability Element.
	var probability string
	probability, token, ok = v.parseToken(ProbabilityToken)
	if ok {
		// Found a single probability Element.
		element = ast.ElementClass().Element(probability)
		return
	}

	// Attempt to parse a single resource Element.
	var resource string
	resource, token, ok = v.parseToken(ResourceToken)
	if ok {
		// Found a single resource Element.
		element = ast.ElementClass().Element(resource)
		return
	}

	// This is not a single Element rule.
	return
}

func (v *parser_) parseEntity() (
	entity ast.EntityLike,
	token TokenLike,
	ok bool,
) {
	// Attempt to parse a single Element Entity.
	var element ast.ElementLike
	element, token, ok = v.parseElement()
	if ok {
		// Found a single Element Entity.
		entity = ast.EntityClass().Entity(element)
		return
	}

	// Attempt to parse a single String Entity.
	var string_ ast.StringLike
	string_, token, ok = v.parseString()
	if ok {
		// Found a single String Entity.
		entity = ast.EntityClass().Entity(string_)
		return
	}

	// Attempt to parse a single Collection Entity.
	var collection ast.CollectionLike
	collection, token, ok = v.parseCollection()
	if ok {
		// Found a single Collection Entity.
		entity = ast.EntityClass().Entity(collection)
		return
	}

	// Attempt to parse a single Procedure Entity.
	var procedure ast.ProcedureLike
	procedure, token, ok = v.parseProcedure()
	if ok {
		// Found a single Procedure Entity.
		entity = ast.EntityClass().Entity(procedure)
		return
	}

	// This is not a single Entity rule.
	return
}

func (v *parser_) parseEvent() (
	event ast.EventLike,
	token TokenLike,
	ok bool,
) {
	var tokens = col.List[TokenLike]()

	// Attempt to parse a single Expression rule.
	var expression ast.ExpressionLike
	expression, token, ok = v.parseExpression()
	switch {
	case ok:
		// No additional put backs allowed at this point.
		tokens = nil
	case uti.IsDefined(tokens):
		// This is not a single Expression rule.
		v.putBack(tokens)
		return
	default:
		// Found a syntax error.
		var message = v.formatError("$Event", token)
		panic(message)
	}

	// Found a single Event rule.
	ok = true
	v.remove(tokens)
	event = ast.EventClass().Event(expression)
	return
}

func (v *parser_) parseException() (
	exception ast.ExceptionLike,
	token TokenLike,
	ok bool,
) {
	var tokens = col.List[TokenLike]()

	// Attempt to parse a single Expression rule.
	var expression ast.ExpressionLike
	expression, token, ok = v.parseExpression()
	switch {
	case ok:
		// No additional put backs allowed at this point.
		tokens = nil
	case uti.IsDefined(tokens):
		// This is not a single Expression rule.
		v.putBack(tokens)
		return
	default:
		// Found a syntax error.
		var message = v.formatError("$Exception", token)
		panic(message)
	}

	// Found a single Exception rule.
	ok = true
	v.remove(tokens)
	exception = ast.ExceptionClass().Exception(expression)
	return
}

func (v *parser_) parseExclusive() (
	exclusive ast.ExclusiveLike,
	token TokenLike,
	ok bool,
) {
	var tokens = col.List[TokenLike]()

	// Attempt to parse a single ")" delimiter.
	_, token, ok = v.parseDelimiter(")")
	if !ok {
		if uti.IsDefined(tokens) {
			// This is not a single Exclusive rule.
			v.putBack(tokens)
			return
		} else {
			// Found a syntax error.
			var message = v.formatError("$Exclusive", token)
			panic(message)
		}
	}
	if uti.IsDefined(tokens) {
		tokens.AppendValue(token)
	}

	// Found a single Exclusive rule.
	ok = true
	v.remove(tokens)
	exclusive = ast.ExclusiveClass().Exclusive()
	return
}

func (v *parser_) parseExclusiveRange() (
	exclusiveRange ast.ExclusiveRangeLike,
	token TokenLike,
	ok bool,
) {
	var tokens = col.List[TokenLike]()

	// Attempt to parse a single "(" delimiter.
	_, token, ok = v.parseDelimiter("(")
	if !ok {
		if uti.IsDefined(tokens) {
			// This is not a single ExclusiveRange rule.
			v.putBack(tokens)
			return
		} else {
			// Found a syntax error.
			var message = v.formatError("$ExclusiveRange", token)
			panic(message)
		}
	}
	if uti.IsDefined(tokens) {
		tokens.AppendValue(token)
	}

	// Attempt to parse a single Primitive rule.
	var primitive1 ast.PrimitiveLike
	primitive1, token, ok = v.parsePrimitive()
	switch {
	case ok:
		// Found a multiexpression token.
		if uti.IsDefined(tokens) {
			tokens.AppendValue(token)
		}
	case uti.IsDefined(tokens):
		// This is not a single Primitive rule.
		v.putBack(tokens)
		return
	default:
		// Found a syntax error.
		var message = v.formatError("$ExclusiveRange", token)
		panic(message)
	}

	// Attempt to parse a single ".." delimiter.
	_, token, ok = v.parseDelimiter("..")
	if !ok {
		if uti.IsDefined(tokens) {
			// This is not a single ExclusiveRange rule.
			v.putBack(tokens)
			return
		} else {
			// Found a syntax error.
			var message = v.formatError("$ExclusiveRange", token)
			panic(message)
		}
	}
	if uti.IsDefined(tokens) {
		tokens.AppendValue(token)
	}

	// Attempt to parse a single Primitive rule.
	var primitive2 ast.PrimitiveLike
	primitive2, token, ok = v.parsePrimitive()
	switch {
	case ok:
		// Found a multiexpression token.
		if uti.IsDefined(tokens) {
			tokens.AppendValue(token)
		}
	case uti.IsDefined(tokens):
		// This is not a single Primitive rule.
		v.putBack(tokens)
		return
	default:
		// Found a syntax error.
		var message = v.formatError("$ExclusiveRange", token)
		panic(message)
	}

	// Attempt to parse a single Bracket rule.
	var bracket ast.BracketLike
	bracket, token, ok = v.parseBracket()
	switch {
	case ok:
		// No additional put backs allowed at this point.
		tokens = nil
	case uti.IsDefined(tokens):
		// This is not a single Bracket rule.
		v.putBack(tokens)
		return
	default:
		// Found a syntax error.
		var message = v.formatError("$ExclusiveRange", token)
		panic(message)
	}

	// Found a single ExclusiveRange rule.
	ok = true
	v.remove(tokens)
	exclusiveRange = ast.ExclusiveRangeClass().ExclusiveRange(
		primitive1,
		primitive2,
		bracket,
	)
	return
}

func (v *parser_) parseExpression() (
	expression ast.ExpressionLike,
	token TokenLike,
	ok bool,
) {
	var tokens = col.List[TokenLike]()

	// Attempt to parse a single Subject rule.
	var subject ast.SubjectLike
	subject, token, ok = v.parseSubject()
	switch {
	case ok:
		// No additional put backs allowed at this point.
		tokens = nil
	case uti.IsDefined(tokens):
		// This is not a single Subject rule.
		v.putBack(tokens)
		return
	default:
		// Found a syntax error.
		var message = v.formatError("$Expression", token)
		panic(message)
	}

	// Attempt to parse multiple Predicate rules.
	var predicates = col.List[ast.PredicateLike]()
predicatesLoop:
	for count := 0; count < mat.MaxInt; count++ {
		var predicate ast.PredicateLike
		predicate, token, ok = v.parsePredicate()
		if !ok {
			switch {
			case count >= 0:
				break predicatesLoop
			case uti.IsDefined(tokens):
				// This is not multiple Predicate rules.
				v.putBack(tokens)
				return
			default:
				// Found a syntax error.
				var message = v.formatError("$Expression", token)
				message += "0 or more Predicate rules are required."
				panic(message)
			}
		}
		// No additional put backs allowed at this point.
		tokens = nil
		predicates.AppendValue(predicate)
	}

	// Found a single Expression rule.
	ok = true
	v.remove(tokens)
	expression = ast.ExpressionClass().Expression(
		subject,
		predicates,
	)
	return
}

func (v *parser_) parseFailure() (
	failure ast.FailureLike,
	token TokenLike,
	ok bool,
) {
	var tokens = col.List[TokenLike]()

	// Attempt to parse a single symbol token.
	var symbol string
	symbol, token, ok = v.parseToken(SymbolToken)
	if !ok {
		if uti.IsDefined(tokens) {
			// This is not a single symbol token.
			v.putBack(tokens)
			return
		} else {
			// Found a syntax error.
			var message = v.formatError("$Failure", token)
			panic(message)
		}
	}
	if uti.IsDefined(tokens) {
		tokens.AppendValue(token)
	}

	// Found a single Failure rule.
	ok = true
	v.remove(tokens)
	failure = ast.FailureClass().Failure(symbol)
	return
}

func (v *parser_) parseFlow() (
	flow ast.FlowLike,
	token TokenLike,
	ok bool,
) {
	// Attempt to parse a single IfClause Flow.
	var ifClause ast.IfClauseLike
	ifClause, token, ok = v.parseIfClause()
	if ok {
		// Found a single IfClause Flow.
		flow = ast.FlowClass().Flow(ifClause)
		return
	}

	// Attempt to parse a single SelectClause Flow.
	var selectClause ast.SelectClauseLike
	selectClause, token, ok = v.parseSelectClause()
	if ok {
		// Found a single SelectClause Flow.
		flow = ast.FlowClass().Flow(selectClause)
		return
	}

	// Attempt to parse a single WhileClause Flow.
	var whileClause ast.WhileClauseLike
	whileClause, token, ok = v.parseWhileClause()
	if ok {
		// Found a single WhileClause Flow.
		flow = ast.FlowClass().Flow(whileClause)
		return
	}

	// Attempt to parse a single WithClause Flow.
	var withClause ast.WithClauseLike
	withClause, token, ok = v.parseWithClause()
	if ok {
		// Found a single WithClause Flow.
		flow = ast.FlowClass().Flow(withClause)
		return
	}

	// Attempt to parse a single ContinueClause Flow.
	var continueClause ast.ContinueClauseLike
	continueClause, token, ok = v.parseContinueClause()
	if ok {
		// Found a single ContinueClause Flow.
		flow = ast.FlowClass().Flow(continueClause)
		return
	}

	// Attempt to parse a single BreakClause Flow.
	var breakClause ast.BreakClauseLike
	breakClause, token, ok = v.parseBreakClause()
	if ok {
		// Found a single BreakClause Flow.
		flow = ast.FlowClass().Flow(breakClause)
		return
	}

	// Attempt to parse a single ReturnClause Flow.
	var returnClause ast.ReturnClauseLike
	returnClause, token, ok = v.parseReturnClause()
	if ok {
		// Found a single ReturnClause Flow.
		flow = ast.FlowClass().Flow(returnClause)
		return
	}

	// Attempt to parse a single ThrowClause Flow.
	var throwClause ast.ThrowClauseLike
	throwClause, token, ok = v.parseThrowClause()
	if ok {
		// Found a single ThrowClause Flow.
		flow = ast.FlowClass().Flow(throwClause)
		return
	}

	// This is not a single Flow rule.
	return
}

func (v *parser_) parseFunction() (
	function ast.FunctionLike,
	token TokenLike,
	ok bool,
) {
	var tokens = col.List[TokenLike]()

	// Attempt to parse a single identifier token.
	var identifier string
	identifier, token, ok = v.parseToken(IdentifierToken)
	if !ok {
		if uti.IsDefined(tokens) {
			// This is not a single identifier token.
			v.putBack(tokens)
			return
		} else {
			// Found a syntax error.
			var message = v.formatError("$Function", token)
			panic(message)
		}
	}
	if uti.IsDefined(tokens) {
		tokens.AppendValue(token)
	}

	// Attempt to parse a single "(" delimiter.
	_, token, ok = v.parseDelimiter("(")
	if !ok {
		if uti.IsDefined(tokens) {
			// This is not a single Function rule.
			v.putBack(tokens)
			return
		} else {
			// Found a syntax error.
			var message = v.formatError("$Function", token)
			panic(message)
		}
	}
	if uti.IsDefined(tokens) {
		tokens.AppendValue(token)
	}

	// Attempt to parse an optional Arguments rule.
	var optionalArguments ast.ArgumentsLike
	optionalArguments, _, ok = v.parseArguments()
	if ok {
		// No additional put backs allowed at this point.
		tokens = nil
	}

	// Attempt to parse a single ")" delimiter.
	_, token, ok = v.parseDelimiter(")")
	if !ok {
		if uti.IsDefined(tokens) {
			// This is not a single Function rule.
			v.putBack(tokens)
			return
		} else {
			// Found a syntax error.
			var message = v.formatError("$Function", token)
			panic(message)
		}
	}
	if uti.IsDefined(tokens) {
		tokens.AppendValue(token)
	}

	// Found a single Function rule.
	ok = true
	v.remove(tokens)
	function = ast.FunctionClass().Function(
		identifier,
		optionalArguments,
	)
	return
}

func (v *parser_) parseIfClause() (
	ifClause ast.IfClauseLike,
	token TokenLike,
	ok bool,
) {
	var tokens = col.List[TokenLike]()

	// Attempt to parse a single "if" delimiter.
	_, token, ok = v.parseDelimiter("if")
	if !ok {
		if uti.IsDefined(tokens) {
			// This is not a single IfClause rule.
			v.putBack(tokens)
			return
		} else {
			// Found a syntax error.
			var message = v.formatError("$IfClause", token)
			panic(message)
		}
	}
	if uti.IsDefined(tokens) {
		tokens.AppendValue(token)
	}

	// Attempt to parse a single Condition rule.
	var condition ast.ConditionLike
	condition, token, ok = v.parseCondition()
	switch {
	case ok:
		// No additional put backs allowed at this point.
		tokens = nil
	case uti.IsDefined(tokens):
		// This is not a single Condition rule.
		v.putBack(tokens)
		return
	default:
		// Found a syntax error.
		var message = v.formatError("$IfClause", token)
		panic(message)
	}

	// Attempt to parse a single "do" delimiter.
	_, token, ok = v.parseDelimiter("do")
	if !ok {
		if uti.IsDefined(tokens) {
			// This is not a single IfClause rule.
			v.putBack(tokens)
			return
		} else {
			// Found a syntax error.
			var message = v.formatError("$IfClause", token)
			panic(message)
		}
	}
	if uti.IsDefined(tokens) {
		tokens.AppendValue(token)
	}

	// Attempt to parse a single Procedure rule.
	var procedure ast.ProcedureLike
	procedure, token, ok = v.parseProcedure()
	switch {
	case ok:
		// No additional put backs allowed at this point.
		tokens = nil
	case uti.IsDefined(tokens):
		// This is not a single Procedure rule.
		v.putBack(tokens)
		return
	default:
		// Found a syntax error.
		var message = v.formatError("$IfClause", token)
		panic(message)
	}

	// Found a single IfClause rule.
	ok = true
	v.remove(tokens)
	ifClause = ast.IfClauseClass().IfClause(
		condition,
		procedure,
	)
	return
}

func (v *parser_) parseInclusive() (
	inclusive ast.InclusiveLike,
	token TokenLike,
	ok bool,
) {
	var tokens = col.List[TokenLike]()

	// Attempt to parse a single "]" delimiter.
	_, token, ok = v.parseDelimiter("]")
	if !ok {
		if uti.IsDefined(tokens) {
			// This is not a single Inclusive rule.
			v.putBack(tokens)
			return
		} else {
			// Found a syntax error.
			var message = v.formatError("$Inclusive", token)
			panic(message)
		}
	}
	if uti.IsDefined(tokens) {
		tokens.AppendValue(token)
	}

	// Found a single Inclusive rule.
	ok = true
	v.remove(tokens)
	inclusive = ast.InclusiveClass().Inclusive()
	return
}

func (v *parser_) parseInclusiveRange() (
	inclusiveRange ast.InclusiveRangeLike,
	token TokenLike,
	ok bool,
) {
	var tokens = col.List[TokenLike]()

	// Attempt to parse a single "[" delimiter.
	_, token, ok = v.parseDelimiter("[")
	if !ok {
		if uti.IsDefined(tokens) {
			// This is not a single InclusiveRange rule.
			v.putBack(tokens)
			return
		} else {
			// Found a syntax error.
			var message = v.formatError("$InclusiveRange", token)
			panic(message)
		}
	}
	if uti.IsDefined(tokens) {
		tokens.AppendValue(token)
	}

	// Attempt to parse a single Primitive rule.
	var primitive1 ast.PrimitiveLike
	primitive1, token, ok = v.parsePrimitive()
	switch {
	case ok:
		// Found a multiexpression token.
		if uti.IsDefined(tokens) {
			tokens.AppendValue(token)
		}
	case uti.IsDefined(tokens):
		// This is not a single Primitive rule.
		v.putBack(tokens)
		return
	default:
		// Found a syntax error.
		var message = v.formatError("$InclusiveRange", token)
		panic(message)
	}

	// Attempt to parse a single ".." delimiter.
	_, token, ok = v.parseDelimiter("..")
	if !ok {
		if uti.IsDefined(tokens) {
			// This is not a single InclusiveRange rule.
			v.putBack(tokens)
			return
		} else {
			// Found a syntax error.
			var message = v.formatError("$InclusiveRange", token)
			panic(message)
		}
	}
	if uti.IsDefined(tokens) {
		tokens.AppendValue(token)
	}

	// Attempt to parse a single Primitive rule.
	var primitive2 ast.PrimitiveLike
	primitive2, token, ok = v.parsePrimitive()
	switch {
	case ok:
		// Found a multiexpression token.
		if uti.IsDefined(tokens) {
			tokens.AppendValue(token)
		}
	case uti.IsDefined(tokens):
		// This is not a single Primitive rule.
		v.putBack(tokens)
		return
	default:
		// Found a syntax error.
		var message = v.formatError("$InclusiveRange", token)
		panic(message)
	}

	// Attempt to parse a single Bracket rule.
	var bracket ast.BracketLike
	bracket, token, ok = v.parseBracket()
	switch {
	case ok:
		// No additional put backs allowed at this point.
		tokens = nil
	case uti.IsDefined(tokens):
		// This is not a single Bracket rule.
		v.putBack(tokens)
		return
	default:
		// Found a syntax error.
		var message = v.formatError("$InclusiveRange", token)
		panic(message)
	}

	// Found a single InclusiveRange rule.
	ok = true
	v.remove(tokens)
	inclusiveRange = ast.InclusiveRangeClass().InclusiveRange(
		primitive1,
		primitive2,
		bracket,
	)
	return
}

func (v *parser_) parseIndex() (
	index ast.IndexLike,
	token TokenLike,
	ok bool,
) {
	var tokens = col.List[TokenLike]()

	// Attempt to parse a single Expression rule.
	var expression ast.ExpressionLike
	expression, token, ok = v.parseExpression()
	switch {
	case ok:
		// No additional put backs allowed at this point.
		tokens = nil
	case uti.IsDefined(tokens):
		// This is not a single Expression rule.
		v.putBack(tokens)
		return
	default:
		// Found a syntax error.
		var message = v.formatError("$Index", token)
		panic(message)
	}

	// Found a single Index rule.
	ok = true
	v.remove(tokens)
	index = ast.IndexClass().Index(expression)
	return
}

func (v *parser_) parseIndices() (
	indices ast.IndicesLike,
	token TokenLike,
	ok bool,
) {
	var tokens = col.List[TokenLike]()

	// Attempt to parse a single Index rule.
	var index ast.IndexLike
	index, token, ok = v.parseIndex()
	switch {
	case ok:
		// No additional put backs allowed at this point.
		tokens = nil
	case uti.IsDefined(tokens):
		// This is not a single Index rule.
		v.putBack(tokens)
		return
	default:
		// Found a syntax error.
		var message = v.formatError("$Indices", token)
		panic(message)
	}

	// Attempt to parse multiple AdditionalIndex rules.
	var additionalIndexes = col.List[ast.AdditionalIndexLike]()
additionalIndexesLoop:
	for count := 0; count < mat.MaxInt; count++ {
		var additionalIndex ast.AdditionalIndexLike
		additionalIndex, token, ok = v.parseAdditionalIndex()
		if !ok {
			switch {
			case count >= 0:
				break additionalIndexesLoop
			case uti.IsDefined(tokens):
				// This is not multiple AdditionalIndex rules.
				v.putBack(tokens)
				return
			default:
				// Found a syntax error.
				var message = v.formatError("$Indices", token)
				message += "0 or more AdditionalIndex rules are required."
				panic(message)
			}
		}
		// No additional put backs allowed at this point.
		tokens = nil
		additionalIndexes.AppendValue(additionalIndex)
	}

	// Found a single Indices rule.
	ok = true
	v.remove(tokens)
	indices = ast.IndicesClass().Indices(
		index,
		additionalIndexes,
	)
	return
}

func (v *parser_) parseIndirect() (
	indirect ast.IndirectLike,
	token TokenLike,
	ok bool,
) {
	// Attempt to parse a single Component Indirect.
	var component ast.ComponentLike
	component, token, ok = v.parseComponent()
	if ok {
		// Found a single Component Indirect.
		indirect = ast.IndirectClass().Indirect(component)
		return
	}

	// Attempt to parse a single Subcomponent Indirect.
	var subcomponent ast.SubcomponentLike
	subcomponent, token, ok = v.parseSubcomponent()
	if ok {
		// Found a single Subcomponent Indirect.
		indirect = ast.IndirectClass().Indirect(subcomponent)
		return
	}

	// Attempt to parse a single Referent Indirect.
	var referent ast.ReferentLike
	referent, token, ok = v.parseReferent()
	if ok {
		// Found a single Referent Indirect.
		indirect = ast.IndirectClass().Indirect(referent)
		return
	}

	// Attempt to parse a single Function Indirect.
	var function ast.FunctionLike
	function, token, ok = v.parseFunction()
	if ok {
		// Found a single Function Indirect.
		indirect = ast.IndirectClass().Indirect(function)
		return
	}

	// Attempt to parse a single Method Indirect.
	var method ast.MethodLike
	method, token, ok = v.parseMethod()
	if ok {
		// Found a single Method Indirect.
		indirect = ast.IndirectClass().Indirect(method)
		return
	}

	// Attempt to parse a single Variable Indirect.
	var variable ast.VariableLike
	variable, token, ok = v.parseVariable()
	if ok {
		// Found a single Variable Indirect.
		indirect = ast.IndirectClass().Indirect(variable)
		return
	}

	// This is not a single Indirect rule.
	return
}

func (v *parser_) parseInduction() (
	induction ast.InductionLike,
	token TokenLike,
	ok bool,
) {
	// Attempt to parse a single DoClause Induction.
	var doClause ast.DoClauseLike
	doClause, token, ok = v.parseDoClause()
	if ok {
		// Found a single DoClause Induction.
		induction = ast.InductionClass().Induction(doClause)
		return
	}

	// Attempt to parse a single LetClause Induction.
	var letClause ast.LetClauseLike
	letClause, token, ok = v.parseLetClause()
	if ok {
		// Found a single LetClause Induction.
		induction = ast.InductionClass().Induction(letClause)
		return
	}

	// This is not a single Induction rule.
	return
}

func (v *parser_) parseInlineAttributes() (
	inlineAttributes ast.InlineAttributesLike,
	token TokenLike,
	ok bool,
) {
	var tokens = col.List[TokenLike]()

	// Attempt to parse a single "[" delimiter.
	_, token, ok = v.parseDelimiter("[")
	if !ok {
		if uti.IsDefined(tokens) {
			// This is not a single InlineAttributes rule.
			v.putBack(tokens)
			return
		} else {
			// Found a syntax error.
			var message = v.formatError("$InlineAttributes", token)
			panic(message)
		}
	}
	if uti.IsDefined(tokens) {
		tokens.AppendValue(token)
	}

	// Attempt to parse a single Association rule.
	var association ast.AssociationLike
	association, token, ok = v.parseAssociation()
	switch {
	case ok:
		// No additional put backs allowed at this point.
		tokens = nil
	case uti.IsDefined(tokens):
		// This is not a single Association rule.
		v.putBack(tokens)
		return
	default:
		// Found a syntax error.
		var message = v.formatError("$InlineAttributes", token)
		panic(message)
	}

	// Attempt to parse multiple AdditionalAssociation rules.
	var additionalAssociations = col.List[ast.AdditionalAssociationLike]()
additionalAssociationsLoop:
	for count := 0; count < mat.MaxInt; count++ {
		var additionalAssociation ast.AdditionalAssociationLike
		additionalAssociation, token, ok = v.parseAdditionalAssociation()
		if !ok {
			switch {
			case count >= 0:
				break additionalAssociationsLoop
			case uti.IsDefined(tokens):
				// This is not multiple AdditionalAssociation rules.
				v.putBack(tokens)
				return
			default:
				// Found a syntax error.
				var message = v.formatError("$InlineAttributes", token)
				message += "0 or more AdditionalAssociation rules are required."
				panic(message)
			}
		}
		// No additional put backs allowed at this point.
		tokens = nil
		additionalAssociations.AppendValue(additionalAssociation)
	}

	// Attempt to parse a single "]" delimiter.
	_, token, ok = v.parseDelimiter("]")
	if !ok {
		if uti.IsDefined(tokens) {
			// This is not a single InlineAttributes rule.
			v.putBack(tokens)
			return
		} else {
			// Found a syntax error.
			var message = v.formatError("$InlineAttributes", token)
			panic(message)
		}
	}
	if uti.IsDefined(tokens) {
		tokens.AppendValue(token)
	}

	// Found a single InlineAttributes rule.
	ok = true
	v.remove(tokens)
	inlineAttributes = ast.InlineAttributesClass().InlineAttributes(
		association,
		additionalAssociations,
	)
	return
}

func (v *parser_) parseInlineParameters() (
	inlineParameters ast.InlineParametersLike,
	token TokenLike,
	ok bool,
) {
	var tokens = col.List[TokenLike]()

	// Attempt to parse a single "(" delimiter.
	_, token, ok = v.parseDelimiter("(")
	if !ok {
		if uti.IsDefined(tokens) {
			// This is not a single InlineParameters rule.
			v.putBack(tokens)
			return
		} else {
			// Found a syntax error.
			var message = v.formatError("$InlineParameters", token)
			panic(message)
		}
	}
	if uti.IsDefined(tokens) {
		tokens.AppendValue(token)
	}

	// Attempt to parse a single Association rule.
	var association ast.AssociationLike
	association, token, ok = v.parseAssociation()
	switch {
	case ok:
		// No additional put backs allowed at this point.
		tokens = nil
	case uti.IsDefined(tokens):
		// This is not a single Association rule.
		v.putBack(tokens)
		return
	default:
		// Found a syntax error.
		var message = v.formatError("$InlineParameters", token)
		panic(message)
	}

	// Attempt to parse multiple AdditionalAssociation rules.
	var additionalAssociations = col.List[ast.AdditionalAssociationLike]()
additionalAssociationsLoop:
	for count := 0; count < mat.MaxInt; count++ {
		var additionalAssociation ast.AdditionalAssociationLike
		additionalAssociation, token, ok = v.parseAdditionalAssociation()
		if !ok {
			switch {
			case count >= 0:
				break additionalAssociationsLoop
			case uti.IsDefined(tokens):
				// This is not multiple AdditionalAssociation rules.
				v.putBack(tokens)
				return
			default:
				// Found a syntax error.
				var message = v.formatError("$InlineParameters", token)
				message += "0 or more AdditionalAssociation rules are required."
				panic(message)
			}
		}
		// No additional put backs allowed at this point.
		tokens = nil
		additionalAssociations.AppendValue(additionalAssociation)
	}

	// Attempt to parse a single ")" delimiter.
	_, token, ok = v.parseDelimiter(")")
	if !ok {
		if uti.IsDefined(tokens) {
			// This is not a single InlineParameters rule.
			v.putBack(tokens)
			return
		} else {
			// Found a syntax error.
			var message = v.formatError("$InlineParameters", token)
			panic(message)
		}
	}
	if uti.IsDefined(tokens) {
		tokens.AppendValue(token)
	}

	// Found a single InlineParameters rule.
	ok = true
	v.remove(tokens)
	inlineParameters = ast.InlineParametersClass().InlineParameters(
		association,
		additionalAssociations,
	)
	return
}

func (v *parser_) parseInlineStatements() (
	inlineStatements ast.InlineStatementsLike,
	token TokenLike,
	ok bool,
) {
	var tokens = col.List[TokenLike]()

	// Attempt to parse a single "{" delimiter.
	_, token, ok = v.parseDelimiter("{")
	if !ok {
		if uti.IsDefined(tokens) {
			// This is not a single InlineStatements rule.
			v.putBack(tokens)
			return
		} else {
			// Found a syntax error.
			var message = v.formatError("$InlineStatements", token)
			panic(message)
		}
	}
	if uti.IsDefined(tokens) {
		tokens.AppendValue(token)
	}

	// Attempt to parse a single Statement rule.
	var statement ast.StatementLike
	statement, token, ok = v.parseStatement()
	switch {
	case ok:
		// No additional put backs allowed at this point.
		tokens = nil
	case uti.IsDefined(tokens):
		// This is not a single Statement rule.
		v.putBack(tokens)
		return
	default:
		// Found a syntax error.
		var message = v.formatError("$InlineStatements", token)
		panic(message)
	}

	// Attempt to parse multiple AdditionalStatement rules.
	var additionalStatements = col.List[ast.AdditionalStatementLike]()
additionalStatementsLoop:
	for count := 0; count < mat.MaxInt; count++ {
		var additionalStatement ast.AdditionalStatementLike
		additionalStatement, token, ok = v.parseAdditionalStatement()
		if !ok {
			switch {
			case count >= 0:
				break additionalStatementsLoop
			case uti.IsDefined(tokens):
				// This is not multiple AdditionalStatement rules.
				v.putBack(tokens)
				return
			default:
				// Found a syntax error.
				var message = v.formatError("$InlineStatements", token)
				message += "0 or more AdditionalStatement rules are required."
				panic(message)
			}
		}
		// No additional put backs allowed at this point.
		tokens = nil
		additionalStatements.AppendValue(additionalStatement)
	}

	// Attempt to parse a single "}" delimiter.
	_, token, ok = v.parseDelimiter("}")
	if !ok {
		if uti.IsDefined(tokens) {
			// This is not a single InlineStatements rule.
			v.putBack(tokens)
			return
		} else {
			// Found a syntax error.
			var message = v.formatError("$InlineStatements", token)
			panic(message)
		}
	}
	if uti.IsDefined(tokens) {
		tokens.AppendValue(token)
	}

	// Found a single InlineStatements rule.
	ok = true
	v.remove(tokens)
	inlineStatements = ast.InlineStatementsClass().InlineStatements(
		statement,
		additionalStatements,
	)
	return
}

func (v *parser_) parseInlineValues() (
	inlineValues ast.InlineValuesLike,
	token TokenLike,
	ok bool,
) {
	var tokens = col.List[TokenLike]()

	// Attempt to parse a single "[" delimiter.
	_, token, ok = v.parseDelimiter("[")
	if !ok {
		if uti.IsDefined(tokens) {
			// This is not a single InlineValues rule.
			v.putBack(tokens)
			return
		} else {
			// Found a syntax error.
			var message = v.formatError("$InlineValues", token)
			panic(message)
		}
	}
	if uti.IsDefined(tokens) {
		tokens.AppendValue(token)
	}

	// Attempt to parse a single Component rule.
	var component ast.ComponentLike
	component, token, ok = v.parseComponent()
	switch {
	case ok:
		// No additional put backs allowed at this point.
		tokens = nil
	case uti.IsDefined(tokens):
		// This is not a single Component rule.
		v.putBack(tokens)
		return
	default:
		// Found a syntax error.
		var message = v.formatError("$InlineValues", token)
		panic(message)
	}

	// Attempt to parse multiple AdditionalValue rules.
	var additionalValues = col.List[ast.AdditionalValueLike]()
additionalValuesLoop:
	for count := 0; count < mat.MaxInt; count++ {
		var additionalValue ast.AdditionalValueLike
		additionalValue, token, ok = v.parseAdditionalValue()
		if !ok {
			switch {
			case count >= 0:
				break additionalValuesLoop
			case uti.IsDefined(tokens):
				// This is not multiple AdditionalValue rules.
				v.putBack(tokens)
				return
			default:
				// Found a syntax error.
				var message = v.formatError("$InlineValues", token)
				message += "0 or more AdditionalValue rules are required."
				panic(message)
			}
		}
		// No additional put backs allowed at this point.
		tokens = nil
		additionalValues.AppendValue(additionalValue)
	}

	// Attempt to parse a single "]" delimiter.
	_, token, ok = v.parseDelimiter("]")
	if !ok {
		if uti.IsDefined(tokens) {
			// This is not a single InlineValues rule.
			v.putBack(tokens)
			return
		} else {
			// Found a syntax error.
			var message = v.formatError("$InlineValues", token)
			panic(message)
		}
	}
	if uti.IsDefined(tokens) {
		tokens.AppendValue(token)
	}

	// Found a single InlineValues rule.
	ok = true
	v.remove(tokens)
	inlineValues = ast.InlineValuesClass().InlineValues(
		component,
		additionalValues,
	)
	return
}

func (v *parser_) parseInverse() (
	inverse ast.InverseLike,
	token TokenLike,
	ok bool,
) {
	// Attempt to parse a single dash Inverse.
	var dash string
	dash, token, ok = v.parseToken(DashToken)
	if ok {
		// Found a single dash Inverse.
		inverse = ast.InverseClass().Inverse(dash)
		return
	}

	// Attempt to parse a single slash Inverse.
	var slash string
	slash, token, ok = v.parseToken(SlashToken)
	if ok {
		// Found a single slash Inverse.
		inverse = ast.InverseClass().Inverse(slash)
		return
	}

	// Attempt to parse a single star Inverse.
	var star string
	star, token, ok = v.parseToken(StarToken)
	if ok {
		// Found a single star Inverse.
		inverse = ast.InverseClass().Inverse(star)
		return
	}

	// This is not a single Inverse rule.
	return
}

func (v *parser_) parseInversion() (
	inversion ast.InversionLike,
	token TokenLike,
	ok bool,
) {
	var tokens = col.List[TokenLike]()

	// Attempt to parse a single Inverse rule.
	var inverse ast.InverseLike
	inverse, token, ok = v.parseInverse()
	switch {
	case ok:
		// Found a multiexpression token.
		if uti.IsDefined(tokens) {
			tokens.AppendValue(token)
		}
	case uti.IsDefined(tokens):
		// This is not a single Inverse rule.
		v.putBack(tokens)
		return
	default:
		// Found a syntax error.
		var message = v.formatError("$Inversion", token)
		panic(message)
	}

	// Attempt to parse a single Numerical rule.
	var numerical ast.NumericalLike
	numerical, token, ok = v.parseNumerical()
	switch {
	case ok:
		// No additional put backs allowed at this point.
		tokens = nil
	case uti.IsDefined(tokens):
		// This is not a single Numerical rule.
		v.putBack(tokens)
		return
	default:
		// Found a syntax error.
		var message = v.formatError("$Inversion", token)
		panic(message)
	}

	// Found a single Inversion rule.
	ok = true
	v.remove(tokens)
	inversion = ast.InversionClass().Inversion(
		inverse,
		numerical,
	)
	return
}

func (v *parser_) parseInvocation() (
	invocation ast.InvocationLike,
	token TokenLike,
	ok bool,
) {
	// Attempt to parse a single Function Invocation.
	var function ast.FunctionLike
	function, token, ok = v.parseFunction()
	if ok {
		// Found a single Function Invocation.
		invocation = ast.InvocationClass().Invocation(function)
		return
	}

	// Attempt to parse a single Method Invocation.
	var method ast.MethodLike
	method, token, ok = v.parseMethod()
	if ok {
		// Found a single Method Invocation.
		invocation = ast.InvocationClass().Invocation(method)
		return
	}

	// This is not a single Invocation rule.
	return
}

func (v *parser_) parseIor() (
	ior ast.IorLike,
	token TokenLike,
	ok bool,
) {
	var tokens = col.List[TokenLike]()

	// Attempt to parse a single "ior" delimiter.
	_, token, ok = v.parseDelimiter("ior")
	if !ok {
		if uti.IsDefined(tokens) {
			// This is not a single Ior rule.
			v.putBack(tokens)
			return
		} else {
			// Found a syntax error.
			var message = v.formatError("$Ior", token)
			panic(message)
		}
	}
	if uti.IsDefined(tokens) {
		tokens.AppendValue(token)
	}

	// Found a single Ior rule.
	ok = true
	v.remove(tokens)
	ior = ast.IorClass().Ior()
	return
}

func (v *parser_) parseIs() (
	is ast.IsLike,
	token TokenLike,
	ok bool,
) {
	var tokens = col.List[TokenLike]()

	// Attempt to parse a single "is" delimiter.
	_, token, ok = v.parseDelimiter("is")
	if !ok {
		if uti.IsDefined(tokens) {
			// This is not a single Is rule.
			v.putBack(tokens)
			return
		} else {
			// Found a syntax error.
			var message = v.formatError("$Is", token)
			panic(message)
		}
	}
	if uti.IsDefined(tokens) {
		tokens.AppendValue(token)
	}

	// Found a single Is rule.
	ok = true
	v.remove(tokens)
	is = ast.IsClass().Is()
	return
}

func (v *parser_) parseItem() (
	item ast.ItemLike,
	token TokenLike,
	ok bool,
) {
	var tokens = col.List[TokenLike]()

	// Attempt to parse a single symbol token.
	var symbol string
	symbol, token, ok = v.parseToken(SymbolToken)
	if !ok {
		if uti.IsDefined(tokens) {
			// This is not a single symbol token.
			v.putBack(tokens)
			return
		} else {
			// Found a syntax error.
			var message = v.formatError("$Item", token)
			panic(message)
		}
	}
	if uti.IsDefined(tokens) {
		tokens.AppendValue(token)
	}

	// Found a single Item rule.
	ok = true
	v.remove(tokens)
	item = ast.ItemClass().Item(symbol)
	return
}

func (v *parser_) parseLetClause() (
	letClause ast.LetClauseLike,
	token TokenLike,
	ok bool,
) {
	var tokens = col.List[TokenLike]()

	// Attempt to parse a single "let" delimiter.
	_, token, ok = v.parseDelimiter("let")
	if !ok {
		if uti.IsDefined(tokens) {
			// This is not a single LetClause rule.
			v.putBack(tokens)
			return
		} else {
			// Found a syntax error.
			var message = v.formatError("$LetClause", token)
			panic(message)
		}
	}
	if uti.IsDefined(tokens) {
		tokens.AppendValue(token)
	}

	// Attempt to parse a single Recipient rule.
	var recipient ast.RecipientLike
	recipient, token, ok = v.parseRecipient()
	switch {
	case ok:
		// No additional put backs allowed at this point.
		tokens = nil
	case uti.IsDefined(tokens):
		// This is not a single Recipient rule.
		v.putBack(tokens)
		return
	default:
		// Found a syntax error.
		var message = v.formatError("$LetClause", token)
		panic(message)
	}

	// Attempt to parse a single Assignment rule.
	var assignment ast.AssignmentLike
	assignment, token, ok = v.parseAssignment()
	switch {
	case ok:
		// No additional put backs allowed at this point.
		tokens = nil
	case uti.IsDefined(tokens):
		// This is not a single Assignment rule.
		v.putBack(tokens)
		return
	default:
		// Found a syntax error.
		var message = v.formatError("$LetClause", token)
		panic(message)
	}

	// Attempt to parse a single Expression rule.
	var expression ast.ExpressionLike
	expression, token, ok = v.parseExpression()
	switch {
	case ok:
		// No additional put backs allowed at this point.
		tokens = nil
	case uti.IsDefined(tokens):
		// This is not a single Expression rule.
		v.putBack(tokens)
		return
	default:
		// Found a syntax error.
		var message = v.formatError("$LetClause", token)
		panic(message)
	}

	// Found a single LetClause rule.
	ok = true
	v.remove(tokens)
	letClause = ast.LetClauseClass().LetClause(
		recipient,
		assignment,
		expression,
	)
	return
}

func (v *parser_) parseLogical() (
	logical ast.LogicalLike,
	token TokenLike,
	ok bool,
) {
	// Attempt to parse a single Component Logical.
	var component ast.ComponentLike
	component, token, ok = v.parseComponent()
	if ok {
		// Found a single Component Logical.
		logical = ast.LogicalClass().Logical(component)
		return
	}

	// Attempt to parse a single Subcomponent Logical.
	var subcomponent ast.SubcomponentLike
	subcomponent, token, ok = v.parseSubcomponent()
	if ok {
		// Found a single Subcomponent Logical.
		logical = ast.LogicalClass().Logical(subcomponent)
		return
	}

	// Attempt to parse a single Precedence Logical.
	var precedence ast.PrecedenceLike
	precedence, token, ok = v.parsePrecedence()
	if ok {
		// Found a single Precedence Logical.
		logical = ast.LogicalClass().Logical(precedence)
		return
	}

	// Attempt to parse a single Referent Logical.
	var referent ast.ReferentLike
	referent, token, ok = v.parseReferent()
	if ok {
		// Found a single Referent Logical.
		logical = ast.LogicalClass().Logical(referent)
		return
	}

	// Attempt to parse a single Complement Logical.
	var complement ast.ComplementLike
	complement, token, ok = v.parseComplement()
	if ok {
		// Found a single Complement Logical.
		logical = ast.LogicalClass().Logical(complement)
		return
	}

	// Attempt to parse a single Function Logical.
	var function ast.FunctionLike
	function, token, ok = v.parseFunction()
	if ok {
		// Found a single Function Logical.
		logical = ast.LogicalClass().Logical(function)
		return
	}

	// Attempt to parse a single Method Logical.
	var method ast.MethodLike
	method, token, ok = v.parseMethod()
	if ok {
		// Found a single Method Logical.
		logical = ast.LogicalClass().Logical(method)
		return
	}

	// Attempt to parse a single Variable Logical.
	var variable ast.VariableLike
	variable, token, ok = v.parseVariable()
	if ok {
		// Found a single Variable Logical.
		logical = ast.LogicalClass().Logical(variable)
		return
	}

	// This is not a single Logical rule.
	return
}

func (v *parser_) parseMagnitude() (
	magnitude ast.MagnitudeLike,
	token TokenLike,
	ok bool,
) {
	var tokens = col.List[TokenLike]()

	// Attempt to parse a single "|" delimiter.
	_, token, ok = v.parseDelimiter("|")
	if !ok {
		if uti.IsDefined(tokens) {
			// This is not a single Magnitude rule.
			v.putBack(tokens)
			return
		} else {
			// Found a syntax error.
			var message = v.formatError("$Magnitude", token)
			panic(message)
		}
	}
	if uti.IsDefined(tokens) {
		tokens.AppendValue(token)
	}

	// Attempt to parse a single Numerical rule.
	var numerical ast.NumericalLike
	numerical, token, ok = v.parseNumerical()
	switch {
	case ok:
		// No additional put backs allowed at this point.
		tokens = nil
	case uti.IsDefined(tokens):
		// This is not a single Numerical rule.
		v.putBack(tokens)
		return
	default:
		// Found a syntax error.
		var message = v.formatError("$Magnitude", token)
		panic(message)
	}

	// Attempt to parse a single "|" delimiter.
	_, token, ok = v.parseDelimiter("|")
	if !ok {
		if uti.IsDefined(tokens) {
			// This is not a single Magnitude rule.
			v.putBack(tokens)
			return
		} else {
			// Found a syntax error.
			var message = v.formatError("$Magnitude", token)
			panic(message)
		}
	}
	if uti.IsDefined(tokens) {
		tokens.AppendValue(token)
	}

	// Found a single Magnitude rule.
	ok = true
	v.remove(tokens)
	magnitude = ast.MagnitudeClass().Magnitude(numerical)
	return
}

func (v *parser_) parseMainClause() (
	mainClause ast.MainClauseLike,
	token TokenLike,
	ok bool,
) {
	// Attempt to parse a single Flow MainClause.
	var flow ast.FlowLike
	flow, token, ok = v.parseFlow()
	if ok {
		// Found a single Flow MainClause.
		mainClause = ast.MainClauseClass().MainClause(flow)
		return
	}

	// Attempt to parse a single Induction MainClause.
	var induction ast.InductionLike
	induction, token, ok = v.parseInduction()
	if ok {
		// Found a single Induction MainClause.
		mainClause = ast.MainClauseClass().MainClause(induction)
		return
	}

	// Attempt to parse a single Messaging MainClause.
	var messaging ast.MessagingLike
	messaging, token, ok = v.parseMessaging()
	if ok {
		// Found a single Messaging MainClause.
		mainClause = ast.MainClauseClass().MainClause(messaging)
		return
	}

	// Attempt to parse a single Repository MainClause.
	var repository ast.RepositoryLike
	repository, token, ok = v.parseRepository()
	if ok {
		// Found a single Repository MainClause.
		mainClause = ast.MainClauseClass().MainClause(repository)
		return
	}

	// This is not a single MainClause rule.
	return
}

func (v *parser_) parseMatchHandler() (
	matchHandler ast.MatchHandlerLike,
	token TokenLike,
	ok bool,
) {
	var tokens = col.List[TokenLike]()

	// Attempt to parse a single "matching" delimiter.
	_, token, ok = v.parseDelimiter("matching")
	if !ok {
		if uti.IsDefined(tokens) {
			// This is not a single MatchHandler rule.
			v.putBack(tokens)
			return
		} else {
			// Found a syntax error.
			var message = v.formatError("$MatchHandler", token)
			panic(message)
		}
	}
	if uti.IsDefined(tokens) {
		tokens.AppendValue(token)
	}

	// Attempt to parse a single Template rule.
	var template ast.TemplateLike
	template, token, ok = v.parseTemplate()
	switch {
	case ok:
		// No additional put backs allowed at this point.
		tokens = nil
	case uti.IsDefined(tokens):
		// This is not a single Template rule.
		v.putBack(tokens)
		return
	default:
		// Found a syntax error.
		var message = v.formatError("$MatchHandler", token)
		panic(message)
	}

	// Attempt to parse a single "do" delimiter.
	_, token, ok = v.parseDelimiter("do")
	if !ok {
		if uti.IsDefined(tokens) {
			// This is not a single MatchHandler rule.
			v.putBack(tokens)
			return
		} else {
			// Found a syntax error.
			var message = v.formatError("$MatchHandler", token)
			panic(message)
		}
	}
	if uti.IsDefined(tokens) {
		tokens.AppendValue(token)
	}

	// Attempt to parse a single Procedure rule.
	var procedure ast.ProcedureLike
	procedure, token, ok = v.parseProcedure()
	switch {
	case ok:
		// No additional put backs allowed at this point.
		tokens = nil
	case uti.IsDefined(tokens):
		// This is not a single Procedure rule.
		v.putBack(tokens)
		return
	default:
		// Found a syntax error.
		var message = v.formatError("$MatchHandler", token)
		panic(message)
	}

	// Found a single MatchHandler rule.
	ok = true
	v.remove(tokens)
	matchHandler = ast.MatchHandlerClass().MatchHandler(
		template,
		procedure,
	)
	return
}

func (v *parser_) parseMatches() (
	matches ast.MatchesLike,
	token TokenLike,
	ok bool,
) {
	var tokens = col.List[TokenLike]()

	// Attempt to parse a single "matches" delimiter.
	_, token, ok = v.parseDelimiter("matches")
	if !ok {
		if uti.IsDefined(tokens) {
			// This is not a single Matches rule.
			v.putBack(tokens)
			return
		} else {
			// Found a syntax error.
			var message = v.formatError("$Matches", token)
			panic(message)
		}
	}
	if uti.IsDefined(tokens) {
		tokens.AppendValue(token)
	}

	// Found a single Matches rule.
	ok = true
	v.remove(tokens)
	matches = ast.MatchesClass().Matches()
	return
}

func (v *parser_) parseMessage() (
	message ast.MessageLike,
	token TokenLike,
	ok bool,
) {
	var tokens = col.List[TokenLike]()

	// Attempt to parse a single Expression rule.
	var expression ast.ExpressionLike
	expression, token, ok = v.parseExpression()
	switch {
	case ok:
		// No additional put backs allowed at this point.
		tokens = nil
	case uti.IsDefined(tokens):
		// This is not a single Expression rule.
		v.putBack(tokens)
		return
	default:
		// Found a syntax error.
		var message = v.formatError("$Message", token)
		panic(message)
	}

	// Found a single Message rule.
	ok = true
	v.remove(tokens)
	message = ast.MessageClass().Message(expression)
	return
}

func (v *parser_) parseMessaging() (
	messaging ast.MessagingLike,
	token TokenLike,
	ok bool,
) {
	// Attempt to parse a single PostClause Messaging.
	var postClause ast.PostClauseLike
	postClause, token, ok = v.parsePostClause()
	if ok {
		// Found a single PostClause Messaging.
		messaging = ast.MessagingClass().Messaging(postClause)
		return
	}

	// Attempt to parse a single RetrieveClause Messaging.
	var retrieveClause ast.RetrieveClauseLike
	retrieveClause, token, ok = v.parseRetrieveClause()
	if ok {
		// Found a single RetrieveClause Messaging.
		messaging = ast.MessagingClass().Messaging(retrieveClause)
		return
	}

	// Attempt to parse a single AcceptClause Messaging.
	var acceptClause ast.AcceptClauseLike
	acceptClause, token, ok = v.parseAcceptClause()
	if ok {
		// Found a single AcceptClause Messaging.
		messaging = ast.MessagingClass().Messaging(acceptClause)
		return
	}

	// Attempt to parse a single RejectClause Messaging.
	var rejectClause ast.RejectClauseLike
	rejectClause, token, ok = v.parseRejectClause()
	if ok {
		// Found a single RejectClause Messaging.
		messaging = ast.MessagingClass().Messaging(rejectClause)
		return
	}

	// Attempt to parse a single PublishClause Messaging.
	var publishClause ast.PublishClauseLike
	publishClause, token, ok = v.parsePublishClause()
	if ok {
		// Found a single PublishClause Messaging.
		messaging = ast.MessagingClass().Messaging(publishClause)
		return
	}

	// This is not a single Messaging rule.
	return
}

func (v *parser_) parseMethod() (
	method ast.MethodLike,
	token TokenLike,
	ok bool,
) {
	var tokens = col.List[TokenLike]()

	// Attempt to parse a single identifier token.
	var identifier1 string
	identifier1, token, ok = v.parseToken(IdentifierToken)
	if !ok {
		if uti.IsDefined(tokens) {
			// This is not a single identifier token.
			v.putBack(tokens)
			return
		} else {
			// Found a syntax error.
			var message = v.formatError("$Method", token)
			panic(message)
		}
	}
	if uti.IsDefined(tokens) {
		tokens.AppendValue(token)
	}

	// Attempt to parse a single Blocking rule.
	var blocking ast.BlockingLike
	blocking, token, ok = v.parseBlocking()
	switch {
	case ok:
		// Found a multiexpression token.
		if uti.IsDefined(tokens) {
			tokens.AppendValue(token)
		}
	case uti.IsDefined(tokens):
		// This is not a single Blocking rule.
		v.putBack(tokens)
		return
	default:
		// Found a syntax error.
		var message = v.formatError("$Method", token)
		panic(message)
	}

	// Attempt to parse a single identifier token.
	var identifier2 string
	identifier2, token, ok = v.parseToken(IdentifierToken)
	if !ok {
		if uti.IsDefined(tokens) {
			// This is not a single identifier token.
			v.putBack(tokens)
			return
		} else {
			// Found a syntax error.
			var message = v.formatError("$Method", token)
			panic(message)
		}
	}
	if uti.IsDefined(tokens) {
		tokens.AppendValue(token)
	}

	// Attempt to parse a single "(" delimiter.
	_, token, ok = v.parseDelimiter("(")
	if !ok {
		if uti.IsDefined(tokens) {
			// This is not a single Method rule.
			v.putBack(tokens)
			return
		} else {
			// Found a syntax error.
			var message = v.formatError("$Method", token)
			panic(message)
		}
	}
	if uti.IsDefined(tokens) {
		tokens.AppendValue(token)
	}

	// Attempt to parse an optional Arguments rule.
	var optionalArguments ast.ArgumentsLike
	optionalArguments, _, ok = v.parseArguments()
	if ok {
		// No additional put backs allowed at this point.
		tokens = nil
	}

	// Attempt to parse a single ")" delimiter.
	_, token, ok = v.parseDelimiter(")")
	if !ok {
		if uti.IsDefined(tokens) {
			// This is not a single Method rule.
			v.putBack(tokens)
			return
		} else {
			// Found a syntax error.
			var message = v.formatError("$Method", token)
			panic(message)
		}
	}
	if uti.IsDefined(tokens) {
		tokens.AppendValue(token)
	}

	// Found a single Method rule.
	ok = true
	v.remove(tokens)
	method = ast.MethodClass().Method(
		identifier1,
		blocking,
		identifier2,
		optionalArguments,
	)
	return
}

func (v *parser_) parseMultilineAttributes() (
	multilineAttributes ast.MultilineAttributesLike,
	token TokenLike,
	ok bool,
) {
	var tokens = col.List[TokenLike]()

	// Attempt to parse a single "[" delimiter.
	_, token, ok = v.parseDelimiter("[")
	if !ok {
		if uti.IsDefined(tokens) {
			// This is not a single MultilineAttributes rule.
			v.putBack(tokens)
			return
		} else {
			// Found a syntax error.
			var message = v.formatError("$MultilineAttributes", token)
			panic(message)
		}
	}
	if uti.IsDefined(tokens) {
		tokens.AppendValue(token)
	}

	// Attempt to parse a single newline token.
	var newline string
	newline, token, ok = v.parseToken(NewlineToken)
	if !ok {
		if uti.IsDefined(tokens) {
			// This is not a single newline token.
			v.putBack(tokens)
			return
		} else {
			// Found a syntax error.
			var message = v.formatError("$MultilineAttributes", token)
			panic(message)
		}
	}
	if uti.IsDefined(tokens) {
		tokens.AppendValue(token)
	}

	// Attempt to parse multiple AnnotatedAssociation rules.
	var annotatedAssociations = col.List[ast.AnnotatedAssociationLike]()
annotatedAssociationsLoop:
	for count := 0; count < mat.MaxInt; count++ {
		var annotatedAssociation ast.AnnotatedAssociationLike
		annotatedAssociation, token, ok = v.parseAnnotatedAssociation()
		if !ok {
			switch {
			case count >= 1:
				break annotatedAssociationsLoop
			case uti.IsDefined(tokens):
				// This is not multiple AnnotatedAssociation rules.
				v.putBack(tokens)
				return
			default:
				// Found a syntax error.
				var message = v.formatError("$MultilineAttributes", token)
				message += "1 or more AnnotatedAssociation rules are required."
				panic(message)
			}
		}
		// No additional put backs allowed at this point.
		tokens = nil
		annotatedAssociations.AppendValue(annotatedAssociation)
	}

	// Attempt to parse a single "]" delimiter.
	_, token, ok = v.parseDelimiter("]")
	if !ok {
		if uti.IsDefined(tokens) {
			// This is not a single MultilineAttributes rule.
			v.putBack(tokens)
			return
		} else {
			// Found a syntax error.
			var message = v.formatError("$MultilineAttributes", token)
			panic(message)
		}
	}
	if uti.IsDefined(tokens) {
		tokens.AppendValue(token)
	}

	// Found a single MultilineAttributes rule.
	ok = true
	v.remove(tokens)
	multilineAttributes = ast.MultilineAttributesClass().MultilineAttributes(
		newline,
		annotatedAssociations,
	)
	return
}

func (v *parser_) parseMultilineParameters() (
	multilineParameters ast.MultilineParametersLike,
	token TokenLike,
	ok bool,
) {
	var tokens = col.List[TokenLike]()

	// Attempt to parse a single "(" delimiter.
	_, token, ok = v.parseDelimiter("(")
	if !ok {
		if uti.IsDefined(tokens) {
			// This is not a single MultilineParameters rule.
			v.putBack(tokens)
			return
		} else {
			// Found a syntax error.
			var message = v.formatError("$MultilineParameters", token)
			panic(message)
		}
	}
	if uti.IsDefined(tokens) {
		tokens.AppendValue(token)
	}

	// Attempt to parse a single newline token.
	var newline string
	newline, token, ok = v.parseToken(NewlineToken)
	if !ok {
		if uti.IsDefined(tokens) {
			// This is not a single newline token.
			v.putBack(tokens)
			return
		} else {
			// Found a syntax error.
			var message = v.formatError("$MultilineParameters", token)
			panic(message)
		}
	}
	if uti.IsDefined(tokens) {
		tokens.AppendValue(token)
	}

	// Attempt to parse multiple AnnotatedAssociation rules.
	var annotatedAssociations = col.List[ast.AnnotatedAssociationLike]()
annotatedAssociationsLoop:
	for count := 0; count < mat.MaxInt; count++ {
		var annotatedAssociation ast.AnnotatedAssociationLike
		annotatedAssociation, token, ok = v.parseAnnotatedAssociation()
		if !ok {
			switch {
			case count >= 1:
				break annotatedAssociationsLoop
			case uti.IsDefined(tokens):
				// This is not multiple AnnotatedAssociation rules.
				v.putBack(tokens)
				return
			default:
				// Found a syntax error.
				var message = v.formatError("$MultilineParameters", token)
				message += "1 or more AnnotatedAssociation rules are required."
				panic(message)
			}
		}
		// No additional put backs allowed at this point.
		tokens = nil
		annotatedAssociations.AppendValue(annotatedAssociation)
	}

	// Attempt to parse a single ")" delimiter.
	_, token, ok = v.parseDelimiter(")")
	if !ok {
		if uti.IsDefined(tokens) {
			// This is not a single MultilineParameters rule.
			v.putBack(tokens)
			return
		} else {
			// Found a syntax error.
			var message = v.formatError("$MultilineParameters", token)
			panic(message)
		}
	}
	if uti.IsDefined(tokens) {
		tokens.AppendValue(token)
	}

	// Found a single MultilineParameters rule.
	ok = true
	v.remove(tokens)
	multilineParameters = ast.MultilineParametersClass().MultilineParameters(
		newline,
		annotatedAssociations,
	)
	return
}

func (v *parser_) parseMultilineStatements() (
	multilineStatements ast.MultilineStatementsLike,
	token TokenLike,
	ok bool,
) {
	var tokens = col.List[TokenLike]()

	// Attempt to parse a single "{" delimiter.
	_, token, ok = v.parseDelimiter("{")
	if !ok {
		if uti.IsDefined(tokens) {
			// This is not a single MultilineStatements rule.
			v.putBack(tokens)
			return
		} else {
			// Found a syntax error.
			var message = v.formatError("$MultilineStatements", token)
			panic(message)
		}
	}
	if uti.IsDefined(tokens) {
		tokens.AppendValue(token)
	}

	// Attempt to parse a single newline token.
	var newline string
	newline, token, ok = v.parseToken(NewlineToken)
	if !ok {
		if uti.IsDefined(tokens) {
			// This is not a single newline token.
			v.putBack(tokens)
			return
		} else {
			// Found a syntax error.
			var message = v.formatError("$MultilineStatements", token)
			panic(message)
		}
	}
	if uti.IsDefined(tokens) {
		tokens.AppendValue(token)
	}

	// Attempt to parse multiple AnnotatedStatement rules.
	var annotatedStatements = col.List[ast.AnnotatedStatementLike]()
annotatedStatementsLoop:
	for count := 0; count < mat.MaxInt; count++ {
		var annotatedStatement ast.AnnotatedStatementLike
		annotatedStatement, token, ok = v.parseAnnotatedStatement()
		if !ok {
			switch {
			case count >= 0:
				break annotatedStatementsLoop
			case uti.IsDefined(tokens):
				// This is not multiple AnnotatedStatement rules.
				v.putBack(tokens)
				return
			default:
				// Found a syntax error.
				var message = v.formatError("$MultilineStatements", token)
				message += "0 or more AnnotatedStatement rules are required."
				panic(message)
			}
		}
		// No additional put backs allowed at this point.
		tokens = nil
		annotatedStatements.AppendValue(annotatedStatement)
	}

	// Attempt to parse a single "}" delimiter.
	_, token, ok = v.parseDelimiter("}")
	if !ok {
		if uti.IsDefined(tokens) {
			// This is not a single MultilineStatements rule.
			v.putBack(tokens)
			return
		} else {
			// Found a syntax error.
			var message = v.formatError("$MultilineStatements", token)
			panic(message)
		}
	}
	if uti.IsDefined(tokens) {
		tokens.AppendValue(token)
	}

	// Found a single MultilineStatements rule.
	ok = true
	v.remove(tokens)
	multilineStatements = ast.MultilineStatementsClass().MultilineStatements(
		newline,
		annotatedStatements,
	)
	return
}

func (v *parser_) parseMultilineValues() (
	multilineValues ast.MultilineValuesLike,
	token TokenLike,
	ok bool,
) {
	var tokens = col.List[TokenLike]()

	// Attempt to parse a single "[" delimiter.
	_, token, ok = v.parseDelimiter("[")
	if !ok {
		if uti.IsDefined(tokens) {
			// This is not a single MultilineValues rule.
			v.putBack(tokens)
			return
		} else {
			// Found a syntax error.
			var message = v.formatError("$MultilineValues", token)
			panic(message)
		}
	}
	if uti.IsDefined(tokens) {
		tokens.AppendValue(token)
	}

	// Attempt to parse a single newline token.
	var newline string
	newline, token, ok = v.parseToken(NewlineToken)
	if !ok {
		if uti.IsDefined(tokens) {
			// This is not a single newline token.
			v.putBack(tokens)
			return
		} else {
			// Found a syntax error.
			var message = v.formatError("$MultilineValues", token)
			panic(message)
		}
	}
	if uti.IsDefined(tokens) {
		tokens.AppendValue(token)
	}

	// Attempt to parse multiple AnnotatedValue rules.
	var annotatedValues = col.List[ast.AnnotatedValueLike]()
annotatedValuesLoop:
	for count := 0; count < mat.MaxInt; count++ {
		var annotatedValue ast.AnnotatedValueLike
		annotatedValue, token, ok = v.parseAnnotatedValue()
		if !ok {
			switch {
			case count >= 1:
				break annotatedValuesLoop
			case uti.IsDefined(tokens):
				// This is not multiple AnnotatedValue rules.
				v.putBack(tokens)
				return
			default:
				// Found a syntax error.
				var message = v.formatError("$MultilineValues", token)
				message += "1 or more AnnotatedValue rules are required."
				panic(message)
			}
		}
		// No additional put backs allowed at this point.
		tokens = nil
		annotatedValues.AppendValue(annotatedValue)
	}

	// Attempt to parse a single "]" delimiter.
	_, token, ok = v.parseDelimiter("]")
	if !ok {
		if uti.IsDefined(tokens) {
			// This is not a single MultilineValues rule.
			v.putBack(tokens)
			return
		} else {
			// Found a syntax error.
			var message = v.formatError("$MultilineValues", token)
			panic(message)
		}
	}
	if uti.IsDefined(tokens) {
		tokens.AppendValue(token)
	}

	// Found a single MultilineValues rule.
	ok = true
	v.remove(tokens)
	multilineValues = ast.MultilineValuesClass().MultilineValues(
		newline,
		annotatedValues,
	)
	return
}

func (v *parser_) parseNoAttributes() (
	noAttributes ast.NoAttributesLike,
	token TokenLike,
	ok bool,
) {
	var tokens = col.List[TokenLike]()

	// Attempt to parse a single "[" delimiter.
	_, token, ok = v.parseDelimiter("[")
	if !ok {
		if uti.IsDefined(tokens) {
			// This is not a single NoAttributes rule.
			v.putBack(tokens)
			return
		} else {
			// Found a syntax error.
			var message = v.formatError("$NoAttributes", token)
			panic(message)
		}
	}
	if uti.IsDefined(tokens) {
		tokens.AppendValue(token)
	}

	// Attempt to parse a single ":" delimiter.
	_, token, ok = v.parseDelimiter(":")
	if !ok {
		if uti.IsDefined(tokens) {
			// This is not a single NoAttributes rule.
			v.putBack(tokens)
			return
		} else {
			// Found a syntax error.
			var message = v.formatError("$NoAttributes", token)
			panic(message)
		}
	}
	if uti.IsDefined(tokens) {
		tokens.AppendValue(token)
	}

	// Attempt to parse a single "]" delimiter.
	_, token, ok = v.parseDelimiter("]")
	if !ok {
		if uti.IsDefined(tokens) {
			// This is not a single NoAttributes rule.
			v.putBack(tokens)
			return
		} else {
			// Found a syntax error.
			var message = v.formatError("$NoAttributes", token)
			panic(message)
		}
	}
	if uti.IsDefined(tokens) {
		tokens.AppendValue(token)
	}

	// Found a single NoAttributes rule.
	ok = true
	v.remove(tokens)
	noAttributes = ast.NoAttributesClass().NoAttributes()
	return
}

func (v *parser_) parseNoStatements() (
	noStatements ast.NoStatementsLike,
	token TokenLike,
	ok bool,
) {
	var tokens = col.List[TokenLike]()

	// Attempt to parse a single "{" delimiter.
	_, token, ok = v.parseDelimiter("{")
	if !ok {
		if uti.IsDefined(tokens) {
			// This is not a single NoStatements rule.
			v.putBack(tokens)
			return
		} else {
			// Found a syntax error.
			var message = v.formatError("$NoStatements", token)
			panic(message)
		}
	}
	if uti.IsDefined(tokens) {
		tokens.AppendValue(token)
	}

	// Attempt to parse a single "}" delimiter.
	_, token, ok = v.parseDelimiter("}")
	if !ok {
		if uti.IsDefined(tokens) {
			// This is not a single NoStatements rule.
			v.putBack(tokens)
			return
		} else {
			// Found a syntax error.
			var message = v.formatError("$NoStatements", token)
			panic(message)
		}
	}
	if uti.IsDefined(tokens) {
		tokens.AppendValue(token)
	}

	// Found a single NoStatements rule.
	ok = true
	v.remove(tokens)
	noStatements = ast.NoStatementsClass().NoStatements()
	return
}

func (v *parser_) parseNoValues() (
	noValues ast.NoValuesLike,
	token TokenLike,
	ok bool,
) {
	var tokens = col.List[TokenLike]()

	// Attempt to parse a single "[" delimiter.
	_, token, ok = v.parseDelimiter("[")
	if !ok {
		if uti.IsDefined(tokens) {
			// This is not a single NoValues rule.
			v.putBack(tokens)
			return
		} else {
			// Found a syntax error.
			var message = v.formatError("$NoValues", token)
			panic(message)
		}
	}
	if uti.IsDefined(tokens) {
		tokens.AppendValue(token)
	}

	// Attempt to parse a single "]" delimiter.
	_, token, ok = v.parseDelimiter("]")
	if !ok {
		if uti.IsDefined(tokens) {
			// This is not a single NoValues rule.
			v.putBack(tokens)
			return
		} else {
			// Found a syntax error.
			var message = v.formatError("$NoValues", token)
			panic(message)
		}
	}
	if uti.IsDefined(tokens) {
		tokens.AppendValue(token)
	}

	// Found a single NoValues rule.
	ok = true
	v.remove(tokens)
	noValues = ast.NoValuesClass().NoValues()
	return
}

func (v *parser_) parseNotarizeClause() (
	notarizeClause ast.NotarizeClauseLike,
	token TokenLike,
	ok bool,
) {
	var tokens = col.List[TokenLike]()

	// Attempt to parse a single "notarize" delimiter.
	_, token, ok = v.parseDelimiter("notarize")
	if !ok {
		if uti.IsDefined(tokens) {
			// This is not a single NotarizeClause rule.
			v.putBack(tokens)
			return
		} else {
			// Found a syntax error.
			var message = v.formatError("$NotarizeClause", token)
			panic(message)
		}
	}
	if uti.IsDefined(tokens) {
		tokens.AppendValue(token)
	}

	// Attempt to parse a single Draft rule.
	var draft ast.DraftLike
	draft, token, ok = v.parseDraft()
	switch {
	case ok:
		// No additional put backs allowed at this point.
		tokens = nil
	case uti.IsDefined(tokens):
		// This is not a single Draft rule.
		v.putBack(tokens)
		return
	default:
		// Found a syntax error.
		var message = v.formatError("$NotarizeClause", token)
		panic(message)
	}

	// Attempt to parse a single "as" delimiter.
	_, token, ok = v.parseDelimiter("as")
	if !ok {
		if uti.IsDefined(tokens) {
			// This is not a single NotarizeClause rule.
			v.putBack(tokens)
			return
		} else {
			// Found a syntax error.
			var message = v.formatError("$NotarizeClause", token)
			panic(message)
		}
	}
	if uti.IsDefined(tokens) {
		tokens.AppendValue(token)
	}

	// Attempt to parse a single Cited rule.
	var cited ast.CitedLike
	cited, token, ok = v.parseCited()
	switch {
	case ok:
		// No additional put backs allowed at this point.
		tokens = nil
	case uti.IsDefined(tokens):
		// This is not a single Cited rule.
		v.putBack(tokens)
		return
	default:
		// Found a syntax error.
		var message = v.formatError("$NotarizeClause", token)
		panic(message)
	}

	// Found a single NotarizeClause rule.
	ok = true
	v.remove(tokens)
	notarizeClause = ast.NotarizeClauseClass().NotarizeClause(
		draft,
		cited,
	)
	return
}

func (v *parser_) parseNotice() (
	notice ast.NoticeLike,
	token TokenLike,
	ok bool,
) {
	var tokens = col.List[TokenLike]()

	// Attempt to parse a single comment token.
	var comment string
	comment, token, ok = v.parseToken(CommentToken)
	if !ok {
		if uti.IsDefined(tokens) {
			// This is not a single comment token.
			v.putBack(tokens)
			return
		} else {
			// Found a syntax error.
			var message = v.formatError("$Notice", token)
			panic(message)
		}
	}
	if uti.IsDefined(tokens) {
		tokens.AppendValue(token)
	}

	// Attempt to parse a single newline token.
	var newline string
	newline, token, ok = v.parseToken(NewlineToken)
	if !ok {
		if uti.IsDefined(tokens) {
			// This is not a single newline token.
			v.putBack(tokens)
			return
		} else {
			// Found a syntax error.
			var message = v.formatError("$Notice", token)
			panic(message)
		}
	}
	if uti.IsDefined(tokens) {
		tokens.AppendValue(token)
	}

	// Found a single Notice rule.
	ok = true
	v.remove(tokens)
	notice = ast.NoticeClass().Notice(
		comment,
		newline,
	)
	return
}

func (v *parser_) parseNumerical() (
	numerical ast.NumericalLike,
	token TokenLike,
	ok bool,
) {
	// Attempt to parse a single Component Numerical.
	var component ast.ComponentLike
	component, token, ok = v.parseComponent()
	if ok {
		// Found a single Component Numerical.
		numerical = ast.NumericalClass().Numerical(component)
		return
	}

	// Attempt to parse a single Subcomponent Numerical.
	var subcomponent ast.SubcomponentLike
	subcomponent, token, ok = v.parseSubcomponent()
	if ok {
		// Found a single Subcomponent Numerical.
		numerical = ast.NumericalClass().Numerical(subcomponent)
		return
	}

	// Attempt to parse a single Precedence Numerical.
	var precedence ast.PrecedenceLike
	precedence, token, ok = v.parsePrecedence()
	if ok {
		// Found a single Precedence Numerical.
		numerical = ast.NumericalClass().Numerical(precedence)
		return
	}

	// Attempt to parse a single Referent Numerical.
	var referent ast.ReferentLike
	referent, token, ok = v.parseReferent()
	if ok {
		// Found a single Referent Numerical.
		numerical = ast.NumericalClass().Numerical(referent)
		return
	}

	// Attempt to parse a single Inversion Numerical.
	var inversion ast.InversionLike
	inversion, token, ok = v.parseInversion()
	if ok {
		// Found a single Inversion Numerical.
		numerical = ast.NumericalClass().Numerical(inversion)
		return
	}

	// Attempt to parse a single Magnitude Numerical.
	var magnitude ast.MagnitudeLike
	magnitude, token, ok = v.parseMagnitude()
	if ok {
		// Found a single Magnitude Numerical.
		numerical = ast.NumericalClass().Numerical(magnitude)
		return
	}

	// Attempt to parse a single Function Numerical.
	var function ast.FunctionLike
	function, token, ok = v.parseFunction()
	if ok {
		// Found a single Function Numerical.
		numerical = ast.NumericalClass().Numerical(function)
		return
	}

	// Attempt to parse a single Method Numerical.
	var method ast.MethodLike
	method, token, ok = v.parseMethod()
	if ok {
		// Found a single Method Numerical.
		numerical = ast.NumericalClass().Numerical(method)
		return
	}

	// Attempt to parse a single Variable Numerical.
	var variable ast.VariableLike
	variable, token, ok = v.parseVariable()
	if ok {
		// Found a single Variable Numerical.
		numerical = ast.NumericalClass().Numerical(variable)
		return
	}

	// This is not a single Numerical rule.
	return
}

func (v *parser_) parseOnClause() (
	onClause ast.OnClauseLike,
	token TokenLike,
	ok bool,
) {
	var tokens = col.List[TokenLike]()

	// Attempt to parse a single "on" delimiter.
	_, token, ok = v.parseDelimiter("on")
	if !ok {
		if uti.IsDefined(tokens) {
			// This is not a single OnClause rule.
			v.putBack(tokens)
			return
		} else {
			// Found a syntax error.
			var message = v.formatError("$OnClause", token)
			panic(message)
		}
	}
	if uti.IsDefined(tokens) {
		tokens.AppendValue(token)
	}

	// Attempt to parse a single Failure rule.
	var failure ast.FailureLike
	failure, token, ok = v.parseFailure()
	switch {
	case ok:
		// No additional put backs allowed at this point.
		tokens = nil
	case uti.IsDefined(tokens):
		// This is not a single Failure rule.
		v.putBack(tokens)
		return
	default:
		// Found a syntax error.
		var message = v.formatError("$OnClause", token)
		panic(message)
	}

	// Attempt to parse multiple MatchHandler rules.
	var matchHandlers = col.List[ast.MatchHandlerLike]()
matchHandlersLoop:
	for count := 0; count < mat.MaxInt; count++ {
		var matchHandler ast.MatchHandlerLike
		matchHandler, token, ok = v.parseMatchHandler()
		if !ok {
			switch {
			case count >= 1:
				break matchHandlersLoop
			case uti.IsDefined(tokens):
				// This is not multiple MatchHandler rules.
				v.putBack(tokens)
				return
			default:
				// Found a syntax error.
				var message = v.formatError("$OnClause", token)
				message += "1 or more MatchHandler rules are required."
				panic(message)
			}
		}
		// No additional put backs allowed at this point.
		tokens = nil
		matchHandlers.AppendValue(matchHandler)
	}

	// Found a single OnClause rule.
	ok = true
	v.remove(tokens)
	onClause = ast.OnClauseClass().OnClause(
		failure,
		matchHandlers,
	)
	return
}

func (v *parser_) parseOperation() (
	operation ast.OperationLike,
	token TokenLike,
	ok bool,
) {
	// Attempt to parse a single Is Operation.
	var is ast.IsLike
	is, token, ok = v.parseIs()
	if ok {
		// Found a single Is Operation.
		operation = ast.OperationClass().Operation(is)
		return
	}

	// Attempt to parse a single Matches Operation.
	var matches ast.MatchesLike
	matches, token, ok = v.parseMatches()
	if ok {
		// Found a single Matches Operation.
		operation = ast.OperationClass().Operation(matches)
		return
	}

	// Attempt to parse a single And Operation.
	var and ast.AndLike
	and, token, ok = v.parseAnd()
	if ok {
		// Found a single And Operation.
		operation = ast.OperationClass().Operation(and)
		return
	}

	// Attempt to parse a single San Operation.
	var san ast.SanLike
	san, token, ok = v.parseSan()
	if ok {
		// Found a single San Operation.
		operation = ast.OperationClass().Operation(san)
		return
	}

	// Attempt to parse a single Ior Operation.
	var ior ast.IorLike
	ior, token, ok = v.parseIor()
	if ok {
		// Found a single Ior Operation.
		operation = ast.OperationClass().Operation(ior)
		return
	}

	// Attempt to parse a single Xor Operation.
	var xor ast.XorLike
	xor, token, ok = v.parseXor()
	if ok {
		// Found a single Xor Operation.
		operation = ast.OperationClass().Operation(xor)
		return
	}

	// This is not a single Operation rule.
	return
}

func (v *parser_) parseOperator() (
	operator ast.OperatorLike,
	token TokenLike,
	ok bool,
) {
	// Attempt to parse a single plus Operator.
	var plus string
	plus, token, ok = v.parseToken(PlusToken)
	if ok {
		// Found a single plus Operator.
		operator = ast.OperatorClass().Operator(plus)
		return
	}

	// Attempt to parse a single dash Operator.
	var dash string
	dash, token, ok = v.parseToken(DashToken)
	if ok {
		// Found a single dash Operator.
		operator = ast.OperatorClass().Operator(dash)
		return
	}

	// Attempt to parse a single star Operator.
	var star string
	star, token, ok = v.parseToken(StarToken)
	if ok {
		// Found a single star Operator.
		operator = ast.OperatorClass().Operator(star)
		return
	}

	// Attempt to parse a single slash Operator.
	var slash string
	slash, token, ok = v.parseToken(SlashToken)
	if ok {
		// Found a single slash Operator.
		operator = ast.OperatorClass().Operator(slash)
		return
	}

	// Attempt to parse a single double Operator.
	var double string
	double, token, ok = v.parseToken(DoubleToken)
	if ok {
		// Found a single double Operator.
		operator = ast.OperatorClass().Operator(double)
		return
	}

	// Attempt to parse a single caret Operator.
	var caret string
	caret, token, ok = v.parseToken(CaretToken)
	if ok {
		// Found a single caret Operator.
		operator = ast.OperatorClass().Operator(caret)
		return
	}

	// Attempt to parse a single ampersand Operator.
	var ampersand string
	ampersand, token, ok = v.parseToken(AmpersandToken)
	if ok {
		// Found a single ampersand Operator.
		operator = ast.OperatorClass().Operator(ampersand)
		return
	}

	// Attempt to parse a single less Operator.
	var less string
	less, token, ok = v.parseToken(LessToken)
	if ok {
		// Found a single less Operator.
		operator = ast.OperatorClass().Operator(less)
		return
	}

	// Attempt to parse a single equal Operator.
	var equal string
	equal, token, ok = v.parseToken(EqualToken)
	if ok {
		// Found a single equal Operator.
		operator = ast.OperatorClass().Operator(equal)
		return
	}

	// Attempt to parse a single more Operator.
	var more string
	more, token, ok = v.parseToken(MoreToken)
	if ok {
		// Found a single more Operator.
		operator = ast.OperatorClass().Operator(more)
		return
	}

	// This is not a single Operator rule.
	return
}

func (v *parser_) parseParameters() (
	parameters ast.ParametersLike,
	token TokenLike,
	ok bool,
) {
	// Attempt to parse a single MultilineParameters Parameters.
	var multilineParameters ast.MultilineParametersLike
	multilineParameters, token, ok = v.parseMultilineParameters()
	if ok {
		// Found a single MultilineParameters Parameters.
		parameters = ast.ParametersClass().Parameters(multilineParameters)
		return
	}

	// Attempt to parse a single InlineParameters Parameters.
	var inlineParameters ast.InlineParametersLike
	inlineParameters, token, ok = v.parseInlineParameters()
	if ok {
		// Found a single InlineParameters Parameters.
		parameters = ast.ParametersClass().Parameters(inlineParameters)
		return
	}

	// This is not a single Parameters rule.
	return
}

func (v *parser_) parsePlusEqual() (
	plusEqual ast.PlusEqualLike,
	token TokenLike,
	ok bool,
) {
	var tokens = col.List[TokenLike]()

	// Attempt to parse a single "+=" delimiter.
	_, token, ok = v.parseDelimiter("+=")
	if !ok {
		if uti.IsDefined(tokens) {
			// This is not a single PlusEqual rule.
			v.putBack(tokens)
			return
		} else {
			// Found a syntax error.
			var message = v.formatError("$PlusEqual", token)
			panic(message)
		}
	}
	if uti.IsDefined(tokens) {
		tokens.AppendValue(token)
	}

	// Found a single PlusEqual rule.
	ok = true
	v.remove(tokens)
	plusEqual = ast.PlusEqualClass().PlusEqual()
	return
}

func (v *parser_) parsePostClause() (
	postClause ast.PostClauseLike,
	token TokenLike,
	ok bool,
) {
	var tokens = col.List[TokenLike]()

	// Attempt to parse a single "post" delimiter.
	_, token, ok = v.parseDelimiter("post")
	if !ok {
		if uti.IsDefined(tokens) {
			// This is not a single PostClause rule.
			v.putBack(tokens)
			return
		} else {
			// Found a syntax error.
			var message = v.formatError("$PostClause", token)
			panic(message)
		}
	}
	if uti.IsDefined(tokens) {
		tokens.AppendValue(token)
	}

	// Attempt to parse a single Message rule.
	var message ast.MessageLike
	message, token, ok = v.parseMessage()
	switch {
	case ok:
		// No additional put backs allowed at this point.
		tokens = nil
	case uti.IsDefined(tokens):
		// This is not a single Message rule.
		v.putBack(tokens)
		return
	default:
		// Found a syntax error.
		var message = v.formatError("$PostClause", token)
		panic(message)
	}

	// Attempt to parse a single "to" delimiter.
	_, token, ok = v.parseDelimiter("to")
	if !ok {
		if uti.IsDefined(tokens) {
			// This is not a single PostClause rule.
			v.putBack(tokens)
			return
		} else {
			// Found a syntax error.
			var message = v.formatError("$PostClause", token)
			panic(message)
		}
	}
	if uti.IsDefined(tokens) {
		tokens.AppendValue(token)
	}

	// Attempt to parse a single Bag rule.
	var bag ast.BagLike
	bag, token, ok = v.parseBag()
	switch {
	case ok:
		// No additional put backs allowed at this point.
		tokens = nil
	case uti.IsDefined(tokens):
		// This is not a single Bag rule.
		v.putBack(tokens)
		return
	default:
		// Found a syntax error.
		var message = v.formatError("$PostClause", token)
		panic(message)
	}

	// Found a single PostClause rule.
	ok = true
	v.remove(tokens)
	postClause = ast.PostClauseClass().PostClause(
		message,
		bag,
	)
	return
}

func (v *parser_) parsePrecedence() (
	precedence ast.PrecedenceLike,
	token TokenLike,
	ok bool,
) {
	var tokens = col.List[TokenLike]()

	// Attempt to parse a single "(" delimiter.
	_, token, ok = v.parseDelimiter("(")
	if !ok {
		if uti.IsDefined(tokens) {
			// This is not a single Precedence rule.
			v.putBack(tokens)
			return
		} else {
			// Found a syntax error.
			var message = v.formatError("$Precedence", token)
			panic(message)
		}
	}
	if uti.IsDefined(tokens) {
		tokens.AppendValue(token)
	}

	// Attempt to parse a single Expression rule.
	var expression ast.ExpressionLike
	expression, token, ok = v.parseExpression()
	switch {
	case ok:
		// No additional put backs allowed at this point.
		tokens = nil
	case uti.IsDefined(tokens):
		// This is not a single Expression rule.
		v.putBack(tokens)
		return
	default:
		// Found a syntax error.
		var message = v.formatError("$Precedence", token)
		panic(message)
	}

	// Attempt to parse a single ")" delimiter.
	_, token, ok = v.parseDelimiter(")")
	if !ok {
		if uti.IsDefined(tokens) {
			// This is not a single Precedence rule.
			v.putBack(tokens)
			return
		} else {
			// Found a syntax error.
			var message = v.formatError("$Precedence", token)
			panic(message)
		}
	}
	if uti.IsDefined(tokens) {
		tokens.AppendValue(token)
	}

	// Found a single Precedence rule.
	ok = true
	v.remove(tokens)
	precedence = ast.PrecedenceClass().Precedence(expression)
	return
}

func (v *parser_) parsePredicate() (
	predicate ast.PredicateLike,
	token TokenLike,
	ok bool,
) {
	var tokens = col.List[TokenLike]()

	// Attempt to parse a single Action rule.
	var action ast.ActionLike
	action, token, ok = v.parseAction()
	switch {
	case ok:
		// No additional put backs allowed at this point.
		tokens = nil
	case uti.IsDefined(tokens):
		// This is not a single Action rule.
		v.putBack(tokens)
		return
	default:
		// Found a syntax error.
		var message = v.formatError("$Predicate", token)
		panic(message)
	}

	// Attempt to parse a single Expression rule.
	var expression ast.ExpressionLike
	expression, token, ok = v.parseExpression()
	switch {
	case ok:
		// No additional put backs allowed at this point.
		tokens = nil
	case uti.IsDefined(tokens):
		// This is not a single Expression rule.
		v.putBack(tokens)
		return
	default:
		// Found a syntax error.
		var message = v.formatError("$Predicate", token)
		panic(message)
	}

	// Found a single Predicate rule.
	ok = true
	v.remove(tokens)
	predicate = ast.PredicateClass().Predicate(
		action,
		expression,
	)
	return
}

func (v *parser_) parsePrimitive() (
	primitive ast.PrimitiveLike,
	token TokenLike,
	ok bool,
) {
	// Attempt to parse a single Element Primitive.
	var element ast.ElementLike
	element, token, ok = v.parseElement()
	if ok {
		// Found a single Element Primitive.
		primitive = ast.PrimitiveClass().Primitive(element)
		return
	}

	// Attempt to parse a single String Primitive.
	var string_ ast.StringLike
	string_, token, ok = v.parseString()
	if ok {
		// Found a single String Primitive.
		primitive = ast.PrimitiveClass().Primitive(string_)
		return
	}

	// This is not a single Primitive rule.
	return
}

func (v *parser_) parseProcedure() (
	procedure ast.ProcedureLike,
	token TokenLike,
	ok bool,
) {
	// Attempt to parse a single MultilineStatements Procedure.
	var multilineStatements ast.MultilineStatementsLike
	multilineStatements, token, ok = v.parseMultilineStatements()
	if ok {
		// Found a single MultilineStatements Procedure.
		procedure = ast.ProcedureClass().Procedure(multilineStatements)
		return
	}

	// Attempt to parse a single InlineStatements Procedure.
	var inlineStatements ast.InlineStatementsLike
	inlineStatements, token, ok = v.parseInlineStatements()
	if ok {
		// Found a single InlineStatements Procedure.
		procedure = ast.ProcedureClass().Procedure(inlineStatements)
		return
	}

	// Attempt to parse a single NoStatements Procedure.
	var noStatements ast.NoStatementsLike
	noStatements, token, ok = v.parseNoStatements()
	if ok {
		// Found a single NoStatements Procedure.
		procedure = ast.ProcedureClass().Procedure(noStatements)
		return
	}

	// This is not a single Procedure rule.
	return
}

func (v *parser_) parsePublishClause() (
	publishClause ast.PublishClauseLike,
	token TokenLike,
	ok bool,
) {
	var tokens = col.List[TokenLike]()

	// Attempt to parse a single "publish" delimiter.
	_, token, ok = v.parseDelimiter("publish")
	if !ok {
		if uti.IsDefined(tokens) {
			// This is not a single PublishClause rule.
			v.putBack(tokens)
			return
		} else {
			// Found a syntax error.
			var message = v.formatError("$PublishClause", token)
			panic(message)
		}
	}
	if uti.IsDefined(tokens) {
		tokens.AppendValue(token)
	}

	// Attempt to parse a single Event rule.
	var event ast.EventLike
	event, token, ok = v.parseEvent()
	switch {
	case ok:
		// No additional put backs allowed at this point.
		tokens = nil
	case uti.IsDefined(tokens):
		// This is not a single Event rule.
		v.putBack(tokens)
		return
	default:
		// Found a syntax error.
		var message = v.formatError("$PublishClause", token)
		panic(message)
	}

	// Found a single PublishClause rule.
	ok = true
	v.remove(tokens)
	publishClause = ast.PublishClauseClass().PublishClause(event)
	return
}

func (v *parser_) parseRecipient() (
	recipient ast.RecipientLike,
	token TokenLike,
	ok bool,
) {
	// Attempt to parse a single Item Recipient.
	var item ast.ItemLike
	item, token, ok = v.parseItem()
	if ok {
		// Found a single Item Recipient.
		recipient = ast.RecipientClass().Recipient(item)
		return
	}

	// Attempt to parse a single Subcomponent Recipient.
	var subcomponent ast.SubcomponentLike
	subcomponent, token, ok = v.parseSubcomponent()
	if ok {
		// Found a single Subcomponent Recipient.
		recipient = ast.RecipientClass().Recipient(subcomponent)
		return
	}

	// This is not a single Recipient rule.
	return
}

func (v *parser_) parseReferent() (
	referent ast.ReferentLike,
	token TokenLike,
	ok bool,
) {
	var tokens = col.List[TokenLike]()

	// Attempt to parse a single "@" delimiter.
	_, token, ok = v.parseDelimiter("@")
	if !ok {
		if uti.IsDefined(tokens) {
			// This is not a single Referent rule.
			v.putBack(tokens)
			return
		} else {
			// Found a syntax error.
			var message = v.formatError("$Referent", token)
			panic(message)
		}
	}
	if uti.IsDefined(tokens) {
		tokens.AppendValue(token)
	}

	// Attempt to parse a single Indirect rule.
	var indirect ast.IndirectLike
	indirect, token, ok = v.parseIndirect()
	switch {
	case ok:
		// No additional put backs allowed at this point.
		tokens = nil
	case uti.IsDefined(tokens):
		// This is not a single Indirect rule.
		v.putBack(tokens)
		return
	default:
		// Found a syntax error.
		var message = v.formatError("$Referent", token)
		panic(message)
	}

	// Found a single Referent rule.
	ok = true
	v.remove(tokens)
	referent = ast.ReferentClass().Referent(indirect)
	return
}

func (v *parser_) parseRejectClause() (
	rejectClause ast.RejectClauseLike,
	token TokenLike,
	ok bool,
) {
	var tokens = col.List[TokenLike]()

	// Attempt to parse a single "reject" delimiter.
	_, token, ok = v.parseDelimiter("reject")
	if !ok {
		if uti.IsDefined(tokens) {
			// This is not a single RejectClause rule.
			v.putBack(tokens)
			return
		} else {
			// Found a syntax error.
			var message = v.formatError("$RejectClause", token)
			panic(message)
		}
	}
	if uti.IsDefined(tokens) {
		tokens.AppendValue(token)
	}

	// Attempt to parse a single Message rule.
	var message ast.MessageLike
	message, token, ok = v.parseMessage()
	switch {
	case ok:
		// No additional put backs allowed at this point.
		tokens = nil
	case uti.IsDefined(tokens):
		// This is not a single Message rule.
		v.putBack(tokens)
		return
	default:
		// Found a syntax error.
		var message = v.formatError("$RejectClause", token)
		panic(message)
	}

	// Found a single RejectClause rule.
	ok = true
	v.remove(tokens)
	rejectClause = ast.RejectClauseClass().RejectClause(message)
	return
}

func (v *parser_) parseRepository() (
	repository ast.RepositoryLike,
	token TokenLike,
	ok bool,
) {
	// Attempt to parse a single CheckoutClause Repository.
	var checkoutClause ast.CheckoutClauseLike
	checkoutClause, token, ok = v.parseCheckoutClause()
	if ok {
		// Found a single CheckoutClause Repository.
		repository = ast.RepositoryClass().Repository(checkoutClause)
		return
	}

	// Attempt to parse a single SaveClause Repository.
	var saveClause ast.SaveClauseLike
	saveClause, token, ok = v.parseSaveClause()
	if ok {
		// Found a single SaveClause Repository.
		repository = ast.RepositoryClass().Repository(saveClause)
		return
	}

	// Attempt to parse a single DiscardClause Repository.
	var discardClause ast.DiscardClauseLike
	discardClause, token, ok = v.parseDiscardClause()
	if ok {
		// Found a single DiscardClause Repository.
		repository = ast.RepositoryClass().Repository(discardClause)
		return
	}

	// Attempt to parse a single NotarizeClause Repository.
	var notarizeClause ast.NotarizeClauseLike
	notarizeClause, token, ok = v.parseNotarizeClause()
	if ok {
		// Found a single NotarizeClause Repository.
		repository = ast.RepositoryClass().Repository(notarizeClause)
		return
	}

	// This is not a single Repository rule.
	return
}

func (v *parser_) parseResult() (
	result ast.ResultLike,
	token TokenLike,
	ok bool,
) {
	var tokens = col.List[TokenLike]()

	// Attempt to parse a single Expression rule.
	var expression ast.ExpressionLike
	expression, token, ok = v.parseExpression()
	switch {
	case ok:
		// No additional put backs allowed at this point.
		tokens = nil
	case uti.IsDefined(tokens):
		// This is not a single Expression rule.
		v.putBack(tokens)
		return
	default:
		// Found a syntax error.
		var message = v.formatError("$Result", token)
		panic(message)
	}

	// Found a single Result rule.
	ok = true
	v.remove(tokens)
	result = ast.ResultClass().Result(expression)
	return
}

func (v *parser_) parseRetrieveClause() (
	retrieveClause ast.RetrieveClauseLike,
	token TokenLike,
	ok bool,
) {
	var tokens = col.List[TokenLike]()

	// Attempt to parse a single "retrieve" delimiter.
	_, token, ok = v.parseDelimiter("retrieve")
	if !ok {
		if uti.IsDefined(tokens) {
			// This is not a single RetrieveClause rule.
			v.putBack(tokens)
			return
		} else {
			// Found a syntax error.
			var message = v.formatError("$RetrieveClause", token)
			panic(message)
		}
	}
	if uti.IsDefined(tokens) {
		tokens.AppendValue(token)
	}

	// Attempt to parse a single Recipient rule.
	var recipient ast.RecipientLike
	recipient, token, ok = v.parseRecipient()
	switch {
	case ok:
		// No additional put backs allowed at this point.
		tokens = nil
	case uti.IsDefined(tokens):
		// This is not a single Recipient rule.
		v.putBack(tokens)
		return
	default:
		// Found a syntax error.
		var message = v.formatError("$RetrieveClause", token)
		panic(message)
	}

	// Attempt to parse a single "from" delimiter.
	_, token, ok = v.parseDelimiter("from")
	if !ok {
		if uti.IsDefined(tokens) {
			// This is not a single RetrieveClause rule.
			v.putBack(tokens)
			return
		} else {
			// Found a syntax error.
			var message = v.formatError("$RetrieveClause", token)
			panic(message)
		}
	}
	if uti.IsDefined(tokens) {
		tokens.AppendValue(token)
	}

	// Attempt to parse a single Bag rule.
	var bag ast.BagLike
	bag, token, ok = v.parseBag()
	switch {
	case ok:
		// No additional put backs allowed at this point.
		tokens = nil
	case uti.IsDefined(tokens):
		// This is not a single Bag rule.
		v.putBack(tokens)
		return
	default:
		// Found a syntax error.
		var message = v.formatError("$RetrieveClause", token)
		panic(message)
	}

	// Found a single RetrieveClause rule.
	ok = true
	v.remove(tokens)
	retrieveClause = ast.RetrieveClauseClass().RetrieveClause(
		recipient,
		bag,
	)
	return
}

func (v *parser_) parseReturnClause() (
	returnClause ast.ReturnClauseLike,
	token TokenLike,
	ok bool,
) {
	var tokens = col.List[TokenLike]()

	// Attempt to parse a single "return" delimiter.
	_, token, ok = v.parseDelimiter("return")
	if !ok {
		if uti.IsDefined(tokens) {
			// This is not a single ReturnClause rule.
			v.putBack(tokens)
			return
		} else {
			// Found a syntax error.
			var message = v.formatError("$ReturnClause", token)
			panic(message)
		}
	}
	if uti.IsDefined(tokens) {
		tokens.AppendValue(token)
	}

	// Attempt to parse a single Result rule.
	var result ast.ResultLike
	result, token, ok = v.parseResult()
	switch {
	case ok:
		// No additional put backs allowed at this point.
		tokens = nil
	case uti.IsDefined(tokens):
		// This is not a single Result rule.
		v.putBack(tokens)
		return
	default:
		// Found a syntax error.
		var message = v.formatError("$ReturnClause", token)
		panic(message)
	}

	// Found a single ReturnClause rule.
	ok = true
	v.remove(tokens)
	returnClause = ast.ReturnClauseClass().ReturnClause(result)
	return
}

func (v *parser_) parseSan() (
	san ast.SanLike,
	token TokenLike,
	ok bool,
) {
	var tokens = col.List[TokenLike]()

	// Attempt to parse a single "san" delimiter.
	_, token, ok = v.parseDelimiter("san")
	if !ok {
		if uti.IsDefined(tokens) {
			// This is not a single San rule.
			v.putBack(tokens)
			return
		} else {
			// Found a syntax error.
			var message = v.formatError("$San", token)
			panic(message)
		}
	}
	if uti.IsDefined(tokens) {
		tokens.AppendValue(token)
	}

	// Found a single San rule.
	ok = true
	v.remove(tokens)
	san = ast.SanClass().San()
	return
}

func (v *parser_) parseSaveClause() (
	saveClause ast.SaveClauseLike,
	token TokenLike,
	ok bool,
) {
	var tokens = col.List[TokenLike]()

	// Attempt to parse a single "save" delimiter.
	_, token, ok = v.parseDelimiter("save")
	if !ok {
		if uti.IsDefined(tokens) {
			// This is not a single SaveClause rule.
			v.putBack(tokens)
			return
		} else {
			// Found a syntax error.
			var message = v.formatError("$SaveClause", token)
			panic(message)
		}
	}
	if uti.IsDefined(tokens) {
		tokens.AppendValue(token)
	}

	// Attempt to parse a single Draft rule.
	var draft ast.DraftLike
	draft, token, ok = v.parseDraft()
	switch {
	case ok:
		// No additional put backs allowed at this point.
		tokens = nil
	case uti.IsDefined(tokens):
		// This is not a single Draft rule.
		v.putBack(tokens)
		return
	default:
		// Found a syntax error.
		var message = v.formatError("$SaveClause", token)
		panic(message)
	}

	// Attempt to parse a single "as" delimiter.
	_, token, ok = v.parseDelimiter("as")
	if !ok {
		if uti.IsDefined(tokens) {
			// This is not a single SaveClause rule.
			v.putBack(tokens)
			return
		} else {
			// Found a syntax error.
			var message = v.formatError("$SaveClause", token)
			panic(message)
		}
	}
	if uti.IsDefined(tokens) {
		tokens.AppendValue(token)
	}

	// Attempt to parse a single Cited rule.
	var cited ast.CitedLike
	cited, token, ok = v.parseCited()
	switch {
	case ok:
		// No additional put backs allowed at this point.
		tokens = nil
	case uti.IsDefined(tokens):
		// This is not a single Cited rule.
		v.putBack(tokens)
		return
	default:
		// Found a syntax error.
		var message = v.formatError("$SaveClause", token)
		panic(message)
	}

	// Found a single SaveClause rule.
	ok = true
	v.remove(tokens)
	saveClause = ast.SaveClauseClass().SaveClause(
		draft,
		cited,
	)
	return
}

func (v *parser_) parseSelectClause() (
	selectClause ast.SelectClauseLike,
	token TokenLike,
	ok bool,
) {
	var tokens = col.List[TokenLike]()

	// Attempt to parse a single "select" delimiter.
	_, token, ok = v.parseDelimiter("select")
	if !ok {
		if uti.IsDefined(tokens) {
			// This is not a single SelectClause rule.
			v.putBack(tokens)
			return
		} else {
			// Found a syntax error.
			var message = v.formatError("$SelectClause", token)
			panic(message)
		}
	}
	if uti.IsDefined(tokens) {
		tokens.AppendValue(token)
	}

	// Attempt to parse a single Target rule.
	var target ast.TargetLike
	target, token, ok = v.parseTarget()
	switch {
	case ok:
		// No additional put backs allowed at this point.
		tokens = nil
	case uti.IsDefined(tokens):
		// This is not a single Target rule.
		v.putBack(tokens)
		return
	default:
		// Found a syntax error.
		var message = v.formatError("$SelectClause", token)
		panic(message)
	}

	// Attempt to parse multiple MatchHandler rules.
	var matchHandlers = col.List[ast.MatchHandlerLike]()
matchHandlersLoop:
	for count := 0; count < mat.MaxInt; count++ {
		var matchHandler ast.MatchHandlerLike
		matchHandler, token, ok = v.parseMatchHandler()
		if !ok {
			switch {
			case count >= 1:
				break matchHandlersLoop
			case uti.IsDefined(tokens):
				// This is not multiple MatchHandler rules.
				v.putBack(tokens)
				return
			default:
				// Found a syntax error.
				var message = v.formatError("$SelectClause", token)
				message += "1 or more MatchHandler rules are required."
				panic(message)
			}
		}
		// No additional put backs allowed at this point.
		tokens = nil
		matchHandlers.AppendValue(matchHandler)
	}

	// Found a single SelectClause rule.
	ok = true
	v.remove(tokens)
	selectClause = ast.SelectClauseClass().SelectClause(
		target,
		matchHandlers,
	)
	return
}

func (v *parser_) parseSequence() (
	sequence ast.SequenceLike,
	token TokenLike,
	ok bool,
) {
	var tokens = col.List[TokenLike]()

	// Attempt to parse a single Expression rule.
	var expression ast.ExpressionLike
	expression, token, ok = v.parseExpression()
	switch {
	case ok:
		// No additional put backs allowed at this point.
		tokens = nil
	case uti.IsDefined(tokens):
		// This is not a single Expression rule.
		v.putBack(tokens)
		return
	default:
		// Found a syntax error.
		var message = v.formatError("$Sequence", token)
		panic(message)
	}

	// Found a single Sequence rule.
	ok = true
	v.remove(tokens)
	sequence = ast.SequenceClass().Sequence(expression)
	return
}

func (v *parser_) parseSlashEqual() (
	slashEqual ast.SlashEqualLike,
	token TokenLike,
	ok bool,
) {
	var tokens = col.List[TokenLike]()

	// Attempt to parse a single "/=" delimiter.
	_, token, ok = v.parseDelimiter("/=")
	if !ok {
		if uti.IsDefined(tokens) {
			// This is not a single SlashEqual rule.
			v.putBack(tokens)
			return
		} else {
			// Found a syntax error.
			var message = v.formatError("$SlashEqual", token)
			panic(message)
		}
	}
	if uti.IsDefined(tokens) {
		tokens.AppendValue(token)
	}

	// Found a single SlashEqual rule.
	ok = true
	v.remove(tokens)
	slashEqual = ast.SlashEqualClass().SlashEqual()
	return
}

func (v *parser_) parseStarEqual() (
	starEqual ast.StarEqualLike,
	token TokenLike,
	ok bool,
) {
	var tokens = col.List[TokenLike]()

	// Attempt to parse a single "*=" delimiter.
	_, token, ok = v.parseDelimiter("*=")
	if !ok {
		if uti.IsDefined(tokens) {
			// This is not a single StarEqual rule.
			v.putBack(tokens)
			return
		} else {
			// Found a syntax error.
			var message = v.formatError("$StarEqual", token)
			panic(message)
		}
	}
	if uti.IsDefined(tokens) {
		tokens.AppendValue(token)
	}

	// Found a single StarEqual rule.
	ok = true
	v.remove(tokens)
	starEqual = ast.StarEqualClass().StarEqual()
	return
}

func (v *parser_) parseStatement() (
	statement ast.StatementLike,
	token TokenLike,
	ok bool,
) {
	var tokens = col.List[TokenLike]()

	// Attempt to parse a single MainClause rule.
	var mainClause ast.MainClauseLike
	mainClause, token, ok = v.parseMainClause()
	switch {
	case ok:
		// No additional put backs allowed at this point.
		tokens = nil
	case uti.IsDefined(tokens):
		// This is not a single MainClause rule.
		v.putBack(tokens)
		return
	default:
		// Found a syntax error.
		var message = v.formatError("$Statement", token)
		panic(message)
	}

	// Attempt to parse an optional OnClause rule.
	var optionalOnClause ast.OnClauseLike
	optionalOnClause, _, ok = v.parseOnClause()
	if ok {
		// No additional put backs allowed at this point.
		tokens = nil
	}

	// Found a single Statement rule.
	ok = true
	v.remove(tokens)
	statement = ast.StatementClass().Statement(
		mainClause,
		optionalOnClause,
	)
	return
}

func (v *parser_) parseStatementLine() (
	statementLine ast.StatementLineLike,
	token TokenLike,
	ok bool,
) {
	var tokens = col.List[TokenLike]()

	// Attempt to parse a single Statement rule.
	var statement ast.StatementLike
	statement, token, ok = v.parseStatement()
	switch {
	case ok:
		// No additional put backs allowed at this point.
		tokens = nil
	case uti.IsDefined(tokens):
		// This is not a single Statement rule.
		v.putBack(tokens)
		return
	default:
		// Found a syntax error.
		var message = v.formatError("$StatementLine", token)
		panic(message)
	}

	// Attempt to parse an optional note token.
	var optionalNote string
	optionalNote, token, ok = v.parseToken(NoteToken)
	if ok && uti.IsDefined(tokens) {
		tokens.AppendValue(token)
	}

	// Found a single StatementLine rule.
	ok = true
	v.remove(tokens)
	statementLine = ast.StatementLineClass().StatementLine(
		statement,
		optionalNote,
	)
	return
}

func (v *parser_) parseString() (
	string_ ast.StringLike,
	token TokenLike,
	ok bool,
) {
	// Attempt to parse a single binary String.
	var binary string
	binary, token, ok = v.parseToken(BinaryToken)
	if ok {
		// Found a single binary String.
		string_ = ast.StringClass().String(binary)
		return
	}

	// Attempt to parse a single bytecode String.
	var bytecode string
	bytecode, token, ok = v.parseToken(BytecodeToken)
	if ok {
		// Found a single bytecode String.
		string_ = ast.StringClass().String(bytecode)
		return
	}

	// Attempt to parse a single name String.
	var name string
	name, token, ok = v.parseToken(NameToken)
	if ok {
		// Found a single name String.
		string_ = ast.StringClass().String(name)
		return
	}

	// Attempt to parse a single narrative String.
	var narrative string
	narrative, token, ok = v.parseToken(NarrativeToken)
	if ok {
		// Found a single narrative String.
		string_ = ast.StringClass().String(narrative)
		return
	}

	// Attempt to parse a single quote String.
	var quote string
	quote, token, ok = v.parseToken(QuoteToken)
	if ok {
		// Found a single quote String.
		string_ = ast.StringClass().String(quote)
		return
	}

	// Attempt to parse a single symbol String.
	var symbol string
	symbol, token, ok = v.parseToken(SymbolToken)
	if ok {
		// Found a single symbol String.
		string_ = ast.StringClass().String(symbol)
		return
	}

	// Attempt to parse a single tag String.
	var tag string
	tag, token, ok = v.parseToken(TagToken)
	if ok {
		// Found a single tag String.
		string_ = ast.StringClass().String(tag)
		return
	}

	// Attempt to parse a single version String.
	var version string
	version, token, ok = v.parseToken(VersionToken)
	if ok {
		// Found a single version String.
		string_ = ast.StringClass().String(version)
		return
	}

	// This is not a single String rule.
	return
}

func (v *parser_) parseSubcomponent() (
	subcomponent ast.SubcomponentLike,
	token TokenLike,
	ok bool,
) {
	var tokens = col.List[TokenLike]()

	// Attempt to parse a single identifier token.
	var identifier string
	identifier, token, ok = v.parseToken(IdentifierToken)
	if !ok {
		if uti.IsDefined(tokens) {
			// This is not a single identifier token.
			v.putBack(tokens)
			return
		} else {
			// Found a syntax error.
			var message = v.formatError("$Subcomponent", token)
			panic(message)
		}
	}
	if uti.IsDefined(tokens) {
		tokens.AppendValue(token)
	}

	// Attempt to parse a single "[" delimiter.
	_, token, ok = v.parseDelimiter("[")
	if !ok {
		if uti.IsDefined(tokens) {
			// This is not a single Subcomponent rule.
			v.putBack(tokens)
			return
		} else {
			// Found a syntax error.
			var message = v.formatError("$Subcomponent", token)
			panic(message)
		}
	}
	if uti.IsDefined(tokens) {
		tokens.AppendValue(token)
	}

	// Attempt to parse a single Indices rule.
	var indices ast.IndicesLike
	indices, token, ok = v.parseIndices()
	switch {
	case ok:
		// No additional put backs allowed at this point.
		tokens = nil
	case uti.IsDefined(tokens):
		// This is not a single Indices rule.
		v.putBack(tokens)
		return
	default:
		// Found a syntax error.
		var message = v.formatError("$Subcomponent", token)
		panic(message)
	}

	// Attempt to parse a single "]" delimiter.
	_, token, ok = v.parseDelimiter("]")
	if !ok {
		if uti.IsDefined(tokens) {
			// This is not a single Subcomponent rule.
			v.putBack(tokens)
			return
		} else {
			// Found a syntax error.
			var message = v.formatError("$Subcomponent", token)
			panic(message)
		}
	}
	if uti.IsDefined(tokens) {
		tokens.AppendValue(token)
	}

	// Found a single Subcomponent rule.
	ok = true
	v.remove(tokens)
	subcomponent = ast.SubcomponentClass().Subcomponent(
		identifier,
		indices,
	)
	return
}

func (v *parser_) parseSubject() (
	subject ast.SubjectLike,
	token TokenLike,
	ok bool,
) {
	// Attempt to parse a single Component Subject.
	var component ast.ComponentLike
	component, token, ok = v.parseComponent()
	if ok {
		// Found a single Component Subject.
		subject = ast.SubjectClass().Subject(component)
		return
	}

	// Attempt to parse a single Subcomponent Subject.
	var subcomponent ast.SubcomponentLike
	subcomponent, token, ok = v.parseSubcomponent()
	if ok {
		// Found a single Subcomponent Subject.
		subject = ast.SubjectClass().Subject(subcomponent)
		return
	}

	// Attempt to parse a single Precedence Subject.
	var precedence ast.PrecedenceLike
	precedence, token, ok = v.parsePrecedence()
	if ok {
		// Found a single Precedence Subject.
		subject = ast.SubjectClass().Subject(precedence)
		return
	}

	// Attempt to parse a single Referent Subject.
	var referent ast.ReferentLike
	referent, token, ok = v.parseReferent()
	if ok {
		// Found a single Referent Subject.
		subject = ast.SubjectClass().Subject(referent)
		return
	}

	// Attempt to parse a single Complement Subject.
	var complement ast.ComplementLike
	complement, token, ok = v.parseComplement()
	if ok {
		// Found a single Complement Subject.
		subject = ast.SubjectClass().Subject(complement)
		return
	}

	// Attempt to parse a single Inversion Subject.
	var inversion ast.InversionLike
	inversion, token, ok = v.parseInversion()
	if ok {
		// Found a single Inversion Subject.
		subject = ast.SubjectClass().Subject(inversion)
		return
	}

	// Attempt to parse a single Magnitude Subject.
	var magnitude ast.MagnitudeLike
	magnitude, token, ok = v.parseMagnitude()
	if ok {
		// Found a single Magnitude Subject.
		subject = ast.SubjectClass().Subject(magnitude)
		return
	}

	// Attempt to parse a single Function Subject.
	var function ast.FunctionLike
	function, token, ok = v.parseFunction()
	if ok {
		// Found a single Function Subject.
		subject = ast.SubjectClass().Subject(function)
		return
	}

	// Attempt to parse a single Method Subject.
	var method ast.MethodLike
	method, token, ok = v.parseMethod()
	if ok {
		// Found a single Method Subject.
		subject = ast.SubjectClass().Subject(method)
		return
	}

	// Attempt to parse a single Variable Subject.
	var variable ast.VariableLike
	variable, token, ok = v.parseVariable()
	if ok {
		// Found a single Variable Subject.
		subject = ast.SubjectClass().Subject(variable)
		return
	}

	// This is not a single Subject rule.
	return
}

func (v *parser_) parseTarget() (
	target ast.TargetLike,
	token TokenLike,
	ok bool,
) {
	// Attempt to parse a single Function Target.
	var function ast.FunctionLike
	function, token, ok = v.parseFunction()
	if ok {
		// Found a single Function Target.
		target = ast.TargetClass().Target(function)
		return
	}

	// Attempt to parse a single Method Target.
	var method ast.MethodLike
	method, token, ok = v.parseMethod()
	if ok {
		// Found a single Method Target.
		target = ast.TargetClass().Target(method)
		return
	}

	// Attempt to parse a single Subcomponent Target.
	var subcomponent ast.SubcomponentLike
	subcomponent, token, ok = v.parseSubcomponent()
	if ok {
		// Found a single Subcomponent Target.
		target = ast.TargetClass().Target(subcomponent)
		return
	}

	// Attempt to parse a single Variable Target.
	var variable ast.VariableLike
	variable, token, ok = v.parseVariable()
	if ok {
		// Found a single Variable Target.
		target = ast.TargetClass().Target(variable)
		return
	}

	// This is not a single Target rule.
	return
}

func (v *parser_) parseTemplate() (
	template ast.TemplateLike,
	token TokenLike,
	ok bool,
) {
	var tokens = col.List[TokenLike]()

	// Attempt to parse a single Expression rule.
	var expression ast.ExpressionLike
	expression, token, ok = v.parseExpression()
	switch {
	case ok:
		// No additional put backs allowed at this point.
		tokens = nil
	case uti.IsDefined(tokens):
		// This is not a single Expression rule.
		v.putBack(tokens)
		return
	default:
		// Found a syntax error.
		var message = v.formatError("$Template", token)
		panic(message)
	}

	// Found a single Template rule.
	ok = true
	v.remove(tokens)
	template = ast.TemplateClass().Template(expression)
	return
}

func (v *parser_) parseThrowClause() (
	throwClause ast.ThrowClauseLike,
	token TokenLike,
	ok bool,
) {
	var tokens = col.List[TokenLike]()

	// Attempt to parse a single "throw" delimiter.
	_, token, ok = v.parseDelimiter("throw")
	if !ok {
		if uti.IsDefined(tokens) {
			// This is not a single ThrowClause rule.
			v.putBack(tokens)
			return
		} else {
			// Found a syntax error.
			var message = v.formatError("$ThrowClause", token)
			panic(message)
		}
	}
	if uti.IsDefined(tokens) {
		tokens.AppendValue(token)
	}

	// Attempt to parse a single Exception rule.
	var exception ast.ExceptionLike
	exception, token, ok = v.parseException()
	switch {
	case ok:
		// No additional put backs allowed at this point.
		tokens = nil
	case uti.IsDefined(tokens):
		// This is not a single Exception rule.
		v.putBack(tokens)
		return
	default:
		// Found a syntax error.
		var message = v.formatError("$ThrowClause", token)
		panic(message)
	}

	// Found a single ThrowClause rule.
	ok = true
	v.remove(tokens)
	throwClause = ast.ThrowClauseClass().ThrowClause(exception)
	return
}

func (v *parser_) parseVariable() (
	variable ast.VariableLike,
	token TokenLike,
	ok bool,
) {
	var tokens = col.List[TokenLike]()

	// Attempt to parse a single identifier token.
	var identifier string
	identifier, token, ok = v.parseToken(IdentifierToken)
	if !ok {
		if uti.IsDefined(tokens) {
			// This is not a single identifier token.
			v.putBack(tokens)
			return
		} else {
			// Found a syntax error.
			var message = v.formatError("$Variable", token)
			panic(message)
		}
	}
	if uti.IsDefined(tokens) {
		tokens.AppendValue(token)
	}

	// Found a single Variable rule.
	ok = true
	v.remove(tokens)
	variable = ast.VariableClass().Variable(identifier)
	return
}

func (v *parser_) parseWhileClause() (
	whileClause ast.WhileClauseLike,
	token TokenLike,
	ok bool,
) {
	var tokens = col.List[TokenLike]()

	// Attempt to parse a single "while" delimiter.
	_, token, ok = v.parseDelimiter("while")
	if !ok {
		if uti.IsDefined(tokens) {
			// This is not a single WhileClause rule.
			v.putBack(tokens)
			return
		} else {
			// Found a syntax error.
			var message = v.formatError("$WhileClause", token)
			panic(message)
		}
	}
	if uti.IsDefined(tokens) {
		tokens.AppendValue(token)
	}

	// Attempt to parse a single Condition rule.
	var condition ast.ConditionLike
	condition, token, ok = v.parseCondition()
	switch {
	case ok:
		// No additional put backs allowed at this point.
		tokens = nil
	case uti.IsDefined(tokens):
		// This is not a single Condition rule.
		v.putBack(tokens)
		return
	default:
		// Found a syntax error.
		var message = v.formatError("$WhileClause", token)
		panic(message)
	}

	// Attempt to parse a single "do" delimiter.
	_, token, ok = v.parseDelimiter("do")
	if !ok {
		if uti.IsDefined(tokens) {
			// This is not a single WhileClause rule.
			v.putBack(tokens)
			return
		} else {
			// Found a syntax error.
			var message = v.formatError("$WhileClause", token)
			panic(message)
		}
	}
	if uti.IsDefined(tokens) {
		tokens.AppendValue(token)
	}

	// Attempt to parse a single Procedure rule.
	var procedure ast.ProcedureLike
	procedure, token, ok = v.parseProcedure()
	switch {
	case ok:
		// No additional put backs allowed at this point.
		tokens = nil
	case uti.IsDefined(tokens):
		// This is not a single Procedure rule.
		v.putBack(tokens)
		return
	default:
		// Found a syntax error.
		var message = v.formatError("$WhileClause", token)
		panic(message)
	}

	// Found a single WhileClause rule.
	ok = true
	v.remove(tokens)
	whileClause = ast.WhileClauseClass().WhileClause(
		condition,
		procedure,
	)
	return
}

func (v *parser_) parseWithClause() (
	withClause ast.WithClauseLike,
	token TokenLike,
	ok bool,
) {
	var tokens = col.List[TokenLike]()

	// Attempt to parse a single "with" delimiter.
	_, token, ok = v.parseDelimiter("with")
	if !ok {
		if uti.IsDefined(tokens) {
			// This is not a single WithClause rule.
			v.putBack(tokens)
			return
		} else {
			// Found a syntax error.
			var message = v.formatError("$WithClause", token)
			panic(message)
		}
	}
	if uti.IsDefined(tokens) {
		tokens.AppendValue(token)
	}

	// Attempt to parse a single "each" delimiter.
	_, token, ok = v.parseDelimiter("each")
	if !ok {
		if uti.IsDefined(tokens) {
			// This is not a single WithClause rule.
			v.putBack(tokens)
			return
		} else {
			// Found a syntax error.
			var message = v.formatError("$WithClause", token)
			panic(message)
		}
	}
	if uti.IsDefined(tokens) {
		tokens.AppendValue(token)
	}

	// Attempt to parse a single Item rule.
	var item ast.ItemLike
	item, token, ok = v.parseItem()
	switch {
	case ok:
		// No additional put backs allowed at this point.
		tokens = nil
	case uti.IsDefined(tokens):
		// This is not a single Item rule.
		v.putBack(tokens)
		return
	default:
		// Found a syntax error.
		var message = v.formatError("$WithClause", token)
		panic(message)
	}

	// Attempt to parse a single "in" delimiter.
	_, token, ok = v.parseDelimiter("in")
	if !ok {
		if uti.IsDefined(tokens) {
			// This is not a single WithClause rule.
			v.putBack(tokens)
			return
		} else {
			// Found a syntax error.
			var message = v.formatError("$WithClause", token)
			panic(message)
		}
	}
	if uti.IsDefined(tokens) {
		tokens.AppendValue(token)
	}

	// Attempt to parse a single Sequence rule.
	var sequence ast.SequenceLike
	sequence, token, ok = v.parseSequence()
	switch {
	case ok:
		// No additional put backs allowed at this point.
		tokens = nil
	case uti.IsDefined(tokens):
		// This is not a single Sequence rule.
		v.putBack(tokens)
		return
	default:
		// Found a syntax error.
		var message = v.formatError("$WithClause", token)
		panic(message)
	}

	// Attempt to parse a single "do" delimiter.
	_, token, ok = v.parseDelimiter("do")
	if !ok {
		if uti.IsDefined(tokens) {
			// This is not a single WithClause rule.
			v.putBack(tokens)
			return
		} else {
			// Found a syntax error.
			var message = v.formatError("$WithClause", token)
			panic(message)
		}
	}
	if uti.IsDefined(tokens) {
		tokens.AppendValue(token)
	}

	// Attempt to parse a single Procedure rule.
	var procedure ast.ProcedureLike
	procedure, token, ok = v.parseProcedure()
	switch {
	case ok:
		// No additional put backs allowed at this point.
		tokens = nil
	case uti.IsDefined(tokens):
		// This is not a single Procedure rule.
		v.putBack(tokens)
		return
	default:
		// Found a syntax error.
		var message = v.formatError("$WithClause", token)
		panic(message)
	}

	// Found a single WithClause rule.
	ok = true
	v.remove(tokens)
	withClause = ast.WithClauseClass().WithClause(
		item,
		sequence,
		procedure,
	)
	return
}

func (v *parser_) parseXor() (
	xor ast.XorLike,
	token TokenLike,
	ok bool,
) {
	var tokens = col.List[TokenLike]()

	// Attempt to parse a single "xor" delimiter.
	_, token, ok = v.parseDelimiter("xor")
	if !ok {
		if uti.IsDefined(tokens) {
			// This is not a single Xor rule.
			v.putBack(tokens)
			return
		} else {
			// Found a syntax error.
			var message = v.formatError("$Xor", token)
			panic(message)
		}
	}
	if uti.IsDefined(tokens) {
		tokens.AppendValue(token)
	}

	// Found a single Xor rule.
	ok = true
	v.remove(tokens)
	xor = ast.XorClass().Xor()
	return
}

func (v *parser_) parseDelimiter(
	expectedValue string,
) (
	value string,
	token TokenLike,
	ok bool,
) {
	// Attempt to parse a single delimiter.
	value, token, ok = v.parseToken(DelimiterToken)
	if ok {
		if value == expectedValue {
			// Found the desired delimiter.
			return
		}
		v.next_.AddValue(token)
		ok = false
	}

	// This is not the desired delimiter.
	return
}

func (v *parser_) parseToken(
	tokenType TokenType,
) (
	value string,
	token TokenLike,
	ok bool,
) {
	// Attempt to parse a specific token type.
	var tokens = col.List[TokenLike]()
	token = v.getNextToken()
	for token != nil {
		tokens.AppendValue(token)
		switch token.GetType() {
		case tokenType:
			// Found the desired token type.
			value = token.GetValue()
			ok = true
			return
		case SpaceToken, NewlineToken:
			// Ignore any unrequested whitespace.
			token = v.getNextToken()
		default:
			// This is not the desired token type.
			v.putBack(tokens)
			return
		}
	}

	// We are at the end-of-file marker.
	return
}

func (v *parser_) formatError(
	ruleName string,
	token TokenLike,
) string {
	// Format the error message.
	var message = fmt.Sprintf(
		"An unexpected token was received by the parser: %v\n",
		ScannerClass().FormatToken(token),
	)
	var line = token.GetLine()
	var lines = sts.Split(v.source_, "\n")

	// Append the source lines with the error in it.
	message += "\033[36m"
	for index := line - 3; index < line; index++ {
		if index > 1 {
			message += fmt.Sprintf("%04d: ", index) + string(lines[index-1]) + "\n"
		}
	}
	message += fmt.Sprintf("%04d: ", line) + string(lines[line-1]) + "\n"

	// Append an arrow pointing to the error.
	message += " \033[32m>>>─"
	var count uint
	for count < token.GetPosition() {
		message += "─"
		count++
	}
	message += "⌃\033[36m\n"

	// Append the following source line for context.
	if line < uint(len(lines)) {
		message += fmt.Sprintf("%04d: ", line+1) + string(lines[line]) + "\n"
	}
	message += "\033[0m\n"
	if uti.IsDefined(ruleName) {
		message += "Was expecting:\n"
		message += fmt.Sprintf(
			"  \033[32m%v: \033[33m%v\033[0m\n\n",
			ruleName,
			v.getDefinition(ruleName),
		)
	}
	return message
}

func (v *parser_) getDefinition(
	ruleName string,
) string {
	var syntax = parserClass().syntax_
	var definition = syntax.GetValue(ruleName)
	return definition
}

func (v *parser_) getNextToken() TokenLike {
	// Check for any read, but unprocessed tokens.
	if !v.next_.IsEmpty() {
		return v.next_.RemoveLast()
	}

	// Read a new token from the token stream.
	var token, ok = v.tokens_.RemoveFirst() // This will wait for a token.
	if !ok {
		// The token channel has been closed.
		return nil
	}

	// Check for an error token.
	if token.GetType() == ErrorToken {
		var message = v.formatError("", token)
		panic(message)
	}

	return token
}

func (v *parser_) putBack(
	tokens col.Sequential[TokenLike],
) {
	var iterator = tokens.GetIterator()
	for iterator.ToEnd(); iterator.HasPrevious(); {
		var token = iterator.GetPrevious()
		v.next_.AddValue(token)
	}
}

func (v *parser_) remove(
	tokens col.Sequential[TokenLike],
) {
	// NOTE: This method does nothing but must exist to satisfy the lint
	// check on the generated parser code.
}

// Instance Structure

type parser_ struct {
	// Declare the instance attributes.
	source_ string                   // The original source code.
	tokens_ col.QueueLike[TokenLike] // A queue of unread tokens from the scanner.
	next_   col.StackLike[TokenLike] // A stack of read, but unprocessed tokens.
}

// Class Structure

type parserClass_ struct {
	// Declare the class constants.
	syntax_ col.CatalogLike[string, string]
}

// Class Reference

func parserClass() *parserClass_ {
	return parserClassReference_
}

var parserClassReference_ = &parserClass_{
	// Initialize the class constants.
	syntax_: col.CatalogFromMap[string, string](
		map[string]string{
			"$Document":  `Notice? Component`,
			"$Notice":    `comment newline`,
			"$Component": `Entity Parameters?`,
			"$Entity": `
    Element
    String
    Collection
    Procedure`,
			"$Parameters": `
    MultilineParameters
    InlineParameters`,
			"$MultilineParameters":   `"(" newline AnnotatedAssociation+ ")"`,
			"$AnnotatedAssociation":  `Association note?`,
			"$InlineParameters":      `"(" Association AdditionalAssociation* ")"`,
			"$AdditionalAssociation": `"," Association`,
			"$Association":           `Primitive ":" Component`,
			"$Primitive": `
    Element
    String`,
			"$Element": `
    angle
    boolean
    citation
    duration
    moment
    number
    pattern
    percentage
    probability
    resource`,
			"$String": `
    binary
    bytecode
    name
    narrative
    quote
    symbol
    tag
    version`,
			"$Collection": `
    MultilineAttributes
    MultilineValues
    InclusiveRange  ! Must be after the multiline rules.
    ExclusiveRange
    InlineAttributes
    InlineValues  ! Must be after the ranges and inline attributes.
    NoAttributes
    NoValues`,
			"$InclusiveRange": `"[" Primitive ".." Primitive Bracket`,
			"$ExclusiveRange": `"(" Primitive ".." Primitive Bracket`,
			"$Bracket": `
    Inclusive
    Exclusive`,
			"$Inclusive":           `"]"`,
			"$Exclusive":           `")"`,
			"$NoAttributes":        `"[" ":" "]"`,
			"$InlineAttributes":    `"[" Association AdditionalAssociation* "]"`,
			"$MultilineAttributes": `"[" newline AnnotatedAssociation+ "]"`,
			"$NoValues":            `"[" "]"`,
			"$InlineValues":        `"[" Component AdditionalValue* "]"`,
			"$AdditionalValue":     `"," Component`,
			"$MultilineValues":     `"[" newline AnnotatedValue+ "]"`,
			"$AnnotatedValue":      `Component note?`,
			"$Procedure": `
    MultilineStatements
    InlineStatements
    NoStatements`,
			"$MultilineStatements": `"{" newline AnnotatedStatement* "}"`,
			"$AnnotatedStatement": `
    CommentLine
    StatementLine`,
			"$CommentLine":         `comment`,
			"$StatementLine":       `Statement note?`,
			"$InlineStatements":    `"{" Statement AdditionalStatement* "}"`,
			"$AdditionalStatement": `";" Statement`,
			"$NoStatements":        `"{" "}"`,
			"$Statement":           `MainClause OnClause?`,
			"$MainClause": `
    Flow
    Induction
    Messaging
    Repository`,
			"$OnClause":     `"on" Failure MatchHandler+`,
			"$MatchHandler": `"matching" Template "do" Procedure`,
			"$Failure":      `symbol`,
			"$Template":     `Expression`,
			"$Flow": `
    IfClause
    SelectClause
    WhileClause
    WithClause
    ContinueClause
    BreakClause
    ReturnClause
    ThrowClause`,
			"$Induction": `
    DoClause
    LetClause`,
			"$Messaging": `
    PostClause
    RetrieveClause
    AcceptClause
    RejectClause
    PublishClause`,
			"$Repository": `
    CheckoutClause
    SaveClause
    DiscardClause
    NotarizeClause`,
			"$IfClause":     `"if" Condition "do" Procedure`,
			"$Condition":    `Expression`,
			"$SelectClause": `"select" Target MatchHandler+`,
			"$Target": `
    Function
    Method
    Subcomponent
    Variable  ! This must be last since others also begin with an identifier.`,
			"$Function":           `identifier "(" Arguments? ")"`,
			"$Arguments":          `Argument AdditionalArgument*`,
			"$AdditionalArgument": `"," Argument`,
			"$Argument":           `identifier`,
			"$Method":             `identifier Blocking identifier "(" Arguments? ")"`,
			"$Blocking": `
    dot
    arrow`,
			"$Subcomponent":    `identifier "[" Indices "]"`,
			"$Indices":         `Index AdditionalIndex*`,
			"$AdditionalIndex": `"," Index`,
			"$Index":           `Expression`,
			"$Variable":        `identifier`,
			"$WhileClause":     `"while" Condition "do" Procedure`,
			"$WithClause":      `"with" "each" Item "in" Sequence "do" Procedure`,
			"$Item":            `symbol`,
			"$Sequence":        `Expression`,
			"$ContinueClause":  `"continue" "loop"`,
			"$BreakClause":     `"break" "loop"`,
			"$ReturnClause":    `"return" Result`,
			"$Result":          `Expression`,
			"$ThrowClause":     `"throw" Exception`,
			"$Exception":       `Expression`,
			"$DoClause":        `"do" Invocation`,
			"$Invocation": `
    Function
    Method`,
			"$LetClause": `"let" Recipient Assignment Expression`,
			"$Assignment": `
    ColonEqual
    DefaultEqual
    PlusEqual
    DashEqual
    StarEqual
    SlashEqual`,
			"$ColonEqual":   `":="`,
			"$DefaultEqual": `"?="`,
			"$PlusEqual":    `"+="`,
			"$DashEqual":    `"-="`,
			"$StarEqual":    `"*="`,
			"$SlashEqual":   `"/="`,
			"$Recipient": `
    Item
    Subcomponent`,
			"$PostClause":     `"post" Message "to" Bag`,
			"$Message":        `Expression`,
			"$Bag":            `Expression`,
			"$RetrieveClause": `"retrieve" Recipient "from" Bag`,
			"$AcceptClause":   `"accept" Message`,
			"$RejectClause":   `"reject" Message`,
			"$PublishClause":  `"publish" Event`,
			"$Event":          `Expression`,
			"$CheckoutClause": `"checkout" Recipient AtLevel? "from" Cited`,
			"$AtLevel":        `"at" "level" Expression`,
			"$Cited":          `Expression`,
			"$SaveClause":     `"save" Draft "as" Cited`,
			"$Draft":          `Expression`,
			"$DiscardClause":  `"discard" Draft`,
			"$NotarizeClause": `"notarize" Draft "as" Cited`,
			"$Expression":     `Subject Predicate*`,
			"$Subject": `
    Component
    Subcomponent
    Precedence
    Referent
    Complement
    Inversion
    Magnitude
    Function
    Method
    Variable  ! This must be last since others also begin with an identifier.`,
			"$Predicate": `Action Expression`,
			"$Action": `
    Operator
    Operation`,
			"$Operator": `
    plus
    dash
    star
    slash
    double
    caret
    ampersand
    less
    equal
    more`,
			"$Operation": `
    Is
    Matches
    And
    San
    Ior
    Xor`,
			"$Is":         `"is"`,
			"$Matches":    `"matches"`,
			"$And":        `"and"`,
			"$San":        `"san"`,
			"$Ior":        `"ior"`,
			"$Xor":        `"xor"`,
			"$Precedence": `"(" Expression ")"`,
			"$Referent":   `"@" Indirect`,
			"$Indirect": `
    Component
    Subcomponent
    Referent
    Function
    Method
    Variable  ! This must be last since others also begin with an identifier.`,
			"$Complement": `"not" Logical`,
			"$Logical": `
    Component
    Subcomponent
    Precedence
    Referent
    Complement
    Function
    Method
    Variable  ! This must be last since others also begin with an identifier.`,
			"$Inversion": `Inverse Numerical`,
			"$Inverse": `
    dash
    slash
    star`,
			"$Numerical": `
    Component
    Subcomponent
    Precedence
    Referent
    Inversion
    Magnitude
    Function
    Method
    Variable  ! This must be last since others also begin with an identifier.`,
			"$Magnitude": `"|" Numerical "|"`,
		},
	),
}
