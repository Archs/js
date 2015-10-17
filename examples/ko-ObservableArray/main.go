package main

import (
	"github.com/Archs/js/gopherjs-ko"
	"math"
)

type ViewModel struct {
	*ko.BaseViewModel
	Values   *ko.Observable `js:"values"`
	MaxValue *ko.Observable `js:"maxValue"`
}

func New() *ViewModel {
	self := new(ViewModel)
	self.BaseViewModel = ko.NewBaseViewModel()
	self.Values = ko.NewObservableArray()
	self.MaxValue = ko.NewComputed(func() interface{} {
		max := float64(0)
		for i := 0; i < self.Values.Length(); i++ {
			max = math.Max(max, self.Values.Index(i).Float())
		}
		return max
	})
	return self
}

func main() {
	vm := New()
	for i := 0; i < 100; i++ {
		vm.Values.Push(i)
	}
	println(vm.MaxValue.Get().Float())
	ko.ApplyBindings(vm)
}
