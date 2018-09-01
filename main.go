package main

import (
	"fmt"
	"net/http"
)

func handleRequest() {
	fmt.Printf("listening on 8080")
}

func main() {
	handleRequest()
	http.ListenAndServe(":8080", nil)
}
