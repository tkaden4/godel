package main

const IP = 0

type Memory interface {
	SetMemory(location uint16, value uint8) error
	GetMemory(location uint16) (uint8, error)
}

type Registers interface {
	GetRegister(register uint8) (uint16, error)
	SetRegister(register uint8, value uint16) error
}

type VM interface {
	Memory() Memory
	Registers() Registers
	Tick()
	Halt()
}

func ReadImmediateUint8(vm VM) (uint8, error) {
	ip, _ := vm.Registers().GetRegister(IP)
	byte, _ := vm.Memory().GetMemory(ip)
	vm.Registers().SetRegister(IP, ip+1)
	return byte, nil
}

func ReadImmediateUint16(vm VM) (uint16, error) {
	beByte, _ := ReadImmediateUint8(vm)
	leByte, _ := ReadImmediateUint8(vm)
	var be uint16 = uint16(beByte)
	var le uint16 = uint16(leByte)
	val := ((be << 8) | le)
	return val, nil
}
