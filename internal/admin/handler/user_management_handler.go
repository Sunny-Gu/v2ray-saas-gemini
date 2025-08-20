package handler

import (
	"net/http"
	"strconv"
	"v2ray-saas-gemini/internal/admin/service"

	"github.com/gin-gonic/gin"
)

// UserManagementHandler handles user management API requests.
type UserManagementHandler struct {
	service service.UserManagementService
}

func NewUserManagementHandler() *UserManagementHandler {
	return &UserManagementHandler{service: service.UserManagementService{}}
}

// ListUsers handles the request to list users.
func (h *UserManagementHandler) ListUsers(c *gin.Context) {
	var input service.ListUsersInput
	if err := c.ShouldBindQuery(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	result, err := h.service.ListUsers(input)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to retrieve users"})
		return
	}
	c.JSON(http.StatusOK, result)
}

// UpdateUserStatus handles the request to update a user's status.
func (h *UserManagementHandler) UpdateUserStatus(c *gin.Context) {
	id, _ := strconv.ParseUint(c.Param("id"), 10, 32)

	var input service.UpdateUserStatusInput
	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	user, err := h.service.UpdateUserStatus(uint(id), input)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to update user status"})
		return
	}
	c.JSON(http.StatusOK, user)
}
