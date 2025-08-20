package handler

import (
	"net/http"
	"strconv"
	"v2ray-saas-gemini/internal/admin/service"

	"github.com/gin-gonic/gin"
)

// OrderHandler handles order management API requests.
type OrderHandler struct {
	service service.OrderService
}

func NewOrderHandler() *OrderHandler {
	return &OrderHandler{service: service.OrderService{}}
}

func (h *OrderHandler) ListRechargeOrders(c *gin.Context) {
	page, _ := strconv.Atoi(c.DefaultQuery("page", "1"))
	pageSize, _ := strconv.Atoi(c.DefaultQuery("page_size", "10"))

	orders, total, err := h.service.ListRechargeOrders(page, pageSize)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to retrieve recharge orders"})
		return
	}
	c.JSON(http.StatusOK, gin.H{
		"orders":     orders,
		"total":      total,
		"page":       page,
		"page_size":  pageSize,
	})
}

func (h *OrderHandler) ListSubscriptionOrders(c *gin.Context) {
	page, _ := strconv.Atoi(c.DefaultQuery("page", "1"))
	pageSize, _ := strconv.Atoi(c.DefaultQuery("page_size", "10"))

	orders, total, err := h.service.ListSubscriptionOrders(page, pageSize)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to retrieve subscription orders"})
		return
	}
	c.JSON(http.StatusOK, gin.H{
		"orders":     orders,
		"total":      total,
		"page":       page,
		"page_size":  pageSize,
	})
}
