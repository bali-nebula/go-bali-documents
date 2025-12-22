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

package documents

import (
	pri "github.com/craterdog/go-essential-primitives/v8"
	uti "github.com/craterdog/go-essential-utilities/v8"
)

// CLASS INTERFACE

// Access Function

func DefineClauseClass() DefineClauseClassLike {
	return defineClauseClass()
}

// Constructor Methods

func (c *defineClauseClass_) DefineClause(
	constant pri.SymbolLike,
	expression ExpressionLike,
) DefineClauseLike {
	if uti.IsUndefined(constant) {
		panic("The \"constant\" attribute is required by this class.")
	}
	if uti.IsUndefined(expression) {
		panic("The \"expression\" attribute is required by this class.")
	}
	var instance = &defineClause_{
		// Initialize the instance attributes.
		constant_:   constant,
		expression_: expression,
	}
	return instance
}

// Constant Methods

// Function Methods

// INSTANCE INTERFACE

// Principal Methods

func (v *defineClause_) GetClass() DefineClauseClassLike {
	return defineClauseClass()
}

// Attribute Methods

func (v *defineClause_) GetConstant() pri.SymbolLike {
	return v.constant_
}

func (v *defineClause_) GetExpression() ExpressionLike {
	return v.expression_
}

// PROTECTED INTERFACE

// Private Methods

// Instance Structure

type defineClause_ struct {
	// Declare the instance attributes.
	constant_   pri.SymbolLike
	expression_ ExpressionLike
}

// Class Structure

type defineClauseClass_ struct {
	// Declare the class constants.
}

// Class Reference

func defineClauseClass() *defineClauseClass_ {
	return defineClauseClassReference_
}

var defineClauseClassReference_ = &defineClauseClass_{
	// Initialize the class constants.
}
