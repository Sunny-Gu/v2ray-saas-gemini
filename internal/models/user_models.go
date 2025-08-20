package models

import (
	"time"

	"gorm.io/gorm"
)

// UserStatusType defines the user account status
type UserStatusType string

const (
	StatusActive     UserStatusType = "active"
	StatusLocked     UserStatusType = "locked"
	StatusBanned     UserStatusType = "banned"
	StatusUnverified UserStatusType = "unverified"
)

// User represents the user account entity
type User struct {
	gorm.Model
	Email        string         `gorm:"type:varchar(255);uniqueIndex;not null"`
	PasswordHash string         `gorm:"type:varchar(255);not null"`
	Status       UserStatusType `gorm:"type:varchar(20);default:'unverified'"`
	LastLoginAt  time.Time
	IPWhitelist  string `gorm:"type:text"` // Store as comma-separated string or JSON
}

// Balance represents the user's financial balance
type Balance struct {
	gorm.Model
	UserID        uint    `gorm:"uniqueIndex;not null"`
	User          User    // Belongs to User
	CurrentBalance float64 `gorm:"type:decimal(10,2);not null;default:0.00"`
	TotalRecharge  float64 `gorm:"type:decimal(10,2);not null;default:0.00"`
	TotalConsumed  float64 `gorm:"type:decimal(10,2);not null;default:0.00"`
}

// SubscriptionStatusType defines the subscription status
type SubscriptionStatusType string

const (
	SubStatusActive  SubscriptionStatusType = "active"
	SubStatusExpired SubscriptionStatusType = "expired"
	SubStatusPending SubscriptionStatusType = "pending"
)

// Subscription represents the user's service subscription
type Subscription struct {
	gorm.Model
	UserID         uint                   `gorm:"not null;index"`
	User           User                   // Belongs to User
	PlanID         uint                   `gorm:"not null"` // Foreign key to a Plan model (to be created)
	StartedAt      time.Time
	ExpiredAt      time.Time
	TotalTrafficGB float64                `gorm:"type:decimal(10,2);not null"`
	UsedTrafficGB  float64                `gorm:"type:decimal(10,2);not null;default:0.00"`
	Status         SubscriptionStatusType `gorm:"type:varchar(20);default:'pending'"`
	SubscriptionURL string                `gorm:"type:varchar(255);uniqueIndex"`
}
