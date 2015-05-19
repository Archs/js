package ko

import (
	"github.com/gopherjs/gopherjs/js"
)

// ViewModel using in the Mapped result
type ViewModel struct {
	*js.Object
}

type Mapper struct {
	*js.Object
	options *js.Object
	target  *js.Object
}

func Mapping() *Mapper {
	return &Mapper{
		Object: Global().Get("mapping"),
	}
}

func (m *Mapper) FromJS(data interface{}) (model *ViewModel) {
	model = new(ViewModel)
	if m.options != nil {
		model.Object = m.Object.Call("fromJS", data, m.options)
	} else {
		model.Object = m.Object.Call("fromJS", data)
	}
	return
}

// Specifying the update target
// TODO implements Target
// func (m *Mapper) Target(obj *js.Object) *Mapper {
// 	m.target = obj
// 	return m
// }

func (v *ViewModel) Update(data interface{}) *ViewModel {
	Mapping().Call("fromJS", data, v.Object)
	return v
}

func (m *Mapper) ToJS(vm *ViewModel) *js.Object {
	return m.Object.Call("toJS", vm.Object)
}

func (m *Mapper) FromJSON(data string) (model *ViewModel) {
	model = new(ViewModel)
	if m.options != nil {
		model.Object = m.Object.Call("fromJSON", data, m.options)
	} else {
		model.Object = m.Object.Call("fromJSON", data)
	}
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

func (v *ViewModel) Set(key string, value interface{}) {
	if v.Get(key) == nil {
		// if isArray(value) {
		// 	v.Set(key, NewObservableArray(value))
		// } else {
		// 	v.Set(key, NewObservable(value))
		// }
		panic("vm has no key: " + key)
	} else {
		v.Call(key, value)
	}
}

func (v *ViewModel) Get(key string) *js.Object {
	return v.Call(key)
}
