package main

import (
	"log"
	"net/http"
	"time"
)

//Logger returns a handler that prints to the log.
func Logger(inner http.Handler, name string) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		start := time.Now()

		//Server the handler
		inner.ServeHTTP(w, r)

		//After the handler is finished print the log
		log.Printf("%s\t%s\t%s\t%s", r.Method, r.RequestURI, name, time.Since(start))

	})
}
