//This tutorial comes from: http://thenewstack.io/make-a-restful-json-api-go/

package main

import (
	"log"
	"net/http"
)

func main() {

	//This receives the router. and the router cointains the handlers
	router := NewRouter()

	//ListenAndServe starts an HTTP server with a given address and handler
	log.Fatal(http.ListenAndServe(":8080", router))

	// run using: go run *.go

	//Using curl instead the browser, in other terminal

	// curl -H "Content-Type: application/json" -d '{"name":"New Todo"}' http://localhost:8080/todos
	// shows: {"id":5,"name":"New Todo","completed":false,"due":"0001-01-01T00:00:00Z"}

	// view all todos
	// curl http://localhost:8080/todos

}
