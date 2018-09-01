package main

import (
	"fmt"
	"net/http"
	"time"
)

type TodoItem struct {
	Done        bool
	Description string
	created     time.Time
	Due         time.Time
}

func handleIndex(resp http.ResponseWriter, r *http.Request) {
	resp.Write([]byte("Hello welcome to the api"))
}

func addTodo(resp http.ResponseWriter, r *http.Request) {
	todo := TodoItem{
		Done:        false,
		Description: "Something",
		created:     time.Now(),
		Due:         time.Now(),
	}

	fmt.Println(todo)
}

func handleRequest() {
	fmt.Printf("listening on 8080")
	http.HandleFunc("/", handleIndex)
	http.HandleFunc("/new", addTodo)
}

func main() {
	handleRequest()
	http.ListenAndServe(":8080", nil)
}
