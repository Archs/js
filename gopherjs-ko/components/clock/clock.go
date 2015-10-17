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
	w float64
	h float64
	r float64 // radius
	// for needle length
	hourNeedleRatio   float64
	minuteNeedleRatio float64
	secondNeedleRatio float64
	// for mark length
	borderRatio float64
	markRatio   float64
}

func newSimClock(ctx *canvas.Context2D, w, h int) *simClock {
	s := new(simClock)
	s.Context2D = ctx
	s.w = float64(w)
	s.h = float64(h)
	if s.w > s.h {
		s.r = s.h / 2
	} else {
		s.r = s.w / 2
	}
	s.hourNeedleRatio = 0.5
	s.minuteNeedleRatio = 0.7
	s.secondNeedleRatio = 0.8
	s.borderRatio = 0.05
	s.markRatio = 0.05
	return s
}

func (s *simClock) drawPane() {
	s.Save()
	defer s.Restore()
	// outer
	s.StrokeStyle = "black"
	s.LineWidth = s.r * s.borderRatio
	// center point
	s.BeginPath()
	s.Arc(0, 0, s.r*s.borderRatio, 0, 2*math.Pi, false)
	s.Fill()
	// circle
	s.BeginPath()
	s.Arc(0, 0, s.r*(1-s.borderRatio), 0, 2*math.Pi, false)
	s.Stroke()
	// hour marks
	s.LineWidth = s.r * s.markRatio
	iv := math.Pi / 6
	for i := 1; i <= 12; i++ {
		r1 := s.r * (1 - 2*s.markRatio - s.borderRatio)
		r2 := s.r * (1 - s.borderRatio)
		r3 := s.r * (1 - 4*s.markRatio - s.borderRatio)
		angle := iv*float64(i) - math.Pi/2
		x1 := r1 * math.Cos(angle)
		y1 := r1 * math.Sin(angle)
		x2 := r2 * math.Cos(angle)
		y2 := r2 * math.Sin(angle)
		x3 := r3 * math.Cos(angle)
		y3 := r3 * math.Sin(angle)
		s.BeginPath()
		s.MoveTo(x1, y1)
		s.LineTo(x2, y2)
		s.Stroke()
		// s.Rotate(-1 * math.Pi / 2)
		s.Font = "bold 20px Arial"
		s.TextAlign = "center"
		s.TextBaseline = "middle"
		// s.FillStyle = "#FFF"
		if s.r > 90 {
			s.FillText(fmt.Sprintf("%d", i), x3, y3, 20)
		}
	}
	// minutes marks
	iv = math.Pi / 30
	s.LineWidth = s.r * s.markRatio * 0.5
	for i := 0; i < 60; i++ {
		s.Rotate(iv)
		s.BeginPath()
		s.MoveTo(s.r*(1-2.5*s.markRatio), 0)
		s.LineTo(s.r*(1-s.borderRatio), 0)
		s.Stroke()
	}
}

func (s *simClock) drawCZ() {
	s.Save()
	defer s.Restore()
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
}

func (s *simClock) drawNeedle(t time.Time) {
	s.Save()
	defer s.Restore()
	needle := func(angle float64, lineWidth float64, ratio float64, color string) {
		s.BeginPath()
		s.StrokeStyle = color
		s.LineWidth = lineWidth
		r := s.r * ratio
		angle = angle - math.Pi/2
		x := r * math.Cos(angle)
		y := r * math.Sin(angle)
		s.MoveTo(0, 0)
		s.LineTo(x, y)
		s.Stroke()
	}
	// houre
	angleHour := float64(t.Hour()) / 6.0 * math.Pi
	needle(angleHour, 5.0, s.hourNeedleRatio, "black")
	// minute
	angleMinute := float64(t.Minute()) / 30.0 * math.Pi
	needle(angleMinute, 3.0, s.minuteNeedleRatio, "black")
	// second
	angleSecond := (float64(t.Second()) + float64(t.Nanosecond())/1000000000.0) / 30.0 * math.Pi
	needle(angleSecond, 1.0, s.secondNeedleRatio, "red")
}

func (s *simClock) drawDay(t time.Time) {
	day := t.Day()
	r := s.r * 0.45
	if r > 30 {
		s.FillStyle = "black"
		s.FillRect(r, -11, 25, 20)
		s.FillStyle = "white"
		s.Font = "20px serif"
		s.FillText(fmt.Sprintf("%d", day), r+3, 6, 30)
	}
}

func (s *simClock) draw(t time.Time) {
	s.Save()
	defer s.Restore()
	s.ClearRect(0, 0, s.w, s.h)
	s.Translate(s.w/2, s.h/2)
	s.StrokeStyle = "black"
	s.LineWidth = 2
	// s.drawCZ()
	s.drawPane()
	s.drawNeedle(t)
	s.drawDay(t)
}

func registerClock() {
	ko.RegisterComponent("clock", func(params *js.Object, info *ko.ComponentInfo) ko.ViewModel {
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
		clock := newSimClock(c.GetContext2D(), c.Width, c.Height)
		clock.draw(time.Now())
		go func() {
			for t := range time.Tick(time.Millisecond * 100) {
				clock.draw(t)
			}
		}()
		return nil
	},
		"<canvas></canvas>",
		"",
	)
}

func init() {
	registerClock()
}
