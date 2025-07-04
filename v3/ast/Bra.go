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

package ast

import (
	uti "github.com/craterdog/go-missing-utilities/v7"
)

// CLASS INTERFACE

// Access Function

func BraClass() BraClassLike {
	return braClass()
}

// Constructor Methods

func (c *braClass_) Bra(
	any_ any,
) BraLike {
	if uti.IsUndefined(any_) {
		panic("The \"any\" attribute is required by this class.")
	}
	var instance = &bra_{
		// Initialize the instance attributes.
		any_: any_,
	}
	return instance
}

// INSTANCE INTERFACE

// Principal Methods

func (v *bra_) GetClass() BraClassLike {
	return braClass()
}

// Attribute Methods

func (v *bra_) GetAny() any {
	return v.any_
}

// PROTECTED INTERFACE

// Instance Structure

type bra_ struct {
	// Declare the instance attributes.
	any_ any
}

// Class Structure

type braClass_ struct {
	// Declare the class constants.
}

// Class Reference

func braClass() *braClass_ {
	return braClassReference_
}

var braClassReference_ = &braClass_{
	// Initialize the class constants.
}
