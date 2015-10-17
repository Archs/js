package clock

import (
	"github.com/Archs/js/canvas"
	"github.com/Archs/js/dom"
	"github.com/Archs/js/vue"
	"github.com/gopherjs/gopherjs/js"
	"math"
	"time"
)

func drawClockPane(ctx *canvas.CanvasRenderingContext2D, r int) {
	ctx.Save()
	// cycle
	ctx.StrokeStyle = "black"
	ctx.LineWidth = 4
	ctx.BeginPath()
	ctx.Arc(0, 0, r, 0., float32(2*math.Pi), true)
	ctx.ClosePath()
	ctx.Stroke()
	// hour tick
	itv := 2 * math.Pi / 12.0
	r0 := float64(r - 10)
	r1 := float64(r)
	for i := 0; i < 12; i++ {
		angle := float64(i) * itv
		ctx.MoveTo(int(r0*math.Sin(angle)), int(r0*math.Cos(angle)))
		ctx.LineTo(int(r1*math.Sin(angle)), int(r1*math.Cos(angle)))
		ctx.Stroke()
	}
	// minute tick
	ctx.LineWidth = 2
	itv = 2 * math.Pi / 60.0
	r0 = float64(r - 5)
	r1 = float64(r)
	for i := 0; i < 60; i++ {
		angle := float64(i) * itv
		ctx.MoveTo(int(r0*math.Sin(angle)), int(r0*math.Cos(angle)))
		ctx.LineTo(int(r1*math.Sin(angle)), int(r1*math.Cos(angle)))
		ctx.Stroke()
	}
	ctx.Restore()
}

func drawNiddles(ctx *canvas.CanvasRenderingContext2D, r int) {
	hourR := float64(r - 30)
	minuteR := float64(r - 20)
	secondR := float64(r - 13)
	now := time.Now()
	hourAngle := (1 - float64(now.Hour()%12)/12) * math.Pi * 2
	minuteAngle := (1 - float64(now.Minute())/60) * math.Pi * 2
	secondAngle := (1 - float64(now.Second())/60) * math.Pi * 2
	rotateAngle := math.Pi
	ctx.Save()
	ctx.Rotate(int(-1 * rotateAngle))

	ctx.StrokeStyle = "black"
	ctx.LineWidth = 4
	ctx.MoveTo(0, 0)
	ctx.LineTo(int(hourR*math.Sin(hourAngle)), int(hourR*math.Cos(hourAngle)))
	ctx.Stroke()

	ctx.LineWidth = 2
	ctx.MoveTo(0, 0)
	ctx.LineTo(int(minuteR*math.Sin(minuteAngle)), int(minuteR*math.Cos(minuteAngle)))
	ctx.Stroke()

	ctx.LineWidth = 1
	ctx.MoveTo(0, 0)
	ctx.LineTo(int(secondR*math.Sin(secondAngle)), int(secondR*math.Cos(secondAngle)))
	ctx.Stroke()
	ctx.Restore()
}

func draw(el *canvas.CanvasElement) {
	// event test
	el.AddEventListener(dom.EvtMousemove, true, func(ev *dom.Event) {
		println("mouse:", ev.MovementX, ev.MovementY)
	})
	// draw
	ctx := el.GetContext2D()
	width := 100
	height := 100
	if el.Width >= width {
		width = el.Width
	} else {
		el.Width = width
	}
	if el.Height >= height {
		height = el.Height
	} else {
		el.Height = height
	}
	startTime := time.Now()
	drawClock := func() {
		ctx.Save()
		ctx.ClearRect(0, 0, width, height)
		ctx.Translate(width/2, height/2)
		radius := int(math.Min(float64(width), float64(height))/2) - 2
		drawClockPane(ctx, radius)
		drawNiddles(ctx, radius)
		ctx.Restore()
	}
	var fn func()
	fn = func() {
		now := time.Now()
		if now.Sub(startTime).Seconds() >= 1 {
			startTime = now
			drawClock()
		}
		vue.NextTick(fn)
	}
	drawClock()
	fn()
}

func init() {
	vue.Component("js-clock", js.M{
		"template": `<canvas></canvas>`,
		"replace":  true,
		"ready": js.MakeFunc(func(this *js.Object, arg []*js.Object) interface{} {
			draw(canvas.NewElement(this.Get("$el")))
			return 0
		}),
	})
}
