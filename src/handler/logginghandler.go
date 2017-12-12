package handler

import (
	"log"
	"net/http"
	"time"
)

// LoggingHandler Log Requests to Std.out
func LoggingHandler(inner http.Handler) http.Handler {
	fn := func(w http.ResponseWriter, r *http.Request) {
		start := time.Now()
		inner.ServeHTTP(w, r)
		log.Printf("[%s]\t %q\t %v\t", r.Method, r.URL.String(), time.Since(start))
	}
	return http.HandlerFunc(fn)
}
