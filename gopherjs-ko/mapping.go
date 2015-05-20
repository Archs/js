package ko

import (
	"github.com/Archs/js/utils/property"
	"github.com/gopherjs/gopherjs/js"
)

// MappedViewModel using in the Mapped result
type MappedViewModel struct {
	*js.Object
}

type Mapper struct {
	*js.Object
	data    interface{}
	options *js.Object
	target  *js.Object
}

func Mapping() *Mapper {
	return &Mapper{
		Object: Global().Get("mapping"),
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

func (m *Mapper) FromJS(data interface{}) (vm *MappedViewModel) {
	vm = new(MappedViewModel)
	m.data = data
	vm.Object = m.Object.Call("fromJS", m.args()...)
	return
}

// Specifying the update target
func (m *Mapper) Target(obj *js.Object) *Mapper {
	m.target = obj
	return m
}

func (m *Mapper) ToJS(vm *MappedViewModel) *js.Object {
	return m.Object.Call("toJS", vm.Object)
}

func (m *Mapper) FromJSON(data string) (vm *MappedViewModel) {
	vm = new(MappedViewModel)
	m.data = data
	vm.Object = m.Object.Call("fromJSON", m.args()...)
	return
}

func (m *Mapper) ToJSON(vm interface{}) string {
	return m.Object.Call("toJSON", vm).String()
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

func (v *MappedViewModel) Set(keyPath string, value interface{}) {
	obj := property.Get(v.Object, keyPath)
	if obj == js.Undefined {
		// if isArray(value) {
		// 	v.Set(key, NewObservableArray(value))
		// } else {
		// 	v.Set(key, NewObservable(value))
		// }
		panic("MappedViewModel has no key: " + keyPath)
	} else {
		obj.Invoke(value)
	}
}

func (v *MappedViewModel) Get(keyPath string) *js.Object {
	obj := property.Get(v.Object, keyPath)
	if obj == js.Undefined {
		return obj
	}
	return obj.Invoke()
}

func (v *MappedViewModel) Update(data interface{}) *MappedViewModel {
	Mapping().Call("fromJS", data, v.Object)
	return v
}
