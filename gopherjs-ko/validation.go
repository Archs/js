package ko

import "github.com/gopherjs/gopherjs/js"

type ValidatedObservable struct {
	*Observable
}

func NewValidatedObservable(data interface{}) *ValidatedObservable {
	return &ValidatedObservable{&Observable{Global().Call("validatedObservable", data)}}
}

func (ob *ValidatedObservable) IsValid() bool {
	return ob.Call("isValid").Bool()
}

type ValidationFuncs struct {
	*js.Object
}

func Validation() *ValidationFuncs {
	return &ValidationFuncs{Object: Global().Get("validation")}
}

func (v *ValidationFuncs) Init(config js.M) {
	v.Call("init", config)
}
