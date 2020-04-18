package main

import "errors"

func Noop(vm VM) error {
	vm.Tick()
	return nil
}

func Halt(vm VM) error {
	vm.Tick()
	vm.Halt()
	return nil
}

func PutImmediate(vm VM) error {
	location, _ := ReadImmediateUint16(vm)
	vm.Tick()
	value, _ := ReadImmediateUint16(vm)
	vm.Tick()
	upperByte := uint8((value & 0xff00) >> 8)
	lowerByte := uint8(value)
	vm.Memory().SetMemory(location, upperByte)
	vm.Memory().SetMemory(location+1, lowerByte)
	return nil
}

func Dispatch(vm VM) error {
	instruction, _ := ReadImmediateUint8(vm)
	switch instruction {
	case 0x00:
		return Noop(vm)
	case 0x01:
		return Halt(vm)
	case 0x02:
		return PutImmediate(vm)
	default:
		return errors.New("Invalid opcode")
	}
}
