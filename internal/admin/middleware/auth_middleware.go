package middleware

import (
	"net/http"
	"strings"
	"v2ray-saas-gemini/internal/utils"

	"github.com/gin-gonic/gin"
)

// AdminAuthMiddleware creates a middleware for admin JWT authentication.
func AdminAuthMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		authHeader := c.GetHeader("Authorization")
		if authHeader == "" {
			c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"error": "Authorization header is required"})
			return
		}

		parts := strings.Split(authHeader, " ")
		if len(parts) != 2 || strings.ToLower(parts[0]) != "bearer" {
			c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"error": "Authorization header format must be Bearer {token}"})
			return
		}

		tokenString := parts[1]
		claims, err := utils.ValidateJWT(tokenString)
		if err != nil {
			c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"error": "Invalid or expired token"})
			return
		}

		// Set admin info in the context
		c.Set("adminID", claims.UserID) // Re-using UserID from claims for admin ID
		c.Set("adminUsername", claims.Email) // Re-using Email from claims for admin username

		c.Next()
	}
}
