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

func ComponentClass() ComponentClassLike {
	return componentClass()
}

// Constructor Methods

func (c *componentClass_) Component(
	entity any,
	optionalParameterization ParameterizationLike,
) ComponentLike {
	if uti.IsUndefined(entity) {
		panic("The \"entity\" attribute is required by this class.")
	}
	var instance = &component_{
		// Initialize the instance attributes.
		entity_:                   entity,
		optionalParameterization_: optionalParameterization,
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

// Attribute Methods

func (v *component_) GetEntity() any {
	return v.entity_
}

func (v *component_) GetOptionalParameterization() ParameterizationLike {
	return v.optionalParameterization_
}

// Declarative Methods

func (v *component_) GetParameter(
	symbol fra.SymbolLike,
) ComponentLike {
	var parameter ComponentLike
	var parameterization = v.GetOptionalParameterization()
	if uti.IsDefined(parameterization) {
		var parameters = parameterization.GetParameters()
		var constraint = parameters.GetValue(symbol)
		if uti.IsDefined(constraint) {
			parameter = componentClass().Component(
				constraint.GetMetadata(),
				constraint.GetOptionalParameterization(),
			)
		}
	}
	return parameter
}

func (v *component_) GetObject(
	indices ...any,
) ObjectLike {
	var object ObjectLike
	if len(indices) > 0 {
		var key = indices[0]
		indices = indices[1:]
		switch collection := v.GetEntity().(type) {
		case ItemsLike:
			var objects = collection.GetObjects()
			var index = key.(uti.Index)
			object = v.getItem(objects, index, indices...)
		case AttributesLike:
			var associations = collection.GetAssociations()
			object = v.getAttribute(associations, key, indices...)
		}
	}
	return object
}

func (v *component_) SetObject(
	class any,
	indices ...any,
) {
	var object ObjectLike
	switch actual := class.(type) {
	case ComponentLike:
		object = ObjectClass().Object(actual, "")
	case DocumentLike:
		object = ObjectClass().Object(actual.GetComponent(), "")
	case ObjectLike:
		object = actual
	default:
		object = ObjectClass().Object(componentClass().Component(class, nil), "")
	}
	if uti.IsDefined(object) && len(indices) > 0 {
		var key = indices[0]
		indices = indices[1:]
		switch collection := v.GetEntity().(type) {
		case ItemsLike:
			switch objects := collection.GetObjects().(type) {
			case fra.ListLike[ObjectLike]:
				var index = key.(uti.Index)
				v.setItem(objects, object, index, indices...)
			default:
				var message = fmt.Sprintf(
					"Attempted to set an item in a non-list type: %T",
					objects,
				)
				panic(message)
			}
		case AttributesLike:
			var associations = collection.GetAssociations()
			v.setAttribute(associations, key, object, indices...)
		}
	}
}

func (v *component_) RemoveObject(
	indices ...any,
) {
	if len(indices) > 0 {
		var key = indices[0]
		indices = indices[1:]
		switch collection := v.GetEntity().(type) {
		case ItemsLike:
			switch objects := collection.GetObjects().(type) {
			case fra.ListLike[ObjectLike]:
				var index = key.(uti.Index)
				v.removeItem(objects, index, indices...)
			default:
				var message = fmt.Sprintf(
					"Attempted to remove an item from a non-list type: %T",
					objects,
				)
				panic(message)
			}
		case AttributesLike:
			var associations = collection.GetAssociations()
			v.removeAttribute(associations, key, indices...)
		}
	}
}

// PROTECTED INTERFACE

// Private Methods

func (v *component_) getItem(
	objects fra.Sequential[ObjectLike],
	index uti.Index,
	indices ...any,
) ObjectLike {
	var object ObjectLike
	var size = uti.Index(objects.GetSize())
	if size == 0 {
		// The list of objects is empty.
		return object
	}
	if index < 0 {
		// Negative indices start from the end of the list.
		index = size + index + 1
	}
	if index > size {
		// The index is out of bounds.
		return object
	}
	object = objects.AsArray()[index-1]
	if len(indices) > 0 {
		var component = object.GetComponent()
		object = component.GetObject(indices...)
	}
	return object
}

func (v *component_) setItem(
	objects fra.ListLike[ObjectLike],
	object ObjectLike,
	index uti.Index,
	indices ...any,
) {
	if index == 0 && len(indices) == 0 {
		// Append the attribute to the end of the list.
		objects.AppendValue(object)
		return
	}
	var size = uti.Index(objects.GetSize())
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
		var component = objects.GetValue(index).GetComponent()
		component.SetObject(object, indices...)
		return
	}
	objects.SetValue(uti.Index(index), object)
}

func (v *component_) removeItem(
	objects fra.ListLike[ObjectLike],
	index uti.Index,
	indices ...any,
) {
	var size = uti.Index(objects.GetSize())
	if size == 0 {
		// The list of objects is empty.
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
		objects.RemoveValue(index)
		return
	}
	var object = objects.GetValue(index)
	var component = object.GetComponent()
	component.RemoveObject(indices...)
}

func (v *component_) getAttribute(
	associations fra.CatalogLike[any, ObjectLike],
	key any,
	indices ...any,
) ObjectLike {
	var object ObjectLike
	var first = fmt.Sprintf("%v", key)
	var iterator = associations.GetIterator()
	for iterator.HasNext() {
		var association = iterator.GetNext()
		var second = fmt.Sprintf("%v", association.GetKey())
		if first == second {
			object = association.GetValue()
			if len(indices) > 0 {
				var component = object.GetComponent()
				object = component.GetObject(indices...)
			}
			break
		}
	}
	return object
}

func (v *component_) setAttribute(
	associations fra.CatalogLike[any, ObjectLike],
	key any,
	object ObjectLike,
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
				component.SetObject(object, indices...)
			} else {
				association.SetValue(object)
			}
			return
		}
	}
	associations.SetValue(key, object)
}

func (v *component_) removeAttribute(
	associations fra.CatalogLike[any, ObjectLike],
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
				var object = association.GetValue()
				var component = object.GetComponent()
				component.RemoveObject(indices...)
			} else {
				associations.RemoveValue(key)
			}
			break
		}
	}
}

// Instance Structure

type component_ struct {
	// Declare the instance attributes.
	entity_                   any
	optionalParameterization_ ParameterizationLike
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
