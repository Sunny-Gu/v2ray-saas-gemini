package models

import (
	"time"

	"gorm.io/gorm"
)

// OrderStatusType defines the status of an order
type OrderStatusType string

const (
	OrderStatusPending   OrderStatusType = "pending"
	OrderStatusPaid      OrderStatusType = "paid"
	OrderStatusCompleted OrderStatusType = "completed"
	OrderStatusCancelled OrderStatusType = "cancelled"
	OrderStatusExpired   OrderStatusType = "expired"
	OrderStatusAbnormal  OrderStatusType = "abnormal"
)

// RechargeOrder represents a user's recharge transaction
type RechargeOrder struct {
	gorm.Model
	UserID         uint            `gorm:"not null;index"`
	User           User            // Belongs to User
	AmountCNY      float64         `gorm:"type:decimal(10,2);not null"` // The final amount credited to the user's balance in CNY
	DiscountRate   float64         `gorm:"type:decimal(3,2);default:1.00"` // e.g., 0.80 for 20% off
	PaidUSDT       float64         `gorm:"type:decimal(16,6)"`           // Actual USDT amount paid by user
	ExchangeRate   float64         `gorm:"type:decimal(10,6)"`           // CNY to USDT exchange rate at the time of order
	PaymentAddress string          `gorm:"type:varchar(255)"`
	TransactionID  string          `gorm:"type:varchar(255);index"` // Blockchain transaction hash
	Status         OrderStatusType `gorm:"type:varchar(20);default:'pending'"`
	CompletedAt    time.Time
}

// SubscriptionOrder represents a user's plan purchase transaction
type SubscriptionOrder struct {
	gorm.Model
	UserID        uint            `gorm:"not null;index"`
	User          User            // Belongs to User
	PlanID        uint            `gorm:"not null"` // Foreign key to a Plan model
	Amount        float64         `gorm:"type:decimal(10,2);not null"`
	Discount      float64         `gorm:"type:decimal(10,2);default:0.00"`
	FinalAmount   float64         `gorm:"type:decimal(10,2);not null"`
	Status        OrderStatusType `gorm:"type:varchar(20);default:'pending'"`
}
