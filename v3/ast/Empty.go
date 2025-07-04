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

package ast

import (
	uti "github.com/craterdog/go-missing-utilities/v7"
)

// CLASS INTERFACE

// Access Function

func EmptyClass() EmptyClassLike {
	return emptyClass()
}

// Constructor Methods

func (c *emptyClass_) Empty(
	delimiter1 string,
	optionalDelimiter string,
	delimiter2 string,
) EmptyLike {
	if uti.IsUndefined(delimiter1) {
		panic("The \"delimiter1\" attribute is required by this class.")
	}
	if uti.IsUndefined(delimiter2) {
		panic("The \"delimiter2\" attribute is required by this class.")
	}
	var instance = &empty_{
		// Initialize the instance attributes.
		delimiter1_:        delimiter1,
		optionalDelimiter_: optionalDelimiter,
		delimiter2_:        delimiter2,
	}
	return instance
}

// INSTANCE INTERFACE

// Principal Methods

func (v *empty_) GetClass() EmptyClassLike {
	return emptyClass()
}

// Attribute Methods

func (v *empty_) GetDelimiter1() string {
	return v.delimiter1_
}

func (v *empty_) GetOptionalDelimiter() string {
	return v.optionalDelimiter_
}

func (v *empty_) GetDelimiter2() string {
	return v.delimiter2_
}

// PROTECTED INTERFACE

// Instance Structure

type empty_ struct {
	// Declare the instance attributes.
	delimiter1_        string
	optionalDelimiter_ string
	delimiter2_        string
}

// Class Structure

type emptyClass_ struct {
	// Declare the class constants.
}

// Class Reference

func emptyClass() *emptyClass_ {
	return emptyClassReference_
}

var emptyClassReference_ = &emptyClass_{
	// Initialize the class constants.
}
