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

func EntitiesClass() EntitiesClassLike {
	return entitiesClass()
}

// Constructor Methods

func (c *entitiesClass_) Entities(
	items fra.ListLike[DocumentLike],
) EntitiesLike {
	if uti.IsUndefined(items) {
		panic("The \"items\" attribute is required by this class.")
	}
	var instance = &entities_{
		// Initialize the instance attributes.
		items_: items,
	}
	return instance
}

// Constant Methods

// Function Methods

// INSTANCE INTERFACE

// Principal Methods

func (v *entities_) GetClass() EntitiesClassLike {
	return entitiesClass()
}

// Attribute Methods

func (v *entities_) GetItems() fra.ListLike[DocumentLike] {
	return v.items_
}

// PROTECTED INTERFACE

// Private Methods

// Instance Structure

type entities_ struct {
	// Declare the instance attributes.
	items_ fra.ListLike[DocumentLike]
}

// Class Structure

type entitiesClass_ struct {
	// Declare the class constants.
}

// Class Reference

func entitiesClass() *entitiesClass_ {
	return entitiesClassReference_
}

var entitiesClassReference_ = &entitiesClass_{
	// Initialize the class constants.
}
