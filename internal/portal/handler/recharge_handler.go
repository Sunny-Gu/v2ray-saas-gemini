package handler

import (
	"net/http"
	"v2ray-saas-gemini/internal/portal/service"

	"github.com/gin-gonic/gin"
)

// RechargeHandler handles recharge-related API requests.
type RechargeHandler struct {
	rechargeService service.RechargeService
}

// NewRechargeHandler creates a new RechargeHandler.
func NewRechargeHandler() *RechargeHandler {
	return &RechargeHandler{
		rechargeService: service.RechargeService{},
	}
}

// RedeemCoupon handles the request to redeem a recharge code.
func (h *RechargeHandler) RedeemCoupon(c *gin.Context) {
	var input service.RedeemCouponInput
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
	if err := h.rechargeService.RedeemCoupon(id, input); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Coupon redeemed successfully"})
}
