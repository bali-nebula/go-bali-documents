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
	delimiter1 string,
	indexes fra.ListLike[IndexLike],
	delimiter2 string,
) SubcomponentLike {
	if uti.IsUndefined(identifier) {
		panic("The \"identifier\" attribute is required by this class.")
	}
	if uti.IsUndefined(delimiter1) {
		panic("The \"delimiter1\" attribute is required by this class.")
	}
	if uti.IsUndefined(indexes) {
		panic("The \"indexes\" attribute is required by this class.")
	}
	if uti.IsUndefined(delimiter2) {
		panic("The \"delimiter2\" attribute is required by this class.")
	}
	var instance = &subcomponent_{
		// Initialize the instance attributes.
		identifier_: identifier,
		delimiter1_: delimiter1,
		indexes_:    indexes,
		delimiter2_: delimiter2,
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

func (v *subcomponent_) GetDelimiter1() string {
	return v.delimiter1_
}

func (v *subcomponent_) GetIndexes() fra.ListLike[IndexLike] {
	return v.indexes_
}

func (v *subcomponent_) GetDelimiter2() string {
	return v.delimiter2_
}

// PROTECTED INTERFACE

// Instance Structure

type subcomponent_ struct {
	// Declare the instance attributes.
	identifier_ string
	delimiter1_ string
	indexes_    fra.ListLike[IndexLike]
	delimiter2_ string
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
