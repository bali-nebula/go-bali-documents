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

func InlineAttributesClass() InlineAttributesClassLike {
	return inlineAttributesClass()
}

// Constructor Methods

func (c *inlineAttributesClass_) InlineAttributes(
	delimiter1 string,
	association AssociationLike,
	additionalAssociations col.Sequential[AdditionalAssociationLike],
	delimiter2 string,
) InlineAttributesLike {
	if uti.IsUndefined(delimiter1) {
		panic("The \"delimiter1\" attribute is required by this class.")
	}
	if uti.IsUndefined(association) {
		panic("The \"association\" attribute is required by this class.")
	}
	if uti.IsUndefined(additionalAssociations) {
		panic("The \"additionalAssociations\" attribute is required by this class.")
	}
	if uti.IsUndefined(delimiter2) {
		panic("The \"delimiter2\" attribute is required by this class.")
	}
	var instance = &inlineAttributes_{
		// Initialize the instance attributes.
		delimiter1_:             delimiter1,
		association_:            association,
		additionalAssociations_: additionalAssociations,
		delimiter2_:             delimiter2,
	}
	return instance
}

// INSTANCE INTERFACE

// Principal Methods

func (v *inlineAttributes_) GetClass() InlineAttributesClassLike {
	return inlineAttributesClass()
}

// Attribute Methods

func (v *inlineAttributes_) GetDelimiter1() string {
	return v.delimiter1_
}

func (v *inlineAttributes_) GetAssociation() AssociationLike {
	return v.association_
}

func (v *inlineAttributes_) GetAdditionalAssociations() col.Sequential[AdditionalAssociationLike] {
	return v.additionalAssociations_
}

func (v *inlineAttributes_) GetDelimiter2() string {
	return v.delimiter2_
}

// PROTECTED INTERFACE

// Instance Structure

type inlineAttributes_ struct {
	// Declare the instance attributes.
	delimiter1_             string
	association_            AssociationLike
	additionalAssociations_ col.Sequential[AdditionalAssociationLike]
	delimiter2_             string
}

// Class Structure

type inlineAttributesClass_ struct {
	// Declare the class constants.
}

// Class Reference

func inlineAttributesClass() *inlineAttributesClass_ {
	return inlineAttributesClassReference_
}

var inlineAttributesClassReference_ = &inlineAttributesClass_{
	// Initialize the class constants.
}
