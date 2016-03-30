//This tutorial comes from: http://thenewstack.io/make-a-restful-json-api-go/

package main

import (
	"encoding/json"
	"fmt"
	"html"
	"log"
	"net/http"
	"time"
	// A powerful URL router and dispatcher for golang. http://www.gorillatoolkit.org/pkg/mux
	"github.com/gorilla/mux"
)

//Todo is the type of each todos.
type Todo struct {
	Name      string
	Completed bool
	Due       time.Time
}

//Todos is a slice of Todo.
type Todos []Todo

//Index is the handler of "/"
func Index(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Hello, %q", html.EscapeString(r.URL.Path))
}

//TodoIndex the index of Todo
func TodoIndex(w http.ResponseWriter, r *http.Request) {
	todos := Todos{
		Todo{Name: "Write presentation"},
		Todo{Name: "Host meetup"},
	}
	//An Encoder writes JSON objects to an output stream.
	//NewEncoder returns a new encoder that writes to w.
	//Encode writes the JSON encoding of v to the stream, followed by a newline character.
	json.NewEncoder(w).Encode(todos)
}

//TodoShow Show todos
func TodoShow(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	todoID := vars["todoId"]
	fmt.Fprintln(w, "Todo show: ", todoID)
}

func main() {

	router := mux.NewRouter().StrictSlash(true)
	//Using router instead of http.HandlFunc, This router only works to localhost:8080/
	router.HandleFunc("/", Index)
	router.HandleFunc("/todos", TodoIndex)

	//we added a variable in the route, called todoId
	router.HandleFunc("/todos/{todoId}", TodoShow)

	log.Fatal(http.ListenAndServe(":8080", router))

	// When runs: http://localhost:8080/test/test
	// comes: 404 page not found
	// When runs:  http://localhost:8080/todos
	// comes:
	/*
		[
			{
				Name: "Write presentation",
				Completed: false,
				Due: "0001-01-01T00:00:00Z"
			},
			{
				Name: "Host meetup",
				Completed: false,
				Due: "0001-01-01T00:00:00Z"
			}
		]
	*/
	// When runs:  http://localhost:8080/todos/myTodo
	// comes: Todo show:  myTodo

}
