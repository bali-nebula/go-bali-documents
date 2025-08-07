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
│              This class file was automatically generated using:              │
│            https://github.com/craterdog/go-development-tools/wiki            │
│                                                                              │
│                     Any updates to it may be overwritten.                    │
└──────────────────────────────────────────────────────────────────────────────┘
*/

package documents

import (
	fra "github.com/craterdog/go-component-framework/v7"
	uti "github.com/craterdog/go-missing-utilities/v7"
)

// CLASS INTERFACE

// Access Function

func SelectClauseClass() SelectClauseClassLike {
	return selectClauseClass()
}

// Constructor Methods

func (c *selectClauseClass_) SelectClause(
	target any,
	matchingClauses fra.ListLike[MatchingClauseLike],
) SelectClauseLike {
	if uti.IsUndefined(target) {
		panic("The \"target\" attribute is required by this class.")
	}
	if uti.IsUndefined(matchingClauses) {
		panic("The \"matchingClauses\" attribute is required by this class.")
	}
	var instance = &selectClause_{
		// Initialize the instance attributes.
		target_:          target,
		matchingClauses_: matchingClauses,
	}
	return instance
}

// INSTANCE INTERFACE

// Principal Methods

func (v *selectClause_) GetClass() SelectClauseClassLike {
	return selectClauseClass()
}

// Attribute Methods

func (v *selectClause_) GetTarget() any {
	return v.target_
}

func (v *selectClause_) GetMatchingClauses() fra.ListLike[MatchingClauseLike] {
	return v.matchingClauses_
}

// PROTECTED INTERFACE

// Instance Structure

type selectClause_ struct {
	// Declare the instance attributes.
	target_          any
	matchingClauses_ fra.ListLike[MatchingClauseLike]
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
