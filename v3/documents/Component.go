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
	literal any,
	optionalGenerics GenericsLike,
	optionalNote string,
) ComponentLike {
	if uti.IsUndefined(literal) {
		panic("The \"literal\" attribute is required by this class.")
	}
	var instance = &component_{
		// Initialize the instance attributes.
		literal_:          literal,
		optionalGenerics_: optionalGenerics,
		optionalNote_:     optionalNote,
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

func (v *component_) GetLiteral() any {
	return v.literal_
}

func (v *component_) GetOptionalGenerics() GenericsLike {
	return v.optionalGenerics_
}

func (v *component_) GetOptionalNote() string {
	return v.optionalNote_
}

func (v *component_) GetParameter(
	symbol pri.SymbolLike,
) ComponentLike {
	var parameter ComponentLike
	var generics = v.GetOptionalGenerics()
	if uti.IsDefined(generics) {
		var parameters = generics.GetParameters()
		var constraint = parameters.GetValue(symbol)
		if uti.IsDefined(constraint) {
			parameter = componentClass().Component(
				constraint.GetEntity(),
				constraint.GetOptionalGenerics(),
				"",
			)
		}
	}
	return parameter
}

func (v *component_) SetSubcomponent(
	value any,
	indices ...any,
) {
	if len(indices) == 0 {
		panic("At least one index must be specified when setting a subcomponent.")
	}

	// Turn the value into a component.
	var component ComponentLike
	switch actual := value.(type) {
	case ComponentLike:
		component = actual
	case DocumentLike:
		component = actual.GetComponent()
	default:
		component = componentClass().Component(value, nil, "")
	}

	// Treat items and attributes differently.
	switch collection := v.GetLiteral().(type) {
	case ItemsLike:
		var components = collection.GetComponents().(com.ListLike[ComponentLike])
		v.setItem(components, component, indices...)
	case AttributesLike:
		var associations = collection.GetAssociations()
		v.setAttribute(associations, component, indices...)
	default:
		var message = fmt.Sprintf(
			"Attempted to set a subcomponent in a simple type: %T",
			collection,
		)
		panic(message)
	}
}

func (v *component_) GetSubcomponent(
	indices ...any,
) ComponentLike {
	var component ComponentLike
	if len(indices) == 0 {
		panic("At least one index must be specified when setting a subcomponent.")
	}
	switch collection := v.GetLiteral().(type) {
	case ItemsLike:
		var components = collection.GetComponents()
		component = v.getItem(components, indices...)
	case AttributesLike:
		var associations = collection.GetAssociations()
		component = v.getAttribute(associations, indices...)
	}
	return component
}

func (v *component_) RemoveSubcomponent(
	indices ...any,
) ComponentLike {
	var component ComponentLike
	if len(indices) > 0 {
		switch collection := v.GetLiteral().(type) {
		case ItemsLike:
			switch components := collection.GetComponents().(type) {
			case com.ListLike[ComponentLike]:
				component = v.removeItem(components, indices...)
			default:
				var message = fmt.Sprintf(
					"Attempted to remove an item from a non-list type: %T",
					components,
				)
				panic(message)
			}
		case AttributesLike:
			var associations = collection.GetAssociations()
			component = v.removeAttribute(associations, indices...)
		}
	}
	return component
}

// PROTECTED INTERFACE

// Private Methods

func (v *component_) getItem(
	components com.Sequential[ComponentLike],
	indices ...any,
) ComponentLike {
	var subcomponent ComponentLike

	// Calculate the index and size constraints.
	var size = int(components.GetSize())
	var index = indices[0].(int)
	if index < 0 {
		// Negative indices start from the end of the list.
		index = size + index + 1
	}
	indices = indices[1:]

	// Check for an out of bounds index.
	if index == 0 || index > size {
		return subcomponent
	}

	// Retrieve the subcomponent.
	subcomponent = components.AsArray()[index-1]
	if len(indices) > 0 {
		subcomponent = subcomponent.GetSubcomponent(indices...)
	}

	return subcomponent
}

func (v *component_) setItem(
	components com.ListLike[ComponentLike],
	component ComponentLike,
	indices ...any,
) {
	// Calculate the index and size constraints.
	var size = int(components.GetSize())
	var index = indices[0].(int)
	if index < 0 {
		// Negative indices start from the end of the list.
		index = size + index + 1
	}
	indices = indices[1:]

	// Check for an out of bounds index.
	if index == 0 || index > size {
		components.AppendValue(component)
		return
	}

	// Update the subcomponent.
	if len(indices) > 0 {
		var subcomponent = components.GetValue(index)
		subcomponent.SetSubcomponent(component, indices...)
		return
	}
	components.SetValue(index, component)
}

func (v *component_) removeItem(
	components com.ListLike[ComponentLike],
	indices ...any,
) ComponentLike {
	var subcomponent ComponentLike

	// Calculate the index and size constraints.
	var size = int(components.GetSize())
	var index = indices[0].(int)
	if index < 0 {
		// Negative indices start from the end of the list.
		index = size + index + 1
	}
	indices = indices[1:]

	// Check for an out of bounds index.
	if index == 0 || index > size {
		return subcomponent
	}

	// Remove the subcomponent.
	subcomponent = components.GetValue(index)
	if len(indices) > 0 {
		subcomponent = subcomponent.RemoveSubcomponent(indices...)
		return subcomponent
	}
	components.RemoveValue(index)

	return subcomponent
}

func (v *component_) getAttribute(
	associations com.CatalogLike[any, ComponentLike],
	indices ...any,
) ComponentLike {
	var subcomponent ComponentLike

	// Calculate the key and size constraints.
	var size = int(associations.GetSize())
	var key = indices[0]
	indices = indices[1:]
	if size == 0 {
		return subcomponent
	}

	// Retrieve the subcomponent.
	subcomponent = associations.GetValue(key)
	if len(indices) > 0 {
		subcomponent = subcomponent.GetSubcomponent(indices...)
	}

	return subcomponent
}

func (v *component_) setAttribute(
	associations com.CatalogLike[any, ComponentLike],
	component ComponentLike,
	indices ...any,
) {
	// Calculate the key and size constraints.
	var size = int(associations.GetSize())
	var key = indices[0]
	indices = indices[1:]
	if size == 0 {
		associations.SetValue(key, component)
		return
	}

	// Update the subcomponent.
	if len(indices) > 0 {
		var subcomponent = associations.GetValue(key)
		subcomponent.SetSubcomponent(component, indices...)
		return
	}

	associations.SetValue(key, component)
}

func (v *component_) removeAttribute(
	associations com.CatalogLike[any, ComponentLike],
	indices ...any,
) ComponentLike {
	var subcomponent ComponentLike

	// Calculate the key and size constraints.
	var size = int(associations.GetSize())
	var key = indices[0]
	indices = indices[1:]
	if size == 0 {
		return subcomponent
	}

	// Remove the subcomponent.
	subcomponent = associations.GetValue(key)
	if len(indices) > 0 {
		subcomponent = subcomponent.RemoveSubcomponent(indices...)
		return subcomponent
	}
	associations.RemoveValue(key)

	return subcomponent
}

// Instance Structure

type component_ struct {
	// Declare the instance attributes.
	literal_          any
	optionalGenerics_ GenericsLike
	optionalNote_     string
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
