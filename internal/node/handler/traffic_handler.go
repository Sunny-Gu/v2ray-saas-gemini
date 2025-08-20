package handler

import (
	"net/http"
	"v2ray-saas-gemini/internal/node/service"

	"github.com/gin-gonic/gin"
)

// TrafficHandler handles traffic reporting requests from nodes.
type TrafficHandler struct {
	trafficService service.TrafficService
}

// NewTrafficHandler creates a new TrafficHandler.
func NewTrafficHandler() *TrafficHandler {
	return &TrafficHandler{
		trafficService: service.TrafficService{},
	}
}

// ReportTraffic handles the traffic usage report from a node.
func (h *TrafficHandler) ReportTraffic(c *gin.Context) {
	var input service.TrafficReportInput
	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid request body"})
		return
	}

	if err := h.trafficService.UpdateUserTraffic(input); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to process traffic report"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"status": "success", "message": "Traffic report processed"})
}
