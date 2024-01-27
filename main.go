package main

import (
	"fmt"
	"net/http"
)

func handlerFunc(w http.ResponseWriter, r *http.Request) {
	fmt.Fprint(w, "Hello, world!")
}

func main() {
	http.HandleFunc("/", handlerFunc)
	fmt.Println("Server started on port 8080")
	http.ListenAndServe(":8080", nil)
}
