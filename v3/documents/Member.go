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

func MemberClass() MemberClassLike {
	return memberClass()
}

// Constructor Methods

func (c *memberClass_) Member(
	component ComponentLike,
	optionalNote string,
) MemberLike {
	if uti.IsUndefined(component) {
		panic("The \"component\" attribute is required by this class.")
	}
	var instance = &member_{
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

func (v *member_) GetClass() MemberClassLike {
	return memberClass()
}

// Attribute Methods

func (v *member_) GetComponent() ComponentLike {
	return v.component_
}

func (v *member_) GetOptionalNote() string {
	return v.optionalNote_
}

// PROTECTED INTERFACE

// Private Methods

// Instance Structure

type member_ struct {
	// Declare the instance attributes.
	component_    ComponentLike
	optionalNote_ string

	// Declare the inherited aspects.
	Declarative
}

// Class Structure

type memberClass_ struct {
	// Declare the class constants.
}

// Class Reference

func memberClass() *memberClass_ {
	return memberClassReference_
}

var memberClassReference_ = &memberClass_{
	// Initialize the class constants.
}
