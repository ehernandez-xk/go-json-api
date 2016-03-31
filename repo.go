package main

//Is incremented to get different ids
var currentID int

//This variable is used in TodoIndex handler and contains all todos
var todos Todos

// Give us some seed data, init function runs before main function
func init() {
	RepoCreateTodo(Todo{Name: "Write presentation"})
	RepoCreateTodo(Todo{Name: "Host meetup"})
	RepoCreateTodo(Todo{Name: "Visit friends"})
}

//RepoCreateTodo receives a todo and appends it to todos
func RepoCreateTodo(t Todo) Todo {
	currentID++
	t.ID = currentID
	todos = append(todos, t)
	return t
}

//RepoFindTodo returns a todo if found its id.
func RepoFindTodo(id int) Todo {
	for _, todo := range todos {
		if id == todo.ID {
			return todo
		}
	}
	return Todo{} //Empty todo if not found
}
