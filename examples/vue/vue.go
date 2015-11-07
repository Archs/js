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
	*js.Object
	El            string     `js:"$el"`
	Data          *js.Object `js:"$data"`
	Options       *js.M      `js:"$options"`
	Parent        *js.Object `js:"$parent"`
	Root          *js.Object `js:"$root"`
	SubComponents *js.Object `js:"$"`  // v-ref, to ref sub component from parent
	Doms          *js.Object `js:"$$"` // v-el, expose dom element to Vue instance
}

func Directive(name string, opt js.M) {
	vue.Call("directive", name, opt)
}

func Extend(opt js.M) *js.Object {
	return vue.Call("extend", opt)
}

func RegisterComponent(name string, cpnt *js.Object) {
	vue.Call("component", name, cpnt)
}

func Component(name string, opt js.M) {
	vue.Call("component", name, opt)
}

func Filter(name string, fn func(*js.Object) *js.Object) {
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
// Example:
// 	vm := vue.New(js.M{
// 		"el": "#demo",
// 		"data": js.M{
// 			"title": "todos",
// 			"todos": []js.M{
// 				js.M{
// 					"done":    true,
// 					"content": "Learn JavaScript",
// 				},
// 				js.M{
// 					"done":    false,
// 					"content": "Learn Vue.js",
// 				},
// 			},
// 		},
// 		"directives": js.M{
// 			"showdone": func(v js.Object) {
// 				println(v)
// 				println("this.expression:", js.This.Get("expression"))
// 			},
// 		},
// 		"filters": js.M{
// 			"testf": func(v js.Object) js.Object {
// 				println("testf:", v, js.This.Get("title"))
// 				return v
// 			},
// 		},
// 		"created": func() {
// 			js.This.Call("$on", "msg", func(msg interface{}) {
// 				println("parent got:", msg.(string))
// 			})
// 		},
// 	)
//
// Options
//
// Data
//
// data
//
// Type: Object | Function
// Restricton: Only accepts Function when used in Vue.extend().
// The data object for the Vue instance. It can be accessed as vm.$data:
// 	var data = { a: 1 }
// 	var vm = new Vue({
// 	  data: data
// 	})
// 	vm.$data === data // -> true
// The Vue instance will proxy access to all its properties, therefore you can manipulate the properties on the Vue instance and the changes get synced back to the actual data object:
//
// 	vm.a   // -> 1
// 	vm.a = 2
// 	data.a // -> 2
// 	data.a = 3
// 	vm.a   // -> 3
// The object must be JSON-compliant (no circular references). You can use it just like an ordinary object, and it will look exactly the same when serialized with JSON.stringify. You can also share it between multiple Vue instances.
//
// A special case here is when using the data option in Vue.extend(). Since we don’t want nested objects to be shared by all instances created from that extended constructor, we must provide a function that returns a fresh copy of the default data:
// 	var MyComponent = Vue.extend({
// 	  data: function () {
// 	    return {
// 	      message: 'some default data.',
// 	      object: {
// 	        fresh: true
// 	      }
// 	    }
// 	  }
// 	})
// Under the hood, Vue.js attaches a hidden property __ob__ and recursively converts the object’s enumerable properties into getters and setters to enable dependency collection. Properties with keys that starts with $ or _ are skipped.
//
// methods
//
// Type: Object
// Methods to be mixed into the Vue instance. You can access these methods directly on the VM instance, or use them in directive expressions. All methods will have their this context automatically bound to the Vue instance.
//
// Example:
// 	var vm = new Vue({
// 	  data: { a: 1 },
// 	  methods: {
// 	    plus: function () {
// 	      this.a++
// 	    }
// 	  }
// 	})
// 	vm.plus()
// 	vm.a // 2
//
// computed
//
// Type: Object
// Computed properties to be mixed into the Vue instance. All getters and setters have their this context automatically bound to the Vue instance.
//
// Example:
// 	var vm = new Vue({
// 	  data: { a: 1 },
// 	  computed: {
// 	    // get only, just need a function
// 	    aDouble: function () {
// 	      return this.a * 2
// 	    },
// 	    // both get and set
// 	    aPlus: {
// 	      get: function () {
// 	        return this.a + 1
// 	      },
// 	      set: function (v) {
// 	        this.a = v - 1
// 	      }
// 	    }
// 	  }
// 	})
// 	vm.aPlus   // -> 2
// 	vm.aPlus = 3
// 	vm.a       // -> 2
// 	vm.aDouble // -> 4
//
// paramAttributes
//
// Type: Array
// An array of attribute names to be set on the Vue instance as initial data. Useful when passing data to a component.
//
// Example:
//
// 	Vue.component('param-demo', {
// 	  paramAttributes: ['size', 'message'],
// 	  compiled: function () {
// 	    console.log(this.size)    // -> 100
// 	    console.log(this.message) // -> 'hello!'
// 	  }
// 	})
//
// 	<param-demo size="100" message="hello!"></param-demo>
// 	Param attributes can also contain interpolation tags. The interpolation will be evaluated against the parent, and under the hood they will be compiled as v-with, which means when the value of the interpolated expression changes, the component’s corresponding property will also be updated:
//
// 	<param-demo message="{{parentMessage}}"></param-demo>
//
// Notes on hyphened attributes
//
// HTML attribute names ignore upper and lower case differences, so we usually use hyphened attributes instead of camel case. There are some special cases when using paramAttributes with attributes that contains hyphens:
//
// If the attribute is a data attribute, the data- prefix will be auto stripped;
//
// If the attribute still contains dashes, it will be camelized. This is because it’s inconvenient to access top level properties containing dashes in templates: the expression my-param will be parsed as a minus expression unless you use the awkward this['my-param'] syntax.
//
// This means a param attribute data-hello will be set on the vm as vm.hello; And my-param will be set as vm.myParam.
//
// DOM
//
// el
//
// Type: String | HTMLElement | Function
// Restriction: only accepts type Function when used in Vue.extend().
// Provide the Vue instance with an existing DOM element. It can be a CSS selector string, an actual HTMLElement, or a function that returns an HTMLElement. The resolved element will be accessible as vm.$el.
//
// When used in Vue.extend, a function must be provided so each instance gets a separately created element.
//
// If the option is available at instantiation, the instance will immediately enter compilation; otherwise, the user will have to explicitly call vm.$mount() to manually start the compilation.
//
// template
//
// Type: String
// A string template to be inserted into vm.$el. Any existing markup inside vm.$el will be overwritten, unless content insertion points are present in the template. If the replace option is true, the template will replace vm.$el entirely.
//
// If it starts with # it will be used as a querySelector and use the selected element’s innerHTML and the template string. This allows the use of the common <script type="x-template"> trick to include templates.
//
// Vue.js uses DOM-based templating. The compiler walks through DOM elements and looks for directives and creates data bindings. This means all Vue.js templates are parsable HTML that can be converted into actual DOM elements by the browser. Vue.js converts string templates into DOM fragments so they can be cloned when creating more Vue instances. If you want your templates to be valid HTML, you can configure the directive prefix to start with data-.
//
// replace
//
// Type: Boolean
// Default: false
// Restriction: only respected if the template option is also present.
// Whether to replace the original vm.$el with the template’s content instead of appending to it.
//
// Lifecycle
//
// All lifecycle hooks have their this context bound to the Vue instance they belong to. The Vue instance will also fire corresponding events for each hook in the form of "hook:<hookName>". e.g. for created, a "hook:created" event will be fired.
//
// created
//
// Type: Function
// Called synchronously after the instance is created. At this stage, the instance has finished processing the options which means the following have been set up: data observation, computed properties, methods, watch/event callbacks. However, DOM compilation has not been started, and the $el property will not be available yet.
//
// beforeCompile
//
// Type: Function
// Called right before the compilation starts.
//
// compiled
//
// Type: Function
// Called after the compilation is finished. At this stage all directives have been linked so data changes will trigger DOM updates. However, $el is not guaranteed to have been inserted into the document yet.
//
// ready
//
// Type: Function
// Called after compilation and the $el is inserted into the document for the first time. Note this insertion must be executed via Vue (with methods like vm.$appendTo() or as a result of a directive update) to trigger the ready hook.
//
// attached
//
// Type: Function
// Called when vm.$el is attached to DOM by a directive or a VM instance method such as $appendTo(). Direct manipulation of vm.$el will not trigger this hook.
//
// detached
//
// Type: Function
// Called when vm.$el is removed from the DOM by a directive or a VM instance method. Direct manipulation of vm.$el will not trigger this hook.
//
// beforeDestroy
//
// Type: Function
// Called right before a Vue instance is destroyed. At this stage the instance is still fully functional.
//
// destroyed
//
// Type: Function
// Called after a Vue instance has been destroyed. When this hook is called, all bindings and directives of the Vue instance have been unbound and all child Vue instances have also been destroyed.
//
// Note if there is a leaving transition, the destroyed hook is called after the transition has finished.
//
// Assets
//
// These are private assets that will be available only to this Vue instance and its children during compilation.
//
// directives
//
// Type: Object
// A hash of directives to be made available to the Vue instance. For details on how to write a custom directive, see Writing Custom Directives.
//
// filters
//
// Type: Object
// A hash of filters to be made available to the Vue instance. For details on how to write a custom filter, see Writing Custom Filters.
//
// components
//
// Type: Object
// A hash of components to be made available to the Vue instance. For details on how to extend and compose Vue instances, see Component System.
//
// partials
//
// Type: Object
// A hash of partials to be made available to the Vue instance. Also see v-partial.
//
// transitions
//
// Type: Object
// A hash of transitions to be made available to the Vue instance. For details see the guide on Transitions.
//
// Others
//
// inherit
//
// Type: Boolean
// Default: false
// Whether to inherit parent scope data. Set it to true if you want to create a component that inherits parent scope. When inherit is set to true, you can:
//
// Bind to parent scope properties in the component template;
// Directly access parent properties on the component instance itself, via prototypal inheritance.
// One important thing to know when using inherit: true is that the child can also set parent properties, because all Vue instance data properties are getter/setters.
//
// Example:
// 	var parent = new Vue({
// 	  data: { a: 1 }
// 	})
// 	var child = parent.$addChild({
// 	  inherit: true,
// 	  data: { b: 2 }
// 	})
// 	child.a  // -> 1
// 	child.b  // -> 2
// 	// the following line modifies parent.a
// 	// instead of creating a new property on child:
// 	child.a = 2
// 	parent.a // -> 2
//
// events
//
// An object where keys are events to listen for and values are the corresponding callbacks. Note these are Vue events rather than DOM events. The value can also be a string of a method name. The Vue instance will call $on() for each entry in the object at instantiation.
//
// Example:
// 	var vm = new Vue({
// 	  events: {
// 	    'hook:created': function () {
// 	      console.log('created!')
// 	    },
// 	    greeting: function (msg) {
// 	      console.log(msg)
// 	    },
// 	    // can also use a string for methods
// 	    bye: 'sayGoodbye'
// 	  },
// 	  methods: {
// 	    sayGoodbye: function () {
// 	      console.log('goodbye!')
// 	    }
// 	  }
// 	}) // -> created!
// 	vm.$emit('greeting', 'hi!') // -> hi!
// 	vm.$emit('bye')             // -> goodbye!
//
// watch
//
// Type: Object
// An object where keys are expressions to watch and values are the corresponding callbacks. The value can also be a string of a method name. The Vue instance will call $watch() for each entry in the object at instantiation.
//
// Example:
// 	var vm = new Vue({
// 	  data: {
// 	    a: 1
// 	  },
// 	  watch: {
// 	    'a': function (val, oldVal) {
// 	      console.log('new: %s, old: %s', val, oldVal)
// 	    }
// 	  }
// 	})
// 	vm.a = 2 // -> new: 2, old: 1
//
// mixins
//
// Type: Array
// The mixins option accepts an array of mixin objects. These mixin objects can contain instance options just like normal instance objects, and they will be merged against the eventual options using the same option merging logic in Vue.extend(). e.g. If your mixin contains a created hook and the component itself also has one, both functions will be called.
//
// Example:
//
// 	var mixin = {
// 	  created: function () { console.log(2) }
// 	}
// 	var vm = new Vue({
// 	  created: function () { console.log(1) },
// 	  mixins: [mixin]
// 	})
// 	// -> 1
// 	// -> 2
//
// name
//
// Type: String
// Restrctions: only respected when used in Vue.extend().
// When inspecting an extended Vue component in the console, the default constructor name is VueComponent, which isn’t very informative. By passing in an optional name option to Vue.extend(), you will get a better inspection output so that you know which component you are looking at. The string will be camelized and used as the component’s constructor name.
//
// Example:
//
// 	var Ctor = Vue.extend({
// 	  name: 'cool-stuff'
// 	})
// 	var vm = new Ctor()
// 	console.log(vm) // -> CoolStuff {$el: null, ...}
func New(opts js.M) *Vue {
	vm := vue.New(opts)
	return &Vue{
		Object: vm,
	}
}

// func (v *Vue) Directive(name string, opt js.M) *Vue {
// 	v.Call("directive", name, opt)
// 	return v
// }

// func (v *Vue) Component(name string, opt js.M) *Vue {
// 	v.Call("component", name, opt)
// 	return v
// }

// func (v *Vue) Filter(name string, fn interface{}) *Vue {
// 	v.Call("filter", name, fn)
// 	return v
// }

type Unwatcher func()

// vm.$watch( expression, callback, [deep, immediate] )
//
// expression String
// callback( newValue, oldValue ) Function
// deep Boolean optional
// immdediate Boolean optional
//
// Watch an expression on the Vue instance for changes.
// The expression can be a single keypath or actual expressions:
func (v *Vue) Watch(expression string, callback func(newVal, oldVal *js.Object), deepWatch bool) Unwatcher {
	obj := v.Call("$watch", expression, callback, deepWatch)
	return func() {
		obj.Invoke()
	}
}

// vm.$eval( expression )
//
// expression String
// Evaluate an expression that can also contain filters.
//
// // assuming vm.msg = 'hello'
// vm.$eval('msg | uppercase') // -> 'HELLO'
func (v *Vue) Eval(expression string) *js.Object {
	return v.Call("$eval", expression)
}

// vm.$get( expression )
//
// expression String
//
// Retrieve a value from the Vue instance given an expression.
// Expressions that throw errors will be suppressed and return undefined.
func (v *Vue) Get(expression string) {
	v.Call("$get", expression)
}

// vm.$set( keypath, value )

// keypath String
// value *
//
// Set a data value on the Vue instance given a valid keypath.
// If the path doesn’t exist it will be created.
func (v *Vue) Set(keypath string, val interface{}) {
	v.Call("$set", keypath, val)
}

// vm.$add( keypath, value )
//
// keypath String
// value *
// Add a root level property to the Vue instance (and also its $data). Due to the limitations of ES5, Vue cannot detect properties directly added to or deleted from an Object, so use this method and vm.$delete when you need to do so. Additionally, all observed objects are augmented with these two methods too.
func (v *Vue) Add(keypath string, val interface{}) {
	v.Call("$add", keypath, val)
}

// vm.$delete( keypath )
//
// keypath String
// Delete a root level property on the Vue instance (and also its $data).
func (v *Vue) Delete(keypath string) {
	v.Call("$delete", keypath)
}

// vm.$interpolate( templateString )
// templateString String
// Evaluate a piece of template string containing mustache interpolations.
// Note that this method simply performs string interpolation;
// attribute directives are not compiled.
//
// // assuming vm.msg = 'hello'
// vm.$interpolate('{{msg}} world!') // -> 'hello world!'
func (v *Vue) Interpolate(templateString string) {
	v.Call("$interpolate", templateString)
}

// Events
// Each vm is also an event emitter.
// When you have multiple nested ViewModels,
// you can use the event system to communicate between them.
//
// vm.$dispatch( event, [args…] )
// event String
// args… optional
//
// Dispatch an event from the current vm that propagates all the way up to its $root.
// If a callback returns false, it will stop the propagation at its owner instance.
func (v *Vue) Dispatch(event string, args interface{}) {
	v.Call("$dispatch", event, args)
}

// vm.$broadcast( event, [args…] )
// event String
// args… optional
//
// Emit an event to all children vms of the current vm,
// which gets further broadcasted to their children all the way down.
// If a callback returns false, its owner instance will not broadcast the event any further.
func (v *Vue) Broadcast(event string, args interface{}) {
	v.Call("$broadcast", event, args)
}

// vm.$emit( event, [args…] )
// event String
// args… optional
//
// Trigger an event on this vm only.
func (v *Vue) Emit(event string, args interface{}) {
	v.Call("$emit", event, args)
}

type EventCallback func(args interface{})

// vm.$on( event, callback )
// event String
// callback Function
//
// Listen for an event on the current vm
func (v *Vue) On(event string, cb EventCallback) {
	v.Call("$on", event, cb)
}

// vm.$once( event, callback )
// event String
// callback Function
//
// Attach a one-time only listener for an event.
func (v *Vue) Once(event string, cb EventCallback) {
	v.Call("$once", event, cb)
}

// vm.$off( [event, callback] )
// event String optional
// callback Function optional
//
// If no arguments are given, stop listening for all events;
// if only the event is given, remove all callbacks for that event;
// if both event and callback are given, remove that specific callback only.
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
//
// Adds a child instance to the current instance.
// The options object is the same in manually instantiating an instance.
// Optionally you can pass in a constructor created from Vue.extend().
//
// There are three implications of a parent-child relationship between instances:
// The parent and child can communicate via the event system.
// The child has access to all parent assets (e.g. custom directives).
// The child, if inheriting parent scope, has access to parent scope data properties.
func (v *Vue) AddChild(options js.M) {
	v.Call("$addChild", options)
}

// vm.$log( [keypath] )
//
// keypath String optional
// Log the current instance data as a plain object, which is more console-inspectable than a bunch of getter/setters. Also accepts an optional key.
//
// vm.$log() // logs entire ViewModel data
// vm.$log('item') // logs vm.item
func (v *Vue) Log() {
	v.Call("$log")
}
