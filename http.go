package main

import (
	"log"
	"net/http"
)

func main() {
	http.HandleFunc("/", hello)

	log.Fatal(http.ListenAndServe(":http", nil))
}

func hello(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Hello, 世界"))
}
