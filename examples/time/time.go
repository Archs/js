package main

import (
	"github.com/Archs/js/examples/time/timefns"
	"github.com/gopherjs/gopherjs/js"
)

type A struct {
	*js.Object
	Str string `js:"str"`
}

func (a *A) Print() {
	println(a.Str)
}

func newA() *A {
	return &A{
		Object: js.Global.Get("Object").New(),
		Str:    "nice 2",
	}
}

func fn(v interface{}) {
	println(v)
}

func main() {
	timefns.Hello(newA())
}
