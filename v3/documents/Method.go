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

func MethodClass() MethodClassLike {
	return methodClass()
}

// Constructor Methods

func (c *methodClass_) Method(
	target string,
	identifier string,
	arguments com.Sequential[any],
	isSynchronous bool,
) MethodLike {
	if uti.IsUndefined(target) {
		panic("The \"target\" attribute is required by this class.")
	}
	if uti.IsUndefined(identifier) {
		panic("The \"identifier\" attribute is required by this class.")
	}
	if uti.IsUndefined(arguments) {
		panic("The \"arguments\" attribute is required by this class.")
	}
	var instance = &method_{
		// Initialize the instance attributes.
		target_:        target,
		identifier_:    identifier,
		arguments_:     arguments,
		isSynchronous_: isSynchronous,
	}
	return instance
}

// Constant Methods

// Function Methods

// INSTANCE INTERFACE

// Principal Methods

func (v *method_) GetClass() MethodClassLike {
	return methodClass()
}

// Attribute Methods

func (v *method_) GetTarget() string {
	return v.target_
}

func (v *method_) GetIdentifier() string {
	return v.identifier_
}

func (v *method_) GetArguments() com.Sequential[any] {
	return v.arguments_
}

func (v *method_) IsSynchronous() bool {
	return v.isSynchronous_
}

// PROTECTED INTERFACE

// Private Methods

// Instance Structure

type method_ struct {
	// Declare the instance attributes.
	target_        string
	identifier_    string
	arguments_     com.Sequential[any]
	isSynchronous_ bool
}

// Class Structure

type methodClass_ struct {
	// Declare the class constants.
}

// Class Reference

func methodClass() *methodClass_ {
	return methodClassReference_
}

var methodClassReference_ = &methodClass_{
	// Initialize the class constants.
}
