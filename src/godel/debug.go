package main

import (
	"fmt"
	"os"
)

const REGISTER_OFFSET uint16 = 2047

type Debug struct {
	memory [2048]uint8
	debug  func(do func())
}

func (self *Debug) Quiet() {
	self.debug = func(_ func()) {}
}

func (self *Debug) Loud() {
	self.debug = func(thunk func()) { thunk() }
}

func (self *Debug) GetRegister(register uint8) (uint8, error) {
	value, err := self.GetMemory(REGISTER_OFFSET - uint16(register))
	self.debug(func() {
		fmt.Printf("Get r%02d -> 0x%04x\n", register, value)
	})
	return value, err
}

func (self *Debug) SetRegister(register uint8, value uint8) error {
	self.debug(func() {
		fmt.Printf("Set r%02d <- 0x%04x\n", register, value)
	})
	return self.SetMemory(REGISTER_OFFSET-uint16(register), value)
}

func (self *Debug) GetMemory(location uint16) (uint8, error) {
	self.debug(func() {
		fmt.Printf("Get 0x%04x -> 0x%04x\n", location, self.memory[location])
	})
	return self.memory[location], nil
}

func (self *Debug) SetMemory(location uint16, value uint8) error {
	self.debug(func() {
		fmt.Printf("Set 0x%04x := 0x%04x\n", location, value)
	})
	self.memory[location] = value
	return nil
}

func (self *Debug) Memory() Memory {
	return self
}

func (self *Debug) Registers() Registers {
	return self
}

func (self *Debug) Halt() {
	self.debug(func() {
		fmt.Println("Halting")
	})
	os.Exit(0)
}
