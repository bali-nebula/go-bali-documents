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
│              This class file was automatically generated using:              │
│            https://github.com/craterdog/go-development-tools/wiki            │
│                                                                              │
│                     Any updates to it may be overwritten.                    │
└──────────────────────────────────────────────────────────────────────────────┘
*/

package ast

import (
	uti "github.com/craterdog/go-missing-utilities/v7"
)

// CLASS INTERFACE

// Access Function

func AssociationClass() AssociationClassLike {
	return associationClass()
}

// Constructor Methods

func (c *associationClass_) Association(
	primitive PrimitiveLike,
	delimiter string,
	document DocumentLike,
) AssociationLike {
	if uti.IsUndefined(primitive) {
		panic("The \"primitive\" attribute is required by this class.")
	}
	if uti.IsUndefined(delimiter) {
		panic("The \"delimiter\" attribute is required by this class.")
	}
	if uti.IsUndefined(document) {
		panic("The \"document\" attribute is required by this class.")
	}
	var instance = &association_{
		// Initialize the instance attributes.
		primitive_: primitive,
		delimiter_: delimiter,
		document_:  document,
	}
	return instance
}

// INSTANCE INTERFACE

// Principal Methods

func (v *association_) GetClass() AssociationClassLike {
	return associationClass()
}

// Attribute Methods

func (v *association_) GetPrimitive() PrimitiveLike {
	return v.primitive_
}

func (v *association_) GetDelimiter() string {
	return v.delimiter_
}

func (v *association_) GetDocument() DocumentLike {
	return v.document_
}

// PROTECTED INTERFACE

// Instance Structure

type association_ struct {
	// Declare the instance attributes.
	primitive_ PrimitiveLike
	delimiter_ string
	document_  DocumentLike
}

// Class Structure

type associationClass_ struct {
	// Declare the class constants.
}

// Class Reference

func associationClass() *associationClass_ {
	return associationClassReference_
}

var associationClassReference_ = &associationClass_{
	// Initialize the class constants.
}
