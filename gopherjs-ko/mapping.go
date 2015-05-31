package ko

import (
	"github.com/Archs/js/utils/property"
	"github.com/gopherjs/gopherjs/js"
)

// ViewModel can be used to wrap ko vm object
type ViewModel struct {
	*js.Object
}

type Mapper struct {
	*js.Object
	data    interface{}
	options *js.Object
	target  interface{}
}

func Mapping() *Mapper {
	return &Mapper{
		Object: ko().Get("mapping"),
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

func (m *Mapper) FromJS(data interface{}) (vm *ViewModel) {
	vm = new(ViewModel)
	m.data = data
	vm.Object = m.Object.Call("fromJS", m.args()...)
	return
}

// Specifying the target to update, can be a *ViewModel or a *js.Object
// or a struct with *js.Object embeded which is a ViewModel then
func (m *Mapper) Target(obj interface{}) *Mapper {
	m.target = obj
	return m
}

func (m *Mapper) ToJS(vm *ViewModel) *js.Object {
	return m.Object.Call("toJS", vm.Object)
}

func (m *Mapper) FromJSON(data string) (vm *ViewModel) {
	vm = new(ViewModel)
	m.data = data
	vm.Object = m.Object.Call("fromJSON", m.args()...)
	return
}

func (m *Mapper) ToJSON(vm *ViewModel) string {
	return m.Object.Call("toJSON", vm.Object).String()
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

func (v *ViewModel) Set(keyPath string, value interface{}) {
	obj := property.Get(v.Object, keyPath)
	if obj == js.Undefined {
		// if isArray(value) {
		// 	v.Set(key, NewObservableArray(value))
		// } else {
		// 	v.Set(key, NewObservable(value))
		// }
		panic("ViewModel has no key: " + keyPath)
	} else {
		obj.Invoke(value)
	}
}

func (v *ViewModel) Get(keyPath string) *js.Object {
	obj := property.Get(v.Object, keyPath)
	if obj == js.Undefined {
		return obj
	}
	return obj.Invoke()
}

func (v *ViewModel) Update(data interface{}) *ViewModel {
	Mapping().Call("fromJS", data, v.Object)
	return v
}
