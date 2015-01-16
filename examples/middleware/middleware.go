package main

import (
	"github.com/msavela/deliver"
	"net/http"
	"log"
	"fmt"
)

type HandlerWithNextMiddleware struct{}
func (r *HandlerWithNextMiddleware) ServeHTTP(res deliver.Response, req *deliver.Request, next func()) {
	fmt.Println("HandlerWithNextMiddleware")
	next()
}

type HttpHandlerMiddleware struct{}
func (r *HttpHandlerMiddleware) ServeHTTP(w http.ResponseWriter, req *http.Request) {
	fmt.Println("HttpHandlerMiddleware")
}

type HttpHandlerWithNextMiddleware struct{}
func (r *HttpHandlerWithNextMiddleware) ServeHTTP(w http.ResponseWriter, req *http.Request, next func()) {
	fmt.Println("HttpHandlerWithNextMiddleware")
	next()
}

func main() {
	d := deliver.New()

	d.Use(&HandlerWithNextMiddleware{})
	d.UseHandler(&HttpHandlerMiddleware{})
	d.UseHandlerNext(&HttpHandlerWithNextMiddleware{})

	d.GET("/", func(res deliver.Response, req *deliver.Request) {
		res.Send(http.StatusText(http.StatusOK))
	})

	log.Fatal(http.ListenAndServe(":8080", d))
}