package main

import (
	"fmt"
	"log"
	"net/http"
)

// main is the entry point for the application.
// It starts a simple HTTP server.
func main() {
	// Register a handler function for the root ("/") URL path.
	// This function will be executed for every incoming request.
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		// Fprintf writes a formatted string to the writer, which is the HTTP response.
		fmt.Fprintf(w, "Hello from Go!")
	})

	// Announce that the server is starting.
	fmt.Println("Go server is starting on port 8081...")

	// ListenAndServe starts an HTTP server with a given address and handler.
	// The handler is nil, which means to use the DefaultServeMux.
	// It blocks until the server is shut down.
	if err := http.ListenAndServe(":8081", nil); err != nil {
		log.Fatalf("Failed to start Go server: %s", err)
	}
}
