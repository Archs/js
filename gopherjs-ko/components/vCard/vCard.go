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
    <img class="img" data-bind="attr:{src:AvatarUrl}" alt="No Pic Available">
    <div class="info">
        <div class="name" data-bind="text: FullName"></div>
        <div class="about" data-bind="text: About"></div>
    </div>
`
	cssRules = `
	ko-vcard {
		display: inline-block;
        border-style: groove;
        box-sizing: border-box;
        // min-width: 20em;
        vertical-align:bottom;
	}
    ko-vcard .img {
        width: 4.8em;
        height: 4.8em;
        float: left;
    }
    ko-vcard .name {
    	font-size: 2em;
		font-weight: bold;
		padding-right: 0.2em;
		text-align: end;
    }
    ko-vcard .info {
		display: inline-block;
        padding:0.2em;
        width:15em;
    }
    ko-vcard .about {
    	margin-top: 0.5em;
  		margin-left: 0.5em;
  		font-size: 0.5em;
  		// text-indent: 2em;
    }
`
)

func init() {
	ko.Components().RegisterEx("ko-vcard", func(params *js.Object, info *ko.ComponentInfo) interface{} {
		println("info.Element:", info.Element)
		vm := newvc()
		ko.Mapping().Target(vm).FromJS(params)
		return vm
	}, template, cssRules)
}
