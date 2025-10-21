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
	ele "github.com/bali-nebula/go-bali-documents/v3/elements"
	uti "github.com/craterdog/go-missing-utilities/v8"
)

// CLASS INTERFACE

// Access Function

func WithClauseClass() WithClauseClassLike {
	return withClauseClass()
}

// Constructor Methods

func (c *withClauseClass_) WithClause(
	symbol ele.SymbolLike,
	sequence ExpressionLike,
	procedure ProcedureLike,
) WithClauseLike {
	if uti.IsUndefined(symbol) {
		panic("The \"symbol\" attribute is required by this class.")
	}
	if uti.IsUndefined(sequence) {
		panic("The \"sequence\" attribute is required by this class.")
	}
	if uti.IsUndefined(procedure) {
		panic("The \"procedure\" attribute is required by this class.")
	}
	var instance = &withClause_{
		// Initialize the instance attributes.
		symbol_:    symbol,
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

func (v *withClause_) GetSymbol() ele.SymbolLike {
	return v.symbol_
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
	symbol_    ele.SymbolLike
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
