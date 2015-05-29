// Package clock implements a KnockoutJS component: clock.
//
// Use it like this in html files after import this package:
// 	<clock params="width:300, height:300"></clock>
// 	<span data-bind="component: {name:'clock',params:{width:200,height:150}}"></span>
package clock

import (
	"fmt"
	"github.com/Archs/js/canvas"
	"github.com/Archs/js/dom"
	"github.com/Archs/js/gopherjs-ko"
	"github.com/gopherjs/gopherjs/js"
	"math"
	"time"
)

type simClock struct {
	*canvas.Context2D
	w int
	h int
	r int // radius
}

func (s *simClock) drawPane() {
	s.Save()
	// outer
	s.StrokeStyle = "black"
	s.LineWidth = 10
	s.BeginPath()
	s.Arc(0, 0, s.r-10, 0, 2*math.Pi, false)
	s.Stroke()
	// hour marks
	s.LineWidth = 6
	iv := math.Pi / 6
	s.Save()
	for i := 1; i <= 12; i++ {
		// s.Rotate(iv)
		s.BeginPath()
		r1 := float64(s.r - 25)
		r2 := float64(s.r - 10)
		r3 := float64(s.r - 40)
		angle := iv*float64(i) - math.Pi/2
		x1 := r1 * math.Cos(angle)
		y1 := r1 * math.Sin(angle)
		x2 := r2 * math.Cos(angle)
		y2 := r2 * math.Sin(angle)
		x3 := r3 * math.Cos(angle)
		y3 := r3 * math.Sin(angle)
		s.MoveTo(x1, y1)
		s.LineTo(x2, y2)
		s.Stroke()
		s.Save()
		// s.Rotate(-1 * math.Pi / 2)
		s.Font = "bold 20px Arial"
		s.TextAlign = "center"
		s.TextBaseline = "middle"
		// s.FillStyle = "#FFF"
		s.FillText(fmt.Sprintf("%d", i), x3, y3, 20)
		s.Restore()
	}
	s.Restore()
	// minutes marks
	iv = math.Pi / 30
	s.LineWidth = 3
	for i := 0; i < 60; i++ {
		s.Rotate(iv)
		s.BeginPath()
		s.MoveTo(s.r-20, 0)
		s.LineTo(s.r-10, 0)
		s.Stroke()
	}
	// end
	s.Restore()
}

func (s *simClock) drawCZ() {
	s.Save()
	// begin
	s.SetLineDash(1, 1)
	// x
	s.BeginPath()
	s.MoveTo(-1*s.r, 0)
	s.LineTo(s.r, 0)
	s.Stroke()
	// y
	s.BeginPath()
	s.MoveTo(0, s.r)
	s.LineTo(0, -1*s.r)
	s.Stroke()
	// end
	s.Restore()
}

func (s *simClock) drawNeedle(t time.Time) {
	s.Save()
	needle := func(angle float64, lineWidth int, ratio float64, color string) {
		s.BeginPath()
		s.StrokeStyle = color
		s.LineWidth = lineWidth
		r := float64(s.r) * ratio
		angle = angle - math.Pi/2
		x := r * math.Cos(angle)
		y := r * math.Sin(angle)
		s.MoveTo(0, 0)
		s.LineTo(x, y)
		s.Stroke()
	}
	// houre
	angleHour := float64(t.Hour()) / 6.0 * math.Pi
	needle(angleHour, 5, 0.5, "black")
	// minute
	angleMinute := float64(t.Minute()) / 30.0 * math.Pi
	needle(angleMinute, 3, 0.7, "black")
	// second
	angleSecond := (float64(t.Second()) + float64(t.Nanosecond())/1000000000.0) / 30.0 * math.Pi
	needle(angleSecond, 1, 0.8, "red")
	// end
	s.Restore()
}

func (s *simClock) drawDay(t time.Time) {
	day := t.Day()
	r := float64(s.r) * 0.45
	s.FillStyle = "black"
	s.FillRect(r, -11, 25, 20)
	s.FillStyle = "white"
	s.Font = "20px serif"
	s.FillText(fmt.Sprintf("%d", day), r+3, 6, 30)
}

func (s *simClock) draw(t time.Time) {
	s.Save()
	s.ClearRect(0, 0, s.w, s.h)
	s.Translate(s.w/2, s.h/2)
	s.StrokeStyle = "black"
	s.LineWidth = 2
	if s.w > s.h {
		s.r = s.h / 2
	} else {
		s.r = s.w / 2
	}
	// s.drawCZ()
	s.drawPane()
	s.drawNeedle(t)
	s.drawDay(t)
	s.Restore()

	s.Save()
	s.Translate(s.w/2, s.h/2)
	s.Font = "20px serif"
	s.FillText(t.Format("15:04:05"), -35, 50, 100)
	s.Restore()
}

func registerClock() {
	ko.Components().Register("clock", js.M{
		"template": "<canvas></canvas>",
		"viewModel": js.M{
			"createViewModel": func(params, info *js.Object) {
				el := dom.Wrap(info.Get("element"))
				c := canvas.New(el.QuerySelector("canvas").Object)
				c.Width = 200
				c.Height = 200
				if params.Get("width") != js.Undefined {
					c.Width = params.Get("width").Int()
				}
				if params.Get("height") != js.Undefined {
					c.Height = params.Get("height").Int()
				}
				if c.Width > c.Height {
					c.Width = c.Height
				} else {
					c.Height = c.Width
				}
				if c.Width < 120 {
					el.InnerHTML = "WARNING! clock width/heigth too small"
					println("WARNING! clock width/heigth too small")
					return
				}
				clock := &simClock{
					Context2D: c.GetContext2D(),
					w:         c.Width,
					h:         c.Height,
				}
				clock.draw(time.Now())
				go func() {
					for t := range time.Tick(time.Millisecond * 100) {
						clock.draw(t)
					}
				}()
			},
		},
	})
}

func init() {
	registerClock()
}
