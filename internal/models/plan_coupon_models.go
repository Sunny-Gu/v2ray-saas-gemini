package models

import "gorm.io/gorm"

// Plan represents a purchasable service plan in the store
type Plan struct {
	gorm.Model
	Name         string  `gorm:"type:varchar(255);not null"`
	Description  string  `gorm:"type:text"`
	Price        float64 `gorm:"type:decimal(10,2);not null"`
	TrafficGB    float64 `gorm:"type:decimal(10,2);not null"` // Total traffic in GB
	DurationDays int     `gorm:"not null"`                  // Duration in days
	IsActive     bool    `gorm:"default:true"`              // Whether the plan is available for purchase
	IsFeatured   bool    `gorm:"default:false"`             // For highlighting popular plans
}

// CouponStatusType defines the status of a coupon/recharge code
type CouponStatusType string

const (
	CouponStatusActive CouponStatusType = "active"
	CouponStatusUsed   CouponStatusType = "used"
	CouponStatusLocked CouponStatusType = "locked"
)

// Coupon represents a recharge code that can be redeemed for balance
type Coupon struct {
	gorm.Model
	Code        string           `gorm:"type:varchar(255);uniqueIndex;not null"`
	Value       float64          `gorm:"type:decimal(10,2);not null"` // The CNY value of the coupon
	Status      CouponStatusType `gorm:"type:varchar(20);default:'active'"`
	GeneratedBy uint             `gorm:"not null"` // 0 for admin, or user ID
	UsedBy      uint             // The user ID who redeemed the coupon
}
