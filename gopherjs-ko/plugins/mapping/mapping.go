// Package mapping provide bindings for knockout.mapping.js
package mapping

import (
	"github.com/Archs/js/gopherjs-ko"
	"github.com/gopherjs/gopherjs/js"
)

type Mapping struct {
	*js.Object
	data    interface{}
	options *js.Object
	target  interface{}
}

func New() *Mapping {
	m := &Mapping{
		Object: ko.Get("mapping"),
	}
	println("mapping:", m.Object)
	return m
}

func (m *Mapping) args() []interface{} {
	args := []interface{}{m.data}
	if m.options != nil {
		args = append(args, m.options)
	}
	if m.target != nil {
		if m.options == nil {
			args = append(args, js.M{})
		}
		args = append(args, m.target)
	}
	return args
}

// Specifying the target to update, can be a *ko.ViewModel or a *js.Object
// or a struct with *js.Object embeded which is a ko.ViewModel then
func (m *Mapping) Target(vm ko.ViewModel) *Mapping {
	m.target = vm.ToJS()
	return m
}

func (m *Mapping) FromJS(data *js.Object) ko.ViewModel {
	vm := ko.NewBaseViewModel()
	m.data = data
	vm.FromJS(m.Object.Call("fromJS", m.args()...))
	return vm
}

func (m *Mapping) FromJSON(data string) ko.ViewModel {
	vm := ko.NewBaseViewModel()
	m.data = data
	vm.FromJS(m.Object.Call("fromJSON", m.args()...))
	return vm
}

func (m *Mapping) ToJS(vm ko.ViewModel) *js.Object {
	return m.Object.Call("toJS", vm.ToJS())
}

func (m *Mapping) ToJSON(vm ko.ViewModel) string {
	return m.Object.Call("toJSON", vm.ToJS()).String()
}

// Set mapping options
func (m *Mapping) Option(key string, value interface{}) *Mapping {
	if m.options == nil {
		m.options = js.Global.Get("Object").New()
	}
	m.options.Set(key, value)
	return m
}

// Ignoring certain properties using “ignore”
func (m *Mapping) Ignore(properties ...string) *Mapping {
	return m.Option("ignore", properties)
}

// Observing only certain properties using “observe”
func (m *Mapping) Observe(properties ...string) *Mapping {
	return m.Option("observe", properties)
}
