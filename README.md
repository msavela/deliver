# deliver [![Build Status](https://github.com/msavela/deliver/actions/workflows/build.yml/badge.svg)](https://github.com/msavela/deliver)

`deliver` is a lightweight, fast HTTP router, middleware and context for Golang. It is designed to be simple, easy to use, and highly performant.

## Installation

```bash
go get github.com/msavela/deliver
```

## Features

- Fast, trie-based router
- Supports middleware
- Request/response context
- Easy to use API

## Routing

`deliver` supports the standard HTTP methods: `GET`, `POST`, `PUT`, `DELETE`, `PATCH`, `HEAD`, `OPTIONS`.

```go
func main() {
    d := deliver.New()

    d.GET("/", func(res deliver.Response, req *deliver.Request) {
        res.Send("Hello, World!")
    })

    d.POST("/users", func(res deliver.Response, req *deliver.Request) {
        // Create a new user
    })

    d.PUT("/users/:id", func(res deliver.Response, req *deliver.Request) {
        // Update a user
    })

    log.Fatal(http.ListenAndServe(":8080", d))
}
```

## Middleware

`deliver` supports middleware that can be executed before or after a request.

```go
func main() {
    d := deliver.New()

    // Logger middleware
    d.UseHandlerNext(middleware.NewLogger())

    d.GET("/", func(res deliver.Response, req *deliver.Request) {
        res.Send("Hello, World!")
    })

    log.Fatal(http.ListenAndServe(":8080", d))
}
```

## Examples

You can find more examples in the `examples` directory:

- [Basic](examples/basic)
- [Logger](examples/logger)
- [Middleware](examples/middleware)
- [Session](examples/session)

## License

`deliver` is licensed under the [MIT License](LICENSE).
