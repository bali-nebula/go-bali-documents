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

func SaveClauseClass() SaveClauseClassLike {
	return saveClauseClass()
}

// Constructor Methods

func (c *saveClauseClass_) SaveClause(
	delimiter1 string,
	draft DraftLike,
	delimiter2 string,
	cited CitedLike,
) SaveClauseLike {
	if uti.IsUndefined(delimiter1) {
		panic("The \"delimiter1\" attribute is required by this class.")
	}
	if uti.IsUndefined(draft) {
		panic("The \"draft\" attribute is required by this class.")
	}
	if uti.IsUndefined(delimiter2) {
		panic("The \"delimiter2\" attribute is required by this class.")
	}
	if uti.IsUndefined(cited) {
		panic("The \"cited\" attribute is required by this class.")
	}
	var instance = &saveClause_{
		// Initialize the instance attributes.
		delimiter1_: delimiter1,
		draft_:      draft,
		delimiter2_: delimiter2,
		cited_:      cited,
	}
	return instance
}

// INSTANCE INTERFACE

// Principal Methods

func (v *saveClause_) GetClass() SaveClauseClassLike {
	return saveClauseClass()
}

// Attribute Methods

func (v *saveClause_) GetDelimiter1() string {
	return v.delimiter1_
}

func (v *saveClause_) GetDraft() DraftLike {
	return v.draft_
}

func (v *saveClause_) GetDelimiter2() string {
	return v.delimiter2_
}

func (v *saveClause_) GetCited() CitedLike {
	return v.cited_
}

// PROTECTED INTERFACE

// Instance Structure

type saveClause_ struct {
	// Declare the instance attributes.
	delimiter1_ string
	draft_      DraftLike
	delimiter2_ string
	cited_      CitedLike
}

// Class Structure

type saveClauseClass_ struct {
	// Declare the class constants.
}

// Class Reference

func saveClauseClass() *saveClauseClass_ {
	return saveClauseClassReference_
}

var saveClauseClassReference_ = &saveClauseClass_{
	// Initialize the class constants.
}
