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
	var entry EntryLike
	switch actual := value.(type) {
	case Composite:
		entry = EntryClass().Entry(actual, "")
	case DocumentLike:
		entry = EntryClass().Entry(actual.GetComposite(), "")
	case EntryLike:
		entry = actual
	default:
		entry = EntryClass().Entry(
			componentClass().Component(value, nil), "",
		)
	}
	if uti.IsDefined(entry) && len(indices) > 0 {
		var key = indices[0]
		indices = indices[1:]
		switch collection := v.GetEntity().(type) {
		case ItemsLike:
			switch entries := collection.GetEntries().(type) {
			case com.ListLike[EntryLike]:
				var index = key.(int)
				v.setItem(entries, entry, index, indices...)
			default:
				var message = fmt.Sprintf(
					"Attempted to set an item in a non-list type: %T",
					entries,
				)
				panic(message)
			}
		case AttributesLike:
			var associations = collection.GetAssociations()
			v.setAttribute(associations, key, entry, indices...)
		}
	}
}

func (v *component_) GetSubcomponent(
	indices ...any,
) EntryLike {
	var entry EntryLike
	if len(indices) > 0 {
		var key = indices[0]
		indices = indices[1:]
		switch collection := v.GetEntity().(type) {
		case ItemsLike:
			var entries = collection.GetEntries()
			var index = key.(int)
			entry = v.getItem(entries, index, indices...)
		case AttributesLike:
			var associations = collection.GetAssociations()
			entry = v.getAttribute(associations, key, indices...)
		}
	}
	return entry
}

func (v *component_) RemoveSubcomponent(
	indices ...any,
) EntryLike {
	var entry EntryLike
	if len(indices) > 0 {
		var key = indices[0]
		indices = indices[1:]
		switch collection := v.GetEntity().(type) {
		case ItemsLike:
			switch entries := collection.GetEntries().(type) {
			case com.ListLike[EntryLike]:
				var index = key.(int)
				entry = v.removeItem(entries, index, indices...)
			default:
				var message = fmt.Sprintf(
					"Attempted to remove an item from a non-list type: %T",
					entries,
				)
				panic(message)
			}
		case AttributesLike:
			var associations = collection.GetAssociations()
			entry = v.removeAttribute(associations, key, indices...)
		}
	}
	return entry
}

// PROTECTED INTERFACE

// Private Methods

func (v *component_) getItem(
	entries com.Sequential[EntryLike],
	index int,
	indices ...any,
) EntryLike {
	var entry EntryLike
	var size = int(entries.GetSize())
	if size == 0 {
		// The list of entries is empty.
		return entry
	}
	if index < 0 {
		// Negative indices start from the end of the list.
		index = size + index + 1
	}
	if index > size {
		// The index is out of bounds.
		return entry
	}
	entry = entries.AsArray()[index-1]
	if len(indices) > 0 {
		var composite = entry.GetComposite()
		entry = composite.GetSubcomponent(indices...)
	}
	return entry
}

func (v *component_) setItem(
	entries com.ListLike[EntryLike],
	entry EntryLike,
	index int,
	indices ...any,
) {
	if index == 0 && len(indices) == 0 {
		// Append the attribute to the end of the list.
		entries.AppendValue(entry)
		return
	}
	var size = int(entries.GetSize())
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
		var composite = entries.GetValue(index).GetComposite()
		composite.SetSubcomponent(entry, indices...)
		return
	}
	entries.SetValue(index, entry)
}

func (v *component_) removeItem(
	entries com.ListLike[EntryLike],
	index int,
	indices ...any,
) EntryLike {
	var entry EntryLike
	var size = int(entries.GetSize())
	if size == 0 {
		// The list of entries is empty.
		return entry
	}
	if index < 0 {
		// Negative indices start from the end of the list.
		index = size + index + 1
	}
	if index > size {
		// The index is out of bounds.
		return entry
	}
	if len(indices) == 0 {
		entry = entries.GetValue(index)
		entries.RemoveValue(index)
		return entry
	}
	entry = entries.GetValue(index)
	var composite = entry.GetComposite()
	entry = composite.RemoveSubcomponent(indices...)
	return entry
}

func (v *component_) getAttribute(
	associations com.CatalogLike[any, EntryLike],
	key any,
	indices ...any,
) EntryLike {
	var entry EntryLike
	var first = fmt.Sprintf("%v", key)
	var iterator = associations.GetIterator()
	for iterator.HasNext() {
		var association = iterator.GetNext()
		var second = fmt.Sprintf("%v", association.GetKey())
		if first == second {
			entry = association.GetValue()
			if len(indices) > 0 {
				var composite = entry.GetComposite()
				entry = composite.GetSubcomponent(indices...)
			}
			break
		}
	}
	return entry
}

func (v *component_) setAttribute(
	associations com.CatalogLike[any, EntryLike],
	key any,
	entry EntryLike,
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
				composite.SetSubcomponent(entry, indices...)
			} else {
				association.SetValue(entry)
			}
			return
		}
	}
	associations.SetValue(key, entry)
}

func (v *component_) removeAttribute(
	associations com.CatalogLike[any, EntryLike],
	key any,
	indices ...any,
) EntryLike {
	var entry EntryLike
	var first = fmt.Sprintf("%v", key)
	var iterator = associations.GetIterator()
	for iterator.HasNext() {
		var association = iterator.GetNext()
		var second = fmt.Sprintf("%v", association.GetKey())
		if first == second {
			if len(indices) > 0 {
				entry = association.GetValue()
				var composite = entry.GetComposite()
				entry = composite.RemoveSubcomponent(indices...)
			} else {
				entry = associations.RemoveValue(key)
			}
			break
		}
	}
	return entry
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
