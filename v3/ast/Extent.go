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

func ExtentClass() ExtentClassLike {
	return extentClass()
}

// Constructor Methods

func (c *extentClass_) Extent(
	leftBracket LeftBracketLike,
	primitive1 PrimitiveLike,
	delimiter string,
	primitive2 PrimitiveLike,
	rightBracket RightBracketLike,
) ExtentLike {
	if uti.IsUndefined(leftBracket) {
		panic("The \"leftBracket\" attribute is required by this class.")
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
	if uti.IsUndefined(rightBracket) {
		panic("The \"rightBracket\" attribute is required by this class.")
	}
	var instance = &extent_{
		// Initialize the instance attributes.
		leftBracket_:  leftBracket,
		primitive1_:   primitive1,
		delimiter_:    delimiter,
		primitive2_:   primitive2,
		rightBracket_: rightBracket,
	}
	return instance
}

// INSTANCE INTERFACE

// Principal Methods

func (v *extent_) GetClass() ExtentClassLike {
	return extentClass()
}

// Attribute Methods

func (v *extent_) GetLeftBracket() LeftBracketLike {
	return v.leftBracket_
}

func (v *extent_) GetPrimitive1() PrimitiveLike {
	return v.primitive1_
}

func (v *extent_) GetDelimiter() string {
	return v.delimiter_
}

func (v *extent_) GetPrimitive2() PrimitiveLike {
	return v.primitive2_
}

func (v *extent_) GetRightBracket() RightBracketLike {
	return v.rightBracket_
}

// PROTECTED INTERFACE

// Instance Structure

type extent_ struct {
	// Declare the instance attributes.
	leftBracket_  LeftBracketLike
	primitive1_   PrimitiveLike
	delimiter_    string
	primitive2_   PrimitiveLike
	rightBracket_ RightBracketLike
}

// Class Structure

type extentClass_ struct {
	// Declare the class constants.
}

// Class Reference

func extentClass() *extentClass_ {
	return extentClassReference_
}

var extentClassReference_ = &extentClass_{
	// Initialize the class constants.
}
