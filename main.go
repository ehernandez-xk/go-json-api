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

	// When runs: http://localhost:8080/test/test
	// comes: 404 page not found
	// When runs:  http://localhost:8080/todos
	// comes:
	/*
		[
			{
				name: "Write presentation",
				completed: false,
				due: "0001-01-01T00:00:00Z"
			},
			{
				name: "Host meetup",
				completed: false,
				due: "0001-01-01T00:00:00Z"
			}
		]
	*/
	// When runs:  http://localhost:8080/todos/myTodo
	// comes: Todo show:  myTodo

}
