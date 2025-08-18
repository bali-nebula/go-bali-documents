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
	fmt "fmt"
	fra "github.com/craterdog/go-component-framework/v7"
	uti "github.com/craterdog/go-missing-utilities/v7"
)

// CLASS INTERFACE

// Access Function

func DocumentClass() DocumentClassLike {
	return documentClass()
}

// Constructor Methods

func (c *documentClass_) Document(
	component any,
	optionalParameters ParametersLike,
	optionalNote string,
) DocumentLike {
	if uti.IsUndefined(component) {
		panic("The \"component\" attribute is required by this class.")
	}
	var instance = &document_{
		// Initialize the instance attributes.
		component_:          component,
		optionalParameters_: optionalParameters,
		optionalNote_:       optionalNote,
	}
	return instance
}

// Constant Methods

// Function Methods

// INSTANCE INTERFACE

// Principal Methods

func (v *document_) GetClass() DocumentClassLike {
	return documentClass()
}

func (v *document_) GetParameter(
	key fra.SymbolLike,
) DocumentLike {
	var parameters = v.GetOptionalParameters()
	if uti.IsUndefined(parameters) {
		return nil
	}
	var associations = parameters.GetAssociations()
	return associations.GetValue(key)
}

func (v *document_) SetParameter(
	key fra.SymbolLike,
	parameter DocumentLike,
) {
	var parameters = v.GetOptionalParameters()
	if uti.IsUndefined(parameters) {
		return
	}
	var associations = parameters.GetAssociations()
	associations.SetValue(key, parameter)
}

func (v *document_) GetAttribute(
	indices ...any,
) DocumentLike {
	if len(indices) == 0 {
		return nil
	}
	switch collection := v.GetComponent().(type) {
	case EntitiesLike:
		var items = collection.GetItems()
		var index = indices[0].(uti.Index)
		return v.getEntity(items, index, indices[1:]...)
	case AttributesLike:
		var associations = collection.GetAssociations()
		var key = indices[0]
		return v.getValue(associations, key, indices[1:]...)
	}
	return nil
}

func (v *document_) SetAttribute(
	attribute DocumentLike,
	indices ...any,
) {
	if uti.IsUndefined(attribute) || len(indices) == 0 {
		return
	}
	switch collection := v.GetComponent().(type) {
	case EntitiesLike:
		var items = collection.GetItems()
		var index = indices[0].(uti.Index)
		v.setEntity(items, index, attribute, indices[1:]...)
	case AttributesLike:
		var associations = collection.GetAssociations()
		var key = indices[0]
		v.setValue(associations, key, attribute, indices[1:]...)
	}
}

func (v *document_) RemoveAttribute(
	indices ...any,
) {
	if len(indices) == 0 {
		return
	}
	switch collection := v.GetComponent().(type) {
	case EntitiesLike:
		var items = collection.GetItems()
		var index = indices[0].(uti.Index)
		v.removeEntity(items, index, indices[1:]...)
	case AttributesLike:
		var associations = collection.GetAssociations()
		var key = indices[0]
		v.removeValue(associations, key, indices[1:]...)
	}
}

// Attribute Methods

func (v *document_) GetComponent() any {
	return v.component_
}

func (v *document_) GetOptionalParameters() ParametersLike {
	return v.optionalParameters_
}

func (v *document_) GetOptionalNote() string {
	return v.optionalNote_
}

// PROTECTED INTERFACE

// Private Methods

func (v *document_) getEntity(
	items fra.ListLike[DocumentLike],
	index uti.Index,
	indices ...any,
) DocumentLike {
	var size = uti.Index(items.GetSize())
	if size == 0 {
		// The list of items is empty.
		return nil
	}
	if index < 0 {
		// Negative indices start from the end of the list.
		index = size + index + 1
	}
	if index > size {
		// The index is out of bounds.
		return nil
	}
	var document = items.GetValue(index)
	if len(indices) > 0 {
		document = document.GetAttribute(indices...)
	}
	return document
}

func (v *document_) setEntity(
	items fra.ListLike[DocumentLike],
	index uti.Index,
	attribute DocumentLike,
	indices ...any,
) {
	if index == 0 && len(indices) == 0 {
		// Append the attribute to the end of the list.
		items.AppendValue(attribute)
		return
	}
	var size = uti.Index(items.GetSize())
	if size == 0 {
		// The list is empty.
		return
	}
	if index < 0 {
		// Negative indices start from the end of the list.
		index = size + index + 1
	}
	if index > size {
		// The index is out of bounds.
		return
	}
	if len(indices) > 0 {
		var document = items.GetValue(index)
		document.SetAttribute(attribute, indices...)
		return
	}
	items.SetValue(uti.Index(index), attribute)
}

func (v *document_) removeEntity(
	items fra.ListLike[DocumentLike],
	index uti.Index,
	indices ...any,
) {
	var size = uti.Index(items.GetSize())
	if size == 0 {
		// The list of items is empty.
		return
	}
	if index < 0 {
		// Negative indices start from the end of the list.
		index = size + index + 1
	}
	if index > size {
		// The index is out of bounds.
		return
	}
	if len(indices) == 0 {
		items.RemoveValue(index)
		return
	}
	var document = items.GetValue(index)
	document.RemoveAttribute(indices...)
}

func (v *document_) getValue(
	associations fra.CatalogLike[any, DocumentLike],
	key any,
	indices ...any,
) DocumentLike {
	var first = fmt.Sprintf("%v", key)
	var iterator = associations.GetIterator()
	for iterator.HasNext() {
		var association = iterator.GetNext()
		var second = fmt.Sprintf("%v", association.GetKey())
		if first == second {
			var document = association.GetValue()
			if len(indices) > 0 {
				document = document.GetAttribute(indices...)
			}
			return document
		}
	}
	return nil
}

func (v *document_) setValue(
	associations fra.CatalogLike[any, DocumentLike],
	key any,
	attribute DocumentLike,
	indices ...any,
) {
	var first = fmt.Sprintf("%v", key)
	var iterator = associations.GetIterator()
	for iterator.HasNext() {
		var association = iterator.GetNext()
		var second = fmt.Sprintf("%v", association.GetKey())
		if first == second {
			if len(indices) > 0 {
				var document = association.GetValue()
				document.SetAttribute(attribute, indices...)
				return
			}
			association.SetValue(attribute)
			return
		}
	}
}

func (v *document_) removeValue(
	associations fra.CatalogLike[any, DocumentLike],
	key any,
	indices ...any,
) {
	var first = fmt.Sprintf("%v", key)
	var iterator = associations.GetIterator()
	for iterator.HasNext() {
		var association = iterator.GetNext()
		var second = fmt.Sprintf("%v", association.GetKey())
		if first == second {
			if len(indices) > 0 {
				var document = association.GetValue()
				document.RemoveAttribute(indices...)
			} else {
				associations.RemoveValue(key)
			}
			return
		}
	}
}

// Instance Structure

type document_ struct {
	// Declare the instance attributes.
	component_          any
	optionalParameters_ ParametersLike
	optionalNote_       string
}

// Class Structure

type documentClass_ struct {
	// Declare the class constants.
}

// Class Reference

func documentClass() *documentClass_ {
	return documentClassReference_
}

var documentClassReference_ = &documentClass_{
	// Initialize the class constants.
}
