package service

import (
	"v2ray-saas-gemini/internal/database"
	"v2ray-saas-gemini/internal/models"
)

// UserManagementService provides services for managing users.
type UserManagementService struct{}

// ListUsersInput defines the query parameters for listing users.
type ListUsersInput struct {
	Email    string `form:"email"`
	Page     int    `form:"page"`
	PageSize int    `form:"page_size"`
}

// ListUsersResult defines the structure for the user list response.
type ListUsersResult struct {
	Users      []models.User `json:"users"`
	Total      int64         `json:"total"`
	Page       int           `json:"page"`
	PageSize   int           `json:"page_size"`
}

// ListUsers retrieves a paginated and filterable list of users.
func (s *UserManagementService) ListUsers(input ListUsersInput) (*ListUsersResult, error) {
	db := database.DB.Model(&models.User{})

	if input.Email != "" {
		db = db.Where("email LIKE ?", "%"+input.Email+"%")
	}

	var total int64
	if err := db.Count(&total).Error; err != nil {
		return nil, err
	}

	if input.Page <= 0 {
		input.Page = 1
	}
	if input.PageSize <= 0 {
		input.PageSize = 10
	}

	offset := (input.Page - 1) * input.PageSize
	var users []models.User
	if err := db.Offset(offset).Limit(input.PageSize).Find(&users).Error; err != nil {
		return nil, err
	}

	return &ListUsersResult{
		Users:      users,
		Total:      total,
		Page:       input.Page,
		PageSize:   input.PageSize,
	}, nil
}

// UpdateUserStatusInput defines the input for updating a user's status.
type UpdateUserStatusInput struct {
	Status models.UserStatusType `json:"status" binding:"required"`
}

// UpdateUserStatus updates the status of a user's account.
func (s *UserManagementService) UpdateUserStatus(userID uint, input UpdateUserStatusInput) (*models.User, error) {
	var user models.User
	if err := database.DB.First(&user, userID).Error; err != nil {
		return nil, err
	}

	user.Status = input.Status
	if err := database.DB.Save(&user).Error; err != nil {
		return nil, err
	}
	return &user, nil
}
