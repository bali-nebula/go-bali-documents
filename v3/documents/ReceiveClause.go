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

func ReceiveClauseClass() ReceiveClauseClassLike {
	return receiveClauseClass()
}

// Constructor Methods

func (c *receiveClauseClass_) ReceiveClause(
	recipient any,
	bag ExpressionLike,
) ReceiveClauseLike {
	if uti.IsUndefined(recipient) {
		panic("The \"recipient\" attribute is required by this class.")
	}
	if uti.IsUndefined(bag) {
		panic("The \"bag\" attribute is required by this class.")
	}
	var instance = &receiveClause_{
		// Initialize the instance attributes.
		recipient_: recipient,
		bag_:       bag,
	}
	return instance
}

// Constant Methods

// Function Methods

// INSTANCE INTERFACE

// Principal Methods

func (v *receiveClause_) GetClass() ReceiveClauseClassLike {
	return receiveClauseClass()
}

// Attribute Methods

func (v *receiveClause_) GetRecipient() any {
	return v.recipient_
}

func (v *receiveClause_) GetBag() ExpressionLike {
	return v.bag_
}

// PROTECTED INTERFACE

// Private Methods

// Instance Structure

type receiveClause_ struct {
	// Declare the instance attributes.
	recipient_ any
	bag_       ExpressionLike
}

// Class Structure

type receiveClauseClass_ struct {
	// Declare the class constants.
}

// Class Reference

func receiveClauseClass() *receiveClauseClass_ {
	return receiveClauseClassReference_
}

var receiveClauseClassReference_ = &receiveClauseClass_{
	// Initialize the class constants.
}
