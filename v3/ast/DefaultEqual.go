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

func DefaultEqualClass() DefaultEqualClassLike {
	return defaultEqualClass()
}

// Constructor Methods

func (c *defaultEqualClass_) DefaultEqual(
	delimiter string,
) DefaultEqualLike {
	if uti.IsUndefined(delimiter) {
		panic("The \"delimiter\" attribute is required by this class.")
	}
	var instance = &defaultEqual_{
		// Initialize the instance attributes.
		delimiter_: delimiter,
	}
	return instance
}

// INSTANCE INTERFACE

// Principal Methods

func (v *defaultEqual_) GetClass() DefaultEqualClassLike {
	return defaultEqualClass()
}

// Attribute Methods

func (v *defaultEqual_) GetDelimiter() string {
	return v.delimiter_
}

// PROTECTED INTERFACE

// Instance Structure

type defaultEqual_ struct {
	// Declare the instance attributes.
	delimiter_ string
}

// Class Structure

type defaultEqualClass_ struct {
	// Declare the class constants.
}

// Class Reference

func defaultEqualClass() *defaultEqualClass_ {
	return defaultEqualClassReference_
}

var defaultEqualClassReference_ = &defaultEqualClass_{
	// Initialize the class constants.
}
