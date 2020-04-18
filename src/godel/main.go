package main

func main() {
	vm := &Debug{}
	program := Program{}
	program.Put(1, uint16('H'))
	program.Cout(1)
	program.Put(1, uint16('e'))
	program.Cout(1)
	program.Put(1, uint16('l'))
	program.Cout(1)
	program.Put(1, uint16('l'))
	program.Cout(1)
	program.Put(1, uint16('o'))
	program.Cout(1)
	program.Put(1, uint16('\n'))
	program.Cout(1)
	copy(vm.memory[:], program.Bytes)
	vm.Quiet()
	Run(vm)
}
