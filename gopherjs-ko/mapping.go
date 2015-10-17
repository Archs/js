package ko

import (
	"github.com/gopherjs/gopherjs/js"
)

type Mapper struct {
	*js.Object
	data    interface{}
	options *js.Object
	target  interface{}
}

func Mapping() *Mapper {
	return &Mapper{
		Object: ko.Get("mapping"),
	}
}

func (m *Mapper) args() []interface{} {
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

// Specifying the target to update, can be a *ViewModel or a *js.Object
// or a struct with *js.Object embeded which is a ViewModel then
func (m *Mapper) Target(vm ViewModel) *Mapper {
	m.target = vm.ToJS()
	return m
}

func (m *Mapper) FromJS(data *js.Object) ViewModel {
	vm := NewBaseViewModel()
	m.data = data
	vm.FromJS(m.Object.Call("fromJS", m.args()...))
	return vm
}

func (m *Mapper) FromJSON(data string) ViewModel {
	vm := NewBaseViewModel()
	m.data = data
	vm.FromJS(m.Object.Call("fromJSON", m.args()...))
	return vm
}

func (m *Mapper) ToJS(vm ViewModel) *js.Object {
	return m.Object.Call("toJS", vm.ToJS())
}

func (m *Mapper) ToJSON(vm ViewModel) string {
	return m.Object.Call("toJSON", vm.ToJS()).String()
}

// Set mapping options
func (m *Mapper) Option(key string, value interface{}) *Mapper {
	if m.options == nil {
		m.options = js.Global.Get("Object").New()
	}
	m.options.Set(key, value)
	return m
}

// Ignoring certain properties using “ignore”
func (m *Mapper) Ignore(properties ...string) *Mapper {
	return m.Option("ignore", properties)
}

// Observing only certain properties using “observe”
func (m *Mapper) Observe(properties ...string) *Mapper {
	return m.Option("observe", properties)
}

// func isArray(i interface{}) bool {
// 	v := reflect.ValueOf(i)
// 	v = reflect.Indirect(v)
// 	if v.Type().Kind() == reflect.Array {
// 		return true
// 	}
// 	return false
// }
