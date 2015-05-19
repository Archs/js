package ko

import "github.com/gopherjs/gopherjs/js"

type ValidatedObservable interface {
	IsValid() bool
}

func NewValidatedObservable(data interface{}) ValidatedObservable {
	return &Object{Global().Call("validatedObservable", data)}
}

func (ob *Object) IsValid() bool {
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
