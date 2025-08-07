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

func ParametersClass() ParametersClassLike {
	return parametersClass()
}

// Constructor Methods

func (c *parametersClass_) Parameters(
	delimiter1 string,
	associations fra.ListLike[AssociationLike],
	delimiter2 string,
) ParametersLike {
	if uti.IsUndefined(delimiter1) {
		panic("The \"delimiter1\" attribute is required by this class.")
	}
	if uti.IsUndefined(associations) {
		panic("The \"associations\" attribute is required by this class.")
	}
	if uti.IsUndefined(delimiter2) {
		panic("The \"delimiter2\" attribute is required by this class.")
	}
	var instance = &parameters_{
		// Initialize the instance attributes.
		delimiter1_:   delimiter1,
		associations_: associations,
		delimiter2_:   delimiter2,
	}
	return instance
}

// INSTANCE INTERFACE

// Principal Methods

func (v *parameters_) GetClass() ParametersClassLike {
	return parametersClass()
}

// Attribute Methods

func (v *parameters_) GetDelimiter1() string {
	return v.delimiter1_
}

func (v *parameters_) GetAssociations() fra.ListLike[AssociationLike] {
	return v.associations_
}

func (v *parameters_) GetDelimiter2() string {
	return v.delimiter2_
}

// PROTECTED INTERFACE

// Instance Structure

type parameters_ struct {
	// Declare the instance attributes.
	delimiter1_   string
	associations_ fra.ListLike[AssociationLike]
	delimiter2_   string
}

// Class Structure

type parametersClass_ struct {
	// Declare the class constants.
}

// Class Reference

func parametersClass() *parametersClass_ {
	return parametersClassReference_
}

var parametersClassReference_ = &parametersClass_{
	// Initialize the class constants.
}
