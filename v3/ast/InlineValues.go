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
	col "github.com/craterdog/go-collection-framework/v7"
	uti "github.com/craterdog/go-missing-utilities/v7"
)

// CLASS INTERFACE

// Access Function

func InlineValuesClass() InlineValuesClassLike {
	return inlineValuesClass()
}

// Constructor Methods

func (c *inlineValuesClass_) InlineValues(
	delimiter1 string,
	component ComponentLike,
	additionalValues col.Sequential[AdditionalValueLike],
	delimiter2 string,
) InlineValuesLike {
	if uti.IsUndefined(delimiter1) {
		panic("The \"delimiter1\" attribute is required by this class.")
	}
	if uti.IsUndefined(component) {
		panic("The \"component\" attribute is required by this class.")
	}
	if uti.IsUndefined(additionalValues) {
		panic("The \"additionalValues\" attribute is required by this class.")
	}
	if uti.IsUndefined(delimiter2) {
		panic("The \"delimiter2\" attribute is required by this class.")
	}
	var instance = &inlineValues_{
		// Initialize the instance attributes.
		delimiter1_:       delimiter1,
		component_:        component,
		additionalValues_: additionalValues,
		delimiter2_:       delimiter2,
	}
	return instance
}

// INSTANCE INTERFACE

// Principal Methods

func (v *inlineValues_) GetClass() InlineValuesClassLike {
	return inlineValuesClass()
}

// Attribute Methods

func (v *inlineValues_) GetDelimiter1() string {
	return v.delimiter1_
}

func (v *inlineValues_) GetComponent() ComponentLike {
	return v.component_
}

func (v *inlineValues_) GetAdditionalValues() col.Sequential[AdditionalValueLike] {
	return v.additionalValues_
}

func (v *inlineValues_) GetDelimiter2() string {
	return v.delimiter2_
}

// PROTECTED INTERFACE

// Instance Structure

type inlineValues_ struct {
	// Declare the instance attributes.
	delimiter1_       string
	component_        ComponentLike
	additionalValues_ col.Sequential[AdditionalValueLike]
	delimiter2_       string
}

// Class Structure

type inlineValuesClass_ struct {
	// Declare the class constants.
}

// Class Reference

func inlineValuesClass() *inlineValuesClass_ {
	return inlineValuesClassReference_
}

var inlineValuesClassReference_ = &inlineValuesClass_{
	// Initialize the class constants.
}
