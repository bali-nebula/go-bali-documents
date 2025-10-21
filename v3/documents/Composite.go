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

func CompositeClass() CompositeClassLike {
	return compositeClass()
}

// Constructor Methods

func (c *compositeClass_) Composite(
	component ComponentLike,
	optionalNote string,
) CompositeLike {
	if uti.IsUndefined(component) {
		panic("The \"component\" attribute is required by this class.")
	}
	var instance = &composite_{
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

func (v *composite_) GetClass() CompositeClassLike {
	return compositeClass()
}

// Attribute Methods

func (v *composite_) GetComponent() ComponentLike {
	return v.component_
}

func (v *composite_) GetOptionalNote() string {
	return v.optionalNote_
}

// PROTECTED INTERFACE

// Private Methods

// Instance Structure

type composite_ struct {
	// Declare the instance attributes.
	component_    ComponentLike
	optionalNote_ string
}

// Class Structure

type compositeClass_ struct {
	// Declare the class constants.
}

// Class Reference

func compositeClass() *compositeClass_ {
	return compositeClassReference_
}

var compositeClassReference_ = &compositeClass_{
	// Initialize the class constants.
}
