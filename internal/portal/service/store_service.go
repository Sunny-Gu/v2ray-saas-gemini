package service

import (
	"errors"
	"time"
	"v2ray-saas-gemini/internal/database"
	"v2ray-saas-gemini/internal/models"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

// StoreService provides store-related services.
type StoreService struct{}

// ListAvailablePlans retrieves all active and featured plans for the storefront.
func (s *StoreService) ListAvailablePlans() ([]models.Plan, error) {
	var plans []models.Plan
	// Find all plans that are marked as active
	if err := database.DB.Where("is_active = ?", true).Order("created_at asc").Find(&plans).Error; err != nil {
		return nil, err
	}
	return plans, nil
}

// PurchasePlanInput defines the input for purchasing a plan.
type PurchasePlanInput struct {
	PlanID uint `json:"plan_id" binding:"required"`
}

// PurchasePlan handles the logic for a user to purchase a new plan.
func (s *StoreService) PurchasePlan(userID uint, input PurchasePlanInput) (*models.Subscription, error) {
	// Start a database transaction
	tx := database.DB.Begin()
	if tx.Error != nil {
		return nil, errors.New("failed to start transaction")
	}

	// 1. Get the plan details
	var plan models.Plan
	if err := tx.First(&plan, input.PlanID).Error; err != nil {
		tx.Rollback()
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, errors.New("plan not found")
		}
		return nil, err
	}

	// 2. Get user's balance and lock the row for update
	var balance models.Balance
	if err := tx.Set("gorm:query_option", "FOR UPDATE").Where("user_id = ?", userID).First(&balance).Error; err != nil {
		tx.Rollback()
		return nil, errors.New("user balance not found")
	}

	// 3. Check if the balance is sufficient
	if balance.CurrentBalance < plan.Price {
		tx.Rollback()
		return nil, errors.New("insufficient balance")
	}

	// 4. Deduct the balance
	balance.CurrentBalance -= plan.Price
	balance.TotalConsumed += plan.Price
	if err := tx.Save(&balance).Error; err != nil {
		tx.Rollback()
		return nil, errors.New("failed to update balance")
	}

	// 5. Create the new subscription
	now := time.Now()
	newSubscription := models.Subscription{
		UserID:         userID,
		PlanID:         plan.ID,
		StartedAt:      &now,
		ExpiredAt:      now.AddDate(0, 0, plan.DurationDays),
		TotalTrafficGB: plan.TrafficGB,
		UsedTrafficGB:  0,
		Status:         models.SubStatusActive,
		SubscriptionURL: uuid.New().String(), // Generate a unique subscription URL
	}
	if err := tx.Create(&newSubscription).Error; err != nil {
		tx.Rollback()
		return nil, errors.New("failed to create subscription")
	}

	// 6. (Optional) Create a subscription order for records
	newOrder := models.SubscriptionOrder{
		UserID:      userID,
		PlanID:      plan.ID,
		Amount:      plan.Price,
		FinalAmount: plan.Price,
		Status:      models.OrderStatusCompleted,
	}
	if err := tx.Create(&newOrder).Error; err != nil {
		tx.Rollback()
		return nil, errors.New("failed to create subscription order")
	}

	// Commit the transaction
	if err := tx.Commit().Error; err != nil {
		return nil, errors.New("failed to commit transaction")
	}

	return &newSubscription, nil
}
