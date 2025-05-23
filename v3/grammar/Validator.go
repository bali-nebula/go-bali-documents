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
	fmt "fmt"
	ast "github.com/bali-nebula/go-bali-documents/v3/ast"
)

// CLASS INTERFACE

// Access Function

func ValidatorClass() ValidatorClassLike {
	return validatorClass()
}

// Constructor Methods

func (c *validatorClass_) Validator() ValidatorLike {
	var instance = &validator_{
		// Initialize the instance attributes.

		// Initialize the inherited aspects.
		Methodical: ProcessorClass().Processor(),
	}
	instance.visitor_ = VisitorClass().Visitor(instance)
	return instance
}

// INSTANCE INTERFACE

// Principal Methods

func (v *validator_) GetClass() ValidatorClassLike {
	return validatorClass()
}

func (v *validator_) ValidateDocument(
	document ast.DocumentLike,
) {
	v.visitor_.VisitDocument(document)
}

// Methodical Methods

func (v *validator_) ProcessAngle(
	angle string,
) {
	v.validateToken(angle, AngleToken)
}

func (v *validator_) ProcessArrow(
	arrow string,
) {
	v.validateToken(arrow, ArrowToken)
}

func (v *validator_) ProcessBinary(
	binary string,
) {
	v.validateToken(binary, BinaryToken)
}

func (v *validator_) ProcessBoolean(
	boolean string,
) {
	v.validateToken(boolean, BooleanToken)
}

func (v *validator_) ProcessBytecode(
	bytecode string,
) {
	v.validateToken(bytecode, BytecodeToken)
}

func (v *validator_) ProcessCaret(
	caret string,
) {
	v.validateToken(caret, CaretToken)
}

func (v *validator_) ProcessCitation(
	citation string,
) {
	v.validateToken(citation, CitationToken)
}

func (v *validator_) ProcessComment(
	comment string,
) {
	v.validateToken(comment, CommentToken)
}

func (v *validator_) ProcessDot(
	dot string,
) {
	v.validateToken(dot, DotToken)
}

func (v *validator_) ProcessDuration(
	duration string,
) {
	v.validateToken(duration, DurationToken)
}

func (v *validator_) ProcessEqual(
	equal string,
) {
	v.validateToken(equal, EqualToken)
}

func (v *validator_) ProcessIdentifier(
	identifier string,
) {
	v.validateToken(identifier, IdentifierToken)
}

func (v *validator_) ProcessLess(
	less string,
) {
	v.validateToken(less, LessToken)
}

func (v *validator_) ProcessMinus(
	minus string,
) {
	v.validateToken(minus, MinusToken)
}

func (v *validator_) ProcessMoment(
	moment string,
) {
	v.validateToken(moment, MomentToken)
}

func (v *validator_) ProcessMore(
	more string,
) {
	v.validateToken(more, MoreToken)
}

func (v *validator_) ProcessName(
	name string,
) {
	v.validateToken(name, NameToken)
}

func (v *validator_) ProcessNarrative(
	narrative string,
) {
	v.validateToken(narrative, NarrativeToken)
}

func (v *validator_) ProcessNewline(
	newline string,
) {
	v.validateToken(newline, NewlineToken)
}

func (v *validator_) ProcessNote(
	note string,
) {
	v.validateToken(note, NoteToken)
}

func (v *validator_) ProcessNumber(
	number string,
) {
	v.validateToken(number, NumberToken)
}

func (v *validator_) ProcessPattern(
	pattern string,
) {
	v.validateToken(pattern, PatternToken)
}

func (v *validator_) ProcessPercent(
	percent string,
) {
	v.validateToken(percent, PercentToken)
}

func (v *validator_) ProcessPercentage(
	percentage string,
) {
	v.validateToken(percentage, PercentageToken)
}

func (v *validator_) ProcessPlus(
	plus string,
) {
	v.validateToken(plus, PlusToken)
}

func (v *validator_) ProcessProbability(
	probability string,
) {
	v.validateToken(probability, ProbabilityToken)
}

func (v *validator_) ProcessQuote(
	quote string,
) {
	v.validateToken(quote, QuoteToken)
}

func (v *validator_) ProcessResource(
	resource string,
) {
	v.validateToken(resource, ResourceToken)
}

func (v *validator_) ProcessSlash(
	slash string,
) {
	v.validateToken(slash, SlashToken)
}

func (v *validator_) ProcessSpace(
	space string,
) {
	v.validateToken(space, SpaceToken)
}

func (v *validator_) ProcessStar(
	star string,
) {
	v.validateToken(star, StarToken)
}

func (v *validator_) ProcessSymbol(
	symbol string,
) {
	v.validateToken(symbol, SymbolToken)
}

func (v *validator_) ProcessTag(
	tag string,
) {
	v.validateToken(tag, TagToken)
}

func (v *validator_) ProcessVersion(
	version string,
) {
	v.validateToken(version, VersionToken)
}

// PROTECTED INTERFACE

// Private Methods

func (v *validator_) validateToken(
	tokenValue string,
	tokenType TokenType,
) {
	var scannerClass = ScannerClass()
	if !scannerClass.MatchesType(tokenValue, tokenType) {
		var message = fmt.Sprintf(
			"The following token value is not of type %v: %v",
			scannerClass.FormatType(tokenType),
			tokenValue,
		)
		panic(message)
	}
}

// Instance Structure

type validator_ struct {
	// Declare the instance attributes.
	visitor_ VisitorLike

	// Declare the inherited aspects.
	Methodical
}

// Class Structure

type validatorClass_ struct {
	// Declare the class constants.
}

// Class Reference

func validatorClass() *validatorClass_ {
	return validatorClassReference_
}

var validatorClassReference_ = &validatorClass_{
	// Initialize the class constants.
}
