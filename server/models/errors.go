package models

// ErrorMessage model for users signup and login
type ErrorMessage struct {
	Field   string `json:"field"`
	Message string `json:"message"`
}
