package service

import (
	"v2ray-saas-gemini/internal/database"
	"v2ray-saas-gemini/internal/models"
)

// ContentService provides services for retrieving public content.
type ContentService struct{}

// ListActiveAnnouncements lists all active announcements.
func (s *ContentService) ListActiveAnnouncements() ([]models.Announcement, error) {
	var announcements []models.Announcement
	err := database.DB.Where("is_active = ?", true).Order("created_at desc").Find(&announcements).Error
	return announcements, err
}

// ListActiveHelpDocuments lists all active help documents.
func (s *ContentService) ListActiveHelpDocuments() ([]models.HelpDocument, error) {
	var documents []models.HelpDocument
	err := database.DB.Where("is_active = ?", true).Order("created_at asc").Find(&documents).Error
	return documents, err
}
