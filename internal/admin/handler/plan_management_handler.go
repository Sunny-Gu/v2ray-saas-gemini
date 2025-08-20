package handler

import (
	"net/http"
	"strconv"
	"v2ray-saas-gemini/internal/admin/service"
	"v2ray-saas-gemini/internal/models"

	"github.com/gin-gonic/gin"
)

// PlanManagementHandler handles plan management API requests.
type PlanManagementHandler struct {
	service service.PlanManagementService
}

func NewPlanManagementHandler() *PlanManagementHandler {
	return &PlanManagementHandler{service: service.PlanManagementService{}}
}

func (h *PlanManagementHandler) CreatePlan(c *gin.Context) {
	var input models.Plan
	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	plan, err := h.service.CreatePlan(input)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to create plan"})
		return
	}
	c.JSON(http.StatusCreated, plan)
}

func (h *PlanManagementHandler) ListPlans(c *gin.Context) {
	plans, err := h.service.ListPlans()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to list plans"})
		return
	}
	c.JSON(http.StatusOK, plans)
}

func (h *PlanManagementHandler) GetPlan(c *gin.Context) {
	id, _ := strconv.ParseUint(c.Param("id"), 10, 32)
	plan, err := h.service.GetPlan(uint(id))
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Plan not found"})
		return
	}
	c.JSON(http.StatusOK, plan)
}

func (h *PlanManagementHandler) UpdatePlan(c *gin.Context) {
	id, _ := strconv.ParseUint(c.Param("id"), 10, 32)
	var input models.Plan
	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	plan, err := h.service.UpdatePlan(uint(id), input)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to update plan"})
		return
	}
	c.JSON(http.StatusOK, plan)
}

func (h *PlanManagementHandler) DeletePlan(c *gin.Context) {
	id, _ := strconv.ParseUint(c.Param("id"), 10, 32)
	if err := h.service.DeletePlan(uint(id)); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to delete plan"})
		return
	}
	c.JSON(http.StatusOK, gin.H{"message": "Plan deleted successfully"})
}
