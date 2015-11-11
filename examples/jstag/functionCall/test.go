package main

import (
	"github.com/Archs/js/dom"
)

func test() {
	body := dom.Body()
	div := dom.Document().CreateElement("div")
	div.SetAttribute("id", "div")
	div.SetAttribute("class", "div")
	body.AppendChild(div)
	body.SetAttribute("aaa", 123)
	// println("getAttr", body.GetAttribute)
	println(body.GetAttribute("aaa"))
	// println("Object:", body.Object)
	// println("Body:", body)
	// println("GetElementsByTagName", body.GetElementsByTagName)
	c := body.GetElementsByTagName("div")
	println(c)
	el := c.Item(0)
	println(el.GetAttribute("id"))

	el = body.QuerySelector("div")
	println(el.GetAttribute("id"))
	c = body.QuerySelectorAll("div")
	println(c, c.Length, c.Item(0).GetAttribute("id"))
	println("el:", el)
	el.Style.SetProperty("width", "100")
	el.Style.SetProperty("height", "100")
	el.Style.SetProperty("border", "dotted")
	el.AddEventListener(dom.EvtClick, func(e *dom.Event) {
		println("click:", e, e.ClientX, e.ClientY, e.ScreenX, e.ClientY)
		println("CurrentTarget", e.CurrentTarget)
	})
}

func main() {
	dom.OnDOMContentLoaded(func() {
		test()
	})
}
