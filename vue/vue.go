// Package vue offers a naive vue.js bindings in gopherjs.
package vue

import (
	"github.com/gopherjs/gopherjs/js"
)

var (
	vue = js.Global.Get("Vue")
)

var (
	// config fields
	// 	{
	//   // print stack trace for warnings?
	//   debug: true,
	//   // attribute prefix for directives
	//   prefix: 'v-',
	//   // interpolation delimiters
	//   // for HTML interpolations, add
	//   // 1 extra outer-most character.
	//   delimiters: ['{{', '}}'],
	//   // suppress warnings?
	//   silent: false,
	//   // interpolate mustache bindings?
	//   interpolate: true,
	//   // use async updates (for directives & watchers)?
	//   async: true,
	//   // allow altering observed Array's prototype chain?
	//   proto: true
	// }
	Config = vue.Get("config")
)

// struct literal not workable yet 2015.1.30
// type VueOption struct {
// 	js.Object
// 	El   string `js:"el"`
// 	Data js.M   `js:"data"`
// 	//混合进vue实例的方法，你可以直接访问VM实例上的方法，
// 	//或者在表达式里面直接使用。每个函数的上下文关系this都绑定在Vue实例上。
// 	Methods js.M `js:"methods"`
// }

type Vue struct {
	js.Object
	El            string    `js:"$el"`
	Data          js.Object `js:"$data"`
	Options       js.M      `js:"$options"`
	Parent        js.Object `js:"$parent"`
	Root          js.Object `js:"$root"`
	SubComponents js.Object `js:"$"`  // v-ref, to ref sub component from parent
	Doms          js.Object `js:"$$"` // v-el, expose dom element to Vue instance
}

func Directive(name string, opt js.M) {
	vue.Call("directive", name, opt)
}

func Extend(opt js.M) js.Object {
	return vue.Call("extend", opt)
}

func RegisterComponent(name string, cpnt js.Object) {
	vue.Call("component", name, cpnt)
}

func Component(name string, opt js.M) {
	vue.Call("component", name, opt)
}

func Filter(name string, fn func(js.Object) js.Object) {
	vue.Call("filter", name, fn)
}

// Vue.partial( id, [definition] )
// id String
// definition String | Node optional
// Register or retrieve a global partial.
// The definition can be a template string, a querySelector that starts with #,
// a DOM element (whose innerHTML will be used as the template string), or a DocumentFragment.
func Partial(name, definition string) {
	vue.Call("partial", name, definition)
}

// Vue.transition( id, [definition] )
// id String
// definition Object optional
// Register or retrieve a global JavaScript transition effect definition.
// For more details see the guide for JavaScript Transitions.
// Example:
// 		Vue.transition('fade', {
// 		  beforeEnter: function (el) {
// 		    // a synchronous function called right before the
// 		    // element is inserted into the document.
// 		    // you can do some pre-styling here to avoid
// 		    // FOC (flash of content).
// 		  },
// 		  enter: function (el, done) {
// 		    // element is already inserted into the DOM
// 		    // call done when animation finishes.
// 		    $(el)
// 		      .css('opacity', 0)
// 		      .animate({ opacity: 1 }, 1000, done)
// 		    // optionally return a "cancel" function
// 		    // to clean up if the animation is cancelled
// 		    return function () {
// 		      $(el).stop()
// 		    }
// 		  },
// 		  leave: function (el, done) {
// 		    // same as enter
// 		    $(el).animate({ opacity: 0 }, 1000, done)
// 		    return function () {
// 		      $(el).stop()
// 		    }
// 		  }
// 		})
// Then you can use it by providing the transition id to v-transition. Note this has higher priority than CSS transitions.
// <p v-transition="fade"></p>
func Transition(name string, definition js.M) {
	vue.Call("transition", name, definition)
}

// Vue.nextTick( callback )
// callback Function
// Vue.js batches view updates and executes them all asynchronously. It uses requestAnimationFrame if available and falls back to setTimeout(fn, 0). This method calls the callback after the next view update, which can be useful when you want to wait until the view has been updated.
func NextTick(callback func()) {
	vue.Call("nextTick", callback)
}

// struct literal not workable yet, using js.M to do initialization
// 		vm := vue.New(js.M{
// 			"el": "#demo",
// 			"data": js.M{
// 				"title": "todos",
// 				"todos": []js.M{
// 					js.M{
// 						"done":    true,
// 						"content": "Learn JavaScript",
// 					},
// 					js.M{
// 						"done":    false,
// 						"content": "Learn Vue.js",
// 					},
// 				},
// 			},
// 		})
func New(opts js.M) *Vue {
	vm := vue.New(opts)
	return &Vue{
		Object: vm,
	}
}

func (v *Vue) Directive(name string, opt js.M) *Vue {
	v.Call("directive", name, opt)
	return v
}

func (v *Vue) Component(name string, opt js.M) *Vue {
	v.Call("component", name, opt)
	return v
}

func (v *Vue) Filter(name string, fn interface{}) *Vue {
	v.Call("filter", name, fn)
	return v
}

type Unwatcher func()

// 你可以监控Vue实例上的数据。注意所有的监控回调是异步的。另外的，修改值再一个事件循环中批量进行的。这就是说在一个事件循环中一个值修改了多次，回调函数只会带最新的值调用一次。
// vm.$watch( expression, callback, [deep, immediate] )
func (v *Vue) Watch(expression string, callback func(newVal, oldVal js.Object), deepWatch bool) Unwatcher {
	obj := v.Call("$watch", expression, callback, deepWatch)
	return func() {
		obj.Invoke()
	}
}

// vm.$eval( expression )
// expression String
// 计算可以包含过滤器的表达式
func (v *Vue) Eval(expression string) js.Object {
	return v.Call("$eval", expression)
}

// Events
// 每个vm也是一个时间触发器。当你有多个嵌套的VM，你可以使用事件系统在它们之间沟通。

// vm.$dispatch( event, [args…] )
// event String
// args… optional
// 从当前vm分发一个消息到它的父元素。如果回调函数返回false。将会停止传播。
func (v *Vue) Dispatch(event string, args interface{}) {
	v.Call("$dispatch", event, args)
}

// vm.$broadcast( event, [args…] )
// event String
// args… optional
// 给所以的子VM广播一条消息，如果返回false，就不在继续往下广播。
func (v *Vue) Broadcast(event string, args interface{}) {
	v.Call("$broadcast", event, args)
}

// vm.$emit( event, [args…] )
// event String
// args… optional
// 触发一条消息给自己。
func (v *Vue) Emit(event string, args interface{}) {
	v.Call("$emit", event, args)
}

type EventCallback func(args interface{})

// vm.$on( event, callback )
// event String
// callback Function
// 在当前的VM上监听消息。
func (v *Vue) On(event string, cb EventCallback) {
	v.Call("$on", event, cb)
}

// vm.$once( event, callback )
// event String
// callback Function
// 监听一次性消息。
func (v *Vue) Once(event string, cb EventCallback) {
	v.Call("$once", event, cb)
}

// vm.$off( [event, callback] )
// event String optional
// callback Function optional
// 如果没有参数，就停止接监听一切消息；如果只有消息给出，删除所有的回调函数；如果既有消息又有回调，就只删除指定的这个回调。
func (v *Vue) OffEvent(event string) {
	v.Call("$off", event)
}
func (v *Vue) OffAllEvent() {
	v.Call("$off")
}

// DOM
// All vm DOM manipulation methods work like their jQuery counterparts - except they also trigger Vue.js transitions if there are any declared on vm’s $el. For more details on transitions see Adding Transition Effects.

// vm.$appendTo( element|selector, [callback] )
// element HTMLElement | selector String
// callback Function optional
// Append the vm’s $el to target element. The argument can be either an element or a querySelector string.
func (v *Vue) AppendTo(elementOrselector string) {
	v.Call("$appendTo", elementOrselector)
}

// vm.$before( element|selector, [callback] )
// element HTMLElement | selector String
// callback Function optional
// Insert the vm’s $el before target element.
func (v *Vue) Before(elementOrselector string) {
	v.Call("$before", elementOrselector)
}

// vm.$after( element|selector, [callback] )
// element HTMLElement | selector String
// callback Function optional
// Insert the vm’s $el after target element.
func (v *Vue) After(elementOrselector string) {
	v.Call("$after", elementOrselector)
}

// vm.$remove( [callback] )
// callback Function optional
// Remove the vm’s $el from the DOM.
func (v *Vue) Remove() {
	v.Call("$remove")
}

// Lifecycle

// vm.$mount( [element|selector] )
// element HTMLElement | selector String optional
// If the Vue instance didn’t get an el option at instantiation, you can manually call $mount() to assign an element to it and start the compilation. If no argument is provided, an empty <div> will be automatically created. Calling $mount() on an already mounted instance will have no effect. The method returns the instance itself so you can chain other instance methods after it.
func (v *Vue) Mount(elementOrselector string) {
	v.Call("$mount", elementOrselector)
}

// vm.$destroy( [remove] )
// remove Boolean optional
// Completely destroy a vm. Clean up its connections with other existing vms, unbind all its directives and remove its $el from the DOM. Also, all $on and $watch listeners will be automatically removed.
func (v *Vue) Destroy(remove bool) {
	v.Call("$destroy", remove)
}

// vm.$compile( element )
// element HTMLElement
// Partially compile a piece of DOM (Element or DocumentFragment). The method returns a decompile function that tearsdown the directives created during the process. Note the decompile function does not remove the DOM. This method is exposed primarily for writing advanced custom directives.
func (v *Vue) Compile(element string) {
	v.Call("$compile", element)
}

// vm.$addChild( [options, constructor] )
// options Object optional
// constructor Function optional
// Adds a child instance to the current instance. The options object is the same in manually instantiating an instance. Optionally you can pass in a constructor created from Vue.extend().

// There are three implications of a parent-child relationship between instances:
// The parent and child can communicate via the event system.
// The child has access to all parent assets (e.g. custom directives).
// The child, if inheriting parent scope, has access to parent scope data properties.
