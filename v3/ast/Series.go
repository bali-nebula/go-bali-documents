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

func SeriesClass() SeriesClassLike {
	return seriesClass()
}

// Constructor Methods

func (c *seriesClass_) Series(
	any_ any,
) SeriesLike {
	if uti.IsUndefined(any_) {
		panic("The \"any\" attribute is required by this class.")
	}
	var instance = &series_{
		// Initialize the instance attributes.
		any_: any_,
	}
	return instance
}

// INSTANCE INTERFACE

// Principal Methods

func (v *series_) GetClass() SeriesClassLike {
	return seriesClass()
}

// Attribute Methods

func (v *series_) GetAny() any {
	return v.any_
}

// PROTECTED INTERFACE

// Instance Structure

type series_ struct {
	// Declare the instance attributes.
	any_ any
}

// Class Structure

type seriesClass_ struct {
	// Declare the class constants.
}

// Class Reference

func seriesClass() *seriesClass_ {
	return seriesClassReference_
}

var seriesClassReference_ = &seriesClass_{
	// Initialize the class constants.
}
