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

func ContentClass() ContentClassLike {
	return contentClass()
}

// Constructor Methods

func (c *contentClass_) Content(
	composite Composite,
	optionalNote string,
) ContentLike {
	if uti.IsUndefined(composite) {
		panic("The \"composite\" attribute is required by this class.")
	}
	var instance = &content_{
		// Initialize the instance attributes.
		composite_:    composite,
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

func (v *content_) GetComposite() Composite {
	return v.composite_
}

func (v *content_) GetOptionalNote() string {
	return v.optionalNote_
}

// PROTECTED INTERFACE

// Private Methods

// Instance Structure

type content_ struct {
	// Declare the instance attributes.
	composite_    Composite
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
