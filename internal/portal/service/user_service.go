package service

import (
	"errors"
	"v2ray-saas-gemini/internal/database"
	"v2ray-saas-gemini/internal/models"
	"v2ray-saas-gemini/internal/utils"

	"gorm.io/gorm"
)

// UserService provides user-related services.
type UserService struct{}

// RegisterUserInput defines the input for the user registration.
type RegisterUserInput struct {
	Email    string `json:"email" binding:"required,email"`
	Password string `json:"password" binding:"required,min=8"`
}

// Register handles the user registration process.
func (s *UserService) Register(input RegisterUserInput) (*models.User, error) {
	// Check if user already exists
	var existingUser models.User
	if err := database.DB.Where("email = ?", input.Email).First(&existingUser).Error; err == nil {
		return nil, errors.New("user with this email already exists")
	} else if !errors.Is(err, gorm.ErrRecordNotFound) {
		// Handle other potential database errors
		return nil, err
	}

	// Hash the password
	hashedPassword, err := utils.HashPassword(input.Password)
	if err != nil {
		return nil, errors.New("failed to hash password")
	}

	// Create the new user
	newUser := models.User{
		Email:        input.Email,
		PasswordHash: hashedPassword,
		Status:       models.StatusUnverified, // Or StatusActive if no email verification
	}

	// Save the user to the database
	if err := database.DB.Create(&newUser).Error; err != nil {
		return nil, errors.New("failed to create user")
	}

	return &newUser, nil
}
