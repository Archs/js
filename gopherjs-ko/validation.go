package ko

import "github.com/gopherjs/gopherjs/js"

var (
	validation = ko.Get("validation")
)

func NewValidatedObservable(data interface{}) *Observable {
	return &Observable{ko.Call("validatedObservable", data)}
}

// Only available when KnockoutJS Validation plugin is loaded
func (v *Observable) IsValid() bool {
	return v.o.Call("isValid").Bool()
}

type ValidationFuncs struct {
	*js.Object
}

func ValidationInit(config js.M) {
	validation.Call("init", config)
}
