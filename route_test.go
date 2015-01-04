package deliver

import (
	"net/http"
	"testing"
)

type routeTest struct {
	title			string
	route			*Route
	request			*http.Request
	params			map[string]string
	splat			[]string
	shouldMatch		bool
}

func TestRoutes(t *testing.T) {
	tests := []routeTest{
		{
			title:			"Root route",
			route:			NewRoute("/", nil).GET(),
			request:		newRequest("GET", "/"),
			shouldMatch:	true,
		},
		{
			title:			"Root route should not match",
			route:			NewRoute("/", nil).GET(),
			request:		newRequest("GET", "/route"),
			shouldMatch:	false,
		},
		{
			title:			"Root route should not match POST request",
			route:			NewRoute("/", nil).GET(),
			request:		newRequest("POST", "/route"),
			shouldMatch:	false,
		},
		{
			title:			"Basic route",
			route:			NewRoute("/hello", nil).GET(),
			request:		newRequest("GET", "/hello"),
			shouldMatch:	true,
		},
		{
			title:			"Basic route should match with / at the end",
			route:			NewRoute("/hello", nil).GET(),
			request:		newRequest("GET", "/hello/"),
			shouldMatch:	true,
		},
		{
			title:			"Basic route should not match",
			route:			NewRoute("/hello", nil).GET(),
			request:		newRequest("GET", "/"),
			shouldMatch:	false,
		},
		{
			title:			"Route with parameter",
			route:			NewRoute("/post/:id", nil).GET(),
			request:		newRequest("GET", "/post/43"),
			shouldMatch:	true,
		},
		{
			title:			"Route with optional parameter",
			route:			NewRoute("/post/:id?", nil).GET(),
			request:		newRequest("GET", "/post/24"),
			shouldMatch:	true,
		},
		{
			title:			"Route with optional parameter not given",
			route:			NewRoute("/post/:id?", nil).GET(),
			request:		newRequest("GET", "/post"),
			shouldMatch:	true,
		},
		{
			title:			"Route with optional param",
			route:			NewRoute("/post/:id?", nil).GET(),
			request:		newRequest("GET", "/post/"),
			shouldMatch:	true,
		},
		{
			title:			"Route with parameter and format",
			route:			NewRoute("/post/:id.:format", nil).GET(),
			request:		newRequest("GET", "/post/45.json"),
			shouldMatch:	true,
		},
		{
			title:			"Route with parameter and optional format",
			route:			NewRoute("/post/:id.:format?", nil).GET(),
			request:		newRequest("GET", "/post/45.json"),
			shouldMatch:	true,
		},
		{
			title:			"Route with parameter and optional format not given",
			route:			NewRoute("/post/:id.:format?", nil).GET(),
			request:		newRequest("GET", "/post/45."),
			shouldMatch:	true,
		},
		{
			title:			"Route with parameter and optional format not given",
			route:			NewRoute("/post/:id.:format?", nil).GET(),
			request:		newRequest("GET", "/post/45"),
			shouldMatch:	true,
		},
		{
			title:			"Route with splat",
			route:			NewRoute("/post/*", nil).GET(),
			request:		newRequest("GET", "/post/45"),
			shouldMatch:	true,
		},
		{
			title:			"Route with splat with format",
			route:			NewRoute("/post/*.*", nil).GET(),
			request:		newRequest("GET", "/post/45.json"),
			shouldMatch:	true,
		},
		{
			title:			"Route with custom regex",
			route:			NewRoute("/lang/:lang([a-z]{2})", nil).GET(),
			request:		newRequest("GET", "/lang/en"),
			shouldMatch:	true,
		},
		{
			title:			"Route with custom regex invalid",
			route:			NewRoute("/lang/:lang([a-z]{2})", nil).GET(),
			request:		newRequest("GET", "/lang/eng"),
			shouldMatch:	false,
		},
		{
			title:			"Route with custom regex invalid number",
			route:			NewRoute("/lang/:lang([a-z]{2})", nil).GET(),
			request:		newRequest("GET", "/lang/23"),
			shouldMatch:	false,
		},
		{
			title:			"Basic route should match lower case url",
			route:			NewRoute("/Hello", nil).GET(),
			request:		newRequest("GET", "/hello"),
			shouldMatch:	true,
		},
	}

	// Test each route
	for _, test := range tests {
		testRoute(t, test)
	}
}

func testRoute(t *testing.T, test routeTest) {
	match := test.route.Match(test.request)

	if match != nil != test.shouldMatch {
		msg := "should match"
		if !test.shouldMatch {
			msg = "should not match"
		}
		t.Errorf("'%v' %v:\nRoute: %#v\nRequest: %#v\n", test.title, msg, test.route, test.request)
		return
	}
}

// Helper method to create a new request.
func newRequest(method, url string) *http.Request {
	req, err := http.NewRequest(method, url, nil)
	if err != nil {
		panic(err)
	}
	return req
}