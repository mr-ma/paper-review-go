package main

import (
	"log"
	"net/http"
)

func main() {
	fs := http.FileServer(http.Dir("/Users/mohsen-tum/Go/src/github.com/mr-ma/frontend/src"))
	http.Handle("/", fs)

	log.Println("Listening...")
	http.ListenAndServe(":3000", nil)
}
