package main

import (
	"fmt"
	"log"
	"net/http"
)

func main() {
	// A simple handler function
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintf(w, "Welcome to the Admin API!")
	})

	// Define the server address
	addr := ":8081"
	log.Printf("Admin API server starting on %s", addr)

	// Start the server
	if err := http.ListenAndServe(addr, nil); err != nil {
		log.Fatalf("Failed to start server: %v", err)
	}
}
