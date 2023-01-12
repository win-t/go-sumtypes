package main

import (
	"fmt"

	"github.com/win-t/go-sumtypes/examples/json/json"
)

func main() {
	v := json.Object{
		"hello": json.String("world").Value(),
		"hai":   json.Null{}.Value(),
		"lala": json.Array{
			json.String("one").Value(),
			json.Number(2).Value(),
			json.Bool(true).Value(),
		}.Value(),
	}.Value()

	fmt.Println(json.Stringify(v))
}
