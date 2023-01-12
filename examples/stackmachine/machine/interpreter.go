package machine

import (
	"fmt"
)

func NewMachine() (ch chan<- Instruction, wait func()) {
	instruction := make(chan Instruction)
	done := make(chan struct{})

	go func() {
		defer close(done)

		var stack []int
		stackPush := func(i int) {
			stack = append(stack, i)
		}
		stackPop := func() int {
			i := stack[len(stack)-1]
			stack = stack[:len(stack)-1]
			return i
		}

		exit := false

		for {
			(<-instruction).Case(
				func(Print) {
					data := stackPop()
					fmt.Println(data)
				},

				func(push PushInt) {
					stack = append(stack, push.Value)
				},

				func(Add) {
					a := stackPop()
					b := stackPop()
					stackPush(a + b)
				},

				func(Exit) {
					exit = true
				},
			)

			if exit {
				return
			}
		}
	}()

	return instruction, func() { <-done }
}
