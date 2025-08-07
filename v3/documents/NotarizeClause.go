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

func NotarizeClauseClass() NotarizeClauseClassLike {
	return notarizeClauseClass()
}

// Constructor Methods

func (c *notarizeClauseClass_) NotarizeClause(
	delimiter1 string,
	draft DraftLike,
	delimiter2 string,
	cited CitedLike,
) NotarizeClauseLike {
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
	var instance = &notarizeClause_{
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

func (v *notarizeClause_) GetClass() NotarizeClauseClassLike {
	return notarizeClauseClass()
}

// Attribute Methods

func (v *notarizeClause_) GetDelimiter1() string {
	return v.delimiter1_
}

func (v *notarizeClause_) GetDraft() DraftLike {
	return v.draft_
}

func (v *notarizeClause_) GetDelimiter2() string {
	return v.delimiter2_
}

func (v *notarizeClause_) GetCited() CitedLike {
	return v.cited_
}

// PROTECTED INTERFACE

// Instance Structure

type notarizeClause_ struct {
	// Declare the instance attributes.
	delimiter1_ string
	draft_      DraftLike
	delimiter2_ string
	cited_      CitedLike
}

// Class Structure

type notarizeClauseClass_ struct {
	// Declare the class constants.
}

// Class Reference

func notarizeClauseClass() *notarizeClauseClass_ {
	return notarizeClauseClassReference_
}

var notarizeClauseClassReference_ = &notarizeClauseClass_{
	// Initialize the class constants.
}
