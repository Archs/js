package main

import (
	"github.com/Archs/js/dom"
	"github.com/Archs/js/gopherjs-ko"
)

type VM struct {
	*ko.BaseViewModel
	List *ko.Observable `js:"list"`
	Doit func()         `js:"doit"`
}

func newVm() *VM {
	vm := new(VM)
	vm.BaseViewModel = ko.NewBaseViewModel()
	vm.List = ko.NewObservableArray([]int{1, 2, 3, 4, 5})
	vm.Doit = func() {
		col := dom.Document().GetElementsByTagName("li")
		println("col:", col.Length, col.Item(0))
		el := dom.Document().QuerySelector("ul")
		els := el.QuerySelectorAll("li")
		println("els:", els.Length, els.Index(0))
	}
	return vm
}

func main() {
	ko.ApplyBindings(newVm())
}
