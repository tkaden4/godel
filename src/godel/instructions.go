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
	location, _ := ReadImmediateUint16(vm)
	value, _ := ReadImmediateUint16(vm)
	vm.Memory().SetMemory(location, uint8((value&0xff00)>>8))
	vm.Memory().SetMemory(location+1, uint8(value))
	return nil
}

func Copy(vm VM) error {
	fst, _ := ReadImmediateUint16(vm)
	fstValue, _ := vm.Memory().GetMemory(fst)
	snd, _ := ReadImmediateUint16(vm)
	return vm.Memory().SetMemory(snd, fstValue)
}

func Cout(vm VM) error {
	location, _ := ReadImmediateUint16(vm)
	value, _ := vm.Memory().GetMemory(location)
	fmt.Print(string(value))
	return nil
}

func flag(t bool) uint8 {
	if t {
		return 1
	}
	return 0
}

func Cmp(vm VM) error {
	locFst, _ := ReadImmediateUint16(vm)
	a, _ := vm.Memory().GetMemory(locFst)
	locSnd, _ := ReadImmediateUint16(vm)
	b, _ := vm.Memory().GetMemory(locSnd)
	vm.Registers().SetRegister(EQ, flag(a == b))
	vm.Registers().SetRegister(LT, flag(a < b))
	vm.Registers().SetRegister(GT, flag(a > b))
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
	case 0x03:
		return Copy(vm)
	case 0x04:
		return Cmp(vm)
	case 0x05:
		return Cout(vm)
	default:
		return errors.New("Invalid opcode")
	}
}

func Run(vm VM) error {
	for true {
		if err := Dispatch(vm); err != nil {
			return err
		}
	}
	return nil
}
