package handler

import (
	"net/http"
	"v2ray-saas-gemini/internal/portal/service"

	"github.com/gin-gonic/gin"
)

// NodeStatusHandler handles node status API requests.
type NodeStatusHandler struct {
	service service.NodeStatusService
}

func NewNodeStatusHandler() *NodeStatusHandler {
	return &NodeStatusHandler{service: service.NodeStatusService{}}
}

// ListAllNodeStatuses handles the request to list all node statuses.
func (h *NodeStatusHandler) ListAllNodeStatuses(c *gin.Context) {
	statuses, err := h.service.ListAllNodeStatuses()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to retrieve node statuses"})
		return
	}
	c.JSON(http.StatusOK, statuses)
}
