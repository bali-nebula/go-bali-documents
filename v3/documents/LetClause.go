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

func LetClauseClass() LetClauseClassLike {
	return letClauseClass()
}

// Constructor Methods

func (c *letClauseClass_) LetClause(
	recipient any,
	assignment Assignment,
	expression ExpressionLike,
) LetClauseLike {
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
		recipient_:  recipient,
		assignment_: assignment,
		expression_: expression,
	}
	return instance
}

// Constant Methods

// Function Methods

// INSTANCE INTERFACE

// Principal Methods

func (v *letClause_) GetClass() LetClauseClassLike {
	return letClauseClass()
}

// Attribute Methods

func (v *letClause_) GetRecipient() any {
	return v.recipient_
}

func (v *letClause_) GetAssignment() Assignment {
	return v.assignment_
}

func (v *letClause_) GetExpression() ExpressionLike {
	return v.expression_
}

// PROTECTED INTERFACE

// Private Methods

// Instance Structure

type letClause_ struct {
	// Declare the instance attributes.
	recipient_  any
	assignment_ Assignment
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
