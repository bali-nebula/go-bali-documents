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
	uti "github.com/craterdog/go-essential-utilities/v8"
)

// CLASS INTERFACE

// Access Function

func RangeClass() RangeClassLike {
	return rangeClass()
}

// Constructor Methods

func (c *rangeClass_) Range(
	left com.Bracket,
	optionalFirst any,
	optionalLast any,
	right com.Bracket,
) RangeLike {
	if uti.IsUndefined(left) {
		panic("The \"left\" attribute is required by this class.")
	}
	if uti.IsUndefined(right) {
		panic("The \"right\" attribute is required by this class.")
	}
	var instance = &range_{
		// Initialize the instance attributes.
		left_:          left,
		optionalFirst_: optionalFirst,
		optionalLast_:  optionalLast,
		right_:         right,
	}
	return instance
}

// Constant Methods

// Function Methods

// INSTANCE INTERFACE

// Principal Methods

func (v *range_) GetClass() RangeClassLike {
	return rangeClass()
}

// Attribute Methods

func (v *range_) GetLeft() com.Bracket {
	return v.left_
}

func (v *range_) GetOptionalFirst() any {
	return v.optionalFirst_
}

func (v *range_) GetOptionalLast() any {
	return v.optionalLast_
}

func (v *range_) GetRight() com.Bracket {
	return v.right_
}

// PROTECTED INTERFACE

// Private Methods

// Instance Structure

type range_ struct {
	// Declare the instance attributes.
	left_          com.Bracket
	optionalFirst_ any
	optionalLast_  any
	right_         com.Bracket
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
