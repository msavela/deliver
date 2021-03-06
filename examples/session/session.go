package main

import (
	"github.com/msavela/deliver"
	"net/http"
	"log"
)

type Session struct {
	Username string
}

type SessionMiddleware struct{}
func (r *SessionMiddleware) ServeHTTP(res deliver.Response, req *deliver.Request, next func()) {
	// Set session with example username
	session := Session{"username"}
	req.Context.Set("session", session)

	next()
}

func main() {
	d := deliver.New()

	// Set global middleware to handle session
	d.Use(&SessionMiddleware{})

	d.GET("/", func(res deliver.Response, req *deliver.Request) {
		if session, ok := req.Context.Get("session").(Session); ok {
			// Session found, respond with session username
			res.Send(session.Username)
		} else {
			// Session not found
			res.Send("Session not found")
		}
	})

	log.Fatal(http.ListenAndServe(":8080", d))
}