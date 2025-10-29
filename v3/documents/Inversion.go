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

func InversionClass() InversionClassLike {
	return inversionClass()
}

// Constructor Methods

func (c *inversionClass_) Inversion(
	inverse Inverse,
	numerical any,
) InversionLike {
	if uti.IsUndefined(inverse) {
		panic("The \"inverse\" attribute is required by this class.")
	}
	if uti.IsUndefined(numerical) {
		panic("The \"numerical\" attribute is required by this class.")
	}
	var instance = &inversion_{
		// Initialize the instance attributes.
		inverse_:   inverse,
		numerical_: numerical,
	}
	return instance
}

// Constant Methods

// Function Methods

// INSTANCE INTERFACE

// Principal Methods

func (v *inversion_) GetClass() InversionClassLike {
	return inversionClass()
}

// Attribute Methods

func (v *inversion_) GetInverse() Inverse {
	return v.inverse_
}

func (v *inversion_) GetNumerical() any {
	return v.numerical_
}

// PROTECTED INTERFACE

// Private Methods

// Instance Structure

type inversion_ struct {
	// Declare the instance attributes.
	inverse_   Inverse
	numerical_ any
}

// Class Structure

type inversionClass_ struct {
	// Declare the class constants.
}

// Class Reference

func inversionClass() *inversionClass_ {
	return inversionClassReference_
}

var inversionClassReference_ = &inversionClass_{
	// Initialize the class constants.
}
