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
	rechargeHandler := handler.NewRechargeHandler()
	storeHandler := handler.NewStoreHandler()

	// API v1 group
	apiV1 := router.Group("/api/v1")
	{
		// Public routes
		public := apiV1.Group("/")
		{
			userRoutes := public.Group("/user")
			{
				userRoutes.POST("/register", userHandler.Register)
				userRoutes.POST("/login", userHandler.Login)
			}
			storeRoutes := public.Group("/store")
			{
				storeRoutes.GET("/plans", storeHandler.ListPlans)
				storeRoutes.GET("/recharge-presets", rechargeHandler.ListRechargePresets)
			}
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

			// Recharge routes
			rechargeRoutes := authenticated.Group("/recharge")
			{
				rechargeRoutes.POST("/redeem", rechargeHandler.RedeemCoupon)
				rechargeRoutes.POST("/create-order", rechargeHandler.CreateUSDTOrder)
			}

			// Store routes (for purchasing)
			storeRoutes := authenticated.Group("/store")
			{
				storeRoutes.POST("/purchase", storeHandler.PurchasePlan)
			}
		}
	}

	return router
}
