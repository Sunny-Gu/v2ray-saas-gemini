package models

import "gorm.io/gorm"

// RechargePreset represents a pre-configured recharge option available to users.
type RechargePreset struct {
	gorm.Model
	AmountCNY   float64 `gorm:"type:decimal(10,2);not null;comment:面额(CNY)"`
	Description string  `gorm:"type:varchar(255);comment:描述"`
	IsEnabled   bool    `gorm:"default:true;comment:是否启用"`
	SortOrder   int     `gorm:"default:0;comment:排序"`
}
