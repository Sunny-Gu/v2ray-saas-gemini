package handler

import (
	"net/http"
	"v2ray-saas-gemini/internal/admin/service"

	"github.com/gin-gonic/gin"
)

// AuditHandler handles financial audit API requests.
type AuditHandler struct {
	service service.AuditService
}

func NewAuditHandler() *AuditHandler {
	return &AuditHandler{service: service.AuditService{}}
}

// GetFinancialOverview handles the request to get a financial overview.
func (h *AuditHandler) GetFinancialOverview(c *gin.Context) {
	overview, err := h.service.GetFinancialOverview()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to generate financial overview"})
		return
	}
	c.JSON(http.StatusOK, overview)
}
