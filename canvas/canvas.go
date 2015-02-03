// Package canvas provides GopherJS bindings for the JavaScript canvas APIs.
//
// Code mainly borrowed from honnef.co/go/js/dom by Dominik Honnef in order to
// be a separated wrapper of the JavaScript canvas API.
package canvas

import (
	"github.com/gopherjs/gopherjs/js"
)

type CanvasElement struct {
	js.Object
	// basic attr
	Id string `js:"id"`
	// width & height
	Width  int `js:"width"`
	Height int `js:"height"`
}

type CanvasRenderingContext2D struct {
	js.Object

	// Colors, Styles, and Shadows
	FillStyle     string `js:"fillStyle"`
	StrokeStyle   string `js:"strokeStyle"`
	ShadowColor   string `js:"shadowColor"`
	ShadowBlur    int    `js:"shadowBlur"`
	ShadowOffsetX int    `js:"shadowOffsetX"`
	ShadowOffsetY int    `js:"shadowOffsetY"`

	// Line Styles
	LineCap    string `js:"lineCap"`
	LineJoin   string `js:"lineJoin"`
	LineWidth  int    `js:"lineWidth"`
	MiterLimit int    `js:"miterLimit"`

	// Text
	Font         string `js:"font"`
	TextAlign    string `js:"textAlign"`
	TextBaseline string `js:"textBaseline"`

	// Compositing
	GlobalAlpha              float64 `js:"globalAlpha"`
	GlobalCompositeOperation string  `js:"globalCompositeOperation"`
}

// el is then html element
func NewCanvas(el js.Object) *CanvasElement {
	return &CanvasElement{Object: el}
}

func (c *CanvasElement) GetContext2D() *CanvasRenderingContext2D {
	ctx := c.Call("getContext", "2d")
	return &CanvasRenderingContext2D{Object: ctx}
}

// canvas.toDataURL("image/jpeg") or canvas.toDataURL()
func (c *CanvasElement) ToDataUrl(mimeType ...string) string {
	var o js.Object
	if len(mimeType) == 0 {
		o = c.Call("toDataURL")
	} else {
		o = c.Call("toDataURL", mimeType)
	}
	return o.String()
}

// Colors, Styles, and Shadows

func (ctx *CanvasRenderingContext2D) CreateLinearGradient(x0, y0, x1, y1 int) {
	ctx.Call("createLinearGradient", x0, y0, x1, y1)
}

// Rectangles

func (ctx *CanvasRenderingContext2D) Rect(x, y, width, height int) {
	ctx.Call("rect", x, y, width, height)
}

func (ctx *CanvasRenderingContext2D) FillRect(x, y, width, height int) {
	ctx.Call("fillRect", x, y, width, height)
}

func (ctx *CanvasRenderingContext2D) StrokeRect(x, y, width, height int) {
	ctx.Call("strokeRect", x, y, width, height)
}

func (ctx *CanvasRenderingContext2D) ClearRect(x, y, width, height int) {
	ctx.Call("clearRect", x, y, width, height)
}

// Paths

func (ctx *CanvasRenderingContext2D) Fill() {
	ctx.Call("fill")
}

func (ctx *CanvasRenderingContext2D) Stroke() {
	ctx.Call("stroke")
}

func (ctx *CanvasRenderingContext2D) BeginPath() {
	ctx.Call("beginPath")
}

func (ctx *CanvasRenderingContext2D) MoveTo(x, y int) {
	ctx.Call("moveTo", x, y)
}

func (ctx *CanvasRenderingContext2D) ClosePath() {
	ctx.Call("closePath")
}

func (ctx *CanvasRenderingContext2D) LineTo(x, y int) {
	ctx.Call("lineTo", x, y)
}

func (ctx *CanvasRenderingContext2D) Clip() {
	ctx.Call("clip")
}

func (ctx *CanvasRenderingContext2D) QuadraticCurveTo(cpx, cpy, x, y int) {
	ctx.Call("quadraticCurveTo", cpx, cpy, x, y)
}

func (ctx *CanvasRenderingContext2D) BezierCurveTo(cp1x, cp1y, cp2x, cp2y, x, y int) {
	ctx.Call("bezierCurveTo", cp1x, cp1y, cp2x, cp2y, x, y)
}

func (ctx *CanvasRenderingContext2D) Arc(x, y, r int, sAngle, eAngle float32, counterclockwise bool) {
	ctx.Call("arc", x, y, r, sAngle, eAngle, counterclockwise)
}

func (ctx *CanvasRenderingContext2D) ArcTo(x1, y1, x2, y2, r int) {
	ctx.Call("arcTo", x1, y1, x2, y2, r)
}

func (ctx *CanvasRenderingContext2D) IsPointInPath(x, y int) bool {
	return ctx.Call("isPointInPath", x, y).Bool()
}

// Transformations

func (ctx *CanvasRenderingContext2D) Scale(scaleWidth, scaleHeight int) {
	ctx.Call("scale", scaleWidth, scaleHeight)
}

func (ctx *CanvasRenderingContext2D) Rotate(angle int) {
	ctx.Call("rotate", angle)
}

func (ctx *CanvasRenderingContext2D) Translate(x, y int) {
	ctx.Call("translate", x, y)
}

func (ctx *CanvasRenderingContext2D) Transform(a, b, c, d, e, f int) {
	ctx.Call("transform", a, b, c, d, e, f)
}

func (ctx *CanvasRenderingContext2D) SetTransform(a, b, c, d, e, f int) {
	ctx.Call("setTransform", a, b, c, d, e, f)
}

// Text

func (ctx *CanvasRenderingContext2D) FillText(text string, x, y, maxWidth int) {
	if maxWidth == -1 {
		ctx.Call("fillText", text, x, y)
		return
	}

	ctx.Call("fillText", text, x, y, maxWidth)
}

func (ctx *CanvasRenderingContext2D) StrokeText(text string, x, y, maxWidth int) {
	if maxWidth == -1 {
		ctx.Call("strokeText", text, x, y)
		return
	}

	ctx.Call("strokeText", text, x, y, maxWidth)
}

// canvas state

func (ctx *CanvasRenderingContext2D) Save() {
	ctx.Call("save")
}

func (ctx *CanvasRenderingContext2D) Restore() {
	ctx.Call("restore")
}
