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

package ast

import (
	uti "github.com/craterdog/go-missing-utilities/v7"
)

// CLASS INTERFACE

// Access Function

func InvokeClass() InvokeClassLike {
	return invokeClass()
}

// Constructor Methods

func (c *invokeClass_) Invoke(
	any_ any,
) InvokeLike {
	if uti.IsUndefined(any_) {
		panic("The \"any\" attribute is required by this class.")
	}
	var instance = &invoke_{
		// Initialize the instance attributes.
		any_: any_,
	}
	return instance
}

// INSTANCE INTERFACE

// Principal Methods

func (v *invoke_) GetClass() InvokeClassLike {
	return invokeClass()
}

// Attribute Methods

func (v *invoke_) GetAny() any {
	return v.any_
}

// PROTECTED INTERFACE

// Instance Structure

type invoke_ struct {
	// Declare the instance attributes.
	any_ any
}

// Class Structure

type invokeClass_ struct {
	// Declare the class constants.
}

// Class Reference

func invokeClass() *invokeClass_ {
	return invokeClassReference_
}

var invokeClassReference_ = &invokeClass_{
	// Initialize the class constants.
}
