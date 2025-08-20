package router

import (
	"net/http"
	"v2ray-saas-gemini/internal/admin/handler"

	"github.com/gin-gonic/gin"
)

// SetupRouter initializes and configures the Gin router for the admin API.
func SetupRouter() *gin.Engine {
	router := gin.Default()

	// Handlers
	nodeHandler := handler.NewNodeHandler()

	// API v1 group for admin
	apiV1 := router.Group("/api/v1/admin")
	// TODO: Add admin-specific authentication middleware
	// apiV1.Use(middleware.AdminAuthMiddleware())
	{
		// Health check for admin
		apiV1.GET("/ping", func(c *gin.Context) {
			c.JSON(http.StatusOK, gin.H{
				"message": "pong from admin",
			})
		})

		// Node management routes
		nodeRoutes := apiV1.Group("/nodes")
		{
			nodeRoutes.POST("/", nodeHandler.CreateNode)
			nodeRoutes.GET("/", nodeHandler.ListNodes)
			nodeRoutes.GET("/:id", nodeHandler.GetNode)
			nodeRoutes.PUT("/:id", nodeHandler.UpdateNode)
			nodeRoutes.DELETE("/:id", nodeHandler.DeleteNode)
		}
	}

	return router
}
