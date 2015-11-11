// Package ko implements bindings to KnockoutJS.
// It also has bindings for the Knockout Validation library found on https://github.com/Knockout-Contrib/Knockout-Validation
// Using EnableSecureBinding make KnockoutJS works under CSP environments.
package ko

import (
	"github.com/Archs/js/dom"
	"github.com/Archs/js/utils/property"
	"github.com/gopherjs/gopherjs/js"
)

func ko() *js.Object {
	return js.Global.Get("ko")
}

// Top level getter for ko instance
func Get(key string) *js.Object {
	return ko().Get(key)
}

// Top level setter for ko instance
func Set(key string, value interface{}) {
	ko().Set(key, value)
}

type Observable struct {
	o *js.Object
}

func (o *Observable) ToJS() *js.Object {
	return o.o
}

func ObservableFromJS(o *js.Object) *Observable {
	return &Observable{o}
}

func NewObservable(data ...interface{}) *Observable {
	if len(data) >= 1 {
		return &Observable{ko().Call("observable", data[0])}
	}
	return &Observable{ko().Call("observable")}
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
	// getter
	Get(keyPath string) *js.Object
	// setter
	Set(keyPath string, value interface{})
}

func ViewModelFromJS(o *js.Object) ViewModel {
	vm := NewBaseViewModel()
	vm.FromJS(o)
	return vm
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

func (b *BaseViewModel) Set(keyPath string, value interface{}) {
	obj := property.Get(b.Object, keyPath)
	if obj == js.Undefined {
		// if isArray(value) {
		//  v.Set(key, NewObservableArray(value))
		// } else {
		//  v.Set(key, NewObservable(value))
		// }
		panic("ViewModel has no key: " + keyPath)
	} else {
		obj.Invoke(value)
	}
}

func (b *BaseViewModel) Get(keyPath string) *js.Object {
	obj := property.Get(b.Object, keyPath)
	if obj == js.Undefined {
		return obj
	}
	return obj.Invoke()
}

func (s *Subscription) Dispose() {
	s.Object.Call("dispose")
}

func (o *Observable) Set(data interface{}) {
	o.o.Invoke(data)
}

func (o *Observable) Get() *js.Object {
	return o.o.Invoke()
}

func (o *Observable) Subscribe(fn func(*js.Object)) *Subscription {
	return &Subscription{
		Object: o.o.Call("subscribe", fn),
	}
}

func (o *Observable) Extend(extenderName string, options interface{}) *Observable {
	o.o.Call("extend", js.M{
		extenderName: options,
	})
	return o
}

// ExtenderFunc is used to extend ko.Observable
// 	target is the observable itself to be extended
//  options is the used to config the extender
//
//  this function should return the extended *ko.Observable
type ExtenderFunc func(target *Observable, options *js.Object) (extended *Observable)

// The function takes in the observable itself as the first argument
// and any options in the second argument.
//
// It can then either return the observable or
// return something new like a computed observable that uses the original observable in some way.
func RegisterExtender(name string, fn ExtenderFunc) {
	ko().Get("extenders").Set(name, func(t *js.Object, options *js.Object) *js.Object {
		target := ObservableFromJS(t)
		o := fn(target, options)
		if o == nil {
			panic("ExtenderFunc should return a valid ko.Observable")
		}
		return o.ToJS()
	})
}

// The rateLimit extender, however, causes an observable to suppress and delay change notifications for a specified period of time. A rate-limited observable therefore updates dependencies asynchronously.
//
// The rateLimit extender can be applied to any type of observable, including observable arrays and computed observables. The main use cases for rate-limiting are:
//
//      1. Making things respond after a certain delay
//      2. Combining multiple changes into a single update
//
// when "notifyWhenChangesStop" is true change envent will be fired only after no change event detects anymore.
// "notifyWhenChangesStop" default is false, then it works under "notifyAtFixedRate" mode, at most one change in one timeframe.
func (o *Observable) RateLimit(timeframeMS int, notifyWhenChangesStop ...bool) {
	method := "notifyAtFixedRate"
	if len(notifyWhenChangesStop) >= 1 && notifyWhenChangesStop[0] {
		method = "notifyWhenChangesStop"
	}
	o.Extend("rateLimit", js.M{
		"timeout": timeframeMS,
		"method":  method,
	})
}

// When a computed observable returns a primitive value (a number, string, boolean, or null),
// the dependencies of the observable are normally only notified if the value actually changed.
// However, it is possible to use the built-in notify extender to ensure
// that a computed observable’s subscribers are always notified on an update,
// even if the value is the same.
func (o *Observable) NotifyAlways() {
	o.Extend("notify", "always")
}

// for observable array

func NewObservableArray(data ...interface{}) *Observable {
	if len(data) >= 1 {
		return &Observable{ko().Call("observableArray", data[0])}
	}
	return &Observable{ko().Call("observableArray")}
}

// adds a new item to the end of array
func (o *Observable) IndexOf(data interface{}) int {
	return o.o.Call("indexOf", data).Int()
}

// removes the last value from the array and returns it
func (o *Observable) Pop() *js.Object {
	return o.o.Call("pop")
}

// inserts a new item at the beginning of the array
func (o *Observable) Unshift(data interface{}) {
	o.o.Call("unshift", data)
}

// removes the first value from the array and returns it
func (o *Observable) Shift() *js.Object {
	return o.o.Call("shift")
}

func (o *Observable) Reverse() {
	o.o.Call("reverse")
}

func (o *Observable) Sort() {
	o.o.Call("sort")
}

func (o *Observable) SortFunc(fn func(*js.Object, *js.Object)) {
	o.o.Call("sort", fn)
}

// removes and returns a given number of elements starting from a given index.
// For example,
//      myObservable.splice(1, 3)
// removes three elements starting from index position 1
// (i.e., the 2nd, 3rd, and 4th elements) and returns them as an array.
func (o *Observable) Splice(i, n int) *js.Object {
	return o.o.Call("splice", i, n)
}

func (o *Observable) RemoveAll(items ...interface{}) *js.Object {
	return o.o.Call("removeAll", items...)
}

func (o *Observable) Index(i int) *js.Object {
	return o.Get().Index(i)
}

func (o *Observable) Length() int {
	return o.Get().Length()
}

func (o *Observable) Push(data interface{}) {
	o.o.Call("push", data)
}

func (o *Observable) Remove(item interface{}) *js.Object {
	return o.o.Call("remove", item)
}

func (o *Observable) RemoveFunc(fn func(*js.Object) bool) *js.Object {
	return o.o.Call("remove", fn)
}

// for computed observable

func NewComputed(fn func() interface{}) *Observable {
	return &Observable{ko().Call("computed", fn)}
}

func NewWritableComputed(r func() interface{}, w func(interface{})) *Observable {
	return &Observable{
		ko().Call("computed", js.M{
			"read":  r,
			"write": w,
		}),
	}
}

func (o *Observable) Dispose() {
	o.o.Call("dispose")
}

// Returns the current value of the computed observable without creating a dependency
func (o *Observable) Peek() *js.Object {
	return o.o.Call("peek")
}

// returns true for observables, observable arrays, and all computed observables.
func IsObservable(o interface{}) bool {
	return ko().Call("isObservable", o).Bool()
}

func (o *Observable) IsComputedObservable() bool {
	return ko().Call("isComputed", o.o).Bool()
}

// returns true for writable computed observables only
func (o *Observable) IsWritableObservable() bool {
	return ko().Call("isWritableObservable", o.o).Bool()
}

// ComponentInfo
//  'params' is an object whose key/value pairs are the parameters
//    passed from the component binding or custom element
//
//  'componentInfo.element' is the element the component is being
//    injected into. When createViewModel is called, the template has
//    already been injected into this element, but isn't yet bound.
//
//  'componentInfo.templateNodes' is an array containing any DOM
//    nodes that have been supplied to the component. See below.
type ComponentInfo struct {
	*js.Object
	Element *dom.Element `js:"element"`
	// DOM nodes inside custom element or component will be stripped out (without being bound to any viewmodel)
	// and replaced by the component’s output.
	// However, those DOM nodes aren’t lost: they are remembered, and are supplied to the component as an array
	TemplateNodes *dom.HTMLCollection `js:"templateNodes"`
}

type ComponentVModelConfig struct {
	// Knockout will invoke your constructor once for each instance of the component,
	// producing a separate viewmodel object for each.
	//
	// 'params' is an object whose key/value pairs are the parameters
	// passed from the component binding or custom element.
	Constructor func(params *js.Object) ViewModel
	// If you want to run any setup logic on the associated element before it is bound to the viewmodel,
	// or use arbitrary logic to decide which viewmodel class to instantiate you can use VmCreator
	Creator func(params *js.Object, info *ComponentInfo) ViewModel
}

type ComponentTemplateConfig struct {
	// A string of markup
	Markup string
	// An existing element ID
	Id string
	// An existing element instance
	Instance *dom.Element
}

type Component struct {
	// The component name can be any nonempty string.
	// It’s recommended, but not mandatory, to use lowercase dash-separated strings (such as your-component-name)
	// so that the component name is valid to use as a custom element (such as <your-component-name>).
	Name string
	// ViewModel is optional
	// if set, only one of its struct member should be provided
	// member checking seq: Constructor > Creator
	ViewModel *ComponentVModelConfig
	// Template must be provided
	// and only one of its struct member should be provided
	// member checking seq: Markup > Id > Instance
	Template *ComponentTemplateConfig
	// If your component configuration has a boolean synchronous property,
	// Knockout uses this to determine whether the component is allowed to be loaded
	// and injected synchronously.
	Synchronous bool // The default is false
	// optional sytle
	// style would be embeded in <style></style> tag
	Style string
}

func NewComponent(name string) *Component {
	return &Component{
		Name:      name,
		ViewModel: new(ComponentVModelConfig),
		Template:  new(ComponentTemplateConfig),
	}
}

func (c *Component) viewModel() interface{} {
	if c.ViewModel.Constructor != nil {
		return func(params *js.Object) *js.Object {
			vm := c.ViewModel.Constructor(params)
			if vm != nil {
				return vm.ToJS()
			}
			return nil
		}
	}
	if c.ViewModel.Creator != nil {
		return js.M{
			"createViewModel": func(params *js.Object, info *ComponentInfo) *js.Object {
				vm := c.ViewModel.Creator(params, info)
				if vm != nil {
					return vm.ToJS()
				}
				return nil
			},
		}
	}
	return nil
}

func (c *Component) template() interface{} {
	if c.Template.Markup != "" {
		return c.Template.Markup
	}
	if c.Template.Id != "" {
		return js.M{
			"element": c.Template.Id,
		}
	}
	if c.Template.Instance != nil {
		return js.M{
			"element": c.Template.Instance,
		}
	}
	panic("Template must be provided for ko.Component")
}

func rawRegisterComponent(name string, params js.M) {
	ko().Get("components").Call("register", name, params)
}

func registerComponent(c *Component) {
	if c.Style != "" {
		style := dom.CreateElement("style")
		style.InnerHTML = c.Style
		dom.Body().AppendChild(style)
	}
	rawRegisterComponent(c.Name, js.M{
		"template":    c.template(),
		"viewModel":   c.viewModel(),
		"synchronous": c.Synchronous,
	})
}

// RegisterComponent is an easy form to create KnockoutJS component
//  name is the component name
//  vmCreator is the ViewModel creator with type: func(paramsMap *js.Object, info *ComponentInfo) (vm ViewModel)
//     vmCreator can be nil which means template only component
//     paramsMap is configured like:
//     <ko-uploader params="uploadUrl:'/uploadUrl', text:'Browser', buttonCls:'button round expand', multiple:true"></ko-uploader>
//  template is the html tempalte for the component
//  cssRules would be directly embeded in the final html page, which can be ""
func RegisterComponent(name string, vmCreator func(params *js.Object, info *ComponentInfo) ViewModel, template, cssRules string) {
	c := NewComponent(name)
	c.ViewModel.Creator = vmCreator
	c.Template.Markup = template
	c.Style = cssRules
	registerComponent(c)
}

type BindingContext struct {
	*js.Object
	parent *js.Object `js:"$parent"`
	// $parents
	// This is an array representing all of the parent view models:
	parents *js.Object `js:"$parents"`

	// $parents[0] is the view model from the parent context (i.e., it’s the same as $parent)

	// $parents[1] is the view model from the grandparent context

	// $parents[2] is the view model from the great-grandparent context

	// … and so on.

	// $root
	// This is the main view model object in the root context, i.e., the topmost parent context. It’s usually the object that was passed to ko().applyBindings. It is equivalent to $parents[$parents.length - 1].
	root *js.Object `js:"$root"`

	// $component
	// If you’re within the context of a particular component template, then $component refers to the viewmodel for that component. It’s the component-specific equivalent to $root. In the case of nested components, $component refers to the viewmodel for the closest component.
	component *js.Object `js:"$component"`

	// This is useful, for example, if a component’s template includes one or more foreach blocks in which you wish to refer to some property or function on the component viewmodel rather than on the current data item.

	// $data
	// This is the view model object in the current context. In the root context, $data and $root are equivalent. Inside a nested binding context, this parameter will be set to the current data item (e.g., inside a with: person binding, $data will be set to person). $data is useful when you want to reference the viewmodel itself, rather than a property on the viewmodel. Example:
	data *js.Object `js:"$data"`

	// <ul data-bind="foreach: ['cats', 'dogs', 'fish']">
	//     <li>The value is <span data-bind="text: $data"></span></li>
	// </ul>
	// $index (only available within foreach bindings)

	// This is the zero-based index of the current array entry being rendered by a foreach binding. Unlike the other binding context properties, $index is an observable and is updated whenever the index of the item changes (e.g., if items are added to or removed from the array).

	// $parentContext
	// This refers to the binding context object at the parent level. This is different from $parent, which refers to the data (not binding context) at the parent level. This is useful, for example, if you need to access the index value of an outer foreach item from an inner context (usage: $parentContext.$index). This is undefined in the root context.
	parentContext *js.Object `js:"$parentContext"`

	// $rawData
	// This is the raw view model value in the current context. Usually this will be the same as $data, but if the view model provided to Knockout is wrapped in an observable, $data will be the unwrapped view model, and $rawData will be the observable itself.
	rawData *js.Object `js:"$rawData"`

	// $componentTemplateNodes
	// If you’re within the context of a particular component template, then $componentTemplateNodes is an array containing any DOM nodes that were passed to that component. This makes it easy to build components that receive templates, for example a grid component that accepts a template to define its output rows. For a complete example, see passing markup into components.
	componentTemplateNodes *js.Object `js:"$componentTemplateNodes"`
}

// This is the view model object in the parent context,
// the one immeditely outside the current context.
// In the root context, this is undefined.
func (b *BindingContext) Parent() ViewModel {
	vm := NewBaseViewModel()
	vm.FromJS(b.parent)
	return vm
}

// This is the main view model object in the root context, i.e., the topmost parent context. It’s usually the object that was passed to ko().applyBindings. It is equivalent to $parents[$parents.length - 1].
func (b *BindingContext) Root() ViewModel {
	vm := NewBaseViewModel()
	vm.FromJS(b.root)
	return vm
}

// This is the view model object in the current context. In the root context, $data and $root are equivalent. Inside a nested binding context, this parameter will be set to the current data item (e.g., inside a with: person binding, $data will be set to person). $data is useful when you want to reference the viewmodel itself, rather than a property on the viewmodel.
func (b *BindingContext) Data() ViewModel {
	vm := NewBaseViewModel()
	vm.FromJS(b.data)
	return vm
}

type AllBindings struct {
	*js.Object
}

func (a *AllBindings) Has(name string) bool {
	return a.Call("has", name).Bool()
}

func (a *AllBindings) Get(name string) *js.Object {
	return a.Call("get", name)
}

// callback funtion used in custom bindings
//
//  element
//    The DOM element involved in this binding
//  valueAccessor
//    A JavaScript function that you can call to get the current model property
//    that is involved in this binding. Call this without passing any parameters
//    (i.e., call valueAccessor()) to get the current model property value.
//    To easily accept both observable and plain values, call ko().unwrap on the returned value.
//    as for this go bindings this is already done
//  allBindings
//    A JavaScript object that you can use to access all the model values
//    bound to this DOM element.
//    Call allBindings.get('name') to retrieve the value of the name binding
//    (returns undefined if the binding doesn’t exist);
//    or allBindings.has('name') to determine if the name binding is present
//    for the current element.
//  viewModel
//    This parameter is deprecated in Knockout 3.x.
//    Use bindingContext.$data or bindingContext.$rawData to access the view model instead.
//  bindingContext
//    An object that holds the binding context available to this element’s bindings.
//    This object includes special properties including $parent, $parents,
//    and $root that can be used to access data that is bound against ancestors of this context.
//
// original javascritp signature
//  function(element, valueAccessor, allBindings, viewModel, bindingContext)
type CustomBindingCallback func(element *dom.Element, valueAccessor func() *js.Object, allBindings *AllBindings, viewmodel *js.Object, bindingContext *BindingContext)

func (c CustomBindingCallback) unwrap() CustomBindingCallback {
	return func(element *dom.Element, valueAccessor func() *js.Object, allBindings *AllBindings, viewmodel *js.Object, bindingContext *BindingContext) {
		va := func() *js.Object {
			return unwrap(valueAccessor())
		}
		c(element, va, allBindings, viewmodel, bindingContext)
	}
}

// unwrap observable or plain js.Object to plain *js.Object
//
// Call this without passing any parameters (i.e., call valueAccessor())
// to get the current MODEL PROPERTY VALUE.
// To easily accept BOTH OBSERVABLE AND PLAIN VALUES,
// call ko().unwrap on the returned value.
func unwrap(o *js.Object) *js.Object {
	return ko().Call("unwrap", o)
}

// RegisterCustomBinding
//
// init
//   This will be called when the binding is first applied to an element
//   Set up any initial state, event handlers, etc. here
// update
//   This will be called once when the binding is first applied to an element,
//   and again whenever any observables/computeds that are accessed change
//   Update the DOM element based on the supplied values here.
func RegisterCustomBinding(name string, init, update CustomBindingCallback) {
	if update == nil {
		panic("update callback must be provided for RegisterCustomBinding")
	}
	if init == nil {
		ko().Get("bindingHandlers").Set(name, js.M{
			"update": update.unwrap(),
		})
	} else {
		ko().Get("bindingHandlers").Set(name, js.M{
			"init":   init.unwrap(),
			"update": update.unwrap(),
		})
	}
}

// easy form of CustomBindingCallback
type BindingCallback func(el *dom.Element, valueAccessor func() *js.Object)

func (b BindingCallback) raw() CustomBindingCallback {
	if b == nil {
		return nil
	}
	return func(element *dom.Element, valueAccessor func() *js.Object, allBindings *AllBindings, viewmodel *js.Object, bindingContext *BindingContext) {
		b(element, valueAccessor)
	}
}

// easy form of RegisterCustomBinding
func RegisterBinding(name string, init, update BindingCallback) {
	RegisterCustomBinding(name, init.raw(), update.raw())
}

// In case you’re wondering what the parameters to ko().applyBindings do,
//
// the first parameter says what view model object you want to use with the declarative bindings it activates
//
// Optionally, you can pass a second parameter to define which part of the document you want to search for data-bind attributes.
//
// For example,
//  ko().applyBindings(myViewModel, document.getElementById('someElementId')).
// This restricts the activation to the element with ID someElementId and its descendants, which is useful if you want to have multiple view models and associate each with a different region of the page.
//
// the vm can be nil, indicating direct apply
func ApplyBindings(vm ViewModel, el ...*dom.Element) {
	if vm == nil {
		ko().Call("applyBindings")
		return
	}
	if len(el) < 1 {
		ko().Call("applyBindings", vm.ToJS())
	} else {
		ko().Call("applyBindings", vm.ToJS(), el[0])
	}
}
