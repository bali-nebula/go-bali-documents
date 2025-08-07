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
	fra "github.com/craterdog/go-component-framework/v7"
	uti "github.com/craterdog/go-missing-utilities/v7"
)

// CLASS INTERFACE

// Access Function

func MethodClass() MethodClassLike {
	return methodClass()
}

// Constructor Methods

func (c *methodClass_) Method(
	identifier1 string,
	invoke InvokeLike,
	identifier2 string,
	delimiter1 string,
	arguments fra.ListLike[ArgumentLike],
	delimiter2 string,
) MethodLike {
	if uti.IsUndefined(identifier1) {
		panic("The \"identifier1\" attribute is required by this class.")
	}
	if uti.IsUndefined(invoke) {
		panic("The \"invoke\" attribute is required by this class.")
	}
	if uti.IsUndefined(identifier2) {
		panic("The \"identifier2\" attribute is required by this class.")
	}
	if uti.IsUndefined(delimiter1) {
		panic("The \"delimiter1\" attribute is required by this class.")
	}
	if uti.IsUndefined(arguments) {
		panic("The \"arguments\" attribute is required by this class.")
	}
	if uti.IsUndefined(delimiter2) {
		panic("The \"delimiter2\" attribute is required by this class.")
	}
	var instance = &method_{
		// Initialize the instance attributes.
		identifier1_: identifier1,
		invoke_:      invoke,
		identifier2_: identifier2,
		delimiter1_:  delimiter1,
		arguments_:   arguments,
		delimiter2_:  delimiter2,
	}
	return instance
}

// INSTANCE INTERFACE

// Principal Methods

func (v *method_) GetClass() MethodClassLike {
	return methodClass()
}

// Attribute Methods

func (v *method_) GetIdentifier1() string {
	return v.identifier1_
}

func (v *method_) GetInvoke() InvokeLike {
	return v.invoke_
}

func (v *method_) GetIdentifier2() string {
	return v.identifier2_
}

func (v *method_) GetDelimiter1() string {
	return v.delimiter1_
}

func (v *method_) GetArguments() fra.ListLike[ArgumentLike] {
	return v.arguments_
}

func (v *method_) GetDelimiter2() string {
	return v.delimiter2_
}

// PROTECTED INTERFACE

// Instance Structure

type method_ struct {
	// Declare the instance attributes.
	identifier1_ string
	invoke_      InvokeLike
	identifier2_ string
	delimiter1_  string
	arguments_   fra.ListLike[ArgumentLike]
	delimiter2_  string
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
