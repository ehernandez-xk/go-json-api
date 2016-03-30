package main

import (
	"encoding/json"
	"fmt"
	"html"
	"io"
	"io/ioutil"
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

//TodoCreate is the handler to create a todo
func TodoCreate(w http.ResponseWriter, r *http.Request) {

	var todo Todo

	//io.LimitReader. This is a good way to protect against malicious attacks on your server.
	body, err := ioutil.ReadAll(io.LimitReader(r.Body, 1048576))
	if err != nil {
		panic(err)
	}

	if err := r.Body.Close(); err != nil {
		panic(err)
	}
	// Unmarshal it to our Todo struct
	if err := json.Unmarshal(body, &todo); err != nil {
		w.Header().Set("Content-Type", "application/json; charset=UTF-8")
		w.WriteHeader(422) // unprocessable entity
		if err := json.NewEncoder(w).Encode(err); err != nil {
			panic(err)
		}
	}

	//t is the todo created.
	t := RepoCreateTodo(todo)
	w.Header().Set("Content-Type", "application/json; charset=UTF-8")

	//201 Created
	w.WriteHeader(http.StatusCreated)
	// Encode t to json
	if err := json.NewEncoder(w).Encode(t); err != nil {
		panic(err)
	}
}
