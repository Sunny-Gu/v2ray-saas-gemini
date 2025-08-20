package service

import (
	"errors"
	"v2ray-saas-gemini/internal/database"
	"v2ray-saas-gemini/internal/models"

	"github.com/google/uuid"
)

// SubscriptionService provides subscription-related services.
type SubscriptionService struct{}

// GetSubscriptionInfo retrieves the active subscription for a user.
func (s *SubscriptionService) GetSubscriptionInfo(userID uint) (*models.Subscription, error) {
	var subscription models.Subscription
	err := database.DB.Where("user_id = ? AND status = ?", userID, models.SubStatusActive).First(&subscription).Error
	if err != nil {
		return nil, errors.New("no active subscription found")
	}
	return &subscription, nil
}

// ResetSubscriptionLink generates a new subscription URL for the user's active subscription.
func (s *SubscriptionService) ResetSubscriptionLink(userID uint) (*models.Subscription, error) {
	subscription, err := s.GetSubscriptionInfo(userID)
	if err != nil {
		return nil, err
	}

	subscription.SubscriptionURL = uuid.New().String()
	if err := database.DB.Save(subscription).Error; err != nil {
		return nil, errors.New("failed to reset subscription link")
	}

	return subscription, nil
}
