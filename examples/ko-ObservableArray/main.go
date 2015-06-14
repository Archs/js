package main

import (
	"github.com/Archs/js/gopherjs-ko"
	"github.com/gopherjs/gopherjs/js"
	"math"
)

type ViewModel struct {
	*js.Object
	Values   *ko.ObservableArray `js:"values"`
	MaxValue *ko.Computed        `js:"maxValue"`
}

// ...

func New() *ViewModel {
	self := new(ViewModel)
	self.Object = js.Global.Get("Object").New()
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
	// ko.ApplyBindings(vm)
	println(vm.MaxValue.Get().Float())
}
