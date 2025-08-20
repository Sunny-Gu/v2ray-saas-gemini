package models

import (
	"time"

	"gorm.io/gorm"
)

type TicketStatus string

const (
	TicketStatusOpen      TicketStatus = "open"
	TicketStatusInProgress TicketStatus = "in_progress"
	TicketStatusResolved  TicketStatus = "resolved"
	TicketStatusClosed    TicketStatus = "closed"
)

// Ticket represents a support ticket submitted by a user.
type Ticket struct {
	gorm.Model
	UserID  uint         `gorm:"not null;index"`
	Title   string       `gorm:"type:varchar(255);not null"`
	Status  TicketStatus `gorm:"type:varchar(20);default:'open'"`
	Replies []TicketReply
}

// TicketReply represents a reply to a support ticket.
type TicketReply struct {
	gorm.Model
	TicketID  uint   `gorm:"not null;index"`
	UserID    uint   `gorm:"not null"` // ID of the user who replied (can be user or admin)
	Message   string `gorm:"type:text;not null"`
	IsAdmin   bool   `gorm:"default:false"` // To distinguish admin replies
	CreatedAt time.Time
}
