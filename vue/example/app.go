// +build ignore

package main

import (
	"github.com/Archs/js/vue"
	_ "github.com/Archs/js/vue/components/js-clock"
	"github.com/gopherjs/gopherjs/js"
)

// type Todo struct {
// 	*js.Object
// 	Done    bool   `js:"done"`
// 	Content string `js:"content"`
// }

// type Data struct {
// 	*js.Object
// 	Title string `js:"title"`
// 	Todos []*Todo
// }

func main() {
	vue.Component("my-cpnt", js.M{
		"template": "<h1>This is my testing component!</h1>" +
			"<content>This will only be displayed if no content is inserted</content>",
		"created": js.MakeFunc(func(this *js.Object, arguments []*js.Object) interface{} {
			println("'An instance of MyComponent has been created!'")
			this.Call("$dispatch", "msg", "hello")
			return 0
		}),
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
			"showdone": js.MakeFunc(func(this *js.Object, arguments []*js.Object) interface{} {
				println("this.expression:", this.Get("expression"))
				return 0
			}),
		},
		"filters": js.M{
			"testf": js.MakeFunc(func(this *js.Object, arguments []*js.Object) interface{} {
				println("testf:", this.Get("title"))
				return 0
			}),
		},
		"created": js.MakeFunc(func(this *js.Object, arguments []*js.Object) interface{} {
			this.Call("$on", "msg", func(msg interface{}) {
				println("parent got:", msg.(string))
			})
			return 0
		}),
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
