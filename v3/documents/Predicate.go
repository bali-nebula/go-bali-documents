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
	uti "github.com/craterdog/go-essential-utilities/v8"
)

// CLASS INTERFACE

// Access Function

func PredicateClass() PredicateClassLike {
	return predicateClass()
}

// Constructor Methods

func (c *predicateClass_) Predicate(
	operation Operation,
	expression ExpressionLike,
) PredicateLike {
	if uti.IsUndefined(operation) {
		panic("The \"operation\" attribute is required by this class.")
	}
	if uti.IsUndefined(expression) {
		panic("The \"expression\" attribute is required by this class.")
	}
	var instance = &predicate_{
		// Initialize the instance attributes.
		operation_:  operation,
		expression_: expression,
	}
	return instance
}

// Constant Methods

// Function Methods

// INSTANCE INTERFACE

// Principal Methods

func (v *predicate_) GetClass() PredicateClassLike {
	return predicateClass()
}

// Attribute Methods

func (v *predicate_) GetOperation() Operation {
	return v.operation_
}

func (v *predicate_) GetExpression() ExpressionLike {
	return v.expression_
}

// PROTECTED INTERFACE

// Private Methods

// Instance Structure

type predicate_ struct {
	// Declare the instance attributes.
	operation_  Operation
	expression_ ExpressionLike
}

// Class Structure

type predicateClass_ struct {
	// Declare the class constants.
}

// Class Reference

func predicateClass() *predicateClass_ {
	return predicateClassReference_
}

var predicateClassReference_ = &predicateClass_{
	// Initialize the class constants.
}
