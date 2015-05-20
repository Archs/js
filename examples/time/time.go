package main

import (
	"github.com/Archs/js/dom"
	"time"
)

func main() {
	dom.OnDOMContentLoaded(func() {
		btn := dom.CreateElement("button")
		btn.InnerText = "hello"
		dom.Body().AppendChild(btn)
		btn.AddEventListener(dom.EvtClick, func(e *dom.Event) {
			println("hello")
			btn.InnerText = time.Now().String()
		}, false)
	})
}
