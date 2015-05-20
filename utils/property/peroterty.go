// Package property adds the ablity of getting/setting by keyPath to Gopherjs.
package property

import (
	"github.com/gopherjs/gopherjs/js"
	"strings"
)

func sep(path string) []string {
	ps := strings.Split(path, ".")
	if len(ps) < 1 {
		panic("no property specified")
	}
	return ps
}

func getProterty(obj *js.Object, ps []string) *js.Object {
	for _, p := range ps {
		obj = obj.Get(p)
		if obj == js.Undefined {
			break
		}
	}
	return obj
}

// return js.Undefined when no property found
func Get(obj *js.Object, path string) *js.Object {
	return getProterty(obj, sep(path))
}

func Set(obj *js.Object, path string, value interface{}) {
	ps := sep(path)
	n := len(ps)
	p := ps[n-1]
	ps = ps[:n-1]
	obj = getProterty(obj, ps)
	if obj == js.Undefined {
		return
	}
	obj.Set(p, value)
}
