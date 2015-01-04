package deliver

import (
	"net/http"
	"regexp"
)

// Request method constants.
const (
	GET    	= "GET"
	POST   	= "POST"
	PUT    	= "PUT"
	DELETE 	= "DELETE"
	OPTIONS = "OPTIONS"
	HEAD	= "HEAD"
	TRACE	= "TRACE"
	CONNECT	= "CONNECT"
	PATCH	= "PATCH"
)

type (
	HandlerFunc func(Response, *Request)
	HttpHandlerFunc func(http.ResponseWriter, *http.Request)

	Route struct {
		re			*regexp.Regexp
		pattern		string
		methods		[]string
		keys		[]string
		params		map[int]string
		handler		HandlerFunc
		middleware	[]*Middleware
	}
)

// Initialize new route with HandlerFunc handler.
func NewRoute(pattern string, handler HandlerFunc) *Route {
	r := &Route{}
	r.pattern	= pattern
	r.re		= r.normalize(false, false) // Case-insensitive and not strict
	r.keys		= r.re.SubexpNames()[1:]
	r.handler	= handler
	r.methods	= []string{}
	return r
}

// Initialize new route with http.HandlerFunc handler.
func NewRouteHandler(pattern string, handler http.HandlerFunc) *Route {
	return NewRoute(pattern, HandlerFunc(func(res Response, req *Request) {
		handler(res.ResponseWriter, req.Request)
	}))
}

func (r *Route) GET() *Route {
	r.methods = append(r.methods, GET)
	return r
}

func (r *Route) POST() *Route {
	r.methods = append(r.methods, POST)
	return r
}

func (r *Route) PUT() *Route {
	r.methods = append(r.methods, PUT)
	return r
}

func (r *Route) DELETE() *Route {
	r.methods = append(r.methods, DELETE)
	return r
}

func (r *Route) Method(methods ...string) *Route {
	r.methods = append(r.methods, methods...)
	return r
}

func (r *Route) Use(handler HandlerNext) *Route {
	r.middleware = append(r.middleware, NewMiddleware(handler))
	return r
}

func (r *Route) UseHandler(handler http.Handler) *Route {
	r.middleware = append(r.middleware, NewMiddlewareHandler(handler))
	return r
}

func (r *Route) UseHandlerNext(handler HttpHandlerNext) *Route {
	r.middleware = append(r.middleware, NewMiddlewareHandlerNext(handler))
	return r
}

// Does route match request.
// Return slice of matches or nil in case no match.
func (r *Route) Match(req *http.Request) []string {
	if r.HasMethod(req.Method) {
		return r.re.FindStringSubmatch(req.URL.Path)
	}
	return nil
}

// Does route repond to particular request method.
func (r *Route) HasMethod(method string) bool {
	for _, m := range r.methods {
		if m == method {
			return true
		}
	}
	return false
}

// Convert string path to regexp.
//
// `sensitive` flag controls case-sensitivity.
// `strict` flag determines whether route accepts / at the end of the url.
func (r* Route) normalize(sensitive bool, strict bool) *regexp.Regexp {
	path := r.pattern

	// Accept / at the end of the url?
	if !strict { path += "/?" }

	path = regexp.MustCompile(`\/\(`).ReplaceAllString(path, "(?:/")

	// Construct regexp for each paramater
	parameterReplace := regexp.MustCompile(`(\/)?(\.)?:(\w+)(?:(\(.*?\)))?(\?)?(\*)?`)
	path = parameterReplace.ReplaceAllStringFunc(path, func(m string) string {

		submatch := parameterReplace.FindStringSubmatch(m)[1:]

		slash		:= submatch[0] // Slash
		format		:= submatch[1] // Format i.e. .json
		key			:= submatch[2] // Param name is always defined
		capture		:= submatch[3] // Custom regex
		optional	:= submatch[4] // Is optional?
		star		:= submatch[5] // Star

		// Begin constructiong single parameter
		param := ""

		if optional == "" { param += slash }

		// Begin named capture group
		param += "(?P<"+ key +">"

		if optional != "" { param += slash }

		param += format

		if capture != "" {
			param += capture
		} else {
			if format != "" {
				param += "(?:[^/.]+?)"
			} else {
				param += "(?:[^/]+?)"
			}
		}

		// End named capture group
		param += ")"

		param += optional

		if star != "" { param += "(/*)?" }

		// End constructiong single parameter
		return param
    })

	path = regexp.MustCompile(`([\/.])`).ReplaceAllString(path, "\\$1")
	path = regexp.MustCompile(`\*`).ReplaceAllString(path, "(.*)")

	// Case-insensitivity
	if !sensitive {
		path = "(?i)" + path
	}

	return regexp.MustCompile("^" + path + "$")
}