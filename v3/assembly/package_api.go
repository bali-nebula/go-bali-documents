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
Package "assembly" provides the class definitions for each part of in "inflated"
Bali Assembly™.  Inflated assembly can be manipulated by the agents defined in
this project.

For detailed documentation on this package refer to the wiki:
  - https://github.com/bali-nebula/go-bali-documents/wiki

This package follows the Crater Dog Technologies™ Go Coding Conventions located
here:
  - https://github.com/craterdog/go-development-tools/wiki/Coding-Conventions

Additional concrete implementations of the classes declared by this package can
be developed and used seamlessly since the interface declarations only depend on
other interfaces and intrinsic types—and the class implementations only depend
on interfaces, not on each other.
*/
package assembly

import (
	fra "github.com/craterdog/go-component-framework/v7"
)

// TYPE DECLARATIONS

/*
Operation is an enumerated type representing the possible BVM operations.
*/
type Operation uint8

const (
	JumpOperation Operation = iota
	PushOperation
	PullOperation
	LoadOperation
	SaveOperation
	DropOperation
	CallOperation
	SendOperation
)

/*
Modifier is an enumerated type representing the possible BVM modifiers.
*/
type Modifier uint8

const (
	OnAnyModifier Modifier = iota
	OnNoneModifier
	OnFalseModifier
	OnEmptyModifier
	LiteralModifier
	ConstantModifier
	ArgumentModifier
	HandlerModifier
	ComponentModifier
	ResultModifier
	ExceptionModifier
	DraftModifier
	ContractModifier
	VariableModifier
	MessageModifier
	With0Modifier
	With1Modifier
	With2Modifier
	With3Modifier
	ComponentWithModifier
	ContractWithModifier
)

/*
Operand is a constrained type representing the possible BVM operands.
*/
type Operand uint16

// FUNCTIONAL DECLARATIONS

// CLASS DECLARATIONS

/*
BytecodeClassLike is a class interface that declares the complete set of
class constructors, constants and functions that must be supported by each
concrete bytecode-like class.
*/
type BytecodeClassLike interface {
	// Constructor Methods
	Bytecode(
		instructions []InstructionLike,
	) BytecodeLike
	BytecodeFromString(
		source string,
	) BytecodeLike
}

/*
InstructionClassLike is a class interface that declares the complete set of
class constructors, constants and functions that must be supported by each
concrete instruction-like class.
*/
type InstructionClassLike interface {
	// Constructor Methods
	Instruction(
		operation Operation,
		modifier Modifier,
		operand Operand,
	) InstructionLike
	InstructionFromInteger(
		integer uint16,
	) InstructionLike

	// Function Methods
	FormatInstructions(
		instructions fra.Sequential[InstructionLike],
	) string
}

// INSTANCE DECLARATIONS

/*
BytecodeLike is an instance interface that declares the complete set of
principal, attribute and aspect methods that must be supported by each instance
of a concrete bytecode-like class.
*/
type BytecodeLike interface {
	// Principal Methods
	GetClass() BytecodeClassLike
	AsString() string
	AsIntrinsic() []InstructionLike

	// Aspect Interfaces
	fra.Sequential[InstructionLike]
}

/*
InstructionLike is an instance interface that declares the complete set of
principal, attribute and aspect methods that must be supported by each instance
of a concrete instruction-like class.
*/
type InstructionLike interface {
	// Principal Methods
	GetClass() InstructionClassLike
	AsIntrinsic() uint16
	AsString() string
	OperationAsString() string
	ModifierAsString() string
	OperandAsString() string

	// Attribute Methods
	GetOperation() Operation
	GetModifier() Modifier
	GetOperand() Operand
}

// ASPECT DECLARATIONS
