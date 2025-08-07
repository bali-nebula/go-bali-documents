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

func RetrieveClauseClass() RetrieveClauseClassLike {
	return retrieveClauseClass()
}

// Constructor Methods

func (c *retrieveClauseClass_) RetrieveClause(
	delimiter1 string,
	recipient RecipientLike,
	delimiter2 string,
	bag BagLike,
) RetrieveClauseLike {
	if uti.IsUndefined(delimiter1) {
		panic("The \"delimiter1\" attribute is required by this class.")
	}
	if uti.IsUndefined(recipient) {
		panic("The \"recipient\" attribute is required by this class.")
	}
	if uti.IsUndefined(delimiter2) {
		panic("The \"delimiter2\" attribute is required by this class.")
	}
	if uti.IsUndefined(bag) {
		panic("The \"bag\" attribute is required by this class.")
	}
	var instance = &retrieveClause_{
		// Initialize the instance attributes.
		delimiter1_: delimiter1,
		recipient_:  recipient,
		delimiter2_: delimiter2,
		bag_:        bag,
	}
	return instance
}

// INSTANCE INTERFACE

// Principal Methods

func (v *retrieveClause_) GetClass() RetrieveClauseClassLike {
	return retrieveClauseClass()
}

// Attribute Methods

func (v *retrieveClause_) GetDelimiter1() string {
	return v.delimiter1_
}

func (v *retrieveClause_) GetRecipient() RecipientLike {
	return v.recipient_
}

func (v *retrieveClause_) GetDelimiter2() string {
	return v.delimiter2_
}

func (v *retrieveClause_) GetBag() BagLike {
	return v.bag_
}

// PROTECTED INTERFACE

// Instance Structure

type retrieveClause_ struct {
	// Declare the instance attributes.
	delimiter1_ string
	recipient_  RecipientLike
	delimiter2_ string
	bag_        BagLike
}

// Class Structure

type retrieveClauseClass_ struct {
	// Declare the class constants.
}

// Class Reference

func retrieveClauseClass() *retrieveClauseClass_ {
	return retrieveClauseClassReference_
}

var retrieveClauseClassReference_ = &retrieveClauseClass_{
	// Initialize the class constants.
}
