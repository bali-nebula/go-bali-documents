/*
................................................................................
.    Copyright (c) 2009-2026 Crater Dog Technologies™.  All Rights Reserved.   .
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

func ComplementClass() ComplementClassLike {
	return complementClass()
}

// Constructor Methods

func (c *complementClass_) Complement(
	reversible any,
) ComplementLike {
	if uti.IsUndefined(reversible) {
		panic("The \"reversible\" attribute is required by this class.")
	}
	var instance = &complement_{
		// Initialize the instance attributes.
		reversible_: reversible,
	}
	return instance
}

// Constant Methods

// Function Methods

// INSTANCE INTERFACE

// Principal Methods

func (v *complement_) GetClass() ComplementClassLike {
	return complementClass()
}

// Attribute Methods

func (v *complement_) GetReversible() any {
	return v.reversible_
}

// PROTECTED INTERFACE

// Private Methods

// Instance Structure

type complement_ struct {
	// Declare the instance attributes.
	reversible_ any
}

// Class Structure

type complementClass_ struct {
	// Declare the class constants.
}

// Class Reference

func complementClass() *complementClass_ {
	return complementClassReference_
}

var complementClassReference_ = &complementClass_{
	// Initialize the class constants.
}
