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

func RejectClauseClass() RejectClauseClassLike {
	return rejectClauseClass()
}

// Constructor Methods

func (c *rejectClauseClass_) RejectClause(
	message ExpressionLike,
	bag ExpressionLike,
) RejectClauseLike {
	if uti.IsUndefined(message) {
		panic("The \"message\" attribute is required by this class.")
	}
	if uti.IsUndefined(bag) {
		panic("The \"bag\" attribute is required by this class.")
	}
	var instance = &rejectClause_{
		// Initialize the instance attributes.
		message_: message,
		bag_:     bag,
	}
	return instance
}

// Constant Methods

// Function Methods

// INSTANCE INTERFACE

// Principal Methods

func (v *rejectClause_) GetClass() RejectClauseClassLike {
	return rejectClauseClass()
}

// Attribute Methods

func (v *rejectClause_) GetMessage() ExpressionLike {
	return v.message_
}
func (v *rejectClause_) GetBag() ExpressionLike {
	return v.bag_
}

// PROTECTED INTERFACE

// Private Methods

// Instance Structure

type rejectClause_ struct {
	// Declare the instance attributes.
	message_ ExpressionLike
	bag_     ExpressionLike
}

// Class Structure

type rejectClauseClass_ struct {
	// Declare the class constants.
}

// Class Reference

func rejectClauseClass() *rejectClauseClass_ {
	return rejectClauseClassReference_
}

var rejectClauseClassReference_ = &rejectClauseClass_{
	// Initialize the class constants.
}
