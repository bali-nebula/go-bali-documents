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

func DocumentClass() DocumentClassLike {
	return documentClass()
}

// Constructor Methods

func (c *documentClass_) Document(
	component any,
	optionalParameters ParametersLike,
	optionalNote string,
) DocumentLike {
	if uti.IsUndefined(component) {
		panic("The \"component\" attribute is required by this class.")
	}
	var instance = &document_{
		// Initialize the instance attributes.
		component_:          component,
		optionalParameters_: optionalParameters,
		optionalNote_:       optionalNote,
	}
	return instance
}

// INSTANCE INTERFACE

// Principal Methods

func (v *document_) GetClass() DocumentClassLike {
	return documentClass()
}

// Attribute Methods

func (v *document_) GetComponent() any {
	return v.component_
}

func (v *document_) GetOptionalParameters() ParametersLike {
	return v.optionalParameters_
}

func (v *document_) GetOptionalNote() string {
	return v.optionalNote_
}

// PROTECTED INTERFACE

// Instance Structure

type document_ struct {
	// Declare the instance attributes.
	component_          any
	optionalParameters_ ParametersLike
	optionalNote_       string
}

// Class Structure

type documentClass_ struct {
	// Declare the class constants.
}

// Class Reference

func documentClass() *documentClass_ {
	return documentClassReference_
}

var documentClassReference_ = &documentClass_{
	// Initialize the class constants.
}
