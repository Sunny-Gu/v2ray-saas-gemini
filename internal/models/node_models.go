package models

import (
	"time"

	"gorm.io/gorm"
)

// NodeStatusType defines the operational status of a node
type NodeStatusType string

const (
	NodeStatusOnline      NodeStatusType = "online"
	NodeStatusOffline     NodeStatusType = "offline"
	NodeStatusMaintenance NodeStatusType = "maintenance"
)

// Node represents a proxy server node
type Node struct {
	gorm.Model
	Name         string         `gorm:"type:varchar(255);not null"`
	ServerIP     string         `gorm:"type:varchar(255);not null"`
	Port         int            `gorm:"not null"`
	ProtocolType string         `gorm:"type:varchar(50);not null"` // e.g., VMESS, VLESS, Trojan
	Encryption   string         `gorm:"type:varchar(100)"`
	Location     string         `gorm:"type:varchar(100)"`
	Weight       int            `gorm:"default:100"`
	Status       NodeStatusType `gorm:"type:varchar(20);default:'offline'"`
	LastCheckAt  time.Time
}

// NodeTraffic records the traffic usage for a user on a specific node
type NodeTraffic struct {
	gorm.Model
	NodeID       uint      `gorm:"not null;index"`
	Node         Node      // Belongs to Node
	UserID       uint      `gorm:"not null;index"`
	User         User      // Belongs to User
	UploadGB     float64   `gorm:"type:decimal(10,4);not null"`
	DownloadGB   float64   `gorm:"type:decimal(10,4);not null"`
	ReportedAt   time.Time // Time the data was reported by the node
}
