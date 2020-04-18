package main

type Program struct {
	Bytes []uint8
}

func (p *Program) appendUint8(value uint8) {
	p.Bytes = append(p.Bytes, value)
}

func (p *Program) appendUint16(value uint16) {
	p.Bytes = append(p.Bytes, uint8((value&0xff00)>>8), uint8(value))
}

func (p *Program) Put(register uint8, value uint16) {
	p.appendUint8(0x02)
	p.appendUint8(register)
	p.appendUint16(value)
}

func (p *Program) Cout(register uint8) {
	p.appendUint8(0x05)
	p.appendUint8(register)
}
