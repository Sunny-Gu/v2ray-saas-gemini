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

// ListRechargePresets lists all enabled recharge presets.
func (s *RechargeService) ListRechargePresets() ([]models.RechargePreset, error) {
	var presets []models.RechargePreset
	if err := database.DB.Where("is_enabled = ?", true).Order("sort_order asc").Find(&presets).Error; err != nil {
		return nil, err
	}
	return presets, nil
}

// CreateUSDTOrderInput defines the input for creating a USDT recharge order.
type CreateUSDTOrderInput struct {
	PresetID uint `json:"preset_id" binding:"required"`
}

// CreateUSDTOrder creates a new recharge order based on a selected preset.
func (s *RechargeService) CreateUSDTOrder(userID uint, input CreateUSDTOrderInput) (*models.RechargeOrder, error) {
	// 1. Get the preset details
	var preset models.RechargePreset
	if err := database.DB.First(&preset, input.PresetID).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, errors.New("recharge preset not found")
		}
		return nil, err
	}

	// 2. Get global recharge config
	rate := config.Cfg.Recharge.ExchangeRate
	discount := config.Cfg.Recharge.DiscountRate
	address := config.Cfg.Recharge.PaymentAddress

	if rate <= 0 {
		return nil, errors.New("exchange rate is not configured correctly")
	}

	// 3. Calculate the required USDT amount
	// Formula: USDT = (AmountCNY * DiscountRate) / ExchangeRate
	paidUSDT := (preset.AmountCNY * discount) / rate

	// 4. Create the recharge order
	order := models.RechargeOrder{
		UserID:         userID,
		AmountCNY:      preset.AmountCNY,
		DiscountRate:   discount,
		PaidUSDT:       paidUSDT,
		ExchangeRate:   rate,
		PaymentAddress: address,
		Status:         models.OrderStatusPending,
	}

	if err := database.DB.Create(&order).Error; err != nil {
		return nil, errors.New("failed to create recharge order")
	}

	return &order, nil
}
