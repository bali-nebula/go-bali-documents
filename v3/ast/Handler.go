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
	uti "github.com/craterdog/go-missing-utilities/v7"
)

// CLASS INTERFACE

// Access Function

func HandlerClass() HandlerClassLike {
	return handlerClass()
}

// Constructor Methods

func (c *handlerClass_) Handler(
	delimiter1 string,
	template TemplateLike,
	delimiter2 string,
	procedure ProcedureLike,
) HandlerLike {
	if uti.IsUndefined(delimiter1) {
		panic("The \"delimiter1\" attribute is required by this class.")
	}
	if uti.IsUndefined(template) {
		panic("The \"template\" attribute is required by this class.")
	}
	if uti.IsUndefined(delimiter2) {
		panic("The \"delimiter2\" attribute is required by this class.")
	}
	if uti.IsUndefined(procedure) {
		panic("The \"procedure\" attribute is required by this class.")
	}
	var instance = &handler_{
		// Initialize the instance attributes.
		delimiter1_: delimiter1,
		template_:   template,
		delimiter2_: delimiter2,
		procedure_:  procedure,
	}
	return instance
}

// INSTANCE INTERFACE

// Principal Methods

func (v *handler_) GetClass() HandlerClassLike {
	return handlerClass()
}

// Attribute Methods

func (v *handler_) GetDelimiter1() string {
	return v.delimiter1_
}

func (v *handler_) GetTemplate() TemplateLike {
	return v.template_
}

func (v *handler_) GetDelimiter2() string {
	return v.delimiter2_
}

func (v *handler_) GetProcedure() ProcedureLike {
	return v.procedure_
}

// PROTECTED INTERFACE

// Instance Structure

type handler_ struct {
	// Declare the instance attributes.
	delimiter1_ string
	template_   TemplateLike
	delimiter2_ string
	procedure_  ProcedureLike
}

// Class Structure

type handlerClass_ struct {
	// Declare the class constants.
}

// Class Reference

func handlerClass() *handlerClass_ {
	return handlerClassReference_
}

var handlerClassReference_ = &handlerClass_{
	// Initialize the class constants.
}
