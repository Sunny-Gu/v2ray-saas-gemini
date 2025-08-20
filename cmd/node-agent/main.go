package main

import (
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
)

func main() {
	// Start the background traffic reporter
	go startTrafficReporter()

	router := gin.Default()

	// Endpoint to receive new v2ray config from node-service
	router.POST("/update-config", func(c *gin.Context) {
		// In a real implementation, this would:
		// 1. Receive the new config JSON in the request body.
		// 2. Write it to the local v2ray config.json file.
		// 3. Execute a command to reload the v2ray service.
		c.JSON(http.StatusOK, gin.H{"message": "Configuration update received."})
	})

	// Endpoint for node-service to query this agent's status
	router.GET("/status", func(c *gin.Context) {
		// In a real implementation, this would gather CPU, memory, etc.
		c.JSON(http.StatusOK, gin.H{
			"cpu_load": "15%",
			"memory_usage": "40%",
		})
	})

	addr := ":8083" // The agent will listen on a different port
	log.Printf("Node Agent server starting on %s", addr)

	if err := router.Run(addr); err != nil {
		log.Fatalf("Failed to start node agent: %v", err)
	}
}
