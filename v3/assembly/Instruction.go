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

package assembly

import (
	fmt "fmt"
	fra "github.com/craterdog/go-component-framework/v7"
	uti "github.com/craterdog/go-missing-utilities/v7"
	sts "strings"
)

// CLASS INTERFACE

// Access Function

func InstructionClass() InstructionClassLike {
	return instructionClass()
}

// Constructor Methods

func (c *instructionClass_) Instruction(
	operation Operation,
	modifier Modifier,
	operand Operand,
) InstructionLike {
	var instruction = c.encodeInstruction(operation, modifier, operand)
	return instruction_(instruction)
}

func (c *instructionClass_) InstructionFromInteger(
	integer uint16,
) InstructionLike {
	return instruction_(integer)
}

// Constant Methods

// Function Methods

func (c *instructionClass_) FormatInstructions(
	instructions fra.Sequential[InstructionLike],
) string {
	var result sts.Builder
	result.WriteString(
		`Address   Bytes   Instruction                Mnemonic
--------------------------------------------------------------------
`,
	)
	var iterator = instructions.GetIterator()
	for iterator.HasNext() {
		var instruction = iterator.GetNext()
		var address = fmt.Sprintf("[x%03x]", iterator.GetSlot())
		var bytes = fmt.Sprintf("x%04x", instruction.AsIntrinsic())
		var operation = (instruction.AsIntrinsic() ^ c.operationMask_) >> 13
		var modifier = (instruction.AsIntrinsic() ^ c.modifierMask_) >> 11
		var operand = instruction.OperandAsString()
		for len(operand) < 5 {
			operand = " " + operand
		}
		if len(operand) < 6 {
			operand += " "
		}
		var bytecode = fmt.Sprintf("%d %d %s", operation, modifier, operand)
		var description = instruction.AsString()
		var line = fmt.Sprintf(
			"%s:   %s   %s   %s\n",
			address,
			bytes,
			bytecode,
			description,
		)
		result.WriteString(line)
	}
	return result.String()
}

// INSTANCE INTERFACE

// Principal Methods

func (v instruction_) GetClass() InstructionClassLike {
	return instructionClass()
}

func (v instruction_) AsIntrinsic() uint16 {
	return uint16(v)
}

func (v instruction_) AsString() string {
	var instruction = v.OperationAsString()
	switch instruction {
	case "SKIP":
	case "JUMP":
		instruction += " TO " + v.OperandAsString()
		var modifier = v.ModifierAsString()
		if uti.IsDefined(modifier) {
			instruction += " " + modifier
		}
	case "CALL", "SEND":
		instruction += " " + v.OperandAsString()
		var modifier = v.ModifierAsString()
		if uti.IsDefined(modifier) {
			instruction += " " + modifier
		}
	case "PULL":
		instruction += " " + v.ModifierAsString()
	case "PUSH", "LOAD", "SAVE", "DROP":
		instruction += " " + v.ModifierAsString()
		instruction += " " + v.OperandAsString()
	}
	return instruction
}

func (v instruction_) OperationAsString() string {
	if v.AsIntrinsic() == 0 {
		return "SKIP"
	}
	var operation = v.GetOperation()
	switch operation {
	case JumpOperation:
		return "JUMP"
	case PushOperation:
		return "PUSH"
	case PullOperation:
		return "PULL"
	case LoadOperation:
		return "LOAD"
	case SaveOperation:
		return "SAVE"
	case DropOperation:
		return "DROP"
	case CallOperation:
		return "CALL"
	case SendOperation:
		return "SEND"
	default:
		var message = fmt.Sprintf(
			"Found an unknown operation type: %v",
			operation,
		)
		panic(message)
	}
}

func (v instruction_) ModifierAsString() string {
	var operation = v.GetOperation()
	var modifier = v.GetModifier()
	switch operation {
	case JumpOperation:
		switch modifier {
		case OnAnyModifier:
			return ""
		case OnEmptyModifier:
			return "ON EMPTY"
		case OnNoneModifier:
			return "ON NONE"
		case OnFalseModifier:
			return "ON FALSE"
		}
	case PushOperation:
		switch modifier {
		case LiteralModifier:
			return "LITERAL"
		case ConstantModifier:
			return "CONSTANT"
		case ArgumentModifier:
			return "ARGUMENT"
		case HandlerModifier:
			return "HANDLER"
		}
	case PullOperation:
		switch modifier {
		case ComponentModifier:
			return "COMPONENT"
		case ResultModifier:
			return "RESULT"
		case ExceptionModifier:
			return "EXCEPTION"
		case HandlerModifier:
			return "HANDLER"
		}
	case LoadOperation, SaveOperation, DropOperation:
		switch modifier {
		case DraftModifier:
			return "DRAFT"
		case ContractModifier:
			return "CONTRACT"
		case VariableModifier:
			return "VARIABLE"
		case MessageModifier:
			return "MESSAGE"
		}
	case CallOperation:
		switch modifier {
		case With0Modifier:
			return ""
		case With1Modifier:
			return "WITH 1 ARGUMENT"
		case With2Modifier:
			return "WITH 2 ARGUMENTS"
		case With3Modifier:
			return "WITH 3 ARGUMENTS"
		}
	case SendOperation:
		switch modifier {
		case ComponentModifier:
			return "TO COMPONENT"
		case ContractModifier:
			return "TO CONTRACT"
		case ComponentWithModifier:
			return "TO COMPONENT WITH ARGUMENTS"
		case ContractWithModifier:
			return "TO CONTRACT WITH ARGUMENTS"
		}
	default:
		var message = fmt.Sprintf(
			"Found an unknown operation type: %v",
			operation,
		)
		panic(message)
	}
	var message = fmt.Sprintf(
		"Found an unknown modifier type: %v",
		modifier,
	)
	panic(message)
}

func (v instruction_) OperandAsString() string {
	var operation = v.GetOperation()
	var modifier = v.GetModifier()
	var operand = v.GetOperand()
	if operation == JumpOperation ||
		(operation == PushOperation && modifier == HandlerModifier) {
		// Treat the operand as an address "[xHEX]".
		return fmt.Sprintf("[x%03x]", operand)
	} else {
		// Treat the operand as an index "DECI".
		return fmt.Sprintf("%d", operand)
	}
}

// Attribute Methods

func (v instruction_) GetOperation() Operation {
	return Operation(v.AsIntrinsic() & instructionClass().operationMask_)
}

func (v instruction_) GetModifier() Modifier {
	return Modifier(v.AsIntrinsic() & instructionClass().modifierMask_)
}

func (v instruction_) GetOperand() Operand {
	return Operand(v.AsIntrinsic() & instructionClass().operandMask_)
}

// PROTECTED INTERFACE

// Private Methods

func (c instructionClass_) encodeInstruction(
	operation Operation,
	modifier Modifier,
	operand Operand,
) uint16 {
	var instruction uint16
	switch operation {
	case JumpOperation:
		switch modifier {
		case OnAnyModifier:
			instruction = c.jumpOperation_ ^ c.onAny_
		case OnNoneModifier:
			instruction = c.jumpOperation_ ^ c.onNone_
		case OnFalseModifier:
			instruction = c.jumpOperation_ ^ c.onFalse_
		case OnEmptyModifier:
			instruction = c.jumpOperation_ ^ c.onEmpty_
		default:
			var message = fmt.Sprintf(
				"An invalid modifier was passed for a JUMP operation: %v",
				modifier,
			)
			panic(message)
		}
	case PushOperation:
		switch modifier {
		case LiteralModifier:
			instruction = c.pushOperation_ ^ c.literal_
		case ConstantModifier:
			instruction = c.pushOperation_ ^ c.constant_
		case ArgumentModifier:
			instruction = c.pushOperation_ ^ c.argument_
		case HandlerModifier:
			instruction = c.pushOperation_ ^ c.handler_
		default:
			var message = fmt.Sprintf(
				"An invalid modifier was passed for a PUSH operation: %v",
				modifier,
			)
			panic(message)
		}
	case PullOperation:
		switch modifier {
		case ComponentModifier:
			instruction = c.pullOperation_ ^ c.component_
		case ResultModifier:
			instruction = c.pullOperation_ ^ c.result_
		case ExceptionModifier:
			instruction = c.pullOperation_ ^ c.exception_
		case HandlerModifier:
			instruction = c.pullOperation_ ^ c.handler_
		default:
			var message = fmt.Sprintf(
				"An invalid modifier was passed for a PULL operation: %v",
				modifier,
			)
			panic(message)
		}
	case LoadOperation:
		switch modifier {
		case DraftModifier:
			instruction = c.loadOperation_ ^ c.draft_
		case ContractModifier:
			instruction = c.loadOperation_ ^ c.contract_
		case VariableModifier:
			instruction = c.loadOperation_ ^ c.variable_
		case MessageModifier:
			instruction = c.loadOperation_ ^ c.message_
		default:
			var message = fmt.Sprintf(
				"An invalid modifier was passed for a LOAD operation: %v",
				modifier,
			)
			panic(message)
		}
	case SaveOperation:
		switch modifier {
		case DraftModifier:
			instruction = c.saveOperation_ ^ c.draft_
		case ContractModifier:
			instruction = c.saveOperation_ ^ c.contract_
		case VariableModifier:
			instruction = c.saveOperation_ ^ c.variable_
		case MessageModifier:
			instruction = c.saveOperation_ ^ c.message_
		default:
			var message = fmt.Sprintf(
				"An invalid modifier was passed for a SAVE operation: %v",
				modifier,
			)
			panic(message)
		}
	case DropOperation:
		switch modifier {
		case DraftModifier:
			instruction = c.dropOperation_ ^ c.draft_
		case ContractModifier:
			instruction = c.dropOperation_ ^ c.contract_
		case VariableModifier:
			instruction = c.dropOperation_ ^ c.variable_
		case MessageModifier:
			instruction = c.dropOperation_ ^ c.message_
		default:
			var message = fmt.Sprintf(
				"An invalid modifier was passed for a DROP operation: %v",
				modifier,
			)
			panic(message)
		}
	case CallOperation:
		switch modifier {
		case With0Modifier:
			instruction = c.callOperation_ ^ c.with0Arguments_
		case With1Modifier:
			instruction = c.callOperation_ ^ c.with1Argument_
		case With2Modifier:
			instruction = c.callOperation_ ^ c.with2Arguments_
		case With3Modifier:
			instruction = c.callOperation_ ^ c.with3Arguments_
		default:
			var message = fmt.Sprintf(
				"An invalid modifier was passed for a CALL operation: %v",
				modifier,
			)
			panic(message)
		}
	case SendOperation:
		switch modifier {
		case ComponentModifier:
			instruction = c.sendOperation_ ^ c.component_
		case ContractModifier:
			instruction = c.sendOperation_ ^ c.contract_
		case ComponentWithModifier:
			instruction = c.sendOperation_ ^ c.componentWithArguments_
		case ContractWithModifier:
			instruction = c.sendOperation_ ^ c.contractWithArguments_
		default:
			var message = fmt.Sprintf(
				"An invalid modifier was passed for a SEND operation: %v",
				modifier,
			)
			panic(message)
		}
	default:
		var message = fmt.Sprintf(
			"An invalid modifier was passed for a SEND operation: %v",
			modifier,
		)
		panic(message)
	}
	return instruction ^ (uint16(operand) & c.operandMask_)
}

// Instance Structure

type instruction_ uint16

// Class Structure

type instructionClass_ struct {
	// Declare the class constants.

	// Masks
	operationMask_ uint16
	modifierMask_  uint16
	operandMask_   uint16

	// Operations
	jumpOperation_ uint16
	pushOperation_ uint16
	pullOperation_ uint16
	loadOperation_ uint16
	saveOperation_ uint16
	dropOperation_ uint16
	callOperation_ uint16
	sendOperation_ uint16

	// JUMP modifiers
	onAny_   uint16
	onEmpty_ uint16
	onNone_  uint16
	onFalse_ uint16

	// PUSH modifiers
	literal_  uint16
	constant_ uint16
	argument_ uint16
	handler_  uint16

	// PULL modifiers
	component_ uint16
	result_    uint16
	exception_ uint16

	// LOAD, SAVE and DROP modifiers
	draft_    uint16
	contract_ uint16
	variable_ uint16
	message_  uint16

	// CALL modifiers
	with0Arguments_ uint16
	with1Argument_  uint16
	with2Arguments_ uint16
	with3Arguments_ uint16

	// SEND modifiers
	componentWithArguments_ uint16
	contractWithArguments_  uint16
}

// Class Reference

func instructionClass() *instructionClass_ {
	return instructionClassReference_
}

var instructionClassReference_ = &instructionClass_{
	// Initialize the class constants.

	// Masks
	operationMask_: 0b1110000000000000,
	modifierMask_:  0b0001100000000000,
	operandMask_:   0b0000011111111111,

	// Operations
	jumpOperation_: 0b0000000000000000,
	pushOperation_: 0b0010000000000000,
	pullOperation_: 0b0100000000000000,
	loadOperation_: 0b0110000000000000,
	saveOperation_: 0b1000000000000000,
	dropOperation_: 0b1010000000000000,
	callOperation_: 0b1100000000000000,
	sendOperation_: 0b1110000000000000,

	// JUMP modifiers
	onAny_:   0b0000000000000000,
	onNone_:  0b0000100000000000,
	onFalse_: 0b0001000000000000,
	onEmpty_: 0b0001100000000000,

	// PUSH modifiers
	literal_:  0b0000000000000000,
	constant_: 0b0000100000000000,
	argument_: 0b0001000000000000,
	handler_:  0b0001100000000000,

	// PULL modifiers
	component_: 0b0000000000000000,
	result_:    0b0000100000000000,
	exception_: 0b0001000000000000,

	// LOAD, SAVE and DROP modifiers
	draft_:    0b0000000000000000,
	contract_: 0b0000100000000000,
	variable_: 0b0001000000000000,
	message_:  0b0001100000000000,

	// CALL modifiers
	with0Arguments_: 0b0000000000000000,
	with1Argument_:  0b0000100000000000,
	with2Arguments_: 0b0001000000000000,
	with3Arguments_: 0b0001100000000000,

	// SEND modifiers
	componentWithArguments_: 0b0001000000000000,
	contractWithArguments_:  0b0001100000000000,
}
