package main

const IP uint8 = 31
const PO uint8 = 30

type Memory interface {
	SetMemory(location uint16, value uint8) error
	GetMemory(location uint16) (uint8, error)
}

type Registers interface {
	GetRegister(register uint8) (uint8, error)
	SetRegister(register uint8, value uint8) error
}

type VM interface {
	Memory() Memory
	Registers() Registers
	Halt()
}

func ReadImmediateUint8(vm VM) (uint8, error) {
	ip, _ := vm.Registers().GetRegister(IP)
	po, _ := vm.Registers().GetRegister(PO)
	location := uint16(po)*0x0100 | uint16(ip)
	value, _ := vm.Memory().GetMemory(location)
	newIP := ip + 1
	newPO := po
	if ip == 0xff {
		newPO++
	}
	vm.Registers().SetRegister(IP, newIP)
	vm.Registers().SetRegister(PO, newPO)
	return value, nil
}

func ReadImmediateUint16(vm VM) (uint16, error) {
	beByte, _ := ReadImmediateUint8(vm)
	leByte, _ := ReadImmediateUint8(vm)
	var be uint16 = uint16(beByte)
	var le uint16 = uint16(leByte)
	val := ((be << 8) | le)
	return val, nil
}
