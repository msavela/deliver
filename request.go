package deliver

import (
	"net/http"
)

type Request struct {
	*http.Request
	Context			*Context
	Params			map[string]string
	Splats			[]string
}

// Initialize a new request.
func NewRequest(req *http.Request, c *Context, params map[string]string, splats []string) *Request {
	r := &Request{}
	r.Request 	= req
	r.Context	= c
	r.Params	= params
	r.Splats	= splats
	return r
}