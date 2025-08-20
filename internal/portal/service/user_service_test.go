package service

import (
	"testing"
	"v2ray-saas-gemini/internal/database"

	"github.com/stretchr/testify/assert"
)

// cleanUserTable deletes all users from the database.
func cleanUserTable() {
	database.DB.Exec("DELETE FROM users")
}

func TestRegister(t *testing.T) {
	defer cleanUserTable() // Clean up after the test

	userService := UserService{}
	input := RegisterUserInput{
		Email:    "test@example.com",
		Password: "password123",
	}

	// Test successful registration
	user, err := userService.Register(input)
	assert.NoError(t, err)
	assert.NotNil(t, user)
	assert.Equal(t, input.Email, user.Email)

	// Test registration with duplicate email
	_, err = userService.Register(input)
	assert.Error(t, err)
	assert.Equal(t, "user with this email already exists", err.Error())
}

func TestLogin(t *testing.T) {
	defer cleanUserTable() // Clean up after the test

	userService := UserService{}
	registerInput := RegisterUserInput{
		Email:    "login@example.com",
		Password: "password123",
	}

	// First, register a user to test login
	_, err := userService.Register(registerInput)
	assert.NoError(t, err)

	// Test successful login
	loginInput := LoginUserInput{
		Email:    "login@example.com",
		Password: "password123",
	}
	token, err := userService.Login(loginInput)
	assert.NoError(t, err)
	assert.NotEmpty(t, token)

	// Test login with wrong password
	loginInput.Password = "wrongpassword"
	_, err = userService.Login(loginInput)
	assert.Error(t, err)
	assert.Equal(t, "invalid email or password", err.Error())

	// Test login with non-existent email
	loginInput.Email = "nonexistent@example.com"
	_, err = userService.Login(loginInput)
	assert.Error(t, err)
	assert.Equal(t, "invalid email or password", err.Error())
}
