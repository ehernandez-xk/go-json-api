//This tutorial comes from: http://thenewstack.io/make-a-restful-json-api-go/

package main

import (
	"fmt"
	"html"
	"log"
	"net/http"
	// A powerful URL router and dispatcher for golang. http://www.gorillatoolkit.org/pkg/mux
	"github.com/gorilla/mux"
)

//Index is the handler of "/"
func Index(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Hello, %q", html.EscapeString(r.URL.Path))
}

func main() {

	router := mux.NewRouter().StrictSlash(true)
	//Using router instead of http.HandlFunc, This router only works to localhost:8080/
	router.HandleFunc("/", Index)
	log.Fatal(http.ListenAndServe(":8080", router))

	// When runs: http://localhost:8080/test/test
	// comes: 404 page not found

}
