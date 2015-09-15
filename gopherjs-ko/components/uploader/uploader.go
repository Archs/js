// Package uploader implements a KnockoutJS component: file uploader.
//
// Use it like this in html files after import this package:
//  <ko-uploader params="uploadUrl:'/uploadUrl', text:'Browser', buttonCls:'button round expand', multiple:true"></ko-uploader>
//  <span data-bind="component: {name:'ko-uploader',params:{uploadUrl:'/uploadUrl', text:'Browser', buttonCls:'button round expand'}}"></span>
package uploader

import (
	"github.com/Archs/js/dom"
	"github.com/Archs/js/gopherjs-ko"
	"github.com/gopherjs/gopherjs/js"
	"honnef.co/go/js/xhr"
)

const (
	template = `
		<button data-bind="click: onUploaderButtonClick, text: text, attr: { class: buttonCls }" ></button>
    	<input type="file" style="display:none" data-bind="event:{change:onFileInputChange}, attr:{multiple: multiple}">
    `
)

var (
	callback    func(url string, files []*dom.File)
	errCallback func(url string, files []*dom.File, statusCode int)
)

func SetCallback(cb func(url string, files []*dom.File)) {
	callback = cb
}

func SetErrorCallback(cb func(url string, files []*dom.File, statusCode int)) {
	errCallback = cb
}

type uploader struct {
	*js.Object
	text                  *ko.Observable                        `js:"text"`
	url                   *ko.Observable                        `js:"uploadUrl"`
	buttonCls             *ko.Observable                        `js:"buttonCls"`
	multiple              *ko.Observable                        `js:"multiple"`
	onFileInputChange     func(data *js.Object, evt *dom.Event) `js:"onFileInputChange"`
	onUploaderButtonClick func(data *js.Object, evt *dom.Event) `js:"onUploaderButtonClick"`
	target                *dom.Element
}

func newUploader() *uploader {
	u := new(uploader)
	u.Object = js.Global.Get("Object").New()
	u.url = ko.NewObservable("/asdafsdf")
	u.text = ko.NewObservable("Browser")
	u.buttonCls = ko.NewObservable("")
	u.multiple = ko.NewObservable()
	u.onFileInputChange = func(data *js.Object, evt *dom.Event) {
		println("onclick event:", evt.Type)
		u.target = evt.Target
		go u.upload()
	}
	u.onUploaderButtonClick = func(data *js.Object, evt *dom.Event) {
		el := evt.Target
		fileInput := el.NextElementSibling
		fileInput.Click()
	}
	return u
}

func (u *uploader) upload() {
	fd := dom.NewFormData()
	// println("fd.set", fd.Set)
	files := u.target.Files()
	for _, file := range files {
		fd.Append(file.Name, file)
	}
	req := xhr.NewRequest("POST", u.url.Get().String())
	println("xhr url:", u.url.Get().String())
	err := req.Send(fd)
	// when network error or response error
	if err != nil || req.Status != 200 {
		if errCallback != nil {
			errCallback(u.url.Get().String(), files, req.Status)
		}
		return
	}
	// upload ok
	if callback != nil {
		callback(u.url.Get().String(), files)
	}
}

func init() {
	ko.Components().RegisterEx("ko-uploader", func(params *js.Object, info *ko.ComponentInfo) interface{} {
		vm := newUploader()
		url := params.Get("uploadUrl")
		if url == js.Undefined {
			panic(info.Element.TagName + " Error:url for uploader must be provided")
		}
		vm.url.Set(url)
		text := params.Get("text")
		if text != js.Undefined {
			vm.text.Set(text)
		}
		cls := params.Get("buttonCls")
		if cls != js.Undefined {
			vm.buttonCls.Set(cls)
		}
		multiple := params.Get("multiple")
		if multiple != js.Undefined && multiple.Bool() {
			vm.multiple.Set(multiple)
		}
		return vm
	}, template, "")
}
