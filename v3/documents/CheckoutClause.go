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

func CheckoutClauseClass() CheckoutClauseClassLike {
	return checkoutClauseClass()
}

// Constructor Methods

func (c *checkoutClauseClass_) CheckoutClause(
	recipient any,
	optionalAtLevel ExpressionLike,
	location ExpressionLike,
) CheckoutClauseLike {
	if uti.IsUndefined(recipient) {
		panic("The \"recipient\" attribute is required by this class.")
	}
	if uti.IsUndefined(location) {
		panic("The \"location\" attribute is required by this class.")
	}
	var instance = &checkoutClause_{
		// Initialize the instance attributes.
		recipient_:       recipient,
		optionalAtLevel_: optionalAtLevel,
		location_:        location,
	}
	return instance
}

// Constant Methods

// Function Methods

// INSTANCE INTERFACE

// Principal Methods

func (v *checkoutClause_) GetClass() CheckoutClauseClassLike {
	return checkoutClauseClass()
}

// Attribute Methods

func (v *checkoutClause_) GetRecipient() any {
	return v.recipient_
}

func (v *checkoutClause_) GetOptionalAtLevel() ExpressionLike {
	return v.optionalAtLevel_
}

func (v *checkoutClause_) GetLocation() ExpressionLike {
	return v.location_
}

// PROTECTED INTERFACE

// Private Methods

// Instance Structure

type checkoutClause_ struct {
	// Declare the instance attributes.
	recipient_       any
	optionalAtLevel_ ExpressionLike
	location_        ExpressionLike
}

// Class Structure

type checkoutClauseClass_ struct {
	// Declare the class constants.
}

// Class Reference

func checkoutClauseClass() *checkoutClauseClass_ {
	return checkoutClauseClassReference_
}

var checkoutClauseClassReference_ = &checkoutClauseClass_{
	// Initialize the class constants.
}
