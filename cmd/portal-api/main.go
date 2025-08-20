package main

import (
	"fmt"
	"log"
	"net/http"
)

func main() {
	// A simple handler function
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintf(w, "Welcome to the Portal API!")
	})

	// Define the server address
	addr := ":8080"
	log.Printf("Portal API server starting on %s", addr)

	// Start the server
	if err := http.ListenAndServe(addr, nil); err != nil {
		log.Fatalf("Failed to start server: %v", err)
	}
}
