package ko

import (
	"github.com/gopherjs/gopherjs/js"
)

var (
	defaultMapping *KoMapping
)

// KoMapping provide bindings for knockout.KoMapping.js
type KoMapping struct {
	*js.Object
	data    interface{}
	options *js.Object
	target  interface{}
}

func Mapping() *KoMapping {
	if defaultMapping == nil {
		defaultMapping = NewKoMapping()
	}
	return defaultMapping
}

func NewKoMapping() *KoMapping {
	m := &KoMapping{
		Object: Get("mapping"),
	}
	return m
}

func (m *KoMapping) args() []interface{} {
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
func (m *KoMapping) Target(vm ViewModel) *KoMapping {
	m.target = vm.ToJS()
	return m
}

func (m *KoMapping) FromJS(data *js.Object) ViewModel {
	vm := NewBaseViewModel()
	m.data = data
	vm.FromJS(m.Object.Call("fromJS", m.args()...))
	return vm
}

func (m *KoMapping) FromJSON(data string) ViewModel {
	vm := NewBaseViewModel()
	m.data = data
	vm.FromJS(m.Object.Call("fromJSON", m.args()...))
	return vm
}

func (m *KoMapping) ToJS(vm ViewModel) *js.Object {
	return m.Object.Call("toJS", vm.ToJS())
}

func (m *KoMapping) ToJSON(vm ViewModel) string {
	return m.Object.Call("toJSON", vm.ToJS()).String()
}

// Set KoMapping options
func (m *KoMapping) Option(key string, value interface{}) *KoMapping {
	if m.options == nil {
		m.options = js.Global.Get("Object").New()
	}
	m.options.Set(key, value)
	return m
}

// Ignoring certain properties using “ignore”
func (m *KoMapping) Ignore(properties ...string) *KoMapping {
	return m.Option("ignore", properties)
}

// Observing only certain properties using “observe”
func (m *KoMapping) Observe(properties ...string) *KoMapping {
	return m.Option("observe", properties)
}

func (m *KoMapping) Reset() *KoMapping {
	m.target = nil
	m.options = nil
	m.data = nil
	return m
}
