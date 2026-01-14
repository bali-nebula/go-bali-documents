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

func SubcomponentClass() SubcomponentClassLike {
	return subcomponentClass()
}

// Constructor Methods

func (c *subcomponentClass_) Subcomponent(
	identifier pri.IdentifierLike,
	indexes com.Sequential[any],
) SubcomponentLike {
	if uti.IsUndefined(identifier) {
		panic("The \"identifier\" attribute is required by this class.")
	}
	if uti.IsUndefined(indexes) {
		panic("The \"indexes\" attribute is required by this class.")
	}
	var instance = &subcomponent_{
		// Initialize the instance attributes.
		identifier_: identifier,
		indexes_:    indexes,
	}
	return instance
}

// Constant Methods

// Function Methods

// INSTANCE INTERFACE

// Principal Methods

func (v *subcomponent_) GetClass() SubcomponentClassLike {
	return subcomponentClass()
}

// Attribute Methods

func (v *subcomponent_) GetIdentifier() pri.IdentifierLike {
	return v.identifier_
}

func (v *subcomponent_) GetIndexes() com.Sequential[any] {
	return v.indexes_
}

// PROTECTED INTERFACE

// Private Methods

// Instance Structure

type subcomponent_ struct {
	// Declare the instance attributes.
	identifier_ pri.IdentifierLike
	indexes_    com.Sequential[any]
}

// Class Structure

type subcomponentClass_ struct {
	// Declare the class constants.
}

// Class Reference

func subcomponentClass() *subcomponentClass_ {
	return subcomponentClassReference_
}

var subcomponentClassReference_ = &subcomponentClass_{
	// Initialize the class constants.
}
