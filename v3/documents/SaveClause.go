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
	uti "github.com/craterdog/go-essential-utilities/v8"
)

// CLASS INTERFACE

// Access Function

func SaveClauseClass() SaveClauseClassLike {
	return saveClauseClass()
}

// Constructor Methods

func (c *saveClauseClass_) SaveClause(
	draft ExpressionLike,
	recipient any,
) SaveClauseLike {
	if uti.IsUndefined(draft) {
		panic("The \"draft\" attribute is required by this class.")
	}
	if uti.IsUndefined(recipient) {
		panic("The \"recipient\" attribute is required by this class.")
	}
	var instance = &saveClause_{
		// Initialize the instance attributes.
		draft_:     draft,
		recipient_: recipient,
	}
	return instance
}

// Constant Methods

// Function Methods

// INSTANCE INTERFACE

// Principal Methods

func (v *saveClause_) GetClass() SaveClauseClassLike {
	return saveClauseClass()
}

// Attribute Methods

func (v *saveClause_) GetDraft() ExpressionLike {
	return v.draft_
}

func (v *saveClause_) GetRecipient() any {
	return v.recipient_
}

// PROTECTED INTERFACE

// Private Methods

// Instance Structure

type saveClause_ struct {
	// Declare the instance attributes.
	draft_     ExpressionLike
	recipient_ any
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
