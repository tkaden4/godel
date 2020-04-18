package main

func main() {
	vm := &Debug{}
	program := Program{}
	program.Put(100, uint16('H'))
	program.Cout(101)
	copy(vm.memory[:], program.Bytes)
	// vm.Loud()
	vm.Quiet()
	Run(vm)
}
