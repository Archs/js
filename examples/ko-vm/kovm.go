package main

import (
	"github.com/Archs/js/gopherjs-ko"
)

type TestVM struct {
	*ko.BaseViewModel
	Str *ko.Observable `js:"str"`
	Int *ko.Observable `js:"int"`
}

func newVM() *TestVM {
	vm := &TestVM{
	// js.Global.Get("Object").New(),
	}
	vm.BaseViewModel = ko.NewBaseViewModel()
	vm.Str = ko.NewObservable("nice to see you")
	vm.Int = ko.NewObservable(1000)
	return vm
}

func main() {
	vm := newVM()
	println("vm:", vm, vm.Str.Get().String())
	ko.ApplyBindings(newVM())
}
