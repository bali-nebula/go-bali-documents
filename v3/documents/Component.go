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
	fra "github.com/craterdog/go-collection-framework/v8"
	ele "github.com/craterdog/go-essential-elements/v8"
	uti "github.com/craterdog/go-missing-utilities/v8"
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

func (v *component_) GetParameter(
	symbol ele.SymbolLike,
) ComponentLike {
	var parameter ComponentLike
	var generics = v.GetOptionalGenerics()
	if uti.IsDefined(generics) {
		var parameters = generics.GetParameters()
		var constraint = parameters.GetValue(symbol)
		if uti.IsDefined(constraint) {
			parameter = componentClass().Component(
				constraint.GetMetadata(),
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
	var composite CompositeLike
	switch actual := value.(type) {
	case ComponentLike:
		composite = CompositeClass().Composite(actual, "")
	case DocumentLike:
		composite = CompositeClass().Composite(actual.GetComponent(), "")
	case CompositeLike:
		composite = actual
	default:
		composite = CompositeClass().Composite(
			componentClass().Component(value, nil), "",
		)
	}
	if uti.IsDefined(composite) && len(indices) > 0 {
		var key = indices[0]
		indices = indices[1:]
		switch collection := v.GetEntity().(type) {
		case ItemsLike:
			switch composites := collection.GetComposites().(type) {
			case fra.ListLike[CompositeLike]:
				var index = key.(int)
				v.setItem(composites, composite, index, indices...)
			default:
				var message = fmt.Sprintf(
					"Attempted to set an item in a non-list type: %T",
					composites,
				)
				panic(message)
			}
		case AttributesLike:
			var associations = collection.GetAssociations()
			v.setAttribute(associations, key, composite, indices...)
		}
	}
}

func (v *component_) GetSubcomponent(
	indices ...any,
) CompositeLike {
	var composite CompositeLike
	if len(indices) > 0 {
		var key = indices[0]
		indices = indices[1:]
		switch collection := v.GetEntity().(type) {
		case ItemsLike:
			var composites = collection.GetComposites()
			var index = key.(int)
			composite = v.getItem(composites, index, indices...)
		case AttributesLike:
			var associations = collection.GetAssociations()
			composite = v.getAttribute(associations, key, indices...)
		}
	}
	return composite
}

func (v *component_) RemoveSubcomponent(
	indices ...any,
) CompositeLike {
	var composite CompositeLike
	if len(indices) > 0 {
		var key = indices[0]
		indices = indices[1:]
		switch collection := v.GetEntity().(type) {
		case ItemsLike:
			switch composites := collection.GetComposites().(type) {
			case fra.ListLike[CompositeLike]:
				var index = key.(int)
				composite = v.removeItem(composites, index, indices...)
			default:
				var message = fmt.Sprintf(
					"Attempted to remove an item from a non-list type: %T",
					composites,
				)
				panic(message)
			}
		case AttributesLike:
			var associations = collection.GetAssociations()
			composite = v.removeAttribute(associations, key, indices...)
		}
	}
	return composite
}

// Attribute Methods

func (v *component_) GetEntity() any {
	return v.entity_
}

func (v *component_) GetOptionalGenerics() GenericsLike {
	return v.optionalGenerics_
}

// PROTECTED INTERFACE

// Private Methods

func (v *component_) getItem(
	composites fra.Sequential[CompositeLike],
	index int,
	indices ...any,
) CompositeLike {
	var composite CompositeLike
	var size = int(composites.GetSize())
	if size == 0 {
		// The list of composites is empty.
		return composite
	}
	if index < 0 {
		// Negative indices start from the end of the list.
		index = size + index + 1
	}
	if index > size {
		// The index is out of bounds.
		return composite
	}
	composite = composites.AsArray()[index-1]
	if len(indices) > 0 {
		var component = composite.GetComponent()
		composite = component.GetSubcomponent(indices...)
	}
	return composite
}

func (v *component_) setItem(
	composites fra.ListLike[CompositeLike],
	composite CompositeLike,
	index int,
	indices ...any,
) {
	if index == 0 && len(indices) == 0 {
		// Append the attribute to the end of the list.
		composites.AppendValue(composite)
		return
	}
	var size = int(composites.GetSize())
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
		var component = composites.GetValue(index).GetComponent()
		component.SetSubcomponent(composite, indices...)
		return
	}
	composites.SetValue(index, composite)
}

func (v *component_) removeItem(
	composites fra.ListLike[CompositeLike],
	index int,
	indices ...any,
) CompositeLike {
	var composite CompositeLike
	var size = int(composites.GetSize())
	if size == 0 {
		// The list of composites is empty.
		return composite
	}
	if index < 0 {
		// Negative indices start from the end of the list.
		index = size + index + 1
	}
	if index > size {
		// The index is out of bounds.
		return composite
	}
	if len(indices) == 0 {
		composite = composites.GetValue(index)
		composites.RemoveValue(index)
		return composite
	}
	composite = composites.GetValue(index)
	var component = composite.GetComponent()
	composite = component.RemoveSubcomponent(indices...)
	return composite
}

func (v *component_) getAttribute(
	associations fra.CatalogLike[any, CompositeLike],
	key any,
	indices ...any,
) CompositeLike {
	var composite CompositeLike
	var first = fmt.Sprintf("%v", key)
	var iterator = associations.GetIterator()
	for iterator.HasNext() {
		var association = iterator.GetNext()
		var second = fmt.Sprintf("%v", association.GetKey())
		if first == second {
			composite = association.GetValue()
			if len(indices) > 0 {
				var component = composite.GetComponent()
				composite = component.GetSubcomponent(indices...)
			}
			break
		}
	}
	return composite
}

func (v *component_) setAttribute(
	associations fra.CatalogLike[any, CompositeLike],
	key any,
	composite CompositeLike,
	indices ...any,
) {
	var first = fmt.Sprintf("%v", key)
	var iterator = associations.GetIterator()
	for iterator.HasNext() {
		var association = iterator.GetNext()
		var second = fmt.Sprintf("%v", association.GetKey())
		if first == second {
			if len(indices) > 0 {
				var component = association.GetValue().GetComponent()
				component.SetSubcomponent(composite, indices...)
			} else {
				association.SetValue(composite)
			}
			return
		}
	}
	associations.SetValue(key, composite)
}

func (v *component_) removeAttribute(
	associations fra.CatalogLike[any, CompositeLike],
	key any,
	indices ...any,
) CompositeLike {
	var composite CompositeLike
	var first = fmt.Sprintf("%v", key)
	var iterator = associations.GetIterator()
	for iterator.HasNext() {
		var association = iterator.GetNext()
		var second = fmt.Sprintf("%v", association.GetKey())
		if first == second {
			if len(indices) > 0 {
				composite = association.GetValue()
				var component = composite.GetComponent()
				composite = component.RemoveSubcomponent(indices...)
			} else {
				composite = associations.RemoveValue(key)
			}
			break
		}
	}
	return composite
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
