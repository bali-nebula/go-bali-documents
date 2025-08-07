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

func OnClauseClass() OnClauseClassLike {
	return onClauseClass()
}

// Constructor Methods

func (c *onClauseClass_) OnClause(
	failure fra.SymbolLike,
	matchingClauses fra.ListLike[MatchingClauseLike],
) OnClauseLike {
	if uti.IsUndefined(failure) {
		panic("The \"failure\" attribute is required by this class.")
	}
	if uti.IsUndefined(matchingClauses) {
		panic("The \"matchingClauses\" attribute is required by this class.")
	}
	var instance = &onClause_{
		// Initialize the instance attributes.
		failure_:         failure,
		matchingClauses_: matchingClauses,
	}
	return instance
}

// INSTANCE INTERFACE

// Principal Methods

func (v *onClause_) GetClass() OnClauseClassLike {
	return onClauseClass()
}

// Attribute Methods

func (v *onClause_) GetFailure() fra.SymbolLike {
	return v.failure_
}

func (v *onClause_) GetMatchingClauses() fra.ListLike[MatchingClauseLike] {
	return v.matchingClauses_
}

// PROTECTED INTERFACE

// Instance Structure

type onClause_ struct {
	// Declare the instance attributes.
	failure_         fra.SymbolLike
	matchingClauses_ fra.ListLike[MatchingClauseLike]
}

// Class Structure

type onClauseClass_ struct {
	// Declare the class constants.
}

// Class Reference

func onClauseClass() *onClauseClass_ {
	return onClauseClassReference_
}

var onClauseClassReference_ = &onClauseClass_{
	// Initialize the class constants.
}
