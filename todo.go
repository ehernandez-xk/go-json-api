package main

import "time"

//Todo is the type of each todos.
type Todo struct {
	ID        int       `json:"id"`
	Name      string    `json:"name"`
	Completed bool      `json:"completed"`
	Due       time.Time `json:"due"`
}

//Todos is a slice of Todo.
type Todos []Todo

//OK helps to validate data to Todo from client request
func (t Todo) OK() error {
	if len(t.Name) == 0 {
		return ErrorMissingField()
	}
	return nil
}
