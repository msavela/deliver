# deliver [![Build Status](https://travis-ci.org/msavela/deliver.svg)](https://travis-ci.org/msavela/deliver)

`deliver` is HTTP router, middleware and context for Golang.

```go
func main() {
	d := deliver.New()

	d.GET("/", func(res deliver.Response, req *deliver.Request) {
		res.Status(http.StatusOK).Send(http.StatusText(http.StatusOK))
	})

	log.Fatal(http.ListenAndServe(":8080", d))
}
```