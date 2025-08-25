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

func ItemsClass() ItemsClassLike {
	return itemsClass()
}

// Constructor Methods

func (c *itemsClass_) Items(
	objects fra.ListLike[ObjectLike],
) ItemsLike {
	if uti.IsUndefined(objects) {
		panic("The \"objects\" attribute is required by this class.")
	}
	var instance = &items_{
		// Initialize the instance attributes.
		objects_: objects,
	}
	return instance
}

// Constant Methods

// Function Methods

// INSTANCE INTERFACE

// Principal Methods

func (v *items_) GetClass() ItemsClassLike {
	return itemsClass()
}

// Attribute Methods

func (v *items_) GetObjects() fra.ListLike[ObjectLike] {
	return v.objects_
}

// PROTECTED INTERFACE

// Private Methods

// Instance Structure

type items_ struct {
	// Declare the instance attributes.
	objects_ fra.ListLike[ObjectLike]
}

// Class Structure

type itemsClass_ struct {
	// Declare the class constants.
}

// Class Reference

func itemsClass() *itemsClass_ {
	return itemsClassReference_
}

var itemsClassReference_ = &itemsClass_{
	// Initialize the class constants.
}
