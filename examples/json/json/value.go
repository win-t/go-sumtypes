package json

import (
	"github.com/win-t/go-sumtypes/sum6"
)

type Value struct {
	sum6.Type[
		String,
		Number,
		Bool,
		Null,
		Array,
		Object,
	]
}

type String string
type Number float64
type Bool bool
type Null struct{}
type Array []Value
type Object map[string]Value

// this is needed, so we don't have to type the type parameter manually
var valueRef Value

func (v String) Value() Value { return Value{valueRef.New0(v)} }
func (v Number) Value() Value { return Value{valueRef.New1(v)} }
func (v Bool) Value() Value   { return Value{valueRef.New2(v)} }
func (v Null) Value() Value   { return Value{valueRef.New3(v)} }
func (v Array) Value() Value  { return Value{valueRef.New4(v)} }
func (v Object) Value() Value { return Value{valueRef.New5(v)} }
