# go-json-api

drone.io Status

[![Build Status](https://drone.io/github.com/ehernandez-xk/go-json-api/status.png)](https://drone.io/github.com/ehernandez-xk/go-json-api/latest)

###To run the program

``go run *.go``

or to generate the binary (in the go-json-api directory)

``go build``

run

``go-json-api``

###View all todos

``http://localhost:8080/todos``

###Create a new Todo
```bash
curl -H "Content-Type: application/json" -d '{"name":"New Todo"}' http://localhost:8080/todos
```

Console ouput
```javascript
{"id":5,"name":"New Todo","completed":false,"due":"0001-01-01T00:00:00Z"}
```
###Find a todo
``http://localhost:8080/todos/3``


```javascript
{
id: 3,
name: "Visit friends",
completed: false,
due: "0001-01-01T00:00:00Z"
}
```

If todo is not found
```javascript
{
code: 404,
text: "Not found"
}
```

*Console log ouput
```bash
2016/03/31 11:54:08 GET	/todos	TodoIndex	167.244µs
2016/03/31 11:54:19 POST	/todos	TodoCreate	61.72µs
2016/03/31 11:54:40 GET	/todos/3	TodoShow	42.378µs
2016/03/31 11:54:53 GET	/todos/5	TodoShow	47.992µs
```
