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

func ObjectClass() ObjectClassLike {
	return objectClass()
}

// Constructor Methods

func (c *objectClass_) Object(
	component ComponentLike,
	optionalNote string,
) ObjectLike {
	if uti.IsUndefined(component) {
		panic("The \"component\" attribute is required by this class.")
	}
	var instance = &object_{
		// Initialize the instance attributes.
		component_:    component,
		optionalNote_: optionalNote,

		// Initialize the inherited aspects.
		Declarative: component,
	}
	return instance
}

// Constant Methods

// Function Methods

// INSTANCE INTERFACE

// Principal Methods

func (v *object_) GetClass() ObjectClassLike {
	return objectClass()
}

// Attribute Methods

func (v *object_) GetComponent() ComponentLike {
	return v.component_
}

func (v *object_) GetOptionalNote() string {
	return v.optionalNote_
}

// PROTECTED INTERFACE

// Private Methods

// Instance Structure

type object_ struct {
	// Declare the instance attributes.
	component_    ComponentLike
	optionalNote_ string

	// Declare the inherited aspects.
	Declarative
}

// Class Structure

type objectClass_ struct {
	// Declare the class constants.
}

// Class Reference

func objectClass() *objectClass_ {
	return objectClassReference_
}

var objectClassReference_ = &objectClass_{
	// Initialize the class constants.
}
