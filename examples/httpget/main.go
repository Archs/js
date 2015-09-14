package main

import (
	"net/http"
)

func testNetHTTP() {
	resp, err := http.Get("http://example.com/")
	println(resp, err)
}

func main() {
	testNetHTTP()
}
