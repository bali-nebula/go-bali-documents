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

package documents

import (
	uti "github.com/craterdog/go-missing-utilities/v8"
)

// CLASS INTERFACE

// Access Function

func DiscardClauseClass() DiscardClauseClassLike {
	return discardClauseClass()
}

// Constructor Methods

func (c *discardClauseClass_) DiscardClause(
	citation ExpressionLike,
) DiscardClauseLike {
	if uti.IsUndefined(citation) {
		panic("The \"citation\" attribute is required by this class.")
	}
	var instance = &discardClause_{
		// Initialize the instance attributes.
		citation_: citation,
	}
	return instance
}

// Constant Methods

// Function Methods

// INSTANCE INTERFACE

// Principal Methods

func (v *discardClause_) GetClass() DiscardClauseClassLike {
	return discardClauseClass()
}

// Attribute Methods

func (v *discardClause_) GetCitation() ExpressionLike {
	return v.citation_
}

// PROTECTED INTERFACE

// Private Methods

// Instance Structure

type discardClause_ struct {
	// Declare the instance attributes.
	citation_ ExpressionLike
}

// Class Structure

type discardClauseClass_ struct {
	// Declare the class constants.
}

// Class Reference

func discardClauseClass() *discardClauseClass_ {
	return discardClauseClassReference_
}

var discardClauseClassReference_ = &discardClauseClass_{
	// Initialize the class constants.
}
