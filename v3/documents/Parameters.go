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

func ParametersClass() ParametersClassLike {
	return parametersClass()
}

// Constructor Methods

func (c *parametersClass_) Parameters(
	associations fra.CatalogLike[fra.SymbolLike, DocumentLike],
) ParametersLike {
	if uti.IsUndefined(associations) {
		panic("The \"associations\" attribute is required by this class.")
	}
	var instance = &parameters_{
		// Initialize the instance attributes.
		associations_: associations,
	}
	return instance
}

// Constant Methods

// Function Methods

// INSTANCE INTERFACE

// Principal Methods

func (v *parameters_) GetClass() ParametersClassLike {
	return parametersClass()
}

// Attribute Methods

func (v *parameters_) GetAssociations() fra.CatalogLike[fra.SymbolLike, DocumentLike] {
	return v.associations_
}

// PROTECTED INTERFACE

// Private Methods

// Instance Structure

type parameters_ struct {
	// Declare the instance attributes.
	associations_ fra.CatalogLike[fra.SymbolLike, DocumentLike]
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
