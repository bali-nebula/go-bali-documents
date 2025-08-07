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

func CheckoutClauseClass() CheckoutClauseClassLike {
	return checkoutClauseClass()
}

// Constructor Methods

func (c *checkoutClauseClass_) CheckoutClause(
	recipient any,
	optionalAtLevel ExpressionLike,
	cited ExpressionLike,
) CheckoutClauseLike {
	if uti.IsUndefined(recipient) {
		panic("The \"recipient\" attribute is required by this class.")
	}
	if uti.IsUndefined(cited) {
		panic("The \"cited\" attribute is required by this class.")
	}
	var instance = &checkoutClause_{
		// Initialize the instance attributes.
		recipient_:       recipient,
		optionalAtLevel_: optionalAtLevel,
		cited_:           cited,
	}
	return instance
}

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

func (v *checkoutClause_) GetCited() ExpressionLike {
	return v.cited_
}

// PROTECTED INTERFACE

// Instance Structure

type checkoutClause_ struct {
	// Declare the instance attributes.
	recipient_       any
	optionalAtLevel_ ExpressionLike
	cited_           ExpressionLike
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
