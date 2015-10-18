package vCard

import (
	"github.com/Archs/js/gopherjs-ko"
	"github.com/Archs/js/gopherjs-ko/plugins/mapping"
	"github.com/gopherjs/gopherjs/js"
)

type vcard struct {
	*ko.BaseViewModel
	FirstName *ko.Observable `js:"FirstName"`
	LastName  *ko.Observable `js:"LastName"`
	AvatarUrl *ko.Observable `js:"AvatarUrl"`
	FullName  *ko.Observable `js:"FullName"`
	About     *ko.Observable `js:"About"`
}

func newvc() *vcard {
	v := new(vcard)
	v.BaseViewModel = ko.NewBaseViewModel()
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
	ko.RegisterComponent("ko-vcard", func(params *js.Object, info *ko.ComponentInfo) ko.ViewModel {
		println("info.Element:", info.Element)
		vm := newvc()
		mapping.New().Target(vm).FromJS(params)
		return vm
	}, template, cssRules)
}
