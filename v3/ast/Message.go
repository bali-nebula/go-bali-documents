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

func MessageClass() MessageClassLike {
	return messageClass()
}

// Constructor Methods

func (c *messageClass_) Message(
	expression ExpressionLike,
) MessageLike {
	if uti.IsUndefined(expression) {
		panic("The \"expression\" attribute is required by this class.")
	}
	var instance = &message_{
		// Initialize the instance attributes.
		expression_: expression,
	}
	return instance
}

// INSTANCE INTERFACE

// Principal Methods

func (v *message_) GetClass() MessageClassLike {
	return messageClass()
}

// Attribute Methods

func (v *message_) GetExpression() ExpressionLike {
	return v.expression_
}

// PROTECTED INTERFACE

// Instance Structure

type message_ struct {
	// Declare the instance attributes.
	expression_ ExpressionLike
}

// Class Structure

type messageClass_ struct {
	// Declare the class constants.
}

// Class Reference

func messageClass() *messageClass_ {
	return messageClassReference_
}

var messageClassReference_ = &messageClass_{
	// Initialize the class constants.
}
