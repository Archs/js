package ko

import "github.com/gopherjs/gopherjs/js"

type ValidatedObservable struct {
	*Observable
}

func NewValidatedObservable(data interface{}) *ValidatedObservable {
	return &ValidatedObservable{&Observable{ko().Call("validatedObservable", data)}}
}

func (v *ValidatedObservable) IsValid() bool {
	return v.o.Call("isValid").Bool()
}

type ValidationFuncs struct {
	*js.Object
}

func Validation() *ValidationFuncs {
	return &ValidationFuncs{Object: ko().Get("validation")}
}

func (v *ValidationFuncs) Init(config js.M) {
	v.Call("init", config)
}
