package deliver

import (
	"testing"
	"net/http"
	"net/http/httptest"
	"reflect"
)

func TestMiddlewareBasic(t *testing.T) {
	d := New()

	d.Use(MiddlewareHandlerFunc(func(res Response, req *Request, next func()) {
		res.Send("Hello")
	}))

	response, body := testMiddleware(t, d)

	expect(t, body, "Hello")
	expect(t, response.Code, http.StatusOK)
}

func TestMiddlewareMultiple(t *testing.T) {
	d := New()

	content := ""
	d.Use(MiddlewareHandlerFunc(func(res Response, req *Request, next func()) {
		content += "Hello"
		next()
	}))

	d.Use(MiddlewareHandlerFunc(func(res Response, req *Request, next func()) {
		content += "World"
		res.SetStatus(http.StatusOK)
	}))

	response, _ := testMiddleware(t, d)

	expect(t, content, "HelloWorld")
	expect(t, response.Code, http.StatusOK)
}

func TestMiddlewareMultipleAfter(t *testing.T) {
	d := New()

	content := ""
	d.Use(MiddlewareHandlerFunc(func(res Response, req *Request, next func()) {
		next()
		content += "Hello"
	}))

	d.Use(MiddlewareHandlerFunc(func(res Response, req *Request, next func()) {
		content += "World"
		res.SetStatus(http.StatusOK)
	}))

	response, _ := testMiddleware(t, d)

	expect(t, content, "WorldHello")
	expect(t, response.Code, http.StatusOK)
}

func TestMiddlewareMultipleInterrupt(t *testing.T) {
	d := New()

	content := ""
	d.Use(MiddlewareHandlerFunc(func(res Response, req *Request, next func()) {
		content += "Hello"
	}))

	d.Use(MiddlewareHandlerFunc(func(res Response, req *Request, next func()) {
		content += "Should not be called"
		res.SetStatus(http.StatusOK)
	}))

	response, _ := testMiddleware(t, d)

	expect(t, content, "Hello")
	expect(t, response.Code, http.StatusNotFound)
}

/* Helpers */
func testMiddleware(t *testing.T, deliver *Deliver) (*httptest.ResponseRecorder, string) {
	response := httptest.NewRecorder()
	deliver.ServeHTTP(response, (*http.Request)(nil))
	return response, response.Body.String()
}

func expect(t *testing.T, a interface{}, b interface{}) {
	if a != b {
		t.Errorf("Expected %v (%v) - Got %v (%v)", b, reflect.TypeOf(b), a, reflect.TypeOf(a))
	}
}