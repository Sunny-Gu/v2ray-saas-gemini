package models

import (
	"time"

	"gorm.io/gorm"
)

// PasswordResetToken stores tokens for the password reset process.
type PasswordResetToken struct {
	gorm.Model
	UserID    uint      `gorm:"not null"`
	Token     string    `gorm:"type:varchar(255);uniqueIndex;not null"`
	ExpiresAt time.Time `gorm:"not null"`
}
