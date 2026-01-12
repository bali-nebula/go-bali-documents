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

func RefinementClass() RefinementClassLike {
	return refinementClass()
}

// Constructor Methods

func (c *refinementClass_) Refinement(
	modifier Modifier,
	subject any,
) RefinementLike {
	if uti.IsUndefined(modifier) {
		panic("The \"modifier\" attribute is required by this class.")
	}
	if uti.IsUndefined(subject) {
		panic("The \"subject\" attribute is required by this class.")
	}
	var instance = &refinement_{
		// Initialize the instance attributes.
		modifier_: modifier,
		subject_:  subject,
	}
	return instance
}

// Constant Methods

// Function Methods

// INSTANCE INTERFACE

// Principal Methods

func (v *refinement_) GetClass() RefinementClassLike {
	return refinementClass()
}

// Attribute Methods

func (v *refinement_) GetModifier() Modifier {
	return v.modifier_
}

func (v *refinement_) GetSubject() any {
	return v.subject_
}

// PROTECTED INTERFACE

// Private Methods

// Instance Structure

type refinement_ struct {
	// Declare the instance attributes.
	modifier_ Modifier
	subject_  any
}

// Class Structure

type refinementClass_ struct {
	// Declare the class constants.
}

// Class Reference

func refinementClass() *refinementClass_ {
	return refinementClassReference_
}

var refinementClassReference_ = &refinementClass_{
	// Initialize the class constants.
}
