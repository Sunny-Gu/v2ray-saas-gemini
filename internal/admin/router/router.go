package router

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

// SetupRouter initializes and configures the Gin router for the admin API.
func SetupRouter() *gin.Engine {
	router := gin.Default()

	// API v1 group for admin
	apiV1 := router.Group("/api/v1/admin")
	// Here we would add an admin-specific authentication middleware
	// apiV1.Use(middleware.AdminAuthMiddleware())
	{
		// Health check for admin
		apiV1.GET("/ping", func(c *gin.Context) {
			c.JSON(http.StatusOK, gin.H{
				"message": "pong from admin",
			})
		})

		// Placeholder for other admin routes
		// e.g., apiV1.GET("/users", userHandler.GetUsers)
	}

	return router
}
