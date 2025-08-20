package handler

import (
	"net/http"
	"v2ray-saas-gemini/internal/portal/service"

	"github.com/gin-gonic/gin"
)

// ProfileHandler handles profile-related API requests.
type ProfileHandler struct {
	profileService service.ProfileService
}

// NewProfileHandler creates a new ProfileHandler.
func NewProfileHandler() *ProfileHandler {
	return &ProfileHandler{
		profileService: service.ProfileService{},
	}
}

// GetProfile handles the request to get user profile data.
func (h *ProfileHandler) GetProfile(c *gin.Context) {
	// Get userID from the context (set by the auth middleware)
	userID, exists := c.Get("userID")
	if !exists {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "User ID not found in context"})
		return
	}

	// Assert userID to the correct type (uint)
	id, ok := userID.(uint)
	if !ok {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Invalid user ID type in context"})
		return
	}

	// Call the service to get profile data
	profileData, err := h.profileService.GetProfileData(id)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, profileData)
}
