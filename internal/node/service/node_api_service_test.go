package service

import (
	"testing"
	"time"
	"v2ray-saas-gemini/internal/database"
	"v2ray-saas-gemini/internal/models"

	"github.com/google/uuid"
	"github.com/stretchr/testify/assert"
)

// Helper to create a user with a subscription for node tests
func createUserWithSubscription(email string, status models.SubscriptionStatusType, usedGB, totalGB float64, expires time.Time) *models.User {
	user := models.User{Email: email, PasswordHash: "test"}
	database.DB.Create(&user)
	sub := models.Subscription{
		UserID:          user.ID,
		Status:          status,
		UsedTrafficGB:   usedGB,
		TotalTrafficGB:  totalGB,
		ExpiredAt:       expires,
		SubscriptionURL: uuid.New().String(), // Ensure unique URL for each sub
	}
	database.DB.Create(&sub)
	return &user
}

func TestGetActiveUsers(t *testing.T) {
	defer func() {
		database.DB.Exec("DELETE FROM subscriptions")
		database.DB.Exec("DELETE FROM users")
	}()

	service := NodeApiService{}

	// --- Setup Test Cases ---
	// 1. Active and valid user
	createUserWithSubscription("active@example.com", models.SubStatusActive, 50, 100, time.Now().Add(24*time.Hour))
	// 2. Expired subscription
	createUserWithSubscription("expired@example.com", models.SubStatusActive, 50, 100, time.Now().Add(-24*time.Hour))
	// 3. Inactive (pending) subscription
	createUserWithSubscription("pending@example.com", models.SubStatusPending, 50, 100, time.Now().Add(24*time.Hour))
	// 4. Traffic used up
	createUserWithSubscription("traffic-out@example.com", models.SubStatusActive, 100, 100, time.Now().Add(24*time.Hour))
	// 5. Another active user
	createUserWithSubscription("active2@example.com", models.SubStatusActive, 10, 100, time.Now().Add(24*time.Hour))

	// --- Run the Service Method ---
	activeUsers, err := service.GetActiveUsers()
	assert.NoError(t, err)

	// --- Assertions ---
	// Should only find the two active users
	assert.Len(t, activeUsers, 2)

	// Verify the correct users were returned
	emails := make(map[string]bool)
	for _, u := range activeUsers {
		var user models.User
		database.DB.First(&user, u.UserID)
		emails[user.Email] = true
	}
	assert.True(t, emails["active@example.com"])
	assert.True(t, emails["active2@example.com"])
	assert.False(t, emails["expired@example.com"])
	assert.False(t, emails["pending@example.com"])
	assert.False(t, emails["traffic-out@example.com"])
}
