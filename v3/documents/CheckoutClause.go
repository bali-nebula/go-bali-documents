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
	delimiter1 string,
	recipient RecipientLike,
	optionalAtLevel AtLevelLike,
	delimiter2 string,
	cited CitedLike,
) CheckoutClauseLike {
	if uti.IsUndefined(delimiter1) {
		panic("The \"delimiter1\" attribute is required by this class.")
	}
	if uti.IsUndefined(recipient) {
		panic("The \"recipient\" attribute is required by this class.")
	}
	if uti.IsUndefined(delimiter2) {
		panic("The \"delimiter2\" attribute is required by this class.")
	}
	if uti.IsUndefined(cited) {
		panic("The \"cited\" attribute is required by this class.")
	}
	var instance = &checkoutClause_{
		// Initialize the instance attributes.
		delimiter1_:      delimiter1,
		recipient_:       recipient,
		optionalAtLevel_: optionalAtLevel,
		delimiter2_:      delimiter2,
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

func (v *checkoutClause_) GetDelimiter1() string {
	return v.delimiter1_
}

func (v *checkoutClause_) GetRecipient() RecipientLike {
	return v.recipient_
}

func (v *checkoutClause_) GetOptionalAtLevel() AtLevelLike {
	return v.optionalAtLevel_
}

func (v *checkoutClause_) GetDelimiter2() string {
	return v.delimiter2_
}

func (v *checkoutClause_) GetCited() CitedLike {
	return v.cited_
}

// PROTECTED INTERFACE

// Instance Structure

type checkoutClause_ struct {
	// Declare the instance attributes.
	delimiter1_      string
	recipient_       RecipientLike
	optionalAtLevel_ AtLevelLike
	delimiter2_      string
	cited_           CitedLike
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
