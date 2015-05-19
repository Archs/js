// Knockout Secure Binding (KSB) is a binding provider for Knockout
// that can be used with a Content Security Policy
// that disables eval and new Function.
//
// Works by importing this package anonymously
package koSecureBindings

import (
	"github.com/gopherjs/gopherjs/js"
)

// var options = {
//    attribute: "data-bind",        // default "data-sbind"
//    globals: window,               // default {}
//    bindings: ko.bindingHandlers,  // default ko.bindingHandlers
//    noVirtualElements: false       // default true
// };
// ko.bindingProvider.instance = new ko.secureBindingsProvider(options);
func ksb() {
	ko := js.Global.Get("ko")
	secureBindingsProvider := ko.Get("secureBindingsProvider")
	ksbp := secureBindingsProvider.New(js.M{
		"attribute":         "data-bind",               // default "data-sbind"
		"globals":           js.Global.Get("window"),   // default {}
		"bindings":          ko.Get("bindingHandlers"), // default ko.bindingHandlers
		"noVirtualElements": false,                     // default true
	})
	ko.Get("bindingProvider").Set("instance", ksbp)
}

func init() {
	ksb()
}
