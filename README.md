# deliver [![Build Status](https://travis-ci.org/msavela/deliver.svg)](https://travis-ci.org/msavela/deliver)

`deliver` is HTTP router, middleware and context for Golang.

```go
func main() {
	d := deliver.New()

	d.GET("/", func(res deliver.Response, req *deliver.Request) {
		res.Status(200).Send("OK")
	})

	log.Fatal(http.ListenAndServe(":8080", d))
}
```