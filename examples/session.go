package main

import (
	"../"
	"net/http"
	"log"
)

type Session struct {
	Username string
}

type SessionMiddleware struct{}
func (r *SessionMiddleware) ServeHTTP(res deliver.Response, req *deliver.Request, next func()) {
	// Set session with example username
	session := Session{"example@somewhere.com"}
	req.Context.Set("session", session)

	next()
}

func main() {
	n := deliver.New()

	// Set global middleware to handle session
	n.Use(&SessionMiddleware{})

	n.GET("/", func(res deliver.Response, req *deliver.Request) {
		if session, ok := req.Context.Get("session").(Session); ok {
			// Session found, respond with session username
			res.Status(200).Send(session.Username)
		} else {
			// Session not found
			res.Status(404).Send("Session not found")
		}
	})

	log.Fatal(http.ListenAndServe(":8080", n))
}