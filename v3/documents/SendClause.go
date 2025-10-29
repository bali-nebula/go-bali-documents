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

func SendClauseClass() SendClauseClassLike {
	return sendClauseClass()
}

// Constructor Methods

func (c *sendClauseClass_) SendClause(
	message ExpressionLike,
	bag ExpressionLike,
) SendClauseLike {
	if uti.IsUndefined(message) {
		panic("The \"message\" attribute is required by this class.")
	}
	if uti.IsUndefined(bag) {
		panic("The \"bag\" attribute is required by this class.")
	}
	var instance = &sendClause_{
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

func (v *sendClause_) GetClass() SendClauseClassLike {
	return sendClauseClass()
}

// Attribute Methods

func (v *sendClause_) GetMessage() ExpressionLike {
	return v.message_
}

func (v *sendClause_) GetBag() ExpressionLike {
	return v.bag_
}

// PROTECTED INTERFACE

// Private Methods

// Instance Structure

type sendClause_ struct {
	// Declare the instance attributes.
	message_ ExpressionLike
	bag_     ExpressionLike
}

// Class Structure

type sendClauseClass_ struct {
	// Declare the class constants.
}

// Class Reference

func sendClauseClass() *sendClauseClass_ {
	return sendClauseClassReference_
}

var sendClauseClassReference_ = &sendClauseClass_{
	// Initialize the class constants.
}
