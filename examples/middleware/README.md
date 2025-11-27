# Middleware Example

This example shows how to use different types of middleware with `deliver`.

It starts a server on port 8080 and responds with "OK" to requests on the root path (`/`).
It uses three different types of middleware to log messages to the console.

## How to Run

```bash
go run middleware.go
```

## How to Use

```bash
curl http://localhost:8080/
```
