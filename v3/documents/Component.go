/*
................................................................................
.    Copyright (c) 2009-2026 Crater Dog Technologies™.  All Rights Reserved.   .
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
	com "github.com/craterdog/go-essential-composites/v8"
	pri "github.com/craterdog/go-essential-primitives/v8"
	uti "github.com/craterdog/go-essential-utilities/v8"
)

// CLASS INTERFACE

// Access Function

func ComponentClass() ComponentClassLike {
	return componentClass()
}

// Constructor Methods

func (c *componentClass_) Component(
	entity any,
	optionalGenerics GenericsLike,
) ComponentLike {
	if uti.IsUndefined(entity) {
		panic("The \"entity\" attribute is required by this class.")
	}
	var instance = &component_{
		// Initialize the instance attributes.
		entity_:           entity,
		optionalGenerics_: optionalGenerics,
	}
	return instance
}

// Constant Methods

// Function Methods

// INSTANCE INTERFACE

// Principal Methods

func (v *component_) GetClass() ComponentClassLike {
	return componentClass()
}

// Composite Methods

func (v *component_) GetEntity() any {
	return v.entity_
}

func (v *component_) GetOptionalGenerics() GenericsLike {
	return v.optionalGenerics_
}

func (v *component_) GetParameter(
	symbol pri.SymbolLike,
) Composite {
	var parameter Composite
	var generics = v.GetOptionalGenerics()
	if uti.IsDefined(generics) {
		var parameters = generics.GetParameters()
		var constraint = parameters.GetValue(symbol)
		if uti.IsDefined(constraint) {
			parameter = componentClass().Component(
				constraint.GetLiteral(),
				constraint.GetOptionalGenerics(),
			)
		}
	}
	return parameter
}

func (v *component_) SetSubcomponent(
	value any,
	indices ...any,
) {
	var content ContentLike
	switch actual := value.(type) {
	case Composite:
		content = ContentClass().Content(actual, "")
	case DocumentLike:
		content = ContentClass().Content(actual.GetComposite(), "")
	case ContentLike:
		content = actual
	default:
		content = ContentClass().Content(
			componentClass().Component(value, nil), "",
		)
	}
	if uti.IsDefined(content) && len(indices) > 0 {
		var key = indices[0]
		indices = indices[1:]
		switch collection := v.GetEntity().(type) {
		case ItemsLike:
			switch contents := collection.GetContents().(type) {
			case com.ListLike[ContentLike]:
				var index = key.(int)
				v.setItem(contents, content, index, indices...)
			default:
				var message = fmt.Sprintf(
					"Attempted to set an item in a non-list type: %T",
					contents,
				)
				panic(message)
			}
		case AttributesLike:
			var associations = collection.GetAssociations()
			v.setAttribute(associations, key, content, indices...)
		}
	}
}

func (v *component_) GetSubcomponent(
	indices ...any,
) ContentLike {
	var content ContentLike
	if len(indices) > 0 {
		var key = indices[0]
		indices = indices[1:]
		switch collection := v.GetEntity().(type) {
		case ItemsLike:
			var contents = collection.GetContents()
			var index = key.(int)
			content = v.getItem(contents, index, indices...)
		case AttributesLike:
			var associations = collection.GetAssociations()
			content = v.getAttribute(associations, key, indices...)
		}
	}
	return content
}

func (v *component_) RemoveSubcomponent(
	indices ...any,
) ContentLike {
	var content ContentLike
	if len(indices) > 0 {
		var key = indices[0]
		indices = indices[1:]
		switch collection := v.GetEntity().(type) {
		case ItemsLike:
			switch contents := collection.GetContents().(type) {
			case com.ListLike[ContentLike]:
				var index = key.(int)
				content = v.removeItem(contents, index, indices...)
			default:
				var message = fmt.Sprintf(
					"Attempted to remove an item from a non-list type: %T",
					contents,
				)
				panic(message)
			}
		case AttributesLike:
			var associations = collection.GetAssociations()
			content = v.removeAttribute(associations, key, indices...)
		}
	}
	return content
}

// PROTECTED INTERFACE

// Private Methods

func (v *component_) getItem(
	contents com.Sequential[ContentLike],
	index int,
	indices ...any,
) ContentLike {
	var content ContentLike
	var size = int(contents.GetSize())
	if size == 0 {
		// The list of contents is empty.
		return content
	}
	if index < 0 {
		// Negative indices start from the end of the list.
		index = size + index + 1
	}
	if index > size {
		// The index is out of bounds.
		return content
	}
	content = contents.AsArray()[index-1]
	if len(indices) > 0 {
		var composite = content.GetComposite()
		content = composite.GetSubcomponent(indices...)
	}
	return content
}

func (v *component_) setItem(
	contents com.ListLike[ContentLike],
	content ContentLike,
	index int,
	indices ...any,
) {
	if index == 0 && len(indices) == 0 {
		// Append the attribute to the end of the list.
		contents.AppendValue(content)
		return
	}
	var size = int(contents.GetSize())
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
		var composite = contents.GetValue(index).GetComposite()
		composite.SetSubcomponent(content, indices...)
		return
	}
	contents.SetValue(index, content)
}

func (v *component_) removeItem(
	contents com.ListLike[ContentLike],
	index int,
	indices ...any,
) ContentLike {
	var content ContentLike
	var size = int(contents.GetSize())
	if size == 0 {
		// The list of contents is empty.
		return content
	}
	if index < 0 {
		// Negative indices start from the end of the list.
		index = size + index + 1
	}
	if index > size {
		// The index is out of bounds.
		return content
	}
	if len(indices) == 0 {
		content = contents.GetValue(index)
		contents.RemoveValue(index)
		return content
	}
	content = contents.GetValue(index)
	var composite = content.GetComposite()
	content = composite.RemoveSubcomponent(indices...)
	return content
}

func (v *component_) getAttribute(
	associations com.CatalogLike[any, ContentLike],
	key any,
	indices ...any,
) ContentLike {
	var content ContentLike
	var first = fmt.Sprintf("%v", key)
	var iterator = associations.GetIterator()
	for iterator.HasNext() {
		var association = iterator.GetNext()
		var second = fmt.Sprintf("%v", association.GetKey())
		if first == second {
			content = association.GetValue()
			if len(indices) > 0 {
				var composite = content.GetComposite()
				content = composite.GetSubcomponent(indices...)
			}
			break
		}
	}
	return content
}

func (v *component_) setAttribute(
	associations com.CatalogLike[any, ContentLike],
	key any,
	content ContentLike,
	indices ...any,
) {
	var first = fmt.Sprintf("%v", key)
	var iterator = associations.GetIterator()
	for iterator.HasNext() {
		var association = iterator.GetNext()
		var second = fmt.Sprintf("%v", association.GetKey())
		if first == second {
			if len(indices) > 0 {
				var composite = association.GetValue().GetComposite()
				composite.SetSubcomponent(content, indices...)
			} else {
				association.SetValue(content)
			}
			return
		}
	}
	associations.SetValue(key, content)
}

func (v *component_) removeAttribute(
	associations com.CatalogLike[any, ContentLike],
	key any,
	indices ...any,
) ContentLike {
	var content ContentLike
	var first = fmt.Sprintf("%v", key)
	var iterator = associations.GetIterator()
	for iterator.HasNext() {
		var association = iterator.GetNext()
		var second = fmt.Sprintf("%v", association.GetKey())
		if first == second {
			if len(indices) > 0 {
				content = association.GetValue()
				var composite = content.GetComposite()
				content = composite.RemoveSubcomponent(indices...)
			} else {
				content = associations.RemoveValue(key)
			}
			break
		}
	}
	return content
}

// Instance Structure

type component_ struct {
	// Declare the instance attributes.
	entity_           any
	optionalGenerics_ GenericsLike
}

// Class Structure

type componentClass_ struct {
	// Declare the class constants.
}

// Class Reference

func componentClass() *componentClass_ {
	return componentClassReference_
}

var componentClassReference_ = &componentClass_{
	// Initialize the class constants.
}
