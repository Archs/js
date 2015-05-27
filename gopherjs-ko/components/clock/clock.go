package clock

import (
	"github.com/Archs/js/canvas"
	"github.com/Archs/js/dom"
	"github.com/Archs/js/gopherjs-ko"
	"github.com/gopherjs/gopherjs/js"
	"time"
)

type ctrl struct {
	*js.Object
	msg   *ko.Observable `js:"msg"`
	click func()         `js:"click"`
}

func newCtrl() *ctrl {
	c := new(ctrl)
	c.Object = js.Global.Get("Object").New()
	c.msg = ko.NewObservable()
	return c
}

func registerTest() {
	ko.Components().Register("test", js.M{
		"template": "<h2>Clock</h2><button data-bind='click:click'>Click</button><span data-bind='text:msg'></span>",
		"viewModel": func(p *js.Object) *ctrl {
			println("viewModel", p)
			c := newCtrl()
			c.msg.Set("A Message From Clock")
			c.click = func() {
				println("click called")
				c.msg.Set(time.Now().String())
			}
			return c
		},
	})
}

func registerClock() {
	ko.Components().Register("clock", js.M{
		"template": "<canvas></canvas>",
		"viewModel": js.M{
			"createViewModel": func(params, info *js.Object) {
				el := dom.Wrap(info.Get("element"))
				el = el.QuerySelector("canvas")
				c := canvas.New(el.Object)
				c.Width = 20
				c.Height = 20
				ctx := c.GetContext2D()
				ctx.FillStyle = "red"
				ctx.FillRect(0, 0, 20, 20)
				println(params, info, el, c, ctx)
			},
		},
	})
}

func init() {
	registerClock()
}
