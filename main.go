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

	// curl http://localhost:8080/todos
	// shows:
	// [{"name":"Write presentation","completed":false,"due":"0001-01-01T00:00:00Z"},{"name":"Host meetup","completed":false,"due":"0001-01-01T00:00:00Z"}]
	// prints log:
	// 2016/03/30 10:00:34 GET	/todos	TodoIndex	57.063µs

}
