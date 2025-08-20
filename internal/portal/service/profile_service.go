package service

import (
	"errors"
	"v2ray-saas-gemini/internal/database"
	"v2ray-saas-gemini/internal/models"

	"gorm.io/gorm"
)

// ProfileService provides profile-related services.
type ProfileService struct{}

// GetProfileDataOutput defines the structure for the profile data response.
type GetProfileDataOutput struct {
	Email          string  `json:"email"`
	Balance        float64 `json:"balance"`
	PlanName       string  `json:"plan_name"`
	PlanExpiryDate string  `json:"plan_expiry_date"`
	UsedTrafficGB  float64 `json:"used_traffic_gb"`
	TotalTrafficGB float64 `json:"total_traffic_gb"`
}

// GetProfileData retrieves comprehensive profile information for a given user ID.
func (s *ProfileService) GetProfileData(userID uint) (*GetProfileDataOutput, error) {
	var user models.User
	if err := database.DB.First(&user, userID).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, errors.New("user not found")
		}
		return nil, err
	}

	var balance models.Balance
	database.DB.Where("user_id = ?", userID).First(&balance)

	var subscription models.Subscription
	database.DB.Where("user_id = ? AND status = ?", userID, models.SubStatusActive).First(&subscription)

	var plan models.Plan
	if subscription.ID != 0 {
		database.DB.First(&plan, subscription.PlanID)
	}

	output := &GetProfileDataOutput{
		Email:          user.Email,
		Balance:        balance.CurrentBalance,
		PlanName:       plan.Name,
		PlanExpiryDate: subscription.ExpiredAt.Format("2006-01-02"),
		UsedTrafficGB:  subscription.UsedTrafficGB,
		TotalTrafficGB: subscription.TotalTrafficGB,
	}

	return output, nil
}

// GetConsumptionHistory retrieves a list of subscription orders for a user.
func (s *ProfileService) GetConsumptionHistory(userID uint) ([]models.SubscriptionOrder, error) {
	var orders []models.SubscriptionOrder
	err := database.DB.Where("user_id = ?", userID).Order("created_at desc").Find(&orders).Error
	if err != nil {
		return nil, errors.New("failed to retrieve consumption history")
	}
	return orders, nil
}
