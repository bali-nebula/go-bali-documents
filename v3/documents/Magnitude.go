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

package documents

import (
	uti "github.com/craterdog/go-missing-utilities/v8"
)

// CLASS INTERFACE

// Access Function

func MagnitudeClass() MagnitudeClassLike {
	return magnitudeClass()
}

// Constructor Methods

func (c *magnitudeClass_) Magnitude(
	expression ExpressionLike,
) MagnitudeLike {
	if uti.IsUndefined(expression) {
		panic("The \"expression\" attribute is required by this class.")
	}
	var instance = &magnitude_{
		// Initialize the instance attributes.
		expression_: expression,
	}
	return instance
}

// Constant Methods

// Function Methods

// INSTANCE INTERFACE

// Principal Methods

func (v *magnitude_) GetClass() MagnitudeClassLike {
	return magnitudeClass()
}

// Attribute Methods

func (v *magnitude_) GetExpression() ExpressionLike {
	return v.expression_
}

// PROTECTED INTERFACE

// Private Methods

// Instance Structure

type magnitude_ struct {
	// Declare the instance attributes.
	expression_ ExpressionLike
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
