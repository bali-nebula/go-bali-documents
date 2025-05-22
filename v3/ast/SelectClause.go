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

func SelectClauseClass() SelectClauseClassLike {
	return selectClauseClass()
}

// Constructor Methods

func (c *selectClauseClass_) SelectClause(
	delimiter string,
	target TargetLike,
	handlers col.Sequential[HandlerLike],
) SelectClauseLike {
	if uti.IsUndefined(delimiter) {
		panic("The \"delimiter\" attribute is required by this class.")
	}
	if uti.IsUndefined(target) {
		panic("The \"target\" attribute is required by this class.")
	}
	if uti.IsUndefined(handlers) {
		panic("The \"handlers\" attribute is required by this class.")
	}
	var instance = &selectClause_{
		// Initialize the instance attributes.
		delimiter_: delimiter,
		target_:    target,
		handlers_:  handlers,
	}
	return instance
}

// INSTANCE INTERFACE

// Principal Methods

func (v *selectClause_) GetClass() SelectClauseClassLike {
	return selectClauseClass()
}

// Attribute Methods

func (v *selectClause_) GetDelimiter() string {
	return v.delimiter_
}

func (v *selectClause_) GetTarget() TargetLike {
	return v.target_
}

func (v *selectClause_) GetHandlers() col.Sequential[HandlerLike] {
	return v.handlers_
}

// PROTECTED INTERFACE

// Instance Structure

type selectClause_ struct {
	// Declare the instance attributes.
	delimiter_ string
	target_    TargetLike
	handlers_  col.Sequential[HandlerLike]
}

// Class Structure

type selectClauseClass_ struct {
	// Declare the class constants.
}

// Class Reference

func selectClauseClass() *selectClauseClass_ {
	return selectClauseClassReference_
}

var selectClauseClassReference_ = &selectClauseClass_{
	// Initialize the class constants.
}
