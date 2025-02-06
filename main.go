package main

import (
	"fmt"
	"net/http"
)

func main() {
	http.HandleFunc("/hello", Hello)
	fmt.Println("Listening on port 8080 ...")
	err := http.ListenAndServe(":8080", nil)
	if err != nil {
		fmt.Println("Error starting server:", err)
	}
}

func Hello(w http.ResponseWriter, r *http.Request) {
	_, err := fmt.Fprint(w, "Hello, world!")
	if err != nil {
		fmt.Println("Error writing response:", err)
		return
	}
}
