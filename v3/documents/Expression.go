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

func ExpressionClass() ExpressionClassLike {
	return expressionClass()
}

// Constructor Methods

func (c *expressionClass_) Expression(
	subject SubjectLike,
	predicates fra.ListLike[PredicateLike],
) ExpressionLike {
	if uti.IsUndefined(subject) {
		panic("The \"subject\" attribute is required by this class.")
	}
	if uti.IsUndefined(predicates) {
		panic("The \"predicates\" attribute is required by this class.")
	}
	var instance = &expression_{
		// Initialize the instance attributes.
		subject_:    subject,
		predicates_: predicates,
	}
	return instance
}

// INSTANCE INTERFACE

// Principal Methods

func (v *expression_) GetClass() ExpressionClassLike {
	return expressionClass()
}

// Attribute Methods

func (v *expression_) GetSubject() SubjectLike {
	return v.subject_
}

func (v *expression_) GetPredicates() fra.ListLike[PredicateLike] {
	return v.predicates_
}

// PROTECTED INTERFACE

// Instance Structure

type expression_ struct {
	// Declare the instance attributes.
	subject_    SubjectLike
	predicates_ fra.ListLike[PredicateLike]
}

// Class Structure

type expressionClass_ struct {
	// Declare the class constants.
}

// Class Reference

func expressionClass() *expressionClass_ {
	return expressionClassReference_
}

var expressionClassReference_ = &expressionClass_{
	// Initialize the class constants.
}
