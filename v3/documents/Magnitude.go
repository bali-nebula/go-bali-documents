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

func MagnitudeClass() MagnitudeClassLike {
	return magnitudeClass()
}

// Constructor Methods

func (c *magnitudeClass_) Magnitude(
	delimiter1 string,
	expression ExpressionLike,
	delimiter2 string,
) MagnitudeLike {
	if uti.IsUndefined(delimiter1) {
		panic("The \"delimiter1\" attribute is required by this class.")
	}
	if uti.IsUndefined(expression) {
		panic("The \"expression\" attribute is required by this class.")
	}
	if uti.IsUndefined(delimiter2) {
		panic("The \"delimiter2\" attribute is required by this class.")
	}
	var instance = &magnitude_{
		// Initialize the instance attributes.
		delimiter1_: delimiter1,
		expression_: expression,
		delimiter2_: delimiter2,
	}
	return instance
}

// INSTANCE INTERFACE

// Principal Methods

func (v *magnitude_) GetClass() MagnitudeClassLike {
	return magnitudeClass()
}

// Attribute Methods

func (v *magnitude_) GetDelimiter1() string {
	return v.delimiter1_
}

func (v *magnitude_) GetExpression() ExpressionLike {
	return v.expression_
}

func (v *magnitude_) GetDelimiter2() string {
	return v.delimiter2_
}

// PROTECTED INTERFACE

// Instance Structure

type magnitude_ struct {
	// Declare the instance attributes.
	delimiter1_ string
	expression_ ExpressionLike
	delimiter2_ string
}

// Class Structure

type magnitudeClass_ struct {
	// Declare the class constants.
}

// Class Reference

func magnitudeClass() *magnitudeClass_ {
	return magnitudeClassReference_
}

var magnitudeClassReference_ = &magnitudeClass_{
	// Initialize the class constants.
}
