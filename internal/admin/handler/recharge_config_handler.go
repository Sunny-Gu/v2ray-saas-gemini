package handler

import (
	"net/http"
	"strconv"
	"v2ray-saas-gemini/internal/admin/service"
	"v2ray-saas-gemini/internal/models"

	"github.com/gin-gonic/gin"
)

// RechargeConfigHandler handles recharge config related API requests.
type RechargeConfigHandler struct {
	service service.RechargeConfigService
}

func NewRechargeConfigHandler() *RechargeConfigHandler {
	return &RechargeConfigHandler{service: service.RechargeConfigService{}}
}

func (h *RechargeConfigHandler) CreatePreset(c *gin.Context) {
	var input models.RechargePreset
	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	preset, err := h.service.CreatePreset(input)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to create preset"})
		return
	}
	c.JSON(http.StatusCreated, preset)
}

func (h *RechargeConfigHandler) ListPresets(c *gin.Context) {
	presets, err := h.service.ListPresets()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to list presets"})
		return
	}
	c.JSON(http.StatusOK, presets)
}

func (h *RechargeConfigHandler) UpdatePreset(c *gin.Context) {
	id, _ := strconv.ParseUint(c.Param("id"), 10, 32)
	var input models.RechargePreset
	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	preset, err := h.service.UpdatePreset(uint(id), input)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to update preset"})
		return
	}
	c.JSON(http.StatusOK, preset)
}

func (h *RechargeConfigHandler) DeletePreset(c *gin.Context) {
	id, _ := strconv.ParseUint(c.Param("id"), 10, 32)
	if err := h.service.DeletePreset(uint(id)); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to delete preset"})
		return
	}
	c.JSON(http.StatusOK, gin.H{"message": "Preset deleted successfully"})
}
