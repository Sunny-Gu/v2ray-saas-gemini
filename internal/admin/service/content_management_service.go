package service

import (
	"v2ray-saas-gemini/internal/database"
	"v2ray-saas-gemini/internal/models"
)

// ContentManagementService provides services for managing content.
type ContentManagementService struct{}

// --- Announcements ---

func (s *ContentManagementService) CreateAnnouncement(input models.Announcement) (*models.Announcement, error) {
	if err := database.DB.Create(&input).Error; err != nil {
		return nil, err
	}
	return &input, nil
}

func (s *ContentManagementService) ListAnnouncements() ([]models.Announcement, error) {
	var items []models.Announcement
	err := database.DB.Order("created_at desc").Find(&items).Error
	return items, err
}

func (s *ContentManagementService) UpdateAnnouncement(id uint, input models.Announcement) (*models.Announcement, error) {
	var item models.Announcement
	if err := database.DB.First(&item, id).Error; err != nil {
		return nil, err
	}
	item.Title = input.Title
	item.Content = input.Content
	item.IsActive = input.IsActive
	if err := database.DB.Save(&item).Error; err != nil {
		return nil, err
	}
	return &item, nil
}

func (s *ContentManagementService) DeleteAnnouncement(id uint) error {
	return database.DB.Delete(&models.Announcement{}, id).Error
}

// --- Help Documents ---

func (s *ContentManagementService) CreateHelpDocument(input models.HelpDocument) (*models.HelpDocument, error) {
	if err := database.DB.Create(&input).Error; err != nil {
		return nil, err
	}
	return &input, nil
}

func (s *ContentManagementService) ListHelpDocuments() ([]models.HelpDocument, error) {
	var items []models.HelpDocument
	err := database.DB.Order("category, created_at asc").Find(&items).Error
	return items, err
}

func (s *ContentManagementService) UpdateHelpDocument(id uint, input models.HelpDocument) (*models.HelpDocument, error) {
	var item models.HelpDocument
	if err := database.DB.First(&item, id).Error; err != nil {
		return nil, err
	}
	item.Title = input.Title
	item.Content = input.Content
	item.Category = input.Category
	item.IsActive = input.IsActive
	if err := database.DB.Save(&item).Error; err != nil {
		return nil, err
	}
	return &item, nil
}

func (s *ContentManagementService) DeleteHelpDocument(id uint) error {
	return database.DB.Delete(&models.HelpDocument{}, id).Error
}
