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
│ Updates to any section other than the Methodical Methods may be overwritten. │
└──────────────────────────────────────────────────────────────────────────────┘
*/

package grammar

import (
	ast "github.com/bali-nebula/go-bali-documents/v3/ast"
	uti "github.com/craterdog/go-missing-utilities/v7"
	sts "strings"
)

// CLASS INTERFACE

// Access Function

func FormatterClass() FormatterClassLike {
	return formatterClass()
}

// Constructor Methods

func (c *formatterClass_) Formatter() FormatterLike {
	var instance = &formatter_{
		// Initialize the instance attributes.

		// Initialize the inherited aspects.
		Methodical: ProcessorClass().Processor(),
	}
	instance.visitor_ = VisitorClass().Visitor(instance)
	return instance
}

// INSTANCE INTERFACE

// Principal Methods

func (v *formatter_) GetClass() FormatterClassLike {
	return formatterClass()
}

func (v *formatter_) FormatDocument(document ast.DocumentLike) string {
	v.visitor_.VisitDocument(document)
	return v.getResult()
}

// Methodical Methods

func (v *formatter_) ProcessAngle(
	angle string,
) {
	v.appendString(angle)
}

func (v *formatter_) ProcessAsynchronous(
	asynchronous string,
) {
	v.appendString(asynchronous)
}

func (v *formatter_) ProcessBinary(
	binary string,
) {
	binary = v.adjustIndentation(binary)
	v.appendString(binary)
}

func (v *formatter_) ProcessBoolean(
	boolean string,
) {
	v.appendString(boolean)
}

func (v *formatter_) ProcessBytecode(
	bytecode string,
) {
	bytecode = v.adjustIndentation(bytecode)
	v.appendString(bytecode)
}

func (v *formatter_) ProcessCaret(
	caret string,
) {
	v.appendString(caret)
}

func (v *formatter_) ProcessComment(
	comment string,
) {
	comment = v.adjustIndentation(comment)
	v.appendString(comment)
}

func (v *formatter_) ProcessDash(
	dash string,
) {
	v.appendString(dash)
}

func (v *formatter_) ProcessDelimiter(
	delimiter string,
) {
	v.appendString(delimiter)
}

func (v *formatter_) ProcessDuration(
	duration string,
) {
	v.appendString(duration)
}

func (v *formatter_) ProcessEqual(
	equal string,
) {
	v.appendString(equal)
}

func (v *formatter_) ProcessGlyph(
	glyph string,
) {
	v.appendString(glyph)
}

func (v *formatter_) ProcessIdentifier(
	identifier string,
) {
	v.appendString(identifier)
}

func (v *formatter_) ProcessLess(
	less string,
) {
	v.appendString(less)
}

func (v *formatter_) ProcessModulo(
	modulo string,
) {
	v.appendString(modulo)
}

func (v *formatter_) ProcessMoment(
	moment string,
) {
	v.appendString(moment)
}

func (v *formatter_) ProcessMore(
	more string,
) {
	v.appendString(more)
}

func (v *formatter_) ProcessName(
	name string,
) {
	v.appendString(name)
}

func (v *formatter_) ProcessNarrative(
	narrative string,
) {
	narrative = v.adjustIndentation(narrative)
	v.appendString(narrative)
}

func (v *formatter_) ProcessNewline(
	newline string,
) {
	v.appendNewline()
}

func (v *formatter_) ProcessNote(
	note string,
) {
	v.appendString(note)
}

func (v *formatter_) ProcessNumber(
	number string,
) {
	v.appendString(number)
}

func (v *formatter_) ProcessPattern(
	pattern string,
) {
	v.appendString(pattern)
}

func (v *formatter_) ProcessPercentage(
	percentage string,
) {
	v.appendString(percentage)
}

func (v *formatter_) ProcessPlus(
	plus string,
) {
	v.appendString(plus)
}

func (v *formatter_) ProcessProbability(
	probability string,
) {
	v.appendString(probability)
}

func (v *formatter_) ProcessQuote(
	quote string,
) {
	v.appendString(quote)
}

func (v *formatter_) ProcessResource(
	resource string,
) {
	v.appendString(resource)
}

func (v *formatter_) ProcessSlash(
	slash string,
) {
	v.appendString(slash)
}

func (v *formatter_) ProcessSpace(
	space string,
) {
	v.appendString(space)
}

func (v *formatter_) ProcessStar(
	star string,
) {
	v.appendString(star)
}

func (v *formatter_) ProcessSymbol(
	symbol string,
) {
	v.appendString(symbol)
}

func (v *formatter_) ProcessSynchronous(
	synchronous string,
) {
	v.appendString(synchronous)
}

func (v *formatter_) ProcessTag(
	tag string,
) {
	v.appendString(tag)
}

func (v *formatter_) ProcessVersion(
	version string,
) {
	v.appendString(version)
}

func (v *formatter_) ProcessAcceptClauseSlot(
	acceptClause ast.AcceptClauseLike,
	slot_ uint,
) {
	v.appendString(" ")
}

func (v *formatter_) PreprocessAnnotation(
	annotation ast.AnnotationLike,
	index_ uint,
	count_ uint,
) {
	v.appendString("\n")
	v.appendNewline()
}

func (v *formatter_) PreprocessArgument(
	argument ast.ArgumentLike,
	index_ uint,
	count_ uint,
) {
	if index_ > 1 {
		v.appendString(" ")
	}
}

func (v *formatter_) ProcessArithmeticSlot(
	arithmetic ast.ArithmeticLike,
	slot_ uint,
) {
	v.appendString(" ")
}

func (v *formatter_) ProcessAssignmentSlot(
	assignment ast.AssignmentLike,
	slot_ uint,
) {
	v.appendString(" ")
}

func (v *formatter_) PreprocessAssociation(
	association ast.AssociationLike,
	index_ uint,
	count_ uint,
) {
	if count_ > 1 {
		v.appendNewline()
	}
}

func (v *formatter_) ProcessAssociationSlot(
	association ast.AssociationLike,
	slot_ uint,
) {
	if slot_ == 2 {
		v.appendString(" ")
	}
}

func (v *formatter_) PreprocessAtLevel(
	atLevel ast.AtLevelLike,
	index_ uint,
	count_ uint,
) {
	v.appendString(" ")
}

func (v *formatter_) ProcessAtLevelSlot(
	atLevel ast.AtLevelLike,
	slot_ uint,
) {
	v.appendString(" ")
}

func (v *formatter_) ProcessAttributesSlot(
	attributes ast.AttributesLike,
	slot_ uint,
) {
	switch slot_ {
	case 1:
		v.depth_++
	case 2:
		v.depth_--
		var associations = attributes.GetAssociations()
		if associations.GetSize() > 1 {
			v.appendNewline()
		}
	}
}

func (v *formatter_) ProcessBreakClauseSlot(
	breakClause ast.BreakClauseLike,
	slot_ uint,
) {
	v.appendString(" ")
}

func (v *formatter_) ProcessCheckoutClauseSlot(
	checkoutClause ast.CheckoutClauseLike,
	slot_ uint,
) {
	switch slot_ {
	case 1, 3, 4:
		v.appendString(" ")
	}
}

func (v *formatter_) ProcessComplementSlot(
	complement ast.ComplementLike,
	slot_ uint,
) {
	v.appendString(" ")
}

func (v *formatter_) ProcessContinueClauseSlot(
	continueClause ast.ContinueClauseLike,
	slot_ uint,
) {
	v.appendString(" ")
}

func (v *formatter_) ProcessDiscardClauseSlot(
	discardClause ast.DiscardClauseLike,
	slot_ uint,
) {
	v.appendString(" ")
}

func (v *formatter_) ProcessDoClauseSlot(
	doClause ast.DoClauseLike,
	slot_ uint,
) {
	v.appendString(" ")
}

func (v *formatter_) PreprocessDocument(
	document ast.DocumentLike,
	index_ uint,
	count_ uint,
) {
	if count_ > 1 {
		v.appendNewline()
	}
}

func (v *formatter_) ProcessDocumentSlot(
	document ast.DocumentLike,
	slot_ uint,
) {
	var note = document.GetOptionalNote()
	if uti.IsDefined(note) && slot_ == 2 {
		v.appendString("  ")
	}
}

func (v *formatter_) PostprocessDocument(
	document ast.DocumentLike,
	index_ uint,
	count_ uint,
) {
	if v.depth_ == 0 {
		v.appendNewline()
	}
}

func (v *formatter_) ProcessEmptySlot(
	empty ast.EmptyLike,
	slot_ uint,
) {
	var colon = empty.GetOptionalDelimiter()
	if uti.IsUndefined(colon) && slot_ == 1 {
		v.appendString(" ")
	}
}

func (v *formatter_) ProcessIfClauseSlot(
	ifClause ast.IfClauseLike,
	slot_ uint,
) {
	v.appendString(" ")
}

func (v *formatter_) PreprocessIndex(
	index ast.IndexLike,
	index_ uint,
	count_ uint,
) {
	if index_ > 1 {
		v.appendString(" ")
	}
}

func (v *formatter_) ProcessItemsSlot(
	items ast.ItemsLike,
	slot_ uint,
) {
	switch slot_ {
	case 1:
		v.depth_++
	case 2:
		v.depth_--
		var documents = items.GetDocuments()
		if documents.GetSize() > 1 {
			v.appendNewline()
		}
	}
}

func (v *formatter_) ProcessLetClauseSlot(
	letClause ast.LetClauseLike,
	slot_ uint,
) {
	v.appendString(" ")
}

func (v *formatter_) PreprocessMatchingClause(
	matchingClause ast.MatchingClauseLike,
	index_ uint,
	count_ uint,
) {
	v.appendString(" ")
}

func (v *formatter_) ProcessMatchingClauseSlot(
	matchingClause ast.MatchingClauseLike,
	slot_ uint,
) {
	v.appendString(" ")
}

func (v *formatter_) ProcessNotarizeClauseSlot(
	notarizeClause ast.NotarizeClauseLike,
	slot_ uint,
) {
	v.appendString(" ")
}

func (v *formatter_) PreprocessOnClause(
	onClause ast.OnClauseLike,
	index_ uint,
	count_ uint,
) {
	v.appendString(" ")
}

func (v *formatter_) ProcessOnClauseSlot(
	onClause ast.OnClauseLike,
	slot_ uint,
) {
	switch slot_ {
	case 1:
		v.appendString(" ")
	}
}

func (v *formatter_) ProcessParametersSlot(
	parameters ast.ParametersLike,
	slot_ uint,
) {
	switch slot_ {
	case 1:
		v.depth_++
	case 2:
		v.depth_--
		var associations = parameters.GetAssociations()
		if associations.GetSize() > 1 {
			v.appendNewline()
		}
	}
}

func (v *formatter_) ProcessPostClauseSlot(
	postClause ast.PostClauseLike,
	slot_ uint,
) {
	v.appendString(" ")
}

func (v *formatter_) PreprocessPredicate(
	predicate ast.PredicateLike,
	index_ uint,
	count_ uint,
) {
	v.appendString(" ")
}

func (v *formatter_) ProcessPredicateSlot(
	predicate ast.PredicateLike,
	slot_ uint,
) {
	v.appendString(" ")
}

func (v *formatter_) ProcessProcedureSlot(
	procedure ast.ProcedureLike,
	slot_ uint,
) {
	switch slot_ {
	case 1:
		v.depth_++
	case 2:
		v.depth_--
		v.appendNewline()
	}
}

func (v *formatter_) ProcessPublishClauseSlot(
	publishClause ast.PublishClauseLike,
	slot_ uint,
) {
	v.appendString(" ")
}

func (v *formatter_) ProcessRejectClauseSlot(
	rejectClause ast.RejectClauseLike,
	slot_ uint,
) {
	v.appendString(" ")
}

func (v *formatter_) ProcessRetrieveClauseSlot(
	retrieveClause ast.RetrieveClauseLike,
	slot_ uint,
) {
	v.appendString(" ")
}

func (v *formatter_) ProcessReturnClauseSlot(
	returnClause ast.ReturnClauseLike,
	slot_ uint,
) {
	v.appendString(" ")
}

func (v *formatter_) ProcessSaveClauseSlot(
	saveClause ast.SaveClauseLike,
	slot_ uint,
) {
	v.appendString(" ")
}

func (v *formatter_) ProcessSelectClauseSlot(
	selectClause ast.SelectClauseLike,
	slot_ uint,
) {
	switch slot_ {
	case 1:
		v.appendString(" ")
	}
}

func (v *formatter_) PreprocessStatement(
	statement ast.StatementLike,
	index_ uint,
	count_ uint,
) {
	v.appendNewline()
}

func (v *formatter_) ProcessThrowClauseSlot(
	throwClause ast.ThrowClauseLike,
	slot_ uint,
) {
	v.appendString(" ")
}

func (v *formatter_) ProcessWhileClauseSlot(
	whileClause ast.WhileClauseLike,
	slot_ uint,
) {
	v.appendString(" ")
}

func (v *formatter_) ProcessWithClauseSlot(
	withClause ast.WithClauseLike,
	slot_ uint,
) {
	v.appendString(" ")
}

const _indentation = "    "

func (v *formatter_) adjustIndentation(
	multiline string,
) string {
	// Calculate the current indentation level.
	var indentation string
	var count uint
	for ; count < v.depth_; count++ {
		indentation += "    "
	}

	// Adjust the indentation level in each line in the multiline string.
	var lines = sts.Split(multiline, "\n")
	var lastIndex = len(lines) - 1
	var cut = len(lines[lastIndex]) - 2 // The previous indentation level.
	for index, line := range lines {
		if index > 0 {
			// The first line is a two character delimiter that is never indented.
			if len(line) < cut {
				// This is a blank line that should still be indented.
				lines[index] = indentation
			} else {
				// Replace previous indentation with the new one.
				lines[index] = indentation + line[cut:]
			}
		}
	}
	multiline = sts.Join(lines, "\n")

	return multiline
}

// PROTECTED INTERFACE

// Private Methods

func (v *formatter_) appendNewline() {
	var newline = "\n"
	var level uint
	for ; level < v.depth_; level++ {
		newline += _indentation
	}
	v.appendString(newline)
}

func (v *formatter_) appendString(
	string_ string,
) {
	v.result_.WriteString(string_)
}

func (v *formatter_) getResult() string {
	var result = v.result_.String()
	v.result_.Reset()
	return result
}

// Instance Structure

type formatter_ struct {
	// Declare the instance attributes.
	visitor_ VisitorLike
	depth_   uint
	result_  sts.Builder

	// Declare the inherited aspects.
	Methodical
}

// Class Structure

type formatterClass_ struct {
	// Declare the class constants.
}

// Class Reference

func formatterClass() *formatterClass_ {
	return formatterClassReference_
}

var formatterClassReference_ = &formatterClass_{
	// Initialize the class constants.
}
