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

func PublishClauseClass() PublishClauseClassLike {
	return publishClauseClass()
}

// Constructor Methods

func (c *publishClauseClass_) PublishClause(
	event ExpressionLike,
) PublishClauseLike {
	if uti.IsUndefined(event) {
		panic("The \"event\" attribute is required by this class.")
	}
	var instance = &publishClause_{
		// Initialize the instance attributes.
		event_: event,
	}
	return instance
}

// Constant Methods

// Function Methods

// INSTANCE INTERFACE

// Principal Methods

func (v *publishClause_) GetClass() PublishClauseClassLike {
	return publishClauseClass()
}

// Attribute Methods

func (v *publishClause_) GetEvent() ExpressionLike {
	return v.event_
}

// PROTECTED INTERFACE

// Private Methods

// Instance Structure

type publishClause_ struct {
	// Declare the instance attributes.
	event_ ExpressionLike
}

// Class Structure

type publishClauseClass_ struct {
	// Declare the class constants.
}

// Class Reference

func publishClauseClass() *publishClauseClass_ {
	return publishClauseClassReference_
}

var publishClauseClassReference_ = &publishClauseClass_{
	// Initialize the class constants.
}
