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
	ele "github.com/bali-nebula/go-bali-documents/v3/elements"
	fra "github.com/craterdog/go-collection-framework/v8"
	uti "github.com/craterdog/go-missing-utilities/v8"
)

// CLASS INTERFACE

// Access Function

func OnClauseClass() OnClauseClassLike {
	return onClauseClass()
}

// Constructor Methods

func (c *onClauseClass_) OnClause(
	symbol ele.SymbolLike,
	matchingClauses fra.Sequential[MatchingClauseLike],
) OnClauseLike {
	if uti.IsUndefined(symbol) {
		panic("The \"symbol\" attribute is required by this class.")
	}
	if uti.IsUndefined(matchingClauses) {
		panic("The \"matchingClauses\" attribute is required by this class.")
	}
	var instance = &onClause_{
		// Initialize the instance attributes.
		symbol_:          symbol,
		matchingClauses_: matchingClauses,
	}
	return instance
}

// Constant Methods

// Function Methods

// INSTANCE INTERFACE

// Principal Methods

func (v *onClause_) GetClass() OnClauseClassLike {
	return onClauseClass()
}

// Attribute Methods

func (v *onClause_) GetSymbol() ele.SymbolLike {
	return v.symbol_
}

func (v *onClause_) GetMatchingClauses() fra.Sequential[MatchingClauseLike] {
	return v.matchingClauses_
}

// PROTECTED INTERFACE

// Private Methods

// Instance Structure

type onClause_ struct {
	// Declare the instance attributes.
	symbol_          ele.SymbolLike
	matchingClauses_ fra.Sequential[MatchingClauseLike]
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
