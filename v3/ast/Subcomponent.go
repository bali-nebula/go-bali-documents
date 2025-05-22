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

func SubcomponentClass() SubcomponentClassLike {
	return subcomponentClass()
}

// Constructor Methods

func (c *subcomponentClass_) Subcomponent(
	identifier string,
	delimiter1 string,
	indices IndicesLike,
	delimiter2 string,
) SubcomponentLike {
	if uti.IsUndefined(identifier) {
		panic("The \"identifier\" attribute is required by this class.")
	}
	if uti.IsUndefined(delimiter1) {
		panic("The \"delimiter1\" attribute is required by this class.")
	}
	if uti.IsUndefined(indices) {
		panic("The \"indices\" attribute is required by this class.")
	}
	if uti.IsUndefined(delimiter2) {
		panic("The \"delimiter2\" attribute is required by this class.")
	}
	var instance = &subcomponent_{
		// Initialize the instance attributes.
		identifier_: identifier,
		delimiter1_: delimiter1,
		indices_:    indices,
		delimiter2_: delimiter2,
	}
	return instance
}

// INSTANCE INTERFACE

// Principal Methods

func (v *subcomponent_) GetClass() SubcomponentClassLike {
	return subcomponentClass()
}

// Attribute Methods

func (v *subcomponent_) GetIdentifier() string {
	return v.identifier_
}

func (v *subcomponent_) GetDelimiter1() string {
	return v.delimiter1_
}

func (v *subcomponent_) GetIndices() IndicesLike {
	return v.indices_
}

func (v *subcomponent_) GetDelimiter2() string {
	return v.delimiter2_
}

// PROTECTED INTERFACE

// Instance Structure

type subcomponent_ struct {
	// Declare the instance attributes.
	identifier_ string
	delimiter1_ string
	indices_    IndicesLike
	delimiter2_ string
}

// Class Structure

type subcomponentClass_ struct {
	// Declare the class constants.
}

// Class Reference

func subcomponentClass() *subcomponentClass_ {
	return subcomponentClassReference_
}

var subcomponentClassReference_ = &subcomponentClass_{
	// Initialize the class constants.
}
