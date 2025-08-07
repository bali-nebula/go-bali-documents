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

func WithClauseClass() WithClauseClassLike {
	return withClauseClass()
}

// Constructor Methods

func (c *withClauseClass_) WithClause(
	delimiter1 string,
	delimiter2 string,
	variable VariableLike,
	delimiter3 string,
	sequence SequenceLike,
	delimiter4 string,
	procedure ProcedureLike,
) WithClauseLike {
	if uti.IsUndefined(delimiter1) {
		panic("The \"delimiter1\" attribute is required by this class.")
	}
	if uti.IsUndefined(delimiter2) {
		panic("The \"delimiter2\" attribute is required by this class.")
	}
	if uti.IsUndefined(variable) {
		panic("The \"variable\" attribute is required by this class.")
	}
	if uti.IsUndefined(delimiter3) {
		panic("The \"delimiter3\" attribute is required by this class.")
	}
	if uti.IsUndefined(sequence) {
		panic("The \"sequence\" attribute is required by this class.")
	}
	if uti.IsUndefined(delimiter4) {
		panic("The \"delimiter4\" attribute is required by this class.")
	}
	if uti.IsUndefined(procedure) {
		panic("The \"procedure\" attribute is required by this class.")
	}
	var instance = &withClause_{
		// Initialize the instance attributes.
		delimiter1_: delimiter1,
		delimiter2_: delimiter2,
		variable_:   variable,
		delimiter3_: delimiter3,
		sequence_:   sequence,
		delimiter4_: delimiter4,
		procedure_:  procedure,
	}
	return instance
}

// INSTANCE INTERFACE

// Principal Methods

func (v *withClause_) GetClass() WithClauseClassLike {
	return withClauseClass()
}

// Attribute Methods

func (v *withClause_) GetDelimiter1() string {
	return v.delimiter1_
}

func (v *withClause_) GetDelimiter2() string {
	return v.delimiter2_
}

func (v *withClause_) GetVariable() VariableLike {
	return v.variable_
}

func (v *withClause_) GetDelimiter3() string {
	return v.delimiter3_
}

func (v *withClause_) GetSequence() SequenceLike {
	return v.sequence_
}

func (v *withClause_) GetDelimiter4() string {
	return v.delimiter4_
}

func (v *withClause_) GetProcedure() ProcedureLike {
	return v.procedure_
}

// PROTECTED INTERFACE

// Instance Structure

type withClause_ struct {
	// Declare the instance attributes.
	delimiter1_ string
	delimiter2_ string
	variable_   VariableLike
	delimiter3_ string
	sequence_   SequenceLike
	delimiter4_ string
	procedure_  ProcedureLike
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
