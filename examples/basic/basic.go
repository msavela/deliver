package main

import (
	"github.com/msavela/deliver"
	"net/http"
	"log"
)

func main() {
	d := deliver.New()

	d.GET("/", func(res deliver.Response, req *deliver.Request) {
		res.Send(http.StatusText(http.StatusOK))
	})

	log.Fatal(http.ListenAndServe(":8080", d))
}