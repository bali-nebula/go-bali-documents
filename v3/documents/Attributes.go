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

package documents

import (
	fra "github.com/craterdog/go-component-framework/v7"
	uti "github.com/craterdog/go-missing-utilities/v7"
)

// CLASS INTERFACE

// Access Function

func AttributesClass() AttributesClassLike {
	return attributesClass()
}

// Constructor Methods

func (c *attributesClass_) Attributes(
	associations fra.CatalogLike[any, DocumentLike],
) AttributesLike {
	if uti.IsUndefined(associations) {
		panic("The \"associations\" attribute is required by this class.")
	}
	var instance = &attributes_{
		// Initialize the instance attributes.
		associations_: associations,
	}
	return instance
}

// INSTANCE INTERFACE

// Principal Methods

func (v *attributes_) GetClass() AttributesClassLike {
	return attributesClass()
}

// Attribute Methods

func (v *attributes_) GetAssociations() fra.CatalogLike[any, DocumentLike] {
	return v.associations_
}

// PROTECTED INTERFACE

// Instance Structure

type attributes_ struct {
	// Declare the instance attributes.
	associations_ fra.CatalogLike[any, DocumentLike]
}

// Class Structure

type attributesClass_ struct {
	// Declare the class constants.
}

// Class Reference

func attributesClass() *attributesClass_ {
	return attributesClassReference_
}

var attributesClassReference_ = &attributesClass_{
	// Initialize the class constants.
}
