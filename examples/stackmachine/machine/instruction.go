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

func (e Print) AsInstruction() Instruction   { var i Instruction; i.Set0(e); return i }
func (e PushInt) AsInstruction() Instruction { var i Instruction; i.Set1(e); return i }
func (e Add) AsInstruction() Instruction     { var i Instruction; i.Set2(e); return i }
func (e Exit) AsInstruction() Instruction    { var i Instruction; i.Set3(e); return i }
