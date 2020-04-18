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

func (p *Program) Put(location uint16, value uint16) {
	p.appendUint8(0x02)
	p.appendUint16(location)
	p.appendUint16(value)
}

func (p *Program) Cout(location uint16) {
	p.appendUint8(0x05)
	p.appendUint16(location)
}
