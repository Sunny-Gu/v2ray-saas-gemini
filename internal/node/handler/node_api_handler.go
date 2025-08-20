package handler

import (
	"net/http"
	"v2ray-saas-gemini/internal/node/service"

	"github.com/gin-gonic/gin"
)

// NodeApiHandler handles requests from v2ray nodes.
type NodeApiHandler struct {
	nodeApiService service.NodeApiService
}

// NewNodeApiHandler creates a new NodeApiHandler.
func NewNodeApiHandler() *NodeApiHandler {
	return &NodeApiHandler{
		nodeApiService: service.NodeApiService{},
	}
}

// GetActiveUsers handles the request from nodes to get the list of active users.
func (h *NodeApiHandler) GetActiveUsers(c *gin.Context) {
	users, err := h.nodeApiService.GetActiveUsers()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to retrieve active users"})
		return
	}
	c.JSON(http.StatusOK, gin.H{
		"status": "success",
		"users":  users,
	})
}
