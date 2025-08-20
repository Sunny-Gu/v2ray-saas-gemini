package handler

import (
	"net/http"
	"v2ray-saas-gemini/internal/admin/service"

	"github.com/gin-gonic/gin"
)

// AuthHandler handles admin authentication requests.
type AuthHandler struct {
	service service.AuthService
}

func NewAuthHandler() *AuthHandler {
	return &AuthHandler{service: service.AuthService{}}
}

// Login handles the admin login request.
func (h *AuthHandler) Login(c *gin.Context) {
	var input service.LoginInput
	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	token, err := h.service.Login(input)
	if err != nil {
		c.JSON(http.StatusUnauthorized, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"token": token})
}
