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

func AdditionalAssociationClass() AdditionalAssociationClassLike {
	return additionalAssociationClass()
}

// Constructor Methods

func (c *additionalAssociationClass_) AdditionalAssociation(
	delimiter string,
	association AssociationLike,
) AdditionalAssociationLike {
	if uti.IsUndefined(delimiter) {
		panic("The \"delimiter\" attribute is required by this class.")
	}
	if uti.IsUndefined(association) {
		panic("The \"association\" attribute is required by this class.")
	}
	var instance = &additionalAssociation_{
		// Initialize the instance attributes.
		delimiter_:   delimiter,
		association_: association,
	}
	return instance
}

// INSTANCE INTERFACE

// Principal Methods

func (v *additionalAssociation_) GetClass() AdditionalAssociationClassLike {
	return additionalAssociationClass()
}

// Attribute Methods

func (v *additionalAssociation_) GetDelimiter() string {
	return v.delimiter_
}

func (v *additionalAssociation_) GetAssociation() AssociationLike {
	return v.association_
}

// PROTECTED INTERFACE

// Instance Structure

type additionalAssociation_ struct {
	// Declare the instance attributes.
	delimiter_   string
	association_ AssociationLike
}

// Class Structure

type additionalAssociationClass_ struct {
	// Declare the class constants.
}

// Class Reference

func additionalAssociationClass() *additionalAssociationClass_ {
	return additionalAssociationClassReference_
}

var additionalAssociationClassReference_ = &additionalAssociationClass_{
	// Initialize the class constants.
}
