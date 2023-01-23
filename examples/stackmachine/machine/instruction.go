package machine

//go:generate go run github.com/win-t/go-sumtypes@latest

type Print struct{}
type PushInt struct{ Value int }
type Add struct{}
type Exit struct{}

type Instruction interface {
	sumtype()
	Case(
		func(Print),
		func(PushInt),
		func(Add),
		func(Exit),
	)
}
