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

func ContentClass() ContentClassLike {
	return contentClass()
}

// Constructor Methods

func (c *contentClass_) Content(
	component Compound,
	optionalNote string,
) ContentLike {
	if uti.IsUndefined(component) {
		panic("The \"component\" attribute is required by this class.")
	}
	var instance = &content_{
		// Initialize the instance attributes.
		component_:    component,
		optionalNote_: optionalNote,
	}
	return instance
}

// Constant Methods

// Function Methods

// INSTANCE INTERFACE

// Principal Methods

func (v *content_) GetClass() ContentClassLike {
	return contentClass()
}

// Attribute Methods

func (v *content_) GetComponent() Compound {
	return v.component_
}

func (v *content_) GetOptionalNote() string {
	return v.optionalNote_
}

// PROTECTED INTERFACE

// Private Methods

// Instance Structure

type content_ struct {
	// Declare the instance attributes.
	component_    Compound
	optionalNote_ string
}

// Class Structure

type contentClass_ struct {
	// Declare the class constants.
}

// Class Reference

func contentClass() *contentClass_ {
	return contentClassReference_
}

var contentClassReference_ = &contentClass_{
	// Initialize the class constants.
}
