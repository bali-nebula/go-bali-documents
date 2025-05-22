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

func ExclusiveClass() ExclusiveClassLike {
	return exclusiveClass()
}

// Constructor Methods

func (c *exclusiveClass_) Exclusive(
	delimiter string,
) ExclusiveLike {
	if uti.IsUndefined(delimiter) {
		panic("The \"delimiter\" attribute is required by this class.")
	}
	var instance = &exclusive_{
		// Initialize the instance attributes.
		delimiter_: delimiter,
	}
	return instance
}

// INSTANCE INTERFACE

// Principal Methods

func (v *exclusive_) GetClass() ExclusiveClassLike {
	return exclusiveClass()
}

// Attribute Methods

func (v *exclusive_) GetDelimiter() string {
	return v.delimiter_
}

// PROTECTED INTERFACE

// Instance Structure

type exclusive_ struct {
	// Declare the instance attributes.
	delimiter_ string
}

// Class Structure

type exclusiveClass_ struct {
	// Declare the class constants.
}

// Class Reference

func exclusiveClass() *exclusiveClass_ {
	return exclusiveClassReference_
}

var exclusiveClassReference_ = &exclusiveClass_{
	// Initialize the class constants.
}
