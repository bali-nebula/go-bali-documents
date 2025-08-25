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
	associations fra.CatalogLike[any, ObjectLike],
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

// Constant Methods

// Function Methods

// INSTANCE INTERFACE

// Principal Methods

func (v *attributes_) GetClass() AttributesClassLike {
	return attributesClass()
}

// Attribute Methods

func (v *attributes_) GetAssociations() fra.CatalogLike[any, ObjectLike] {
	return v.associations_
}

// PROTECTED INTERFACE

// Private Methods

// Instance Structure

type attributes_ struct {
	// Declare the instance attributes.
	associations_ fra.CatalogLike[any, ObjectLike]
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
