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

func MultilineValuesClass() MultilineValuesClassLike {
	return multilineValuesClass()
}

// Constructor Methods

func (c *multilineValuesClass_) MultilineValues(
	delimiter1 string,
	newline string,
	annotatedValues col.Sequential[AnnotatedValueLike],
	delimiter2 string,
) MultilineValuesLike {
	if uti.IsUndefined(delimiter1) {
		panic("The \"delimiter1\" attribute is required by this class.")
	}
	if uti.IsUndefined(newline) {
		panic("The \"newline\" attribute is required by this class.")
	}
	if uti.IsUndefined(annotatedValues) {
		panic("The \"annotatedValues\" attribute is required by this class.")
	}
	if uti.IsUndefined(delimiter2) {
		panic("The \"delimiter2\" attribute is required by this class.")
	}
	var instance = &multilineValues_{
		// Initialize the instance attributes.
		delimiter1_:      delimiter1,
		newline_:         newline,
		annotatedValues_: annotatedValues,
		delimiter2_:      delimiter2,
	}
	return instance
}

// INSTANCE INTERFACE

// Principal Methods

func (v *multilineValues_) GetClass() MultilineValuesClassLike {
	return multilineValuesClass()
}

// Attribute Methods

func (v *multilineValues_) GetDelimiter1() string {
	return v.delimiter1_
}

func (v *multilineValues_) GetNewline() string {
	return v.newline_
}

func (v *multilineValues_) GetAnnotatedValues() col.Sequential[AnnotatedValueLike] {
	return v.annotatedValues_
}

func (v *multilineValues_) GetDelimiter2() string {
	return v.delimiter2_
}

// PROTECTED INTERFACE

// Instance Structure

type multilineValues_ struct {
	// Declare the instance attributes.
	delimiter1_      string
	newline_         string
	annotatedValues_ col.Sequential[AnnotatedValueLike]
	delimiter2_      string
}

// Class Structure

type multilineValuesClass_ struct {
	// Declare the class constants.
}

// Class Reference

func multilineValuesClass() *multilineValuesClass_ {
	return multilineValuesClassReference_
}

var multilineValuesClassReference_ = &multilineValuesClass_{
	// Initialize the class constants.
}
