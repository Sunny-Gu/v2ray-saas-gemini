package service

import (
	"errors"
	"v2ray-saas-gemini/internal/database"
	"v2ray-saas-gemini/internal/models"

	"gorm.io/gorm"
)

// RechargeService provides recharge-related services.
type RechargeService struct{}

// RedeemCouponInput defines the input for redeeming a coupon.
type RedeemCouponInput struct {
	Code string `json:"code" binding:"required"`
}

// RedeemCoupon handles the logic for a user to redeem a recharge code.
func (s *RechargeService) RedeemCoupon(userID uint, input RedeemCouponInput) error {
	// Start a database transaction
	tx := database.DB.Begin()
	if tx.Error != nil {
		return errors.New("failed to start transaction")
	}

	// Find the coupon by code, and lock it for update
	var coupon models.Coupon
	if err := tx.Set("gorm:query_option", "FOR UPDATE").Where("code = ?", input.Code).First(&coupon).Error; err != nil {
		tx.Rollback()
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return errors.New("coupon code not found")
		}
		return err
	}

	// Check coupon status
	if coupon.Status != models.CouponStatusActive {
		tx.Rollback()
		return errors.New("coupon is not active or has already been used")
	}

	// Update coupon status to 'used'
	coupon.Status = models.CouponStatusUsed
	coupon.UsedBy = userID
	if err := tx.Save(&coupon).Error; err != nil {
		tx.Rollback()
		return errors.New("failed to update coupon status")
	}

	// Find user's balance and lock it for update
	var balance models.Balance
	if err := tx.Set("gorm:query_option", "FOR UPDATE").Where("user_id = ?", userID).First(&balance).Error; err != nil {
		tx.Rollback()
		// This case should ideally not happen for a logged-in user
		return errors.New("user balance not found")
	}

	// Update user's balance
	balance.CurrentBalance += coupon.Value
	balance.TotalRecharge += coupon.Value
	if err := tx.Save(&balance).Error; err != nil {
		tx.Rollback()
		return errors.New("failed to update user balance")
	}

	// Commit the transaction
	if err := tx.Commit().Error; err != nil {
		return errors.New("failed to commit transaction")
	}

	return nil
}
