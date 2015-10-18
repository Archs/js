package main

import (
	"github.com/Archs/js/dom"
	"github.com/Archs/js/gopherjs-ko"
	"github.com/Archs/js/gopherjs-ko/components/uploader"
	"github.com/Archs/js/gopherjs-ko/plugins/async"
)

func main() {
	dom.OnDOMContentLoaded(func() {
		println(`ko.Get("utils"):`, ko.Get("utils"))
		async.AsyncTemplateConfig().SetDefaultPath(".")
		uploader.SetCallback(func(url string, files []*dom.File) {
			for _, file := range files {
				println("upload", file.Name, "to", url, "done")
			}
		})
		ko.ApplyBindings(nil)
	})
}
