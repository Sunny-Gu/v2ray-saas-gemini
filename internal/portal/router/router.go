package router

import (
	"net/http"
	"v2ray-saas-gemini/internal/portal/handler"
	"v2ray-saas-gemini/internal/portal/middleware"

	"github.com/gin-gonic/gin"
)

// SetupRouter initializes and configures the Gin router.
func SetupRouter() *gin.Engine {
	router := gin.Default()

	// Health check
	router.GET("/ping", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{"message": "pong"})
	})

	// Handlers
	userHandler := handler.NewUserHandler()
	profileHandler := handler.NewProfileHandler()

	// API v1 group
	apiV1 := router.Group("/api/v1")
	{
		// Public routes (no authentication required)
		public := apiV1.Group("/user")
		{
			public.POST("/register", userHandler.Register)
			public.POST("/login", userHandler.Login)
		}

		// Authenticated routes
		authenticated := apiV1.Group("/")
		authenticated.Use(middleware.AuthMiddleware())
		{
			// Profile routes
			profileRoutes := authenticated.Group("/profile")
			{
				profileRoutes.GET("/", profileHandler.GetProfile)
			}
		}
	}

	return router
}
