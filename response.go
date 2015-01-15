package deliver

import (
	"net/http"
)

type (
	Response interface {
		http.ResponseWriter
		Written()		bool
		Status() 		int
		SetStatus(int)	*response
		Send(string)	*response
	}

	response struct {
		http.ResponseWriter
		status	int
	}
)

// Initialize a new response.
func NewResponse(res http.ResponseWriter) Response {
	return &response{res, 0}
}

// WriteHeader stores the status code and writes header.
func (r *response) WriteHeader(status int) {
	r.status = status
	r.ResponseWriter.WriteHeader(status)
}

// Writes response and sets the status code in case not already set.
func (r *response) Write(b []byte) (int, error) {
	if !r.Written() {
		// Set StatusOK in case status not set
		r.WriteHeader(http.StatusOK)
	}
	size, err := r.ResponseWriter.Write(b)
	return size, err
}

// Is status already written.
func (r *response) Written() bool {
	return r.status != 0
}

// Get response status code.
func (r *response) Status() int {
	return r.status
}

// Set response status code.
func (r *response) SetStatus(status int) *response {
	r.status = status
	r.WriteHeader(status)
	return r
}

// Send response body.
func (r *response) Send(response string) *response {
	r.Write([]byte(response))
	return r
}