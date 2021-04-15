package models

import "time"

type User struct {
	ID        int        `json:"id"`
	CreatedAt *time.Time `json:"createdAt"`
	UpdatedAt *time.Time `json:"updatedAt"`
	Username  string     `json:"username"`
	Email     string     `json:"email"`
	Password  string
}

type UserInput struct {
	UsernameOrEmail string `json:"usernameOrEmail"`
	Password        string `json:"password"`
}
