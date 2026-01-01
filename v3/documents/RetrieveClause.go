/*
................................................................................
.    Copyright (c) 2009-2026 Crater Dog Technologies™.  All Rights Reserved.   .
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

func RetrieveClauseClass() RetrieveClauseClassLike {
	return retrieveClauseClass()
}

// Constructor Methods

func (c *retrieveClauseClass_) RetrieveClause(
	recipient any,
	citation ExpressionLike,
) RetrieveClauseLike {
	if uti.IsUndefined(recipient) {
		panic("The \"recipient\" attribute is required by this class.")
	}
	if uti.IsUndefined(citation) {
		panic("The \"citation\" attribute is required by this class.")
	}
	var instance = &retrieveClause_{
		// Initialize the instance attributes.
		recipient_: recipient,
		citation_:  citation,
	}
	return instance
}

// Constant Methods

// Function Methods

// INSTANCE INTERFACE

// Principal Methods

func (v *retrieveClause_) GetClass() RetrieveClauseClassLike {
	return retrieveClauseClass()
}

// Attribute Methods

func (v *retrieveClause_) GetRecipient() any {
	return v.recipient_
}

func (v *retrieveClause_) GetCitation() ExpressionLike {
	return v.citation_
}

// PROTECTED INTERFACE

// Private Methods

// Instance Structure

type retrieveClause_ struct {
	// Declare the instance attributes.
	recipient_ any
	citation_  ExpressionLike
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
