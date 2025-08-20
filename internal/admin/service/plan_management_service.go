package service

import (
	"v2ray-saas-gemini/internal/database"
	"v2ray-saas-gemini/internal/models"
)

// PlanManagementService provides services for managing service plans.
type PlanManagementService struct{}

// CreatePlan creates a new service plan.
func (s *PlanManagementService) CreatePlan(input models.Plan) (*models.Plan, error) {
	if err := database.DB.Create(&input).Error; err != nil {
		return nil, err
	}
	return &input, nil
}

// ListPlans lists all service plans.
func (s *PlanManagementService) ListPlans() ([]models.Plan, error) {
	var plans []models.Plan
	if err := database.DB.Order("created_at desc").Find(&plans).Error; err != nil {
		return nil, err
	}
	return plans, nil
}

// GetPlan retrieves a single plan by its ID.
func (s *PlanManagementService) GetPlan(id uint) (*models.Plan, error) {
	var plan models.Plan
	if err := database.DB.First(&plan, id).Error; err != nil {
		return nil, err
	}
	return &plan, nil
}

// UpdatePlan updates an existing service plan.
func (s *PlanManagementService) UpdatePlan(id uint, input models.Plan) (*models.Plan, error) {
	var plan models.Plan
	if err := database.DB.First(&plan, id).Error; err != nil {
		return nil, err
	}

	plan.Name = input.Name
	plan.Description = input.Description
	plan.Price = input.Price
	plan.TrafficGB = input.TrafficGB
	plan.DurationDays = input.DurationDays
	plan.IsActive = input.IsActive
	plan.IsFeatured = input.IsFeatured

	if err := database.DB.Save(&plan).Error; err != nil {
		return nil, err
	}
	return &plan, nil
}

// DeletePlan deletes a service plan.
func (s *PlanManagementService) DeletePlan(id uint) error {
	if err := database.DB.Delete(&models.Plan{}, id).Error; err != nil {
		return err
	}
	return nil
}
