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
	uti "github.com/craterdog/go-missing-utilities/v7"
)

// CLASS INTERFACE

// Access Function

func RangeClass() RangeClassLike {
	return rangeClass()
}

// Constructor Methods

func (c *rangeClass_) Range(
	bra Extent,
	first any,
	last any,
	ket Extent,
) RangeLike {
	if uti.IsUndefined(bra) {
		panic("The \"bra\" attribute is required by this class.")
	}
	if uti.IsUndefined(first) {
		panic("The \"first\" attribute is required by this class.")
	}
	if uti.IsUndefined(last) {
		panic("The \"last\" attribute is required by this class.")
	}
	if uti.IsUndefined(ket) {
		panic("The \"ket\" attribute is required by this class.")
	}
	var instance = &range_{
		// Initialize the instance attributes.
		bra_:   bra,
		first_: first,
		last_:  last,
		ket_:   ket,
	}
	return instance
}

// INSTANCE INTERFACE

// Principal Methods

func (v *range_) GetClass() RangeClassLike {
	return rangeClass()
}

// Attribute Methods

func (v *range_) GetBra() Extent {
	return v.bra_
}

func (v *range_) GetFirst() any {
	return v.first_
}

func (v *range_) GetLast() any {
	return v.last_
}

func (v *range_) GetKet() Extent {
	return v.ket_
}

// PROTECTED INTERFACE

// Instance Structure

type range_ struct {
	// Declare the instance attributes.
	bra_   Extent
	first_ any
	last_  any
	ket_   Extent
}

// Class Structure

type rangeClass_ struct {
	// Declare the class constants.
}

// Class Reference

func rangeClass() *rangeClass_ {
	return rangeClassReference_
}

var rangeClassReference_ = &rangeClass_{
	// Initialize the class constants.
}
