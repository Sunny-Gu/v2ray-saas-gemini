package router

import (
	"net/http"
	"v2ray-saas-gemini/internal/portal/handler"

	"github.com/gin-gonic/gin"
)

// SetupRouter initializes and configures the Gin router.
func SetupRouter() *gin.Engine {
	// Create a new Gin router with default middleware (logger, recovery).
	router := gin.Default()

	// Health check endpoint
	router.GET("/ping", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"message": "pong",
		})
	})

	// Create handlers
	userHandler := handler.NewUserHandler()

	// Group API routes under /api/v1
	apiV1 := router.Group("/api/v1")
	{
		// User routes
		userRoutes := apiV1.Group("/user")
		{
			userRoutes.POST("/register", userHandler.Register)
			// Placeholder for user login
			userRoutes.POST("/login", func(c *gin.Context) {
				c.JSON(http.StatusOK, gin.H{"message": "login endpoint"})
			})
		}
	}

	return router
}
