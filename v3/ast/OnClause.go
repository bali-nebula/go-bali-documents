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

func OnClauseClass() OnClauseClassLike {
	return onClauseClass()
}

// Constructor Methods

func (c *onClauseClass_) OnClause(
	delimiter string,
	failure FailureLike,
	handlers col.Sequential[HandlerLike],
) OnClauseLike {
	if uti.IsUndefined(delimiter) {
		panic("The \"delimiter\" attribute is required by this class.")
	}
	if uti.IsUndefined(failure) {
		panic("The \"failure\" attribute is required by this class.")
	}
	if uti.IsUndefined(handlers) {
		panic("The \"handlers\" attribute is required by this class.")
	}
	var instance = &onClause_{
		// Initialize the instance attributes.
		delimiter_: delimiter,
		failure_:   failure,
		handlers_:  handlers,
	}
	return instance
}

// INSTANCE INTERFACE

// Principal Methods

func (v *onClause_) GetClass() OnClauseClassLike {
	return onClauseClass()
}

// Attribute Methods

func (v *onClause_) GetDelimiter() string {
	return v.delimiter_
}

func (v *onClause_) GetFailure() FailureLike {
	return v.failure_
}

func (v *onClause_) GetHandlers() col.Sequential[HandlerLike] {
	return v.handlers_
}

// PROTECTED INTERFACE

// Instance Structure

type onClause_ struct {
	// Declare the instance attributes.
	delimiter_ string
	failure_   FailureLike
	handlers_  col.Sequential[HandlerLike]
}

// Class Structure

type onClauseClass_ struct {
	// Declare the class constants.
}

// Class Reference

func onClauseClass() *onClauseClass_ {
	return onClauseClassReference_
}

var onClauseClassReference_ = &onClauseClass_{
	// Initialize the class constants.
}
