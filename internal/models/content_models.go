package models

import "gorm.io/gorm"

// Announcement represents a system-wide announcement.
type Announcement struct {
	gorm.Model
	Title    string `gorm:"type:varchar(255);not null"`
	Content  string `gorm:"type:text;not null"`
	IsActive bool   `gorm:"default:true"`
}

// HelpDocument represents an article in the help center.
type HelpDocument struct {
	gorm.Model
	Title    string `gorm:"type:varchar(255);not null"`
	Content  string `gorm:"type:text;not null"`
	Category string `gorm:"type:varchar(100)"`
	IsActive bool   `gorm:"default:true"`
}
