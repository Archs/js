// Package clock implements a KnockoutJS component: clock.
//
// Use it like this in html files after import this package:
// 	<clock params="width:300, height:300"></clock>
// 	<span data-bind="component: {name:'clock',params:{width:200,height:150}}"></span>
package clock

import (
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
	s.LineWidth = 2
	s.BeginPath()
	s.Arc(0, 0, s.r-2, 0, 2*math.Pi, false)
	s.Stroke()
	// hour needle
	s.LineWidth = 6
	iv := math.Pi / 6
	s.Save()
	for i := 0; i < 12; i++ {
		s.Rotate(iv)
		s.BeginPath()
		s.MoveTo(s.r-7, 0)
		s.LineTo(s.r-1, 0)
		s.Stroke()
	}
	s.Restore()
	// minutes needle
	iv = math.Pi / 30
	s.LineWidth = 3
	for i := 0; i < 60; i++ {
		s.Rotate(iv)
		s.BeginPath()
		s.MoveTo(s.r-5, 0)
		s.LineTo(s.r-1, 0)
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
	s.Rotate(-1 * math.Pi / 2)
	// houre
	s.Save()
	angleHour := float32(t.Hour()) / 12.0 * math.Pi
	s.LineWidth = 5
	s.Rotate(angleHour)
	s.MoveTo(-15, 0)
	s.LineTo(s.r-30, 0)
	s.Stroke()
	s.Restore()
	// minute
	angleMinute := float32(t.Minute()) / 30.0 * math.Pi
	s.LineWidth = 2
	s.Rotate(angleMinute)
	s.MoveTo(-25, 0)
	s.LineTo(s.r-20, 0)
	s.Stroke()
	// end
	s.Restore()
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
					for t := range time.Tick(time.Second) {
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
