package main

func (p *Program) helloWorld() {
	for _, x := range "Hello, World\n" {
		p.Put(100, uint16(x))
		p.Cout(101)
	}
}

func main() {
	vm := CreateDebug()
	program := Program{}
	program.helloWorld()
	copy(vm.memory[:], program.Bytes)
	vm.Quiet()
	Run(vm)
}
