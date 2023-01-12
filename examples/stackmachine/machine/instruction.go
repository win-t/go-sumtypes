package machine

import (
	"github.com/win-t/go-sumtypes/sum4"
)

type Print struct{}
type PushInt struct{ Value int }
type Add struct{}
type Exit struct{}

type Instruction struct {
	sum4.Type[
		Print,
		PushInt,
		Add,
		Exit,
	]
}

func (e Print) AsInstruction() Instruction   { return Instruction{Instruction{}.New0(e)} }
func (e PushInt) AsInstruction() Instruction { return Instruction{Instruction{}.New1(e)} }
func (e Add) AsInstruction() Instruction     { return Instruction{Instruction{}.New2(e)} }
func (e Exit) AsInstruction() Instruction    { return Instruction{Instruction{}.New3(e)} }
