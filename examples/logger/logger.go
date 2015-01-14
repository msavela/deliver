package main

import (
	"github.com/msavela/deliver"
	"github.com/msavela/deliver/middleware"
	"net/http"
	"log"
)

func main() {
	d := deliver.New()

	d.UseHandlerNext(middleware.NewLogger())

	d.GET("/", func(res deliver.Response, req *deliver.Request) {
		res.Send(http.StatusText(http.StatusOK))
	})

	log.Fatal(http.ListenAndServe(":8080", d))
}