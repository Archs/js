package main

import (
	"fmt"
	"github.com/Archs/js/canvas"
	"github.com/Archs/js/dom"
)

func test() {
	span := dom.Document().CreateElement("span")
	dom.Body().AppendChild(span)
	el := dom.Document().GetElementById("canvas")
	c := canvas.New(el.Object)
	ctx := c.GetContext2D()
	img := dom.Document().CreateElement("img")
	img.Src = "img/bull.png"
	img.AddEventListener(dom.EvtLoad, func(e *dom.Event) {
		// p := ctx.CreatePattern(img, canvas.PatternNoRepeat)
		// ctx.FillStyle = p.Value()
		// ctx.FillRect(0, 0, 400, 400)
		// im := ctx.GetImageData(0, 0, 20, 20)
		// println("im20:", im, im.Data, im.Width, im.Height)
		// println(im.At(0, 0))
		// println(im.At(0, 10))
		// println(im, im.At(0, 0))
		println(img, 0, 0, img.Width, img.Height)
		el.Width = img.Width
		el.Height = img.Height
		ctx.DrawImage(img, 0, 0, float64(img.Width), float64(img.Height))
	})
	el.AddEventListener(dom.EvtMousemove, func(e *dom.Event) {
		println("mouse:", e.LayerX, e.LayerY)
		im := ctx.GetImageData(e.LayerX, e.LayerY, 1, 1)
		println("im1:", im, im.Data, im.Width, im.Height)
		c := im.At(0, 0)
		println(c)
		value := fmt.Sprintf("rgba(%d,%d,%d,%d)", c.R, c.G, c.B, c.A)
		span.Style.SetProperty("background", value)
		span.InnerText = value
	})
}

func main() {
	dom.OnDOMContentLoaded(test)
}
