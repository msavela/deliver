package main

import (
	"../"
	"net/http"
	"log"
)

func main() {
	d := deliver.New()

	d.GET("/", func(res deliver.Response, req *deliver.Request) {
		res.Status(http.StatusOK).Send(http.StatusText(http.StatusOK))
	})

	log.Fatal(http.ListenAndServe(":8080", d))
}