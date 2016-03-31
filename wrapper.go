package main

import (
	"log"
	"net/http"
	"time"
)

//commonHeaders using the wrapper technique adds common headers to the handlers
//https://medium.com/@matryer/the-http-handlerfunc-wrapper-technique-in-golang-c60bf76e6124
func commonHeaders(fn http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json; charset-UTF-8")
		fn(w, r)
	}
}

//Logger returns a handler that prints to the log.
func Logger(inner http.Handler, name string) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		start := time.Now()

		//Before the handler
		inner.ServeHTTP(w, r)

		//After the handler is finished print the log
		log.Printf("%s\t%s\t%s\t%s", r.Method, r.RequestURI, name, time.Since(start))

	})
}

/*
functions that take a http.HandlerFunc and return a new one can do things
before and/or after the handler is called, and even decide whether to call
the original handler at all.
*/
