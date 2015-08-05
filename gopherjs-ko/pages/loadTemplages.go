package pages

import (
	"github.com/Archs/js/dom"
	"honnef.co/go/js/xhr"
	"log"
	"path"
	"strings"
)

func rmExt(name string) string {
	idx := strings.Index(name, ".")
	return name[:idx]
}

func wrapIntoScriptTag(id string, text string) *dom.Element {
	basename := path.Base(id)
	t := dom.CreateElement("script")
	t.Id = rmExt(basename)
	t.SetAttribute("type", "text/html")
	t.InnerHTML = text
	return t
}

func LoadTemplages(endpoint string, names ...string) error {
	body := dom.Body()
	for _, name := range names {
		url := endpoint + "/" + name
		req := xhr.NewRequest("GET", url)
		req.ResponseType = xhr.Text
		if err := req.Send(nil); err != nil {
			log.Printf("Loading %s faild: %s", url, err.Error())
			continue
		}
		body.AppendChild(wrapIntoScriptTag(name, req.ResponseText))
	}
	return nil
}
