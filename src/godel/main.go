package main

func basicProgram(vm VM) {
	PutImmediate(vm)
}

func main() {
	basicProgram(&Debug{})
}
