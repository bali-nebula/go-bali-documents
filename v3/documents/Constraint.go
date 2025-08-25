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
	uti "github.com/craterdog/go-missing-utilities/v7"
)

// CLASS INTERFACE

// Access Function

func ConstraintClass() ConstraintClassLike {
	return constraintClass()
}

// Constructor Methods

func (c *constraintClass_) Constraint(
	metadata any,
	optionalParameterization ParameterizationLike,
) ConstraintLike {
	if uti.IsUndefined(metadata) {
		panic("The \"type\" attribute is required by this class.")
	}
	var instance = &constraint_{
		// Initialize the instance attributes.
		metadata:                  metadata,
		optionalParameterization_: optionalParameterization,
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

func (v *constraint_) GetMetadata() any {
	return v.metadata
}

func (v *constraint_) GetOptionalParameterization() ParameterizationLike {
	return v.optionalParameterization_
}

// PROTECTED INTERFACE

// Private Methods

// Instance Structure

type constraint_ struct {
	// Declare the instance attributes.
	metadata                  any
	optionalParameterization_ ParameterizationLike
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
