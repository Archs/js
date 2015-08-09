package main

import (
	"github.com/Archs/js/dom"
	"github.com/Archs/js/gopherjs-ko"
	"github.com/Archs/js/gopherjs-ko/components/uploader"
)

func main() {
	println(`ko.Get("utils"):`, ko.Get("utils"))
	uploader.SetCallback(func(url string, files []*dom.File) {
		for _, file := range files {
			println("upload", file.Name, "to", url, "done")
		}
	})
	ko.ApplyBindings(nil)
}
