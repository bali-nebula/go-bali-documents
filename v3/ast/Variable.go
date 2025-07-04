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

func VariableClass() VariableClassLike {
	return variableClass()
}

// Constructor Methods

func (c *variableClass_) Variable(
	symbol string,
) VariableLike {
	if uti.IsUndefined(symbol) {
		panic("The \"symbol\" attribute is required by this class.")
	}
	var instance = &variable_{
		// Initialize the instance attributes.
		symbol_: symbol,
	}
	return instance
}

// INSTANCE INTERFACE

// Principal Methods

func (v *variable_) GetClass() VariableClassLike {
	return variableClass()
}

// Attribute Methods

func (v *variable_) GetSymbol() string {
	return v.symbol_
}

// PROTECTED INTERFACE

// Instance Structure

type variable_ struct {
	// Declare the instance attributes.
	symbol_ string
}

// Class Structure

type variableClass_ struct {
	// Declare the class constants.
}

// Class Reference

func variableClass() *variableClass_ {
	return variableClassReference_
}

var variableClassReference_ = &variableClass_{
	// Initialize the class constants.
}
