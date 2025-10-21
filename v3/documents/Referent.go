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
	uti "github.com/craterdog/go-missing-utilities/v8"
)

// CLASS INTERFACE

// Access Function

func ReferentClass() ReferentClassLike {
	return referentClass()
}

// Constructor Methods

func (c *referentClass_) Referent(
	reference any,
) ReferentLike {
	if uti.IsUndefined(reference) {
		panic("The \"reference\" attribute is required by this class.")
	}
	var instance = &referent_{
		// Initialize the instance attributes.
		reference_: reference,
	}
	return instance
}

// Constant Methods

// Function Methods

// INSTANCE INTERFACE

// Principal Methods

func (v *referent_) GetClass() ReferentClassLike {
	return referentClass()
}

// Attribute Methods

func (v *referent_) GetReference() any {
	return v.reference_
}

// PROTECTED INTERFACE

// Private Methods

// Instance Structure

type referent_ struct {
	// Declare the instance attributes.
	reference_ any
}

// Class Structure

type referentClass_ struct {
	// Declare the class constants.
}

// Class Reference

func referentClass() *referentClass_ {
	return referentClassReference_
}

var referentClassReference_ = &referentClass_{
	// Initialize the class constants.
}
