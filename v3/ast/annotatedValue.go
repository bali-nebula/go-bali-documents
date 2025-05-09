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

func AnnotatedValueClass() AnnotatedValueClassLike {
	return annotatedValueClass()
}

// Constructor Methods

func (c *annotatedValueClass_) AnnotatedValue(
	component ComponentLike,
	optionalNote string,
) AnnotatedValueLike {
	if uti.IsUndefined(component) {
		panic("The \"component\" attribute is required by this class.")
	}
	var instance = &annotatedValue_{
		// Initialize the instance attributes.
		component_:    component,
		optionalNote_: optionalNote,
	}
	return instance
}

// INSTANCE INTERFACE

// Principal Methods

func (v *annotatedValue_) GetClass() AnnotatedValueClassLike {
	return annotatedValueClass()
}

// Attribute Methods

func (v *annotatedValue_) GetComponent() ComponentLike {
	return v.component_
}

func (v *annotatedValue_) GetOptionalNote() string {
	return v.optionalNote_
}

// PROTECTED INTERFACE

// Instance Structure

type annotatedValue_ struct {
	// Declare the instance attributes.
	component_    ComponentLike
	optionalNote_ string
}

// Class Structure

type annotatedValueClass_ struct {
	// Declare the class constants.
}

// Class Reference

func annotatedValueClass() *annotatedValueClass_ {
	return annotatedValueClassReference_
}

var annotatedValueClassReference_ = &annotatedValueClass_{
	// Initialize the class constants.
}
