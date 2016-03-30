# go-json-api

###To run the program

``go run *.go``

###View all todos

``http://localhost:8080/todos``

###Create a new Todo
``curl -H "Content-Type: application/json" -d '{"name":"New Todo"}' http://localhost:8080/todos``

Console ouput

{"id":5,"name":"New Todo","completed":false,"due":"0001-01-01T00:00:00Z"}

###Find a todo
``http://localhost:8080/todos/3``

If todo is not found

```
{
code: 404,
text: "Not found"
}
```

*Console log ouput
```
2016/03/30 15:39:15 GET	/todos/4	TodoShow	132.27µs
2016/03/30 15:39:19 POST	/todos	TodoCreate	118.788µs
2016/03/30 15:39:21 GET	/todos/4	TodoShow	22.51µs
2016/03/30 15:39:23 GET	/todos/5	TodoShow	13.613µs
```
