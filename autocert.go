package main

import (
	"log"
	"net/http"

	"golang.org/x/crypto/acme/autocert"
)

func main() {
	http.HandleFunc("/", hello)

	l := autocert.NewListener("example.com")
	log.Fatal(http.Serve(l, nil))
}

func hello(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Hello, 世界"))
}
