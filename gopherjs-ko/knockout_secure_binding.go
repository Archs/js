package ko

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

// Knockout Secure Binding (KSB) is a binding provider for Knockout
// that can be used with a Content Security Policy (CSP)
// that disables eval and new Function.
//
// Use this function to make gopherjs-ko works under chrome app/extensions.
//
// Must load knockout-secure-binding.min.js first:
// https://github.com/brianmhunt/knockout-secure-binding
func EnableSecureBinding() {
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
