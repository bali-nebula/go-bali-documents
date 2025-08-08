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
	uti "github.com/craterdog/go-missing-utilities/v7"
)

// CLASS INTERFACE

// Access Function

func NotarizeClauseClass() NotarizeClauseClassLike {
	return notarizeClauseClass()
}

// Constructor Methods

func (c *notarizeClauseClass_) NotarizeClause(
	draft ExpressionLike,
	cited ExpressionLike,
) NotarizeClauseLike {
	if uti.IsUndefined(draft) {
		panic("The \"draft\" attribute is required by this class.")
	}
	if uti.IsUndefined(cited) {
		panic("The \"cited\" attribute is required by this class.")
	}
	var instance = &notarizeClause_{
		// Initialize the instance attributes.
		draft_: draft,
		cited_: cited,
	}
	return instance
}

// Constant Methods

// Function Methods

// INSTANCE INTERFACE

// Principal Methods

func (v *notarizeClause_) GetClass() NotarizeClauseClassLike {
	return notarizeClauseClass()
}

// Attribute Methods

func (v *notarizeClause_) GetDraft() ExpressionLike {
	return v.draft_
}

func (v *notarizeClause_) GetCited() ExpressionLike {
	return v.cited_
}

// PROTECTED INTERFACE

// Private Methods

// Instance Structure

type notarizeClause_ struct {
	// Declare the instance attributes.
	draft_ ExpressionLike
	cited_ ExpressionLike
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
