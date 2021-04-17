package models

import "time"

type Expenses struct {
	ID           int        `json:"id"`
	CreatedAt    *time.Time `json:"createdAt"`
	UpdatedAt    *time.Time `json:"updatedAt"`
	Description  string     `json:"description"`
	Amount       int64      `json:"amount"`
	CategoryName string     `json:"category,omitempty"`
	UserID       int        `json:"userID"`
}

type PaginatedExpenses struct {
	Expenses []Expenses `json:"expenses"`
	HasMore  bool       `json:"hasMore"`
}

type Category struct {
	ID       int        `json:"id" gorm:"primaryKey"`
	Name     string     `json:"name" gorm:"unique"`
	Expenses []Expenses `json:"-" gorm:"foreignKey:CategoryName;references:Name"`
}
