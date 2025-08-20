package service

import (
	"testing"
	"time"
	"v2ray-saas-gemini/internal/database"
	"v2ray-saas-gemini/internal/models"

	"github.com/stretchr/testify/assert"
)

func TestUpdateUserTraffic(t *testing.T) {
	defer func() {
		database.DB.Exec("DELETE FROM subscriptions")
		database.DB.Exec("DELETE FROM users")
	}()

	service := TrafficService{}

	// Setup: Create users with active subscriptions
	user1 := createUserWithSubscription("user1@traffic.com", models.SubStatusActive, 10, 100, time.Now().Add(24*time.Hour))
	user2 := createUserWithSubscription("user2@traffic.com", models.SubStatusActive, 20, 100, time.Now().Add(24*time.Hour))

	// --- Prepare Traffic Report ---
	report := TrafficReportInput{
		{UserID: user1.ID, UploadGB: 5.5, DownloadGB: 4.5}, // Total 10 GB
		{UserID: user2.ID, UploadGB: 1, DownloadGB: 1},     // Total 2 GB
		{UserID: 999, UploadGB: 1, DownloadGB: 1},        // Non-existent user, should be ignored
	}

	// --- Run the Service Method ---
	err := service.UpdateUserTraffic(report)
	assert.NoError(t, err)

	// --- Assertions ---
	// Verify user1's traffic
	var sub1 models.Subscription
	database.DB.Where("user_id = ?", user1.ID).First(&sub1)
	assert.InDelta(t, 20.0, sub1.UsedTrafficGB, 0.001) // 10 (initial) + 10 (report) = 20

	// Verify user2's traffic
	var sub2 models.Subscription
	database.DB.Where("user_id = ?", user2.ID).First(&sub2)
	assert.InDelta(t, 22.0, sub2.UsedTrafficGB, 0.001) // 20 (initial) + 2 (report) = 22
}
