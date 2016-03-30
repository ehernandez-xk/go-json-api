package main

import (
	"encoding/json"
	"fmt"
	"html"
	"net/http"
	// A powerful URL router and dispatcher for golang. http://www.gorillatoolkit.org/pkg/mux
	"github.com/gorilla/mux"
)

//Index is the handler of "/"
func Index(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Hello, %q", html.EscapeString(r.URL.Path))
}

//TodoIndex the index of Todo
func TodoIndex(w http.ResponseWriter, r *http.Request) {

	// we are sending back our content type and telling the client to expect json.
	w.Header().Set("Content-Type", "application/json; charset-UTF-8")
	// we are explicitly setting the status code.
	w.WriteHeader(http.StatusOK)

	//An Encoder writes JSON objects to an output stream.
	//NewEncoder returns a new encoder that writes to w.
	//Encode writes the JSON encoding of v to the stream, followed by a newline character.
	if err := json.NewEncoder(w).Encode(todos); err != nil {
		panic(err)
	}
}

//TodoShow Show todos
func TodoShow(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	todoID := vars["todoId"]
	fmt.Fprintln(w, "Todo show: ", todoID)
}
