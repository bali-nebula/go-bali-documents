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

func GenericsClass() GenericsClassLike {
	return genericsClass()
}

// Constructor Methods

func (c *genericsClass_) Generics(
	parameters fra.CatalogLike[fra.SymbolLike, ConstraintLike],
) GenericsLike {
	if uti.IsUndefined(parameters) {
		panic("The \"parameters\" attribute is required by this class.")
	}
	var instance = &generics_{
		// Initialize the instance attributes.
		parameters_: parameters,
	}
	return instance
}

// Constant Methods

// Function Methods

// INSTANCE INTERFACE

// Principal Methods

func (v *generics_) GetClass() GenericsClassLike {
	return genericsClass()
}

// Attribute Methods

func (v *generics_) GetParameters() fra.CatalogLike[fra.SymbolLike, ConstraintLike] {
	return v.parameters_
}

// PROTECTED INTERFACE

// Private Methods

// Instance Structure

type generics_ struct {
	// Declare the instance attributes.
	parameters_ fra.CatalogLike[fra.SymbolLike, ConstraintLike]
}

// Class Structure

type genericsClass_ struct {
	// Declare the class constants.
}

// Class Reference

func genericsClass() *genericsClass_ {
	return genericsClassReference_
}

var genericsClassReference_ = &genericsClass_{
	// Initialize the class constants.
}
