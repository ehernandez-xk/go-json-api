//This tutorial comes from: http://thenewstack.io/make-a-restful-json-api-go/

package main

import (
	"fmt"
	"html"
	"log"
	"net/http"
)

//This is the handler
func handler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Hello, %q", html.EscapeString(r.URL.Path))
}

func main() {
	//Everything that comes to the request is handled by handler func.
	http.HandleFunc("/", handler)
	log.Fatal(http.ListenAndServe(":8080", nil))

	// When runs: http://localhost:8080/test/test
	// comes: Hello, "/test/test"

}
