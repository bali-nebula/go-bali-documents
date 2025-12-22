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

func AssignClauseClass() AssignClauseClassLike {
	return assignClauseClass()
}

// Constructor Methods

func (c *assignClauseClass_) AssignClause(
	recipient any,
	assignment Assignment,
	expression ExpressionLike,
) AssignClauseLike {
	if uti.IsUndefined(recipient) {
		panic("The \"recipient\" attribute is required by this class.")
	}
	if uti.IsUndefined(assignment) {
		panic("The \"assignment\" attribute is required by this class.")
	}
	if uti.IsUndefined(expression) {
		panic("The \"expression\" attribute is required by this class.")
	}
	var instance = &assignClause_{
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

func (v *assignClause_) GetClass() AssignClauseClassLike {
	return assignClauseClass()
}

// Attribute Methods

func (v *assignClause_) GetRecipient() any {
	return v.recipient_
}

func (v *assignClause_) GetAssignment() Assignment {
	return v.assignment_
}

func (v *assignClause_) GetExpression() ExpressionLike {
	return v.expression_
}

// PROTECTED INTERFACE

// Private Methods

// Instance Structure

type assignClause_ struct {
	// Declare the instance attributes.
	recipient_  any
	assignment_ Assignment
	expression_ ExpressionLike
}

// Class Structure

type assignClauseClass_ struct {
	// Declare the class constants.
}

// Class Reference

func assignClauseClass() *assignClauseClass_ {
	return assignClauseClassReference_
}

var assignClauseClassReference_ = &assignClauseClass_{
	// Initialize the class constants.
}
