package service

import (
	"errors"
	"v2ray-saas-gemini/internal/database"
	"v2ray-saas-gemini/internal/models"
	"v2ray-saas-gemini/internal/utils"

	"gorm.io/gorm"
)

// AuthService provides authentication services for admins.
type AuthService struct{}

// LoginInput defines the input for admin login.
type LoginInput struct {
	Username string `json:"username" binding:"required"`
	Password string `json:"password" binding:"required"`
}

// Login handles the admin login process and returns a JWT.
func (s *AuthService) Login(input LoginInput) (string, error) {
	var admin models.Admin
	if err := database.DB.Where("username = ? AND is_active = ?", input.Username, true).First(&admin).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return "", errors.New("invalid username or password")
		}
		return "", err
	}

	if !utils.CheckPasswordHash(input.Password, admin.PasswordHash) {
		return "", errors.New("invalid username or password")
	}

	// Using the same JWT utility as the portal for simplicity.
	// In a real-world scenario, you might want a separate JWT setup for admins.
	token, err := utils.GenerateJWT(admin.ID, admin.Username)
	if err != nil {
		return "", errors.New("failed to generate token")
	}

	return token, nil
}
