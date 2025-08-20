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

// ListRechargePresets handles the request to list available recharge presets.
func (h *RechargeHandler) ListRechargePresets(c *gin.Context) {
	presets, err := h.rechargeService.ListRechargePresets()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to retrieve recharge presets"})
		return
	}
	c.JSON(http.StatusOK, presets)
}

// CreateUSDTOrder handles the request to create a new USDT recharge order.
func (h *RechargeHandler) CreateUSDTOrder(c *gin.Context) {
	var input service.CreateUSDTOrderInput
	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	userID, _ := c.Get("userID")
	id, _ := userID.(uint)

	order, err := h.rechargeService.CreateUSDTOrder(id, input)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusCreated, order)
}

// GenerateCoupon handles the request for a user to generate a coupon.
func (h *RechargeHandler) GenerateCoupon(c *gin.Context) {
	var input service.GenerateCouponInput
	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	userID, _ := c.Get("userID")
	id, _ := userID.(uint)

	coupon, err := h.rechargeService.GenerateCoupon(id, input)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusCreated, coupon)
}
