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
	com "github.com/craterdog/go-essential-composites/v8"
	uti "github.com/craterdog/go-essential-utilities/v8"
)

// CLASS INTERFACE

// Access Function

func SelectClauseClass() SelectClauseClassLike {
	return selectClauseClass()
}

// Constructor Methods

func (c *selectClauseClass_) SelectClause(
	template ExpressionLike,
	matchingClauses com.Sequential[MatchingClauseLike],
) SelectClauseLike {
	if uti.IsUndefined(template) {
		panic("The \"template\" attribute is required by this class.")
	}
	if uti.IsUndefined(matchingClauses) {
		panic("The \"matchingClauses\" attribute is required by this class.")
	}
	var instance = &selectClause_{
		// Initialize the instance attributes.
		template_:        template,
		matchingClauses_: matchingClauses,
	}
	return instance
}

// Constant Methods

// Function Methods

// INSTANCE INTERFACE

// Principal Methods

func (v *selectClause_) GetClass() SelectClauseClassLike {
	return selectClauseClass()
}

// Attribute Methods

func (v *selectClause_) GetTemplate() ExpressionLike {
	return v.template_
}

func (v *selectClause_) GetMatchingClauses() com.Sequential[MatchingClauseLike] {
	return v.matchingClauses_
}

// PROTECTED INTERFACE

// Private Methods

// Instance Structure

type selectClause_ struct {
	// Declare the instance attributes.
	template_        ExpressionLike
	matchingClauses_ com.Sequential[MatchingClauseLike]
}

// Class Structure

type selectClauseClass_ struct {
	// Declare the class constants.
}

// Class Reference

func selectClauseClass() *selectClauseClass_ {
	return selectClauseClassReference_
}

var selectClauseClassReference_ = &selectClauseClass_{
	// Initialize the class constants.
}
