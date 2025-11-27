# Logger Example

This example shows how to use the `Logger` middleware to log requests.

It starts a server on port 8080 and responds with "OK" to requests on the root path (`/`).
Each request will be logged to the console.

## How to Run

```bash
go run logger.go
```

## How to Use

```bash
curl http://localhost:8080/
```
