package a

import (
	_ "github.com/Archs/js/gopherjs-ko"
	"github.com/gopherjs/gopherjs/js"
)

func Hello() {
	println("Hello from package a")
}

func T() {
	ko := js.Global.Get("ko")
	println(ko.Get("testVar").String())
}
