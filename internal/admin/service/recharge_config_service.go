package service

import (
	"v2ray-saas-gemini/internal/database"
	"v2ray-saas-gemini/internal/models"
)

// RechargeConfigService provides services for managing recharge presets.
type RechargeConfigService struct{}

// CreatePreset creates a new recharge preset.
func (s *RechargeConfigService) CreatePreset(input models.RechargePreset) (*models.RechargePreset, error) {
	if err := database.DB.Create(&input).Error; err != nil {
		return nil, err
	}
	return &input, nil
}

// ListPresets lists all recharge presets.
func (s *RechargeConfigService) ListPresets() ([]models.RechargePreset, error) {
	var presets []models.RechargePreset
	if err := database.DB.Order("sort_order asc").Find(&presets).Error; err != nil {
		return nil, err
	}
	return presets, nil
}

// UpdatePreset updates an existing recharge preset.
func (s *RechargeConfigService) UpdatePreset(id uint, input models.RechargePreset) (*models.RechargePreset, error) {
	var preset models.RechargePreset
	if err := database.DB.First(&preset, id).Error; err != nil {
		return nil, err
	}
	preset.AmountCNY = input.AmountCNY
	preset.Description = input.Description
	preset.IsEnabled = input.IsEnabled
	preset.SortOrder = input.SortOrder
	if err := database.DB.Save(&preset).Error; err != nil {
		return nil, err
	}
	return &preset, nil
}

// DeletePreset deletes a recharge preset.
func (s *RechargeConfigService) DeletePreset(id uint) error {
	if err := database.DB.Delete(&models.RechargePreset{}, id).Error; err != nil {
		return err
	}
	return nil
}
