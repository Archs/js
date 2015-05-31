package vCard

import (
	"github.com/Archs/js/gopherjs-ko"
	"github.com/gopherjs/gopherjs/js"
)

type vcard struct {
	*js.Object
	FirstName *ko.Observable `js:"FirstName"`
	LastName  *ko.Observable `js:"LastName"`
	AvatarUrl *ko.Observable `js:"AvatarUrl"`
	FullName  *ko.Computed   `js:"FullName"`
	About     *ko.Observable `js:"About"`
}

func newvc() *vcard {
	v := new(vcard)
	v.Object = js.Global.Get("Object").New()
	v.FirstName = ko.NewObservable("Knockout")
	v.LastName = ko.NewObservable("JS")
	v.AvatarUrl = ko.NewObservable()
	v.About = ko.NewObservable("bla bla bla ...")
	v.FullName = ko.NewComputed(func() interface{} {
		return v.FirstName.Get().String() + " " + v.LastName.Get().String()
	})
	return v
}

var (
	template = `
<div class="container">
    <img data-bind="attr:{src:AvatarUrl}, visibale:AvatarUrl">
    <div>
        <div>
            <span data-bind="text: FullName"></span>
        </div>
        <div data-bind="text: About">
            
        </div>
    </div>
</div>
`
	cssRules = `
    .container {
      display: inline-block;
      width: 19em;
      height: 5.3em;
      position: relative;
      border-style: groove;
      box-sizing: border-box;
    }
    .container > div {
        position: absolute;
        left: 5.1em;
        top:0;
    }
    ko-vcard img {
        width: 4.8em;
        height: 4.8em;
    }
`
)

func init() {
	ko.Components().RegisterEx("ko-vcard", func(params *js.Object) interface{} {
		vm := newvc()
		ko.Mapping().Target(vm).FromJS(params)
		return vm
	}, template, cssRules)
}
