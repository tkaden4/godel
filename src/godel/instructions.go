package main

import (
	"errors"
	"fmt"
)

func Noop(vm VM) error {
	return nil
}

func Halt(vm VM) error {
	vm.Halt()
	return nil
}

func PutImmediate(vm VM) error {
	register, _ := ReadImmediateUint8(vm)
	value, _ := ReadImmediateUint16(vm)
	vm.Registers().SetRegister(register, value)
	return nil
}

func Cout(vm VM) error {
	register, _ := ReadImmediateUint8(vm)
	value, _ := vm.Registers().GetRegister(register)
	fmt.Print(string(value))
	return nil
}

func Dispatch(vm VM) error {
	instruction, _ := ReadImmediateUint8(vm)
	switch instruction {
	case 0x00:
		return Halt(vm)
	case 0x01:
		return Noop(vm)
	case 0x02:
		return PutImmediate(vm)
	case 0x05:
		return Cout(vm)
	default:
		return errors.New("Invalid opcode")
	}
}

func Run(vm VM) error {
	for true {
		Dispatch(vm)
	}
	return nil
}
