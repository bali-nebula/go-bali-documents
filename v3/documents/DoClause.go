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

func DoClauseClass() DoClauseClassLike {
	return doClauseClass()
}

// Constructor Methods

func (c *doClauseClass_) DoClause(
	method MethodLike,
) DoClauseLike {
	if uti.IsUndefined(method) {
		panic("The \"method\" attribute is required by this class.")
	}
	var instance = &doClause_{
		// Initialize the instance attributes.
		method_: method,
	}
	return instance
}

// Constant Methods

// Function Methods

// INSTANCE INTERFACE

// Principal Methods

func (v *doClause_) GetClass() DoClauseClassLike {
	return doClauseClass()
}

// Attribute Methods

func (v *doClause_) GetMethod() MethodLike {
	return v.method_
}

// PROTECTED INTERFACE

// Private Methods

// Instance Structure

type doClause_ struct {
	// Declare the instance attributes.
	method_ MethodLike
}

// Class Structure

type doClauseClass_ struct {
	// Declare the class constants.
}

// Class Reference

func doClauseClass() *doClauseClass_ {
	return doClauseClassReference_
}

var doClauseClassReference_ = &doClauseClass_{
	// Initialize the class constants.
}
