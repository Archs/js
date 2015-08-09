package main

import (
	"github.com/Archs/js/gopherjs-ko"
	"github.com/gopherjs/gopherjs/js"
)

func main() {
	c := ko.NewComponent("test")
	c.Template.Markup = "<h3>Component Test</h3>"
	c.ViewModel.Creator = func(params *js.Object, info *ko.ComponentInfo) ko.ViewModel {
		println("creator called")
		println(info.TemplateNodes)
		for i := 0; i < info.TemplateNodes.Length; i++ {
			el := info.TemplateNodes.Item(i)
			println(i, el, el.NodeType, el.TagName, el.InnerText)
		}
		return nil
	}
	ko.RegisterComponent(c)
	ko.ApplyBindings(nil)
}
