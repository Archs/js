package fn

import (
	"encoding/json"
	"github.com/Archs/js/dom"
	"github.com/gopherjs/gopherjs/js"
	"reflect"
)

// Wrap returns a wrapper func that handles the conversion from native JavaScript *js.Object parameters
// to the following types.
//
// It supports *js.Object (left unmodified), dom.Document, dom.Element, dom.Event, dom.HTMLElement, dom.Node.
//
// For other types, the input is assumed to be a JSON string which is then unmarshalled into that type.
func Wrap(fn interface{}) func(...*js.Object) {
	v := reflect.ValueOf(fn)
	return func(args ...*js.Object) {
		in := make([]reflect.Value, v.Type().NumIn())
		for i := range in {
			switch t := v.Type().In(i); t {
			// *js.Object is passed through.
			case typeOf((**js.Object)(nil)):
				in[i] = reflect.ValueOf(args[i])

			// dom types are wrapped.
			case typeOf((*dom.Element)(nil)):
				in[i] = reflect.ValueOf(dom.Wrap(args[i]))
			case typeOf((*dom.Event)(nil)):
				in[i] = reflect.ValueOf(dom.WrapEvent(args[i]))

			// Unmarshal incoming encoded JSON into the Go type.
			default:
				p := reflect.New(t)
				err := json.Unmarshal([]byte(args[i].String()), p.Interface())
				if err != nil {
					panic(err)
				}
				in[i] = reflect.Indirect(p)
			}
		}
		v.Call(in)
	}
}

// typeOf returns the reflect.Type of what the pointer points to.
func typeOf(pointer interface{}) reflect.Type {
	return reflect.TypeOf(pointer).Elem()
}
