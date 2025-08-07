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

func PostClauseClass() PostClauseClassLike {
	return postClauseClass()
}

// Constructor Methods

func (c *postClauseClass_) PostClause(
	delimiter1 string,
	message MessageLike,
	delimiter2 string,
	bag BagLike,
) PostClauseLike {
	if uti.IsUndefined(delimiter1) {
		panic("The \"delimiter1\" attribute is required by this class.")
	}
	if uti.IsUndefined(message) {
		panic("The \"message\" attribute is required by this class.")
	}
	if uti.IsUndefined(delimiter2) {
		panic("The \"delimiter2\" attribute is required by this class.")
	}
	if uti.IsUndefined(bag) {
		panic("The \"bag\" attribute is required by this class.")
	}
	var instance = &postClause_{
		// Initialize the instance attributes.
		delimiter1_: delimiter1,
		message_:    message,
		delimiter2_: delimiter2,
		bag_:        bag,
	}
	return instance
}

// INSTANCE INTERFACE

// Principal Methods

func (v *postClause_) GetClass() PostClauseClassLike {
	return postClauseClass()
}

// Attribute Methods

func (v *postClause_) GetDelimiter1() string {
	return v.delimiter1_
}

func (v *postClause_) GetMessage() MessageLike {
	return v.message_
}

func (v *postClause_) GetDelimiter2() string {
	return v.delimiter2_
}

func (v *postClause_) GetBag() BagLike {
	return v.bag_
}

// PROTECTED INTERFACE

// Instance Structure

type postClause_ struct {
	// Declare the instance attributes.
	delimiter1_ string
	message_    MessageLike
	delimiter2_ string
	bag_        BagLike
}

// Class Structure

type postClauseClass_ struct {
	// Declare the class constants.
}

// Class Reference

func postClauseClass() *postClauseClass_ {
	return postClauseClassReference_
}

var postClauseClassReference_ = &postClauseClass_{
	// Initialize the class constants.
}
