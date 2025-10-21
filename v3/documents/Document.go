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

func DocumentClass() DocumentClassLike {
	return documentClass()
}

// Constructor Methods

func (c *documentClass_) Document(
	optionalAnnotation string,
	component ComponentLike,
) DocumentLike {
	if uti.IsUndefined(component) {
		panic("The \"component\" attribute is required by this class.")
	}
	var instance = &document_{
		// Initialize the instance attributes.
		optionalAnnotation_: optionalAnnotation,
		component_:          component,
	}
	return instance
}

// Constant Methods

// Function Methods

// INSTANCE INTERFACE

// Principal Methods

func (v *document_) GetClass() DocumentClassLike {
	return documentClass()
}

// Attribute Methods

func (v *document_) GetOptionalAnnotation() string {
	return v.optionalAnnotation_
}

func (v *document_) GetComponent() ComponentLike {
	return v.component_
}

// PROTECTED INTERFACE

// Private Methods

// Instance Structure

type document_ struct {
	// Declare the instance attributes.
	optionalAnnotation_ string
	component_          ComponentLike
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
