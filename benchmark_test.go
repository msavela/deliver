package deliver

import (
	"net/http"
	"testing"
)

func BenchmarkRouter(b *testing.B) {
	d := New()
	d.GET("/v1/:param", func(res Response, req *Request) {})

	request, _ := http.NewRequest("GET", "/v1/anything", nil)
	for i := 0; i < b.N; i++ {
		d.ServeHTTP(nil, request)
	}
}