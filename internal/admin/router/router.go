package router

import (
	"net/http"
	"v2ray-saas-gemini/internal/admin/handler"
	"v2ray-saas-gemini/internal/admin/middleware"

	"github.com/gin-gonic/gin"
)

// SetupRouter initializes and configures the Gin router for the admin API.
func SetupRouter() *gin.Engine {
	router := gin.Default()

	// Handlers
	authHandler := handler.NewAuthHandler()
	nodeHandler := handler.NewNodeHandler()
	rechargeConfigHandler := handler.NewRechargeConfigHandler()
	userManagementHandler := handler.NewUserManagementHandler()
	planManagementHandler := handler.NewPlanManagementHandler()

	// API v1 group for admin
	apiV1 := router.Group("/api/v1/admin")
	{
		// Public route for login
		apiV1.POST("/login", authHandler.Login)

		// Authenticated routes
		authRequired := apiV1.Group("/")
		authRequired.Use(middleware.AdminAuthMiddleware())
		{
			// Health check for admin
			authRequired.GET("/ping", func(c *gin.Context) {
				c.JSON(http.StatusOK, gin.H{
					"message": "pong from admin",
				})
			})

			// Node management routes
			nodeRoutes := authRequired.Group("/nodes")
			{
				nodeRoutes.POST("/", nodeHandler.CreateNode)
				nodeRoutes.GET("/", nodeHandler.ListNodes)
				nodeRoutes.GET("/:id", nodeHandler.GetNode)
				nodeRoutes.PUT("/:id", nodeHandler.UpdateNode)
				nodeRoutes.DELETE("/:id", nodeHandler.DeleteNode)
			}

			// Recharge preset management routes
			rechargeRoutes := authRequired.Group("/recharge-presets")
			{
				rechargeRoutes.POST("/", rechargeConfigHandler.CreatePreset)
				rechargeRoutes.GET("/", rechargeConfigHandler.ListPresets)
				rechargeRoutes.PUT("/:id", rechargeConfigHandler.UpdatePreset)
				rechargeRoutes.DELETE("/:id", rechargeConfigHandler.DeletePreset)
			}

			// User management routes
			userRoutes := authRequired.Group("/users")
			{
				userRoutes.GET("/", userManagementHandler.ListUsers)
				userRoutes.PATCH("/:id/status", userManagementHandler.UpdateUserStatus)
			}

			// Plan management routes
			planRoutes := authRequired.Group("/plans")
			{
				planRoutes.POST("/", planManagementHandler.CreatePlan)
				planRoutes.GET("/", planManagementHandler.ListPlans)
				planRoutes.GET("/:id", planManagementHandler.GetPlan)
				planRoutes.PUT("/:id", planManagementHandler.UpdatePlan)
				planRoutes.DELETE("/:id", planManagementHandler.DeletePlan)
			}

			// Order management routes
			orderRoutes := authRequired.Group("/orders")
			{
				orderRoutes.GET("/recharge", orderHandler.ListRechargeOrders)
				orderRoutes.GET("/subscription", orderHandler.ListSubscriptionOrders)
			}

			// Content management routes
			contentRoutes := authRequired.Group("/content")
			{
				// Announcements
				contentRoutes.POST("/announcements", contentManagementHandler.CreateAnnouncement)
				contentRoutes.GET("/announcements", contentManagementHandler.ListAnnouncements)
				contentRoutes.PUT("/announcements/:id", contentManagementHandler.UpdateAnnouncement)
				contentRoutes.DELETE("/announcements/:id", contentManagementHandler.DeleteAnnouncement)
				// Help Documents
				contentRoutes.POST("/help-documents", contentManagementHandler.CreateHelpDocument)
				contentRoutes.GET("/help-documents", contentManagementHandler.ListHelpDocuments)
				contentRoutes.PUT("/help-documents/:id", contentManagementHandler.UpdateHelpDocument)
				contentRoutes.DELETE("/help-documents/:id", contentManagementHandler.DeleteHelpDocument)
			}

			// Audit routes
			auditRoutes := authRequired.Group("/audit")
			{
				auditRoutes.GET("/financial-overview", auditHandler.GetFinancialOverview)
			}
		}
	}

	return router
}
