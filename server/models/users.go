package models

import "time"

// User definition
type User struct {
	ID        int        `json:"id"`
	CreatedAt *time.Time `json:"createdAt"`
	UpdatedAt *time.Time `json:"updatedAt"`
	Username  string     `json:"username"`
	Email     string     `json:"email"`
	Password  string     `json:"-"`
	Expenses  []Expenses `json:"-" gorm:"foreignKey:UserID;references:ID"`
	Categorys []Category `json:"-" gorm:"foreignKey:UserID;references:ID"`
}

// User input for login
type UserInput struct {
	UsernameOrEmail string `json:"usernameOrEmail"`
	Password        string `json:"password"`
}

// User input for register
type UserRegisterInput struct {
	Username string `json:"username"`
	Email    string `json:"email"`
	Password string `json:"password"`
}
