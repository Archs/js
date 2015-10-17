package main

import (
	"github.com/go-martini/martini"
	"github.com/martini-contrib/render"
	"log"
	"net/http"
)

func get(r render.Render, req *http.Request) {
	if err := req.ParseForm(); err != nil {
		r.JSON(http.StatusBadRequest, err.Error())
		log.Println("ParseForm failed:", err)
		return
	}
	log.Printf("reqeust params: %v\n", req.Form)
	r.JSON(http.StatusOK, req.Form)
}

func post(r render.Render, req *http.Request) {
	if err := req.ParseMultipartForm(2 * 1024 * 2014); err != nil {
		// r.JSON(http.StatusBadRequest, err.Error())
		log.Println("ParseMultipartForm failed:", err)
		// return
	}
	log.Printf("post params: %v\n", req.Form)
	if req.MultipartForm != nil && req.MultipartForm.File != nil {
		for name, f := range req.MultipartForm.File {
			log.Printf("Uploading file: %s with header %+v\n", name, f[0].Header)
		}
	}
	r.JSON(http.StatusOK, req.Form)
}

func main() {
	m := martini.Classic()
	m.Use(martini.Static("."))
	m.Use(render.Renderer())
	// url map
	m.Get("/get", get)
	m.Post("/post", post)
	// run
	m.RunOnAddr(":3000")
}
