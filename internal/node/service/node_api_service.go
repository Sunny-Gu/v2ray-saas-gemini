package service

import (
	"v2ray-saas-gemini/internal/database"
	"v2ray-saas-gemini/internal/models"
)

// NodeApiService provides services for the node-facing API.
type NodeApiService struct{}

// ActiveUser represents the data structure needed by v2ray nodes.
type ActiveUser struct {
	UserID          uint    `json:"user_id"`
	SubscriptionUUID string  `json:"subscription_uuid"` // The unique part of the subscription URL
	TotalTrafficGB  float64 `json:"total_traffic_gb"`
	UsedTrafficGB   float64 `json:"used_traffic_gb"`
}

// GetActiveUsers retrieves all users with an active subscription.
// This is the core function for providing user configs to v2ray nodes.
func (s *NodeApiService) GetActiveUsers() ([]ActiveUser, error) {
	var activeSubscriptions []models.Subscription
	if err := database.DB.
		Where("status = ? AND expired_at > NOW()", models.SubStatusActive).
		Find(&activeSubscriptions).Error; err != nil {
		return nil, err
	}

	var activeUsers []ActiveUser
	for _, sub := range activeSubscriptions {
		// Check if traffic limit is exceeded
		if sub.UsedTrafficGB >= sub.TotalTrafficGB {
			continue // Skip users who have used up their traffic
		}

		activeUsers = append(activeUsers, ActiveUser{
			UserID:          sub.UserID,
			SubscriptionUUID: sub.SubscriptionURL, // Assuming this holds the UUID
			TotalTrafficGB:  sub.TotalTrafficGB,
			UsedTrafficGB:   sub.UsedTrafficGB,
		})
	}

	return activeUsers, nil
}
