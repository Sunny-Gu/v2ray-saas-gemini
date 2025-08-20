package handler

import (
	"net/http"
	"strconv"
	"v2ray-saas-gemini/internal/admin/service"

	"github.com/gin-gonic/gin"
)

// NodeHandler handles node-related API requests.
type NodeHandler struct {
	nodeService service.NodeService
}

// NewNodeHandler creates a new NodeHandler.
func NewNodeHandler() *NodeHandler {
	return &NodeHandler{
		nodeService: service.NodeService{},
	}
}

// CreateNode handles the request to create a new node.
func (h *NodeHandler) CreateNode(c *gin.Context) {
	var input service.CreateNodeInput
	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	node, err := h.nodeService.CreateNode(input)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to create node"})
		return
	}
	c.JSON(http.StatusCreated, node)
}

// GetNode handles the request to get a single node.
func (h *NodeHandler) GetNode(c *gin.Context) {
	id, err := strconv.ParseUint(c.Param("id"), 10, 32)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid node ID"})
		return
	}
	node, err := h.nodeService.GetNode(uint(id))
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, node)
}

// ListNodes handles the request to list all nodes.
func (h *NodeHandler) ListNodes(c *gin.Context) {
	nodes, err := h.nodeService.ListNodes()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to retrieve nodes"})
		return
	}
	c.JSON(http.StatusOK, nodes)
}

// UpdateNode handles the request to update a node.
func (h *NodeHandler) UpdateNode(c *gin.Context) {
	id, err := strconv.ParseUint(c.Param("id"), 10, 32)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid node ID"})
		return
	}
	var input service.UpdateNodeInput
	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	node, err := h.nodeService.UpdateNode(uint(id), input)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, node)
}

// DeleteNode handles the request to delete a node.
func (h *NodeHandler) DeleteNode(c *gin.Context) {
	id, err := strconv.ParseUint(c.Param("id"), 10, 32)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid node ID"})
		return
	}
	if err := h.nodeService.DeleteNode(uint(id)); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to delete node"})
		return
	}
	c.JSON(http.StatusOK, gin.H{"message": "Node deleted successfully"})
}
