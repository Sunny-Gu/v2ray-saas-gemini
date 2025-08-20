package service

import (
	"testing"
	"v2ray-saas-gemini/internal/database"
	"v2ray-saas-gemini/internal/models"

	"github.com/stretchr/testify/assert"
)

// Helper to create a user for admin tests
func createTestUser(email string, status models.UserStatusType) *models.User {
	user := models.User{Email: email, PasswordHash: "test", Status: status}
	database.DB.Create(&user)
	return &user
}

func TestListUsers(t *testing.T) {
	defer database.DB.Exec("DELETE FROM users")

	service := UserManagementService{}
	createTestUser("test1@example.com", models.StatusActive)
	createTestUser("test2@example.com", models.StatusLocked)
	createTestUser("another@example.com", models.StatusActive)

	// --- Test Case 1: List all users (paginated) ---
	listInput1 := ListUsersInput{Page: 1, PageSize: 2}
	result1, err := service.ListUsers(listInput1)
	assert.NoError(t, err)
	assert.Equal(t, int64(3), result1.Total)
	assert.Len(t, result1.Users, 2)

	// --- Test Case 2: Search by email ---
	listInput2 := ListUsersInput{Email: "test"}
	result2, err := service.ListUsers(listInput2)
	assert.NoError(t, err)
	assert.Equal(t, int64(2), result2.Total)
	assert.Len(t, result2.Users, 2)
}

func TestUpdateUserStatus(t *testing.T) {
	defer database.DB.Exec("DELETE FROM users")

	service := UserManagementService{}
	user := createTestUser("status-update@example.com", models.StatusActive)

	// --- Test Case 1: Update status to locked ---
	updateInput := UpdateUserStatusInput{Status: models.StatusLocked}
	updatedUser, err := service.UpdateUserStatus(user.ID, updateInput)
	assert.NoError(t, err)
	assert.Equal(t, models.StatusLocked, updatedUser.Status)

	// Verify in DB
	var dbUser models.User
	database.DB.First(&dbUser, user.ID)
	assert.Equal(t, models.StatusLocked, dbUser.Status)
}
