package main

import (
	"github.com/gopherjs/gopherjs/js"
	"github.com/neelance/dom"
	"github.com/neelance/dom/bind"
	"github.com/neelance/dom/elem"
	"github.com/neelance/dom/prop"
)

// SetBody replaces the page's body with the given aspects.
func ApplyToNode(node *js.Object, aspects ...dom.Aspect) {
	dom.Group(aspects...).Apply(node, 0, 1)
}

func getElementById(id string) *js.Object {
	return js.Global.Get("document").Call("getElementById", id)
}

func main() {
	input := "init value"
	scope := bind.NewScope()
	scope.NewListener(func() {
		println("input value:", input)
	})
	ApplyToNode(getElementById("inp"), bind.Value(&input, scope))
	ApplyToNode(getElementById("main"),
		elem.Div(elem.Header2(dom.Text("Nice to see you"))),
		elem.Input(prop.Type(prop.TypePassword)),
	)
}
