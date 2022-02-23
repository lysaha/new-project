package main

import (
	"log"
	"net/http"
)

func Handler(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Hello, world!\nThis is my first local web server on GO!"))
}

func main() {
	http.HandleFunc("/", Handler)
	log.Println("Start HTTP server on port 8082")
	log.Fatal(http.ListenAndServe(":8082", nil))
}
