// +build ignore

package main

import (
	"github.com/Archs/js/vue"
	_ "github.com/Archs/js/vue/components/js-clock"
	"github.com/gopherjs/gopherjs/js"
)

func main() {
	vue.Component("my-cpnt", js.M{
		"template": "<h1>This is my testing component!</h1>" +
			"<content>This will only be displayed if no content is inserted</content>",
		"created": func() {
			println("'An instance of MyComponent has been created!'")
			js.This.Call("$dispatch", "msg", "hello")
		},
	})

	vm := vue.New(js.M{
		"el": "#demo",
		"data": js.M{
			"title": "todos",
			"todos": []js.M{
				js.M{
					"done":    true,
					"content": "Learn JavaScript",
				},
				js.M{
					"done":    false,
					"content": "Learn Vue.js",
				},
			},
		},
		"directives": js.M{
			"showdone": func(v js.Object) {
				println(v)
				println("this.expression:", js.This.Get("expression"))
			},
		},
		"filters": js.M{
			"testf": func(v js.Object) js.Object {
				println("testf:", v, js.This.Get("title"))
				return v
			},
		},
		"created": func() {
			js.This.Call("$on", "msg", func(msg interface{}) {
				println("parent got:", msg.(string))
			})
		},
	})
	// v := vue.New(vue.VueOption{
	// 	El: "#demo",
	// 	Data: js.M{
	// 		"title": "todos",
	// 		"todos": []js.M{
	// 			js.M{
	// 				"done":    true,
	// 				"content": "Learn JavaScript",
	// 			},
	// 			js.M{
	// 				"done":    false,
	// 				"content": "Learn Vue.js",
	// 			},
	// 		},
	// 	},
	// })
	println(vm.Object)
	println(vm.Options)
	vm.Watch("todos", func(newVal, oldVal js.Object) {
		println(newVal.Index(0).Get("done"))
	}, true)

	// vm.On("msg", func(msg interface{}) {
	// 	println("parent got:", msg.(string))
	// })
}
