// +build ignore

package main

import (
	"github.com/Archs/js/vue"
	_ "github.com/Archs/js/vue/components/js-clock"
	"github.com/Archs/structs"
	"github.com/gopherjs/gopherjs/js"
)

type Todo struct {
	Done    bool   `js:"done" json:"done"`
	Content string `js:"content" json:"content"`
}

type Data struct {
	Title string  `js:"title" json:"title"`
	Todos []*Todo `js:"todos" json:"todos"`
}

func newData() *Data {
	return &Data{
		// Object: js.Global.Get("Object").New(),
		Title: "todos",
		Todos: []*Todo{
			&Todo{
				// Object:  js.Global.Get("Object").New(),
				Done:    false,
				Content: "asdfadadfa",
			},
			&Todo{
				// Object:  js.Global.Get("Object").New(),
				Done:    false,
				Content: "asdfadadfa",
			},
			&Todo{
				// Object:  js.Global.Get("Object").New(),
				Done:    false,
				Content: "asdfadadfa",
			},
		},
	}
	// dat := new(Data)
	// dat.Object = js.Global.Get("Object").New()
	// dat.title = "Todos"
	// t1 := new(Todo)
	// t1.Object = js.Global.Get("Object").New()
	// t1.done = false
	// t1.content = "asdfadfa"
	// t2 := new(Todo)
	// t2.Object = js.Global.Get("Object").New()
	// t2.done = false
	// t2.content = "asdfadfa"
	// dat.todos = []*Todo{t1, t2}
	// return dat
}

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
	data := newData()
	datam := structs.New(data).Tag("js").Map()
	println("data:", data)
	println("datam:", datam)
	vm := vue.New(js.M{
		"el": "#demo",
		// "data": js.M{
		// 	"title": "todos",
		// 	"todos": []js.M{
		// 		js.M{
		// 			"done":    true,
		// 			"content": "Learn JavaScript",
		// 		},
		// 		js.M{
		// 			"done":    false,
		// 			"content": "Learn Vue.js",
		// 		},
		// 	},
		// },
		"data": data,
		"methods": js.M{
			"change": func() {
				// data.Title = "Funcked"
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
	vm.Watch("todos", func(newVal, oldVal *js.Object) {
		println(newVal.Index(0).Get("done"))
	}, true)

	// vm.On("msg", func(msg interface{}) {
	// 	println("parent got:", msg.(string))
	// })
}
