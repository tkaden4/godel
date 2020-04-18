package main

import (
	"fmt"
	"os"
)

type Debug struct {
	memory    [256]uint8
	registers [32]uint16
	ticks     uint64
	debug     func(do func())
}

func (self *Debug) Quiet() {
	self.debug = func(_ func()) {}
}

func (self *Debug) Loud() {
	self.debug = func(f func()) { f() }
}

func (self *Debug) GetRegister(register uint8) (uint16, error) {
	self.debug(func() {
		fmt.Printf("Get r%02d -> 0x%04x\n", register, self.registers[register])
	})
	return self.registers[register], nil
}

func (self *Debug) SetRegister(register uint8, value uint16) error {
	self.debug(func() {
		fmt.Printf("Set r%02d <- 0x%04x\n", register, value)
	})
	self.registers[register] = value
	return nil
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

func (self *Debug) Tick() {
	self.ticks++
	self.debug(func() {
		fmt.Printf("Tick %04d\n", self.ticks)

	})
}

func (self *Debug) Halt() {
	self.debug(func() {
		fmt.Println("Halting")
	})
	os.Exit(0)
}
