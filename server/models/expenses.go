package models

import "time"

// Expense model
type Expenses struct {
	ID           int        `json:"id"`
	CreatedAt    *time.Time `json:"createdAt"`
	UpdatedAt    *time.Time `json:"updatedAt"`
	Description  string     `json:"description"`
	Amount       int64      `json:"amount"`
	CategoryName string     `json:"category,omitempty"`
	UserID       int        `json:"user_id"`
}

// expense response for pagination
type PaginatedExpenses struct {
	Expenses []Expenses `json:"expenses"`
	HasMore  bool       `json:"hasMore"`
}

// category model
type Category struct {
	ID       int        `json:"id" gorm:"primaryKey"`
	Name     string     `json:"name" gorm:"unique"`
	Expenses []Expenses `json:"-" gorm:"foreignKey:CategoryName;references:Name"`
}
