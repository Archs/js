// Package canvas provides GopherJS bindings for the JavaScript canvas APIs.
//
// The code is mainly based on package honnef.co/go/js/dom by Dominik Honnef
// in order to create a thin wrapper of the JavaScript canvas API.
package canvas

import (
	"github.com/Archs/js/dom"
	"github.com/gopherjs/gopherjs/js"
)

// canvas元素也可以通过应用CSS的方式来增加边框，设置内边距、外边距等，
// 而且一些CSS属性还可以被canvas内的元素继承。
// 比如字体样式，在canvas内添加的文字，其样式默认同canvas元素本身是一样的。
//
// canvas是行内元素
type Canvas struct {
	*dom.Element
}

// 在canvas中为context设置属性同样要遵从CSS语法
type Context2D struct {
	*js.Object

	// 线条的颜色，默认为”#000000”，其值可以设置为CSS颜色值、渐变对象或者模式对象。
	StrokeStyle interface{} `js:"strokeStyle"`
	// 填充的颜色，默认为”#000000”，与strokeStyle一样，值也可以设置为CSS颜色值、渐变对象或者模式对象。
	FillStyle     interface{} `js:"fillStyle"`
	ShadowColor   string      `js:"shadowColor"`
	ShadowBlur    int         `js:"shadowBlur"`
	ShadowOffsetX int         `js:"shadowOffsetX"`
	ShadowOffsetY int         `js:"shadowOffsetY"`

	// 线条的端点样式，有butt（无）、round（圆头）、square（方头）三种类型可供选择，默认为butt。
	LineCap string `js:"lineCap"`
	// 线条的转折处样式，有round（圆角）、bevel（平角）、miter（尖角）三种；类型可供选择，默认为miter。
	LineJoin string `js:"lineJoin"`
	// 线条的宽度，单位是像素（px），默认为1.0。
	LineWidth int `js:"lineWidth"`
	// 线条尖角折角的锐利程序，默认为10。
	MiterLimit int `js:"miterLimit"`

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

// 用于描绘一个已知左上角顶点位置以及宽和高的矩形，描绘完成后Context的绘制起点会移动到该矩形的左上角顶点。
//
// 参数表示矩形左上角顶点的x、y坐标以及矩形的宽和高。
func (ctx *Context2D) Rect(x, y, width, height int) {
	ctx.Call("rect", x, y, width, height)
}

// 用于使用当前的fillStyle（默认为”#000000”，黑色）样式
// 填充一个左上角顶点在(left, top)点、宽为width、高为height的矩形。
func (ctx *Context2D) FillRect(left, top, width, height int) {
	ctx.Call("fillRect", left, top, width, height)
}

// 用于使用当前的线条风格绘制一个左上角顶点在(left, top)点、宽为width、高为height的矩形边框。
func (ctx *Context2D) StrokeRect(left, top, width, height int) {
	ctx.Call("strokeRect", left, top, width, height)
}

// clearRect的作用是清除矩形区域内的所有内容并将它恢复到初始状态，即透明色
// 用于清除左上角顶点在(left,top)点、宽为width、高为height的矩形区域内的所有内容。
func (ctx *Context2D) ClearRect(left, top, width, height int) {
	ctx.Call("clearRect", left, top, width, height)
}

// Paths

// 用于使用当前的填充风格来填充路径的区域。
func (ctx *Context2D) Fill() {
	ctx.Call("fill")
}

// 用于按照已有的路径绘制线条。
func (ctx *Context2D) Stroke() {
	ctx.Call("stroke")
}

// canvas中很多用于设置样式和外观的函数也同样不会直接修改显示结果。
// HTML5 Canvas的基本图形都是以路径为基础的。通常使用Context对象的moveTo()、lineTo()、rect()、arc()等方法先在画布中描出图形的路径点，然后使用fill()或者stroke()方法依照路径点来填充图形或者绘制线条。
//
// 通常，在开始描绘路径之前需要调用Context对象的beginPath()方法，其作用是清除之前的路径并提醒Context开始绘制一条新的路径，否则当调用stroke()方法的时候会绘制之前所有的路径，影响绘制效果，同时也因为重复多次操作而影响网页性能。另外，调用Context对象的closePath()方法可以显式地关闭当前路径，不过不会清除路径。
// 只有当对路径应用绘制（stroke）或填充（fill）方法时，结果才会显示出来
func (ctx *Context2D) BeginPath() {
	ctx.Call("beginPath")
}

// 用于显式地指定路径的起点。默认状态下，第一条路径的起点是画布的(0, 0)点，之后的起点是上一条路径的终点。
// 两个参数分为表示起点的x、y坐标值。
func (ctx *Context2D) MoveTo(x, y int) {
	ctx.Call("moveTo", x, y)
}

// 这个函数的行为同lineTo很像，唯一的差别在于closePath会将路径的起始坐标自动作为目标坐标。
// closePath还会通知canvas当前绘制的图形已经闭合或者形成了完全封闭的区域，
// 这对将来的填充和描边都非常有用。
func (ctx *Context2D) ClosePath() {
	ctx.Call("closePath")
}

// 用于描绘一条从起点从指定位置的直线路径，描绘完成后绘制的起点会移动到该指定位置。
//
// 参数表示指定位置的x、y坐标值。
func (ctx *Context2D) LineTo(x, y int) {
	ctx.Call("lineTo", x, y)
}

// 用于按照已有的路线在画布中设置剪辑区域。
//
// 调用clip()方法之后，图形绘制代码只对剪辑区域有效而不再影响区域外的画布。
//
// 如调用之前没有描绘路径（即默认状态下），则得到的剪辑区域为整个Canvas区域。
func (ctx *Context2D) Clip() {
	ctx.Call("clip")
}

// 用于描绘一个以(x, y)点为圆心，radius为半径，startAngle为起始弧度，endAngle为终止弧度的圆弧。
// anticlockwise为布尔型的参数，true表示逆时针，false表示顺时针。
//
// quadraticCurveTo 函数绘制曲线的起点是当前坐标，带有两组（x,y）参数。第二组是指曲线的终点。
// 第一组代表控制点（control point）。
// 所谓的控制点位于曲线的旁边（不是曲线之上），其作用相当于对曲线产生一个拉力。
// 通过调整控制点的位置，就可以改变曲线的曲率。
func (ctx *Context2D) QuadraticCurveTo(cpx, cpy, x, y int) {
	ctx.Call("quadraticCurveTo", cpx, cpy, x, y)
}

// 用于描绘以当前Context绘制起点为起点，(cpx1,cpy1)点和(cpx2, cpy2)点为两个控制点，
// (x, y)点为终点的贝塞尔曲线路径。
func (ctx *Context2D) BezierCurveTo(cp1x, cp1y, cp2x, cp2y, x, y int) {
	ctx.Call("bezierCurveTo", cp1x, cp1y, cp2x, cp2y, x, y)
}

// 参数中的两个弧度以0表示0°，位置在3点钟方向；Math.PI值表示180°，位置在9点钟方向。
func (ctx *Context2D) Arc(x, y, radius int, sAngle, eAngle float32, counterclockwise bool) {
	ctx.Call("arc", x, y, radius, sAngle, eAngle, counterclockwise)
}

// 用于描绘一个与两条线段相切的圆弧，两条线段分别以当前Context绘制起点和(x2, y2)点为起点，都以(x1, y1)点为终点，圆弧的半径为radius。
// 描绘完成后绘制起点会移动到以(x2, y2)为起点的线段与圆弧的切点。
func (ctx *Context2D) ArcTo(x1, y1, x2, y2, r int) {
	ctx.Call("arcTo", x1, y1, x2, y2, r)
}

func (ctx *Context2D) IsPointInPath(x, y int) bool {
	return ctx.Call("isPointInPath", x, y).Bool()
}

// 缩放
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

// fillText()方法能够在画布中绘制字符串
//
// 需绘制的字符串，绘制到画布中时左上角在画布中的横坐标及纵坐标，绘制的字符串的最大长度。其中最大长度maxWidth是可选参数。另外，可以通过改变Context对象的font属性来调整字符串的字体以及大小，默认为”10px sans-serif”。
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

// 保存当前绘图状态
func (ctx *Context2D) Save() {
	ctx.Call("save")
}

// 恢复原有的绘图状态
func (ctx *Context2D) Restore() {
	ctx.Call("restore")
}

// Context对象中拥有drawImage()方法可以将外部图片绘制到Canvas中。
//
// image参数可以是HTMLImageElement、HTMLCanvasElement或者HTMLVideoElement。
func (ctx *Context2D) DrawImage(image *dom.Element, dx, dy, dw, dh int) {
	ctx.Call("drawImage")
}
