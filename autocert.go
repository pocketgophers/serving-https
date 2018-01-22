package main

import (
	"log"
	"net/http"

	"golang.org/x/crypto/acme/autocert"
)

func main() {
	http.HandleFunc("/", hello)

	m := &autocert.Manager{
		Prompt:     autocert.AcceptTOS,
		HostPolicy: autocert.HostWhitelist("example.com"),
		Cache:      autocert.DirCache("golang-autocert"),
	}

	// HTTP Server
	go func() {
		log.Fatal(http.ListenAndServe(":http",
			m.HTTPHandler(nil)))
	}()

	// HTTPS Server
	log.Fatal(http.Serve(m.Listener(), nil))
}

func hello(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Hello, 世界"))
}
