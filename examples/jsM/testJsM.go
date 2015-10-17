package main

import (
	"github.com/gopherjs/gopherjs/js"
	"github.com/gopherjs/jquery"
	"io/ioutil"
	"log"
	"net/http"
)

func main() {
	jquery.GetJSON("/json", func(data js.M) {
		println("data:", data, data["Str"].(string), data["Float"].(float64), data["Int"].(float64))
	})
	resp, err := http.Get("/json")
	if err != nil {
		log.Fatalln(err.Error())
	}
	defer resp.Body.Close()
	dat, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		log.Fatalln(err.Error())
	}
	log.Println("json data:", string(dat))
}
