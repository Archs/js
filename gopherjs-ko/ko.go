// Package ko implements bindings to KnockoutJS.
// It also has bindings for the Knockout Validation library found on https://github.com/Knockout-Contrib/Knockout-Validation
// Using EnableSecureBinding make KnockoutJS works under CSP environments.
package ko

import (
	"github.com/Archs/js/dom"
	"github.com/Archs/js/utils/property"
	"github.com/gopherjs/gopherjs/js"
)

type Observable struct {
	o *js.Object
}

type ObservableArray struct {
	*Observable
}

type Computed struct {
	*Observable
}

type WritableComputed struct {
	*Computed
}

type Subscription struct {
	*js.Object
}

// ViewModel is used to wrap ko vm object
type ViewModel interface {
	// return the vm for js side
	ToJS() *js.Object
	// set the real vm from the js side
	FromJS(*js.Object)
}

type BaseViewModel struct {
	*js.Object
}

func NewBaseViewModel() *BaseViewModel {
	return &BaseViewModel{
		Object: js.Global.Get("Object").New(),
	}
}

func (b *BaseViewModel) ToJS() *js.Object {
	return b.Object
}

func (b *BaseViewModel) FromJS(vm *js.Object) {
	b.Object = vm
}

func (v *BaseViewModel) Set(keyPath string, value interface{}) {
	obj := property.Get(v.Object, keyPath)
	if obj == js.Undefined {
		// if isArray(value) {
		// 	v.Set(key, NewObservableArray(value))
		// } else {
		// 	v.Set(key, NewObservable(value))
		// }
		panic("ViewModel has no key: " + keyPath)
	} else {
		obj.Invoke(value)
	}
}

func (v *BaseViewModel) Get(keyPath string) *js.Object {
	obj := property.Get(v.Object, keyPath)
	if obj == js.Undefined {
		return obj
	}
	return obj.Invoke()
}

type Disposer func()

func (s *Subscription) Dispose() {
	s.Object.Call("dispose")
}

func (ob *Observable) Set(data interface{}) {
	ob.o.Invoke(data)
}

func (ob *Observable) Get() *js.Object {
	return ob.o.Invoke()
}

func (ob *Observable) Subscribe(fn func(*js.Object)) *Subscription {
	o := ob.o.Call("subscribe", fn)
	return &Subscription{
		Object: o,
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
// when "notifyWhenChangesStop" is true change envent will be fired only after no change event detects anymore.
// "notifyWhenChangesStop" default is false, then it works under "notifyAtFixedRate" mode, at most one change in one timeframe.
func (ob *Observable) RateLimit(timeframeMS int, notifyWhenChangesStop ...bool) {
	method := "notifyAtFixedRate"
	if len(notifyWhenChangesStop) >= 1 && notifyWhenChangesStop[0] {
		method = "notifyWhenChangesStop"
	}
	ob.Extend(js.M{
		"rateLimit": js.M{
			"timeout": timeframeMS,
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

// adds a new item to the end of array
func (ob *ObservableArray) IndexOf(data interface{}) int {
	return ob.o.Call("indexOf", data).Int()
}

// removes the last value from the array and returns it
func (ob *ObservableArray) Pop() *js.Object {
	return ob.o.Call("pop")
}

// inserts a new item at the beginning of the array
func (ob *ObservableArray) Unshift(data interface{}) {
	ob.o.Call("unshift", data)
}

// removes the first value from the array and returns it
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

// removes and returns a given number of elements starting from a given index.
// For example,
// 		myObservableArray.splice(1, 3)
// removes three elements starting from index position 1
// (i.e., the 2nd, 3rd, and 4th elements) and returns them as an array.
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

type ComponentManager struct {
	o *js.Object
}

func Components() *ComponentManager {
	return &ComponentManager{
		o: ko().Get("components"),
	}
}

func (co *ComponentManager) rawRegister(name string, params js.M) {
	co.o.Call("register", name, params)
}

// - 'params' is an object whose key/value pairs are the parameters
//   passed from the component binding or custom element
// - 'componentInfo.element' is the element the component is being
//   injected into. When createViewModel is called, the template has
//   already been injected into this element, but isn't yet bound.
// - 'componentInfo.templateNodes' is an array containing any DOM
//   nodes that have been supplied to the component. See below.
type ComponentInfo struct {
	*js.Object
	Element *dom.Element `js:"element"`
}

// Register is an easy form to create KnockoutJS component
//  name is the component name
//  vmCreator is the ViewModel creator with type: func(paramsMap *js.Object, info *ComponentInfo) (vm ViewModel)
// 	   vmCreator can be nil which means template only component
//     paramsMap is configured like:
//     <ko-uploader params="uploadUrl:'/uploadUrl', text:'Browser', buttonCls:'button round expand', multiple:true"></ko-uploader>
//  template is the html tempalte for the component
//  cssRules would be directly embeded in the final html page, which can be ""
func (co *ComponentManager) Register(name string, vmCreator func(params *js.Object, info *ComponentInfo) ViewModel, template, cssRules string) {
	// embed the cssRules
	if cssRules != "" {
		style := dom.CreateElement("style")
		style.InnerHTML = cssRules
		dom.Body().AppendChild(style)
	}
	// template only component
	if vmCreator == nil {
		co.rawRegister(name, js.M{
			"template": template,
		})
		return
	}
	// register the component
	co.rawRegister(name, js.M{
		"viewModel": js.M{
			"createViewModel": func(params *js.Object, info *ComponentInfo) *js.Object {
				vm := vmCreator(params, info)
				return vm.ToJS()
			},
		},
		"template": template,
	})
}

func ko() *js.Object {
	return js.Global.Get("ko")
}

func NewObservable(data ...interface{}) *Observable {
	if len(data) >= 1 {
		return &Observable{ko().Call("observable", data[0])}
	}
	return &Observable{ko().Call("observable")}
}

func NewObservableArray(data ...interface{}) *ObservableArray {
	if len(data) >= 1 {
		return &ObservableArray{&Observable{ko().Call("observableArray", data[0])}}
	}
	return &ObservableArray{&Observable{ko().Call("observableArray")}}
}

func NewComputed(fn func() interface{}) *Computed {
	return &Computed{&Observable{ko().Call("computed", fn)}}
}

func NewWritableComputed(r func() interface{}, w func(interface{})) *WritableComputed {
	return &WritableComputed{
		&Computed{
			&Observable{
				ko().Call("computed", js.M{
					"read":  r,
					"write": w,
				}),
			},
		},
	}
}

func (ob *Computed) Dispose() {
	ob.o.Call("dispose")
}

// Returns the current value of the computed observable without creating a dependency
func (ob *Computed) Peek() *js.Object {
	return ob.o.Call("peek")
}

// returns true for observables, observable arrays, and all computed observables.
func IsObservable(data interface{}) bool {
	return ko().Call("isObservable", data).Bool()
}

func IsComputed(data interface{}) bool {
	return ko().Call("isComputed", data).Bool()
}

// returns true for observables, observable arrays, and writable computed observables
func IsWritableObservable(data interface{}) bool {
	return ko().Call("isWritableObservable", data).Bool()
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

// In case you’re wondering what the parameters to ko.applyBindings do,
//
// the first parameter says what view model object you want to use with the declarative bindings it activates
//
// Optionally, you can pass a second parameter to define which part of the document you want to search for data-bind attributes.
//
// For example,
// 	ko.applyBindings(myViewModel, document.getElementById('someElementId')).
// This restricts the activation to the element with ID someElementId and its descendants, which is useful if you want to have multiple view models and associate each with a different region of the page.
func ApplyBindings(vm ViewModel, el ...*dom.Element) {
	if len(el) < 1 {
		ko().Call("applyBindings", vm.ToJS())
	} else {
		ko().Call("applyBindings", vm.ToJS(), el[0])
	}
}
