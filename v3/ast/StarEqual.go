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

func StarEqualClass() StarEqualClassLike {
	return starEqualClass()
}

// Constructor Methods

func (c *starEqualClass_) StarEqual(
	delimiter string,
) StarEqualLike {
	if uti.IsUndefined(delimiter) {
		panic("The \"delimiter\" attribute is required by this class.")
	}
	var instance = &starEqual_{
		// Initialize the instance attributes.
		delimiter_: delimiter,
	}
	return instance
}

// INSTANCE INTERFACE

// Principal Methods

func (v *starEqual_) GetClass() StarEqualClassLike {
	return starEqualClass()
}

// Attribute Methods

func (v *starEqual_) GetDelimiter() string {
	return v.delimiter_
}

// PROTECTED INTERFACE

// Instance Structure

type starEqual_ struct {
	// Declare the instance attributes.
	delimiter_ string
}

// Class Structure

type starEqualClass_ struct {
	// Declare the class constants.
}

// Class Reference

func starEqualClass() *starEqualClass_ {
	return starEqualClassReference_
}

var starEqualClassReference_ = &starEqualClass_{
	// Initialize the class constants.
}
