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
│                 This class file was automatically generated.                 │
│                     Any updates to it may be overwritten.                    │
└──────────────────────────────────────────────────────────────────────────────┘
*/

package ast

import (
	uti "github.com/craterdog/go-missing-utilities/v7"
)

// CLASS INTERFACE

// Access Function

func NoStatementsClass() NoStatementsClassLike {
	return noStatementsClass()
}

// Constructor Methods

func (c *noStatementsClass_) NoStatements(
	delimiter1 string,
	delimiter2 string,
) NoStatementsLike {
	if uti.IsUndefined(delimiter1) {
		panic("The \"delimiter1\" attribute is required by this class.")
	}
	if uti.IsUndefined(delimiter2) {
		panic("The \"delimiter2\" attribute is required by this class.")
	}
	var instance = &noStatements_{
		// Initialize the instance attributes.
		delimiter1_: delimiter1,
		delimiter2_: delimiter2,
	}
	return instance
}

// INSTANCE INTERFACE

// Principal Methods

func (v *noStatements_) GetClass() NoStatementsClassLike {
	return noStatementsClass()
}

// Attribute Methods

func (v *noStatements_) GetDelimiter1() string {
	return v.delimiter1_
}

func (v *noStatements_) GetDelimiter2() string {
	return v.delimiter2_
}

// PROTECTED INTERFACE

// Instance Structure

type noStatements_ struct {
	// Declare the instance attributes.
	delimiter1_ string
	delimiter2_ string
}

// Class Structure

type noStatementsClass_ struct {
	// Declare the class constants.
}

// Class Reference

func noStatementsClass() *noStatementsClass_ {
	return noStatementsClassReference_
}

var noStatementsClassReference_ = &noStatementsClass_{
	// Initialize the class constants.
}
