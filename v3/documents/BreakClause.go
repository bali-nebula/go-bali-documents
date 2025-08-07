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
	uti "github.com/craterdog/go-missing-utilities/v7"
)

// CLASS INTERFACE

// Access Function

func BreakClauseClass() BreakClauseClassLike {
	return breakClauseClass()
}

// Constructor Methods

func (c *breakClauseClass_) BreakClause(
	delimiter1 string,
	delimiter2 string,
) BreakClauseLike {
	if uti.IsUndefined(delimiter1) {
		panic("The \"delimiter1\" attribute is required by this class.")
	}
	if uti.IsUndefined(delimiter2) {
		panic("The \"delimiter2\" attribute is required by this class.")
	}
	var instance = &breakClause_{
		// Initialize the instance attributes.
		delimiter1_: delimiter1,
		delimiter2_: delimiter2,
	}
	return instance
}

// INSTANCE INTERFACE

// Principal Methods

func (v *breakClause_) GetClass() BreakClauseClassLike {
	return breakClauseClass()
}

// Attribute Methods

func (v *breakClause_) GetDelimiter1() string {
	return v.delimiter1_
}

func (v *breakClause_) GetDelimiter2() string {
	return v.delimiter2_
}

// PROTECTED INTERFACE

// Instance Structure

type breakClause_ struct {
	// Declare the instance attributes.
	delimiter1_ string
	delimiter2_ string
}

// Class Structure

type breakClauseClass_ struct {
	// Declare the class constants.
}

// Class Reference

func breakClauseClass() *breakClauseClass_ {
	return breakClauseClassReference_
}

var breakClauseClassReference_ = &breakClauseClass_{
	// Initialize the class constants.
}
