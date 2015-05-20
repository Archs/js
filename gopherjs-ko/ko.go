// Package ko implements bindings to KnockoutJS.
// It also has bindings for the Knockout Validation library found on https://github.com/Knockout-Contrib/Knockout-Validation
// Using EnableSecureBinding make KnockoutJS works under CSP environments.
package ko

import "github.com/gopherjs/gopherjs/js"

type Disposer func()

type Observable struct {
	o *js.Object
}

type ObservableArray struct {
	*Observable
}

type Computed struct {
	*Observable
}

func (ob *Observable) Set(data interface{}) {
	ob.o.Invoke(data)
}

func (ob *Observable) Get() *js.Object {
	return ob.o.Invoke()
}

func (ob *Observable) Subscribe(fn func(*js.Object)) Disposer {
	o := ob.o.Call("subscribe", fn)
	return func() {
		o.Invoke()
	}
}

func (ob *Observable) Extend(params js.M) *Observable {
	ob.o.Call("extend", params)
	return ob
}

// The rateLimit extender, however, causes an observable to suppress and delay change notifications for a specified period of time. A rate-limited observable therefore updates dependencies asynchronously.
//
// The rateLimit extender can be applied to any type of observable, including observable arrays and computed observables. The main use cases for rate-limiting are:
//
// 		1. Making things respond after a certain delay
// 		2. Combining multiple changes into a single update
//
// fixedRate default is true
func (ob *Observable) RateLimit(timeout int, fixedRate ...bool) {
	method := "notifyAtFixedRate"
	if len(fixedRate) >= 1 && !fixedRate[0] {
		method = "notifyWhenChangesStop"
	}
	ob.Extend(js.M{
		"rateLimit": js.M{
			"timeout": timeout,
			"method":  method,
		},
	})
}

// When a computed observable returns a primitive value (a number, string, boolean, or null),
// the dependencies of the observable are normally only notified if the value actually changed.
// However, it is possible to use the built-in notify extender to ensure
// that a computed observable’s subscribers are always notified on an update,
// even if the value is the same.
func (ob *Observable) NotifyAlways() {
	ob.Extend(js.M{
		"notify": "always",
	})
}

func (ob *ObservableArray) IndexOf(data interface{}) int {
	return ob.o.Call("indexOf", data).Int()
}

func (ob *ObservableArray) Pop() *js.Object {
	return ob.o.Call("pop")
}

func (ob *ObservableArray) Unshift(data interface{}) {
	ob.o.Call("unshift", data)
}

func (ob *ObservableArray) Shift() *js.Object {
	return ob.o.Call("shift")
}

func (ob *ObservableArray) Reverse() {
	ob.o.Call("reverse")
}

func (ob *ObservableArray) Sort() {
	ob.o.Call("sort")
}

func (ob *ObservableArray) SortFunc(fn func(*js.Object, *js.Object)) {
	ob.o.Call("sort", fn)
}

func (ob *ObservableArray) Splice(i, n int) *js.Object {
	return ob.o.Call("splice", i, n)
}

func (ob *ObservableArray) RemoveAll(items ...interface{}) *js.Object {
	return ob.o.Call("removeAll", items...)
}

func (ob *ObservableArray) Index(i int) *js.Object {
	return ob.Get().Index(i)
}

func (ob *ObservableArray) Length() int {
	return ob.Get().Length()
}

func (ob *ObservableArray) Push(data interface{}) {
	ob.o.Call("push", data)
}

func (ob *ObservableArray) Remove(item interface{}) *js.Object {
	return ob.o.Call("remove", item)
}

func (ob *ObservableArray) RemoveFunc(fn func(*js.Object) bool) *js.Object {
	return ob.o.Call("remove", fn)
}

type ComponentsFuncs struct {
	o *js.Object
}

func Components() *ComponentsFuncs {
	return &ComponentsFuncs{
		o: ko().Get("components"),
	}
}

func (co *ComponentsFuncs) Register(name string, params js.M) {
	co.o.Call("register", name, params)
}

func ko() *js.Object {
	return js.Global.Get("ko")
}

func NewObservable(data interface{}) *Observable {
	return &Observable{
		o: ko().Call("observable", data),
	}
}

func NewObservableArray(data interface{}) *ObservableArray {
	return &ObservableArray{&Observable{ko().Call("observableArray", data)}}
}

func NewComputed(fn func() interface{}) *Computed {
	return &Computed{&Observable{ko().Call("computed", fn)}}
}

// RegisterURLTemplateLoader register a new template loader which can be used to load
// template files from a webserver.
// To use it you need to pass a map with a `url` key as template argument to your component:
//   "template":  js.M{"url": "form.html"}
// This loader requires jQuery.
func RegisterURLTemplateLoader() {
	loader := func(name string, config *js.Object, callback func(*js.Object)) {
		url := config.Get("url")
		if url != nil && url != js.Undefined {
			// Some browsers are caching these requests too aggressively
			urlStr := url.String()
			urlStr += "?_=" + js.Global.Call("eval", `Date.now()`).String()

			js.Global.Get("jQuery").Call("get", urlStr, func(data *js.Object) {
				// We need an array of DOM nodes, not a string.
				// We can use the default loader to convert to the
				// required format.
				Components().o.Get("defaultLoader").Call("loadTemplate", name, data, callback)
			})
		} else {
			// Unrecognized config format. Let another loader handle it.
			callback(nil)
		}
	}

	Components().o.Get("loaders").Call("unshift", js.M{
		"loadTemplate": loader,
	})
}

func Unwrap(ob *js.Object) *js.Object {
	return ko().Call("unwrap", ob)
}

func ApplyBindings(model interface{}) {
	ko().Call("applyBindings", model)
}
