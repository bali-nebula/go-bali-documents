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

func HeaderClass() HeaderClassLike {
	return headerClass()
}

// Constructor Methods

func (c *headerClass_) Header(
	comment string,
) HeaderLike {
	if uti.IsUndefined(comment) {
		panic("The \"comment\" attribute is required by this class.")
	}
	var instance = &header_{
		// Initialize the instance attributes.
		comment_: comment,
	}
	return instance
}

// INSTANCE INTERFACE

// Principal Methods

func (v *header_) GetClass() HeaderClassLike {
	return headerClass()
}

// Attribute Methods

func (v *header_) GetComment() string {
	return v.comment_
}

// PROTECTED INTERFACE

// Instance Structure

type header_ struct {
	// Declare the instance attributes.
	comment_ string
}

// Class Structure

type headerClass_ struct {
	// Declare the class constants.
}

// Class Reference

func headerClass() *headerClass_ {
	return headerClassReference_
}

var headerClassReference_ = &headerClass_{
	// Initialize the class constants.
}
