package main

import (
	"fmt"
	"log"
	"net/http"
)

func handler(w http.ResponseWriter, r *http.Request) {
	userInput := r.URL.Path[1:]
	log.Printf("Endpoint requested by user: %s", userInput)
	fmt.Fprintf(w, "Hi there, I love %s!", userInput)
}

func main() {
	http.HandleFunc("/", handler)
	log.Fatal(http.ListenAndServe(":8080", nil))
}
