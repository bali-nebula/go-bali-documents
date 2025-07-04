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

func MessageHandlingClass() MessageHandlingClassLike {
	return messageHandlingClass()
}

// Constructor Methods

func (c *messageHandlingClass_) MessageHandling(
	any_ any,
) MessageHandlingLike {
	if uti.IsUndefined(any_) {
		panic("The \"any\" attribute is required by this class.")
	}
	var instance = &messageHandling_{
		// Initialize the instance attributes.
		any_: any_,
	}
	return instance
}

// INSTANCE INTERFACE

// Principal Methods

func (v *messageHandling_) GetClass() MessageHandlingClassLike {
	return messageHandlingClass()
}

// Attribute Methods

func (v *messageHandling_) GetAny() any {
	return v.any_
}

// PROTECTED INTERFACE

// Instance Structure

type messageHandling_ struct {
	// Declare the instance attributes.
	any_ any
}

// Class Structure

type messageHandlingClass_ struct {
	// Declare the class constants.
}

// Class Reference

func messageHandlingClass() *messageHandlingClass_ {
	return messageHandlingClassReference_
}

var messageHandlingClassReference_ = &messageHandlingClass_{
	// Initialize the class constants.
}
