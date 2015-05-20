// Package canvas provides GopherJS bindings for the JavaScript canvas APIs.
//
// The code is mainly based on package honnef.co/go/js/dom by Dominik Honnef
// in order to create a thin wrapper of the JavaScript canvas API.
package canvas

import (
	"github.com/Archs/js/dom"
	"github.com/gopherjs/gopherjs/js"
)

type Canvas struct {
	*dom.Element
}

type Context2D struct {
	*js.Object

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

// New creates a Canvas instance
// el is the html element
func New(el *js.Object) *Canvas {
	return &Canvas{dom.Wrap(el)}
}

func (c *Canvas) GetContext2D() *Context2D {
	ctx := c.Call("getContext", "2d")
	return &Context2D{Object: ctx}
}

// canvas.toDataURL("image/jpeg") or canvas.toDataURL()
func (c *Canvas) ToDataUrl(mimeType ...string) string {
	var o *js.Object
	if len(mimeType) == 0 {
		o = c.Call("toDataURL")
	} else {
		o = c.Call("toDataURL", mimeType)
	}
	return o.String()
}

// Colors, Styles, and Shadows

func (ctx *Context2D) CreateLinearGradient(x0, y0, x1, y1 int) {
	ctx.Call("createLinearGradient", x0, y0, x1, y1)
}

// Rectangles

func (ctx *Context2D) Rect(x, y, width, height int) {
	ctx.Call("rect", x, y, width, height)
}

func (ctx *Context2D) FillRect(x, y, width, height int) {
	ctx.Call("fillRect", x, y, width, height)
}

func (ctx *Context2D) StrokeRect(x, y, width, height int) {
	ctx.Call("strokeRect", x, y, width, height)
}

func (ctx *Context2D) ClearRect(x, y, width, height int) {
	ctx.Call("clearRect", x, y, width, height)
}

// Paths

func (ctx *Context2D) Fill() {
	ctx.Call("fill")
}

func (ctx *Context2D) Stroke() {
	ctx.Call("stroke")
}

func (ctx *Context2D) BeginPath() {
	ctx.Call("beginPath")
}

func (ctx *Context2D) MoveTo(x, y int) {
	ctx.Call("moveTo", x, y)
}

func (ctx *Context2D) ClosePath() {
	ctx.Call("closePath")
}

func (ctx *Context2D) LineTo(x, y int) {
	ctx.Call("lineTo", x, y)
}

func (ctx *Context2D) Clip() {
	ctx.Call("clip")
}

func (ctx *Context2D) QuadraticCurveTo(cpx, cpy, x, y int) {
	ctx.Call("quadraticCurveTo", cpx, cpy, x, y)
}

func (ctx *Context2D) BezierCurveTo(cp1x, cp1y, cp2x, cp2y, x, y int) {
	ctx.Call("bezierCurveTo", cp1x, cp1y, cp2x, cp2y, x, y)
}

func (ctx *Context2D) Arc(x, y, r int, sAngle, eAngle float32, counterclockwise bool) {
	ctx.Call("arc", x, y, r, sAngle, eAngle, counterclockwise)
}

func (ctx *Context2D) ArcTo(x1, y1, x2, y2, r int) {
	ctx.Call("arcTo", x1, y1, x2, y2, r)
}

func (ctx *Context2D) IsPointInPath(x, y int) bool {
	return ctx.Call("isPointInPath", x, y).Bool()
}

// Transformations

func (ctx *Context2D) Scale(scaleWidth, scaleHeight int) {
	ctx.Call("scale", scaleWidth, scaleHeight)
}

func (ctx *Context2D) Rotate(angle int) {
	ctx.Call("rotate", angle)
}

func (ctx *Context2D) Translate(x, y int) {
	ctx.Call("translate", x, y)
}

func (ctx *Context2D) Transform(a, b, c, d, e, f int) {
	ctx.Call("transform", a, b, c, d, e, f)
}

func (ctx *Context2D) SetTransform(a, b, c, d, e, f int) {
	ctx.Call("setTransform", a, b, c, d, e, f)
}

// Text

func (ctx *Context2D) FillText(text string, x, y, maxWidth int) {
	if maxWidth == -1 {
		ctx.Call("fillText", text, x, y)
		return
	}

	ctx.Call("fillText", text, x, y, maxWidth)
}

func (ctx *Context2D) StrokeText(text string, x, y, maxWidth int) {
	if maxWidth == -1 {
		ctx.Call("strokeText", text, x, y)
		return
	}

	ctx.Call("strokeText", text, x, y, maxWidth)
}

// canvas state

func (ctx *Context2D) Save() {
	ctx.Call("save")
}

func (ctx *Context2D) Restore() {
	ctx.Call("restore")
}
