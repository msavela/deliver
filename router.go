package deliver

import (
	"net/http"
)

type Router struct {
	// Slice of routes
	routes		[]*Route
	// Slice of global middleware
	middleware	[]*Middleware
}

// Initialize new router.
func NewRouter() *Router {
	return &Router{
		routes: []*Route{},
	}
}

// Add new route with HandlerFunc handler.
func (r *Router) Route(path string, handler HandlerFunc) *Route {
	route := NewRoute(path, handler)
	r.routes = append(r.routes, route)
	return route
}

// Add new route with http.HandlerFunc handler.
func (r *Router) RouteHandler(path string, handler http.HandlerFunc) *Route {
	route := NewRouteHandler(path, handler)
	r.routes = append(r.routes, route)
	return route
}

// Add new GET route.
func (r *Router) GET(path string, handler HandlerFunc) *Route {
	return r.Route(path, handler).GET()
}

// Add new POST route.
func (r *Router) POST(path string, handler HandlerFunc) *Route {
	return r.Route(path, handler).POST()
}

// Add new PUT route.
func (r *Router) PUT(path string, handler HandlerFunc) *Route {
	return r.Route(path, handler).PUT()
}

// Add new DELETE route.
func (r *Router) DELETE(path string, handler HandlerFunc) *Route {
	return r.Route(path, handler).DELETE()
}

func (r *Router) Use(handler HandlerNext) *Router {
	r.middleware = append(r.middleware, NewMiddleware(handler))
	return r
}

func (r *Router) UseHandler(handler http.Handler) *Router {
	r.middleware = append(r.middleware, NewMiddlewareHandler(handler))
	return r
}

func (r *Router) UseHandlerNext(handler HttpHandlerNext) *Router {
	r.middleware = append(r.middleware, NewMiddlewareHandlerNext(handler))
	return r
}

// ServeHTTP makes the router implement the http.Handler interface.
func (r *Router) ServeHTTP(w http.ResponseWriter, req *http.Request) {

	// Initialize request specific context
	context := &Context{}

	// Collect params and splats
	params := map[string]string{}
	splats := []string{}

	// Store matched route
	var routeMatch *Route

	// Loop through available routes for a match
	for _, route := range r.routes {
		if matches := route.Match(req); matches != nil && len(matches) > 0 {

			// Collect matches to params and splats
			for i, n := range matches[1:] {
				if route.keys[i] != "" {
					// Found a key, append to params
					params[route.keys[i]] = n
				} else {
					// No key found, append to splats
					splats = append(splats, n)
				}
			}

			// Match found, store route and break the loop
			routeMatch = route
			break
		}
	}

	response := NewResponse(w)
	request  := NewRequest(req, context, params, splats)

	// Middleware to handle
	// Include route specific middleware if any
	var middleware []*Middleware
	if routeMatch != nil {
		// Include both global and local middleware
		middleware = append(r.middleware, routeMatch.middleware...)
	} else {
		// Include global middleware only
		middleware = r.middleware
	}

	// Handle middleware
	r.handleMiddleware(response, request, middleware, func() {
		// Middleware done, proceed to route handler
		if routeMatch != nil {
			routeMatch.handler(response, request)
		}
	})

	// NotFound handler in case response not written
	if !response.Written() {
		http.NotFound(response, req)
	}
}

// Handle request middleware.
func (r *Router) handleMiddleware(res Response, req *Request, middleware []*Middleware, proceed func()) {
	index := 0

	var next func()
	next = func() {
		if len(middleware) == 0 || index > len(middleware) - 1 {
			// No more middleware available, proceed
			proceed()
			return
		}

		handler := middleware[index].handler
		index += 1

		// Call middleware
		handler.ServeHTTP(res, req, next)
	}

	// Loop through available middleware
	next()
}