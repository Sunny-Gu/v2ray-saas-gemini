package service

import (
	"testing"
	"time"
	"v2ray-saas-gemini/internal/database"
	"v2ray-saas-gemini/internal/models"

	"github.com/stretchr/testify/assert"
)

// Helper function to create a user and their balance for testing
func createUserWithBalance(email string, balance float64) *models.User {
	user := models.User{Email: email, PasswordHash: "test"}
	database.DB.Create(&user)
	bal := models.Balance{UserID: user.ID, CurrentBalance: balance}
	database.DB.Create(&bal)
	return &user
}

// Helper function to create a plan for testing
func createPlan(name string, price float64, traffic float64, duration int) *models.Plan {
	plan := models.Plan{
		Name:         name,
		Price:        price,
		TrafficGB:    traffic,
		DurationDays: duration,
		IsActive:     true,
	}
	database.DB.Create(&plan)
	return &plan
}

func TestPurchasePlan(t *testing.T) {
	// Clean up tables after test
	defer func() {
		database.DB.Exec("DELETE FROM subscription_orders")
		database.DB.Exec("DELETE FROM subscriptions")
		database.DB.Exec("DELETE FROM balances")
		database.DB.Exec("DELETE FROM users")
		database.DB.Exec("DELETE FROM plans")
	}()

	storeService := StoreService{}

	// Setup: Create a user and a plan
	user := createUserWithBalance("purchase@example.com", 100.0)
	plan := createPlan("Basic Plan", 50.0, 200.0, 30)

	// --- Test Case 1: Successful Purchase ---
	purchaseInput := PurchasePlanInput{PlanID: plan.ID}
	subscription, err := storeService.PurchasePlan(user.ID, purchaseInput)

	assert.NoError(t, err)
	assert.NotNil(t, subscription)
	assert.Equal(t, user.ID, subscription.UserID)
	assert.Equal(t, plan.ID, subscription.PlanID)
	assert.Equal(t, models.SubStatusActive, subscription.Status)
	assert.WithinDuration(t, time.Now().AddDate(0, 0, 30), subscription.ExpiredAt, time.Second*5)

	// Verify user balance was deducted
	var updatedBalance models.Balance
	database.DB.Where("user_id = ?", user.ID).First(&updatedBalance)
	assert.Equal(t, 50.0, updatedBalance.CurrentBalance) // 100 - 50 = 50
	assert.Equal(t, 50.0, updatedBalance.TotalConsumed)

	// Verify order was created
	var order models.SubscriptionOrder
	database.DB.Where("user_id = ? AND plan_id = ?", user.ID, plan.ID).First(&order)
	assert.Equal(t, plan.Price, order.FinalAmount)

	// --- Test Case 2: Insufficient Balance ---
	expensivePlan := createPlan("Expensive Plan", 60.0, 1000.0, 30)
	purchaseInput2 := PurchasePlanInput{PlanID: expensivePlan.ID}
	// User now has 50, so they can't buy a plan that costs 60
	_, err = storeService.PurchasePlan(user.ID, purchaseInput2)
	assert.Error(t, err)
	assert.Equal(t, "insufficient balance", err.Error())

	// Verify balance did not change
	var finalBalance models.Balance
	database.DB.Where("user_id = ?", user.ID).First(&finalBalance)
	assert.Equal(t, 50.0, finalBalance.CurrentBalance)
}
