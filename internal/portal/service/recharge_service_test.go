package service

import (
	"testing"
	"v2ray-saas-gemini/internal/database"
	"v2ray-saas-gemini/internal/models"

	"github.com/stretchr/testify/assert"
)

// Helper function to create a coupon for testing
func createCoupon(code string, value float64, status models.CouponStatusType) *models.Coupon {
	coupon := models.Coupon{
		Code:   code,
		Value:  value,
		Status: status,
	}
	database.DB.Create(&coupon)
	return &coupon
}

func TestRedeemCoupon(t *testing.T) {
	defer func() {
		database.DB.Exec("DELETE FROM balances")
		database.DB.Exec("DELETE FROM users")
		database.DB.Exec("DELETE FROM coupons")
	}()

	rechargeService := RechargeService{}
	user := createUserWithBalance("redeem@example.com", 10.0)
	coupon := createCoupon("VALID-CODE", 50.0, models.CouponStatusActive)
	usedCoupon := createCoupon("USED-CODE", 50.0, models.CouponStatusUsed)

	// --- Test Case 1: Successful Redemption ---
	err := rechargeService.RedeemCoupon(user.ID, RedeemCouponInput{Code: coupon.Code})
	assert.NoError(t, err)

	// Verify balance
	var balance models.Balance
	database.DB.Where("user_id = ?", user.ID).First(&balance)
	assert.Equal(t, 60.0, balance.CurrentBalance) // 10 + 50 = 60

	// Verify coupon status
	var updatedCoupon models.Coupon
	database.DB.Where("code = ?", coupon.Code).First(&updatedCoupon)
	assert.Equal(t, models.CouponStatusUsed, updatedCoupon.Status)
	assert.Equal(t, user.ID, updatedCoupon.UsedBy)

	// --- Test Case 2: Redeem a non-existent coupon ---
	err = rechargeService.RedeemCoupon(user.ID, RedeemCouponInput{Code: "FAKE-CODE"})
	assert.Error(t, err)
	assert.Equal(t, "coupon code not found", err.Error())

	// --- Test Case 3: Redeem an already used coupon ---
	err = rechargeService.RedeemCoupon(user.ID, RedeemCouponInput{Code: usedCoupon.Code})
	assert.Error(t, err)
	assert.Equal(t, "coupon is not active or has already been used", err.Error())
}

func TestGenerateCoupon(t *testing.T) {
	defer func() {
		database.DB.Exec("DELETE FROM balances")
		database.DB.Exec("DELETE FROM users")
		database.DB.Exec("DELETE FROM coupons")
	}()

	rechargeService := RechargeService{}
	user := createUserWithBalance("generate@example.com", 100.0)

	// --- Test Case 1: Successful Generation ---
	input := GenerateCouponInput{Value: 30.0}
	coupon, err := rechargeService.GenerateCoupon(user.ID, input)
	assert.NoError(t, err)
	assert.NotNil(t, coupon)
	assert.Equal(t, 30.0, coupon.Value)
	assert.Equal(t, user.ID, coupon.GeneratedBy)

	// Verify balance
	var balance models.Balance
	database.DB.Where("user_id = ?", user.ID).First(&balance)
	assert.Equal(t, 70.0, balance.CurrentBalance) // 100 - 30 = 70

	// --- Test Case 2: Insufficient Balance ---
	input.Value = 80.0 // User only has 70 left
	_, err = rechargeService.GenerateCoupon(user.ID, input)
	assert.Error(t, err)
	assert.Equal(t, "insufficient balance", err.Error())
}
