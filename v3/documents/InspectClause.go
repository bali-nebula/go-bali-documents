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

func InspectClauseClass() InspectClauseClassLike {
	return inspectClauseClass()
}

// Constructor Methods

func (c *inspectClauseClass_) InspectClause(
	recipient any,
	location ExpressionLike,
) InspectClauseLike {
	if uti.IsUndefined(recipient) {
		panic("The \"recipient\" attribute is required by this class.")
	}
	if uti.IsUndefined(location) {
		panic("The \"location\" attribute is required by this class.")
	}
	var instance = &inspectClause_{
		// Initialize the instance attributes.
		recipient_: recipient,
		location_:  location,
	}
	return instance
}

// Constant Methods

// Function Methods

// INSTANCE INTERFACE

// Principal Methods

func (v *inspectClause_) GetClass() InspectClauseClassLike {
	return inspectClauseClass()
}

// Attribute Methods

func (v *inspectClause_) GetRecipient() any {
	return v.recipient_
}

func (v *inspectClause_) GetLocation() ExpressionLike {
	return v.location_
}

// PROTECTED INTERFACE

// Private Methods

// Instance Structure

type inspectClause_ struct {
	// Declare the instance attributes.
	recipient_ any
	location_  ExpressionLike
}

// Class Structure

type inspectClauseClass_ struct {
	// Declare the class constants.
}

// Class Reference

func inspectClauseClass() *inspectClauseClass_ {
	return inspectClauseClassReference_
}

var inspectClauseClassReference_ = &inspectClauseClass_{
	// Initialize the class constants.
}
