package main

import (
	"log"
	"net/http"
)

func main() {
	http.HandleFunc("/", hello)

	log.Fatal(http.ListenAndServeTLS(
		":https", "cert.pem", "key.pem", nil))
}

func hello(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Hello, 世界"))
}
