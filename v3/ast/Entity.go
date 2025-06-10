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

func EntityClass() EntityClassLike {
	return entityClass()
}

// Constructor Methods

func (c *entityClass_) Entity(
	component ComponentLike,
	optionalParameters ParametersLike,
	optionalNote string,
) EntityLike {
	if uti.IsUndefined(component) {
		panic("The \"component\" attribute is required by this class.")
	}
	var instance = &entity_{
		// Initialize the instance attributes.
		component_:          component,
		optionalParameters_: optionalParameters,
		optionalNote_:       optionalNote,
	}
	return instance
}

// INSTANCE INTERFACE

// Principal Methods

func (v *entity_) GetClass() EntityClassLike {
	return entityClass()
}

// Attribute Methods

func (v *entity_) GetComponent() ComponentLike {
	return v.component_
}

func (v *entity_) GetOptionalParameters() ParametersLike {
	return v.optionalParameters_
}

func (v *entity_) GetOptionalNote() string {
	return v.optionalNote_
}

// PROTECTED INTERFACE

// Instance Structure

type entity_ struct {
	// Declare the instance attributes.
	component_          ComponentLike
	optionalParameters_ ParametersLike
	optionalNote_       string
}

// Class Structure

type entityClass_ struct {
	// Declare the class constants.
}

// Class Reference

func entityClass() *entityClass_ {
	return entityClassReference_
}

var entityClassReference_ = &entityClass_{
	// Initialize the class constants.
}
