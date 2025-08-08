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
	fra "github.com/craterdog/go-component-framework/v7"
	uti "github.com/craterdog/go-missing-utilities/v7"
)

// CLASS INTERFACE

// Access Function

func WithClauseClass() WithClauseClassLike {
	return withClauseClass()
}

// Constructor Methods

func (c *withClauseClass_) WithClause(
	variable fra.SymbolLike,
	sequence ExpressionLike,
	procedure ProcedureLike,
) WithClauseLike {
	if uti.IsUndefined(variable) {
		panic("The \"variable\" attribute is required by this class.")
	}
	if uti.IsUndefined(sequence) {
		panic("The \"sequence\" attribute is required by this class.")
	}
	if uti.IsUndefined(procedure) {
		panic("The \"procedure\" attribute is required by this class.")
	}
	var instance = &withClause_{
		// Initialize the instance attributes.
		variable_:  variable,
		sequence_:  sequence,
		procedure_: procedure,
	}
	return instance
}

// Constant Methods

// Function Methods

// INSTANCE INTERFACE

// Principal Methods

func (v *withClause_) GetClass() WithClauseClassLike {
	return withClauseClass()
}

// Attribute Methods

func (v *withClause_) GetVariable() fra.SymbolLike {
	return v.variable_
}

func (v *withClause_) GetSequence() ExpressionLike {
	return v.sequence_
}

func (v *withClause_) GetProcedure() ProcedureLike {
	return v.procedure_
}

// PROTECTED INTERFACE

// Private Methods

// Instance Structure

type withClause_ struct {
	// Declare the instance attributes.
	variable_  fra.SymbolLike
	sequence_  ExpressionLike
	procedure_ ProcedureLike
}

// Class Structure

type withClauseClass_ struct {
	// Declare the class constants.
}

// Class Reference

func withClauseClass() *withClauseClass_ {
	return withClauseClassReference_
}

var withClauseClassReference_ = &withClauseClass_{
	// Initialize the class constants.
}
