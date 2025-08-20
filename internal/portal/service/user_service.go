package service

import (
	"errors"
	"time"
	"v2ray-saas-gemini/internal/database"
	"v2ray-saas-gemini/internal/models"
	"v2ray-saas-gemini/internal/utils"

	"github.com/google/uuid"
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

// LoginUserInput defines the input for the user login.
type LoginUserInput struct {
	Email    string `json:"email" binding:"required,email"`
	Password string `json:"password" binding:"required"`
}

// Login handles the user login process and returns a JWT upon success.
func (s *UserService) Login(input LoginUserInput) (string, error) {
	// Find the user by email
	var user models.User
	if err := database.DB.Where("email = ?", input.Email).First(&user).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return "", errors.New("invalid email or password")
		}
		return "", err
	}

	// Check the password
	if !utils.CheckPasswordHash(input.Password, user.PasswordHash) {
		return "", errors.New("invalid email or password")
	}

	// Generate JWT
	token, err := utils.GenerateJWT(user.ID, user.Email)
	if err != nil {
		return "", errors.New("failed to generate token")
	}

	return token, nil
}

// RequestPasswordResetInput defines the input for requesting a password reset.
type RequestPasswordResetInput struct {
	Email string `json:"email" binding:"required,email"`
}

// RequestPasswordReset generates a password reset token for a user.
// In a real application, this would also trigger an email.
func (s *UserService) RequestPasswordReset(input RequestPasswordResetInput) (string, error) {
	var user models.User
	if err := database.DB.Where("email = ?", input.Email).First(&user).Error; err != nil {
		// Do not reveal if the user exists or not for security reasons.
		return "", nil
	}

	// Generate a unique token
	token := uuid.New().String()
	expiresAt := time.Now().Add(1 * time.Hour) // Token valid for 1 hour

	resetToken := models.PasswordResetToken{
		UserID:    user.ID,
		Token:     token,
		ExpiresAt: expiresAt,
	}

	if err := database.DB.Create(&resetToken).Error; err != nil {
		return "", errors.New("failed to create reset token")
	}

	// In a real app, you would send an email with the token here.
	// For now, we return the token directly for testing purposes.
	return token, nil
}

// ResetPasswordInput defines the input for resetting a password.
type ResetPasswordInput struct {
	Token       string `json:"token" binding:"required"`
	NewPassword string `json:"new_password" binding:"required,min=8"`
}

// ResetPassword validates the token and resets the user's password.
func (s *UserService) ResetPassword(input ResetPasswordInput) error {
	var resetToken models.PasswordResetToken
	if err := database.DB.Where("token = ? AND expires_at > ?", input.Token, time.Now()).First(&resetToken).Error; err != nil {
		return errors.New("invalid or expired token")
	}

	var user models.User
	if err := database.DB.First(&user, resetToken.UserID).Error; err != nil {
		return errors.New("user not found")
	}

	hashedPassword, err := utils.HashPassword(input.NewPassword)
	if err != nil {
		return errors.New("failed to hash new password")
	}

	user.PasswordHash = hashedPassword
	if err := database.DB.Save(&user).Error; err != nil {
		return errors.New("failed to update password")
	}

	// Invalidate the token after use
	database.DB.Delete(&resetToken)

	return nil
}
