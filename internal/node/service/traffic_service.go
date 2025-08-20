package service

import (
	"v2ray-saas-gemini/internal/database"
	"v2ray-saas-gemini/internal/models"

	"gorm.io/gorm"
)

// TrafficService provides services for handling traffic reports.
type TrafficService struct{}

// TrafficReportInput defines the structure for traffic data reported by nodes.
// It's a slice to allow nodes to report usage for multiple users in one request.
type TrafficReportInput []struct {
	UserID     uint    `json:"user_id" binding:"required"`
	UploadGB   float64 `json:"upload_gb"`
	DownloadGB float64 `json:"download_gb"`
}

// UpdateUserTraffic updates the traffic usage for multiple users based on a report.
func (s *TrafficService) UpdateUserTraffic(input TrafficReportInput) error {
	// Use a transaction to ensure all updates are processed or none are.
	tx := database.DB.Begin()
	if tx.Error != nil {
		return tx.Error
	}

	for _, report := range input {
		// Find the user's active subscription.
		var subscription models.Subscription
		err := tx.Where("user_id = ? AND status = ?", report.UserID, models.SubStatusActive).
			First(&subscription).Error

		if err != nil {
			// If a user's subscription is not found, just skip them.
			// This could happen if the subscription expired between the node's sync and the report.
			if err == gorm.ErrRecordNotFound {
				continue
			}
			tx.Rollback()
			return err
		}

		// Update the used traffic.
		// Using gorm.Expr to perform the update atomically in the database.
		err = tx.Model(&subscription).Update("used_traffic_gb", gorm.Expr("used_traffic_gb + ?", report.UploadGB+report.DownloadGB)).Error
		if err != nil {
			tx.Rollback()
			return err
		}
	}

	return tx.Commit().Error
}
