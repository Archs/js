package main

import (
	// "github.com/gopherjs/gopherjs/js"
	// "github.com/gopherjs/jquery"
	// "log"
	// "strings"
	// "time"
	// "honnef.co/go/js/xhr"
	"net/http"
	"net/url"
)

func testNetHTTP() {
	// resp, err := http.Get("/get")
	// println(resp, err)
	// xhr
	// r := xhr.NewRequest("GET", "/get")
	// r.ResponseType = xhr.JSON
	// r.SetRequestHeader("Content-Type", "application/x-www-form-urlencoded")
	// if err := r.Send(js.M{
	// 	"a": 123,
	// }); err != nil {
	// 	println("xhr failed:", err.Error())
	// 	return
	// }
	// println("xhr ret:", r.Response)
	// jquery
	// jquery.Get("/get", js.M{
	// 	"a": strings.Join([]string{"a", "b"}, ","),
	// }, func(data *js.Object) {
	// 	log.Println(time.Now().String(), "jquery done:", data)
	// })
}

func testPost() {
	req, _ := http.NewRequest("POST", "/post", nil)
	req.SetBasicAuth("api", "key-3ax6xnjp29jd6fds4gc373sgvjxteol0")
	req.Header.Set("Content-Type", "application/x-www-form-urlencoded")
	//生成Form 因为Form是个Values对象来着
	form := url.Values{
		"from":    {"Excited User <me@samples.mailgun.org>"},
		"to":      {"mail@example.com"},
		"subject": {"Hello"},
		"text":    {"Testing some Mailgun awesomness!"}}
	//对form进行编码
	resp, err := http.PostForm("/post", form)
	if err != nil {
		println("PostForm failed:", err.Error())
	}
	println("PostForm ret:", resp)
}

func main() {
	testNetHTTP()
	testPost()
}
