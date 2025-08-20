package service

import (
	"v2ray-saas-gemini/internal/database"
	"v2ray-saas-gemini/internal/models"
)

// OrderService provides services for managing orders.
type OrderService struct{}

// ListRechargeOrders retrieves a paginated list of recharge orders.
func (s *OrderService) ListRechargeOrders(page, pageSize int) ([]models.RechargeOrder, int64, error) {
	var orders []models.RechargeOrder
	var total int64

	db := database.DB.Model(&models.RechargeOrder{})
	db.Count(&total)

	offset := (page - 1) * pageSize
	err := db.Offset(offset).Limit(pageSize).Order("created_at desc").Find(&orders).Error

	return orders, total, err
}

// ListSubscriptionOrders retrieves a paginated list of subscription orders.
func (s *OrderService) ListSubscriptionOrders(page, pageSize int) ([]models.SubscriptionOrder, int64, error) {
	var orders []models.SubscriptionOrder
	var total int64

	db := database.DB.Model(&models.SubscriptionOrder{})
	db.Count(&total)

	offset := (page - 1) * pageSize
	err := db.Offset(offset).Limit(pageSize).Order("created_at desc").Find(&orders).Error

	return orders, total, err
}
