package handler

import (
	"net/http"
	"strconv"
	"v2ray-saas-gemini/internal/admin/service"
	"v2ray-saas-gemini/internal/models"

	"github.com/gin-gonic/gin"
)

// ContentManagementHandler handles content management API requests.
type ContentManagementHandler struct {
	service service.ContentManagementService
}

func NewContentManagementHandler() *ContentManagementHandler {
	return &ContentManagementHandler{service: service.ContentManagementService{}}
}

// --- Announcements ---

func (h *ContentManagementHandler) CreateAnnouncement(c *gin.Context) {
	var input models.Announcement
	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	item, err := h.service.CreateAnnouncement(input)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to create announcement"})
		return
	}
	c.JSON(http.StatusCreated, item)
}

func (h *ContentManagementHandler) ListAnnouncements(c *gin.Context) {
	items, err := h.service.ListAnnouncements()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to list announcements"})
		return
	}
	c.JSON(http.StatusOK, items)
}

func (h *ContentManagementHandler) UpdateAnnouncement(c *gin.Context) {
	id, _ := strconv.ParseUint(c.Param("id"), 10, 32)
	var input models.Announcement
	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	item, err := h.service.UpdateAnnouncement(uint(id), input)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to update announcement"})
		return
	}
	c.JSON(http.StatusOK, item)
}

func (h *ContentManagementHandler) DeleteAnnouncement(c *gin.Context) {
	id, _ := strconv.ParseUint(c.Param("id"), 10, 32)
	if err := h.service.DeleteAnnouncement(uint(id)); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to delete announcement"})
		return
	}
	c.JSON(http.StatusOK, gin.H{"message": "Announcement deleted successfully"})
}

// --- Help Documents ---

func (h *ContentManagementHandler) CreateHelpDocument(c *gin.Context) {
	var input models.HelpDocument
	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	item, err := h.service.CreateHelpDocument(input)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to create help document"})
		return
	}
	c.JSON(http.StatusCreated, item)
}

func (h *ContentManagementHandler) ListHelpDocuments(c *gin.Context) {
	items, err := h.service.ListHelpDocuments()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to list help documents"})
		return
	}
	c.JSON(http.StatusOK, items)
}

func (h *ContentManagementHandler) UpdateHelpDocument(c *gin.Context) {
	id, _ := strconv.ParseUint(c.Param("id"), 10, 32)
	var input models.HelpDocument
	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	item, err := h.service.UpdateHelpDocument(uint(id), input)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to update help document"})
		return
	}
	c.JSON(http.StatusOK, item)
}

func (h *ContentManagementHandler) DeleteHelpDocument(c *gin.Context) {
	id, _ := strconv.ParseUint(c.Param("id"), 10, 32)
	if err := h.service.DeleteHelpDocument(uint(id)); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to delete help document"})
		return
	}
	c.JSON(http.StatusOK, gin.H{"message": "Help document deleted successfully"})
}
