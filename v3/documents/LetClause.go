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

func LetClauseClass() LetClauseClassLike {
	return letClauseClass()
}

// Constructor Methods

func (c *letClauseClass_) LetClause(
	delimiter string,
	recipient RecipientLike,
	assignment AssignmentLike,
	expression ExpressionLike,
) LetClauseLike {
	if uti.IsUndefined(delimiter) {
		panic("The \"delimiter\" attribute is required by this class.")
	}
	if uti.IsUndefined(recipient) {
		panic("The \"recipient\" attribute is required by this class.")
	}
	if uti.IsUndefined(assignment) {
		panic("The \"assignment\" attribute is required by this class.")
	}
	if uti.IsUndefined(expression) {
		panic("The \"expression\" attribute is required by this class.")
	}
	var instance = &letClause_{
		// Initialize the instance attributes.
		delimiter_:  delimiter,
		recipient_:  recipient,
		assignment_: assignment,
		expression_: expression,
	}
	return instance
}

// INSTANCE INTERFACE

// Principal Methods

func (v *letClause_) GetClass() LetClauseClassLike {
	return letClauseClass()
}

// Attribute Methods

func (v *letClause_) GetDelimiter() string {
	return v.delimiter_
}

func (v *letClause_) GetRecipient() RecipientLike {
	return v.recipient_
}

func (v *letClause_) GetAssignment() AssignmentLike {
	return v.assignment_
}

func (v *letClause_) GetExpression() ExpressionLike {
	return v.expression_
}

// PROTECTED INTERFACE

// Instance Structure

type letClause_ struct {
	// Declare the instance attributes.
	delimiter_  string
	recipient_  RecipientLike
	assignment_ AssignmentLike
	expression_ ExpressionLike
}

// Class Structure

type letClauseClass_ struct {
	// Declare the class constants.
}

// Class Reference

func letClauseClass() *letClauseClass_ {
	return letClauseClassReference_
}

var letClauseClassReference_ = &letClauseClass_{
	// Initialize the class constants.
}
