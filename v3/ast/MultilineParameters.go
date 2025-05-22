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

func MultilineParametersClass() MultilineParametersClassLike {
	return multilineParametersClass()
}

// Constructor Methods

func (c *multilineParametersClass_) MultilineParameters(
	delimiter1 string,
	newline string,
	annotatedAssociations col.Sequential[AnnotatedAssociationLike],
	delimiter2 string,
) MultilineParametersLike {
	if uti.IsUndefined(delimiter1) {
		panic("The \"delimiter1\" attribute is required by this class.")
	}
	if uti.IsUndefined(newline) {
		panic("The \"newline\" attribute is required by this class.")
	}
	if uti.IsUndefined(annotatedAssociations) {
		panic("The \"annotatedAssociations\" attribute is required by this class.")
	}
	if uti.IsUndefined(delimiter2) {
		panic("The \"delimiter2\" attribute is required by this class.")
	}
	var instance = &multilineParameters_{
		// Initialize the instance attributes.
		delimiter1_:            delimiter1,
		newline_:               newline,
		annotatedAssociations_: annotatedAssociations,
		delimiter2_:            delimiter2,
	}
	return instance
}

// INSTANCE INTERFACE

// Principal Methods

func (v *multilineParameters_) GetClass() MultilineParametersClassLike {
	return multilineParametersClass()
}

// Attribute Methods

func (v *multilineParameters_) GetDelimiter1() string {
	return v.delimiter1_
}

func (v *multilineParameters_) GetNewline() string {
	return v.newline_
}

func (v *multilineParameters_) GetAnnotatedAssociations() col.Sequential[AnnotatedAssociationLike] {
	return v.annotatedAssociations_
}

func (v *multilineParameters_) GetDelimiter2() string {
	return v.delimiter2_
}

// PROTECTED INTERFACE

// Instance Structure

type multilineParameters_ struct {
	// Declare the instance attributes.
	delimiter1_            string
	newline_               string
	annotatedAssociations_ col.Sequential[AnnotatedAssociationLike]
	delimiter2_            string
}

// Class Structure

type multilineParametersClass_ struct {
	// Declare the class constants.
}

// Class Reference

func multilineParametersClass() *multilineParametersClass_ {
	return multilineParametersClassReference_
}

var multilineParametersClassReference_ = &multilineParametersClass_{
	// Initialize the class constants.
}
