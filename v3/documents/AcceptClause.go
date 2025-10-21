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
	uti "github.com/craterdog/go-missing-utilities/v8"
)

// CLASS INTERFACE

// Access Function

func AcceptClauseClass() AcceptClauseClassLike {
	return acceptClauseClass()
}

// Constructor Methods

func (c *acceptClauseClass_) AcceptClause(
	message ExpressionLike,
	bag ExpressionLike,
) AcceptClauseLike {
	if uti.IsUndefined(message) {
		panic("The \"message\" attribute is required by this class.")
	}
	if uti.IsUndefined(bag) {
		panic("The \"bag\" attribute is required by this class.")
	}
	var instance = &acceptClause_{
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

func (v *acceptClause_) GetClass() AcceptClauseClassLike {
	return acceptClauseClass()
}

// Attribute Methods

func (v *acceptClause_) GetMessage() ExpressionLike {
	return v.message_
}

func (v *acceptClause_) GetBag() ExpressionLike {
	return v.bag_
}

// PROTECTED INTERFACE

// Private Methods

// Instance Structure

type acceptClause_ struct {
	// Declare the instance attributes.
	message_ ExpressionLike
	bag_     ExpressionLike
}

// Class Structure

type acceptClauseClass_ struct {
	// Declare the class constants.
}

// Class Reference

func acceptClauseClass() *acceptClauseClass_ {
	return acceptClauseClassReference_
}

var acceptClauseClassReference_ = &acceptClauseClass_{
	// Initialize the class constants.
}
