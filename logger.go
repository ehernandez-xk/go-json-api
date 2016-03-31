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
