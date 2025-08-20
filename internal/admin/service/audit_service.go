package service

import "v2ray-saas-gemini/internal/database"

// AuditService provides services for financial auditing.
type AuditService struct{}

// FinancialOverview represents a snapshot of the platform's finances.
type FinancialOverview struct {
	TotalUserBalance      float64 `json:"total_user_balance"`
	TotalRechargeAmount   float64 `json:"total_recharge_amount"`
	TotalConsumedAmount   float64 `json:"total_consumed_amount"`
	TotalUsers            int64   `json:"total_users"`
	TotalActiveSubs       int64   `json:"total_active_subs"`
}

// GetFinancialOverview calculates and returns a financial overview.
func (s *AuditService) GetFinancialOverview() (*FinancialOverview, error) {
	var overview FinancialOverview
	var err error

	// Sum of all users' current balances
	err = database.DB.Table("balances").Select("sum(current_balance)").Row().Scan(&overview.TotalUserBalance)
	if err != nil { return nil, err }

	// Sum of all successful recharge orders
	err = database.DB.Table("recharge_orders").Where("status = ?", "completed").Select("sum(amount_cny)").Row().Scan(&overview.TotalRechargeAmount)
	if err != nil { return nil, err }

	// Sum of all consumed amounts from balances
	err = database.DB.Table("balances").Select("sum(total_consumed)").Row().Scan(&overview.TotalConsumedAmount)
	if err != nil { return nil, err }

	// Total number of users
	err = database.DB.Table("users").Count(&overview.TotalUsers).Error
	if err != nil { return nil, err }

	// Total number of active subscriptions
	err = database.DB.Table("subscriptions").Where("status = ?", "active").Count(&overview.TotalActiveSubs).Error
	if err != nil { return nil, err }

	return &overview, nil
}
