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

func ConstraintClass() ConstraintClassLike {
	return constraintClass()
}

// Constructor Methods

func (c *constraintClass_) Constraint(
	entity any,
	optionalGenerics GenericsLike,
) ConstraintLike {
	if uti.IsUndefined(entity) {
		panic("The \"entity\" attribute is required by this class.")
	}
	var instance = &constraint_{
		// Initialize the instance attributes.
		entity_:           entity,
		optionalGenerics_: optionalGenerics,
	}
	return instance
}

// Constant Methods

// Function Methods

// INSTANCE INTERFACE

// Principal Methods

func (v *constraint_) GetClass() ConstraintClassLike {
	return constraintClass()
}

// Attribute Methods

func (v *constraint_) GetEntity() any {
	return v.entity_
}

func (v *constraint_) GetOptionalGenerics() GenericsLike {
	return v.optionalGenerics_
}

// PROTECTED INTERFACE

// Private Methods

// Instance Structure

type constraint_ struct {
	// Declare the instance attributes.
	entity_           any
	optionalGenerics_ GenericsLike
}

// Class Structure

type constraintClass_ struct {
	// Declare the class constants.
}

// Class Reference

func constraintClass() *constraintClass_ {
	return constraintClassReference_
}

var constraintClassReference_ = &constraintClass_{
	// Initialize the class constants.
}
