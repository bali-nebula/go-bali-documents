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
	uti "github.com/craterdog/go-essential-utilities/v8"
)

// CLASS INTERFACE

// Access Function

func InvokeClauseClass() InvokeClauseClassLike {
	return invokeClauseClass()
}

// Constructor Methods

func (c *invokeClauseClass_) InvokeClause(
	method MethodLike,
) InvokeClauseLike {
	if uti.IsUndefined(method) {
		panic("The \"method\" attribute is required by this class.")
	}
	var instance = &invokeClause_{
		// Initialize the instance attributes.
		method_: method,
	}
	return instance
}

// Constant Methods

// Function Methods

// INSTANCE INTERFACE

// Principal Methods

func (v *invokeClause_) GetClass() InvokeClauseClassLike {
	return invokeClauseClass()
}

// Attribute Methods

func (v *invokeClause_) GetMethod() MethodLike {
	return v.method_
}

// PROTECTED INTERFACE

// Private Methods

// Instance Structure

type invokeClause_ struct {
	// Declare the instance attributes.
	method_ MethodLike
}

// Class Structure

type invokeClauseClass_ struct {
	// Declare the class constants.
}

// Class Reference

func invokeClauseClass() *invokeClauseClass_ {
	return invokeClauseClassReference_
}

var invokeClauseClassReference_ = &invokeClauseClass_{
	// Initialize the class constants.
}
