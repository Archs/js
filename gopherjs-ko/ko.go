// Package ko implements bindings to KnockoutJS.
// It also has bindings for the Knockout Validation library found on https://github.com/Knockout-Contrib/Knockout-Validation
// Using EnableSecureBinding make KnockoutJS works under CSP environments.
package ko

import "github.com/gopherjs/gopherjs/js"

type Disposable func()

type Observable struct {
	*js.Object
}

type ObservableArray struct {
	*Observable
}

type Computed struct {
	*Observable
}

func (ob *Observable) Set(data interface{}) {
	ob.Object.Invoke(data)
}

func (ob *Observable) Get() *js.Object {
	return ob.Invoke()
}

func (ob *Observable) Subscribe(fn func(*js.Object)) Disposable {
	o := ob.Call("subscribe", fn)
	return func() {
		o.Invoke()
	}
}

func (ob *Observable) Extend(params js.M) *Observable {
	ob.Call("extend", params)
	return ob
}

func (ob *ObservableArray) IndexOf(data interface{}) int {
	return ob.Call("indexOf", data).Int()
}

func (ob *ObservableArray) Pop() *js.Object {
	return ob.Call("pop")
}

func (ob *ObservableArray) Unshift(data interface{}) {
	ob.Call("unshift", data)
}

func (ob *ObservableArray) Shift() *js.Object {
	return ob.Call("shift")
}

func (ob *ObservableArray) Reverse() {
	ob.Call("reverse")
}

func (ob *ObservableArray) Sort() {
	ob.Call("sort")
}

func (ob *ObservableArray) SortFunc(fn func(*js.Object, *js.Object)) {
	ob.Call("sort", fn)
}

func (ob *ObservableArray) Splice(i, n int) *js.Object {
	return ob.Call("splice", i, n)
}

func (ob *ObservableArray) RemoveAll(items ...interface{}) *js.Object {
	return ob.Call("removeAll", items...)
}

func (ob *ObservableArray) Index(i int) *js.Object {
	return ob.Get().Index(i)
}

func (ob *ObservableArray) Length() int {
	return ob.Get().Length()
}

func (ob *ObservableArray) Push(data interface{}) {
	ob.Call("push", data)
}

func (ob *ObservableArray) Remove(item interface{}) *js.Object {
	return ob.Call("remove", item)
}

func (ob *ObservableArray) RemoveFunc(fn func(*js.Object) bool) *js.Object {
	return ob.Call("remove", fn)
}

type ComponentsFuncs struct {
	*js.Object
}

func Components() *ComponentsFuncs {
	return &ComponentsFuncs{
		Object: Global().Get("components"),
	}
}

func (co *ComponentsFuncs) Register(name string, params js.M) {
	co.Call("register", name, params)
}

func Global() *js.Object {
	return js.Global.Get("ko")
}

func NewObservable(data interface{}) *Observable {
	return &Observable{
		Object: Global().Call("observable", data),
	}
}

func NewObservableArray(data interface{}) *ObservableArray {
	return &ObservableArray{&Observable{Global().Call("observableArray", data)}}
}

func NewComputed(fn func() interface{}) *Computed {
	return &Computed{&Observable{Global().Call("computed", fn)}}
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
				Components().Get("defaultLoader").Call("loadTemplate", name, data, callback)
			})
		} else {
			// Unrecognized config format. Let another loader handle it.
			callback(nil)
		}
	}

	Components().Get("loaders").Call("unshift", js.M{
		"loadTemplate": loader,
	})
}

func Unwrap(ob *js.Object) *js.Object {
	return Global().Call("unwrap", ob)
}

func ApplyBindings(model interface{}) {
	Global().Call("applyBindings", model)
}
