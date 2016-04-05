//This tutorial comes from: http://thenewstack.io/make-a-restful-json-api-go/

package main

import (
	"log"
)

func main() {

	//This receives the router. and the router cointains the handlers
	router := NewRouter()

	//ListenAndServe starts an HTTP server with a given address and handler
	log.Fatal(http.ListenAndServe(":8080", router))

}
