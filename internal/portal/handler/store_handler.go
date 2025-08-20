package handler

import (
	"net/http"
	"v2ray-saas-gemini/internal/portal/service"

	"github.com/gin-gonic/gin"
)

// StoreHandler handles store-related API requests.
type StoreHandler struct {
	storeService service.StoreService
}

// NewStoreHandler creates a new StoreHandler.
func NewStoreHandler() *StoreHandler {
	return &StoreHandler{
		storeService: service.StoreService{},
	}
}

// ListPlans handles the request to list all available plans.
func (h *StoreHandler) ListPlans(c *gin.Context) {
	plans, err := h.storeService.ListAvailablePlans()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to retrieve plans"})
		return
	}
	c.JSON(http.StatusOK, plans)
}

// PurchasePlan handles the request to purchase a plan.
func (h *StoreHandler) PurchasePlan(c *gin.Context) {
	var input service.PurchasePlanInput
	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// Get userID from the context
	userID, exists := c.Get("userID")
	if !exists {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "User ID not found in context"})
		return
	}
	id, _ := userID.(uint)

	// Call the service
	subscription, err := h.storeService.PurchasePlan(id, input)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusCreated, gin.H{
		"message":      "Plan purchased successfully",
		"subscription": subscription,
	})
}
