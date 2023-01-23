package main

import (
	"github.com/win-t/go-sumtypes/examples/stackmachine/machine"
)

func main() {
	instruction, wait := machine.NewMachine()

	instruction <- machine.PushInt{Value: 10}
	instruction <- machine.Print{}

	instruction <- machine.PushInt{Value: 13}
	instruction <- machine.PushInt{Value: 29}
	instruction <- machine.Add{}
	instruction <- machine.Print{}

	instruction <- machine.Exit{}

	wait()
}
