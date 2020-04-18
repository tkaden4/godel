package main

func main() {
	vm := CreateDebug()
	program := Program{}
	program.Put(100, uint16('H'))
	program.Cout(101)
	copy(vm.memory[:], program.Bytes)
	vm.Quiet()
	Run(vm)
}
