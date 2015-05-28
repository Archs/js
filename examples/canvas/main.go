package main

import (
	"github.com/Archs/js/canvas"
	"github.com/Archs/js/dom"
)

func main() {
	el := dom.GetElementById("canvas")
	c := canvas.New(el.Object)
	c.Width = 400
	c.Height = 400
	ctx := c.GetContext2D()
	img := dom.CreateElement("img")
	img.Src = "img/Canvas_createpattern.png"
	img.AddEventListener(dom.EvtLoad, func(e *dom.Event) {
		p := ctx.CreatePattern(img, "repeat")
		ctx.FillStyle = p.Value()
		ctx.FillRect(0, 0, 400, 400)
		im := ctx.GetImage(0, 0, 20, 20)
		println(im.At(0, 0).RGBA())
		println(im.At(0, 10).RGBA())
		println(im, im.At(0, 0), im.Bounds())
		// g := ctx.CreateRadialGradient(100, 100, 100, 100, 100, 0)
		// g.AddColorStop(0.0, "white")
		// g.AddColorStop(1.0, "green")
		// ctx.FillStyle = g.Value()
		// ctx.FillRect(0, 0, 200, 200)
	})
}
