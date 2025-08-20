package main

import (
	"log"
	"os"
	"v2ray-saas-gemini/internal/config"
	"v2ray-saas-gemini/internal/database"
	"v2ray-saas-gemini/internal/node/router"
)

func main() {
	// Initialize configuration
	config.Init()

	// Initialize database
	os.Setenv("DB_DSN", config.Cfg.Database.DSN)
	database.Init()

	// Setup the router
	r := router.SetupRouter()

	// Get server address from config
	addr := config.Cfg.Server.NodeServicePort
	log.Printf("Node Service server starting on %s", addr)

	// Start the server
	if err := r.Run(addr); err != nil {
		log.Fatalf("Failed to start server: %v", err)
	}
}
