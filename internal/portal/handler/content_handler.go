package handler

import (
	"net/http"
	"v2ray-saas-gemini/internal/portal/service"

	"github.com/gin-gonic/gin"
)

// ContentHandler handles content-related API requests.
type ContentHandler struct {
	service service.ContentService
}

func NewContentHandler() *ContentHandler {
	return &ContentHandler{service: service.ContentService{}}
}

// ListAnnouncements handles the request to list announcements.
func (h *ContentHandler) ListAnnouncements(c *gin.Context) {
	announcements, err := h.service.ListActiveAnnouncements()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to retrieve announcements"})
		return
	}
	c.JSON(http.StatusOK, announcements)
}

// ListHelpDocuments handles the request to list help documents.
func (h *ContentHandler) ListHelpDocuments(c *gin.Context) {
	documents, err := h.service.ListActiveHelpDocuments()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to retrieve help documents"})
		return
	}
	c.JSON(http.StatusOK, documents)
}
