# Session Example

This example shows how to use middleware to simulate a session.

It starts a server on port 8080 and responds with the username from the session.
A middleware is used to set a `session` object in the request context.

## How to Run

```bash
go run session.go
```

## How to Use

```bash
curl http://localhost:8080/
```
