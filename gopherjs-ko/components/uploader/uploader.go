// Package uploader implements a KnockoutJS component: file uploader.
//
// Use it like this in html files after import this package:
//  <ko-uploader params="uploadUrl:'/upload'"></ko-uploader>
//  <span data-bind="component: {name:'ko-uploader',params:{width:200,height:150}}"></span>
package uploader

import (
	"github.com/Archs/js/dom"
	"github.com/Archs/js/gopherjs-ko"
	"github.com/gopherjs/gopherjs/js"
	"honnef.co/go/js/xhr"
)

const (
	template = `
    <input type="file" data-bind="event:{change:onFileInputChange}">
    `
)

type uploader struct {
	*js.Object
	onFileInputChange func(data *js.Object, evt *dom.Event) `js:"onFileInputChange"`
	target            *dom.Element
	url               string
}

func newUploader(url string) *uploader {
	u := &uploader{
		Object: js.Global.Get("Object").New(),
		url:    url,
	}
	u.onFileInputChange = func(data *js.Object, evt *dom.Event) {
		println("onclick event:", evt.Type)
		u.target = evt.Target
		go u.upload()
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
	req := xhr.NewRequest("POST", u.url)
	req.Send(fd)
}

func init() {
	ko.Components().RegisterEx("ko-uploader", func(params *js.Object, info *ko.ComponentInfo) interface{} {
		url := params.Get("url")
		if url == nil {
			panic("url for uploader must be provided")
		}
		vm := newUploader(url.String())
		return vm
	}, template, "")
}
