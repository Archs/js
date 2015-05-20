package main

import (
	"github.com/Archs/js/dom"
	"github.com/Archs/js/gopherjs-ko"
	_ "github.com/Archs/js/gopherjs-ko/components/clock"
)

func main() {
	el := dom.Wrap(info.Get("element"))
	el = el.QuerySelector("canvas")
	c := canvas.New(el.Object)
	c.Width = 200
	c.Height = 200
	ctx := c.GetContext2D()
	img := dom.CreateElement("img")
	img.Src = "https://mdn.mozillademos.org/files/222/Canvas_createpattern.png"
	img.AddEventListener(dom.EvtLoad, func(e *dom.Event) {
		p := ctx.Cp(img, "repeat")
		ctx.FillStyle = p.Value()
		ctx.FillRect(0, 0, 200, 200)
		g := ctx.CreateRadialGradient(100, 100, 100, 100, 100, 0)
		g.AddColorStop(0.0, "white")
		g.AddColorStop(1.0, "green")
		ctx.FillStyle = g.Value()
		ctx.FillRect(0, 0, 200, 200)
		println(params, info, el, c, ctx)
	}, true)
	ko.ApplyBindings(100)
}
