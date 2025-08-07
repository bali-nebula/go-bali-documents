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
	bra BraLike,
	primitive1 PrimitiveLike,
	delimiter string,
	primitive2 PrimitiveLike,
	ket KetLike,
) RangeLike {
	if uti.IsUndefined(bra) {
		panic("The \"bra\" attribute is required by this class.")
	}
	if uti.IsUndefined(primitive1) {
		panic("The \"primitive1\" attribute is required by this class.")
	}
	if uti.IsUndefined(delimiter) {
		panic("The \"delimiter\" attribute is required by this class.")
	}
	if uti.IsUndefined(primitive2) {
		panic("The \"primitive2\" attribute is required by this class.")
	}
	if uti.IsUndefined(ket) {
		panic("The \"ket\" attribute is required by this class.")
	}
	var instance = &range_{
		// Initialize the instance attributes.
		bra_:        bra,
		primitive1_: primitive1,
		delimiter_:  delimiter,
		primitive2_: primitive2,
		ket_:        ket,
	}
	return instance
}

// INSTANCE INTERFACE

// Principal Methods

func (v *range_) GetClass() RangeClassLike {
	return rangeClass()
}

// Attribute Methods

func (v *range_) GetBra() BraLike {
	return v.bra_
}

func (v *range_) GetPrimitive1() PrimitiveLike {
	return v.primitive1_
}

func (v *range_) GetDelimiter() string {
	return v.delimiter_
}

func (v *range_) GetPrimitive2() PrimitiveLike {
	return v.primitive2_
}

func (v *range_) GetKet() KetLike {
	return v.ket_
}

// PROTECTED INTERFACE

// Instance Structure

type range_ struct {
	// Declare the instance attributes.
	bra_        BraLike
	primitive1_ PrimitiveLike
	delimiter_  string
	primitive2_ PrimitiveLike
	ket_        KetLike
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
