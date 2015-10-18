package main

import (
	"github.com/Archs/js/examples/incTest/a"
	"github.com/Archs/js/gopherjs-ko"
)

func main() {
	println("main loaded!")
	// async.AsyncTemplateConfig().SetDefaultPath("path")
	a.Hello()
	ko.ApplyBindings(nil)
}
