package models

import "gorm.io/gorm"

// Admin represents an administrator account.
type Admin struct {
	gorm.Model
	Username     string `gorm:"type:varchar(100);uniqueIndex;not null"`
	PasswordHash string `gorm:"type:varchar(255);not null"`
	Role         string `gorm:"type:varchar(50)"` // e.g., superadmin, finance, support
	IsActive     bool   `gorm:"default:true"`
}
