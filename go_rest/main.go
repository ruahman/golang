package main

import (
	"net/http"
)

// setup routes that we handle
type routes struct{}

// handles web requests
type handler struct{}

// ServeHTTP is the method to implement the http.Handler interface
// this processes all incoming requests
func (h *handler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Hello, World!"))
}

func main() {
	// multiplexer for requests
	mux := http.NewServeMux()

	// assign handler to the multiplexer
	mux.Handle("/", &handler{})

	// start the server
	http.ListenAndServe(":8080", mux)
}
