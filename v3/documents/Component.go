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
		parameter = componentClass().Component(
			constraint.GetType(),
			constraint.GetOptionalParameterization(),
		)
	}
	return parameter
}

func (v *component_) GetMember(
	indices ...any,
) MemberLike {
	var member MemberLike
	if len(indices) > 0 {
		var key = indices[0]
		indices = indices[1:]
		switch collection := v.GetEntity().(type) {
		case ItemsLike:
			var members = collection.GetMembers()
			var index = key.(uti.Index)
			member = v.getItem(members, index, indices...)
		case AttributesLike:
			var associations = collection.GetAssociations()
			member = v.getAttribute(associations, key, indices...)
		}
	}
	return member
}

func (v *component_) SetMember(
	class any,
	indices ...any,
) {
	var member MemberLike
	switch actual := class.(type) {
	case ComponentLike:
		member = MemberClass().Member(actual, "")
	case DocumentLike:
		member = MemberClass().Member(actual.GetComponent(), "")
	case MemberLike:
		member = actual
	default:
		member = MemberClass().Member(componentClass().Component(class, nil), "")
	}
	if uti.IsDefined(member) && len(indices) > 0 {
		var key = indices[0]
		indices = indices[1:]
		switch collection := v.GetEntity().(type) {
		case ItemsLike:
			var members = collection.GetMembers()
			var index = key.(uti.Index)
			v.setItem(members, member, index, indices...)
		case AttributesLike:
			var associations = collection.GetAssociations()
			v.setAttribute(associations, key, member, indices...)
		}
	}
}

func (v *component_) RemoveMember(
	indices ...any,
) {
	if len(indices) > 0 {
		var key = indices[0]
		indices = indices[1:]
		switch collection := v.GetEntity().(type) {
		case ItemsLike:
			var members = collection.GetMembers()
			var index = key.(uti.Index)
			v.removeItem(members, index, indices...)
		case AttributesLike:
			var associations = collection.GetAssociations()
			v.removeAttribute(associations, key, indices...)
		}
	}
}

// PROTECTED INTERFACE

// Private Methods

func (v *component_) getItem(
	members fra.ListLike[MemberLike],
	index uti.Index,
	indices ...any,
) MemberLike {
	var member MemberLike
	var size = uti.Index(members.GetSize())
	if size == 0 {
		// The list of members is empty.
		return member
	}
	if index < 0 {
		// Negative indices start from the end of the list.
		index = size + index + 1
	}
	if index > size {
		// The index is out of bounds.
		return member
	}
	member = members.GetValue(index)
	if len(indices) > 0 {
		var component = member.GetComponent()
		member = component.GetMember(indices...)
	}
	return member
}

func (v *component_) setItem(
	members fra.ListLike[MemberLike],
	member MemberLike,
	index uti.Index,
	indices ...any,
) {
	if index == 0 && len(indices) == 0 {
		// Append the attribute to the end of the list.
		members.AppendValue(member)
		return
	}
	var size = uti.Index(members.GetSize())
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
		var component = members.GetValue(index).GetComponent()
		component.SetMember(member, indices...)
		return
	}
	members.SetValue(uti.Index(index), member)
}

func (v *component_) removeItem(
	members fra.ListLike[MemberLike],
	index uti.Index,
	indices ...any,
) {
	var size = uti.Index(members.GetSize())
	if size == 0 {
		// The list of members is empty.
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
		members.RemoveValue(index)
		return
	}
	var member = members.GetValue(index)
	var component = member.GetComponent()
	component.RemoveMember(indices...)
}

func (v *component_) getAttribute(
	associations fra.CatalogLike[any, MemberLike],
	key any,
	indices ...any,
) MemberLike {
	var member MemberLike
	var first = fmt.Sprintf("%v", key)
	var iterator = associations.GetIterator()
	for iterator.HasNext() {
		var association = iterator.GetNext()
		var second = fmt.Sprintf("%v", association.GetKey())
		if first == second {
			member = association.GetValue()
			if len(indices) > 0 {
				var component = member.GetComponent()
				member = component.GetMember(indices...)
			}
			break
		}
	}
	return member
}

func (v *component_) setAttribute(
	associations fra.CatalogLike[any, MemberLike],
	key any,
	member MemberLike,
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
				component.SetMember(member, indices...)
			} else {
				association.SetValue(member)
			}
			break
		}
	}
}

func (v *component_) removeAttribute(
	associations fra.CatalogLike[any, MemberLike],
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
				var member = association.GetValue()
				var component = member.GetComponent()
				component.RemoveMember(indices...)
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
