package models

import "time"

type Expenses struct {
	ID           int        `json:"id"`
	CreatedAt    *time.Time `json:"createdAt"`
	UpdatedAt    *time.Time `json:"updatedAt"`
	Description  string     `json:"description"`
	Amount       int64      `json:"amount"`
	CategoryName string     `json:"category"`
	UserID       int
}

type Category struct {
	ID       int    `json:"id"`
	Name     string `json:"name"`
	UserID   int
	Expenses []Expenses `gorm:"foreignKey:CategoryName;references:Name"`
}
