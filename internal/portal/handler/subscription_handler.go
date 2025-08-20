package handler

import (
	"net/http"
	"v2ray-saas-gemini/internal/portal/service"

	"github.com/gin-gonic/gin"
)

// SubscriptionHandler handles subscription-related API requests.
type SubscriptionHandler struct {
	service service.SubscriptionService
}

func NewSubscriptionHandler() *SubscriptionHandler {
	return &SubscriptionHandler{service: service.SubscriptionService{}}
}

// GetSubscriptionInfo handles the request to get the user's subscription info.
func (h *SubscriptionHandler) GetSubscriptionInfo(c *gin.Context) {
	userID, _ := c.Get("userID")
	id, _ := userID.(uint)

	sub, err := h.service.GetSubscriptionInfo(id)
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, sub)
}

// ResetSubscriptionLink handles the request to reset the subscription link.
func (h *SubscriptionHandler) ResetSubscriptionLink(c *gin.Context) {
	userID, _ := c.Get("userID")
	id, _ := userID.(uint)

	sub, err := h.service.ResetSubscriptionLink(id)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{
		"message":      "Subscription link reset successfully",
		"subscription": sub,
	})
}
