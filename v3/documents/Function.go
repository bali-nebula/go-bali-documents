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

func FunctionClass() FunctionClassLike {
	return functionClass()
}

// Constructor Methods

func (c *functionClass_) Function(
	identifier string,
	delimiter1 string,
	arguments fra.ListLike[ArgumentLike],
	delimiter2 string,
) FunctionLike {
	if uti.IsUndefined(identifier) {
		panic("The \"identifier\" attribute is required by this class.")
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
	var instance = &function_{
		// Initialize the instance attributes.
		identifier_: identifier,
		delimiter1_: delimiter1,
		arguments_:  arguments,
		delimiter2_: delimiter2,
	}
	return instance
}

// INSTANCE INTERFACE

// Principal Methods

func (v *function_) GetClass() FunctionClassLike {
	return functionClass()
}

// Attribute Methods

func (v *function_) GetIdentifier() string {
	return v.identifier_
}

func (v *function_) GetDelimiter1() string {
	return v.delimiter1_
}

func (v *function_) GetArguments() fra.ListLike[ArgumentLike] {
	return v.arguments_
}

func (v *function_) GetDelimiter2() string {
	return v.delimiter2_
}

// PROTECTED INTERFACE

// Instance Structure

type function_ struct {
	// Declare the instance attributes.
	identifier_ string
	delimiter1_ string
	arguments_  fra.ListLike[ArgumentLike]
	delimiter2_ string
}

// Class Structure

type functionClass_ struct {
	// Declare the class constants.
}

// Class Reference

func functionClass() *functionClass_ {
	return functionClassReference_
}

var functionClassReference_ = &functionClass_{
	// Initialize the class constants.
}
