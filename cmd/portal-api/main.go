package main

import (
	"fmt"
	"log"
	"net/http"
	"os"
	"v2ray-saas-gemini/internal/config"
	"v2ray-saas-gemini/internal/database"
)

func main() {
	// Initialize configuration
	config.Init()

	// Set database DSN from config
	// This is a simple way to pass the DSN. A more robust solution might involve
	// passing the config object or DSN directly to the database.Init() function.
	os.Setenv("DB_DSN", config.Cfg.Database.DSN)
	database.Init()

	// A simple handler function
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintf(w, "Welcome to the Portal API!")
	})

	// Define the server address from config
	addr := config.Cfg.Server.PortalAPIPort
	log.Printf("Portal API server starting on %s", addr)

	// Start the server
	if err := http.ListenAndServe(addr, nil); err != nil {
		log.Fatalf("Failed to start server: %v", err)
	}
}
