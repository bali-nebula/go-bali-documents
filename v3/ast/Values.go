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
│                 This class file was automatically generated.                 │
│                     Any updates to it may be overwritten.                    │
└──────────────────────────────────────────────────────────────────────────────┘
*/

package ast

import (
	col "github.com/craterdog/go-collection-framework/v7"
	uti "github.com/craterdog/go-missing-utilities/v7"
)

// CLASS INTERFACE

// Access Function

func ValuesClass() ValuesClassLike {
	return valuesClass()
}

// Constructor Methods

func (c *valuesClass_) Values(
	delimiter1 string,
	components col.Sequential[ComponentLike],
	delimiter2 string,
) ValuesLike {
	if uti.IsUndefined(delimiter1) {
		panic("The \"delimiter1\" attribute is required by this class.")
	}
	if uti.IsUndefined(components) {
		panic("The \"components\" attribute is required by this class.")
	}
	if uti.IsUndefined(delimiter2) {
		panic("The \"delimiter2\" attribute is required by this class.")
	}
	var instance = &values_{
		// Initialize the instance attributes.
		delimiter1_: delimiter1,
		components_: components,
		delimiter2_: delimiter2,
	}
	return instance
}

// INSTANCE INTERFACE

// Principal Methods

func (v *values_) GetClass() ValuesClassLike {
	return valuesClass()
}

// Attribute Methods

func (v *values_) GetDelimiter1() string {
	return v.delimiter1_
}

func (v *values_) GetComponents() col.Sequential[ComponentLike] {
	return v.components_
}

func (v *values_) GetDelimiter2() string {
	return v.delimiter2_
}

// PROTECTED INTERFACE

// Instance Structure

type values_ struct {
	// Declare the instance attributes.
	delimiter1_ string
	components_ col.Sequential[ComponentLike]
	delimiter2_ string
}

// Class Structure

type valuesClass_ struct {
	// Declare the class constants.
}

// Class Reference

func valuesClass() *valuesClass_ {
	return valuesClassReference_
}

var valuesClassReference_ = &valuesClass_{
	// Initialize the class constants.
}
