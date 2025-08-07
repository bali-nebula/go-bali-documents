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

func SubcomponentClass() SubcomponentClassLike {
	return subcomponentClass()
}

// Constructor Methods

func (c *subcomponentClass_) Subcomponent(
	identifier string,
	indexes fra.ListLike[any],
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

// INSTANCE INTERFACE

// Principal Methods

func (v *subcomponent_) GetClass() SubcomponentClassLike {
	return subcomponentClass()
}

// Attribute Methods

func (v *subcomponent_) GetIdentifier() string {
	return v.identifier_
}

func (v *subcomponent_) GetIndexes() fra.ListLike[any] {
	return v.indexes_
}

// PROTECTED INTERFACE

// Instance Structure

type subcomponent_ struct {
	// Declare the instance attributes.
	identifier_ string
	indexes_    fra.ListLike[any]
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
