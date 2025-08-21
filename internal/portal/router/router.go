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
	subscriptionHandler := handler.NewSubscriptionHandler()
	nodeStatusHandler := handler.NewNodeStatusHandler()
	ticketHandler := handler.NewTicketHandler()
	contentHandler := handler.NewContentHandler()

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
				userRoutes.POST("/request-password-reset", userHandler.RequestPasswordReset)
				userRoutes.POST("/reset-password", userHandler.ResetPassword)
			}
			storeRoutes := public.Group("/store")
			{
				storeRoutes.GET("/plans", storeHandler.ListPlans)
				storeRoutes.GET("/recharge-presets", rechargeHandler.ListRechargePresets)
			}
			// Public node status route
			public.GET("/nodes/status", nodeStatusHandler.ListAllNodeStatuses)
			// Public content routes
			public.GET("/announcements", contentHandler.ListAnnouncements)
			public.GET("/help-documents", contentHandler.ListHelpDocuments)
		}

		// Authenticated routes
		authenticated := apiV1.Group("/")
		authenticated.Use(middleware.AuthMiddleware())
		{
			// Profile routes
			profileRoutes := authenticated.Group("/profile")
			{
				profileRoutes.GET("", profileHandler.GetProfile)
				profileRoutes.GET("/history", profileHandler.GetConsumptionHistory)
			}

			// Recharge routes
			rechargeRoutes := authenticated.Group("/recharge")
			{
				rechargeRoutes.POST("/redeem", rechargeHandler.RedeemCoupon)
				rechargeRoutes.POST("/create-order", rechargeHandler.CreateUSDTOrder)
				rechargeRoutes.POST("/generate-coupon", rechargeHandler.GenerateCoupon)
			}

			// Store routes (for purchasing)
			storeRoutes := authenticated.Group("/store")
			{
				storeRoutes.POST("/purchase", storeHandler.PurchasePlan)
			}

			// Subscription routes
			subscriptionRoutes := authenticated.Group("/subscription")
			{
				subscriptionRoutes.GET("/", subscriptionHandler.GetSubscriptionInfo)
				subscriptionRoutes.POST("/reset-link", subscriptionHandler.ResetSubscriptionLink)
			}

			// Ticket routes
			ticketRoutes := authenticated.Group("/tickets")
			{
				ticketRoutes.POST("/", ticketHandler.CreateTicket)
				ticketRoutes.GET("/", ticketHandler.ListUserTickets)
				ticketRoutes.GET("/:id", ticketHandler.GetTicketDetails)
				ticketRoutes.POST("/:id/reply", ticketHandler.ReplyToTicket)
			}
		}
	}

	return router
}
