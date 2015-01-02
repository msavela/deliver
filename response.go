package deliver

import (
	"net/http"
)

type Response struct {
	http.ResponseWriter
}

// Initialize new response.
func NewResponse(res http.ResponseWriter) Response {
	r := Response{res}
	return r
}

// Set response status code.
func (r *Response) Status(status int) *Response {
	r.WriteHeader(status)
	return r
}

// Send response body.
func (r *Response) Send(response string) *Response {
	r.Write([]byte(response))
	return r
}