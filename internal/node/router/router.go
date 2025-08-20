package router

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

// SetupRouter initializes and configures the Gin router for the node service.
func SetupRouter() *gin.Engine {
	router := gin.Default()

	// Here we would add a middleware for authenticating requests from v2ray nodes,
	// likely using a pre-shared secret key.
	// router.Use(middleware.NodeAuthMiddleware())

	// API v1 group for node communications
	apiV1 := router.Group("/api/v1/node")
	{
		// Endpoint for v2ray nodes to fetch user configurations
		apiV1.GET("/users", func(c *gin.Context) {
			c.JSON(http.StatusOK, gin.H{"message": "user config endpoint"})
		})

		// Endpoint for v2ray nodes to report traffic usage
		apiV1.POST("/traffic", func(c *gin.Context) {
			c.JSON(http.StatusOK, gin.H{"message": "traffic report endpoint"})
		})
	}

	return router
}
