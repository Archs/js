package localStorage

import (
	"github.com/gopherjs/gopherjs/js"
)

var (
	localStorage = js.Global.Get("localStorage")
)

// Save val into localStorage under key
func SetItem(key string, val string) {
	localStorage.Call("setItem", key, val)
}

func GetItem(key string) string {
	return localStorage.Call("getItem", key).String()
}
