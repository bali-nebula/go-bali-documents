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

func ActionInductionClass() ActionInductionClassLike {
	return actionInductionClass()
}

// Constructor Methods

func (c *actionInductionClass_) ActionInduction(
	any_ any,
) ActionInductionLike {
	if uti.IsUndefined(any_) {
		panic("The \"any\" attribute is required by this class.")
	}
	var instance = &actionInduction_{
		// Initialize the instance attributes.
		any_: any_,
	}
	return instance
}

// INSTANCE INTERFACE

// Principal Methods

func (v *actionInduction_) GetClass() ActionInductionClassLike {
	return actionInductionClass()
}

// Attribute Methods

func (v *actionInduction_) GetAny() any {
	return v.any_
}

// PROTECTED INTERFACE

// Instance Structure

type actionInduction_ struct {
	// Declare the instance attributes.
	any_ any
}

// Class Structure

type actionInductionClass_ struct {
	// Declare the class constants.
}

// Class Reference

func actionInductionClass() *actionInductionClass_ {
	return actionInductionClassReference_
}

var actionInductionClassReference_ = &actionInductionClass_{
	// Initialize the class constants.
}
