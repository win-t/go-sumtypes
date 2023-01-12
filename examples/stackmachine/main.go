package main

import (
	"github.com/win-t/go-sumtypes/examples/stackmachine/machine"
)

func main() {
	instruction, wait := machine.NewMachine()

	instruction <- machine.PushInt{Value: 10}.AsInstruction()
	instruction <- machine.Print{}.AsInstruction()

	instruction <- machine.PushInt{Value: 13}.AsInstruction()
	instruction <- machine.PushInt{Value: 29}.AsInstruction()
	instruction <- machine.Add{}.AsInstruction()
	instruction <- machine.Print{}.AsInstruction()

	instruction <- machine.Exit{}.AsInstruction()

	wait()
}
