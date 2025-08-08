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

func MatchingClauseClass() MatchingClauseClassLike {
	return matchingClauseClass()
}

// Constructor Methods

func (c *matchingClauseClass_) MatchingClause(
	template ExpressionLike,
	procedure ProcedureLike,
) MatchingClauseLike {
	if uti.IsUndefined(template) {
		panic("The \"template\" attribute is required by this class.")
	}
	if uti.IsUndefined(procedure) {
		panic("The \"procedure\" attribute is required by this class.")
	}
	var instance = &matchingClause_{
		// Initialize the instance attributes.
		template_:  template,
		procedure_: procedure,
	}
	return instance
}

// Constant Methods

// Function Methods

// INSTANCE INTERFACE

// Principal Methods

func (v *matchingClause_) GetClass() MatchingClauseClassLike {
	return matchingClauseClass()
}

// Attribute Methods

func (v *matchingClause_) GetTemplate() ExpressionLike {
	return v.template_
}

func (v *matchingClause_) GetProcedure() ProcedureLike {
	return v.procedure_
}

// PROTECTED INTERFACE

// Private Methods

// Instance Structure

type matchingClause_ struct {
	// Declare the instance attributes.
	template_  ExpressionLike
	procedure_ ProcedureLike
}

// Class Structure

type matchingClauseClass_ struct {
	// Declare the class constants.
}

// Class Reference

func matchingClauseClass() *matchingClauseClass_ {
	return matchingClauseClassReference_
}

var matchingClauseClassReference_ = &matchingClauseClass_{
	// Initialize the class constants.
}
