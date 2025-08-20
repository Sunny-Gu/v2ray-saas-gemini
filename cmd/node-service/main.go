package main

import (
	"fmt"
	"log"
	"net/http"
)

func main() {
	// A simple handler function
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintf(w, "Welcome to the Node Service!")
	})

	// Define the server address
	addr := ":8082"
	log.Printf("Node Service server starting on %s", addr)

	// Start the server
	if err := http.ListenAndServe(addr, nil); err != nil {
		log.Fatalf("Failed to start server: %v", err)
	}
}
