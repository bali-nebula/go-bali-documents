/*
................................................................................
.    Copyright (c) 2009-2026 Crater Dog Technologies™.  All Rights Reserved.   .
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
	com "github.com/craterdog/go-essential-composites/v8"
	pri "github.com/craterdog/go-essential-primitives/v8"
	uti "github.com/craterdog/go-essential-utilities/v8"
)

// CLASS INTERFACE

// Access Function

func ParametersClass() ParametersClassLike {
	return parametersClass()
}

// Constructor Methods

func (c *parametersClass_) Parameters(
	constraints com.CatalogLike[pri.SymbolLike, ComponentLike],
) ParametersLike {
	if uti.IsUndefined(constraints) {
		panic("The \"constraints\" attribute is required by this class.")
	}
	var instance = &parameters_{
		// Initialize the instance attributes.
		constraints_: constraints,
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

func (v *parameters_) GetConstraints() com.CatalogLike[pri.SymbolLike, ComponentLike] {
	return v.constraints_
}

// PROTECTED INTERFACE

// Private Methods

// Instance Structure

type parameters_ struct {
	// Declare the instance attributes.
	constraints_ com.CatalogLike[pri.SymbolLike, ComponentLike]
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
