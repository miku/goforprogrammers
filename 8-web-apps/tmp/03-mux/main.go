package main

import (
	"fmt"
	"log"
	"net/http"
)

// Index is a basic HTTP Handler.
func Index(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(w, "Welcome")
}

// Echo is a basic HTTP Handler.
func Echo(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "You asked to %s %s\n", r.Method, r.URL.Path)
}

func main() {

	// Uses the http.DefaultServeMux to dispatch handlers.
	http.HandleFunc("/", Index)
	http.HandleFunc("/echo", Echo)

	// Start a server listening on port 8000 and responding using Echo.
	if err := http.ListenAndServe("localhost:8000", nil); err != nil {
		log.Fatalf("error: listening and serving: %s", err)
	}
}
