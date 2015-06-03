package main

import (
	"github.com/Archs/js/dom"
	"github.com/gopherjs/gopherjs/js"
)

type A struct {
	*js.Object
	Method func() `js:"Method"` // js tag with method invoke, panic
	// Method func() // this is fine
}

func main() {
	a := &A{
		Object: js.Global.Get("Object").New(),
		Method: func() {
			println("hello")
		},
	}
	println(a.Method) // this is fine
	// a.Method()        // this cause panic
	dom.OnDOMContentLoaded(func() {
		body := dom.Body()
		body.SetAttribute("aaa", 123)
		println(body.GetAttribute("aaa"))
	})
}
