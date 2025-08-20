package utils

import (
	"time"
	"v2ray-saas-gemini/internal/config"

	"github.com/dgrijalva/jwt-go"
)

// Claims defines the JWT claims structure.
type Claims struct {
	UserID uint   `json:"user_id"`
	Email  string `json:"email"`
	jwt.StandardClaims
}

// GenerateJWT generates a new JWT for a given user.
func GenerateJWT(userID uint, email string) (string, error) {
	// Get the expiration time from the config.
	expirationTime := time.Now().Add(time.Duration(config.Cfg.JWT.ExpireHours) * time.Hour)

	// Create the JWT claims.
	claims := &Claims{
		UserID: userID,
		Email:  email,
		StandardClaims: jwt.StandardClaims{
			ExpiresAt: expirationTime.Unix(),
			Issuer:    "v2ray-saas-gemini",
		},
	}

	// Create the token with the specified algorithm and claims.
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)

	// Sign the token with the secret key.
	tokenString, err := token.SignedString([]byte(config.Cfg.JWT.Secret))
	if err != nil {
		return "", err
	}

	return tokenString, nil
}

// ValidateJWT validates a given JWT string.
func ValidateJWT(tokenString string) (*Claims, error) {
	claims := &Claims{}
	token, err := jwt.ParseWithClaims(tokenString, claims, func(token *jwt.Token) (interface{}, error) {
		return []byte(config.Cfg.JWT.Secret), nil
	})

	if err != nil {
		return nil, err
	}

	if !token.Valid {
		return nil, err
	}

	return claims, nil
}
