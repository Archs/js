// +build ignore

package main

import (
	"github.com/Archs/js/dom"
	// "github.com/gopherjs/gopherjs/js"
)

// k=document.getElementById('kkk')

func main() {
	div := dom.CreateElement("div")
	div.InnerHTML = "asdf"
	div.Id = "kkk"
	div.SetAttribute("num", 123)
	div.SetAttribute("str", "hello")
	// div.ContentEditable = true
	dom.Body.AppendChild(div)

	div.AddEventListener("click", true, func(ev *dom.Event) {
		println("div clicked", ev)
	})

	div.AddEventListener("mousemove", true, func(ev *dom.Event) {
		println("mouse move:", ev.MovementX, ev.MovementY, ev.ScreenX, ev.ScreenY)
	})

	dom.Body.AddEventListener("keydown", true, func(ev *dom.Event) {
		println("key:", ev.Key, ev.KeyCode)
	})

	// println(div.GetAttribute("num"))
	// js.Global.Call("alert", div.Attributes.Get("str"))

	dom.Body.AppendChild(dom.CreateElement("a"))
	a := dom.CreateElement("a")
	a.Id = "a2"
	a.InnerText = "text for aaa"
	dom.Body.AppendChild(a)
	dom.Body.AppendChild(dom.CreateElement("img"))
	println(dom.Body.QuerySelector("img"))
	a = dom.Body.QuerySelectorAll("a")[0]
	println(0, a.TagName, a.Id, a.InnerHTML, a.InnerText, a.NextElementSibling.TagName)
	a = dom.Body.QuerySelectorAll("a")[1]
	println(1, a.TagName, a.Id, a.InnerHTML, a.InnerText, a.NextElementSibling.TagName)
}
