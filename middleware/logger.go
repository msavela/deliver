package middleware

import (
	"log"
	"net/http"
	"os"
	"time"
)

type Logger struct {
	*log.Logger
}

// New logger instance.
func NewLogger() *Logger {
	return NewLoggerPrefix("")
}

// New logger instance with custom prefix.
func NewLoggerPrefix(prefix string) *Logger {
	return &Logger{log.New(os.Stdout, prefix, log.Ldate | log.Ltime)}
}

func (l *Logger) ServeHTTP(rw http.ResponseWriter, r *http.Request, next func()) {
	start := time.Now()
	l.Printf("%s %s", r.Method, r.URL.Path)

	// Proceed to next handler
	next()

	// Printed once the request is completed
	l.Printf("Completed in %v", time.Since(start))
}