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

func RepositoryAccessClass() RepositoryAccessClassLike {
	return repositoryAccessClass()
}

// Constructor Methods

func (c *repositoryAccessClass_) RepositoryAccess(
	any_ any,
) RepositoryAccessLike {
	if uti.IsUndefined(any_) {
		panic("The \"any\" attribute is required by this class.")
	}
	var instance = &repositoryAccess_{
		// Initialize the instance attributes.
		any_: any_,
	}
	return instance
}

// INSTANCE INTERFACE

// Principal Methods

func (v *repositoryAccess_) GetClass() RepositoryAccessClassLike {
	return repositoryAccessClass()
}

// Attribute Methods

func (v *repositoryAccess_) GetAny() any {
	return v.any_
}

// PROTECTED INTERFACE

// Instance Structure

type repositoryAccess_ struct {
	// Declare the instance attributes.
	any_ any
}

// Class Structure

type repositoryAccessClass_ struct {
	// Declare the class constants.
}

// Class Reference

func repositoryAccessClass() *repositoryAccessClass_ {
	return repositoryAccessClassReference_
}

var repositoryAccessClassReference_ = &repositoryAccessClass_{
	// Initialize the class constants.
}
