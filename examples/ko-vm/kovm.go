package main

import (
	"github.com/Archs/js/dom"
	"github.com/Archs/js/gopherjs-ko"
	"github.com/gopherjs/gopherjs/js"
)

type TestVM struct {
	*ko.BaseViewModel
	Str   *ko.Observable `js:"str"`
	Int   *ko.Observable `js:"int"`
	Array *ko.Observable `js:"arr"`
}

func newVM() *TestVM {
	vm := &TestVM{
	// js.Global.Get("Object").New(),
	}
	vm.BaseViewModel = ko.NewBaseViewModel()
	vm.Str = ko.NewObservable("nice to see you")
	vm.Int = ko.NewObservable(1000)
	vm.Array = ko.NewObservableArray([]int{1, 2, 3, 4, 5})
	return vm
}

func main() {
	vm := newVM()
	println("vm:", vm, vm.Str.Get().String())
	ko.RegisterBinding("mytext", nil, func(el *dom.Element, valueAccessor func() *js.Object) {
		println("my binding called:", valueAccessor())
		el.InnerHTML = valueAccessor().String()
	})
	ko.RegisterCustomBinding("test", nil, func(el *dom.Element, valueAccessor func() *js.Object, allBindings *ko.AllBindings, viewmodel *js.Object, bindingContext *ko.BindingContext) {
		el.InnerHTML = valueAccessor().String()
		println("test custom binding :", el, valueAccessor(), allBindings, viewmodel, bindingContext)
		println("allbindigns.has('param'):", allBindings.Has("param"), allBindings.Get("param"))
		println("binding context:", bindingContext.Data(), bindingContext.Root())
	})
	ko.ApplyBindings(newVM())
}
