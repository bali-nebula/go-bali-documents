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

func DoClauseClass() DoClauseClassLike {
	return doClauseClass()
}

// Constructor Methods

func (c *doClauseClass_) DoClause(
	delimiter string,
	invocation InvocationLike,
) DoClauseLike {
	if uti.IsUndefined(delimiter) {
		panic("The \"delimiter\" attribute is required by this class.")
	}
	if uti.IsUndefined(invocation) {
		panic("The \"invocation\" attribute is required by this class.")
	}
	var instance = &doClause_{
		// Initialize the instance attributes.
		delimiter_:  delimiter,
		invocation_: invocation,
	}
	return instance
}

// INSTANCE INTERFACE

// Principal Methods

func (v *doClause_) GetClass() DoClauseClassLike {
	return doClauseClass()
}

// Attribute Methods

func (v *doClause_) GetDelimiter() string {
	return v.delimiter_
}

func (v *doClause_) GetInvocation() InvocationLike {
	return v.invocation_
}

// PROTECTED INTERFACE

// Instance Structure

type doClause_ struct {
	// Declare the instance attributes.
	delimiter_  string
	invocation_ InvocationLike
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
