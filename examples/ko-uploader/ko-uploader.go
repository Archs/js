package main

import (
	"github.com/Archs/js/dom"
	"github.com/Archs/js/gopherjs-ko"
	"github.com/Archs/js/gopherjs-ko/components/uploader"
	"github.com/Archs/js/gopherjs-ko/plugins/csp"
)

func main() {
	csp.EnableSecureBinding()
	uploader.SetCallback(func(url string, files []*dom.File) {
		for _, file := range files {
			println("upload", file.Name, "to", url, "done")
		}
	})
	ko.ApplyBindings(nil)
}
